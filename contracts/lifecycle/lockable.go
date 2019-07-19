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

// LockableABI is the input ABI used to generate the binding from.
const LockableABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"locked\",\"type\":\"bool\"}],\"name\":\"setLocked\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwner\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ZERO_ADDRESS\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isLocked\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"oldOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnerTransferred\",\"type\":\"event\"}]"

// LockableBin is the compiled bytecode used for deploying new contracts.
const LockableBin = `6080604052600180546001600160a01b031916905534801561002057600080fd5b50600080546001600160a01b031916331790556001805460ff60a01b191690556104388061004f6000396000f3fe608060405234801561001057600080fd5b50600436106100675760003560e01c8063538ba4f911610050578063538ba4f9146100b35780638da5cb5b146100d7578063a4e2d634146100df57610067565b8063211e28b61461006c5780634fb2e45d1461008d575b600080fd5b61008b6004803603602081101561008257600080fd5b503515156100fb565b005b61008b600480360360208110156100a357600080fd5b50356001600160a01b03166101fe565b6100bb610388565b604080516001600160a01b039092168252519081900360200190f35b6100bb610397565b6100e76103a6565b604080519115158252519081900360200190f35b6000546001600160a01b031633148061012457506001546000546001600160a01b039081169116145b610175576040805162461bcd60e51b815260206004820152601960248201527f4f776e6572206163636f756e7420697320726571756972656400000000000000604482015290519081900360640190fd5b60015460ff600160a01b90910416151581151514156101c55760405162461bcd60e51b81526004018080602001828103825260288152602001806103dc6028913960400191505060405180910390fd5b60018054911515600160a01b027fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff909216919091179055565b6000546001600160a01b031633148061022757506001546000546001600160a01b039081169116145b610278576040805162461bcd60e51b815260206004820152601960248201527f4f776e6572206163636f756e7420697320726571756972656400000000000000604482015290519081900360640190fd5b6000546001600160a01b03828116911614156102c55760405162461bcd60e51b81526004018080602001828103825260258152602001806103b76025913960400191505060405180910390fd5b6001600160a01b038116610320576040805162461bcd60e51b815260206004820181905260248201527f4e6577204f776e65722063616e6e6f74206265207a65726f2061646472657373604482015290519081900360640190fd5b600080546001600160a01b038381167fffffffffffffffffffffffff0000000000000000000000000000000000000000831681178455604051919092169283917f8934ce4adea8d9ce0d714d2c22b86790e41b7731c84b926fbbdc1d40ff6533c99190a35050565b6001546001600160a01b031681565b6000546001600160a01b031681565b600154600160a01b900460ff168156fe4e6577204f776e65722063616e6e6f74206265207468652063757272656e74206f776e6572436f6e747261637420616c726561647920696e20726571756573746564206c6f636b207374617465a265627a7a723058203bf8d23237d3c4bfb9c86dc1b4b2e9a9a2a4c5c2b1d4ae45cb1b291418bd8e0b64736f6c63430005090032`

