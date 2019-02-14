package investor

import (
	"math/big"
	"strconv"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/spf13/cobra"

	"github.com/tzero-dev/go-t0ken/cli"
	"github.com/tzero-dev/go-t0ken/commands/destroyable"
	"github.com/tzero-dev/go-t0ken/commands/lockable"
	"github.com/tzero-dev/go-t0ken/commands/ownable"
)

var SetterCommands = []*cobra.Command{
	&cobra.Command{
		Use:    "add <investor> <hash> <country> <accreditation>",
		Short:  "Adds the <investor> address with the given <hash>, 2 character <country> code, and <accreditation> UTC date",
		Args:   cli.ChainArgs(cli.AddressArgFunc("investor", 0), cli.HexArgFunc("hash", 1, 32), cli.CountryCodeArgFunc("country", 2), cli.DateArgFunc("accreditation", 3)),
		PreRun: connectTransactor,
		Run: func(cmd *cobra.Command, args []string) {
			investor := common.HexToAddress(args[0])
			h, _ := hexutil.Decode(args[1])
			country, _ := cli.CountryFromArg(args[2])
			accreditation, _ := cli.DateFromArg(args[3])

			// Convert has to fixed size byte array
			var hash [32]byte
			copy(hash[:], h)

			cli.PrintTransFn(cmd)(transSession.Add(investor, hash, country, big.NewInt(accreditation.Unix())))
		},
	},
	&cobra.Command{
		Use:    "remove <investor>",
		Short:  "Removes the <investor> address",
		Args:   cli.ChainArgs(cli.AddressArgFunc("investor", 0)),
		PreRun: connectTransactor,
		Run: func(cmd *cobra.Command, args []string) {
			investor := common.HexToAddress(args[0])
			cli.PrintTransFn(cmd)(transSession.Remove(investor))
		},
	},
	&cobra.Command{
		Use:    "setAccreditation <investor> <accreditation>",
		Short:  "Sets the <investor> <accreditation> UTC date.",
		Args:   cli.ChainArgs(cli.AddressArgFunc("investor", 0), cli.DateArgFunc("accreditation", 1)),
		PreRun: connectTransactor,
		Run: func(cmd *cobra.Command, args []string) {
			investor := common.HexToAddress(args[0])
			accreditation, _ := cli.DateFromArg(args[2])
			cli.PrintTransFn(cmd)(transSession.SetAccreditation(investor, big.NewInt(accreditation.Unix())))
		},
	},
	&cobra.Command{
		Use:    "setCountry <inverstor> <country>",
		Short:  "Sets the <investor> 2 character <country> code",
		Args:   cli.ChainArgs(cli.AddressArgFunc("investor", 0), cli.CountryCodeArgFunc("country", 1)),
		PreRun: connectTransactor,
		Run: func(cmd *cobra.Command, args []string) {
			investor := common.HexToAddress(args[0])
			country, _ := cli.CountryFromArg(args[2])
			cli.PrintTransFn(cmd)(transSession.SetCountry(investor, country))
		},
	},
	&cobra.Command{
		Use:    "setFrozen <inverstor> <frozen>",
		Short:  "Sets the <investor> to the <frozen> state",
		Args:   cli.ChainArgs(cli.AddressArgFunc("investor", 0), cli.BoolArgFunc("frozen", 1)),
		PreRun: connectTransactor,
		Run: func(cmd *cobra.Command, args []string) {
			investor := common.HexToAddress(args[0])
			frozen, _ := strconv.ParseBool(args[1])
			cli.PrintTransFn(cmd)(transSession.SetFrozen(investor, frozen))
		},
	},
	&cobra.Command{
		Use:    "setHash <inverstor> <hash>",
		Short:  "Sets the <investor> <hash> value",
		Args:   cli.ChainArgs(cli.AddressArgFunc("investor", 0), cli.HexArgFunc("hash", 1, 16)),
		PreRun: connectTransactor,
		Run: func(cmd *cobra.Command, args []string) {
			investor := common.HexToAddress(args[0])
			h, _ := hexutil.Decode(args[1])

			// Convert data to fixed size byte array
			var hash [32]byte
			copy(hash[:], h)

			cli.PrintTransFn(cmd)(transSession.SetHash(investor, hash))
		},
	},
}

func init() {
	// Add the Administrable, Destroyable, Lockable contract getter commands
	SetterCommands = append(SetterCommands, destroyable.NewSetterCommands(contractKey)...)
	SetterCommands = append(SetterCommands, lockable.NewSetterCommands(contractKey)...)
	SetterCommands = append(SetterCommands, ownable.NewSetterCommands(contractKey)...)

	for _, cmd := range SetterCommands {
		// Allow providing contract 'address' flag
		cmd.Flags().String("address", "", `address of the BrokerDealer registry contract (default "[`+contractKey+`] value from config")`)
	}
}
