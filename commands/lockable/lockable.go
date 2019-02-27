package lockable

import (
	"strconv"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/spf13/cobra"

	"github.com/tzero-dev/go-t0ken/cli"
	"github.com/tzero-dev/go-t0ken/commands"
	"github.com/tzero-dev/go-t0ken/contracts/lifecycle"
)

func callerSessionFn(addr common.Address, caller bind.ContractCaller) (interface{}, error) {
	return lifecycle.NewLockableCaller(addr, caller)
}

func transactorSessionFn(addr common.Address, transactor bind.ContractTransactor) (interface{}, error) {
	return lifecycle.NewLockableTransactor(addr, transactor)
}

func NewGetterCommands(contractConfigKey string) []*cobra.Command {
	var session *lifecycle.LockableCallerSession

	connect := func(cmd *cobra.Command, args []string) {
		caller, callOpts := commands.ConnectWithCallerSessionFunc(cmd, args, contractConfigKey, callerSessionFn)
		session = &lifecycle.LockableCallerSession{caller.(*lifecycle.LockableCaller), callOpts}
	}

	return []*cobra.Command{
		&cobra.Command{
			Use:    "isLocked",
			Short:  "Checks if the contract is currently in a locked state",
			Args:   cobra.NoArgs,
			PreRun: connect,
			Run: func(cmd *cobra.Command, args []string) {
				cli.CheckGetter(cmd)(session.IsLocked())
			},
		},
	}
}

func NewSetterCommands(contractConfigKey string) []*cobra.Command {
	var session *lifecycle.LockableTransactorSession

	connect := func(cmd *cobra.Command, args []string) {
		transactor, transactOpts := commands.ConnectWithTransactorSessionFunc(cmd, args, contractConfigKey, transactorSessionFn)
		session = &lifecycle.LockableTransactorSession{transactor.(*lifecycle.LockableTransactor), transactOpts}
	}

	return []*cobra.Command{
		&cobra.Command{
			Use:    "setLocked <bool>",
			Short:  "Sets the contract to the locked <bool> state",
			Args:   cli.BoolArgFunc("bool", 0),
			PreRun: connect,
			Run: func(cmd *cobra.Command, args []string) {
				locked, _ := strconv.ParseBool(args[0])
				cli.PrintTransactionFn(cmd)(session.SetLocked(locked))
			},
		},
	}
}
