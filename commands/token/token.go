package token

import (
	"strconv"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/spf13/cobra"

	"github.com/tzero-dev/go-t0ken/cli"
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
		Use:    "deploy",
		Short:  "Deploys a new t0ken contract",
		Args:   cli.ChainArgs(cobra.ExactArgs(3), cli.UintArgFunc(2, 8)),
		PreRun: cli.ConnectWithKeyStore,
		Run: func(cmd *cobra.Command, args []string) {
			// Get the token data args
			name := args[0]
			symbol := args[1]
			decimals, err := strconv.ParseInt(args[2], 10, 8)

			// Deploy the token using for the symbol/name/decimals
			addr, tx, _, err := erc20.DeployT0ken(cli.Conn.Opts, cli.Conn.Client, name, symbol, uint8(decimals))
			cli.CheckErr(cmd, err)
			cmd.Println("   Contract:", addr.String())
			cli.PrintTransaction(cmd, tx)
		},
	}

	address      string
	callSession  *erc20.T0kenCallerSession
	transSession *erc20.T0kenTransactorSession
)

type t0kenGetterFunc func(s *erc20.T0kenCallerSession) (interface{}, error)

// newTokenTransactorSession returns a transactor session for the T0ken contract.
func newT0kenTransactorSession(addr common.Address) (*erc20.T0kenTransactorSession, error) {
	transactor, err := erc20.NewT0kenTransactor(addr, cli.Conn.Client)
	if err != nil {
		return nil, err
	}
	session := &erc20.T0kenTransactorSession{transactor, *cli.Conn.Opts}
	return session, err
}

// newTokenCallerSession returns a transactor caller session for the T0ken contract.
func newT0kenCallerSession(addr common.Address) (*erc20.T0kenCallerSession, error) {
	caller, err := erc20.NewT0kenCaller(addr, cli.Conn.Client)
	if err != nil {
		return nil, err
	}
	session := &erc20.T0kenCallerSession{caller, bind.CallOpts{}}
	return session, err
}

// connectWithT0kenTransactorSession establishes a t0ken caller session.
func connectWithT0kenCallerSession(cmd *cobra.Command, args []string) {
	cli.Connect(cmd, args)
	addr, err := cli.FlagOrConfigAddress("address", "t0ken")
	if err == nil {
		callSession, err = newT0kenCallerSession(addr)
	}
	cli.CheckErr(cmd, err)
}

// connectWithT0kenTransactorSession establishes a t0ken transactor session.
func connectWithT0kenTransactorSession(cmd *cobra.Command, args []string) {
	cli.ConnectWithKeyStore(cmd, args)

	// Set gas price
	gasPrice, err := gas.GetPrice(cmd, args)
	cli.CheckErr(cmd, err)
	cli.Conn.SetGasPrice(gasPrice)

	// Set nonce
	n, err := nonce.Get()
	cli.CheckErr(cmd, err)
	cli.Conn.SetNonce(n)

	// Create session
	addr, err := cli.FlagOrConfigAddress("address", "t0ken")
	if err == nil {
		transSession, err = newT0kenTransactorSession(addr)
	}
	cli.CheckErr(cmd, err)
}

func init() {
	// Add both the 'gasPrice' and 'nonce' flags to the deploy function
	gas.Flag(DeployCommand)
	nonce.Flag(DeployCommand)
}
