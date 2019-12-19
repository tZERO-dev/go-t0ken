package cli

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"math/big"
	"strconv"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/spf13/cobra"

	"github.com/tzero-dev/go-t0ken/units"
)

var (
	waitDuration     = time.Duration(5 * time.Second)
	zeroAddressBytes = new(common.Address).Bytes()
)

// PrintBlock prints the block info.
func PrintBlock(w io.Writer, latest *types.Header, block *types.Block) error {
	// Header info
	header := block.Header()
	confirmations := new(big.Int).Sub(latest.Number, header.Number).Uint64()
	used := header.GasUsed
	limit := header.GasLimit
	percent := float64(0)
	if limit > 0 {
		percent = (float64(used) / float64(limit)) * 100
	}
	rlpData := []string{}
	rlp.DecodeBytes(header.Extra, &rlpData)

	// Print the block information
	_, err := fmt.Fprintf(w, `        Block: %s (%s)
Confirmations: %d
    Gas Limit: %d
     Gas Used: %d (%.2f%%)
   Difficulty: %s
         Size: %s
      Receipt: %s
       Parent: %s
    Sha3Uncle: %s
        Miner: %s
   Extra Data: %s
`,
		header.Number, header.Hash().String(),
		confirmations,
		limit,
		used, percent,
		header.Difficulty,
		header.Size().String(),
		header.ReceiptHash.String(),
		header.ParentHash.String(),
		header.UncleHash.String(),
		header.Coinbase.String(),
		strings.Join(rlpData, " "))
	return err
}

// PrintTransaction prints the transaction info.
func PrintTransaction(w io.Writer, tx *types.Transaction) error {
	value := units.ConvertInt(tx.Value(), units.Wei, units.Ether)
	cost := units.ConvertInt(tx.Cost(), units.Wei, units.Ether)
	price := units.ConvertInt(tx.GasPrice(), units.Wei, units.Gwei)
	_, err := fmt.Fprintf(w, `  Transaction: %s
         Data: %x
        Value: %s Ether
      Tx Cost: %s Ether
    Gas Price: %s Gwei
        Nonce: %d
`,
		tx.Hash().String(),
		tx.Data(),
		value.Text('f', 9),
		cost.Text('f', 9),
		price.Text('f', 4),
		tx.Nonce())

	return err
}

type transactionReceiptBlock struct {
	BlockHash   common.Hash    `json:"blockHash"`
	BlockNumber hexutil.Uint64 `json:"blockNumber"`
}

func receiptBlock(hash common.Hash) (*types.Block, error) {
	ctx := context.Background()
	r, err := Conn.TransactionReceipt(ctx, hash)
	if err != nil {
		return nil, err
	}
	return Conn.BlockByHash(ctx, r.BlockHash)
}

func WaitOnTransaction(cmd *cobra.Command, tx common.Hash, timeout int) error {
	// Wait for the transaction to be mined
	ctx := context.Background()
	elapsed := 0
	for {
		_, pending, err := Conn.TransactionByHash(ctx, tx)
		if err != nil && err != ethereum.NotFound {
			return err
		} else if timeout > 0 && elapsed-timeout > 0 {
			return errors.New("Exceederd timeout")
		} else if !pending {
			if elapsed > 0 {
				cmd.Println()
			}
			break
		}

		<-time.After(waitDuration)
		elapsed += 1
		if err == ethereum.NotFound {
			cmd.Print("!")
		} else {
			cmd.Print(".")
		}
	}

	// Ensure the transaction is successful
	tr, err := Conn.TransactionReceipt(ctx, tx)
	if err != nil {
		return err
	}

	// Output the status and contract address
	cmd.Printf("       Status: %d\n", tr.Status)
	if bytes.Compare(tr.ContractAddress.Bytes(), zeroAddressBytes) != 0 {
		cmd.Printf("     Contract: %s\n", tr.ContractAddress.String())
	}

	// Get latest block
	latest, err := Conn.HeaderByNumber(ctx, nil)
	if err != nil {
		return err
	}

	// Output block info from receipt
	block, err := receiptBlock(tx)
	if err == nil {
		err = PrintBlock(cmd.OutOrStdout(), latest, block)
	}
	return err
}

