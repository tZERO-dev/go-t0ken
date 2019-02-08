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

// RestrictToContractABI is the input ABI used to generate the binding from.
const RestrictToContractABI = "[{\"constant\":false,\"inputs\":[],\"name\":\"kill\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwner\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"initiator\",\"type\":\"address\"},{\"name\":\"from\",\"type\":\"address\"},{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"toKind\",\"type\":\"uint8\"},{\"name\":\"tokens\",\"type\":\"uint256\"},{\"name\":\"store\",\"type\":\"address\"}],\"name\":\"check\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"oldOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnerTransferred\",\"type\":\"event\"}]"

// RestrictToContractBin is the compiled bytecode used for deploying new contracts.
const RestrictToContractBin = `608060405260008054600160a060020a031916331790556104c5806100256000396000f3fe608060405234801561001057600080fd5b5060043610610068577c0100000000000000000000000000000000000000000000000000000000600035046341c0e1b5811461006d5780634fb2e45d146100775780638da5cb5b146100aa578063b762c76d146100db575b600080fd5b610075610135565b005b6100756004803603602081101561008d57600080fd5b503573ffffffffffffffffffffffffffffffffffffffff166101d6565b6100b26103c9565b6040805173ffffffffffffffffffffffffffffffffffffffff9092168252519081900360200190f35b610075600480360360c08110156100f157600080fd5b5073ffffffffffffffffffffffffffffffffffffffff81358116916020810135821691604082013581169160ff6060820135169160808201359160a00135166103e5565b60005473ffffffffffffffffffffffffffffffffffffffff1633146101bb57604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601960248201527f4f776e6572206163636f756e7420697320726571756972656400000000000000604482015290519081900360640190fd5b60005473ffffffffffffffffffffffffffffffffffffffff16ff5b60005473ffffffffffffffffffffffffffffffffffffffff16331461025c57604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601960248201527f4f776e6572206163636f756e7420697320726571756972656400000000000000604482015290519081900360640190fd5b60005473ffffffffffffffffffffffffffffffffffffffff828116911614156102d0576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252602581526020018061044e6025913960400191505060405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff8116151561035457604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820181905260248201527f4e6577204f776e65722063616e6e6f74206265207a65726f2061646472657373604482015290519081900360640190fd5b6000805473ffffffffffffffffffffffffffffffffffffffff8381167fffffffffffffffffffffffff0000000000000000000000000000000000000000831681178455604051919092169283917f8934ce4adea8d9ce0d714d2c22b86790e41b7731c84b926fbbdc1d40ff6533c99190a35050565b60005473ffffffffffffffffffffffffffffffffffffffff1681565b833b63ffffffff811615610444576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260278152602001806104736027913960400191505060405180910390fd5b5050505050505056fe4e6577204f776e65722063616e6e6f74206265207468652063757272656e74206f776e65725472616e736665727320746f20636f6e74726163747320617265206e6f7420616c6c6f7765642ea165627a7a723058204551cdbebcd3d9475aa9c5305fe48c0016349f5b9148150748b308768aa083b20029`

