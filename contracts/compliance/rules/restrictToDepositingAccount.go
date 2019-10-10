// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package rules

import (
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// RestrictToDepositingAccountABI is the input ABI used to generate the binding from.
const RestrictToDepositingAccountABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnerTransferred\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[],\"name\":\"ZERO_ADDRESS\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"contractIT0ken\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"initiator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokens\",\"type\":\"uint256\"}],\"name\":\"check\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"kill\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"contractICompliance\",\"name\":\"compliance\",\"type\":\"address\"},{\"internalType\":\"contractIT0ken\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"initiator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokens\",\"type\":\"uint256\"}],\"name\":\"test\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"s\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwner\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// RestrictToDepositingAccountBin is the compiled bytecode used for deploying new contracts.
const RestrictToDepositingAccountBin = `6080604052600180546001600160a01b031990811690915560008054909116331790556109b1806100316000396000f3fe608060405234801561001057600080fd5b506004361061007d5760003560e01c80634fb2e45d1161005b5780634fb2e45d14610155578063538ba4f91461017b578063803fcd431461019f5780638da5cb5b146101e55761007d565b806306fdde03146100825780630b6dedfd146100ff57806341c0e1b51461014b575b600080fd5b61008a6101ed565b6040805160208082528351818301528351919283929083019185019080838360005b838110156100c45781810151838201526020016100ac565b50505050905090810190601f1680156100f15780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b61008a600480360360c081101561011557600080fd5b506001600160a01b0381358116916020810135821691604082013581169160608101358216916080820135169060a00135610226565b610153610423565b005b6101536004803603602081101561016b57600080fd5b50356001600160a01b03166104ab565b610183610635565b604080516001600160a01b039092168252519081900360200190f35b610153600480360360a08110156101b557600080fd5b506001600160a01b0381358116916020810135821691604082013581169160608101359091169060800135610644565b610183610857565b6040518060400160405280601e81526020017f526573747269637420546f204465706f736974696e67204163636f756e74000081525081565b60606000610232610866565b6001600160a01b03166331aaa74a866040518263ffffffff1660e01b815260040180826001600160a01b03166001600160a01b0316815260200191505060206040518083038186803b15801561028757600080fd5b505afa15801561029b573d6000803e3d6000fd5b505050506040513d60208110156102b157600080fd5b50516040805160608101909152602380825291925060009161090e602083013982866040516020018084805190602001908083835b602083106103055780518252601f1990920191602091820191016102e6565b6001836020036101000a038019825116818451168082178552505050505050905001836001600160a01b03166001600160a01b031660601b8152601401826001600160a01b03166001600160a01b031660601b8152601401935050505060405160208183030381529060405280519060200120905060006103846108d2565b6001600160a01b0316637ae1cfca836040518263ffffffff1660e01b81526004018082815260200191505060206040518083038186803b1580156103c757600080fd5b505afa1580156103db573d6000803e3d6000fd5b505050506040513d60208110156103f157600080fd5b5051905080610416576040518060600160405280602781526020016109566027913993505b5050509695505050505050565b6000546001600160a01b031633148061044c57506001546000546001600160a01b039081169116145b61049d576040805162461bcd60e51b815260206004820152601960248201527f4f776e6572206163636f756e7420697320726571756972656400000000000000604482015290519081900360640190fd5b6000546001600160a01b0316ff5b6000546001600160a01b03163314806104d457506001546000546001600160a01b039081169116145b610525576040805162461bcd60e51b815260206004820152601960248201527f4f776e6572206163636f756e7420697320726571756972656400000000000000604482015290519081900360640190fd5b6000546001600160a01b03828116911614156105725760405162461bcd60e51b81526004018080602001828103825260258152602001806109316025913960400191505060405180910390fd5b6001600160a01b0381166105cd576040805162461bcd60e51b815260206004820181905260248201527f4e6577204f776e65722063616e6e6f74206265207a65726f2061646472657373604482015290519081900360640190fd5b600080546001600160a01b038381167fffffffffffffffffffffffff0000000000000000000000000000000000000000831681178455604051919092169283917f8934ce4adea8d9ce0d714d2c22b86790e41b7731c84b926fbbdc1d40ff6533c99190a35050565b6001546001600160a01b031681565b600061064e610866565b6001600160a01b03166331aaa74a856040518263ffffffff1660e01b815260040180826001600160a01b03166001600160a01b0316815260200191505060206040518083038186803b1580156106a357600080fd5b505afa1580156106b7573d6000803e3d6000fd5b505050506040513d60208110156106cd57600080fd5b50516040805160608101909152602380825291925060009161090e602083013982856040516020018084805190602001908083835b602083106107215780518252601f199092019160209182019101610702565b6001836020036101000a038019825116818451168082178552505050505050905001836001600160a01b03166001600160a01b031660601b8152601401826001600160a01b03166001600160a01b031660601b8152601401935050505060405160208183030381529060405280519060200120905060006107a06108d2565b6001600160a01b0316637ae1cfca836040518263ffffffff1660e01b81526004018082815260200191505060206040518083038186803b1580156107e357600080fd5b505afa1580156107f7573d6000803e3d6000fd5b505050506040513d602081101561080d57600080fd5b505190508061084d5760405162461bcd60e51b81526004018080602001828103825260278152602001806109566027913960400191505060405180910390fd5b5050505050505050565b6000546001600160a01b031681565b6000336001600160a01b0316637b1039996040518163ffffffff1660e01b815260040160206040518083038186803b1580156108a157600080fd5b505afa1580156108b5573d6000803e3d6000fd5b505050506040513d60208110156108cb57600080fd5b5051905090565b6000336001600160a01b031663975057e76040518163ffffffff1660e01b815260040160206040518083038186803b1580156108a157600080fdfe5265737472696374546f4465706f736974696e674163636f756e742e616464726573734e6577204f776e65722063616e6e6f74206265207468652063757272656e74206f776e65725472616e7366657273207265737472696374656420746f206465706f736974206163636f756e74a265627a7a72315820cd5fbb19fdfb90ef4c19b020a72efccf348b7659df0836c700ad9557bfee8de564736f6c634300050c0032`

// DeployRestrictToDepositingAccount deploys a new Ethereum contract, binding an instance of RestrictToDepositingAccount to it.
func DeployRestrictToDepositingAccount(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *RestrictToDepositingAccount, error) {
	parsed, err := abi.JSON(strings.NewReader(RestrictToDepositingAccountABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(RestrictToDepositingAccountBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &RestrictToDepositingAccount{RestrictToDepositingAccountCaller: RestrictToDepositingAccountCaller{contract: contract}, RestrictToDepositingAccountTransactor: RestrictToDepositingAccountTransactor{contract: contract}, RestrictToDepositingAccountFilterer: RestrictToDepositingAccountFilterer{contract: contract}}, nil
}

// RestrictToDepositingAccount is an auto generated Go binding around an Ethereum contract.
type RestrictToDepositingAccount struct {
	RestrictToDepositingAccountCaller     // Read-only binding to the contract
	RestrictToDepositingAccountTransactor // Write-only binding to the contract
	RestrictToDepositingAccountFilterer   // Log filterer for contract events
}

// RestrictToDepositingAccountCaller is an auto generated read-only Go binding around an Ethereum contract.
type RestrictToDepositingAccountCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RestrictToDepositingAccountTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RestrictToDepositingAccountTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RestrictToDepositingAccountFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RestrictToDepositingAccountFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RestrictToDepositingAccountSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RestrictToDepositingAccountSession struct {
	Contract     *RestrictToDepositingAccount // Generic contract binding to set the session for
	CallOpts     bind.CallOpts                // Call options to use throughout this session
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// RestrictToDepositingAccountCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RestrictToDepositingAccountCallerSession struct {
	Contract *RestrictToDepositingAccountCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                      // Call options to use throughout this session
}

// RestrictToDepositingAccountTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RestrictToDepositingAccountTransactorSession struct {
	Contract     *RestrictToDepositingAccountTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                      // Transaction auth options to use throughout this session
}

// RestrictToDepositingAccountRaw is an auto generated low-level Go binding around an Ethereum contract.
type RestrictToDepositingAccountRaw struct {
	Contract *RestrictToDepositingAccount // Generic contract binding to access the raw methods on
}

// RestrictToDepositingAccountCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RestrictToDepositingAccountCallerRaw struct {
	Contract *RestrictToDepositingAccountCaller // Generic read-only contract binding to access the raw methods on
}

// RestrictToDepositingAccountTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RestrictToDepositingAccountTransactorRaw struct {
	Contract *RestrictToDepositingAccountTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRestrictToDepositingAccount creates a new instance of RestrictToDepositingAccount, bound to a specific deployed contract.
func NewRestrictToDepositingAccount(address common.Address, backend bind.ContractBackend) (*RestrictToDepositingAccount, error) {
	contract, err := bindRestrictToDepositingAccount(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &RestrictToDepositingAccount{RestrictToDepositingAccountCaller: RestrictToDepositingAccountCaller{contract: contract}, RestrictToDepositingAccountTransactor: RestrictToDepositingAccountTransactor{contract: contract}, RestrictToDepositingAccountFilterer: RestrictToDepositingAccountFilterer{contract: contract}}, nil
}

// NewRestrictToDepositingAccountCaller creates a new read-only instance of RestrictToDepositingAccount, bound to a specific deployed contract.
func NewRestrictToDepositingAccountCaller(address common.Address, caller bind.ContractCaller) (*RestrictToDepositingAccountCaller, error) {
	contract, err := bindRestrictToDepositingAccount(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RestrictToDepositingAccountCaller{contract: contract}, nil
}

// NewRestrictToDepositingAccountTransactor creates a new write-only instance of RestrictToDepositingAccount, bound to a specific deployed contract.
func NewRestrictToDepositingAccountTransactor(address common.Address, transactor bind.ContractTransactor) (*RestrictToDepositingAccountTransactor, error) {
	contract, err := bindRestrictToDepositingAccount(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RestrictToDepositingAccountTransactor{contract: contract}, nil
}

// NewRestrictToDepositingAccountFilterer creates a new log filterer instance of RestrictToDepositingAccount, bound to a specific deployed contract.
func NewRestrictToDepositingAccountFilterer(address common.Address, filterer bind.ContractFilterer) (*RestrictToDepositingAccountFilterer, error) {
	contract, err := bindRestrictToDepositingAccount(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RestrictToDepositingAccountFilterer{contract: contract}, nil
}

// bindRestrictToDepositingAccount binds a generic wrapper to an already deployed contract.
func bindRestrictToDepositingAccount(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(RestrictToDepositingAccountABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RestrictToDepositingAccount *RestrictToDepositingAccountRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _RestrictToDepositingAccount.Contract.RestrictToDepositingAccountCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RestrictToDepositingAccount *RestrictToDepositingAccountRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RestrictToDepositingAccount.Contract.RestrictToDepositingAccountTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RestrictToDepositingAccount *RestrictToDepositingAccountRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RestrictToDepositingAccount.Contract.RestrictToDepositingAccountTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RestrictToDepositingAccount *RestrictToDepositingAccountCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _RestrictToDepositingAccount.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RestrictToDepositingAccount *RestrictToDepositingAccountTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RestrictToDepositingAccount.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RestrictToDepositingAccount *RestrictToDepositingAccountTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RestrictToDepositingAccount.Contract.contract.Transact(opts, method, params...)
}

// ZEROADDRESS is a free data retrieval call binding the contract method 0x538ba4f9.
//
// Solidity: function ZERO_ADDRESS() constant returns(address)
func (_RestrictToDepositingAccount *RestrictToDepositingAccountCaller) ZEROADDRESS(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _RestrictToDepositingAccount.contract.Call(opts, out, "ZERO_ADDRESS")
	return *ret0, err
}

// ZEROADDRESS is a free data retrieval call binding the contract method 0x538ba4f9.
//
// Solidity: function ZERO_ADDRESS() constant returns(address)
func (_RestrictToDepositingAccount *RestrictToDepositingAccountSession) ZEROADDRESS() (common.Address, error) {
	return _RestrictToDepositingAccount.Contract.ZEROADDRESS(&_RestrictToDepositingAccount.CallOpts)
}

// ZEROADDRESS is a free data retrieval call binding the contract method 0x538ba4f9.
//
// Solidity: function ZERO_ADDRESS() constant returns(address)
func (_RestrictToDepositingAccount *RestrictToDepositingAccountCallerSession) ZEROADDRESS() (common.Address, error) {
	return _RestrictToDepositingAccount.Contract.ZEROADDRESS(&_RestrictToDepositingAccount.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_RestrictToDepositingAccount *RestrictToDepositingAccountCaller) Name(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _RestrictToDepositingAccount.contract.Call(opts, out, "name")
	return *ret0, err
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_RestrictToDepositingAccount *RestrictToDepositingAccountSession) Name() (string, error) {
	return _RestrictToDepositingAccount.Contract.Name(&_RestrictToDepositingAccount.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_RestrictToDepositingAccount *RestrictToDepositingAccountCallerSession) Name() (string, error) {
	return _RestrictToDepositingAccount.Contract.Name(&_RestrictToDepositingAccount.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_RestrictToDepositingAccount *RestrictToDepositingAccountCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _RestrictToDepositingAccount.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_RestrictToDepositingAccount *RestrictToDepositingAccountSession) Owner() (common.Address, error) {
	return _RestrictToDepositingAccount.Contract.Owner(&_RestrictToDepositingAccount.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_RestrictToDepositingAccount *RestrictToDepositingAccountCallerSession) Owner() (common.Address, error) {
	return _RestrictToDepositingAccount.Contract.Owner(&_RestrictToDepositingAccount.CallOpts)
}

// Test is a free data retrieval call binding the contract method 0x0b6dedfd.
//
// Solidity: function test(compliance address, token address, initiator address, from address, to address, tokens uint256) constant returns(s string)
func (_RestrictToDepositingAccount *RestrictToDepositingAccountCaller) Test(opts *bind.CallOpts, compliance common.Address, token common.Address, initiator common.Address, from common.Address, to common.Address, tokens *big.Int) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _RestrictToDepositingAccount.contract.Call(opts, out, "test", compliance, token, initiator, from, to, tokens)
	return *ret0, err
}

// Test is a free data retrieval call binding the contract method 0x0b6dedfd.
//
// Solidity: function test(compliance address, token address, initiator address, from address, to address, tokens uint256) constant returns(s string)
func (_RestrictToDepositingAccount *RestrictToDepositingAccountSession) Test(compliance common.Address, token common.Address, initiator common.Address, from common.Address, to common.Address, tokens *big.Int) (string, error) {
	return _RestrictToDepositingAccount.Contract.Test(&_RestrictToDepositingAccount.CallOpts, compliance, token, initiator, from, to, tokens)
}

// Test is a free data retrieval call binding the contract method 0x0b6dedfd.
//
// Solidity: function test(compliance address, token address, initiator address, from address, to address, tokens uint256) constant returns(s string)
func (_RestrictToDepositingAccount *RestrictToDepositingAccountCallerSession) Test(compliance common.Address, token common.Address, initiator common.Address, from common.Address, to common.Address, tokens *big.Int) (string, error) {
	return _RestrictToDepositingAccount.Contract.Test(&_RestrictToDepositingAccount.CallOpts, compliance, token, initiator, from, to, tokens)
}

// Check is a paid mutator transaction binding the contract method 0x803fcd43.
//
// Solidity: function check(token address, initiator address, from address, to address, tokens uint256) returns()
func (_RestrictToDepositingAccount *RestrictToDepositingAccountTransactor) Check(opts *bind.TransactOpts, token common.Address, initiator common.Address, from common.Address, to common.Address, tokens *big.Int) (*types.Transaction, error) {
	return _RestrictToDepositingAccount.contract.Transact(opts, "check", token, initiator, from, to, tokens)
}

// Check is a paid mutator transaction binding the contract method 0x803fcd43.
//
// Solidity: function check(token address, initiator address, from address, to address, tokens uint256) returns()
func (_RestrictToDepositingAccount *RestrictToDepositingAccountSession) Check(token common.Address, initiator common.Address, from common.Address, to common.Address, tokens *big.Int) (*types.Transaction, error) {
	return _RestrictToDepositingAccount.Contract.Check(&_RestrictToDepositingAccount.TransactOpts, token, initiator, from, to, tokens)
}

// Check is a paid mutator transaction binding the contract method 0x803fcd43.
//
// Solidity: function check(token address, initiator address, from address, to address, tokens uint256) returns()
func (_RestrictToDepositingAccount *RestrictToDepositingAccountTransactorSession) Check(token common.Address, initiator common.Address, from common.Address, to common.Address, tokens *big.Int) (*types.Transaction, error) {
	return _RestrictToDepositingAccount.Contract.Check(&_RestrictToDepositingAccount.TransactOpts, token, initiator, from, to, tokens)
}

// Kill is a paid mutator transaction binding the contract method 0x41c0e1b5.
//
// Solidity: function kill() returns()
func (_RestrictToDepositingAccount *RestrictToDepositingAccountTransactor) Kill(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RestrictToDepositingAccount.contract.Transact(opts, "kill")
}

// Kill is a paid mutator transaction binding the contract method 0x41c0e1b5.
//
// Solidity: function kill() returns()
func (_RestrictToDepositingAccount *RestrictToDepositingAccountSession) Kill() (*types.Transaction, error) {
	return _RestrictToDepositingAccount.Contract.Kill(&_RestrictToDepositingAccount.TransactOpts)
}

// Kill is a paid mutator transaction binding the contract method 0x41c0e1b5.
//
// Solidity: function kill() returns()
func (_RestrictToDepositingAccount *RestrictToDepositingAccountTransactorSession) Kill() (*types.Transaction, error) {
	return _RestrictToDepositingAccount.Contract.Kill(&_RestrictToDepositingAccount.TransactOpts)
}

// TransferOwner is a paid mutator transaction binding the contract method 0x4fb2e45d.
//
// Solidity: function transferOwner(newOwner address) returns()
func (_RestrictToDepositingAccount *RestrictToDepositingAccountTransactor) TransferOwner(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _RestrictToDepositingAccount.contract.Transact(opts, "transferOwner", newOwner)
}

// TransferOwner is a paid mutator transaction binding the contract method 0x4fb2e45d.
//
// Solidity: function transferOwner(newOwner address) returns()
func (_RestrictToDepositingAccount *RestrictToDepositingAccountSession) TransferOwner(newOwner common.Address) (*types.Transaction, error) {
	return _RestrictToDepositingAccount.Contract.TransferOwner(&_RestrictToDepositingAccount.TransactOpts, newOwner)
}

// TransferOwner is a paid mutator transaction binding the contract method 0x4fb2e45d.
//
// Solidity: function transferOwner(newOwner address) returns()
func (_RestrictToDepositingAccount *RestrictToDepositingAccountTransactorSession) TransferOwner(newOwner common.Address) (*types.Transaction, error) {
	return _RestrictToDepositingAccount.Contract.TransferOwner(&_RestrictToDepositingAccount.TransactOpts, newOwner)
}

// RestrictToDepositingAccountOwnerTransferredIterator is returned from FilterOwnerTransferred and is used to iterate over the raw logs and unpacked data for OwnerTransferred events raised by the RestrictToDepositingAccount contract.
type RestrictToDepositingAccountOwnerTransferredIterator struct {
	Event *RestrictToDepositingAccountOwnerTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *RestrictToDepositingAccountOwnerTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RestrictToDepositingAccountOwnerTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(RestrictToDepositingAccountOwnerTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *RestrictToDepositingAccountOwnerTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RestrictToDepositingAccountOwnerTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RestrictToDepositingAccountOwnerTransferred represents a OwnerTransferred event raised by the RestrictToDepositingAccount contract.
type RestrictToDepositingAccountOwnerTransferred struct {
	OldOwner common.Address
	NewOwner common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOwnerTransferred is a free log retrieval operation binding the contract event 0x8934ce4adea8d9ce0d714d2c22b86790e41b7731c84b926fbbdc1d40ff6533c9.
//
// Solidity: e OwnerTransferred(oldOwner indexed address, newOwner indexed address)
func (_RestrictToDepositingAccount *RestrictToDepositingAccountFilterer) FilterOwnerTransferred(opts *bind.FilterOpts, oldOwner []common.Address, newOwner []common.Address) (*RestrictToDepositingAccountOwnerTransferredIterator, error) {

	var oldOwnerRule []interface{}
	for _, oldOwnerItem := range oldOwner {
		oldOwnerRule = append(oldOwnerRule, oldOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _RestrictToDepositingAccount.contract.FilterLogs(opts, "OwnerTransferred", oldOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &RestrictToDepositingAccountOwnerTransferredIterator{contract: _RestrictToDepositingAccount.contract, event: "OwnerTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnerTransferred is a free log subscription operation binding the contract event 0x8934ce4adea8d9ce0d714d2c22b86790e41b7731c84b926fbbdc1d40ff6533c9.
//
// Solidity: e OwnerTransferred(oldOwner indexed address, newOwner indexed address)
func (_RestrictToDepositingAccount *RestrictToDepositingAccountFilterer) WatchOwnerTransferred(opts *bind.WatchOpts, sink chan<- *RestrictToDepositingAccountOwnerTransferred, oldOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var oldOwnerRule []interface{}
	for _, oldOwnerItem := range oldOwner {
		oldOwnerRule = append(oldOwnerRule, oldOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _RestrictToDepositingAccount.contract.WatchLogs(opts, "OwnerTransferred", oldOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RestrictToDepositingAccountOwnerTransferred)
				if err := _RestrictToDepositingAccount.contract.UnpackLog(event, "OwnerTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}
