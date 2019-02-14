package broker

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/spf13/cobra"

	"github.com/tzero-dev/go-t0ken/cli"
	"github.com/tzero-dev/go-t0ken/commands/destroyable"
	"github.com/tzero-dev/go-t0ken/commands/lockable"
	"github.com/tzero-dev/go-t0ken/commands/ownable"
)

var SetterCommands = []*cobra.Command{
	&cobra.Command{
		Use:    "add <broker>",
		Short:  "Adds the <broker> address",
		Args:   cli.AddressArgFunc("broker", 0),
		PreRun: connectTransactor,
		Run: func(cmd *cobra.Command, args []string) {
			broker := common.HexToAddress(args[0])
			cli.PrintTransFn(cmd)(transSession.Add(broker))
		},
	},
	&cobra.Command{
		Use:    "addAccount <broker> <acccount>",
		Short:  "Adds the custodial-<account> to the <broker>",
		Args:   cli.ChainArgs(cli.AddressArgFunc("broker", 0), cli.AddressArgFunc("account", 1)),
		PreRun: connectTransactor,
		Run: func(cmd *cobra.Command, args []string) {
			broker := common.HexToAddress(args[0])
			custodialAccount := common.HexToAddress(args[1])
			cli.PrintTransFn(cmd)(transSession.AddAccount(broker, custodialAccount))
		},
	},
	&cobra.Command{
		Use:    "remove <broker>",
		Short:  "Removes the <broker> address",
		Args:   cli.ChainArgs(cli.AddressArgFunc("broker", 0)),
		PreRun: connectTransactor,
		Run: func(cmd *cobra.Command, args []string) {
			addr := common.HexToAddress(args[0])
			cli.PrintTransFn(cmd)(transSession.Remove(addr))
		},
	},
	&cobra.Command{
		Use:    "removeAccount <acccount>",
		Short:  "Removes the custodial-<account> of a broker",
		Args:   cli.AddressArgFunc("account", 0),
		PreRun: connectTransactor,
		Run: func(cmd *cobra.Command, args []string) {
			custodialAccount := common.HexToAddress(args[0])
			cli.PrintTransFn(cmd)(transSession.RemoveAccount(custodialAccount))
		},
	},
	&cobra.Command{
		Use:    "setStorage <address>",
		Short:  "Sets the storage contract to <address>",
		Args:   cli.AddressArgFunc("address", 0),
		PreRun: connectTransactor,
		Run: func(cmd *cobra.Command, args []string) {
			addr := common.HexToAddress(args[0])
			cli.PrintTransFn(cmd)(transSession.SetStorage(addr))
		},
	},
}

func init() {
	// Add the Administrable, Destroyable, Lockable contract getter commands
	SetterCommands = append(SetterCommands, destroyable.NewSetterCommands(contractKey)...)
	SetterCommands = append(SetterCommands, lockable.NewSetterCommands(contractKey)...)
	SetterCommands = append(SetterCommands, ownable.NewSetterCommands(contractKey)...)

	for _, cmd := range SetterCommands {
		// Allow providing contract 'address' flag
		cmd.Flags().String("address", "", `address of the BrokerDealer registry contract (default "[`+contractKey+`] value from config")`)
	}
}
