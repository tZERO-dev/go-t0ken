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
const DestroyableABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnerTransferred\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[],\"name\":\"ZERO_ADDRESS\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"kill\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwner\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// DestroyableBin is the compiled bytecode used for deploying new contracts.
const DestroyableBin = `6080604052600180546001600160a01b03199081169091556000805490911633179055610337806100316000396000f3fe608060405234801561001057600080fd5b506004361061004c5760003560e01c806341c0e1b5146100515780634fb2e45d1461005b578063538ba4f9146100815780638da5cb5b146100a5575b600080fd5b6100596100ad565b005b6100596004803603602081101561007157600080fd5b50356001600160a01b0316610135565b6100896102bf565b604080516001600160a01b039092168252519081900360200190f35b6100896102ce565b6000546001600160a01b03163314806100d657506001546000546001600160a01b039081169116145b610127576040805162461bcd60e51b815260206004820152601960248201527f4f776e6572206163636f756e7420697320726571756972656400000000000000604482015290519081900360640190fd5b6000546001600160a01b0316ff5b6000546001600160a01b031633148061015e57506001546000546001600160a01b039081169116145b6101af576040805162461bcd60e51b815260206004820152601960248201527f4f776e6572206163636f756e7420697320726571756972656400000000000000604482015290519081900360640190fd5b6000546001600160a01b03828116911614156101fc5760405162461bcd60e51b81526004018080602001828103825260258152602001806102de6025913960400191505060405180910390fd5b6001600160a01b038116610257576040805162461bcd60e51b815260206004820181905260248201527f4e6577204f776e65722063616e6e6f74206265207a65726f2061646472657373604482015290519081900360640190fd5b600080546001600160a01b038381167fffffffffffffffffffffffff0000000000000000000000000000000000000000831681178455604051919092169283917f8934ce4adea8d9ce0d714d2c22b86790e41b7731c84b926fbbdc1d40ff6533c99190a35050565b6001546001600160a01b031681565b6000546001600160a01b03168156fe4e6577204f776e65722063616e6e6f74206265207468652063757272656e74206f776e6572a265627a7a723158207c2e660c0a62d9d97298b291c71777f38bfc1fb6138ac42e174a8b7c46fdd95f64736f6c634300050c0032`

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

// ZEROADDRESS is a free data retrieval call binding the contract method 0x538ba4f9.
//
// Solidity: function ZERO_ADDRESS() constant returns(address)
func (_Destroyable *DestroyableCaller) ZEROADDRESS(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Destroyable.contract.Call(opts, out, "ZERO_ADDRESS")
	return *ret0, err
}

// ZEROADDRESS is a free data retrieval call binding the contract method 0x538ba4f9.
//
// Solidity: function ZERO_ADDRESS() constant returns(address)
func (_Destroyable *DestroyableSession) ZEROADDRESS() (common.Address, error) {
	return _Destroyable.Contract.ZEROADDRESS(&_Destroyable.CallOpts)
}

// ZEROADDRESS is a free data retrieval call binding the contract method 0x538ba4f9.
//
// Solidity: function ZERO_ADDRESS() constant returns(address)
func (_Destroyable *DestroyableCallerSession) ZEROADDRESS() (common.Address, error) {
	return _Destroyable.Contract.ZEROADDRESS(&_Destroyable.CallOpts)
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
