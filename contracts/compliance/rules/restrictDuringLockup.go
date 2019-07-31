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

// RestrictDuringLockupABI is the input ABI used to generate the binding from.
const RestrictDuringLockupABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"kill\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwner\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ZERO_ADDRESS\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"token\",\"type\":\"address\"},{\"name\":\"initiator\",\"type\":\"address\"},{\"name\":\"from\",\"type\":\"address\"},{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"toKind\",\"type\":\"uint8\"},{\"name\":\"tokens\",\"type\":\"uint256\"}],\"name\":\"check\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"compliance\",\"type\":\"address\"},{\"name\":\"token\",\"type\":\"address\"},{\"name\":\"initiator\",\"type\":\"address\"},{\"name\":\"from\",\"type\":\"address\"},{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"toKind\",\"type\":\"uint8\"},{\"name\":\"tokens\",\"type\":\"uint256\"}],\"name\":\"test\",\"outputs\":[{\"name\":\"s\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"oldOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnerTransferred\",\"type\":\"event\"}]"

// RestrictDuringLockupBin is the compiled bytecode used for deploying new contracts.
const RestrictDuringLockupBin = `600180546001600160a01b031916905560c0604052601660808190527f526573747269637420447572696e67204c6f636b75700000000000000000000060a090815261004e9160029190610066565b50600080546001600160a01b03191633179055610101565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106100a757805160ff19168380011785556100d4565b828001600101855582156100d4579182015b828111156100d45782518255916020019190600101906100b9565b506100e09291506100e4565b5090565b6100fe91905b808211156100e057600081556001016100ea565b90565b610af3806101106000396000f3fe608060405234801561001057600080fd5b506004361061007d5760003560e01c8063538ba4f91161005b578063538ba4f91461012f5780635a47e1c7146101535780638da5cb5b146101a2578063a9189562146101aa5761007d565b806306fdde031461008257806341c0e1b5146100ff5780634fb2e45d14610109575b600080fd5b61008a6101ff565b6040805160208082528351818301528351919283929083019185019080838360005b838110156100c45781810151838201526020016100ac565b50505050905090810190601f1680156100f15780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b61010761028a565b005b6101076004803603602081101561011f57600080fd5b50356001600160a01b0316610312565b61013761049c565b604080516001600160a01b039092168252519081900360200190f35b610107600480360360c081101561016957600080fd5b506001600160a01b0381358116916020810135821691604082013581169160608101359091169060ff6080820135169060a001356104ab565b610137610721565b61008a600480360360e08110156101c057600080fd5b506001600160a01b0381358116916020810135821691604082013581169160608101358216916080820135169060ff60a0820135169060c00135610730565b6002805460408051602060018416156101000260001901909316849004601f810184900484028201840190925281815292918301828280156102825780601f1061025757610100808354040283529160200191610282565b820191906000526020600020905b81548152906001019060200180831161026557829003601f168201915b505050505081565b6000546001600160a01b03163314806102b357506001546000546001600160a01b039081169116145b610304576040805162461bcd60e51b815260206004820152601960248201527f4f776e6572206163636f756e7420697320726571756972656400000000000000604482015290519081900360640190fd5b6000546001600160a01b0316ff5b6000546001600160a01b031633148061033b57506001546000546001600160a01b039081169116145b61038c576040805162461bcd60e51b815260206004820152601960248201527f4f776e6572206163636f756e7420697320726571756972656400000000000000604482015290519081900360640190fd5b6000546001600160a01b03828116911614156103d95760405162461bcd60e51b8152600401808060200182810382526025815260200180610a796025913960400191505060405180910390fd5b6001600160a01b038116610434576040805162461bcd60e51b815260206004820181905260248201527f4e6577204f776e65722063616e6e6f74206265207a65726f2061646472657373604482015290519081900360640190fd5b600080546001600160a01b038381167fffffffffffffffffffffffff0000000000000000000000000000000000000000831681178455604051919092169283917f8934ce4adea8d9ce0d714d2c22b86790e41b7731c84b926fbbdc1d40ff6533c99190a35050565b6001546001600160a01b031681565b60006040518060400160405280601e81526020017f5265737472696374447572696e674c6f636b75702e6475726174696f6e730000815250876001600160a01b03166395d89b416040518163ffffffff1660e01b815260040160006040518083038186803b15801561051c57600080fd5b505afa158015610530573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052602081101561055957600080fd5b81019080805164010000000081111561057157600080fd5b8201602081018481111561058457600080fd5b815164010000000081118282018710171561059e57600080fd5b50509291905050506040516020018083805190602001908083835b602083106105d85780518252601f1990920191602091820191016105b9565b51815160209384036101000a600019018019909216911617905285519190930192850191508083835b602083106106205780518252601f199092019160209182019101610601565b6001836020036101000a038019825116818451168082178552505050505050905001925050506040516020818303038152906040528051906020012090506000610668610a0c565b6001600160a01b03166333598b00836040518263ffffffff1660e01b81526004018082815260200191505060206040518083038186803b1580156106ab57600080fd5b505afa1580156106bf573d6000803e3d6000fd5b505050506040513d60208110156106d557600080fd5b505190504281106107175760405162461bcd60e51b8152600401808060200182810382526021815260200180610a9e6021913960400191505060405180910390fd5b5050505050505050565b6000546001600160a01b031681565b606060006040518060400160405280601e81526020017f5265737472696374447572696e674c6f636b75702e6475726174696f6e730000815250886001600160a01b03166395d89b416040518163ffffffff1660e01b815260040160006040518083038186803b1580156107a357600080fd5b505afa1580156107b7573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f1916820160405260208110156107e057600080fd5b8101908080516401000000008111156107f857600080fd5b8201602081018481111561080b57600080fd5b815164010000000081118282018710171561082557600080fd5b50509291905050506040516020018083805190602001908083835b6020831061085f5780518252601f199092019160209182019101610840565b51815160209384036101000a600019018019909216911617905285519190930192850191508083835b602083106108a75780518252601f199092019160209182019101610888565b6001836020036101000a038019825116818451168082178552505050505050905001925050506040516020818303038152906040528051906020012090506000896001600160a01b031663975057e76040518163ffffffff1660e01b815260040160206040518083038186803b15801561092057600080fd5b505afa158015610934573d6000803e3d6000fd5b505050506040513d602081101561094a57600080fd5b5051604080517f33598b000000000000000000000000000000000000000000000000000000000081526004810185905290516001600160a01b03909216916333598b0091602480820192602092909190829003018186803b1580156109ae57600080fd5b505afa1580156109c2573d6000803e3d6000fd5b505050506040513d60208110156109d857600080fd5b505190504281106109ff57604051806060016040528060218152602001610a9e6021913992505b5050979650505050505050565b6000336001600160a01b031663975057e76040518163ffffffff1660e01b815260040160206040518083038186803b158015610a4757600080fd5b505afa158015610a5b573d6000803e3d6000fd5b505050506040513d6020811015610a7157600080fd5b505190509056fe4e6577204f776e65722063616e6e6f74206265207468652063757272656e74206f776e6572486f6c6465722063757272656e746c7920696e206c6f636b757020706572696f64a265627a7a72305820deadcf44d50a19d13732c1004181fe37857b5cf4029d259395768f20ff7a7ab064736f6c634300050a0032`

