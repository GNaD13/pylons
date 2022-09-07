package cli_test

import (
	"testing"

	"github.com/Pylons-tech/pylons/x/pylons/client/cli"
	"github.com/Pylons-tech/pylons/x/pylons/types"
	clitestutil "github.com/cosmos/cosmos-sdk/testutil/cli"
	"github.com/stretchr/testify/require"
)

func TestCmdListRecipesByCookbook(t *testing.T) {
	net, recipe := settingRecipeObject(t)
	ctx := net.Validators[0].ClientCtx

	for _, tc := range []struct {
		desc       string
		cookbookId string
		id         string
		shouldErr  bool
	}{
		{
			desc:       "Valid",
			cookbookId: recipe.CookbookId,
			id:         recipe.Id,
			shouldErr:  false,
		},
		{
			desc:       "Invalid - cookbookId",
			cookbookId: "Invalid",
			id:         recipe.Id,
			shouldErr:  true,
		},
	} {
		tc := tc
		t.Run(tc.desc, func(t *testing.T) {
			args := []string{}
			args = append(args, tc.cookbookId)
			out, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdListRecipesByCookbook(), args)
			if tc.shouldErr {
				var resp types.QueryListRecipesByCookbookResponse
				require.NoError(t, net.Config.Codec.UnmarshalJSON(out.Bytes(), &resp))
				require.Equal(t, len(resp.Recipes), 0)
			} else {
				require.NoError(t, err)
				var resp types.QueryListRecipesByCookbookResponse
				require.NoError(t, net.Config.Codec.UnmarshalJSON(out.Bytes(), &resp))
				require.NotNil(t, resp.Recipes[0])
				require.Equal(t, resp.Recipes[0].CookbookId, tc.cookbookId)
				require.Equal(t, resp.Recipes[0].Id, tc.id)
			}
		})
	}
}
