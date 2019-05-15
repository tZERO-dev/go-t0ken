package cli

import (
	"encoding/hex"
	"errors"
	"fmt"
	"math/big"
	"strconv"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
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
		return addr, fmt.Errorf("requires address arg")
	}

	if !common.IsHexAddress(args[index]) {
		return addr, errors.New("invalid address for <address> arg")
	}
	return common.HexToAddress(args[index]), nil

}

// IsAddress checks if the given string is a valid address.
func IsAddress(s string) error {
	if !common.IsHexAddress(s) {
		return fmt.Errorf("received '%s' but expected a valid address", s)
	}
	return nil
}

// IsAddressArg validates that the args contains enough values, and if the given argument position is a valid address.
func IsAddressArg(key string, args []string, i int) error {
	if len(args) < i+1 {
		return fmt.Errorf("requires <%s> arg", key)
	}
	return IsAddress(args[i])
}

// IsIntArg validates that the args contains enough values, and if the given argument position is an int.
func IsIntArg(key string, args []string, i int) error {
	if len(args) < i+1 {
		return fmt.Errorf("requires <%s> arg", key)
	}
	_, err := strconv.Atoi(args[i])
	return err
}

// BoolArgFunc returns a cobra.Command Arg function to validate the given argument index is a bool.
func BoolArgFunc(key string, index int) func(*cobra.Command, []string) error {
	return func(cmd *cobra.Command, args []string) error {
		if len(args) < index+1 {
			return fmt.Errorf("requires <%s> arg", key)
		}
		_, err := strconv.ParseBool(args[index])
		return err
	}
}

// UintArgFunc returns a cobra.Command Arg function to validate the given argument index is an uint of the specified bit size.
func UintArgFunc(key string, index, bitSize int) func(*cobra.Command, []string) error {
	return func(cmd *cobra.Command, args []string) error {
		if len(args) < index+1 {
			return fmt.Errorf("requires <%s> arg", key)
		}
		_, err := strconv.ParseInt(args[index], 10, bitSize)
		return err
	}
}

// BoolArgFunc returns a cobra.Command Arg function to validate the given argument index is an int.
func IntArgFunc(key string, index int) func(*cobra.Command, []string) error {
	return func(cmd *cobra.Command, args []string) error {
		return IsIntArg(key, args, index)
	}
}

// BoolArgFunc returns a cobra.Command Arg function to validate the given argument index is a big.Int.
func BigIntArgFunc(key string, index int) func(*cobra.Command, []string) error {
	return func(cmd *cobra.Command, args []string) error {
		if len(args) < index+1 {
			return fmt.Errorf("requires <%s> arg", key)
		}
		_, ok := new(big.Int).SetString(args[index], 10)
		if !ok {
			return fmt.Errorf("requires valid <%s> arg", key)
		}
		return nil
	}
}

// BoolArgFunc returns a cobra.Command Arg function to validate the given argument index is a big.Float.
func BigFloatArgFunc(key string, index int) func(*cobra.Command, []string) error {
	return func(cmd *cobra.Command, args []string) error {
		if len(args) < index+1 {
			return fmt.Errorf("requires <%s> arg", key)
		}
		_, ok := new(big.Float).SetString(args[index])
		if !ok {
			return fmt.Errorf("requires valid <%s> arg", key)
		}
		return nil
	}
}

// BoolArgFunc returns a cobra.Command Arg function to validate the given argument index is a common.Address.
func AddressArgFunc(key string, index int) func(*cobra.Command, []string) error {
	return func(cmd *cobra.Command, args []string) error {
		return IsAddressArg(key, args, index)
	}
}

// AddressArgOrKeystoreAddressFunc returns a cobra.Command Arg function to validate that either the 'keystoreAddress' arg is set, or a config file value is set.
func AddressArgOrKeystoreAddressFunc(key string, index int) func(*cobra.Command, []string) error {
	return func(cmd *cobra.Command, args []string) error {
		err := IsAddressArg(key, args, index)
		if err != nil {
			err = IsAddress(viper.GetString("keystoreAddress"))
		}
		return err
	}
}

// CountryFromArg returns the country code in bytes.
func CountryFromArg(s string) ([2]byte, error) {
	var country [2]byte
	var b = []byte(s)

	if len(country) != len(b) {
		return country, fmt.Errorf("country code '%s' is not valid", b)
	}
	copy(country[:], b)
	return country, nil
}

// CountryCodeArgFunc returns a cobra.Command Arg function that validates the argument at the given index is a valid country code.
func CountryCodeArgFunc(key string, index int) func(*cobra.Command, []string) error {
	return func(cmd *cobra.Command, args []string) error {
		if len(args) < index+1 {
			return fmt.Errorf("requires <%s> arg", key)
		}

		_, err := CountryFromArg(args[index])
		return err
	}
}

// HexArgFunc returns a cobra.Command Arg function that validates the argument at the given index is a hex string of the given length.
func HexArgFunc(key string, index int, length int) func(*cobra.Command, []string) error {
	return func(cmd *cobra.Command, args []string) error {
		if len(args) < index+1 {
			return fmt.Errorf("requires <%s> arg", key)
		}

		b, err := hexutil.Decode(args[index])
		l := hex.DecodedLen(len(b))
		if err == nil && l != length {
			err = fmt.Errorf("hex arg <%s> data size %d is different than the required length %d", key, l, length)
		}
		return err
	}
}

// DateFromArg returns the parsed string as a date.
func DateFromArg(s string) (time.Time, error) {
	return time.Parse("2006-01-02", s)
}

// DateArgFunc returns a cobra.Command Arg function that validates the argument at the given index is a valid date.
func DateArgFunc(key string, index int) func(*cobra.Command, []string) error {
	return func(cmd *cobra.Command, args []string) error {
		if len(args) < index+1 {
			return fmt.Errorf("requires <%s> arg", key)
		}
		_, err := DateFromArg(args[index])
		return err
	}
}
