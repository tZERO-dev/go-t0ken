package cli

import (
	"errors"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/console"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/tzero-dev/go-t0ken/connection"
)

// Conn is a connection to the node
var Conn *connection.Connection

// getAddress retrieves either the 'keystoreAddress' or the matching address from 'keystoreAddressAliases'
func getAddress() (common.Address, error) {
	s := viper.GetString("keystoreAddress")
	if common.IsHexAddress(s) {
		return common.HexToAddress(s), nil
	}
	return addressForKeystoreAlias(s)
}

// Connect creates a connection to the node
func Connect(cmd *cobra.Command, args []string) {
	var err error
	Conn, err = connection.New(viper.GetString("url"))
	CheckErr(cmd, err)
}

// Connect creates a connection to the node, using the address from the keystore as the transactor
func ConnectWithKeyStore(cmd *cobra.Command, args []string) {
	if viper.GetString("url") == "" || viper.GetString("keystore") == "" || viper.GetString("keystoreAddress") == "" {
		CheckErr(cmd, errors.New("`url`, `keystore` and `keystoreAddress` flags are required"))
	}

	// Get keystore address and password
	var password string
	addr, err := getAddress()
	if err == nil {
		password, err = console.Stdin.PromptPassword("Password: ")
	}

	// Connect using the keystore
	if err == nil {
		url := viper.GetString("url")
		keystore := viper.GetString("keystore")
		Conn, err = connection.NewForKeyStore(url, keystore, addr, password)
	}
	CheckErr(cmd, err)
}
