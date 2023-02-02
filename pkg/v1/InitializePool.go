// Code generated by https://github.com/dcaf-mocha/anchor-go. DO NOT EDIT.

package whirlpool

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// InitializePool is the `initializePool` instruction.
type InitializePool struct {
	Bumps            *WhirlpoolBumps
	TickSpacing      *uint16
	InitialSqrtPrice *ag_binary.Uint128

	// [0] = [] whirlpoolsConfig
	//
	// [1] = [] tokenMintA
	//
	// [2] = [] tokenMintB
	//
	// [3] = [WRITE, SIGNER] funder
	//
	// [4] = [WRITE] whirlpool
	//
	// [5] = [WRITE, SIGNER] tokenVaultA
	//
	// [6] = [WRITE, SIGNER] tokenVaultB
	//
	// [7] = [] feeTier
	//
	// [8] = [] tokenProgram
	//
	// [9] = [] systemProgram
	//
	// [10] = [] rent
	ag_solanago.AccountMetaSlice `bin:"-"`
}

type InitializePoolAccounts struct {
	WhirlpoolsConfig ag_solanago.PublicKey
	TokenMintA       ag_solanago.PublicKey
	TokenMintB       ag_solanago.PublicKey
	Funder           ag_solanago.PublicKey
	Whirlpool        ag_solanago.PublicKey
	TokenVaultA      ag_solanago.PublicKey
	TokenVaultB      ag_solanago.PublicKey
	FeeTier          ag_solanago.PublicKey
	TokenProgram     ag_solanago.PublicKey
	SystemProgram    ag_solanago.PublicKey
	Rent             ag_solanago.PublicKey
}

// NewInitializePoolInstructionBuilder creates a new `InitializePool` instruction builder.
func NewInitializePoolInstructionBuilder() *InitializePool {
	nd := &InitializePool{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 11),
	}
	return nd
}

func (inst *InitializePool) GetInitializePoolAccounts() *InitializePoolAccounts {
	res := &InitializePoolAccounts{}
	if inst.AccountMetaSlice[0] != nil {
		res.WhirlpoolsConfig = inst.AccountMetaSlice[0].PublicKey
	}
	if inst.AccountMetaSlice[1] != nil {
		res.TokenMintA = inst.AccountMetaSlice[1].PublicKey
	}
	if inst.AccountMetaSlice[2] != nil {
		res.TokenMintB = inst.AccountMetaSlice[2].PublicKey
	}
	if inst.AccountMetaSlice[3] != nil {
		res.Funder = inst.AccountMetaSlice[3].PublicKey
	}
	if inst.AccountMetaSlice[4] != nil {
		res.Whirlpool = inst.AccountMetaSlice[4].PublicKey
	}
	if inst.AccountMetaSlice[5] != nil {
		res.TokenVaultA = inst.AccountMetaSlice[5].PublicKey
	}
	if inst.AccountMetaSlice[6] != nil {
		res.TokenVaultB = inst.AccountMetaSlice[6].PublicKey
	}
	if inst.AccountMetaSlice[7] != nil {
		res.FeeTier = inst.AccountMetaSlice[7].PublicKey
	}
	if inst.AccountMetaSlice[8] != nil {
		res.TokenProgram = inst.AccountMetaSlice[8].PublicKey
	}
	if inst.AccountMetaSlice[9] != nil {
		res.SystemProgram = inst.AccountMetaSlice[9].PublicKey
	}
	if inst.AccountMetaSlice[10] != nil {
		res.Rent = inst.AccountMetaSlice[10].PublicKey
	}
	return res
}

// SetBumps sets the "bumps" parameter.
func (inst *InitializePool) SetBumps(bumps WhirlpoolBumps) *InitializePool {
	inst.Bumps = &bumps
	return inst
}

// SetTickSpacing sets the "tickSpacing" parameter.
func (inst *InitializePool) SetTickSpacing(tickSpacing uint16) *InitializePool {
	inst.TickSpacing = &tickSpacing
	return inst
}

// SetInitialSqrtPrice sets the "initialSqrtPrice" parameter.
func (inst *InitializePool) SetInitialSqrtPrice(initialSqrtPrice ag_binary.Uint128) *InitializePool {
	inst.InitialSqrtPrice = &initialSqrtPrice
	return inst
}

// SetWhirlpoolsConfigAccount sets the "whirlpoolsConfig" account.
func (inst *InitializePool) SetWhirlpoolsConfigAccount(whirlpoolsConfig ag_solanago.PublicKey) *InitializePool {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(whirlpoolsConfig)
	return inst
}

// GetWhirlpoolsConfigAccount gets the "whirlpoolsConfig" account.
func (inst *InitializePool) GetWhirlpoolsConfigAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetTokenMintAAccount sets the "tokenMintA" account.
func (inst *InitializePool) SetTokenMintAAccount(tokenMintA ag_solanago.PublicKey) *InitializePool {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(tokenMintA)
	return inst
}

