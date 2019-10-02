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

// EscrowRecorderABI is the input ABI used to generate the binding from.
const EscrowRecorderABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"contractICompliance\",\"name\":\"compliance\",\"type\":\"address\"},{\"internalType\":\"contractIT0ken\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"initiator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokens\",\"type\":\"uint256\"}],\"name\":\"test\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"s\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"kill\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwner\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ZERO_ADDRESS\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"contractIT0ken\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"initiator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokens\",\"type\":\"uint256\"}],\"name\":\"check\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnerTransferred\",\"type\":\"event\"}]"

// EscrowRecorderBin is the compiled bytecode used for deploying new contracts.
const EscrowRecorderBin = `600180546001600160a01b031916905560c0604052600f60808190527f457363726f77205265636f72646572000000000000000000000000000000000060a090815261004e9160029190610066565b50600080546001600160a01b03191633179055610101565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106100a757805160ff19168380011785556100d4565b828001600101855582156100d4579182015b828111156100d45782518255916020019190600101906100b9565b506100e09291506100e4565b5090565b6100fe91905b808211156100e057600081556001016100ea565b90565b6107ed806101106000396000f3fe608060405234801561001057600080fd5b506004361061007d5760003560e01c80634fb2e45d1161005b5780634fb2e45d14610155578063538ba4f91461017b578063803fcd431461019f5780638da5cb5b146101e55761007d565b806306fdde03146100825780630b6dedfd146100ff57806341c0e1b51461014b575b600080fd5b61008a6101ed565b6040805160208082528351818301528351919283929083019185019080838360005b838110156100c45781810151838201526020016100ac565b50505050905090810190601f1680156100f15780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b61008a600480360360c081101561011557600080fd5b506001600160a01b0381358116916020810135821691604082013581169160608101358216916080820135169060a00135610278565b610153610284565b005b6101536004803603602081101561016b57600080fd5b50356001600160a01b031661030c565b610183610496565b604080516001600160a01b039092168252519081900360200190f35b610153600480360360a08110156101b557600080fd5b506001600160a01b03813581169160208101358216916040820135811691606081013590911690608001356104a5565b61018361055d565b6002805460408051602060018416156101000260001901909316849004601f810184900484028201840190925281815292918301828280156102705780601f1061024557610100808354040283529160200191610270565b820191906000526020600020905b81548152906001019060200180831161025357829003601f168201915b505050505081565b60609695505050505050565b6000546001600160a01b03163314806102ad57506001546000546001600160a01b039081169116145b6102fe576040805162461bcd60e51b815260206004820152601960248201527f4f776e6572206163636f756e7420697320726571756972656400000000000000604482015290519081900360640190fd5b6000546001600160a01b0316ff5b6000546001600160a01b031633148061033557506001546000546001600160a01b039081169116145b610386576040805162461bcd60e51b815260206004820152601960248201527f4f776e6572206163636f756e7420697320726571756972656400000000000000604482015290519081900360640190fd5b6000546001600160a01b03828116911614156103d35760405162461bcd60e51b81526004018080602001828103825260258152602001806107946025913960400191505060405180910390fd5b6001600160a01b03811661042e576040805162461bcd60e51b815260206004820181905260248201527f4e6577204f776e65722063616e6e6f74206265207a65726f2061646472657373604482015290519081900360640190fd5b600080546001600160a01b038381167fffffffffffffffffffffffff0000000000000000000000000000000000000000831681178455604051919092169283917f8934ce4adea8d9ce0d714d2c22b86790e41b7731c84b926fbbdc1d40ff6533c99190a35050565b6001546001600160a01b031681565b60006104b86104b261056c565b856105d8565b90506001600160a01b0381161561055557604080517f172a93fb0000000000000000000000000000000000000000000000000000000081526001600160a01b03888116600483015286811660248301526044820185905291519183169163172a93fb9160648082019260009290919082900301818387803b15801561053c57600080fd5b505af1158015610550573d6000803e3d6000fd5b505050505b505050505050565b6000546001600160a01b031681565b6000336001600160a01b031663975057e76040518163ffffffff1660e01b815260040160206040518083038186803b1580156105a757600080fd5b505afa1580156105bb573d6000803e3d6000fd5b505050506040513d60208110156105d157600080fd5b5051905090565b6000806105e3610758565b6001600160a01b03166331aaa74a846040518263ffffffff1660e01b815260040180826001600160a01b03166001600160a01b0316815260200191505060206040518083038186803b15801561063857600080fd5b505afa15801561064c573d6000803e3d6000fd5b505050506040513d602081101561066257600080fd5b5051604080517f457363726f775265636f726465722e62726f6b657200000000000000000000006020828101919091526bffffffffffffffffffffffff19606085901b16603583015282516029818403018152604983018085528151918301919091207f21f8a72100000000000000000000000000000000000000000000000000000000909152604d8301819052925193945091926001600160a01b038816926321f8a72192606d8082019391829003018186803b15801561072357600080fd5b505afa158015610737573d6000803e3d6000fd5b505050506040513d602081101561074d57600080fd5b505195945050505050565b6000336001600160a01b0316637b1039996040518163ffffffff1660e01b815260040160206040518083038186803b1580156105a757600080fdfe4e6577204f776e65722063616e6e6f74206265207468652063757272656e74206f776e6572a265627a7a723158201583ae0d4953e119193938133ecdea810eb5bf934f677209effe3079109fb86864736f6c634300050b0032`

