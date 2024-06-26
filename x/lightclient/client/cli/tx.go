package cli

import (
	"fmt"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"strconv"
	"time"

	"github.com/spf13/cobra"

	// "github.com/cosmos/cosmos-sdk/client/flags"
	"foochain/x/lightclient/types"
)

var (
	DefaultRelativePacketTimeoutTimestamp = uint64((time.Duration(10) * time.Minute).Nanoseconds())
)

const (
	flagPacketTimeoutTimestamp = "packet-timeout-timestamp"
	listSeparator              = ","
)

// GetTxCmd returns the transaction commands for this module
func GetTxCmd() *cobra.Command {
	//cmd := &cobra.Command{
	//	Use:                        types.ModuleName,
	//	Short:                      fmt.Sprintf("%s transactions subcommands", types.ModuleName),
	//	DisableFlagParsing:         true,
	//	SuggestionsMinimumDistance: 2,
	//	RunE:                       client.ValidateCmd,
	//}
	//
	//// this line is used by starport scaffolding # 1
	//
	//return cmd
	shortMsg := "%s module transaction commands"
	lightclientTxCmd := &cobra.Command{
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

	lightclientTxCmd.AddCommand(
		NewStoreHeadersCmd(),
	)
	// this line is used by starport scaffolding # 1

	return lightclientTxCmd

}

func NewStoreHeadersCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "store-tx-headers [blockHeight] [merkleRoot]",
		Short: "Store transaction header",
		Long:  "Store transaction header",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			blockHeight, _ := strconv.ParseInt(args[0], 10, 64)
			merkleRoot := args[1]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.MsgStoreExternalChain1TxHeadersRequest{}
			msg.Sender = clientCtx.GetFromAddress().String()
			msg.BlockHeight = blockHeight
			msg.MerkleRoot = merkleRoot

			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), &msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	return cmd
}
