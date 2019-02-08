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

// RestrictToBrokerOrCustodialAccountABI is the input ABI used to generate the binding from.
const RestrictToBrokerOrCustodialAccountABI = "[{\"constant\":false,\"inputs\":[],\"name\":\"kill\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwner\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"initiator\",\"type\":\"address\"},{\"name\":\"from\",\"type\":\"address\"},{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"toKind\",\"type\":\"uint8\"},{\"name\":\"tokens\",\"type\":\"uint256\"},{\"name\":\"store\",\"type\":\"address\"}],\"name\":\"check\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"oldOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnerTransferred\",\"type\":\"event\"}]"

// RestrictToBrokerOrCustodialAccountBin is the compiled bytecode used for deploying new contracts.
const RestrictToBrokerOrCustodialAccountBin = `608060405260008054600160a060020a031916331790556104ec806100256000396000f3fe608060405234801561001057600080fd5b5060043610610068577c0100000000000000000000000000000000000000000000000000000000600035046341c0e1b5811461006d5780634fb2e45d146100775780638da5cb5b146100aa578063b762c76d146100db575b600080fd5b610075610135565b005b6100756004803603602081101561008d57600080fd5b503573ffffffffffffffffffffffffffffffffffffffff166101d6565b6100b26103c9565b6040805173ffffffffffffffffffffffffffffffffffffffff9092168252519081900360200190f35b610075600480360360c08110156100f157600080fd5b5073ffffffffffffffffffffffffffffffffffffffff81358116916020810135821691604082013581169160ff6060820135169160808201359160a00135166103e5565b60005473ffffffffffffffffffffffffffffffffffffffff1633146101bb57604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601960248201527f4f776e6572206163636f756e7420697320726571756972656400000000000000604482015290519081900360640190fd5b60005473ffffffffffffffffffffffffffffffffffffffff16ff5b60005473ffffffffffffffffffffffffffffffffffffffff16331461025c57604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601960248201527f4f776e6572206163636f756e7420697320726571756972656400000000000000604482015290519081900360640190fd5b60005473ffffffffffffffffffffffffffffffffffffffff828116911614156102d0576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252602581526020018061045a6025913960400191505060405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff8116151561035457604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820181905260248201527f4e6577204f776e65722063616e6e6f74206265207a65726f2061646472657373604482015290519081900360640190fd5b6000805473ffffffffffffffffffffffffffffffffffffffff8381167fffffffffffffffffffffffff0000000000000000000000000000000000000000831681178455604051919092169283917f8934ce4adea8d9ce0d714d2c22b86790e41b7731c84b926fbbdc1d40ff6533c99190a35050565b60005473ffffffffffffffffffffffffffffffffffffffff1681565b60ff8316600214806103fa575060ff83166003145b1515610451576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252604281526020018061047f6042913960600191505060405180910390fd5b50505050505056fe4e6577204f776e65722063616e6e6f74206265207468652063757272656e74206f776e657254686520746f2061646472657373206d75737420626520656974686572206120637573746f6469616c2d6163636f756e74206f722062726f6b65722d6465616c6572a165627a7a72305820b65c45eedb13776e5fc41da084b5c6eb7099637b1c4d61fe620cb422b35b38110029`

