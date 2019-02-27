package custodian

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/spf13/cobra"

	"github.com/tzero-dev/go-t0ken/cli"
	"github.com/tzero-dev/go-t0ken/commands"
	"github.com/tzero-dev/go-t0ken/commands/gas"
	"github.com/tzero-dev/go-t0ken/commands/nonce"
	"github.com/tzero-dev/go-t0ken/contracts/registry"
)

var (
	Command = &cobra.Command{
		Use:   "custodian",
		Short: "Custodian utilities",
	}

	DeployCommand = &cobra.Command{
		Use:     "deploy",
		Short:   "Deploys a new custodian registry contract",
		Example: "t0ken custodian deploy --keystoreAddress owner",
		Args:    cobra.NoArgs,
		PreRun:  commands.ConnectWithKeyStore,
		Run: func(cmd *cobra.Command, args []string) {
			// Deploy the custodian using for the symbol/name/decimals
			addr, tx, _, err := registry.DeployCustodian(cli.Conn.Opts, cli.Conn.Client)
			cli.CheckErr(cmd, err)
			cmd.Println("   Contract:", addr.String())
			cli.PrintTransactionFn(cmd)(tx, nil)
		},
	}

	contractKey  = "custodianRegistry"
	callSession  *registry.CustodianCallerSession
	transSession *registry.CustodianTransactorSession
)

func callerSessionFn(addr common.Address, caller bind.ContractCaller) (interface{}, error) {
	return registry.NewCustodianCaller(addr, caller)
}

func transactorSessionFn(addr common.Address, transactor bind.ContractTransactor) (interface{}, error) {
	return registry.NewCustodianTransactor(addr, transactor)
}

func connectCaller(cmd *cobra.Command, args []string) {
	o, callOpts := commands.ConnectWithCallerSessionFunc(cmd, args, contractKey, callerSessionFn)
	caller := o.(*registry.CustodianCaller)
	callSession = &registry.CustodianCallerSession{caller, callOpts}
}

func connectTransactor(cmd *cobra.Command, args []string) {
	o, transactOpts := commands.ConnectWithTransactorSessionFunc(cmd, args, contractKey, transactorSessionFn)
	transactor := o.(*registry.CustodianTransactor)
	transSession = &registry.CustodianTransactorSession{transactor, transactOpts}
}

func init() {
	// Add both the 'gasPrice' and 'nonce' flags to the deploy function
	gas.Flag(DeployCommand)
	nonce.Flag(DeployCommand)

	// Add the 'wait' flag
	DeployCommand.Flags().Int("wait", -1, "waits the provided number of seconds for the transaction to be mined ('0' waits indefinitely)")
}