// DeployRestrictToContract deploys a new Ethereum contract, binding an instance of RestrictToContract to it.
func DeployRestrictToContract(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *RestrictToContract, error) {
	parsed, err := abi.JSON(strings.NewReader(RestrictToContractABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(RestrictToContractBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &RestrictToContract{RestrictToContractCaller: RestrictToContractCaller{contract: contract}, RestrictToContractTransactor: RestrictToContractTransactor{contract: contract}, RestrictToContractFilterer: RestrictToContractFilterer{contract: contract}}, nil
}

// RestrictToContract is an auto generated Go binding around an Ethereum contract.
type RestrictToContract struct {
	RestrictToContractCaller     // Read-only binding to the contract
	RestrictToContractTransactor // Write-only binding to the contract
	RestrictToContractFilterer   // Log filterer for contract events
}

// RestrictToContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type RestrictToContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RestrictToContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RestrictToContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RestrictToContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RestrictToContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RestrictToContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RestrictToContractSession struct {
	Contract     *RestrictToContract // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// RestrictToContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RestrictToContractCallerSession struct {
	Contract *RestrictToContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// RestrictToContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RestrictToContractTransactorSession struct {
	Contract     *RestrictToContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// RestrictToContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type RestrictToContractRaw struct {
	Contract *RestrictToContract // Generic contract binding to access the raw methods on
}

// RestrictToContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RestrictToContractCallerRaw struct {
	Contract *RestrictToContractCaller // Generic read-only contract binding to access the raw methods on
}

// RestrictToContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RestrictToContractTransactorRaw struct {
	Contract *RestrictToContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRestrictToContract creates a new instance of RestrictToContract, bound to a specific deployed contract.
func NewRestrictToContract(address common.Address, backend bind.ContractBackend) (*RestrictToContract, error) {
	contract, err := bindRestrictToContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &RestrictToContract{RestrictToContractCaller: RestrictToContractCaller{contract: contract}, RestrictToContractTransactor: RestrictToContractTransactor{contract: contract}, RestrictToContractFilterer: RestrictToContractFilterer{contract: contract}}, nil
}

// NewRestrictToContractCaller creates a new read-only instance of RestrictToContract, bound to a specific deployed contract.
func NewRestrictToContractCaller(address common.Address, caller bind.ContractCaller) (*RestrictToContractCaller, error) {
	contract, err := bindRestrictToContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RestrictToContractCaller{contract: contract}, nil
}

// NewRestrictToContractTransactor creates a new write-only instance of RestrictToContract, bound to a specific deployed contract.
func NewRestrictToContractTransactor(address common.Address, transactor bind.ContractTransactor) (*RestrictToContractTransactor, error) {
	contract, err := bindRestrictToContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RestrictToContractTransactor{contract: contract}, nil
}

// NewRestrictToContractFilterer creates a new log filterer instance of RestrictToContract, bound to a specific deployed contract.
func NewRestrictToContractFilterer(address common.Address, filterer bind.ContractFilterer) (*RestrictToContractFilterer, error) {
	contract, err := bindRestrictToContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RestrictToContractFilterer{contract: contract}, nil
}

// bindRestrictToContract binds a generic wrapper to an already deployed contract.
func bindRestrictToContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(RestrictToContractABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RestrictToContract *RestrictToContractRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _RestrictToContract.Contract.RestrictToContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RestrictToContract *RestrictToContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RestrictToContract.Contract.RestrictToContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RestrictToContract *RestrictToContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RestrictToContract.Contract.RestrictToContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RestrictToContract *RestrictToContractCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _RestrictToContract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RestrictToContract *RestrictToContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RestrictToContract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RestrictToContract *RestrictToContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RestrictToContract.Contract.contract.Transact(opts, method, params...)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_RestrictToContract *RestrictToContractCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _RestrictToContract.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_RestrictToContract *RestrictToContractSession) Owner() (common.Address, error) {
	return _RestrictToContract.Contract.Owner(&_RestrictToContract.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_RestrictToContract *RestrictToContractCallerSession) Owner() (common.Address, error) {
	return _RestrictToContract.Contract.Owner(&_RestrictToContract.CallOpts)
}

// Check is a paid mutator transaction binding the contract method 0xb762c76d.
//
// Solidity: function check(initiator address, from address, to address, toKind uint8, tokens uint256, store address) returns()
func (_RestrictToContract *RestrictToContractTransactor) Check(opts *bind.TransactOpts, initiator common.Address, from common.Address, to common.Address, toKind uint8, tokens *big.Int, store common.Address) (*types.Transaction, error) {
	return _RestrictToContract.contract.Transact(opts, "check", initiator, from, to, toKind, tokens, store)
}

// Check is a paid mutator transaction binding the contract method 0xb762c76d.
//
// Solidity: function check(initiator address, from address, to address, toKind uint8, tokens uint256, store address) returns()
func (_RestrictToContract *RestrictToContractSession) Check(initiator common.Address, from common.Address, to common.Address, toKind uint8, tokens *big.Int, store common.Address) (*types.Transaction, error) {
	return _RestrictToContract.Contract.Check(&_RestrictToContract.TransactOpts, initiator, from, to, toKind, tokens, store)
}

// Check is a paid mutator transaction binding the contract method 0xb762c76d.
//
// Solidity: function check(initiator address, from address, to address, toKind uint8, tokens uint256, store address) returns()
func (_RestrictToContract *RestrictToContractTransactorSession) Check(initiator common.Address, from common.Address, to common.Address, toKind uint8, tokens *big.Int, store common.Address) (*types.Transaction, error) {
	return _RestrictToContract.Contract.Check(&_RestrictToContract.TransactOpts, initiator, from, to, toKind, tokens, store)
}

// Kill is a paid mutator transaction binding the contract method 0x41c0e1b5.
//
// Solidity: function kill() returns()
func (_RestrictToContract *RestrictToContractTransactor) Kill(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RestrictToContract.contract.Transact(opts, "kill")
}

// Kill is a paid mutator transaction binding the contract method 0x41c0e1b5.
//
// Solidity: function kill() returns()
func (_RestrictToContract *RestrictToContractSession) Kill() (*types.Transaction, error) {
	return _RestrictToContract.Contract.Kill(&_RestrictToContract.TransactOpts)
}

// Kill is a paid mutator transaction binding the contract method 0x41c0e1b5.
//
// Solidity: function kill() returns()
func (_RestrictToContract *RestrictToContractTransactorSession) Kill() (*types.Transaction, error) {
	return _RestrictToContract.Contract.Kill(&_RestrictToContract.TransactOpts)
}

// TransferOwner is a paid mutator transaction binding the contract method 0x4fb2e45d.
//
// Solidity: function transferOwner(newOwner address) returns()
func (_RestrictToContract *RestrictToContractTransactor) TransferOwner(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _RestrictToContract.contract.Transact(opts, "transferOwner", newOwner)
}

// TransferOwner is a paid mutator transaction binding the contract method 0x4fb2e45d.
//
// Solidity: function transferOwner(newOwner address) returns()
func (_RestrictToContract *RestrictToContractSession) TransferOwner(newOwner common.Address) (*types.Transaction, error) {
	return _RestrictToContract.Contract.TransferOwner(&_RestrictToContract.TransactOpts, newOwner)
}

// TransferOwner is a paid mutator transaction binding the contract method 0x4fb2e45d.
//
// Solidity: function transferOwner(newOwner address) returns()
func (_RestrictToContract *RestrictToContractTransactorSession) TransferOwner(newOwner common.Address) (*types.Transaction, error) {
	return _RestrictToContract.Contract.TransferOwner(&_RestrictToContract.TransactOpts, newOwner)
}

// RestrictToContractOwnerTransferredIterator is returned from FilterOwnerTransferred and is used to iterate over the raw logs and unpacked data for OwnerTransferred events raised by the RestrictToContract contract.
type RestrictToContractOwnerTransferredIterator struct {
	Event *RestrictToContractOwnerTransferred // Event containing the contract specifics and raw log

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
func (it *RestrictToContractOwnerTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RestrictToContractOwnerTransferred)
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
		it.Event = new(RestrictToContractOwnerTransferred)
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
func (it *RestrictToContractOwnerTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RestrictToContractOwnerTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RestrictToContractOwnerTransferred represents a OwnerTransferred event raised by the RestrictToContract contract.
type RestrictToContractOwnerTransferred struct {
	OldOwner common.Address
	NewOwner common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOwnerTransferred is a free log retrieval operation binding the contract event 0x8934ce4adea8d9ce0d714d2c22b86790e41b7731c84b926fbbdc1d40ff6533c9.
//
// Solidity: e OwnerTransferred(oldOwner indexed address, newOwner indexed address)
func (_RestrictToContract *RestrictToContractFilterer) FilterOwnerTransferred(opts *bind.FilterOpts, oldOwner []common.Address, newOwner []common.Address) (*RestrictToContractOwnerTransferredIterator, error) {

	var oldOwnerRule []interface{}
	for _, oldOwnerItem := range oldOwner {
		oldOwnerRule = append(oldOwnerRule, oldOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _RestrictToContract.contract.FilterLogs(opts, "OwnerTransferred", oldOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &RestrictToContractOwnerTransferredIterator{contract: _RestrictToContract.contract, event: "OwnerTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnerTransferred is a free log subscription operation binding the contract event 0x8934ce4adea8d9ce0d714d2c22b86790e41b7731c84b926fbbdc1d40ff6533c9.
//
// Solidity: e OwnerTransferred(oldOwner indexed address, newOwner indexed address)
func (_RestrictToContract *RestrictToContractFilterer) WatchOwnerTransferred(opts *bind.WatchOpts, sink chan<- *RestrictToContractOwnerTransferred, oldOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var oldOwnerRule []interface{}
	for _, oldOwnerItem := range oldOwner {
		oldOwnerRule = append(oldOwnerRule, oldOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _RestrictToContract.contract.WatchLogs(opts, "OwnerTransferred", oldOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RestrictToContractOwnerTransferred)
				if err := _RestrictToContract.contract.UnpackLog(event, "OwnerTransferred", log); err != nil {
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
