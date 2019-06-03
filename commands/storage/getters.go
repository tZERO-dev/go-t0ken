package storage

import (
	"fmt"
	"math/big"
	"strconv"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/spf13/cobra"

	"github.com/tzero-dev/go-t0ken/cli"
	"github.com/tzero-dev/go-t0ken/commands/lockable"
	"github.com/tzero-dev/go-t0ken/commands/ownable"
	"github.com/tzero-dev/go-t0ken/contracts/registry"
)

var GetterCommands = []*cobra.Command{
	&cobra.Command{
		Use:     "abi",
		Short:   "Outputs the Storage ABI",
		Example: "t0ken investor abi",
		Args:    cobra.NoArgs,
		Run:     func(cmd *cobra.Command, args []string) { cmd.Println(registry.StorageABI) },
	},
	&cobra.Command{
		Use:     "bin",
		Short:   "Outputs the Storage Binary",
		Example: "t0ken investor bin",
		Args:    cobra.NoArgs,
		Run:     func(cmd *cobra.Command, args []string) { cmd.Println(registry.StorageBin) },
	},
	&cobra.Command{
		Use:     "accountAt <index>",
		Short:   "Gets account at the given <index>",
		Example: "t0ken storage accountAt 5",
		Args:    cli.BigIntArgFunc("index", 0),
		PreRun:  connectCaller,
		Run: func(cmd *cobra.Command, args []string) {
			index, _ := new(big.Int).SetString(args[0], 10)
			cli.CheckGetter(cmd)(accountStr(callSession.AccountAt(index)))
		},
	},
	&cobra.Command{
		Use:     "accountExists <address>",
		Short:   "Checks if the <address> exists",
		Example: "t0ken storage accountExists 0xf01ff29dcbee147e9ca151a281bfdf136f66a45b",
		Args:    cli.AddressArgFunc("address", 0),
		PreRun:  connectCaller,
		Run: func(cmd *cobra.Command, args []string) {
			addr := common.HexToAddress(args[0])
			cli.CheckGetter(cmd)(callSession.AccountExists(addr))
		},
	},
	&cobra.Command{
		Use:     "accountFrozen <address>",
		Short:   "Checks if the <address> is frozen",
		Example: "t0ken storage accountFrozen 0xf01ff29dcbee147e9ca151a281bfdf136f66a45b",
		Args:    cli.AddressArgFunc("address", 0),
		PreRun:  connectCaller,
		Run: func(cmd *cobra.Command, args []string) {
			addr := common.HexToAddress(args[0])
			cli.CheckGetter(cmd)(callSession.AccountFrozen(addr))
		},
	},
	&cobra.Command{
		Use:     "accountGet <address>",
		Short:   "Gets the account for the given <address>",
		Example: "t0ken storage accountGet 0xf01ff29dcbee147e9ca151a281bfdf136f66a45b",
		Args:    cli.AddressArgFunc("address", 0),
		PreRun:  connectCaller,
		Run: func(cmd *cobra.Command, args []string) {
			addr := common.HexToAddress(args[0])
			cli.CheckGetter(cmd)(accountGetStr(callSession.AccountGet(addr)))
		},
	},
	&cobra.Command{
		Use:     "accountIndexOf <address>",
		Short:   "Gets the index of the <address>",
		Example: "t0ken storage indexOf 0xf01ff29dcbee147e9ca151a281bfdf136f66a45b",
		Args:    cli.AddressArgFunc("address", 0),
		PreRun:  connectCaller,
		Run: func(cmd *cobra.Command, args []string) {
			addr := common.HexToAddress(args[0])
			cli.CheckGetter(cmd)(callSession.AccountIndexOf(addr))
		},
	},
	&cobra.Command{
		Use:     "accountKind <address>",
		Short:   "Gets the kind for the <address>",
		Example: "t0ken storage accountKind 0xf01ff29dcbee147e9ca151a281bfdf136f66a45b",
		Args:    cli.AddressArgFunc("address", 0),
		PreRun:  connectCaller,
		Run: func(cmd *cobra.Command, args []string) {
			addr := common.HexToAddress(args[0])
			cli.CheckGetter(cmd)(callSession.AccountKind(addr))
		},
	},
	&cobra.Command{
		Use:     "accountParent <address>",
		Short:   "Gets the parent of the <address>",
		Example: "t0ken storage accountParent 0xf01ff29dcbee147e9ca151a281bfdf136f66a45b",
		Args:    cli.AddressArgFunc("address", 0),
		PreRun:  connectCaller,
		Run: func(cmd *cobra.Command, args []string) {
			addr := common.HexToAddress(args[0])
			cli.CheckGetter(cmd)(callSession.AccountParent(addr))
		},
	},
	&cobra.Command{
		Use:     "accounts",
		Short:   "Gets the total number of accounts",
		Example: "t0ken storage accounts",
		Args:    cobra.NoArgs,
		PreRun:  connectCaller,
		Run: func(cmd *cobra.Command, args []string) {
			cli.CheckGetter(cmd)(callSession.Accounts())
		},
	},
	&cobra.Command{
		Use:     "data <address> <index>",
		Short:   "Gets the data for the <address> at the given <index>",
		Example: "t0ken storage data 0xf01ff29dcbee147e9ca151a281bfdf136f66a45b 1",
		Args:    cli.ChainArgs(cli.AddressArgFunc("address", 0), cli.UintArgFunc("index", 1, 8)),
		PreRun:  connectCaller,
		Run: func(cmd *cobra.Command, args []string) {
			addr := common.HexToAddress(args[0])
			index, _ := strconv.ParseInt(args[1], 10, 8)
			cli.CheckGetter(cmd)(dataStr(callSession.Data(addr, uint8(index))))
		},
	},
	&cobra.Command{
		Use:     "permissionAt <kind> <index>",
		Short:   "Gets the address for the <kind> permission at <index>",
		Example: "t0ken permissionAt 4, 1",
		Args:    cli.ChainArgs(cli.UintArgFunc("kind", 0, 8), cli.BigIntArgFunc("index", 1)),
		PreRun:  connectCaller,
		Run: func(cmd *cobra.Command, args []string) {
			kind, _ := strconv.ParseInt(args[0], 10, 8)
			index, _ := new(big.Int).SetString(args[1], 10)
			cli.CheckAddressGetter(cmd)(callSession.PermissionAt(uint8(kind), index))
		},
	},
	&cobra.Command{
		Use:     "permissionExists <kind> <address>",
		Short:   "Checks if <kind> permission exists for the <address>",
		Example: "t0ken storage permissionExists 4 0xf01ff29dcbee147e9ca151a281bfdf136f66a45b",
		Args:    cli.ChainArgs(cli.UintArgFunc("kind", 0, 8), cli.AddressArgFunc("address", 1)),
		PreRun:  connectCaller,
		Run: func(cmd *cobra.Command, args []string) {
			kind, _ := strconv.ParseInt(args[0], 10, 8)
			addr := common.HexToAddress(args[1])
			cli.CheckGetter(cmd)(callSession.PermissionExists(uint8(kind), addr))
		},
	},
	&cobra.Command{
		Use:     "permissionIndexOf <kind> <address>",
		Short:   "Checks if <kind> permission exists for the <address>",
		Example: "t0ken storage permissionIndexOf 4 0xf01ff29dcbee147e9ca151a281bfdf136f66a45b",
		Args:    cli.ChainArgs(cli.UintArgFunc("kind", 0, 8), cli.AddressArgFunc("address", 1)),
		PreRun:  connectCaller,
		Run: func(cmd *cobra.Command, args []string) {
			kind, _ := strconv.ParseInt(args[0], 10, 8)
			addr := common.HexToAddress(args[1])
			cli.CheckGetter(cmd)(callSession.PermissionIndexOf(uint8(kind), addr))
		},
	},
	&cobra.Command{
		Use:     "permissions <kind>",
		Short:   "Gets the total number of permissions for the <kind>",
		Example: "t0ken storage permissions 4",
		Args:    cli.UintArgFunc("kind", 0, 8),
		PreRun:  connectCaller,
		Run: func(cmd *cobra.Command, args []string) {
			kind, _ := strconv.ParseInt(args[0], 10, 8)
			cli.CheckGetter(cmd)(callSession.Permissions(uint8(kind)))
		},
	},
}

func accountGetStr(kind uint8, frozen bool, parent common.Address, err error) (interface{}, error) {
	if err != nil {
		return nil, err
	}

	return fmt.Sprintf(`   kind: %d
 frozen: %t
 parent: %s`, kind, frozen, parent.String()), err
}

func accountStr(addr common.Address, kind uint8, frozen bool, parent common.Address, err error) (interface{}, error) {
	if err != nil {
		return nil, err
	}

	return fmt.Sprintf(` address: %s
    kind: %d
  frozen: %t
  parent: %s`, addr.String(), kind, frozen, parent.String()), err
}

func dataStr(data [32]byte, err error) (interface{}, error) {
	if err != nil {
		return nil, err
	}

	// Add the hex value first
	h := hexutil.Encode(data[:])
	s := []string{"Hex:\n  ", h, "\nBinary:\n"}

	// Add the binary value second
	for i, b := range data {
		if i > 0 && i%4 == 0 {
			s = append(s, "\n")
		}
		s = append(s, fmt.Sprintf("  %08b", b))
	}
	return strings.Join(s[:], ""), err
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
		cmd.Flags().String("address", "", `address of the Storage contract (default "[`+contractKey+`] value from config")`)
	}
}
