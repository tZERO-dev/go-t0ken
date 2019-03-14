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
	"github.com/tzero-dev/go-t0ken/contracts/token/erc20"
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
			addr, tx, _, err := erc20.DeployT0ken(cli.Conn.Opts, cli.Conn.Client, name, symbol, uint8(decimals))
			cli.CheckErr(cmd, err)
			cmd.Println("   Contract:", addr.String())
			cli.PrintTransactionFn(cmd)(tx, nil)
		},
	}

	AuditCommand = &cobra.Command{
		Use:    "audit",
		Short:  "Audit the holders of the t0ken, outputting CSV to <stdout>",
		PreRun: connectCaller,
		Run: func(cmd *cobra.Command, arts []string) {
			// Get the current block
			header, err := cli.Conn.HeaderByNumber(context.Background(), nil)
			cli.CheckErr(cmd, err)

			// Pin the requests to the same block (avoid any state changes during audit)
			callSession.CallOpts.BlockNumber = header.Number
			n, err := callSession.Shareholders()
			cli.CheckErr(cmd, err)

			one := big.NewInt(1)
			index := big.NewInt(0)
			cmd.Printf("Block: %s\n", header.Number.String())
			fmt.Println("address,balance")
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
	callSession  *erc20.T0kenCallerSession
	transSession *erc20.T0kenTransactorSession
)

func callerSessionFn(addr common.Address, caller bind.ContractCaller) (interface{}, error) {
	return erc20.NewT0kenCaller(addr, caller)
}

func transactorSessionFn(addr common.Address, transactor bind.ContractTransactor) (interface{}, error) {
	return erc20.NewT0kenTransactor(addr, transactor)
}

func connectCaller(cmd *cobra.Command, args []string) {
	o, callOpts := commands.ConnectWithCallerSessionFunc(cmd, args, contractKey, callerSessionFn)
	caller := o.(*erc20.T0kenCaller)
	callSession = &erc20.T0kenCallerSession{caller, callOpts}
}

func connectTransactor(cmd *cobra.Command, args []string) {
	o, transactOpts := commands.ConnectWithTransactorSessionFunc(cmd, args, contractKey, transactorSessionFn)
	transactor := o.(*erc20.T0kenTransactor)
	transSession = &erc20.T0kenTransactorSession{transactor, transactOpts}
}

func init() {
	// Add both the 'gasPrice' and 'nonce' flags to the deploy function
	gas.Flag(DeployCommand)
	nonce.Flag(DeployCommand)

	// Add the 'wait' flag
	DeployCommand.Flags().Int("wait", -1, "waits the provided number of seconds for the transaction to be mined ('0' waits indefinitely)")
}
