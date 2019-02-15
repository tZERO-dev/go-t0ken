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
	used := tx.Gas()
	limit := Conn.Opts.GasLimit
	limitStr := string(limit)
	percent := uint64(0)
	if limit > 0 {
		percent = used / limit
	} else {
		limitStr = "~"
	}

	cmd.Printf(`
Transaction: %s
       Cost: %s Ether
  Gas Limit: %s
   Gas Used: %d (%d%%)
  Gas Price: %s Gwei
      Nonce: %d
`,
		tx.Hash().String(),
		cost.Text('f', 9),
		limitStr,
		used, percent,
		price.Text('f', 4),
		tx.Nonce())
}
