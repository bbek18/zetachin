package keeper

import (
	"context"
	"fmt"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/zeta-chain/zetacore/x/crosschain/types"
	observertypes "github.com/zeta-chain/zetacore/x/observer/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func getOutTrackerIndex(chainID int64, nonce uint64) string {
	return fmt.Sprintf("%d-%d", chainID, nonce)
}

// SetOutTxTracker set a specific outTxTracker in the store from its index
func (k Keeper) SetOutTxTracker(ctx sdk.Context, outTxTracker types.OutTxTracker) {
	outTxTracker.Index = getOutTrackerIndex(outTxTracker.ChainId, outTxTracker.Nonce)
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.OutTxTrackerKeyPrefix))
	b := k.cdc.MustMarshal(&outTxTracker)
	store.Set(types.OutTxTrackerKey(
		outTxTracker.Index,
	), b)
}

// GetOutTxTracker returns a outTxTracker from its index
func (k Keeper) GetOutTxTracker(
	ctx sdk.Context,
	chainID int64,
	nonce uint64,

) (val types.OutTxTracker, found bool) {
	index := getOutTrackerIndex(chainID, nonce)
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.OutTxTrackerKeyPrefix))

	b := store.Get(types.OutTxTrackerKey(
		index,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveOutTxTracker removes a outTxTracker from the store
func (k Keeper) RemoveOutTxTracker(
	ctx sdk.Context,
	chainID int64,
	nonce uint64,

) {
	index := getOutTrackerIndex(chainID, nonce)
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.OutTxTrackerKeyPrefix))
	store.Delete(types.OutTxTrackerKey(
		index,
	))
}

// GetAllOutTxTracker returns all outTxTracker
func (k Keeper) GetAllOutTxTracker(ctx sdk.Context) (list []types.OutTxTracker) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.OutTxTrackerKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.OutTxTracker
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// Queries

func (k Keeper) OutTxTrackerAll(c context.Context, req *types.QueryAllOutTxTrackerRequest) (*types.QueryAllOutTxTrackerResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var outTxTrackers []types.OutTxTracker
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	outTxTrackerStore := prefix.NewStore(store, types.KeyPrefix(types.OutTxTrackerKeyPrefix))
	pageRes, err := query.Paginate(outTxTrackerStore, req.Pagination, func(key []byte, value []byte) error {
		var outTxTracker types.OutTxTracker
		if err := k.cdc.Unmarshal(value, &outTxTracker); err != nil {
			return err
		}

		outTxTrackers = append(outTxTrackers, outTxTracker)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &types.QueryAllOutTxTrackerResponse{OutTxTracker: outTxTrackers, Pagination: pageRes}, nil
}

func (k Keeper) OutTxTrackerAllByChain(c context.Context, req *types.QueryAllOutTxTrackerByChainRequest) (*types.QueryAllOutTxTrackerByChainResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var outTxTrackers []types.OutTxTracker
	ctx := sdk.UnwrapSDKContext(c)

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.OutTxTrackerKeyPrefix))
	chainStore := prefix.NewStore(store, types.KeyPrefix(fmt.Sprintf("%d-", req.Chain)))

	pageRes, err := query.Paginate(chainStore, req.Pagination, func(key []byte, value []byte) error {
		var outTxTracker types.OutTxTracker
		if err := k.cdc.Unmarshal(value, &outTxTracker); err != nil {
			return err
		}
		outTxTrackers = append(outTxTrackers, outTxTracker)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllOutTxTrackerByChainResponse{OutTxTracker: outTxTrackers, Pagination: pageRes}, nil
}

func (k Keeper) OutTxTracker(c context.Context, req *types.QueryGetOutTxTrackerRequest) (*types.QueryGetOutTxTrackerResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)
	val, found := k.GetOutTxTracker(
		ctx,
		req.ChainID,
		req.Nonce,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetOutTxTrackerResponse{OutTxTracker: val}, nil
}

// Messages

// RemoveFromOutTxTracker removes a record from the outbound transaction tracker by chain ID and nonce.
// only the admin policy account is authorized to broadcast this message.
func (k msgServer) RemoveFromOutTxTracker(goCtx context.Context, msg *types.MsgRemoveFromOutTxTracker) (*types.MsgRemoveFromOutTxTrackerResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	if msg.Creator != k.zetaObserverKeeper.GetParams(ctx).GetAdminPolicyAccount(observertypes.Policy_Type_group1) {
		return &types.MsgRemoveFromOutTxTrackerResponse{}, observertypes.ErrNotAuthorizedPolicy
	}

	k.RemoveOutTxTracker(ctx, msg.ChainId, msg.Nonce)
	return &types.MsgRemoveFromOutTxTrackerResponse{}, nil
}
