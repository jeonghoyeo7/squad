<!-- order: 7 -->

# Parameters

The `liquidfarming` module contains the following parameters:

| Key                    | Type             | Example                   |
| ---------------------- | ---------------- | ------------------------- |
| LiquidFarms            | []LiquidFarm     | []LiquidFarm{}            |
| RewardsAuctionDuration | string (time ns) | 43200000000000 (12 hours) |
| FeeCollector           | string           | "cosmos1..."              |

## LiquidFarms

`LiquidFarms` is a list of `LiquidFarm`, where a `LiquidFarm` is corresponding to a specific pool with `PoolId`.
A single `LiquidFarm` can exist for a given pool.

```go
type LiquidFarm struct {
	PoolId           uint64  // the pool id
	MinDepositAmount sdk.Int // the minimum deposit amount; it allows zero value
	MinBidAmount     sdk.Int // the minimum bid amount; it allows zero value
	FeeRate          sdk.Dec // the fee rate that deducts from auction winner's rewards;
}
```

## RewardsAuctionDuration

`RewardsAuctionDuration` is the duration that triggers the module to create new `RewardsAuction`.
If there is an ongoing `RewardsAuction`, then it finishes it and it creates next one.


## FeeCollector

A fee collector is a module account address that collects fees generated in the module.
