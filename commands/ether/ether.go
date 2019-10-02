package ether

import (
	"context"
	"errors"
	"math/big"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/tzero-dev/go-t0ken/cli"
	"github.com/tzero-dev/go-t0ken/commands/gas"
	"github.com/tzero-dev/go-t0ken/commands/nonce"
	"github.com/tzero-dev/go-t0ken/units"
)

var Command = &cobra.Command{
	Use:   "ether",
	Short: "Ether utilities",
}

var SendCommand = &cobra.Command{
	Use:    "send <address> <ether> [message]",
	Short:  "sends an <address> <ether> [message]",
	PreRun: cli.ConnectWithKeyStore,
	Args:   cli.ChainArgs(cobra.RangeArgs(2, 3), cli.AddressArgFunc("address", 0), cli.BigFloatArgFunc("ether", 1)),
	Run: func(cmd *cobra.Command, args []string) {
		to := common.HexToAddress(args[0])
		wei := getAmount(args[1])
		gasPrice, err := gas.GetPrice(cmd, args)
		cli.CheckErr(cmd, err)

		var data []byte
		if len(args) == 3 {
			data = []byte(args[2])
		}

		// Get the next, or supplied nonce
		n, err := nonce.Get()
		cli.CheckErr(cmd, err)

		// Estimate the gas (2100 for non-message transaction)
		gas, err := cli.Conn.EstimateGas(context.Background(), ethereum.CallMsg{
			From:     cli.Conn.Opts.From,
			To:       &to,
			Data:     data,
			Value:    wei,
			GasPrice: gasPrice,
		})
		cli.CheckErr(cmd, err)

		// Create and sign the transaction
		tx := types.NewTransaction(n, to, wei, gas, gasPrice, data)
		signedTx, err := cli.Conn.Opts.Signer(types.HomesteadSigner{}, cli.Conn.Opts.From, tx)
		cli.CheckErr(cmd, err)

		// Send it
		err = cli.Conn.SendTransaction(context.Background(), signedTx)
		cli.PrintTransactionFn(cmd)(signedTx, err)
	},
}

var BalanceCommand = &cobra.Command{
	Use:    "balance [address]",
	Short:  "gets the balance of the [address]",
	PreRun: cli.Connect,
	Args:   cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		var addr common.Address
		var err error

		if len(args) == 0 {
			s := viper.GetString("keystoreAddress")
			addr = common.HexToAddress(s)
			if !common.IsHexAddress(s) {
				err = errors.New("No [address] provided, and no valid 'keystoreAddress' address found")
			}
		} else if common.IsHexAddress(args[0]) {
			addr, err = cli.GetArgAddress(0, args)
		} else {
			addr, _, _, err = cli.AddressForKeystoreAlias(args[0])
		}
		cli.CheckErr(cmd, err)

		// Get the block flag, if set
		var block *big.Int
		i, err := cmd.Flags().GetInt64("block")
		cli.CheckErr(cmd, err)
		if i > -1 {
			block = big.NewInt(i)
		}
		balance, err := cli.Conn.BalanceAt(context.Background(), addr, block)
		cli.CheckErr(cmd, err)

		ether := units.ConvertInt(balance, units.Wei, units.Ether)
		cmd.Printf("Ξ %s\n", ether.Text('f', 4))
		usd, btc, sat, err := getExchangePrice(ether)
		if err == nil {
			cmd.Printf("\n(%s)\n", exchange)
			cmd.Printf("₿ %s (%s SAT)\n", btc.Text('f', 4), sat.Text('f', 0))
			cmd.Printf("$ %s\n", usd.Text('f', 4))
		}
	},
}

// getAmount returns the given Ether amount in Wei.
func getAmount(ether string) *big.Int {
	amount, _ := new(big.Float).SetString(ether)
	wei := new(big.Int)
	units.ConvertFloat(amount, units.Ether, units.Wei).Int(wei)
	return wei
}

func init() {
	// Add the 'gasPrice', 'nonce' and 'wait' flags to the deploy function
	gas.Flag(SendCommand)
	nonce.Flag(SendCommand)
	cli.WaitFlag(SendCommand)
	cli.BlockFlag(BalanceCommand)
}
