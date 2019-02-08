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

// RestrictABI is the input ABI used to generate the binding from.
const RestrictABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"isRestricted\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"restricted\",\"type\":\"bool\"}],\"name\":\"setRestricted\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"kill\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwner\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"administrable\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"initiator\",\"type\":\"address\"},{\"name\":\"from\",\"type\":\"address\"},{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"toKind\",\"type\":\"uint8\"},{\"name\":\"tokens\",\"type\":\"uint256\"},{\"name\":\"store\",\"type\":\"address\"}],\"name\":\"check\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"a\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"oldOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnerTransferred\",\"type\":\"event\"}]"

// RestrictBin is the compiled bytecode used for deploying new contracts.
const RestrictBin = `60806040526000805460a060020a60ff021916905534801561002057600080fd5b506040516020806107618339810180604052602081101561004057600080fd5b505160008054600160a060020a0319908116331790915560018054600160a060020a03909316929091169190911790556106e28061007f6000396000f3fe608060405234801561001057600080fd5b506004361061009a576000357c0100000000000000000000000000000000000000000000000000000000900480634fb2e45d116100785780634fb2e45d146100e4578063834b1ab0146101175780638da5cb5b14610148578063b762c76d146101505761009a565b80631f5e8f4c1461009f57806324daddc5146100bb57806341c0e1b5146100dc575b600080fd5b6100a76101aa565b604080519115158252519081900360200190f35b6100da600480360360208110156100d157600080fd5b503515156101cb565b005b6100da610309565b6100da600480360360208110156100fa57600080fd5b503573ffffffffffffffffffffffffffffffffffffffff166103aa565b61011f61059d565b6040805173ffffffffffffffffffffffffffffffffffffffff9092168252519081900360200190f35b61011f6105b9565b6100da600480360360c081101561016657600080fd5b5073ffffffffffffffffffffffffffffffffffffffff81358116916020810135821691604082013581169160ff6060820135169160808201359160a00135166105d5565b60005474010000000000000000000000000000000000000000900460ff1681565b600154604080517f24d7806c000000000000000000000000000000000000000000000000000000008152336004820152905173ffffffffffffffffffffffffffffffffffffffff909216916324d7806c91602480820192602092909190829003018186803b15801561023c57600080fd5b505afa158015610250573d6000803e3d6000fd5b505050506040513d602081101561026657600080fd5b505115156102bf576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252602a81526020018061068d602a913960400191505060405180910390fd5b6000805491151574010000000000000000000000000000000000000000027fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff909216919091179055565b60005473ffffffffffffffffffffffffffffffffffffffff16331461038f57604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601960248201527f4f776e6572206163636f756e7420697320726571756972656400000000000000604482015290519081900360640190fd5b60005473ffffffffffffffffffffffffffffffffffffffff16ff5b60005473ffffffffffffffffffffffffffffffffffffffff16331461043057604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601960248201527f4f776e6572206163636f756e7420697320726571756972656400000000000000604482015290519081900360640190fd5b60005473ffffffffffffffffffffffffffffffffffffffff828116911614156104a4576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260258152602001806106686025913960400191505060405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff8116151561052857604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820181905260248201527f4e6577204f776e65722063616e6e6f74206265207a65726f2061646472657373604482015290519081900360640190fd5b6000805473ffffffffffffffffffffffffffffffffffffffff8381167fffffffffffffffffffffffff0000000000000000000000000000000000000000831681178455604051919092169283917f8934ce4adea8d9ce0d714d2c22b86790e41b7731c84b926fbbdc1d40ff6533c99190a35050565b60015473ffffffffffffffffffffffffffffffffffffffff1681565b60005473ffffffffffffffffffffffffffffffffffffffff1681565b60005474010000000000000000000000000000000000000000900460ff161561065f57604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820181905260248201527f5265737472696374696f6e2069732063757272656e746c7920656e61626c6564604482015290519081900360640190fd5b50505050505056fe4e6577204f776e65722063616e6e6f74206265207468652063757272656e74206f776e65724f6e6c7920616e2061646d696e2063616e20736574207468652072657374726963746564207374617465a165627a7a723058208b0d10614075b7f18b9e892a212f6aa5d04c0c89fa20a5fec366c6de57ed7a620029`

