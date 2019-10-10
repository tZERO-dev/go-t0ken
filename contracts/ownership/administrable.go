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
const AdministrableABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"}],\"name\":\"AdminAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"}],\"name\":\"AdminRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnerTransferred\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[],\"name\":\"ZERO_ADDRESS\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"int256\",\"name\":\"index\",\"type\":\"int256\"}],\"name\":\"adminAt\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"admins\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"count\",\"type\":\"int256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"isAdmin\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"add\",\"type\":\"bool\"}],\"name\":\"setAdmin\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwner\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// AdministrableBin is the compiled bytecode used for deploying new contracts.
const AdministrableBin = `6080604052600180546001600160a01b031990811690915560008054909116331790556107de806100316000396000f3fe608060405234801561001057600080fd5b506004361061007d5760003560e01c80634fb2e45d1161005b5780634fb2e45d14610125578063538ba4f91461014b5780638da5cb5b14610153578063a5de36191461015b5761007d565b806324d7806c146100825780632bdbc56f146100bc5780634b0bddd2146100f5575b600080fd5b6100a86004803603602081101561009857600080fd5b50356001600160a01b0316610175565b604080519115158252519081900360200190f35b6100d9600480360360208110156100d257600080fd5b50356101a4565b604080516001600160a01b039092168252519081900360200190f35b6101236004803603604081101561010b57600080fd5b506001600160a01b03813516906020013515156101b7565b005b6101236004803603602081101561013b57600080fd5b50356001600160a01b031661036d565b6100d96104ec565b6100d96104fb565b61016361050a565b60408051918252519081900360200190f35b600080546001600160a01b038381169116148061019e575061019e60028363ffffffff61051016565b92915050565b600061019e60028363ffffffff61054716565b6000546001600160a01b03163314806101e057506001546000546001600160a01b039081169116145b610231576040805162461bcd60e51b815260206004820152601960248201527f4f776e6572206163636f756e7420697320726571756972656400000000000000604482015290519081900360640190fd5b80156102d25761024860028363ffffffff6105cd16565b610299576040805162461bcd60e51b815260206004820152601360248201527f556e61626c6520746f206164642061646d696e00000000000000000000000000604482015290519081900360640190fd5b6040516001600160a01b038316907f44d6d25963f097ad14f29f06854a01f575648a1ef82f30e562ccd3889717e33990600090a2610369565b6102e360028363ffffffff61067f16565b610334576040805162461bcd60e51b815260206004820152601660248201527f556e61626c6520746f2072656d6f76652061646d696e00000000000000000000604482015290519081900360640190fd5b6040516001600160a01b038316907fa3b62bc36326052d97ea62d63c3d60308ed4c3ea8ac079dd8499f1e9c4f80c0f90600090a25b5050565b6000546001600160a01b031633148061039657506001546000546001600160a01b039081169116145b6103e7576040805162461bcd60e51b815260206004820152601960248201527f4f776e6572206163636f756e7420697320726571756972656400000000000000604482015290519081900360640190fd5b6000546001600160a01b03828116911614156104345760405162461bcd60e51b81526004018080602001828103825260258152602001806107856025913960400191505060405180910390fd5b6001600160a01b03811661048f576040805162461bcd60e51b815260206004820181905260248201527f4e6577204f776e65722063616e6e6f74206265207a65726f2061646472657373604482015290519081900360640190fd5b600080546001600160a01b0383811673ffffffffffffffffffffffffffffffffffffffff19831681178455604051919092169283917f8934ce4adea8d9ce0d714d2c22b86790e41b7731c84b926fbbdc1d40ff6533c99190a35050565b6001546001600160a01b031681565b6000546001600160a01b031681565b60025481565b6001600160a01b03811660009081526001830160205260408120546000190181811280159061053f5750835481125b949350505050565b60008082121580156105595750825482125b6105aa576040805162461bcd60e51b815260206004820152601860248201527f496e646578206f757473696465206f6620626f756e64732e0000000000000000604482015290519081900360640190fd5b50600101600090815260029190910160205260409020546001600160a01b031690565b60006001600160a01b0382166105e55750600061019e565b6001600160a01b0382166000908152600184016020526040812054600019019081128015906106145750835481125b1561062357600091505061019e565b5050815460019081018084556001600160a01b03831660008181528386016020908152604080832085905593825260028701905291909120805473ffffffffffffffffffffffffffffffffffffffff1916909117905592915050565b6001600160a01b038116600090815260018084016020526040822054908112806106a95750835481135b156106b857600091505061019e565b835481121561072c57835460009081526002850160208181526040808420546001600160a01b031680855260018901835281852086905585855292909152808320805473ffffffffffffffffffffffffffffffffffffffff1990811690931790558654835290912080549091169055610758565b60008181526002850160205260409020805473ffffffffffffffffffffffffffffffffffffffff191690555b50506001600160a01b03166000908152600182810160205260408220919091558154600019019091559056fe4e6577204f776e65722063616e6e6f74206265207468652063757272656e74206f776e6572a265627a7a723158209891364a13f5c3fec247b5b2b803c42514b36107c85606817a9d913beab4c94264736f6c634300050c0032`

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

