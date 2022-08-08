package keeper_test

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	utils "github.com/cosmosquad-labs/squad/v2/types"
	"github.com/cosmosquad-labs/squad/v2/x/liquidfarming"
	"github.com/cosmosquad-labs/squad/v2/x/liquidfarming/types"

	_ "github.com/stretchr/testify/suite"
)

func (s *KeeperTestSuite) TestFarm() {
	farmerAddr := s.addr(0)

	s.createPair(farmerAddr, "denom1", "denom2", true)
	s.createPool(s.addr(0), 1, utils.ParseCoins("100_000_000denom1, 100_000_000denom2"), true)
	s.createLiquidFarm(types.NewLiquidFarm(1, sdk.ZeroInt(), sdk.ZeroInt()))
	s.Require().Len(s.keeper.GetParams(s.ctx).LiquidFarms, 1)

	pool, found := s.app.LiquidityKeeper.GetPool(s.ctx, 1)
	s.Require().True(found)

	amount1, amount2, amount3 := sdk.NewInt(100000000), sdk.NewInt(200000000), sdk.NewInt(300000000)

	s.farm(pool.Id, farmerAddr, sdk.NewCoin(pool.PoolCoinDenom, amount1), true)
	s.nextBlock()

	s.farm(pool.Id, farmerAddr, sdk.NewCoin(pool.PoolCoinDenom, amount2), true)
	s.nextBlock()

	s.farm(pool.Id, farmerAddr, sdk.NewCoin(pool.PoolCoinDenom, amount3), true)
	s.nextBlock()

	// Check if the liquid farm reserve account staked in the farming module
	reserveAcc := types.LiquidFarmReserveAddress(pool.Id)
	queuedCoins := s.app.FarmingKeeper.GetAllQueuedCoinsByFarmer(s.ctx, reserveAcc)
	s.Require().Equal(amount1.Add(amount2).Add(amount3), queuedCoins.AmountOf(pool.PoolCoinDenom))

	// Check queued farmings farmed by the farmer
	queuedFarmings := s.keeper.GetQueuedFarmingsByFarmer(s.ctx, farmerAddr)
	s.Require().Len(queuedFarmings, 3)
}

func (s *KeeperTestSuite) TestFarm_MoreCases() {
	pair := s.createPair(s.addr(0), "denom1", "denom2", true)
	pool := s.createPool(s.addr(0), pair.Id, utils.ParseCoins("100_000_000denom1, 100_000_000denom2"), true)
	s.createLiquidFarm(types.NewLiquidFarm(1, sdk.ZeroInt(), sdk.ZeroInt()))

	for _, tc := range []struct {
		name        string
		msg         *types.MsgFarm
		postRun     func(ctx sdk.Context, qfs []types.QueuedFarming)
		expectedErr string
	}{
		{
			"happy case",
			types.NewMsgFarm(
				pool.Id,
				s.addr(0).String(),
				sdk.NewInt64Coin(pool.PoolCoinDenom, 1_000_000_000),
			),
			func(ctx sdk.Context, qfs []types.QueuedFarming) {
				for _, qf := range qfs {
					fmt.Println("QueuedFarming: ", qf)
				}
			},
			"",
		},
		// {
		// 	"insufficient balance",
		// 	types.NewMsgFarm(
		// 		pool.Id,
		// 		s.addr(0).String(),
		// 		sdk.NewInt64Coin(pool.PoolCoinDenom, 5_000_000_000_000_000),
		// 	),
		// 	nil,
		// 	"1000000000000 is smaller than 5000000000000000: insufficient funds",
		// },
	} {
		s.Run(tc.name, func() {
			s.Require().NoError(tc.msg.ValidateBasic())
			cacheCtx, _ := s.ctx.CacheContext()
			err := s.keeper.Farm(cacheCtx, tc.msg)
			if tc.expectedErr == "" {
				s.Require().NoError(err)
				tc.postRun(cacheCtx, s.keeper.GetQueuedFarmingsByFarmer(s.ctx, tc.msg.GetFarmer()))
			} else {
				s.Require().EqualError(err, tc.expectedErr)
			}
		})
	}
}

