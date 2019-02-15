package storage

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
		Use:   "storage",
		Short: "Storage utilities",
	}

	DeployCommand = &cobra.Command{
		Use:     "deploy",
		Short:   "Deploys a new storage contract",
		Example: "t0ken storage deploy --keystoreAddress owner",
		Args:    cobra.NoArgs,
		PreRun:  connectTransactor,
		Run: func(cmd *cobra.Command, args []string) {
			// Deploy the storage using for the symbol/name/decimals
			addr, tx, _, err := registry.DeployStorage(cli.Conn.Opts, cli.Conn.Client)
			cli.CheckErr(cmd, err)
			cmd.Println("   Contract:", addr.String())
			cli.PrintTransaction(cmd, tx)
		},
	}

	contractKey  = "storage"
	callSession  *registry.StorageCallerSession
	transSession *registry.StorageTransactorSession
)

func callerSessionFn(addr common.Address, caller bind.ContractCaller) (interface{}, error) {
	return registry.NewStorageCaller(addr, caller)
}

func transactorSessionFn(addr common.Address, transactor bind.ContractTransactor) (interface{}, error) {
	return registry.NewStorageTransactor(addr, transactor)
}

func connectCaller(cmd *cobra.Command, args []string) {
	o, callOpts := commands.ConnectWithCallerSessionFunc(cmd, args, contractKey, callerSessionFn)
	caller := o.(*registry.StorageCaller)
	callSession = &registry.StorageCallerSession{caller, callOpts}
}

func connectTransactor(cmd *cobra.Command, args []string) {
	o, transactOpts := commands.ConnectWithTransactorSessionFunc(cmd, args, contractKey, transactorSessionFn)
	transactor := o.(*registry.StorageTransactor)
	transSession = &registry.StorageTransactorSession{transactor, transactOpts}
}

func init() {
	// Add both the 'gasPrice' and 'nonce' flags to the deploy function
	gas.Flag(DeployCommand)
	nonce.Flag(DeployCommand)
}
