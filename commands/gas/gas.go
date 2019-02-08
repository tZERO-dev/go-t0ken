package gas

import (
	"context"
	"math/big"

	"github.com/spf13/cobra"

	"github.com/tzero-dev/go-t0ken/cli"
	"github.com/tzero-dev/go-t0ken/units"
)

var gasPrice float64

var Command = &cobra.Command{
	Use:   "gas",
	Short: "Gas utilities",
}

var PriceCommand = &cobra.Command{
	Use:    "price",
	Short:  "Gets the current gas price in Gwei",
	PreRun: cli.Connect,
	Args:   cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		price, err := cli.Conn.SuggestGasPrice(context.Background())
		cli.CheckErr(cmd, err)
		gwei := units.ConvertInt(price, units.Wei, units.Gwei)
		cmd.Printf("%s gwei\n", gwei.Text('f', 2))
	},
}

// GasPrice Returns the gas price flag value, or the suggested gas price when unset.
func GetPrice(cmd *cobra.Command, args []string) (*big.Int, error) {
	f := cmd.Flag("gasPrice")
	if f == nil {
		return cli.Conn.SuggestGasPrice(context.Background())
	}
	g := big.NewFloat(gasPrice)
	wei := new(big.Int)
	units.ConvertFloat(g, units.Gwei, units.Wei).Int(wei)
	return wei, nil
}

// Flag adds the 'gasPrice' flag to the given command.
func Flag(cmd *cobra.Command) {
	cmd.Flags().Float64Var(&gasPrice, "gasPrice", 0, "manually set the gas price for the transaction (gwei)")
}
