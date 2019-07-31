package dynamicSplittable

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/spf13/cobra"

	"github.com/tzero-dev/go-t0ken/cli"
	"github.com/tzero-dev/go-t0ken/commands"
	"github.com/tzero-dev/go-t0ken/commands/gas"
	"github.com/tzero-dev/go-t0ken/commands/nonce"
	"github.com/tzero-dev/go-t0ken/contracts/token/erc20"
)

var (
	Command = &cobra.Command{
		Use:   "token-dynamicSplittable",
		Short: "DynamicSplittable-T0ken utilities",
	}

	DeployCommand = &cobra.Command{
		Use:     "deploy <name> <symbol> <decimals>",
		Short:   "Deploys a new migrateable-t0ken contract",
		Example: "t0ken token-migrateable deploy --keystoreAddress owner",
		Args:    cli.ChainArgs(cobra.ExactArgs(2)),
		PreRun:  commands.ConnectWithKeyStore,
		Run: func(cmd *cobra.Command, args []string) {
			// Get the token data args
			name := args[0]
			symbol := args[1]

			// Deploy the token using for the symbol/name/decimals
			addr, tx, _, err := erc20.DeployDynamicSplittableT0ken(cli.Conn.Opts, cli.Conn.Client, name, symbol)
			cli.CheckErr(cmd, err)
			cmd.Println("   Contract:", addr.String())
			cli.PrintTransactionFn(cmd)(tx, nil)
		},
	}

	contractKey  = "token"
	callSession  *erc20.DynamicSplittableT0kenCallerSession
	transSession *erc20.DynamicSplittableT0kenTransactorSession
)

func callerSessionFn(addr common.Address, caller bind.ContractCaller) (interface{}, error) {
	return erc20.NewDynamicSplittableT0kenCaller(addr, caller)
}

func transactorSessionFn(addr common.Address, transactor bind.ContractTransactor) (interface{}, error) {
	return erc20.NewDynamicSplittableT0kenTransactor(addr, transactor)
}

func connectCaller(cmd *cobra.Command, args []string) {
	o, callOpts := commands.ConnectWithCallerSessionFunc(cmd, args, contractKey, callerSessionFn)
	caller := o.(*erc20.DynamicSplittableT0kenCaller)
	callSession = &erc20.DynamicSplittableT0kenCallerSession{caller, callOpts}
}

func connectTransactor(cmd *cobra.Command, args []string) {
	o, transactOpts := commands.ConnectWithTransactorSessionFunc(cmd, args, contractKey, transactorSessionFn)
	transactor := o.(*erc20.DynamicSplittableT0kenTransactor)
	transSession = &erc20.DynamicSplittableT0kenTransactorSession{transactor, transactOpts}
}

func init() {
	// Add the 'gasPrice', 'nonce' and 'wait' flags to the deploy function
	gas.Flag(DeployCommand)
	nonce.Flag(DeployCommand)
	cli.WaitFlag(DeployCommand)
}
