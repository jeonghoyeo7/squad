package testutil

import (
	"github.com/stretchr/testify/suite"

	"github.com/cosmos/cosmos-sdk/testutil/network"
	sdk "github.com/cosmos/cosmos-sdk/types"
	tmdb "github.com/tendermint/tm-db"

	chain "github.com/cosmosquad-labs/squad/v2/app"
	farmingkeeper "github.com/cosmosquad-labs/squad/v2/x/farming/keeper"
	"github.com/cosmosquad-labs/squad/v2/x/liquidfarming/types"
)

type IntegrationTestSuite struct {
	suite.Suite

	cfg     network.Config
	network *network.Network
}

// SetupTest creates a new network for _each_ integration test. We create a new
// network for each test because there are some state modifications that are
// needed to be made in order to make useful queries. However, we don't want
// these state changes to be present in other tests.
func (s *IntegrationTestSuite) SetupTest() {
	s.T().Log("setting up integration test suite")

	farmingkeeper.EnableAdvanceEpoch = true
	farmingkeeper.EnableRatioPlan = true

	db := tmdb.NewMemDB()
	cfg := chain.NewConfig(db)
	cfg.NumValidators = 1

	var genesisState types.GenesisState
	err := cfg.Codec.UnmarshalJSON(cfg.GenesisState[types.ModuleName], &genesisState)
	s.Require().NoError(err)

	genesisState.Params = types.DefaultParams()
	cfg.GenesisState[types.ModuleName] = cfg.Codec.MustMarshalJSON(&genesisState)
	cfg.AccountTokens = sdk.NewInt(100_000_000_000) // node0token denom
	cfg.StakingTokens = sdk.NewInt(100_000_000_000) // stake denom

	s.cfg = cfg
	s.network = network.New(s.T(), cfg)

	_, err = s.network.WaitForHeight(1)
	s.Require().NoError(err)
}

// TearDownTest cleans up the current test network after each test in the suite.
func (s *IntegrationTestSuite) TearDownTest() {
	s.T().Log("tearing down integration test suite")
	s.network.Cleanup()
}

func (s *IntegrationTestSuite) TestNewFarmCmd() {
	// TODO: not implemented yet
}

func (s *IntegrationTestSuite) TestNewUnfarmCmd() {
	// TODO: not implemented yet
}

func (s *IntegrationTestSuite) TestNewPlaceBidCmd() {
	// TODO: not implemented yet
}

func (s *IntegrationTestSuite) TestNewRefundBidCmd() {
	// TODO: not implemented yet
}
