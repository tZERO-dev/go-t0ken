package dynamicSplittable

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/spf13/cobra"

	"github.com/tzero-dev/go-t0ken/cli"
	"github.com/tzero-dev/go-t0ken/commands/token"
	t "github.com/tzero-dev/go-t0ken/contracts/token"
)

var GetterCommands = []*cobra.Command{
	&cobra.Command{
		Use:     "abi",
		Short:   "Outputs the DynamicSplittable-T0ken ABI",
		Example: "t0ken token-dynamicSplittable investor abi",
		Args:    cobra.NoArgs,
		Run:     func(cmd *cobra.Command, args []string) { cmd.Println(t.T0kenABI) },
	},
	&cobra.Command{
		Use:     "bin",
		Short:   "Outputs the DynamicSplittable-T0ken Binary",
		Example: "t0ken token-dynamicSplittable investor bin",
		Args:    cobra.NoArgs,
		Run:     func(cmd *cobra.Command, args []string) { cmd.Println(t.T0kenBin) },
	},
	&cobra.Command{
		Use:     "holdersSplit <holder>",
		Short:   "Gets the <holder> split value",
		Example: "t0ken token-dynamicSplittable predecessor",
		Args:    cli.AddressArgFunc("holder", 0),
		PreRun:  connectCaller,
		Run: func(cmd *cobra.Command, args []string) {
			holder := common.HexToAddress(args[0])
			cli.CheckGetter(cmd)(callSession.HoldersSplit(holder))
		},
	},
	&cobra.Command{
		Use:     "splits <index>",
		Short:   "Gets the split for the given <index>",
		Example: "t0ken token-dynamicSplittable sucessor",
		Args:    cli.BigIntArgFunc("index", 0),
		PreRun:  connectCaller,
		Run: func(cmd *cobra.Command, args []string) {
			index, _ := new(big.Int).SetString(args[0], 10)
			cli.CheckGetter(cmd)(callSession.Splits(index))
		},
	},
	&cobra.Command{
		Use:     "currentSplitIndex",
		Short:   "Gets the current split index",
		Example: "t0ken token-dynamicSplittable currentSplitIndex",
		Args:    cobra.NoArgs,
		PreRun:  connectCaller,
		Run: func(cmd *cobra.Command, args []string) {
			cli.CheckGetter(cmd)(callSession.CurrentSplitIndex())
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
