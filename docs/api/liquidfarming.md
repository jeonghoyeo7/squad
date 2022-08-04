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
  - [QueuedFarmings](#queuedfarmings)
  - [QueuedFarmingsByFarmer](#queuedfarmingsbyfarmer)
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

```

## LiquidFarm

Example Request

<!-- markdown-link-check-disable -->
```bash
http://localhost:1317/sqaud/liquidfarming/v1beta1/liquidfarms/1
```

Example Response

```json

```

## QueuedFarmings

Example Request

<!-- markdown-link-check-disable -->
```bash
http://localhost:1317/sqaud/liquidfarming/v1beta1/pools/1/queued_farmings
```

Example Response

```json
{
  "queued_farmings": [
    {
      "pool_id": "1",
      "amount": "500000000000"
    }
  ],
}
```

## QueuedFarmingsByFarmer

Example Request

<!-- markdown-link-check-disable -->
```bash
http://localhost:1317/sqaud/liquidfarming/v1beta1/pools/1/queued_farmings/farmer/{farmer_address}
```

Example Response

```json
{
  "queued_farmings": [
    {
      "pool_id": "1",
      "amount": "500000000000"
    }
  ],
}
```

## RewardsAuctions

Example Request

<!-- markdown-link-check-disable -->
```bash
http://localhost:1317/sqaud/liquidfarming/v1beta1/pools/1/rewards_auctions
```

Example Response

```json

```

## RewardsAuction

Example Request

<!-- markdown-link-check-disable -->
```bash
http://localhost:1317/sqaud/liquidfarming/v1beta1/pools/1/rewards_auctions/1
```

Example Response

```json

```


## Bids

Example Request

<!-- markdown-link-check-disable -->
```bash
http://localhost:1317/sqaud/liquidfarming/v1beta1/pools/1/bids
```

Example Response

```json

```