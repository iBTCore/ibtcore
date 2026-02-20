package keeper

import (
    "context"
    sdk "github.com/cosmos/cosmos-sdk/types"
    "github.com/ibtcore/ibtcore/x/pow/types"
)

type queryServer struct {
    Keeper
}

func NewQueryServer(k Keeper) types.QueryServer {
    return &queryServer{Keeper: k}
}

// RewardInfo returns current reward split and blockshare address
func (q queryServer) RewardInfo(ctx context.Context, req *types.QueryRewardInfoRequest) (*types.QueryRewardInfoResponse, error) {
    return &types.QueryRewardInfoResponse{
        MinerPercent:      99,
        BlocksharePercent: 1,
        BlockshareAddress: q.Keeper.GetBlockshareAddress(ctx),
    }, nil
}
