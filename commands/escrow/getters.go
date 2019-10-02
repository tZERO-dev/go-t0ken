package escrow

import (
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
		Use:     "holders",
		Short:   "Gets the total number of holders in escrow",
		Example: "t0ken escrow holders",
		Args:    cobra.NoArgs,
		PreRun:  connectCaller,
		Run: func(cmd *cobra.Command, args []string) {
			cli.CheckGetter(cmd)(callSession.Holders())
		},
	},
	&cobra.Command{
		Use:     "ledger <symbol> <holder>",
		Short:   "Gets the total number of tokens in escrow for the <holder> of the token <symbol>",
		Example: "t0ken escrow ledger TZROP 0xf01ff29dcbee147e9ca151a281bfdf136f66a45b",
		Args:    cli.ChainArgs(cli.AddressArgFunc("token", 0), cli.BigIntArgFunc("tokens", 1)),
		PreRun:  connectCaller,
		Run: func(cmd *cobra.Command, args []string) {
			holder := common.HexToAddress(args[1])
			cli.CheckGetter(cmd)(callSession.Ledger(args[0], holder))
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
}
