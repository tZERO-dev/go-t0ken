package complianceStorage

import (
	"encoding/hex"
	"math/big"
	"strconv"

	"github.com/ethereum/go-ethereum/common"
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
		Use:    "setPermission <address> <grant>",
		Short:  "Grants/Revokes permission for the <addr>",
		Args:   cli.ChainArgs(cli.AddressArgFunc("address", 0), cli.BoolArgFunc("grant", 1)),
		PreRun: connectTransactor,
		Run: func(cmd *cobra.Command, args []string) {
			addr := common.HexToAddress(args[0])
			grant, _ := strconv.ParseBool(args[1])
			cli.PrintTransactionFn(cmd)(transSession.SetPermission(addr, grant))
		},
	},

	&cobra.Command{
		Use:    "setAddress <key> <value>",
		Short:  "Sets the address at <key> to the <value>",
		Args:   cli.ChainArgs(cli.HexArgLenFunc("key", 0, 16), cli.AddressArgFunc("value", 1)),
		PreRun: connectTransactor,
		Run: func(cmd *cobra.Command, args []string) {
			key, _ := cli.Bytes32FromArg(args[0])
			value := common.HexToAddress(args[1])
			cli.PrintTransactionFn(cmd)(transSession.SetAddress(key, value))
		},
	},
	&cobra.Command{
		Use:    "deleteAddress <key>",
		Short:  "Deletes the address for the <key>",
		Args:   cli.HexArgLenFunc("key", 0, 16),
		PreRun: connectTransactor,
		Run: func(cmd *cobra.Command, args []string) {
			key, _ := cli.Bytes32FromArg(args[0])
			cli.PrintTransactionFn(cmd)(transSession.DeleteAddress(key))
		},
	},

	&cobra.Command{
		Use:    "setBool <key> <value>",
		Short:  "Sets the bool at <key> to the <value>",
		Args:   cli.ChainArgs(cli.HexArgLenFunc("key", 0, 16), cli.BoolArgFunc("value", 1)),
		PreRun: connectTransactor,
		Run: func(cmd *cobra.Command, args []string) {
			key, _ := cli.Bytes32FromArg(args[0])
			value, _ := strconv.ParseBool(args[1])
			cli.PrintTransactionFn(cmd)(transSession.SetBool(key, value))
		},
	},
	&cobra.Command{
		Use:    "deleteBool <key>",
		Short:  "Deletes the bool for the <key>",
		Args:   cli.HexArgLenFunc("key", 0, 16),
		PreRun: connectTransactor,
		Run: func(cmd *cobra.Command, args []string) {
			key, _ := cli.Bytes32FromArg(args[0])
			cli.PrintTransactionFn(cmd)(transSession.DeleteBool(key))
		},
	},

	&cobra.Command{
		Use:    "setBytes32 <key> <value>",
		Short:  "Sets the bytes32 at <key> to the <value>",
		Args:   cli.ChainArgs(cli.HexArgLenFunc("key", 0, 16), cli.HexArgLenFunc("value", 1, 16)),
		PreRun: connectTransactor,
		Run: func(cmd *cobra.Command, args []string) {
			key, _ := cli.Bytes32FromArg(args[0])
			value, _ := cli.Bytes32FromArg(args[1])
			cli.PrintTransactionFn(cmd)(transSession.SetBytes32(key, value))
		},
	},
	&cobra.Command{
		Use:    "deleteBytes32 <key>",
		Short:  "Deletes the bytes32 for the <key>",
		Args:   cli.HexArgLenFunc("key", 0, 16),
		PreRun: connectTransactor,
		Run: func(cmd *cobra.Command, args []string) {
			key, _ := cli.Bytes32FromArg(args[0])
			cli.PrintTransactionFn(cmd)(transSession.DeleteBytes32(key))
		},
	},

	&cobra.Command{
		Use:    "setBytes <key> <value>",
		Short:  "Sets the bytes at <key> to the <value>",
		Args:   cli.ChainArgs(cli.HexArgLenFunc("key", 0, 16), cli.HexArgFunc("value", 1)),
		PreRun: connectTransactor,
		Run: func(cmd *cobra.Command, args []string) {
			key, _ := cli.Bytes32FromArg(args[0])
			value, _ := hex.DecodeString(args[0])
			cli.PrintTransactionFn(cmd)(transSession.SetBytes(key, value))
		},
	},
	&cobra.Command{
		Use:    "deleteBytes <key>",
		Short:  "Deletes the bytes for the <key>",
		Args:   cli.HexArgLenFunc("key", 0, 16),
		PreRun: connectTransactor,
		Run: func(cmd *cobra.Command, args []string) {
			key, _ := cli.Bytes32FromArg(args[0])
			cli.PrintTransactionFn(cmd)(transSession.DeleteBytes(key))
		},
	},

	&cobra.Command{
		Use:    "setint256 <key> <value>",
		Short:  "Sets the int256 at <key> to the <value>",
		Args:   cli.ChainArgs(cli.HexArgLenFunc("key", 0, 16), cli.BigIntArgFunc("value", 1)),
		PreRun: connectTransactor,
		Run: func(cmd *cobra.Command, args []string) {
			key, _ := cli.Bytes32FromArg(args[0])
			value, _ := new(big.Int).SetString(args[1], 10)
			cli.PrintTransactionFn(cmd)(transSession.SetInt256(key, value))
		},
	},
	&cobra.Command{
		Use:    "deleteint256 <key>",
		Short:  "Deletes the int256 for the <key>",
		Args:   cli.HexArgLenFunc("key", 0, 16),
		PreRun: connectTransactor,
		Run: func(cmd *cobra.Command, args []string) {
			key, _ := cli.Bytes32FromArg(args[0])
			cli.PrintTransactionFn(cmd)(transSession.DeleteInt256(key))
		},
	},

	&cobra.Command{
		Use:    "setString <key> <value>",
		Short:  "Sets the string at <key> to the <value>",
		Args:   cli.ChainArgs(cobra.ExactArgs(2), cli.HexArgLenFunc("key", 0, 16)),
		PreRun: connectTransactor,
		Run: func(cmd *cobra.Command, args []string) {
			key, _ := cli.Bytes32FromArg(args[0])
			value := args[1]
			cli.PrintTransactionFn(cmd)(transSession.SetString(key, value))
		},
	},
	&cobra.Command{
		Use:    "deleteString <key>",
		Short:  "Deletes the string for the <key>",
		Args:   cli.HexArgLenFunc("key", 0, 16),
		PreRun: connectTransactor,
		Run: func(cmd *cobra.Command, args []string) {
			key, _ := cli.Bytes32FromArg(args[0])
			cli.PrintTransactionFn(cmd)(transSession.DeleteString(key))
		},
	},

	&cobra.Command{
		Use:    "setUint256 <key> <value>",
		Short:  "Sets the uint256 at <key> to the <value>",
		Args:   cli.ChainArgs(cli.HexArgLenFunc("key", 0, 16), cli.BigIntArgFunc("value", 1)),
		PreRun: connectTransactor,
		Run: func(cmd *cobra.Command, args []string) {
			key, _ := cli.Bytes32FromArg(args[0])
			value, _ := new(big.Int).SetString(args[1], 10)
			cli.PrintTransactionFn(cmd)(transSession.SetUint256(key, value))
		},
	},
	&cobra.Command{
		Use:    "deleteUint256 <key>",
		Short:  "Deletes the uint256 for the <key>",
		Args:   cli.HexArgLenFunc("key", 0, 16),
		PreRun: connectTransactor,
		Run: func(cmd *cobra.Command, args []string) {
			key, _ := cli.Bytes32FromArg(args[0])
			cli.PrintTransactionFn(cmd)(transSession.DeleteUint256(key))
		},
	},
}

func init() {
	// Add the Administrable, Destroyable, Lockable, Ownable contract getter commands
	SetterCommands = append(SetterCommands, destroyable.NewSetterCommands(contractKey)...)
	SetterCommands = append(SetterCommands, lockable.NewSetterCommands(contractKey)...)
	SetterCommands = append(SetterCommands, ownable.NewSetterCommands(contractKey)...)

	for _, cmd := range SetterCommands {
		// Add the 'gasPrice', 'nonce' and 'wait' flags to the deploy function
		gas.Flag(cmd)
		nonce.Flag(cmd)
		cli.WaitFlag(cmd)

		// Allow providing contract 'address' flag
		cmd.Flags().String("address", "", `address of the ComplianceStorage contract (default "[`+contractKey+`] value from config")`)
	}
}
