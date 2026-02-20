package pow

import (
    sdk "github.com/cosmos/cosmos-sdk/types"
)

// CreateCoinbaseTx creates a coinbase transaction with 99% miner reward and 1% blockshare reward.
func CreateCoinbaseTx(minerAddr sdk.AccAddress, blockshareAddr sdk.AccAddress, reward sdk.Int) sdk.Tx {
    // Calculate shares
    minerShare := reward.MulRaw(99).QuoRaw(100)   // 99%
    blockshareShare := reward.Sub(minerShare)     // 1%

    // Define outputs
    outputs := []sdk.Output{
        {Address: minerAddr, Coins: sdk.Coins{sdk.NewCoin("ibt", minerShare)}},
        {Address: blockshareAddr, Coins: sdk.Coins{sdk.NewCoin("ibt", blockshareShare)}},
    }

    // Return transaction (placeholder, adapt to Cosmos SDK Tx structure)
    return sdk.NewTx(outputs)
}
