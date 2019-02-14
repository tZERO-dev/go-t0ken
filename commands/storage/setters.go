package storage

import (
	"strconv"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/spf13/cobra"

	"github.com/tzero-dev/go-t0ken/cli"
	"github.com/tzero-dev/go-t0ken/commands/destroyable"
	"github.com/tzero-dev/go-t0ken/commands/lockable"
	"github.com/tzero-dev/go-t0ken/commands/ownable"
)

var SetterCommands = []*cobra.Command{
	&cobra.Command{
		Use:    "addAccount <address> <kind> <frozen>, <parent>",
		Short:  "Adds the <address> under the <kind>, set as <frozen>, for the <parent>",
		Args:   cli.ChainArgs(cli.AddressArgFunc("address", 0), cli.UintArgFunc("kind", 1, 8), cli.BoolArgFunc("frozen", 2), cli.AddressArgFunc("parent", 3)),
		PreRun: connectTransactor,
		Run: func(cmd *cobra.Command, args []string) {
			addr := common.HexToAddress(args[0])
			kind, _ := strconv.ParseInt(args[1], 10, 8)
			frozen, _ := strconv.ParseBool(args[2])
			parent := common.HexToAddress(args[3])
			cli.PrintTransFn(cmd)(transSession.AddAccount(addr, uint8(kind), frozen, parent))
		},
	},
	&cobra.Command{
		Use:    "grantPermission <kind> <address>",
		Short:  "Grants <kind> permissions for <address>",
		Args:   cli.ChainArgs(cli.UintArgFunc("kind", 0, 8), cli.AddressArgFunc("address", 1)),
		PreRun: connectTransactor,
		Run: func(cmd *cobra.Command, args []string) {
			kind, _ := strconv.ParseInt(args[0], 10, 8)
			addr := common.HexToAddress(args[1])
			cli.PrintTransFn(cmd)(transSession.GrantPermission(uint8(kind), addr))
		},
	},
	&cobra.Command{
		Use:    "removeAccount <address>",
		Short:  "Removes the <address>",
		Args:   cli.ChainArgs(cli.AddressArgFunc("address", 0)),
		PreRun: connectTransactor,
		Run: func(cmd *cobra.Command, args []string) {
			addr := common.HexToAddress(args[0])
			cli.PrintTransFn(cmd)(transSession.RemoveAccount(addr))
		},
	},
	&cobra.Command{
		Use:    "revokePermission <kind> <address>",
		Short:  "Revokes the <kind> permissions from the <address>",
		Args:   cli.ChainArgs(cli.UintArgFunc("kind", 0, 8), cli.AddressArgFunc("address", 1)),
		PreRun: connectTransactor,
		Run: func(cmd *cobra.Command, args []string) {
			kind, _ := strconv.ParseInt(args[0], 10, 8)
			addr := common.HexToAddress(args[1])
			cli.PrintTransFn(cmd)(transSession.RevokePermission(uint8(kind), addr))
		},
	},
	&cobra.Command{
		Use:    "setAccountData <address> <index> <data>",
		Short:  "Sets the hex <data>, excluding '0x' prefix, at the <index> for the <address>",
		Args:   cli.ChainArgs(cli.AddressArgFunc("address", 0), cli.UintArgFunc("index", 1, 8), cli.HexArgFunc("data", 2, 32)),
		PreRun: connectTransactor,
		Run: func(cmd *cobra.Command, args []string) {
			addr := common.HexToAddress(args[0])
			index, _ := strconv.ParseInt(args[0], 10, 8)
			h, _ := hexutil.Decode(args[2])

			// Convert data to fixed size byte array
			var data [32]byte
			copy(data[:], h)

			cli.PrintTransFn(cmd)(transSession.SetAccountData(addr, uint8(index), data))
		},
	},
	&cobra.Command{
		Use:    "setAccountFrozen <address> <frozen>",
		Short:  "Set the <address> state to <frozen>",
		Args:   cli.ChainArgs(cli.AddressArgFunc("address", 0), cli.BoolArgFunc("frozen", 1)),
		PreRun: connectTransactor,
		Run: func(cmd *cobra.Command, args []string) {
			addr := common.HexToAddress(args[0])
			frozen, _ := strconv.ParseBool(args[1])
			cli.PrintTransFn(cmd)(transSession.SetAccountFrozen(addr, frozen))
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
