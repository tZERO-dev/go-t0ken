package storage

import (
	"context"
	"fmt"
	"math/big"
	"strconv"

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
		PreRun:  commands.ConnectWithKeyStore,
		Run: func(cmd *cobra.Command, args []string) {
			// Deploy the storage using for the symbol/name/decimals
			addr, tx, _, err := registry.DeployStorage(cli.Conn.Opts, cli.Conn.Client)
			cli.CheckErr(cmd, err)
			cmd.Println("   Contract:", addr.String())
			cli.PrintTransactionFn(cmd)(tx, nil)
		},
	}

	AuditCommand = &cobra.Command{
		Use:    "audit <kind>",
		Short:  "Audit the accounts of <kind>, outputting CSV to <stdout>",
		Args:   cli.IntArgFunc("kind", 0),
		PreRun: connectCaller,
		Run: func(cmd *cobra.Command, args []string) {
			var kind uint8
			if cli.IsIntArg("kind", args, 0) == nil {
				k, _ := strconv.Atoi(args[0])
				kind = uint8(k)
			}

			// Get the current block
			header, err := cli.Conn.HeaderByNumber(context.Background(), nil)
			cli.CheckErr(cmd, err)

			// Pin the requests to the same block (avoid any state changes during audit)
			callSession.CallOpts.BlockNumber = header.Number
			n, err := callSession.Accounts()
			cli.CheckErr(cmd, err)

			one := big.NewInt(1)
			index := big.NewInt(0)
			cmd.Printf("Block: %s\n", header.Number.String())
			fmt.Println("address,kind,frozen,parent")
			for i := int64(0); i < n.Int64(); i++ {
				addr, k, frozen, parent, err := callSession.AccountAt(index)
				cli.CheckErr(cmd, err)

				if kind == 0 || k == kind {
					fmt.Printf("%s,%d,%t,%s\n", addr.String(), k, frozen, parent.String())
				}
				index.Add(index, one)
			}
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
	// Add the 'gasPrice', 'nonce' and 'wait' flags to the deploy function
	gas.Flag(DeployCommand)
	nonce.Flag(DeployCommand)
	cli.WaitFlag(DeployCommand)

	// Allow providing contract 'address' flag
	AuditCommand.Flags().String("address", "", `address of the BrokerDealer registry contract (default "[`+contractKey+`] value from config")`)
}
