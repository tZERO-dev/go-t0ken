// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package rules

import (
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// RuleABI is the input ABI used to generate the binding from.
const RuleABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"initiator\",\"type\":\"address\"},{\"name\":\"from\",\"type\":\"address\"},{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"toKind\",\"type\":\"uint8\"},{\"name\":\"tokens\",\"type\":\"uint256\"},{\"name\":\"store\",\"type\":\"address\"}],\"name\":\"check\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// Rule is an auto generated Go binding around an Ethereum contract.
type Rule struct {
	RuleCaller     // Read-only binding to the contract
	RuleTransactor // Write-only binding to the contract
	RuleFilterer   // Log filterer for contract events
}

// RuleCaller is an auto generated read-only Go binding around an Ethereum contract.
type RuleCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RuleTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RuleTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RuleFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RuleFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RuleSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RuleSession struct {
	Contract     *Rule             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RuleCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RuleCallerSession struct {
	Contract *RuleCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// RuleTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RuleTransactorSession struct {
	Contract     *RuleTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RuleRaw is an auto generated low-level Go binding around an Ethereum contract.
type RuleRaw struct {
	Contract *Rule // Generic contract binding to access the raw methods on
}

// RuleCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RuleCallerRaw struct {
	Contract *RuleCaller // Generic read-only contract binding to access the raw methods on
}

// RuleTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RuleTransactorRaw struct {
	Contract *RuleTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRule creates a new instance of Rule, bound to a specific deployed contract.
func NewRule(address common.Address, backend bind.ContractBackend) (*Rule, error) {
	contract, err := bindRule(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Rule{RuleCaller: RuleCaller{contract: contract}, RuleTransactor: RuleTransactor{contract: contract}, RuleFilterer: RuleFilterer{contract: contract}}, nil
}

// NewRuleCaller creates a new read-only instance of Rule, bound to a specific deployed contract.
func NewRuleCaller(address common.Address, caller bind.ContractCaller) (*RuleCaller, error) {
	contract, err := bindRule(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RuleCaller{contract: contract}, nil
}

// NewRuleTransactor creates a new write-only instance of Rule, bound to a specific deployed contract.
func NewRuleTransactor(address common.Address, transactor bind.ContractTransactor) (*RuleTransactor, error) {
	contract, err := bindRule(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RuleTransactor{contract: contract}, nil
}

// NewRuleFilterer creates a new log filterer instance of Rule, bound to a specific deployed contract.
func NewRuleFilterer(address common.Address, filterer bind.ContractFilterer) (*RuleFilterer, error) {
	contract, err := bindRule(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RuleFilterer{contract: contract}, nil
}

// bindRule binds a generic wrapper to an already deployed contract.
func bindRule(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(RuleABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Rule *RuleRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Rule.Contract.RuleCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Rule *RuleRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Rule.Contract.RuleTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Rule *RuleRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Rule.Contract.RuleTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Rule *RuleCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Rule.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Rule *RuleTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Rule.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Rule *RuleTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Rule.Contract.contract.Transact(opts, method, params...)
}

// Check is a paid mutator transaction binding the contract method 0xb762c76d.
//
// Solidity: function check(initiator address, from address, to address, toKind uint8, tokens uint256, store address) returns()
func (_Rule *RuleTransactor) Check(opts *bind.TransactOpts, initiator common.Address, from common.Address, to common.Address, toKind uint8, tokens *big.Int, store common.Address) (*types.Transaction, error) {
	return _Rule.contract.Transact(opts, "check", initiator, from, to, toKind, tokens, store)
}

// Check is a paid mutator transaction binding the contract method 0xb762c76d.
//
// Solidity: function check(initiator address, from address, to address, toKind uint8, tokens uint256, store address) returns()
func (_Rule *RuleSession) Check(initiator common.Address, from common.Address, to common.Address, toKind uint8, tokens *big.Int, store common.Address) (*types.Transaction, error) {
	return _Rule.Contract.Check(&_Rule.TransactOpts, initiator, from, to, toKind, tokens, store)
}

// Check is a paid mutator transaction binding the contract method 0xb762c76d.
//
// Solidity: function check(initiator address, from address, to address, toKind uint8, tokens uint256, store address) returns()
func (_Rule *RuleTransactorSession) Check(initiator common.Address, from common.Address, to common.Address, toKind uint8, tokens *big.Int, store common.Address) (*types.Transaction, error) {
	return _Rule.Contract.Check(&_Rule.TransactOpts, initiator, from, to, toKind, tokens, store)
}
