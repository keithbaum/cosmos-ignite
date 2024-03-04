package cli

import (
	"fmt"
	lightclientmoduletypes "foochain/x/lightclient/types"
	"strconv"
	"strings"
	"time"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/tx"
	// "github.com/cosmos/cosmos-sdk/client/flags"
	"foochain/x/foo/types"
)

var (
	DefaultRelativePacketTimeoutTimestamp = uint64((time.Duration(10) * time.Minute).Nanoseconds())
)

const (
	flagPacketTimeoutTimestamp = "packet-timeout-timestamp"
	listSeparator              = ","
)

func GetTxCmd() *cobra.Command {
	shortMsg := "%s module transaction commands"
	fooTxCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf(shortMsg, types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE: func(cmd *cobra.Command, args []string) error {
			usageTemplate := `Usage:{{if .HasAvailableSubCommands}}
  {{.CommandPath}} [command]{{end}}
  
{{if .HasAvailableSubCommands}}Available Commands:{{range .Commands}}{{if .IsAvailableCommand}}
  {{rpad .Name .NamePadding }} {{.Short}}{{end}}{{end}}
Use "{{.CommandPath}} [command] --help" for more information about a command.{{end}}
`
			cmd.SetUsageTemplate(usageTemplate)
			return cmd.Help()
		},
	}

	fooTxCmd.AddCommand(
		NewVerifyTxCmd(),
	)
	// this line is used by starport scaffolding # 1

	return fooTxCmd
}

func NewVerifyTxCmd() *cobra.Command {
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

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.MsgVerifyTxRequest{}
			msg.TxHash = txHash
			msg.BlockHeight = blockHeight
			msg.TxData = &lightclientmoduletypes.TxData{
				TxIdx: txIdx,
			}
			msg.Proof = &lightclientmoduletypes.Proof{
				MerkleRoot: merkleRoot,
				MerklePath: merklePath,
			}

			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), &msg)
		},
	}

	return cmd
}
