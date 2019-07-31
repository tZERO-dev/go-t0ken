package dynamicSplittable

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/spf13/cobra"

	"github.com/tzero-dev/go-t0ken/cli"
	"github.com/tzero-dev/go-t0ken/commands/gas"
	"github.com/tzero-dev/go-t0ken/commands/nonce"
	"github.com/tzero-dev/go-t0ken/commands/token"
)

var SetterCommands = []*cobra.Command{
	&cobra.Command{
		Use:     "updateSplit <numerator> <denominator>",
		Short:   "Updates the tokens split value",
		Example: "t0ken token-dynamicSplittable updateSplit 3 5 --keystoreAddress issuer",
		Args:    cli.ChainArgs(cli.BigIntArgFunc("numverator", 0), cli.BigIntArgFunc("denominator", 1)),
		PreRun:  connectTransactor,
		Run: func(cmd *cobra.Command, args []string) {
			numerator, _ := new(big.Int).SetString(args[0], 10)
			denominator, _ := new(big.Int).SetString(args[1], 10)
			cli.PrintTransactionFn(cmd)(transSession.UpdateSplit(numerator, denominator))
		},
	},
	&cobra.Command{
		Use:     "splitTotalSupply <cachedOutQuantity>",
		Short:   "Processes the split based on the current multiplier, subtracting the <cachedOutQuantity>",
		Example: "t0ken token-dynamicSplittable splitTotalSupply 10 --keystoreAddress issuer",
		Args:    cli.BigIntArgFunc("cachedOutQuantity", 0),
		PreRun:  connectTransactor,
		Run: func(cmd *cobra.Command, args []string) {
			qty, _ := new(big.Int).SetString(args[0], 10)
			cli.PrintTransactionFn(cmd)(transSession.SplitTotalSupply(qty))
		},
	},
	&cobra.Command{
		Use:     "removeShareholder <holder>",
		Short:   "Removes the shareholder when a split results in a balance of zero",
		Example: "t0ken token-dynamicSplittable removeShareholder 0xf01ff29dcbee147e9ca151a281bfdf136f66a45b --keystoreAddress issuer",
		Args:    cli.AddressArgFunc("holder", 0),
		PreRun:  connectTransactor,
		Run: func(cmd *cobra.Command, args []string) {
			holder := common.HexToAddress(args[0])
			cli.PrintTransactionFn(cmd)(transSession.RemoveShareholder(holder))
		},
	},
}

func init() {
	for _, cmd := range SetterCommands {
		// Add the 'gasPrice', 'nonce' and 'wait' flags to the deploy function
		gas.Flag(cmd)
		nonce.Flag(cmd)
		cli.WaitFlag(cmd)

		// Allow providing contract 'address' flag
		cmd.Flags().String("address", "", `address of the token contract (default "[`+contractKey+`] value from config")`)
	}

	for _, cmd := range token.SetterCommands {
		if cmd.Use == "abi" || cmd.Use == "bin" {
			continue
		}
		SetterCommands = append(SetterCommands, cmd)
	}
}
