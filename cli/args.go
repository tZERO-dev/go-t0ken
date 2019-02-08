package cli

import (
	"errors"
	"math/big"
	"strconv"

	"github.com/ethereum/go-ethereum/common"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// ChainArgs chains the given args functions for use as a single cobra.Command Args.
func ChainArgs(funcs ...func(cmd *cobra.Command, args []string) error) func(*cobra.Command, []string) error {
	return func(cmd *cobra.Command, args []string) error {
		var err error
		for _, fn := range funcs {
			err = fn(cmd, args)
			if err != nil {
				break
			}
		}
		return err
	}
}

// GetArgAddress returns the address for the argument at the specified index.
func GetArgAddress(index int, args []string) (common.Address, error) {
	var addr common.Address
	if len(args) < index+1 {
		if !common.IsHexAddress(viper.GetString("keystoreAddress")) {
			return addr, errors.New("missing one of <address> or <keystoreAddress> args")
		}
		return common.HexToAddress(viper.GetString("keystoreAddress")), nil
	}
	if !common.IsHexAddress(args[index]) {
		return addr, errors.New("invalid address for <address> arg")
	}
	return common.HexToAddress(args[index]), nil

}

// IsAddress checks if the given string is a valid address.
func IsAddress(s string) error {
	if !common.IsHexAddress(s) {
		return errors.New("argument must be a valid address")
	}
	return nil
}

// IsAddressArg validates that the args contains enough values, and if the given argument position is a valid address.
func IsAddressArg(args []string, i int) error {
	if len(args) < i+1 {
		return errors.New("requires <address> arg")
	}
	return IsAddress(args[i])
}

// IsAddressArg validates that the args contains enough values, and if the given argument position is an int.
func IsIntArg(args []string, i int) error {
	if len(args) < i+1 {
		return errors.New("requires <int> arg")
	}
	_, err := strconv.Atoi(args[i])
	return err
}

// BoolArgFunc returns a cobra.Command Arg function to validate the given argument index is a bool.
func BoolArgFunc(index int) func(*cobra.Command, []string) error {
	return func(cmd *cobra.Command, args []string) error {
		if len(args) < index+1 {
			return errors.New("requires <bool> arg")
		}
		_, err := strconv.ParseBool(args[index])
		return err
	}
}

// UintArgFunc returns a cobra.Command Arg function to validate the given argument index is an uint of the specified bit size.
func UintArgFunc(index, bitSize int) func(*cobra.Command, []string) error {
	return func(cmd *cobra.Command, args []string) error {
		if len(args) < index+1 {
			return errors.New("requires <int> arg")
		}
		_, err := strconv.ParseInt(args[index], 10, bitSize)
		return err
	}
}

// BoolArgFunc returns a cobra.Command Arg function to validate the given argument index is an int.
func IntArgFunc(index int) func(*cobra.Command, []string) error {
	return func(cmd *cobra.Command, args []string) error {
		return IsIntArg(args, index)
	}
}

// BoolArgFunc returns a cobra.Command Arg function to validate the given argument index is a big.Int.
func BigIntArgFunc(index int) func(*cobra.Command, []string) error {
	return func(cmd *cobra.Command, args []string) error {
		if len(args) < index+1 {
			return errors.New("requires <int> arg")
		}
		_, ok := new(big.Int).SetString(args[index], 10)
		if !ok {
			return errors.New("requires valid <int> arg")
		}
		return nil
	}
}

// BoolArgFunc returns a cobra.Command Arg function to validate the given argument index is a big.Float.
func BigFloatArgFunc(index int) func(*cobra.Command, []string) error {
	return func(cmd *cobra.Command, args []string) error {
		if len(args) < index+1 {
			return errors.New("requires <float> arg")
		}
		_, ok := new(big.Float).SetString(args[index])
		if !ok {
			return errors.New("requires valid <float> arg")
		}
		return nil
	}
}

// BoolArgFunc returns a cobra.Command Arg function to validate the given argument index is a common.Address.
func AddressArgFunc(index int) func(*cobra.Command, []string) error {
	return func(cmd *cobra.Command, args []string) error {
		return IsAddressArg(args, index)
	}
}

// AddressArgOrKeystoreAddressFunc returns a cobra.Command Arg function to validate that either the 'keystoreAddress' arg is set, or a config file value is set.
func AddressArgOrKeystoreAddressFunc(index int) func(*cobra.Command, []string) error {
	return func(cmd *cobra.Command, args []string) error {
		err := IsAddressArg(args, index)
		if err != nil {
			err = IsAddress(viper.GetString("keystoreAddress"))
		}
		return err
	}
}
