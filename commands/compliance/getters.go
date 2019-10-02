package compliance

import (
	"bytes"
	"math/big"
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
		Short:   "Outputs the Compliance ABI",
		Example: "t0ken investor abi",
		Args:    cobra.NoArgs,
		Run:     func(cmd *cobra.Command, args []string) { cmd.Println(compliance.ComplianceABI) },
	},
	&cobra.Command{
		Use:     "bin",
		Short:   "Outputs the Compliance Binary",
		Example: "t0ken investor bin",
		Args:    cobra.NoArgs,
		Run:     func(cmd *cobra.Command, args []string) { cmd.Println(compliance.ComplianceBin) },
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
	&cobra.Command{
		Use:     "getRules <token> <kind>",
		Short:   "Gets the addresses that exist for the given kind",
		Example: "t0ken compliance getRules 0x47acdf4eac45ba7d819be0c6c96c3ebda5283405 4",
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
		Use:   "canTransfer <token> <initiator> <from> <to> <quantity>",
		Short: "Checks if the <initiator> can transfer <from> <to> in the amount of <tokens>",
		Example: `t0ken compliance canTransferTest \
	0x0000d6abc75370f48b42024371b07b4506885a55 \
	0xf01ff29dcbee147e9ca151a281bfdf136f66a45b \
	0xf02f537578d03f6aece28f249eac19542d848f20 \
	5`,
		Args:   cli.ChainArgs(cli.AddressArgFunc("token", 0), cli.AddressArgFunc("from", 1), cli.AddressArgFunc("to", 2), cli.BigIntArgFunc("quantity", 3)),
		PreRun: connectCaller,
		Run: func(cmd *cobra.Command, args []string) {
			token := common.HexToAddress(args[0])
			from := common.HexToAddress(args[1])
			to := common.HexToAddress(args[2])
			qty, _ := new(big.Int).SetString(args[3], 10)

			// Set the initiator when the flag is present, otherwise set it to the from address
			initiator, err := cli.OptionalFlagAddress(cmd, "initiator")
			cli.CheckErr(cmd, err)
			if bytes.Equal(initiator.Bytes(), make([]byte, 20)) {
				initiator = from
			}

			s, err := callSession.CanTransferTest(token, initiator, from, to, qty)
			cli.CheckErr(cmd, err)
			cmd.Println(s)
		},
	},
}

func init() {
	GetterCommands[len(GetterCommands)-1].Flags().String("initiator", "", "initiating address for the request")

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
		cmd.Flags().String("address", "", `address of the Compliance contract (default "[`+contractKey+`] value from config")`)
		cli.BlockFlag(cmd)
	}
}
