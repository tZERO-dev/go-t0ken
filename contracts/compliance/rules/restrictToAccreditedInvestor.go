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

// RestrictToAccreditedInvestorABI is the input ABI used to generate the binding from.
const RestrictToAccreditedInvestorABI = "[{\"constant\":false,\"inputs\":[],\"name\":\"kill\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwner\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"initiator\",\"type\":\"address\"},{\"name\":\"from\",\"type\":\"address\"},{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"toKind\",\"type\":\"uint8\"},{\"name\":\"tokens\",\"type\":\"uint256\"},{\"name\":\"store\",\"type\":\"address\"}],\"name\":\"check\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"oldOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnerTransferred\",\"type\":\"event\"}]"

// RestrictToAccreditedInvestorBin is the compiled bytecode used for deploying new contracts.
const RestrictToAccreditedInvestorBin = `608060405260008054600160a060020a03191633179055610582806100256000396000f3fe608060405234801561001057600080fd5b5060043610610068577c0100000000000000000000000000000000000000000000000000000000600035046341c0e1b5811461006d5780634fb2e45d146100775780638da5cb5b146100aa578063b762c76d146100db575b600080fd5b610075610135565b005b6100756004803603602081101561008d57600080fd5b503573ffffffffffffffffffffffffffffffffffffffff166101d6565b6100b26103c9565b6040805173ffffffffffffffffffffffffffffffffffffffff9092168252519081900360200190f35b610075600480360360c08110156100f157600080fd5b5073ffffffffffffffffffffffffffffffffffffffff81358116916020810135821691604082013581169160ff6060820135169160808201359160a00135166103e5565b60005473ffffffffffffffffffffffffffffffffffffffff1633146101bb57604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601960248201527f4f776e6572206163636f756e7420697320726571756972656400000000000000604482015290519081900360640190fd5b60005473ffffffffffffffffffffffffffffffffffffffff16ff5b60005473ffffffffffffffffffffffffffffffffffffffff16331461025c57604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601960248201527f4f776e6572206163636f756e7420697320726571756972656400000000000000604482015290519081900360640190fd5b60005473ffffffffffffffffffffffffffffffffffffffff828116911614156102d0576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260258152602001806105086025913960400191505060405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff8116151561035457604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820181905260248201527f4e6577204f776e65722063616e6e6f74206265207a65726f2061646472657373604482015290519081900360640190fd5b6000805473ffffffffffffffffffffffffffffffffffffffff8381167fffffffffffffffffffffffff0000000000000000000000000000000000000000831681178455604051919092169283917f8934ce4adea8d9ce0d714d2c22b86790e41b7731c84b926fbbdc1d40ff6533c99190a35050565b60005473ffffffffffffffffffffffffffffffffffffffff1681565b60ff8316600414156104ff57604080517f572bc6b100000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff868116600483015260016024830152915142926010929085169163572bc6b191604480820192602092909190829003018186803b15801561046d57600080fd5b505afa158015610481573d6000803e3d6000fd5b505050506040513d602081101561049757600080fd5b505165ffffffffffff60029290920a900416116104ff576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252602a81526020018061052d602a913960400191505060405180910390fd5b50505050505056fe4e6577204f776e65722063616e6e6f74206265207468652063757272656e74206f776e657254686520746f2061646472657373206973206e6f742063757272656e746c792061636372656469746564a165627a7a72305820726c3f89ad6a45e60ccae28e8f6a22448d869583030e94134c363d54b69fd2e70029`

