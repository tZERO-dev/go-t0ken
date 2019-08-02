package administrable

import (
	"math/big"
	"strconv"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/spf13/cobra"

	"github.com/tzero-dev/go-t0ken/cli"
	"github.com/tzero-dev/go-t0ken/commands"
	"github.com/tzero-dev/go-t0ken/contracts/ownership"
)

func callerSessionFn(addr common.Address, caller bind.ContractCaller) (interface{}, error) {
	return ownership.NewAdministrableCaller(addr, caller)
}

func transactorSessionFn(addr common.Address, transactor bind.ContractTransactor) (interface{}, error) {
	return ownership.NewAdministrableTransactor(addr, transactor)
}

func NewGetterCommands(contractConfigKey string) []*cobra.Command {
	var session *ownership.AdministrableCallerSession

	connect := func(cmd *cobra.Command, args []string) {
		o, callOpts := commands.ConnectWithCallerSessionFunc(cmd, args, contractConfigKey, callerSessionFn)
		caller := o.(*ownership.AdministrableCaller)
		session = &ownership.AdministrableCallerSession{caller, callOpts}
	}

	return []*cobra.Command{
		&cobra.Command{
			Use:    "admins",
			Short:  "Gets the total number of admin addresses that exist",
			Args:   cobra.NoArgs,
			PreRun: connect,
			Run: func(cmd *cobra.Command, args []string) {
				cli.CheckGetter(cmd)(session.Admins())
			},
		},
		&cobra.Command{
			Use:    "isAdmin <address>",
			Short:  "Checks if the given <address> is a contract admin",
			Args:   cli.AddressArgFunc("address", 0),
			PreRun: connect,
			Run: func(cmd *cobra.Command, args []string) {
				addr := common.HexToAddress(args[0])
				cli.CheckGetter(cmd)(session.IsAdmin(addr))
			},
		},
		&cobra.Command{
			Use:    "adminAt <index>",
			Short:  "Gets the admin address at the given <index>",
			Args:   cli.IntArgFunc("index", 0),
			PreRun: connect,
			Run: func(cmd *cobra.Command, args []string) {
				index, _ := new(big.Int).SetString(args[0], 10)
				cli.CheckGetter(cmd)(session.AdminAt(index))
			},
		},
	}
}

func NewSetterCommands(contractConfigKey string) []*cobra.Command {
	var session *ownership.AdministrableTransactorSession

	connect := func(cmd *cobra.Command, args []string) {
		o, transactOpts := commands.ConnectWithTransactorSessionFunc(cmd, args, contractConfigKey, transactorSessionFn)
		transactor := o.(*ownership.AdministrableTransactor)
		session = &ownership.AdministrableTransactorSession{transactor, transactOpts}
	}

	return []*cobra.Command{
		&cobra.Command{
			Use:    "setAdmin <address> <isAdmin>",
			Short:  "Adds/removes the given <address> as an admin",
			Args:   cli.ChainArgs(cli.AddressArgFunc("address", 0), cli.BoolArgFunc("isAdmin", 1)),
			PreRun: connect,
			Run: func(cmd *cobra.Command, args []string) {
				addr := common.HexToAddress(args[0])
				isAdmin, _ := strconv.ParseBool(args[1])
				cli.PrintTransactionFn(cmd)(session.SetAdmin(addr, isAdmin))
			},
		},
	}
}
