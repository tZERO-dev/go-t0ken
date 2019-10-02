package token

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
	t "github.com/tzero-dev/go-t0ken/contracts/token"
)

var (
	Command = &cobra.Command{
		Use:   "token",
		Short: "T0ken utilities",
	}

	DeployCommand = &cobra.Command{
		Use:     "deploy <name> <symbol> <decimals>",
		Short:   "Deploys a new t0ken contract",
		Example: "t0ken token deploy --keystoreAddress owner",
		Args:    cli.ChainArgs(cobra.ExactArgs(3), cli.UintArgFunc("decimals", 2, 8)),
		PreRun:  commands.ConnectWithKeyStore,
		Run: func(cmd *cobra.Command, args []string) {
			// Get the token data args
			name := args[0]
			symbol := args[1]
			decimals, err := strconv.ParseInt(args[2], 10, 8)

			// Deploy the token using for the symbol/name/decimals
			addr, tx, _, err := t.DeployT0ken(cli.Conn.Opts, cli.Conn.Client, name, symbol, uint8(decimals))
			cli.CheckErr(cmd, err)
			cmd.Println("   Contract:", addr.String())
			cli.PrintTransactionFn(cmd)(tx, nil)
		},
	}

	AuditCommand = &cobra.Command{
		Use:   "audit",
		Short: "Audit the holders of the t0ken, outputting CSV to <stdout>",
		Example: `t0ken token audit
t0ken token audit --block 8658083`,
		PreRun: connectCaller,
		Run: func(cmd *cobra.Command, arts []string) {
			if callSession.CallOpts.BlockNumber == nil {
				// Get the current block
				header, err := cli.Conn.HeaderByNumber(context.Background(), nil)
				cli.CheckErr(cmd, err)

				// Pin the requests to the same block (avoid any state changes during audit)
				callSession.CallOpts.BlockNumber = header.Number
			}
			n, err := callSession.Holders()
			cli.CheckErr(cmd, err)

			one := big.NewInt(1)
			index := big.NewInt(0)
			//cmd.Printf("Block: %s\n", header.Number.String())
			cmd.Printf("Block: %s\n", callSession.CallOpts.BlockNumber.String())
			fmt.Println("address,tokens")
			for i := int64(0); i < n.Int64(); i++ {
				holder, err := callSession.HolderAt(index)
				cli.CheckErr(cmd, err)

				balance, err := callSession.BalanceOf(holder)
				cli.CheckErr(cmd, err)

				fmt.Printf("%s,%s\n", holder.String(), balance.String())
				index.Add(index, one)
			}
		},
	}

	contractKey  = "t0ken"
	filterer     *t.T0kenFilterer
	callSession  *t.T0kenCallerSession
	transSession *t.T0kenTransactorSession
)

func callerSessionFn(addr common.Address, caller bind.ContractCaller) (interface{}, error) {
	return t.NewT0kenCaller(addr, caller)
}

func filterSessionFn(addr common.Address, filter bind.ContractFilterer) (interface{}, error) {
	return t.NewT0kenFilterer(addr, filter)
}

func transactorSessionFn(addr common.Address, transactor bind.ContractTransactor) (interface{}, error) {
	return t.NewT0kenTransactor(addr, transactor)
}

func connectCaller(cmd *cobra.Command, args []string) {
	o, callOpts := commands.ConnectWithCallerSessionFunc(cmd, args, contractKey, callerSessionFn)
	caller := o.(*t.T0kenCaller)
	callSession = &t.T0kenCallerSession{caller, callOpts}
}

func connectFilterer(cmd *cobra.Command, args []string) {
	f := commands.ConnectWithFiltererSessionFunc(cmd, args, contractKey, filterSessionFn)
	filterer = f.(*t.T0kenFilterer)
}

func connectTransactor(cmd *cobra.Command, args []string) {
	o, transactOpts := commands.ConnectWithTransactorSessionFunc(cmd, args, contractKey, transactorSessionFn)
	transactor := o.(*t.T0kenTransactor)
	transSession = &t.T0kenTransactorSession{transactor, transactOpts}
}

func init() {
	// Add the 'gasPrice', 'nonce' and 'wait' flags to the deploy function
	gas.Flag(DeployCommand)
	nonce.Flag(DeployCommand)
	cli.WaitFlag(DeployCommand)

	// Allow providing contract 'address' flag
	AuditCommand.Flags().String("address", "", `address of the token contract (default "[`+contractKey+`] value from config")`)
	cli.BlockFlag(AuditCommand)
}
