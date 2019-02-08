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
const RestrictFromInvestorABI = "[{\"constant\":false,\"inputs\":[],\"name\":\"kill\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwner\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"initiator\",\"type\":\"address\"},{\"name\":\"from\",\"type\":\"address\"},{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"toKind\",\"type\":\"uint8\"},{\"name\":\"tokens\",\"type\":\"uint256\"},{\"name\":\"store\",\"type\":\"address\"}],\"name\":\"check\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"oldOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnerTransferred\",\"type\":\"event\"}]"

// RestrictFromInvestorBin is the compiled bytecode used for deploying new contracts.
const RestrictFromInvestorBin = `6080604052600080543360a060020a60ff0219909116740400000000000000000000000000000000000000001760a860020a60ff021916750500000000000000000000000000000000000000000017600160a060020a03191617815561050a90819061006b90396000f3fe608060405234801561001057600080fd5b5060043610610068577c0100000000000000000000000000000000000000000000000000000000600035046341c0e1b5811461006d5780634fb2e45d146100775780638da5cb5b146100aa578063b762c76d146100db575b600080fd5b610075610135565b005b6100756004803603602081101561008d57600080fd5b503573ffffffffffffffffffffffffffffffffffffffff166101d6565b6100b26103c9565b6040805173ffffffffffffffffffffffffffffffffffffffff9092168252519081900360200190f35b610075600480360360c08110156100f157600080fd5b5073ffffffffffffffffffffffffffffffffffffffff81358116916020810135821691604082013581169160ff6060820135169160808201359160a00135166103e5565b60005473ffffffffffffffffffffffffffffffffffffffff1633146101bb57604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601960248201527f4f776e6572206163636f756e7420697320726571756972656400000000000000604482015290519081900360640190fd5b60005473ffffffffffffffffffffffffffffffffffffffff16ff5b60005473ffffffffffffffffffffffffffffffffffffffff16331461025c57604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601960248201527f4f776e6572206163636f756e7420697320726571756972656400000000000000604482015290519081900360640190fd5b60005473ffffffffffffffffffffffffffffffffffffffff828116911614156102d0576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260258152602001806104966025913960400191505060405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff8116151561035457604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820181905260248201527f4e6577204f776e65722063616e6e6f74206265207a65726f2061646472657373604482015290519081900360640190fd5b6000805473ffffffffffffffffffffffffffffffffffffffff8381167fffffffffffffffffffffffff0000000000000000000000000000000000000000831681178455604051919092169283917f8934ce4adea8d9ce0d714d2c22b86790e41b7731c84b926fbbdc1d40ff6533c99190a35050565b60005473ffffffffffffffffffffffffffffffffffffffff1681565b60005460ff848116740100000000000000000000000000000000000000009092041614801590610436575060005460ff84811675010000000000000000000000000000000000000000009092041614155b151561048d576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260248152602001806104bb6024913960400191505060405180910390fd5b50505050505056fe4e6577204f776e65722063616e6e6f74206265207468652063757272656e74206f776e657254686520746f20616464726573732063616e6e6f7420626520616e20696e766573746f72a165627a7a72305820150478c9fc9f8928c6d5fb201e1c7db35bb4f60ab3a5d3b471aa827ca997094c0029`

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

// Check is a paid mutator transaction binding the contract method 0xb762c76d.
//
// Solidity: function check(initiator address, from address, to address, toKind uint8, tokens uint256, store address) returns()
func (_RestrictFromInvestor *RestrictFromInvestorTransactor) Check(opts *bind.TransactOpts, initiator common.Address, from common.Address, to common.Address, toKind uint8, tokens *big.Int, store common.Address) (*types.Transaction, error) {
	return _RestrictFromInvestor.contract.Transact(opts, "check", initiator, from, to, toKind, tokens, store)
}

// Check is a paid mutator transaction binding the contract method 0xb762c76d.
//
// Solidity: function check(initiator address, from address, to address, toKind uint8, tokens uint256, store address) returns()
func (_RestrictFromInvestor *RestrictFromInvestorSession) Check(initiator common.Address, from common.Address, to common.Address, toKind uint8, tokens *big.Int, store common.Address) (*types.Transaction, error) {
	return _RestrictFromInvestor.Contract.Check(&_RestrictFromInvestor.TransactOpts, initiator, from, to, toKind, tokens, store)
}

// Check is a paid mutator transaction binding the contract method 0xb762c76d.
//
// Solidity: function check(initiator address, from address, to address, toKind uint8, tokens uint256, store address) returns()
func (_RestrictFromInvestor *RestrictFromInvestorTransactorSession) Check(initiator common.Address, from common.Address, to common.Address, toKind uint8, tokens *big.Int, store common.Address) (*types.Transaction, error) {
	return _RestrictFromInvestor.Contract.Check(&_RestrictFromInvestor.TransactOpts, initiator, from, to, toKind, tokens, store)
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
