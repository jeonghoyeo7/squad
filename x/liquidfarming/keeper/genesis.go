package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/cosmosquad-labs/squad/v2/x/liquidfarming/types"
)

// InitGenesis initializes the capability module's state from a provided genesis state.
func (k Keeper) InitGenesis(ctx sdk.Context, genState types.GenesisState) {
	if err := genState.Validate(); err != nil {
		panic(err)
	}

	// Initialize objects to prevent from having nil slice
	if genState.Params.LiquidFarms == nil {
		genState.Params.LiquidFarms = []types.LiquidFarm{}
	}
	if genState.RewardsAuctions == nil {
		genState.RewardsAuctions = []types.RewardsAuction{}
	}

	k.SetParams(ctx, genState.Params)

	for _, liquidFarm := range genState.LiquidFarms {
		k.SetLiquidFarm(ctx, liquidFarm)
	}

	for _, record := range genState.LastRewardsAuctionIdRecord {
		k.SetRewardsAuctionId(ctx, record.PoolId, record.AuctionId)
	}

	for _, auction := range genState.RewardsAuctions {
		k.SetRewardsAuction(ctx, auction)
	}

	for _, bid := range genState.Bids {
		k.SetBid(ctx, bid)
	}

	for _, record := range genState.WinningBidRecords {
		k.SetWinningBid(ctx, record.WinningBid, record.AuctionId)
	}
}

// ExportGenesis returns the module's exported genesis.
func (k Keeper) ExportGenesis(ctx sdk.Context) *types.GenesisState {
	params := k.GetParams(ctx)

	// Initialize objects to prevent from having nil slice
	rewardsAuctions := k.GetAllRewardsAuctions(ctx)

	if params.LiquidFarms == nil {
		params.LiquidFarms = []types.LiquidFarm{}
	}

	liquidFarms := k.GetAllLiquidFarms(ctx)

	poolIds := []uint64{}
	for _, liquidFarm := range params.LiquidFarms {
		poolIds = append(poolIds, liquidFarm.PoolId)
	}

	bids := []types.Bid{}
	lastRewardsAuctionIdRecords := []types.LastRewardsAuctionIdRecord{}
	winningBidRecords := []types.WinningBidRecord{}
	for _, poolId := range poolIds {
		auctionId := k.GetLastRewardsAuctionId(ctx, poolId)
		winningBid, found := k.GetWinningBid(ctx, poolId, auctionId)
		if found {
			winningBidRecords = append(winningBidRecords, types.WinningBidRecord{
				AuctionId:  auctionId,
				WinningBid: winningBid,
			})
		}
		lastRewardsAuctionIdRecords = append(lastRewardsAuctionIdRecords, types.LastRewardsAuctionIdRecord{
			PoolId:    poolId,
			AuctionId: k.GetLastRewardsAuctionId(ctx, poolId),
		})
		bids = append(bids, k.GetBidsByPoolId(ctx, poolId)...)
	}

	return &types.GenesisState{
		Params:                     params,
		LastRewardsAuctionIdRecord: lastRewardsAuctionIdRecords,
		LiquidFarms:                liquidFarms,
		RewardsAuctions:            rewardsAuctions,
		Bids:                       bids,
		WinningBidRecords:          winningBidRecords,
	}
}
