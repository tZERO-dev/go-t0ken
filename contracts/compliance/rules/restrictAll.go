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

// RestrictAllABI is the input ABI used to generate the binding from.
const RestrictAllABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"compliance\",\"type\":\"address\"},{\"name\":\"token\",\"type\":\"address\"},{\"name\":\"initiator\",\"type\":\"address\"},{\"name\":\"from\",\"type\":\"address\"},{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"tokens\",\"type\":\"uint256\"}],\"name\":\"test\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"kill\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwner\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ZERO_ADDRESS\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"token\",\"type\":\"address\"},{\"name\":\"initiator\",\"type\":\"address\"},{\"name\":\"from\",\"type\":\"address\"},{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"tokens\",\"type\":\"uint256\"}],\"name\":\"check\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"oldOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnerTransferred\",\"type\":\"event\"}]"

// RestrictAllBin is the compiled bytecode used for deploying new contracts.
const RestrictAllBin = `600180546001600160a01b031916905560c0604052600c60808190527f526573747269637420416c6c000000000000000000000000000000000000000060a090815261004e9160029190610066565b50600080546001600160a01b03191633179055610101565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106100a757805160ff19168380011785556100d4565b828001600101855582156100d4579182015b828111156100d45782518255916020019190600101906100b9565b506100e09291506100e4565b5090565b6100fe91905b808211156100e057600081556001016100ea565b90565b610585806101106000396000f3fe608060405234801561001057600080fd5b506004361061007d5760003560e01c80634fb2e45d1161005b5780634fb2e45d14610155578063538ba4f91461017b578063803fcd431461019f5780638da5cb5b146101e55761007d565b806306fdde03146100825780630b6dedfd146100ff57806341c0e1b51461014b575b600080fd5b61008a6101ed565b6040805160208082528351818301528351919283929083019185019080838360005b838110156100c45781810151838201526020016100ac565b50505050905090810190601f1680156100f15780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b61008a600480360360c081101561011557600080fd5b506001600160a01b0381358116916020810135821691604082013581169160608101358216916080820135169060a00135610278565b61015361029e565b005b6101536004803603602081101561016b57600080fd5b50356001600160a01b0316610326565b6101836104b0565b604080516001600160a01b039092168252519081900360200190f35b610153600480360360a08110156101b557600080fd5b506001600160a01b03813581169160208101358216916040820135811691606081013590911690608001356104bf565b6101836104f6565b6002805460408051602060018416156101000260001901909316849004601f810184900484028201840190925281815292918301828280156102705780601f1061024557610100808354040283529160200191610270565b820191906000526020600020905b81548152906001019060200180831161025357829003601f168201915b505050505081565b606060405180606001604052806026815260200161050660269139979650505050505050565b6000546001600160a01b03163314806102c757506001546000546001600160a01b039081169116145b610318576040805162461bcd60e51b815260206004820152601960248201527f4f776e6572206163636f756e7420697320726571756972656400000000000000604482015290519081900360640190fd5b6000546001600160a01b0316ff5b6000546001600160a01b031633148061034f57506001546000546001600160a01b039081169116145b6103a0576040805162461bcd60e51b815260206004820152601960248201527f4f776e6572206163636f756e7420697320726571756972656400000000000000604482015290519081900360640190fd5b6000546001600160a01b03828116911614156103ed5760405162461bcd60e51b815260040180806020018281038252602581526020018061052c6025913960400191505060405180910390fd5b6001600160a01b038116610448576040805162461bcd60e51b815260206004820181905260248201527f4e6577204f776e65722063616e6e6f74206265207a65726f2061646472657373604482015290519081900360640190fd5b600080546001600160a01b038381167fffffffffffffffffffffffff0000000000000000000000000000000000000000831681178455604051919092169283917f8934ce4adea8d9ce0d714d2c22b86790e41b7731c84b926fbbdc1d40ff6533c99190a35050565b6001546001600160a01b031681565b60405162461bcd60e51b81526004018080602001828103825260268152602001806105066026913960400191505060405180910390fd5b6000546001600160a01b03168156fe416c6c207472616e7366657273206172652063757272656e746c7920726573747269637465644e6577204f776e65722063616e6e6f74206265207468652063757272656e74206f776e6572a265627a7a72305820e7d1e9985d33a9787c016fd84b29705538c5c30f67b13f0cc9679a205c439c6264736f6c634300050a0032`