// DeployRestrictToBrokerOrCustodialAccount deploys a new Ethereum contract, binding an instance of RestrictToBrokerOrCustodialAccount to it.
func DeployRestrictToBrokerOrCustodialAccount(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *RestrictToBrokerOrCustodialAccount, error) {
	parsed, err := abi.JSON(strings.NewReader(RestrictToBrokerOrCustodialAccountABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(RestrictToBrokerOrCustodialAccountBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &RestrictToBrokerOrCustodialAccount{RestrictToBrokerOrCustodialAccountCaller: RestrictToBrokerOrCustodialAccountCaller{contract: contract}, RestrictToBrokerOrCustodialAccountTransactor: RestrictToBrokerOrCustodialAccountTransactor{contract: contract}, RestrictToBrokerOrCustodialAccountFilterer: RestrictToBrokerOrCustodialAccountFilterer{contract: contract}}, nil
}

// RestrictToBrokerOrCustodialAccount is an auto generated Go binding around an Ethereum contract.
type RestrictToBrokerOrCustodialAccount struct {
	RestrictToBrokerOrCustodialAccountCaller     // Read-only binding to the contract
	RestrictToBrokerOrCustodialAccountTransactor // Write-only binding to the contract
	RestrictToBrokerOrCustodialAccountFilterer   // Log filterer for contract events
}

// RestrictToBrokerOrCustodialAccountCaller is an auto generated read-only Go binding around an Ethereum contract.
type RestrictToBrokerOrCustodialAccountCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RestrictToBrokerOrCustodialAccountTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RestrictToBrokerOrCustodialAccountTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RestrictToBrokerOrCustodialAccountFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RestrictToBrokerOrCustodialAccountFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RestrictToBrokerOrCustodialAccountSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RestrictToBrokerOrCustodialAccountSession struct {
	Contract     *RestrictToBrokerOrCustodialAccount // Generic contract binding to set the session for
	CallOpts     bind.CallOpts                       // Call options to use throughout this session
	TransactOpts bind.TransactOpts                   // Transaction auth options to use throughout this session
}

// RestrictToBrokerOrCustodialAccountCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RestrictToBrokerOrCustodialAccountCallerSession struct {
	Contract *RestrictToBrokerOrCustodialAccountCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                             // Call options to use throughout this session
}

// RestrictToBrokerOrCustodialAccountTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RestrictToBrokerOrCustodialAccountTransactorSession struct {
	Contract     *RestrictToBrokerOrCustodialAccountTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                             // Transaction auth options to use throughout this session
}

// RestrictToBrokerOrCustodialAccountRaw is an auto generated low-level Go binding around an Ethereum contract.
type RestrictToBrokerOrCustodialAccountRaw struct {
	Contract *RestrictToBrokerOrCustodialAccount // Generic contract binding to access the raw methods on
}

// RestrictToBrokerOrCustodialAccountCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RestrictToBrokerOrCustodialAccountCallerRaw struct {
	Contract *RestrictToBrokerOrCustodialAccountCaller // Generic read-only contract binding to access the raw methods on
}

// RestrictToBrokerOrCustodialAccountTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RestrictToBrokerOrCustodialAccountTransactorRaw struct {
	Contract *RestrictToBrokerOrCustodialAccountTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRestrictToBrokerOrCustodialAccount creates a new instance of RestrictToBrokerOrCustodialAccount, bound to a specific deployed contract.
