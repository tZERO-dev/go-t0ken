package token

import (
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/spf13/cobra"

	"github.com/tzero-dev/go-t0ken/cli"
	"github.com/tzero-dev/go-t0ken/commands/lockable"
	"github.com/tzero-dev/go-t0ken/commands/ownable"
	t "github.com/tzero-dev/go-t0ken/contracts/token"
)

var GetterCommands = []*cobra.Command{
	&cobra.Command{
		Use:     "abi",
		Short:   "Outputs the T0ken ABI",
		Example: "t0ken investor abi",
		Args:    cobra.NoArgs,
		Run:     func(cmd *cobra.Command, args []string) { cmd.Println(t.T0kenABI) },
	},
	&cobra.Command{
		Use:     "bin",
		Short:   "Outputs the T0ken Binary",
		Example: "t0ken investor bin",
		Args:    cobra.NoArgs,
		Run:     func(cmd *cobra.Command, args []string) { cmd.Println(t.T0kenBin) },
	},
	&cobra.Command{
		Use:     "allowance <owner> <spender>",
		Short:   "Gets the amount of tokens the <owner> has approved the <spender> to transfer",
		Example: "t0ken token allowance 0xf01ff29dcbee147e9ca151a281bfdf136f66a45b 0xa01a0a93716633058d69a28fbd472fd40e7c6b79",
		Args:    cli.ChainArgs(cobra.ExactArgs(2), cli.AddressArgFunc("owner", 0), cli.AddressArgFunc("spender", 1)),
		PreRun:  connectCaller,
		Run: func(cmd *cobra.Command, args []string) {
			owner := common.HexToAddress(args[0])
			spender := common.HexToAddress(args[1])
			cli.CheckGetter(cmd)(callSession.Allowance(owner, spender))
		},
	},
	&cobra.Command{
		Use:     "balanceOf <address>",
		Short:   "Gets the balance of the given <address>",
		Example: "t0ken token balanceOf 0xf01ff29dcbee147e9ca151a281bfdf136f66a45b",
		Args:    cli.ChainArgs(cobra.MaximumNArgs(1), cli.AddressArgFunc("address", 0)),
		PreRun:  connectCaller,
		Run: func(cmd *cobra.Command, args []string) {
			addr := common.HexToAddress(args[0])
			cli.CheckGetter(cmd)(callSession.BalanceOf(addr))
		},
	},
	&cobra.Command{
		Use:     "compliance",
		Short:   "Gets the compliance contract address for the t0ken",
		Example: "t0ken token compliance",
		Args:    cobra.NoArgs,
		PreRun:  connectCaller,
		Run: func(cmd *cobra.Command, args []string) {
			cli.CheckAddressGetter(cmd)(callSession.Compliance())
		},
	},
	&cobra.Command{
		Use:     "decimals",
		Short:   "Gets the number of decimals the t0ken is set to",
		Example: "t0ken token decimals",
		Args:    cobra.NoArgs,
		PreRun:  connectCaller,
		Run: func(cmd *cobra.Command, args []string) {
			cli.CheckGetter(cmd)(callSession.Decimals())
		},
	},
	&cobra.Command{
		Use:     "holderAt <index>",
		Short:   "Gets the holder address at the given <index>",
		Example: "t0ken token holderAt 5",
		Args:    cli.ChainArgs(cobra.MaximumNArgs(1), cli.IntArgFunc("index", 0)),
		PreRun:  connectCaller,
		Run: func(cmd *cobra.Command, args []string) {
			index, _ := new(big.Int).SetString(args[0], 10)
			cli.CheckAddressGetter(cmd)(callSession.HolderAt(index))
		},
	},
	&cobra.Command{
		Use:     "isHolder <address>",
		Short:   "Checks if the given <address> is a current holder",
		Example: "t0ken token isHolder 0xf01ff29dcbee147e9ca151a281bfdf136f66a45b",
		Args:    cli.ChainArgs(cobra.MaximumNArgs(1), cli.AddressArgFunc("address", 0)),
		PreRun:  connectCaller,
		Run: func(cmd *cobra.Command, args []string) {
			addr := common.HexToAddress(args[0])
			cli.CheckGetter(cmd)(callSession.IsHolder(addr))
		},
	},
	&cobra.Command{
		Use:     "issuer",
		Short:   "Gets the issuer of the t0ken",
		Example: "t0ken token issuer",
		Args:    cobra.NoArgs,
		PreRun:  connectCaller,
		Run: func(cmd *cobra.Command, args []string) {
			cli.CheckAddressGetter(cmd)(callSession.Issuer())
		},
	},
	&cobra.Command{
		Use:     "issuanceFinished",
		Short:   "Returns if issuance has been finished",
		Example: "t0ken token issuanceFinished",
		Args:    cobra.NoArgs,
		PreRun:  connectCaller,
		Run: func(cmd *cobra.Command, args []string) {
			cli.CheckGetter(cmd)(callSession.IssuanceFinished())
		},
	},
	&cobra.Command{
		Use:     "name",
		Short:   "Gets the name of the t0ken",
		Example: "t0ken token name",
		Args:    cobra.NoArgs,
		PreRun:  connectCaller,
		Run: func(cmd *cobra.Command, args []string) {
			cli.CheckGetter(cmd)(callSession.Name())
		},
	},
	&cobra.Command{
		Use:     "holders",
		Short:   "Gets the total number of holders",
		Example: "t0ken token holders",
		Args:    cobra.NoArgs,
		PreRun:  connectCaller,
		Run: func(cmd *cobra.Command, args []string) {
			cli.CheckGetter(cmd)(callSession.Holders())
		},
	},
	&cobra.Command{
		Use:     "symbol",
		Short:   "Gets the symbol of the t0ken",
		Example: "t0ken token symbol",
		Args:    cobra.NoArgs,
		PreRun:  connectCaller,
		Run: func(cmd *cobra.Command, args []string) {
			cli.CheckGetter(cmd)(callSession.Symbol())
		},
	},
	&cobra.Command{
		Use:     "totalSupply",
		Short:   "Gets the total supply",
		Example: "t0ken token totalSupply",
		Args:    cobra.NoArgs,
		PreRun:  connectCaller,
		Run: func(cmd *cobra.Command, args []string) {
			cli.CheckGetter(cmd)(callSession.TotalSupply())
		},
	},
}

