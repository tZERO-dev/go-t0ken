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

// AddressMapABI is the input ABI used to generate the binding from.
const AddressMapABI = "[]"

// AddressMapBin is the compiled bytecode used for deploying new contracts.
const AddressMapBin = `60556023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea265627a7a72305820f639c956e6dea8c8377d510b36be6be530a7064fc63e9b0c66ef8535677c187b64736f6c634300050a0032`

// DeployAddressMap deploys a new Ethereum contract, binding an instance of AddressMap to it.
func DeployAddressMap(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *AddressMap, error) {
	parsed, err := abi.JSON(strings.NewReader(AddressMapABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(AddressMapBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &AddressMap{AddressMapCaller: AddressMapCaller{contract: contract}, AddressMapTransactor: AddressMapTransactor{contract: contract}, AddressMapFilterer: AddressMapFilterer{contract: contract}}, nil
}

// AddressMap is an auto generated Go binding around an Ethereum contract.
type AddressMap struct {
	AddressMapCaller     // Read-only binding to the contract
	AddressMapTransactor // Write-only binding to the contract
	AddressMapFilterer   // Log filterer for contract events
}

// AddressMapCaller is an auto generated read-only Go binding around an Ethereum contract.
type AddressMapCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AddressMapTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AddressMapTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AddressMapFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AddressMapFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AddressMapSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AddressMapSession struct {
	Contract     *AddressMap       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AddressMapCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AddressMapCallerSession struct {
	Contract *AddressMapCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// AddressMapTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AddressMapTransactorSession struct {
	Contract     *AddressMapTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// AddressMapRaw is an auto generated low-level Go binding around an Ethereum contract.
type AddressMapRaw struct {
	Contract *AddressMap // Generic contract binding to access the raw methods on
}

// AddressMapCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AddressMapCallerRaw struct {
	Contract *AddressMapCaller // Generic read-only contract binding to access the raw methods on
}

// AddressMapTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AddressMapTransactorRaw struct {
	Contract *AddressMapTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAddressMap creates a new instance of AddressMap, bound to a specific deployed contract.
func NewAddressMap(address common.Address, backend bind.ContractBackend) (*AddressMap, error) {
	contract, err := bindAddressMap(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AddressMap{AddressMapCaller: AddressMapCaller{contract: contract}, AddressMapTransactor: AddressMapTransactor{contract: contract}, AddressMapFilterer: AddressMapFilterer{contract: contract}}, nil
}

// NewAddressMapCaller creates a new read-only instance of AddressMap, bound to a specific deployed contract.
func NewAddressMapCaller(address common.Address, caller bind.ContractCaller) (*AddressMapCaller, error) {
	contract, err := bindAddressMap(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AddressMapCaller{contract: contract}, nil
}

// NewAddressMapTransactor creates a new write-only instance of AddressMap, bound to a specific deployed contract.
func NewAddressMapTransactor(address common.Address, transactor bind.ContractTransactor) (*AddressMapTransactor, error) {
	contract, err := bindAddressMap(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AddressMapTransactor{contract: contract}, nil
}

// NewAddressMapFilterer creates a new log filterer instance of AddressMap, bound to a specific deployed contract.
func NewAddressMapFilterer(address common.Address, filterer bind.ContractFilterer) (*AddressMapFilterer, error) {
	contract, err := bindAddressMap(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AddressMapFilterer{contract: contract}, nil
}

// bindAddressMap binds a generic wrapper to an already deployed contract.
func bindAddressMap(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AddressMapABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AddressMap *AddressMapRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _AddressMap.Contract.AddressMapCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AddressMap *AddressMapRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AddressMap.Contract.AddressMapTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AddressMap *AddressMapRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AddressMap.Contract.AddressMapTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AddressMap *AddressMapCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _AddressMap.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AddressMap *AddressMapTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AddressMap.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AddressMap *AddressMapTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AddressMap.Contract.contract.Transact(opts, method, params...)
}