func (s *KeeperTestSuite) TestFarm_AfterStaked() {
	s.createPair(s.addr(0), "denom1", "denom2", true)
	s.createPool(s.addr(0), 1, utils.ParseCoins("100_000_000denom1, 100_000_000denom2"), true)
	s.createLiquidFarm(types.NewLiquidFarm(1, sdk.ZeroInt(), sdk.ZeroInt()))
	s.Require().Len(s.keeper.GetParams(s.ctx).LiquidFarms, 1)

	pool, found := s.app.LiquidityKeeper.GetPool(s.ctx, 1)
	s.Require().True(found)

	farmerAddr := s.addr(1)
	amount1, amount2 := sdk.NewInt(100_000_000), sdk.NewInt(400_000_000)

	s.farm(pool.Id, farmerAddr, sdk.NewCoin(pool.PoolCoinDenom, amount1), true)
	s.nextBlock()
	s.advanceEpochDays()

	// Staked amount must be 100
	totalStakings, found := s.app.FarmingKeeper.GetTotalStakings(s.ctx, pool.PoolCoinDenom)
	s.Require().True(found)
	s.Require().Equal(amount1, totalStakings.Amount)

	s.farm(pool.Id, farmerAddr, sdk.NewCoin(pool.PoolCoinDenom, amount2), true)
	s.advanceEpochDays()

	// Staked amount must be 500 (amount1 + amount2)
	totalStakings, found = s.app.FarmingKeeper.GetTotalStakings(s.ctx, pool.PoolCoinDenom)
	s.Require().True(found)
	s.Require().Equal(amount1.Add(amount2), totalStakings.Amount)

	// Check minted LFCoin
	lfCoinDenom := types.LiquidFarmCoinDenom(pool.Id)
	lfCoinBalance := s.getBalance(farmerAddr, lfCoinDenom)
	s.Require().Equal(sdk.NewCoin(lfCoinDenom, amount1.Add(amount2)), lfCoinBalance)

	// Queued farmings must be deleted
	queuedFarmings := s.keeper.GetQueuedFarmingsByFarmer(s.ctx, farmerAddr)
	s.Require().Len(queuedFarmings, 0)
}

func (s *KeeperTestSuite) TestFarm_AfterStaked_MultipleFarmers() {
	s.createPair(s.addr(0), "denom1", "denom2", true)
	s.createPool(s.addr(0), 1, utils.ParseCoins("100_000_000denom1, 100_000_000denom2"), true)
	s.createLiquidFarm(types.NewLiquidFarm(1, sdk.ZeroInt(), sdk.ZeroInt()))

	pool, found := s.app.LiquidityKeeper.GetPool(s.ctx, 1)
	s.Require().True(found)

	farmerAddr1, farmerAddr2, farmerAddr3 := s.addr(1), s.addr(2), s.addr(3)
	amount1, amount2, amount3 := sdk.NewInt(100000000), sdk.NewInt(200000000), sdk.NewInt(666)

	// Multiple farmers farm in the same block
	s.farm(pool.Id, farmerAddr1, sdk.NewCoin(pool.PoolCoinDenom, amount1), true)
	s.farm(pool.Id, farmerAddr2, sdk.NewCoin(pool.PoolCoinDenom, amount2), true)
	s.advanceEpochDays()

	queuedFarmings := s.keeper.GetQueuedFarmingsByFarmer(s.ctx, farmerAddr1)
	s.Require().Len(queuedFarmings, 0)

	queuedFarmings = s.keeper.GetQueuedFarmingsByFarmer(s.ctx, farmerAddr2)
	s.Require().Len(queuedFarmings, 0)

	s.Require().Equal(amount1, s.getBalance(farmerAddr1, types.LiquidFarmCoinDenom(pool.Id)).Amount)
	s.Require().Equal(amount2, s.getBalance(farmerAddr2, types.LiquidFarmCoinDenom(pool.Id)).Amount)

	s.farm(pool.Id, farmerAddr3, sdk.NewCoin(pool.PoolCoinDenom, amount3), true)
	s.advanceEpochDays()

	queuedFarmings = s.keeper.GetQueuedFarmingsByFarmer(s.ctx, farmerAddr3)
	s.Require().Len(queuedFarmings, 0)

	reserveAddr := types.LiquidFarmReserveAddress(pool.Id)
	lfCoinTotalSupply := s.app.BankKeeper.GetSupply(s.ctx, types.LiquidFarmCoinDenom(pool.Id)).Amount
	lpCoinTotalStaked := s.app.FarmingKeeper.GetAllStakedCoinsByFarmer(s.ctx, reserveAddr).AmountOf(pool.PoolCoinDenom)
	mintingAmt := lfCoinTotalSupply.Mul(amount3).Quo(lpCoinTotalStaked)
	s.Require().Equal(mintingAmt, s.getBalance(farmerAddr3, types.LiquidFarmCoinDenom(pool.Id)).Amount)
}

