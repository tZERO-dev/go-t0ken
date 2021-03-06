package cli

import (
	"errors"
	"math/big"
	"os"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/spf13/cobra"
)

// Error signature - first four bytes of `Error(string)` keccak256
var errorSig = []byte{0x08, 0xc3, 0x79, 0xa0}

// CheckErr prints the given error and exits, if the error is not nil.
func CheckErr(cmd *cobra.Command, err error) {
	if err != nil {
		cmd.Println(err)
		os.Exit(1)
	}
}

// CheckGetter prints the given error and exits, if the error is not nil, otherwise the return value is printed.
func CheckGetter(cmd *cobra.Command) func(interface{}, error) {
	return func(o interface{}, err error) {
		if err != nil {
			s := err.Error()
			if len(s) > 39 && s[34:38] == string(errorSig) {
				// Line 85 of github.com/ethereum/go-ethereum/accounts/abi/abi.go, for the `Unpack` function has the following error format string:
				//   `abi: improperly formatted output: %s - Bytes: [%+v]`
				// We're looking for the `%s` value, which is the error message
				i := strings.Index(s, " - Bytes: [")
				err = errors.New(s[40:i])
			}
		}
		CheckErr(cmd, err)
		cmd.Println(o)
	}
}

// CheckAddressesGetter prints the address.
func CheckAddressGetter(cmd *cobra.Command) func(common.Address, error) {
	return func(addr common.Address, err error) {
		CheckErr(cmd, err)
		cmd.Println(addr.String())
	}
}

// CheckAddressesGetter prints the collection of addresses.
func CheckAddressesGetter(cmd *cobra.Command) func([]common.Address, error) {
	return func(s []common.Address, err error) {
		CheckErr(cmd, err)
		for _, addr := range s {
			cmd.Println(addr.String())
		}
	}
}

// CheckAccreditationGetter prints the UTC date and EPOCH.
func CheckAccreditationGetter(cmd *cobra.Command) func(*big.Int, error) {
	return func(accreditation *big.Int, err error) {
		CheckErr(cmd, err)

		var t time.Time
		a := accreditation.Int64()
		if a > 999999999999 {
			t = time.Unix(0, a*int64(time.Millisecond)).UTC()
		} else {
			t = time.Unix(a, 0).UTC()
		}
		cmd.Printf(" date: %s\nepoch: %s\n", t, accreditation)
	}
}

// CheckCountryGetter prints the country code.
func CheckCountryGetter(cmd *cobra.Command) func([2]byte, error) {
	return func(country [2]byte, err error) {
		CheckErr(cmd, err)
		cmd.Println(string(country[:]))
	}
}

// CheckBytes32Getter prints the return hex
func CheckBytes32Getter(cmd *cobra.Command) func([32]byte, error) {
	return func(b [32]byte, err error) {
		CheckErr(cmd, err)
		cmd.Printf("0x%x\n", b[:])
	}
}

// CheckBytes32Getter prints the return hex
func CheckBytesGetter(cmd *cobra.Command) func([]byte, error) {
	return func(b []byte, err error) {
		CheckErr(cmd, err)
		cmd.Printf("0x%x\n", b)
	}
}