// DeployRestrict deploys a new Ethereum contract, binding an instance of Restrict to it.
func DeployRestrict(auth *bind.TransactOpts, backend bind.ContractBackend, a common.Address) (common.Address, *types.Transaction, *Restrict, error) {
	parsed, err := abi.JSON(strings.NewReader(RestrictABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(RestrictBin), backend, a)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Restrict{RestrictCaller: RestrictCaller{contract: contract}, RestrictTransactor: RestrictTransactor{contract: contract}, RestrictFilterer: RestrictFilterer{contract: contract}}, nil
}

// Restrict is an auto generated Go binding around an Ethereum contract.
type Restrict struct {
	RestrictCaller     // Read-only binding to the contract
	RestrictTransactor // Write-only binding to the contract
	RestrictFilterer   // Log filterer for contract events
}

// RestrictCaller is an auto generated read-only Go binding around an Ethereum contract.
type RestrictCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RestrictTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RestrictTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RestrictFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RestrictFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RestrictSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RestrictSession struct {
	Contract     *Restrict         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RestrictCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RestrictCallerSession struct {
	Contract *RestrictCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// RestrictTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RestrictTransactorSession struct {
	Contract     *RestrictTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// RestrictRaw is an auto generated low-level Go binding around an Ethereum contract.
type RestrictRaw struct {
	Contract *Restrict // Generic contract binding to access the raw methods on
}

// RestrictCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RestrictCallerRaw struct {
	Contract *RestrictCaller // Generic read-only contract binding to access the raw methods on
}

// RestrictTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RestrictTransactorRaw struct {
	Contract *RestrictTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRestrict creates a new instance of Restrict, bound to a specific deployed contract.
func NewRestrict(address common.Address, backend bind.ContractBackend) (*Restrict, error) {
	contract, err := bindRestrict(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Restrict{RestrictCaller: RestrictCaller{contract: contract}, RestrictTransactor: RestrictTransactor{contract: contract}, RestrictFilterer: RestrictFilterer{contract: contract}}, nil
}

// NewRestrictCaller creates a new read-only instance of Restrict, bound to a specific deployed contract.
func NewRestrictCaller(address common.Address, caller bind.ContractCaller) (*RestrictCaller, error) {
	contract, err := bindRestrict(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RestrictCaller{contract: contract}, nil
}

// NewRestrictTransactor creates a new write-only instance of Restrict, bound to a specific deployed contract.
func NewRestrictTransactor(address common.Address, transactor bind.ContractTransactor) (*RestrictTransactor, error) {
	contract, err := bindRestrict(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RestrictTransactor{contract: contract}, nil
}

// NewRestrictFilterer creates a new log filterer instance of Restrict, bound to a specific deployed contract.
func NewRestrictFilterer(address common.Address, filterer bind.ContractFilterer) (*RestrictFilterer, error) {
	contract, err := bindRestrict(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RestrictFilterer{contract: contract}, nil
}

// bindRestrict binds a generic wrapper to an already deployed contract.
func bindRestrict(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(RestrictABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Restrict *RestrictRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Restrict.Contract.RestrictCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Restrict *RestrictRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Restrict.Contract.RestrictTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Restrict *RestrictRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Restrict.Contract.RestrictTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Restrict *RestrictCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Restrict.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Restrict *RestrictTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Restrict.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Restrict *RestrictTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Restrict.Contract.contract.Transact(opts, method, params...)
}

// Administrable is a free data retrieval call binding the contract method 0x834b1ab0.
//
// Solidity: function administrable() constant returns(address)
func (_Restrict *RestrictCaller) Administrable(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Restrict.contract.Call(opts, out, "administrable")
	return *ret0, err
}

// Administrable is a free data retrieval call binding the contract method 0x834b1ab0.
//
// Solidity: function administrable() constant returns(address)
func (_Restrict *RestrictSession) Administrable() (common.Address, error) {
	return _Restrict.Contract.Administrable(&_Restrict.CallOpts)
}

// Administrable is a free data retrieval call binding the contract method 0x834b1ab0.
//
// Solidity: function administrable() constant returns(address)
func (_Restrict *RestrictCallerSession) Administrable() (common.Address, error) {
	return _Restrict.Contract.Administrable(&_Restrict.CallOpts)
}

// IsRestricted is a free data retrieval call binding the contract method 0x1f5e8f4c.
//
// Solidity: function isRestricted() constant returns(bool)
func (_Restrict *RestrictCaller) IsRestricted(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Restrict.contract.Call(opts, out, "isRestricted")
	return *ret0, err
}

// IsRestricted is a free data retrieval call binding the contract method 0x1f5e8f4c.
//
// Solidity: function isRestricted() constant returns(bool)
func (_Restrict *RestrictSession) IsRestricted() (bool, error) {
	return _Restrict.Contract.IsRestricted(&_Restrict.CallOpts)
}

// IsRestricted is a free data retrieval call binding the contract method 0x1f5e8f4c.
//
// Solidity: function isRestricted() constant returns(bool)
func (_Restrict *RestrictCallerSession) IsRestricted() (bool, error) {
	return _Restrict.Contract.IsRestricted(&_Restrict.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Restrict *RestrictCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Restrict.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Restrict *RestrictSession) Owner() (common.Address, error) {
	return _Restrict.Contract.Owner(&_Restrict.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Restrict *RestrictCallerSession) Owner() (common.Address, error) {
	return _Restrict.Contract.Owner(&_Restrict.CallOpts)
}

// Check is a paid mutator transaction binding the contract method 0xb762c76d.
//
// Solidity: function check(initiator address, from address, to address, toKind uint8, tokens uint256, store address) returns()
func (_Restrict *RestrictTransactor) Check(opts *bind.TransactOpts, initiator common.Address, from common.Address, to common.Address, toKind uint8, tokens *big.Int, store common.Address) (*types.Transaction, error) {
	return _Restrict.contract.Transact(opts, "check", initiator, from, to, toKind, tokens, store)
}

// Check is a paid mutator transaction binding the contract method 0xb762c76d.
//
// Solidity: function check(initiator address, from address, to address, toKind uint8, tokens uint256, store address) returns()
func (_Restrict *RestrictSession) Check(initiator common.Address, from common.Address, to common.Address, toKind uint8, tokens *big.Int, store common.Address) (*types.Transaction, error) {
	return _Restrict.Contract.Check(&_Restrict.TransactOpts, initiator, from, to, toKind, tokens, store)
}

// Check is a paid mutator transaction binding the contract method 0xb762c76d.
//
// Solidity: function check(initiator address, from address, to address, toKind uint8, tokens uint256, store address) returns()
func (_Restrict *RestrictTransactorSession) Check(initiator common.Address, from common.Address, to common.Address, toKind uint8, tokens *big.Int, store common.Address) (*types.Transaction, error) {
	return _Restrict.Contract.Check(&_Restrict.TransactOpts, initiator, from, to, toKind, tokens, store)
}

// Kill is a paid mutator transaction binding the contract method 0x41c0e1b5.
//
// Solidity: function kill() returns()
func (_Restrict *RestrictTransactor) Kill(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Restrict.contract.Transact(opts, "kill")
}

// Kill is a paid mutator transaction binding the contract method 0x41c0e1b5.
//
// Solidity: function kill() returns()
func (_Restrict *RestrictSession) Kill() (*types.Transaction, error) {
	return _Restrict.Contract.Kill(&_Restrict.TransactOpts)
}

// Kill is a paid mutator transaction binding the contract method 0x41c0e1b5.
//
// Solidity: function kill() returns()
func (_Restrict *RestrictTransactorSession) Kill() (*types.Transaction, error) {
	return _Restrict.Contract.Kill(&_Restrict.TransactOpts)
}

// SetRestricted is a paid mutator transaction binding the contract method 0x24daddc5.
//
// Solidity: function setRestricted(restricted bool) returns()
func (_Restrict *RestrictTransactor) SetRestricted(opts *bind.TransactOpts, restricted bool) (*types.Transaction, error) {
	return _Restrict.contract.Transact(opts, "setRestricted", restricted)
}

// SetRestricted is a paid mutator transaction binding the contract method 0x24daddc5.
//
// Solidity: function setRestricted(restricted bool) returns()
func (_Restrict *RestrictSession) SetRestricted(restricted bool) (*types.Transaction, error) {
	return _Restrict.Contract.SetRestricted(&_Restrict.TransactOpts, restricted)
}

// SetRestricted is a paid mutator transaction binding the contract method 0x24daddc5.
//
// Solidity: function setRestricted(restricted bool) returns()
func (_Restrict *RestrictTransactorSession) SetRestricted(restricted bool) (*types.Transaction, error) {
	return _Restrict.Contract.SetRestricted(&_Restrict.TransactOpts, restricted)
}

// TransferOwner is a paid mutator transaction binding the contract method 0x4fb2e45d.
//
// Solidity: function transferOwner(newOwner address) returns()
func (_Restrict *RestrictTransactor) TransferOwner(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Restrict.contract.Transact(opts, "transferOwner", newOwner)
}

// TransferOwner is a paid mutator transaction binding the contract method 0x4fb2e45d.
//
// Solidity: function transferOwner(newOwner address) returns()
func (_Restrict *RestrictSession) TransferOwner(newOwner common.Address) (*types.Transaction, error) {
	return _Restrict.Contract.TransferOwner(&_Restrict.TransactOpts, newOwner)
}

// TransferOwner is a paid mutator transaction binding the contract method 0x4fb2e45d.
//
// Solidity: function transferOwner(newOwner address) returns()
func (_Restrict *RestrictTransactorSession) TransferOwner(newOwner common.Address) (*types.Transaction, error) {
	return _Restrict.Contract.TransferOwner(&_Restrict.TransactOpts, newOwner)
}

// RestrictOwnerTransferredIterator is returned from FilterOwnerTransferred and is used to iterate over the raw logs and unpacked data for OwnerTransferred events raised by the Restrict contract.
type RestrictOwnerTransferredIterator struct {
	Event *RestrictOwnerTransferred // Event containing the contract specifics and raw log

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
func (it *RestrictOwnerTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RestrictOwnerTransferred)
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
		it.Event = new(RestrictOwnerTransferred)
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
func (it *RestrictOwnerTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RestrictOwnerTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RestrictOwnerTransferred represents a OwnerTransferred event raised by the Restrict contract.
type RestrictOwnerTransferred struct {
	OldOwner common.Address
	NewOwner common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOwnerTransferred is a free log retrieval operation binding the contract event 0x8934ce4adea8d9ce0d714d2c22b86790e41b7731c84b926fbbdc1d40ff6533c9.
//
// Solidity: e OwnerTransferred(oldOwner indexed address, newOwner indexed address)
func (_Restrict *RestrictFilterer) FilterOwnerTransferred(opts *bind.FilterOpts, oldOwner []common.Address, newOwner []common.Address) (*RestrictOwnerTransferredIterator, error) {

	var oldOwnerRule []interface{}
	for _, oldOwnerItem := range oldOwner {
		oldOwnerRule = append(oldOwnerRule, oldOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Restrict.contract.FilterLogs(opts, "OwnerTransferred", oldOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &RestrictOwnerTransferredIterator{contract: _Restrict.contract, event: "OwnerTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnerTransferred is a free log subscription operation binding the contract event 0x8934ce4adea8d9ce0d714d2c22b86790e41b7731c84b926fbbdc1d40ff6533c9.
//
// Solidity: e OwnerTransferred(oldOwner indexed address, newOwner indexed address)
func (_Restrict *RestrictFilterer) WatchOwnerTransferred(opts *bind.WatchOpts, sink chan<- *RestrictOwnerTransferred, oldOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var oldOwnerRule []interface{}
	for _, oldOwnerItem := range oldOwner {
		oldOwnerRule = append(oldOwnerRule, oldOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Restrict.contract.WatchLogs(opts, "OwnerTransferred", oldOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RestrictOwnerTransferred)
				if err := _Restrict.contract.UnpackLog(event, "OwnerTransferred", log); err != nil {
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
