package liquidfarming

import (
	"fmt"
	"time"

	"github.com/cosmos/cosmos-sdk/telemetry"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/cosmosquad-labs/squad/v2/x/liquidfarming/keeper"
	"github.com/cosmosquad-labs/squad/v2/x/liquidfarming/types"
)

func BeginBlocker(ctx sdk.Context, k keeper.Keeper) {
	defer telemetry.ModuleMeasureSince(types.ModuleName, time.Now(), telemetry.MetricKeyBeginBlocker)

	liquidFarmSet := map[uint64]types.LiquidFarm{}
	for _, liquidFarm := range k.GetAllLiquidFarms(ctx) {
		liquidFarmSet[liquidFarm.PoolId] = liquidFarm
	}

	paramsLiquidFarmSet := map[uint64]types.LiquidFarm{}
	for _, liquidFarm := range k.GetParams(ctx).LiquidFarms {
		liquidFarmSet[liquidFarm.PoolId] = liquidFarm
	}

	for poolId := range paramsLiquidFarmSet {
		delete(liquidFarmSet, poolId)
	}

	if len(liquidFarmSet) != 0 {
		// Means that LiquidFarm is removed in params
		fmt.Println("Removed")
	}
	if len(paramsLiquidFarmSet) != 0 {
		// Means that LiquidFarm is newly added in params
		fmt.Println("Added")
	}
}
