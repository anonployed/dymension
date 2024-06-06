package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/dymensionxyz/dymension/v3/x/delayedack/types"
)

// GetParams get all parameters as types.Params
func (k Keeper) GetParams(ctx sdk.Context) types.Params {
	return types.NewParams(k.EpochIdentifier(ctx), k.BridgingFee(ctx))
}

// SetParams set the params
func (k Keeper) SetParams(ctx sdk.Context, params types.Params) {
	k.paramstore.SetParamSet(ctx, &params)
}

func (k Keeper) EpochIdentifier(ctx sdk.Context) (res string) {
	k.paramstore.Get(ctx, types.KeyEpochIdentifier, &res)
	return
}

func (k Keeper) BridgingFee(ctx sdk.Context) (res sdk.Dec) {
	k.paramstore.Get(ctx, types.KeyBridgeFee, &res)
	return
}

func (k Keeper) BridgingFeeFromAmt(ctx sdk.Context, amt sdk.Int) (res sdk.Int) {
	return k.BridgingFee(ctx).MulInt(amt).TruncateInt()
}