// DeployLockable deploys a new Ethereum contract, binding an instance of Lockable to it.
func DeployLockable(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Lockable, error) {
	parsed, err := abi.JSON(strings.NewReader(LockableABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(LockableBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Lockable{LockableCaller: LockableCaller{contract: contract}, LockableTransactor: LockableTransactor{contract: contract}, LockableFilterer: LockableFilterer{contract: contract}}, nil
}

// Lockable is an auto generated Go binding around an Ethereum contract.
type Lockable struct {
	LockableCaller     // Read-only binding to the contract
	LockableTransactor // Write-only binding to the contract
	LockableFilterer   // Log filterer for contract events
}

// LockableCaller is an auto generated read-only Go binding around an Ethereum contract.
type LockableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LockableTransactor is an auto generated write-only Go binding around an Ethereum contract.
type LockableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LockableFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type LockableFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LockableSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type LockableSession struct {
	Contract     *Lockable         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// LockableCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type LockableCallerSession struct {
	Contract *LockableCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// LockableTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type LockableTransactorSession struct {
	Contract     *LockableTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// LockableRaw is an auto generated low-level Go binding around an Ethereum contract.
type LockableRaw struct {
	Contract *Lockable // Generic contract binding to access the raw methods on
}

// LockableCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type LockableCallerRaw struct {
	Contract *LockableCaller // Generic read-only contract binding to access the raw methods on
}

// LockableTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type LockableTransactorRaw struct {
	Contract *LockableTransactor // Generic write-only contract binding to access the raw methods on
}

// NewLockable creates a new instance of Lockable, bound to a specific deployed contract.
func NewLockable(address common.Address, backend bind.ContractBackend) (*Lockable, error) {
	contract, err := bindLockable(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Lockable{LockableCaller: LockableCaller{contract: contract}, LockableTransactor: LockableTransactor{contract: contract}, LockableFilterer: LockableFilterer{contract: contract}}, nil
}

// NewLockableCaller creates a new read-only instance of Lockable, bound to a specific deployed contract.
func NewLockableCaller(address common.Address, caller bind.ContractCaller) (*LockableCaller, error) {
	contract, err := bindLockable(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &LockableCaller{contract: contract}, nil
}

// NewLockableTransactor creates a new write-only instance of Lockable, bound to a specific deployed contract.
func NewLockableTransactor(address common.Address, transactor bind.ContractTransactor) (*LockableTransactor, error) {
	contract, err := bindLockable(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &LockableTransactor{contract: contract}, nil
}

// NewLockableFilterer creates a new log filterer instance of Lockable, bound to a specific deployed contract.
func NewLockableFilterer(address common.Address, filterer bind.ContractFilterer) (*LockableFilterer, error) {
	contract, err := bindLockable(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &LockableFilterer{contract: contract}, nil
}

// bindLockable binds a generic wrapper to an already deployed contract.
func bindLockable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(LockableABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Lockable *LockableRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Lockable.Contract.LockableCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Lockable *LockableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Lockable.Contract.LockableTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Lockable *LockableRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Lockable.Contract.LockableTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Lockable *LockableCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Lockable.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Lockable *LockableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Lockable.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Lockable *LockableTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Lockable.Contract.contract.Transact(opts, method, params...)
}

// ZEROADDRESS is a free data retrieval call binding the contract method 0x538ba4f9.
//
// Solidity: function ZERO_ADDRESS() constant returns(address)
func (_Lockable *LockableCaller) ZEROADDRESS(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Lockable.contract.Call(opts, out, "ZERO_ADDRESS")
	return *ret0, err
}

// ZEROADDRESS is a free data retrieval call binding the contract method 0x538ba4f9.
//
// Solidity: function ZERO_ADDRESS() constant returns(address)
func (_Lockable *LockableSession) ZEROADDRESS() (common.Address, error) {
	return _Lockable.Contract.ZEROADDRESS(&_Lockable.CallOpts)
}

// ZEROADDRESS is a free data retrieval call binding the contract method 0x538ba4f9.
//
// Solidity: function ZERO_ADDRESS() constant returns(address)
func (_Lockable *LockableCallerSession) ZEROADDRESS() (common.Address, error) {
	return _Lockable.Contract.ZEROADDRESS(&_Lockable.CallOpts)
}

// IsLocked is a free data retrieval call binding the contract method 0xa4e2d634.
//
// Solidity: function isLocked() constant returns(bool)
func (_Lockable *LockableCaller) IsLocked(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Lockable.contract.Call(opts, out, "isLocked")
	return *ret0, err
}

// IsLocked is a free data retrieval call binding the contract method 0xa4e2d634.
//
// Solidity: function isLocked() constant returns(bool)
func (_Lockable *LockableSession) IsLocked() (bool, error) {
	return _Lockable.Contract.IsLocked(&_Lockable.CallOpts)
}

// IsLocked is a free data retrieval call binding the contract method 0xa4e2d634.
//
// Solidity: function isLocked() constant returns(bool)
func (_Lockable *LockableCallerSession) IsLocked() (bool, error) {
	return _Lockable.Contract.IsLocked(&_Lockable.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Lockable *LockableCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Lockable.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Lockable *LockableSession) Owner() (common.Address, error) {
	return _Lockable.Contract.Owner(&_Lockable.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Lockable *LockableCallerSession) Owner() (common.Address, error) {
	return _Lockable.Contract.Owner(&_Lockable.CallOpts)
}

// SetLocked is a paid mutator transaction binding the contract method 0x211e28b6.
//
// Solidity: function setLocked(locked bool) returns()
func (_Lockable *LockableTransactor) SetLocked(opts *bind.TransactOpts, locked bool) (*types.Transaction, error) {
	return _Lockable.contract.Transact(opts, "setLocked", locked)
}

// SetLocked is a paid mutator transaction binding the contract method 0x211e28b6.
//
// Solidity: function setLocked(locked bool) returns()
func (_Lockable *LockableSession) SetLocked(locked bool) (*types.Transaction, error) {
	return _Lockable.Contract.SetLocked(&_Lockable.TransactOpts, locked)
}

// SetLocked is a paid mutator transaction binding the contract method 0x211e28b6.
//
// Solidity: function setLocked(locked bool) returns()
func (_Lockable *LockableTransactorSession) SetLocked(locked bool) (*types.Transaction, error) {
	return _Lockable.Contract.SetLocked(&_Lockable.TransactOpts, locked)
}

// TransferOwner is a paid mutator transaction binding the contract method 0x4fb2e45d.
//
// Solidity: function transferOwner(newOwner address) returns()
func (_Lockable *LockableTransactor) TransferOwner(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Lockable.contract.Transact(opts, "transferOwner", newOwner)
}

// TransferOwner is a paid mutator transaction binding the contract method 0x4fb2e45d.
//
// Solidity: function transferOwner(newOwner address) returns()
func (_Lockable *LockableSession) TransferOwner(newOwner common.Address) (*types.Transaction, error) {
	return _Lockable.Contract.TransferOwner(&_Lockable.TransactOpts, newOwner)
}

// TransferOwner is a paid mutator transaction binding the contract method 0x4fb2e45d.
//
// Solidity: function transferOwner(newOwner address) returns()
func (_Lockable *LockableTransactorSession) TransferOwner(newOwner common.Address) (*types.Transaction, error) {
	return _Lockable.Contract.TransferOwner(&_Lockable.TransactOpts, newOwner)
}

// LockableOwnerTransferredIterator is returned from FilterOwnerTransferred and is used to iterate over the raw logs and unpacked data for OwnerTransferred events raised by the Lockable contract.
type LockableOwnerTransferredIterator struct {
	Event *LockableOwnerTransferred // Event containing the contract specifics and raw log

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
func (it *LockableOwnerTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LockableOwnerTransferred)
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
		it.Event = new(LockableOwnerTransferred)
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
func (it *LockableOwnerTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LockableOwnerTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LockableOwnerTransferred represents a OwnerTransferred event raised by the Lockable contract.
type LockableOwnerTransferred struct {
	OldOwner common.Address
	NewOwner common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOwnerTransferred is a free log retrieval operation binding the contract event 0x8934ce4adea8d9ce0d714d2c22b86790e41b7731c84b926fbbdc1d40ff6533c9.
//
// Solidity: e OwnerTransferred(oldOwner indexed address, newOwner indexed address)
func (_Lockable *LockableFilterer) FilterOwnerTransferred(opts *bind.FilterOpts, oldOwner []common.Address, newOwner []common.Address) (*LockableOwnerTransferredIterator, error) {

	var oldOwnerRule []interface{}
	for _, oldOwnerItem := range oldOwner {
		oldOwnerRule = append(oldOwnerRule, oldOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Lockable.contract.FilterLogs(opts, "OwnerTransferred", oldOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &LockableOwnerTransferredIterator{contract: _Lockable.contract, event: "OwnerTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnerTransferred is a free log subscription operation binding the contract event 0x8934ce4adea8d9ce0d714d2c22b86790e41b7731c84b926fbbdc1d40ff6533c9.
//
// Solidity: e OwnerTransferred(oldOwner indexed address, newOwner indexed address)
func (_Lockable *LockableFilterer) WatchOwnerTransferred(opts *bind.WatchOpts, sink chan<- *LockableOwnerTransferred, oldOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var oldOwnerRule []interface{}
	for _, oldOwnerItem := range oldOwner {
		oldOwnerRule = append(oldOwnerRule, oldOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Lockable.contract.WatchLogs(opts, "OwnerTransferred", oldOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LockableOwnerTransferred)
				if err := _Lockable.contract.UnpackLog(event, "OwnerTransferred", log); err != nil {
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
