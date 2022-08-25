package keeper

import (
	"context"
	"fmt"
	"strings"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"

	"github.com/kava-labs/kava/x/earn/types"
)

type queryServer struct {
	keeper Keeper
}

// NewQueryServerImpl creates a new server for handling gRPC queries.
func NewQueryServerImpl(k Keeper) types.QueryServer {
	return &queryServer{keeper: k}
}

var _ types.QueryServer = queryServer{}

// Params implements the gRPC service handler for querying x/earn parameters.
func (s queryServer) Params(
	ctx context.Context,
	req *types.QueryParamsRequest,
) (*types.QueryParamsResponse, error) {
	if req == nil {
		return nil, status.Errorf(codes.InvalidArgument, "empty request")
	}

	sdkCtx := sdk.UnwrapSDKContext(ctx)
	params := s.keeper.GetParams(sdkCtx)

	return &types.QueryParamsResponse{Params: params}, nil
}

// Vaults implements the gRPC service handler for querying x/earn vaults.
func (s queryServer) Vaults(
	ctx context.Context,
	req *types.QueryVaultsRequest,
) (*types.QueryVaultsResponse, error) {
	if req == nil {
		return nil, status.Errorf(codes.InvalidArgument, "empty request")
	}

	sdkCtx := sdk.UnwrapSDKContext(ctx)

	vaults := []types.VaultResponse{}

	var vaultRecordsErr error

	// Iterate over vault records instead of AllowedVaults to get all bkava-*
	// vaults
	s.keeper.IterateVaultRecords(sdkCtx, func(record types.VaultRecord) bool {
		allowedVault, found := s.keeper.GetAllowedVault(sdkCtx, record.TotalShares.Denom)
		if !found {
			vaultRecordsErr = fmt.Errorf("vault record not found for vault record denom %s", record.TotalShares.Denom)
		}

		totalValue, err := s.keeper.GetVaultTotalValue(sdkCtx, record.TotalShares.Denom)
		if err != nil {
			vaultRecordsErr = err
			// Stop iterating if error
			return true
		}

		vaults = append(vaults, types.VaultResponse{
			Denom:             record.TotalShares.Denom,
			Strategies:        allowedVault.Strategies,
			IsPrivateVault:    allowedVault.IsPrivateVault,
			AllowedDepositors: addressSliceToStringSlice(allowedVault.AllowedDepositors),
			TotalShares:       record.TotalShares.Amount.String(),
			TotalValue:        totalValue.Amount,
		})

		return false
	})

	if vaultRecordsErr != nil {
		return nil, vaultRecordsErr
	}

	// Does not include vaults that have no deposits, only iterates over vault
	// records which exists only for those with deposits.
	return &types.QueryVaultsResponse{
		Vaults: vaults,
	}, nil
}

// Vaults implements the gRPC service handler for querying x/earn vaults.
func (s queryServer) Vault(
	ctx context.Context,
	req *types.QueryVaultRequest,
) (*types.QueryVaultResponse, error) {
	if req == nil {
		return nil, status.Errorf(codes.InvalidArgument, "empty request")
	}

	sdkCtx := sdk.UnwrapSDKContext(ctx)

	if req.Denom == "" {
		return nil, status.Errorf(codes.InvalidArgument, "empty denom")
	}

	// Only 1 vault
	allowedVault, found := s.keeper.GetAllowedVault(sdkCtx, req.Denom)
	if !found {
		return nil, status.Errorf(codes.NotFound, "vault not found with specified denom")
	}

	// Handle bkava separately to get total of **all** bkava vaults
	if req.Denom == "bkava" {
		return s.getAggregateBkavaVault(sdkCtx, allowedVault)
	}

	// Must be req.Denom and not allowedVault.Denom to get full "bkava" denom
	vaultRecord, found := s.keeper.GetVaultRecord(sdkCtx, req.Denom)
	if !found {
		// No supply yet, no error just set it to zero
		vaultRecord.TotalShares = types.NewVaultShare(req.Denom, sdk.ZeroDec())
	}

	totalValue, err := s.keeper.GetVaultTotalValue(sdkCtx, req.Denom)
	if err != nil {
		return nil, err
	}

	vault := types.VaultResponse{
		// VaultRecord denom instead of AllowedVault.Denom for full bkava denom
		Denom:             vaultRecord.TotalShares.Denom,
		Strategies:        allowedVault.Strategies,
		IsPrivateVault:    allowedVault.IsPrivateVault,
		AllowedDepositors: addressSliceToStringSlice(allowedVault.AllowedDepositors),
		TotalShares:       vaultRecord.TotalShares.Amount.String(),
		TotalValue:        totalValue.Amount,
	}

	return &types.QueryVaultResponse{
		Vault: vault,
	}, nil
}

