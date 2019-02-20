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
		Use:     "addAccount <address> <kind> <frozen>, <parent>",
		Short:   "Adds the <address> under the <kind>, set as <frozen>, for the <parent>",
		Example: "t0ken storage addAccount 0xf01ff29dcbee147e9ca151a281bfdf136f66a45b 4 false 0xb01ba0d19cc9cd613253bad489b69e583dbfd4da --keystoreAddress owner",
		Args:    cli.ChainArgs(cli.AddressArgFunc("address", 0), cli.UintArgFunc("kind", 1, 8), cli.BoolArgFunc("frozen", 2), cli.AddressArgFunc("parent", 3)),
		PreRun:  connectTransactor,
		Run: func(cmd *cobra.Command, args []string) {
			addr := common.HexToAddress(args[0])
			kind, _ := strconv.ParseInt(args[1], 10, 8)
			frozen, _ := strconv.ParseBool(args[2])
			parent := common.HexToAddress(args[3])
			cli.PrintTransFn(cmd)(transSession.AddAccount(addr, uint8(kind), frozen, parent))
		},
	},
	&cobra.Command{
		Use:     "grantPermission <kind> <address>",
		Short:   "Grants <kind> permissions for <address>",
		Example: "t0ken storage grantPermission 4 0xf01ff29dcbee147e9ca151a281bfdf136f66a45b --keystoreAddress owner",
		Args:    cli.ChainArgs(cli.UintArgFunc("kind", 0, 8), cli.AddressArgFunc("address", 1)),
		PreRun:  connectTransactor,
		Run: func(cmd *cobra.Command, args []string) {
			kind, _ := strconv.ParseInt(args[0], 10, 8)
			addr := common.HexToAddress(args[1])
			cli.PrintTransFn(cmd)(transSession.GrantPermission(uint8(kind), addr))
		},
	},
	&cobra.Command{
		Use:     "removeAccount <address>",
		Short:   "Removes the <address>",
		Example: "t0ken storage removeAccount 0xf01ff29dcbee147e9ca151a281bfdf136f66a45b --keystoreAddress owner",
		Args:    cli.ChainArgs(cli.AddressArgFunc("address", 0)),
		PreRun:  connectTransactor,
		Run: func(cmd *cobra.Command, args []string) {
			addr := common.HexToAddress(args[0])
			cli.PrintTransFn(cmd)(transSession.RemoveAccount(addr))
		},
	},
	&cobra.Command{
		Use:     "revokePermission <kind> <address>",
		Short:   "Revokes the <kind> permissions from the <address>",
		Example: "t0ken storage revokePermission 4 0xf01ff29dcbee147e9ca151a281bfdf136f66a45b --keystoreAddress owner",
		Args:    cli.ChainArgs(cli.UintArgFunc("kind", 0, 8), cli.AddressArgFunc("address", 1)),
		PreRun:  connectTransactor,
		Run: func(cmd *cobra.Command, args []string) {
			kind, _ := strconv.ParseInt(args[0], 10, 8)
			addr := common.HexToAddress(args[1])
			cli.PrintTransFn(cmd)(transSession.RevokePermission(uint8(kind), addr))
		},
	},
	&cobra.Command{
		Use:     "setAccountData <address> <index> <data>",
		Short:   "Sets the hex <data>, excluding '0x' prefix, at the <index> for the <address>",
		Example: "t0ken storage setAccountData 0xf01ff29dcbee147e9ca151a281bfdf136f66a45b 0xa1896382c22b03c562b0241324cfca19505cc5c78eb06751d9cee690e21ed6a1 --keystoreAddress owner",
		Args:    cli.ChainArgs(cli.AddressArgFunc("address", 0), cli.UintArgFunc("index", 1, 8), cli.HexArgFunc("data", 2, 32)),
		PreRun:  connectTransactor,
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
		Use:     "setAccountFrozen <address> <frozen>",
		Short:   "Set the <address> state to <frozen>",
		Example: "t0ken storage setAccountFrozen 0xf01ff29dcbee147e9ca151a281bfdf136f66a45b true --keystoreAddress owner",
		Args:    cli.ChainArgs(cli.AddressArgFunc("address", 0), cli.BoolArgFunc("frozen", 1)),
		PreRun:  connectTransactor,
		Run: func(cmd *cobra.Command, args []string) {
			addr := common.HexToAddress(args[0])
			frozen, _ := strconv.ParseBool(args[1])
			cli.PrintTransFn(cmd)(transSession.SetAccountFrozen(addr, frozen))
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
