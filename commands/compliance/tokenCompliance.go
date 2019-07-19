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
		Example: "t0ken compliance deploy 0x5bd5b4e1a2c9b12812795e7217201b78c8c10b78 0x0f18bbc1ae1c5ab891a7feb06d65388ba6b4cd07 --keystoreAddress owner",
		Args:    cli.ChainArgs(cli.AddressArgFunc("registry", 0), cli.AddressArgFunc("complianceStorage", 1)),
		PreRun:  commands.ConnectWithKeyStore,
		// TODO: refactor these deploy funcs into something...
		Run: func(cmd *cobra.Command, args []string) {
			// Deploy the token-compliance using for the symbol/name/decimals
			r := common.HexToAddress(args[0])
			s := common.HexToAddress(args[1])
			addr, tx, _, err := compliance.DeployT0kenCompliance(cli.Conn.Opts, cli.Conn.Client, r, s)
			cli.CheckErr(cmd, err)
			cmd.Println("   Contract:", addr.String())
			cli.PrintTransactionFn(cmd)(tx, nil)
		},
	}

	//contractKey  = "t0kenCompliance"
	contractKey  = "compliance"
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
	// Add the 'gasPrice', 'nonce' and 'wait' flags to the deploy function
	gas.Flag(DeployCommand)
	nonce.Flag(DeployCommand)
	cli.WaitFlag(DeployCommand)
}
