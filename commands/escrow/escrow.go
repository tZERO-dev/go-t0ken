package escrow

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/spf13/cobra"

	"github.com/tzero-dev/go-t0ken/cli"
	"github.com/tzero-dev/go-t0ken/commands"
	"github.com/tzero-dev/go-t0ken/commands/gas"
	"github.com/tzero-dev/go-t0ken/commands/nonce"
	"github.com/tzero-dev/go-t0ken/contracts/escrow"
)

var (
	Command = &cobra.Command{
		Use:   "escrow",
		Short: "Escrow utilities",
	}

	DeployCommand = &cobra.Command{
		Use:     "deploy <registry>",
		Short:   "Deploys a new Escrow contract",
		Example: "t0ken escrow deploy 0x397e7b9c15ff22ba67ec6e78f46f1e21540bcb36 --keystoreAddress owner",
		Args:    cli.AddressArgFunc("registry", 0),
		PreRun:  commands.ConnectWithKeyStore,
		Run: func(cmd *cobra.Command, args []string) {
			// Deploy the escrow, pointing to the registry
			registryAddress := common.HexToAddress(args[0])
			addr, tx, _, err := escrow.DeployEscrow(cli.Conn.Opts, cli.Conn.Client, registryAddress)
			cli.CheckErr(cmd, err)
			cmd.Println("   Contract:", addr.String())
			cli.PrintTransactionFn(cmd)(tx, nil)
		},
	}

	contractKey  = "escrow"
	callSession  *escrow.EscrowCallerSession
	transSession *escrow.EscrowTransactorSession
)

func callerSessionFn(addr common.Address, caller bind.ContractCaller) (interface{}, error) {
	return escrow.NewEscrowCaller(addr, caller)
}

func transactorSessionFn(addr common.Address, transactor bind.ContractTransactor) (interface{}, error) {
	return escrow.NewEscrowTransactor(addr, transactor)
}

func connectCaller(cmd *cobra.Command, args []string) {
	o, callOpts := commands.ConnectWithCallerSessionFunc(cmd, args, contractKey, callerSessionFn)
	caller := o.(*escrow.EscrowCaller)
	callSession = &escrow.EscrowCallerSession{caller, callOpts}
}

func connectTransactor(cmd *cobra.Command, args []string) {
	o, transactOpts := commands.ConnectWithTransactorSessionFunc(cmd, args, contractKey, transactorSessionFn)
	transactor := o.(*escrow.EscrowTransactor)
	transSession = &escrow.EscrowTransactorSession{transactor, transactOpts}
}

func init() {
	// Add the 'gasPrice', 'nonce' and 'wait' flags to the deploy function
	gas.Flag(DeployCommand)
	nonce.Flag(DeployCommand)
	cli.WaitFlag(DeployCommand)
}