// DeployEscrowRecorder deploys a new Ethereum contract, binding an instance of EscrowRecorder to it.
func DeployEscrowRecorder(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *EscrowRecorder, error) {
	parsed, err := abi.JSON(strings.NewReader(EscrowRecorderABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(EscrowRecorderBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &EscrowRecorder{EscrowRecorderCaller: EscrowRecorderCaller{contract: contract}, EscrowRecorderTransactor: EscrowRecorderTransactor{contract: contract}, EscrowRecorderFilterer: EscrowRecorderFilterer{contract: contract}}, nil
}

// EscrowRecorder is an auto generated Go binding around an Ethereum contract.
type EscrowRecorder struct {
	EscrowRecorderCaller     // Read-only binding to the contract
	EscrowRecorderTransactor // Write-only binding to the contract
	EscrowRecorderFilterer   // Log filterer for contract events
}

// EscrowRecorderCaller is an auto generated read-only Go binding around an Ethereum contract.
type EscrowRecorderCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EscrowRecorderTransactor is an auto generated write-only Go binding around an Ethereum contract.
type EscrowRecorderTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EscrowRecorderFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type EscrowRecorderFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EscrowRecorderSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type EscrowRecorderSession struct {
	Contract     *EscrowRecorder   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// EscrowRecorderCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type EscrowRecorderCallerSession struct {
	Contract *EscrowRecorderCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// EscrowRecorderTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type EscrowRecorderTransactorSession struct {
	Contract     *EscrowRecorderTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// EscrowRecorderRaw is an auto generated low-level Go binding around an Ethereum contract.
type EscrowRecorderRaw struct {
	Contract *EscrowRecorder // Generic contract binding to access the raw methods on
}

// EscrowRecorderCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type EscrowRecorderCallerRaw struct {
	Contract *EscrowRecorderCaller // Generic read-only contract binding to access the raw methods on
}

// EscrowRecorderTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type EscrowRecorderTransactorRaw struct {
	Contract *EscrowRecorderTransactor // Generic write-only contract binding to access the raw methods on
}

// NewEscrowRecorder creates a new instance of EscrowRecorder, bound to a specific deployed contract.
func NewEscrowRecorder(address common.Address, backend bind.ContractBackend) (*EscrowRecorder, error) {
	contract, err := bindEscrowRecorder(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &EscrowRecorder{EscrowRecorderCaller: EscrowRecorderCaller{contract: contract}, EscrowRecorderTransactor: EscrowRecorderTransactor{contract: contract}, EscrowRecorderFilterer: EscrowRecorderFilterer{contract: contract}}, nil
}

// NewEscrowRecorderCaller creates a new read-only instance of EscrowRecorder, bound to a specific deployed contract.
func NewEscrowRecorderCaller(address common.Address, caller bind.ContractCaller) (*EscrowRecorderCaller, error) {
	contract, err := bindEscrowRecorder(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &EscrowRecorderCaller{contract: contract}, nil
}

// NewEscrowRecorderTransactor creates a new write-only instance of EscrowRecorder, bound to a specific deployed contract.
func NewEscrowRecorderTransactor(address common.Address, transactor bind.ContractTransactor) (*EscrowRecorderTransactor, error) {
	contract, err := bindEscrowRecorder(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &EscrowRecorderTransactor{contract: contract}, nil
}

// NewEscrowRecorderFilterer creates a new log filterer instance of EscrowRecorder, bound to a specific deployed contract.
func NewEscrowRecorderFilterer(address common.Address, filterer bind.ContractFilterer) (*EscrowRecorderFilterer, error) {
	contract, err := bindEscrowRecorder(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &EscrowRecorderFilterer{contract: contract}, nil
}

// bindEscrowRecorder binds a generic wrapper to an already deployed contract.
func bindEscrowRecorder(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(EscrowRecorderABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EscrowRecorder *EscrowRecorderRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _EscrowRecorder.Contract.EscrowRecorderCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EscrowRecorder *EscrowRecorderRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EscrowRecorder.Contract.EscrowRecorderTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EscrowRecorder *EscrowRecorderRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EscrowRecorder.Contract.EscrowRecorderTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EscrowRecorder *EscrowRecorderCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _EscrowRecorder.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EscrowRecorder *EscrowRecorderTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EscrowRecorder.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EscrowRecorder *EscrowRecorderTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EscrowRecorder.Contract.contract.Transact(opts, method, params...)
}

// ZEROADDRESS is a free data retrieval call binding the contract method 0x538ba4f9.
//
// Solidity: function ZERO_ADDRESS() constant returns(address)
func (_EscrowRecorder *EscrowRecorderCaller) ZEROADDRESS(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _EscrowRecorder.contract.Call(opts, out, "ZERO_ADDRESS")
	return *ret0, err
}

// ZEROADDRESS is a free data retrieval call binding the contract method 0x538ba4f9.
//
// Solidity: function ZERO_ADDRESS() constant returns(address)
func (_EscrowRecorder *EscrowRecorderSession) ZEROADDRESS() (common.Address, error) {
	return _EscrowRecorder.Contract.ZEROADDRESS(&_EscrowRecorder.CallOpts)
}

// ZEROADDRESS is a free data retrieval call binding the contract method 0x538ba4f9.
//
// Solidity: function ZERO_ADDRESS() constant returns(address)
func (_EscrowRecorder *EscrowRecorderCallerSession) ZEROADDRESS() (common.Address, error) {
	return _EscrowRecorder.Contract.ZEROADDRESS(&_EscrowRecorder.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_EscrowRecorder *EscrowRecorderCaller) Name(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _EscrowRecorder.contract.Call(opts, out, "name")
	return *ret0, err
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_EscrowRecorder *EscrowRecorderSession) Name() (string, error) {
	return _EscrowRecorder.Contract.Name(&_EscrowRecorder.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_EscrowRecorder *EscrowRecorderCallerSession) Name() (string, error) {
	return _EscrowRecorder.Contract.Name(&_EscrowRecorder.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_EscrowRecorder *EscrowRecorderCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _EscrowRecorder.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_EscrowRecorder *EscrowRecorderSession) Owner() (common.Address, error) {
	return _EscrowRecorder.Contract.Owner(&_EscrowRecorder.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_EscrowRecorder *EscrowRecorderCallerSession) Owner() (common.Address, error) {
	return _EscrowRecorder.Contract.Owner(&_EscrowRecorder.CallOpts)
}

// Test is a free data retrieval call binding the contract method 0x0b6dedfd.
//
// Solidity: function test(compliance address, token address, initiator address, from address, to address, tokens uint256) constant returns(s string)
func (_EscrowRecorder *EscrowRecorderCaller) Test(opts *bind.CallOpts, compliance common.Address, token common.Address, initiator common.Address, from common.Address, to common.Address, tokens *big.Int) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _EscrowRecorder.contract.Call(opts, out, "test", compliance, token, initiator, from, to, tokens)
	return *ret0, err
}

// Test is a free data retrieval call binding the contract method 0x0b6dedfd.
//
// Solidity: function test(compliance address, token address, initiator address, from address, to address, tokens uint256) constant returns(s string)
func (_EscrowRecorder *EscrowRecorderSession) Test(compliance common.Address, token common.Address, initiator common.Address, from common.Address, to common.Address, tokens *big.Int) (string, error) {
	return _EscrowRecorder.Contract.Test(&_EscrowRecorder.CallOpts, compliance, token, initiator, from, to, tokens)
}

// Test is a free data retrieval call binding the contract method 0x0b6dedfd.
//
// Solidity: function test(compliance address, token address, initiator address, from address, to address, tokens uint256) constant returns(s string)
func (_EscrowRecorder *EscrowRecorderCallerSession) Test(compliance common.Address, token common.Address, initiator common.Address, from common.Address, to common.Address, tokens *big.Int) (string, error) {
	return _EscrowRecorder.Contract.Test(&_EscrowRecorder.CallOpts, compliance, token, initiator, from, to, tokens)
}

// Check is a paid mutator transaction binding the contract method 0x803fcd43.
//
// Solidity: function check(token address, initiator address, from address, to address, tokens uint256) returns()
func (_EscrowRecorder *EscrowRecorderTransactor) Check(opts *bind.TransactOpts, token common.Address, initiator common.Address, from common.Address, to common.Address, tokens *big.Int) (*types.Transaction, error) {
	return _EscrowRecorder.contract.Transact(opts, "check", token, initiator, from, to, tokens)
}

// Check is a paid mutator transaction binding the contract method 0x803fcd43.
//
// Solidity: function check(token address, initiator address, from address, to address, tokens uint256) returns()
func (_EscrowRecorder *EscrowRecorderSession) Check(token common.Address, initiator common.Address, from common.Address, to common.Address, tokens *big.Int) (*types.Transaction, error) {
	return _EscrowRecorder.Contract.Check(&_EscrowRecorder.TransactOpts, token, initiator, from, to, tokens)
}

// Check is a paid mutator transaction binding the contract method 0x803fcd43.
//
// Solidity: function check(token address, initiator address, from address, to address, tokens uint256) returns()
func (_EscrowRecorder *EscrowRecorderTransactorSession) Check(token common.Address, initiator common.Address, from common.Address, to common.Address, tokens *big.Int) (*types.Transaction, error) {
	return _EscrowRecorder.Contract.Check(&_EscrowRecorder.TransactOpts, token, initiator, from, to, tokens)
}

// Kill is a paid mutator transaction binding the contract method 0x41c0e1b5.
//
// Solidity: function kill() returns()
func (_EscrowRecorder *EscrowRecorderTransactor) Kill(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EscrowRecorder.contract.Transact(opts, "kill")
}

// Kill is a paid mutator transaction binding the contract method 0x41c0e1b5.
//
// Solidity: function kill() returns()
func (_EscrowRecorder *EscrowRecorderSession) Kill() (*types.Transaction, error) {
	return _EscrowRecorder.Contract.Kill(&_EscrowRecorder.TransactOpts)
}

// Kill is a paid mutator transaction binding the contract method 0x41c0e1b5.
//
// Solidity: function kill() returns()
func (_EscrowRecorder *EscrowRecorderTransactorSession) Kill() (*types.Transaction, error) {
	return _EscrowRecorder.Contract.Kill(&_EscrowRecorder.TransactOpts)
}

// TransferOwner is a paid mutator transaction binding the contract method 0x4fb2e45d.
//
// Solidity: function transferOwner(newOwner address) returns()
func (_EscrowRecorder *EscrowRecorderTransactor) TransferOwner(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _EscrowRecorder.contract.Transact(opts, "transferOwner", newOwner)
}

// TransferOwner is a paid mutator transaction binding the contract method 0x4fb2e45d.
//
// Solidity: function transferOwner(newOwner address) returns()
func (_EscrowRecorder *EscrowRecorderSession) TransferOwner(newOwner common.Address) (*types.Transaction, error) {
	return _EscrowRecorder.Contract.TransferOwner(&_EscrowRecorder.TransactOpts, newOwner)
}

// TransferOwner is a paid mutator transaction binding the contract method 0x4fb2e45d.
//
// Solidity: function transferOwner(newOwner address) returns()
func (_EscrowRecorder *EscrowRecorderTransactorSession) TransferOwner(newOwner common.Address) (*types.Transaction, error) {
	return _EscrowRecorder.Contract.TransferOwner(&_EscrowRecorder.TransactOpts, newOwner)
}

// EscrowRecorderOwnerTransferredIterator is returned from FilterOwnerTransferred and is used to iterate over the raw logs and unpacked data for OwnerTransferred events raised by the EscrowRecorder contract.
type EscrowRecorderOwnerTransferredIterator struct {
	Event *EscrowRecorderOwnerTransferred // Event containing the contract specifics and raw log

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
func (it *EscrowRecorderOwnerTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EscrowRecorderOwnerTransferred)
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
		it.Event = new(EscrowRecorderOwnerTransferred)
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
func (it *EscrowRecorderOwnerTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EscrowRecorderOwnerTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EscrowRecorderOwnerTransferred represents a OwnerTransferred event raised by the EscrowRecorder contract.
type EscrowRecorderOwnerTransferred struct {
	OldOwner common.Address
	NewOwner common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOwnerTransferred is a free log retrieval operation binding the contract event 0x8934ce4adea8d9ce0d714d2c22b86790e41b7731c84b926fbbdc1d40ff6533c9.
//
// Solidity: e OwnerTransferred(oldOwner indexed address, newOwner indexed address)
func (_EscrowRecorder *EscrowRecorderFilterer) FilterOwnerTransferred(opts *bind.FilterOpts, oldOwner []common.Address, newOwner []common.Address) (*EscrowRecorderOwnerTransferredIterator, error) {

	var oldOwnerRule []interface{}
	for _, oldOwnerItem := range oldOwner {
		oldOwnerRule = append(oldOwnerRule, oldOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _EscrowRecorder.contract.FilterLogs(opts, "OwnerTransferred", oldOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &EscrowRecorderOwnerTransferredIterator{contract: _EscrowRecorder.contract, event: "OwnerTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnerTransferred is a free log subscription operation binding the contract event 0x8934ce4adea8d9ce0d714d2c22b86790e41b7731c84b926fbbdc1d40ff6533c9.
//
// Solidity: e OwnerTransferred(oldOwner indexed address, newOwner indexed address)
func (_EscrowRecorder *EscrowRecorderFilterer) WatchOwnerTransferred(opts *bind.WatchOpts, sink chan<- *EscrowRecorderOwnerTransferred, oldOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var oldOwnerRule []interface{}
	for _, oldOwnerItem := range oldOwner {
		oldOwnerRule = append(oldOwnerRule, oldOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _EscrowRecorder.contract.WatchLogs(opts, "OwnerTransferred", oldOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EscrowRecorderOwnerTransferred)
				if err := _EscrowRecorder.contract.UnpackLog(event, "OwnerTransferred", log); err != nil {
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
