// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package lifecycle

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

// AdminLockableDestroyableABI is the input ABI used to generate the binding from.
const AdminLockableDestroyableABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"admin\",\"type\":\"address\"}],\"name\":\"removeAdmin\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"locked\",\"type\":\"bool\"}],\"name\":\"setLocked\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"isAdmin\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"kill\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwner\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"admin\",\"type\":\"address\"}],\"name\":\"addAdmin\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isLocked\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"admins\",\"outputs\":[{\"name\":\"count\",\"type\":\"int256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"admin\",\"type\":\"address\"}],\"name\":\"AdminAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"admin\",\"type\":\"address\"}],\"name\":\"AdminRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"oldOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnerTransferred\",\"type\":\"event\"}]"

// AdminLockableDestroyableBin is the compiled bytecode used for deploying new contracts.
const AdminLockableDestroyableBin = `608060405260008054600160a060020a031916331790556004805460ff19169055610b798061002f6000396000f3fe608060405234801561001057600080fd5b50600436106100c0576000357c0100000000000000000000000000000000000000000000000000000000900480634fb2e45d116100935780638da5cb5b116100785780638da5cb5b146101ce578063a4e2d634146101ff578063a5de361914610207576100c0565b80634fb2e45d14610168578063704802751461019b576100c0565b80631785f53c146100c5578063211e28b6146100fa57806324d7806c1461011957806341c0e1b514610160575b600080fd5b6100f8600480360360208110156100db57600080fd5b503573ffffffffffffffffffffffffffffffffffffffff16610221565b005b6100f86004803603602081101561011057600080fd5b50351515610369565b61014c6004803603602081101561012f57600080fd5b503573ffffffffffffffffffffffffffffffffffffffff1661047d565b604080519115158252519081900360200190f35b6100f86104b9565b6100f86004803603602081101561017e57600080fd5b503573ffffffffffffffffffffffffffffffffffffffff1661055a565b6100f8600480360360208110156101b157600080fd5b503573ffffffffffffffffffffffffffffffffffffffff16610742565b6101d661088a565b6040805173ffffffffffffffffffffffffffffffffffffffff9092168252519081900360200190f35b61014c6108a6565b61020f6108af565b60408051918252519081900360200190f35b60005473ffffffffffffffffffffffffffffffffffffffff1633146102a757604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601960248201527f4f776e6572206163636f756e7420697320726571756972656400000000000000604482015290519081900360640190fd5b6102b860018263ffffffff6108b516565b151561032557604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f556e61626c6520746f2072656d6f76652061646d696e00000000000000000000604482015290519081900360640190fd5b60405173ffffffffffffffffffffffffffffffffffffffff8216907fa3b62bc36326052d97ea62d63c3d60308ed4c3ea8ac079dd8499f1e9c4f80c0f90600090a250565b60005473ffffffffffffffffffffffffffffffffffffffff1633148061039b575061039b60013363ffffffff6109e116565b151561040857604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601960248201527f41646d696e206163636f756e7420697320726571756972656400000000000000604482015290519081900360640190fd5b60045460ff161515811515141561046a576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401808060200182810382526028815260200180610b266028913960400191505060405180910390fd5b6004805460ff1916911515919091179055565b6000805473ffffffffffffffffffffffffffffffffffffffff838116911614806104b357506104b360018363ffffffff6109e116565b92915050565b60005473ffffffffffffffffffffffffffffffffffffffff16331461053f57604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601960248201527f4f776e6572206163636f756e7420697320726571756972656400000000000000604482015290519081900360640190fd5b60005473ffffffffffffffffffffffffffffffffffffffff16ff5b60005473ffffffffffffffffffffffffffffffffffffffff1633146105e057604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601960248201527f4f776e6572206163636f756e7420697320726571756972656400000000000000604482015290519081900360640190fd5b60005473ffffffffffffffffffffffffffffffffffffffff82811691161415610654576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401808060200182810382526025815260200180610b016025913960400191505060405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff811615156106d857604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820181905260248201527f4e6577204f776e65722063616e6e6f74206265207a65726f2061646472657373604482015290519081900360640190fd5b6000805473ffffffffffffffffffffffffffffffffffffffff83811673ffffffffffffffffffffffffffffffffffffffff19831681178455604051919092169283917f8934ce4adea8d9ce0d714d2c22b86790e41b7731c84b926fbbdc1d40ff6533c99190a35050565b60005473ffffffffffffffffffffffffffffffffffffffff1633146107c857604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601960248201527f4f776e6572206163636f756e7420697320726571756972656400000000000000604482015290519081900360640190fd5b6107d960018263ffffffff610a2516565b151561084657604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601360248201527f556e61626c6520746f206164642061646d696e00000000000000000000000000604482015290519081900360640190fd5b60405173ffffffffffffffffffffffffffffffffffffffff8216907f44d6d25963f097ad14f29f06854a01f575648a1ef82f30e562ccd3889717e33990600090a250565b60005473ffffffffffffffffffffffffffffffffffffffff1681565b60045460ff1681565b60015481565b73ffffffffffffffffffffffffffffffffffffffff8116600090815260018084016020526040822054908112806108ec5750835481135b156108fb5760009150506104b3565b835481121561097c578354600090815260028501602081815260408084205473ffffffffffffffffffffffffffffffffffffffff1680855260018901835281852086905585855292909152808320805473ffffffffffffffffffffffffffffffffffffffff19908116909317905586548352909120805490911690556109a8565b60008181526002850160205260409020805473ffffffffffffffffffffffffffffffffffffffff191690555b505073ffffffffffffffffffffffffffffffffffffffff1660009081526001828101602052604082209190915581546000190190915590565b73ffffffffffffffffffffffffffffffffffffffff8116600090815260018301602052604081205460001901818112801590610a1d5750835481125b949350505050565b600073ffffffffffffffffffffffffffffffffffffffff82161515610a4c575060006104b3565b73ffffffffffffffffffffffffffffffffffffffff8216600090815260018401602052604081205460001901908112801590610a885750835481125b15610a975760009150506104b3565b50508154600190810180845573ffffffffffffffffffffffffffffffffffffffff831660008181528386016020908152604080832085905593825260028701905291909120805473ffffffffffffffffffffffffffffffffffffffff191690911790559291505056fe4e6577204f776e65722063616e6e6f74206265207468652063757272656e74206f776e6572436f6e747261637420616c726561647920696e20726571756573746564206c6f636b207374617465a165627a7a72305820fdc149bbe106d8c1131ff9601f955abe16d25f44b27cd2a72e282c1b28db58aa0029`

