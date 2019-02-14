package cli

import (
	"errors"

	"github.com/ethereum/go-ethereum/common"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// FlagOrConfigAddress returns the address of the given flag, when set, or the matching address for the given key, within the config file.
func FlagOrConfigAddress(cmd *cobra.Command, flag, configKey string) (common.Address, error) {
	// Get the flag or config value
	s, err := cmd.Flags().GetString(flag)
	if s == "" {
		s = viper.GetString(configKey)
	}
	// If neither the flag or config exists, and we have an error
	if s == "" && err != nil {
		return common.Address{}, err
	}

	// Convert to an address
	err = IsAddress(s)
	if err != nil {
		return common.Address{}, err
	}
	return common.HexToAddress(s), err
}

// addressForKeystoreAlias returns the address in the config 'keystoreAddressAliases' secion of the given alias
func addressForKeystoreAlias(alias string) (common.Address, error) {
	var addr common.Address

	// Make sure configuration contains a valid entry for 'keystoreAddressAliases'
	o := viper.Get("keystoreAddressAliases")
	m, ok := o.(map[string]interface{})
	if !ok {
		return addr, errors.New("Invalid, or improperly setup 'keystoreAddressAliases' YAML configuration")
	}

	// Check if the alias exists
	v, ok := m[alias]
	if !ok {
		return addr, errors.New("Missing key '%s' within the 'keystoreAddressAliases' YAML configuration")
	}

	// Verify the value is a valid address
	s, ok := v.(string)
	if !ok {
		return addr, errors.New("Values for 'keystoreAddressAliases' must be strings")
	}
	err := addr.UnmarshalText([]byte(s))
	return addr, err
}
