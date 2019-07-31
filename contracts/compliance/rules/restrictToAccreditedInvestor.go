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

// RestrictToAccreditedInvestorABI is the input ABI used to generate the binding from.
const RestrictToAccreditedInvestorABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"kill\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwner\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ZERO_ADDRESS\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"token\",\"type\":\"address\"},{\"name\":\"initiator\",\"type\":\"address\"},{\"name\":\"from\",\"type\":\"address\"},{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"toKind\",\"type\":\"uint8\"},{\"name\":\"tokens\",\"type\":\"uint256\"}],\"name\":\"check\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"compliance\",\"type\":\"address\"},{\"name\":\"token\",\"type\":\"address\"},{\"name\":\"initiator\",\"type\":\"address\"},{\"name\":\"from\",\"type\":\"address\"},{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"toKind\",\"type\":\"uint8\"},{\"name\":\"tokens\",\"type\":\"uint256\"}],\"name\":\"test\",\"outputs\":[{\"name\":\"s\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"oldOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnerTransferred\",\"type\":\"event\"}]"

// RestrictToAccreditedInvestorBin is the compiled bytecode used for deploying new contracts.
const RestrictToAccreditedInvestorBin = `600180546001600160a01b031916905560c0604052601f60808190527f526573747269637420546f204163637265646974656420496e766573746f720060a090815261004e9160029190610066565b50600080546001600160a01b03191633179055610101565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106100a757805160ff19168380011785556100d4565b828001600101855582156100d4579182015b828111156100d45782518255916020019190600101906100b9565b506100e09291506100e4565b5090565b6100fe91905b808211156100e057600081556001016100ea565b90565b6107b3806101106000396000f3fe608060405234801561001057600080fd5b506004361061007d5760003560e01c8063538ba4f91161005b578063538ba4f91461012f5780635a47e1c7146101535780638da5cb5b146101a2578063a9189562146101aa5761007d565b806306fdde031461008257806341c0e1b5146100ff5780634fb2e45d14610109575b600080fd5b61008a6101ff565b6040805160208082528351818301528351919283929083019185019080838360005b838110156100c45781810151838201526020016100ac565b50505050905090810190601f1680156100f15780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b61010761028a565b005b6101076004803603602081101561011f57600080fd5b50356001600160a01b0316610312565b61013761049c565b604080516001600160a01b039092168252519081900360200190f35b610107600480360360c081101561016957600080fd5b506001600160a01b0381358116916020810135821691604082013581169160608101359091169060ff6080820135169060a001356104ab565b61013761051b565b61008a600480360360e08110156101c057600080fd5b506001600160a01b0381358116916020810135821691604082013581169160608101358216916080820135169060ff60a0820135169060c0013561052a565b6002805460408051602060018416156101000260001901909316849004601f810184900484028201840190925281815292918301828280156102825780601f1061025757610100808354040283529160200191610282565b820191906000526020600020905b81548152906001019060200180831161026557829003601f168201915b505050505081565b6000546001600160a01b03163314806102b357506001546000546001600160a01b039081169116145b610304576040805162461bcd60e51b815260206004820152601960248201527f4f776e6572206163636f756e7420697320726571756972656400000000000000604482015290519081900360640190fd5b6000546001600160a01b0316ff5b6000546001600160a01b031633148061033b57506001546000546001600160a01b039081169116145b61038c576040805162461bcd60e51b815260206004820152601960248201527f4f776e6572206163636f756e7420697320726571756972656400000000000000604482015290519081900360640190fd5b6000546001600160a01b03828116911614156103d95760405162461bcd60e51b81526004018080602001828103825260258152602001806107046025913960400191505060405180910390fd5b6001600160a01b038116610434576040805162461bcd60e51b815260206004820181905260248201527f4e6577204f776e65722063616e6e6f74206265207a65726f2061646472657373604482015290519081900360640190fd5b600080546001600160a01b038381167fffffffffffffffffffffffff0000000000000000000000000000000000000000831681178455604051919092169283917f8934ce4adea8d9ce0d714d2c22b86790e41b7731c84b926fbbdc1d40ff6533c99190a35050565b6001546001600160a01b031681565b60ff821660041415610513576104d8836104c36105e6565b6001600160a01b03169063ffffffff61065216565b6105135760405162461bcd60e51b815260040180806020018281038252602a815260200180610729602a913960400191505060405180910390fd5b505050505050565b6000546001600160a01b031681565b606060ff831660041480156105ba57506105b884896001600160a01b0316637b1039996040518163ffffffff1660e01b815260040160206040518083038186803b15801561057757600080fd5b505afa15801561058b573d6000803e3d6000fd5b505050506040513d60208110156105a157600080fd5b50516001600160a01b03169063ffffffff61065216565b155b156105db576040518060600160405280602c8152602001610753602c913990505b979650505050505050565b6000336001600160a01b0316637b1039996040518163ffffffff1660e01b815260040160206040518083038186803b15801561062157600080fd5b505afa158015610635573d6000803e3d6000fd5b505050506040513d602081101561064b57600080fd5b5051905090565b604080517f572bc6b10000000000000000000000000000000000000000000000000000000081526001600160a01b038381166004830152600060248301819052925183929186169163572bc6b1916044808301926020929190829003018186803b1580156106bf57600080fd5b505afa1580156106d3573d6000803e3d6000fd5b505050506040513d60208110156106e957600080fd5b50514260109190911c65ffffffffffff161194935050505056fe4e6577204f776e65722063616e6e6f74206265207468652063757272656e74206f776e657254686520746f2061646472657373206973206e6f742063757272656e746c7920616363726564697465645468652027746f272061646472657373206973206e6f742063757272656e746c792061636372656469746564a265627a7a723058204a543e9b407b60253f32f75c542f4b37ca019674d648076df7c4f9209d598ec664736f6c634300050a0032`

