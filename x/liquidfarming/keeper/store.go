package keeper

import (
	"time"

	gogotypes "github.com/gogo/protobuf/types"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/cosmosquad-labs/squad/v3/x/liquidfarming/types"
)

// GetLiquidFarm returns liquid farm object by the given pool id.
func (k Keeper) GetLiquidFarm(ctx sdk.Context, poolId uint64) (liquidFarm types.LiquidFarm, found bool) {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.GetLiquidFarmKey(poolId))
	if bz == nil {
		return
	}
	k.cdc.MustUnmarshal(bz, &liquidFarm)
	found = true
	return
}

// GetAllLiquidFarms returns all liquid farm objects stored in the store.
func (k Keeper) GetAllLiquidFarms(ctx sdk.Context) (liquidFarms []types.LiquidFarm) {
	liquidFarms = []types.LiquidFarm{}
	k.IterateLiquidFarms(ctx, func(liquidFarm types.LiquidFarm) (stop bool) {
		liquidFarms = append(liquidFarms, liquidFarm)
		return false
	})
	return liquidFarms
}

// SetLiquidFarm stores liquid farm object with the given pool id.
func (k Keeper) SetLiquidFarm(ctx sdk.Context, liquidFarm types.LiquidFarm) {
	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshal(&liquidFarm)
	store.Set(types.GetLiquidFarmKey(liquidFarm.PoolId), bz)
}

// DeleteLiquidFarm deletes the liquid farm object from the store.
func (k Keeper) DeleteLiquidFarm(ctx sdk.Context, liquidFarm types.LiquidFarm) {
	store := ctx.KVStore(k.storeKey)
	store.Delete(types.GetLiquidFarmKey(liquidFarm.PoolId))
}

// GetLastRewardsAuctionId returns the last rewards auction id.
func (k Keeper) GetLastRewardsAuctionId(ctx sdk.Context, poolId uint64) uint64 {
	var id uint64
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.GetLastRewardsAuctionIdKey(poolId))
	if bz == nil {
		id = 0 // initialize the auction id
	} else {
		val := gogotypes.UInt64Value{}
		k.cdc.MustUnmarshal(bz, &val)
		id = val.GetValue()
	}
	return id
}

// SetLastRewardsAuctionId stores the last rewards auction id.
func (k Keeper) SetLastRewardsAuctionId(ctx sdk.Context, id uint64, poolId uint64) {
	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshal(&gogotypes.UInt64Value{Value: id})
	store.Set(types.GetLastRewardsAuctionIdKey(poolId), bz)
}

// GetLastRewardsAuctionEndTime returns the last rewards auction end time.
func (k Keeper) GetLastRewardsAuctionEndTime(ctx sdk.Context) time.Time {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.RewardsAuctionTimeKey)
	if bz == nil {
		return time.Time{}
	}
	ts, err := sdk.ParseTimeBytes(bz)
	if err != nil {
		panic(err)
	}
	return ts
}

// SetLastRewardsAuctionEndTime stores the last rewards auction end time.
func (k Keeper) SetLastRewardsAuctionEndTime(ctx sdk.Context, auctionTime time.Time) {
	store := ctx.KVStore(k.storeKey)
	bz := sdk.FormatTimeBytes(auctionTime)
	store.Set(types.RewardsAuctionTimeKey, bz)
}

// GetRewardsAuction returns the reward auction object by the given pool id and auction id.
func (k Keeper) GetRewardsAuction(ctx sdk.Context, auctionId uint64, poolId uint64) (auction types.RewardsAuction, found bool) {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.GetRewardsAuctionKey(auctionId, poolId))
	if bz == nil {
		return auction, false
	}

	auction = types.MustUnmarshalRewardsAuction(k.cdc, bz)

	return auction, true
}

// GetAllRewardsAuctions returns all rewards auctions in the store.
func (k Keeper) GetAllRewardsAuctions(ctx sdk.Context) (auctions []types.RewardsAuction) {
	auctions = []types.RewardsAuction{}
	k.IterateRewardsAuctions(ctx, func(auction types.RewardsAuction) (stop bool) {
		auctions = append(auctions, auction)
		return false
	})
	return auctions
}

// SetRewardsAuction stores rewards auction.
func (k Keeper) SetRewardsAuction(ctx sdk.Context, auction types.RewardsAuction) {
	store := ctx.KVStore(k.storeKey)
	bz := types.MustMarshalRewardsAuction(k.cdc, auction)
	store.Set(types.GetRewardsAuctionKey(auction.Id, auction.PoolId), bz)
}

// GetCompoundingRewards returns the last farming rewards by the given pool id.
func (k Keeper) GetCompoundingRewards(ctx sdk.Context, poolId uint64) (rewards types.CompoundingRewards, found bool) {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.GetCompoundingRewardsKey(poolId))
	if bz == nil {
		return
	}
	k.cdc.MustUnmarshal(bz, &rewards)
	found = true
	return
}

