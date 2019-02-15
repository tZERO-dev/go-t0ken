package token

import (
	"math/big"
	//"strconv"

	"github.com/ethereum/go-ethereum/common"
	"github.com/spf13/cobra"

	"github.com/tzero-dev/go-t0ken/cli"
	"github.com/tzero-dev/go-t0ken/commands/administrable"
	"github.com/tzero-dev/go-t0ken/commands/destroyable"
	"github.com/tzero-dev/go-t0ken/commands/gas"
	"github.com/tzero-dev/go-t0ken/commands/lockable"
	"github.com/tzero-dev/go-t0ken/commands/nonce"
	"github.com/tzero-dev/go-t0ken/commands/ownable"
)

var SetterCommands = []*cobra.Command{
	&cobra.Command{
		Use:     "approve <address> <tokens>",
		Short:   "Approves <address> to transfer <tokens> on your behalf",
		Example: "t0ken token approve 0xf01ff29dcbee147e9ca151a281bfdf136f66a45b 0xa01a0a93716633058d69a28fbd472fd40e7c6b79 --keystoreAddress 0xf01ff29dcbee147e9ca151a281bfdf136f66a45b",
		Args:    cli.ChainArgs(cobra.MaximumNArgs(2), cli.AddressArgFunc("address", 0), cli.BigIntArgFunc("tokens", 1)),
		PreRun:  connectTransactor,
		Run: func(cmd *cobra.Command, args []string) {
			addr := common.HexToAddress(args[0])
			quantity, _ := new(big.Int).SetString(args[1], 10)
			cli.PrintTransFn(cmd)(transSession.Approve(addr, quantity))
		},
	},
	&cobra.Command{
		Use:     "cancelAndReissue <originalAddress> <replacementAddress>",
		Short:   "Cancels the <originalAddress> and replaces it with <replacemenmtAddress>",
		Example: "t0ken token cancelAndReissue 0xf01ff29dcbee147e9ca151a281bfdf136f66a45b 0xa01a0a93716633058d69a28fbd472fd40e7c6b79 --keystoreAddress issuer",
		Args:    cli.ChainArgs(cobra.ExactArgs(2), cli.AddressArgFunc("originalAddress", 0), cli.AddressArgFunc("replacementAddress", 1)),
		PreRun:  connectTransactor,
		Run: func(cmd *cobra.Command, args []string) {
			addr := common.HexToAddress(args[0])
			replacement := common.HexToAddress(args[1])
			cli.PrintTransFn(cmd)(transSession.CancelAndReissue(addr, replacement))
		},
	},
	&cobra.Command{
		Use:     "finishIssuing",
		Short:   "Finishes issuing for the token (can't be undone)",
		Example: "t0ken token finishIssuing --keystoreAddress issuer",
		Args:    cobra.NoArgs,
		PreRun:  connectTransactor,
		Run: func(cmd *cobra.Command, args []string) {
			cli.PrintTransFn(cmd)(transSession.FinishIssuing())
		},
	},
	&cobra.Command{
		Use:     "issueTokens <quantity>",
		Short:   "Issues <quantity> of tokens to the issuer",
		Example: "t0ken token issueTokens 5000 --keystoreAddress issuer",
		Args:    cli.IntArgFunc("quantity", 0),
		PreRun:  connectTransactor,
		Run: func(cmd *cobra.Command, args []string) {
			quantity, _ := new(big.Int).SetString(args[0], 10)
			cli.PrintTransFn(cmd)(transSession.IssueTokens(quantity))
		},
	},
	&cobra.Command{
		Use:     "setCompliance <address>",
		Short:   "Sets the compliance contract <address>",
		Example: "t0ken token setCompliance 0x397e7b9c15ff22ba67ec6e78f46f1e21540bcb36 --keystoreAddress owner",
		Args:    cli.ChainArgs(cobra.ExactArgs(1), cli.AddressArgFunc("address", 0)),
		PreRun:  connectTransactor,
		Run: func(cmd *cobra.Command, args []string) {
			addr := common.HexToAddress(args[0])
			cli.PrintTransFn(cmd)(transSession.SetCompliance(addr))
		},
	},
	&cobra.Command{
		Use:     "setIssuer <address>",
		Short:   "Sets the issuer to the <address>",
		Example: "t0ken token setIssuer 0xf01ff29dcbee147e9ca151a281bfdf136f66a45b --keystoreAddress owner",
		Args:    cli.ChainArgs(cobra.ExactArgs(1), cli.AddressArgFunc("address", 0)),
		PreRun:  connectTransactor,
		Run: func(cmd *cobra.Command, args []string) {
			addr := common.HexToAddress(args[0])
			cli.PrintTransFn(cmd)(transSession.SetIssuer(addr))
		},
	},
	&cobra.Command{
		Use:     "transfer <address> <quantity>",
		Short:   "Transfers <quantity> of tokens from your address to <address>",
		Example: "t0ken token transfer 0xf01ff29dcbee147e9ca151a281bfdf136f66a45b 15 --keystoreAddress 0xa01a0a93716633058d69a28fbd472fd40e7c6b79",
		Args:    cli.ChainArgs(cobra.MaximumNArgs(2), cli.AddressArgFunc("address", 0), cli.BigIntArgFunc("quantity", 1)),
		PreRun:  connectTransactor,
		Run: func(cmd *cobra.Command, args []string) {
			addr := common.HexToAddress(args[0])
			qty, _ := new(big.Int).SetString(args[1], 10)
			cli.PrintTransFn(cmd)(transSession.Transfer(addr, qty))
		},
	},
	&cobra.Command{
		Use:     "transferFrom <sender> <recipient> <quantity>",
		Short:   "Transfers <quantity> of tokens from the <sender> address to the <recipient> address (requires approval)",
		Example: "t0ken token transferFrom 0xf01ff29dcbee147e9ca151a281bfdf136f66a45b 0xf02f537578d03f6aece28f249eac19542d848f20 15 --keystoreAddress 0xa01a0a93716633058d69a28fbd472fd40e7c6b79",
		Args:    cli.ChainArgs(cobra.MaximumNArgs(3), cli.AddressArgFunc("sender", 0), cli.AddressArgFunc("recipient", 1), cli.BigIntArgFunc("quantity", 2)),
		PreRun:  connectTransactor,
		Run: func(cmd *cobra.Command, args []string) {
			sender := common.HexToAddress(args[0])
			recipient := common.HexToAddress(args[1])
			qty, _ := new(big.Int).SetString(args[2], 10)
			cli.PrintTransFn(cmd)(transSession.TransferFrom(sender, recipient, qty))
		},
	},
	&cobra.Command{
		Use:     "transferOverride <sender> <recipient> <quantity>",
		Short:   "Transfers <quantity> of tokens from the <sender> address to the <recipient> address (invoker must be an admin, and t0ken must have compliance set)",
		Example: "t0ken token transferOverride 0xf01ff29dcbee147e9ca151a281bfdf136f66a45b 0xf02f537578d03f6aece28f249eac19542d848f20 15 --keystoreAddress admin",
		Args:    cli.ChainArgs(cobra.MaximumNArgs(3), cli.AddressArgFunc("sender", 0), cli.AddressArgFunc("recipient", 1), cli.BigIntArgFunc("quantity", 2)),
		PreRun:  connectTransactor,
		Run: func(cmd *cobra.Command, args []string) {
			sender := common.HexToAddress(args[0])
			recipient := common.HexToAddress(args[1])
			qty, _ := new(big.Int).SetString(args[2], 10)
			cli.PrintTransFn(cmd)(transSession.TransferOverride(sender, recipient, qty))
		},
	},
}

func init() {
	// Add the Administrable, Destroyable, Lockable contract getter commands
	SetterCommands = append(SetterCommands, administrable.NewSetterCommands(contractKey)...)
	SetterCommands = append(SetterCommands, destroyable.NewSetterCommands(contractKey)...)
	SetterCommands = append(SetterCommands, ownable.NewSetterCommands(contractKey)...)
	SetterCommands = append(SetterCommands, lockable.NewSetterCommands(contractKey)...)

	for _, cmd := range SetterCommands {
		// Allow 'gasPrice' and 'nonce' flags
		gas.Flag(cmd)
		nonce.Flag(cmd)

		// Allow providing contract 'address' flag
		cmd.Flags().String("address", "", `address of the BrokerDealer registry contract (default "[`+contractKey+`] value from config")`)
	}
}
