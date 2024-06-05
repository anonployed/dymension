package transfergenesis

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/dymensionxyz/dymension/v3/utils/gerr"

	"github.com/dymensionxyz/dymension/v3/utils"

	delayedacktypes "github.com/dymensionxyz/dymension/v3/x/delayedack/types"

	transfertypes "github.com/cosmos/ibc-go/v6/modules/apps/transfer/types"

	errorsmod "cosmossdk.io/errors"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	channeltypes "github.com/cosmos/ibc-go/v6/modules/core/04-channel/types"
	porttypes "github.com/cosmos/ibc-go/v6/modules/core/05-port/types"
	"github.com/cosmos/ibc-go/v6/modules/core/exported"
	dymerror "github.com/dymensionxyz/dymension/v3/x/common/errors"
	delayedackkeeper "github.com/dymensionxyz/dymension/v3/x/delayedack/keeper"
	rollappkeeper "github.com/dymensionxyz/dymension/v3/x/rollapp/keeper"
)

var _ porttypes.Middleware = &IBCMiddleware{}

type DenomMetadataKeeper interface {
	CreateDenomMetadata(ctx sdk.Context, metadata banktypes.Metadata) error
}

type TransferKeeper interface {
	SetDenomTrace(ctx sdk.Context, denomTrace transfertypes.DenomTrace)
}

type IBCMiddleware struct {
	porttypes.Middleware // next one
	delayedackKeeper     delayedackkeeper.Keeper
	rollappKeeper        rollappkeeper.Keeper
	transferKeeper       TransferKeeper
	denomKeeper          DenomMetadataKeeper
}

func NewIBCMiddleware(
	next porttypes.Middleware,
	delayedAckKeeper delayedackkeeper.Keeper,
	rollappKeeper rollappkeeper.Keeper,
	transferKeeper TransferKeeper,
	denomKeeper DenomMetadataKeeper,
) IBCMiddleware {
	return IBCMiddleware{
		Middleware:       next,
		delayedackKeeper: delayedAckKeeper,
		rollappKeeper:    rollappKeeper,
		transferKeeper:   transferKeeper,
		denomKeeper:      denomKeeper,
	}
}

// OnRecvPacket ..
// Happy path: the genesis transfer window is open. In this case, record the transfer.
/*
TODO: what was I doing yesterday afternoon?
	I need to finish implementing this - essentially
	If the transfer does not contain the memo, need to check that transfers are enabled
		If they are, forward it
		Else, return a protocol violation (TODO: or an error ack?)
	If it does contain the memo, record it, check for a violation
		If valid:
			register denom data
		    forward it with eibc skipper
		If the last one, need to disable transfers and start the clock

Probably easiest to inline the sub method,
	/
*/
func (im IBCMiddleware) OnRecvPacket(
	ctx sdk.Context,
	packet channeltypes.Packet,
	relayer sdk.AccAddress,
) exported.Acknowledgement {
	ctx, err := im.handleGenesisTransfers(ctx, packet)
	if err != nil {
		return channeltypes.NewErrorAcknowledgement(err)
	}

	return im.Middleware.OnRecvPacket(ctx, packet, relayer)
}

type memo struct {
	Denom banktypes.Metadata `json:"denom"`
	// How many transfers in total will be sent in the transfer genesis period
	TotalNumTransfers uint64 `json:"total_num_transfers"`
	// Which transfer is this? If there are 5 transfers total, they will be numbered 0,1,2,3,4.
	ThisTransferIx uint64 `json:"this_transfer_ix"`
}

type genesisTransferDenomMemo struct {
	Data memo `json:"genesis_transfer"`
}

