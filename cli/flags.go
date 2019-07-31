package cli

import (
	"errors"
	"fmt"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// FlagAddress returns the address of the given flag, when set.
func OptionalFlagAddress(cmd *cobra.Command, flag string) (common.Address, error) {
	// Get the flag or config value
	s, err := cmd.Flags().GetString(flag)
	if s == "" || err != nil {
		return common.Address{}, err
	}
	if !common.IsHexAddress(s) {
		return common.Address{}, err
	}
	return common.HexToAddress(s), err
}

// FlagOrConfigAddress returns the address of the given flag, when set, or the matching address for the given key, within the config file.
func FlagOrConfigAddress(cmd *cobra.Command, flag, configKey string) (common.Address, error) {
	// Get the flag or config value
	s, err := cmd.Flags().GetString(flag)
	if s == "" {
		s = viper.GetString(configKey)
	}
	// If neither the flag or config exists, and we have an error
	if s == "" && err != nil {
		return common.Address{}, fmt.Errorf("missing required flag '%s' or config key '%s'", flag, configKey)
	}

	// Convert to an address
	err = IsAddress(s)
	if err != nil {
		return common.Address{}, fmt.Errorf("'%s' is not a valid address for flag '%s' or '%s' config key", s, flag, configKey)
	}
	return common.HexToAddress(s), err
}

// AddressForKeystoreAlias returns the address in the config 'keystoreAddressAliases' secion of the given alias
func AddressForKeystoreAlias(alias string) (common.Address, string, bool, error) {
	var addr common.Address
	var passphrase string
	var hasPassphrase bool

	// Make sure configuration contains a valid entry for 'keystoreAddressAliases'
	o := viper.Get("keystoreAddressAliases")
	m, ok := o.(map[string]interface{})
	if !ok {
		return addr, passphrase, hasPassphrase, errors.New("Invalid, or improperly setup 'keystoreAddressAliases' YAML configuration")
	}

	// Check if the alias exists
	v, ok := m[alias]
	if !ok {
		return addr, passphrase, hasPassphrase, fmt.Errorf("Missing key '%s' within the 'keystoreAddressAliases' YAML configuration", alias)
	}

	// Verify the value is a valid address
	s, ok := v.(string)
	if !ok {
		return addr, passphrase, hasPassphrase, errors.New("Values for 'keystoreAddressAliases' must be strings")
	}

	p := strings.SplitN(s, ",", 2)
	hasPassphrase = len(p) == 2
	if hasPassphrase {
		s = p[0]
		passphrase = p[1]
	}
	err := addr.UnmarshalText([]byte(s))
	return addr, passphrase, hasPassphrase, err
}

// WaitFlag adds the 'wait' flag to the given command, to wait for 'n' confirmations..
func WaitFlag(cmd *cobra.Command) {
	cmd.Flags().Int("wait", -1, "waits the provided number of confirmations ('0' for accepted only)")
}
