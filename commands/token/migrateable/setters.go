package migrateable

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/spf13/cobra"

	"github.com/tzero-dev/go-t0ken/cli"
	"github.com/tzero-dev/go-t0ken/commands/gas"
	"github.com/tzero-dev/go-t0ken/commands/nonce"
	"github.com/tzero-dev/go-t0ken/commands/token"
)

var SetterCommands = []*cobra.Command{
	&cobra.Command{
		Use:     "setPredecessor <token>",
		Short:   "Sets the tokens predecessor to the <token>",
		Example: "t0ken token-migrateable setPredecessor 0x693944c37ef055864bc070e159ed1f1a719580c7 --keystoreAddress issuer",
		Args:    cli.ChainArgs(cobra.MaximumNArgs(1), cli.AddressArgFunc("token", 0)),
		PreRun:  connectTransactor,
		Run: func(cmd *cobra.Command, args []string) {
			predecessor := common.HexToAddress(args[0])
			cli.PrintTransactionFn(cmd)(transSession.SetPredecessor(predecessor))
		},
	},
	&cobra.Command{
		Use:     "setSuccessor <token>",
		Short:   "Sets the tokens successor to the <token>",
		Example: "t0ken token-migrateable setSuccessor 0x693944c37ef055864bc070e159ed1f1a719580c7 --keystoreAddress issuer",
		Args:    cli.ChainArgs(cobra.MaximumNArgs(1), cli.AddressArgFunc("token", 0)),
		PreRun:  connectTransactor,
		Run: func(cmd *cobra.Command, args []string) {
			successor := common.HexToAddress(args[0])
			cli.PrintTransactionFn(cmd)(transSession.SetSuccessor(successor))
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
