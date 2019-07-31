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

// LockableDestroyableABI is the input ABI used to generate the binding from.
const LockableDestroyableABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"locked\",\"type\":\"bool\"}],\"name\":\"setLocked\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"kill\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwner\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ZERO_ADDRESS\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isLocked\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"oldOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnerTransferred\",\"type\":\"event\"}]"

// LockableDestroyableBin is the compiled bytecode used for deploying new contracts.
const LockableDestroyableBin = `608060405260018054600080546001600160a01b031916331790556001600160a81b03191690556104d3806100356000396000f3fe608060405234801561001057600080fd5b50600436106100725760003560e01c8063538ba4f911610050578063538ba4f9146100c65780638da5cb5b146100ea578063a4e2d634146100f257610072565b8063211e28b61461007757806341c0e1b5146100985780634fb2e45d146100a0575b600080fd5b6100966004803603602081101561008d57600080fd5b5035151561010e565b005b610096610211565b610096600480360360208110156100b657600080fd5b50356001600160a01b0316610299565b6100ce610423565b604080516001600160a01b039092168252519081900360200190f35b6100ce610432565b6100fa610441565b604080519115158252519081900360200190f35b6000546001600160a01b031633148061013757506001546000546001600160a01b039081169116145b610188576040805162461bcd60e51b815260206004820152601960248201527f4f776e6572206163636f756e7420697320726571756972656400000000000000604482015290519081900360640190fd5b60015460ff600160a01b90910416151581151514156101d85760405162461bcd60e51b81526004018080602001828103825260288152602001806104776028913960400191505060405180910390fd5b60018054911515600160a01b027fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff909216919091179055565b6000546001600160a01b031633148061023a57506001546000546001600160a01b039081169116145b61028b576040805162461bcd60e51b815260206004820152601960248201527f4f776e6572206163636f756e7420697320726571756972656400000000000000604482015290519081900360640190fd5b6000546001600160a01b0316ff5b6000546001600160a01b03163314806102c257506001546000546001600160a01b039081169116145b610313576040805162461bcd60e51b815260206004820152601960248201527f4f776e6572206163636f756e7420697320726571756972656400000000000000604482015290519081900360640190fd5b6000546001600160a01b03828116911614156103605760405162461bcd60e51b81526004018080602001828103825260258152602001806104526025913960400191505060405180910390fd5b6001600160a01b0381166103bb576040805162461bcd60e51b815260206004820181905260248201527f4e6577204f776e65722063616e6e6f74206265207a65726f2061646472657373604482015290519081900360640190fd5b600080546001600160a01b038381167fffffffffffffffffffffffff0000000000000000000000000000000000000000831681178455604051919092169283917f8934ce4adea8d9ce0d714d2c22b86790e41b7731c84b926fbbdc1d40ff6533c99190a35050565b6001546001600160a01b031681565b6000546001600160a01b031681565b600154600160a01b900460ff168156fe4e6577204f776e65722063616e6e6f74206265207468652063757272656e74206f776e6572436f6e747261637420616c726561647920696e20726571756573746564206c6f636b207374617465a265627a7a723058204bcdd4056a1b345614d795b43a66ce3f5d4f922f63e0b82c00f18cf6cfd71d6864736f6c634300050a0032`