var FilterCommands = []*cobra.Command{
	&cobra.Command{
		Use:   "filterTransfer",
		Short: "Filters transfer events within the given block range",
		Example: `t0ken token filterTransfer --start 8658083
t0ken token filterTransfer --start 8658083 --end 8700000
t0ken token filterTransfer --from 0xf01fF29DCbEE147e9cA151a281bFdf136f66A45b
t0ken token filterTransfer --to 0xf01fF29DCbEE147e9cA151a281bFdf136f66A45b,0xf02f537578d03f6AeCE28F249eaC19542D848F20`,
		PreRun: connectFilterer,
		Run: func(cmd *cobra.Command, args []string) {
			from, err := cli.AddressesFlag(cmd, "from", false)
			cli.CheckErr(cmd, err)
			to, err := cli.AddressesFlag(cmd, "to", false)
			cli.CheckErr(cmd, err)

			opts := cli.FilterOpts(cmd)
			i, err := filterer.FilterTransfer(&opts, from, to)
			cli.CheckErr(cmd, err)

			// Output matching events
			defer i.Close()
			fmt.Println("block,from,to,value")
			for i.Next() {
				cli.CheckErr(cmd, i.Error())
				fmt.Printf("%d,%s,%s,%s\n", i.Event.Raw.BlockNumber, i.Event.From.String(), i.Event.To.String(), i.Event.Value)
			}
		},
	},
}

func init() {
	// Add the Lockable, Ownable contract getter commands
	GetterCommands = append(GetterCommands, lockable.NewGetterCommands(contractKey)...)
	GetterCommands = append(GetterCommands, ownable.NewGetterCommands(contractKey)...)

	for i, cmd := range GetterCommands {
		// Skip ABI/Bin
		if i < 2 {
			continue
		}

		// Allow providing contract 'address' flag
		cmd.Flags().String("address", "", `address of the token contract (default "[`+contractKey+`] value from config")`)
		cli.BlockFlag(cmd)
	}

	for _, cmd := range FilterCommands {
		// Allow providing contract 'address' flag
		cmd.Flags().String("address", "", `address of the Investor registry contract (default "[`+contractKey+`] value from config")`)
		cmd.Flags().Uint64("start", 0, "start block of the filter")
		cmd.Flags().Uint64("end", 0, "end block of the filter")

		cmd.MarkFlagRequired("start")
	}
	FilterCommands[0].Flags().StringSlice("from", nil, "comma separated addresses to filter")
	FilterCommands[0].Flags().StringSlice("to", nil, "comma separated addresses to filter")
}