func (s *KeeperTestSuite) TestCancelQueuedFarming_All() {
	s.createPair(s.addr(0), "denom1", "denom2", true)
	s.createPool(s.addr(0), 1, utils.ParseCoins("100_000_000denom1, 100_000_000denom2"), true)
	s.createLiquidFarm(types.NewLiquidFarm(1, sdk.ZeroInt(), sdk.ZeroInt()))
	s.Require().Len(s.keeper.GetParams(s.ctx).LiquidFarms, 1)

	pool, found := s.app.LiquidityKeeper.GetPool(s.ctx, 1)
	s.Require().True(found)

	farmerAddr := s.addr(1)

	s.farm(pool.Id, farmerAddr, sdk.NewCoin(pool.PoolCoinDenom, sdk.NewInt(500_000_000)), true)
	s.nextBlock()

	s.farm(pool.Id, farmerAddr, sdk.NewCoin(pool.PoolCoinDenom, sdk.NewInt(500_000_000)), true)
	s.nextBlock()

	queuedFarmings := s.keeper.GetQueuedFarmingsByFarmer(s.ctx, farmerAddr)
	s.Require().Len(queuedFarmings, 2)

	// Cancel all amounts
	unfarmingCoin := sdk.NewInt64Coin(pool.PoolCoinDenom, 1_000_000_000)
	s.cancelQueuedFarming(pool.Id, farmerAddr, unfarmingCoin)

	// QueuedFarmings must be deleted
	queuedFarmings = s.keeper.GetQueuedFarmingsByFarmer(s.ctx, farmerAddr)
	s.Require().Len(queuedFarmings, 0)

	// Unfarming amount must be returned to the farmer
	balance := s.getBalance(farmerAddr, pool.PoolCoinDenom)
	s.Require().Equal(unfarmingCoin, balance)
}

