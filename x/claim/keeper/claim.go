package keeper

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"

	"github.com/cosmosquad-labs/squad/x/claim/types"
)

func (k Keeper) Claim(ctx sdk.Context, msg *types.MsgClaim) (types.ClaimRecord, error) {
	airdrop, found := k.GetAirdrop(ctx, msg.AirdropId)
	if !found {
		return types.ClaimRecord{}, sdkerrors.Wrap(sdkerrors.ErrNotFound, "airdrop not found")
	}

	if !airdrop.EndTime.After(ctx.BlockTime()) {
		return types.ClaimRecord{}, types.ErrTerminatedAirdrop
	}

	record, found := k.GetClaimRecordByRecipient(ctx, airdrop.Id, msg.GetRecipient())
	if !found {
		return types.ClaimRecord{}, sdkerrors.Wrap(sdkerrors.ErrNotFound, "claim record not found")
	}

	for _, c := range record.ClaimedConditions {
		if c == msg.ConditionType {
			return types.ClaimRecord{}, types.ErrAlreadyClaimed
		}
	}

	// Validate whether or not the recipient has executed the condition
	if err := k.ValidateCondition(ctx, record.GetRecipient(), msg.ConditionType); err != nil {
		return types.ClaimRecord{}, err
	}

	claimableCoins := record.GetClaimableCoinsForCondition(airdrop.Conditions)

	if err := k.bankKeeper.SendCoins(ctx, airdrop.GetSourceAddress(), record.GetRecipient(), claimableCoins); err != nil {
		return types.ClaimRecord{}, sdkerrors.Wrap(err, "failed to transfer coins to the recipient")
	}

	record.ClaimableCoins = record.ClaimableCoins.Sub(claimableCoins)
	record.ClaimedConditions = append(record.ClaimedConditions, msg.ConditionType)
	k.SetClaimRecord(ctx, record)

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeClaim,
			sdk.NewAttribute(types.AttributeKeyAirdropId, fmt.Sprint(record.AirdropId)),
			sdk.NewAttribute(types.AttributeKeyRecipient, record.Recipient),
			sdk.NewAttribute(types.AttributeKeyInitialClaimableCoins, record.InitialClaimableCoins.String()),
			sdk.NewAttribute(types.AttributeKeyClaimableCoins, record.ClaimableCoins.String()),
			sdk.NewAttribute(types.AttributeKeyConditionType, msg.ConditionType.String()),
		),
	})

	return record, nil
}

// ValidateCondition validates if the recipient has executed the condition.
func (k Keeper) ValidateCondition(ctx sdk.Context, recipient sdk.AccAddress, ct types.ConditionType) error {
	ok := false

	switch ct {
	case types.ConditionTypeDeposit:
		if len(k.liquidityKeeper.GetDepositRequestsByDepositor(ctx, recipient)) != 0 {
			ok = true
		}

	case types.ConditionTypeSwap:
		if len(k.liquidityKeeper.GetOrdersByOrderer(ctx, recipient)) != 0 {
			ok = true
		}

	case types.ConditionTypeLiquidStake:
		params := k.liquidStakingKeeper.GetParams(ctx)
		spendable := k.bankKeeper.SpendableCoins(ctx, recipient)
		bTokenBalance := spendable.AmountOf(params.LiquidBondDenom)
		if !bTokenBalance.IsZero() {
			ok = true
		}

	case types.ConditionTypeVote:
		k.govKeeper.IterateAllVotes(ctx, func(vote govtypes.Vote) (stop bool) {
			if vote.Voter == recipient.String() {
				ok = true
				return true
			}
			return false
		})
	}

	if !ok {
		return types.ErrConditionRequired
	}

	return nil
}

// TerminateAirdrop terminates the airdrop and transfer the remaining coins to the community pool.
func (k Keeper) TerminateAirdrop(ctx sdk.Context, airdrop types.Airdrop) error {
	amt := k.bankKeeper.SpendableCoins(ctx, airdrop.GetSourceAddress())
	if !amt.IsZero() {
		if err := k.distrKeeper.FundCommunityPool(ctx, amt, airdrop.GetSourceAddress()); err != nil {
			return sdkerrors.Wrap(err, "failed to transfer the remaining coins to the community pool")
		}
	}
	return nil
}
