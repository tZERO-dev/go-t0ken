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
const AdminLockableABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"admin\",\"type\":\"address\"}],\"name\":\"removeAdmin\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"locked\",\"type\":\"bool\"}],\"name\":\"setLocked\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"isAdmin\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwner\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"admin\",\"type\":\"address\"}],\"name\":\"addAdmin\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isLocked\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"admins\",\"outputs\":[{\"name\":\"count\",\"type\":\"int256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"admin\",\"type\":\"address\"}],\"name\":\"AdminAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"admin\",\"type\":\"address\"}],\"name\":\"AdminRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"oldOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnerTransferred\",\"type\":\"event\"}]"

// AdminLockableBin is the compiled bytecode used for deploying new contracts.
const AdminLockableBin = `608060405234801561001057600080fd5b5060008054600160a060020a031916331790556004805460ff19169055610ab58061003c6000396000f3fe608060405234801561001057600080fd5b50600436106100a5576000357c010000000000000000000000000000000000000000000000000000000090048063704802751161007857806370480275146101785780638da5cb5b146101ab578063a4e2d634146101dc578063a5de3619146101e4576100a5565b80631785f53c146100aa578063211e28b6146100df57806324d7806c146100fe5780634fb2e45d14610145575b600080fd5b6100dd600480360360208110156100c057600080fd5b503573ffffffffffffffffffffffffffffffffffffffff166101fe565b005b6100dd600480360360208110156100f557600080fd5b50351515610346565b6101316004803603602081101561011457600080fd5b503573ffffffffffffffffffffffffffffffffffffffff1661045a565b604080519115158252519081900360200190f35b6100dd6004803603602081101561015b57600080fd5b503573ffffffffffffffffffffffffffffffffffffffff16610496565b6100dd6004803603602081101561018e57600080fd5b503573ffffffffffffffffffffffffffffffffffffffff1661067e565b6101b36107c6565b6040805173ffffffffffffffffffffffffffffffffffffffff9092168252519081900360200190f35b6101316107e2565b6101ec6107eb565b60408051918252519081900360200190f35b60005473ffffffffffffffffffffffffffffffffffffffff16331461028457604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601960248201527f4f776e6572206163636f756e7420697320726571756972656400000000000000604482015290519081900360640190fd5b61029560018263ffffffff6107f116565b151561030257604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f556e61626c6520746f2072656d6f76652061646d696e00000000000000000000604482015290519081900360640190fd5b60405173ffffffffffffffffffffffffffffffffffffffff8216907fa3b62bc36326052d97ea62d63c3d60308ed4c3ea8ac079dd8499f1e9c4f80c0f90600090a250565b60005473ffffffffffffffffffffffffffffffffffffffff16331480610378575061037860013363ffffffff61091d16565b15156103e557604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601960248201527f41646d696e206163636f756e7420697320726571756972656400000000000000604482015290519081900360640190fd5b60045460ff1615158115151415610447576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401808060200182810382526028815260200180610a626028913960400191505060405180910390fd5b6004805460ff1916911515919091179055565b6000805473ffffffffffffffffffffffffffffffffffffffff83811691161480610490575061049060018363ffffffff61091d16565b92915050565b60005473ffffffffffffffffffffffffffffffffffffffff16331461051c57604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601960248201527f4f776e6572206163636f756e7420697320726571756972656400000000000000604482015290519081900360640190fd5b60005473ffffffffffffffffffffffffffffffffffffffff82811691161415610590576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401808060200182810382526025815260200180610a3d6025913960400191505060405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff8116151561061457604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820181905260248201527f4e6577204f776e65722063616e6e6f74206265207a65726f2061646472657373604482015290519081900360640190fd5b6000805473ffffffffffffffffffffffffffffffffffffffff83811673ffffffffffffffffffffffffffffffffffffffff19831681178455604051919092169283917f8934ce4adea8d9ce0d714d2c22b86790e41b7731c84b926fbbdc1d40ff6533c99190a35050565b60005473ffffffffffffffffffffffffffffffffffffffff16331461070457604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601960248201527f4f776e6572206163636f756e7420697320726571756972656400000000000000604482015290519081900360640190fd5b61071560018263ffffffff61096116565b151561078257604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601360248201527f556e61626c6520746f206164642061646d696e00000000000000000000000000604482015290519081900360640190fd5b60405173ffffffffffffffffffffffffffffffffffffffff8216907f44d6d25963f097ad14f29f06854a01f575648a1ef82f30e562ccd3889717e33990600090a250565b60005473ffffffffffffffffffffffffffffffffffffffff1681565b60045460ff1681565b60015481565b73ffffffffffffffffffffffffffffffffffffffff8116600090815260018084016020526040822054908112806108285750835481135b15610837576000915050610490565b83548112156108b8578354600090815260028501602081815260408084205473ffffffffffffffffffffffffffffffffffffffff1680855260018901835281852086905585855292909152808320805473ffffffffffffffffffffffffffffffffffffffff19908116909317905586548352909120805490911690556108e4565b60008181526002850160205260409020805473ffffffffffffffffffffffffffffffffffffffff191690555b505073ffffffffffffffffffffffffffffffffffffffff1660009081526001828101602052604082209190915581546000190190915590565b73ffffffffffffffffffffffffffffffffffffffff81166000908152600183016020526040812054600019018181128015906109595750835481125b949350505050565b600073ffffffffffffffffffffffffffffffffffffffff8216151561098857506000610490565b73ffffffffffffffffffffffffffffffffffffffff82166000908152600184016020526040812054600019019081128015906109c45750835481125b156109d3576000915050610490565b50508154600190810180845573ffffffffffffffffffffffffffffffffffffffff831660008181528386016020908152604080832085905593825260028701905291909120805473ffffffffffffffffffffffffffffffffffffffff191690911790559291505056fe4e6577204f776e65722063616e6e6f74206265207468652063757272656e74206f776e6572436f6e747261637420616c726561647920696e20726571756573746564206c6f636b207374617465a165627a7a7230582008597d3f24bd2f0e739fa774b6acad047909fc8cf4697a2c85ec29c2eefb8c2f0029`

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