func (im IBCMiddleware) handleGenesisTransfers(
	ctx sdk.Context,
	packet channeltypes.Packet,
) (sdk.Context, error) {
	if !im.delayedackKeeper.IsRollappsEnabled(ctx) {
		// TODO: makes sense?
		return ctx, nil
	}

	l := ctx.Logger().With(
		"middleware", "transferGenesis",
		"packet_source_port", packet.SourcePort,
		"packet_destination_port", packet.DestinationPort,
		"packet_sequence", packet.Sequence)

	chaID, raID, err := im.getChannelAndRollappID(ctx, packet)
	if err != nil {
		// TODO:
	}

	err = im.ensureRollappExists(ctx, raID)
	if err != nil {
		// TODO:
	}

	memo, err := getMemo(packet)
	if err != nil {
		// TODO:
	}

	nTransfersDone, err := im.rollappKeeper.VerifyAndRecordGenesisTransfer(ctx, raID, memo.ThisTransferIx, memo.TotalNumTransfers)
	if errorsmod.IsOf(err, dymerror.ErrProtocolViolation) {
		// TODO: emit event or freeze rollapp, or something else?
	}
	if err != nil {
		err = fmt.Errorf("verify and record genesis transfer: %w", err)
		l.Error("OnRecvPacket", "err", err)
		panic(err)
	}

	err = im.registerDenomMetadata(ctx, raID, chaID, memo.Denom)
	if err != nil {
		err = fmt.Errorf("register denom meta: %w", err)
		l.Error("OnRecvPacket", "err", err)
		panic(err)
	}

	if nTransfersDone == memo.TotalNumTransfers {
		err = im.rollappKeeper.AddRollappToGenesisTransferFinalizationQueue(ctx, raID)
		if err != nil {
			err = fmt.Errorf("register denom meta: %w", err)
			l.Error("OnRecvPacket", "err", err)
			panic(err)
		}
		// The transfer window is finished!
		// TODO: emit event
	}

	l.Info("Registered denom meta data from genesis transfer.")

	return delayedacktypes.SkipContext(ctx), nil
}

func getMemo(packet channeltypes.Packet) (memo, error) {
	var data transfertypes.FungibleTokenPacketData
	if err := transfertypes.ModuleCdc.UnmarshalJSON(packet.GetData(), &data); err != nil {
		return memo{}, errorsmod.Wrap(sdkerrors.ErrJSONUnmarshal, "fungible token packet")
	}

	rawMemo := data.GetMemo()
	var m genesisTransferDenomMemo
	err := json.Unmarshal([]byte(rawMemo), &m)
	if err != nil {
		return memo{}, errorsmod.Wrap(sdkerrors.ErrJSONUnmarshal, "rawMemo")
	}
	return m.Data, nil
}

func (im IBCMiddleware) getChannelAndRollappID(ctx sdk.Context, packet channeltypes.Packet,
) (string, string, error) {
	chaID := "channel-0"
	raID := "rollappevm_1234-1"
	return chaID, raID, nil
}

func (im IBCMiddleware) ensureRollappExists(ctx sdk.Context, raID string) error {
	ra, ok := im.rollappKeeper.GetRollapp(ctx, raID) // TODO: necessary?
	if !ok {
		panic(errors.New("must find rollapp"))
	}

	_ = ra
	// TODO:
	return nil
}

func (im IBCMiddleware) registerDenomMetadata(ctx sdk.Context, rollappID, channelID string, m banktypes.Metadata) error {
	// TODO: only do it if it hasn't been done before?

	trace := utils.GetForeignDenomTrace(channelID, m.Base)

	im.transferKeeper.SetDenomTrace(ctx, trace)

	ibcDenom := trace.IBCDenom()

	/*
		Change the base to the ibc denom, and add an alias to the original
	*/
	m.Description = fmt.Sprintf("auto-generated ibc denom for rollapp: base: %s: rollapp: %s", ibcDenom, rollappID)
	m.Base = ibcDenom
	for i, u := range m.DenomUnits {
		if u.Exponent == 0 {
			m.DenomUnits[i].Aliases = append(m.DenomUnits[i].Aliases, u.Denom)
			m.DenomUnits[i].Denom = ibcDenom
		}
	}

	if err := m.Validate(); err != nil {
		// TODO: errorsmod with nice wrapping
		return fmt.Errorf("invalid denom metadata on genesis event: %w", err)
	}

	// We go by the denom keeper instead of calling bank directly, as something might happen in-between
	err := im.denomKeeper.CreateDenomMetadata(ctx, m)
	if errorsmod.IsOf(err, gerr.ErrAlreadyExist) {
		return nil
	}
	if err != nil {
		return errorsmod.Wrap(err, "create denom metadata")
	}

	return nil
}
