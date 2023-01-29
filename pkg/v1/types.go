// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package whirlpool

import (
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
)

type OpenPositionBumps struct {
	PositionBump uint8
}

func (obj OpenPositionBumps) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `PositionBump` param:
	err = encoder.Encode(obj.PositionBump)
	if err != nil {
		return err
	}
	return nil
}

func (obj *OpenPositionBumps) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `PositionBump`:
	err = decoder.Decode(&obj.PositionBump)
	if err != nil {
		return err
	}
	return nil
}

type OpenPositionWithMetadataBumps struct {
	PositionBump uint8
	MetadataBump uint8
}

func (obj OpenPositionWithMetadataBumps) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `PositionBump` param:
	err = encoder.Encode(obj.PositionBump)
	if err != nil {
		return err
	}
	// Serialize `MetadataBump` param:
	err = encoder.Encode(obj.MetadataBump)
	if err != nil {
		return err
	}
	return nil
}

func (obj *OpenPositionWithMetadataBumps) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `PositionBump`:
	err = decoder.Decode(&obj.PositionBump)
	if err != nil {
		return err
	}
	// Deserialize `MetadataBump`:
	err = decoder.Decode(&obj.MetadataBump)
	if err != nil {
		return err
	}
	return nil
}

type PositionRewardInfo struct {
	GrowthInsideCheckpoint ag_binary.Uint128
	AmountOwed             uint64
}

func (obj PositionRewardInfo) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `GrowthInsideCheckpoint` param:
	err = encoder.Encode(obj.GrowthInsideCheckpoint)
	if err != nil {
		return err
	}
	// Serialize `AmountOwed` param:
	err = encoder.Encode(obj.AmountOwed)
	if err != nil {
		return err
	}
	return nil
}

func (obj *PositionRewardInfo) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `GrowthInsideCheckpoint`:
	err = decoder.Decode(&obj.GrowthInsideCheckpoint)
	if err != nil {
		return err
	}
	// Deserialize `AmountOwed`:
	err = decoder.Decode(&obj.AmountOwed)
	if err != nil {
		return err
	}
	return nil
}

type Tick struct {
	Initialized          bool
	LiquidityNet         ag_binary.Int128
	LiquidityGross       ag_binary.Uint128
	FeeGrowthOutsideA    ag_binary.Uint128
	FeeGrowthOutsideB    ag_binary.Uint128
	RewardGrowthsOutside [3]ag_binary.Uint128
}

func (obj Tick) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `Initialized` param:
	err = encoder.Encode(obj.Initialized)
	if err != nil {
		return err
	}
	// Serialize `LiquidityNet` param:
	err = encoder.Encode(obj.LiquidityNet)
	if err != nil {
		return err
	}
	// Serialize `LiquidityGross` param:
	err = encoder.Encode(obj.LiquidityGross)
	if err != nil {
		return err
	}
	// Serialize `FeeGrowthOutsideA` param:
	err = encoder.Encode(obj.FeeGrowthOutsideA)
	if err != nil {
		return err
	}
	// Serialize `FeeGrowthOutsideB` param:
	err = encoder.Encode(obj.FeeGrowthOutsideB)
	if err != nil {
		return err
	}
	// Serialize `RewardGrowthsOutside` param:
	err = encoder.Encode(obj.RewardGrowthsOutside)
	if err != nil {
		return err
	}
	return nil
}

func (obj *Tick) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `Initialized`:
	err = decoder.Decode(&obj.Initialized)
	if err != nil {
		return err
	}
	// Deserialize `LiquidityNet`:
	err = decoder.Decode(&obj.LiquidityNet)
	if err != nil {
		return err
	}
	// Deserialize `LiquidityGross`:
	err = decoder.Decode(&obj.LiquidityGross)
	if err != nil {
		return err
	}
	// Deserialize `FeeGrowthOutsideA`:
	err = decoder.Decode(&obj.FeeGrowthOutsideA)
	if err != nil {
		return err
	}
	// Deserialize `FeeGrowthOutsideB`:
	err = decoder.Decode(&obj.FeeGrowthOutsideB)
	if err != nil {
		return err
	}
	// Deserialize `RewardGrowthsOutside`:
	err = decoder.Decode(&obj.RewardGrowthsOutside)
	if err != nil {
		return err
	}
	return nil
}

type WhirlpoolRewardInfo struct {
	Mint                  ag_solanago.PublicKey
	Vault                 ag_solanago.PublicKey
	Authority             ag_solanago.PublicKey
	EmissionsPerSecondX64 ag_binary.Uint128
	GrowthGlobalX64       ag_binary.Uint128
}

func (obj WhirlpoolRewardInfo) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `Mint` param:
	err = encoder.Encode(obj.Mint)
	if err != nil {
		return err
	}
	// Serialize `Vault` param:
	err = encoder.Encode(obj.Vault)
	if err != nil {
		return err
	}
	// Serialize `Authority` param:
	err = encoder.Encode(obj.Authority)
	if err != nil {
		return err
	}
	// Serialize `EmissionsPerSecondX64` param:
	err = encoder.Encode(obj.EmissionsPerSecondX64)
	if err != nil {
		return err
	}
	// Serialize `GrowthGlobalX64` param:
	err = encoder.Encode(obj.GrowthGlobalX64)
	if err != nil {
		return err
	}
	return nil
}

func (obj *WhirlpoolRewardInfo) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `Mint`:
	err = decoder.Decode(&obj.Mint)
	if err != nil {
		return err
	}
	// Deserialize `Vault`:
	err = decoder.Decode(&obj.Vault)
	if err != nil {
		return err
	}
	// Deserialize `Authority`:
	err = decoder.Decode(&obj.Authority)
	if err != nil {
		return err
	}
	// Deserialize `EmissionsPerSecondX64`:
	err = decoder.Decode(&obj.EmissionsPerSecondX64)
	if err != nil {
		return err
	}
	// Deserialize `GrowthGlobalX64`:
	err = decoder.Decode(&obj.GrowthGlobalX64)
	if err != nil {
		return err
	}
	return nil
}

type WhirlpoolBumps struct {
	WhirlpoolBump uint8
}

func (obj WhirlpoolBumps) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `WhirlpoolBump` param:
	err = encoder.Encode(obj.WhirlpoolBump)
	if err != nil {
		return err
	}
	return nil
}

func (obj *WhirlpoolBumps) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `WhirlpoolBump`:
	err = decoder.Decode(&obj.WhirlpoolBump)
	if err != nil {
		return err
	}
	return nil
}

type CurrIndex ag_binary.BorshEnum

const (
	CurrIndexBelow CurrIndex = iota
	CurrIndexInside
	CurrIndexAbove
)

func (value CurrIndex) String() string {
	switch value {
	case CurrIndexBelow:
		return "Below"
	case CurrIndexInside:
		return "Inside"
	case CurrIndexAbove:
		return "Above"
	default:
		return ""
	}
}

type TickLabel ag_binary.BorshEnum

const (
	TickLabelUpper TickLabel = iota
	TickLabelLower
)

func (value TickLabel) String() string {
	switch value {
	case TickLabelUpper:
		return "Upper"
	case TickLabelLower:
		return "Lower"
	default:
		return ""
	}
}

type Direction ag_binary.BorshEnum

const (
	DirectionLeft Direction = iota
	DirectionRight
)

func (value Direction) String() string {
	switch value {
	case DirectionLeft:
		return "Left"
	case DirectionRight:
		return "Right"
	default:
		return ""
	}
}
