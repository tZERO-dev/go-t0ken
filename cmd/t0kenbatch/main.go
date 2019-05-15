package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"math/big"
	"os"
	"strconv"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/tzero-dev/go-t0ken/cli"
	"github.com/tzero-dev/go-t0ken/commands"
	"github.com/tzero-dev/go-t0ken/contracts/registry"
	"github.com/tzero-dev/go-t0ken/contracts/token/erc20"
)

var one = new(big.Int).SetInt64(1)

var configFile string

var rootCmd = &cobra.Command{
	Use:   "t0kenbatch [OPTIONS] COMMAND [ARG...]",
	Short: "Batch functions for use with tZERO Ethereum smart contracts",
}

var completionCmd = &cobra.Command{
	Use:   "completion [shell] (default 'bash' or 'zsh')",
	Short: "Generates shell completion scripts",
	Long: `To load completion run

# --- bash ---
. <(t0kenbatch completion)

To configure your bash shell to load completions for each session add to your bashrc

# ~/.bashrc or ~/.profile
. <(t0kenbatch completion)

# --- zsh ---
t0kenbatch completion zsh > /fpath/location/_t0kenbatch
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

type header int

const (
	ADDRESS header = iota
	HASH
	COUNTRY
	ACCREDITATION
	TOKENS
)

type headers [5]int

type InvestorMsg struct {
	Investor Investor
	Error    error
}
type Investor struct {
	Address       common.Address
	Hash          [32]byte
	Country       [2]byte
	Accreditation *big.Int
	Tokens        *big.Int
}

func (i Investor) String() string {
	return fmt.Sprintf("%s, Hash: %#x, Country: %s, Accreditation: %s, Tokens: %s", i.Address.String(), i.Hash[:], string(i.Country[:]), i.Accreditation, i.Tokens.String())
}

func investorHeaders(s []string) (headers, error) {
	var h headers
	c := 0
	for i := 0; i < len(s); i++ {
		switch strings.ToUpper(s[i]) {
		case "ADDRESS":
			h[ADDRESS] = i
			c++
		case "HASH":
			h[HASH] = i
			c++
		case "COUNTRY":
			h[COUNTRY] = i
			c++
		case "ACCREDITATION":
			h[ACCREDITATION] = i
			c++
		case "TOKENS":
			h[TOKENS] = i
			c++
		}
	}
	var err error
	return h, err
}

func newInvestor(row []string, h headers, isIssuance bool) (Investor, error) {
	var i Investor

	if !common.IsHexAddress(row[h[ADDRESS]]) {
		return i, fmt.Errorf("invalid address: '%s'", row[h[ADDRESS]])
	}

	hash, err := hexutil.Decode(row[h[HASH]])
	if err != nil {
		return i, fmt.Errorf("invalid hash: '%s', %s", row[h[HASH]], err)
	}

	c, err := cli.CountryFromArg(row[h[COUNTRY]])
	if err != nil {
		return i, fmt.Errorf("invalid country: '%s', %s", row[h[COUNTRY]], err)
	}

	a, err := cli.DateFromArg(row[h[ACCREDITATION]])
	if err != nil {
		return i, fmt.Errorf("invalid accreditation: '%s', %s", row[h[ACCREDITATION]], err)
	}

	var tokens int
	if isIssuance {
		tokens, err = strconv.Atoi(row[h[TOKENS]])
		if err != nil {
			return i, fmt.Errorf("invalid tokens: '%s', %s", row[h[TOKENS]], err)
		}
	}

	var b [32]byte
	copy(b[:], hash)
	return Investor{
		Address:       common.HexToAddress(row[h[ADDRESS]]),
		Hash:          b,
		Country:       c,
		Accreditation: big.NewInt(a.Unix()),
		Tokens:        big.NewInt(int64(tokens)),
	}, nil
}

func parseInvestorCSV(r io.Reader, isIssuance bool) <-chan InvestorMsg {
	c := make(chan InvestorMsg)
	go func(ch chan<- InvestorMsg) {
		reader := csv.NewReader(r)
		index := -1
		var h headers
		for {
			index++
			row, err := reader.Read()
			switch {
			case err == io.EOF:
				close(ch)
				return
			case err != nil:
				ch <- InvestorMsg{Error: err}
				return
			case index == 0:
				h, err = investorHeaders(row)
				if err != nil {
					ch <- InvestorMsg{Error: err}
				}
				continue
			}

			i, err := newInvestor(row, h, isIssuance)
			if err != nil {
				err = fmt.Errorf("error processing at row: %d, %s", index, err)
			}
			ch <- InvestorMsg{Investor: i, Error: err}
		}
		close(ch)
	}(c)
	return c
}

func investorTransactorSessionFn(addr common.Address, transactor bind.ContractTransactor) (interface{}, error) {
	return registry.NewInvestorTransactor(addr, transactor)
}

var onboardInvestorCommand = &cobra.Command{
	Use:   "onboardInvestors <file>",
	Short: "Onboards the investors from the given CSV <file>",
	Long: `The CSV file must be in the following format:
---
address,hash,country,accreditation
0xf01ff29dcbee147e9ca151a281bfdf136f66a45b,0xa1896382c22b03c562b0241324cfca19505cc5c78eb06751d9cee690e21ed6a1,US,2012-08-24
---

 - The header column must be present, with all values being required.
 - The country must be in "ISO 3166-1 alpha-2" format
 - The accreditation must UTC and be in "YYYY-MM-DD" format`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		// Open file
		f, err := os.Open(args[0])
		cli.CheckErr(cmd, err)

		// Connect to node
		o, transactOpts := commands.ConnectWithTransactorSessionFunc(cmd, args, "investorRegistry", investorTransactorSessionFn)
		transactor := o.(*registry.InvestorTransactor)
		transSession := &registry.InvestorTransactorSession{transactor, transactOpts}
		nonce := cli.Conn.Opts.Nonce

		// Onboard
		var transactions []common.Hash
		for msg := range parseInvestorCSV(f, false) {
			if msg.Error != nil {
				cmd.Println(msg.Error)
				break
			}

			i := msg.Investor
			tx, err := transSession.Add(i.Address, i.Hash, i.Country, i.Accreditation)
			if err != nil {
				cmd.Printf("Failed onboarding '%s', %s\n", i.Address.String(), err)
				continue
			}

			nonce.Add(nonce, one)
			transactions = append(transactions, tx.Hash())
			cli.PrintTransaction(cmd.OutOrStdout(), tx)
		}

		// Wait on the transaction
		if len(transactions) == 1 {
			err = cli.WaitOnTransaction(cmd, transactions[0], 0)
		} else {
			err = cli.WaitOnTransactions(cmd, transactions, 0)
		}
		cli.CheckErr(cmd, err)
	},
}

func tokenTransactorSessionFn(addr common.Address, transactor bind.ContractTransactor) (interface{}, error) {
	return erc20.NewT0kenTransactor(addr, transactor)
}

var issueTokensCommand = &cobra.Command{
	Use:   "issuance <file>",
	Short: "Issues tokens to investors from the given CSV <file>",
	Long: `The CSV file must be in the following format:
---
address,tokens
0xf01ff29dcbee147e9ca151a281bfdf136f66a45b,66
---

 - The header column must be present, with all values being required.
`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		// Open file
		f, err := os.Open(args[0])
		cli.CheckErr(cmd, err)

		// Connect to node
		o, transactOpts := commands.ConnectWithTransactorSessionFunc(cmd, args, "t0ken", tokenTransactorSessionFn)
		transactor := o.(*erc20.T0kenTransactor)
		transSession := &erc20.T0kenTransactorSession{transactor, transactOpts}
		nonce := cli.Conn.Opts.Nonce

		// Onboard
		var transactions []common.Hash
		for msg := range parseInvestorCSV(f, true) {
			if msg.Error != nil {
				cmd.Println(msg.Error)
				break
			}

			i := msg.Investor
			tx, err := transSession.Transfer(i.Address, i.Tokens)
			if err != nil {
				cmd.Printf("Failed onboarding '%s', %s\n", i.Address.String(), err)
				continue
			}

			nonce.Add(nonce, one)
			transactions = append(transactions, tx.Hash())
			cli.PrintTransaction(cmd.OutOrStdout(), tx)
		}

		// Wait on the transaction
		if len(transactions) == 1 {
			err = cli.WaitOnTransaction(cmd, transactions[0], 0)
		} else {
			err = cli.WaitOnTransactions(cmd, transactions, 0)
		}
		cli.CheckErr(cmd, err)
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
	rootCmd.AddCommand(commands.Version)
	rootCmd.AddCommand(completionCmd)

	onboardInvestorCommand.Flags().String("address", "", `address of the BrokerDealer registry contract (default "[investorRegistry] value from config")`)
	onboardInvestorCommand.Flags().Int("wait", -1, "waits the provided number of seconds for the transaction to be mined ('0' waits indefinitely)")
	issueTokensCommand.Flags().String("address", "", `address of the BrokerDealer registry contract (default "[investorRegistry] value from config")`)
	issueTokensCommand.Flags().Int("wait", -1, "waits the provided number of seconds for the transaction to be mined ('0' waits indefinitely)")

	rootCmd.AddCommand(onboardInvestorCommand)
	rootCmd.AddCommand(issueTokensCommand)

	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
