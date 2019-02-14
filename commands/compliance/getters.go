package compliance

import (
	"strconv"

	"github.com/ethereum/go-ethereum/common"
	"github.com/spf13/cobra"

	"github.com/tzero-dev/go-t0ken/cli"
	"github.com/tzero-dev/go-t0ken/commands/administrable"
	"github.com/tzero-dev/go-t0ken/commands/lockable"
	"github.com/tzero-dev/go-t0ken/commands/ownable"
	"github.com/tzero-dev/go-t0ken/contracts/compliance"
)

var GetterCommands = []*cobra.Command{
	&cobra.Command{
		Use:   "abi",
		Short: "Outputs the T0kenCompliance ABI",
		Args:  cobra.NoArgs,
		Run:   func(cmd *cobra.Command, args []string) { cmd.Println(compliance.T0kenComplianceABI) },
	},
	&cobra.Command{
		Use:   "bin",
		Short: "Outputs the T0kenCompliance Binary",
		Args:  cobra.NoArgs,
		Run:   func(cmd *cobra.Command, args []string) { cmd.Println(compliance.T0kenComplianceBin) },
	},
	&cobra.Command{
		Use:    "addressFreezes <address>",
		Short:  "Gets wether the given address is currently set as frozen",
		Args:   cli.AddressArgFunc("address", 0),
		PreRun: connectCaller,
		Run: func(cmd *cobra.Command, args []string) {
			addr := common.HexToAddress(args[0])
			cli.CheckGetter(cmd)(callSession.AddressFreezes(addr))
		},
	},
	&cobra.Command{
		Use:    "getRules <kind>",
		Short:  "Gets the addresses that exist for the given kind",
		Args:   cli.UintArgFunc("kind", 0, 8),
		PreRun: connectCaller,
		Run: func(cmd *cobra.Command, args []string) {
			k, _ := strconv.ParseInt(args[0], 10, 8)
			kind := uint8(k)
			cli.CheckGetter(cmd)(callSession.GetRules(kind))
		},
	},
	&cobra.Command{
		Use:    "maxRules",
		Short:  "Gets wether the given address is currently set as frozen",
		Args:   cobra.NoArgs,
		PreRun: connectCaller,
		Run: func(cmd *cobra.Command, args []string) {
			cli.CheckGetter(cmd)(callSession.MaxRules())
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
	GetterCommands = append(GetterCommands, administrable.NewGetterCommands(contractKey)...)
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