func (s *KeeperTestSuite) TestCancelQueuedFarming_Partial() {
	s.createPair(s.addr(0), "denom1", "denom2", true)
	s.createPool(s.addr(0), 1, utils.ParseCoins("100_000_000denom1, 100_000_000denom2"), true)
	s.createLiquidFarm(types.NewLiquidFarm(1, sdk.ZeroInt(), sdk.ZeroInt()))
	s.Require().Len(s.keeper.GetParams(s.ctx).LiquidFarms, 1)

	pool, found := s.app.LiquidityKeeper.GetPool(s.ctx, 1)
	s.Require().True(found)

	farmerAddr := s.addr(1)

	// Farm 1000 in total (400, 300, 400)
	amount1, amount2, amount3 := sdk.NewInt(400_000_000), sdk.NewInt(300_000_000), sdk.NewInt(300_000_000)

	s.farm(pool.Id, farmerAddr, sdk.NewCoin(pool.PoolCoinDenom, amount1), true)
	s.nextBlock()

	s.farm(pool.Id, farmerAddr, sdk.NewCoin(pool.PoolCoinDenom, amount2), true)
	s.nextBlock()

	s.farm(pool.Id, farmerAddr, sdk.NewCoin(pool.PoolCoinDenom, amount3), true)
	s.nextBlock()

	queuedFarmings := s.keeper.GetQueuedFarmingsByFarmer(s.ctx, farmerAddr)
	s.Require().Len(queuedFarmings, 3)

	// Cancel partial amounts
	unfarmingCoin := sdk.NewInt64Coin(pool.PoolCoinDenom, 650_000_000)
	s.cancelQueuedFarming(pool.Id, farmerAddr, unfarmingCoin)

	// Two queuedFarmings must be deleted and the last one must be in store
	queuedFarmings = s.keeper.GetQueuedFarmingsByFarmer(s.ctx, farmerAddr)
	s.Require().Len(queuedFarmings, 1)

	// Check the remaining queued coin
	total := amount1.Add(amount2).Add(amount3)
	s.Require().Equal(total.Sub(unfarmingCoin.Amount), queuedFarmings[0].Amount)

	// Unfarming amount must be returned to the farmer
	balance := s.getBalance(farmerAddr, pool.PoolCoinDenom)
	s.Require().Equal(unfarmingCoin, balance)
}

func (s *KeeperTestSuite) TestCancelQueuedFarming_Exceed() {
	s.createPair(s.addr(0), "denom1", "denom2", true)
	s.createPool(s.addr(0), 1, utils.ParseCoins("100_000_000denom1, 100_000_000denom2"), true)
	s.createLiquidFarm(types.NewLiquidFarm(1, sdk.ZeroInt(), sdk.ZeroInt()))
	s.Require().Len(s.keeper.GetParams(s.ctx).LiquidFarms, 1)

	pool, found := s.app.LiquidityKeeper.GetPool(s.ctx, 1)
	s.Require().True(found)

	farmerAddr := s.addr(1)

	// Farm 1000 in total (400, 300, 400)
	amount1, amount2, amount3 := sdk.NewInt(400_000_000), sdk.NewInt(300_000_000), sdk.NewInt(300_000_000)

	s.farm(pool.Id, farmerAddr, sdk.NewCoin(pool.PoolCoinDenom, amount1), true)
	s.nextBlock()

	s.farm(pool.Id, farmerAddr, sdk.NewCoin(pool.PoolCoinDenom, amount2), true)
	s.nextBlock()

	s.farm(pool.Id, farmerAddr, sdk.NewCoin(pool.PoolCoinDenom, amount3), true)
	s.nextBlock()

	queuedFarmings := s.keeper.GetQueuedFarmingsByFarmer(s.ctx, farmerAddr)
	s.Require().Len(queuedFarmings, 3)

	// Cancel exceeding amount
	unfarmingCoin := sdk.NewInt64Coin(pool.PoolCoinDenom, 1_500_000_000)
	err := s.keeper.CancelQueuedFarming(s.ctx, &types.MsgCancelQueuedFarming{
		PoolId:        pool.Id,
		Farmer:        farmerAddr.String(),
		UnfarmingCoin: unfarmingCoin,
	})
	s.Require().ErrorIs(err, sdkerrors.ErrInsufficientFunds)
}