// DeployAdminLockableDestroyable deploys a new Ethereum contract, binding an instance of AdminLockableDestroyable to it.
func DeployAdminLockableDestroyable(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *AdminLockableDestroyable, error) {
	parsed, err := abi.JSON(strings.NewReader(AdminLockableDestroyableABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(AdminLockableDestroyableBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &AdminLockableDestroyable{AdminLockableDestroyableCaller: AdminLockableDestroyableCaller{contract: contract}, AdminLockableDestroyableTransactor: AdminLockableDestroyableTransactor{contract: contract}, AdminLockableDestroyableFilterer: AdminLockableDestroyableFilterer{contract: contract}}, nil
}

// AdminLockableDestroyable is an auto generated Go binding around an Ethereum contract.
type AdminLockableDestroyable struct {
	AdminLockableDestroyableCaller     // Read-only binding to the contract
	AdminLockableDestroyableTransactor // Write-only binding to the contract
	AdminLockableDestroyableFilterer   // Log filterer for contract events
}

// AdminLockableDestroyableCaller is an auto generated read-only Go binding around an Ethereum contract.
type AdminLockableDestroyableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AdminLockableDestroyableTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AdminLockableDestroyableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AdminLockableDestroyableFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AdminLockableDestroyableFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AdminLockableDestroyableSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AdminLockableDestroyableSession struct {
	Contract     *AdminLockableDestroyable // Generic contract binding to set the session for
	CallOpts     bind.CallOpts             // Call options to use throughout this session
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// AdminLockableDestroyableCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AdminLockableDestroyableCallerSession struct {
	Contract *AdminLockableDestroyableCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                   // Call options to use throughout this session
}

// AdminLockableDestroyableTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AdminLockableDestroyableTransactorSession struct {
	Contract     *AdminLockableDestroyableTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                   // Transaction auth options to use throughout this session
}

// AdminLockableDestroyableRaw is an auto generated low-level Go binding around an Ethereum contract.
type AdminLockableDestroyableRaw struct {
	Contract *AdminLockableDestroyable // Generic contract binding to access the raw methods on
}

// AdminLockableDestroyableCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AdminLockableDestroyableCallerRaw struct {
	Contract *AdminLockableDestroyableCaller // Generic read-only contract binding to access the raw methods on
}

// AdminLockableDestroyableTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AdminLockableDestroyableTransactorRaw struct {
	Contract *AdminLockableDestroyableTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAdminLockableDestroyable creates a new instance of AdminLockableDestroyable, bound to a specific deployed contract.
func NewAdminLockableDestroyable(address common.Address, backend bind.ContractBackend) (*AdminLockableDestroyable, error) {
	contract, err := bindAdminLockableDestroyable(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AdminLockableDestroyable{AdminLockableDestroyableCaller: AdminLockableDestroyableCaller{contract: contract}, AdminLockableDestroyableTransactor: AdminLockableDestroyableTransactor{contract: contract}, AdminLockableDestroyableFilterer: AdminLockableDestroyableFilterer{contract: contract}}, nil
}

// NewAdminLockableDestroyableCaller creates a new read-only instance of AdminLockableDestroyable, bound to a specific deployed contract.
func NewAdminLockableDestroyableCaller(address common.Address, caller bind.ContractCaller) (*AdminLockableDestroyableCaller, error) {
	contract, err := bindAdminLockableDestroyable(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AdminLockableDestroyableCaller{contract: contract}, nil
}

// NewAdminLockableDestroyableTransactor creates a new write-only instance of AdminLockableDestroyable, bound to a specific deployed contract.
func NewAdminLockableDestroyableTransactor(address common.Address, transactor bind.ContractTransactor) (*AdminLockableDestroyableTransactor, error) {
	contract, err := bindAdminLockableDestroyable(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AdminLockableDestroyableTransactor{contract: contract}, nil
}

// NewAdminLockableDestroyableFilterer creates a new log filterer instance of AdminLockableDestroyable, bound to a specific deployed contract.
func NewAdminLockableDestroyableFilterer(address common.Address, filterer bind.ContractFilterer) (*AdminLockableDestroyableFilterer, error) {
	contract, err := bindAdminLockableDestroyable(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AdminLockableDestroyableFilterer{contract: contract}, nil
}

// bindAdminLockableDestroyable binds a generic wrapper to an already deployed contract.
func bindAdminLockableDestroyable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AdminLockableDestroyableABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AdminLockableDestroyable *AdminLockableDestroyableRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _AdminLockableDestroyable.Contract.AdminLockableDestroyableCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AdminLockableDestroyable *AdminLockableDestroyableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AdminLockableDestroyable.Contract.AdminLockableDestroyableTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AdminLockableDestroyable *AdminLockableDestroyableRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AdminLockableDestroyable.Contract.AdminLockableDestroyableTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AdminLockableDestroyable *AdminLockableDestroyableCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _AdminLockableDestroyable.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AdminLockableDestroyable *AdminLockableDestroyableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AdminLockableDestroyable.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AdminLockableDestroyable *AdminLockableDestroyableTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AdminLockableDestroyable.Contract.contract.Transact(opts, method, params...)
}

// Admins is a free data retrieval call binding the contract method 0xa5de3619.
//
// Solidity: function admins() constant returns(count int256)
func (_AdminLockableDestroyable *AdminLockableDestroyableCaller) Admins(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AdminLockableDestroyable.contract.Call(opts, out, "admins")
	return *ret0, err
}

// Admins is a free data retrieval call binding the contract method 0xa5de3619.
//
// Solidity: function admins() constant returns(count int256)
func (_AdminLockableDestroyable *AdminLockableDestroyableSession) Admins() (*big.Int, error) {
	return _AdminLockableDestroyable.Contract.Admins(&_AdminLockableDestroyable.CallOpts)
}

// Admins is a free data retrieval call binding the contract method 0xa5de3619.
//
// Solidity: function admins() constant returns(count int256)
func (_AdminLockableDestroyable *AdminLockableDestroyableCallerSession) Admins() (*big.Int, error) {
	return _AdminLockableDestroyable.Contract.Admins(&_AdminLockableDestroyable.CallOpts)
}

// IsAdmin is a free data retrieval call binding the contract method 0x24d7806c.
//
// Solidity: function isAdmin(addr address) constant returns(bool)
func (_AdminLockableDestroyable *AdminLockableDestroyableCaller) IsAdmin(opts *bind.CallOpts, addr common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _AdminLockableDestroyable.contract.Call(opts, out, "isAdmin", addr)
	return *ret0, err
}

// IsAdmin is a free data retrieval call binding the contract method 0x24d7806c.
//
// Solidity: function isAdmin(addr address) constant returns(bool)
func (_AdminLockableDestroyable *AdminLockableDestroyableSession) IsAdmin(addr common.Address) (bool, error) {
	return _AdminLockableDestroyable.Contract.IsAdmin(&_AdminLockableDestroyable.CallOpts, addr)
}

// IsAdmin is a free data retrieval call binding the contract method 0x24d7806c.
//
// Solidity: function isAdmin(addr address) constant returns(bool)
func (_AdminLockableDestroyable *AdminLockableDestroyableCallerSession) IsAdmin(addr common.Address) (bool, error) {
	return _AdminLockableDestroyable.Contract.IsAdmin(&_AdminLockableDestroyable.CallOpts, addr)
}

// IsLocked is a free data retrieval call binding the contract method 0xa4e2d634.
//
// Solidity: function isLocked() constant returns(bool)
func (_AdminLockableDestroyable *AdminLockableDestroyableCaller) IsLocked(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _AdminLockableDestroyable.contract.Call(opts, out, "isLocked")
	return *ret0, err
}

// IsLocked is a free data retrieval call binding the contract method 0xa4e2d634.
//
// Solidity: function isLocked() constant returns(bool)
func (_AdminLockableDestroyable *AdminLockableDestroyableSession) IsLocked() (bool, error) {
	return _AdminLockableDestroyable.Contract.IsLocked(&_AdminLockableDestroyable.CallOpts)
}

// IsLocked is a free data retrieval call binding the contract method 0xa4e2d634.
//
// Solidity: function isLocked() constant returns(bool)
func (_AdminLockableDestroyable *AdminLockableDestroyableCallerSession) IsLocked() (bool, error) {
	return _AdminLockableDestroyable.Contract.IsLocked(&_AdminLockableDestroyable.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_AdminLockableDestroyable *AdminLockableDestroyableCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _AdminLockableDestroyable.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_AdminLockableDestroyable *AdminLockableDestroyableSession) Owner() (common.Address, error) {
	return _AdminLockableDestroyable.Contract.Owner(&_AdminLockableDestroyable.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_AdminLockableDestroyable *AdminLockableDestroyableCallerSession) Owner() (common.Address, error) {
	return _AdminLockableDestroyable.Contract.Owner(&_AdminLockableDestroyable.CallOpts)
}

// AddAdmin is a paid mutator transaction binding the contract method 0x70480275.
//
// Solidity: function addAdmin(admin address) returns()
func (_AdminLockableDestroyable *AdminLockableDestroyableTransactor) AddAdmin(opts *bind.TransactOpts, admin common.Address) (*types.Transaction, error) {
	return _AdminLockableDestroyable.contract.Transact(opts, "addAdmin", admin)
}

// AddAdmin is a paid mutator transaction binding the contract method 0x70480275.
//
// Solidity: function addAdmin(admin address) returns()
func (_AdminLockableDestroyable *AdminLockableDestroyableSession) AddAdmin(admin common.Address) (*types.Transaction, error) {
	return _AdminLockableDestroyable.Contract.AddAdmin(&_AdminLockableDestroyable.TransactOpts, admin)
}

// AddAdmin is a paid mutator transaction binding the contract method 0x70480275.
//
// Solidity: function addAdmin(admin address) returns()
func (_AdminLockableDestroyable *AdminLockableDestroyableTransactorSession) AddAdmin(admin common.Address) (*types.Transaction, error) {
	return _AdminLockableDestroyable.Contract.AddAdmin(&_AdminLockableDestroyable.TransactOpts, admin)
}

// Kill is a paid mutator transaction binding the contract method 0x41c0e1b5.
//
// Solidity: function kill() returns()
func (_AdminLockableDestroyable *AdminLockableDestroyableTransactor) Kill(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AdminLockableDestroyable.contract.Transact(opts, "kill")
}

// Kill is a paid mutator transaction binding the contract method 0x41c0e1b5.
//
// Solidity: function kill() returns()
func (_AdminLockableDestroyable *AdminLockableDestroyableSession) Kill() (*types.Transaction, error) {
	return _AdminLockableDestroyable.Contract.Kill(&_AdminLockableDestroyable.TransactOpts)
}

// Kill is a paid mutator transaction binding the contract method 0x41c0e1b5.
//
// Solidity: function kill() returns()
func (_AdminLockableDestroyable *AdminLockableDestroyableTransactorSession) Kill() (*types.Transaction, error) {
	return _AdminLockableDestroyable.Contract.Kill(&_AdminLockableDestroyable.TransactOpts)
}

// RemoveAdmin is a paid mutator transaction binding the contract method 0x1785f53c.
//
// Solidity: function removeAdmin(admin address) returns()
func (_AdminLockableDestroyable *AdminLockableDestroyableTransactor) RemoveAdmin(opts *bind.TransactOpts, admin common.Address) (*types.Transaction, error) {
	return _AdminLockableDestroyable.contract.Transact(opts, "removeAdmin", admin)
}

// RemoveAdmin is a paid mutator transaction binding the contract method 0x1785f53c.
//
// Solidity: function removeAdmin(admin address) returns()
func (_AdminLockableDestroyable *AdminLockableDestroyableSession) RemoveAdmin(admin common.Address) (*types.Transaction, error) {
	return _AdminLockableDestroyable.Contract.RemoveAdmin(&_AdminLockableDestroyable.TransactOpts, admin)
}

// RemoveAdmin is a paid mutator transaction binding the contract method 0x1785f53c.
//
// Solidity: function removeAdmin(admin address) returns()
func (_AdminLockableDestroyable *AdminLockableDestroyableTransactorSession) RemoveAdmin(admin common.Address) (*types.Transaction, error) {
	return _AdminLockableDestroyable.Contract.RemoveAdmin(&_AdminLockableDestroyable.TransactOpts, admin)
}

// SetLocked is a paid mutator transaction binding the contract method 0x211e28b6.
//
// Solidity: function setLocked(locked bool) returns()
func (_AdminLockableDestroyable *AdminLockableDestroyableTransactor) SetLocked(opts *bind.TransactOpts, locked bool) (*types.Transaction, error) {
	return _AdminLockableDestroyable.contract.Transact(opts, "setLocked", locked)
}

// SetLocked is a paid mutator transaction binding the contract method 0x211e28b6.
//
// Solidity: function setLocked(locked bool) returns()
func (_AdminLockableDestroyable *AdminLockableDestroyableSession) SetLocked(locked bool) (*types.Transaction, error) {
	return _AdminLockableDestroyable.Contract.SetLocked(&_AdminLockableDestroyable.TransactOpts, locked)
}

// SetLocked is a paid mutator transaction binding the contract method 0x211e28b6.
//
// Solidity: function setLocked(locked bool) returns()
func (_AdminLockableDestroyable *AdminLockableDestroyableTransactorSession) SetLocked(locked bool) (*types.Transaction, error) {
	return _AdminLockableDestroyable.Contract.SetLocked(&_AdminLockableDestroyable.TransactOpts, locked)
}

// TransferOwner is a paid mutator transaction binding the contract method 0x4fb2e45d.
//
// Solidity: function transferOwner(newOwner address) returns()
func (_AdminLockableDestroyable *AdminLockableDestroyableTransactor) TransferOwner(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _AdminLockableDestroyable.contract.Transact(opts, "transferOwner", newOwner)
}

// TransferOwner is a paid mutator transaction binding the contract method 0x4fb2e45d.
//
// Solidity: function transferOwner(newOwner address) returns()
func (_AdminLockableDestroyable *AdminLockableDestroyableSession) TransferOwner(newOwner common.Address) (*types.Transaction, error) {
	return _AdminLockableDestroyable.Contract.TransferOwner(&_AdminLockableDestroyable.TransactOpts, newOwner)
}

// TransferOwner is a paid mutator transaction binding the contract method 0x4fb2e45d.
//
// Solidity: function transferOwner(newOwner address) returns()
func (_AdminLockableDestroyable *AdminLockableDestroyableTransactorSession) TransferOwner(newOwner common.Address) (*types.Transaction, error) {
	return _AdminLockableDestroyable.Contract.TransferOwner(&_AdminLockableDestroyable.TransactOpts, newOwner)
}

// AdminLockableDestroyableAdminAddedIterator is returned from FilterAdminAdded and is used to iterate over the raw logs and unpacked data for AdminAdded events raised by the AdminLockableDestroyable contract.
type AdminLockableDestroyableAdminAddedIterator struct {
	Event *AdminLockableDestroyableAdminAdded // Event containing the contract specifics and raw log

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
func (it *AdminLockableDestroyableAdminAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AdminLockableDestroyableAdminAdded)
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
		it.Event = new(AdminLockableDestroyableAdminAdded)
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
func (it *AdminLockableDestroyableAdminAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AdminLockableDestroyableAdminAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AdminLockableDestroyableAdminAdded represents a AdminAdded event raised by the AdminLockableDestroyable contract.
type AdminLockableDestroyableAdminAdded struct {
	Admin common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterAdminAdded is a free log retrieval operation binding the contract event 0x44d6d25963f097ad14f29f06854a01f575648a1ef82f30e562ccd3889717e339.
//
// Solidity: e AdminAdded(admin indexed address)
func (_AdminLockableDestroyable *AdminLockableDestroyableFilterer) FilterAdminAdded(opts *bind.FilterOpts, admin []common.Address) (*AdminLockableDestroyableAdminAddedIterator, error) {

	var adminRule []interface{}
	for _, adminItem := range admin {
		adminRule = append(adminRule, adminItem)
	}

	logs, sub, err := _AdminLockableDestroyable.contract.FilterLogs(opts, "AdminAdded", adminRule)
	if err != nil {
		return nil, err
	}
	return &AdminLockableDestroyableAdminAddedIterator{contract: _AdminLockableDestroyable.contract, event: "AdminAdded", logs: logs, sub: sub}, nil
}

// WatchAdminAdded is a free log subscription operation binding the contract event 0x44d6d25963f097ad14f29f06854a01f575648a1ef82f30e562ccd3889717e339.
//
// Solidity: e AdminAdded(admin indexed address)
func (_AdminLockableDestroyable *AdminLockableDestroyableFilterer) WatchAdminAdded(opts *bind.WatchOpts, sink chan<- *AdminLockableDestroyableAdminAdded, admin []common.Address) (event.Subscription, error) {

	var adminRule []interface{}
	for _, adminItem := range admin {
		adminRule = append(adminRule, adminItem)
	}

	logs, sub, err := _AdminLockableDestroyable.contract.WatchLogs(opts, "AdminAdded", adminRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AdminLockableDestroyableAdminAdded)
				if err := _AdminLockableDestroyable.contract.UnpackLog(event, "AdminAdded", log); err != nil {
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

// AdminLockableDestroyableAdminRemovedIterator is returned from FilterAdminRemoved and is used to iterate over the raw logs and unpacked data for AdminRemoved events raised by the AdminLockableDestroyable contract.
type AdminLockableDestroyableAdminRemovedIterator struct {
	Event *AdminLockableDestroyableAdminRemoved // Event containing the contract specifics and raw log

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
func (it *AdminLockableDestroyableAdminRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AdminLockableDestroyableAdminRemoved)
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
		it.Event = new(AdminLockableDestroyableAdminRemoved)
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
func (it *AdminLockableDestroyableAdminRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AdminLockableDestroyableAdminRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AdminLockableDestroyableAdminRemoved represents a AdminRemoved event raised by the AdminLockableDestroyable contract.
type AdminLockableDestroyableAdminRemoved struct {
	Admin common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterAdminRemoved is a free log retrieval operation binding the contract event 0xa3b62bc36326052d97ea62d63c3d60308ed4c3ea8ac079dd8499f1e9c4f80c0f.
//
// Solidity: e AdminRemoved(admin indexed address)
func (_AdminLockableDestroyable *AdminLockableDestroyableFilterer) FilterAdminRemoved(opts *bind.FilterOpts, admin []common.Address) (*AdminLockableDestroyableAdminRemovedIterator, error) {

	var adminRule []interface{}
	for _, adminItem := range admin {
		adminRule = append(adminRule, adminItem)
	}

	logs, sub, err := _AdminLockableDestroyable.contract.FilterLogs(opts, "AdminRemoved", adminRule)
	if err != nil {
		return nil, err
	}
	return &AdminLockableDestroyableAdminRemovedIterator{contract: _AdminLockableDestroyable.contract, event: "AdminRemoved", logs: logs, sub: sub}, nil
}

// WatchAdminRemoved is a free log subscription operation binding the contract event 0xa3b62bc36326052d97ea62d63c3d60308ed4c3ea8ac079dd8499f1e9c4f80c0f.
//
// Solidity: e AdminRemoved(admin indexed address)
func (_AdminLockableDestroyable *AdminLockableDestroyableFilterer) WatchAdminRemoved(opts *bind.WatchOpts, sink chan<- *AdminLockableDestroyableAdminRemoved, admin []common.Address) (event.Subscription, error) {

	var adminRule []interface{}
	for _, adminItem := range admin {
		adminRule = append(adminRule, adminItem)
	}

	logs, sub, err := _AdminLockableDestroyable.contract.WatchLogs(opts, "AdminRemoved", adminRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AdminLockableDestroyableAdminRemoved)
				if err := _AdminLockableDestroyable.contract.UnpackLog(event, "AdminRemoved", log); err != nil {
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

// AdminLockableDestroyableOwnerTransferredIterator is returned from FilterOwnerTransferred and is used to iterate over the raw logs and unpacked data for OwnerTransferred events raised by the AdminLockableDestroyable contract.
type AdminLockableDestroyableOwnerTransferredIterator struct {
	Event *AdminLockableDestroyableOwnerTransferred // Event containing the contract specifics and raw log

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
func (it *AdminLockableDestroyableOwnerTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AdminLockableDestroyableOwnerTransferred)
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
		it.Event = new(AdminLockableDestroyableOwnerTransferred)
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
func (it *AdminLockableDestroyableOwnerTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AdminLockableDestroyableOwnerTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AdminLockableDestroyableOwnerTransferred represents a OwnerTransferred event raised by the AdminLockableDestroyable contract.
type AdminLockableDestroyableOwnerTransferred struct {
	OldOwner common.Address
	NewOwner common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOwnerTransferred is a free log retrieval operation binding the contract event 0x8934ce4adea8d9ce0d714d2c22b86790e41b7731c84b926fbbdc1d40ff6533c9.
//
// Solidity: e OwnerTransferred(oldOwner indexed address, newOwner indexed address)
func (_AdminLockableDestroyable *AdminLockableDestroyableFilterer) FilterOwnerTransferred(opts *bind.FilterOpts, oldOwner []common.Address, newOwner []common.Address) (*AdminLockableDestroyableOwnerTransferredIterator, error) {

	var oldOwnerRule []interface{}
	for _, oldOwnerItem := range oldOwner {
		oldOwnerRule = append(oldOwnerRule, oldOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _AdminLockableDestroyable.contract.FilterLogs(opts, "OwnerTransferred", oldOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &AdminLockableDestroyableOwnerTransferredIterator{contract: _AdminLockableDestroyable.contract, event: "OwnerTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnerTransferred is a free log subscription operation binding the contract event 0x8934ce4adea8d9ce0d714d2c22b86790e41b7731c84b926fbbdc1d40ff6533c9.
//
// Solidity: e OwnerTransferred(oldOwner indexed address, newOwner indexed address)
func (_AdminLockableDestroyable *AdminLockableDestroyableFilterer) WatchOwnerTransferred(opts *bind.WatchOpts, sink chan<- *AdminLockableDestroyableOwnerTransferred, oldOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var oldOwnerRule []interface{}
	for _, oldOwnerItem := range oldOwner {
		oldOwnerRule = append(oldOwnerRule, oldOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _AdminLockableDestroyable.contract.WatchLogs(opts, "OwnerTransferred", oldOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AdminLockableDestroyableOwnerTransferred)
				if err := _AdminLockableDestroyable.contract.UnpackLog(event, "OwnerTransferred", log); err != nil {
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
