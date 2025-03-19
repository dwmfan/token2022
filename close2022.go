// Copyright 2025 github.com/dwnfan
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package token2022

import (
	"errors"
	"fmt"

	bin "github.com/gagliardetto/binary"
	solana "github.com/gagliardetto/solana-go"
	format "github.com/gagliardetto/solana-go/text/format"
	treeout "github.com/gagliardetto/treeout"
)

// Close2022 closes a token2022 mint account by transferring all its SOL to the destination account.
// Non-native accounts may only be closed if its token amount is zero.
type Close2022 struct {
	Account     solana.PublicKey `bin:"-" borsh_skip:"true"`
	Destination solana.PublicKey `bin:"-" borsh_skip:"true"`
	Owner       solana.PublicKey `bin:"-" borsh_skip:"true"`

	// [0] = [WRITE] account
	// ··········· The account to close.
	//
	// [1] = [WRITE] destination
	// ··········· The destination account.
	//
	// [2] = [SIGNER] owner
	// ··········· The account's owner.
	//
	// [3] = [] TokenProgram
	// ··········· Token 2022 program ID
	//
	// [4...] = [SIGNER] signers
	// ··········· M signer accounts.
	solana.AccountMetaSlice `bin:"-" borsh_skip:"true"`
	Signers  solana.AccountMetaSlice `bin:"-" borsh_skip:"true"`
}

// SetAccounts sets the accounts for the instruction.
func (obj *Close2022) SetAccounts(accounts []*solana.AccountMeta) error {
	obj.AccountMetaSlice, obj.Signers = solana.AccountMetaSlice(accounts).SplitFrom(4)
	return nil
}

// GetAccounts implements the AccountMetaGettable interface
func (slice Close2022) GetAccounts() (accounts []*solana.AccountMeta) {
	accounts = append(accounts, slice.AccountMetaSlice...)
	accounts = append(accounts, slice.Signers...)
	return
}

// NewClose2022InstructionBuilder creates a new `Close2022` instruction builder.
func NewClose2022InstructionBuilder() *Close2022 {
	nd := &Close2022{
		AccountMetaSlice: make(solana.AccountMetaSlice, 4),
		Signers:  make(solana.AccountMetaSlice, 0),
	}
	return nd
}

// SetAccount sets the "account" account.
// The account to close.
func (inst *Close2022) SetAccount(account solana.PublicKey) *Close2022 {
	inst.Account = account
	inst.AccountMetaSlice[0] = solana.Meta(account).WRITE()
	return inst
}

// GetAccount gets the "account" account.
// The account to close.
func (inst *Close2022) GetAccount() *solana.AccountMeta {
	return inst.AccountMetaSlice[0]
}

// SetDestinationAccount sets the "destination" account.
// The destination account.
func (inst *Close2022) SetDestinationAccount(destination solana.PublicKey) *Close2022 {
	inst.Destination = destination
	inst.AccountMetaSlice[1] = solana.Meta(destination).WRITE()
	return inst
}

// GetDestinationAccount gets the "destination" account.
// The destination account.
func (inst *Close2022) GetDestinationAccount() *solana.AccountMeta {
	return inst.AccountMetaSlice[1]
}

// SetOwnerAccount sets the "owner" account.
// The account's owner.
func (inst *Close2022) SetOwnerAccount(owner solana.PublicKey, multisigSigners ...solana.PublicKey) *Close2022 {
	inst.Owner = owner
	inst.AccountMetaSlice[2] = solana.Meta(owner).SIGNER()
	
	// Add token program
	inst.AccountMetaSlice[3] = solana.Meta(solana.Token2022ProgramID)
	
	for _, signer := range multisigSigners {
		inst.Signers = append(inst.Signers, solana.Meta(signer).SIGNER())
	}
	return inst
}

// GetOwnerAccount gets the "owner" account.
// The account's owner.
func (inst *Close2022) GetOwnerAccount() *solana.AccountMeta {
	return inst.AccountMetaSlice[2]
}

// Build builds the instruction.
func (inst Close2022) Build() *Instruction {
	return &Instruction{BaseVariant: bin.BaseVariant{
		Impl:   inst,
		TypeID: bin.NoTypeIDDefaultID,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst Close2022) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

// Validate validates the instruction parameters and accounts.
func (inst *Close2022) Validate() error {
	// Check whether all (required) accounts are set:
	{
		if inst.Account.IsZero() || inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.Account is not set")
		}
		if inst.Destination.IsZero() || inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.Destination is not set")
		}
		if inst.Owner.IsZero() || inst.AccountMetaSlice[2] == nil {
			return errors.New("accounts.Owner is not set")
		}
		if inst.AccountMetaSlice[3] == nil {
			return errors.New("accounts.TokenProgram is not set")
		}
		if !inst.AccountMetaSlice[2].IsSigner && len(inst.Signers) == 0 {
			return fmt.Errorf("accounts.Signers is not set")
		}
		if len(inst.Signers) > 11 {
			return fmt.Errorf("too many signers; got %v, but max is 11", len(inst.Signers))
		}
	}
	return nil
}

// EncodeToTree encodes the instruction to a tree.
func (inst *Close2022) EncodeToTree(parent treeout.Branches) {
	parent.Child(format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch treeout.Branches) {
			programBranch.Child(format.Instruction("Close2022")).
				//
				ParentFunc(func(instructionBranch treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params").ParentFunc(func(paramsBranch treeout.Branches) {})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts").ParentFunc(func(accountsBranch treeout.Branches) {
						accountsBranch.Child(format.Meta("         account", inst.AccountMetaSlice[0]))
						accountsBranch.Child(format.Meta("    destination", inst.AccountMetaSlice[1]))
						accountsBranch.Child(format.Meta("          owner", inst.AccountMetaSlice[2]))
						accountsBranch.Child(format.Meta("   tokenProgram", inst.AccountMetaSlice[3]))

						signersBranch := accountsBranch.Child(fmt.Sprintf("signers[len=%v]", len(inst.Signers)))
						for i, v := range inst.Signers {
							if len(inst.Signers) > 9 && i < 10 {
								signersBranch.Child(format.Meta(fmt.Sprintf(" [%v]", i), v))
							} else {
								signersBranch.Child(format.Meta(fmt.Sprintf("[%v]", i), v))
							}
						}
					})
				})
		})
}

// MarshalWithEncoder implements the bin.EncoderDecoder interface
func (obj Close2022) MarshalWithEncoder(encoder *bin.Encoder) error {
	return encoder.WriteBytes([]byte{}, false)
}

// UnmarshalWithDecoder implements the bin.EncoderDecoder interface
func (obj *Close2022) UnmarshalWithDecoder(decoder *bin.Decoder) error {
	return nil
}

// NewClose2022Instruction declares a new Close2022 instruction with the provided parameters and accounts.
func NewClose2022Instruction(
	// Accounts:
	account solana.PublicKey,
	destination solana.PublicKey,
	owner solana.PublicKey,
	multisigSigners []solana.PublicKey,
) *Close2022 {
	return NewClose2022InstructionBuilder().
		SetAccount(account).
		SetDestinationAccount(destination).
		SetOwnerAccount(owner, multisigSigners...)
}
