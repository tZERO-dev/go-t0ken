// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ownership

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

// AdministrableABI is the input ABI used to generate the binding from.
const AdministrableABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"admin\",\"type\":\"address\"}],\"name\":\"removeAdmin\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"isAdmin\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwner\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"admin\",\"type\":\"address\"}],\"name\":\"addAdmin\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"admins\",\"outputs\":[{\"name\":\"count\",\"type\":\"int256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"admin\",\"type\":\"address\"}],\"name\":\"AdminAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"admin\",\"type\":\"address\"}],\"name\":\"AdminRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"oldOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnerTransferred\",\"type\":\"event\"}]"

// AdministrableBin is the compiled bytecode used for deploying new contracts.
const AdministrableBin = `608060405260008054600160a060020a03191633179055610933806100256000396000f3fe608060405234801561001057600080fd5b506004361061008f576000357c010000000000000000000000000000000000000000000000000000000090048063704802751161006d57806370480275146101435780638da5cb5b14610176578063a5de3619146101a75761008f565b80631785f53c1461009457806324d7806c146100c95780634fb2e45d14610110575b600080fd5b6100c7600480360360208110156100aa57600080fd5b503573ffffffffffffffffffffffffffffffffffffffff166101c1565b005b6100fc600480360360208110156100df57600080fd5b503573ffffffffffffffffffffffffffffffffffffffff16610309565b604080519115158252519081900360200190f35b6100c76004803603602081101561012657600080fd5b503573ffffffffffffffffffffffffffffffffffffffff16610345565b6100c76004803603602081101561015957600080fd5b503573ffffffffffffffffffffffffffffffffffffffff1661052d565b61017e610675565b6040805173ffffffffffffffffffffffffffffffffffffffff9092168252519081900360200190f35b6101af610691565b60408051918252519081900360200190f35b60005473ffffffffffffffffffffffffffffffffffffffff16331461024757604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601960248201527f4f776e6572206163636f756e7420697320726571756972656400000000000000604482015290519081900360640190fd5b61025860018263ffffffff61069716565b15156102c557604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f556e61626c6520746f2072656d6f76652061646d696e00000000000000000000604482015290519081900360640190fd5b60405173ffffffffffffffffffffffffffffffffffffffff8216907fa3b62bc36326052d97ea62d63c3d60308ed4c3ea8ac079dd8499f1e9c4f80c0f90600090a250565b6000805473ffffffffffffffffffffffffffffffffffffffff8381169116148061033f575061033f60018363ffffffff6107c316565b92915050565b60005473ffffffffffffffffffffffffffffffffffffffff1633146103cb57604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601960248201527f4f776e6572206163636f756e7420697320726571756972656400000000000000604482015290519081900360640190fd5b60005473ffffffffffffffffffffffffffffffffffffffff8281169116141561043f576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260258152602001806108e36025913960400191505060405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff811615156104c357604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820181905260248201527f4e6577204f776e65722063616e6e6f74206265207a65726f2061646472657373604482015290519081900360640190fd5b6000805473ffffffffffffffffffffffffffffffffffffffff83811673ffffffffffffffffffffffffffffffffffffffff19831681178455604051919092169283917f8934ce4adea8d9ce0d714d2c22b86790e41b7731c84b926fbbdc1d40ff6533c99190a35050565b60005473ffffffffffffffffffffffffffffffffffffffff1633146105b357604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601960248201527f4f776e6572206163636f756e7420697320726571756972656400000000000000604482015290519081900360640190fd5b6105c460018263ffffffff61080716565b151561063157604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601360248201527f556e61626c6520746f206164642061646d696e00000000000000000000000000604482015290519081900360640190fd5b60405173ffffffffffffffffffffffffffffffffffffffff8216907f44d6d25963f097ad14f29f06854a01f575648a1ef82f30e562ccd3889717e33990600090a250565b60005473ffffffffffffffffffffffffffffffffffffffff1681565b60015481565b73ffffffffffffffffffffffffffffffffffffffff8116600090815260018084016020526040822054908112806106ce5750835481135b156106dd57600091505061033f565b835481121561075e578354600090815260028501602081815260408084205473ffffffffffffffffffffffffffffffffffffffff1680855260018901835281852086905585855292909152808320805473ffffffffffffffffffffffffffffffffffffffff199081169093179055865483529091208054909116905561078a565b60008181526002850160205260409020805473ffffffffffffffffffffffffffffffffffffffff191690555b505073ffffffffffffffffffffffffffffffffffffffff1660009081526001828101602052604082209190915581546000190190915590565b73ffffffffffffffffffffffffffffffffffffffff81166000908152600183016020526040812054600019018181128015906107ff5750835481125b949350505050565b600073ffffffffffffffffffffffffffffffffffffffff8216151561082e5750600061033f565b73ffffffffffffffffffffffffffffffffffffffff821660009081526001840160205260408120546000190190811280159061086a5750835481125b1561087957600091505061033f565b50508154600190810180845573ffffffffffffffffffffffffffffffffffffffff831660008181528386016020908152604080832085905593825260028701905291909120805473ffffffffffffffffffffffffffffffffffffffff191690911790559291505056fe4e6577204f776e65722063616e6e6f74206265207468652063757272656e74206f776e6572a165627a7a7230582048df02d79a930d3d46104a168989f1786d5bdb9a74e554c3ad0a69426de402770029`

