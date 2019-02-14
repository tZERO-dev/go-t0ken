package cli

import (
	"math/big"
	"os"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/spf13/cobra"
)

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
		CheckErr(cmd, err)
		cmd.Println(o)
	}
}

// CheckGetter prints the given error and exits, if the error is not nil, otherwise the returned address is printed.
func CheckAddressGetter(cmd *cobra.Command) func(common.Address, error) {
	return func(addr common.Address, err error) {
		CheckErr(cmd, err)
		cmd.Println(addr.String())
	}
}

func CheckAccreditationGetter(cmd *cobra.Command) func(*big.Int, error) {
	return func(accreditation *big.Int, err error) {
		CheckErr(cmd, err)
		t := time.Unix(accreditation.Int64(), 0).UTC()
		cmd.Printf(" date: %s\nepoch: %s\n", t, accreditation)
	}
}

func CheckCountryGetter(cmd *cobra.Command) func([2]byte, error) {
	return func(country [2]byte, err error) {
		CheckErr(cmd, err)
		cmd.Println(string(country[:]))
	}
}

func CheckHashGetter(cmd *cobra.Command) func([32]byte, error) {
	return func(hash [32]byte, err error) {
		CheckErr(cmd, err)
		cmd.Printf("0x%x\n", hash[:])
	}
}
