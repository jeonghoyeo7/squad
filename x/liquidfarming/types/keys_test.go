package types_test

import (
	fmt "fmt"
	"testing"
	time "time"

	"github.com/stretchr/testify/suite"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/tendermint/tendermint/crypto"

	utils "github.com/cosmosquad-labs/squad/v2/types"
	"github.com/cosmosquad-labs/squad/v2/x/liquidfarming/types"
)

type keysTestSuite struct {
	suite.Suite
}

func TestKeysTestSuite(t *testing.T) {
	suite.Run(t, new(keysTestSuite))
}

func (s *keysTestSuite) TestGetLastRewardsAuctionIdKey() {
	s.Require().Equal([]byte{0xe1, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0}, types.GetLastRewardsAuctionIdKey(0))
	s.Require().Equal([]byte{0xe1, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x9}, types.GetLastRewardsAuctionIdKey(9))
	s.Require().Equal([]byte{0xe1, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0xa}, types.GetLastRewardsAuctionIdKey(10))
}

func (s *keysTestSuite) TestGetLiquidFarmKey() {
	s.Require().Equal([]byte{0xe3, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0}, types.GetLiquidFarmKey(0))
	s.Require().Equal([]byte{0xe3, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x9}, types.GetLiquidFarmKey(9))
	s.Require().Equal([]byte{0xe3, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0xa}, types.GetLiquidFarmKey(10))
}

func (s *keysTestSuite) TestGetQueuedFarmingKey() {
	testCases := []struct {
		endTime          time.Time
		farmingCoinDenom string
		farmerAcc        sdk.AccAddress
		expected         []byte
	}{
		{
			utils.ParseTime("2022-01-01T00:00:00Z"),
			sdk.DefaultBondDenom,
			sdk.AccAddress(crypto.AddressHash([]byte("farmer1"))),
			[]byte{0xe4, 0x1d, 0x32, 0x30, 0x32, 0x32, 0x2d, 0x30, 0x31, 0x2d, 0x30, 0x31, 0x54, 0x30,
				0x30, 0x3a, 0x30, 0x30, 0x3a, 0x30, 0x30, 0x2e, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30,
				0x30, 0x30, 0x30, 0x5, 0x73, 0x74, 0x61, 0x6b, 0x65, 0xd3, 0x7a, 0x85, 0xec, 0x75,
				0xf, 0x3, 0xaa, 0xe5, 0x36, 0xcf, 0x1b, 0xb7, 0x59, 0xb7, 0xbc, 0xbd, 0x5c, 0xfe, 0x3d},
		},
		{
			utils.ParseTime("2022-06-10T00:02:10Z"),
			sdk.DefaultBondDenom,
			sdk.AccAddress(crypto.AddressHash([]byte("farmer2"))),
			[]byte{0xe4, 0x1d, 0x32, 0x30, 0x32, 0x32, 0x2d, 0x30, 0x36, 0x2d, 0x31, 0x30, 0x54, 0x30,
				0x30, 0x3a, 0x30, 0x32, 0x3a, 0x31, 0x30, 0x2e, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30,
				0x30, 0x30, 0x30, 0x5, 0x73, 0x74, 0x61, 0x6b, 0x65, 0x15, 0x1, 0x20, 0x25, 0x5a, 0x5d,
				0xe8, 0x6b, 0xa1, 0xed, 0xfb, 0x6f, 0x45, 0x48, 0xcb, 0xfb, 0x6f, 0x28, 0x66, 0xf3},
		},
		{
			utils.ParseTime("2022-08-15T00:05:05Z"),
			sdk.DefaultBondDenom,
			sdk.AccAddress(crypto.AddressHash([]byte("farmer3"))),
			[]byte{0xe4, 0x1d, 0x32, 0x30, 0x32, 0x32, 0x2d, 0x30, 0x38, 0x2d, 0x31, 0x35, 0x54, 0x30,
				0x30, 0x3a, 0x30, 0x35, 0x3a, 0x30, 0x35, 0x2e, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30,
				0x30, 0x30, 0x30, 0x5, 0x73, 0x74, 0x61, 0x6b, 0x65, 0xdf, 0xb0, 0x6d, 0xbf, 0xc6, 0x9a,
				0xcd, 0xf5, 0x7b, 0xb, 0xe7, 0x69, 0x75, 0x50, 0x9e, 0x69, 0x54, 0xa6, 0x1e, 0xe2},
		},
	}

	for i, tc := range testCases {
		s.Run(fmt.Sprint(i), func() {
			key := types.GetQueuedFarmingKey(tc.endTime, tc.farmingCoinDenom, tc.farmerAcc)
			s.Require().Equal(tc.expected, key)

			endTime, farmingCoinDenom, farmerAcc := types.ParseQueuedFarmingKey(key)
			s.Require().True(tc.endTime.Equal(endTime))
			s.Require().Equal(tc.farmingCoinDenom, farmingCoinDenom)
			s.Require().Equal(tc.farmerAcc, farmerAcc)
		})
	}
}

