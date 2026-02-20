package types

import "fmt"

func (p *Params) Validate() error {
	if p.BlockReward < 0 {
		return fmt.Errorf("block reward must be non-negative")
	}
	if p.FounderShare < 0 || p.FounderShare > 100 {
		return fmt.Errorf("founder share must be between 0 and 100")
	}
	if p.MinerShare < 0 || p.MinerShare > 100 {
		return fmt.Errorf("miner share must be between 0 and 100")
	}
	if p.FounderShare+p.MinerShare != 100 {
		return fmt.Errorf("founder share + miner share must equal 100")
	}
	if p.TxFeeFlat < 0 {
		return fmt.Errorf("tx fee flat must be non-negative")
	}
	if p.TxFeeFounder < 0 || p.TxFeeFounder > 100 {
		return fmt.Errorf("tx fee founder must be between 0 and 100")
	}
	if p.TxFeeMiner < 0 || p.TxFeeMiner > 100 {
		return fmt.Errorf("tx fee miner must be between 0 and 100")
	}
	return nil
}
