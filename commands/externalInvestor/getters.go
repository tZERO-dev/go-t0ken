package externalInvestor

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/ethereum/go-ethereum/common"
	"github.com/spf13/cobra"

	"github.com/tzero-dev/go-t0ken/cli"
	"github.com/tzero-dev/go-t0ken/commands/lockable"
	"github.com/tzero-dev/go-t0ken/commands/ownable"
	"github.com/tzero-dev/go-t0ken/contracts/registry"
)

var GetterCommands = []*cobra.Command{
	&cobra.Command{
		Use:     "abi",
		Short:   "Outputs the External Investor Registry ABI",
		Example: "t0ken externalInvestor abi",
		Args:    cobra.NoArgs,
		Run:     func(cmd *cobra.Command, args []string) { cmd.Println(registry.ExternalInvestorABI) },
	},
	&cobra.Command{
		Use:     "bin",
		Short:   "Outputs the External Investor Registry Binary",
		Example: "t0ken externalInvestor bin",
		Args:    cobra.NoArgs,
		Run:     func(cmd *cobra.Command, args []string) { cmd.Println(registry.ExternalInvestorBin) },
	},
	&cobra.Command{
		Use:     "getAccreditation <investor>",
		Short:   "Gets the accreditation date of the <investor>",
		Example: "t0ken externalInvestor getAccreditation 0xf01ff29dcbee147e9ca151a281bfdf136f66a45b",
		Args:    cli.AddressArgFunc("investor", 0),
		PreRun:  connectCaller,
		Run: func(cmd *cobra.Command, args []string) {
			investor := common.HexToAddress(args[0])
			cli.CheckAccreditationGetter(cmd)(callSession.GetAccreditation(investor))
		},
	},
	&cobra.Command{
		Use:     "getCountry <investor>",
		Short:   "Gets the country of the <investor>",
		Example: "t0ken externalInvestor getCountry 0xf01ff29dcbee147e9ca151a281bfdf136f66a45b",
		Args:    cli.AddressArgFunc("investor", 0),
		PreRun:  connectCaller,
		Run: func(cmd *cobra.Command, args []string) {
			investor := common.HexToAddress(args[0])
			cli.CheckCountryGetter(cmd)(callSession.GetCountry(investor))
		},
	},
	&cobra.Command{
		Use:     "getHash <investor>",
		Short:   "Gets the hash of the <investor>",
		Example: "t0ken externalInvestor getHash 0xf01ff29dcbee147e9ca151a281bfdf136f66a45b",
		Args:    cli.AddressArgFunc("investor", 0),
		PreRun:  connectCaller,
		Run: func(cmd *cobra.Command, args []string) {
			investor := common.HexToAddress(args[0])
			cli.CheckBytes32Getter(cmd)(callSession.GetHash(investor))
		},
	},
	&cobra.Command{
		Use:     "registry",
		Short:   "Gets the Registry contract address",
		Example: "t0ken externalInvestor registry",
		Args:    cobra.NoArgs,
		PreRun:  connectCaller,
		Run: func(cmd *cobra.Command, args []string) {
			cli.CheckAddressGetter(cmd)(callSession.Registry())
		},
	},
}

func filterArgs(cmd *cobra.Command) ([]common.Address, []common.Address) {
	investors, err := cli.AddressesFlag(cmd, "investors", false)
	cli.CheckErr(cmd, err)
	owners, err := cli.AddressesFlag(cmd, "owners", false)
	cli.CheckErr(cmd, err)
	return investors, owners
}

