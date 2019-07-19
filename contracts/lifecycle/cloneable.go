// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package lifecycle

import (
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// CloneableABI is the input ABI used to generate the binding from.
const CloneableABI = "[{\"constant\":false,\"inputs\":[],\"name\":\"clone\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"}]"

// CloneableBin is the compiled bytecode used for deploying new contracts.
const CloneableBin = `608060405234801561001057600080fd5b5060c58061001f6000396000f3fe608060405260043610601c5760003560e01c806309ed4607146021575b600080fd5b60276050565b6040805173ffffffffffffffffffffffffffffffffffffffff9092168252519081900360200190f35b6000803060481b7f5880730000000000000000000000000000000000000000803b80938091923cf3176000526020600034f09050803b608b57fe5b90509056fea265627a7a72305820d5703d51057e73c2174ee42142589d410dc4191a1a0df61da1915ca4bb99008c64736f6c63430005090032`

// DeployCloneable deploys a new Ethereum contract, binding an instance of Cloneable to it.
func DeployCloneable(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Cloneable, error) {
	parsed, err := abi.JSON(strings.NewReader(CloneableABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(CloneableBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Cloneable{CloneableCaller: CloneableCaller{contract: contract}, CloneableTransactor: CloneableTransactor{contract: contract}, CloneableFilterer: CloneableFilterer{contract: contract}}, nil
}

// Cloneable is an auto generated Go binding around an Ethereum contract.
type Cloneable struct {
	CloneableCaller     // Read-only binding to the contract
	CloneableTransactor // Write-only binding to the contract
	CloneableFilterer   // Log filterer for contract events
}

// CloneableCaller is an auto generated read-only Go binding around an Ethereum contract.
type CloneableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CloneableTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CloneableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CloneableFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CloneableFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CloneableSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CloneableSession struct {
	Contract     *Cloneable        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CloneableCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CloneableCallerSession struct {
	Contract *CloneableCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// CloneableTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CloneableTransactorSession struct {
	Contract     *CloneableTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// CloneableRaw is an auto generated low-level Go binding around an Ethereum contract.
type CloneableRaw struct {
	Contract *Cloneable // Generic contract binding to access the raw methods on
}

// CloneableCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CloneableCallerRaw struct {
	Contract *CloneableCaller // Generic read-only contract binding to access the raw methods on
}

// CloneableTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CloneableTransactorRaw struct {
	Contract *CloneableTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCloneable creates a new instance of Cloneable, bound to a specific deployed contract.
func NewCloneable(address common.Address, backend bind.ContractBackend) (*Cloneable, error) {
	contract, err := bindCloneable(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Cloneable{CloneableCaller: CloneableCaller{contract: contract}, CloneableTransactor: CloneableTransactor{contract: contract}, CloneableFilterer: CloneableFilterer{contract: contract}}, nil
}

// NewCloneableCaller creates a new read-only instance of Cloneable, bound to a specific deployed contract.
func NewCloneableCaller(address common.Address, caller bind.ContractCaller) (*CloneableCaller, error) {
	contract, err := bindCloneable(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CloneableCaller{contract: contract}, nil
}

// NewCloneableTransactor creates a new write-only instance of Cloneable, bound to a specific deployed contract.
func NewCloneableTransactor(address common.Address, transactor bind.ContractTransactor) (*CloneableTransactor, error) {
	contract, err := bindCloneable(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CloneableTransactor{contract: contract}, nil
}

// NewCloneableFilterer creates a new log filterer instance of Cloneable, bound to a specific deployed contract.
func NewCloneableFilterer(address common.Address, filterer bind.ContractFilterer) (*CloneableFilterer, error) {
	contract, err := bindCloneable(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CloneableFilterer{contract: contract}, nil
}

// bindCloneable binds a generic wrapper to an already deployed contract.
func bindCloneable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(CloneableABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Cloneable *CloneableRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Cloneable.Contract.CloneableCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Cloneable *CloneableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Cloneable.Contract.CloneableTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Cloneable *CloneableRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Cloneable.Contract.CloneableTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Cloneable *CloneableCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Cloneable.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Cloneable *CloneableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Cloneable.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Cloneable *CloneableTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Cloneable.Contract.contract.Transact(opts, method, params...)
}

// Clone is a paid mutator transaction binding the contract method 0x09ed4607.
//
// Solidity: function clone() returns(address)
func (_Cloneable *CloneableTransactor) Clone(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Cloneable.contract.Transact(opts, "clone")
}

// Clone is a paid mutator transaction binding the contract method 0x09ed4607.
//
// Solidity: function clone() returns(address)
func (_Cloneable *CloneableSession) Clone() (*types.Transaction, error) {
	return _Cloneable.Contract.Clone(&_Cloneable.TransactOpts)
}

// Clone is a paid mutator transaction binding the contract method 0x09ed4607.
//
// Solidity: function clone() returns(address)
func (_Cloneable *CloneableTransactorSession) Clone() (*types.Transaction, error) {
	return _Cloneable.Contract.Clone(&_Cloneable.TransactOpts)
}
