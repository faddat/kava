package keeper_test

import (
	"math/big"
	"testing"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	vestingtypes "github.com/cosmos/cosmos-sdk/x/auth/vesting/types"
	"github.com/stretchr/testify/suite"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/tharsis/ethermint/tests"
	etherminttypes "github.com/tharsis/ethermint/types"
	"github.com/tharsis/ethermint/x/evm/statedb"
	"github.com/tharsis/ethermint/x/evm/types"

	"github.com/kava-labs/kava/x/evmutil/testutil"
)

type evmKeeperTestSuite struct {
	testutil.Suite
}

func (suite *evmKeeperTestSuite) SetupTest() {
	suite.Suite.SetupTest()
}

func (suite *evmKeeperTestSuite) StateDB() *statedb.StateDB {
	return statedb.New(suite.Ctx, suite.App.GetEvmKeeper(), statedb.NewEmptyTxConfig(common.BytesToHash(suite.Ctx.HeaderHash().Bytes())))
}

func (suite *evmKeeperTestSuite) TestEvmKeeper_SetAccount() {
	baseAddr := tests.GenerateAddress()
	baseAcc := &authtypes.BaseAccount{Address: sdk.AccAddress(baseAddr.Bytes()).String()}
	ethAddr := tests.GenerateAddress()
	ethAcc := &etherminttypes.EthAccount{BaseAccount: &authtypes.BaseAccount{Address: sdk.AccAddress(ethAddr.Bytes()).String()}, CodeHash: common.BytesToHash(types.EmptyCodeHash).String()}
	vestingAddr := tests.GenerateAddress()
	vestingAcc := vestingtypes.NewBaseVestingAccount(&authtypes.BaseAccount{Address: sdk.AccAddress(vestingAddr.Bytes()).String()}, sdk.NewCoins(), time.Now().Unix())

	testCases := []struct {
		name        string
		address     common.Address
		account     statedb.Account
		expectedErr error
	}{
		{
			"new account, non-contract account",
			tests.GenerateAddress(),
			statedb.Account{10, big.NewInt(100), types.EmptyCodeHash},
			nil,
		},
		{
			"new account, contract account",
			tests.GenerateAddress(),
			statedb.Account{10, big.NewInt(100), crypto.Keccak256Hash([]byte("some code hash")).Bytes()},
			nil,
		},
		{
			"existing eth account, non-contract account",
			ethAddr,
			statedb.Account{10, big.NewInt(1), types.EmptyCodeHash},
			nil,
		},
		{
			"existing eth account, contract account",
			ethAddr,
			statedb.Account{10, big.NewInt(0), crypto.Keccak256Hash([]byte("some code hash")).Bytes()},
			nil,
		},
		{
			"existing base account, non-contract account",
			baseAddr,
			statedb.Account{10, big.NewInt(10), types.EmptyCodeHash},
			nil,
		},
		{
			"existing base account, contract account",
			baseAddr,
			statedb.Account{10, big.NewInt(99), crypto.Keccak256Hash([]byte("some code hash")).Bytes()},
			nil,
		},
		{
			"existing vesting account, non-contract account",
			vestingAddr,
			statedb.Account{10, big.NewInt(1000), types.EmptyCodeHash},
			nil,
		},
		{
			"existing vesting account, contract account",
			vestingAddr,
			statedb.Account{10, big.NewInt(1001), crypto.Keccak256Hash([]byte("some code hash")).Bytes()},
			types.ErrInvalidAccount,
		},
	}

	for _, tc := range testCases {
		suite.Run(tc.name, func() {
			suite.SetupTest()

			if tc.address == baseAddr {
				suite.AccountKeeper.SetAccount(suite.Ctx, baseAcc)
			}
			if tc.address == ethAddr {
				suite.AccountKeeper.SetAccount(suite.Ctx, ethAcc)
			}
			if tc.address == vestingAddr {
				suite.AccountKeeper.SetAccount(suite.Ctx, vestingAcc)
			}

			vmdb := suite.StateDB()
			err := vmdb.Keeper().SetAccount(suite.Ctx, tc.address, tc.account)

			if tc.expectedErr == nil {
				suite.Require().NoError(err)
			} else {
				suite.Require().Error(tc.expectedErr)
				return
			}

			nonce := vmdb.GetNonce(tc.address)
			suite.Equal(nonce, tc.account.Nonce, "expected nonce to be set")

			hash := vmdb.GetCodeHash(tc.address)
			suite.Equal(common.BytesToHash(tc.account.CodeHash), hash, "expected code hash to be set")

			balance := vmdb.GetBalance(tc.address)
			suite.Equal(balance, tc.account.Balance, "expected balance to be set")
		})
	}
}

func TestEvmKeeperTestSuite(t *testing.T) {
	suite.Run(t, new(evmKeeperTestSuite))
}