// DeployRestrictAll deploys a new Ethereum contract, binding an instance of RestrictAll to it.
func DeployRestrictAll(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *RestrictAll, error) {
	parsed, err := abi.JSON(strings.NewReader(RestrictAllABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(RestrictAllBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &RestrictAll{RestrictAllCaller: RestrictAllCaller{contract: contract}, RestrictAllTransactor: RestrictAllTransactor{contract: contract}, RestrictAllFilterer: RestrictAllFilterer{contract: contract}}, nil
}

// RestrictAll is an auto generated Go binding around an Ethereum contract.
type RestrictAll struct {
	RestrictAllCaller     // Read-only binding to the contract
	RestrictAllTransactor // Write-only binding to the contract
	RestrictAllFilterer   // Log filterer for contract events
}

// RestrictAllCaller is an auto generated read-only Go binding around an Ethereum contract.
type RestrictAllCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RestrictAllTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RestrictAllTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RestrictAllFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RestrictAllFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RestrictAllSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RestrictAllSession struct {
	Contract     *RestrictAll      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RestrictAllCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RestrictAllCallerSession struct {
	Contract *RestrictAllCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// RestrictAllTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RestrictAllTransactorSession struct {
	Contract     *RestrictAllTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// RestrictAllRaw is an auto generated low-level Go binding around an Ethereum contract.
type RestrictAllRaw struct {
	Contract *RestrictAll // Generic contract binding to access the raw methods on
}

// RestrictAllCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RestrictAllCallerRaw struct {
	Contract *RestrictAllCaller // Generic read-only contract binding to access the raw methods on
}

// RestrictAllTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RestrictAllTransactorRaw struct {
	Contract *RestrictAllTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRestrictAll creates a new instance of RestrictAll, bound to a specific deployed contract.
func NewRestrictAll(address common.Address, backend bind.ContractBackend) (*RestrictAll, error) {
	contract, err := bindRestrictAll(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &RestrictAll{RestrictAllCaller: RestrictAllCaller{contract: contract}, RestrictAllTransactor: RestrictAllTransactor{contract: contract}, RestrictAllFilterer: RestrictAllFilterer{contract: contract}}, nil
}

// NewRestrictAllCaller creates a new read-only instance of RestrictAll, bound to a specific deployed contract.
func NewRestrictAllCaller(address common.Address, caller bind.ContractCaller) (*RestrictAllCaller, error) {
	contract, err := bindRestrictAll(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RestrictAllCaller{contract: contract}, nil
}

// NewRestrictAllTransactor creates a new write-only instance of RestrictAll, bound to a specific deployed contract.
func NewRestrictAllTransactor(address common.Address, transactor bind.ContractTransactor) (*RestrictAllTransactor, error) {
	contract, err := bindRestrictAll(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RestrictAllTransactor{contract: contract}, nil
}

// NewRestrictAllFilterer creates a new log filterer instance of RestrictAll, bound to a specific deployed contract.
func NewRestrictAllFilterer(address common.Address, filterer bind.ContractFilterer) (*RestrictAllFilterer, error) {
	contract, err := bindRestrictAll(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RestrictAllFilterer{contract: contract}, nil
}

// bindRestrictAll binds a generic wrapper to an already deployed contract.
func bindRestrictAll(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(RestrictAllABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RestrictAll *RestrictAllRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _RestrictAll.Contract.RestrictAllCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RestrictAll *RestrictAllRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RestrictAll.Contract.RestrictAllTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RestrictAll *RestrictAllRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RestrictAll.Contract.RestrictAllTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RestrictAll *RestrictAllCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _RestrictAll.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RestrictAll *RestrictAllTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RestrictAll.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RestrictAll *RestrictAllTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RestrictAll.Contract.contract.Transact(opts, method, params...)
}

// ZEROADDRESS is a free data retrieval call binding the contract method 0x538ba4f9.
//
// Solidity: function ZERO_ADDRESS() constant returns(address)
func (_RestrictAll *RestrictAllCaller) ZEROADDRESS(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _RestrictAll.contract.Call(opts, out, "ZERO_ADDRESS")
	return *ret0, err
}

// ZEROADDRESS is a free data retrieval call binding the contract method 0x538ba4f9.
//
// Solidity: function ZERO_ADDRESS() constant returns(address)
func (_RestrictAll *RestrictAllSession) ZEROADDRESS() (common.Address, error) {
	return _RestrictAll.Contract.ZEROADDRESS(&_RestrictAll.CallOpts)
}

// ZEROADDRESS is a free data retrieval call binding the contract method 0x538ba4f9.
//
// Solidity: function ZERO_ADDRESS() constant returns(address)
func (_RestrictAll *RestrictAllCallerSession) ZEROADDRESS() (common.Address, error) {
	return _RestrictAll.Contract.ZEROADDRESS(&_RestrictAll.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_RestrictAll *RestrictAllCaller) Name(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _RestrictAll.contract.Call(opts, out, "name")
	return *ret0, err
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_RestrictAll *RestrictAllSession) Name() (string, error) {
	return _RestrictAll.Contract.Name(&_RestrictAll.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_RestrictAll *RestrictAllCallerSession) Name() (string, error) {
	return _RestrictAll.Contract.Name(&_RestrictAll.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_RestrictAll *RestrictAllCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _RestrictAll.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_RestrictAll *RestrictAllSession) Owner() (common.Address, error) {
	return _RestrictAll.Contract.Owner(&_RestrictAll.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_RestrictAll *RestrictAllCallerSession) Owner() (common.Address, error) {
	return _RestrictAll.Contract.Owner(&_RestrictAll.CallOpts)
}

// Test is a free data retrieval call binding the contract method 0x0b6dedfd.
//
// Solidity: function test(compliance address, token address, initiator address, from address, to address, tokens uint256) constant returns(string)
func (_RestrictAll *RestrictAllCaller) Test(opts *bind.CallOpts, compliance common.Address, token common.Address, initiator common.Address, from common.Address, to common.Address, tokens *big.Int) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _RestrictAll.contract.Call(opts, out, "test", compliance, token, initiator, from, to, tokens)
	return *ret0, err
}

// Test is a free data retrieval call binding the contract method 0x0b6dedfd.
//
// Solidity: function test(compliance address, token address, initiator address, from address, to address, tokens uint256) constant returns(string)
func (_RestrictAll *RestrictAllSession) Test(compliance common.Address, token common.Address, initiator common.Address, from common.Address, to common.Address, tokens *big.Int) (string, error) {
	return _RestrictAll.Contract.Test(&_RestrictAll.CallOpts, compliance, token, initiator, from, to, tokens)
}

// Test is a free data retrieval call binding the contract method 0x0b6dedfd.
//
// Solidity: function test(compliance address, token address, initiator address, from address, to address, tokens uint256) constant returns(string)
func (_RestrictAll *RestrictAllCallerSession) Test(compliance common.Address, token common.Address, initiator common.Address, from common.Address, to common.Address, tokens *big.Int) (string, error) {
	return _RestrictAll.Contract.Test(&_RestrictAll.CallOpts, compliance, token, initiator, from, to, tokens)
}

// Check is a paid mutator transaction binding the contract method 0x803fcd43.
//
// Solidity: function check(token address, initiator address, from address, to address, tokens uint256) returns()
func (_RestrictAll *RestrictAllTransactor) Check(opts *bind.TransactOpts, token common.Address, initiator common.Address, from common.Address, to common.Address, tokens *big.Int) (*types.Transaction, error) {
	return _RestrictAll.contract.Transact(opts, "check", token, initiator, from, to, tokens)
}

// Check is a paid mutator transaction binding the contract method 0x803fcd43.
//
// Solidity: function check(token address, initiator address, from address, to address, tokens uint256) returns()
func (_RestrictAll *RestrictAllSession) Check(token common.Address, initiator common.Address, from common.Address, to common.Address, tokens *big.Int) (*types.Transaction, error) {
	return _RestrictAll.Contract.Check(&_RestrictAll.TransactOpts, token, initiator, from, to, tokens)
}

// Check is a paid mutator transaction binding the contract method 0x803fcd43.
//
// Solidity: function check(token address, initiator address, from address, to address, tokens uint256) returns()
func (_RestrictAll *RestrictAllTransactorSession) Check(token common.Address, initiator common.Address, from common.Address, to common.Address, tokens *big.Int) (*types.Transaction, error) {
	return _RestrictAll.Contract.Check(&_RestrictAll.TransactOpts, token, initiator, from, to, tokens)
}

// Kill is a paid mutator transaction binding the contract method 0x41c0e1b5.
//
// Solidity: function kill() returns()
func (_RestrictAll *RestrictAllTransactor) Kill(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RestrictAll.contract.Transact(opts, "kill")
}

// Kill is a paid mutator transaction binding the contract method 0x41c0e1b5.
//
// Solidity: function kill() returns()
func (_RestrictAll *RestrictAllSession) Kill() (*types.Transaction, error) {
	return _RestrictAll.Contract.Kill(&_RestrictAll.TransactOpts)
}

// Kill is a paid mutator transaction binding the contract method 0x41c0e1b5.
//
// Solidity: function kill() returns()
func (_RestrictAll *RestrictAllTransactorSession) Kill() (*types.Transaction, error) {
	return _RestrictAll.Contract.Kill(&_RestrictAll.TransactOpts)
}

// TransferOwner is a paid mutator transaction binding the contract method 0x4fb2e45d.
//
// Solidity: function transferOwner(newOwner address) returns()
func (_RestrictAll *RestrictAllTransactor) TransferOwner(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _RestrictAll.contract.Transact(opts, "transferOwner", newOwner)
}

// TransferOwner is a paid mutator transaction binding the contract method 0x4fb2e45d.
//
// Solidity: function transferOwner(newOwner address) returns()
func (_RestrictAll *RestrictAllSession) TransferOwner(newOwner common.Address) (*types.Transaction, error) {
	return _RestrictAll.Contract.TransferOwner(&_RestrictAll.TransactOpts, newOwner)
}

// TransferOwner is a paid mutator transaction binding the contract method 0x4fb2e45d.
//
// Solidity: function transferOwner(newOwner address) returns()
func (_RestrictAll *RestrictAllTransactorSession) TransferOwner(newOwner common.Address) (*types.Transaction, error) {
	return _RestrictAll.Contract.TransferOwner(&_RestrictAll.TransactOpts, newOwner)
}

// RestrictAllOwnerTransferredIterator is returned from FilterOwnerTransferred and is used to iterate over the raw logs and unpacked data for OwnerTransferred events raised by the RestrictAll contract.
type RestrictAllOwnerTransferredIterator struct {
	Event *RestrictAllOwnerTransferred // Event containing the contract specifics and raw log

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
func (it *RestrictAllOwnerTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RestrictAllOwnerTransferred)
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
		it.Event = new(RestrictAllOwnerTransferred)
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
func (it *RestrictAllOwnerTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RestrictAllOwnerTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RestrictAllOwnerTransferred represents a OwnerTransferred event raised by the RestrictAll contract.
type RestrictAllOwnerTransferred struct {
	OldOwner common.Address
	NewOwner common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOwnerTransferred is a free log retrieval operation binding the contract event 0x8934ce4adea8d9ce0d714d2c22b86790e41b7731c84b926fbbdc1d40ff6533c9.
//
// Solidity: e OwnerTransferred(oldOwner indexed address, newOwner indexed address)
func (_RestrictAll *RestrictAllFilterer) FilterOwnerTransferred(opts *bind.FilterOpts, oldOwner []common.Address, newOwner []common.Address) (*RestrictAllOwnerTransferredIterator, error) {

	var oldOwnerRule []interface{}
	for _, oldOwnerItem := range oldOwner {
		oldOwnerRule = append(oldOwnerRule, oldOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _RestrictAll.contract.FilterLogs(opts, "OwnerTransferred", oldOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &RestrictAllOwnerTransferredIterator{contract: _RestrictAll.contract, event: "OwnerTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnerTransferred is a free log subscription operation binding the contract event 0x8934ce4adea8d9ce0d714d2c22b86790e41b7731c84b926fbbdc1d40ff6533c9.
//
// Solidity: e OwnerTransferred(oldOwner indexed address, newOwner indexed address)
func (_RestrictAll *RestrictAllFilterer) WatchOwnerTransferred(opts *bind.WatchOpts, sink chan<- *RestrictAllOwnerTransferred, oldOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var oldOwnerRule []interface{}
	for _, oldOwnerItem := range oldOwner {
		oldOwnerRule = append(oldOwnerRule, oldOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _RestrictAll.contract.WatchLogs(opts, "OwnerTransferred", oldOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RestrictAllOwnerTransferred)
				if err := _RestrictAll.contract.UnpackLog(event, "OwnerTransferred", log); err != nil {
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
