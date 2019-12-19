package custodian

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
		Use:     "add <custodian> <hash>",
		Short:   "Adds the <custodian> <hash> to the registry",
		Example: "t0ken custodian add 0xc01c68ac7dd1cc48f95e3187a33a170fdb9023e7 0x6ea3939dd4721de598d003ee4796ea2a59ab095cf6b28d49fadd43a5e34db548 --keystoreAddress owner",
		Args:    cli.ChainArgs(cli.AddressArgFunc("custodian", 0), cli.HexArgLenFunc("hash", 1, 16)),
		PreRun:  connectTransactor,
		Run: func(cmd *cobra.Command, args []string) {
			custodian := common.HexToAddress(args[0])
			hash, _ := cli.Bytes32FromArg(args[1])
			cli.PrintTransactionFn(cmd)(transSession.Add(custodian, hash))
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
			cli.PrintTransactionFn(cmd)(transSession.Remove(addr))
		},
	},
	&cobra.Command{
		Use:     "setRegistry <address>",
		Short:   "Sets the registry contract to <address>",
		Example: "t0ken custodian setRegistry 0x397e7b9c15ff22ba67ec6e78f46f1e21540bcb36 --keystoreAddress owner",
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
		cmd.Flags().String("address", "", `address of the Custodian registry contract (default "[`+contractKey+`] value from config")`)
	}
}