// ZEROADDRESS is a free data retrieval call binding the contract method 0x538ba4f9.
//
// Solidity: function ZERO_ADDRESS() constant returns(address)
func (_Administrable *AdministrableCaller) ZEROADDRESS(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Administrable.contract.Call(opts, out, "ZERO_ADDRESS")
	return *ret0, err
}

// ZEROADDRESS is a free data retrieval call binding the contract method 0x538ba4f9.
//
// Solidity: function ZERO_ADDRESS() constant returns(address)
func (_Administrable *AdministrableSession) ZEROADDRESS() (common.Address, error) {
	return _Administrable.Contract.ZEROADDRESS(&_Administrable.CallOpts)
}

// ZEROADDRESS is a free data retrieval call binding the contract method 0x538ba4f9.
//
// Solidity: function ZERO_ADDRESS() constant returns(address)
func (_Administrable *AdministrableCallerSession) ZEROADDRESS() (common.Address, error) {
	return _Administrable.Contract.ZEROADDRESS(&_Administrable.CallOpts)
}

// AdminAt is a free data retrieval call binding the contract method 0x2bdbc56f.
//
// Solidity: function adminAt(index int256) constant returns(address)
func (_Administrable *AdministrableCaller) AdminAt(opts *bind.CallOpts, index *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Administrable.contract.Call(opts, out, "adminAt", index)
	return *ret0, err
}

// AdminAt is a free data retrieval call binding the contract method 0x2bdbc56f.
//
// Solidity: function adminAt(index int256) constant returns(address)
func (_Administrable *AdministrableSession) AdminAt(index *big.Int) (common.Address, error) {
	return _Administrable.Contract.AdminAt(&_Administrable.CallOpts, index)
}

// AdminAt is a free data retrieval call binding the contract method 0x2bdbc56f.
//
// Solidity: function adminAt(index int256) constant returns(address)
func (_Administrable *AdministrableCallerSession) AdminAt(index *big.Int) (common.Address, error) {
	return _Administrable.Contract.AdminAt(&_Administrable.CallOpts, index)
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

// SetAdmin is a paid mutator transaction binding the contract method 0x4b0bddd2.
//
// Solidity: function setAdmin(addr address, add bool) returns()
func (_Administrable *AdministrableTransactor) SetAdmin(opts *bind.TransactOpts, addr common.Address, add bool) (*types.Transaction, error) {
	return _Administrable.contract.Transact(opts, "setAdmin", addr, add)
}

// SetAdmin is a paid mutator transaction binding the contract method 0x4b0bddd2.
//
// Solidity: function setAdmin(addr address, add bool) returns()
func (_Administrable *AdministrableSession) SetAdmin(addr common.Address, add bool) (*types.Transaction, error) {
	return _Administrable.Contract.SetAdmin(&_Administrable.TransactOpts, addr, add)
}

// SetAdmin is a paid mutator transaction binding the contract method 0x4b0bddd2.
//
// Solidity: function setAdmin(addr address, add bool) returns()
func (_Administrable *AdministrableTransactorSession) SetAdmin(addr common.Address, add bool) (*types.Transaction, error) {
	return _Administrable.Contract.SetAdmin(&_Administrable.TransactOpts, addr, add)
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
