package migrateable

import (
	"strconv"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/spf13/cobra"

	"github.com/tzero-dev/go-t0ken/cli"
	"github.com/tzero-dev/go-t0ken/commands"
	"github.com/tzero-dev/go-t0ken/commands/gas"
	"github.com/tzero-dev/go-t0ken/commands/nonce"
	"github.com/tzero-dev/go-t0ken/contracts/token"
)

var (
	Command = &cobra.Command{
		Use:   "token-migrateable",
		Short: "Migrateable-T0ken utilities",
	}

	DeployCommand = &cobra.Command{
		Use:     "deploy <name> <symbol> <decimals>",
		Short:   "Deploys a new migrateable-t0ken contract",
		Example: "t0ken token-migrateable deploy --keystoreAddress owner",
		Args:    cli.ChainArgs(cobra.ExactArgs(3), cli.UintArgFunc("decimals", 2, 8)),
		PreRun:  commands.ConnectWithKeyStore,
		Run: func(cmd *cobra.Command, args []string) {
			// Get the token data args
			name := args[0]
			symbol := args[1]
			decimals, err := strconv.ParseInt(args[2], 10, 8)

			// Deploy the token using for the symbol/name/decimals
			addr, tx, _, err := token.DeployT0kenMigrateable(cli.Conn.Opts, cli.Conn.Client, name, symbol, uint8(decimals))
			cli.CheckErr(cmd, err)
			cmd.Println("   Contract:", addr.String())
			cli.PrintTransactionFn(cmd)(tx, nil)
		},
	}

	contractKey  = "token"
	callSession  *token.T0kenMigrateableCallerSession
	transSession *token.T0kenMigrateableTransactorSession
)

func callerSessionFn(addr common.Address, caller bind.ContractCaller) (interface{}, error) {
	return token.NewT0kenMigrateableCaller(addr, caller)
}

func transactorSessionFn(addr common.Address, transactor bind.ContractTransactor) (interface{}, error) {
	return token.NewT0kenMigrateableTransactor(addr, transactor)
}

func connectCaller(cmd *cobra.Command, args []string) {
	o, callOpts := commands.ConnectWithCallerSessionFunc(cmd, args, contractKey, callerSessionFn)
	caller := o.(*token.T0kenMigrateableCaller)
	callSession = &token.T0kenMigrateableCallerSession{caller, callOpts}
}

func connectTransactor(cmd *cobra.Command, args []string) {
	o, transactOpts := commands.ConnectWithTransactorSessionFunc(cmd, args, contractKey, transactorSessionFn)
	transactor := o.(*token.T0kenMigrateableTransactor)
	transSession = &token.T0kenMigrateableTransactorSession{transactor, transactOpts}
}

func init() {
	// Add the 'gasPrice', 'nonce' and 'wait' flags to the deploy function
	gas.Flag(DeployCommand)
	nonce.Flag(DeployCommand)
	cli.WaitFlag(DeployCommand)
}