// GetTokenMintAAccount gets the "tokenMintA" account.
func (inst *InitializePool) GetTokenMintAAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

// SetTokenMintBAccount sets the "tokenMintB" account.
func (inst *InitializePool) SetTokenMintBAccount(tokenMintB ag_solanago.PublicKey) *InitializePool {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(tokenMintB)
	return inst
}

// GetTokenMintBAccount gets the "tokenMintB" account.
func (inst *InitializePool) GetTokenMintBAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(2)
}

// SetFunderAccount sets the "funder" account.
func (inst *InitializePool) SetFunderAccount(funder ag_solanago.PublicKey) *InitializePool {
	inst.AccountMetaSlice[3] = ag_solanago.Meta(funder).WRITE().SIGNER()
	return inst
}

// GetFunderAccount gets the "funder" account.
func (inst *InitializePool) GetFunderAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(3)
}

// SetWhirlpoolAccount sets the "whirlpool" account.
func (inst *InitializePool) SetWhirlpoolAccount(whirlpool ag_solanago.PublicKey) *InitializePool {
	inst.AccountMetaSlice[4] = ag_solanago.Meta(whirlpool).WRITE()
	return inst
}

// GetWhirlpoolAccount gets the "whirlpool" account.
func (inst *InitializePool) GetWhirlpoolAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(4)
}

// SetTokenVaultAAccount sets the "tokenVaultA" account.
func (inst *InitializePool) SetTokenVaultAAccount(tokenVaultA ag_solanago.PublicKey) *InitializePool {
	inst.AccountMetaSlice[5] = ag_solanago.Meta(tokenVaultA).WRITE().SIGNER()
	return inst
}

// GetTokenVaultAAccount gets the "tokenVaultA" account.
func (inst *InitializePool) GetTokenVaultAAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(5)
}

// SetTokenVaultBAccount sets the "tokenVaultB" account.
func (inst *InitializePool) SetTokenVaultBAccount(tokenVaultB ag_solanago.PublicKey) *InitializePool {
	inst.AccountMetaSlice[6] = ag_solanago.Meta(tokenVaultB).WRITE().SIGNER()
	return inst
}

// GetTokenVaultBAccount gets the "tokenVaultB" account.
func (inst *InitializePool) GetTokenVaultBAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(6)
}

// SetFeeTierAccount sets the "feeTier" account.
func (inst *InitializePool) SetFeeTierAccount(feeTier ag_solanago.PublicKey) *InitializePool {
	inst.AccountMetaSlice[7] = ag_solanago.Meta(feeTier)
	return inst
}

// GetFeeTierAccount gets the "feeTier" account.
func (inst *InitializePool) GetFeeTierAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(7)
}

// SetTokenProgramAccount sets the "tokenProgram" account.
func (inst *InitializePool) SetTokenProgramAccount(tokenProgram ag_solanago.PublicKey) *InitializePool {
	inst.AccountMetaSlice[8] = ag_solanago.Meta(tokenProgram)
	return inst
}

// GetTokenProgramAccount gets the "tokenProgram" account.
func (inst *InitializePool) GetTokenProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(8)
}

// SetSystemProgramAccount sets the "systemProgram" account.
func (inst *InitializePool) SetSystemProgramAccount(systemProgram ag_solanago.PublicKey) *InitializePool {
	inst.AccountMetaSlice[9] = ag_solanago.Meta(systemProgram)
	return inst
}

// GetSystemProgramAccount gets the "systemProgram" account.
func (inst *InitializePool) GetSystemProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(9)
}

// SetRentAccount sets the "rent" account.
func (inst *InitializePool) SetRentAccount(rent ag_solanago.PublicKey) *InitializePool {
	inst.AccountMetaSlice[10] = ag_solanago.Meta(rent)
	return inst
}

// GetRentAccount gets the "rent" account.
func (inst *InitializePool) GetRentAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(10)
}

func (inst InitializePool) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_InitializePool,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst InitializePool) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *InitializePool) Validate() error {
	// Check whether all (required) parameters are set:
	{
		if inst.Bumps == nil {
			return errors.New("Bumps parameter is not set")
		}
		if inst.TickSpacing == nil {
			return errors.New("TickSpacing parameter is not set")
		}
		if inst.InitialSqrtPrice == nil {
			return errors.New("InitialSqrtPrice parameter is not set")
		}
	}

	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.WhirlpoolsConfig is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.TokenMintA is not set")
		}
		if inst.AccountMetaSlice[2] == nil {
			return errors.New("accounts.TokenMintB is not set")
		}
		if inst.AccountMetaSlice[3] == nil {
			return errors.New("accounts.Funder is not set")
		}
		if inst.AccountMetaSlice[4] == nil {
			return errors.New("accounts.Whirlpool is not set")
		}
		if inst.AccountMetaSlice[5] == nil {
			return errors.New("accounts.TokenVaultA is not set")
		}
		if inst.AccountMetaSlice[6] == nil {
			return errors.New("accounts.TokenVaultB is not set")
		}
		if inst.AccountMetaSlice[7] == nil {
			return errors.New("accounts.FeeTier is not set")
		}
		if inst.AccountMetaSlice[8] == nil {
			return errors.New("accounts.TokenProgram is not set")
		}
		if inst.AccountMetaSlice[9] == nil {
			return errors.New("accounts.SystemProgram is not set")
		}
		if inst.AccountMetaSlice[10] == nil {
			return errors.New("accounts.Rent is not set")
		}
	}
	return nil
}

