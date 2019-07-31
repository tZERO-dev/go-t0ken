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

// RestrictFromInvestorABI is the input ABI used to generate the binding from.
const RestrictFromInvestorABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"kill\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwner\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ZERO_ADDRESS\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"token\",\"type\":\"address\"},{\"name\":\"initiator\",\"type\":\"address\"},{\"name\":\"from\",\"type\":\"address\"},{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"toKind\",\"type\":\"uint8\"},{\"name\":\"tokens\",\"type\":\"uint256\"}],\"name\":\"check\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"compliance\",\"type\":\"address\"},{\"name\":\"token\",\"type\":\"address\"},{\"name\":\"initiator\",\"type\":\"address\"},{\"name\":\"from\",\"type\":\"address\"},{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"toKind\",\"type\":\"uint8\"},{\"name\":\"tokens\",\"type\":\"uint256\"}],\"name\":\"test\",\"outputs\":[{\"name\":\"s\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"oldOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnerTransferred\",\"type\":\"event\"}]"

// RestrictFromInvestorBin is the compiled bytecode used for deploying new contracts.
const RestrictFromInvestorBin = `600180546001600160a81b031916740400000000000000000000000000000000000000001760ff60a81b1916750500000000000000000000000000000000000000000017905560c0604052601660808190527f52657374726963742046726f6d20496e766573746f720000000000000000000060a0908152610084916002919061009c565b50600080546001600160a01b03191633179055610137565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106100dd57805160ff191683800117855561010a565b8280016001018555821561010a579182015b8281111561010a5782518255916020019190600101906100ef565b5061011692915061011a565b5090565b61013491905b808211156101165760008155600101610120565b90565b61062b806101466000396000f3fe608060405234801561001057600080fd5b506004361061007d5760003560e01c8063538ba4f91161005b578063538ba4f91461012f5780635a47e1c7146101535780638da5cb5b146101a2578063a9189562146101aa5761007d565b806306fdde031461008257806341c0e1b5146100ff5780634fb2e45d14610109575b600080fd5b61008a6101ff565b6040805160208082528351818301528351919283929083019185019080838360005b838110156100c45781810151838201526020016100ac565b50505050905090810190601f1680156100f15780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b61010761028a565b005b6101076004803603602081101561011f57600080fd5b50356001600160a01b0316610312565b61013761049c565b604080516001600160a01b039092168252519081900360200190f35b610107600480360360c081101561016957600080fd5b506001600160a01b0381358116916020810135821691604082013581169160608101359091169060ff6080820135169060a001356104ab565b61013761051c565b61008a600480360360e08110156101c057600080fd5b506001600160a01b0381358116916020810135821691604082013581169160608101358216916080820135169060ff60a0820135169060c0013561052b565b6002805460408051602060018416156101000260001901909316849004601f810184900484028201840190925281815292918301828280156102825780601f1061025757610100808354040283529160200191610282565b820191906000526020600020905b81548152906001019060200180831161026557829003601f168201915b505050505081565b6000546001600160a01b03163314806102b357506001546000546001600160a01b039081169116145b610304576040805162461bcd60e51b815260206004820152601960248201527f4f776e6572206163636f756e7420697320726571756972656400000000000000604482015290519081900360640190fd5b6000546001600160a01b0316ff5b6000546001600160a01b031633148061033b57506001546000546001600160a01b039081169116145b61038c576040805162461bcd60e51b815260206004820152601960248201527f4f776e6572206163636f756e7420697320726571756972656400000000000000604482015290519081900360640190fd5b6000546001600160a01b03828116911614156103d95760405162461bcd60e51b81526004018080602001828103825260258152602001806105886025913960400191505060405180910390fd5b6001600160a01b038116610434576040805162461bcd60e51b815260206004820181905260248201527f4e6577204f776e65722063616e6e6f74206265207a65726f2061646472657373604482015290519081900360640190fd5b600080546001600160a01b038381167fffffffffffffffffffffffff0000000000000000000000000000000000000000831681178455604051919092169283917f8934ce4adea8d9ce0d714d2c22b86790e41b7731c84b926fbbdc1d40ff6533c99190a35050565b6001546001600160a01b031681565b60015460ff838116600160a01b90920416148015906104d9575060015460ff838116600160a81b9092041614155b6105145760405162461bcd60e51b81526004018080602001828103825260248152602001806105ad6024913960400191505060405180910390fd5b505050505050565b6000546001600160a01b031681565b60015460609060ff848116600160a01b909204161480159061055c575060015460ff848116600160a81b9092041614155b61057c576040518060600160405280602681526020016105d16026913990505b97965050505050505056fe4e6577204f776e65722063616e6e6f74206265207468652063757272656e74206f776e657254686520746f20616464726573732063616e6e6f7420626520616e20696e766573746f725468652027746f2720616464726573732063616e6e6f7420626520616e20696e766573746f72a265627a7a72305820d1085418f283395406c519c101b45f83c7014b85181f1036fa053f09acbf4ffe64736f6c634300050a0032`

