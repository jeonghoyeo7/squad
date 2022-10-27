---
Title: Liquidfarming
Description: A high-level overview of what gRPC-gateway REST routes are supported in the liquidfarming module.
---

# Liquidfarming Module

## Synopsis

This document provides a high-level overview of what gRPC-gateway REST routes are supported in the `liquidfarming` module.

## gRPC-gateway REST Routes

<!-- markdown-link-check-disable -->

++https://github.com/cosmosquad-labs/squad/blob/main/proto/squad/liquidfarming/v1beta1/query.proto


- [Params](#Params)
- [LiquidFarms](#Liquidfarms)
- [LiquidFarm](#Liquidfarm)
- [RewardsAuctions](#Rewardsauctions)
- [RewardsAuction](#Rewardsauction)
- [Bids](#Bids)
- [Rewards](#Rewards)
- [ExchangeRate](#ExchangeRate)

## Params

Example Request

<!-- markdown-link-check-disable -->

```bash
http://localhost:1317/squad/liquidfarming/v1beta1/params
```

Example Response

```json
{
  "params": {
    "fee_collector": "cosmos1lsvtflq2gau8ha7zvlethfy85qus59eserphyhc3tumua7upx6eq59trlz",
    "rewards_auction_duration": "120s",
    "liquid_farms": [
      {
        "pool_id": "1",
        "min_farm_amount": "1",
        "min_bid_amount": "1",
        "fee_rate": "0.000000000000000000"
      }
    ]
  }
}
```

## LiquidFarms

Example Request

<!-- markdown-link-check-disable -->

```bash
http://localhost:1317/squad/liquidfarming/v1beta1/liquidfarms
```

Example Response

```json
{
  "liquid_farms": [
    {
      "pool_id": "1",
      "liquid_farm_reserve_address": "cosmos1zyyf855slxure4c8dr06p00qjnkem95d2lgv8wgvry2rt437x6tsaf9tcf",
      "lf_coin_denom": "lf1",
      "min_farm_amount": "1",
      "min_bid_amount": "1",
      "total_farming_amount": "500000000000"
    }
  ]
}
```

## LiquidFarm

Example Request

<!-- markdown-link-check-disable -->

```bash
http://localhost:1317/squad/liquidfarming/v1beta1/liquidfarms/1
```

Example Response

```json
{
  "liquid_farm": {
    "pool_id": "1",
    "liquid_farm_reserve_address": "cosmos1zyyf855slxure4c8dr06p00qjnkem95d2lgv8wgvry2rt437x6tsaf9tcf",
    "lf_coin_denom": "lf1",
    "minimum_farm_amount": "1",
    "minimum_bid_amount": "1",
    "total_farming_amount": "500000000000"
  }
}
```

## RewardsAuctions

Example Request

<!-- markdown-link-check-disable -->

```bash
http://localhost:1317/squad/liquidfarming/v1beta1/liquidfarms/1/rewards_auctions
http://localhost:1317/squad/liquidfarming/v1beta1/liquidfarms/1/rewards_auctions?status=AUCTION_STATUS_STARTED
http://localhost:1317/squad/liquidfarming/v1beta1/liquidfarms/1/rewards_auctions?status=AUCTION_STATUS_FINISHED
http://localhost:1317/squad/liquidfarming/v1beta1/liquidfarms/1/rewards_auctions?status=AUCTION_STATUS_SKIPPED
```

Example Response

```json
{
  "reward_auctions": [
    {
      "id": "1",
      "pool_id": "1",
      "bidding_coin_denom": "pool1",
      "paying_reserve_address": "cosmos1h72q3pkvsz537kj08hyv20tun3apampxhpgad97t3ls47nukgtxqeq6eu2",
      "start_time": "2022-10-27T00:00:00Z",
      "end_time": "2022-10-27T00:02:00Z",
      "status": "AUCTION_STATUS_STARTED",
      "winner": "",
      "winning_amount": {
        "denom": "",
        "amount": "0"
      },
      "rewards": [
      ]
    }
  ]
}
```

## RewardsAuction

Example Request

<!-- markdown-link-check-disable -->

```bash
http://localhost:1317/squad/liquidfarming/v1beta1/liquidfarms/1/rewards_auctions/1
```

Example Response

```json
{
  "reward_auction": {
    "id": "1",
    "pool_id": "1",
    "bidding_coin_denom": "pool1",
    "paying_reserve_address": "cosmos1h72q3pkvsz537kj08hyv20tun3apampxhpgad97t3ls47nukgtxqeq6eu2",
    "start_time": "2022-10-27T00:00:00Z",
    "end_time": "2022-10-27T00:02:00Z",
    "status": "AUCTION_STATUS_SKIPPED",
    "winner": "",
    "winning_amount": {
      "denom": "",
      "amount": "0"
    },
    "rewards": [
      {
        "denom": "stake",
        "amount": "25369"
      }
    ]
  }
}
```

## Bids

Example Request

<!-- markdown-link-check-disable -->

```bash
http://localhost:1317/squad/liquidfarming/v1beta1/liquidfarms/1/bids
```

Example Response

```json
{
  "bids": [
    {
      "pool_id": "1",
      "bidder": "cosmos1zaavvzxez0elundtn32qnk9lkm8kmcszzsv80v",
      "amount": {
        "denom": "pool1",
        "amount": "1000000000"
      }
    }
  ]
}
```

## Rewards

Example Request

<!-- markdown-link-check-disable -->

```bash
http://localhost:1317/squad/liquidfarming/v1beta1/liquidfarms/1/rewards
```

Example Response

```json
{
  "rewards": [
    {
      "denom": "stake",
      "amount": "1903"
    }
  ]
}
```

## ExchangeRate

Example Request

<!-- markdown-link-check-disable -->

```bash
http://localhost:1317/squad/liquidfarming/v1beta1/liquidfarms/1/exchange_rate
```

Example Response

```json
{
  "exchange_rate": {
    "mint_rate": "1.000000000000000000",
    "burn_rate": "1.000000000000000000"
  }
}
```