func (s *keysTestSuite) TestGetQueuedFarmingIndexKey() {
	testCases := []struct {
		farmerAcc        sdk.AccAddress
		farmingCoinDenom string
		endTime          time.Time
		expected         []byte
	}{
		{
			sdk.AccAddress(crypto.AddressHash([]byte("farmer1"))),
			sdk.DefaultBondDenom,
			utils.ParseTime("2022-01-01T00:00:00Z"),
			[]byte{0xe5, 0x14, 0xd3, 0x7a, 0x85, 0xec, 0x75, 0xf, 0x3, 0xaa, 0xe5, 0x36, 0xcf, 0x1b,
				0xb7, 0x59, 0xb7, 0xbc, 0xbd, 0x5c, 0xfe, 0x3d, 0x5, 0x73, 0x74, 0x61, 0x6b, 0x65, 0x32,
				0x30, 0x32, 0x32, 0x2d, 0x30, 0x31, 0x2d, 0x30, 0x31, 0x54, 0x30, 0x30, 0x3a, 0x30, 0x30,
				0x3a, 0x30, 0x30, 0x2e, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30},
		},
		{
			sdk.AccAddress(crypto.AddressHash([]byte("farmer5"))),
			sdk.DefaultBondDenom,
			utils.ParseTime("2022-07-01T00:10:00Z"),
			[]byte{0xe5, 0x14, 0xbc, 0xf3, 0x38, 0xd4, 0xac, 0xfa, 0x9b, 0x37, 0x5b, 0xf0, 0x54, 0x10,
				0xdf, 0x3b, 0x4a, 0x8e, 0xe3, 0x39, 0xb1, 0xc6, 0x5, 0x73, 0x74, 0x61, 0x6b, 0x65, 0x32,
				0x30, 0x32, 0x32, 0x2d, 0x30, 0x37, 0x2d, 0x30, 0x31, 0x54, 0x30, 0x30, 0x3a, 0x31, 0x30,
				0x3a, 0x30, 0x30, 0x2e, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30},
		},
		{
			sdk.AccAddress(crypto.AddressHash([]byte("farmer7"))),
			sdk.DefaultBondDenom,
			utils.ParseTime("2022-10-01T10:05:00Z"),
			[]byte{0xe5, 0x14, 0xd3, 0x6e, 0x93, 0x5b, 0x2b, 0x18, 0x63, 0x1, 0xe4, 0xaf, 0xb8, 0x6d,
				0x8c, 0xe5, 0x7b, 0xf4, 0x9e, 0x5f, 0xf9, 0x26, 0x5, 0x73, 0x74, 0x61, 0x6b, 0x65, 0x32,
				0x30, 0x32, 0x32, 0x2d, 0x31, 0x30, 0x2d, 0x30, 0x31, 0x54, 0x31, 0x30, 0x3a, 0x30, 0x35,
				0x3a, 0x30, 0x30, 0x2e, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30},
		},
	}

	for i, tc := range testCases {
		s.Run(fmt.Sprint(i), func() {
			key := types.GetQueuedFarmingIndexKey(tc.farmerAcc, tc.farmingCoinDenom, tc.endTime)
			s.Require().Equal(tc.expected, key)

			farmerAcc, farmingCoinDenom, endTime := types.ParseQueuedFarmingIndexKey(key)
			s.Require().Equal(tc.farmerAcc, farmerAcc)
			s.Require().Equal(tc.farmingCoinDenom, farmingCoinDenom)
			s.Require().True(tc.endTime.Equal(endTime))
		})
	}
}

