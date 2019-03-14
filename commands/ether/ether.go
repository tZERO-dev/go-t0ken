package ether

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/spf13/cobra"

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

		if common.IsHexAddress(args[0]) {
			addr, err = cli.GetArgAddress(0, args)
		} else {
			addr, _, _, err = cli.AddressForKeystoreAlias(args[0])
		}
		cli.CheckErr(cmd, err)

		balance, err := cli.Conn.BalanceAt(context.Background(), addr, nil)
		cli.CheckErr(cmd, err)

		ether := units.ConvertInt(balance, units.Wei, units.Ether)
		cmd.Println(ether.Text('f', 4), "Ether")
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
	gas.Flag(SendCommand)
	nonce.Flag(SendCommand)

	SendCommand.Flags().Int("wait", -1, "waits the provided number of seconds for the transaction to be mined ('0' waits indefinitely)")
}
