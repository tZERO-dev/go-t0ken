package destroyable

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/spf13/cobra"

	"github.com/tzero-dev/go-t0ken/cli"
	"github.com/tzero-dev/go-t0ken/commands"
	"github.com/tzero-dev/go-t0ken/contracts/lifecycle"
)

func transactorSessionFn(addr common.Address, transactor bind.ContractTransactor) (interface{}, error) {
	return lifecycle.NewDestroyableTransactor(addr, transactor)
}

func NewSetterCommands(contractConfigKey string) []*cobra.Command {
	var session *lifecycle.DestroyableTransactorSession

	connect := func(cmd *cobra.Command, args []string) {
		transactor, transactOpts := commands.ConnectWithTransactorSessionFunc(cmd, args, contractConfigKey, transactorSessionFn)
		session = &lifecycle.DestroyableTransactorSession{transactor.(*lifecycle.DestroyableTransactor), transactOpts}
	}

	return []*cobra.Command{
		&cobra.Command{
			Use:    "kill",
			Short:  "Invokes 'selfdestruct' on the contract",
			Args:   cobra.NoArgs,
			PreRun: connect,
			Run: func(cmd *cobra.Command, args []string) {
				cli.PrintTransFn(cmd)(session.Kill())
			},
		},
	}
}
