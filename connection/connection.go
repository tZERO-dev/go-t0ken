package connection

import (
	"context"
	"errors"
	"fmt"
	"io/ioutil"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

var (
	accountKey      string
	accountPassword *string
	clientURL       *string
)

// Connection wraps 'etherclient.Client' to hold transactor options and a key-store.
type Connection struct {
	*ethclient.Client
	Opts     *bind.TransactOpts
	KeyStore *keystore.KeyStore
}

// accountFor returns the matching account within the key-store.
func accountFor(address common.Address, ks *keystore.KeyStore) (accounts.Account, error) {
	for _, a := range ks.Accounts() {
		if a.Address == address {
			return a, nil
		}
	}
	return accounts.Account{}, fmt.Errorf("Address doesn't exist within keystore")
}

// NewTransactor returns a new transactor for the given keyFile.
func NewTransactor(keyFile string, password string) (*bind.TransactOpts, error) {
	b, err := ioutil.ReadFile(keyFile)
	if err != nil {
		return nil, err
	}
	key := strings.NewReader(string(b))
	auth, err := bind.NewTransactor(key, password)
	return auth, err
}

// SetNextNonce sets the next nonce of the transactor address.
func SetNextNonce(client *ethclient.Client, opts *bind.TransactOpts) error {
	nonce, err := client.PendingNonceAt(context.Background(), opts.From)
	if err == nil {
		opts.Nonce = new(big.Int).SetUint64(nonce)
	}
	return err
}

// SetGasPrice
func (c *Connection) SetGasPrice(wei *big.Int) {
	c.Opts.GasPrice = wei
}

// SetSuggestedGasPrice sets the suggested gas price of the connection.
func (c *Connection) SetSuggestedGasPrice() error {
	wei, err := c.SuggestGasPrice(context.Background())
	if err != nil {
		c.Opts.GasPrice = wei
	}
	return err
}

// SetNonce sets the nonce to the provided value. This can be used to replay a stuck transaction.
func (c *Connection) SetNonce(nonce uint64) {
	c.Opts.Nonce = new(big.Int).SetUint64(nonce)
}

// SetNextNonce sets nonce to the next available for the transactor's address.
func (c *Connection) SetNextNonce() error {
	return SetNextNonce(c.Client, c.Opts)
}

// SetKeyStore sets the connections keystore to the path provided.
func (c *Connection) SetKeyStore(keystorePath string) error {
	c.KeyStore = keystore.NewKeyStore(keystorePath, keystore.LightScryptN, keystore.LightScryptP)
	return nil
}

// SetTransactor sets the connections transactor to that of the given keyFile.
func (c *Connection) SetTransactor(keyPath, password string) error {
	o, err := NewTransactor(keyPath, password)
	if err == nil {
		c.Opts = o
	}
	return err
}

// SetTransactorFromKeyStore sets the connections transactor to the address within the key-tore.
func (c *Connection) SetTransactorFromKeyStore(a common.Address, password string) error {
	if c.KeyStore == nil {
		return errors.New("KeyStore must be set to set the transactor")
	}
	key, err := accountFor(a, c.KeyStore)
	if err != nil {
		return err
	}
	return c.SetTransactor(key.URL.Path, password)
}

// New creates a new connection
func New(url string) (*Connection, error) {
	c, err := ethclient.Dial(url)
	if err != nil {
		return nil, err
	}
	return &Connection{c, nil, nil}, err
}

// NewForKey creates a new connection using the given key-store file and password.
func NewForKey(url, keyFile, password string) (*Connection, error) {
	c, err := New(url)
	if err != nil {
		return c, err
	}
	c.Opts, err = NewTransactor(keyFile, password)
	return c, err
}

// NewForKeyStore creates a new connection using the given key-store and password.
func NewForKeyStore(url, keystorePath string, a common.Address, password string) (*Connection, error) {
	c, err := New(url)
	if err != nil {
		return c, err
	}
	c.KeyStore = keystore.NewKeyStore(keystorePath, keystore.LightScryptN, keystore.LightScryptP)
	err = c.SetTransactorFromKeyStore(a, password)
	return c, err
}
