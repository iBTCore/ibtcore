package types

import "encoding/json"

type GenesisState struct {
	Params Params `json:"params"`
	// add other state like difficulty, total_minted, epoch, etc.
}

func DefaultGenesis() GenesisState {
	return GenesisState{
		Params: DefaultParams(),
	}
}

func ValidateGenesis(bz json.RawMessage) error {
	var gs GenesisState
	if err := json.Unmarshal(bz, &gs); err != nil {
		return err
	}
	return gs.Params.Validate()
}
