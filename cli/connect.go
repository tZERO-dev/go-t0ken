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

type password string

func getAddress(cmd *cobra.Command) (common.Address, string, bool, error) {
	// If an address was specified as a flag, don't perform password checking
	f, err := cmd.Flags().GetString("keystoreAddress")
	if common.IsHexAddress(f) || err != nil {
		return common.HexToAddress(f), "", false, err
	}

	// If the config value is an address then we have no alias/password to process
	f = viper.GetString("keystoreAddress")
	if common.IsHexAddress(f) {
		return common.HexToAddress(f), "", false, err
	}

	// Get the address and optional password, for the alias
	return AddressForKeystoreAlias(f)
}

// Connect creates a connection to the node
func Connect(cmd *cobra.Command, args []string) {
	var err error
	cmd.Println(viper.GetString("Url"))
	Conn, err = connection.New(viper.GetString("url"))
	CheckErr(cmd, err)
}

// Connect creates a connection to the node, using the address from the keystore as the transactor
func ConnectWithKeyStore(cmd *cobra.Command, args []string) {
	if viper.GetString("url") == "" || viper.GetString("keystore") == "" || viper.GetString("keystoreAddress") == "" {
		CheckErr(cmd, errors.New("`url`, `keystore` and `keystoreAddress` flags are required"))
	}

	// Get keystore address and password
	addr, password, hasPassword, err := getAddress(cmd)
	if err == nil && !hasPassword {
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

// NewConnection returns a connection using the URL from 't0ken.yaml'
func NewConnection(cmd *cobra.Command, args []string) (*connection.Connection, error) {
	url := viper.GetString("url")
	if url == "" {
		return nil, errors.New("missing 'url' flag or configuration value")
	}
	return connection.New(url)
}

// NewKeystoreConnection returns a connection using the URL, address, and optional password from 't0ken.yaml'
func NewKeystoreConnection(cmd *cobra.Command, args []string) (*connection.Connection, error) {
	conn, err := NewConnection(cmd, args)
	if err != nil {
		return conn, err
	}

	keystore, keystoreAddress := viper.GetString("keystore"), viper.GetString("keystoreAddress")
	if keystore == "" || keystoreAddress == "" {
		return conn, errors.New("missing one, or both of, 'keystore' and 'keystoreAddress' flags or configuration values")
	}

	// Get keystore address and password
	addr, password, hasPassword, err := getAddress(cmd)
	if err == nil && !hasPassword {
		password, err = console.Stdin.PromptPassword("Password: ")
	}
	if err != nil {
		return conn, err
	}

	// Set the transactor to the given keystore address/password
	err = conn.SetTransactorFromKeyStore(addr, password)
	return conn, err
}
