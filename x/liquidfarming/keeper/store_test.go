package keeper_test

import (
	"fmt"
	"math/rand"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"

	utils "github.com/cosmosquad-labs/squad/v2/types"
	"github.com/cosmosquad-labs/squad/v2/x/liquidfarming/types"

	_ "github.com/stretchr/testify/suite"
)

func (s *KeeperTestSuite) TestQueuedFarming() {
	// TODO: not implemented yet
	// Set | Get | Delete
}

func (s *KeeperTestSuite) TestIterateQueuedFarmingsByFarmerAndDenomReverse() {
	poolId := uint64(1)
	poolCoinDenom := "pool1"
	farmerAcc := s.addr(0)

	s.createPair(farmerAcc, "denom1", "denom2", true)
	s.createPool(farmerAcc, 1, sdk.NewCoins(sdk.NewInt64Coin("denom1", 100000000), sdk.NewInt64Coin("denom2", 100000000)), true)
	s.createLiquidFarm(types.NewLiquidFarm(poolId, sdk.ZeroInt(), sdk.ZeroInt()))
	s.Require().Len(s.keeper.GetParams(s.ctx).LiquidFarms, 1)

	for seed := int64(0); seed <= 5; seed++ {
		r := rand.New(rand.NewSource(seed))

		s.farm(poolId, farmerAcc, sdk.NewInt64Coin(poolCoinDenom, r.Int63()+1), true)
		s.nextBlock()
	}

	skip := true // first item
	prevEndTime := time.Time{}
	s.keeper.IterateQueuedFarmingsByFarmerAndDenomReverse(s.ctx, farmerAcc, poolCoinDenom, func(endTime time.Time, queuedFarming types.QueuedFarming) (stop bool) {
		if skip {
			skip = false
		} else {
			s.Require().True(prevEndTime.After(endTime))
		}
		prevEndTime = endTime

		return false
	})
}

func (s *KeeperTestSuite) TestIterateMatureQueuedFarmings() {
	pair := s.createPair(s.addr(0), "denom1", "denom2", true)
	pool := s.createPool(s.addr(0), pair.Id, utils.ParseCoins("100000000denom1,100000000denom2"), true)
	s.createLiquidFarm(types.NewLiquidFarm(pool.Id, sdk.ZeroInt(), sdk.ZeroInt()))

	s.farm(pool.Id, s.addr(0), utils.ParseCoin("5000000000pool1"), true)
	// s.nextBlock()
	// s.nextBlock()
	// s.nextBlock()

	queuedFarmings := s.keeper.GetQueuedFarmingsByFarmer(s.ctx, s.addr(0))
	fmt.Println("length: ", len(queuedFarmings))

	s.advanceEpochDays()

}
