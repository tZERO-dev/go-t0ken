package token

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/tzero-dev/go-t0ken/cli"
	"github.com/tzero-dev/go-t0ken/contracts/token/erc20"
)

var GetterCommands = []*cobra.Command{
	&cobra.Command{
		Use:   "abi",
		Short: "Outputs the T0ken ABI",
		Args:  cobra.NoArgs,
		Run:   func(cmd *cobra.Command, args []string) { cmd.Println(erc20.T0kenABI) },
	},
	&cobra.Command{
		Use:   "bin",
		Short: "Outputs the T0ken Binary",
		Args:  cobra.NoArgs,
		Run:   func(cmd *cobra.Command, args []string) { cmd.Println(erc20.T0kenBin) },
	},
	&cobra.Command{
		Use:    "owner",
		Short:  "Gets the owner address of the t0ken",
		Args:   cobra.NoArgs,
		PreRun: connectWithT0kenCallerSession,
		Run: func(cmd *cobra.Command, args []string) {
			cli.CheckAddressGetter(cmd)(callSession.Owner())
		},
	},
	&cobra.Command{
		Use:   "balanceOf <address>",
		Short: "gets the balance of the given <address>",
		Args:  cli.ChainArgs(cobra.MaximumNArgs(1), cli.AddressArgFunc(0)),
		Run: func(cmd *cobra.Command, args []string) {
			addr := common.HexToAddress(args[0])
			cli.CheckGetter(cmd)(callSession.BalanceOf(addr))
		},
	},
	&cobra.Command{
		Use:   "compliance",
		Short: "Gets the compliance contract address for the t0ken",
		Args:  cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {
			cli.CheckAddressGetter(cmd)(callSession.Compliance())
		},
	},
	&cobra.Command{
		Use:   "decimals",
		Short: "Gets the number of decimals the t0ken is set to",
		Args:  cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {
			cli.CheckGetter(cmd)(callSession.Decimals())
		},
	},
	&cobra.Command{
		Use:   "superseded",
		Short: "Gets the superseded address of the given address",
		Args:  cli.ChainArgs(cobra.MaximumNArgs(1), cli.AddressArgFunc(0)),
		Run: func(cmd *cobra.Command, args []string) {
			addr := common.HexToAddress(args[0])
			cli.CheckAddressGetter(cmd)(callSession.GetSuperseded(addr))
		},
	},
	&cobra.Command{
		Use:   "holderAt <index>",
		Short: "Gets the holder address at the given <index>",
		Args:  cli.IntArgFunc(0),
		Run: func(cmd *cobra.Command, args []string) {
			index, _ := new(big.Int).SetString(args[0], 10)
			cli.CheckAddressGetter(cmd)(callSession.HolderAt(index))
		},
	},
	&cobra.Command{
		Use:   "issuer",
		Short: "Gets the issuer of the t0ken",
		Args:  cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {
			cli.CheckAddressGetter(cmd)(callSession.Issuer())
		},
	},
	&cobra.Command{
		Use:   "issuingFinished",
		Short: "Returns if issuing has been finished",
		Args:  cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {
			cli.CheckGetter(cmd)(callSession.IssuingFinished())
		},
	},
	&cobra.Command{
		Use:   "name",
		Short: "Gets the name of the t0ken",
		Args:  cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {
			cli.CheckGetter(cmd)(callSession.Name())
		},
	},
	&cobra.Command{
		Use:   "symbol",
		Short: "Gets the symbol of the t0ken",
		Args:  cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {
			cli.CheckGetter(cmd)(callSession.Symbol())
		},
	},
	&cobra.Command{
		Use:   "totalSupply",
		Short: "Gets the total supply",
		Args:  cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {
			cli.CheckGetter(cmd)(callSession.TotalSupply())
		},
	},
}

func init() {
	for i, cmd := range GetterCommands {
		// Skip ABI/Bin
		if i < 2 {
			continue
		}

		// Add pre-run connect
		cmd.PreRun = connectWithT0kenCallerSession

		// Allow providing contract 'address' flag
		cmd.Flags().StringVar(&address, "address", "", `address of the T0ken contract (default "[t0ken] value from config")`)
		viper.BindPFlag("address", cmd.Flags().Lookup("address"))
	}
}