func (s *KeeperTestSuite) TestUnfarm_All() {
	s.createPair(s.addr(0), "denom1", "denom2", true)
	s.createPool(s.addr(0), 1, utils.ParseCoins("100_000_000denom1, 100_000_000denom2"), true)

	s.createLiquidFarm(types.NewLiquidFarm(1, sdk.ZeroInt(), sdk.ZeroInt()))
	s.Require().Len(s.keeper.GetParams(s.ctx).LiquidFarms, 1)

	pool, found := s.app.LiquidityKeeper.GetPool(s.ctx, 1)
	s.Require().True(found)

	farmerAddr := s.addr(1)
	amount1 := sdk.NewInt(100_000_000)

	s.farm(pool.Id, farmerAddr, sdk.NewCoin(pool.PoolCoinDenom, amount1), true)
	s.advanceEpochDays()

	// Staked amount must be 100
	totalStakings, found := s.app.FarmingKeeper.GetTotalStakings(s.ctx, pool.PoolCoinDenom)
	s.Require().True(found)
	s.Require().Equal(amount1, totalStakings.Amount)

	queuedFarmings := s.keeper.GetQueuedFarmingsByFarmer(s.ctx, farmerAddr)
	s.Require().Len(queuedFarmings, 0)

	// Check minted LFCoin
	lfCoinDenom := types.LiquidFarmCoinDenom(pool.Id)
	lfCoinBalance := s.getBalance(farmerAddr, lfCoinDenom)
	s.Require().Equal(sdk.NewCoin(lfCoinDenom, amount1), lfCoinBalance)

	// Unfarm all amounts
	_, err := s.keeper.Unfarm(s.ctx, pool.Id, farmerAddr, lfCoinBalance)
	s.Require().NoError(err)

	// Verify
	_, found = s.app.FarmingKeeper.GetTotalStakings(s.ctx, pool.PoolCoinDenom)
	s.Require().False(found)

	supply := s.app.BankKeeper.GetSupply(s.ctx, lfCoinDenom)
	s.Require().Equal(sdk.ZeroInt(), supply.Amount)
}

func (s *KeeperTestSuite) TestUnfarmAndWithdraw() {
	pair := s.createPair(s.addr(0), "denom1", "denom2", true)
	depositCoins := utils.ParseCoins("100_000_000denom1, 100_000_000denom2")
	pool := s.createPool(s.addr(0), pair.Id, depositCoins, true)

	s.createLiquidFarm(types.NewLiquidFarm(pool.Id, sdk.ZeroInt(), sdk.ZeroInt()))
	s.Require().Len(s.keeper.GetParams(s.ctx).LiquidFarms, 1)

	poolAmt := sdk.NewInt(1_000_000_000_000)

	s.farm(pool.Id, s.addr(0), sdk.NewCoin(pool.PoolCoinDenom, poolAmt), false)
	s.advanceEpochDays()

	// Staked amount must be 100
	totalStakings, found := s.app.FarmingKeeper.GetTotalStakings(s.ctx, pool.PoolCoinDenom)
	s.Require().True(found)
	s.Require().Equal(poolAmt, totalStakings.Amount)

	// Check minted LFCoin
	lfCoinDenom := types.LiquidFarmCoinDenom(pool.Id)
	lfCoinBalance := s.getBalance(s.addr(0), lfCoinDenom)
	s.Require().Equal(sdk.NewCoin(lfCoinDenom, poolAmt), lfCoinBalance)

	// Call UnfarmAndWithdraw
	err := s.keeper.UnfarmAndWithdraw(s.ctx, types.NewMsgUnfarmAndWithdraw(pool.Id, s.addr(0).String(), lfCoinBalance))
	s.Require().NoError(err)

	// Call nextBlock as Withdraw is executed in batch
	s.nextBlock()

	// Verify
	s.Require().EqualValues(depositCoins, s.getBalances(s.addr(0)))
}

