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
		Use:     "abi",
		Short:   "Outputs the T0kenCompliance ABI",
		Example: "t0ken investor abi",
		Args:    cobra.NoArgs,
		Run:     func(cmd *cobra.Command, args []string) { cmd.Println(compliance.T0kenComplianceABI) },
	},
	&cobra.Command{
		Use:     "bin",
		Short:   "Outputs the T0kenCompliance Binary",
		Example: "t0ken investor bin",
		Args:    cobra.NoArgs,
		Run:     func(cmd *cobra.Command, args []string) { cmd.Println(compliance.T0kenComplianceBin) },
	},
	//&cobra.Command{
	//	Use:     "addressFreezes <address>",
	//	Short:   "Gets wether the given address is currently set as frozen",
	//	Example: "t0ken compliance addressesFreezes 0xf01ff29dcbee147e9ca151a281bfdf136f66a45b",
	//	Args:    cli.AddressArgFunc("address", 0),
	//	PreRun:  connectCaller,
	//	Run: func(cmd *cobra.Command, args []string) {
	//		addr := common.HexToAddress(args[0])
	//		cli.CheckGetter(cmd)(callSession.AddressFreezes(addr))
	//	},
	//},
	&cobra.Command{
		Use:     "getRules <symbol> <kind>",
		Short:   "Gets the addresses that exist for the given kind",
		Example: "t0ken compliance getRules 4",
		Args:    cli.ChainArgs(cli.AddressArgFunc("token", 0), cli.UintArgFunc("kind", 1, 8)),
		PreRun:  connectCaller,
		Run: func(cmd *cobra.Command, args []string) {
			token := common.HexToAddress(args[0])
			k, _ := strconv.ParseInt(args[1], 10, 8)
			kind := uint8(k)
			cli.CheckAddressesGetter(cmd)(callSession.GetRules(token, kind))
		},
	},
	&cobra.Command{
		Use:     "registry",
		Short:   "Gets the Registry contract address",
		Example: "t0ken compliance registry",
		Args:    cobra.NoArgs,
		PreRun:  connectCaller,
		Run: func(cmd *cobra.Command, args []string) {
			cli.CheckAddressGetter(cmd)(callSession.Registry())
		},
	},
	&cobra.Command{
		Use:     "store",
		Short:   "Gets the ComplianceStorage contract address",
		Example: "t0ken compliance store",
		Args:    cobra.NoArgs,
		PreRun:  connectCaller,
		Run: func(cmd *cobra.Command, args []string) {
			cli.CheckAddressGetter(cmd)(callSession.Store())
		},
	},
}

func init() {
	// Add the Administrable, Ownable, Lockable contract getter commands
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
