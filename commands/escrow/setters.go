package escrow

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/spf13/cobra"

	"github.com/tzero-dev/go-t0ken/cli"
	"github.com/tzero-dev/go-t0ken/commands/administrable"
	"github.com/tzero-dev/go-t0ken/commands/gas"
	"github.com/tzero-dev/go-t0ken/commands/nonce"
)

var SetterCommands = []*cobra.Command{
	&cobra.Command{
		Use:     "setRegistry <address>",
		Short:   "Sets the registry contract to <address>",
		Example: "t0ken escrow setRegistry 0x397e7b9c15ff22ba67ec6e78f46f1e21540bcb36 --keystoreAddress owner",
		Args:    cli.AddressArgFunc("address", 0),
		PreRun:  connectTransactor,
		Run: func(cmd *cobra.Command, args []string) {
			addr := common.HexToAddress(args[0])
			cli.PrintTransactionFn(cmd)(transSession.SetRegistry(addr))
		},
	},
	&cobra.Command{
		Use:     "accept <token> <holder> <quantity> <grantee>",
		Short:   "Accepts the escrow for the <token>/<holder> pair, sending the <quantity> to the <grantee>",
		Example: "t0ken escrow accept 0x5bd5b4e1a2c9b12812795e7217201b78c8c10b78 0xf01ff29dcbee147e9ca151a281bfdf136f66a45b 5 0xa01a0a93716633058d69a28fbd472fd40e7c6b79 ",
		Args:    cli.ChainArgs(cli.AddressArgFunc("token", 0), cli.AddressArgFunc("holder", 1), cli.BigIntArgFunc("tokens", 2), cli.AddressArgFunc("grantee", 3)),
		PreRun:  connectTransactor,
		Run: func(cmd *cobra.Command, args []string) {
			token := common.HexToAddress(args[0])
			holder := common.HexToAddress(args[1])
			quantity, _ := new(big.Int).SetString(args[2], 10)
			grantee := common.HexToAddress(args[3])
			cli.PrintTransactionFn(cmd)(transSession.Accept(token, holder, quantity, grantee))
		},
	},
	&cobra.Command{
		Use:     "reject <token> <holder> <quantity>",
		Short:   "Rejects the escrow for the <token>/<holder> pair, sending the <quantity> back to the <holder>",
		Example: "t0ken escrow accept 0x5bd5b4e1a2c9b12812795e7217201b78c8c10b78 0xf01ff29dcbee147e9ca151a281bfdf136f66a45b 5",
		Args:    cli.ChainArgs(cli.AddressArgFunc("token", 0), cli.AddressArgFunc("holder", 1), cli.BigIntArgFunc("tokens", 2)),
		PreRun:  connectTransactor,
		Run: func(cmd *cobra.Command, args []string) {
			token := common.HexToAddress(args[0])
			holder := common.HexToAddress(args[1])
			quantity, _ := new(big.Int).SetString(args[2], 10)
			cli.PrintTransactionFn(cmd)(transSession.Reject(token, holder, quantity))
		},
	},
	&cobra.Command{
		Use:     "record <token> <holder> <quantity>",
		Short:   "Records the escrow for the <token>/<holder> pair for the <quantity> of tokens",
		Example: "t0ken escrow accept 0x5bd5b4e1a2c9b12812795e7217201b78c8c10b78 0xf01ff29dcbee147e9ca151a281bfdf136f66a45b 5",
		Args:    cli.ChainArgs(cli.AddressArgFunc("token", 0), cli.AddressArgFunc("holder", 1), cli.BigIntArgFunc("tokens", 2)),
		PreRun:  connectTransactor,
		Run: func(cmd *cobra.Command, args []string) {
			token := common.HexToAddress(args[0])
			holder := common.HexToAddress(args[1])
			quantity, _ := new(big.Int).SetString(args[2], 10)
			cli.PrintTransactionFn(cmd)(transSession.Record(token, holder, quantity))
		},
	},
}

func init() {
	// Add Administrable contract getter commands
	SetterCommands = append(SetterCommands, administrable.NewSetterCommands(contractKey)...)

	for _, cmd := range SetterCommands {
		// Add the 'gasPrice', 'nonce' and 'wait' flags to the deploy function
		gas.Flag(cmd)
		nonce.Flag(cmd)
		cli.WaitFlag(cmd)

		// Allow providing contract 'address' flag
		cmd.Flags().String("address", "", `address of the Escrow contract (default "[`+contractKey+`] value from config")`)
	}
}
