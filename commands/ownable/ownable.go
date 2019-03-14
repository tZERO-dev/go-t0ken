package ownable

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/spf13/cobra"

	"github.com/tzero-dev/go-t0ken/cli"
	"github.com/tzero-dev/go-t0ken/commands"
	"github.com/tzero-dev/go-t0ken/contracts/ownership"
)

func callerSessionFn(addr common.Address, caller bind.ContractCaller) (interface{}, error) {
	return ownership.NewOwnableCaller(addr, caller)
}

func transactorSessionFn(addr common.Address, transactor bind.ContractTransactor) (interface{}, error) {
	return ownership.NewOwnableTransactor(addr, transactor)
}

func NewGetterCommands(contractConfigKey string) []*cobra.Command {
	var session *ownership.OwnableCallerSession

	connect := func(cmd *cobra.Command, args []string) {
		caller, callOpts := commands.ConnectWithCallerSessionFunc(cmd, args, contractConfigKey, callerSessionFn)
		session = &ownership.OwnableCallerSession{caller.(*ownership.OwnableCaller), callOpts}
	}

	return []*cobra.Command{
		&cobra.Command{
			Use:    "owner",
			Short:  "Gets the owner address of the t0ken",
			Args:   cobra.NoArgs,
			PreRun: connect,
			Run: func(cmd *cobra.Command, args []string) {
				cli.CheckAddressGetter(cmd)(session.Owner())
			},
		},
	}
}

func NewSetterCommands(contractConfigKey string) []*cobra.Command {
	var session *ownership.OwnableTransactorSession

	connect := func(cmd *cobra.Command, args []string) {
		transactor, transactOpts := commands.ConnectWithTransactorSessionFunc(cmd, args, contractConfigKey, transactorSessionFn)
		session = &ownership.OwnableTransactorSession{transactor.(*ownership.OwnableTransactor), transactOpts}
	}

	return []*cobra.Command{
		&cobra.Command{
			Use:    "transferOwner <address>",
			Short:  "Transfers ownership of the contract to <address>",
			Args:   cli.ChainArgs(cobra.ExactArgs(1), cli.AddressArgFunc("address", 0)),
			PreRun: connect,
			Run: func(cmd *cobra.Command, args []string) {
				addr := common.HexToAddress(args[0])
				cli.PrintTransactionFn(cmd)(session.TransferOwner(addr))
			},
		},
	}
}