// getAggregateBkavaVault returns a VaultResponse of the total of all bkava
// vaults.
func (s queryServer) getAggregateBkavaVault(
	ctx sdk.Context,
	allowedVault types.AllowedVault,
) (*types.QueryVaultResponse, error) {
	totalValue := sdk.NewInt(0)

	var iterErr error
	s.keeper.IterateVaultRecords(ctx, func(record types.VaultRecord) (stop bool) {
		// Skip non bkava vaults
		if !strings.HasPrefix(record.TotalShares.Denom, "bkava") {
			return false
		}

		vaultValue, err := s.keeper.GetVaultTotalValue(ctx, record.TotalShares.Denom)
		if err != nil {
			iterErr = err
			return false
		}

		totalValue = totalValue.Add(vaultValue.Amount)

		return false
	})

	if iterErr != nil {
		return nil, iterErr
	}

	return &types.QueryVaultResponse{
		Vault: types.VaultResponse{
			Denom:             "bkava",
			Strategies:        allowedVault.Strategies,
			IsPrivateVault:    allowedVault.IsPrivateVault,
			AllowedDepositors: addressSliceToStringSlice(allowedVault.AllowedDepositors),
			// Empty for shares, as adding up all shares is not useful information
			TotalShares: "0",
			TotalValue:  totalValue,
		},
	}, nil
}

// Deposits implements the gRPC service handler for querying x/earn deposits.
func (s queryServer) Deposits(
	ctx context.Context,
	req *types.QueryDepositsRequest,
) (*types.QueryDepositsResponse, error) {
	if req == nil {
		return nil, status.Errorf(codes.InvalidArgument, "empty request")
	}

	sdkCtx := sdk.UnwrapSDKContext(ctx)

	// 1. 1 account and "bkava" vault
	if req.Depositor != "" && req.Denom == "bkava" {
		return s.get1AccountBkavaVaultDeposit(sdkCtx, req)
	}

	// 1. 1 account and 1 vault
	if req.Depositor != "" && req.Denom != "" {
		return s.get1Account1VaultDeposit(sdkCtx, req)
	}

	// 2. All accounts, "bkava" vault
	if req.Depositor == "" && req.Denom == "bkava" {
		return s.getBkavaVaultAllDeposits(sdkCtx, req)
	}

	// 2. All accounts, 1 vault
	if req.Depositor == "" && req.Denom != "" {
		return s.get1VaultAllDeposits(sdkCtx, req)
	}

	// 3. 1 account, all vaults
	if req.Depositor != "" && req.Denom == "" {
		return s.get1AccountAllDeposits(sdkCtx, req)
	}

	// 4. All accounts, all vaults
	return s.getAllDeposits(sdkCtx, req)
}

// get1Account1VaultDeposit returns deposits for a specific vault and a specific
// account
func (s queryServer) get1Account1VaultDeposit(
	ctx sdk.Context,
	req *types.QueryDepositsRequest,
) (*types.QueryDepositsResponse, error) {
	depositor, err := sdk.AccAddressFromBech32(req.Depositor)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "Invalid address")
	}

	shareRecord, found := s.keeper.GetVaultShareRecord(ctx, depositor)
	if !found {
		return nil, status.Error(codes.NotFound, "No deposit found for owner")
	}

	// Only requesting the value of the specified denom
	value, err := s.keeper.GetVaultAccountValue(ctx, req.Denom, depositor)
	if err != nil {
		return nil, status.Error(codes.NotFound, err.Error())
	}

	return &types.QueryDepositsResponse{
		Deposits: []types.DepositResponse{
			{
				Depositor: depositor.String(),
				// Only respond with requested denom shares
				Shares: types.NewVaultShares(
					types.NewVaultShare(req.Denom, shareRecord.Shares.AmountOf(req.Denom)),
				),
				Value: sdk.NewCoins(value),
			},
		},
		Pagination: nil,
	}, nil
}

