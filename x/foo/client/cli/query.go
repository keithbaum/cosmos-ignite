package cli

import (
	"context"
	"fmt"
	lightclientmoduletypes "foochain/x/lightclient/types"
	"strconv"
	"strings"

	// "strings"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	// "github.com/cosmos/cosmos-sdk/client/flags"
	// sdk "github.com/cosmos/cosmos-sdk/types"

	"foochain/x/foo/types"
)

// GetQueryCmd returns the cli query commands for this module
func GetQueryCmd(queryRoute string) *cobra.Command {
	// Group foo queries under a subcommand
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("Querying commands for the %s module", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(CmdQueryParams(),
		VerifyTxCmd())
	// this line is used by starport scaffolding # 1

	return cmd
}

func VerifyTxCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "verify-proof [txHash] [blockHeight] [merkleRoot] [merklePath] [txIdx]",
		Short: "verify a proof",
		Long:  "verify a proof",
		Args:  cobra.ExactArgs(5),
		RunE: func(cmd *cobra.Command, args []string) error {
			txHash := args[0]
			blockHeight, _ := strconv.ParseInt(args[1], 10, 64)
			merkleRoot := args[2]
			merklePath := strings.Split(args[3], ",")
			txIdx, _ := strconv.ParseInt(args[4], 10, 64)

			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}
			queryClient := types.NewQueryClient(clientCtx)

			req := types.QueryVerifyTxRequest{}
			req.TxHash = txHash
			req.BlockHeight = blockHeight
			req.TxData = &lightclientmoduletypes.TxData{
				TxIdx: txIdx,
			}
			req.Proof = &lightclientmoduletypes.Proof{
				MerkleRoot: merkleRoot,
				MerklePath: merklePath,
			}

			res, err := queryClient.VerifyTx(context.Background(), &req)
			if err != nil {
				return err
			}
			return clientCtx.PrintProto(res)
		},
	}

	return cmd
}
