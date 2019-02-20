package custodian

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
		Use:     "add <custodian>",
		Short:   "Adds the <custodian> address",
		Example: "t0ken custodian add 0xc01c68ac7dd1cc48f95e3187a33a170fdb9023e7 --keystoreAddress owner",
		Args:    cli.AddressArgFunc("custodian", 0),
		PreRun:  connectTransactor,
		Run: func(cmd *cobra.Command, args []string) {
			custodian := common.HexToAddress(args[0])
			cli.PrintTransFn(cmd)(transSession.Add(custodian))
		},
	},
	&cobra.Command{
		Use:     "remove <custodian>",
		Short:   "Removes the <custodian> address",
		Example: "t0ken custodian remove 0xc01c68ac7dd1cc48f95e3187a33a170fdb9023e7 --keystoreAddress owner",
		Args:    cli.ChainArgs(cli.AddressArgFunc("custodian", 0)),
		PreRun:  connectTransactor,
		Run: func(cmd *cobra.Command, args []string) {
			addr := common.HexToAddress(args[0])
			cli.PrintTransFn(cmd)(transSession.Remove(addr))
		},
	},
	&cobra.Command{
		Use:     "setStorage <address>",
		Short:   "Sets the storage contract to <address>",
		Example: "t0ken custodian setStorage 0x397e7b9c15ff22ba67ec6e78f46f1e21540bcb36 --keystoreAddress owner",
		Args:    cli.AddressArgFunc("address", 0),
		PreRun:  connectTransactor,
		Run: func(cmd *cobra.Command, args []string) {
			addr := common.HexToAddress(args[0])
			cli.PrintTransFn(cmd)(transSession.SetStorage(addr))
		},
	},
}

func init() {
	// Add the Destroyable, Lockable, Ownable contract getter commands
	SetterCommands = append(SetterCommands, destroyable.NewSetterCommands(contractKey)...)
	SetterCommands = append(SetterCommands, lockable.NewSetterCommands(contractKey)...)
	SetterCommands = append(SetterCommands, ownable.NewSetterCommands(contractKey)...)

	for _, cmd := range SetterCommands {
		// Allow providing contract 'address' flag
		cmd.Flags().String("address", "", `address of the BrokerDealer registry contract (default "[`+contractKey+`] value from config")`)
		cmd.Flags().Int("wait", -1, "waits the provided number of seconds for the transaction to be mined ('0' waits indefinitely)")
	}
}
