package types

import (
	"encoding/json"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Ensure MsgSubmitAggregate implements sdk.Msg
var _ sdk.Msg = &MsgSubmitAggregate{}

// Route returns the module name
func (m *MsgSubmitAggregate) Route() string { return ModuleName }

// Type returns the action type
func (m *MsgSubmitAggregate) Type() string { return "submit_aggregate" }

// GetSigners returns the addresses of signers
func (m *MsgSubmitAggregate) GetSigners() []sdk.AccAddress {
	if m == nil {
		return nil
	}
	addr, err := sdk.AccAddressFromBech32(m.Miner)
	if err != nil {
		// if miner is not a bech32 address, return empty slice to let ValidateBasic catch it
		return []sdk.AccAddress{}
	}
	return []sdk.AccAddress{addr}
}

// GetSignBytes returns the bytes to sign over
func (m *MsgSubmitAggregate) GetSignBytes() []byte {
	// Use JSON canonical encoding via ModuleCdc (legacy Amino) for now.
	// For production, switch to proto-based sign bytes (e.g., using tx builder).
	bz, err := json.Marshal(m)
	if err != nil {
		// fallback to empty bytes on error
		return []byte{}
	}
	return sdk.MustSortJSON(bz)
}

// ValidateBasic performs stateless checks
func (m *MsgSubmitAggregate) ValidateBasic() error {
	if m == nil {
		return fmt.Errorf("nil message")
	}
	if len(m.Miner) == 0 {
		return fmt.Errorf("miner cannot be empty")
	}
	// check miner bech32 validity
	if _, err := sdk.AccAddressFromBech32(m.Miner); err != nil {
		return fmt.Errorf("invalid miner address: %w", err)
	}
	if m.ShareCount == 0 {
		return fmt.Errorf("share_count must be > 0")
	}
	if len(m.AggregateHash) == 0 {
		return fmt.Errorf("aggregate_hash cannot be empty")
	}
	return nil
}
