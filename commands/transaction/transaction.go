package transaction

import (
	"context"
	"errors"
	"fmt"
	"math/big"
	"strconv"

	//"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/spf13/cobra"

	"github.com/tzero-dev/go-t0ken/cli"
	"github.com/tzero-dev/go-t0ken/commands/gas"
	"github.com/tzero-dev/go-t0ken/commands/nonce"
	"github.com/tzero-dev/go-t0ken/units"
)

var Command = &cobra.Command{
	Use:   "transaction",
	Short: "Transaction utilities",
}

var ReplayCommand = &cobra.Command{
	Use:    "replay <transaction>...",
	Short:  "Replays the pending <transaction>... by doubling the gas price (can override this with the 'gasPrice' flag)",
	PreRun: cli.ConnectWithKeyStore,
	Run: func(cmd *cobra.Command, args []string) {
		hashes, err := getHashes(args)
		cli.CheckErr(cmd, err)
		for i := range hashes {
			// Get the transaction and new gas-price
			tx, gasPrice, err := transAndGas(cmd, hashes[i])
			cli.CheckErr(cmd, err)

			// Get the latest block, for gasLimit
			b, err := cli.Conn.BlockByNumber(context.Background(), nil)
			cli.CheckErr(cmd, err)

			// Create the replay transaction or contract creation
			if tx.To() != nil {
				tx = types.NewTransaction(tx.Nonce(), *tx.To(), tx.Value(), b.GasLimit(), gasPrice, tx.Data())
			} else {
				// NOTE: Replaying contract creation transactions doesn't currently work, as Geth always returns the error: 'known transaction: <tx>'.
				//       To work around this, cancel the transaction first, then resubmit the contract creation.
				types.NewContractCreation(tx.Nonce(), tx.Value(), b.GasLimit(), gasPrice, tx.Data())
			}

			// Sign & replay the transaction with the new gas-price
			signedTx, err := cli.Conn.Opts.Signer(types.HomesteadSigner{}, cli.Conn.Opts.From, tx)
			if err == nil {
				err = cli.Conn.SendTransaction(context.Background(), signedTx)
			}
			cli.CheckErr(cmd, err)
			cli.PrintTransactionFn(cmd)(signedTx, err)
		}
	},
}

var CancelCommand = &cobra.Command{
	Use:    "cancel <transaction>...",
	Short:  "Cancels the pending <transaction>...",
	PreRun: cli.ConnectWithKeyStore,
	Run: func(cmd *cobra.Command, args []string) {
		hashes, err := getHashes(args)
		cli.CheckErr(cmd, err)
		for i := range hashes {
			// Get the transaction and new gas-price
			tx, gasPrice, err := transAndGas(cmd, hashes[i])
			cli.CheckErr(cmd, err)

			// Send 0 eth with the new gas value to try an cancel the pending transaction with matching nonce.
			to := cli.Conn.Opts.From
			tx = types.NewTransaction(tx.Nonce(), to, big.NewInt(0), uint64(21000), gasPrice, nil)
			signedTx, err := cli.Conn.Opts.Signer(types.HomesteadSigner{}, cli.Conn.Opts.From, tx)
			cli.CheckErr(cmd, err)
			err = cli.Conn.SendTransaction(context.Background(), signedTx)
			cli.PrintTransactionFn(cmd)(signedTx, err)
		}
	},
}

var WaitCommand = &cobra.Command{
	Use:    "wait <transaction>...",
	Short:  "Waits for the <transaction>... to no longer be in a pending state",
	PreRun: cli.Connect,
	Run: func(cmd *cobra.Command, args []string) {
		// Get wait flag/value
		wait := 0
		f := cmd.Flag("wait")
		if f != nil {
			w, _ := strconv.Atoi(f.Value.String())
			if w > -1 {
				wait = w
			}
		}

		// Get hash args and wait on transactions
		hashes, err := getHashes(args)
		cli.CheckErr(cmd, err)

		// Print transaction info
		for i := range hashes {
			tx, _, err := cli.Conn.TransactionByHash(context.Background(), hashes[i])
			cli.CheckErr(cmd, err)
			cli.PrintTransaction(cmd.OutOrStdout(), tx)
		}

		// Wait on the transactions, and print block info
		if len(hashes) == 1 {
			err = cli.WaitOnTransaction(cmd, hashes[0], wait)
		} else {
			err = cli.WaitOnTransactions(cmd, hashes, wait)
		}
		cli.CheckErr(cmd, err)
	},
}

var NextAddressCommand = &cobra.Command{
	Use:    "nextAddress [flags]",
	Short:  "Outputs the next contract address, supply the 'nonce' flag when you want a future address",
	PreRun: cli.ConnectWithKeyStore,
	Run: func(cmd *cobra.Command, args []string) {
		n, err := nonce.Get()
		cli.CheckErr(cmd, err)
		addr := crypto.CreateAddress(cli.Conn.Opts.From, n)
		cmd.Println(addr.String())
	},
}

func getHashes(args []string) ([]common.Hash, error) {
	hashes := make([]common.Hash, len(args), len(args))
	for i := range args {
		h := common.HexToHash(args[i])
		if len(h.Bytes()) != 32 {
			return hashes, fmt.Errorf("invalid hash length of %d", len(h))
		}
		hashes[i] = h
	}
	return hashes, nil
}

func transAndGas(cmd *cobra.Command, h common.Hash) (*types.Transaction, *big.Int, error) {
	// Get the transaction for the hash, ensuring it's valid and not in a pending state
	tx, pending, err := cli.Conn.TransactionByHash(context.Background(), h)
	if err != nil {
		return tx, nil, err
	} else if !pending {
		return tx, nil, errors.New("transaction must be in 'pending' state")
	}

	// Get gas price (use the value from 'gasPrice' flag or add 50% to the existing transaction value)
	gasPrice := gas.GetFlagValue(cmd)
	if gasPrice == nil {
		// Get the original transaction's gas price
		gasPrice = tx.GasPrice()

		// Convert from Wei to Gwei and add 20% (must increase by more than 10% for the replacement to be accepted)
		g := units.ConvertInt(gasPrice, units.Wei, units.Gwei)
		g.Mul(g, big.NewFloat(1.2))

		// Convert Gwei + 20% back to Wei
		units.ConvertFloat(g, units.Gwei, units.Wei).Int(gasPrice)
	}

	return tx, gasPrice, nil
}

func init() {
	gas.Flag(ReplayCommand)
	gas.Flag(CancelCommand)

	nonce.Flag(NextAddressCommand)
	WaitCommand.Flags().Int("wait", -1, "waits the provided number of seconds for the transaction to be mined ('0' waits indefinitely)")
}
