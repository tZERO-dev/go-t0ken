package ether

import (
	"context"
	"errors"
	"math/big"

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
	Use:    "send <address> <ether>",
	Short:  "sends an <address> <ether>",
	PreRun: cli.ConnectWithKeyStore,
	Args:   cli.ChainArgs(cobra.ExactArgs(2), cli.AddressArgFunc("address", 0), cli.BigFloatArgFunc("ether", 1)),
	Run: func(cmd *cobra.Command, args []string) {
		to := common.HexToAddress(args[0])
		wei := getAmount(args[1])
		gasPrice, err := gas.GetPrice(cmd, args)
		cli.CheckErr(cmd, err)

		n, err := nonce.Get()
		cli.CheckErr(cmd, err)

		tx := types.NewTransaction(n, to, wei, uint64(21000), gasPrice, nil)
		signedTx, err := cli.Conn.Opts.Signer(types.HomesteadSigner{}, cli.Conn.Opts.From, tx)
		cli.CheckErr(cmd, err)

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

		balance, err := cli.Conn.BalanceAt(context.Background(), addr, nil)
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
}
