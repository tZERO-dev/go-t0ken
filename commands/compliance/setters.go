package compliance

import (
	"errors"
	"math/big"
	"strconv"

	"github.com/ethereum/go-ethereum/common"
	"github.com/spf13/cobra"

	"github.com/tzero-dev/go-t0ken/cli"
	"github.com/tzero-dev/go-t0ken/commands/administrable"
	"github.com/tzero-dev/go-t0ken/commands/destroyable"
	"github.com/tzero-dev/go-t0ken/commands/lockable"
	"github.com/tzero-dev/go-t0ken/commands/ownable"
)

var SetterCommands = []*cobra.Command{
	&cobra.Command{
		Use:    "canIssue <issuer> <from> <to> <quantity>",
		Short:  "Checks if the <issuer> can issue <quantity> of tokens",
		Args:   cli.ChainArgs(cli.AddressArgFunc("issuer", 0), cli.AddressArgFunc("from", 1), cli.AddressArgFunc("to", 2), cli.BigIntArgFunc("quantity", 3)),
		PreRun: connectTransactor,
		Run: func(cmd *cobra.Command, args []string) {
			cli.PrintTransFn(cmd)(transSession.CanIssue(getIssuerFromToQty(args)))
		},
	},
	&cobra.Command{
		Use:    "canOverride <admin> <from> <to> <quantity>",
		Short:  "Checks if the <admin> can override a transfer <from> <to> in the amount of <tokens>",
		Args:   cli.ChainArgs(cli.AddressArgFunc("admin", 0), cli.AddressArgFunc("from", 1), cli.AddressArgFunc("to", 2), cli.BigIntArgFunc("tokens", 3)),
		PreRun: connectTransactor,
		Run: func(cmd *cobra.Command, args []string) {
			cli.CheckGetter(cmd)(transSession.CanOverride(getIssuerFromToQty(args)))
		},
	},
	&cobra.Command{
		Use:    "canTransfer <initiator> <from> <to> <quantity>",
		Short:  "Checks if the <initiator> can transfer <from> <to> in the amount of <tokens>",
		Args:   cli.ChainArgs(cli.AddressArgFunc("initiator", 0), cli.AddressArgFunc("from", 1), cli.AddressArgFunc("to", 2), cli.BigIntArgFunc("quantity", 3)),
		PreRun: connectTransactor,
		Run: func(cmd *cobra.Command, args []string) {
			cli.PrintTransFn(cmd)(transSession.CanTransfer(getIssuerFromToQty(args)))
		},
	},
	&cobra.Command{
		Use:    "setFrozen <address> <frozen>",
		Short:  "Sets the <address> frozen state to <frozen>",
		Args:   cli.ChainArgs(cli.AddressArgFunc("address", 0), cli.BoolArgFunc("frozen", 1)),
		PreRun: connectTransactor,
		Run: func(cmd *cobra.Command, args []string) {
			addr := common.HexToAddress(args[0])
			frozen, _ := strconv.ParseBool(args[1])
			cli.PrintTransFn(cmd)(transSession.SetFrozen(addr, frozen))
		},
	},
	&cobra.Command{
		Use:     "setMaxRules <limit>",
		Short:   "Sets the maximum number of compliance rules to <limit>",
		Example: "t0ken compliance setMaxRules 25 --keystoreAddress admin",
		Args:    cli.UintArgFunc("limit", 0, 8),
		PreRun:  connectTransactor,
		Run: func(cmd *cobra.Command, args []string) {
			limit, _ := strconv.ParseInt(args[0], 10, 8)
			cli.PrintTransFn(cmd)(transSession.SetMaxRules(uint8(limit)))
		},
	},
	&cobra.Command{
		Use:     "setRules <kind> [address args]",
		Short:   "Sets the <kind> rules to [address args]",
		Example: "t0ken compliance setRules 4 0x397e7b9c15ff22ba67ec6e78f46f1e21540bcb36--keystoreAddress admin",
		Args:    cli.UintArgFunc("kind", 0, 8),
		PreRun:  connectTransactor,
		Run: func(cmd *cobra.Command, args []string) {
			// Read in all rule address args
			var rules []common.Address
			var err error
			for i := 1; i < len(args); i++ {
				if !common.IsHexAddress(args[i]) {
					err = errors.New("Address args must be valid addrtesses")
				}
				rules = append(rules, common.HexToAddress(args[i]))
			}

			// Ensure all args are valid addresses
			cli.CheckErr(cmd, err)
			kind, _ := strconv.ParseInt(args[0], 10, 8)
			cli.PrintTransFn(cmd)(transSession.SetRules(uint8(kind), rules))
		},
	},
	&cobra.Command{
		Use:     "setStorage <address>",
		Short:   "Sets the storage contract to <address>",
		Example: "t0ken compliance setStorage 0x397e7b9c15ff22ba67ec6e78f46f1e21540bcb36 --keystoreAddress owner",
		Args:    cli.AddressArgFunc("address", 0),
		PreRun:  connectTransactor,
		Run: func(cmd *cobra.Command, args []string) {
			addr := common.HexToAddress(args[0])
			cli.PrintTransFn(cmd)(transSession.SetStorage(addr))
		},
	},
}

func getIssuerFromToQty(args []string) (common.Address, common.Address, common.Address, *big.Int) {
	issuer := common.HexToAddress(args[0])
	from := common.HexToAddress(args[1])
	to := common.HexToAddress(args[2])
	qty, _ := new(big.Int).SetString(args[3], 10)
	return issuer, from, to, qty
}

func init() {
	// Add the Administrable, Destroyable, Lockable contract getter commands
	SetterCommands = append(SetterCommands, administrable.NewSetterCommands(contractKey)...)
	SetterCommands = append(SetterCommands, destroyable.NewSetterCommands(contractKey)...)
	SetterCommands = append(SetterCommands, lockable.NewSetterCommands(contractKey)...)
	SetterCommands = append(SetterCommands, ownable.NewSetterCommands(contractKey)...)

	for _, cmd := range SetterCommands {
		// Allow providing contract 'address' flag
		cmd.Flags().String("address", "", `address of the BrokerDealer registry contract (default "[`+contractKey+`] value from config")`)
	}
}
