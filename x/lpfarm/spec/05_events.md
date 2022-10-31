<!-- order: 5 -->

# Events

## Handlers

### MsgCreatePrivatePlan

| Type                                        | Attribute Key        | Attribute Value                             |
|---------------------------------------------|----------------------|---------------------------------------------|
| message                                     | action               | /squad.lpfarm.v1beta1.Msg/CreatePrivatePlan |
| squad.lpfarm.v1beta1.EventCreatePrivatePlan | creator              | {planCreatorAddress}                        |
| squad.lpfarm.v1beta1.EventCreatePrivatePlan | plan_id              | {planId}                                    |
| squad.lpfarm.v1beta1.EventCreatePrivatePlan | farming_pool_address | {farmingPoolAddress}                        |

### MsgFarm

| Type                           | Attribute Key     | Attribute Value                |
|--------------------------------|-------------------|--------------------------------|
| message                        | action            | /squad.lpfarm.v1beta1.Msg/Farm |
| squad.lpfarm.v1beta1.EventFarm | farmer            | {farmerAddress}                |
| squad.lpfarm.v1beta1.EventFarm | coin              | {coin}                         |
| squad.lpfarm.v1beta1.EventFarm | withdrawn_rewards | {withdrawnRewards}             |

### MsgUnfarm

| Type                             | Attribute Key     | Attribute Value                  |
|----------------------------------|-------------------|----------------------------------|
| message                          | action            | /squad.lpfarm.v1beta1.Msg/Unfarm |
| squad.lpfarm.v1beta1.EventUnfarm | farmer            | {farmerAddress}                  |
| squad.lpfarm.v1beta1.EventUnfarm | coin              | {coin}                           |
| squad.lpfarm.v1beta1.EventUnfarm | withdrawn_rewards | {withdrawnRewards}               |

### MsgHarvest

| Type                              | Attribute Key     | Attribute Value                   |
|-----------------------------------|-------------------|-----------------------------------|
| message                           | action            | /squad.lpfarm.v1beta1.Msg/Harvest |
| squad.lpfarm.v1beta1.EventHarvest | farmer            | {farmerAddress}                   |
| squad.lpfarm.v1beta1.EventHarvest | denom             | {farmingAssetDenom}               |
| squad.lpfarm.v1beta1.EventHarvest | withdrawn_rewards | {withdrawnRewards}                |
