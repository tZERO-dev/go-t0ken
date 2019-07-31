package migrateable

import (
	//"github.com/ethereum/go-ethereum/common"
	"github.com/spf13/cobra"

	"github.com/tzero-dev/go-t0ken/cli"
	"github.com/tzero-dev/go-t0ken/commands/token"
	"github.com/tzero-dev/go-t0ken/contracts/token/erc20"
)

var GetterCommands = []*cobra.Command{
	&cobra.Command{
		Use:     "abi",
		Short:   "Outputs the Migrateable-T0ken ABI",
		Example: "t0ken token-migrateable investor abi",
		Args:    cobra.NoArgs,
		Run:     func(cmd *cobra.Command, args []string) { cmd.Println(erc20.T0kenABI) },
	},
	&cobra.Command{
		Use:     "bin",
		Short:   "Outputs the Migrateable-T0ken Binary",
		Example: "t0ken token-migrateable investor bin",
		Args:    cobra.NoArgs,
		Run:     func(cmd *cobra.Command, args []string) { cmd.Println(erc20.T0kenBin) },
	},
	&cobra.Command{
		Use:     "predecessor",
		Short:   "Gets the tokens predecessor",
		Example: "t0ken token-migrateable predecessor",
		Args:    cobra.NoArgs,
		PreRun:  connectCaller,
		Run: func(cmd *cobra.Command, args []string) {
			cli.CheckGetter(cmd)(callSession.Predecessor())
		},
	},
	&cobra.Command{
		Use:     "successor",
		Short:   "Gets the tokens sucessor",
		Example: "t0ken token-migrateable sucessor",
		Args:    cobra.NoArgs,
		PreRun:  connectCaller,
		Run: func(cmd *cobra.Command, args []string) {
			cli.CheckGetter(cmd)(callSession.Successor())
		},
	},
}

func init() {
	for i, cmd := range GetterCommands {
		// Skip ABI/Bin
		if i < 2 {
			continue
		}

		// Allow providing contract 'address' flag
		cmd.Flags().String("address", "", `address of the token contract (default "[`+contractKey+`] value from config")`)
	}

	for _, cmd := range token.GetterCommands {
		if cmd.Use == "abi" || cmd.Use == "bin" {
			continue
		}
		GetterCommands = append(GetterCommands, cmd)
	}
}
