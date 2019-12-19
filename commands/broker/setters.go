package broker

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/spf13/cobra"

	"github.com/tzero-dev/go-t0ken/cli"
	"github.com/tzero-dev/go-t0ken/commands/destroyable"
	"github.com/tzero-dev/go-t0ken/commands/gas"
	"github.com/tzero-dev/go-t0ken/commands/lockable"
	"github.com/tzero-dev/go-t0ken/commands/nonce"
	"github.com/tzero-dev/go-t0ken/commands/ownable"
)

var SetterCommands = []*cobra.Command{
	&cobra.Command{
		Use:     "add <broker> <hash>",
		Short:   "Adds the <broker> <hash> to the registry",
		Example: "t0ken broker add 0xb01ba0d19cc9cd613253bad489b69e583dbfd4da 0x6ea3939dd4721de598d003ee4796ea2a59ab095cf6b28d49fadd43a5e34db548 --keystoreAddress custodian",
		Args:    cli.ChainArgs(cli.AddressArgFunc("broker", 0), cli.HexArgLenFunc("hash", 1, 16)),
		PreRun:  connectTransactor,
		Run: func(cmd *cobra.Command, args []string) {
			broker := common.HexToAddress(args[0])
			hash, _ := cli.Bytes32FromArg(args[1])
			cli.PrintTransactionFn(cmd)(transSession.Add(broker, hash))
		},
	},
	&cobra.Command{
		Use:     "addAccount <broker> <acccount>",
		Short:   "Adds the custodial-<account> to the <broker>",
		Example: "t0ken broker addAccount 0xb01ba0d19cc9cd613253bad489b69e583dbfd4da 0xa01a0a93716633058d69a28fbd472fd40e7c6b79 --keystoreAddress custodian",
		Args:    cli.ChainArgs(cli.AddressArgFunc("broker", 0), cli.AddressArgFunc("account", 1)),
		PreRun:  connectTransactor,
		Run: func(cmd *cobra.Command, args []string) {
			broker := common.HexToAddress(args[0])
			custodialAccount := common.HexToAddress(args[1])
			cli.PrintTransactionFn(cmd)(transSession.AddAccount(broker, custodialAccount))
		},
	},
	&cobra.Command{
		Use:     "remove <broker>",
		Short:   "Removes the <broker> address",
		Example: "t0ken broker remove 0xb01ba0d19cc9cd613253bad489b69e583dbfd4da --keystoreAddress custodian",
		Args:    cli.ChainArgs(cli.AddressArgFunc("broker", 0)),
		PreRun:  connectTransactor,
		Run: func(cmd *cobra.Command, args []string) {
			addr := common.HexToAddress(args[0])
			cli.PrintTransactionFn(cmd)(transSession.Remove(addr))
		},
	},
	&cobra.Command{
		Use:     "removeAccount <acccount>",
		Short:   "Removes the custodial-<account> of a broker",
		Example: "t0ken broker removeAccount 0xa01a0a93716633058d69a28fbd472fd40e7c6b79 --keystoreAddress custodian",
		Args:    cli.AddressArgFunc("account", 0),
		PreRun:  connectTransactor,
		Run: func(cmd *cobra.Command, args []string) {
			custodialAccount := common.HexToAddress(args[0])
			cli.PrintTransactionFn(cmd)(transSession.RemoveAccount(custodialAccount))
		},
	},
	&cobra.Command{
		Use:     "setRegistry <address>",
		Short:   "Sets the registry contract to <address>",
		Example: "t0ken broker setRegistry 0x397e7b9c15ff22ba67ec6e78f46f1e21540bcb36 --keystoreAddress owner",
		Args:    cli.AddressArgFunc("address", 0),
		PreRun:  connectTransactor,
		Run: func(cmd *cobra.Command, args []string) {
			addr := common.HexToAddress(args[0])
			cli.PrintTransactionFn(cmd)(transSession.SetRegistry(addr))
		},
	},
}

func init() {
	// Add the Destroyable, Lockable, Ownable contract getter commands
	SetterCommands = append(SetterCommands, destroyable.NewSetterCommands(contractKey)...)
	SetterCommands = append(SetterCommands, lockable.NewSetterCommands(contractKey)...)
	SetterCommands = append(SetterCommands, ownable.NewSetterCommands(contractKey)...)

	for _, cmd := range SetterCommands {
		// Add the 'gasPrice', 'nonce' and 'wait' flags to the deploy function
		gas.Flag(cmd)
		nonce.Flag(cmd)
		cli.WaitFlag(cmd)

		// Allow providing contract 'address' flag
		cmd.Flags().String("address", "", `address of the BrokerDealer registry contract (default "[`+contractKey+`] value from config")`)
	}
}
