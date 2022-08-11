---
Title: LiquidFarming
Description: A high-level overview of how the command-line interfaces (CLI) works for the liquidfarming module.
---

# LiquidFarming Module

## Synopsis

This document provides a high-level overview of how the command line (CLI) interface works for the `liquidfarming` module. 
To set up a local testing environment, it requires the latest [Ignite CLI](https://docs.ignite.com/). 
If you don't have Ignite CLI set up in your local machine, see [this guide](https://docs.ignite.com/guide/install.html) to install it. 
Run this command under the project root directory `$ ignite chain serve -c config-test.yml`.

Note that [jq](https://stedolan.github.io/jq/) is recommended to be installed as it is used to process JSON throughout the document.

- [LiquidFarming Module](#liquidfarming-module)
  - [Synopsis](#synopsis)
- [Transaction](#transaction)
  - [Farm](#farm)
  - [Unfarm](#unfarm)
- [Query](#query)
  - [Params](#params)
  - [Liquidfarms](#liquidfarms)
  - [Liquidfarm](#liquidfarm)
  - [Rewards-Auctions](#rewards-auctions)
  - [Rewards-Auction](#rewards-auction)
  - [Bids](#bids)

# Transaction

## Farm

Farm pool coin for liquid farming. 
It is important to note that the farmer receives corresponding LFCoin after 1 epoch is passed. 
It is because their pool coin is reserved in liquid farm reserve account and it stakes the amount in the farming module for them. 
When an epoch is passed, the module mints the LFCoin and send them to the farmer.

Usage

```bash
farm [pool-id] [amount]
```

| **Argument** |  **Description**                                          |
| :----------- | :-------------------------------------------------------- |
| pool-id      | target pool id of the liquid farm                         |
| amount       | amount of pool coin of the target pool to liquid farm     |

Example

```bash
squad tx liquidfarming farm 1 500000000000pool1 \
--chain-id localnet \
--from bob \
--keyring-backend test \
--gas 1000000 \
--broadcast-mode block \
--yes \
--output json | jq

#
# Tips
#
# Query account balances
# Notice the newly minted bToken
squad q liquidfarming queued-farmings 1 -o json | jq
```

## Unfarm

Unfarm liquid farming coin to receive the corresponding pool coin.

Usage

```bash
unfarm [pool-id] [amount]
```

| **Argument**  |  **Description**                                      |
| :------------ | :---------------------------------------------------- |
| pool-id       | target pool id of the liquid unfarm                   |
| amount        | amount of lf coin to liquid unfarm                    |

Example

```bash
squad tx liquidfarming unfarm 1 300000000000lf1 \
--chain-id localnet \
--from alice \
--keyring-backend test \
--broadcast-mode block \
--yes \
--output json | jq

#
# Tips
#
# Query account balances
# Notice the newly minted bToken
squad q bank balances cosmos1mzgucqnfr2l8cj5apvdpllhzt4zeuh2cshz5xu -o json | jq
```

# Query

## Params

Query the current liquidfarming parameters information.

Usage

```bash
params
```

Example

```bash
squad query liquidfarming params -o json | jq
```

## Liquidfarms

 Query for all liquidfarms.

Usage

```bash
liquidfarms
```

Example

```bash
squad query liquidfarming liquidfarms -o json | jq
```

## Liquidfarm

Query the specific liquidfarm with pool id.

Usage

```bash
liquidfarm [pool-id]
```

| **Argument**  |  **Description**                                      |
| :------------ | :---------------------------------------------------- |
| pool-id       | target pool id of the liquidfarm                      |

Example

```bash
squad query liquidfarming liquidfarm 1 -o json | jq
```

## Rewards-Auctions

Query all rewards auctions for the liquidfarm

Usage

```bash
rewards-auctions
```

Example

```bash
squad query liquidfarming rewards-auctions -o json | jq
```

## Rewards-Auction 

Query the specific reward auction

Usage

```bash
rewards-auction [pool-id] [auction-id]
```

| **Argument**  |  **Description**                                      |
| :------------ | :---------------------------------------------------- |
| pool-id       | target pool id of the liquidfarm                      |
| auction-id    | target auction id of the liquidfarm with the pool id  |

Example

```bash
squad query liquidfarming rewards-auction 1 1 -o json | jq
```

## Bids 

Query all bids for the rewards auction

Usage

```bash
bids [pool-id]
```

| **Argument**  |  **Description**                                      |
| :------------ | :---------------------------------------------------- |
| pool-id       | target pool id of the liquidfarm                      |

Example

```bash
squad query liquidfarming bids 1 -o json | jq
```
