package token

import (
	"math/big"
	"strconv"

	"github.com/ethereum/go-ethereum/common"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/tzero-dev/go-t0ken/cli"
	"github.com/tzero-dev/go-t0ken/commands/gas"
	"github.com/tzero-dev/go-t0ken/commands/nonce"
)

var SetterCommands = []*cobra.Command{
	&cobra.Command{
		Use:   "approve <address> <tokens>",
		Short: "Approves <address> to transfer <tokens> on your behalf",
		Args:  cli.ChainArgs(cobra.MaximumNArgs(2), cli.AddressArgFunc(0), cli.BigIntArgFunc(1)),
		Run: func(cmd *cobra.Command, args []string) {
			addr := common.HexToAddress(args[0])
			quantity, _ := new(big.Int).SetString(args[1], 10)
			cli.PrintTransFn(cmd)(transSession.Approve(addr, quantity))
		},
	},
	&cobra.Command{
		Use:   "cancelAndReissue <originalAddress> <replacementAddress>",
		Short: "Cancels the <originalAddress> and replaces it with <replacemenmtAddress>",
		Args:  cli.ChainArgs(cobra.ExactArgs(2), cli.AddressArgFunc(0), cli.AddressArgFunc(1)),
		Run: func(cmd *cobra.Command, args []string) {
			addr := common.HexToAddress(args[0])
			replacement := common.HexToAddress(args[1])
			cli.PrintTransFn(cmd)(transSession.CancelAndReissue(addr, replacement))
		},
	},
	&cobra.Command{
		Use:   "finishIssuing",
		Short: "Finishes issuing for the token (can't be undone)",
		Args:  cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {
			cli.PrintTransFn(cmd)(transSession.FinishIssuing())
		},
	},
	&cobra.Command{
		Use:   "issueTokens <quantity>",
		Short: "Issues <quantity> of tokens to the issuer",
		Args:  cli.IntArgFunc(0),
		Run: func(cmd *cobra.Command, args []string) {
			quantity, _ := new(big.Int).SetString(args[0], 10)
			cli.PrintTransFn(cmd)(transSession.IssueTokens(quantity))
		},
	},
	&cobra.Command{
		Use:   "kill",
		Short: "Destroys the contract",
		Run: func(cmd *cobra.Command, args []string) {
			cli.PrintTransFn(cmd)(transSession.Kill())
		},
	},
	&cobra.Command{
		Use:   "setCompliance <address>",
		Short: "Sets the compliance contract <address>",
		Args:  cli.ChainArgs(cobra.ExactArgs(1), cli.AddressArgFunc(0)),
		Run: func(cmd *cobra.Command, args []string) {
			addr := common.HexToAddress(args[0])
			cli.PrintTransFn(cmd)(transSession.SetCompliance(addr))
		},
	},
	&cobra.Command{
		Use:   "setIssuer <address>",
		Short: "Sets the issuer to the <address>",
		Args:  cli.ChainArgs(cobra.ExactArgs(1), cli.AddressArgFunc(0)),
		Run: func(cmd *cobra.Command, args []string) {
			addr := common.HexToAddress(args[0])
			cli.PrintTransFn(cmd)(transSession.SetIssuer(addr))
		},
	},
	&cobra.Command{
		Use:   "setLocked <bool>",
		Short: "Locks/Unlocks the contract, preventing any transfers from occuring",
		Args:  cli.ChainArgs(cobra.ExactArgs(1), cli.BoolArgFunc(0)),
		Run: func(cmd *cobra.Command, args []string) {
			locked, _ := strconv.ParseBool(args[0])
			cli.PrintTransFn(cmd)(transSession.SetLocked(locked))
		},
	},
	&cobra.Command{
		Use:   "transfer <address> <quantity>",
		Short: "Transfers <quantity> of tokens from your address to <address>",
		Args:  cli.ChainArgs(cobra.MaximumNArgs(2), cli.AddressArgFunc(0), cli.BigIntArgFunc(1)),
		Run: func(cmd *cobra.Command, args []string) {
			addr := common.HexToAddress(args[0])
			qty, _ := new(big.Int).SetString(args[1], 10)
			cli.PrintTransFn(cmd)(transSession.Transfer(addr, qty))
		},
	},
	&cobra.Command{
		Use:   "transferFrom <sender> <recipient> <quantity>",
		Short: "Transfers <quantity> of tokens from the <sender> address to the <recipient> address (requires approval)",
		Args:  cli.ChainArgs(cobra.MaximumNArgs(3), cli.AddressArgFunc(0), cli.AddressArgFunc(1), cli.BigIntArgFunc(2)),
		Run: func(cmd *cobra.Command, args []string) {
			sender := common.HexToAddress(args[0])
			recipient := common.HexToAddress(args[1])
			qty, _ := new(big.Int).SetString(args[2], 10)
			cli.PrintTransFn(cmd)(transSession.TransferFrom(sender, recipient, qty))
		},
	},
	&cobra.Command{
		Use:   "transferOverride <sender> <recipient> <quantity>",
		Short: "Transfers <quantity> of tokens from the <sender> address to the <recipient> address (invoker must be an admin, and t0ken must have compliance set)",
		Args:  cli.ChainArgs(cobra.MaximumNArgs(3), cli.AddressArgFunc(0), cli.AddressArgFunc(1), cli.BigIntArgFunc(2)),
		Run: func(cmd *cobra.Command, args []string) {
			sender := common.HexToAddress(args[0])
			recipient := common.HexToAddress(args[1])
			qty, _ := new(big.Int).SetString(args[2], 10)
			cli.PrintTransFn(cmd)(transSession.TransferOverride(sender, recipient, qty))
		},
	},
	&cobra.Command{
		Use:   "transferOwner <address>",
		Short: "Transfers ownership of the contract to <address>",
		Args:  cli.ChainArgs(cobra.ExactArgs(1), cli.AddressArgFunc(0)),
		Run: func(cmd *cobra.Command, args []string) {
			addr := common.HexToAddress(args[0])
			cli.PrintTransFn(cmd)(transSession.TransferOwner(addr))
		},
	},
}

func init() {
	for _, cmd := range SetterCommands {
		// Add pre-run connect
		cmd.PreRun = connectWithT0kenTransactorSession

		// Allow providing contract 'address' flag
		cmd.Flags().StringVar(&address, "address", "", `address of the T0ken contract (default "[t0ken] value from config")`)
		viper.BindPFlag("address", cmd.Flags().Lookup("address"))

		// Allow 'gasPrice' and 'nonce' flags
		gas.Flag(cmd)
		nonce.Flag(cmd)
	}
}