// DeployRestrictToAccreditedInvestor deploys a new Ethereum contract, binding an instance of RestrictToAccreditedInvestor to it.
func DeployRestrictToAccreditedInvestor(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *RestrictToAccreditedInvestor, error) {
	parsed, err := abi.JSON(strings.NewReader(RestrictToAccreditedInvestorABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(RestrictToAccreditedInvestorBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &RestrictToAccreditedInvestor{RestrictToAccreditedInvestorCaller: RestrictToAccreditedInvestorCaller{contract: contract}, RestrictToAccreditedInvestorTransactor: RestrictToAccreditedInvestorTransactor{contract: contract}, RestrictToAccreditedInvestorFilterer: RestrictToAccreditedInvestorFilterer{contract: contract}}, nil
}

// RestrictToAccreditedInvestor is an auto generated Go binding around an Ethereum contract.
type RestrictToAccreditedInvestor struct {
	RestrictToAccreditedInvestorCaller     // Read-only binding to the contract
	RestrictToAccreditedInvestorTransactor // Write-only binding to the contract
	RestrictToAccreditedInvestorFilterer   // Log filterer for contract events
}

// RestrictToAccreditedInvestorCaller is an auto generated read-only Go binding around an Ethereum contract.
type RestrictToAccreditedInvestorCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RestrictToAccreditedInvestorTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RestrictToAccreditedInvestorTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RestrictToAccreditedInvestorFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RestrictToAccreditedInvestorFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RestrictToAccreditedInvestorSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RestrictToAccreditedInvestorSession struct {
	Contract     *RestrictToAccreditedInvestor // Generic contract binding to set the session for
	CallOpts     bind.CallOpts                 // Call options to use throughout this session
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// RestrictToAccreditedInvestorCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RestrictToAccreditedInvestorCallerSession struct {
	Contract *RestrictToAccreditedInvestorCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                       // Call options to use throughout this session
}

// RestrictToAccreditedInvestorTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RestrictToAccreditedInvestorTransactorSession struct {
	Contract     *RestrictToAccreditedInvestorTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                       // Transaction auth options to use throughout this session
}

// RestrictToAccreditedInvestorRaw is an auto generated low-level Go binding around an Ethereum contract.
type RestrictToAccreditedInvestorRaw struct {
	Contract *RestrictToAccreditedInvestor // Generic contract binding to access the raw methods on
}

// RestrictToAccreditedInvestorCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RestrictToAccreditedInvestorCallerRaw struct {
	Contract *RestrictToAccreditedInvestorCaller // Generic read-only contract binding to access the raw methods on
}

// RestrictToAccreditedInvestorTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RestrictToAccreditedInvestorTransactorRaw struct {
	Contract *RestrictToAccreditedInvestorTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRestrictToAccreditedInvestor creates a new instance of RestrictToAccreditedInvestor, bound to a specific deployed contract.
func NewRestrictToAccreditedInvestor(address common.Address, backend bind.ContractBackend) (*RestrictToAccreditedInvestor, error) {
	contract, err := bindRestrictToAccreditedInvestor(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &RestrictToAccreditedInvestor{RestrictToAccreditedInvestorCaller: RestrictToAccreditedInvestorCaller{contract: contract}, RestrictToAccreditedInvestorTransactor: RestrictToAccreditedInvestorTransactor{contract: contract}, RestrictToAccreditedInvestorFilterer: RestrictToAccreditedInvestorFilterer{contract: contract}}, nil
}

// NewRestrictToAccreditedInvestorCaller creates a new read-only instance of RestrictToAccreditedInvestor, bound to a specific deployed contract.
func NewRestrictToAccreditedInvestorCaller(address common.Address, caller bind.ContractCaller) (*RestrictToAccreditedInvestorCaller, error) {
	contract, err := bindRestrictToAccreditedInvestor(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RestrictToAccreditedInvestorCaller{contract: contract}, nil
}

// NewRestrictToAccreditedInvestorTransactor creates a new write-only instance of RestrictToAccreditedInvestor, bound to a specific deployed contract.
func NewRestrictToAccreditedInvestorTransactor(address common.Address, transactor bind.ContractTransactor) (*RestrictToAccreditedInvestorTransactor, error) {
	contract, err := bindRestrictToAccreditedInvestor(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RestrictToAccreditedInvestorTransactor{contract: contract}, nil
}

// NewRestrictToAccreditedInvestorFilterer creates a new log filterer instance of RestrictToAccreditedInvestor, bound to a specific deployed contract.
func NewRestrictToAccreditedInvestorFilterer(address common.Address, filterer bind.ContractFilterer) (*RestrictToAccreditedInvestorFilterer, error) {
	contract, err := bindRestrictToAccreditedInvestor(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RestrictToAccreditedInvestorFilterer{contract: contract}, nil
}

// bindRestrictToAccreditedInvestor binds a generic wrapper to an already deployed contract.
func bindRestrictToAccreditedInvestor(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(RestrictToAccreditedInvestorABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RestrictToAccreditedInvestor *RestrictToAccreditedInvestorRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _RestrictToAccreditedInvestor.Contract.RestrictToAccreditedInvestorCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RestrictToAccreditedInvestor *RestrictToAccreditedInvestorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RestrictToAccreditedInvestor.Contract.RestrictToAccreditedInvestorTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RestrictToAccreditedInvestor *RestrictToAccreditedInvestorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RestrictToAccreditedInvestor.Contract.RestrictToAccreditedInvestorTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RestrictToAccreditedInvestor *RestrictToAccreditedInvestorCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _RestrictToAccreditedInvestor.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RestrictToAccreditedInvestor *RestrictToAccreditedInvestorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RestrictToAccreditedInvestor.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RestrictToAccreditedInvestor *RestrictToAccreditedInvestorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RestrictToAccreditedInvestor.Contract.contract.Transact(opts, method, params...)
}

// ZEROADDRESS is a free data retrieval call binding the contract method 0x538ba4f9.
//
// Solidity: function ZERO_ADDRESS() constant returns(address)
func (_RestrictToAccreditedInvestor *RestrictToAccreditedInvestorCaller) ZEROADDRESS(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _RestrictToAccreditedInvestor.contract.Call(opts, out, "ZERO_ADDRESS")
	return *ret0, err
}

// ZEROADDRESS is a free data retrieval call binding the contract method 0x538ba4f9.
//
// Solidity: function ZERO_ADDRESS() constant returns(address)
func (_RestrictToAccreditedInvestor *RestrictToAccreditedInvestorSession) ZEROADDRESS() (common.Address, error) {
	return _RestrictToAccreditedInvestor.Contract.ZEROADDRESS(&_RestrictToAccreditedInvestor.CallOpts)
}

// ZEROADDRESS is a free data retrieval call binding the contract method 0x538ba4f9.
//
// Solidity: function ZERO_ADDRESS() constant returns(address)
func (_RestrictToAccreditedInvestor *RestrictToAccreditedInvestorCallerSession) ZEROADDRESS() (common.Address, error) {
	return _RestrictToAccreditedInvestor.Contract.ZEROADDRESS(&_RestrictToAccreditedInvestor.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_RestrictToAccreditedInvestor *RestrictToAccreditedInvestorCaller) Name(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _RestrictToAccreditedInvestor.contract.Call(opts, out, "name")
	return *ret0, err
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_RestrictToAccreditedInvestor *RestrictToAccreditedInvestorSession) Name() (string, error) {
	return _RestrictToAccreditedInvestor.Contract.Name(&_RestrictToAccreditedInvestor.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_RestrictToAccreditedInvestor *RestrictToAccreditedInvestorCallerSession) Name() (string, error) {
	return _RestrictToAccreditedInvestor.Contract.Name(&_RestrictToAccreditedInvestor.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_RestrictToAccreditedInvestor *RestrictToAccreditedInvestorCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _RestrictToAccreditedInvestor.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_RestrictToAccreditedInvestor *RestrictToAccreditedInvestorSession) Owner() (common.Address, error) {
	return _RestrictToAccreditedInvestor.Contract.Owner(&_RestrictToAccreditedInvestor.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_RestrictToAccreditedInvestor *RestrictToAccreditedInvestorCallerSession) Owner() (common.Address, error) {
	return _RestrictToAccreditedInvestor.Contract.Owner(&_RestrictToAccreditedInvestor.CallOpts)
}

// Test is a free data retrieval call binding the contract method 0xa9189562.
//
// Solidity: function test(compliance address, token address, initiator address, from address, to address, toKind uint8, tokens uint256) constant returns(s string)
func (_RestrictToAccreditedInvestor *RestrictToAccreditedInvestorCaller) Test(opts *bind.CallOpts, compliance common.Address, token common.Address, initiator common.Address, from common.Address, to common.Address, toKind uint8, tokens *big.Int) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _RestrictToAccreditedInvestor.contract.Call(opts, out, "test", compliance, token, initiator, from, to, toKind, tokens)
	return *ret0, err
}

// Test is a free data retrieval call binding the contract method 0xa9189562.
//
// Solidity: function test(compliance address, token address, initiator address, from address, to address, toKind uint8, tokens uint256) constant returns(s string)
func (_RestrictToAccreditedInvestor *RestrictToAccreditedInvestorSession) Test(compliance common.Address, token common.Address, initiator common.Address, from common.Address, to common.Address, toKind uint8, tokens *big.Int) (string, error) {
	return _RestrictToAccreditedInvestor.Contract.Test(&_RestrictToAccreditedInvestor.CallOpts, compliance, token, initiator, from, to, toKind, tokens)
}

// Test is a free data retrieval call binding the contract method 0xa9189562.
//
// Solidity: function test(compliance address, token address, initiator address, from address, to address, toKind uint8, tokens uint256) constant returns(s string)
func (_RestrictToAccreditedInvestor *RestrictToAccreditedInvestorCallerSession) Test(compliance common.Address, token common.Address, initiator common.Address, from common.Address, to common.Address, toKind uint8, tokens *big.Int) (string, error) {
	return _RestrictToAccreditedInvestor.Contract.Test(&_RestrictToAccreditedInvestor.CallOpts, compliance, token, initiator, from, to, toKind, tokens)
}

// Check is a paid mutator transaction binding the contract method 0x5a47e1c7.
//
// Solidity: function check(token address, initiator address, from address, to address, toKind uint8, tokens uint256) returns()
func (_RestrictToAccreditedInvestor *RestrictToAccreditedInvestorTransactor) Check(opts *bind.TransactOpts, token common.Address, initiator common.Address, from common.Address, to common.Address, toKind uint8, tokens *big.Int) (*types.Transaction, error) {
	return _RestrictToAccreditedInvestor.contract.Transact(opts, "check", token, initiator, from, to, toKind, tokens)
}

// Check is a paid mutator transaction binding the contract method 0x5a47e1c7.
//
// Solidity: function check(token address, initiator address, from address, to address, toKind uint8, tokens uint256) returns()
func (_RestrictToAccreditedInvestor *RestrictToAccreditedInvestorSession) Check(token common.Address, initiator common.Address, from common.Address, to common.Address, toKind uint8, tokens *big.Int) (*types.Transaction, error) {
	return _RestrictToAccreditedInvestor.Contract.Check(&_RestrictToAccreditedInvestor.TransactOpts, token, initiator, from, to, toKind, tokens)
}

// Check is a paid mutator transaction binding the contract method 0x5a47e1c7.
//
// Solidity: function check(token address, initiator address, from address, to address, toKind uint8, tokens uint256) returns()
func (_RestrictToAccreditedInvestor *RestrictToAccreditedInvestorTransactorSession) Check(token common.Address, initiator common.Address, from common.Address, to common.Address, toKind uint8, tokens *big.Int) (*types.Transaction, error) {
	return _RestrictToAccreditedInvestor.Contract.Check(&_RestrictToAccreditedInvestor.TransactOpts, token, initiator, from, to, toKind, tokens)
}

// Kill is a paid mutator transaction binding the contract method 0x41c0e1b5.
//
// Solidity: function kill() returns()
func (_RestrictToAccreditedInvestor *RestrictToAccreditedInvestorTransactor) Kill(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RestrictToAccreditedInvestor.contract.Transact(opts, "kill")
}

// Kill is a paid mutator transaction binding the contract method 0x41c0e1b5.
//
// Solidity: function kill() returns()
func (_RestrictToAccreditedInvestor *RestrictToAccreditedInvestorSession) Kill() (*types.Transaction, error) {
	return _RestrictToAccreditedInvestor.Contract.Kill(&_RestrictToAccreditedInvestor.TransactOpts)
}

// Kill is a paid mutator transaction binding the contract method 0x41c0e1b5.
//
// Solidity: function kill() returns()
func (_RestrictToAccreditedInvestor *RestrictToAccreditedInvestorTransactorSession) Kill() (*types.Transaction, error) {
	return _RestrictToAccreditedInvestor.Contract.Kill(&_RestrictToAccreditedInvestor.TransactOpts)
}

// TransferOwner is a paid mutator transaction binding the contract method 0x4fb2e45d.
//
// Solidity: function transferOwner(newOwner address) returns()
func (_RestrictToAccreditedInvestor *RestrictToAccreditedInvestorTransactor) TransferOwner(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _RestrictToAccreditedInvestor.contract.Transact(opts, "transferOwner", newOwner)
}

// TransferOwner is a paid mutator transaction binding the contract method 0x4fb2e45d.
//
// Solidity: function transferOwner(newOwner address) returns()
func (_RestrictToAccreditedInvestor *RestrictToAccreditedInvestorSession) TransferOwner(newOwner common.Address) (*types.Transaction, error) {
	return _RestrictToAccreditedInvestor.Contract.TransferOwner(&_RestrictToAccreditedInvestor.TransactOpts, newOwner)
}

// TransferOwner is a paid mutator transaction binding the contract method 0x4fb2e45d.
//
// Solidity: function transferOwner(newOwner address) returns()
func (_RestrictToAccreditedInvestor *RestrictToAccreditedInvestorTransactorSession) TransferOwner(newOwner common.Address) (*types.Transaction, error) {
	return _RestrictToAccreditedInvestor.Contract.TransferOwner(&_RestrictToAccreditedInvestor.TransactOpts, newOwner)
}

// RestrictToAccreditedInvestorOwnerTransferredIterator is returned from FilterOwnerTransferred and is used to iterate over the raw logs and unpacked data for OwnerTransferred events raised by the RestrictToAccreditedInvestor contract.
type RestrictToAccreditedInvestorOwnerTransferredIterator struct {
	Event *RestrictToAccreditedInvestorOwnerTransferred // Event containing the contract specifics and raw log

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
func (it *RestrictToAccreditedInvestorOwnerTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RestrictToAccreditedInvestorOwnerTransferred)
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
		it.Event = new(RestrictToAccreditedInvestorOwnerTransferred)
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
func (it *RestrictToAccreditedInvestorOwnerTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RestrictToAccreditedInvestorOwnerTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RestrictToAccreditedInvestorOwnerTransferred represents a OwnerTransferred event raised by the RestrictToAccreditedInvestor contract.
type RestrictToAccreditedInvestorOwnerTransferred struct {
	OldOwner common.Address
	NewOwner common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOwnerTransferred is a free log retrieval operation binding the contract event 0x8934ce4adea8d9ce0d714d2c22b86790e41b7731c84b926fbbdc1d40ff6533c9.
//
// Solidity: e OwnerTransferred(oldOwner indexed address, newOwner indexed address)
func (_RestrictToAccreditedInvestor *RestrictToAccreditedInvestorFilterer) FilterOwnerTransferred(opts *bind.FilterOpts, oldOwner []common.Address, newOwner []common.Address) (*RestrictToAccreditedInvestorOwnerTransferredIterator, error) {

	var oldOwnerRule []interface{}
	for _, oldOwnerItem := range oldOwner {
		oldOwnerRule = append(oldOwnerRule, oldOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _RestrictToAccreditedInvestor.contract.FilterLogs(opts, "OwnerTransferred", oldOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &RestrictToAccreditedInvestorOwnerTransferredIterator{contract: _RestrictToAccreditedInvestor.contract, event: "OwnerTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnerTransferred is a free log subscription operation binding the contract event 0x8934ce4adea8d9ce0d714d2c22b86790e41b7731c84b926fbbdc1d40ff6533c9.
//
// Solidity: e OwnerTransferred(oldOwner indexed address, newOwner indexed address)
func (_RestrictToAccreditedInvestor *RestrictToAccreditedInvestorFilterer) WatchOwnerTransferred(opts *bind.WatchOpts, sink chan<- *RestrictToAccreditedInvestorOwnerTransferred, oldOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var oldOwnerRule []interface{}
	for _, oldOwnerItem := range oldOwner {
		oldOwnerRule = append(oldOwnerRule, oldOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _RestrictToAccreditedInvestor.contract.WatchLogs(opts, "OwnerTransferred", oldOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RestrictToAccreditedInvestorOwnerTransferred)
				if err := _RestrictToAccreditedInvestor.contract.UnpackLog(event, "OwnerTransferred", log); err != nil {
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