// DeployAdministrable deploys a new Ethereum contract, binding an instance of Administrable to it.
func DeployAdministrable(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Administrable, error) {
	parsed, err := abi.JSON(strings.NewReader(AdministrableABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(AdministrableBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Administrable{AdministrableCaller: AdministrableCaller{contract: contract}, AdministrableTransactor: AdministrableTransactor{contract: contract}, AdministrableFilterer: AdministrableFilterer{contract: contract}}, nil
}

// Administrable is an auto generated Go binding around an Ethereum contract.
type Administrable struct {
	AdministrableCaller     // Read-only binding to the contract
	AdministrableTransactor // Write-only binding to the contract
	AdministrableFilterer   // Log filterer for contract events
}

// AdministrableCaller is an auto generated read-only Go binding around an Ethereum contract.
type AdministrableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AdministrableTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AdministrableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AdministrableFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AdministrableFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AdministrableSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AdministrableSession struct {
	Contract     *Administrable    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AdministrableCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AdministrableCallerSession struct {
	Contract *AdministrableCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// AdministrableTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AdministrableTransactorSession struct {
	Contract     *AdministrableTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// AdministrableRaw is an auto generated low-level Go binding around an Ethereum contract.
type AdministrableRaw struct {
	Contract *Administrable // Generic contract binding to access the raw methods on
}

// AdministrableCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AdministrableCallerRaw struct {
	Contract *AdministrableCaller // Generic read-only contract binding to access the raw methods on
}

// AdministrableTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AdministrableTransactorRaw struct {
	Contract *AdministrableTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAdministrable creates a new instance of Administrable, bound to a specific deployed contract.
func NewAdministrable(address common.Address, backend bind.ContractBackend) (*Administrable, error) {
	contract, err := bindAdministrable(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Administrable{AdministrableCaller: AdministrableCaller{contract: contract}, AdministrableTransactor: AdministrableTransactor{contract: contract}, AdministrableFilterer: AdministrableFilterer{contract: contract}}, nil
}

// NewAdministrableCaller creates a new read-only instance of Administrable, bound to a specific deployed contract.
func NewAdministrableCaller(address common.Address, caller bind.ContractCaller) (*AdministrableCaller, error) {
	contract, err := bindAdministrable(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AdministrableCaller{contract: contract}, nil
}

// NewAdministrableTransactor creates a new write-only instance of Administrable, bound to a specific deployed contract.
func NewAdministrableTransactor(address common.Address, transactor bind.ContractTransactor) (*AdministrableTransactor, error) {
	contract, err := bindAdministrable(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AdministrableTransactor{contract: contract}, nil
}

// NewAdministrableFilterer creates a new log filterer instance of Administrable, bound to a specific deployed contract.
func NewAdministrableFilterer(address common.Address, filterer bind.ContractFilterer) (*AdministrableFilterer, error) {
	contract, err := bindAdministrable(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AdministrableFilterer{contract: contract}, nil
}

// bindAdministrable binds a generic wrapper to an already deployed contract.
func bindAdministrable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AdministrableABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Administrable *AdministrableRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Administrable.Contract.AdministrableCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Administrable *AdministrableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Administrable.Contract.AdministrableTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Administrable *AdministrableRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Administrable.Contract.AdministrableTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Administrable *AdministrableCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Administrable.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Administrable *AdministrableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Administrable.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Administrable *AdministrableTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Administrable.Contract.contract.Transact(opts, method, params...)
}

// Admins is a free data retrieval call binding the contract method 0xa5de3619.
//
// Solidity: function admins() constant returns(count int256)
func (_Administrable *AdministrableCaller) Admins(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Administrable.contract.Call(opts, out, "admins")
	return *ret0, err
}

// Admins is a free data retrieval call binding the contract method 0xa5de3619.
//
// Solidity: function admins() constant returns(count int256)
func (_Administrable *AdministrableSession) Admins() (*big.Int, error) {
	return _Administrable.Contract.Admins(&_Administrable.CallOpts)
}

// Admins is a free data retrieval call binding the contract method 0xa5de3619.
//
// Solidity: function admins() constant returns(count int256)
func (_Administrable *AdministrableCallerSession) Admins() (*big.Int, error) {
	return _Administrable.Contract.Admins(&_Administrable.CallOpts)
}

// IsAdmin is a free data retrieval call binding the contract method 0x24d7806c.
//
// Solidity: function isAdmin(addr address) constant returns(bool)
func (_Administrable *AdministrableCaller) IsAdmin(opts *bind.CallOpts, addr common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Administrable.contract.Call(opts, out, "isAdmin", addr)
	return *ret0, err
}

// IsAdmin is a free data retrieval call binding the contract method 0x24d7806c.
//
// Solidity: function isAdmin(addr address) constant returns(bool)
func (_Administrable *AdministrableSession) IsAdmin(addr common.Address) (bool, error) {
	return _Administrable.Contract.IsAdmin(&_Administrable.CallOpts, addr)
}

// IsAdmin is a free data retrieval call binding the contract method 0x24d7806c.
//
// Solidity: function isAdmin(addr address) constant returns(bool)
func (_Administrable *AdministrableCallerSession) IsAdmin(addr common.Address) (bool, error) {
	return _Administrable.Contract.IsAdmin(&_Administrable.CallOpts, addr)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Administrable *AdministrableCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Administrable.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Administrable *AdministrableSession) Owner() (common.Address, error) {
	return _Administrable.Contract.Owner(&_Administrable.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Administrable *AdministrableCallerSession) Owner() (common.Address, error) {
	return _Administrable.Contract.Owner(&_Administrable.CallOpts)
}

// AddAdmin is a paid mutator transaction binding the contract method 0x70480275.
//
// Solidity: function addAdmin(admin address) returns()
func (_Administrable *AdministrableTransactor) AddAdmin(opts *bind.TransactOpts, admin common.Address) (*types.Transaction, error) {
	return _Administrable.contract.Transact(opts, "addAdmin", admin)
}

// AddAdmin is a paid mutator transaction binding the contract method 0x70480275.
//
// Solidity: function addAdmin(admin address) returns()
func (_Administrable *AdministrableSession) AddAdmin(admin common.Address) (*types.Transaction, error) {
	return _Administrable.Contract.AddAdmin(&_Administrable.TransactOpts, admin)
}

// AddAdmin is a paid mutator transaction binding the contract method 0x70480275.
//
// Solidity: function addAdmin(admin address) returns()
func (_Administrable *AdministrableTransactorSession) AddAdmin(admin common.Address) (*types.Transaction, error) {
	return _Administrable.Contract.AddAdmin(&_Administrable.TransactOpts, admin)
}

// RemoveAdmin is a paid mutator transaction binding the contract method 0x1785f53c.
//
// Solidity: function removeAdmin(admin address) returns()
func (_Administrable *AdministrableTransactor) RemoveAdmin(opts *bind.TransactOpts, admin common.Address) (*types.Transaction, error) {
	return _Administrable.contract.Transact(opts, "removeAdmin", admin)
}

// RemoveAdmin is a paid mutator transaction binding the contract method 0x1785f53c.
//
// Solidity: function removeAdmin(admin address) returns()
func (_Administrable *AdministrableSession) RemoveAdmin(admin common.Address) (*types.Transaction, error) {
	return _Administrable.Contract.RemoveAdmin(&_Administrable.TransactOpts, admin)
}

// RemoveAdmin is a paid mutator transaction binding the contract method 0x1785f53c.
//
// Solidity: function removeAdmin(admin address) returns()
func (_Administrable *AdministrableTransactorSession) RemoveAdmin(admin common.Address) (*types.Transaction, error) {
	return _Administrable.Contract.RemoveAdmin(&_Administrable.TransactOpts, admin)
}

// TransferOwner is a paid mutator transaction binding the contract method 0x4fb2e45d.
//
// Solidity: function transferOwner(newOwner address) returns()
func (_Administrable *AdministrableTransactor) TransferOwner(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Administrable.contract.Transact(opts, "transferOwner", newOwner)
}

// TransferOwner is a paid mutator transaction binding the contract method 0x4fb2e45d.
//
// Solidity: function transferOwner(newOwner address) returns()
func (_Administrable *AdministrableSession) TransferOwner(newOwner common.Address) (*types.Transaction, error) {
	return _Administrable.Contract.TransferOwner(&_Administrable.TransactOpts, newOwner)
}

// TransferOwner is a paid mutator transaction binding the contract method 0x4fb2e45d.
//
// Solidity: function transferOwner(newOwner address) returns()
func (_Administrable *AdministrableTransactorSession) TransferOwner(newOwner common.Address) (*types.Transaction, error) {
	return _Administrable.Contract.TransferOwner(&_Administrable.TransactOpts, newOwner)
}

// AdministrableAdminAddedIterator is returned from FilterAdminAdded and is used to iterate over the raw logs and unpacked data for AdminAdded events raised by the Administrable contract.
type AdministrableAdminAddedIterator struct {
	Event *AdministrableAdminAdded // Event containing the contract specifics and raw log

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
func (it *AdministrableAdminAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AdministrableAdminAdded)
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
		it.Event = new(AdministrableAdminAdded)
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
func (it *AdministrableAdminAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AdministrableAdminAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AdministrableAdminAdded represents a AdminAdded event raised by the Administrable contract.
type AdministrableAdminAdded struct {
	Admin common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterAdminAdded is a free log retrieval operation binding the contract event 0x44d6d25963f097ad14f29f06854a01f575648a1ef82f30e562ccd3889717e339.
//
// Solidity: e AdminAdded(admin indexed address)
func (_Administrable *AdministrableFilterer) FilterAdminAdded(opts *bind.FilterOpts, admin []common.Address) (*AdministrableAdminAddedIterator, error) {

	var adminRule []interface{}
	for _, adminItem := range admin {
		adminRule = append(adminRule, adminItem)
	}

	logs, sub, err := _Administrable.contract.FilterLogs(opts, "AdminAdded", adminRule)
	if err != nil {
		return nil, err
	}
	return &AdministrableAdminAddedIterator{contract: _Administrable.contract, event: "AdminAdded", logs: logs, sub: sub}, nil
}

// WatchAdminAdded is a free log subscription operation binding the contract event 0x44d6d25963f097ad14f29f06854a01f575648a1ef82f30e562ccd3889717e339.
//
// Solidity: e AdminAdded(admin indexed address)
func (_Administrable *AdministrableFilterer) WatchAdminAdded(opts *bind.WatchOpts, sink chan<- *AdministrableAdminAdded, admin []common.Address) (event.Subscription, error) {

	var adminRule []interface{}
	for _, adminItem := range admin {
		adminRule = append(adminRule, adminItem)
	}

	logs, sub, err := _Administrable.contract.WatchLogs(opts, "AdminAdded", adminRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AdministrableAdminAdded)
				if err := _Administrable.contract.UnpackLog(event, "AdminAdded", log); err != nil {
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

// AdministrableAdminRemovedIterator is returned from FilterAdminRemoved and is used to iterate over the raw logs and unpacked data for AdminRemoved events raised by the Administrable contract.
type AdministrableAdminRemovedIterator struct {
	Event *AdministrableAdminRemoved // Event containing the contract specifics and raw log

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
func (it *AdministrableAdminRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AdministrableAdminRemoved)
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
		it.Event = new(AdministrableAdminRemoved)
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
func (it *AdministrableAdminRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AdministrableAdminRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AdministrableAdminRemoved represents a AdminRemoved event raised by the Administrable contract.
type AdministrableAdminRemoved struct {
	Admin common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterAdminRemoved is a free log retrieval operation binding the contract event 0xa3b62bc36326052d97ea62d63c3d60308ed4c3ea8ac079dd8499f1e9c4f80c0f.
//
// Solidity: e AdminRemoved(admin indexed address)
func (_Administrable *AdministrableFilterer) FilterAdminRemoved(opts *bind.FilterOpts, admin []common.Address) (*AdministrableAdminRemovedIterator, error) {

	var adminRule []interface{}
	for _, adminItem := range admin {
		adminRule = append(adminRule, adminItem)
	}

	logs, sub, err := _Administrable.contract.FilterLogs(opts, "AdminRemoved", adminRule)
	if err != nil {
		return nil, err
	}
	return &AdministrableAdminRemovedIterator{contract: _Administrable.contract, event: "AdminRemoved", logs: logs, sub: sub}, nil
}

// WatchAdminRemoved is a free log subscription operation binding the contract event 0xa3b62bc36326052d97ea62d63c3d60308ed4c3ea8ac079dd8499f1e9c4f80c0f.
//
// Solidity: e AdminRemoved(admin indexed address)
func (_Administrable *AdministrableFilterer) WatchAdminRemoved(opts *bind.WatchOpts, sink chan<- *AdministrableAdminRemoved, admin []common.Address) (event.Subscription, error) {

	var adminRule []interface{}
	for _, adminItem := range admin {
		adminRule = append(adminRule, adminItem)
	}

	logs, sub, err := _Administrable.contract.WatchLogs(opts, "AdminRemoved", adminRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AdministrableAdminRemoved)
				if err := _Administrable.contract.UnpackLog(event, "AdminRemoved", log); err != nil {
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

// AdministrableOwnerTransferredIterator is returned from FilterOwnerTransferred and is used to iterate over the raw logs and unpacked data for OwnerTransferred events raised by the Administrable contract.
type AdministrableOwnerTransferredIterator struct {
	Event *AdministrableOwnerTransferred // Event containing the contract specifics and raw log

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
func (it *AdministrableOwnerTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AdministrableOwnerTransferred)
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
		it.Event = new(AdministrableOwnerTransferred)
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
func (it *AdministrableOwnerTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AdministrableOwnerTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AdministrableOwnerTransferred represents a OwnerTransferred event raised by the Administrable contract.
type AdministrableOwnerTransferred struct {
	OldOwner common.Address
	NewOwner common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOwnerTransferred is a free log retrieval operation binding the contract event 0x8934ce4adea8d9ce0d714d2c22b86790e41b7731c84b926fbbdc1d40ff6533c9.
//
// Solidity: e OwnerTransferred(oldOwner indexed address, newOwner indexed address)
func (_Administrable *AdministrableFilterer) FilterOwnerTransferred(opts *bind.FilterOpts, oldOwner []common.Address, newOwner []common.Address) (*AdministrableOwnerTransferredIterator, error) {

	var oldOwnerRule []interface{}
	for _, oldOwnerItem := range oldOwner {
		oldOwnerRule = append(oldOwnerRule, oldOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Administrable.contract.FilterLogs(opts, "OwnerTransferred", oldOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &AdministrableOwnerTransferredIterator{contract: _Administrable.contract, event: "OwnerTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnerTransferred is a free log subscription operation binding the contract event 0x8934ce4adea8d9ce0d714d2c22b86790e41b7731c84b926fbbdc1d40ff6533c9.
//
// Solidity: e OwnerTransferred(oldOwner indexed address, newOwner indexed address)
func (_Administrable *AdministrableFilterer) WatchOwnerTransferred(opts *bind.WatchOpts, sink chan<- *AdministrableOwnerTransferred, oldOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var oldOwnerRule []interface{}
	for _, oldOwnerItem := range oldOwner {
		oldOwnerRule = append(oldOwnerRule, oldOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Administrable.contract.WatchLogs(opts, "OwnerTransferred", oldOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AdministrableOwnerTransferred)
				if err := _Administrable.contract.UnpackLog(event, "OwnerTransferred", log); err != nil {
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