func NewRestrictToBrokerOrCustodialAccount(address common.Address, backend bind.ContractBackend) (*RestrictToBrokerOrCustodialAccount, error) {
	contract, err := bindRestrictToBrokerOrCustodialAccount(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &RestrictToBrokerOrCustodialAccount{RestrictToBrokerOrCustodialAccountCaller: RestrictToBrokerOrCustodialAccountCaller{contract: contract}, RestrictToBrokerOrCustodialAccountTransactor: RestrictToBrokerOrCustodialAccountTransactor{contract: contract}, RestrictToBrokerOrCustodialAccountFilterer: RestrictToBrokerOrCustodialAccountFilterer{contract: contract}}, nil
}

// NewRestrictToBrokerOrCustodialAccountCaller creates a new read-only instance of RestrictToBrokerOrCustodialAccount, bound to a specific deployed contract.
func NewRestrictToBrokerOrCustodialAccountCaller(address common.Address, caller bind.ContractCaller) (*RestrictToBrokerOrCustodialAccountCaller, error) {
	contract, err := bindRestrictToBrokerOrCustodialAccount(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RestrictToBrokerOrCustodialAccountCaller{contract: contract}, nil
}

// NewRestrictToBrokerOrCustodialAccountTransactor creates a new write-only instance of RestrictToBrokerOrCustodialAccount, bound to a specific deployed contract.
func NewRestrictToBrokerOrCustodialAccountTransactor(address common.Address, transactor bind.ContractTransactor) (*RestrictToBrokerOrCustodialAccountTransactor, error) {
	contract, err := bindRestrictToBrokerOrCustodialAccount(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RestrictToBrokerOrCustodialAccountTransactor{contract: contract}, nil
}

// NewRestrictToBrokerOrCustodialAccountFilterer creates a new log filterer instance of RestrictToBrokerOrCustodialAccount, bound to a specific deployed contract.
func NewRestrictToBrokerOrCustodialAccountFilterer(address common.Address, filterer bind.ContractFilterer) (*RestrictToBrokerOrCustodialAccountFilterer, error) {
	contract, err := bindRestrictToBrokerOrCustodialAccount(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RestrictToBrokerOrCustodialAccountFilterer{contract: contract}, nil
}

// bindRestrictToBrokerOrCustodialAccount binds a generic wrapper to an already deployed contract.
func bindRestrictToBrokerOrCustodialAccount(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(RestrictToBrokerOrCustodialAccountABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RestrictToBrokerOrCustodialAccount *RestrictToBrokerOrCustodialAccountRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _RestrictToBrokerOrCustodialAccount.Contract.RestrictToBrokerOrCustodialAccountCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RestrictToBrokerOrCustodialAccount *RestrictToBrokerOrCustodialAccountRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RestrictToBrokerOrCustodialAccount.Contract.RestrictToBrokerOrCustodialAccountTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RestrictToBrokerOrCustodialAccount *RestrictToBrokerOrCustodialAccountRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RestrictToBrokerOrCustodialAccount.Contract.RestrictToBrokerOrCustodialAccountTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RestrictToBrokerOrCustodialAccount *RestrictToBrokerOrCustodialAccountCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _RestrictToBrokerOrCustodialAccount.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RestrictToBrokerOrCustodialAccount *RestrictToBrokerOrCustodialAccountTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RestrictToBrokerOrCustodialAccount.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RestrictToBrokerOrCustodialAccount *RestrictToBrokerOrCustodialAccountTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RestrictToBrokerOrCustodialAccount.Contract.contract.Transact(opts, method, params...)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_RestrictToBrokerOrCustodialAccount *RestrictToBrokerOrCustodialAccountCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _RestrictToBrokerOrCustodialAccount.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_RestrictToBrokerOrCustodialAccount *RestrictToBrokerOrCustodialAccountSession) Owner() (common.Address, error) {
	return _RestrictToBrokerOrCustodialAccount.Contract.Owner(&_RestrictToBrokerOrCustodialAccount.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_RestrictToBrokerOrCustodialAccount *RestrictToBrokerOrCustodialAccountCallerSession) Owner() (common.Address, error) {
	return _RestrictToBrokerOrCustodialAccount.Contract.Owner(&_RestrictToBrokerOrCustodialAccount.CallOpts)
}

// Check is a paid mutator transaction binding the contract method 0xb762c76d.
//
// Solidity: function check(initiator address, from address, to address, toKind uint8, tokens uint256, store address) returns()
func (_RestrictToBrokerOrCustodialAccount *RestrictToBrokerOrCustodialAccountTransactor) Check(opts *bind.TransactOpts, initiator common.Address, from common.Address, to common.Address, toKind uint8, tokens *big.Int, store common.Address) (*types.Transaction, error) {
	return _RestrictToBrokerOrCustodialAccount.contract.Transact(opts, "check", initiator, from, to, toKind, tokens, store)
}

// Check is a paid mutator transaction binding the contract method 0xb762c76d.
//
// Solidity: function check(initiator address, from address, to address, toKind uint8, tokens uint256, store address) returns()
func (_RestrictToBrokerOrCustodialAccount *RestrictToBrokerOrCustodialAccountSession) Check(initiator common.Address, from common.Address, to common.Address, toKind uint8, tokens *big.Int, store common.Address) (*types.Transaction, error) {
	return _RestrictToBrokerOrCustodialAccount.Contract.Check(&_RestrictToBrokerOrCustodialAccount.TransactOpts, initiator, from, to, toKind, tokens, store)
}

// Check is a paid mutator transaction binding the contract method 0xb762c76d.
//
// Solidity: function check(initiator address, from address, to address, toKind uint8, tokens uint256, store address) returns()
func (_RestrictToBrokerOrCustodialAccount *RestrictToBrokerOrCustodialAccountTransactorSession) Check(initiator common.Address, from common.Address, to common.Address, toKind uint8, tokens *big.Int, store common.Address) (*types.Transaction, error) {
	return _RestrictToBrokerOrCustodialAccount.Contract.Check(&_RestrictToBrokerOrCustodialAccount.TransactOpts, initiator, from, to, toKind, tokens, store)
}

// Kill is a paid mutator transaction binding the contract method 0x41c0e1b5.
//
// Solidity: function kill() returns()
func (_RestrictToBrokerOrCustodialAccount *RestrictToBrokerOrCustodialAccountTransactor) Kill(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RestrictToBrokerOrCustodialAccount.contract.Transact(opts, "kill")
}

// Kill is a paid mutator transaction binding the contract method 0x41c0e1b5.
//
// Solidity: function kill() returns()
func (_RestrictToBrokerOrCustodialAccount *RestrictToBrokerOrCustodialAccountSession) Kill() (*types.Transaction, error) {
	return _RestrictToBrokerOrCustodialAccount.Contract.Kill(&_RestrictToBrokerOrCustodialAccount.TransactOpts)
}

// Kill is a paid mutator transaction binding the contract method 0x41c0e1b5.
//
// Solidity: function kill() returns()
func (_RestrictToBrokerOrCustodialAccount *RestrictToBrokerOrCustodialAccountTransactorSession) Kill() (*types.Transaction, error) {
	return _RestrictToBrokerOrCustodialAccount.Contract.Kill(&_RestrictToBrokerOrCustodialAccount.TransactOpts)
}

// TransferOwner is a paid mutator transaction binding the contract method 0x4fb2e45d.
//
// Solidity: function transferOwner(newOwner address) returns()
func (_RestrictToBrokerOrCustodialAccount *RestrictToBrokerOrCustodialAccountTransactor) TransferOwner(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _RestrictToBrokerOrCustodialAccount.contract.Transact(opts, "transferOwner", newOwner)
}

// TransferOwner is a paid mutator transaction binding the contract method 0x4fb2e45d.
//
// Solidity: function transferOwner(newOwner address) returns()
func (_RestrictToBrokerOrCustodialAccount *RestrictToBrokerOrCustodialAccountSession) TransferOwner(newOwner common.Address) (*types.Transaction, error) {
	return _RestrictToBrokerOrCustodialAccount.Contract.TransferOwner(&_RestrictToBrokerOrCustodialAccount.TransactOpts, newOwner)
}

// TransferOwner is a paid mutator transaction binding the contract method 0x4fb2e45d.
//
// Solidity: function transferOwner(newOwner address) returns()
func (_RestrictToBrokerOrCustodialAccount *RestrictToBrokerOrCustodialAccountTransactorSession) TransferOwner(newOwner common.Address) (*types.Transaction, error) {
	return _RestrictToBrokerOrCustodialAccount.Contract.TransferOwner(&_RestrictToBrokerOrCustodialAccount.TransactOpts, newOwner)
}

// RestrictToBrokerOrCustodialAccountOwnerTransferredIterator is returned from FilterOwnerTransferred and is used to iterate over the raw logs and unpacked data for OwnerTransferred events raised by the RestrictToBrokerOrCustodialAccount contract.
type RestrictToBrokerOrCustodialAccountOwnerTransferredIterator struct {
	Event *RestrictToBrokerOrCustodialAccountOwnerTransferred // Event containing the contract specifics and raw log

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
func (it *RestrictToBrokerOrCustodialAccountOwnerTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RestrictToBrokerOrCustodialAccountOwnerTransferred)
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
		it.Event = new(RestrictToBrokerOrCustodialAccountOwnerTransferred)
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
func (it *RestrictToBrokerOrCustodialAccountOwnerTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RestrictToBrokerOrCustodialAccountOwnerTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RestrictToBrokerOrCustodialAccountOwnerTransferred represents a OwnerTransferred event raised by the RestrictToBrokerOrCustodialAccount contract.
type RestrictToBrokerOrCustodialAccountOwnerTransferred struct {
	OldOwner common.Address
	NewOwner common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOwnerTransferred is a free log retrieval operation binding the contract event 0x8934ce4adea8d9ce0d714d2c22b86790e41b7731c84b926fbbdc1d40ff6533c9.
//
// Solidity: e OwnerTransferred(oldOwner indexed address, newOwner indexed address)
func (_RestrictToBrokerOrCustodialAccount *RestrictToBrokerOrCustodialAccountFilterer) FilterOwnerTransferred(opts *bind.FilterOpts, oldOwner []common.Address, newOwner []common.Address) (*RestrictToBrokerOrCustodialAccountOwnerTransferredIterator, error) {

	var oldOwnerRule []interface{}
	for _, oldOwnerItem := range oldOwner {
		oldOwnerRule = append(oldOwnerRule, oldOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _RestrictToBrokerOrCustodialAccount.contract.FilterLogs(opts, "OwnerTransferred", oldOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &RestrictToBrokerOrCustodialAccountOwnerTransferredIterator{contract: _RestrictToBrokerOrCustodialAccount.contract, event: "OwnerTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnerTransferred is a free log subscription operation binding the contract event 0x8934ce4adea8d9ce0d714d2c22b86790e41b7731c84b926fbbdc1d40ff6533c9.
//
// Solidity: e OwnerTransferred(oldOwner indexed address, newOwner indexed address)
func (_RestrictToBrokerOrCustodialAccount *RestrictToBrokerOrCustodialAccountFilterer) WatchOwnerTransferred(opts *bind.WatchOpts, sink chan<- *RestrictToBrokerOrCustodialAccountOwnerTransferred, oldOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var oldOwnerRule []interface{}
	for _, oldOwnerItem := range oldOwner {
		oldOwnerRule = append(oldOwnerRule, oldOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _RestrictToBrokerOrCustodialAccount.contract.WatchLogs(opts, "OwnerTransferred", oldOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RestrictToBrokerOrCustodialAccountOwnerTransferred)
				if err := _RestrictToBrokerOrCustodialAccount.contract.UnpackLog(event, "OwnerTransferred", log); err != nil {
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