// get1AccountBkavaVaultDeposit returns deposits for the aggregated bkava vault
// and a specific account
func (s queryServer) get1AccountBkavaVaultDeposit(
	ctx sdk.Context,
	req *types.QueryDepositsRequest,
) (*types.QueryDepositsResponse, error) {
	depositor, err := sdk.AccAddressFromBech32(req.Depositor)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "Invalid address")
	}

	shareRecord, found := s.keeper.GetVaultShareRecord(ctx, depositor)
	if !found {
		return nil, status.Error(codes.NotFound, "No deposit found for owner")
	}

	// Get all account deposit values to add up bkava
	totalAccountValue, err := getAccountTotalValue(ctx, s.keeper, depositor, shareRecord.Shares)
	if err != nil {
		return nil, err
	}

	// Use account value with only the aggregate bkava
	bkavaValue := getTotalBkava(totalAccountValue)

	return &types.QueryDepositsResponse{
		Deposits: []types.DepositResponse{
			{
				Depositor: depositor.String(),
				// Only respond with requested denom shares
				Shares: types.NewVaultShares(
					types.NewVaultShare(req.Denom, shareRecord.Shares.AmountOf(req.Denom)),
				),
				Value: sdk.NewCoins(bkavaValue),
			},
		},
		Pagination: nil,
	}, nil
}

// get1VaultAllDeposits returns all deposits for a specific vault
func (s queryServer) get1VaultAllDeposits(
	ctx sdk.Context,
	req *types.QueryDepositsRequest,
) (*types.QueryDepositsResponse, error) {
	_, found := s.keeper.GetVaultRecord(ctx, req.Denom)
	if !found {
		return nil, status.Error(codes.NotFound, "Vault record for denom not found")
	}

	deposits := []types.DepositResponse{}
	store := prefix.NewStore(ctx.KVStore(s.keeper.key), types.VaultShareRecordKeyPrefix)

	pageRes, err := query.FilteredPaginate(
		store,
		req.Pagination,
		func(key []byte, value []byte, accumulate bool) (bool, error) {
			var record types.VaultShareRecord
			err := s.keeper.cdc.Unmarshal(value, &record)
			if err != nil {
				return false, err
			}

			// Only those that have amount of requested denom
			if record.Shares.AmountOf(req.Denom).IsZero() {
				// inform paginate that there was no match on this key
				return false, nil
			}

			if accumulate {
				// Only get the value for the requested vault denom
				vaultValue, err := s.keeper.GetVaultAccountValue(ctx, req.Denom, record.Depositor)
				if err != nil {
					return false, err
				}

				// only add to results if paginate tells us to
				deposits = append(deposits, types.DepositResponse{
					Depositor: record.Depositor.String(),
					// Only the specified shares of requested denom
					Shares: types.NewVaultShares(
						types.NewVaultShare(req.Denom, record.Shares.AmountOf(req.Denom)),
					),
					Value: sdk.NewCoins(vaultValue),
				})
			}

			// inform paginate that were was a match on this key
			return true, nil
		},
	)

	if err != nil {
		return nil, err
	}

	return &types.QueryDepositsResponse{
		Deposits:   deposits,
		Pagination: pageRes,
	}, nil
}

// getBkavaVaultAllDeposits returns all deposits for the aggregated bkava vault
func (s queryServer) getBkavaVaultAllDeposits(
	ctx sdk.Context,
	req *types.QueryDepositsRequest,
) (*types.QueryDepositsResponse, error) {
	deposits := []types.DepositResponse{}
	store := prefix.NewStore(ctx.KVStore(s.keeper.key), types.VaultShareRecordKeyPrefix)

	pageRes, err := query.FilteredPaginate(
		store,
		req.Pagination,
		func(key []byte, value []byte, accumulate bool) (bool, error) {
			var record types.VaultShareRecord
			err := s.keeper.cdc.Unmarshal(value, &record)
			if err != nil {
				return false, err
			}

			// Only those that have bkava deposits
			if !vaultSharesContainBkava(record.Shares) {
				// inform paginate that there was no match on this key
				return false, nil
			}

			if accumulate {
				// Get total value for all bkava
				totalAccountValue, err := getAccountTotalValue(ctx, s.keeper, record.Depositor, record.Shares)
				if err != nil {
					return false, err
				}

				bkavaTotal := getTotalBkava(totalAccountValue)

				// only add to results if paginate tells us to
				deposits = append(deposits, types.DepositResponse{
					Depositor: record.Depositor.String(),
					// Only the specified shares of requested denom
					Shares: nil,
					Value:  sdk.NewCoins(bkavaTotal),
				})
			}

			// inform paginate that were was a match on this key
			return true, nil
		},
	)

	if err != nil {
		return nil, err
	}

	return &types.QueryDepositsResponse{
		Deposits:   deposits,
		Pagination: pageRes,
	}, nil
}

