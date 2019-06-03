package externalInvestor

import (
	"math/big"
	"strconv"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/spf13/cobra"

	"github.com/tzero-dev/go-t0ken/cli"
	"github.com/tzero-dev/go-t0ken/commands/destroyable"
	"github.com/tzero-dev/go-t0ken/commands/gas"
	"github.com/tzero-dev/go-t0ken/commands/lockable"
	"github.com/tzero-dev/go-t0ken/commands/nonce"
	"github.com/tzero-dev/go-t0ken/commands/ownable"
)

var SetterCommands = []*cobra.Command{
	&cobra.Command{
		Use:   "add <investor> <hash> <country> <accreditation>",
		Short: "Adds the <investor> address with the given <hash>, 2 character <country> code, and <accreditation> UTC date",
		Example: `t0ken externalInvestor add \
	0xf01ff29dcbee147e9ca151a281bfdf136f66a45b \
	0xa1896382c22b03c562b0241324cfca19505cc5c78eb06751d9cee690e21ed6a1 \
	US \
	2020-07-11 --keystoreAddress broker`,
		Args:   cli.ChainArgs(cli.AddressArgFunc("investor", 0), cli.HexArgFunc("hash", 1, 16), cli.CountryCodeArgFunc("country", 2), cli.DateArgFunc("accreditation", 3)),
		PreRun: connectTransactor,
		Run: func(cmd *cobra.Command, args []string) {
			investor := common.HexToAddress(args[0])
			h, _ := hexutil.Decode(args[1])
			country, _ := cli.CountryFromArg(args[2])
			accreditation, _ := cli.DateFromArg(args[3])

			// Convert has to fixed size byte array
			var hash [32]byte
			copy(hash[:], h)

			cli.PrintTransactionFn(cmd)(transSession.Add(investor, hash, country, big.NewInt(accreditation.Unix())))
		},
	},
	&cobra.Command{
		Use:     "remove <investor>",
		Short:   "Removes the <investor> address --keystoreAddress broker",
		Example: "t0ken externalInvestor remove 0xf01ff29dcbee147e9ca151a281bfdf136f66a45b --keystoreAddress broker",
		Args:    cli.ChainArgs(cli.AddressArgFunc("investor", 0)),
		PreRun:  connectTransactor,
		Run: func(cmd *cobra.Command, args []string) {
			investor := common.HexToAddress(args[0])
			cli.PrintTransactionFn(cmd)(transSession.Remove(investor))
		},
	},
	&cobra.Command{
		Use:     "setAccreditation <investor> <accreditation>",
		Short:   "Sets the <investor> <accreditation> UTC date.",
		Example: "t0ken externalInvestor setAccreditation 0xf01ff29dcbee147e9ca151a281bfdf136f66a45b 2020-07-11 --keystoreAddress broker",
		Args:    cli.ChainArgs(cli.AddressArgFunc("investor", 0), cli.DateArgFunc("accreditation", 1)),
		PreRun:  connectTransactor,
		Run: func(cmd *cobra.Command, args []string) {
			investor := common.HexToAddress(args[0])
			accreditation, _ := cli.DateFromArg(args[2])
			cli.PrintTransactionFn(cmd)(transSession.SetAccreditation(investor, big.NewInt(accreditation.Unix())))
		},
	},
	&cobra.Command{
		Use:     "setCountry <inverstor> <country>",
		Short:   "Sets the <investor> 2 character <country> code",
		Example: "t0ken externalInvestor setCountry 0xf01ff29dcbee147e9ca151a281bfdf136f66a45b US --keystoreAddress broker",
		Args:    cli.ChainArgs(cli.AddressArgFunc("investor", 0), cli.CountryCodeArgFunc("country", 1)),
		PreRun:  connectTransactor,
		Run: func(cmd *cobra.Command, args []string) {
			investor := common.HexToAddress(args[0])
			country, _ := cli.CountryFromArg(args[2])
			cli.PrintTransactionFn(cmd)(transSession.SetCountry(investor, country))
		},
	},
	&cobra.Command{
		Use:     "setFrozen <inverstor> <frozen>",
		Short:   "Sets the <investor> to the <frozen> state",
		Example: "t0ken externalInvestor setFrozen 0xf01ff29dcbee147e9ca151a281bfdf136f66a45b true --keystoreAddress broker",
		Args:    cli.ChainArgs(cli.AddressArgFunc("investor", 0), cli.BoolArgFunc("frozen", 1)),
		PreRun:  connectTransactor,
		Run: func(cmd *cobra.Command, args []string) {
			investor := common.HexToAddress(args[0])
			frozen, _ := strconv.ParseBool(args[1])
			cli.PrintTransactionFn(cmd)(transSession.SetFrozen(investor, frozen))
		},
	},
	&cobra.Command{
		Use:     "setHash <inverstor> <hash>",
		Short:   "Sets the <investor> <hash> value",
		Example: "t0ken externalInvestor setHash 0xf01ff29dcbee147e9ca151a281bfdf136f66a45b 0xa1896382c22b03c562b0241324cfca19505cc5c78eb06751d9cee690e21ed6a1 --keystoreAddress broker",
		Args:    cli.ChainArgs(cli.AddressArgFunc("investor", 0), cli.HexArgFunc("hash", 1, 16)),
		PreRun:  connectTransactor,
		Run: func(cmd *cobra.Command, args []string) {
			investor := common.HexToAddress(args[0])
			h, _ := hexutil.Decode(args[1])

			// Convert data to fixed size byte array
			var hash [32]byte
			copy(hash[:], h)

			cli.PrintTransactionFn(cmd)(transSession.SetHash(investor, hash))
		},
	},
	&cobra.Command{
		Use:     "setStorage <address>",
		Short:   "Sets the storage contract to <address>",
		Example: "t0ken externalInvestor setStorage 0x397e7b9c15ff22ba67ec6e78f46f1e21540bcb36 --keystoreAddress owner",
		Args:    cli.AddressArgFunc("address", 0),
		PreRun:  connectTransactor,
		Run: func(cmd *cobra.Command, args []string) {
			addr := common.HexToAddress(args[0])
			cli.PrintTransactionFn(cmd)(transSession.SetStorage(addr))
		},
	},
}

func init() {
	// Add the Destroyable, Lockable, Ownable contract getter commands
	SetterCommands = append(SetterCommands, destroyable.NewSetterCommands(contractKey)...)
	SetterCommands = append(SetterCommands, lockable.NewSetterCommands(contractKey)...)
	SetterCommands = append(SetterCommands, ownable.NewSetterCommands(contractKey)...)

	for _, cmd := range SetterCommands {
		// Allow 'gasPrice' and 'nonce' flags
		gas.Flag(cmd)
		nonce.Flag(cmd)

		// Allow providing contract 'address' flag
		cmd.Flags().String("address", "", `address of the External-Investor registry contract (default "[`+contractKey+`] value from config")`)
		cmd.Flags().Int("wait", -1, "waits the provided number of seconds for the transaction to be mined ('0' waits indefinitely)")
	}
}
