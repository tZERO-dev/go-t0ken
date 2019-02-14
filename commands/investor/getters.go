package investor

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
		Use:   "abi",
		Short: "Outputs the Storage ABI",
		Args:  cobra.NoArgs,
		Run:   func(cmd *cobra.Command, args []string) { cmd.Println(registry.StorageABI) },
	},
	&cobra.Command{
		Use:   "bin",
		Short: "Outputs the Storage Binary",
		Args:  cobra.NoArgs,
		Run:   func(cmd *cobra.Command, args []string) { cmd.Println(registry.StorageBin) },
	},
	&cobra.Command{
		Use:    "getAccreditation <investor>",
		Short:  "Gets the accreditation date of the <investor>",
		Args:   cli.AddressArgFunc("investor", 0),
		PreRun: connectCaller,
		Run: func(cmd *cobra.Command, args []string) {
			investor := common.HexToAddress(args[0])
			cli.CheckAccreditationGetter(cmd)(callSession.GetAccreditation(investor))
		},
	},
	&cobra.Command{
		Use:    "getCountry <investor>",
		Short:  "Gets the country of the <investor>",
		Args:   cli.AddressArgFunc("investor", 0),
		PreRun: connectCaller,
		Run: func(cmd *cobra.Command, args []string) {
			investor := common.HexToAddress(args[0])
			cli.CheckCountryGetter(cmd)(callSession.GetCountry(investor))
		},
	},
	&cobra.Command{
		Use:    "getHash <investor>",
		Short:  "Gets the hash of the <investor>",
		Args:   cli.AddressArgFunc("investor", 0),
		PreRun: connectCaller,
		Run: func(cmd *cobra.Command, args []string) {
			investor := common.HexToAddress(args[0])
			cli.CheckHashGetter(cmd)(callSession.GetHash(investor))
		},
	},
	&cobra.Command{
		Use:    "store",
		Short:  "Gets the Storage contract address",
		Args:   cobra.NoArgs,
		PreRun: connectCaller,
		Run: func(cmd *cobra.Command, args []string) {
			cli.CheckAddressGetter(cmd)(callSession.Store())
		},
	},
}

func init() {
	// Add the Administrable, Lockable contract getter commands
	GetterCommands = append(GetterCommands, ownable.NewGetterCommands(contractKey)...)
	GetterCommands = append(GetterCommands, lockable.NewGetterCommands(contractKey)...)

	for i, cmd := range GetterCommands {
		// Skip ABI/Bin
		if i < 2 {
			continue
		}

		// Allow providing contract 'address' flag
		cmd.Flags().String("address", "", `address of the BrokerDealer registry contract (default "[`+contractKey+`] value from config")`)
	}
}
