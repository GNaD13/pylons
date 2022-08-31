package cli_test

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"testing"

	util "github.com/Pylons-tech/pylons/testutil/cli"
	"github.com/Pylons-tech/pylons/testutil/network"
	"github.com/Pylons-tech/pylons/x/pylons/client/cli"
	"github.com/Pylons-tech/pylons/x/pylons/types"
	"github.com/btcsuite/btcutil/base58"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	clitestutil "github.com/cosmos/cosmos-sdk/testutil/cli"
	"github.com/stretchr/testify/require"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func settingGenesis(t *testing.T) (client.Context, []types.UserMap, *types.Item) {
	t.Helper()
	cfg := network.DefaultConfig()
	state := types.GenesisState{}
	net := network.New(t, cfg)
	ctx := net.Validators[0].ClientCtx
	require.NoError(t, cfg.Codec.UnmarshalJSON(cfg.GenesisState[types.ModuleName], &state))
	// setup account
	accountAddress := types.GenTestBech32List(3)
	address, err := util.GenerateAddressWithAccount(ctx, t, net)
	cookbookCreator := types.UserMap{
		AccountAddr: accountAddress[0],
		Username:    "cookbookCreator",
	}
	state.AccountList = append(state.AccountList, cookbookCreator)

	itemSender := types.UserMap{
		AccountAddr: address,
		Username:    "itemSender",
	}
	state.AccountList = append(state.AccountList, itemSender)

	itemReceiver := types.UserMap{
		AccountAddr: accountAddress[2],
		Username:    "itemReceiver",
	}
	state.AccountList = append(state.AccountList, itemReceiver)
	// setup cookbook
	cookbook := types.Cookbook{
		Creator: cookbookCreator.AccountAddr,
		Id:      "testCookbook",
	}
	state.CookbookList = append(state.CookbookList, cookbook)
	// setup item
	coin := []sdk.Coin{sdk.NewCoin(types.PylonsCoinDenom, sdk.NewInt(0))}
	id := make([]byte, 8)
	binary.LittleEndian.PutUint64(id, 10)
	item := types.Item{
		Owner:       itemSender.AccountAddr,
		CookbookId:  cookbook.Id,
		Id:          base58.Encode(id),
		Tradeable:   true,
		TransferFee: coin,
	}
	state.ItemList = append(state.ItemList, item)
	// setup param
	itemTransferFeePercentage, _ := sdk.NewDecFromStr("0.0")
	param := types.DefaultParams()
	param.MinTransferFee = sdk.ZeroInt()
	param.MaxTransferFee = sdk.ZeroInt()
	param.ItemTransferFeePercentage = itemTransferFeePercentage
	state.Params = param
	buffer, err := cfg.Codec.MarshalJSON(&state)
	require.NoError(t, err)
	cfg.GenesisState[types.ModuleName] = buffer

	return ctx, state.AccountList, &item
}

func TestCmdSendItems(t *testing.T) {
	ctx, userMaps, item := settingGenesis(t)

	common := []string{
		fmt.Sprintf("--%s=%s", flags.FlagFrom, userMaps[1].AccountAddr),
		fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
		fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastBlock),
	}

	itemRefs := make([]types.ItemRef, 1)

	itemRefs[0] = types.ItemRef{
		CookbookId: item.CookbookId,
		ItemId:     item.Id,
	}

	for _, tc := range []struct {
		desc      string
		receiver  string
		itemRef   []types.ItemRef
		shouldErr bool
	}{
		{
			desc:      "Valid",
			receiver:  userMaps[2].AccountAddr,
			itemRef:   itemRefs,
			shouldErr: false,
		},
	} {
		tc := tc
		t.Run(tc.desc, func(t *testing.T) {
			args := []string{}
			itemRefByte, _ := json.Marshal(tc.itemRef)
			args = append(args, tc.receiver)
			args = append(args, string(itemRefByte))
			args = append(args, common...)
			_, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdSendItems(), args)
			if tc.shouldErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}
