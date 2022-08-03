package types_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/tendermint/tendermint/crypto"

	utils "github.com/cosmosquad-labs/squad/v2/types"
	"github.com/cosmosquad-labs/squad/v2/x/liquidfarming/types"
)

func TestGenesisState_Validate(t *testing.T) {
	validFarmer := sdk.AccAddress(crypto.AddressHash([]byte("validFarmer")))
	validPayingReserveAddr := sdk.AccAddress(crypto.AddressHash([]byte("validPayingReserveAddr")))
	validBidder := sdk.AccAddress(crypto.AddressHash([]byte("validBidder")))
	validDenom := "denom1"
	validPoolId := uint64(1)
	validAuctionId := uint64(1)

	for _, tc := range []struct {
		name        string
		malleate    func(genState *types.GenesisState)
		expectedErr string
	}{
		{
			"default is valid",
			func(genState *types.GenesisState) {},
			"",
		},
		{
			"invalid queued farming record: farming coin denom",
			func(genState *types.GenesisState) {
				genState.QueuedFarmingRecords = []types.QueuedFarmingRecord{
					{
						EndTime:          utils.ParseTime("2022-08-01T00:00:00Z"),
						Farmer:           validFarmer.String(),
						FarmingCoinDenom: "123!@#$%",
						QueuedFarming: types.QueuedFarming{
							PoolId: validPoolId,
							Amount: sdk.NewInt(100),
						},
					},
				}
			},
			"invalid farming coin denom",
		},
		{
			"invalid queued farming record: pool id",
			func(genState *types.GenesisState) {
				genState.QueuedFarmingRecords = []types.QueuedFarmingRecord{
					{
						EndTime:          utils.ParseTime("2022-08-01T00:00:00Z"),
						Farmer:           validFarmer.String(),
						FarmingCoinDenom: validDenom,
						QueuedFarming: types.QueuedFarming{
							PoolId: 0,
							Amount: sdk.NewInt(100),
						},
					},
				}
			},
			"pool id must not be 0",
		},
		{
			"invalid queued farming record: queued farming amount",
			func(genState *types.GenesisState) {
				genState.QueuedFarmingRecords = []types.QueuedFarmingRecord{
					{
						EndTime:          utils.ParseTime("2022-08-01T00:00:00Z"),
						Farmer:           validFarmer.String(),
						FarmingCoinDenom: validDenom,
						QueuedFarming: types.QueuedFarming{
							PoolId: validPoolId,
							Amount: sdk.NewInt(0),
						},
					},
				}
			},
			"amount must be positive value",
		},
		{
			"invalid rewards auction: pool id",
			func(genState *types.GenesisState) {
				genState.RewardsAuctions = []types.RewardsAuction{
					{
						PoolId:               0,
						BiddingCoinDenom:     validDenom,
						PayingReserveAddress: validPayingReserveAddr.String(),
						StartTime:            utils.ParseTime("2022-08-01T00:00:00Z"),
						EndTime:              utils.ParseTime("2022-08-02T00:00:00Z"),
						Status:               types.AuctionStatusStarted,
					},
				}
			},
			"pool id must not be 0",
		},
		{
			"invalid rewards auction: bidding coin denom",
			func(genState *types.GenesisState) {
				genState.RewardsAuctions = []types.RewardsAuction{
					{
						PoolId:               validPoolId,
						BiddingCoinDenom:     "123!@#$%",
						PayingReserveAddress: validPayingReserveAddr.String(),
						StartTime:            utils.ParseTime("2022-08-01T00:00:00Z"),
						EndTime:              utils.ParseTime("2022-08-02T00:00:00Z"),
						Status:               types.AuctionStatusStarted,
					},
				}
			},
			"invalid coin denom",
		},
		{
			"invalid rewards auction: auction status",
			func(genState *types.GenesisState) {
				genState.RewardsAuctions = []types.RewardsAuction{
					{
						PoolId:               validPoolId,
						BiddingCoinDenom:     validDenom,
						PayingReserveAddress: validPayingReserveAddr.String(),
						StartTime:            utils.ParseTime("2022-08-01T00:00:00Z"),
						EndTime:              utils.ParseTime("2022-08-02T00:00:00Z"),
						Status:               types.AuctionStatusNil,
					},
				}
			},
			"invalid auction status",
		},
		{
			"invalid bids: pool id",
			func(genState *types.GenesisState) {
				genState.Bids = []types.Bid{
					{
						PoolId: 0,
						Bidder: validBidder.String(),
						Amount: utils.ParseCoin("1000000pool1"),
					},
				}
			},
			"pool id must not be 0",
		},
		{
			"invalid bids: bid amount",
			func(genState *types.GenesisState) {
				genState.Bids = []types.Bid{
					{
						PoolId: validPoolId,
						Bidder: validBidder.String(),
						Amount: utils.ParseCoin("0pool1"),
					},
				}
			},
			"amount must be positive value",
		},
		{
			"invalid winning bid records: auction id",
			func(genState *types.GenesisState) {
				genState.WinningBidRecords = []types.WinningBidRecord{
					{
						AuctionId: 0,
						WinningBid: types.Bid{
							PoolId: validPoolId,
							Bidder: validBidder.String(),
							Amount: utils.ParseCoin("1000000pool1"),
						},
					},
				}
			},
			"auction id must not be 0",
		},
		{
			"invalid winning bid records: pool id",
			func(genState *types.GenesisState) {
				genState.WinningBidRecords = []types.WinningBidRecord{
					{
						AuctionId: validAuctionId,
						WinningBid: types.Bid{
							PoolId: 0,
							Bidder: validBidder.String(),
							Amount: utils.ParseCoin("1000000pool1"),
						},
					},
				}
			},
			"invalid winning bid: pool id must not be 0",
		},
		{
			"invalid winning bid records: pool id",
			func(genState *types.GenesisState) {
				genState.WinningBidRecords = []types.WinningBidRecord{
					{
						AuctionId: validAuctionId,
						WinningBid: types.Bid{
							PoolId: validPoolId,
							Bidder: validBidder.String(),
							Amount: utils.ParseCoin("0pool1"),
						},
					},
				}
			},
			"invalid winning bid: amount must be positive value",
		},
		{
			"invalid winning bid records: duplicate winning bid",
			func(genState *types.GenesisState) {
				genState.WinningBidRecords = []types.WinningBidRecord{
					{
						AuctionId: validAuctionId,
						WinningBid: types.Bid{
							PoolId: validPoolId,
							Bidder: validBidder.String(),
							Amount: utils.ParseCoin("1000000pool1"),
						},
					},
					{
						AuctionId: validAuctionId,
						WinningBid: types.Bid{
							PoolId: validPoolId,
							Bidder: validBidder.String(),
							Amount: utils.ParseCoin("1000000pool1"),
						},
					},
				}
			},
			"multiple winning bids at auction id: 1",
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			genState := types.DefaultGenesis()
			genState.QueuedFarmingRecords = []types.QueuedFarmingRecord{}
			genState.RewardsAuctions = []types.RewardsAuction{}
			genState.Bids = []types.Bid{}
			genState.WinningBidRecords = []types.WinningBidRecord{}
			tc.malleate(genState)
			err := genState.Validate()
			if tc.expectedErr == "" {
				require.NoError(t, err)
			} else {
				require.EqualError(t, err, tc.expectedErr)
			}
		})
	}
}
