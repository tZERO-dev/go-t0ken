// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package registry

import (
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// CustodianABI is the input ABI used to generate the binding from.
const CustodianABI = "[{\"constant\":false,\"inputs\":[{\"internalType\":\"bool\",\"name\":\"locked\",\"type\":\"bool\"}],\"name\":\"setLocked\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"custodian\",\"type\":\"address\"}],\"name\":\"remove\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"custodian\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"}],\"name\":\"add\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"kill\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwner\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ZERO_ADDRESS\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"registry\",\"outputs\":[{\"internalType\":\"contractIRegistry\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isLocked\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"contractIRegistry\",\"name\":\"r\",\"type\":\"address\"}],\"name\":\"setRegistry\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"custodian\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"frozen\",\"type\":\"bool\"}],\"name\":\"setFrozen\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIRegistry\",\"name\":\"r\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnerTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"custodian\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"CustodianAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"custodian\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"CustodianRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"custodian\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bool\",\"name\":\"frozen\",\"type\":\"bool\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"CustodianFrozen\",\"type\":\"event\"}]"

// CustodianBin is the compiled bytecode used for deploying new contracts.
const CustodianBin = `6080604052600180546001600160a01b031916905534801561002057600080fd5b50604051610bb0380380610bb08339818101604052602081101561004357600080fd5b5051600080546001600160a01b031990811633179091556001805460ff60a01b19169055600280546001600160a01b0390931692909116919091179055610b218061008f6000396000f3fe608060405234801561001057600080fd5b50600436106100c95760003560e01c8063538ba4f911610081578063a4e2d6341161005b578063a4e2d634146101a3578063a91ee0dc146101bf578063ac869cd8146101e5576100c9565b8063538ba4f91461016f5780637b103999146101935780638da5cb5b1461019b576100c9565b80632d580ef6116100b25780632d580ef61461011557806341c0e1b5146101415780634fb2e45d14610149576100c9565b8063211e28b6146100ce57806329092d0e146100ef575b600080fd5b6100ed600480360360208110156100e457600080fd5b50351515610213565b005b6100ed6004803603602081101561010557600080fd5b50356001600160a01b0316610316565b6100ed6004803603604081101561012b57600080fd5b506001600160a01b038135169060200135610492565b6100ed610629565b6100ed6004803603602081101561015f57600080fd5b50356001600160a01b03166106b1565b610177610830565b604080516001600160a01b039092168252519081900360200190f35b61017761083f565b61017761084e565b6101ab61085d565b604080519115158252519081900360200190f35b6100ed600480360360208110156101d557600080fd5b50356001600160a01b031661086d565b6100ed600480360360408110156101fb57600080fd5b506001600160a01b0381351690602001351515610916565b6000546001600160a01b031633148061023c57506001546000546001600160a01b039081169116145b61028d576040805162461bcd60e51b815260206004820152601960248201527f4f776e6572206163636f756e7420697320726571756972656400000000000000604482015290519081900360640190fd5b60015460ff600160a01b90910416151581151514156102dd5760405162461bcd60e51b8152600401808060200182810382526028815260200180610ac56028913960400191505060405180910390fd5b60018054911515600160a01b027fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff909216919091179055565b60025460408051639d44ac4f60e01b81526001600482015233602482015290516001600160a01b0390921691639d44ac4f91604480820192602092909190829003018186803b15801561036857600080fd5b505afa15801561037c573d6000803e3d6000fd5b505050506040513d602081101561039257600080fd5b50516103da576040805162461bcd60e51b815260206004820152601260248201527126b4b9b9b4b733903832b936b4b9b9b4b7b760711b604482015290519081900360640190fd5b600254604080517fc4740a950000000000000000000000000000000000000000000000000000000081526001600160a01b0384811660048301529151919092169163c4740a9591602480830192600092919082900301818387803b15801561044157600080fd5b505af1158015610455573d6000803e3d6000fd5b50506040513392506001600160a01b03841691507f1c0cdcc74010449d4477f9576aaf31cee9e18d2611031462fcd5bf5329dec88e90600090a350565b60025460408051639d44ac4f60e01b81526001600482015233602482015290516001600160a01b0390921691639d44ac4f91604480820192602092909190829003018186803b1580156104e457600080fd5b505afa1580156104f8573d6000803e3d6000fd5b505050506040513d602081101561050e57600080fd5b5051610556576040805162461bcd60e51b815260206004820152601260248201527126b4b9b9b4b733903832b936b4b9b9b4b7b760711b604482015290519081900360640190fd5b600254604080517fe8d74d1f0000000000000000000000000000000000000000000000000000000081526001600160a01b038581166004830152600160248301526000604483018190526064830181905260848301869052925193169263e8d74d1f9260a48084019391929182900301818387803b1580156105d757600080fd5b505af11580156105eb573d6000803e3d6000fd5b50506040513392506001600160a01b03851691507f2c79891dd909910a9ed2116eb868a4783a4abeb6da7c8dc217cd5b198b329b5d90600090a35050565b6000546001600160a01b031633148061065257506001546000546001600160a01b039081169116145b6106a3576040805162461bcd60e51b815260206004820152601960248201527f4f776e6572206163636f756e7420697320726571756972656400000000000000604482015290519081900360640190fd5b6000546001600160a01b0316ff5b6000546001600160a01b03163314806106da57506001546000546001600160a01b039081169116145b61072b576040805162461bcd60e51b815260206004820152601960248201527f4f776e6572206163636f756e7420697320726571756972656400000000000000604482015290519081900360640190fd5b6000546001600160a01b03828116911614156107785760405162461bcd60e51b8152600401808060200182810382526025815260200180610aa06025913960400191505060405180910390fd5b6001600160a01b0381166107d3576040805162461bcd60e51b815260206004820181905260248201527f4e6577204f776e65722063616e6e6f74206265207a65726f2061646472657373604482015290519081900360640190fd5b600080546001600160a01b0383811673ffffffffffffffffffffffffffffffffffffffff19831681178455604051919092169283917f8934ce4adea8d9ce0d714d2c22b86790e41b7731c84b926fbbdc1d40ff6533c99190a35050565b6001546001600160a01b031681565b6002546001600160a01b031681565b6000546001600160a01b031681565b600154600160a01b900460ff1681565b6000546001600160a01b031633148061089657506001546000546001600160a01b039081169116145b6108e7576040805162461bcd60e51b815260206004820152601960248201527f4f776e6572206163636f756e7420697320726571756972656400000000000000604482015290519081900360640190fd5b6002805473ffffffffffffffffffffffffffffffffffffffff19166001600160a01b0392909216919091179055565b60025460408051639d44ac4f60e01b81526001600482015233602482015290516001600160a01b0390921691639d44ac4f91604480820192602092909190829003018186803b15801561096857600080fd5b505afa15801561097c573d6000803e3d6000fd5b505050506040513d602081101561099257600080fd5b50516109da576040805162461bcd60e51b815260206004820152601260248201527126b4b9b9b4b733903832b936b4b9b9b4b7b760711b604482015290519081900360640190fd5b600254604080517fcbe5404f0000000000000000000000000000000000000000000000000000000081526001600160a01b03858116600483015284151560248301529151919092169163cbe5404f91604480830192600092919082900301818387803b158015610a4957600080fd5b505af1158015610a5d573d6000803e3d6000fd5b505060405133925083151591506001600160a01b038516907f41c9c2bee97c3cc15014e3280012a23b12872bc0485aebfb86d0a0ea9c3afc9c90600090a4505056fe4e6577204f776e65722063616e6e6f74206265207468652063757272656e74206f776e6572436f6e747261637420616c726561647920696e20726571756573746564206c6f636b207374617465a265627a7a72315820f92d68445f69fb46d6fcdb7fd789539e2f187d46e5811350e6cd66300e621b4564736f6c634300050b0032`

// DeployCustodian deploys a new Ethereum contract, binding an instance of Custodian to it.
func DeployCustodian(auth *bind.TransactOpts, backend bind.ContractBackend, r common.Address) (common.Address, *types.Transaction, *Custodian, error) {
	parsed, err := abi.JSON(strings.NewReader(CustodianABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(CustodianBin), backend, r)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Custodian{CustodianCaller: CustodianCaller{contract: contract}, CustodianTransactor: CustodianTransactor{contract: contract}, CustodianFilterer: CustodianFilterer{contract: contract}}, nil
}

// Custodian is an auto generated Go binding around an Ethereum contract.
type Custodian struct {
	CustodianCaller     // Read-only binding to the contract
	CustodianTransactor // Write-only binding to the contract
	CustodianFilterer   // Log filterer for contract events
}

// CustodianCaller is an auto generated read-only Go binding around an Ethereum contract.
type CustodianCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CustodianTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CustodianTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CustodianFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CustodianFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CustodianSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CustodianSession struct {
	Contract     *Custodian        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CustodianCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CustodianCallerSession struct {
	Contract *CustodianCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// CustodianTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CustodianTransactorSession struct {
	Contract     *CustodianTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// CustodianRaw is an auto generated low-level Go binding around an Ethereum contract.
type CustodianRaw struct {
	Contract *Custodian // Generic contract binding to access the raw methods on
}

// CustodianCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CustodianCallerRaw struct {
	Contract *CustodianCaller // Generic read-only contract binding to access the raw methods on
}

// CustodianTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CustodianTransactorRaw struct {
	Contract *CustodianTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCustodian creates a new instance of Custodian, bound to a specific deployed contract.
func NewCustodian(address common.Address, backend bind.ContractBackend) (*Custodian, error) {
	contract, err := bindCustodian(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Custodian{CustodianCaller: CustodianCaller{contract: contract}, CustodianTransactor: CustodianTransactor{contract: contract}, CustodianFilterer: CustodianFilterer{contract: contract}}, nil
}

// NewCustodianCaller creates a new read-only instance of Custodian, bound to a specific deployed contract.
func NewCustodianCaller(address common.Address, caller bind.ContractCaller) (*CustodianCaller, error) {
	contract, err := bindCustodian(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CustodianCaller{contract: contract}, nil
}

// NewCustodianTransactor creates a new write-only instance of Custodian, bound to a specific deployed contract.
func NewCustodianTransactor(address common.Address, transactor bind.ContractTransactor) (*CustodianTransactor, error) {
	contract, err := bindCustodian(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CustodianTransactor{contract: contract}, nil
}

// NewCustodianFilterer creates a new log filterer instance of Custodian, bound to a specific deployed contract.
func NewCustodianFilterer(address common.Address, filterer bind.ContractFilterer) (*CustodianFilterer, error) {
	contract, err := bindCustodian(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CustodianFilterer{contract: contract}, nil
}

// bindCustodian binds a generic wrapper to an already deployed contract.
func bindCustodian(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(CustodianABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Custodian *CustodianRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Custodian.Contract.CustodianCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Custodian *CustodianRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Custodian.Contract.CustodianTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Custodian *CustodianRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Custodian.Contract.CustodianTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Custodian *CustodianCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Custodian.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Custodian *CustodianTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Custodian.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Custodian *CustodianTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Custodian.Contract.contract.Transact(opts, method, params...)
}

// ZEROADDRESS is a free data retrieval call binding the contract method 0x538ba4f9.
//
// Solidity: function ZERO_ADDRESS() constant returns(address)
func (_Custodian *CustodianCaller) ZEROADDRESS(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Custodian.contract.Call(opts, out, "ZERO_ADDRESS")
	return *ret0, err
}

// ZEROADDRESS is a free data retrieval call binding the contract method 0x538ba4f9.
//
// Solidity: function ZERO_ADDRESS() constant returns(address)
func (_Custodian *CustodianSession) ZEROADDRESS() (common.Address, error) {
	return _Custodian.Contract.ZEROADDRESS(&_Custodian.CallOpts)
}

// ZEROADDRESS is a free data retrieval call binding the contract method 0x538ba4f9.
//
// Solidity: function ZERO_ADDRESS() constant returns(address)
func (_Custodian *CustodianCallerSession) ZEROADDRESS() (common.Address, error) {
	return _Custodian.Contract.ZEROADDRESS(&_Custodian.CallOpts)
}

// IsLocked is a free data retrieval call binding the contract method 0xa4e2d634.
//
// Solidity: function isLocked() constant returns(bool)
func (_Custodian *CustodianCaller) IsLocked(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Custodian.contract.Call(opts, out, "isLocked")
	return *ret0, err
}

// IsLocked is a free data retrieval call binding the contract method 0xa4e2d634.
//
// Solidity: function isLocked() constant returns(bool)
func (_Custodian *CustodianSession) IsLocked() (bool, error) {
	return _Custodian.Contract.IsLocked(&_Custodian.CallOpts)
}

// IsLocked is a free data retrieval call binding the contract method 0xa4e2d634.
//
// Solidity: function isLocked() constant returns(bool)
func (_Custodian *CustodianCallerSession) IsLocked() (bool, error) {
	return _Custodian.Contract.IsLocked(&_Custodian.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Custodian *CustodianCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Custodian.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Custodian *CustodianSession) Owner() (common.Address, error) {
	return _Custodian.Contract.Owner(&_Custodian.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Custodian *CustodianCallerSession) Owner() (common.Address, error) {
	return _Custodian.Contract.Owner(&_Custodian.CallOpts)
}

// Registry is a free data retrieval call binding the contract method 0x7b103999.
//
// Solidity: function registry() constant returns(address)
func (_Custodian *CustodianCaller) Registry(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Custodian.contract.Call(opts, out, "registry")
	return *ret0, err
}

// Registry is a free data retrieval call binding the contract method 0x7b103999.
//
// Solidity: function registry() constant returns(address)
func (_Custodian *CustodianSession) Registry() (common.Address, error) {
	return _Custodian.Contract.Registry(&_Custodian.CallOpts)
}

// Registry is a free data retrieval call binding the contract method 0x7b103999.
//
// Solidity: function registry() constant returns(address)
func (_Custodian *CustodianCallerSession) Registry() (common.Address, error) {
	return _Custodian.Contract.Registry(&_Custodian.CallOpts)
}

// Add is a paid mutator transaction binding the contract method 0x2d580ef6.
//
// Solidity: function add(custodian address, hash bytes32) returns()
func (_Custodian *CustodianTransactor) Add(opts *bind.TransactOpts, custodian common.Address, hash [32]byte) (*types.Transaction, error) {
	return _Custodian.contract.Transact(opts, "add", custodian, hash)
}

// Add is a paid mutator transaction binding the contract method 0x2d580ef6.
//
// Solidity: function add(custodian address, hash bytes32) returns()
func (_Custodian *CustodianSession) Add(custodian common.Address, hash [32]byte) (*types.Transaction, error) {
	return _Custodian.Contract.Add(&_Custodian.TransactOpts, custodian, hash)
}

// Add is a paid mutator transaction binding the contract method 0x2d580ef6.
//
// Solidity: function add(custodian address, hash bytes32) returns()
func (_Custodian *CustodianTransactorSession) Add(custodian common.Address, hash [32]byte) (*types.Transaction, error) {
	return _Custodian.Contract.Add(&_Custodian.TransactOpts, custodian, hash)
}

// Kill is a paid mutator transaction binding the contract method 0x41c0e1b5.
//
// Solidity: function kill() returns()
func (_Custodian *CustodianTransactor) Kill(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Custodian.contract.Transact(opts, "kill")
}

// Kill is a paid mutator transaction binding the contract method 0x41c0e1b5.
//
// Solidity: function kill() returns()
func (_Custodian *CustodianSession) Kill() (*types.Transaction, error) {
	return _Custodian.Contract.Kill(&_Custodian.TransactOpts)
}

// Kill is a paid mutator transaction binding the contract method 0x41c0e1b5.
//
// Solidity: function kill() returns()
func (_Custodian *CustodianTransactorSession) Kill() (*types.Transaction, error) {
	return _Custodian.Contract.Kill(&_Custodian.TransactOpts)
}

// Remove is a paid mutator transaction binding the contract method 0x29092d0e.
//
// Solidity: function remove(custodian address) returns()
func (_Custodian *CustodianTransactor) Remove(opts *bind.TransactOpts, custodian common.Address) (*types.Transaction, error) {
	return _Custodian.contract.Transact(opts, "remove", custodian)
}

// Remove is a paid mutator transaction binding the contract method 0x29092d0e.
//
// Solidity: function remove(custodian address) returns()
func (_Custodian *CustodianSession) Remove(custodian common.Address) (*types.Transaction, error) {
	return _Custodian.Contract.Remove(&_Custodian.TransactOpts, custodian)
}

// Remove is a paid mutator transaction binding the contract method 0x29092d0e.
//
// Solidity: function remove(custodian address) returns()
func (_Custodian *CustodianTransactorSession) Remove(custodian common.Address) (*types.Transaction, error) {
	return _Custodian.Contract.Remove(&_Custodian.TransactOpts, custodian)
}

// SetFrozen is a paid mutator transaction binding the contract method 0xac869cd8.
//
// Solidity: function setFrozen(custodian address, frozen bool) returns()
func (_Custodian *CustodianTransactor) SetFrozen(opts *bind.TransactOpts, custodian common.Address, frozen bool) (*types.Transaction, error) {
	return _Custodian.contract.Transact(opts, "setFrozen", custodian, frozen)
}

// SetFrozen is a paid mutator transaction binding the contract method 0xac869cd8.
//
// Solidity: function setFrozen(custodian address, frozen bool) returns()
func (_Custodian *CustodianSession) SetFrozen(custodian common.Address, frozen bool) (*types.Transaction, error) {
	return _Custodian.Contract.SetFrozen(&_Custodian.TransactOpts, custodian, frozen)
}

// SetFrozen is a paid mutator transaction binding the contract method 0xac869cd8.
//
// Solidity: function setFrozen(custodian address, frozen bool) returns()
func (_Custodian *CustodianTransactorSession) SetFrozen(custodian common.Address, frozen bool) (*types.Transaction, error) {
	return _Custodian.Contract.SetFrozen(&_Custodian.TransactOpts, custodian, frozen)
}

// SetLocked is a paid mutator transaction binding the contract method 0x211e28b6.
//
// Solidity: function setLocked(locked bool) returns()
func (_Custodian *CustodianTransactor) SetLocked(opts *bind.TransactOpts, locked bool) (*types.Transaction, error) {
	return _Custodian.contract.Transact(opts, "setLocked", locked)
}

// SetLocked is a paid mutator transaction binding the contract method 0x211e28b6.
//
// Solidity: function setLocked(locked bool) returns()
func (_Custodian *CustodianSession) SetLocked(locked bool) (*types.Transaction, error) {
	return _Custodian.Contract.SetLocked(&_Custodian.TransactOpts, locked)
}

// SetLocked is a paid mutator transaction binding the contract method 0x211e28b6.
//
// Solidity: function setLocked(locked bool) returns()
func (_Custodian *CustodianTransactorSession) SetLocked(locked bool) (*types.Transaction, error) {
	return _Custodian.Contract.SetLocked(&_Custodian.TransactOpts, locked)
}

// SetRegistry is a paid mutator transaction binding the contract method 0xa91ee0dc.
//
// Solidity: function setRegistry(r address) returns()
func (_Custodian *CustodianTransactor) SetRegistry(opts *bind.TransactOpts, r common.Address) (*types.Transaction, error) {
	return _Custodian.contract.Transact(opts, "setRegistry", r)
}

// SetRegistry is a paid mutator transaction binding the contract method 0xa91ee0dc.
//
// Solidity: function setRegistry(r address) returns()
func (_Custodian *CustodianSession) SetRegistry(r common.Address) (*types.Transaction, error) {
	return _Custodian.Contract.SetRegistry(&_Custodian.TransactOpts, r)
}

// SetRegistry is a paid mutator transaction binding the contract method 0xa91ee0dc.
//
// Solidity: function setRegistry(r address) returns()
func (_Custodian *CustodianTransactorSession) SetRegistry(r common.Address) (*types.Transaction, error) {
	return _Custodian.Contract.SetRegistry(&_Custodian.TransactOpts, r)
}

// TransferOwner is a paid mutator transaction binding the contract method 0x4fb2e45d.
//
// Solidity: function transferOwner(newOwner address) returns()
func (_Custodian *CustodianTransactor) TransferOwner(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Custodian.contract.Transact(opts, "transferOwner", newOwner)
}

// TransferOwner is a paid mutator transaction binding the contract method 0x4fb2e45d.
//
// Solidity: function transferOwner(newOwner address) returns()
func (_Custodian *CustodianSession) TransferOwner(newOwner common.Address) (*types.Transaction, error) {
	return _Custodian.Contract.TransferOwner(&_Custodian.TransactOpts, newOwner)
}

// TransferOwner is a paid mutator transaction binding the contract method 0x4fb2e45d.
//
// Solidity: function transferOwner(newOwner address) returns()
func (_Custodian *CustodianTransactorSession) TransferOwner(newOwner common.Address) (*types.Transaction, error) {
	return _Custodian.Contract.TransferOwner(&_Custodian.TransactOpts, newOwner)
}

// CustodianCustodianAddedIterator is returned from FilterCustodianAdded and is used to iterate over the raw logs and unpacked data for CustodianAdded events raised by the Custodian contract.
type CustodianCustodianAddedIterator struct {
	Event *CustodianCustodianAdded // Event containing the contract specifics and raw log

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
func (it *CustodianCustodianAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CustodianCustodianAdded)
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
		it.Event = new(CustodianCustodianAdded)
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
func (it *CustodianCustodianAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CustodianCustodianAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CustodianCustodianAdded represents a CustodianAdded event raised by the Custodian contract.
type CustodianCustodianAdded struct {
	Custodian common.Address
	Owner     common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterCustodianAdded is a free log retrieval operation binding the contract event 0x2c79891dd909910a9ed2116eb868a4783a4abeb6da7c8dc217cd5b198b329b5d.
//
// Solidity: e CustodianAdded(custodian indexed address, owner indexed address)
func (_Custodian *CustodianFilterer) FilterCustodianAdded(opts *bind.FilterOpts, custodian []common.Address, owner []common.Address) (*CustodianCustodianAddedIterator, error) {

	var custodianRule []interface{}
	for _, custodianItem := range custodian {
		custodianRule = append(custodianRule, custodianItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Custodian.contract.FilterLogs(opts, "CustodianAdded", custodianRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return &CustodianCustodianAddedIterator{contract: _Custodian.contract, event: "CustodianAdded", logs: logs, sub: sub}, nil
}

// WatchCustodianAdded is a free log subscription operation binding the contract event 0x2c79891dd909910a9ed2116eb868a4783a4abeb6da7c8dc217cd5b198b329b5d.
//
// Solidity: e CustodianAdded(custodian indexed address, owner indexed address)
func (_Custodian *CustodianFilterer) WatchCustodianAdded(opts *bind.WatchOpts, sink chan<- *CustodianCustodianAdded, custodian []common.Address, owner []common.Address) (event.Subscription, error) {

	var custodianRule []interface{}
	for _, custodianItem := range custodian {
		custodianRule = append(custodianRule, custodianItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Custodian.contract.WatchLogs(opts, "CustodianAdded", custodianRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CustodianCustodianAdded)
				if err := _Custodian.contract.UnpackLog(event, "CustodianAdded", log); err != nil {
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

// CustodianCustodianFrozenIterator is returned from FilterCustodianFrozen and is used to iterate over the raw logs and unpacked data for CustodianFrozen events raised by the Custodian contract.
type CustodianCustodianFrozenIterator struct {
	Event *CustodianCustodianFrozen // Event containing the contract specifics and raw log

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
func (it *CustodianCustodianFrozenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CustodianCustodianFrozen)
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
		it.Event = new(CustodianCustodianFrozen)
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
func (it *CustodianCustodianFrozenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CustodianCustodianFrozenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CustodianCustodianFrozen represents a CustodianFrozen event raised by the Custodian contract.
type CustodianCustodianFrozen struct {
	Custodian common.Address
	Frozen    bool
	Owner     common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterCustodianFrozen is a free log retrieval operation binding the contract event 0x41c9c2bee97c3cc15014e3280012a23b12872bc0485aebfb86d0a0ea9c3afc9c.
//
// Solidity: e CustodianFrozen(custodian indexed address, frozen indexed bool, owner indexed address)
func (_Custodian *CustodianFilterer) FilterCustodianFrozen(opts *bind.FilterOpts, custodian []common.Address, frozen []bool, owner []common.Address) (*CustodianCustodianFrozenIterator, error) {

	var custodianRule []interface{}
	for _, custodianItem := range custodian {
		custodianRule = append(custodianRule, custodianItem)
	}
	var frozenRule []interface{}
	for _, frozenItem := range frozen {
		frozenRule = append(frozenRule, frozenItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Custodian.contract.FilterLogs(opts, "CustodianFrozen", custodianRule, frozenRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return &CustodianCustodianFrozenIterator{contract: _Custodian.contract, event: "CustodianFrozen", logs: logs, sub: sub}, nil
}

// WatchCustodianFrozen is a free log subscription operation binding the contract event 0x41c9c2bee97c3cc15014e3280012a23b12872bc0485aebfb86d0a0ea9c3afc9c.
//
// Solidity: e CustodianFrozen(custodian indexed address, frozen indexed bool, owner indexed address)
func (_Custodian *CustodianFilterer) WatchCustodianFrozen(opts *bind.WatchOpts, sink chan<- *CustodianCustodianFrozen, custodian []common.Address, frozen []bool, owner []common.Address) (event.Subscription, error) {

	var custodianRule []interface{}
	for _, custodianItem := range custodian {
		custodianRule = append(custodianRule, custodianItem)
	}
	var frozenRule []interface{}
	for _, frozenItem := range frozen {
		frozenRule = append(frozenRule, frozenItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Custodian.contract.WatchLogs(opts, "CustodianFrozen", custodianRule, frozenRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CustodianCustodianFrozen)
				if err := _Custodian.contract.UnpackLog(event, "CustodianFrozen", log); err != nil {
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

// CustodianCustodianRemovedIterator is returned from FilterCustodianRemoved and is used to iterate over the raw logs and unpacked data for CustodianRemoved events raised by the Custodian contract.
type CustodianCustodianRemovedIterator struct {
	Event *CustodianCustodianRemoved // Event containing the contract specifics and raw log

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
func (it *CustodianCustodianRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CustodianCustodianRemoved)
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
		it.Event = new(CustodianCustodianRemoved)
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
func (it *CustodianCustodianRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CustodianCustodianRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CustodianCustodianRemoved represents a CustodianRemoved event raised by the Custodian contract.
type CustodianCustodianRemoved struct {
	Custodian common.Address
	Owner     common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterCustodianRemoved is a free log retrieval operation binding the contract event 0x1c0cdcc74010449d4477f9576aaf31cee9e18d2611031462fcd5bf5329dec88e.
//
// Solidity: e CustodianRemoved(custodian indexed address, owner indexed address)
func (_Custodian *CustodianFilterer) FilterCustodianRemoved(opts *bind.FilterOpts, custodian []common.Address, owner []common.Address) (*CustodianCustodianRemovedIterator, error) {

	var custodianRule []interface{}
	for _, custodianItem := range custodian {
		custodianRule = append(custodianRule, custodianItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Custodian.contract.FilterLogs(opts, "CustodianRemoved", custodianRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return &CustodianCustodianRemovedIterator{contract: _Custodian.contract, event: "CustodianRemoved", logs: logs, sub: sub}, nil
}

// WatchCustodianRemoved is a free log subscription operation binding the contract event 0x1c0cdcc74010449d4477f9576aaf31cee9e18d2611031462fcd5bf5329dec88e.
//
// Solidity: e CustodianRemoved(custodian indexed address, owner indexed address)
func (_Custodian *CustodianFilterer) WatchCustodianRemoved(opts *bind.WatchOpts, sink chan<- *CustodianCustodianRemoved, custodian []common.Address, owner []common.Address) (event.Subscription, error) {

	var custodianRule []interface{}
	for _, custodianItem := range custodian {
		custodianRule = append(custodianRule, custodianItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Custodian.contract.WatchLogs(opts, "CustodianRemoved", custodianRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CustodianCustodianRemoved)
				if err := _Custodian.contract.UnpackLog(event, "CustodianRemoved", log); err != nil {
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

// CustodianOwnerTransferredIterator is returned from FilterOwnerTransferred and is used to iterate over the raw logs and unpacked data for OwnerTransferred events raised by the Custodian contract.
type CustodianOwnerTransferredIterator struct {
	Event *CustodianOwnerTransferred // Event containing the contract specifics and raw log

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
func (it *CustodianOwnerTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CustodianOwnerTransferred)
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
		it.Event = new(CustodianOwnerTransferred)
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
func (it *CustodianOwnerTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CustodianOwnerTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CustodianOwnerTransferred represents a OwnerTransferred event raised by the Custodian contract.
type CustodianOwnerTransferred struct {
	OldOwner common.Address
	NewOwner common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOwnerTransferred is a free log retrieval operation binding the contract event 0x8934ce4adea8d9ce0d714d2c22b86790e41b7731c84b926fbbdc1d40ff6533c9.
//
// Solidity: e OwnerTransferred(oldOwner indexed address, newOwner indexed address)
func (_Custodian *CustodianFilterer) FilterOwnerTransferred(opts *bind.FilterOpts, oldOwner []common.Address, newOwner []common.Address) (*CustodianOwnerTransferredIterator, error) {

	var oldOwnerRule []interface{}
	for _, oldOwnerItem := range oldOwner {
		oldOwnerRule = append(oldOwnerRule, oldOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Custodian.contract.FilterLogs(opts, "OwnerTransferred", oldOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &CustodianOwnerTransferredIterator{contract: _Custodian.contract, event: "OwnerTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnerTransferred is a free log subscription operation binding the contract event 0x8934ce4adea8d9ce0d714d2c22b86790e41b7731c84b926fbbdc1d40ff6533c9.
//
// Solidity: e OwnerTransferred(oldOwner indexed address, newOwner indexed address)
func (_Custodian *CustodianFilterer) WatchOwnerTransferred(opts *bind.WatchOpts, sink chan<- *CustodianOwnerTransferred, oldOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var oldOwnerRule []interface{}
	for _, oldOwnerItem := range oldOwner {
		oldOwnerRule = append(oldOwnerRule, oldOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Custodian.contract.WatchLogs(opts, "OwnerTransferred", oldOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CustodianOwnerTransferred)
				if err := _Custodian.contract.UnpackLog(event, "OwnerTransferred", log); err != nil {
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
