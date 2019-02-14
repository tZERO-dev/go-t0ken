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
		Use:    "approve <address> <tokens>",
		Short:  "Approves <address> to transfer <tokens> on your behalf",
		Args:   cli.ChainArgs(cobra.MaximumNArgs(2), cli.AddressArgFunc("address", 0), cli.BigIntArgFunc("tokens", 1)),
		PreRun: connectTransactor,
		Run: func(cmd *cobra.Command, args []string) {
			addr := common.HexToAddress(args[0])
			quantity, _ := new(big.Int).SetString(args[1], 10)
			cli.PrintTransFn(cmd)(transSession.Approve(addr, quantity))
		},
	},
	&cobra.Command{
		Use:    "cancelAndReissue <originalAddress> <replacementAddress>",
		Short:  "Cancels the <originalAddress> and replaces it with <replacemenmtAddress>",
		Args:   cli.ChainArgs(cobra.ExactArgs(2), cli.AddressArgFunc("originalAddress", 0), cli.AddressArgFunc("replacementAddress", 1)),
		PreRun: connectTransactor,
		Run: func(cmd *cobra.Command, args []string) {
			addr := common.HexToAddress(args[0])
			replacement := common.HexToAddress(args[1])
			cli.PrintTransFn(cmd)(transSession.CancelAndReissue(addr, replacement))
		},
	},
	&cobra.Command{
		Use:    "finishIssuing",
		Short:  "Finishes issuing for the token (can't be undone)",
		Args:   cobra.NoArgs,
		PreRun: connectTransactor,
		Run: func(cmd *cobra.Command, args []string) {
			cli.PrintTransFn(cmd)(transSession.FinishIssuing())
		},
	},
	&cobra.Command{
		Use:    "issueTokens <quantity>",
		Short:  "Issues <quantity> of tokens to the issuer",
		Args:   cli.IntArgFunc("quantity", 0),
		PreRun: connectTransactor,
		Run: func(cmd *cobra.Command, args []string) {
			quantity, _ := new(big.Int).SetString(args[0], 10)
			cli.PrintTransFn(cmd)(transSession.IssueTokens(quantity))
		},
	},
	&cobra.Command{
		Use:    "setCompliance <address>",
		Short:  "Sets the compliance contract <address>",
		Args:   cli.ChainArgs(cobra.ExactArgs(1), cli.AddressArgFunc("address", 0)),
		PreRun: connectTransactor,
		Run: func(cmd *cobra.Command, args []string) {
			addr := common.HexToAddress(args[0])
			cli.PrintTransFn(cmd)(transSession.SetCompliance(addr))
		},
	},
	&cobra.Command{
		Use:    "setIssuer <address>",
		Short:  "Sets the issuer to the <address>",
		Args:   cli.ChainArgs(cobra.ExactArgs(1), cli.AddressArgFunc("address", 0)),
		PreRun: connectTransactor,
		Run: func(cmd *cobra.Command, args []string) {
			addr := common.HexToAddress(args[0])
			cli.PrintTransFn(cmd)(transSession.SetIssuer(addr))
		},
	},
	&cobra.Command{
		Use:    "transfer <address> <quantity>",
		Short:  "Transfers <quantity> of tokens from your address to <address>",
		Args:   cli.ChainArgs(cobra.MaximumNArgs(2), cli.AddressArgFunc("address", 0), cli.BigIntArgFunc("quantity", 1)),
		PreRun: connectTransactor,
		Run: func(cmd *cobra.Command, args []string) {
			addr := common.HexToAddress(args[0])
			qty, _ := new(big.Int).SetString(args[1], 10)
			cli.PrintTransFn(cmd)(transSession.Transfer(addr, qty))
		},
	},
	&cobra.Command{
		Use:    "transferFrom <sender> <recipient> <quantity>",
		Short:  "Transfers <quantity> of tokens from the <sender> address to the <recipient> address (requires approval)",
		Args:   cli.ChainArgs(cobra.MaximumNArgs(3), cli.AddressArgFunc("sender", 0), cli.AddressArgFunc("recipient", 1), cli.BigIntArgFunc("quantity", 2)),
		PreRun: connectTransactor,
		Run: func(cmd *cobra.Command, args []string) {
			sender := common.HexToAddress(args[0])
			recipient := common.HexToAddress(args[1])
			qty, _ := new(big.Int).SetString(args[2], 10)
			cli.PrintTransFn(cmd)(transSession.TransferFrom(sender, recipient, qty))
		},
	},
	&cobra.Command{
		Use:    "transferOverride <sender> <recipient> <quantity>",
		Short:  "Transfers <quantity> of tokens from the <sender> address to the <recipient> address (invoker must be an admin, and t0ken must have compliance set)",
		Args:   cli.ChainArgs(cobra.MaximumNArgs(3), cli.AddressArgFunc("sender", 0), cli.AddressArgFunc("recipient", 1), cli.BigIntArgFunc("quantity", 2)),
		PreRun: connectTransactor,
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
