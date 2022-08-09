package liquidfarming

import (
	"sort"
	"time"

	"github.com/cosmos/cosmos-sdk/telemetry"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/cosmosquad-labs/squad/v2/x/liquidfarming/keeper"
	"github.com/cosmosquad-labs/squad/v2/x/liquidfarming/types"
)

// BeginBlocker compares all LiquidFarms stored in the store with all LiquidFarms registered in params.
// Execute an appropriate operation when either adding new LiquidFarm or removing one through governance proposal.
func BeginBlocker(ctx sdk.Context, k keeper.Keeper) {
	defer telemetry.ModuleMeasureSince(types.ModuleName, time.Now(), telemetry.MetricKeyBeginBlocker)

	liquidFarmSet := map[uint64]types.LiquidFarm{} // PoolId => LiquidFarm
	for _, liquidFarm := range k.GetAllLiquidFarms(ctx) {
		liquidFarmSet[liquidFarm.PoolId] = liquidFarm
	}

	// Compare all liquid farms stored in the store with all liquid farms registered in params
	// Store if new liquid farm is added and delete from the liquidFarmSet if it exists
	for _, liquidFarm := range k.GetParams(ctx).LiquidFarms {
		_, found := liquidFarmSet[liquidFarm.PoolId]
		if !found { // new LiquidFarm is added
			k.SetLiquidFarm(ctx, liquidFarm)
		} else {
			delete(liquidFarmSet, liquidFarm.PoolId)
		}
	}

	// Sort map keys for deterministic execution
	var pairIds []uint64
	for pairId := range liquidFarmSet {
		pairIds = append(pairIds, pairId)
	}
	sort.Slice(pairIds, func(i, j int) bool {
		return pairIds[i] < pairIds[j]
	})

	// Remove liquid farm when it is removed in params
	for _, pairId := range pairIds {
		k.RemoveLiquidFarm(ctx, liquidFarmSet[pairId])
	}
}