// DeployRestrictDuringLockup deploys a new Ethereum contract, binding an instance of RestrictDuringLockup to it.
func DeployRestrictDuringLockup(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *RestrictDuringLockup, error) {
	parsed, err := abi.JSON(strings.NewReader(RestrictDuringLockupABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(RestrictDuringLockupBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &RestrictDuringLockup{RestrictDuringLockupCaller: RestrictDuringLockupCaller{contract: contract}, RestrictDuringLockupTransactor: RestrictDuringLockupTransactor{contract: contract}, RestrictDuringLockupFilterer: RestrictDuringLockupFilterer{contract: contract}}, nil
}

// RestrictDuringLockup is an auto generated Go binding around an Ethereum contract.
type RestrictDuringLockup struct {
	RestrictDuringLockupCaller     // Read-only binding to the contract
	RestrictDuringLockupTransactor // Write-only binding to the contract
	RestrictDuringLockupFilterer   // Log filterer for contract events
}

// RestrictDuringLockupCaller is an auto generated read-only Go binding around an Ethereum contract.
type RestrictDuringLockupCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RestrictDuringLockupTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RestrictDuringLockupTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RestrictDuringLockupFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RestrictDuringLockupFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RestrictDuringLockupSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RestrictDuringLockupSession struct {
	Contract     *RestrictDuringLockup // Generic contract binding to set the session for
	CallOpts     bind.CallOpts         // Call options to use throughout this session
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// RestrictDuringLockupCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RestrictDuringLockupCallerSession struct {
	Contract *RestrictDuringLockupCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts               // Call options to use throughout this session
}

// RestrictDuringLockupTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RestrictDuringLockupTransactorSession struct {
	Contract     *RestrictDuringLockupTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts               // Transaction auth options to use throughout this session
}

// RestrictDuringLockupRaw is an auto generated low-level Go binding around an Ethereum contract.
type RestrictDuringLockupRaw struct {
	Contract *RestrictDuringLockup // Generic contract binding to access the raw methods on
}

// RestrictDuringLockupCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RestrictDuringLockupCallerRaw struct {
	Contract *RestrictDuringLockupCaller // Generic read-only contract binding to access the raw methods on
}

// RestrictDuringLockupTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RestrictDuringLockupTransactorRaw struct {
	Contract *RestrictDuringLockupTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRestrictDuringLockup creates a new instance of RestrictDuringLockup, bound to a specific deployed contract.
func NewRestrictDuringLockup(address common.Address, backend bind.ContractBackend) (*RestrictDuringLockup, error) {
	contract, err := bindRestrictDuringLockup(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &RestrictDuringLockup{RestrictDuringLockupCaller: RestrictDuringLockupCaller{contract: contract}, RestrictDuringLockupTransactor: RestrictDuringLockupTransactor{contract: contract}, RestrictDuringLockupFilterer: RestrictDuringLockupFilterer{contract: contract}}, nil
}

// NewRestrictDuringLockupCaller creates a new read-only instance of RestrictDuringLockup, bound to a specific deployed contract.
func NewRestrictDuringLockupCaller(address common.Address, caller bind.ContractCaller) (*RestrictDuringLockupCaller, error) {
	contract, err := bindRestrictDuringLockup(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RestrictDuringLockupCaller{contract: contract}, nil
}

// NewRestrictDuringLockupTransactor creates a new write-only instance of RestrictDuringLockup, bound to a specific deployed contract.
func NewRestrictDuringLockupTransactor(address common.Address, transactor bind.ContractTransactor) (*RestrictDuringLockupTransactor, error) {
	contract, err := bindRestrictDuringLockup(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RestrictDuringLockupTransactor{contract: contract}, nil
}

// NewRestrictDuringLockupFilterer creates a new log filterer instance of RestrictDuringLockup, bound to a specific deployed contract.
func NewRestrictDuringLockupFilterer(address common.Address, filterer bind.ContractFilterer) (*RestrictDuringLockupFilterer, error) {
	contract, err := bindRestrictDuringLockup(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RestrictDuringLockupFilterer{contract: contract}, nil
}

// bindRestrictDuringLockup binds a generic wrapper to an already deployed contract.
func bindRestrictDuringLockup(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(RestrictDuringLockupABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RestrictDuringLockup *RestrictDuringLockupRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _RestrictDuringLockup.Contract.RestrictDuringLockupCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RestrictDuringLockup *RestrictDuringLockupRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RestrictDuringLockup.Contract.RestrictDuringLockupTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RestrictDuringLockup *RestrictDuringLockupRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RestrictDuringLockup.Contract.RestrictDuringLockupTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RestrictDuringLockup *RestrictDuringLockupCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _RestrictDuringLockup.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RestrictDuringLockup *RestrictDuringLockupTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RestrictDuringLockup.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RestrictDuringLockup *RestrictDuringLockupTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RestrictDuringLockup.Contract.contract.Transact(opts, method, params...)
}

// ZEROADDRESS is a free data retrieval call binding the contract method 0x538ba4f9.
//
// Solidity: function ZERO_ADDRESS() constant returns(address)
func (_RestrictDuringLockup *RestrictDuringLockupCaller) ZEROADDRESS(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _RestrictDuringLockup.contract.Call(opts, out, "ZERO_ADDRESS")
	return *ret0, err
}

// ZEROADDRESS is a free data retrieval call binding the contract method 0x538ba4f9.
//
// Solidity: function ZERO_ADDRESS() constant returns(address)
func (_RestrictDuringLockup *RestrictDuringLockupSession) ZEROADDRESS() (common.Address, error) {
	return _RestrictDuringLockup.Contract.ZEROADDRESS(&_RestrictDuringLockup.CallOpts)
}

// ZEROADDRESS is a free data retrieval call binding the contract method 0x538ba4f9.
//
// Solidity: function ZERO_ADDRESS() constant returns(address)
func (_RestrictDuringLockup *RestrictDuringLockupCallerSession) ZEROADDRESS() (common.Address, error) {
	return _RestrictDuringLockup.Contract.ZEROADDRESS(&_RestrictDuringLockup.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_RestrictDuringLockup *RestrictDuringLockupCaller) Name(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _RestrictDuringLockup.contract.Call(opts, out, "name")
	return *ret0, err
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_RestrictDuringLockup *RestrictDuringLockupSession) Name() (string, error) {
	return _RestrictDuringLockup.Contract.Name(&_RestrictDuringLockup.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_RestrictDuringLockup *RestrictDuringLockupCallerSession) Name() (string, error) {
	return _RestrictDuringLockup.Contract.Name(&_RestrictDuringLockup.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_RestrictDuringLockup *RestrictDuringLockupCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _RestrictDuringLockup.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_RestrictDuringLockup *RestrictDuringLockupSession) Owner() (common.Address, error) {
	return _RestrictDuringLockup.Contract.Owner(&_RestrictDuringLockup.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_RestrictDuringLockup *RestrictDuringLockupCallerSession) Owner() (common.Address, error) {
	return _RestrictDuringLockup.Contract.Owner(&_RestrictDuringLockup.CallOpts)
}

// Test is a free data retrieval call binding the contract method 0xa9189562.
//
// Solidity: function test(compliance address, token address, initiator address, from address, to address, toKind uint8, tokens uint256) constant returns(s string)
func (_RestrictDuringLockup *RestrictDuringLockupCaller) Test(opts *bind.CallOpts, compliance common.Address, token common.Address, initiator common.Address, from common.Address, to common.Address, toKind uint8, tokens *big.Int) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _RestrictDuringLockup.contract.Call(opts, out, "test", compliance, token, initiator, from, to, toKind, tokens)
	return *ret0, err
}

// Test is a free data retrieval call binding the contract method 0xa9189562.
//
// Solidity: function test(compliance address, token address, initiator address, from address, to address, toKind uint8, tokens uint256) constant returns(s string)
func (_RestrictDuringLockup *RestrictDuringLockupSession) Test(compliance common.Address, token common.Address, initiator common.Address, from common.Address, to common.Address, toKind uint8, tokens *big.Int) (string, error) {
	return _RestrictDuringLockup.Contract.Test(&_RestrictDuringLockup.CallOpts, compliance, token, initiator, from, to, toKind, tokens)
}

// Test is a free data retrieval call binding the contract method 0xa9189562.
//
// Solidity: function test(compliance address, token address, initiator address, from address, to address, toKind uint8, tokens uint256) constant returns(s string)
func (_RestrictDuringLockup *RestrictDuringLockupCallerSession) Test(compliance common.Address, token common.Address, initiator common.Address, from common.Address, to common.Address, toKind uint8, tokens *big.Int) (string, error) {
	return _RestrictDuringLockup.Contract.Test(&_RestrictDuringLockup.CallOpts, compliance, token, initiator, from, to, toKind, tokens)
}

// Check is a paid mutator transaction binding the contract method 0x5a47e1c7.
//
// Solidity: function check(token address, initiator address, from address, to address, toKind uint8, tokens uint256) returns()
func (_RestrictDuringLockup *RestrictDuringLockupTransactor) Check(opts *bind.TransactOpts, token common.Address, initiator common.Address, from common.Address, to common.Address, toKind uint8, tokens *big.Int) (*types.Transaction, error) {
	return _RestrictDuringLockup.contract.Transact(opts, "check", token, initiator, from, to, toKind, tokens)
}

// Check is a paid mutator transaction binding the contract method 0x5a47e1c7.
//
// Solidity: function check(token address, initiator address, from address, to address, toKind uint8, tokens uint256) returns()
func (_RestrictDuringLockup *RestrictDuringLockupSession) Check(token common.Address, initiator common.Address, from common.Address, to common.Address, toKind uint8, tokens *big.Int) (*types.Transaction, error) {
	return _RestrictDuringLockup.Contract.Check(&_RestrictDuringLockup.TransactOpts, token, initiator, from, to, toKind, tokens)
}

// Check is a paid mutator transaction binding the contract method 0x5a47e1c7.
//
// Solidity: function check(token address, initiator address, from address, to address, toKind uint8, tokens uint256) returns()
func (_RestrictDuringLockup *RestrictDuringLockupTransactorSession) Check(token common.Address, initiator common.Address, from common.Address, to common.Address, toKind uint8, tokens *big.Int) (*types.Transaction, error) {
	return _RestrictDuringLockup.Contract.Check(&_RestrictDuringLockup.TransactOpts, token, initiator, from, to, toKind, tokens)
}

// Kill is a paid mutator transaction binding the contract method 0x41c0e1b5.
//
// Solidity: function kill() returns()
func (_RestrictDuringLockup *RestrictDuringLockupTransactor) Kill(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RestrictDuringLockup.contract.Transact(opts, "kill")
}

// Kill is a paid mutator transaction binding the contract method 0x41c0e1b5.
//
// Solidity: function kill() returns()
func (_RestrictDuringLockup *RestrictDuringLockupSession) Kill() (*types.Transaction, error) {
	return _RestrictDuringLockup.Contract.Kill(&_RestrictDuringLockup.TransactOpts)
}

// Kill is a paid mutator transaction binding the contract method 0x41c0e1b5.
//
// Solidity: function kill() returns()
func (_RestrictDuringLockup *RestrictDuringLockupTransactorSession) Kill() (*types.Transaction, error) {
	return _RestrictDuringLockup.Contract.Kill(&_RestrictDuringLockup.TransactOpts)
}

// TransferOwner is a paid mutator transaction binding the contract method 0x4fb2e45d.
//
// Solidity: function transferOwner(newOwner address) returns()
func (_RestrictDuringLockup *RestrictDuringLockupTransactor) TransferOwner(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _RestrictDuringLockup.contract.Transact(opts, "transferOwner", newOwner)
}

// TransferOwner is a paid mutator transaction binding the contract method 0x4fb2e45d.
//
// Solidity: function transferOwner(newOwner address) returns()
func (_RestrictDuringLockup *RestrictDuringLockupSession) TransferOwner(newOwner common.Address) (*types.Transaction, error) {
	return _RestrictDuringLockup.Contract.TransferOwner(&_RestrictDuringLockup.TransactOpts, newOwner)
}

// TransferOwner is a paid mutator transaction binding the contract method 0x4fb2e45d.
//
// Solidity: function transferOwner(newOwner address) returns()
func (_RestrictDuringLockup *RestrictDuringLockupTransactorSession) TransferOwner(newOwner common.Address) (*types.Transaction, error) {
	return _RestrictDuringLockup.Contract.TransferOwner(&_RestrictDuringLockup.TransactOpts, newOwner)
}

// RestrictDuringLockupOwnerTransferredIterator is returned from FilterOwnerTransferred and is used to iterate over the raw logs and unpacked data for OwnerTransferred events raised by the RestrictDuringLockup contract.
type RestrictDuringLockupOwnerTransferredIterator struct {
	Event *RestrictDuringLockupOwnerTransferred // Event containing the contract specifics and raw log

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
func (it *RestrictDuringLockupOwnerTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RestrictDuringLockupOwnerTransferred)
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
		it.Event = new(RestrictDuringLockupOwnerTransferred)
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
func (it *RestrictDuringLockupOwnerTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RestrictDuringLockupOwnerTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RestrictDuringLockupOwnerTransferred represents a OwnerTransferred event raised by the RestrictDuringLockup contract.
type RestrictDuringLockupOwnerTransferred struct {
	OldOwner common.Address
	NewOwner common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOwnerTransferred is a free log retrieval operation binding the contract event 0x8934ce4adea8d9ce0d714d2c22b86790e41b7731c84b926fbbdc1d40ff6533c9.
//
// Solidity: e OwnerTransferred(oldOwner indexed address, newOwner indexed address)
func (_RestrictDuringLockup *RestrictDuringLockupFilterer) FilterOwnerTransferred(opts *bind.FilterOpts, oldOwner []common.Address, newOwner []common.Address) (*RestrictDuringLockupOwnerTransferredIterator, error) {

	var oldOwnerRule []interface{}
	for _, oldOwnerItem := range oldOwner {
		oldOwnerRule = append(oldOwnerRule, oldOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _RestrictDuringLockup.contract.FilterLogs(opts, "OwnerTransferred", oldOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &RestrictDuringLockupOwnerTransferredIterator{contract: _RestrictDuringLockup.contract, event: "OwnerTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnerTransferred is a free log subscription operation binding the contract event 0x8934ce4adea8d9ce0d714d2c22b86790e41b7731c84b926fbbdc1d40ff6533c9.
//
// Solidity: e OwnerTransferred(oldOwner indexed address, newOwner indexed address)
func (_RestrictDuringLockup *RestrictDuringLockupFilterer) WatchOwnerTransferred(opts *bind.WatchOpts, sink chan<- *RestrictDuringLockupOwnerTransferred, oldOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var oldOwnerRule []interface{}
	for _, oldOwnerItem := range oldOwner {
		oldOwnerRule = append(oldOwnerRule, oldOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _RestrictDuringLockup.contract.WatchLogs(opts, "OwnerTransferred", oldOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RestrictDuringLockupOwnerTransferred)
				if err := _RestrictDuringLockup.contract.UnpackLog(event, "OwnerTransferred", log); err != nil {
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
