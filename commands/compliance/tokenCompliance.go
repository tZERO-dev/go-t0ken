package compliance

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/spf13/cobra"

	"github.com/tzero-dev/go-t0ken/cli"
	"github.com/tzero-dev/go-t0ken/commands"
	"github.com/tzero-dev/go-t0ken/commands/gas"
	"github.com/tzero-dev/go-t0ken/commands/nonce"
	"github.com/tzero-dev/go-t0ken/contracts/compliance"
)

var (
	Command = &cobra.Command{
		Use:   "compliance",
		Short: "T0ken-Compliance utilities",
	}

	DeployCommand = &cobra.Command{
		Use:     "deploy",
		Short:   "Deploys a new t0ken-compliance contract",
		Example: "t0ken compliance deploy --keystoreAddress owner",
		Args:    cobra.NoArgs,
		PreRun:  commands.ConnectWithKeyStore,
		// TODO: refactor these deploy funcs into something...
		Run: func(cmd *cobra.Command, args []string) {
			// Deploy the token-compliance using for the symbol/name/decimals
			addr, tx, _, err := compliance.DeployT0kenCompliance(cli.Conn.Opts, cli.Conn.Client)
			cli.CheckErr(cmd, err)
			cmd.Println("   Contract:", addr.String())
			cli.PrintTransactionFn(cmd)(tx, nil)
		},
	}

	contractKey  = "t0kenCompliance"
	callSession  *compliance.T0kenComplianceCallerSession
	transSession *compliance.T0kenComplianceTransactorSession
)

func callerSessionFn(addr common.Address, caller bind.ContractCaller) (interface{}, error) {
	return compliance.NewT0kenComplianceCaller(addr, caller)
}

func transactorSessionFn(addr common.Address, transactor bind.ContractTransactor) (interface{}, error) {
	return compliance.NewT0kenComplianceTransactor(addr, transactor)
}

func connectCaller(cmd *cobra.Command, args []string) {
	o, callOpts := commands.ConnectWithCallerSessionFunc(cmd, args, contractKey, callerSessionFn)
	caller := o.(*compliance.T0kenComplianceCaller)
	callSession = &compliance.T0kenComplianceCallerSession{caller, callOpts}
}

func connectTransactor(cmd *cobra.Command, args []string) {
	o, transactOpts := commands.ConnectWithTransactorSessionFunc(cmd, args, contractKey, transactorSessionFn)
	transactor := o.(*compliance.T0kenComplianceTransactor)
	transSession = &compliance.T0kenComplianceTransactorSession{transactor, transactOpts}
}

func init() {
	// Add both the 'gasPrice' and 'nonce' flags to the deploy function
	gas.Flag(DeployCommand)
	nonce.Flag(DeployCommand)

	// Add the 'wait' flag
	DeployCommand.Flags().Int("wait", -1, "waits the provided number of seconds for the transaction to be mined ('0' waits indefinitely)")
}
