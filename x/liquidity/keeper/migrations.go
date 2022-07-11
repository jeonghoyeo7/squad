package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	v2 "github.com/cosmosquad-labs/squad/v2/x/liquidity/legacy/v2"
)

type Migrator struct {
	keeper Keeper
}

func NewMigrator(keeper Keeper) Migrator {
	return Migrator{keeper: keeper}
}

func (m Migrator) Migrate1to2(ctx sdk.Context) error {
	return v2.MigrateStore(ctx, m.keeper.storeKey, m.keeper.cdc)
}
