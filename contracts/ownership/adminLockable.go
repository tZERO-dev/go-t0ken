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

// AdminLockableABI is the input ABI used to generate the binding from.
const AdminLockableABI = "[{\"constant\":false,\"inputs\":[{\"internalType\":\"bool\",\"name\":\"locked\",\"type\":\"bool\"}],\"name\":\"setLocked\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"isAdmin\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"int256\",\"name\":\"index\",\"type\":\"int256\"}],\"name\":\"adminAt\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"add\",\"type\":\"bool\"}],\"name\":\"setAdmin\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwner\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ZERO_ADDRESS\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isLocked\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"admins\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"count\",\"type\":\"int256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"}],\"name\":\"AdminAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"}],\"name\":\"AdminRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnerTransferred\",\"type\":\"event\"}]"

// AdminLockableBin is the compiled bytecode used for deploying new contracts.
const AdminLockableBin = `6080604052600180546001600160a01b031916905534801561002057600080fd5b50600080546001600160a01b031916331790556005805460ff1916905561092d8061004c6000396000f3fe608060405234801561001057600080fd5b50600436106100a35760003560e01c80634fb2e45d116100765780638da5cb5b1161005b5780638da5cb5b14610198578063a4e2d634146101a0578063a5de3619146101a8576100a3565b80634fb2e45d1461016a578063538ba4f914610190576100a3565b8063211e28b6146100a857806324d7806c146100c95780632bdbc56f146101035780634b0bddd21461013c575b600080fd5b6100c7600480360360208110156100be57600080fd5b503515156101c2565b005b6100ef600480360360208110156100df57600080fd5b50356001600160a01b0316610293565b604080519115158252519081900360200190f35b6101206004803603602081101561011957600080fd5b50356102c2565b604080516001600160a01b039092168252519081900360200190f35b6100c76004803603604081101561015257600080fd5b506001600160a01b03813516906020013515156102d5565b6100c76004803603602081101561018057600080fd5b50356001600160a01b031661048b565b61012061060a565b610120610619565b6100ef610628565b6101b0610631565b60408051918252519081900360200190f35b6000546001600160a01b03163314806101e757506101e760023363ffffffff61063716565b610238576040805162461bcd60e51b815260206004820152601960248201527f41646d696e206163636f756e7420697320726571756972656400000000000000604482015290519081900360640190fd5b60055460ff16151581151514156102805760405162461bcd60e51b81526004018080602001828103825260288152602001806108d16028913960400191505060405180910390fd5b6005805460ff1916911515919091179055565b600080546001600160a01b03838116911614806102bc57506102bc60028363ffffffff61063716565b92915050565b60006102bc60028363ffffffff61066e16565b6000546001600160a01b03163314806102fe57506001546000546001600160a01b039081169116145b61034f576040805162461bcd60e51b815260206004820152601960248201527f4f776e6572206163636f756e7420697320726571756972656400000000000000604482015290519081900360640190fd5b80156103f05761036660028363ffffffff6106f416565b6103b7576040805162461bcd60e51b815260206004820152601360248201527f556e61626c6520746f206164642061646d696e00000000000000000000000000604482015290519081900360640190fd5b6040516001600160a01b038316907f44d6d25963f097ad14f29f06854a01f575648a1ef82f30e562ccd3889717e33990600090a2610487565b61040160028363ffffffff6107a616565b610452576040805162461bcd60e51b815260206004820152601660248201527f556e61626c6520746f2072656d6f76652061646d696e00000000000000000000604482015290519081900360640190fd5b6040516001600160a01b038316907fa3b62bc36326052d97ea62d63c3d60308ed4c3ea8ac079dd8499f1e9c4f80c0f90600090a25b5050565b6000546001600160a01b03163314806104b457506001546000546001600160a01b039081169116145b610505576040805162461bcd60e51b815260206004820152601960248201527f4f776e6572206163636f756e7420697320726571756972656400000000000000604482015290519081900360640190fd5b6000546001600160a01b03828116911614156105525760405162461bcd60e51b81526004018080602001828103825260258152602001806108ac6025913960400191505060405180910390fd5b6001600160a01b0381166105ad576040805162461bcd60e51b815260206004820181905260248201527f4e6577204f776e65722063616e6e6f74206265207a65726f2061646472657373604482015290519081900360640190fd5b600080546001600160a01b0383811673ffffffffffffffffffffffffffffffffffffffff19831681178455604051919092169283917f8934ce4adea8d9ce0d714d2c22b86790e41b7731c84b926fbbdc1d40ff6533c99190a35050565b6001546001600160a01b031681565b6000546001600160a01b031681565b60055460ff1681565b60025481565b6001600160a01b0381166000908152600183016020526040812054600019018181128015906106665750835481125b949350505050565b60008082121580156106805750825482125b6106d1576040805162461bcd60e51b815260206004820152601860248201527f496e646578206f757473696465206f6620626f756e64732e0000000000000000604482015290519081900360640190fd5b50600101600090815260029190910160205260409020546001600160a01b031690565b60006001600160a01b03821661070c575060006102bc565b6001600160a01b03821660009081526001840160205260408120546000190190811280159061073b5750835481125b1561074a5760009150506102bc565b5050815460019081018084556001600160a01b03831660008181528386016020908152604080832085905593825260028701905291909120805473ffffffffffffffffffffffffffffffffffffffff1916909117905592915050565b6001600160a01b038116600090815260018084016020526040822054908112806107d05750835481135b156107df5760009150506102bc565b835481121561085357835460009081526002850160208181526040808420546001600160a01b031680855260018901835281852086905585855292909152808320805473ffffffffffffffffffffffffffffffffffffffff199081169093179055865483529091208054909116905561087f565b60008181526002850160205260409020805473ffffffffffffffffffffffffffffffffffffffff191690555b50506001600160a01b03166000908152600182810160205260408220919091558154600019019091559056fe4e6577204f776e65722063616e6e6f74206265207468652063757272656e74206f776e6572436f6e747261637420616c726561647920696e20726571756573746564206c6f636b207374617465a265627a7a7231582046160fc31272502ab8101ea185baddc02871a371a4716f5f528bf4256b4472a864736f6c634300050b0032`

