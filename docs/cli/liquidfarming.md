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
  - [LiquidUnstake](#liquidunstake)
- [Query](#query)
  - [Params](#params)
  - [LiquidValidators](#liquidvalidators)
  - [States](#states)
  - [VotingPower](#votingpower)

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
| amount       | amount of coin to liquid stake; it must be the bond denom |

Example

```bash
squad tx liquidfarming farm 1 100000000pool1 \
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
squad q bank balances cosmos1mzgucqnfr2l8cj5apvdpllhzt4zeuh2cshz5xu -o json | jq

# Query the voter's liquid staking voting power
squad q liquidstaking voting-power cosmos1mzgucqnfr2l8cj5apvdpllhzt4zeuh2cshz5xu -o json | jq
```

## LiquidUnstake

Unstake coin.

Usage

```bash
liquid-unstake [amount]
```

| **Argument**  |  **Description**                                      |
| :------------ | :---------------------------------------------------- |
| amount        | amount of coin to unstake; it must be the bToken denom|

Example

```bash
squad tx liquidstaking liquid-unstake 1000000000bstake \
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
squad q bank balances cosmos1mzgucqnfr2l8cj5apvdpllhzt4zeuh2cshz5xu -o json | jq

# Query the voter's liquid staking voting power
squad q liquidstaking voting-power cosmos1mzgucqnfr2l8cj5apvdpllhzt4zeuh2cshz5xu -o json | jq
```

# Query

## Params

Query the current liquidstaking parameters information.

Usage

```bash
params
```

Example

```bash
squad query liquidstaking params -o json | jq
```

## LiquidValidators

Query all liquid validators.

Usage

```bash
liquid-validators
```

Example

```bash
squad query liquidstaking liquid-validators -o json | jq
```
## States

Query net amount state.

Usage

```bash
states
```

Example

```bash
squad query liquidstaking states -o json | jq
```

## VotingPower

Query the voter’s staking and liquid staking voting power. 

Usage

```bash
voting-power [voter]
```

| **Argument** |  **Description**      |
| :----------- | :-------------------- |
| voter        | voter account address |

Example

```bash
squad query liquidstaking voting-power cosmos1mzgucqnfr2l8cj5apvdpllhzt4zeuh2cshz5xu -o json | jq
```