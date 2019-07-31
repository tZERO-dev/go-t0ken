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
	"github.com/tzero-dev/go-t0ken/commands/gas"
	"github.com/tzero-dev/go-t0ken/commands/lockable"
	"github.com/tzero-dev/go-t0ken/commands/nonce"
	"github.com/tzero-dev/go-t0ken/commands/ownable"
)

var SetterCommands = []*cobra.Command{
	&cobra.Command{
		Use:     "setRegistryAndStore <registry> <storage>",
		Short:   "Sets the registry and compaliance-storage addresses",
		Example: "t0ken compliance setRegistryAndStore 0x397e7b9c15ff22ba67ec6e78f46f1e21540bcb36 0x0f18bbc1ae1c5ab891a7feb06d65388ba6b4cd07 --keystoreAddress owner",
		Args:    cli.ChainArgs(cli.AddressArgFunc("registry", 0), cli.AddressArgFunc("storage", 0)),
		PreRun:  connectTransactor,
		Run: func(cmd *cobra.Command, args []string) {
			r := common.HexToAddress(args[0])
			s := common.HexToAddress(args[1])
			cli.PrintTransactionFn(cmd)(transSession.SetRegistryAndComplianceStore(r, s))
		},
	},
	&cobra.Command{
		Use:    "setFrozen <token> <address> <frozen>",
		Short:  "Sets the <token> <address> frozen state to <frozen>",
		Args:   cli.ChainArgs(cli.AddressArgFunc("token", 0), cli.AddressArgFunc("address", 1), cli.BoolArgFunc("frozen", 2)),
		PreRun: connectTransactor,
		Run: func(cmd *cobra.Command, args []string) {
			token := common.HexToAddress(args[0])
			addr := common.HexToAddress(args[1])
			frozen, _ := strconv.ParseBool(args[1])
			cli.PrintTransactionFn(cmd)(transSession.SetFrozen(token, addr, frozen))
		},
	},
	&cobra.Command{
		Use:     "setRules <token> <kind> [address args]",
		Short:   "Sets the <token> <kind> rules to [address args]",
		Example: "t0ken compliance setRules 0x5bd5b4e1a2c9b12812795e7217201b78c8c10b78 4 0x397e7b9c15ff22ba67ec6e78f46f1e21540bcb36 --keystoreAddress admin",
		Args:    cli.ChainArgs(cli.AddressArgFunc("token", 0), cli.UintArgFunc("kind", 1, 8)),
		PreRun:  connectTransactor,
		Run: func(cmd *cobra.Command, args []string) {
			// Read in all rule address args
			var rules []common.Address
			var err error
			for i := 2; i < len(args); i++ {
				if !common.IsHexAddress(args[i]) {
					err = errors.New("Address args must be valid addrtesses")
				}
				rules = append(rules, common.HexToAddress(args[i]))
			}
			// Ensure all args are valid addresses
			cli.CheckErr(cmd, err)

			token := common.HexToAddress(args[0])
			kind, _ := strconv.ParseInt(args[1], 10, 8)
			cli.PrintTransactionFn(cmd)(transSession.SetRules(token, uint8(kind), rules))
		},
	},
}

func getInitiatorFromToQty(args []string) (common.Address, common.Address, common.Address, *big.Int) {
	initiator := common.HexToAddress(args[0])
	from := common.HexToAddress(args[1])
	to := common.HexToAddress(args[2])
	qty, _ := new(big.Int).SetString(args[3], 10)
	return initiator, from, to, qty
}

func init() {
	// Add the Administrable, Destroyable, Lockable, Ownable contract getter commands
	SetterCommands = append(SetterCommands, administrable.NewSetterCommands(contractKey)...)
	SetterCommands = append(SetterCommands, destroyable.NewSetterCommands(contractKey)...)
	SetterCommands = append(SetterCommands, lockable.NewSetterCommands(contractKey)...)
	SetterCommands = append(SetterCommands, ownable.NewSetterCommands(contractKey)...)

	for _, cmd := range SetterCommands {
		// Add the 'gasPrice', 'nonce' and 'wait' flags to the deploy function
		gas.Flag(cmd)
		nonce.Flag(cmd)
		cli.WaitFlag(cmd)

		// Allow providing contract 'address' flag
		cmd.Flags().String("address", "", `address of the BrokerDealer registry contract (default "[`+contractKey+`] value from config")`)
	}
}
