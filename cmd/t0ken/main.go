package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/tzero-dev/go-t0ken"
	"github.com/tzero-dev/go-t0ken/commands/broker"
	"github.com/tzero-dev/go-t0ken/commands/compliance"
	"github.com/tzero-dev/go-t0ken/commands/custodian"
	"github.com/tzero-dev/go-t0ken/commands/ether"
	"github.com/tzero-dev/go-t0ken/commands/gas"
	"github.com/tzero-dev/go-t0ken/commands/investor"
	"github.com/tzero-dev/go-t0ken/commands/nonce"
	"github.com/tzero-dev/go-t0ken/commands/storage"
	"github.com/tzero-dev/go-t0ken/commands/token"
	"github.com/tzero-dev/go-t0ken/commands/transaction"
)

var configFile string

var rootCmd = &cobra.Command{
	Use:   "t0ken [OPTIONS] COMMAND [ARG...]",
	Short: "Utility functions for use with tZERO Ethereum smart contracts",
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Displays the t0ken version",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Println(t0ken.VERSION)
	},
}

var completionCmd = &cobra.Command{
	Use:   "completion [shell] (default 'bash' or 'zsh')",
	Short: "Generates shell completion scripts",
	Long: `To load completion run

# --- bash ---
. <(t0ken completion)

To configure your bash shell to load completions for each session add to your bashrc

# ~/.bashrc or ~/.profile
. <(t0ken completion)

# --- zsh ---
t0ken completion zsh > /fpath/location/_t0ken
`,
	Run: func(cmd *cobra.Command, args []string) {
		genFn := rootCmd.GenBashCompletion
		if len(args) > 0 {
			switch args[0] {
			case "bash":
			case "zsh":
				genFn = rootCmd.GenZshCompletion
			default:
				cmd.Println("Invalid shell name, expecting one of 'bash', 'zsh'")
				os.Exit(1)
			}
		}
		genFn(os.Stdout)
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
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Printf("Failed to read config: %s - %s\n", configFile, err)
		os.Exit(1)
	}
}

func main() {
	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(completionCmd)

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

	// Transaction
	transaction.Command.AddCommand(transaction.ReplayCommand)
	transaction.Command.AddCommand(transaction.CancelCommand)
	transaction.Command.AddCommand(transaction.NextAddressCommand)
	transaction.Command.AddCommand(transaction.WaitCommand)
	rootCmd.AddCommand(transaction.Command)

	// Token
	token.Command.AddCommand(token.DeployCommand)
	token.Command.AddCommand(token.AuditCommand)
	token.Command.AddCommand(token.GetterCommands...)
	token.Command.AddCommand(token.SetterCommands...)
	rootCmd.AddCommand(token.Command)

	// Token-Compliance
	compliance.Command.AddCommand(compliance.DeployCommand)
	compliance.Command.AddCommand(compliance.GetterCommands...)
	compliance.Command.AddCommand(compliance.SetterCommands...)
	rootCmd.AddCommand(compliance.Command)

	// Registry, Storage
	storage.Command.AddCommand(storage.DeployCommand)
	storage.Command.AddCommand(storage.AuditCommand)
	storage.Command.AddCommand(storage.GetterCommands...)
	storage.Command.AddCommand(storage.SetterCommands...)
	rootCmd.AddCommand(storage.Command)

	// Registry, Custodian
	custodian.Command.AddCommand(custodian.DeployCommand)
	custodian.Command.AddCommand(custodian.GetterCommands...)
	custodian.Command.AddCommand(custodian.SetterCommands...)
	rootCmd.AddCommand(custodian.Command)

	// Registry, Broker-Dealer
	broker.Command.AddCommand(broker.DeployCommand)
	broker.Command.AddCommand(broker.GetterCommands...)
	broker.Command.AddCommand(broker.SetterCommands...)
	rootCmd.AddCommand(broker.Command)

	// Registry, Investor
	investor.Command.AddCommand(investor.DeployCommand)
	investor.Command.AddCommand(investor.GetterCommands...)
	investor.Command.AddCommand(investor.SetterCommands...)
	rootCmd.AddCommand(investor.Command)

	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
