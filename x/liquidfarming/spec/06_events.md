<!-- order: 6 -->

# Events

The `liquidfarming` module emits the following events:

## Handlers

### MsgFarm

| Type       | Attribute Key      | Attribute Value        |
| ---------- | ------------------ | ---------------------- |
| farm       | pool_id            | {poolId}               |
| farm       | farmer             | {farmer}               |
| farm       | farming_coin       | {farmingCoin}          |
| farm       | minted_coin        | {mintedCoin}           |
| message    | module             | liquidfarming          |
| message    | action             | farm                   |
| message    | farmer             | {farmerAddress}        |

### MsgUnfarm

| Type       | Attribute Key      | Attribute Value        |
| ---------- | ------------------ | ---------------------- |
| unfarm     | pool_id            | {poolId}               |
| unfarm     | farmer             | {farmer}               |
| unfarm     | burning_coin       | {unfarmingCoin}        |
| unfarm     | unfarmed_coin      | {unfarmCoin}           |
| message    | module             | liquidfarming          |
| message    | action             | unfarm                 |
| message    | farmer             | {farmerAddress}        |

### MsgUnfarmAndWithdraw

| Type       | Attribute Key      | Attribute Value        |
| ---------- | ------------------ | ---------------------- |
| unfarm     | pool_id            | {poolId}               |
| unfarm     | farmer             | {farmer}               |
| unfarm     | burning_coin       | {unfarmingCoin}        |
| unfarm     | unfarmed_coin      | {unfarmCoin}           |
| message    | module             | liquidfarming          |
| message    | action             | unfarmandwithdraw      |
| message    | farmer             | {farmerAddress}        |

### MsgPlaceBid

| Type       | Attribute Key      | Attribute Value        |
| ---------- | ------------------ | ---------------------- |
| place_bid  | auction_id         | {auctionId}            |
| place_bid  | bidder             | {bidder}               |
| place_bid  | bidding_coin       | {biddingCoin}          |
| message    | module             | liquidfarming          |
| message    | action             | placebid               |
| message    | bidder             | {bidderAddress}        |

### MsgRefundBid 

| Type       | Attribute Key      | Attribute Value        |
| ---------- | ------------------ | ---------------------- |
| refund_bid | pool_id            | {auctionId}            |
| refund_bid | bidder             | {bidder}               |
| refund_bid | refund_coin        | {bid.amount}           |
| message    | module             | liquidfarming          |
| message    | action             | refundbid              |
| message    | bidder             | {bidderAddress}        |