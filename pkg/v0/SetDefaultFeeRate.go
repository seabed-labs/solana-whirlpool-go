// Code generated by https://github.com/dcaf-mocha/anchor-go. DO NOT EDIT.

package whirlpool

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// SetDefaultFeeRate is the `setDefaultFeeRate` instruction.
type SetDefaultFeeRate struct {
	DefaultFeeRate *uint16

	// [0] = [] whirlpoolsConfig
	//
	// [1] = [WRITE] feeTier
	//
	// [2] = [SIGNER] feeAuthority
	ag_solanago.AccountMetaSlice `bin:"-"`
}

type SetDefaultFeeRateAccounts struct {
	WhirlpoolsConfig ag_solanago.PublicKey
	FeeTier          ag_solanago.PublicKey
	FeeAuthority     ag_solanago.PublicKey
}

// NewSetDefaultFeeRateInstructionBuilder creates a new `SetDefaultFeeRate` instruction builder.
func NewSetDefaultFeeRateInstructionBuilder() *SetDefaultFeeRate {
	nd := &SetDefaultFeeRate{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 3),
	}
	return nd
}

func (inst *SetDefaultFeeRate) GetSetDefaultFeeRateAccounts() *SetDefaultFeeRateAccounts {
	res := &SetDefaultFeeRateAccounts{}
	if inst.AccountMetaSlice[0] != nil {
		res.WhirlpoolsConfig = inst.AccountMetaSlice[0].PublicKey
	}
	if inst.AccountMetaSlice[1] != nil {
		res.FeeTier = inst.AccountMetaSlice[1].PublicKey
	}
	if inst.AccountMetaSlice[2] != nil {
		res.FeeAuthority = inst.AccountMetaSlice[2].PublicKey
	}
	return res
}

// SetDefaultFeeRate sets the "defaultFeeRate" parameter.
func (inst *SetDefaultFeeRate) SetDefaultFeeRate(defaultFeeRate uint16) *SetDefaultFeeRate {
	inst.DefaultFeeRate = &defaultFeeRate
	return inst
}

// SetWhirlpoolsConfigAccount sets the "whirlpoolsConfig" account.
func (inst *SetDefaultFeeRate) SetWhirlpoolsConfigAccount(whirlpoolsConfig ag_solanago.PublicKey) *SetDefaultFeeRate {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(whirlpoolsConfig)
	return inst
}

// GetWhirlpoolsConfigAccount gets the "whirlpoolsConfig" account.
func (inst *SetDefaultFeeRate) GetWhirlpoolsConfigAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetFeeTierAccount sets the "feeTier" account.
func (inst *SetDefaultFeeRate) SetFeeTierAccount(feeTier ag_solanago.PublicKey) *SetDefaultFeeRate {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(feeTier).WRITE()
	return inst
}

// GetFeeTierAccount gets the "feeTier" account.
func (inst *SetDefaultFeeRate) GetFeeTierAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

// SetFeeAuthorityAccount sets the "feeAuthority" account.
func (inst *SetDefaultFeeRate) SetFeeAuthorityAccount(feeAuthority ag_solanago.PublicKey) *SetDefaultFeeRate {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(feeAuthority).SIGNER()
	return inst
}

// GetFeeAuthorityAccount gets the "feeAuthority" account.
func (inst *SetDefaultFeeRate) GetFeeAuthorityAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(2)
}

func (inst SetDefaultFeeRate) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_SetDefaultFeeRate,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst SetDefaultFeeRate) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *SetDefaultFeeRate) Validate() error {
	// Check whether all (required) parameters are set:
	{
		if inst.DefaultFeeRate == nil {
			return errors.New("DefaultFeeRate parameter is not set")
		}
	}

	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.WhirlpoolsConfig is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.FeeTier is not set")
		}
		if inst.AccountMetaSlice[2] == nil {
			return errors.New("accounts.FeeAuthority is not set")
		}
	}
	return nil
}

func (inst *SetDefaultFeeRate) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("SetDefaultFeeRate")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=1]").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param("DefaultFeeRate", *inst.DefaultFeeRate))
					})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=3]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("whirlpoolsConfig", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta("         feeTier", inst.AccountMetaSlice.Get(1)))
						accountsBranch.Child(ag_format.Meta("    feeAuthority", inst.AccountMetaSlice.Get(2)))
					})
				})
		})
}

func (obj SetDefaultFeeRate) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `DefaultFeeRate` param:
	err = encoder.Encode(obj.DefaultFeeRate)
	if err != nil {
		return err
	}
	return nil
}
func (obj *SetDefaultFeeRate) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `DefaultFeeRate`:
	err = decoder.Decode(&obj.DefaultFeeRate)
	if err != nil {
		return err
	}
	return nil
}

// NewSetDefaultFeeRateInstruction declares a new SetDefaultFeeRate instruction with the provided parameters and accounts.
func NewSetDefaultFeeRateInstruction(
	// Parameters:
	defaultFeeRate uint16,
	// Accounts:
	whirlpoolsConfig ag_solanago.PublicKey,
	feeTier ag_solanago.PublicKey,
	feeAuthority ag_solanago.PublicKey) *SetDefaultFeeRate {
	return NewSetDefaultFeeRateInstructionBuilder().
		SetDefaultFeeRate(defaultFeeRate).
		SetWhirlpoolsConfigAccount(whirlpoolsConfig).
		SetFeeTierAccount(feeTier).
		SetFeeAuthorityAccount(feeAuthority)
}