// AddAdmin is a paid mutator transaction binding the contract method 0x70480275.
//
// Solidity: function addAdmin(admin address) returns()
func (_AdminLockable *AdminLockableTransactor) AddAdmin(opts *bind.TransactOpts, admin common.Address) (*types.Transaction, error) {
	return _AdminLockable.contract.Transact(opts, "addAdmin", admin)
}

// AddAdmin is a paid mutator transaction binding the contract method 0x70480275.
//
// Solidity: function addAdmin(admin address) returns()
func (_AdminLockable *AdminLockableSession) AddAdmin(admin common.Address) (*types.Transaction, error) {
	return _AdminLockable.Contract.AddAdmin(&_AdminLockable.TransactOpts, admin)
}

// AddAdmin is a paid mutator transaction binding the contract method 0x70480275.
//
// Solidity: function addAdmin(admin address) returns()
func (_AdminLockable *AdminLockableTransactorSession) AddAdmin(admin common.Address) (*types.Transaction, error) {
	return _AdminLockable.Contract.AddAdmin(&_AdminLockable.TransactOpts, admin)
}

// RemoveAdmin is a paid mutator transaction binding the contract method 0x1785f53c.
//
// Solidity: function removeAdmin(admin address) returns()
func (_AdminLockable *AdminLockableTransactor) RemoveAdmin(opts *bind.TransactOpts, admin common.Address) (*types.Transaction, error) {
	return _AdminLockable.contract.Transact(opts, "removeAdmin", admin)
}

// RemoveAdmin is a paid mutator transaction binding the contract method 0x1785f53c.
//
// Solidity: function removeAdmin(admin address) returns()
func (_AdminLockable *AdminLockableSession) RemoveAdmin(admin common.Address) (*types.Transaction, error) {
	return _AdminLockable.Contract.RemoveAdmin(&_AdminLockable.TransactOpts, admin)
}

// RemoveAdmin is a paid mutator transaction binding the contract method 0x1785f53c.
//
// Solidity: function removeAdmin(admin address) returns()
func (_AdminLockable *AdminLockableTransactorSession) RemoveAdmin(admin common.Address) (*types.Transaction, error) {
	return _AdminLockable.Contract.RemoveAdmin(&_AdminLockable.TransactOpts, admin)
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
