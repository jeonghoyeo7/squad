package keeper_test

import (
	"encoding/binary"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/suite"

	sdk "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"

	chain "github.com/cosmosquad-labs/squad/v2/app"
	farmingtypes "github.com/cosmosquad-labs/squad/v2/x/farming/types"
	"github.com/cosmosquad-labs/squad/v2/x/liquidfarming/keeper"
	"github.com/cosmosquad-labs/squad/v2/x/liquidfarming/types"
	liquiditytypes "github.com/cosmosquad-labs/squad/v2/x/liquidity/types"
)

type KeeperTestSuite struct {
	suite.Suite

	app       *chain.App
	ctx       sdk.Context
	keeper    keeper.Keeper
	querier   keeper.Querier
	msgServer types.MsgServer
}

func TestKeeperTestSuite(t *testing.T) {
	suite.Run(t, new(KeeperTestSuite))
}

func (s *KeeperTestSuite) SetupTest() {
	s.app = chain.Setup(false)
	s.ctx = s.app.BaseApp.NewContext(false, tmproto.Header{})
	s.keeper = s.app.LiquidFarmingKeeper
	s.querier = keeper.Querier{Keeper: s.keeper}
	s.msgServer = keeper.NewMsgServerImpl(s.keeper)
}

//
// Below are just shortcuts to frequently-used functions.
//

func (s *KeeperTestSuite) fundAddr(addr sdk.AccAddress, amt sdk.Coins) {
	s.T().Helper()
	err := s.app.BankKeeper.MintCoins(s.ctx, types.ModuleName, amt)
	s.Require().NoError(err)
	err = s.app.BankKeeper.SendCoinsFromModuleToAccount(s.ctx, types.ModuleName, addr, amt)
	s.Require().NoError(err)
}

func (s *KeeperTestSuite) createPrivateFixedAmountPlan(
	creator sdk.AccAddress, stakingCoinWeights sdk.DecCoins,
	startTime, endTime time.Time, epochAmt sdk.Coins) (farmingtypes.PlanI, error) {
	msg := farmingtypes.NewMsgCreateFixedAmountPlan(
		fmt.Sprintf("plan%d", s.app.FarmingKeeper.GetGlobalPlanId(s.ctx)+1),
		creator, stakingCoinWeights,
		startTime, endTime, epochAmt,
	)
	farmingPoolAcc, err := s.app.FarmingKeeper.DerivePrivatePlanFarmingPoolAcc(s.ctx, msg.Name)
	if err != nil {
		return nil, err
	}
	plan, err := s.app.FarmingKeeper.CreateFixedAmountPlan(s.ctx, msg, farmingPoolAcc, creator, farmingtypes.PlanTypePrivate)
	if err != nil {
		return nil, err
	}
	return plan, nil
}

func (s *KeeperTestSuite) stake(farmerAcc sdk.AccAddress, amt sdk.Coins) {
	s.T().Helper()
	err := s.app.FarmingKeeper.Stake(s.ctx, farmerAcc, amt)
	s.Require().NoError(err)
}

func (s *KeeperTestSuite) unstake(farmerAcc sdk.AccAddress, amt sdk.Coins) {
	s.T().Helper()
	err := s.app.FarmingKeeper.Unstake(s.ctx, farmerAcc, amt)
	s.Require().NoError(err)
}

func (s *KeeperTestSuite) harvest(farmerAcc sdk.AccAddress, stakingCoinDenoms []string) {
	s.T().Helper()
	err := s.app.FarmingKeeper.Harvest(s.ctx, farmerAcc, stakingCoinDenoms)
	s.Require().NoError(err)
}

func (s *KeeperTestSuite) createPair(creator sdk.AccAddress, baseCoinDenom, quoteCoinDenom string, fund bool) liquiditytypes.Pair {
	s.T().Helper()
	params := s.app.LiquidityKeeper.GetParams(s.ctx)
	if fund {
		s.fundAddr(creator, params.PairCreationFee)
	}
	pair, err := s.app.LiquidityKeeper.CreatePair(s.ctx, liquiditytypes.NewMsgCreatePair(creator, baseCoinDenom, quoteCoinDenom))
	s.Require().NoError(err)
	return pair
}

