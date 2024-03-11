package cli

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"

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

	// this line is used by starport scaffolding # 1

	return fooTxCmd
}
