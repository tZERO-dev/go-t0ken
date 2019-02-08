package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/tzero-dev/go-t0ken/commands/ether"
	"github.com/tzero-dev/go-t0ken/commands/gas"
	"github.com/tzero-dev/go-t0ken/commands/nonce"
	"github.com/tzero-dev/go-t0ken/commands/token"
)

const VERSION = "0.0.1"

var configFile string

var rootCmd = &cobra.Command{
	Use:   "t0ken [OPTIONS] COMMAND [ARG...]",
	Short: "Utility functions for use with tZERO Ethereum smart contracts",
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Displays the t0ken version",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Println(VERSION)
	},
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&configFile, "config", "t0ken.yaml", "config file")
	rootCmd.PersistentFlags().String("url", "http://localhost:8545", "URL to an Ethereum node")
	rootCmd.PersistentFlags().String("keystore", ".keystore", "keystore path")
	rootCmd.PersistentFlags().String("keystoreAddress", "", "address to use within the keystore")

	viper.BindPFlag("url", rootCmd.PersistentFlags().Lookup("url"))
	viper.BindPFlag("keystore", rootCmd.PersistentFlags().Lookup("keystore"))
	viper.BindPFlag("keystoreAddress", rootCmd.PersistentFlags().Lookup("keystoreAddress"))
}

func initConfig() {
	if configFile != "" {
		viper.SetConfigFile(configFile)
	}
	//viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Printf("Failed to read config: %s - %s\n", configFile, err)
		os.Exit(1)
	}
}

func main() {
	rootCmd.AddCommand(versionCmd)

	// Gas
	gas.Command.AddCommand(gas.PriceCommand)
	rootCmd.AddCommand(gas.Command)

	// Nonce
	nonce.Command.AddCommand(nonce.NextCommand)
	rootCmd.AddCommand(nonce.Command)

	// Ether
	ether.Command.AddCommand(ether.SendCommand)
	ether.Command.AddCommand(ether.BalanceCommand)
	rootCmd.AddCommand(ether.Command)

	// Token
	token.Command.AddCommand(token.GetterCommands...)
	token.Command.AddCommand(token.SetterCommands...)
	token.Command.AddCommand(token.DeployCommand)
	rootCmd.AddCommand(token.Command)

	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
