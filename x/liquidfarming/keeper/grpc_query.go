package keeper

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"

	"github.com/cosmosquad-labs/squad/v3/x/liquidfarming/types"
	liquiditytypes "github.com/cosmosquad-labs/squad/v3/x/liquidity/types"
)

// Querier is used as Keeper will have duplicate methods if used directly, and gRPC names take precedence over keeper.
type Querier struct {
	Keeper
}

var _ types.QueryServer = Querier{}

// Params queries the parameters of the module.
func (k Querier) Params(c context.Context, req *types.QueryParamsRequest) (*types.QueryParamsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}
	ctx := sdk.UnwrapSDKContext(c)
	params := k.GetParams(ctx)
	return &types.QueryParamsResponse{Params: params}, nil
}

// LiquidFarms queries all LiquidFarm objects.
func (k Querier) LiquidFarms(c context.Context, req *types.QueryLiquidFarmsRequest) (*types.QueryLiquidFarmsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}

	ctx := sdk.UnwrapSDKContext(c)

	res := []types.LiquidFarmResponse{}
	for _, liquidFarm := range k.GetLiquidFarmsInStore(ctx) {
		reserveAddr := types.LiquidFarmReserveAddress(liquidFarm.PoolId)
		poolCoinDenom := liquiditytypes.PoolCoinDenom(liquidFarm.PoolId)
		lfCoinDenom := types.LiquidFarmCoinDenom(liquidFarm.PoolId)
		farm, found := k.farmKeeper.GetFarm(ctx, poolCoinDenom)
		if !found {
			farm.TotalFarmingAmount = sdk.ZeroInt()
		}

		res = append(res, types.LiquidFarmResponse{
			PoolId:                   liquidFarm.PoolId,
			LiquidFarmReserveAddress: reserveAddr.String(),
			LFCoinDenom:              lfCoinDenom,
			MinFarmAmount:            liquidFarm.MinFarmAmount,
			MinBidAmount:             liquidFarm.MinBidAmount,
			TotalFarmingAmount:       farm.TotalFarmingAmount,
		})
	}

	return &types.QueryLiquidFarmsResponse{LiquidFarms: res}, nil
}

// LiquidFarm queries the particular LiquidFarm object.
func (k Querier) LiquidFarm(c context.Context, req *types.QueryLiquidFarmRequest) (*types.QueryLiquidFarmResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}

	if req.PoolId == 0 {
		return nil, status.Error(codes.InvalidArgument, "pool id cannot be 0")
	}

	ctx := sdk.UnwrapSDKContext(c)

	liquidFarm, found := k.GetLiquidFarm(ctx, req.PoolId)
	if !found {
		return nil, status.Errorf(codes.NotFound, "liquid farm by pool id %d not found", req.PoolId)
	}

	reserveAddr := types.LiquidFarmReserveAddress(liquidFarm.PoolId)
	poolCoinDenom := liquiditytypes.PoolCoinDenom(liquidFarm.PoolId)
	lfCoinDenom := types.LiquidFarmCoinDenom(liquidFarm.PoolId)
	farm, found := k.farmKeeper.GetFarm(ctx, poolCoinDenom)
	if !found {
		farm.TotalFarmingAmount = sdk.ZeroInt()
	}

	res := types.LiquidFarmResponse{
		PoolId:                   liquidFarm.PoolId,
		LiquidFarmReserveAddress: reserveAddr.String(),
		LFCoinDenom:              lfCoinDenom,
		MinFarmAmount:            liquidFarm.MinFarmAmount,
		MinBidAmount:             liquidFarm.MinBidAmount,
		TotalFarmingAmount:       farm.TotalFarmingAmount,
	}

	return &types.QueryLiquidFarmResponse{LiquidFarm: res}, nil
}

// RewardsAuctions queries all RewardsAuction objects.
func (k Querier) RewardsAuctions(c context.Context, req *types.QueryRewardsAuctionsRequest) (*types.QueryRewardsAuctionsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}

	if req.PoolId == 0 {
		return nil, status.Error(codes.InvalidArgument, "pool id cannot be 0")
	}

	if req.Status != "" && !(req.Status == types.AuctionStatusStarted.String() ||
		req.Status == types.AuctionStatusFinished.String() ||
		req.Status == types.AuctionStatusSkipped.String()) {
		return nil, status.Errorf(codes.InvalidArgument, "invalid auction status %s", req.Status)
	}

	ctx := sdk.UnwrapSDKContext(c)
	store := ctx.KVStore(k.storeKey)
	auctionStore := prefix.NewStore(store, types.RewardsAuctionKeyPrefix)

	// Filter auctions by descending order to show an ongoing auction first
	req.Pagination = &query.PageRequest{
		Reverse: true,
	}

	var auctions []types.RewardsAuction
	pageRes, err := query.FilteredPaginate(auctionStore, req.Pagination, func(key []byte, value []byte, accumulate bool) (bool, error) {
		auction, err := types.UnmarshalRewardsAuction(k.cdc, value)
		if err != nil {
			return false, err
		}

		if auction.PoolId != req.PoolId {
			return false, err
		}

		// Return all rewards auctions by default
		if req.Status != "" && auction.Status.String() != req.Status {
			return false, err
		}

		if accumulate {
			auctions = append(auctions, auction)
		}

		return true, nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryRewardsAuctionsResponse{RewardAuctions: auctions, Pagination: pageRes}, nil
}

// RewardsAuction queries the particular RewardsAuction object.
func (k Querier) RewardsAuction(c context.Context, req *types.QueryRewardsAuctionRequest) (*types.QueryRewardsAuctionResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}

	if req.AuctionId == 0 {
		return nil, status.Error(codes.InvalidArgument, "auction id cannot be 0")
	}

	if req.PoolId == 0 {
		return nil, status.Error(codes.InvalidArgument, "pool id cannot be 0")
	}

	ctx := sdk.UnwrapSDKContext(c)

	auction, found := k.GetRewardsAuction(ctx, req.AuctionId, req.PoolId)
	if !found {
		return nil, status.Errorf(codes.NotFound, "auction by auction %d and pool id %d not found", req.AuctionId, req.PoolId)
	}

	return &types.QueryRewardsAuctionResponse{RewardAuction: auction}, nil
}

