// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package lifecycle

import (
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// DestroyableABI is the input ABI used to generate the binding from.
const DestroyableABI = "[{\"constant\":false,\"inputs\":[],\"name\":\"kill\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwner\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"oldOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnerTransferred\",\"type\":\"event\"}]"

// DestroyableBin is the compiled bytecode used for deploying new contracts.
const DestroyableBin = `608060405260008054600160a060020a031916331790556103d1806100256000396000f3fe608060405234801561001057600080fd5b506004361061005d577c0100000000000000000000000000000000000000000000000000000000600035046341c0e1b581146100625780634fb2e45d1461006c5780638da5cb5b1461009f575b600080fd5b61006a6100d0565b005b61006a6004803603602081101561008257600080fd5b503573ffffffffffffffffffffffffffffffffffffffff16610171565b6100a7610364565b6040805173ffffffffffffffffffffffffffffffffffffffff9092168252519081900360200190f35b60005473ffffffffffffffffffffffffffffffffffffffff16331461015657604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601960248201527f4f776e6572206163636f756e7420697320726571756972656400000000000000604482015290519081900360640190fd5b60005473ffffffffffffffffffffffffffffffffffffffff16ff5b60005473ffffffffffffffffffffffffffffffffffffffff1633146101f757604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601960248201527f4f776e6572206163636f756e7420697320726571756972656400000000000000604482015290519081900360640190fd5b60005473ffffffffffffffffffffffffffffffffffffffff8281169116141561026b576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260258152602001806103816025913960400191505060405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff811615156102ef57604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820181905260248201527f4e6577204f776e65722063616e6e6f74206265207a65726f2061646472657373604482015290519081900360640190fd5b6000805473ffffffffffffffffffffffffffffffffffffffff8381167fffffffffffffffffffffffff0000000000000000000000000000000000000000831681178455604051919092169283917f8934ce4adea8d9ce0d714d2c22b86790e41b7731c84b926fbbdc1d40ff6533c99190a35050565b60005473ffffffffffffffffffffffffffffffffffffffff168156fe4e6577204f776e65722063616e6e6f74206265207468652063757272656e74206f776e6572a165627a7a723058205d8d6c552ea55224cad69f2718de646dc4bdf1178ad2024bd019c2930d790ee00029`

// DeployDestroyable deploys a new Ethereum contract, binding an instance of Destroyable to it.
func DeployDestroyable(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Destroyable, error) {
	parsed, err := abi.JSON(strings.NewReader(DestroyableABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(DestroyableBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Destroyable{DestroyableCaller: DestroyableCaller{contract: contract}, DestroyableTransactor: DestroyableTransactor{contract: contract}, DestroyableFilterer: DestroyableFilterer{contract: contract}}, nil
}

// Destroyable is an auto generated Go binding around an Ethereum contract.
type Destroyable struct {
	DestroyableCaller     // Read-only binding to the contract
	DestroyableTransactor // Write-only binding to the contract
	DestroyableFilterer   // Log filterer for contract events
}

// DestroyableCaller is an auto generated read-only Go binding around an Ethereum contract.
type DestroyableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DestroyableTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DestroyableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DestroyableFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DestroyableFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DestroyableSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DestroyableSession struct {
	Contract     *Destroyable      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DestroyableCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DestroyableCallerSession struct {
	Contract *DestroyableCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// DestroyableTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DestroyableTransactorSession struct {
	Contract     *DestroyableTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// DestroyableRaw is an auto generated low-level Go binding around an Ethereum contract.
type DestroyableRaw struct {
	Contract *Destroyable // Generic contract binding to access the raw methods on
}

// DestroyableCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DestroyableCallerRaw struct {
	Contract *DestroyableCaller // Generic read-only contract binding to access the raw methods on
}

// DestroyableTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DestroyableTransactorRaw struct {
	Contract *DestroyableTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDestroyable creates a new instance of Destroyable, bound to a specific deployed contract.
func NewDestroyable(address common.Address, backend bind.ContractBackend) (*Destroyable, error) {
	contract, err := bindDestroyable(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Destroyable{DestroyableCaller: DestroyableCaller{contract: contract}, DestroyableTransactor: DestroyableTransactor{contract: contract}, DestroyableFilterer: DestroyableFilterer{contract: contract}}, nil
}

// NewDestroyableCaller creates a new read-only instance of Destroyable, bound to a specific deployed contract.
func NewDestroyableCaller(address common.Address, caller bind.ContractCaller) (*DestroyableCaller, error) {
	contract, err := bindDestroyable(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DestroyableCaller{contract: contract}, nil
}

// NewDestroyableTransactor creates a new write-only instance of Destroyable, bound to a specific deployed contract.
func NewDestroyableTransactor(address common.Address, transactor bind.ContractTransactor) (*DestroyableTransactor, error) {
	contract, err := bindDestroyable(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DestroyableTransactor{contract: contract}, nil
}

// NewDestroyableFilterer creates a new log filterer instance of Destroyable, bound to a specific deployed contract.
func NewDestroyableFilterer(address common.Address, filterer bind.ContractFilterer) (*DestroyableFilterer, error) {
	contract, err := bindDestroyable(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DestroyableFilterer{contract: contract}, nil
}

// bindDestroyable binds a generic wrapper to an already deployed contract.
func bindDestroyable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DestroyableABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Destroyable *DestroyableRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Destroyable.Contract.DestroyableCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Destroyable *DestroyableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Destroyable.Contract.DestroyableTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Destroyable *DestroyableRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Destroyable.Contract.DestroyableTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Destroyable *DestroyableCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Destroyable.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Destroyable *DestroyableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Destroyable.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Destroyable *DestroyableTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Destroyable.Contract.contract.Transact(opts, method, params...)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Destroyable *DestroyableCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Destroyable.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Destroyable *DestroyableSession) Owner() (common.Address, error) {
	return _Destroyable.Contract.Owner(&_Destroyable.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Destroyable *DestroyableCallerSession) Owner() (common.Address, error) {
	return _Destroyable.Contract.Owner(&_Destroyable.CallOpts)
}

// Kill is a paid mutator transaction binding the contract method 0x41c0e1b5.
//
// Solidity: function kill() returns()
func (_Destroyable *DestroyableTransactor) Kill(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Destroyable.contract.Transact(opts, "kill")
}

// Kill is a paid mutator transaction binding the contract method 0x41c0e1b5.
//
// Solidity: function kill() returns()
func (_Destroyable *DestroyableSession) Kill() (*types.Transaction, error) {
	return _Destroyable.Contract.Kill(&_Destroyable.TransactOpts)
}

// Kill is a paid mutator transaction binding the contract method 0x41c0e1b5.
//
// Solidity: function kill() returns()
func (_Destroyable *DestroyableTransactorSession) Kill() (*types.Transaction, error) {
	return _Destroyable.Contract.Kill(&_Destroyable.TransactOpts)
}

// TransferOwner is a paid mutator transaction binding the contract method 0x4fb2e45d.
//
// Solidity: function transferOwner(newOwner address) returns()
func (_Destroyable *DestroyableTransactor) TransferOwner(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Destroyable.contract.Transact(opts, "transferOwner", newOwner)
}

// TransferOwner is a paid mutator transaction binding the contract method 0x4fb2e45d.
//
// Solidity: function transferOwner(newOwner address) returns()
func (_Destroyable *DestroyableSession) TransferOwner(newOwner common.Address) (*types.Transaction, error) {
	return _Destroyable.Contract.TransferOwner(&_Destroyable.TransactOpts, newOwner)
}

// TransferOwner is a paid mutator transaction binding the contract method 0x4fb2e45d.
//
// Solidity: function transferOwner(newOwner address) returns()
func (_Destroyable *DestroyableTransactorSession) TransferOwner(newOwner common.Address) (*types.Transaction, error) {
	return _Destroyable.Contract.TransferOwner(&_Destroyable.TransactOpts, newOwner)
}

// DestroyableOwnerTransferredIterator is returned from FilterOwnerTransferred and is used to iterate over the raw logs and unpacked data for OwnerTransferred events raised by the Destroyable contract.
type DestroyableOwnerTransferredIterator struct {
	Event *DestroyableOwnerTransferred // Event containing the contract specifics and raw log

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
func (it *DestroyableOwnerTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DestroyableOwnerTransferred)
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
		it.Event = new(DestroyableOwnerTransferred)
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
func (it *DestroyableOwnerTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DestroyableOwnerTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DestroyableOwnerTransferred represents a OwnerTransferred event raised by the Destroyable contract.
type DestroyableOwnerTransferred struct {
	OldOwner common.Address
	NewOwner common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOwnerTransferred is a free log retrieval operation binding the contract event 0x8934ce4adea8d9ce0d714d2c22b86790e41b7731c84b926fbbdc1d40ff6533c9.
//
// Solidity: e OwnerTransferred(oldOwner indexed address, newOwner indexed address)
func (_Destroyable *DestroyableFilterer) FilterOwnerTransferred(opts *bind.FilterOpts, oldOwner []common.Address, newOwner []common.Address) (*DestroyableOwnerTransferredIterator, error) {

	var oldOwnerRule []interface{}
	for _, oldOwnerItem := range oldOwner {
		oldOwnerRule = append(oldOwnerRule, oldOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Destroyable.contract.FilterLogs(opts, "OwnerTransferred", oldOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &DestroyableOwnerTransferredIterator{contract: _Destroyable.contract, event: "OwnerTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnerTransferred is a free log subscription operation binding the contract event 0x8934ce4adea8d9ce0d714d2c22b86790e41b7731c84b926fbbdc1d40ff6533c9.
//
// Solidity: e OwnerTransferred(oldOwner indexed address, newOwner indexed address)
func (_Destroyable *DestroyableFilterer) WatchOwnerTransferred(opts *bind.WatchOpts, sink chan<- *DestroyableOwnerTransferred, oldOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var oldOwnerRule []interface{}
	for _, oldOwnerItem := range oldOwner {
		oldOwnerRule = append(oldOwnerRule, oldOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Destroyable.contract.WatchLogs(opts, "OwnerTransferred", oldOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DestroyableOwnerTransferred)
				if err := _Destroyable.contract.UnpackLog(event, "OwnerTransferred", log); err != nil {
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
