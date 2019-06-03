package externalInvestor

import (
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
			cli.CheckHashGetter(cmd)(callSession.GetHash(investor))
		},
	},
	&cobra.Command{
		Use:     "storage",
		Short:   "Gets the Storage contract address",
		Example: "t0ken externalInvestor storage",
		Args:    cobra.NoArgs,
		PreRun:  connectCaller,
		Run: func(cmd *cobra.Command, args []string) {
			cli.CheckAddressGetter(cmd)(callSession.Store())
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
	}
}
