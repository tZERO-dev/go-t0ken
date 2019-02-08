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

// RestrictFromAffiliateABI is the input ABI used to generate the binding from.
const RestrictFromAffiliateABI = "[{\"constant\":false,\"inputs\":[],\"name\":\"kill\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"affiliates\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwner\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"},{\"name\":\"status\",\"type\":\"bool\"}],\"name\":\"setAffiliate\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"initiator\",\"type\":\"address\"},{\"name\":\"from\",\"type\":\"address\"},{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"toKind\",\"type\":\"uint8\"},{\"name\":\"tokens\",\"type\":\"uint256\"},{\"name\":\"store\",\"type\":\"address\"}],\"name\":\"check\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"oldOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnerTransferred\",\"type\":\"event\"}]"

// RestrictFromAffiliateBin is the compiled bytecode used for deploying new contracts.
const RestrictFromAffiliateBin = `608060405260008054600160a060020a03191633179055610679806100256000396000f3fe608060405234801561001057600080fd5b506004361061008f576000357c0100000000000000000000000000000000000000000000000000000000900480638863ebbb1161006d5780638863ebbb146101185780638da5cb5b14610153578063b762c76d146101845761008f565b806341c0e1b5146100945780634f51e2941461009e5780634fb2e45d146100e5575b600080fd5b61009c6101de565b005b6100d1600480360360208110156100b457600080fd5b503573ffffffffffffffffffffffffffffffffffffffff1661027f565b604080519115158252519081900360200190f35b61009c600480360360208110156100fb57600080fd5b503573ffffffffffffffffffffffffffffffffffffffff16610294565b61009c6004803603604081101561012e57600080fd5b5073ffffffffffffffffffffffffffffffffffffffff81351690602001351515610487565b61015b610545565b6040805173ffffffffffffffffffffffffffffffffffffffff9092168252519081900360200190f35b61009c600480360360c081101561019a57600080fd5b5073ffffffffffffffffffffffffffffffffffffffff81358116916020810135821691604082013581169160ff6060820135169160808201359160a0013516610561565b60005473ffffffffffffffffffffffffffffffffffffffff16331461026457604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601960248201527f4f776e6572206163636f756e7420697320726571756972656400000000000000604482015290519081900360640190fd5b60005473ffffffffffffffffffffffffffffffffffffffff16ff5b60016020526000908152604090205460ff1681565b60005473ffffffffffffffffffffffffffffffffffffffff16331461031a57604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601960248201527f4f776e6572206163636f756e7420697320726571756972656400000000000000604482015290519081900360640190fd5b60005473ffffffffffffffffffffffffffffffffffffffff8281169116141561038e576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260258152602001806105e96025913960400191505060405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff8116151561041257604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820181905260248201527f4e6577204f776e65722063616e6e6f74206265207a65726f2061646472657373604482015290519081900360640190fd5b6000805473ffffffffffffffffffffffffffffffffffffffff8381167fffffffffffffffffffffffff0000000000000000000000000000000000000000831681178455604051919092169283917f8934ce4adea8d9ce0d714d2c22b86790e41b7731c84b926fbbdc1d40ff6533c99190a35050565b60005473ffffffffffffffffffffffffffffffffffffffff16331461050d57604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601960248201527f4f776e6572206163636f756e7420697320726571756972656400000000000000604482015290519081900360640190fd5b73ffffffffffffffffffffffffffffffffffffffff919091166000908152600160205260409020805460ff1916911515919091179055565b60005473ffffffffffffffffffffffffffffffffffffffff1681565b73ffffffffffffffffffffffffffffffffffffffff851660009081526001602052604090205460ff16156105e0576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252604081526020018061060e6040913960400191505060405180910390fd5b50505050505056fe4e6577204f776e65722063616e6e6f74206265207468652063757272656e74206f776e65725468652066726f6d206164647265737320697320616e20616666696c6961746520616e64206e6f7420616c6c6f77656420746f2073656e6420746f6b656e732ea165627a7a72305820581692f7e59a333eb41b86638b61d7503ac54246d96b963cb8b79ac167bc82640029`

