package escrow

import (
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/spf13/cobra"

	"github.com/tzero-dev/go-t0ken/cli"
	"github.com/tzero-dev/go-t0ken/commands/administrable"
	e "github.com/tzero-dev/go-t0ken/contracts/escrow"
)

var GetterCommands = []*cobra.Command{
	&cobra.Command{
		Use:     "abi",
		Short:   "Outputs the Escrow ABI",
		Example: "t0ken escrow abi",
		Args:    cobra.NoArgs,
		Run:     func(cmd *cobra.Command, args []string) { cmd.Println(e.EscrowABI) },
	},
	&cobra.Command{
		Use:     "bin",
		Short:   "Outputs the Escrow Binary",
		Example: "t0ken escrow bin",
		Args:    cobra.NoArgs,
		Run:     func(cmd *cobra.Command, args []string) { cmd.Println(e.EscrowBin) },
	},
	&cobra.Command{
		Use:     "registry",
		Short:   "Gets the Registry contract address",
		Example: "t0ken escrow registry",
		Args:    cobra.NoArgs,
		PreRun:  connectCaller,
		Run: func(cmd *cobra.Command, args []string) {
			cli.CheckAddressGetter(cmd)(callSession.Registry())
		},
	},
	&cobra.Command{
		Use:     "ledger <token> <grantor>",
		Short:   "Gets the total number of tokens in escrow for the <grantor> of the <token>",
		Example: "t0ken escrow ledger 0x2c8091e6f63d9efde70cb278e3c91934d85462d2 0xf01ff29dcbee147e9ca151a281bfdf136f66a45b",
		Args:    cli.ChainArgs(cli.AddressArgFunc("token", 0), cli.AddressArgFunc("grantor", 1)),
		PreRun:  connectCaller,
		Run: func(cmd *cobra.Command, args []string) {
			token := common.HexToAddress(args[0])
			grantor := common.HexToAddress(args[1])
			cli.CheckGetter(cmd)(callSession.Ledger(token, grantor))
		},
	},
}

var FilterCommands = []*cobra.Command{
	&cobra.Command{
		Use:   "filterAdded",
		Short: "Filters escrow removed events within the given block range",
		Example: `t0ken token filterAdded --start 8658083
t0ken token filterAdded --start 8658083 --end 8700000
t0ken token filterAdded --symbols TZROP,OSTKO
t0ken token filterAdded --tokens 0xf01fF29DCbEE147e9cA151a281bFdf136f66A45b
t0ken token filterAdded --grantors 0xf01fF29DCbEE147e9cA151a281bFdf136f66A45b,0xf02f537578d03f6AeCE28F249eaC19542D848F20`,
		PreRun: connectFilterer,
		Run: func(cmd *cobra.Command, args []string) {
			symbols, err := cli.StringsFlag(cmd, "symbols", false)
			cli.CheckErr(cmd, err)
			tokens, err := cli.AddressesFlag(cmd, "tokens", false)
			cli.CheckErr(cmd, err)
			grantors, err := cli.AddressesFlag(cmd, "grantors", false)
			cli.CheckErr(cmd, err)

			opts := cli.FilterOpts(cmd)
			i, err := filterer.FilterEscrowAdded(&opts, symbols, tokens, grantors)
			cli.CheckErr(cmd, err)

			// Output matching events
			defer i.Close()
			fmt.Println("block,symbol-hash,token,grantor")
			for i.Next() {
				cli.CheckErr(cmd, i.Error())
				e := i.Event
				fmt.Printf("%d,%s,%s,%s\n", e.Raw.BlockNumber, e.Symbol.String(), e.Token.String(), e.Grantor.String())
			}
		},
	},
	&cobra.Command{
		Use:   "filterRemoved",
		Short: "Filters escrow added events within the given block range",
		Example: `t0ken token filterRemoved --start 8658083
t0ken token filterRemoved --start 8658083 --end 8700000
t0ken token filterRemoved --symbols TZROP,OSTKO
t0ken token filterRemoved --tokens 0xf01fF29DCbEE147e9cA151a281bFdf136f66A45b
t0ken token filterRemoved --grantors 0xf01fF29DCbEE147e9cA151a281bFdf136f66A45b,0xf02f537578d03f6AeCE28F249eaC19542D848F20`,
		PreRun: connectFilterer,
		Run: func(cmd *cobra.Command, args []string) {
			symbols, err := cli.StringsFlag(cmd, "symbols", false)
			cli.CheckErr(cmd, err)
			tokens, err := cli.AddressesFlag(cmd, "tokens", false)
			cli.CheckErr(cmd, err)
			grantors, err := cli.AddressesFlag(cmd, "grantors", false)
			cli.CheckErr(cmd, err)

			opts := cli.FilterOpts(cmd)
			i, err := filterer.FilterEscrowRemoved(&opts, symbols, tokens, grantors)
			cli.CheckErr(cmd, err)

			// Output matching events
			defer i.Close()
			fmt.Println("block,symbol-hash,token,grantor")
			for i.Next() {
				cli.CheckErr(cmd, i.Error())
				e := i.Event
				fmt.Printf("%d,%s,%s,%s\n", e.Raw.BlockNumber, e.Symbol.String(), e.Token.String(), e.Grantor.String())
			}
		},
	},
}

func init() {
	// Add Administrable contract getter commands
	GetterCommands = append(GetterCommands, administrable.NewGetterCommands(contractKey)...)

	for i, cmd := range GetterCommands {
		// Skip ABI/Bin
		if i < 2 {
			continue
		}

		// Allow providing contract 'address' flag
		cmd.Flags().String("address", "", `address of the Escrow contract (default "[`+contractKey+`] value from config")`)
		cli.BlockFlag(cmd)
	}

	for _, cmd := range FilterCommands {
		// Allow providing contract 'address' flag
		cmd.Flags().String("address", "", `address of the Escrow contract (default "[`+contractKey+`] value from config")`)
		cmd.Flags().Uint64("start", 0, "start block of the filter")
		cmd.Flags().Uint64("end", 0, "end block of the filter")

		cmd.MarkFlagRequired("start")
		cmd.Flags().StringSlice("symbols", nil, "comma separated addresses to filter")
		cmd.Flags().StringSlice("tokens", nil, "comma separated addresses to filter")
		cmd.Flags().StringSlice("grantors", nil, "comma separated addresses to filter")
	}
}