var FilterCommands = []*cobra.Command{
	&cobra.Command{
		Use:   "filterAdded",
		Short: "Filters investor added events within the given block range",
		Example: `t0ken externalInvestor filterAdded --start 8658083
t0ken externalInvestor filterAdded --start 8658083 --end 8700000
t0ken externalInvestor filterAdded --owners 0x0000d6AbC75370F48B42024371B07B4506885a55
t0ken externalInvestor filterAdded --investors 0xf01fF29DCbEE147e9cA151a281bFdf136f66A45b,0xf02f537578d03f6AeCE28F249eaC19542D848F20`,
		PreRun: connectFilterer,
		Run: func(cmd *cobra.Command, args []string) {
			investors, owners := filterArgs(cmd)
			opts := cli.FilterOpts(cmd)
			i, err := filterer.FilterInvestorAdded(&opts, investors, owners)
			cli.CheckErr(cmd, err)

			// Output matching events
			defer i.Close()
			fmt.Println("block,owner,investor")
			for i.Next() {
				cli.CheckErr(cmd, i.Error())
				fmt.Printf("%d,%s,%s\n", i.Event.Raw.BlockNumber, i.Event.Owner.String(), i.Event.Investor.String())
			}
		},
	},
	&cobra.Command{
		Use:   "filterRemoved",
		Short: "Filters investor removed events within the given block range",
		Example: `t0ken externalInvestor filterRemoved --start 8658083,
t0ken externalInvestor filterRemoved --start 8658083 --end 8700000
t0ken externalInvestor filterRemoved --owners 0x0000d6AbC75370F48B42024371B07B4506885a55
t0ken externalInvestor filterRemoved --investors 0xf01fF29DCbEE147e9cA151a281bFdf136f66A45b,0xf02f537578d03f6AeCE28F249eaC19542D848F20`,
		PreRun: connectFilterer,
		Run: func(cmd *cobra.Command, args []string) {
			investors, owners := filterArgs(cmd)
			opts := cli.FilterOpts(cmd)
			i, err := filterer.FilterInvestorRemoved(&opts, investors, owners)
			cli.CheckErr(cmd, err)

			// Output matching events
			defer i.Close()
			fmt.Println("block,owner,investor")
			for i.Next() {
				cli.CheckErr(cmd, i.Error())
				fmt.Printf("%d,%s,%s\n", i.Event.Raw.BlockNumber, i.Event.Owner.String(), i.Event.Investor.String())
			}
		},
	},
	&cobra.Command{
		Use:   "filterUpdated",
		Short: "Filters investor updated events within the given block range",
		Example: `t0ken externalInvestor filterUpdated --start 8658083
t0ken externalInvestor filterUpdated --start 8658083 --end 8700000
t0ken externalInvestor filterUpdated --owners 0x0000d6AbC75370F48B42024371B07B4506885a55
t0ken externalInvestor filterUpdated --investors 0xf01fF29DCbEE147e9cA151a281bFdf136f66A45b,0xf02f537578d03f6AeCE28F249eaC19542D848F20`,
		PreRun: connectFilterer,
		Run: func(cmd *cobra.Command, args []string) {
			investors, owners := filterArgs(cmd)
			opts := cli.FilterOpts(cmd)
			i, err := filterer.FilterInvestorUpdated(&opts, investors, owners)
			cli.CheckErr(cmd, err)

			// Output matching events
			defer i.Close()
			fmt.Println("block,owner,investor")
			for i.Next() {
				cli.CheckErr(cmd, i.Error())
				fmt.Printf("%d,%s,%s\n", i.Event.Raw.BlockNumber, i.Event.Owner.String(), i.Event.Investor.String())
			}
		},
	},
	&cobra.Command{
		Use:   "filterFrozen [frozen]",
		Short: "Filters investor frozen events within the given block range",
		Example: `t0ken externalInvestor filterFrozen true --start 8658083
t0ken externalInvestor filterFrozen --start 8658083 --end 8700000
t0ken externalInvestor filterFrozen --owners 0x0000d6AbC75370F48B42024371B07B4506885a55
t0ken externalInvestor filterFrozen false --investors 0xf01fF29DCbEE147e9cA151a281bFdf136f66A45b,0xf02f537578d03f6AeCE28F249eaC19542D848F20`,
		Args:   cobra.MaximumNArgs(1),
		PreRun: connectFilterer,
		Run: func(cmd *cobra.Command, args []string) {
			var frozen []bool
			if len(args) > 0 {
				b, err := strconv.ParseBool(args[0])
				if err != nil {
					cli.CheckErr(cmd, errors.New("invalid value for [frozen] arg"))
				}
				frozen = append(frozen, b)
			}

			investors, owners := filterArgs(cmd)
			opts := cli.FilterOpts(cmd)
			i, err := filterer.FilterInvestorFrozen(&opts, investors, frozen, owners)
			cli.CheckErr(cmd, err)

			// Output matching events
			defer i.Close()
			fmt.Println("block,owner,investor,frozen")
			for i.Next() {
				cli.CheckErr(cmd, i.Error())
				fmt.Printf("%d,%s,%s\n", i.Event.Raw.BlockNumber, i.Event.Owner.String(), i.Event.Investor.String(), i.Event.Frozen)
			}
		},
	},
}

func init() {
	// Add the Ownable, Lockable contract getter commands
	GetterCommands = append(GetterCommands, ownable.NewGetterCommands(contractKey)...)
	GetterCommands = append(GetterCommands, lockable.NewGetterCommands(contractKey)...)

	for i, cmd := range GetterCommands {
		// Skip ABI/Bin
		if i < 2 {
			continue
		}

		// Allow providing contract 'address' flag
		cmd.Flags().String("address", "", `address of the External-Investor registry contract (default "[`+contractKey+`] value from config")`)
		cli.BlockFlag(cmd)
	}

	for _, cmd := range FilterCommands {
		// Allow providing contract 'address' flag
		cmd.Flags().String("address", "", `address of the Investor registry contract (default "[`+contractKey+`] value from config")`)
		cmd.Flags().Uint64("start", 0, "start block of the filter")
		cmd.Flags().Uint64("end", 0, "end block of the filter")

		cmd.MarkFlagRequired("start")
		cmd.Flags().StringSlice("investors", nil, "comma separated addresses to filter")
		cmd.Flags().StringSlice("owners", nil, "comma separated addresses to filter")
	}
}
