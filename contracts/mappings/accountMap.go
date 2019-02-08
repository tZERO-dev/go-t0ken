// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package mappings

import (
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// AccountMapABI is the input ABI used to generate the binding from.
const AccountMapABI = "[]"

// AccountMapBin is the compiled bytecode used for deploying new contracts.
const AccountMapBin = `604c602c600b82828239805160001a60731460008114601c57601e565bfe5b5030600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea165627a7a72305820a86d4e2877fc87fd582585d0a28437d88110a3ba05a5700418248214a90586fb0029`

// DeployAccountMap deploys a new Ethereum contract, binding an instance of AccountMap to it.
func DeployAccountMap(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *AccountMap, error) {
	parsed, err := abi.JSON(strings.NewReader(AccountMapABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(AccountMapBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &AccountMap{AccountMapCaller: AccountMapCaller{contract: contract}, AccountMapTransactor: AccountMapTransactor{contract: contract}, AccountMapFilterer: AccountMapFilterer{contract: contract}}, nil
}

// AccountMap is an auto generated Go binding around an Ethereum contract.
type AccountMap struct {
	AccountMapCaller     // Read-only binding to the contract
	AccountMapTransactor // Write-only binding to the contract
	AccountMapFilterer   // Log filterer for contract events
}

// AccountMapCaller is an auto generated read-only Go binding around an Ethereum contract.
type AccountMapCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AccountMapTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AccountMapTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AccountMapFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AccountMapFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AccountMapSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AccountMapSession struct {
	Contract     *AccountMap       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AccountMapCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AccountMapCallerSession struct {
	Contract *AccountMapCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// AccountMapTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AccountMapTransactorSession struct {
	Contract     *AccountMapTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// AccountMapRaw is an auto generated low-level Go binding around an Ethereum contract.
type AccountMapRaw struct {
	Contract *AccountMap // Generic contract binding to access the raw methods on
}

// AccountMapCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AccountMapCallerRaw struct {
	Contract *AccountMapCaller // Generic read-only contract binding to access the raw methods on
}

// AccountMapTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AccountMapTransactorRaw struct {
	Contract *AccountMapTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAccountMap creates a new instance of AccountMap, bound to a specific deployed contract.
func NewAccountMap(address common.Address, backend bind.ContractBackend) (*AccountMap, error) {
	contract, err := bindAccountMap(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AccountMap{AccountMapCaller: AccountMapCaller{contract: contract}, AccountMapTransactor: AccountMapTransactor{contract: contract}, AccountMapFilterer: AccountMapFilterer{contract: contract}}, nil
}

// NewAccountMapCaller creates a new read-only instance of AccountMap, bound to a specific deployed contract.
func NewAccountMapCaller(address common.Address, caller bind.ContractCaller) (*AccountMapCaller, error) {
	contract, err := bindAccountMap(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AccountMapCaller{contract: contract}, nil
}

// NewAccountMapTransactor creates a new write-only instance of AccountMap, bound to a specific deployed contract.
func NewAccountMapTransactor(address common.Address, transactor bind.ContractTransactor) (*AccountMapTransactor, error) {
	contract, err := bindAccountMap(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AccountMapTransactor{contract: contract}, nil
}

// NewAccountMapFilterer creates a new log filterer instance of AccountMap, bound to a specific deployed contract.
func NewAccountMapFilterer(address common.Address, filterer bind.ContractFilterer) (*AccountMapFilterer, error) {
	contract, err := bindAccountMap(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AccountMapFilterer{contract: contract}, nil
}

// bindAccountMap binds a generic wrapper to an already deployed contract.
func bindAccountMap(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AccountMapABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AccountMap *AccountMapRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _AccountMap.Contract.AccountMapCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AccountMap *AccountMapRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AccountMap.Contract.AccountMapTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AccountMap *AccountMapRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AccountMap.Contract.AccountMapTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AccountMap *AccountMapCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _AccountMap.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AccountMap *AccountMapTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AccountMap.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AccountMap *AccountMapTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AccountMap.Contract.contract.Transact(opts, method, params...)
}
