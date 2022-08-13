package testutil

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	servertypes "github.com/cosmos/cosmos-sdk/server/types"
	"github.com/cosmos/cosmos-sdk/simapp"
	store "github.com/cosmos/cosmos-sdk/store/types"
	utilcli "github.com/cosmos/cosmos-sdk/testutil/cli"
	"github.com/cosmos/cosmos-sdk/testutil/network"
	sdk "github.com/cosmos/cosmos-sdk/types"
	tmcli "github.com/tendermint/tendermint/libs/cli"
	dbm "github.com/tendermint/tm-db"

	chain "github.com/cosmosquad-labs/squad/v2/app"
	"github.com/cosmosquad-labs/squad/v2/app/params"
	farmingkeeper "github.com/cosmosquad-labs/squad/v2/x/farming/keeper"
	"github.com/cosmosquad-labs/squad/v2/x/liquidfarming/client/cli"
	"github.com/cosmosquad-labs/squad/v2/x/liquidfarming/types"
	liquiditytestutil "github.com/cosmosquad-labs/squad/v2/x/liquidity/client/testutil"
)

type IntegrationTestSuite struct {
	suite.Suite

	cfg       network.Config
	network   *network.Network
	val       *network.Validator
	clientCtx client.Context

	denom1, denom2 string
}

func NewAppConstructor(encodingCfg params.EncodingConfig) network.AppConstructor {
	return func(val network.Validator) servertypes.Application {
		return chain.NewApp(
			val.Ctx.Logger, dbm.NewMemDB(), nil, true, make(map[int64]bool), val.Ctx.Config.RootDir, 0,
			encodingCfg,
			simapp.EmptyAppOptions{},
			baseapp.SetPruning(store.NewPruningOptionsFromString(val.AppConfig.Pruning)),
			baseapp.SetMinGasPrices(val.AppConfig.MinGasPrices),
		)
	}
}

// SetupTest creates a new network for _each_ integration test. We create a new
// network for each test because there are some state modifications that are
// needed to be made in order to make useful queries. However, we don't want
// these state changes to be present in other tests.
func (s *IntegrationTestSuite) SetupTest() {
	s.T().Log("setting up integration test suite")

	if testing.Short() {
		s.T().Skip("skipping test in unit-tests mode.")
	}

	farmingkeeper.EnableAdvanceEpoch = true
	farmingkeeper.EnableRatioPlan = true

	encCfg := chain.MakeTestEncodingConfig()

	cfg := network.DefaultConfig()
	cfg.AppConstructor = NewAppConstructor(encCfg)
	cfg.GenesisState = chain.ModuleBasics.DefaultGenesis(cfg.Codec)
	cfg.NumValidators = 1

	var genesisState types.GenesisState
	err := cfg.Codec.UnmarshalJSON(cfg.GenesisState[types.ModuleName], &genesisState)
	s.Require().NoError(err)

	genesisState.Params = types.DefaultParams()
	genesisState.Params.LiquidFarms = []types.LiquidFarm{
		{
			PoolId:        1,
			MinFarmAmount: sdk.NewInt(100_000),
			MinBidAmount:  sdk.NewInt(100_000),
		},
	}
	cfg.GenesisState[types.ModuleName] = cfg.Codec.MustMarshalJSON(&genesisState)

	s.cfg = cfg
	s.network = network.New(s.T(), cfg)

	s.val = s.network.Validators[0]
	s.clientCtx = s.val.ClientCtx

	s.denom1, s.denom2 = fmt.Sprintf("%stoken", s.val.Moniker), s.cfg.BondDenom

	s.createPair(s.denom1, s.denom2)
	s.createPool(1, sdk.NewCoins(sdk.NewInt64Coin(s.denom1, 10000000), sdk.NewInt64Coin(s.denom2, 10000000)))

	_, err = s.network.WaitForHeight(1)
	s.Require().NoError(err)
}

// TearDownTest cleans up the current test network after each test in the suite.
func (s *IntegrationTestSuite) TearDownTest() {
	s.T().Log("tearing down integration test suite")
	s.network.Cleanup()
}

//
// Helper functions
//

func (s *IntegrationTestSuite) createPair(baseCoinDenom, quoteCoinDenom string) {
	_, err := liquiditytestutil.MsgCreatePair(s.clientCtx, s.val.Address.String(), baseCoinDenom, quoteCoinDenom)
	s.Require().NoError(err)

	err = s.network.WaitForNextBlock()
	s.Require().NoError(err)
}

func (s *IntegrationTestSuite) createPool(pairId uint64, depositCoins sdk.Coins) {
	_, err := liquiditytestutil.MsgCreatePool(s.clientCtx, s.val.Address.String(), pairId, depositCoins)
	s.Require().NoError(err)

	err = s.network.WaitForNextBlock()
	s.Require().NoError(err)
}

//
// Query CLI Integration Tests
//

func (s *IntegrationTestSuite) TestCmdQueryParams() {
	val := s.network.Validators[0]
	clientCtx := val.ClientCtx

	testCases := []struct {
		name      string
		args      []string
		expectErr bool
	}{
		{
			"happy case",
			[]string{fmt.Sprintf("--%s=json", tmcli.OutputFlag)},
			false,
		},
		{
			"with specific height",
			[]string{fmt.Sprintf("--%s=1", flags.FlagHeight), fmt.Sprintf("--%s=json", tmcli.OutputFlag)},
			false,
		},
	}

	for _, tc := range testCases {
		s.Run(tc.name, func() {
			cmd := cli.NewQueryParamsCmd()

			out, err := utilcli.ExecTestCLICmd(clientCtx, cmd, tc.args)
			if tc.expectErr {
				s.Require().Error(err)
				s.Require().NotEqual("internal", err.Error())
			} else {
				s.Require().NoError(err)

				var params types.Params
				s.Require().NoError(clientCtx.Codec.UnmarshalJSON(out.Bytes(), &params))
				s.Require().NotEmpty(params.LiquidFarms)
			}
		})
	}
}
