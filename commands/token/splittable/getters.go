package splittable

import (
	"github.com/spf13/cobra"

	"github.com/tzero-dev/go-t0ken/cli"
	"github.com/tzero-dev/go-t0ken/commands/token"
	"github.com/tzero-dev/go-t0ken/contracts/token/erc20"
)

var GetterCommands = []*cobra.Command{
	&cobra.Command{
		Use:     "abi",
		Short:   "Outputs the Splittable-T0ken ABI",
		Example: "t0ken token-splittable investor abi",
		Args:    cobra.NoArgs,
		Run:     func(cmd *cobra.Command, args []string) { cmd.Println(erc20.T0kenABI) },
	},
	&cobra.Command{
		Use:     "bin",
		Short:   "Outputs the Splittable-T0ken Binary",
		Example: "t0ken token-splittable investor bin",
		Args:    cobra.NoArgs,
		Run:     func(cmd *cobra.Command, args []string) { cmd.Println(erc20.T0kenBin) },
	},
	&cobra.Command{
		Use:     "split",
		Short:   "Gets the split value",
		Example: "t0ken token-splittable",
		Args:    cobra.NoArgs,
		PreRun:  connectCaller,
		Run: func(cmd *cobra.Command, args []string) {
			cli.CheckGetter(cmd)(callSession.Split())
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