// get1AccountAllDeposits returns deposits for all vaults for a specific account
func (s queryServer) get1AccountAllDeposits(
	ctx sdk.Context,
	req *types.QueryDepositsRequest,
) (*types.QueryDepositsResponse, error) {
	depositor, err := sdk.AccAddressFromBech32(req.Depositor)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "Invalid address")
	}

	deposits := []types.DepositResponse{}

	accountShare, found := s.keeper.GetVaultShareRecord(ctx, depositor)
	if !found {
		return nil, status.Error(codes.NotFound, "No deposit found for depositor")
	}

	value, err := getAccountTotalValue(ctx, s.keeper, depositor, accountShare.Shares)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	deposits = append(deposits, types.DepositResponse{
		Depositor: depositor.String(),
		Shares:    accountShare.Shares,
		Value:     value,
	})

	return &types.QueryDepositsResponse{
		Deposits:   deposits,
		Pagination: nil,
	}, nil
}

// getAllDeposits returns all deposits for all vaults
func (s queryServer) getAllDeposits(
	ctx sdk.Context,
	req *types.QueryDepositsRequest,
) (*types.QueryDepositsResponse, error) {
	deposits := []types.DepositResponse{}
	store := prefix.NewStore(ctx.KVStore(s.keeper.key), types.VaultShareRecordKeyPrefix)

	pageRes, err := query.Paginate(
		store,
		req.Pagination,
		func(key []byte, value []byte) error {
			var record types.VaultShareRecord
			err := s.keeper.cdc.Unmarshal(value, &record)
			if err != nil {
				return err
			}

			accValue, err := getAccountTotalValue(ctx, s.keeper, record.Depositor, record.Shares)
			if err != nil {
				return err
			}

			// only add to results if paginate tells us to
			deposits = append(deposits, types.DepositResponse{
				Depositor: record.Depositor.String(),
				Shares:    record.Shares,
				Value:     accValue,
			})

			return nil
		},
	)

	if err != nil {
		return nil, err
	}

	return &types.QueryDepositsResponse{
		Deposits:   deposits,
		Pagination: pageRes,
	}, nil
}

// getAccountTotalValue returns the total value for all vaults for a specific
// account based on their shares.
func getAccountTotalValue(
	ctx sdk.Context,
	keeper Keeper,
	account sdk.AccAddress,
	shares types.VaultShares,
) (sdk.Coins, error) {
	value := sdk.NewCoins()

	for _, share := range shares {
		accValue, err := keeper.GetVaultAccountValue(ctx, share.Denom, account)
		if err != nil {
			return nil, err
		}

		value = value.Add(sdk.NewCoin(share.Denom, accValue.Amount))
	}

	return value, nil
}

func addressSliceToStringSlice(addresses []sdk.AccAddress) []string {
	var strings []string
	for _, address := range addresses {
		strings = append(strings, address.String())
	}

	return strings
}

func vaultSharesContainBkava(shares types.VaultShares) bool {
	for _, share := range shares {
		if strings.HasPrefix(share.Denom, "bkava") {
			return true
		}
	}

	return false
}

func getTotalBkava(coins sdk.Coins) sdk.Coin {
	bkavaTotal := sdk.NewCoin("bkava", sdk.ZeroInt())

	for _, coin := range coins {
		if strings.HasPrefix(coin.Denom, "bkava") {
			bkavaTotal = bkavaTotal.AddAmount(coin.Amount)
		}
	}

	return bkavaTotal
}