func (s *keysTestSuite) TestGetQueuedFarmingsByFarmerPrefix() {
	farmer0 := sdk.AccAddress(crypto.AddressHash([]byte("farmer1")))
	farmer1 := sdk.AccAddress(crypto.AddressHash([]byte("farmer2")))
	farmer2 := sdk.AccAddress(crypto.AddressHash([]byte("farmer3")))
	farmer3 := sdk.AccAddress(crypto.AddressHash([]byte("farmer4")))
	s.Require().Equal([]byte{0xe5, 0x14, 0xd3, 0x7a, 0x85, 0xec, 0x75,
		0xf, 0x3, 0xaa, 0xe5, 0x36, 0xcf, 0x1b, 0xb7, 0x59, 0xb7, 0xbc,
		0xbd, 0x5c, 0xfe, 0x3d}, types.GetQueuedFarmingsByFarmerPrefix(farmer0))
	s.Require().Equal([]byte{0xe5, 0x14, 0x15, 0x1, 0x20, 0x25, 0x5a,
		0x5d, 0xe8, 0x6b, 0xa1, 0xed, 0xfb, 0x6f, 0x45, 0x48, 0xcb,
		0xfb, 0x6f, 0x28, 0x66, 0xf3}, types.GetQueuedFarmingsByFarmerPrefix(farmer1))
	s.Require().Equal([]byte{0xe5, 0x14, 0xdf, 0xb0, 0x6d, 0xbf, 0xc6,
		0x9a, 0xcd, 0xf5, 0x7b, 0xb, 0xe7, 0x69, 0x75, 0x50, 0x9e, 0x69,
		0x54, 0xa6, 0x1e, 0xe2}, types.GetQueuedFarmingsByFarmerPrefix(farmer2))
	s.Require().Equal([]byte{0xe5, 0x14, 0x98, 0x94, 0x3f, 0x57, 0x25,
		0xab, 0x66, 0xef, 0x46, 0x63, 0x4a, 0xfe, 0xeb, 0x8, 0xc0, 0x4a,
		0x53, 0x25, 0x2c, 0x9f}, types.GetQueuedFarmingsByFarmerPrefix(farmer3))
}

func (s *keysTestSuite) TestGetQueuedFarmingsByFarmerAndDenomPrefix() {
	testCases := []struct {
		farmerAcc        sdk.AccAddress
		farmingCoinDenom string
		expected         []byte
	}{
		{
			sdk.AccAddress(crypto.AddressHash([]byte("farmer1"))),
			sdk.DefaultBondDenom,
			[]byte{0xe5, 0x14, 0xd3, 0x7a, 0x85, 0xec, 0x75, 0xf, 0x3,
				0xaa, 0xe5, 0x36, 0xcf, 0x1b, 0xb7, 0x59, 0xb7, 0xbc,
				0xbd, 0x5c, 0xfe, 0x3d, 0x5, 0x73, 0x74, 0x61, 0x6b, 0x65},
		},
		{
			sdk.AccAddress(crypto.AddressHash([]byte("farmer3"))),
			sdk.DefaultBondDenom,
			[]byte{0xe5, 0x14, 0xdf, 0xb0, 0x6d, 0xbf, 0xc6, 0x9a, 0xcd,
				0xf5, 0x7b, 0xb, 0xe7, 0x69, 0x75, 0x50, 0x9e, 0x69, 0x54,
				0xa6, 0x1e, 0xe2, 0x5, 0x73, 0x74, 0x61, 0x6b, 0x65},
		},
		{
			sdk.AccAddress(crypto.AddressHash([]byte("farmer10"))),
			sdk.DefaultBondDenom,
			[]byte{0xe5, 0x14, 0x82, 0x64, 0xac, 0x54, 0xac, 0xa0, 0xc8,
				0xa5, 0xe5, 0xd2, 0xea, 0x5e, 0x5d, 0x1b, 0xe4, 0x32,
				0xd8, 0xb1, 0x88, 0xb0, 0x5, 0x73, 0x74, 0x61, 0x6b, 0x65},
		},
	}

	for i, tc := range testCases {
		s.Run(fmt.Sprint(i), func() {
			key := types.GetQueuedFarmingsByFarmerAndDenomPrefix(tc.farmerAcc, tc.farmingCoinDenom)
			s.Require().Equal(tc.expected, key)
		})
	}
}