func (inst *InitializePool) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("InitializePool")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=3]").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param("           Bumps", *inst.Bumps))
						paramsBranch.Child(ag_format.Param("     TickSpacing", *inst.TickSpacing))
						paramsBranch.Child(ag_format.Param("InitialSqrtPrice", *inst.InitialSqrtPrice))
					})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=11]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("whirlpoolsConfig", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta("      tokenMintA", inst.AccountMetaSlice.Get(1)))
						accountsBranch.Child(ag_format.Meta("      tokenMintB", inst.AccountMetaSlice.Get(2)))
						accountsBranch.Child(ag_format.Meta("          funder", inst.AccountMetaSlice.Get(3)))
						accountsBranch.Child(ag_format.Meta("       whirlpool", inst.AccountMetaSlice.Get(4)))
						accountsBranch.Child(ag_format.Meta("     tokenVaultA", inst.AccountMetaSlice.Get(5)))
						accountsBranch.Child(ag_format.Meta("     tokenVaultB", inst.AccountMetaSlice.Get(6)))
						accountsBranch.Child(ag_format.Meta("         feeTier", inst.AccountMetaSlice.Get(7)))
						accountsBranch.Child(ag_format.Meta("    tokenProgram", inst.AccountMetaSlice.Get(8)))
						accountsBranch.Child(ag_format.Meta("   systemProgram", inst.AccountMetaSlice.Get(9)))
						accountsBranch.Child(ag_format.Meta("            rent", inst.AccountMetaSlice.Get(10)))
					})
				})
		})
}

func (obj InitializePool) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `Bumps` param:
	err = encoder.Encode(obj.Bumps)
	if err != nil {
		return err
	}
	// Serialize `TickSpacing` param:
	err = encoder.Encode(obj.TickSpacing)
	if err != nil {
		return err
	}
	// Serialize `InitialSqrtPrice` param:
	err = encoder.Encode(obj.InitialSqrtPrice)
	if err != nil {
		return err
	}
	return nil
}
func (obj *InitializePool) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `Bumps`:
	err = decoder.Decode(&obj.Bumps)
	if err != nil {
		return err
	}
	// Deserialize `TickSpacing`:
	err = decoder.Decode(&obj.TickSpacing)
	if err != nil {
		return err
	}
	// Deserialize `InitialSqrtPrice`:
	err = decoder.Decode(&obj.InitialSqrtPrice)
	if err != nil {
		return err
	}
	return nil
}

// NewInitializePoolInstruction declares a new InitializePool instruction with the provided parameters and accounts.
func NewInitializePoolInstruction(
	// Parameters:
	bumps WhirlpoolBumps,
	tickSpacing uint16,
	initialSqrtPrice ag_binary.Uint128,
	// Accounts:
	whirlpoolsConfig ag_solanago.PublicKey,
	tokenMintA ag_solanago.PublicKey,
	tokenMintB ag_solanago.PublicKey,
	funder ag_solanago.PublicKey,
	whirlpool ag_solanago.PublicKey,
	tokenVaultA ag_solanago.PublicKey,
	tokenVaultB ag_solanago.PublicKey,
	feeTier ag_solanago.PublicKey,
	tokenProgram ag_solanago.PublicKey,
	systemProgram ag_solanago.PublicKey,
	rent ag_solanago.PublicKey) *InitializePool {
	return NewInitializePoolInstructionBuilder().
		SetBumps(bumps).
		SetTickSpacing(tickSpacing).
		SetInitialSqrtPrice(initialSqrtPrice).
		SetWhirlpoolsConfigAccount(whirlpoolsConfig).
		SetTokenMintAAccount(tokenMintA).
		SetTokenMintBAccount(tokenMintB).
		SetFunderAccount(funder).
		SetWhirlpoolAccount(whirlpool).
		SetTokenVaultAAccount(tokenVaultA).
		SetTokenVaultBAccount(tokenVaultB).
		SetFeeTierAccount(feeTier).
		SetTokenProgramAccount(tokenProgram).
		SetSystemProgramAccount(systemProgram).
		SetRentAccount(rent)
}
