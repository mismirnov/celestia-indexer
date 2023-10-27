// SPDX-FileCopyrightText: 2023 PK Lab AG <contact@pklab.io>
// SPDX-License-Identifier: MIT

package types

import (
	"database/sql"
	"database/sql/driver"

	"github.com/pkg/errors"
)

type MsgTypeBits struct {
	Bits
}

const (
	MsgTypeBitsUnknown uint64 = 1 << iota

	MsgTypeBitsSetWithdrawAddress
	MsgTypeBitsWithdrawDelegatorReward
	MsgTypeBitsWithdrawValidatorCommission
	MsgTypeBitsFundCommunityPool

	MsgTypeBitsCreateValidator
	MsgTypeBitsEditValidator
	MsgTypeBitsDelegate
	MsgTypeBitsBeginRedelegate
	MsgTypeBitsUndelegate
	MsgTypeBitsCancelUnbondingDelegation

	MsgTypeBitsUnjail

	MsgTypeBitsSend
	MsgTypeBitsMultiSend

	MsgTypeBitsCreateVestingAccount
	MsgTypeBitsCreatePermanentLockedAccount
	MsgTypeBitsCreatePeriodicVestingAccount

	MsgTypeBitsPayForBlobs

	MsgTypeBitsGrant
	MsgTypeBitsExec
	MsgTypeBitsRevoke

	MsgTypeBitsGrantAllowance
	MsgTypeBitsRevokeAllowance

	MsgTypeBitsRegisterEVMAddress

	MsgTypeBitsSubmitProposal
	MsgTypeBitsExecLegacyContent
	MsgTypeBitsVote
	MsgTypeBitsVoteWeighted
	MsgTypeBitsDeposit

	MsgTypeBitsIBCTransfer

	MsgTypeBitsVerifyInvariant

	MsgTypeBitsSubmitEvidence

	MsgTypeBitsSendNFT

	MsgTypeBitsCreateGroup
	MsgTypeBitsUpdateGroupMembers
	MsgTypeBitsUpdateGroupAdmin
	MsgTypeBitsUpdateGroupMetadata
	MsgTypeBitsCreateGroupPolicy
	MsgTypeBitsUpdateGroupPolicyAdmin
	MsgTypeBitsCreateGroupWithPolicy
	MsgTypeBitsUpdateGroupPolicyDecisionPolicy
	MsgTypeBitsUpdateGroupPolicyMetadata
	MsgTypeBitsSubmitProposalGroup
	MsgTypeBitsWithdrawProposal
	MsgTypeBitsVoteGroup
	MsgTypeBitsExecGroup
	MsgTypeBitsLeaveGroup
)

func NewMsgTypeBitMask(values ...MsgType) MsgTypeBits {
	var mask MsgTypeBits
	for i := range values {
		mask.SetBit(values[i])
	}
	return mask
}

