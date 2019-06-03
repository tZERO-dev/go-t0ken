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
	//// Ensure the transaction is successful
	//tr, err := Conn.TransactionReceipt(context.Background(), hash)
	//if err != nil {
	//	return nil, err
	//} else if tr.Status != 1 {
	//	return nil, fmt.Errorf("transaction failed with status:", tr.Status)
	//}
	//
	//// Only attempting to get logs so we can natively get to the block number, since we can't with current Geth 1.8.
	//// The issue is that not every transaction includes logs, so this is unreliable.
	//// We can use a filter query, but without knowing the block range this could be intensive.
	//if len(tr.Logs) == 0 {
	//	return nil, errors.New("Receipt contains no logs to get block, awaiting on go-ethereum PR-17662 to remedy this")
	//}
	//h := tr.Logs[0].BlockHash
	//return Conn.BlockByHash(context.Background(), h)

	// https://github.com/ethereum/go-ethereum/pull/17662#pullrequestreview-206299201
	// We should be able to use the above PR, which will hopefully be included in the 1.9.x release to get the block number.
	// But for now, issuing a raw query and parsing the bits we need will suffice.
	var r *transactionReceiptBlock
	err := Conn.RawCall(&r, "eth_getTransactionReceipt", hash)
	if err != nil {
		return nil, err
	}
	return Conn.BlockByHash(context.Background(), r.BlockHash)
}

//func WaitOnTransaction(cmd *cobra.Command, tx common.Hash, confirmations int) error {
//	// Wait for the transaction to be mined
//	for {
//		_, pending, err := Conn.TransactionByHash(context.Background(), tx)
//		if err != nil {
//			return err
//		} else if !pending {
//			break
//		}
//
//		<-time.After(waitDuration)
//		cmd.Print(".")
//	}
//
//	// Ensure the transaction is successful
//	tr, err := Conn.TransactionReceipt(context.Background(), tx)
//	if err != nil {
//		return err
//	}
//
//	// Output the status and contract address
//	cmd.Printf("       Status: %d\n", tr.Status)
//	if bytes.Compare(tr.ContractAddress.Bytes(), zeroAddressBytes) != 0 {
//		cmd.Printf("     Contract: %s\n", tr.ContractAddress.String())
//	}
//
//	c := big.NewInt(int64(confirmations))
//	total := int64(0)
//	for {
//		// Get latest block
//		latest, err := Conn.HeaderByNumber(context.Background(), nil)
//		if err != nil {
//			return err
//		}
//
//		// Output block info from receipt
//		block, err := receiptBlock(tx)
//		if err == nil {
//			err = PrintBlock(cmd.OutOrStdout(), latest, block)
//		}
//		if err != nil {
//			return err
//		}
//
//		// Wait for the block to have the number of given confirmations
//		n := latest.Number.Sub(latest.Number, block.Number())
//		if n.Int64() != total {
//			total = n.Int64()
//			cmd.Printf("\n%d confirmations", total)
//		}
//		if n.Cmp(c) >= 0 {
//			break
//		}
//		cmd.Print(".")
//		<-time.After(waitDuration)
//	}
//	return nil
//}

func WaitOnTransaction(cmd *cobra.Command, tx common.Hash, timeout int) error {
	// Wait for the transaction to be mined
	elapsed := 0
	for {
		_, pending, err := Conn.TransactionByHash(context.Background(), tx)
		if err != nil {
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
		cmd.Print(".")
	}

	// Ensure the transaction is successful
	tr, err := Conn.TransactionReceipt(context.Background(), tx)
	if err != nil {
		return err
	}

	// Output the status and contract address
	cmd.Printf("       Status: %d\n", tr.Status)
	if bytes.Compare(tr.ContractAddress.Bytes(), zeroAddressBytes) != 0 {
		cmd.Printf("     Contract: %s\n", tr.ContractAddress.String())
	}

	// Get latest block
	latest, err := Conn.HeaderByNumber(context.Background(), nil)
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
	confirmedTransactions := make([]common.Hash, len(transactions))

	// Wait for all transaction confirmations
	for i := 0; i < len(transactions); {
		elapsed := 0
		_, pending, err := Conn.TransactionByHash(context.Background(), transactions[i])
		if err != nil {
			cmd.Printf("failed to retrieve transaction '%s', %s\n", transactions[i].String(), err)
			i++
		} else if !pending {
			confirmedTransactions = append(confirmedTransactions, transactions[i])
			i++
		} else if timeout > 0 && elapsed-timeout > 0 {
			return errors.New("exceederd timeout")
		}
		<-time.After(waitDuration)
		elapsed += 1
		cmd.Print(".")
	}

	// Get the blocks the transactions are included in
	var blocks []*types.Block
	blockTx := make(map[uint64][]string)
	for _, tx := range transactions {
		// Ensure the transaction is successful
		tr, err := Conn.TransactionReceipt(context.Background(), tx)
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
	latest, err := Conn.HeaderByNumber(context.Background(), nil)
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
	tx, pending, err := Conn.TransactionByHash(context.Background(), h)
	if err != nil {
		return err
	} else if pending {
		return errors.New("transaction must be in 'pending' state")
	}

	tx = types.NewTransaction(tx.Nonce(), *tx.To(), tx.Value(), Conn.Opts.GasLimit, gasPrice, tx.Data())
	return Conn.SendTransaction(context.Background(), tx)
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