func WaitOnTransactions(cmd *cobra.Command, transactions []common.Hash, timeout int) error {
	// Wait for all transaction confirmations
	ctx := context.Background()
	n := len(transactions)
	confirmed := make([]common.Hash, 0, n)
	elapsed := 0
	for i := 0; len(confirmed) < n; i++ {
		_, pending, err := Conn.TransactionByHash(ctx, transactions[i])
		if err != nil && err != ethereum.NotFound {
			cmd.Printf("failed to retrieve transaction '%s', %s\n", transactions[i].String(), err)
			n--
			transactions[i] = transactions[n]
		} else if !pending {
			confirmed = append(confirmed, transactions[i])
		}

		if i == n {
			elapsed++
			if elapsed > timeout {
				return errors.New("exceederd timeout")
			}
			<-time.After(waitDuration)
			if err == ethereum.NotFound {
				cmd.Print("!")
			} else {
				cmd.Print(".")
			}
			i = 0
		}
	}

	// Get the blocks the transactions are included in
	var blocks []*types.Block
	blockTx := make(map[uint64][]string)
	for _, tx := range confirmed {
		// Ensure the transaction is successful
		tr, err := Conn.TransactionReceipt(ctx, tx)
		if err != nil {
			cmd.Printf("failed retrieving receipt for '%s', %s\n", tx.String(), err)
			continue
		} else if tr.Status != 1 {
			cmd.Printf("transaction '%s' failed with status: %d\n", tx.String(), tr.Status)
			continue
		}
		block, err := receiptBlock(tx)
		if err != nil {
			cmd.Printf("failed to get block for transaction '%s', %s\n", tx.String(), err)
			continue
		}
		n := block.NumberU64()
		blocks = append(blocks, block)
		blockTx[n] = append(blockTx[n], tx.String())
	}

	// Get latest block
	latest, err := Conn.HeaderByNumber(ctx, nil)
	if err != nil {
		return err
	}

	// Print block(s)
	for _, block := range blocks {
		s, ok := blockTx[block.NumberU64()]
		if !ok {
			continue
		}

		cmd.Println("--------------------------------------------------------------------------------")
		t := strings.Join(s, "\n               ")
		cmd.Printf("Block %s\n Transactions: %s\n               ...\n", block.Number().String(), t)
		err := PrintBlock(cmd.OutOrStdout(), latest, block)
		if err != nil {
			return err
		}
		delete(blockTx, block.NumberU64())
	}

	return nil
}

// ReplayTransaction attempts to replay an existing, pending transaction.
func ReplayTransaction(cmd *cobra.Command, h common.Hash, gasPrice *big.Int) error {
	ctx := context.Background()
	tx, pending, err := Conn.TransactionByHash(ctx, h)
	if err != nil {
		return err
	} else if pending {
		return errors.New("transaction must be in 'pending' state")
	}

	tx = types.NewTransaction(tx.Nonce(), *tx.To(), tx.Value(), Conn.Opts.GasLimit, gasPrice, tx.Data())
	return Conn.SendTransaction(ctx, tx)
}

// PrintTransactionFn returns a function that checks for an error, printing the transaction when successful.
func PrintTransactionFn(cmd *cobra.Command) func(*types.Transaction, error) {
	return func(tx *types.Transaction, err error) {
		CheckErr(cmd, err)
		PrintTransaction(cmd.OutOrStdout(), tx)
		f := cmd.Flag("wait")
		if f == nil {
			return
		}
		seconds, err := strconv.Atoi(f.Value.String())
		if err != nil || seconds < 0 {
			return
		}
		err = WaitOnTransaction(cmd, tx.Hash(), seconds)
		CheckErr(cmd, err)
	}
}
