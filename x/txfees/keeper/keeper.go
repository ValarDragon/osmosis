package keeper

import (
	"fmt"

	"github.com/tendermint/tendermint/libs/log"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"

	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"

	"github.com/osmosis-labs/osmosis/v7/x/txfees/types"
)

type (
	Keeper struct {
		cdc      codec.Codec
		storeKey sdk.StoreKey

		accountKeeper    types.AccountKeeper
		bankKeeper       *bankkeeper.BaseKeeper
		epochKeeper      types.EpochKeeper

		spotPriceCalculator types.SpotPriceCalculator

		feeCollectorName string
		fooCollectorName string
	}
)

func NewKeeper(
	cdc codec.Codec,
	accountKeeper types.AccountKeeper,
	bankKeeper *bankkeeper.BaseKeeper,
	epochKeeper types.EpochKeeper,
	storeKey sdk.StoreKey,
	spotPriceCalculator types.SpotPriceCalculator,
	feeCollectorName string,
	fooCollectorName string,
) Keeper {
	return Keeper{
		cdc:                 cdc,
		accountKeeper: accountKeeper,
		bankKeeper: bankKeeper,
		epochKeeper: epochKeeper,
		storeKey:            storeKey,
		spotPriceCalculator: spotPriceCalculator,
		feeCollectorName: feeCollectorName,
		fooCollectorName: fooCollectorName,
	}
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

func (k Keeper) GetFeeTokensStore(ctx sdk.Context) sdk.KVStore {
	store := ctx.KVStore(k.storeKey)
	return prefix.NewStore(store, []byte(types.FeeTokensStorePrefix))
}