// DeployAdminLockable deploys a new Ethereum contract, binding an instance of AdminLockable to it.
func DeployAdminLockable(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *AdminLockable, error) {
	parsed, err := abi.JSON(strings.NewReader(AdminLockableABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(AdminLockableBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &AdminLockable{AdminLockableCaller: AdminLockableCaller{contract: contract}, AdminLockableTransactor: AdminLockableTransactor{contract: contract}, AdminLockableFilterer: AdminLockableFilterer{contract: contract}}, nil
}

// AdminLockable is an auto generated Go binding around an Ethereum contract.
type AdminLockable struct {
	AdminLockableCaller     // Read-only binding to the contract
	AdminLockableTransactor // Write-only binding to the contract
	AdminLockableFilterer   // Log filterer for contract events
}

// AdminLockableCaller is an auto generated read-only Go binding around an Ethereum contract.
type AdminLockableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AdminLockableTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AdminLockableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AdminLockableFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AdminLockableFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AdminLockableSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AdminLockableSession struct {
	Contract     *AdminLockable    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AdminLockableCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AdminLockableCallerSession struct {
	Contract *AdminLockableCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// AdminLockableTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AdminLockableTransactorSession struct {
	Contract     *AdminLockableTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// AdminLockableRaw is an auto generated low-level Go binding around an Ethereum contract.
type AdminLockableRaw struct {
	Contract *AdminLockable // Generic contract binding to access the raw methods on
}

// AdminLockableCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AdminLockableCallerRaw struct {
	Contract *AdminLockableCaller // Generic read-only contract binding to access the raw methods on
}

// AdminLockableTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AdminLockableTransactorRaw struct {
	Contract *AdminLockableTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAdminLockable creates a new instance of AdminLockable, bound to a specific deployed contract.
func NewAdminLockable(address common.Address, backend bind.ContractBackend) (*AdminLockable, error) {
	contract, err := bindAdminLockable(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AdminLockable{AdminLockableCaller: AdminLockableCaller{contract: contract}, AdminLockableTransactor: AdminLockableTransactor{contract: contract}, AdminLockableFilterer: AdminLockableFilterer{contract: contract}}, nil
}

// NewAdminLockableCaller creates a new read-only instance of AdminLockable, bound to a specific deployed contract.
func NewAdminLockableCaller(address common.Address, caller bind.ContractCaller) (*AdminLockableCaller, error) {
	contract, err := bindAdminLockable(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AdminLockableCaller{contract: contract}, nil
}

// NewAdminLockableTransactor creates a new write-only instance of AdminLockable, bound to a specific deployed contract.
func NewAdminLockableTransactor(address common.Address, transactor bind.ContractTransactor) (*AdminLockableTransactor, error) {
	contract, err := bindAdminLockable(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AdminLockableTransactor{contract: contract}, nil
}

// NewAdminLockableFilterer creates a new log filterer instance of AdminLockable, bound to a specific deployed contract.
func NewAdminLockableFilterer(address common.Address, filterer bind.ContractFilterer) (*AdminLockableFilterer, error) {
	contract, err := bindAdminLockable(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AdminLockableFilterer{contract: contract}, nil
}

// bindAdminLockable binds a generic wrapper to an already deployed contract.
func bindAdminLockable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AdminLockableABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AdminLockable *AdminLockableRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _AdminLockable.Contract.AdminLockableCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AdminLockable *AdminLockableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AdminLockable.Contract.AdminLockableTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AdminLockable *AdminLockableRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AdminLockable.Contract.AdminLockableTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AdminLockable *AdminLockableCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _AdminLockable.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AdminLockable *AdminLockableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AdminLockable.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AdminLockable *AdminLockableTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AdminLockable.Contract.contract.Transact(opts, method, params...)
}

// ZEROADDRESS is a free data retrieval call binding the contract method 0x538ba4f9.
//
// Solidity: function ZERO_ADDRESS() constant returns(address)
func (_AdminLockable *AdminLockableCaller) ZEROADDRESS(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _AdminLockable.contract.Call(opts, out, "ZERO_ADDRESS")
	return *ret0, err
}

// ZEROADDRESS is a free data retrieval call binding the contract method 0x538ba4f9.
//
// Solidity: function ZERO_ADDRESS() constant returns(address)
func (_AdminLockable *AdminLockableSession) ZEROADDRESS() (common.Address, error) {
	return _AdminLockable.Contract.ZEROADDRESS(&_AdminLockable.CallOpts)
}

// ZEROADDRESS is a free data retrieval call binding the contract method 0x538ba4f9.
//
// Solidity: function ZERO_ADDRESS() constant returns(address)
func (_AdminLockable *AdminLockableCallerSession) ZEROADDRESS() (common.Address, error) {
	return _AdminLockable.Contract.ZEROADDRESS(&_AdminLockable.CallOpts)
}

// AdminAt is a free data retrieval call binding the contract method 0x2bdbc56f.
//
// Solidity: function adminAt(index int256) constant returns(address)
func (_AdminLockable *AdminLockableCaller) AdminAt(opts *bind.CallOpts, index *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _AdminLockable.contract.Call(opts, out, "adminAt", index)
	return *ret0, err
}

// AdminAt is a free data retrieval call binding the contract method 0x2bdbc56f.
//
// Solidity: function adminAt(index int256) constant returns(address)
func (_AdminLockable *AdminLockableSession) AdminAt(index *big.Int) (common.Address, error) {
	return _AdminLockable.Contract.AdminAt(&_AdminLockable.CallOpts, index)
}

// AdminAt is a free data retrieval call binding the contract method 0x2bdbc56f.
//
// Solidity: function adminAt(index int256) constant returns(address)
func (_AdminLockable *AdminLockableCallerSession) AdminAt(index *big.Int) (common.Address, error) {
	return _AdminLockable.Contract.AdminAt(&_AdminLockable.CallOpts, index)
}

// Admins is a free data retrieval call binding the contract method 0xa5de3619.
//
// Solidity: function admins() constant returns(count int256)
func (_AdminLockable *AdminLockableCaller) Admins(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AdminLockable.contract.Call(opts, out, "admins")
	return *ret0, err
}

// Admins is a free data retrieval call binding the contract method 0xa5de3619.
//
// Solidity: function admins() constant returns(count int256)
func (_AdminLockable *AdminLockableSession) Admins() (*big.Int, error) {
	return _AdminLockable.Contract.Admins(&_AdminLockable.CallOpts)
}

// Admins is a free data retrieval call binding the contract method 0xa5de3619.
//
// Solidity: function admins() constant returns(count int256)
func (_AdminLockable *AdminLockableCallerSession) Admins() (*big.Int, error) {
	return _AdminLockable.Contract.Admins(&_AdminLockable.CallOpts)
}

// IsAdmin is a free data retrieval call binding the contract method 0x24d7806c.
//
// Solidity: function isAdmin(addr address) constant returns(bool)
func (_AdminLockable *AdminLockableCaller) IsAdmin(opts *bind.CallOpts, addr common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _AdminLockable.contract.Call(opts, out, "isAdmin", addr)
	return *ret0, err
}

// IsAdmin is a free data retrieval call binding the contract method 0x24d7806c.
//
// Solidity: function isAdmin(addr address) constant returns(bool)
func (_AdminLockable *AdminLockableSession) IsAdmin(addr common.Address) (bool, error) {
	return _AdminLockable.Contract.IsAdmin(&_AdminLockable.CallOpts, addr)
}

// IsAdmin is a free data retrieval call binding the contract method 0x24d7806c.
//
// Solidity: function isAdmin(addr address) constant returns(bool)
func (_AdminLockable *AdminLockableCallerSession) IsAdmin(addr common.Address) (bool, error) {
	return _AdminLockable.Contract.IsAdmin(&_AdminLockable.CallOpts, addr)
}

// IsLocked is a free data retrieval call binding the contract method 0xa4e2d634.
//
// Solidity: function isLocked() constant returns(bool)
func (_AdminLockable *AdminLockableCaller) IsLocked(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _AdminLockable.contract.Call(opts, out, "isLocked")
	return *ret0, err
}

// IsLocked is a free data retrieval call binding the contract method 0xa4e2d634.
//
// Solidity: function isLocked() constant returns(bool)
func (_AdminLockable *AdminLockableSession) IsLocked() (bool, error) {
	return _AdminLockable.Contract.IsLocked(&_AdminLockable.CallOpts)
}

// IsLocked is a free data retrieval call binding the contract method 0xa4e2d634.
//
// Solidity: function isLocked() constant returns(bool)
func (_AdminLockable *AdminLockableCallerSession) IsLocked() (bool, error) {
	return _AdminLockable.Contract.IsLocked(&_AdminLockable.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_AdminLockable *AdminLockableCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _AdminLockable.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_AdminLockable *AdminLockableSession) Owner() (common.Address, error) {
	return _AdminLockable.Contract.Owner(&_AdminLockable.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_AdminLockable *AdminLockableCallerSession) Owner() (common.Address, error) {
	return _AdminLockable.Contract.Owner(&_AdminLockable.CallOpts)
}

// SetAdmin is a paid mutator transaction binding the contract method 0x4b0bddd2.
//
// Solidity: function setAdmin(addr address, add bool) returns()
func (_AdminLockable *AdminLockableTransactor) SetAdmin(opts *bind.TransactOpts, addr common.Address, add bool) (*types.Transaction, error) {
	return _AdminLockable.contract.Transact(opts, "setAdmin", addr, add)
}

// SetAdmin is a paid mutator transaction binding the contract method 0x4b0bddd2.
//
// Solidity: function setAdmin(addr address, add bool) returns()
func (_AdminLockable *AdminLockableSession) SetAdmin(addr common.Address, add bool) (*types.Transaction, error) {
	return _AdminLockable.Contract.SetAdmin(&_AdminLockable.TransactOpts, addr, add)
}

// SetAdmin is a paid mutator transaction binding the contract method 0x4b0bddd2.
//
// Solidity: function setAdmin(addr address, add bool) returns()
func (_AdminLockable *AdminLockableTransactorSession) SetAdmin(addr common.Address, add bool) (*types.Transaction, error) {
	return _AdminLockable.Contract.SetAdmin(&_AdminLockable.TransactOpts, addr, add)
}

// SetLocked is a paid mutator transaction binding the contract method 0x211e28b6.
//
// Solidity: function setLocked(locked bool) returns()
func (_AdminLockable *AdminLockableTransactor) SetLocked(opts *bind.TransactOpts, locked bool) (*types.Transaction, error) {
	return _AdminLockable.contract.Transact(opts, "setLocked", locked)
}

// SetLocked is a paid mutator transaction binding the contract method 0x211e28b6.
//
// Solidity: function setLocked(locked bool) returns()
func (_AdminLockable *AdminLockableSession) SetLocked(locked bool) (*types.Transaction, error) {
	return _AdminLockable.Contract.SetLocked(&_AdminLockable.TransactOpts, locked)
}

// SetLocked is a paid mutator transaction binding the contract method 0x211e28b6.
//
// Solidity: function setLocked(locked bool) returns()
func (_AdminLockable *AdminLockableTransactorSession) SetLocked(locked bool) (*types.Transaction, error) {
	return _AdminLockable.Contract.SetLocked(&_AdminLockable.TransactOpts, locked)
}

// TransferOwner is a paid mutator transaction binding the contract method 0x4fb2e45d.
//
// Solidity: function transferOwner(newOwner address) returns()
func (_AdminLockable *AdminLockableTransactor) TransferOwner(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _AdminLockable.contract.Transact(opts, "transferOwner", newOwner)
}

// TransferOwner is a paid mutator transaction binding the contract method 0x4fb2e45d.
//
// Solidity: function transferOwner(newOwner address) returns()
func (_AdminLockable *AdminLockableSession) TransferOwner(newOwner common.Address) (*types.Transaction, error) {
	return _AdminLockable.Contract.TransferOwner(&_AdminLockable.TransactOpts, newOwner)
}

// TransferOwner is a paid mutator transaction binding the contract method 0x4fb2e45d.
//
// Solidity: function transferOwner(newOwner address) returns()
func (_AdminLockable *AdminLockableTransactorSession) TransferOwner(newOwner common.Address) (*types.Transaction, error) {
	return _AdminLockable.Contract.TransferOwner(&_AdminLockable.TransactOpts, newOwner)
}

// AdminLockableAdminAddedIterator is returned from FilterAdminAdded and is used to iterate over the raw logs and unpacked data for AdminAdded events raised by the AdminLockable contract.
type AdminLockableAdminAddedIterator struct {
	Event *AdminLockableAdminAdded // Event containing the contract specifics and raw log

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
func (it *AdminLockableAdminAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AdminLockableAdminAdded)
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
		it.Event = new(AdminLockableAdminAdded)
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
func (it *AdminLockableAdminAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AdminLockableAdminAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AdminLockableAdminAdded represents a AdminAdded event raised by the AdminLockable contract.
type AdminLockableAdminAdded struct {
	Admin common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterAdminAdded is a free log retrieval operation binding the contract event 0x44d6d25963f097ad14f29f06854a01f575648a1ef82f30e562ccd3889717e339.
//
// Solidity: e AdminAdded(admin indexed address)
func (_AdminLockable *AdminLockableFilterer) FilterAdminAdded(opts *bind.FilterOpts, admin []common.Address) (*AdminLockableAdminAddedIterator, error) {

	var adminRule []interface{}
	for _, adminItem := range admin {
		adminRule = append(adminRule, adminItem)
	}

	logs, sub, err := _AdminLockable.contract.FilterLogs(opts, "AdminAdded", adminRule)
	if err != nil {
		return nil, err
	}
	return &AdminLockableAdminAddedIterator{contract: _AdminLockable.contract, event: "AdminAdded", logs: logs, sub: sub}, nil
}

// WatchAdminAdded is a free log subscription operation binding the contract event 0x44d6d25963f097ad14f29f06854a01f575648a1ef82f30e562ccd3889717e339.
//
// Solidity: e AdminAdded(admin indexed address)
func (_AdminLockable *AdminLockableFilterer) WatchAdminAdded(opts *bind.WatchOpts, sink chan<- *AdminLockableAdminAdded, admin []common.Address) (event.Subscription, error) {

	var adminRule []interface{}
	for _, adminItem := range admin {
		adminRule = append(adminRule, adminItem)
	}

	logs, sub, err := _AdminLockable.contract.WatchLogs(opts, "AdminAdded", adminRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AdminLockableAdminAdded)
				if err := _AdminLockable.contract.UnpackLog(event, "AdminAdded", log); err != nil {
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

// AdminLockableAdminRemovedIterator is returned from FilterAdminRemoved and is used to iterate over the raw logs and unpacked data for AdminRemoved events raised by the AdminLockable contract.
type AdminLockableAdminRemovedIterator struct {
	Event *AdminLockableAdminRemoved // Event containing the contract specifics and raw log

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
func (it *AdminLockableAdminRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AdminLockableAdminRemoved)
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
		it.Event = new(AdminLockableAdminRemoved)
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
func (it *AdminLockableAdminRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AdminLockableAdminRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AdminLockableAdminRemoved represents a AdminRemoved event raised by the AdminLockable contract.
type AdminLockableAdminRemoved struct {
	Admin common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterAdminRemoved is a free log retrieval operation binding the contract event 0xa3b62bc36326052d97ea62d63c3d60308ed4c3ea8ac079dd8499f1e9c4f80c0f.
//
// Solidity: e AdminRemoved(admin indexed address)
func (_AdminLockable *AdminLockableFilterer) FilterAdminRemoved(opts *bind.FilterOpts, admin []common.Address) (*AdminLockableAdminRemovedIterator, error) {

	var adminRule []interface{}
	for _, adminItem := range admin {
		adminRule = append(adminRule, adminItem)
	}

	logs, sub, err := _AdminLockable.contract.FilterLogs(opts, "AdminRemoved", adminRule)
	if err != nil {
		return nil, err
	}
	return &AdminLockableAdminRemovedIterator{contract: _AdminLockable.contract, event: "AdminRemoved", logs: logs, sub: sub}, nil
}

// WatchAdminRemoved is a free log subscription operation binding the contract event 0xa3b62bc36326052d97ea62d63c3d60308ed4c3ea8ac079dd8499f1e9c4f80c0f.
//
// Solidity: e AdminRemoved(admin indexed address)
func (_AdminLockable *AdminLockableFilterer) WatchAdminRemoved(opts *bind.WatchOpts, sink chan<- *AdminLockableAdminRemoved, admin []common.Address) (event.Subscription, error) {

	var adminRule []interface{}
	for _, adminItem := range admin {
		adminRule = append(adminRule, adminItem)
	}

	logs, sub, err := _AdminLockable.contract.WatchLogs(opts, "AdminRemoved", adminRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AdminLockableAdminRemoved)
				if err := _AdminLockable.contract.UnpackLog(event, "AdminRemoved", log); err != nil {
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

// AdminLockableOwnerTransferredIterator is returned from FilterOwnerTransferred and is used to iterate over the raw logs and unpacked data for OwnerTransferred events raised by the AdminLockable contract.
type AdminLockableOwnerTransferredIterator struct {
	Event *AdminLockableOwnerTransferred // Event containing the contract specifics and raw log

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
func (it *AdminLockableOwnerTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AdminLockableOwnerTransferred)
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
		it.Event = new(AdminLockableOwnerTransferred)
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
func (it *AdminLockableOwnerTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AdminLockableOwnerTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AdminLockableOwnerTransferred represents a OwnerTransferred event raised by the AdminLockable contract.
type AdminLockableOwnerTransferred struct {
	OldOwner common.Address
	NewOwner common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOwnerTransferred is a free log retrieval operation binding the contract event 0x8934ce4adea8d9ce0d714d2c22b86790e41b7731c84b926fbbdc1d40ff6533c9.
//
// Solidity: e OwnerTransferred(oldOwner indexed address, newOwner indexed address)
func (_AdminLockable *AdminLockableFilterer) FilterOwnerTransferred(opts *bind.FilterOpts, oldOwner []common.Address, newOwner []common.Address) (*AdminLockableOwnerTransferredIterator, error) {

	var oldOwnerRule []interface{}
	for _, oldOwnerItem := range oldOwner {
		oldOwnerRule = append(oldOwnerRule, oldOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _AdminLockable.contract.FilterLogs(opts, "OwnerTransferred", oldOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &AdminLockableOwnerTransferredIterator{contract: _AdminLockable.contract, event: "OwnerTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnerTransferred is a free log subscription operation binding the contract event 0x8934ce4adea8d9ce0d714d2c22b86790e41b7731c84b926fbbdc1d40ff6533c9.
//
// Solidity: e OwnerTransferred(oldOwner indexed address, newOwner indexed address)
func (_AdminLockable *AdminLockableFilterer) WatchOwnerTransferred(opts *bind.WatchOpts, sink chan<- *AdminLockableOwnerTransferred, oldOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var oldOwnerRule []interface{}
	for _, oldOwnerItem := range oldOwner {
		oldOwnerRule = append(oldOwnerRule, oldOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _AdminLockable.contract.WatchLogs(opts, "OwnerTransferred", oldOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AdminLockableOwnerTransferred)
				if err := _AdminLockable.contract.UnpackLog(event, "OwnerTransferred", log); err != nil {
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
