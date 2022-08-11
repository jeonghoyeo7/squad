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

- [Liquidfarming Module](#liquidfarming-module)
  - [Synopsis](#synopsis)
  - [gRPC-gateway REST Routes](#grpc-gateway-rest-routes)
  - [Params](#params)
  - [LiquidFarms](#liquidfarms)
  - [LiquidFarm](#liquidfarm)
  - [RewardsAuctions](#rewardsauctions)
  - [RewardsAuction](#rewardsauction)
  - [Bids](#bids)

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
    "liquid_farms": [
      {
        "pool_id": "1",
        "minimum_farm_amount": "1",
        "minimum_bid_amount": "1"
      },
      {
        "pool_id": "2",
        "minimum_farm_amount": "1",
        "minimum_bid_amount": "1"
      }
    ],
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
      "minimum_farm_amount": "1",
      "minimum_bid_amount": "1",
      "staked_coin": {
        "denom": "pool1",
        "amount": "0"
      },
      "queued_coin": {
        "denom": "pool1",
        "amount": "500000000000"
      }
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
    "staked_coin": {
      "denom": "pool1",
      "amount": "0"
    },
    "queued_coin": {
      "denom": "pool1",
      "amount": "500000000000"
    }
  }
}
```

## RewardsAuctions

Example Request

<!-- markdown-link-check-disable -->
```bash
http://localhost:1317/squad/liquidfarming/v1beta1/pools/1/rewards_auctions
```

Example Response

```json
{
  "reward_auctions": [
    {
      "id": "2",
      "pool_id": "1",
      "bidding_coin_denom": "pool1",
      "paying_reserve_address": "cosmos1h72q3pkvsz537kj08hyv20tun3apampxhpgad97t3ls47nukgtxqeq6eu2",
      "start_time": "2022-08-05T08:56:55.820787Z",
      "end_time": "2022-08-06T08:56:55.820787Z",
      "status": "AUCTION_STATUS_STARTED",
      "winner": "",
      "rewards": [
        
      ]
    },
    {
      "id": "1",
      "pool_id": "1",
      "bidding_coin_denom": "pool1",
      "paying_reserve_address": "cosmos1h72q3pkvsz537kj08hyv20tun3apampxhpgad97t3ls47nukgtxqeq6eu2",
      "start_time": "2022-08-05T08:56:22.237454Z",
      "end_time": "2022-08-06T08:56:22.237454Z",
      "status": "AUCTION_STATUS_FINISHED",
      "winner": "",
      "rewards": [
        
      ]
    }
  ],
}
```

## RewardsAuction

Example Request

<!-- markdown-link-check-disable -->
```bash
http://localhost:1317/squad/liquidfarming/v1beta1/pools/1/rewards_auctions/1
```

Example Response

```json
{
  "reward_auction": {
    "id": "1",
    "pool_id": "1",
    "bidding_coin_denom": "pool1",
    "paying_reserve_address": "cosmos1h72q3pkvsz537kj08hyv20tun3apampxhpgad97t3ls47nukgtxqeq6eu2",
    "start_time": "2022-08-05T08:56:22.237454Z",
    "end_time": "2022-08-06T08:56:22.237454Z",
    "status": "AUCTION_STATUS_FINISHED",
    "winner": "",
    "rewards": [
      
    ]
  }
}
```


## Bids

Example Request

<!-- markdown-link-check-disable -->
```bash
http://localhost:1317/squad/liquidfarming/v1beta1/pools/1/bids
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
        "amount": "10000000"
      }
    }
  ],
}
```