// Bids queries all Bid objects.
func (k Querier) Bids(c context.Context, req *types.QueryBidsRequest) (*types.QueryBidsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}

	if req.PoolId == 0 {
		return nil, status.Error(codes.InvalidArgument, "pool id cannot be 0")
	}

	ctx := sdk.UnwrapSDKContext(c)
	store := ctx.KVStore(k.storeKey)
	bidStore := prefix.NewStore(store, types.BidKeyPrefix)

	var bids []types.Bid
	pageRes, err := query.FilteredPaginate(bidStore, req.Pagination, func(key []byte, value []byte, accumulate bool) (bool, error) {
		bid, err := types.UnmarshalBid(k.cdc, value)
		if err != nil {
			return false, err
		}

		if bid.PoolId != req.PoolId {
			return false, nil
		}

		if accumulate {
			bids = append(bids, bid)
		}

		return true, nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryBidsResponse{Bids: bids, Pagination: pageRes}, nil
}

// Rewards queries all farming rewards accumulated for the liquid farm.
func (k Querier) Rewards(c context.Context, req *types.QueryRewardsRequest) (*types.QueryRewardsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}

	if req.PoolId == 0 {
		return nil, status.Error(codes.InvalidArgument, "pool id cannot be 0")
	}

	ctx := sdk.UnwrapSDKContext(c)

	// Currently accumulated rewards from the farm module + all withdrawn rewards in the WithdrawnRewardsReserve account
	liquidFarmReserveAddr := types.LiquidFarmReserveAddress(req.PoolId)
	poolCoinDenom := liquiditytypes.PoolCoinDenom(req.PoolId)
	withdrawnRewards := k.farmKeeper.Rewards(ctx, liquidFarmReserveAddr, poolCoinDenom)
	truncatedRewards, _ := withdrawnRewards.TruncateDecimal()
	withdrawnRewardsReserveAddr := types.WithdrawnRewardsReserveAddress(req.PoolId)
	spendableCoins := k.bankKeeper.SpendableCoins(ctx, withdrawnRewardsReserveAddr)

	return &types.QueryRewardsResponse{Rewards: truncatedRewards.Add(spendableCoins...)}, nil
}

// ExchangeRate queries exchange rate, such as mint rate and burn rate per 1 LFCoin.
func (k Querier) ExchangeRate(c context.Context, req *types.QueryExchangeRateRequest) (*types.QueryExchangeRateResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}

	if req.PoolId == 0 {
		return nil, status.Error(codes.InvalidArgument, "pool id cannot be 0")
	}

	ctx := sdk.UnwrapSDKContext(c)

	res := types.ExchangeRateResponse{
		MintRate: sdk.ZeroDec(),
		BurnRate: sdk.ZeroDec(),
	}
	lfCoinTotalSupplyAmt := k.bankKeeper.GetSupply(ctx, types.LiquidFarmCoinDenom(req.PoolId)).Amount

	if !lfCoinTotalSupplyAmt.IsZero() {
		farm, found := k.farmKeeper.GetFarm(ctx, liquiditytypes.PoolCoinDenom(req.PoolId))
		if !found {
			farm.TotalFarmingAmount = sdk.ZeroInt()
		}

		compoundingRewards, found := k.GetCompoundingRewards(ctx, req.PoolId)
		if !found {
			compoundingRewards.Amount = sdk.ZeroInt()
		}

		// MintRate = LFCoinTotalSupply / LPCoinTotalFarmingAmount
		res.MintRate = lfCoinTotalSupplyAmt.ToDec().Quo(farm.TotalFarmingAmount.ToDec())

		// BurnRate = LPCoinTotalFarmingAmount - CompoundingRewards / LFCoinTotalSupply
		lpCoinTotalFarmingAmt := farm.TotalFarmingAmount.Sub(compoundingRewards.Amount)
		res.BurnRate = lpCoinTotalFarmingAmt.ToDec().Quo(lfCoinTotalSupplyAmt.ToDec())
	}

	return &types.QueryExchangeRateResponse{ExchangeRate: res}, nil
}