// DeployRestrictFromInvestor deploys a new Ethereum contract, binding an instance of RestrictFromInvestor to it.
func DeployRestrictFromInvestor(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *RestrictFromInvestor, error) {
	parsed, err := abi.JSON(strings.NewReader(RestrictFromInvestorABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(RestrictFromInvestorBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &RestrictFromInvestor{RestrictFromInvestorCaller: RestrictFromInvestorCaller{contract: contract}, RestrictFromInvestorTransactor: RestrictFromInvestorTransactor{contract: contract}, RestrictFromInvestorFilterer: RestrictFromInvestorFilterer{contract: contract}}, nil
}

// RestrictFromInvestor is an auto generated Go binding around an Ethereum contract.
type RestrictFromInvestor struct {
	RestrictFromInvestorCaller     // Read-only binding to the contract
	RestrictFromInvestorTransactor // Write-only binding to the contract
	RestrictFromInvestorFilterer   // Log filterer for contract events
}

// RestrictFromInvestorCaller is an auto generated read-only Go binding around an Ethereum contract.
type RestrictFromInvestorCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RestrictFromInvestorTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RestrictFromInvestorTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RestrictFromInvestorFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RestrictFromInvestorFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RestrictFromInvestorSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RestrictFromInvestorSession struct {
	Contract     *RestrictFromInvestor // Generic contract binding to set the session for
	CallOpts     bind.CallOpts         // Call options to use throughout this session
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// RestrictFromInvestorCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RestrictFromInvestorCallerSession struct {
	Contract *RestrictFromInvestorCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts               // Call options to use throughout this session
}

// RestrictFromInvestorTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RestrictFromInvestorTransactorSession struct {
	Contract     *RestrictFromInvestorTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts               // Transaction auth options to use throughout this session
}

// RestrictFromInvestorRaw is an auto generated low-level Go binding around an Ethereum contract.
type RestrictFromInvestorRaw struct {
	Contract *RestrictFromInvestor // Generic contract binding to access the raw methods on
}

// RestrictFromInvestorCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RestrictFromInvestorCallerRaw struct {
	Contract *RestrictFromInvestorCaller // Generic read-only contract binding to access the raw methods on
}

// RestrictFromInvestorTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RestrictFromInvestorTransactorRaw struct {
	Contract *RestrictFromInvestorTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRestrictFromInvestor creates a new instance of RestrictFromInvestor, bound to a specific deployed contract.
func NewRestrictFromInvestor(address common.Address, backend bind.ContractBackend) (*RestrictFromInvestor, error) {
	contract, err := bindRestrictFromInvestor(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &RestrictFromInvestor{RestrictFromInvestorCaller: RestrictFromInvestorCaller{contract: contract}, RestrictFromInvestorTransactor: RestrictFromInvestorTransactor{contract: contract}, RestrictFromInvestorFilterer: RestrictFromInvestorFilterer{contract: contract}}, nil
}

// NewRestrictFromInvestorCaller creates a new read-only instance of RestrictFromInvestor, bound to a specific deployed contract.
func NewRestrictFromInvestorCaller(address common.Address, caller bind.ContractCaller) (*RestrictFromInvestorCaller, error) {
	contract, err := bindRestrictFromInvestor(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RestrictFromInvestorCaller{contract: contract}, nil
}

// NewRestrictFromInvestorTransactor creates a new write-only instance of RestrictFromInvestor, bound to a specific deployed contract.
func NewRestrictFromInvestorTransactor(address common.Address, transactor bind.ContractTransactor) (*RestrictFromInvestorTransactor, error) {
	contract, err := bindRestrictFromInvestor(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RestrictFromInvestorTransactor{contract: contract}, nil
}

// NewRestrictFromInvestorFilterer creates a new log filterer instance of RestrictFromInvestor, bound to a specific deployed contract.
func NewRestrictFromInvestorFilterer(address common.Address, filterer bind.ContractFilterer) (*RestrictFromInvestorFilterer, error) {
	contract, err := bindRestrictFromInvestor(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RestrictFromInvestorFilterer{contract: contract}, nil
}

// bindRestrictFromInvestor binds a generic wrapper to an already deployed contract.
func bindRestrictFromInvestor(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(RestrictFromInvestorABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RestrictFromInvestor *RestrictFromInvestorRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _RestrictFromInvestor.Contract.RestrictFromInvestorCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RestrictFromInvestor *RestrictFromInvestorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RestrictFromInvestor.Contract.RestrictFromInvestorTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RestrictFromInvestor *RestrictFromInvestorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RestrictFromInvestor.Contract.RestrictFromInvestorTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RestrictFromInvestor *RestrictFromInvestorCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _RestrictFromInvestor.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RestrictFromInvestor *RestrictFromInvestorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RestrictFromInvestor.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RestrictFromInvestor *RestrictFromInvestorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RestrictFromInvestor.Contract.contract.Transact(opts, method, params...)
}

// ZEROADDRESS is a free data retrieval call binding the contract method 0x538ba4f9.
//
// Solidity: function ZERO_ADDRESS() constant returns(address)
func (_RestrictFromInvestor *RestrictFromInvestorCaller) ZEROADDRESS(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _RestrictFromInvestor.contract.Call(opts, out, "ZERO_ADDRESS")
	return *ret0, err
}

// ZEROADDRESS is a free data retrieval call binding the contract method 0x538ba4f9.
//
// Solidity: function ZERO_ADDRESS() constant returns(address)
func (_RestrictFromInvestor *RestrictFromInvestorSession) ZEROADDRESS() (common.Address, error) {
	return _RestrictFromInvestor.Contract.ZEROADDRESS(&_RestrictFromInvestor.CallOpts)
}

// ZEROADDRESS is a free data retrieval call binding the contract method 0x538ba4f9.
//
// Solidity: function ZERO_ADDRESS() constant returns(address)
func (_RestrictFromInvestor *RestrictFromInvestorCallerSession) ZEROADDRESS() (common.Address, error) {
	return _RestrictFromInvestor.Contract.ZEROADDRESS(&_RestrictFromInvestor.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_RestrictFromInvestor *RestrictFromInvestorCaller) Name(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _RestrictFromInvestor.contract.Call(opts, out, "name")
	return *ret0, err
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_RestrictFromInvestor *RestrictFromInvestorSession) Name() (string, error) {
	return _RestrictFromInvestor.Contract.Name(&_RestrictFromInvestor.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_RestrictFromInvestor *RestrictFromInvestorCallerSession) Name() (string, error) {
	return _RestrictFromInvestor.Contract.Name(&_RestrictFromInvestor.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_RestrictFromInvestor *RestrictFromInvestorCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _RestrictFromInvestor.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_RestrictFromInvestor *RestrictFromInvestorSession) Owner() (common.Address, error) {
	return _RestrictFromInvestor.Contract.Owner(&_RestrictFromInvestor.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_RestrictFromInvestor *RestrictFromInvestorCallerSession) Owner() (common.Address, error) {
	return _RestrictFromInvestor.Contract.Owner(&_RestrictFromInvestor.CallOpts)
}

// Test is a free data retrieval call binding the contract method 0xa9189562.
//
// Solidity: function test(compliance address, token address, initiator address, from address, to address, toKind uint8, tokens uint256) constant returns(s string)
func (_RestrictFromInvestor *RestrictFromInvestorCaller) Test(opts *bind.CallOpts, compliance common.Address, token common.Address, initiator common.Address, from common.Address, to common.Address, toKind uint8, tokens *big.Int) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _RestrictFromInvestor.contract.Call(opts, out, "test", compliance, token, initiator, from, to, toKind, tokens)
	return *ret0, err
}

// Test is a free data retrieval call binding the contract method 0xa9189562.
//
// Solidity: function test(compliance address, token address, initiator address, from address, to address, toKind uint8, tokens uint256) constant returns(s string)
func (_RestrictFromInvestor *RestrictFromInvestorSession) Test(compliance common.Address, token common.Address, initiator common.Address, from common.Address, to common.Address, toKind uint8, tokens *big.Int) (string, error) {
	return _RestrictFromInvestor.Contract.Test(&_RestrictFromInvestor.CallOpts, compliance, token, initiator, from, to, toKind, tokens)
}

// Test is a free data retrieval call binding the contract method 0xa9189562.
//
// Solidity: function test(compliance address, token address, initiator address, from address, to address, toKind uint8, tokens uint256) constant returns(s string)
func (_RestrictFromInvestor *RestrictFromInvestorCallerSession) Test(compliance common.Address, token common.Address, initiator common.Address, from common.Address, to common.Address, toKind uint8, tokens *big.Int) (string, error) {
	return _RestrictFromInvestor.Contract.Test(&_RestrictFromInvestor.CallOpts, compliance, token, initiator, from, to, toKind, tokens)
}

// Check is a paid mutator transaction binding the contract method 0x5a47e1c7.
//
// Solidity: function check(token address, initiator address, from address, to address, toKind uint8, tokens uint256) returns()
func (_RestrictFromInvestor *RestrictFromInvestorTransactor) Check(opts *bind.TransactOpts, token common.Address, initiator common.Address, from common.Address, to common.Address, toKind uint8, tokens *big.Int) (*types.Transaction, error) {
	return _RestrictFromInvestor.contract.Transact(opts, "check", token, initiator, from, to, toKind, tokens)
}

// Check is a paid mutator transaction binding the contract method 0x5a47e1c7.
//
// Solidity: function check(token address, initiator address, from address, to address, toKind uint8, tokens uint256) returns()
func (_RestrictFromInvestor *RestrictFromInvestorSession) Check(token common.Address, initiator common.Address, from common.Address, to common.Address, toKind uint8, tokens *big.Int) (*types.Transaction, error) {
	return _RestrictFromInvestor.Contract.Check(&_RestrictFromInvestor.TransactOpts, token, initiator, from, to, toKind, tokens)
}

// Check is a paid mutator transaction binding the contract method 0x5a47e1c7.
//
// Solidity: function check(token address, initiator address, from address, to address, toKind uint8, tokens uint256) returns()
func (_RestrictFromInvestor *RestrictFromInvestorTransactorSession) Check(token common.Address, initiator common.Address, from common.Address, to common.Address, toKind uint8, tokens *big.Int) (*types.Transaction, error) {
	return _RestrictFromInvestor.Contract.Check(&_RestrictFromInvestor.TransactOpts, token, initiator, from, to, toKind, tokens)
}

// Kill is a paid mutator transaction binding the contract method 0x41c0e1b5.
//
// Solidity: function kill() returns()
func (_RestrictFromInvestor *RestrictFromInvestorTransactor) Kill(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RestrictFromInvestor.contract.Transact(opts, "kill")
}

// Kill is a paid mutator transaction binding the contract method 0x41c0e1b5.
//
// Solidity: function kill() returns()
func (_RestrictFromInvestor *RestrictFromInvestorSession) Kill() (*types.Transaction, error) {
	return _RestrictFromInvestor.Contract.Kill(&_RestrictFromInvestor.TransactOpts)
}

// Kill is a paid mutator transaction binding the contract method 0x41c0e1b5.
//
// Solidity: function kill() returns()
func (_RestrictFromInvestor *RestrictFromInvestorTransactorSession) Kill() (*types.Transaction, error) {
	return _RestrictFromInvestor.Contract.Kill(&_RestrictFromInvestor.TransactOpts)
}

// TransferOwner is a paid mutator transaction binding the contract method 0x4fb2e45d.
//
// Solidity: function transferOwner(newOwner address) returns()
func (_RestrictFromInvestor *RestrictFromInvestorTransactor) TransferOwner(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _RestrictFromInvestor.contract.Transact(opts, "transferOwner", newOwner)
}

// TransferOwner is a paid mutator transaction binding the contract method 0x4fb2e45d.
//
// Solidity: function transferOwner(newOwner address) returns()
func (_RestrictFromInvestor *RestrictFromInvestorSession) TransferOwner(newOwner common.Address) (*types.Transaction, error) {
	return _RestrictFromInvestor.Contract.TransferOwner(&_RestrictFromInvestor.TransactOpts, newOwner)
}

// TransferOwner is a paid mutator transaction binding the contract method 0x4fb2e45d.
//
// Solidity: function transferOwner(newOwner address) returns()
func (_RestrictFromInvestor *RestrictFromInvestorTransactorSession) TransferOwner(newOwner common.Address) (*types.Transaction, error) {
	return _RestrictFromInvestor.Contract.TransferOwner(&_RestrictFromInvestor.TransactOpts, newOwner)
}

// RestrictFromInvestorOwnerTransferredIterator is returned from FilterOwnerTransferred and is used to iterate over the raw logs and unpacked data for OwnerTransferred events raised by the RestrictFromInvestor contract.
type RestrictFromInvestorOwnerTransferredIterator struct {
	Event *RestrictFromInvestorOwnerTransferred // Event containing the contract specifics and raw log

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
func (it *RestrictFromInvestorOwnerTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RestrictFromInvestorOwnerTransferred)
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
		it.Event = new(RestrictFromInvestorOwnerTransferred)
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
func (it *RestrictFromInvestorOwnerTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RestrictFromInvestorOwnerTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RestrictFromInvestorOwnerTransferred represents a OwnerTransferred event raised by the RestrictFromInvestor contract.
type RestrictFromInvestorOwnerTransferred struct {
	OldOwner common.Address
	NewOwner common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOwnerTransferred is a free log retrieval operation binding the contract event 0x8934ce4adea8d9ce0d714d2c22b86790e41b7731c84b926fbbdc1d40ff6533c9.
//
// Solidity: e OwnerTransferred(oldOwner indexed address, newOwner indexed address)
func (_RestrictFromInvestor *RestrictFromInvestorFilterer) FilterOwnerTransferred(opts *bind.FilterOpts, oldOwner []common.Address, newOwner []common.Address) (*RestrictFromInvestorOwnerTransferredIterator, error) {

	var oldOwnerRule []interface{}
	for _, oldOwnerItem := range oldOwner {
		oldOwnerRule = append(oldOwnerRule, oldOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _RestrictFromInvestor.contract.FilterLogs(opts, "OwnerTransferred", oldOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &RestrictFromInvestorOwnerTransferredIterator{contract: _RestrictFromInvestor.contract, event: "OwnerTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnerTransferred is a free log subscription operation binding the contract event 0x8934ce4adea8d9ce0d714d2c22b86790e41b7731c84b926fbbdc1d40ff6533c9.
//
// Solidity: e OwnerTransferred(oldOwner indexed address, newOwner indexed address)
func (_RestrictFromInvestor *RestrictFromInvestorFilterer) WatchOwnerTransferred(opts *bind.WatchOpts, sink chan<- *RestrictFromInvestorOwnerTransferred, oldOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var oldOwnerRule []interface{}
	for _, oldOwnerItem := range oldOwner {
		oldOwnerRule = append(oldOwnerRule, oldOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _RestrictFromInvestor.contract.WatchLogs(opts, "OwnerTransferred", oldOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RestrictFromInvestorOwnerTransferred)
				if err := _RestrictFromInvestor.contract.UnpackLog(event, "OwnerTransferred", log); err != nil {
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
