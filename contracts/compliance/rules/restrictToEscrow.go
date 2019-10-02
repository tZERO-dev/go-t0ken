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

// RestrictToEscrowABI is the input ABI used to generate the binding from.
const RestrictToEscrowABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"contractICompliance\",\"name\":\"compliance\",\"type\":\"address\"},{\"internalType\":\"contractIT0ken\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"initiator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokens\",\"type\":\"uint256\"}],\"name\":\"test\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"s\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"kill\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwner\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ZERO_ADDRESS\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"contractIT0ken\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"initiator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokens\",\"type\":\"uint256\"}],\"name\":\"check\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnerTransferred\",\"type\":\"event\"}]"

// RestrictToEscrowBin is the compiled bytecode used for deploying new contracts.
const RestrictToEscrowBin = `600180546001600160a01b031916905560c0604052600f60808190527f457363726f77205265636f72646572000000000000000000000000000000000060a090815261004e91600291906100ab565b506040805180820190915260128082527f526573747269637420546f20457363726f7700000000000000000000000000006020909201918252610093916003916100ab565b50600080546001600160a01b03191633179055610146565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106100ec57805160ff1916838001178555610119565b82800160010185558215610119579182015b828111156101195782518255916020019190600101906100fe565b50610125929150610129565b5090565b61014391905b80821115610125576000815560010161012f565b90565b610990806101556000396000f3fe608060405234801561001057600080fd5b506004361061007d5760003560e01c80634fb2e45d1161005b5780634fb2e45d14610155578063538ba4f91461017b578063803fcd431461019f5780638da5cb5b146101e55761007d565b806306fdde03146100825780630b6dedfd146100ff57806341c0e1b51461014b575b600080fd5b61008a6101ed565b6040805160208082528351818301528351919283929083019185019080838360005b838110156100c45781810151838201526020016100ac565b50505050905090810190601f1680156100f15780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b61008a600480360360c081101561011557600080fd5b506001600160a01b0381358116916020810135821691604082013581169160608101358216916080820135169060a0013561027b565b61015361035a565b005b6101536004803603602081101561016b57600080fd5b50356001600160a01b03166103e2565b61018361056c565b604080516001600160a01b039092168252519081900360200190f35b610153600480360360a08110156101b557600080fd5b506001600160a01b038135811691602081013582169160408201358116916060810135909116906080013561057b565b610183610700565b6003805460408051602060026001851615610100026000190190941693909304601f810184900484028201840190925281815292918301828280156102735780601f1061024857610100808354040283529160200191610273565b820191906000526020600020905b81548152906001019060200180831161025657829003601f168201915b505050505081565b6060600061029061028a61070f565b8661077b565b90506001600160a01b0381166102dd576040518060400160405280601681526020017f496e76616c69642f4d697373696e6720657363726f7700000000000000000000815250915061034f565b836001600160a01b0316816001600160a01b0316141580156103115750846001600160a01b0316816001600160a01b031614155b1561034f576040518060400160405280601e81526020017f5472616e7366657273207265737472696374656420746f20657363726f77000081525091505b509695505050505050565b6000546001600160a01b031633148061038357506001546000546001600160a01b039081169116145b6103d4576040805162461bcd60e51b815260206004820152601960248201527f4f776e6572206163636f756e7420697320726571756972656400000000000000604482015290519081900360640190fd5b6000546001600160a01b0316ff5b6000546001600160a01b031633148061040b57506001546000546001600160a01b039081169116145b61045c576040805162461bcd60e51b815260206004820152601960248201527f4f776e6572206163636f756e7420697320726571756972656400000000000000604482015290519081900360640190fd5b6000546001600160a01b03828116911614156104a95760405162461bcd60e51b81526004018080602001828103825260258152602001806109376025913960400191505060405180910390fd5b6001600160a01b038116610504576040805162461bcd60e51b815260206004820181905260248201527f4e6577204f776e65722063616e6e6f74206265207a65726f2061646472657373604482015290519081900360640190fd5b600080546001600160a01b038381167fffffffffffffffffffffffff0000000000000000000000000000000000000000831681178455604051919092169283917f8934ce4adea8d9ce0d714d2c22b86790e41b7731c84b926fbbdc1d40ff6533c99190a35050565b6001546001600160a01b031681565b600061058e61058861070f565b8561077b565b90506001600160a01b0381166105eb576040805162461bcd60e51b815260206004820152601660248201527f496e76616c69642f4d697373696e6720657363726f7700000000000000000000604482015290519081900360640190fd5b826001600160a01b0316816001600160a01b0316148061061c5750836001600160a01b0316816001600160a01b0316145b61066d576040805162461bcd60e51b815260206004820152601e60248201527f5472616e7366657273207265737472696374656420746f20657363726f770000604482015290519081900360640190fd5b604080517f172a93fb0000000000000000000000000000000000000000000000000000000081526001600160a01b03888116600483015286811660248301526044820185905291519183169163172a93fb9160648082019260009290919082900301818387803b1580156106e057600080fd5b505af11580156106f4573d6000803e3d6000fd5b50505050505050505050565b6000546001600160a01b031681565b6000336001600160a01b031663975057e76040518163ffffffff1660e01b815260040160206040518083038186803b15801561074a57600080fd5b505afa15801561075e573d6000803e3d6000fd5b505050506040513d602081101561077457600080fd5b5051905090565b6000806107866108fb565b6001600160a01b03166331aaa74a846040518263ffffffff1660e01b815260040180826001600160a01b03166001600160a01b0316815260200191505060206040518083038186803b1580156107db57600080fd5b505afa1580156107ef573d6000803e3d6000fd5b505050506040513d602081101561080557600080fd5b5051604080517f457363726f775265636f726465722e62726f6b657200000000000000000000006020828101919091526bffffffffffffffffffffffff19606085901b16603583015282516029818403018152604983018085528151918301919091207f21f8a72100000000000000000000000000000000000000000000000000000000909152604d8301819052925193945091926001600160a01b038816926321f8a72192606d8082019391829003018186803b1580156108c657600080fd5b505afa1580156108da573d6000803e3d6000fd5b505050506040513d60208110156108f057600080fd5b505195945050505050565b6000336001600160a01b0316637b1039996040518163ffffffff1660e01b815260040160206040518083038186803b15801561074a57600080fdfe4e6577204f776e65722063616e6e6f74206265207468652063757272656e74206f776e6572a265627a7a723158209aa43a9d2164802185de1678581f315cad65140a6bd526fef3d57a73ad2f342064736f6c634300050b0032`