// DeployRestrictFromAffiliate deploys a new Ethereum contract, binding an instance of RestrictFromAffiliate to it.
func DeployRestrictFromAffiliate(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *RestrictFromAffiliate, error) {
	parsed, err := abi.JSON(strings.NewReader(RestrictFromAffiliateABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(RestrictFromAffiliateBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &RestrictFromAffiliate{RestrictFromAffiliateCaller: RestrictFromAffiliateCaller{contract: contract}, RestrictFromAffiliateTransactor: RestrictFromAffiliateTransactor{contract: contract}, RestrictFromAffiliateFilterer: RestrictFromAffiliateFilterer{contract: contract}}, nil
}

// RestrictFromAffiliate is an auto generated Go binding around an Ethereum contract.
type RestrictFromAffiliate struct {
	RestrictFromAffiliateCaller     // Read-only binding to the contract
	RestrictFromAffiliateTransactor // Write-only binding to the contract
	RestrictFromAffiliateFilterer   // Log filterer for contract events
}

// RestrictFromAffiliateCaller is an auto generated read-only Go binding around an Ethereum contract.
type RestrictFromAffiliateCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RestrictFromAffiliateTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RestrictFromAffiliateTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RestrictFromAffiliateFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RestrictFromAffiliateFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RestrictFromAffiliateSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RestrictFromAffiliateSession struct {
	Contract     *RestrictFromAffiliate // Generic contract binding to set the session for
	CallOpts     bind.CallOpts          // Call options to use throughout this session
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// RestrictFromAffiliateCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RestrictFromAffiliateCallerSession struct {
	Contract *RestrictFromAffiliateCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                // Call options to use throughout this session
}

// RestrictFromAffiliateTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RestrictFromAffiliateTransactorSession struct {
	Contract     *RestrictFromAffiliateTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                // Transaction auth options to use throughout this session
}

// RestrictFromAffiliateRaw is an auto generated low-level Go binding around an Ethereum contract.
type RestrictFromAffiliateRaw struct {
	Contract *RestrictFromAffiliate // Generic contract binding to access the raw methods on
}

// RestrictFromAffiliateCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RestrictFromAffiliateCallerRaw struct {
	Contract *RestrictFromAffiliateCaller // Generic read-only contract binding to access the raw methods on
}

// RestrictFromAffiliateTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RestrictFromAffiliateTransactorRaw struct {
	Contract *RestrictFromAffiliateTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRestrictFromAffiliate creates a new instance of RestrictFromAffiliate, bound to a specific deployed contract.
func NewRestrictFromAffiliate(address common.Address, backend bind.ContractBackend) (*RestrictFromAffiliate, error) {
	contract, err := bindRestrictFromAffiliate(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &RestrictFromAffiliate{RestrictFromAffiliateCaller: RestrictFromAffiliateCaller{contract: contract}, RestrictFromAffiliateTransactor: RestrictFromAffiliateTransactor{contract: contract}, RestrictFromAffiliateFilterer: RestrictFromAffiliateFilterer{contract: contract}}, nil
}

// NewRestrictFromAffiliateCaller creates a new read-only instance of RestrictFromAffiliate, bound to a specific deployed contract.
func NewRestrictFromAffiliateCaller(address common.Address, caller bind.ContractCaller) (*RestrictFromAffiliateCaller, error) {
	contract, err := bindRestrictFromAffiliate(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RestrictFromAffiliateCaller{contract: contract}, nil
}

// NewRestrictFromAffiliateTransactor creates a new write-only instance of RestrictFromAffiliate, bound to a specific deployed contract.
func NewRestrictFromAffiliateTransactor(address common.Address, transactor bind.ContractTransactor) (*RestrictFromAffiliateTransactor, error) {
	contract, err := bindRestrictFromAffiliate(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RestrictFromAffiliateTransactor{contract: contract}, nil
}

// NewRestrictFromAffiliateFilterer creates a new log filterer instance of RestrictFromAffiliate, bound to a specific deployed contract.
func NewRestrictFromAffiliateFilterer(address common.Address, filterer bind.ContractFilterer) (*RestrictFromAffiliateFilterer, error) {
	contract, err := bindRestrictFromAffiliate(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RestrictFromAffiliateFilterer{contract: contract}, nil
}

// bindRestrictFromAffiliate binds a generic wrapper to an already deployed contract.
func bindRestrictFromAffiliate(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(RestrictFromAffiliateABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RestrictFromAffiliate *RestrictFromAffiliateRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _RestrictFromAffiliate.Contract.RestrictFromAffiliateCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RestrictFromAffiliate *RestrictFromAffiliateRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RestrictFromAffiliate.Contract.RestrictFromAffiliateTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RestrictFromAffiliate *RestrictFromAffiliateRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RestrictFromAffiliate.Contract.RestrictFromAffiliateTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RestrictFromAffiliate *RestrictFromAffiliateCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _RestrictFromAffiliate.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RestrictFromAffiliate *RestrictFromAffiliateTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RestrictFromAffiliate.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RestrictFromAffiliate *RestrictFromAffiliateTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RestrictFromAffiliate.Contract.contract.Transact(opts, method, params...)
}

// Affiliates is a free data retrieval call binding the contract method 0x4f51e294.
//
// Solidity: function affiliates( address) constant returns(bool)
func (_RestrictFromAffiliate *RestrictFromAffiliateCaller) Affiliates(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _RestrictFromAffiliate.contract.Call(opts, out, "affiliates", arg0)
	return *ret0, err
}

// Affiliates is a free data retrieval call binding the contract method 0x4f51e294.
//
// Solidity: function affiliates( address) constant returns(bool)
func (_RestrictFromAffiliate *RestrictFromAffiliateSession) Affiliates(arg0 common.Address) (bool, error) {
	return _RestrictFromAffiliate.Contract.Affiliates(&_RestrictFromAffiliate.CallOpts, arg0)
}

// Affiliates is a free data retrieval call binding the contract method 0x4f51e294.
//
// Solidity: function affiliates( address) constant returns(bool)
func (_RestrictFromAffiliate *RestrictFromAffiliateCallerSession) Affiliates(arg0 common.Address) (bool, error) {
	return _RestrictFromAffiliate.Contract.Affiliates(&_RestrictFromAffiliate.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_RestrictFromAffiliate *RestrictFromAffiliateCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _RestrictFromAffiliate.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_RestrictFromAffiliate *RestrictFromAffiliateSession) Owner() (common.Address, error) {
	return _RestrictFromAffiliate.Contract.Owner(&_RestrictFromAffiliate.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_RestrictFromAffiliate *RestrictFromAffiliateCallerSession) Owner() (common.Address, error) {
	return _RestrictFromAffiliate.Contract.Owner(&_RestrictFromAffiliate.CallOpts)
}

// Check is a paid mutator transaction binding the contract method 0xb762c76d.
//
// Solidity: function check(initiator address, from address, to address, toKind uint8, tokens uint256, store address) returns()
func (_RestrictFromAffiliate *RestrictFromAffiliateTransactor) Check(opts *bind.TransactOpts, initiator common.Address, from common.Address, to common.Address, toKind uint8, tokens *big.Int, store common.Address) (*types.Transaction, error) {
	return _RestrictFromAffiliate.contract.Transact(opts, "check", initiator, from, to, toKind, tokens, store)
}

// Check is a paid mutator transaction binding the contract method 0xb762c76d.
//
// Solidity: function check(initiator address, from address, to address, toKind uint8, tokens uint256, store address) returns()
func (_RestrictFromAffiliate *RestrictFromAffiliateSession) Check(initiator common.Address, from common.Address, to common.Address, toKind uint8, tokens *big.Int, store common.Address) (*types.Transaction, error) {
	return _RestrictFromAffiliate.Contract.Check(&_RestrictFromAffiliate.TransactOpts, initiator, from, to, toKind, tokens, store)
}

// Check is a paid mutator transaction binding the contract method 0xb762c76d.
//
// Solidity: function check(initiator address, from address, to address, toKind uint8, tokens uint256, store address) returns()
func (_RestrictFromAffiliate *RestrictFromAffiliateTransactorSession) Check(initiator common.Address, from common.Address, to common.Address, toKind uint8, tokens *big.Int, store common.Address) (*types.Transaction, error) {
	return _RestrictFromAffiliate.Contract.Check(&_RestrictFromAffiliate.TransactOpts, initiator, from, to, toKind, tokens, store)
}

// Kill is a paid mutator transaction binding the contract method 0x41c0e1b5.
//
// Solidity: function kill() returns()
func (_RestrictFromAffiliate *RestrictFromAffiliateTransactor) Kill(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RestrictFromAffiliate.contract.Transact(opts, "kill")
}

// Kill is a paid mutator transaction binding the contract method 0x41c0e1b5.
//
// Solidity: function kill() returns()
func (_RestrictFromAffiliate *RestrictFromAffiliateSession) Kill() (*types.Transaction, error) {
	return _RestrictFromAffiliate.Contract.Kill(&_RestrictFromAffiliate.TransactOpts)
}

// Kill is a paid mutator transaction binding the contract method 0x41c0e1b5.
//
// Solidity: function kill() returns()
func (_RestrictFromAffiliate *RestrictFromAffiliateTransactorSession) Kill() (*types.Transaction, error) {
	return _RestrictFromAffiliate.Contract.Kill(&_RestrictFromAffiliate.TransactOpts)
}

// SetAffiliate is a paid mutator transaction binding the contract method 0x8863ebbb.
//
// Solidity: function setAffiliate(addr address, status bool) returns()
func (_RestrictFromAffiliate *RestrictFromAffiliateTransactor) SetAffiliate(opts *bind.TransactOpts, addr common.Address, status bool) (*types.Transaction, error) {
	return _RestrictFromAffiliate.contract.Transact(opts, "setAffiliate", addr, status)
}

// SetAffiliate is a paid mutator transaction binding the contract method 0x8863ebbb.
//
// Solidity: function setAffiliate(addr address, status bool) returns()
func (_RestrictFromAffiliate *RestrictFromAffiliateSession) SetAffiliate(addr common.Address, status bool) (*types.Transaction, error) {
	return _RestrictFromAffiliate.Contract.SetAffiliate(&_RestrictFromAffiliate.TransactOpts, addr, status)
}

// SetAffiliate is a paid mutator transaction binding the contract method 0x8863ebbb.
//
// Solidity: function setAffiliate(addr address, status bool) returns()
func (_RestrictFromAffiliate *RestrictFromAffiliateTransactorSession) SetAffiliate(addr common.Address, status bool) (*types.Transaction, error) {
	return _RestrictFromAffiliate.Contract.SetAffiliate(&_RestrictFromAffiliate.TransactOpts, addr, status)
}

// TransferOwner is a paid mutator transaction binding the contract method 0x4fb2e45d.
//
// Solidity: function transferOwner(newOwner address) returns()
func (_RestrictFromAffiliate *RestrictFromAffiliateTransactor) TransferOwner(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _RestrictFromAffiliate.contract.Transact(opts, "transferOwner", newOwner)
}

// TransferOwner is a paid mutator transaction binding the contract method 0x4fb2e45d.
//
// Solidity: function transferOwner(newOwner address) returns()
func (_RestrictFromAffiliate *RestrictFromAffiliateSession) TransferOwner(newOwner common.Address) (*types.Transaction, error) {
	return _RestrictFromAffiliate.Contract.TransferOwner(&_RestrictFromAffiliate.TransactOpts, newOwner)
}

// TransferOwner is a paid mutator transaction binding the contract method 0x4fb2e45d.
//
// Solidity: function transferOwner(newOwner address) returns()
func (_RestrictFromAffiliate *RestrictFromAffiliateTransactorSession) TransferOwner(newOwner common.Address) (*types.Transaction, error) {
	return _RestrictFromAffiliate.Contract.TransferOwner(&_RestrictFromAffiliate.TransactOpts, newOwner)
}

// RestrictFromAffiliateOwnerTransferredIterator is returned from FilterOwnerTransferred and is used to iterate over the raw logs and unpacked data for OwnerTransferred events raised by the RestrictFromAffiliate contract.
type RestrictFromAffiliateOwnerTransferredIterator struct {
	Event *RestrictFromAffiliateOwnerTransferred // Event containing the contract specifics and raw log

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
func (it *RestrictFromAffiliateOwnerTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RestrictFromAffiliateOwnerTransferred)
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
		it.Event = new(RestrictFromAffiliateOwnerTransferred)
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
func (it *RestrictFromAffiliateOwnerTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RestrictFromAffiliateOwnerTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RestrictFromAffiliateOwnerTransferred represents a OwnerTransferred event raised by the RestrictFromAffiliate contract.
type RestrictFromAffiliateOwnerTransferred struct {
	OldOwner common.Address
	NewOwner common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOwnerTransferred is a free log retrieval operation binding the contract event 0x8934ce4adea8d9ce0d714d2c22b86790e41b7731c84b926fbbdc1d40ff6533c9.
//
// Solidity: e OwnerTransferred(oldOwner indexed address, newOwner indexed address)
func (_RestrictFromAffiliate *RestrictFromAffiliateFilterer) FilterOwnerTransferred(opts *bind.FilterOpts, oldOwner []common.Address, newOwner []common.Address) (*RestrictFromAffiliateOwnerTransferredIterator, error) {

	var oldOwnerRule []interface{}
	for _, oldOwnerItem := range oldOwner {
		oldOwnerRule = append(oldOwnerRule, oldOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _RestrictFromAffiliate.contract.FilterLogs(opts, "OwnerTransferred", oldOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &RestrictFromAffiliateOwnerTransferredIterator{contract: _RestrictFromAffiliate.contract, event: "OwnerTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnerTransferred is a free log subscription operation binding the contract event 0x8934ce4adea8d9ce0d714d2c22b86790e41b7731c84b926fbbdc1d40ff6533c9.
//
// Solidity: e OwnerTransferred(oldOwner indexed address, newOwner indexed address)
func (_RestrictFromAffiliate *RestrictFromAffiliateFilterer) WatchOwnerTransferred(opts *bind.WatchOpts, sink chan<- *RestrictFromAffiliateOwnerTransferred, oldOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var oldOwnerRule []interface{}
	for _, oldOwnerItem := range oldOwner {
		oldOwnerRule = append(oldOwnerRule, oldOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _RestrictFromAffiliate.contract.WatchLogs(opts, "OwnerTransferred", oldOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RestrictFromAffiliateOwnerTransferred)
				if err := _RestrictFromAffiliate.contract.UnpackLog(event, "OwnerTransferred", log); err != nil {
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
