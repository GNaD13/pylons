package v4_test

import (
	"testing"

	"cosmossdk.io/math"
	"github.com/Pylons-tech/pylons/app/apptesting"
	v4 "github.com/Pylons-tech/pylons/app/upgrades/v4"
	sdk "github.com/cosmos/cosmos-sdk/types"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	"github.com/stretchr/testify/suite"
)

const (
	stakingCoinDenom = "ubedrock"
	stripeCoinDenom  = "ustripeusd"
	fundAmount       = 10_000_000
)

var (
	defaultAcctFunds sdk.Coins = sdk.NewCoins(
		sdk.NewCoin(stakingCoinDenom, sdk.NewInt(fundAmount)),
	)
	defaultAcctFundsStripeCoin sdk.Coins = sdk.NewCoins(
		sdk.NewCoin(stripeCoinDenom, sdk.NewInt(fundAmount)),
	)
)

type UpgradeTestSuite struct {
	apptesting.KeeperTestHelper
}

func TestUpgradeTestSuite(t *testing.T) {
	s := new(UpgradeTestSuite)
	suite.Run(t, s)
}

func (suite *UpgradeTestSuite) TestBurnToken_Ubedrock() {
	suite.Setup()
	// Fund ubedrock to test account
	for _, acc := range suite.TestAccs {
		suite.FundAcc(acc, defaultAcctFunds)
	}
	// Check genesis delegation
	delegations := suite.App.StakingKeeper.GetAllDelegations(suite.Ctx)
	suite.Require().Equal(1, len(delegations))
	bondedAmount := suite.App.StakingKeeper.GetDelegatorBonded(suite.Ctx, sdk.MustAccAddressFromBech32(delegations[0].DelegatorAddress))
	suite.Require().Equal(bondedAmount, math.NewInt(1_000_000))
	// Create new delegation
	valAddress, err := sdk.ValAddressFromBech32(delegations[0].ValidatorAddress)
	val, found := suite.App.StakingKeeper.GetValidator(suite.Ctx, valAddress)
	suite.Require().True(found)
	suite.Require().NoError(err)
	_, err = suite.App.StakingKeeper.Delegate(
		suite.Ctx,
		suite.TestAccs[0],
		math.NewInt(1_000_000),
		stakingtypes.Unbonded,
		val,
		true,
	)
	suite.Require().NoError(err)
	// Check all delegation
	delegations = suite.App.StakingKeeper.GetAllDelegations(suite.Ctx)
	suite.Require().Equal(2, len(delegations))
	bondedModuleAddress := suite.App.AccountKeeper.GetModuleAddress(stakingtypes.BondedPoolName)
	bondedModuleAmount := suite.App.BankKeeper.GetBalance(suite.Ctx, bondedModuleAddress, stakingCoinDenom)
	suite.Require().Equal(bondedModuleAmount.Amount, math.NewInt(2_000_000))
	// Get ubedrock total supply
	totalAmount := suite.App.BankKeeper.GetSupply(suite.Ctx, stakingCoinDenom)
	suite.Require().Equal(totalAmount.Amount, math.NewInt(31_000_000))
	// Burn ubedrock
	bankBaseKeeper, _ := suite.App.BankKeeper.(bankkeeper.BaseKeeper)
	v4.BurnToken(suite.Ctx, stakingCoinDenom, &suite.App.AccountKeeper, &bankBaseKeeper, &suite.App.StakingKeeper)
	// Check ubedrock total supply (should equal 0)
	totalAmount = suite.App.BankKeeper.GetSupply(suite.Ctx, stakingCoinDenom)
	suite.Require().Equal(totalAmount.Amount, math.ZeroInt())
}

func (suite *UpgradeTestSuite) TestBurnToken_Ustripeusd() {
	suite.Setup()
	// Fund Ustripeusd to test account
	for _, acc := range suite.TestAccs {
		suite.FundAcc(acc, defaultAcctFundsStripeCoin)
	}
	// Get Ustripeusd total supply
	totalAmount := suite.App.BankKeeper.GetSupply(suite.Ctx, stripeCoinDenom)
	suite.Require().Equal(totalAmount.Amount, math.NewInt(30_000_000))
	// Burn Ustripeusd
	bankBaseKeeper, _ := suite.App.BankKeeper.(bankkeeper.BaseKeeper)
	v4.BurnToken(suite.Ctx, stripeCoinDenom, &suite.App.AccountKeeper, &bankBaseKeeper, &suite.App.StakingKeeper)
	// Check Ustripeusd total supply (should equal 0)
	totalAmount = suite.App.BankKeeper.GetSupply(suite.Ctx, stripeCoinDenom)
	suite.Require().Equal(totalAmount.Amount, math.ZeroInt())
}

func (suite *UpgradeTestSuite) TestMintUbedrockForInitialAccount() {
	suite.Setup()

	// Burn ubedrock
	bankBaseKeeper, _ := suite.App.BankKeeper.(bankkeeper.BaseKeeper)
	v4.BurnToken(suite.Ctx, stakingCoinDenom, &suite.App.AccountKeeper, &bankBaseKeeper, &suite.App.StakingKeeper)
	totalAmount := suite.App.BankKeeper.GetSupply(suite.Ctx, stakingCoinDenom)
	suite.Require().Equal(totalAmount.Amount, math.ZeroInt())

	// Mint ubedrock for initial account
	v4.MintUbedrockForInitialAccount(suite.Ctx, &bankBaseKeeper, &suite.App.StakingKeeper)
	totalAmount = suite.App.BankKeeper.GetSupply(suite.Ctx, stakingCoinDenom)
	suite.Require().Equal(totalAmount.Amount, math.NewIntFromUint64(1_000_000_000_000_000))
	// Check vals - each vals should have 1 bedrock = 1_000_000 ubedrock
	vals := suite.App.StakingKeeper.GetAllValidators(suite.Ctx)
	for _, val := range vals {
		token := val.GetTokens()
		suite.Require().Equal(math.NewIntFromUint64(1_000_000), token)
	}
	// Check token in all initial account
	totalDelegation := len(vals)
	totalDelegateAmount := math.NewIntFromUint64(1_000_000).MulRaw(int64(totalDelegation))
	for _, acc := range v4.Accounts {
		if acc == v4.EngineHotWal {
			accAmount := suite.App.BankKeeper.GetBalance(suite.Ctx, sdk.MustAccAddressFromBech32(acc), stakingCoinDenom)
			amount := v4.UbedrockDistribute[acc].Sub(totalDelegateAmount)
			suite.Require().Equal(accAmount.Amount, amount)
			continue
		}
		accAmount := suite.App.BankKeeper.GetBalance(suite.Ctx, sdk.MustAccAddressFromBech32(acc), stakingCoinDenom)
		suite.Require().Equal(accAmount.Amount, v4.UbedrockDistribute[acc])
	}
	// Check token in bonded pool module, should equal total delegate amount
	bondedModuleAddress := suite.App.AccountKeeper.GetModuleAddress(stakingtypes.BondedPoolName)
	bondedModuleAmount := suite.App.BankKeeper.GetBalance(suite.Ctx, bondedModuleAddress, stakingCoinDenom)
	suite.Require().Equal(bondedModuleAmount.Amount, totalDelegateAmount)
}