func (s *keysTestSuite) TestGetQueuedFarmingEndBytes() {
	endTime0 := utils.ParseTime("2022-08-01T00:00:00Z")
	endTime1 := utils.ParseTime("2022-08-01T00:00:01Z")
	endTime2 := utils.ParseTime("2022-08-02T00:00:00Z")
	s.Require().Equal([]byte{0xe4, 0x1d, 0x32, 0x30, 0x32,
		0x32, 0x2d, 0x30, 0x38, 0x2d, 0x30, 0x31, 0x54, 0x30,
		0x30, 0x3a, 0x30, 0x30, 0x3a, 0x30, 0x30, 0x2e, 0x30,
		0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x31}, types.GetQueuedFarmingEndBytes(endTime0))
	s.Require().Equal([]byte{0xe4, 0x1d, 0x32, 0x30, 0x32,
		0x32, 0x2d, 0x30, 0x38, 0x2d, 0x30, 0x31, 0x54, 0x30,
		0x30, 0x3a, 0x30, 0x30, 0x3a, 0x30, 0x31, 0x2e, 0x30,
		0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x31}, types.GetQueuedFarmingEndBytes(endTime1))
	s.Require().Equal([]byte{0xe4, 0x1d, 0x32, 0x30, 0x32,
		0x32, 0x2d, 0x30, 0x38, 0x2d, 0x30, 0x32, 0x54, 0x30,
		0x30, 0x3a, 0x30, 0x30, 0x3a, 0x30, 0x30, 0x2e, 0x30,
		0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x31}, types.GetQueuedFarmingEndBytes(endTime2))
}

func (s *keysTestSuite) TestGetRewardsAuctionKey() {
	testCases := []struct {
		poolId    uint64
		auctionId uint64
		expected  []byte
	}{
		{
			1,
			1,
			[]byte{0xe7, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0,
				0x1, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x1},
		},
		{
			1,
			5,
			[]byte{0xe7, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0,
				0x1, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x5},
		},
		{
			5,
			5,
			[]byte{0xe7, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0,
				0x5, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x5},
		},
	}

	for i, tc := range testCases {
		s.Run(fmt.Sprint(i), func() {
			key := types.GetRewardsAuctionKey(tc.poolId, tc.auctionId)
			s.Require().Equal(tc.expected, key)
		})
	}
}

func (s *keysTestSuite) TestGetBidKey() {
	testCases := []struct {
		poolId   uint64
		bidder   sdk.AccAddress
		expected []byte
	}{
		{
			1,
			sdk.AccAddress(crypto.AddressHash([]byte("bidder1"))),
			[]byte{0xea, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x1, 0x14,
				0x20, 0x5c, 0xa, 0x82, 0xa, 0xf1, 0xed, 0x98, 0x39, 0x6a,
				0x35, 0xfe, 0xe3, 0x5d, 0x5, 0x2c, 0xd7, 0x96, 0x5a, 0x37},
		},
		{
			1,
			sdk.AccAddress(crypto.AddressHash([]byte("bidder3"))),
			[]byte{0xea, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x1, 0x14,
				0xe, 0x99, 0x7b, 0x9b, 0x5c, 0xef, 0x81, 0x2f, 0xc6, 0x3f,
				0xb6, 0x8b, 0x27, 0x42, 0x8a, 0xab, 0x7a, 0x58, 0xbc, 0x5e},
		},
		{
			5,
			sdk.AccAddress(crypto.AddressHash([]byte("bidder22"))),
			[]byte{0xea, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x5, 0x14,
				0x4c, 0xf1, 0xbd, 0x90, 0x1, 0x70, 0x78, 0xfb, 0xfc, 0x87,
				0x51, 0x9d, 0x40, 0x4, 0x39, 0x9f, 0x4d, 0xe3, 0xc9, 0x43},
		},
	}

	for i, tc := range testCases {
		s.Run(fmt.Sprint(i), func() {
			key := types.GetBidKey(tc.poolId, tc.bidder)
			s.Require().Equal(tc.expected, key)
		})
	}
}

func (s *keysTestSuite) TestGetBidByPoolIdPrefix() {
	s.Require().Equal([]byte{0xea, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0}, types.GetBidByPoolIdPrefix(0))
	s.Require().Equal([]byte{0xea, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x9}, types.GetBidByPoolIdPrefix(9))
	s.Require().Equal([]byte{0xea, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0xa}, types.GetBidByPoolIdPrefix(10))
}

func (s *keysTestSuite) TestGetWinningBidKey() {
	testCases := []struct {
		poolId    uint64
		auctionId uint64
		expected  []byte
	}{
		{
			1,
			1,
			[]byte{0xeb, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0,
				0x1, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x1},
		},
		{
			1,
			5,
			[]byte{0xeb, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0,
				0x1, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x5},
		},
		{
			5,
			5,
			[]byte{0xeb, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0,
				0x5, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x5},
		},
	}

	for i, tc := range testCases {
		s.Run(fmt.Sprint(i), func() {
			key := types.GetWinningBidKey(tc.poolId, tc.auctionId)
			s.Require().Equal(tc.expected, key)
		})
	}
}
