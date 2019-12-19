package commands

import (
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/spf13/cobra"

	"github.com/tzero-dev/go-t0ken/cli"
	"github.com/tzero-dev/go-t0ken/commands/gas"
	"github.com/tzero-dev/go-t0ken/commands/nonce"
)

type CallerSessionFunc func(common.Address, bind.ContractCaller) (interface{}, error)
type FilterSessionFunc func(common.Address, bind.ContractFilterer) (interface{}, error)
type TransactorSessionFunc func(common.Address, bind.ContractTransactor) (interface{}, error)

// ConnectWithKeyStore establishes a connection with gasPrice and nonce values set.
func ConnectWithKeyStore(cmd *cobra.Command, args []string) {
	cli.ConnectWithKeyStore(cmd, args)

	// Set gas price
	gasPrice, err := gas.GetPrice(cmd, args)
	cli.CheckErr(cmd, err)
	cli.Conn.SetGasPrice(gasPrice)

	// Set nonce
	n, err := nonce.Get()
	cli.CheckErr(cmd, err)
	cli.Conn.SetNonce(n)
}

// ConnectWithCallerSession establishes a caller session.
func ConnectWithCallerSessionFunc(cmd *cobra.Command, args []string, contractConfigKey string, fn CallerSessionFunc) (interface{}, bind.CallOpts) {
	// Geth the contract flag or config address
	var caller interface{}
	cli.Connect(cmd, args)
	addr, err := cli.FlagOrConfigAddress(cmd, "address", contractConfigKey)
	if err == nil {
		caller, err = fn(addr, cli.Conn.Client)
	}
	cli.CheckErr(cmd, err)

	// Get the blocknumber to use, if provided
	block, err := cmd.Flags().GetInt64("block")
	cli.CheckErr(cmd, err)

	// Set blocknumber to the 'block' flag
	o := bind.CallOpts{}
	if block > 0 {
		o.BlockNumber = big.NewInt(block)
	}
	return caller, o
}

// ConnectWithCallerSession establishes a filter session.
func ConnectWithFiltererSessionFunc(cmd *cobra.Command, args []string, contractConfigKey string, fn FilterSessionFunc) interface{} {
	var filter interface{}
	cli.Connect(cmd, nil)
	addr, err := cli.FlagOrConfigAddress(cmd, "address", contractConfigKey)
	if err == nil {
		filter, err = fn(addr, cli.Conn.Client)
	}
	cli.CheckErr(cmd, err)
	return filter
}

// ConnectWithTransactorSession establishes a transactor session.
func ConnectWithTransactorSessionFunc(cmd *cobra.Command, args []string, contractConfigKey string, fn TransactorSessionFunc) (interface{}, bind.TransactOpts) {
	var transactor interface{}
	ConnectWithKeyStore(cmd, args)
	// Create session
	addr, err := cli.FlagOrConfigAddress(cmd, "address", contractConfigKey)
	if err == nil {
		transactor, err = fn(addr, cli.Conn.Client)
	}
	cli.CheckErr(cmd, err)
	return transactor, *cli.Conn.Opts
}
