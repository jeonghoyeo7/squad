<!-- order: 6 -->

# Parameters

The `liquidfarming` module contains the following parameters:

| Key                        | Type         | Example                                        |
| -------------------------- | ------------ | ---------------------------------------------- |
| LiquidFarmCreationFee      | sdk.Coins    | [{"denom":"stake","amount":"100000000"}]       |
| LiquidFarms                | []LiquidFarm | TBD                                            |


## LiquidFarmCreationFee

`LiquidFarmCreationFee` is a fee required to create new liquid farm.

## LiquidFarms

`LiquidFarms` is a list of `LiquidFarm`, where a `LiquidFarm` is corresponding to a specific pool with `PoolId`. 
A single `LiquidFarm` can exist for a given pool. 


```go
type LiquidFarm struct {
	PoolId               uint64
	MinimumFarmAmount sdk.Int
	MinimumBidAmount     sdk.Int
}
```
