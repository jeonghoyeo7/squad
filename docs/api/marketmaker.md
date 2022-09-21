---
Title: Marketmaker
Description: A high-level overview of what gRPC-gateway REST routes are supported in the marketmaker module.
---

# Marketmaker Module

## Synopsis

This document provides a high-level overview of what gRPC-gateway REST routes are supported in the `marketmaker` module.

## gRPC-gateway REST Routes

<!-- markdown-link-check-disable -->

++https://github.com/cosmosquad-labs/squad/blob/main/proto/squad/marketmaker/v1beta1/query.proto

- [Params](#Params)
- [MarketMakers](#MarketMakers)
- [Incentive](#Incentive)

### Params

Query the values set as marketmaker parameters:

Example Request

<!-- markdown-link-check-disable -->

```bash
http://localhost:1317/squad/marketmaker/v1beta1/params
```

```json
{
  "params": {
    "incentive_budget_address": "cosmos1ddn66jv0sjpmck0ptegmhmqtn35qsg2vxyk2hn9sqf4qxtzqz3sqanrtcm",
    "deposit_amount": [
      {
        "denom": "stake",
        "amount": "1000000000"
      }
    ],
    "common": {
      "min_open_ratio": "0.500000000000000000",
      "min_open_depth_ratio": "0.100000000000000000",
      "max_downtime": 20,
      "max_total_downtime": 100,
      "min_hours": 16,
      "min_days": 22
    },
    "incentive_pairs": [
      {
        "pair_id": "1",
        "update_time": "2022-09-10T00:00:00Z",
        "incentive_weight": "0.000000000000000000",
        "max_spread": "0.000000000000000000",
        "min_width": "0.000000000000000000",
        "min_depth": "0"
      },
      {
        "pair_id": "2",
        "update_time": "2022-09-10T00:00:00Z",
        "incentive_weight": "0.000000000000000000",
        "max_spread": "0.000000000000000000",
        "min_width": "0.000000000000000000",
        "min_depth": "0"
      }
    ]
  }
}
```

### MarketMakers

Query all the market makers in the network:

Example Request

<!-- markdown-link-check-disable -->

```bash
http://localhost:1317/squad/marketmaker/v1beta1/marketmakers
```

```json
{
  "marketmakers": [
    {
      "address": "cosmos1zaavvzxez0elundtn32qnk9lkm8kmcszzsv80v",
      "pair_id": "1",
      "eligible": true
    }
  ],
  "pagination": {
    "next_key": null,
    "total": "1"
  }
}
```

### Incentive

Query claimable incentive of a market maker

Example Request

<!-- markdown-link-check-disable -->

```bash
http://localhost:1317/squad/marketmaker/v1beta1/incentive/cosmos1zaavvzxez0elundtn32qnk9lkm8kmcszzsv80v
```

```json
{
  "incentive": {
    "address": "cosmos1zaavvzxez0elundtn32qnk9lkm8kmcszzsv80v",
    "claimable": [
      {
        "denom": "stake",
        "amount": "100000000"
      }
    ]
  }
}
```
