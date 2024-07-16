package v4

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	upgradetypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"
	"github.com/tendermint/tendermint/crypto"

	"github.com/dymensionxyz/dymension/v3/app/keepers"
	"github.com/dymensionxyz/dymension/v3/app/upgrades"
	"github.com/dymensionxyz/dymension/v3/app/upgrades/v4/types"
	delayedackkeeper "github.com/dymensionxyz/dymension/v3/x/delayedack/keeper"
	delayedacktypes "github.com/dymensionxyz/dymension/v3/x/delayedack/types"
	rollappkeeper "github.com/dymensionxyz/dymension/v3/x/rollapp/keeper"
	rollapptypes "github.com/dymensionxyz/dymension/v3/x/rollapp/types"
)

// CreateUpgradeHandler creates an SDK upgrade handler for v4
func CreateUpgradeHandler(
	mm *module.Manager,
	appCodec codec.Codec,
	configurator module.Configurator,
	_ upgrades.BaseAppParamManager,
	keepers *keepers.AppKeepers,
) upgradetypes.UpgradeHandler {
	return func(ctx sdk.Context, _ upgradetypes.Plan, fromVM module.VersionMap) (module.VersionMap, error) {
		logger := ctx.Logger().With("upgrade", UpgradeName)

		migrateDelayedAckParams(ctx, keepers.DelayedAckKeeper)
		migrateRollappParams(ctx, keepers.RollappKeeper)
		if err := migrateRollapps(ctx, keepers.GetKey(rollapptypes.ModuleName), appCodec, keepers.RollappKeeper); err != nil {
			return nil, err
		}

		// Start running the module migrations
		logger.Debug("running module migrations ...")
		return mm.RunMigrations(ctx, configurator, fromVM)
	}
}

func migrateDelayedAckParams(ctx sdk.Context, delayedAckKeeper delayedackkeeper.Keeper) {
	// overwrite params for delayedack module due to added parameters
	params := delayedacktypes.DefaultParams()
	delayedAckKeeper.SetParams(ctx, params)
}

func migrateRollappParams(ctx sdk.Context, rollappkeeper rollappkeeper.Keeper) {
	// overwrite params for rollapp module due to proto change
	params := rollapptypes.DefaultParams()
	params.DisputePeriodInBlocks = rollappkeeper.DisputePeriodInBlocks(ctx)
	rollappkeeper.SetParams(ctx, params)
}

func migrateRollapps(ctx sdk.Context, rollappStoreKey *storetypes.KVStoreKey, appCodec codec.Codec, rollappkeeper rollappkeeper.Keeper) error {
	for _, oldRollapp := range getAllOldRollapps(ctx, rollappStoreKey, appCodec) {
		newRollapp := ConvertOldRollappToNew(oldRollapp)
		if err := newRollapp.ValidateBasic(); err != nil {
			return err
		}
		rollappkeeper.SetRollapp(ctx, newRollapp)
	}
	return nil
}

func ConvertOldRollappToNew(oldRollapp types.Rollapp) rollapptypes.Rollapp {
	bech32Prefix := oldRollapp.RollappId[:5]
	return rollapptypes.Rollapp{
		RollappId: oldRollapp.RollappId,
		Creator:   oldRollapp.Creator,
		GenesisState: rollapptypes.RollappGenesisState{
			TransfersEnabled: oldRollapp.GenesisState.TransfersEnabled,
		},
		ChannelId:               oldRollapp.ChannelId,
		Frozen:                  oldRollapp.Frozen,
		RegisteredDenoms:        oldRollapp.RegisteredDenoms,
		InitialSequencerAddress: oldRollapp.Creator,                                  // whatever, just to make test pass
		Bech32Prefix:            bech32Prefix,                                        // whatever, just to make test pass
		GenesisChecksum:         string(crypto.Sha256([]byte(oldRollapp.RollappId))), // whatever, just to make test pass
		Alias:                   fmt.Sprintf("rol%s", bech32Prefix),                  // whatever, just to make test pass
		Metadata: &rollapptypes.RollappMetadata{
			Website:      "", // TODO
			Description:  "", // TODO
			LogoDataUri:  "", // TODO
			TokenLogoUri: "", // TODO
		},
	}
}

func getAllOldRollapps(ctx sdk.Context, storeKey *storetypes.KVStoreKey, appCodec codec.Codec) (list []types.Rollapp) {
	store := prefix.NewStore(ctx.KVStore(storeKey), rollapptypes.KeyPrefix(rollapptypes.RollappKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close() // nolint: errcheck

	for ; iterator.Valid(); iterator.Next() {
		var val types.Rollapp
		bz := iterator.Value()
		appCodec.MustUnmarshalJSON(bz, &val)
		list = append(list, val)
	}

	return
}