func (mask *MsgTypeBits) SetBit(value MsgType) {
	switch value {
	case MsgUnknown:
		mask.Set(Bits(MsgTypeBitsUnknown))

	case MsgSetWithdrawAddress:
		mask.Set(Bits(MsgTypeBitsSetWithdrawAddress))
	case MsgWithdrawDelegatorReward:
		mask.Set(Bits(MsgTypeBitsWithdrawDelegatorReward))
	case MsgWithdrawValidatorCommission:
		mask.Set(Bits(MsgTypeBitsWithdrawValidatorCommission))
	case MsgFundCommunityPool:
		mask.Set(Bits(MsgTypeBitsFundCommunityPool))

	case MsgCreateValidator:
		mask.Set(Bits(MsgTypeBitsCreateValidator))
	case MsgEditValidator:
		mask.Set(Bits(MsgTypeBitsEditValidator))
	case MsgDelegate:
		mask.Set(Bits(MsgTypeBitsDelegate))
	case MsgBeginRedelegate:
		mask.Set(Bits(MsgTypeBitsBeginRedelegate))
	case MsgUndelegate:
		mask.Set(Bits(MsgTypeBitsUndelegate))
	case MsgCancelUnbondingDelegation:
		mask.Set(Bits(MsgTypeBitsCancelUnbondingDelegation))

	case MsgUnjail:
		mask.Set(Bits(MsgTypeBitsUnjail))

	case MsgSend:
		mask.Set(Bits(MsgTypeBitsSend))
	case MsgMultiSend:
		mask.Set(Bits(MsgTypeBitsMultiSend))

	case MsgCreateVestingAccount:
		mask.Set(Bits(MsgTypeBitsCreateVestingAccount))
	case MsgCreatePermanentLockedAccount:
		mask.Set(Bits(MsgTypeBitsCreatePermanentLockedAccount))
	case MsgCreatePeriodicVestingAccount:
		mask.Set(Bits(MsgTypeBitsCreatePeriodicVestingAccount))

	case MsgPayForBlobs:
		mask.Set(Bits(MsgTypeBitsPayForBlobs))

	case MsgGrant:
		mask.Set(Bits(MsgTypeBitsGrant))
	case MsgExec:
		mask.Set(Bits(MsgTypeBitsExec))
	case MsgRevoke:
		mask.Set(Bits(MsgTypeBitsRevoke))

	case MsgGrantAllowance:
		mask.Set(Bits(MsgTypeBitsGrantAllowance))
	case MsgRevokeAllowance:
		mask.Set(Bits(MsgTypeBitsRevokeAllowance))

	case MsgRegisterEVMAddress:
		mask.Set(Bits(MsgTypeBitsRegisterEVMAddress))

	case MsgSubmitProposal:
		mask.Set(Bits(MsgTypeBitsSubmitProposal))
	case MsgExecLegacyContent:
		mask.Set(Bits(MsgTypeBitsExecLegacyContent))
	case MsgVote:
		mask.Set(Bits(MsgTypeBitsVote))
	case MsgVoteWeighted:
		mask.Set(Bits(MsgTypeBitsVoteWeighted))
	case MsgDeposit:
		mask.Set(Bits(MsgTypeBitsDeposit))
	case IBCTransfer:
		mask.Set(Bits(MsgTypeBitsIBCTransfer))
	case MsgVerifyInvariant:
		mask.Set(Bits(MsgTypeBitsVerifyInvariant))
	case MsgSubmitEvidence:
		mask.Set(Bits(MsgTypeBitsSubmitEvidence))
	case MsgSendNFT:
		mask.Set(Bits(MsgTypeBitsSendNFT))

	case MsgCreateGroup:
		mask.Set(Bits(MsgTypeBitsCreateGroup))
	case MsgUpdateGroupMembers:
		mask.Set(Bits(MsgTypeBitsUpdateGroupMembers))
	case MsgUpdateGroupAdmin:
		mask.Set(Bits(MsgTypeBitsUpdateGroupAdmin))
	case MsgUpdateGroupMetadata:
		mask.Set(Bits(MsgTypeBitsUpdateGroupMetadata))
	case MsgCreateGroupPolicy:
		mask.Set(Bits(MsgTypeBitsCreateGroupPolicy))
	case MsgUpdateGroupPolicyAdmin:
		mask.Set(Bits(MsgTypeBitsUpdateGroupPolicyAdmin))
	case MsgCreateGroupWithPolicy:
		mask.Set(Bits(MsgTypeBitsCreateGroupWithPolicy))
	case MsgUpdateGroupPolicyDecisionPolicy:
		mask.Set(Bits(MsgTypeBitsUpdateGroupPolicyDecisionPolicy))
	case MsgUpdateGroupPolicyMetadata:
		mask.Set(Bits(MsgTypeBitsUpdateGroupPolicyMetadata))
	case MsgSubmitProposalGroup:
		mask.Set(Bits(MsgTypeBitsSubmitProposalGroup))
	case MsgWithdrawProposal:
		mask.Set(Bits(MsgTypeBitsWithdrawProposal))
	case MsgVoteGroup:
		mask.Set(Bits(MsgTypeBitsVoteGroup))
	case MsgExecGroup:
		mask.Set(Bits(MsgTypeBitsExecGroup))
	case MsgLeaveGroup:
		mask.Set(Bits(MsgTypeBitsLeaveGroup))
	}
}