// DeployRestrictToAccreditedInvestor deploys a new Ethereum contract, binding an instance of RestrictToAccreditedInvestor to it.
func DeployRestrictToAccreditedInvestor(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *RestrictToAccreditedInvestor, error) {
	parsed, err := abi.JSON(strings.NewReader(RestrictToAccreditedInvestorABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(RestrictToAccreditedInvestorBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &RestrictToAccreditedInvestor{RestrictToAccreditedInvestorCaller: RestrictToAccreditedInvestorCaller{contract: contract}, RestrictToAccreditedInvestorTransactor: RestrictToAccreditedInvestorTransactor{contract: contract}, RestrictToAccreditedInvestorFilterer: RestrictToAccreditedInvestorFilterer{contract: contract}}, nil
}

// RestrictToAccreditedInvestor is an auto generated Go binding around an Ethereum contract.
type RestrictToAccreditedInvestor struct {
	RestrictToAccreditedInvestorCaller     // Read-only binding to the contract
	RestrictToAccreditedInvestorTransactor // Write-only binding to the contract
	RestrictToAccreditedInvestorFilterer   // Log filterer for contract events
}

// RestrictToAccreditedInvestorCaller is an auto generated read-only Go binding around an Ethereum contract.
type RestrictToAccreditedInvestorCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RestrictToAccreditedInvestorTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RestrictToAccreditedInvestorTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RestrictToAccreditedInvestorFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RestrictToAccreditedInvestorFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RestrictToAccreditedInvestorSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RestrictToAccreditedInvestorSession struct {
	Contract     *RestrictToAccreditedInvestor // Generic contract binding to set the session for
	CallOpts     bind.CallOpts                 // Call options to use throughout this session
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// RestrictToAccreditedInvestorCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RestrictToAccreditedInvestorCallerSession struct {
	Contract *RestrictToAccreditedInvestorCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                       // Call options to use throughout this session
}

// RestrictToAccreditedInvestorTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RestrictToAccreditedInvestorTransactorSession struct {
	Contract     *RestrictToAccreditedInvestorTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                       // Transaction auth options to use throughout this session
}

// RestrictToAccreditedInvestorRaw is an auto generated low-level Go binding around an Ethereum contract.
type RestrictToAccreditedInvestorRaw struct {
	Contract *RestrictToAccreditedInvestor // Generic contract binding to access the raw methods on
}

// RestrictToAccreditedInvestorCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RestrictToAccreditedInvestorCallerRaw struct {
	Contract *RestrictToAccreditedInvestorCaller // Generic read-only contract binding to access the raw methods on
}

// RestrictToAccreditedInvestorTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RestrictToAccreditedInvestorTransactorRaw struct {
	Contract *RestrictToAccreditedInvestorTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRestrictToAccreditedInvestor creates a new instance of RestrictToAccreditedInvestor, bound to a specific deployed contract.
func NewRestrictToAccreditedInvestor(address common.Address, backend bind.ContractBackend) (*RestrictToAccreditedInvestor, error) {
	contract, err := bindRestrictToAccreditedInvestor(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &RestrictToAccreditedInvestor{RestrictToAccreditedInvestorCaller: RestrictToAccreditedInvestorCaller{contract: contract}, RestrictToAccreditedInvestorTransactor: RestrictToAccreditedInvestorTransactor{contract: contract}, RestrictToAccreditedInvestorFilterer: RestrictToAccreditedInvestorFilterer{contract: contract}}, nil
}

// NewRestrictToAccreditedInvestorCaller creates a new read-only instance of RestrictToAccreditedInvestor, bound to a specific deployed contract.
func NewRestrictToAccreditedInvestorCaller(address common.Address, caller bind.ContractCaller) (*RestrictToAccreditedInvestorCaller, error) {
	contract, err := bindRestrictToAccreditedInvestor(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RestrictToAccreditedInvestorCaller{contract: contract}, nil
}

// NewRestrictToAccreditedInvestorTransactor creates a new write-only instance of RestrictToAccreditedInvestor, bound to a specific deployed contract.
func NewRestrictToAccreditedInvestorTransactor(address common.Address, transactor bind.ContractTransactor) (*RestrictToAccreditedInvestorTransactor, error) {
	contract, err := bindRestrictToAccreditedInvestor(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RestrictToAccreditedInvestorTransactor{contract: contract}, nil
}

// NewRestrictToAccreditedInvestorFilterer creates a new log filterer instance of RestrictToAccreditedInvestor, bound to a specific deployed contract.
func NewRestrictToAccreditedInvestorFilterer(address common.Address, filterer bind.ContractFilterer) (*RestrictToAccreditedInvestorFilterer, error) {
	contract, err := bindRestrictToAccreditedInvestor(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RestrictToAccreditedInvestorFilterer{contract: contract}, nil
}

// bindRestrictToAccreditedInvestor binds a generic wrapper to an already deployed contract.
func bindRestrictToAccreditedInvestor(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(RestrictToAccreditedInvestorABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RestrictToAccreditedInvestor *RestrictToAccreditedInvestorRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _RestrictToAccreditedInvestor.Contract.RestrictToAccreditedInvestorCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RestrictToAccreditedInvestor *RestrictToAccreditedInvestorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RestrictToAccreditedInvestor.Contract.RestrictToAccreditedInvestorTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RestrictToAccreditedInvestor *RestrictToAccreditedInvestorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RestrictToAccreditedInvestor.Contract.RestrictToAccreditedInvestorTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RestrictToAccreditedInvestor *RestrictToAccreditedInvestorCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _RestrictToAccreditedInvestor.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RestrictToAccreditedInvestor *RestrictToAccreditedInvestorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RestrictToAccreditedInvestor.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RestrictToAccreditedInvestor *RestrictToAccreditedInvestorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RestrictToAccreditedInvestor.Contract.contract.Transact(opts, method, params...)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_RestrictToAccreditedInvestor *RestrictToAccreditedInvestorCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _RestrictToAccreditedInvestor.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_RestrictToAccreditedInvestor *RestrictToAccreditedInvestorSession) Owner() (common.Address, error) {
	return _RestrictToAccreditedInvestor.Contract.Owner(&_RestrictToAccreditedInvestor.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_RestrictToAccreditedInvestor *RestrictToAccreditedInvestorCallerSession) Owner() (common.Address, error) {
	return _RestrictToAccreditedInvestor.Contract.Owner(&_RestrictToAccreditedInvestor.CallOpts)
}

// Check is a paid mutator transaction binding the contract method 0xb762c76d.
//
// Solidity: function check(initiator address, from address, to address, toKind uint8, tokens uint256, store address) returns()
func (_RestrictToAccreditedInvestor *RestrictToAccreditedInvestorTransactor) Check(opts *bind.TransactOpts, initiator common.Address, from common.Address, to common.Address, toKind uint8, tokens *big.Int, store common.Address) (*types.Transaction, error) {
	return _RestrictToAccreditedInvestor.contract.Transact(opts, "check", initiator, from, to, toKind, tokens, store)
}

// Check is a paid mutator transaction binding the contract method 0xb762c76d.
//
// Solidity: function check(initiator address, from address, to address, toKind uint8, tokens uint256, store address) returns()
func (_RestrictToAccreditedInvestor *RestrictToAccreditedInvestorSession) Check(initiator common.Address, from common.Address, to common.Address, toKind uint8, tokens *big.Int, store common.Address) (*types.Transaction, error) {
	return _RestrictToAccreditedInvestor.Contract.Check(&_RestrictToAccreditedInvestor.TransactOpts, initiator, from, to, toKind, tokens, store)
}

// Check is a paid mutator transaction binding the contract method 0xb762c76d.
//
// Solidity: function check(initiator address, from address, to address, toKind uint8, tokens uint256, store address) returns()
func (_RestrictToAccreditedInvestor *RestrictToAccreditedInvestorTransactorSession) Check(initiator common.Address, from common.Address, to common.Address, toKind uint8, tokens *big.Int, store common.Address) (*types.Transaction, error) {
	return _RestrictToAccreditedInvestor.Contract.Check(&_RestrictToAccreditedInvestor.TransactOpts, initiator, from, to, toKind, tokens, store)
}

// Kill is a paid mutator transaction binding the contract method 0x41c0e1b5.
//
// Solidity: function kill() returns()
func (_RestrictToAccreditedInvestor *RestrictToAccreditedInvestorTransactor) Kill(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RestrictToAccreditedInvestor.contract.Transact(opts, "kill")
}

// Kill is a paid mutator transaction binding the contract method 0x41c0e1b5.
//
// Solidity: function kill() returns()
func (_RestrictToAccreditedInvestor *RestrictToAccreditedInvestorSession) Kill() (*types.Transaction, error) {
	return _RestrictToAccreditedInvestor.Contract.Kill(&_RestrictToAccreditedInvestor.TransactOpts)
}

// Kill is a paid mutator transaction binding the contract method 0x41c0e1b5.
//
// Solidity: function kill() returns()
func (_RestrictToAccreditedInvestor *RestrictToAccreditedInvestorTransactorSession) Kill() (*types.Transaction, error) {
	return _RestrictToAccreditedInvestor.Contract.Kill(&_RestrictToAccreditedInvestor.TransactOpts)
}

// TransferOwner is a paid mutator transaction binding the contract method 0x4fb2e45d.
//
// Solidity: function transferOwner(newOwner address) returns()
func (_RestrictToAccreditedInvestor *RestrictToAccreditedInvestorTransactor) TransferOwner(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _RestrictToAccreditedInvestor.contract.Transact(opts, "transferOwner", newOwner)
}

// TransferOwner is a paid mutator transaction binding the contract method 0x4fb2e45d.
//
// Solidity: function transferOwner(newOwner address) returns()
func (_RestrictToAccreditedInvestor *RestrictToAccreditedInvestorSession) TransferOwner(newOwner common.Address) (*types.Transaction, error) {
	return _RestrictToAccreditedInvestor.Contract.TransferOwner(&_RestrictToAccreditedInvestor.TransactOpts, newOwner)
}

// TransferOwner is a paid mutator transaction binding the contract method 0x4fb2e45d.
//
// Solidity: function transferOwner(newOwner address) returns()
func (_RestrictToAccreditedInvestor *RestrictToAccreditedInvestorTransactorSession) TransferOwner(newOwner common.Address) (*types.Transaction, error) {
	return _RestrictToAccreditedInvestor.Contract.TransferOwner(&_RestrictToAccreditedInvestor.TransactOpts, newOwner)
}

// RestrictToAccreditedInvestorOwnerTransferredIterator is returned from FilterOwnerTransferred and is used to iterate over the raw logs and unpacked data for OwnerTransferred events raised by the RestrictToAccreditedInvestor contract.
type RestrictToAccreditedInvestorOwnerTransferredIterator struct {
	Event *RestrictToAccreditedInvestorOwnerTransferred // Event containing the contract specifics and raw log

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
func (it *RestrictToAccreditedInvestorOwnerTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RestrictToAccreditedInvestorOwnerTransferred)
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
		it.Event = new(RestrictToAccreditedInvestorOwnerTransferred)
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
func (it *RestrictToAccreditedInvestorOwnerTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RestrictToAccreditedInvestorOwnerTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RestrictToAccreditedInvestorOwnerTransferred represents a OwnerTransferred event raised by the RestrictToAccreditedInvestor contract.
type RestrictToAccreditedInvestorOwnerTransferred struct {
	OldOwner common.Address
	NewOwner common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOwnerTransferred is a free log retrieval operation binding the contract event 0x8934ce4adea8d9ce0d714d2c22b86790e41b7731c84b926fbbdc1d40ff6533c9.
//
// Solidity: e OwnerTransferred(oldOwner indexed address, newOwner indexed address)
func (_RestrictToAccreditedInvestor *RestrictToAccreditedInvestorFilterer) FilterOwnerTransferred(opts *bind.FilterOpts, oldOwner []common.Address, newOwner []common.Address) (*RestrictToAccreditedInvestorOwnerTransferredIterator, error) {

	var oldOwnerRule []interface{}
	for _, oldOwnerItem := range oldOwner {
		oldOwnerRule = append(oldOwnerRule, oldOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _RestrictToAccreditedInvestor.contract.FilterLogs(opts, "OwnerTransferred", oldOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &RestrictToAccreditedInvestorOwnerTransferredIterator{contract: _RestrictToAccreditedInvestor.contract, event: "OwnerTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnerTransferred is a free log subscription operation binding the contract event 0x8934ce4adea8d9ce0d714d2c22b86790e41b7731c84b926fbbdc1d40ff6533c9.
//
// Solidity: e OwnerTransferred(oldOwner indexed address, newOwner indexed address)
func (_RestrictToAccreditedInvestor *RestrictToAccreditedInvestorFilterer) WatchOwnerTransferred(opts *bind.WatchOpts, sink chan<- *RestrictToAccreditedInvestorOwnerTransferred, oldOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var oldOwnerRule []interface{}
	for _, oldOwnerItem := range oldOwner {
		oldOwnerRule = append(oldOwnerRule, oldOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _RestrictToAccreditedInvestor.contract.WatchLogs(opts, "OwnerTransferred", oldOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RestrictToAccreditedInvestorOwnerTransferred)
				if err := _RestrictToAccreditedInvestor.contract.UnpackLog(event, "OwnerTransferred", log); err != nil {
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
