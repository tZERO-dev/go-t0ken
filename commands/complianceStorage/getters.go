package complianceStorage

import (
	"encoding/hex"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/spf13/cobra"

	"github.com/tzero-dev/go-t0ken/cli"
	"github.com/tzero-dev/go-t0ken/commands/lockable"
	"github.com/tzero-dev/go-t0ken/commands/ownable"
	"github.com/tzero-dev/go-t0ken/contracts/compliance"
)

var GetterCommands = []*cobra.Command{
	&cobra.Command{
		Use:     "abi",
		Short:   "Outputs the ComplianceStorage ABI",
		Example: "t0ken complianceStorage abi",
		Args:    cobra.NoArgs,
		Run:     func(cmd *cobra.Command, args []string) { cmd.Println(compliance.StorageABI) },
	},
	&cobra.Command{
		Use:     "bin",
		Short:   "Outputs the ComplianceStorage Binary",
		Example: "t0ken complianceStorage bin",
		Args:    cobra.NoArgs,
		Run:     func(cmd *cobra.Command, args []string) { cmd.Println(compliance.StorageBin) },
	},
	&cobra.Command{
		Use:   "key <params>...",
		Short: "Generates the key for the given <params>",
		Example: `t0ken complianceStorage key "Compliance.rules" 8166409948e8827a8fdcb0e7365013e968a6dfa1 04 
t0ken complianceStorage key "Compliance.rules" 8166409948e8827a8fdcb0e7365013e968a6dfa1 04 | xargs t0ken complianceStorage getBytes`,
		Args: cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			var data [][]byte
			for _, s := range args {
				b, err := hex.DecodeString(s)
				if err != nil {
					b = []byte(s)
				}
				data = append(data, b)
			}
			cmd.Printf("0x%x\n", crypto.Keccak256(data...))
		},
	},
	&cobra.Command{
		Use:     "permissions",
		Short:   "Gets the total number of permissions",
		Example: "t0ken complianceStorage permissions",
		Args:    cobra.NoArgs,
		PreRun:  connectCaller,
		Run: func(cmd *cobra.Command, args []string) {
			cli.CheckGetter(cmd)(callSession.Permissions())
		},
	},
	&cobra.Command{
		Use:     "permissionAt <index>",
		Short:   "Gets the address with permission at the given <index>",
		Example: "t0ken complianceStorage permissionAt 0",
		Args:    cli.ChainArgs(cli.BigIntArgFunc("index", 0)),
		PreRun:  connectCaller,
		Run: func(cmd *cobra.Command, args []string) {
			index, _ := new(big.Int).SetString(args[0], 10)
			cli.CheckAddressGetter(cmd)(callSession.PermissionAt(index))
		},
	},
	&cobra.Command{
		Use:     "permissionIndexOf <address>",
		Short:   "Gets the permission index of the given <address>",
		Example: "t0ken complianceStorage permissionIndexOf 0x82ABc653f36Ce78E6290AF8D7e83823c8613bc47",
		Args:    cli.ChainArgs(cli.AddressArgFunc("address", 0)),
		PreRun:  connectCaller,
		Run: func(cmd *cobra.Command, args []string) {
			addr := common.HexToAddress(args[0])
			cli.CheckGetter(cmd)(callSession.PermissionIndexOf(addr))
		},
	},
	&cobra.Command{
		Use:    "getAddress <key>",
		Short:  "Gets the address for the given key",
		Args:   cli.HexArgLenFunc("key", 0, 16),
		PreRun: connectCaller,
		Run: func(cmd *cobra.Command, args []string) {
			key, _ := cli.Bytes32FromArg(args[0])
			cli.CheckAddressGetter(cmd)(callSession.GetAddress(key))
		},
	},
	&cobra.Command{
		Use:    "getBool <key>",
		Short:  "Gets the bool for the given key",
		Args:   cli.HexArgLenFunc("hash", 0, 16),
		PreRun: connectCaller,
		Run: func(cmd *cobra.Command, args []string) {
			key, _ := cli.Bytes32FromArg(args[0])
			cli.CheckGetter(cmd)(callSession.GetBool(key))
		},
	},
	&cobra.Command{
		Use:    "getBytes <key>",
		Short:  "Gets the bytes for the given key",
		Args:   cli.HexArgLenFunc("hash", 0, 16),
		PreRun: connectCaller,
		Run: func(cmd *cobra.Command, args []string) {
			key, _ := cli.Bytes32FromArg(args[0])
			cli.CheckBytesGetter(cmd)(callSession.GetBytes(key))
		},
	},
	&cobra.Command{
		Use:    "getBytes32 <key>",
		Short:  "Gets the bytes32 for the given key",
		Args:   cli.HexArgLenFunc("hash", 0, 16),
		PreRun: connectCaller,
		Run: func(cmd *cobra.Command, args []string) {
			key, _ := cli.Bytes32FromArg(args[0])
			cli.CheckGetter(cmd)(callSession.GetBytes32(key))
		},
	},
	&cobra.Command{
		Use:    "getString <key>",
		Short:  "Gets the string for the given key",
		Args:   cli.HexArgLenFunc("hash", 0, 16),
		PreRun: connectCaller,
		Run: func(cmd *cobra.Command, args []string) {
			key, _ := cli.Bytes32FromArg(args[0])
			cli.CheckGetter(cmd)(callSession.GetString(key))
		},
	},
	&cobra.Command{
		Use:    "getInt256 <key>",
		Short:  "Gets the int256 for the given key",
		Args:   cli.HexArgLenFunc("hash", 0, 16),
		PreRun: connectCaller,
		Run: func(cmd *cobra.Command, args []string) {
			key, _ := cli.Bytes32FromArg(args[0])
			cli.CheckGetter(cmd)(callSession.GetInt256(key))
		},
	},
	&cobra.Command{
		Use:    "getUint256 <key>",
		Short:  "Gets the uint256 for the given key",
		Args:   cli.HexArgLenFunc("hash", 0, 16),
		PreRun: connectCaller,
		Run: func(cmd *cobra.Command, args []string) {
			key, _ := cli.Bytes32FromArg(args[0])
			cli.CheckGetter(cmd)(callSession.GetUint256(key))
		},
	},
}

func init() {
	// Add the Administrable, Ownable, Lockable contract getter commands
	GetterCommands = append(GetterCommands, lockable.NewGetterCommands(contractKey)...)
	GetterCommands = append(GetterCommands, ownable.NewGetterCommands(contractKey)...)

	for i, cmd := range GetterCommands {
		// Skip ABI/Bin
		if i < 2 {
			continue
		}

		// Allow providing contract 'address' flag
		cmd.Flags().String("address", "", `address of the ComplianceStorage contract (default "[`+contractKey+`] value from config")`)
		cli.BlockFlag(cmd)
	}
}
