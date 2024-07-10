package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/suite"

	"github.com/dymensionxyz/dymension/v3/app/apptesting"
	"github.com/dymensionxyz/dymension/v3/x/rollapp/keeper"
	"github.com/dymensionxyz/dymension/v3/x/rollapp/types"

	"github.com/cosmos/cosmos-sdk/baseapp"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

const (
	alice           = "dym1wg8p6j0pxpnsvhkwfu54ql62cnrumf0v634mft"
	bob             = "dym1d0wlmz987qlurs6e3kc6zd25z6wsdmnwx8tafy"
	registrationFee = "1000000000000000000adym"
	bech32Prefix    = "eth"
	bech32Prefix2   = "rax"
)

type RollappTestSuite struct {
	apptesting.KeeperTestHelper
	msgServer   types.MsgServer
	queryClient types.QueryClient
}

func (suite *RollappTestSuite) SetupTest() {
	app := apptesting.Setup(suite.T(), false)
	ctx := app.GetBaseApp().NewContext(false, tmproto.Header{})

	app.AccountKeeper.SetParams(ctx, authtypes.DefaultParams())
	app.BankKeeper.SetParams(ctx, banktypes.DefaultParams())
	regFee, _ := sdk.ParseCoinNormalized(registrationFee)
	app.RollappKeeper.SetParams(ctx, types.NewParams(2, regFee))

	aliceBal := sdk.NewCoins(regFee.AddAmount(regFee.Amount.Mul(sdk.NewInt(10))))
	apptesting.FundAccount(app, ctx, sdk.MustAccAddressFromBech32(alice), aliceBal)

	queryHelper := baseapp.NewQueryServerTestHelper(ctx, app.InterfaceRegistry())
	types.RegisterQueryServer(queryHelper, app.RollappKeeper)
	queryClient := types.NewQueryClient(queryHelper)

	suite.App = app
	suite.msgServer = keeper.NewMsgServerImpl(app.RollappKeeper)
	suite.Ctx = ctx
	suite.queryClient = queryClient
}

func TestRollappKeeperTestSuite(t *testing.T) {
	suite.Run(t, new(RollappTestSuite))
}

func createNRollapp(keeper *keeper.Keeper, ctx sdk.Context, n int) (items []types.Rollapp, rollappSummaries []types.RollappSummary) {
	items, rollappSummaries = make([]types.Rollapp, n), make([]types.RollappSummary, n)

	for i := range items {
		items[i].RollappId = strconv.Itoa(i)
		keeper.SetRollapp(ctx, items[i])

		rollappSummaries[i] = types.RollappSummary{
			RollappId: items[i].RollappId,
		}
	}

	return
}