// SetCompoundingRewards stores compounding rewards with the given pool id.
func (k Keeper) SetCompoundingRewards(ctx sdk.Context, poolId uint64, rewards types.CompoundingRewards) {
	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshal(&rewards)
	store.Set(types.GetCompoundingRewardsKey(poolId), bz)
}

// GetBid returns the bid object by the given pool id and bidder address.
func (k Keeper) GetBid(ctx sdk.Context, poolId uint64, bidder sdk.AccAddress) (bid types.Bid, found bool) {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.GetBidKey(poolId, bidder))
	if bz == nil {
		return bid, false
	}
	k.cdc.MustUnmarshal(bz, &bid)
	return bid, true
}

// SetBid stores a bid object.
func (k Keeper) SetBid(ctx sdk.Context, bid types.Bid) {
	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshal(&bid)
	store.Set(types.GetBidKey(bid.PoolId, bid.GetBidder()), bz)
}

// DeleteBid deletes the bid object.
func (k Keeper) DeleteBid(ctx sdk.Context, bid types.Bid) {
	store := ctx.KVStore(k.storeKey)
	store.Delete(types.GetBidKey(bid.PoolId, bid.GetBidder()))
}

// GetBidsByPoolId returns all bid objects by the pool id.
func (k Keeper) GetBidsByPoolId(ctx sdk.Context, poolId uint64) []types.Bid {
	bids := []types.Bid{}
	k.IterateBidsByPoolId(ctx, poolId, func(bid types.Bid) (stop bool) {
		bids = append(bids, bid)
		return false
	})
	return bids
}

// GetWinningBid returns the winning bid object by the given pool id and auction id.
func (k Keeper) GetWinningBid(ctx sdk.Context, auctionId uint64, poolId uint64) (bid types.Bid, found bool) {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.GetWinningBidKey(auctionId, poolId))
	if bz == nil {
		return bid, false
	}
	k.cdc.MustUnmarshal(bz, &bid)
	return bid, true
}

// SetWinningBid stores the winning bid with the auction id.
func (k Keeper) SetWinningBid(ctx sdk.Context, auctionId uint64, bid types.Bid) {
	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshal(&bid)
	store.Set(types.GetWinningBidKey(auctionId, bid.PoolId), bz)
}

// DeleteWinningBid deletes the winning bid from the store.
func (k Keeper) DeleteWinningBid(ctx sdk.Context, auctionId, poolId uint64) {
	store := ctx.KVStore(k.storeKey)
	store.Delete(types.GetWinningBidKey(auctionId, poolId))
}

// IterateLiquidFarms iterates through all liquid farm objects
// stored in the store and invokes callback function for each item.
// Stops the iteration when the callback function for each time.
func (k Keeper) IterateLiquidFarms(ctx sdk.Context, cb func(liquidFarm types.LiquidFarm) (stop bool)) {
	store := ctx.KVStore(k.storeKey)
	iter := sdk.KVStorePrefixIterator(store, types.LiquidFarmKeyPrefix)
	defer iter.Close()
	for ; iter.Valid(); iter.Next() {
		var liquidFarm types.LiquidFarm
		k.cdc.MustUnmarshal(iter.Value(), &liquidFarm)
		if cb(liquidFarm) {
			break
		}
	}
}

// IterateRewardsAuctions iterates over all the stored auctions and performs a callback function.
// Stops iteration when callback returns true.
func (k Keeper) IterateRewardsAuctions(ctx sdk.Context, cb func(auction types.RewardsAuction) (stop bool)) {
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.RewardsAuctionKeyPrefix)
	defer iterator.Close()
	for ; iterator.Valid(); iterator.Next() {
		auction := types.MustUnmarshalRewardsAuction(k.cdc, iterator.Value())
		if cb(auction) {
			break
		}
	}
}

// IterateBidsBy PoolId iterates through all bids by pool id stored in the store and
// invokes callback function for each item.
// Stops the iteration when the callback function returns true.
func (k Keeper) IterateBidsByPoolId(ctx sdk.Context, poolId uint64, cb func(bid types.Bid) (stop bool)) {
	store := ctx.KVStore(k.storeKey)
	iter := sdk.KVStorePrefixIterator(store, types.GetBidByPoolIdPrefix(poolId))
	defer iter.Close()
	for ; iter.Valid(); iter.Next() {
		var bid types.Bid
		k.cdc.MustUnmarshal(iter.Value(), &bid)
		if cb(bid) {
			break
		}
	}
}

type LiquidFarm struct {
	PoolId           uint64  // the pool id
	MinDepositAmount sdk.Int // the minimum deposit amount; it allows zero value
	MinBidAmount     sdk.Int // the minimum bid amount; it allows zero value
	FeeRate          sdk.Dec // the fee rate for the liquid farm
}