func (s *KeeperTestSuite) createPool(creator sdk.AccAddress, pairId uint64, depositCoins sdk.Coins, fund bool) liquiditytypes.Pool {
	s.T().Helper()
	params := s.app.LiquidityKeeper.GetParams(s.ctx)
	if fund {
		s.fundAddr(creator, depositCoins.Add(params.PoolCreationFee...))
	}
	pool, err := s.app.LiquidityKeeper.CreatePool(s.ctx, liquiditytypes.NewMsgCreatePool(creator, pairId, depositCoins))
	s.Require().NoError(err)
	return pool
}

func (s *KeeperTestSuite) depositLiquidity(depositor sdk.AccAddress, poolId uint64, depositCoins sdk.Coins, fund bool) liquiditytypes.DepositRequest {
	s.T().Helper()
	if fund {
		s.fundAddr(depositor, depositCoins)
	}
	req, err := s.app.LiquidityKeeper.Deposit(s.ctx, liquiditytypes.NewMsgDeposit(depositor, poolId, depositCoins))
	s.Require().NoError(err)
	return req
}

func (s *KeeperTestSuite) createLiquidFarm(liquidFarms []types.LiquidFarm) {
	params := s.keeper.GetParams(s.ctx)
	params.LiquidFarms = liquidFarms
	s.keeper.SetParams(s.ctx, params)
}

func (s *KeeperTestSuite) farm(poolId uint64, farmer sdk.AccAddress, farmingCoin sdk.Coin, fund bool) {
	s.T().Helper()
	if fund {
		s.fundAddr(farmer, sdk.NewCoins(farmingCoin))
	}
	err := s.keeper.Farm(s.ctx, &types.MsgFarm{
		PoolId:      poolId,
		Farmer:      farmer.String(),
		FarmingCoin: farmingCoin,
	})
	s.Require().NoError(err)
}

func (s *KeeperTestSuite) unfarm(poolId uint64, farmer sdk.AccAddress, lfCoin sdk.Coin, fund bool) {
	s.T().Helper()
	if fund {
		s.fundAddr(farmer, sdk.NewCoins(lfCoin))
	}
	err := s.keeper.Unfarm(s.ctx, &types.MsgUnfarm{
		PoolId: poolId,
		Farmer: farmer.String(),
		LFCoin: lfCoin,
	})
	s.Require().NoError(err)
}

//
// Below are helper functions to write test code easily
//

func (s *KeeperTestSuite) addr(addrNum int) sdk.AccAddress {
	addr := make(sdk.AccAddress, 20)
	binary.PutVarint(addr, int64(addrNum))
	return addr
}

func (s *KeeperTestSuite) getBalances(addr sdk.AccAddress) sdk.Coins {
	return s.app.BankKeeper.GetAllBalances(s.ctx, addr)
}

func (s *KeeperTestSuite) getBalance(addr sdk.AccAddress, denom string) sdk.Coin {
	return s.app.BankKeeper.GetBalance(s.ctx, addr, denom)
}

func (s *KeeperTestSuite) sendCoins(fromAddr, toAddr sdk.AccAddress, amt sdk.Coins) {
	s.T().Helper()
	err := s.app.BankKeeper.SendCoins(s.ctx, fromAddr, toAddr, amt)
	s.Require().NoError(err)
}

func (s *KeeperTestSuite) nextBlock() {
	s.T().Helper()
	s.app.EndBlock(abci.RequestEndBlock{})
	s.app.Commit()
	hdr := tmproto.Header{
		Height: s.app.LastBlockHeight() + 1,
		Time:   s.ctx.BlockTime().Add(5 * time.Second),
	}
	s.app.BeginBlock(abci.RequestBeginBlock{Header: hdr})
	s.ctx = s.app.BaseApp.NewContext(false, hdr)
	s.app.BeginBlocker(s.ctx, abci.RequestBeginBlock{Header: hdr})
}
