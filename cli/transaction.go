package cli

import (
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/spf13/cobra"

	"github.com/tzero-dev/go-t0ken/units"
)

// PrintTransFn returns a function that checks for an error, printing the transaction when successful.
func PrintTransFn(cmd *cobra.Command) func(*types.Transaction, error) {
	return func(tx *types.Transaction, err error) {
		CheckErr(cmd, err)
		PrintTransaction(cmd, tx)
	}
}

// PrintTransaction prints the transaction info.
func PrintTransaction(cmd *cobra.Command, tx *types.Transaction) {
	cost := units.ConvertInt(tx.Cost(), units.Wei, units.Ether)
	price := units.ConvertInt(tx.GasPrice(), units.Wei, units.Gwei)

	cmd.Printf(`
Transaction: %s
       Cost: %s Ether
   Gas Used: %d
  Gas Price: %s Gwei
      Nonce: %d
`,
		tx.Hash().String(),
		cost.Text('f', 9),
		tx.Gas(),
		price.Text('f', 4),
		tx.Nonce())

	/*cmd.Println(conn.Opts.GasLimit)
	data := t.Data()
	for i := range data {
		cmd.Print(string(data[i]))
	}*/
}
