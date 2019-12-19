package registry

import (
	"context"
	"errors"
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
		Use:   "registry",
		Short: "Registry utilities",
	}

	DeployCommand = &cobra.Command{
		Use:     "deploy",
		Short:   "Deploys a new registry contract",
		Example: "t0ken registry deploy --keystoreAddress owner",
		Args:    cobra.NoArgs,
		PreRun:  commands.ConnectWithKeyStore,
		Run: func(cmd *cobra.Command, args []string) {
			// Deploy the registry using for the symbol/name/decimals
			addr, tx, _, err := registry.DeployRegistry(cli.Conn.Opts, cli.Conn.Client)
			cli.CheckErr(cmd, err)
			cmd.Println("   Contract:", addr.String())
			cli.PrintTransactionFn(cmd)(tx, nil)
		},
	}

	AuditCommand = &cobra.Command{
		Use:   "audit [kind]",
		Short: "Audit the accounts of <kind>, outputting CSV to <stdout>",
		Example: `t0ken registry audit
t0ken token audit 4,
t0ken token audit 4 --block 8658083`,
		Args:   cobra.MaximumNArgs(1),
		PreRun: connectCaller,
		Run: func(cmd *cobra.Command, args []string) {
			var kind *uint8
			if len(args) > 0 {
				if cli.IsIntArg("kind", args, 0) == nil {
					k, err := strconv.Atoi(args[0])
					if err != nil || k == 0 {
						cli.CheckErr(cmd, errors.New("invalid value for [kind] arg"))
					}
					kind = new(uint8)
					*kind = uint8(k)
				}
			}

			if callSession.CallOpts.BlockNumber == nil {
				// Get the current block
				header, err := cli.Conn.HeaderByNumber(context.Background(), nil)
				cli.CheckErr(cmd, err)

				// Pin the requests to the same block (avoid any state changes during audit)
				callSession.CallOpts.BlockNumber = header.Number
			}
			n, err := callSession.Accounts()
			cli.CheckErr(cmd, err)

			one := big.NewInt(1)
			index := big.NewInt(0)
			cmd.Printf("Block: %s\n", callSession.CallOpts.BlockNumber.String())
			fmt.Println("index,address,kind,frozen,parent")
			for i := int64(0); i < n.Int64(); i++ {
				addr, k, frozen, parent, hash, err := callSession.AccountAt(index)
				cli.CheckErr(cmd, err)

				if kind == nil || k == *kind {
					fmt.Printf("%d,%s,%d,%t,%s,%x\n", i, addr.String(), k, frozen, parent.String(), hash)
				}
				index.Add(index, one)
			}
		},
	}

	contractKey  = "registry"
	callSession  *registry.RegistryCallerSession
	transSession *registry.RegistryTransactorSession
)

func callerSessionFn(addr common.Address, caller bind.ContractCaller) (interface{}, error) {
	return registry.NewRegistryCaller(addr, caller)
}

func transactorSessionFn(addr common.Address, transactor bind.ContractTransactor) (interface{}, error) {
	return registry.NewRegistryTransactor(addr, transactor)
}

func connectCaller(cmd *cobra.Command, args []string) {
	o, callOpts := commands.ConnectWithCallerSessionFunc(cmd, args, contractKey, callerSessionFn)
	caller := o.(*registry.RegistryCaller)
	callSession = &registry.RegistryCallerSession{caller, callOpts}
}

func connectTransactor(cmd *cobra.Command, args []string) {
	o, transactOpts := commands.ConnectWithTransactorSessionFunc(cmd, args, contractKey, transactorSessionFn)
	transactor := o.(*registry.RegistryTransactor)
	transSession = &registry.RegistryTransactorSession{transactor, transactOpts}
}

func init() {
	// Add the 'gasPrice', 'nonce' and 'wait' flags to the deploy function
	gas.Flag(DeployCommand)
	nonce.Flag(DeployCommand)
	cli.WaitFlag(DeployCommand)

	// Allow providing contract 'address' flag
	AuditCommand.Flags().String("address", "", `address of the Registry contract (default "[`+contractKey+`] value from config")`)
	cli.BlockFlag(AuditCommand)
}
