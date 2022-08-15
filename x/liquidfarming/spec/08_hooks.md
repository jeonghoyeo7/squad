<!-- order: 8 -->

# Hooks

The `Liquidfarming` module uses the following hooks registered in `farming` module.

## AfterAllocateRewards Hook

```go
AfterAllocateRewards(ctx sdk.Context)
```

When `AfterAllocateRewards` hook is delivered, the following operations are performed.
- If the auction currently going on exists, the current auction becomes finished. And, 
  - the winner is chosen,
  - the rewards is harvested and sent to the winner,
  - the pool coins from the winner in the paying reserve address is sent to the module account,
    - the module stakes the pool coins from the auction, the amount of these pool coins is saved to `CompoundingRewards`
  - the pool coins from the others not winner in the paying reserve address is refunded to each bidder’s account.
- A new auction is created.
  