// DeployLockableDestroyable deploys a new Ethereum contract, binding an instance of LockableDestroyable to it.
func DeployLockableDestroyable(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *LockableDestroyable, error) {
	parsed, err := abi.JSON(strings.NewReader(LockableDestroyableABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(LockableDestroyableBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &LockableDestroyable{LockableDestroyableCaller: LockableDestroyableCaller{contract: contract}, LockableDestroyableTransactor: LockableDestroyableTransactor{contract: contract}, LockableDestroyableFilterer: LockableDestroyableFilterer{contract: contract}}, nil
}

// LockableDestroyable is an auto generated Go binding around an Ethereum contract.
type LockableDestroyable struct {
	LockableDestroyableCaller     // Read-only binding to the contract
	LockableDestroyableTransactor // Write-only binding to the contract
	LockableDestroyableFilterer   // Log filterer for contract events
}

// LockableDestroyableCaller is an auto generated read-only Go binding around an Ethereum contract.
type LockableDestroyableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LockableDestroyableTransactor is an auto generated write-only Go binding around an Ethereum contract.
type LockableDestroyableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LockableDestroyableFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type LockableDestroyableFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LockableDestroyableSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type LockableDestroyableSession struct {
	Contract     *LockableDestroyable // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// LockableDestroyableCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type LockableDestroyableCallerSession struct {
	Contract *LockableDestroyableCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// LockableDestroyableTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type LockableDestroyableTransactorSession struct {
	Contract     *LockableDestroyableTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// LockableDestroyableRaw is an auto generated low-level Go binding around an Ethereum contract.
type LockableDestroyableRaw struct {
	Contract *LockableDestroyable // Generic contract binding to access the raw methods on
}

// LockableDestroyableCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type LockableDestroyableCallerRaw struct {
	Contract *LockableDestroyableCaller // Generic read-only contract binding to access the raw methods on
}

// LockableDestroyableTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type LockableDestroyableTransactorRaw struct {
	Contract *LockableDestroyableTransactor // Generic write-only contract binding to access the raw methods on
}

// NewLockableDestroyable creates a new instance of LockableDestroyable, bound to a specific deployed contract.
func NewLockableDestroyable(address common.Address, backend bind.ContractBackend) (*LockableDestroyable, error) {
	contract, err := bindLockableDestroyable(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &LockableDestroyable{LockableDestroyableCaller: LockableDestroyableCaller{contract: contract}, LockableDestroyableTransactor: LockableDestroyableTransactor{contract: contract}, LockableDestroyableFilterer: LockableDestroyableFilterer{contract: contract}}, nil
}

// NewLockableDestroyableCaller creates a new read-only instance of LockableDestroyable, bound to a specific deployed contract.
func NewLockableDestroyableCaller(address common.Address, caller bind.ContractCaller) (*LockableDestroyableCaller, error) {
	contract, err := bindLockableDestroyable(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &LockableDestroyableCaller{contract: contract}, nil
}

// NewLockableDestroyableTransactor creates a new write-only instance of LockableDestroyable, bound to a specific deployed contract.
func NewLockableDestroyableTransactor(address common.Address, transactor bind.ContractTransactor) (*LockableDestroyableTransactor, error) {
	contract, err := bindLockableDestroyable(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &LockableDestroyableTransactor{contract: contract}, nil
}

// NewLockableDestroyableFilterer creates a new log filterer instance of LockableDestroyable, bound to a specific deployed contract.
func NewLockableDestroyableFilterer(address common.Address, filterer bind.ContractFilterer) (*LockableDestroyableFilterer, error) {
	contract, err := bindLockableDestroyable(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &LockableDestroyableFilterer{contract: contract}, nil
}

// bindLockableDestroyable binds a generic wrapper to an already deployed contract.
func bindLockableDestroyable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(LockableDestroyableABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LockableDestroyable *LockableDestroyableRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _LockableDestroyable.Contract.LockableDestroyableCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LockableDestroyable *LockableDestroyableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LockableDestroyable.Contract.LockableDestroyableTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LockableDestroyable *LockableDestroyableRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LockableDestroyable.Contract.LockableDestroyableTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LockableDestroyable *LockableDestroyableCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _LockableDestroyable.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LockableDestroyable *LockableDestroyableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LockableDestroyable.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LockableDestroyable *LockableDestroyableTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LockableDestroyable.Contract.contract.Transact(opts, method, params...)
}

// ZEROADDRESS is a free data retrieval call binding the contract method 0x538ba4f9.
//
// Solidity: function ZERO_ADDRESS() constant returns(address)
func (_LockableDestroyable *LockableDestroyableCaller) ZEROADDRESS(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _LockableDestroyable.contract.Call(opts, out, "ZERO_ADDRESS")
	return *ret0, err
}

// ZEROADDRESS is a free data retrieval call binding the contract method 0x538ba4f9.
//
// Solidity: function ZERO_ADDRESS() constant returns(address)
func (_LockableDestroyable *LockableDestroyableSession) ZEROADDRESS() (common.Address, error) {
	return _LockableDestroyable.Contract.ZEROADDRESS(&_LockableDestroyable.CallOpts)
}

// ZEROADDRESS is a free data retrieval call binding the contract method 0x538ba4f9.
//
// Solidity: function ZERO_ADDRESS() constant returns(address)
func (_LockableDestroyable *LockableDestroyableCallerSession) ZEROADDRESS() (common.Address, error) {
	return _LockableDestroyable.Contract.ZEROADDRESS(&_LockableDestroyable.CallOpts)
}

// IsLocked is a free data retrieval call binding the contract method 0xa4e2d634.
//
// Solidity: function isLocked() constant returns(bool)
func (_LockableDestroyable *LockableDestroyableCaller) IsLocked(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _LockableDestroyable.contract.Call(opts, out, "isLocked")
	return *ret0, err
}

// IsLocked is a free data retrieval call binding the contract method 0xa4e2d634.
//
// Solidity: function isLocked() constant returns(bool)
func (_LockableDestroyable *LockableDestroyableSession) IsLocked() (bool, error) {
	return _LockableDestroyable.Contract.IsLocked(&_LockableDestroyable.CallOpts)
}

// IsLocked is a free data retrieval call binding the contract method 0xa4e2d634.
//
// Solidity: function isLocked() constant returns(bool)
func (_LockableDestroyable *LockableDestroyableCallerSession) IsLocked() (bool, error) {
	return _LockableDestroyable.Contract.IsLocked(&_LockableDestroyable.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_LockableDestroyable *LockableDestroyableCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _LockableDestroyable.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_LockableDestroyable *LockableDestroyableSession) Owner() (common.Address, error) {
	return _LockableDestroyable.Contract.Owner(&_LockableDestroyable.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_LockableDestroyable *LockableDestroyableCallerSession) Owner() (common.Address, error) {
	return _LockableDestroyable.Contract.Owner(&_LockableDestroyable.CallOpts)
}

// Kill is a paid mutator transaction binding the contract method 0x41c0e1b5.
//
// Solidity: function kill() returns()
func (_LockableDestroyable *LockableDestroyableTransactor) Kill(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LockableDestroyable.contract.Transact(opts, "kill")
}

// Kill is a paid mutator transaction binding the contract method 0x41c0e1b5.
//
// Solidity: function kill() returns()
func (_LockableDestroyable *LockableDestroyableSession) Kill() (*types.Transaction, error) {
	return _LockableDestroyable.Contract.Kill(&_LockableDestroyable.TransactOpts)
}

// Kill is a paid mutator transaction binding the contract method 0x41c0e1b5.
//
// Solidity: function kill() returns()
func (_LockableDestroyable *LockableDestroyableTransactorSession) Kill() (*types.Transaction, error) {
	return _LockableDestroyable.Contract.Kill(&_LockableDestroyable.TransactOpts)
}

// SetLocked is a paid mutator transaction binding the contract method 0x211e28b6.
//
// Solidity: function setLocked(locked bool) returns()
func (_LockableDestroyable *LockableDestroyableTransactor) SetLocked(opts *bind.TransactOpts, locked bool) (*types.Transaction, error) {
	return _LockableDestroyable.contract.Transact(opts, "setLocked", locked)
}

// SetLocked is a paid mutator transaction binding the contract method 0x211e28b6.
//
// Solidity: function setLocked(locked bool) returns()
func (_LockableDestroyable *LockableDestroyableSession) SetLocked(locked bool) (*types.Transaction, error) {
	return _LockableDestroyable.Contract.SetLocked(&_LockableDestroyable.TransactOpts, locked)
}

// SetLocked is a paid mutator transaction binding the contract method 0x211e28b6.
//
// Solidity: function setLocked(locked bool) returns()
func (_LockableDestroyable *LockableDestroyableTransactorSession) SetLocked(locked bool) (*types.Transaction, error) {
	return _LockableDestroyable.Contract.SetLocked(&_LockableDestroyable.TransactOpts, locked)
}

// TransferOwner is a paid mutator transaction binding the contract method 0x4fb2e45d.
//
// Solidity: function transferOwner(newOwner address) returns()
func (_LockableDestroyable *LockableDestroyableTransactor) TransferOwner(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _LockableDestroyable.contract.Transact(opts, "transferOwner", newOwner)
}

// TransferOwner is a paid mutator transaction binding the contract method 0x4fb2e45d.
//
// Solidity: function transferOwner(newOwner address) returns()
func (_LockableDestroyable *LockableDestroyableSession) TransferOwner(newOwner common.Address) (*types.Transaction, error) {
	return _LockableDestroyable.Contract.TransferOwner(&_LockableDestroyable.TransactOpts, newOwner)
}

// TransferOwner is a paid mutator transaction binding the contract method 0x4fb2e45d.
//
// Solidity: function transferOwner(newOwner address) returns()
func (_LockableDestroyable *LockableDestroyableTransactorSession) TransferOwner(newOwner common.Address) (*types.Transaction, error) {
	return _LockableDestroyable.Contract.TransferOwner(&_LockableDestroyable.TransactOpts, newOwner)
}

// LockableDestroyableOwnerTransferredIterator is returned from FilterOwnerTransferred and is used to iterate over the raw logs and unpacked data for OwnerTransferred events raised by the LockableDestroyable contract.
type LockableDestroyableOwnerTransferredIterator struct {
	Event *LockableDestroyableOwnerTransferred // Event containing the contract specifics and raw log

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
func (it *LockableDestroyableOwnerTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LockableDestroyableOwnerTransferred)
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
		it.Event = new(LockableDestroyableOwnerTransferred)
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
func (it *LockableDestroyableOwnerTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LockableDestroyableOwnerTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LockableDestroyableOwnerTransferred represents a OwnerTransferred event raised by the LockableDestroyable contract.
type LockableDestroyableOwnerTransferred struct {
	OldOwner common.Address
	NewOwner common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOwnerTransferred is a free log retrieval operation binding the contract event 0x8934ce4adea8d9ce0d714d2c22b86790e41b7731c84b926fbbdc1d40ff6533c9.
//
// Solidity: e OwnerTransferred(oldOwner indexed address, newOwner indexed address)
func (_LockableDestroyable *LockableDestroyableFilterer) FilterOwnerTransferred(opts *bind.FilterOpts, oldOwner []common.Address, newOwner []common.Address) (*LockableDestroyableOwnerTransferredIterator, error) {

	var oldOwnerRule []interface{}
	for _, oldOwnerItem := range oldOwner {
		oldOwnerRule = append(oldOwnerRule, oldOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _LockableDestroyable.contract.FilterLogs(opts, "OwnerTransferred", oldOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &LockableDestroyableOwnerTransferredIterator{contract: _LockableDestroyable.contract, event: "OwnerTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnerTransferred is a free log subscription operation binding the contract event 0x8934ce4adea8d9ce0d714d2c22b86790e41b7731c84b926fbbdc1d40ff6533c9.
//
// Solidity: e OwnerTransferred(oldOwner indexed address, newOwner indexed address)
func (_LockableDestroyable *LockableDestroyableFilterer) WatchOwnerTransferred(opts *bind.WatchOpts, sink chan<- *LockableDestroyableOwnerTransferred, oldOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var oldOwnerRule []interface{}
	for _, oldOwnerItem := range oldOwner {
		oldOwnerRule = append(oldOwnerRule, oldOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _LockableDestroyable.contract.WatchLogs(opts, "OwnerTransferred", oldOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LockableDestroyableOwnerTransferred)
				if err := _LockableDestroyable.contract.UnpackLog(event, "OwnerTransferred", log); err != nil {
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
