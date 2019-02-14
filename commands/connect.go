package commands

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/spf13/cobra"

	"github.com/tzero-dev/go-t0ken/cli"
	"github.com/tzero-dev/go-t0ken/commands/gas"
	"github.com/tzero-dev/go-t0ken/commands/nonce"
)

type CallerSessionFunc func(common.Address, bind.ContractCaller) (interface{}, error)
type TransactorSessionFunc func(common.Address, bind.ContractTransactor) (interface{}, error)

// ConnectWithCallerSession establishes a T0kenCompliance caller session.
func ConnectWithCallerSessionFunc(cmd *cobra.Command, args []string, contractConfigKey string, fn CallerSessionFunc) (interface{}, bind.CallOpts) {
	var caller interface{}
	cli.Connect(cmd, args)
	addr, err := cli.FlagOrConfigAddress(cmd, "address", contractConfigKey)
	if err == nil {
		caller, err = fn(addr, cli.Conn.Client)
	}
	cli.CheckErr(cmd, err)
	return caller, bind.CallOpts{}
}

// ConnectWithTransactorSession establishes a T0kenCompliance transactor session.
func ConnectWithTransactorSessionFunc(cmd *cobra.Command, args []string, contractConfigKey string, fn TransactorSessionFunc) (interface{}, bind.TransactOpts) {
	cli.ConnectWithKeyStore(cmd, args)

	// Set gas price
	gasPrice, err := gas.GetPrice(cmd, args)
	cli.CheckErr(cmd, err)
	cli.Conn.SetGasPrice(gasPrice)

	// Set nonce
	n, err := nonce.Get()
	cli.CheckErr(cmd, err)
	cli.Conn.SetNonce(n)

	var transactor interface{}
	// Create session
	addr, err := cli.FlagOrConfigAddress(cmd, "address", contractConfigKey)
	if err == nil {
		transactor, err = fn(addr, cli.Conn.Client)
	}
	cli.CheckErr(cmd, err)
	return transactor, *cli.Conn.Opts
}
