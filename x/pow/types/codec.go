package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const ModuleName = "pow"

var ModuleCdc = codec.NewLegacyAmino()

func RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgSubmitAggregate{}, "pow/MsgSubmitAggregate", nil)
}

func RegisterInterfaces(reg cdctypes.InterfaceRegistry) {
	reg.RegisterImplementations((*sdk.Msg)(nil),
		&MsgSubmitAggregate{},
	)
}

func DefaultParams() Params {
	return Params{
		BlockReward:  499429244, // 0.499429244 IBTC (in smallest units)
		FounderShare: 1,         // percent
		MinerShare:   99,        // percent
		TxFeeFlat:    50000,     // example smallest unit per tx
		TxFeeFounder: 10,        // percent
		TxFeeMiner:   90,        // percent
	}
}