func (s *KeeperTestSuite) TestTerminateLiquidFarm() {
	pair1 := s.createPair(s.addr(0), "denom1", "denom2", true)
	pool1 := s.createPool(s.addr(0), pair1.Id, utils.ParseCoins("100_000_000denom1, 100_000_000denom2"), true)
	pair2 := s.createPair(s.addr(0), "denom3", "denom4", true)
	pool2 := s.createPool(s.addr(0), pair2.Id, utils.ParseCoins("100_000_000denom3, 100_000_000denom4"), true)

	// Add liquid farms in params
	params := s.keeper.GetParams(s.ctx)
	params.LiquidFarms = append(params.LiquidFarms, types.NewLiquidFarm(pool1.Id, sdk.NewInt(100), sdk.NewInt(100)))
	params.LiquidFarms = append(params.LiquidFarms, types.NewLiquidFarm(pool2.Id, sdk.NewInt(500), sdk.NewInt(500)))
	s.keeper.SetParams(s.ctx, params)
	s.Require().Len(s.keeper.GetParams(s.ctx).LiquidFarms, len(params.LiquidFarms))

	// Execute BeginBlocker to store all registered LiquidFarms in params
	liquidfarming.BeginBlocker(s.ctx, s.keeper)

	// Ensure the length of liquid farms in the store and params are the same
	s.Require().Len(s.keeper.GetAllLiquidFarms(s.ctx), 2)
	s.Require().Len(s.keeper.GetParams(s.ctx).LiquidFarms, 2)

	// Forcefully remove one liquid farm from the store
	liquidFarm, found := s.keeper.GetLiquidFarm(s.ctx, pool1.Id)
	s.Require().True(found)
	s.keeper.DeleteLiquidFarm(s.ctx, liquidFarm)
	s.Require().Len(s.keeper.GetAllLiquidFarms(s.ctx), 1)
	s.Require().Len(s.keeper.GetParams(s.ctx).LiquidFarms, 2)

	// Execute BeginBlocker again to store all registered LiquidFarms in params
	liquidfarming.BeginBlocker(s.ctx, s.keeper)

	// Ensure the length of liquid farms in the store and params are the same
	s.Require().Len(s.keeper.GetAllLiquidFarms(s.ctx), 2)
	s.Require().Len(s.keeper.GetParams(s.ctx).LiquidFarms, 2)

	// Now, in order to test adding new liquid farm in the params
	// Farm some coins
	s.farm(pool1.Id, s.addr(1), sdk.NewCoin(pool1.PoolCoinDenom, sdk.NewInt(200_000)), true)
	s.farm(pool1.Id, s.addr(2), sdk.NewCoin(pool1.PoolCoinDenom, sdk.NewInt(200_000)), true)
	s.advanceEpochDays() // trigger AfterStaked hook to mint LFCoin and delete QueuedFarming
	s.advanceEpochDays() // trigger AllocateRewards hook to create rewards auctions

	s.placeBid(pool1.Id, s.addr(3), utils.ParseCoin("100_000pool1"), true)
	s.placeBid(pool1.Id, s.addr(4), utils.ParseCoin("200_000pool1"), true)
	s.placeBid(pool1.Id, s.addr(5), utils.ParseCoin("500_000pool1"), true)

	// Remove the first liquid farm object in params
	params = s.keeper.GetParams(s.ctx)
	params.LiquidFarms = []types.LiquidFarm{types.NewLiquidFarm(pool2.Id, sdk.NewInt(100), sdk.NewInt(100))}
	s.keeper.SetParams(s.ctx, params)

	// Execute BeginBlocker again to store all registered LiquidFarms in params
	liquidfarming.BeginBlocker(s.ctx, s.keeper)

	// Ensure that bidders got their funds back
	s.Require().Equal(utils.ParseCoin("100_000pool1"), s.getBalance(s.addr(3), pool1.PoolCoinDenom))
	s.Require().Equal(utils.ParseCoin("200_000pool1"), s.getBalance(s.addr(4), pool1.PoolCoinDenom))
	s.Require().Equal(utils.ParseCoin("500_000pool1"), s.getBalance(s.addr(5), pool1.PoolCoinDenom))
	s.Require().Equal(utils.ParseCoin("400_000pool1"), s.getBalance(types.LiquidFarmReserveAddress(1), pool1.PoolCoinDenom))

	// Ensure the auction status
	auction, found := s.keeper.GetRewardsAuction(s.ctx, pool1.Id, 1)
	s.Require().True(found)
	s.Require().Equal(types.AuctionStatusFinished, auction.Status)

	// Ensure all staked coins is zero
	s.Require().True(s.app.FarmingKeeper.GetAllStakedCoinsByFarmer(s.ctx, types.LiquidFarmReserveAddress(1)).IsZero())
}