// DeployRestrictToEscrow deploys a new Ethereum contract, binding an instance of RestrictToEscrow to it.
func DeployRestrictToEscrow(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *RestrictToEscrow, error) {
	parsed, err := abi.JSON(strings.NewReader(RestrictToEscrowABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(RestrictToEscrowBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &RestrictToEscrow{RestrictToEscrowCaller: RestrictToEscrowCaller{contract: contract}, RestrictToEscrowTransactor: RestrictToEscrowTransactor{contract: contract}, RestrictToEscrowFilterer: RestrictToEscrowFilterer{contract: contract}}, nil
}

// RestrictToEscrow is an auto generated Go binding around an Ethereum contract.
type RestrictToEscrow struct {
	RestrictToEscrowCaller     // Read-only binding to the contract
	RestrictToEscrowTransactor // Write-only binding to the contract
	RestrictToEscrowFilterer   // Log filterer for contract events
}

// RestrictToEscrowCaller is an auto generated read-only Go binding around an Ethereum contract.
type RestrictToEscrowCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RestrictToEscrowTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RestrictToEscrowTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RestrictToEscrowFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RestrictToEscrowFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RestrictToEscrowSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RestrictToEscrowSession struct {
	Contract     *RestrictToEscrow // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RestrictToEscrowCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RestrictToEscrowCallerSession struct {
	Contract *RestrictToEscrowCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// RestrictToEscrowTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RestrictToEscrowTransactorSession struct {
	Contract     *RestrictToEscrowTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// RestrictToEscrowRaw is an auto generated low-level Go binding around an Ethereum contract.
type RestrictToEscrowRaw struct {
	Contract *RestrictToEscrow // Generic contract binding to access the raw methods on
}

// RestrictToEscrowCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RestrictToEscrowCallerRaw struct {
	Contract *RestrictToEscrowCaller // Generic read-only contract binding to access the raw methods on
}

// RestrictToEscrowTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RestrictToEscrowTransactorRaw struct {
	Contract *RestrictToEscrowTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRestrictToEscrow creates a new instance of RestrictToEscrow, bound to a specific deployed contract.
func NewRestrictToEscrow(address common.Address, backend bind.ContractBackend) (*RestrictToEscrow, error) {
	contract, err := bindRestrictToEscrow(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &RestrictToEscrow{RestrictToEscrowCaller: RestrictToEscrowCaller{contract: contract}, RestrictToEscrowTransactor: RestrictToEscrowTransactor{contract: contract}, RestrictToEscrowFilterer: RestrictToEscrowFilterer{contract: contract}}, nil
}

// NewRestrictToEscrowCaller creates a new read-only instance of RestrictToEscrow, bound to a specific deployed contract.
func NewRestrictToEscrowCaller(address common.Address, caller bind.ContractCaller) (*RestrictToEscrowCaller, error) {
	contract, err := bindRestrictToEscrow(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RestrictToEscrowCaller{contract: contract}, nil
}

// NewRestrictToEscrowTransactor creates a new write-only instance of RestrictToEscrow, bound to a specific deployed contract.
func NewRestrictToEscrowTransactor(address common.Address, transactor bind.ContractTransactor) (*RestrictToEscrowTransactor, error) {
	contract, err := bindRestrictToEscrow(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RestrictToEscrowTransactor{contract: contract}, nil
}

// NewRestrictToEscrowFilterer creates a new log filterer instance of RestrictToEscrow, bound to a specific deployed contract.
func NewRestrictToEscrowFilterer(address common.Address, filterer bind.ContractFilterer) (*RestrictToEscrowFilterer, error) {
	contract, err := bindRestrictToEscrow(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RestrictToEscrowFilterer{contract: contract}, nil
}

// bindRestrictToEscrow binds a generic wrapper to an already deployed contract.
func bindRestrictToEscrow(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(RestrictToEscrowABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RestrictToEscrow *RestrictToEscrowRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _RestrictToEscrow.Contract.RestrictToEscrowCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RestrictToEscrow *RestrictToEscrowRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RestrictToEscrow.Contract.RestrictToEscrowTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RestrictToEscrow *RestrictToEscrowRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RestrictToEscrow.Contract.RestrictToEscrowTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RestrictToEscrow *RestrictToEscrowCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _RestrictToEscrow.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RestrictToEscrow *RestrictToEscrowTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RestrictToEscrow.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RestrictToEscrow *RestrictToEscrowTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RestrictToEscrow.Contract.contract.Transact(opts, method, params...)
}

// ZEROADDRESS is a free data retrieval call binding the contract method 0x538ba4f9.
//
// Solidity: function ZERO_ADDRESS() constant returns(address)
func (_RestrictToEscrow *RestrictToEscrowCaller) ZEROADDRESS(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _RestrictToEscrow.contract.Call(opts, out, "ZERO_ADDRESS")
	return *ret0, err
}

// ZEROADDRESS is a free data retrieval call binding the contract method 0x538ba4f9.
//
// Solidity: function ZERO_ADDRESS() constant returns(address)
func (_RestrictToEscrow *RestrictToEscrowSession) ZEROADDRESS() (common.Address, error) {
	return _RestrictToEscrow.Contract.ZEROADDRESS(&_RestrictToEscrow.CallOpts)
}

// ZEROADDRESS is a free data retrieval call binding the contract method 0x538ba4f9.
//
// Solidity: function ZERO_ADDRESS() constant returns(address)
func (_RestrictToEscrow *RestrictToEscrowCallerSession) ZEROADDRESS() (common.Address, error) {
	return _RestrictToEscrow.Contract.ZEROADDRESS(&_RestrictToEscrow.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_RestrictToEscrow *RestrictToEscrowCaller) Name(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _RestrictToEscrow.contract.Call(opts, out, "name")
	return *ret0, err
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_RestrictToEscrow *RestrictToEscrowSession) Name() (string, error) {
	return _RestrictToEscrow.Contract.Name(&_RestrictToEscrow.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_RestrictToEscrow *RestrictToEscrowCallerSession) Name() (string, error) {
	return _RestrictToEscrow.Contract.Name(&_RestrictToEscrow.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_RestrictToEscrow *RestrictToEscrowCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _RestrictToEscrow.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_RestrictToEscrow *RestrictToEscrowSession) Owner() (common.Address, error) {
	return _RestrictToEscrow.Contract.Owner(&_RestrictToEscrow.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_RestrictToEscrow *RestrictToEscrowCallerSession) Owner() (common.Address, error) {
	return _RestrictToEscrow.Contract.Owner(&_RestrictToEscrow.CallOpts)
}

// Test is a free data retrieval call binding the contract method 0x0b6dedfd.
//
// Solidity: function test(compliance address, token address, initiator address, from address, to address, tokens uint256) constant returns(s string)
func (_RestrictToEscrow *RestrictToEscrowCaller) Test(opts *bind.CallOpts, compliance common.Address, token common.Address, initiator common.Address, from common.Address, to common.Address, tokens *big.Int) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _RestrictToEscrow.contract.Call(opts, out, "test", compliance, token, initiator, from, to, tokens)
	return *ret0, err
}

// Test is a free data retrieval call binding the contract method 0x0b6dedfd.
//
// Solidity: function test(compliance address, token address, initiator address, from address, to address, tokens uint256) constant returns(s string)
func (_RestrictToEscrow *RestrictToEscrowSession) Test(compliance common.Address, token common.Address, initiator common.Address, from common.Address, to common.Address, tokens *big.Int) (string, error) {
	return _RestrictToEscrow.Contract.Test(&_RestrictToEscrow.CallOpts, compliance, token, initiator, from, to, tokens)
}

// Test is a free data retrieval call binding the contract method 0x0b6dedfd.
//
// Solidity: function test(compliance address, token address, initiator address, from address, to address, tokens uint256) constant returns(s string)
func (_RestrictToEscrow *RestrictToEscrowCallerSession) Test(compliance common.Address, token common.Address, initiator common.Address, from common.Address, to common.Address, tokens *big.Int) (string, error) {
	return _RestrictToEscrow.Contract.Test(&_RestrictToEscrow.CallOpts, compliance, token, initiator, from, to, tokens)
}

// Check is a paid mutator transaction binding the contract method 0x803fcd43.
//
// Solidity: function check(token address, initiator address, from address, to address, tokens uint256) returns()
func (_RestrictToEscrow *RestrictToEscrowTransactor) Check(opts *bind.TransactOpts, token common.Address, initiator common.Address, from common.Address, to common.Address, tokens *big.Int) (*types.Transaction, error) {
	return _RestrictToEscrow.contract.Transact(opts, "check", token, initiator, from, to, tokens)
}

// Check is a paid mutator transaction binding the contract method 0x803fcd43.
//
// Solidity: function check(token address, initiator address, from address, to address, tokens uint256) returns()
func (_RestrictToEscrow *RestrictToEscrowSession) Check(token common.Address, initiator common.Address, from common.Address, to common.Address, tokens *big.Int) (*types.Transaction, error) {
	return _RestrictToEscrow.Contract.Check(&_RestrictToEscrow.TransactOpts, token, initiator, from, to, tokens)
}

// Check is a paid mutator transaction binding the contract method 0x803fcd43.
//
// Solidity: function check(token address, initiator address, from address, to address, tokens uint256) returns()
func (_RestrictToEscrow *RestrictToEscrowTransactorSession) Check(token common.Address, initiator common.Address, from common.Address, to common.Address, tokens *big.Int) (*types.Transaction, error) {
	return _RestrictToEscrow.Contract.Check(&_RestrictToEscrow.TransactOpts, token, initiator, from, to, tokens)
}

// Kill is a paid mutator transaction binding the contract method 0x41c0e1b5.
//
// Solidity: function kill() returns()
func (_RestrictToEscrow *RestrictToEscrowTransactor) Kill(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RestrictToEscrow.contract.Transact(opts, "kill")
}

// Kill is a paid mutator transaction binding the contract method 0x41c0e1b5.
//
// Solidity: function kill() returns()
func (_RestrictToEscrow *RestrictToEscrowSession) Kill() (*types.Transaction, error) {
	return _RestrictToEscrow.Contract.Kill(&_RestrictToEscrow.TransactOpts)
}

// Kill is a paid mutator transaction binding the contract method 0x41c0e1b5.
//
// Solidity: function kill() returns()
func (_RestrictToEscrow *RestrictToEscrowTransactorSession) Kill() (*types.Transaction, error) {
	return _RestrictToEscrow.Contract.Kill(&_RestrictToEscrow.TransactOpts)
}

// TransferOwner is a paid mutator transaction binding the contract method 0x4fb2e45d.
//
// Solidity: function transferOwner(newOwner address) returns()
func (_RestrictToEscrow *RestrictToEscrowTransactor) TransferOwner(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _RestrictToEscrow.contract.Transact(opts, "transferOwner", newOwner)
}

// TransferOwner is a paid mutator transaction binding the contract method 0x4fb2e45d.
//
// Solidity: function transferOwner(newOwner address) returns()
func (_RestrictToEscrow *RestrictToEscrowSession) TransferOwner(newOwner common.Address) (*types.Transaction, error) {
	return _RestrictToEscrow.Contract.TransferOwner(&_RestrictToEscrow.TransactOpts, newOwner)
}

// TransferOwner is a paid mutator transaction binding the contract method 0x4fb2e45d.
//
// Solidity: function transferOwner(newOwner address) returns()
func (_RestrictToEscrow *RestrictToEscrowTransactorSession) TransferOwner(newOwner common.Address) (*types.Transaction, error) {
	return _RestrictToEscrow.Contract.TransferOwner(&_RestrictToEscrow.TransactOpts, newOwner)
}

// RestrictToEscrowOwnerTransferredIterator is returned from FilterOwnerTransferred and is used to iterate over the raw logs and unpacked data for OwnerTransferred events raised by the RestrictToEscrow contract.
type RestrictToEscrowOwnerTransferredIterator struct {
	Event *RestrictToEscrowOwnerTransferred // Event containing the contract specifics and raw log

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
func (it *RestrictToEscrowOwnerTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RestrictToEscrowOwnerTransferred)
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
		it.Event = new(RestrictToEscrowOwnerTransferred)
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
func (it *RestrictToEscrowOwnerTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RestrictToEscrowOwnerTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RestrictToEscrowOwnerTransferred represents a OwnerTransferred event raised by the RestrictToEscrow contract.
type RestrictToEscrowOwnerTransferred struct {
	OldOwner common.Address
	NewOwner common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOwnerTransferred is a free log retrieval operation binding the contract event 0x8934ce4adea8d9ce0d714d2c22b86790e41b7731c84b926fbbdc1d40ff6533c9.
//
// Solidity: e OwnerTransferred(oldOwner indexed address, newOwner indexed address)
func (_RestrictToEscrow *RestrictToEscrowFilterer) FilterOwnerTransferred(opts *bind.FilterOpts, oldOwner []common.Address, newOwner []common.Address) (*RestrictToEscrowOwnerTransferredIterator, error) {

	var oldOwnerRule []interface{}
	for _, oldOwnerItem := range oldOwner {
		oldOwnerRule = append(oldOwnerRule, oldOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _RestrictToEscrow.contract.FilterLogs(opts, "OwnerTransferred", oldOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &RestrictToEscrowOwnerTransferredIterator{contract: _RestrictToEscrow.contract, event: "OwnerTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnerTransferred is a free log subscription operation binding the contract event 0x8934ce4adea8d9ce0d714d2c22b86790e41b7731c84b926fbbdc1d40ff6533c9.
//
// Solidity: e OwnerTransferred(oldOwner indexed address, newOwner indexed address)
func (_RestrictToEscrow *RestrictToEscrowFilterer) WatchOwnerTransferred(opts *bind.WatchOpts, sink chan<- *RestrictToEscrowOwnerTransferred, oldOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var oldOwnerRule []interface{}
	for _, oldOwnerItem := range oldOwner {
		oldOwnerRule = append(oldOwnerRule, oldOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _RestrictToEscrow.contract.WatchLogs(opts, "OwnerTransferred", oldOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RestrictToEscrowOwnerTransferred)
				if err := _RestrictToEscrow.contract.UnpackLog(event, "OwnerTransferred", log); err != nil {
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