func (mask MsgTypeBits) Names() []MsgType {
	names := make([]MsgType, mask.CountBits())
	var i int

	if mask.Has(Bits(MsgTypeBitsUnknown)) {
		names[i] = MsgUnknown
		i++
	}

	if mask.Has(Bits(MsgTypeBitsSetWithdrawAddress)) {
		names[i] = MsgSetWithdrawAddress
	}
	if mask.Has(Bits(MsgTypeBitsWithdrawDelegatorReward)) {
		names[i] = MsgWithdrawDelegatorReward
		i++
	}
	if mask.Has(Bits(MsgTypeBitsWithdrawValidatorCommission)) {
		names[i] = MsgWithdrawValidatorCommission
		i++
	}
	if mask.Has(Bits(MsgTypeBitsFundCommunityPool)) {
		names[i] = MsgFundCommunityPool
		i++
	}

	if mask.Has(Bits(MsgTypeBitsCreateValidator)) {
		names[i] = MsgCreateValidator
		i++
	}
	if mask.Has(Bits(MsgTypeBitsEditValidator)) {
		names[i] = MsgEditValidator
		i++
	}
	if mask.Has(Bits(MsgTypeBitsDelegate)) {
		names[i] = MsgDelegate
		i++
	}
	if mask.Has(Bits(MsgTypeBitsBeginRedelegate)) {
		names[i] = MsgBeginRedelegate
		i++
	}
	if mask.Has(Bits(MsgTypeBitsUndelegate)) {
		names[i] = MsgUndelegate
		i++
	}
	if mask.Has(Bits(MsgTypeBitsCancelUnbondingDelegation)) {
		names[i] = MsgCancelUnbondingDelegation
		i++
	}

	if mask.Has(Bits(MsgTypeBitsUnjail)) {
		names[i] = MsgUnjail
		i++
	}

	if mask.Has(Bits(MsgTypeBitsSend)) {
		names[i] = MsgSend
		i++
	}
	if mask.Has(Bits(MsgTypeBitsMultiSend)) {
		names[i] = MsgMultiSend
		i++
	}

	if mask.Has(Bits(MsgTypeBitsCreateVestingAccount)) {
		names[i] = MsgCreateVestingAccount
		i++
	}
	if mask.Has(Bits(MsgTypeBitsCreatePermanentLockedAccount)) {
		names[i] = MsgCreatePermanentLockedAccount
		i++
	}
	if mask.Has(Bits(MsgTypeBitsCreatePeriodicVestingAccount)) {
		names[i] = MsgCreatePeriodicVestingAccount
		i++
	}

	if mask.Has(Bits(MsgTypeBitsPayForBlobs)) {
		names[i] = MsgPayForBlobs
		i++
	}

	if mask.Has(Bits(MsgTypeBitsGrant)) {
		names[i] = MsgGrant
		i++
	}
	if mask.Has(Bits(MsgTypeBitsExec)) {
		names[i] = MsgExec
		i++
	}
	if mask.Has(Bits(MsgTypeBitsRevoke)) {
		names[i] = MsgRevoke
		i++
	}

	if mask.Has(Bits(MsgTypeBitsGrantAllowance)) {
		names[i] = MsgGrantAllowance
		i++
	}
	if mask.Has(Bits(MsgTypeBitsRevokeAllowance)) {
		names[i] = MsgRevokeAllowance
		i++
	}

	if mask.Has(Bits(MsgTypeBitsRegisterEVMAddress)) {
		names[i] = MsgRegisterEVMAddress
		i++
	}

	if mask.Has(Bits(MsgTypeBitsSubmitProposal)) {
		names[i] = MsgSubmitProposal
		i++
	}
	if mask.Has(Bits(MsgTypeBitsExecLegacyContent)) {
		names[i] = MsgExecLegacyContent
		i++
	}
	if mask.Has(Bits(MsgTypeBitsVote)) {
		names[i] = MsgVote
		i++
	}
	if mask.Has(Bits(MsgTypeBitsVoteWeighted)) {
		names[i] = MsgVoteWeighted
		i++
	}
	if mask.Has(Bits(MsgTypeBitsDeposit)) {
		names[i] = MsgDeposit
		i++
	}

	if mask.Has(Bits(MsgTypeBitsIBCTransfer)) {
		names[i] = IBCTransfer
		i++
	}

	if mask.Has(Bits(MsgTypeBitsVerifyInvariant)) {
		names[i] = MsgVerifyInvariant
		i++
	}

	if mask.Has(Bits(MsgTypeBitsSubmitEvidence)) {
		names[i] = MsgSubmitEvidence
		i++
	}

	if mask.Has(Bits(MsgTypeBitsSendNFT)) {
		names[i] = MsgSendNFT
		i++
	}

	if mask.Has(Bits(MsgTypeBitsCreateGroup)) {
		names[i] = MsgCreateGroup
		i++
	}
	if mask.Has(Bits(MsgTypeBitsUpdateGroupMembers)) {
		names[i] = MsgUpdateGroupMembers
		i++
	}
	if mask.Has(Bits(MsgTypeBitsUpdateGroupAdmin)) {
		names[i] = MsgUpdateGroupAdmin
		i++
	}
	if mask.Has(Bits(MsgTypeBitsUpdateGroupMetadata)) {
		names[i] = MsgUpdateGroupMetadata
		i++
	}
	if mask.Has(Bits(MsgTypeBitsCreateGroupPolicy)) {
		names[i] = MsgCreateGroupPolicy
		i++
	}
	if mask.Has(Bits(MsgTypeBitsUpdateGroupPolicyAdmin)) {
		names[i] = MsgUpdateGroupPolicyAdmin
		i++
	}
	if mask.Has(Bits(MsgTypeBitsCreateGroupWithPolicy)) {
		names[i] = MsgCreateGroupWithPolicy
		i++
	}
	if mask.Has(Bits(MsgTypeBitsUpdateGroupPolicyDecisionPolicy)) {
		names[i] = MsgUpdateGroupPolicyDecisionPolicy
		i++
	}
	if mask.Has(Bits(MsgTypeBitsUpdateGroupPolicyMetadata)) {
		names[i] = MsgUpdateGroupPolicyMetadata
		i++
	}
	if mask.Has(Bits(MsgTypeBitsSubmitProposalGroup)) {
		names[i] = MsgSubmitProposalGroup
		i++
	}
	if mask.Has(Bits(MsgTypeBitsWithdrawProposal)) {
		names[i] = MsgWithdrawProposal
		i++
	}
	if mask.Has(Bits(MsgTypeBitsVoteGroup)) {
		names[i] = MsgVoteGroup
		i++
	}
	if mask.Has(Bits(MsgTypeBitsExecGroup)) {
		names[i] = MsgExecGroup
		i++
	}
	if mask.Has(Bits(MsgTypeBitsLeaveGroup)) {
		names[i] = MsgLeaveGroup
		// i++
	}

	return names
}

func (mask MsgTypeBits) HasOne(value MsgTypeBits) bool {
	return mask.Bits&value.Bits > 0
}

var _ sql.Scanner = (*MsgTypeBits)(nil)

func (mask *MsgTypeBits) Scan(src interface{}) (err error) {
	switch val := src.(type) {
	case int64:
		mask.Bits = Bits(val)
	case nil:
		mask.Bits = 0
	default:
		return errors.Errorf("unknown bits database type: %T", src)
	}
	return nil
}

var _ driver.Valuer = (*MsgTypeBits)(nil)

func (mask MsgTypeBits) Value() (driver.Value, error) {
	return uint64(mask.Bits), nil
}
