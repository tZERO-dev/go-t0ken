package investor

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
		Use:   "investor",
		Short: "Investor utilities",
	}

	DeployCommand = &cobra.Command{
		Use:     "deploy <registry>",
		Short:   "Deploys a new investor contract",
		Example: "t0ken investor deploy 0x397e7b9c15ff22ba67ec6e78f46f1e21540bcb36 --keystoreAddress owner",
		Args:    cli.AddressArgFunc("registry", 0),
		PreRun:  commands.ConnectWithKeyStore,
		Run: func(cmd *cobra.Command, args []string) {
			// Deploy the investor registry, pointing to the registry
			registryAddress := common.HexToAddress(args[0])
			addr, tx, _, err := registry.DeployInvestor(cli.Conn.Opts, cli.Conn.Client, registryAddress)
			cli.CheckErr(cmd, err)
			cmd.Println("   Contract:", addr.String())
			cli.PrintTransactionFn(cmd)(tx, nil)
		},
	}

	contractKey  = "investorRegistry"
	filterer     *registry.InvestorFilterer
	callSession  *registry.InvestorCallerSession
	transSession *registry.InvestorTransactorSession
)

func callerSessionFn(addr common.Address, caller bind.ContractCaller) (interface{}, error) {
	return registry.NewInvestorCaller(addr, caller)
}

func filterSessionFn(addr common.Address, filter bind.ContractFilterer) (interface{}, error) {
	return registry.NewInvestorFilterer(addr, filter)
}

func transactorSessionFn(addr common.Address, transactor bind.ContractTransactor) (interface{}, error) {
	return registry.NewInvestorTransactor(addr, transactor)
}

func connectCaller(cmd *cobra.Command, args []string) {
	c, callOpts := commands.ConnectWithCallerSessionFunc(cmd, args, contractKey, callerSessionFn)
	caller := c.(*registry.InvestorCaller)
	callSession = &registry.InvestorCallerSession{caller, callOpts}
}

func connectFilterer(cmd *cobra.Command, args []string) {
	f := commands.ConnectWithFiltererSessionFunc(cmd, args, contractKey, filterSessionFn)
	filterer = f.(*registry.InvestorFilterer)
}

func connectTransactor(cmd *cobra.Command, args []string) {
	t, transactOpts := commands.ConnectWithTransactorSessionFunc(cmd, args, contractKey, transactorSessionFn)
	transactor := t.(*registry.InvestorTransactor)
	transSession = &registry.InvestorTransactorSession{transactor, transactOpts}
}

func init() {
	// Add both the 'gasPrice' and 'nonce' flags to the deploy function
	gas.Flag(DeployCommand)
	nonce.Flag(DeployCommand)

	// Add the 'wait' flag
	DeployCommand.Flags().Int("wait", -1, "waits the provided number of seconds for the transaction to be mined ('0' waits indefinitely)")
}
