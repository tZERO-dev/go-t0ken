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

// RestrictFromAffiliateABI is the input ABI used to generate the binding from.
const RestrictFromAffiliateABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"contractICompliance\",\"name\":\"compliance\",\"type\":\"address\"},{\"internalType\":\"contractIT0ken\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"initiator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokens\",\"type\":\"uint256\"}],\"name\":\"test\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"s\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"kill\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwner\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ZERO_ADDRESS\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"contractIT0ken\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"initiator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokens\",\"type\":\"uint256\"}],\"name\":\"check\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnerTransferred\",\"type\":\"event\"}]"

// RestrictFromAffiliateBin is the compiled bytecode used for deploying new contracts.
const RestrictFromAffiliateBin = `600180546001600160a01b031916905560c0604052601760808190527f52657374726963742046726f6d20416666696c6961746500000000000000000060a090815261004e9160029190610066565b50600080546001600160a01b03191633179055610101565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106100a757805160ff19168380011785556100d4565b828001600101855582156100d4579182015b828111156100d45782518255916020019190600101906100b9565b506100e09291506100e4565b5090565b6100fe91905b808211156100e057600081556001016100ea565b90565b610b11806101106000396000f3fe608060405234801561001057600080fd5b506004361061007d5760003560e01c80634fb2e45d1161005b5780634fb2e45d14610155578063538ba4f91461017b578063803fcd431461019f5780638da5cb5b146101e55761007d565b806306fdde03146100825780630b6dedfd146100ff57806341c0e1b51461014b575b600080fd5b61008a6101ed565b6040805160208082528351818301528351919283929083019185019080838360005b838110156100c45781810151838201526020016100ac565b50505050905090810190601f1680156100f15780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b61008a600480360360c081101561011557600080fd5b506001600160a01b0381358116916020810135821691604082013581169160608101358216916080820135169060a00135610278565b610153610558565b005b6101536004803603602081101561016b57600080fd5b50356001600160a01b03166105e0565b61018361076a565b604080516001600160a01b039092168252519081900360200190f35b610153600480360360a08110156101b557600080fd5b506001600160a01b0381358116916020810135821691604082013581169160608101359091169060800135610779565b6101836109f3565b6002805460408051602060018416156101000260001901909316849004601f810184900484028201840190925281815292918301828280156102705780601f1061024557610100808354040283529160200191610270565b820191906000526020600020905b81548152906001019060200180831161025357829003601f168201915b505050505081565b60606000866001600160a01b03166395d89b416040518163ffffffff1660e01b815260040160006040518083038186803b1580156102b557600080fd5b505afa1580156102c9573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f1916820160405260208110156102f257600080fd5b810190808051604051939291908464010000000082111561031257600080fd5b90830190602082018581111561032757600080fd5b825164010000000081118282018810171561034157600080fd5b82525081516020918201929091019080838360005b8381101561036e578181015183820152602001610356565b50505050905090810190601f16801561039b5780820380516001836020036101000a031916815260200191505b50604052505050856040516020018080610a946021913960210183805190602001908083835b602083106103e05780518252601f1990920191602091820191016103c1565b6001836020036101000a038019825116818451168082178552505050505050905001826001600160a01b03166001600160a01b031660601b815260140192505050604051602081830303815290604052805190602001209050876001600160a01b031663975057e76040518163ffffffff1660e01b815260040160206040518083038186803b15801561047257600080fd5b505afa158015610486573d6000803e3d6000fd5b505050506040513d602081101561049c57600080fd5b5051604080517f7ae1cfca0000000000000000000000000000000000000000000000000000000081526004810184905290516001600160a01b0390921691637ae1cfca91602480820192602092909190829003018186803b15801561050057600080fd5b505afa158015610514573d6000803e3d6000fd5b505050506040513d602081101561052a57600080fd5b50511561054d57604051806060016040528060288152602001610ab56028913991505b509695505050505050565b6000546001600160a01b031633148061058157506001546000546001600160a01b039081169116145b6105d2576040805162461bcd60e51b815260206004820152601960248201527f4f776e6572206163636f756e7420697320726571756972656400000000000000604482015290519081900360640190fd5b6000546001600160a01b0316ff5b6000546001600160a01b031633148061060957506001546000546001600160a01b039081169116145b61065a576040805162461bcd60e51b815260206004820152601960248201527f4f776e6572206163636f756e7420697320726571756972656400000000000000604482015290519081900360640190fd5b6000546001600160a01b03828116911614156106a75760405162461bcd60e51b8152600401808060200182810382526025815260200180610a6f6025913960400191505060405180910390fd5b6001600160a01b038116610702576040805162461bcd60e51b815260206004820181905260248201527f4e6577204f776e65722063616e6e6f74206265207a65726f2061646472657373604482015290519081900360640190fd5b600080546001600160a01b038381167fffffffffffffffffffffffff0000000000000000000000000000000000000000831681178455604051919092169283917f8934ce4adea8d9ce0d714d2c22b86790e41b7731c84b926fbbdc1d40ff6533c99190a35050565b6001546001600160a01b031681565b6000856001600160a01b03166395d89b416040518163ffffffff1660e01b815260040160006040518083038186803b1580156107b457600080fd5b505afa1580156107c8573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f1916820160405260208110156107f157600080fd5b810190808051604051939291908464010000000082111561081157600080fd5b90830190602082018581111561082657600080fd5b825164010000000081118282018810171561084057600080fd5b82525081516020918201929091019080838360005b8381101561086d578181015183820152602001610855565b50505050905090810190601f16801561089a5780820380516001836020036101000a031916815260200191505b50604052505050846040516020018080610a946021913960210183805190602001908083835b602083106108df5780518252601f1990920191602091820191016108c0565b6001836020036101000a038019825116818451168082178552505050505050905001826001600160a01b03166001600160a01b031660601b815260140192505050604051602081830303815290604052805190602001209050610940610a02565b6001600160a01b0316637ae1cfca826040518263ffffffff1660e01b81526004018082815260200191505060206040518083038186803b15801561098357600080fd5b505afa158015610997573d6000803e3d6000fd5b505050506040513d60208110156109ad57600080fd5b5051156109eb5760405162461bcd60e51b8152600401808060200182810382526028815260200180610ab56028913960400191505060405180910390fd5b505050505050565b6000546001600160a01b031681565b6000336001600160a01b031663975057e76040518163ffffffff1660e01b815260040160206040518083038186803b158015610a3d57600080fd5b505afa158015610a51573d6000803e3d6000fd5b505050506040513d6020811015610a6757600080fd5b505190509056fe4e6577204f776e65722063616e6e6f74206265207468652063757272656e74206f776e6572526573747269637446726f6d416666696c696174652e6973416666696c696174655472616e73666572732066726f6d20616666696c6961746573206172652072657374726963746564a265627a7a7231582032c011709829290b5db040a44162a2284989051811263561979b24949cd98d8e64736f6c634300050b0032`

// DeployRestrictFromAffiliate deploys a new Ethereum contract, binding an instance of RestrictFromAffiliate to it.
func DeployRestrictFromAffiliate(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *RestrictFromAffiliate, error) {
	parsed, err := abi.JSON(strings.NewReader(RestrictFromAffiliateABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(RestrictFromAffiliateBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &RestrictFromAffiliate{RestrictFromAffiliateCaller: RestrictFromAffiliateCaller{contract: contract}, RestrictFromAffiliateTransactor: RestrictFromAffiliateTransactor{contract: contract}, RestrictFromAffiliateFilterer: RestrictFromAffiliateFilterer{contract: contract}}, nil
}

// RestrictFromAffiliate is an auto generated Go binding around an Ethereum contract.
type RestrictFromAffiliate struct {
	RestrictFromAffiliateCaller     // Read-only binding to the contract
	RestrictFromAffiliateTransactor // Write-only binding to the contract
	RestrictFromAffiliateFilterer   // Log filterer for contract events
}

// RestrictFromAffiliateCaller is an auto generated read-only Go binding around an Ethereum contract.
type RestrictFromAffiliateCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RestrictFromAffiliateTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RestrictFromAffiliateTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RestrictFromAffiliateFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RestrictFromAffiliateFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RestrictFromAffiliateSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RestrictFromAffiliateSession struct {
	Contract     *RestrictFromAffiliate // Generic contract binding to set the session for
	CallOpts     bind.CallOpts          // Call options to use throughout this session
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// RestrictFromAffiliateCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RestrictFromAffiliateCallerSession struct {
	Contract *RestrictFromAffiliateCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                // Call options to use throughout this session
}

// RestrictFromAffiliateTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RestrictFromAffiliateTransactorSession struct {
	Contract     *RestrictFromAffiliateTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                // Transaction auth options to use throughout this session
}

// RestrictFromAffiliateRaw is an auto generated low-level Go binding around an Ethereum contract.
type RestrictFromAffiliateRaw struct {
	Contract *RestrictFromAffiliate // Generic contract binding to access the raw methods on
}

// RestrictFromAffiliateCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RestrictFromAffiliateCallerRaw struct {
	Contract *RestrictFromAffiliateCaller // Generic read-only contract binding to access the raw methods on
}

// RestrictFromAffiliateTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RestrictFromAffiliateTransactorRaw struct {
	Contract *RestrictFromAffiliateTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRestrictFromAffiliate creates a new instance of RestrictFromAffiliate, bound to a specific deployed contract.
func NewRestrictFromAffiliate(address common.Address, backend bind.ContractBackend) (*RestrictFromAffiliate, error) {
	contract, err := bindRestrictFromAffiliate(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &RestrictFromAffiliate{RestrictFromAffiliateCaller: RestrictFromAffiliateCaller{contract: contract}, RestrictFromAffiliateTransactor: RestrictFromAffiliateTransactor{contract: contract}, RestrictFromAffiliateFilterer: RestrictFromAffiliateFilterer{contract: contract}}, nil
}

// NewRestrictFromAffiliateCaller creates a new read-only instance of RestrictFromAffiliate, bound to a specific deployed contract.
func NewRestrictFromAffiliateCaller(address common.Address, caller bind.ContractCaller) (*RestrictFromAffiliateCaller, error) {
	contract, err := bindRestrictFromAffiliate(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RestrictFromAffiliateCaller{contract: contract}, nil
}

// NewRestrictFromAffiliateTransactor creates a new write-only instance of RestrictFromAffiliate, bound to a specific deployed contract.
func NewRestrictFromAffiliateTransactor(address common.Address, transactor bind.ContractTransactor) (*RestrictFromAffiliateTransactor, error) {
	contract, err := bindRestrictFromAffiliate(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RestrictFromAffiliateTransactor{contract: contract}, nil
}

// NewRestrictFromAffiliateFilterer creates a new log filterer instance of RestrictFromAffiliate, bound to a specific deployed contract.
func NewRestrictFromAffiliateFilterer(address common.Address, filterer bind.ContractFilterer) (*RestrictFromAffiliateFilterer, error) {
	contract, err := bindRestrictFromAffiliate(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RestrictFromAffiliateFilterer{contract: contract}, nil
}

// bindRestrictFromAffiliate binds a generic wrapper to an already deployed contract.
func bindRestrictFromAffiliate(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(RestrictFromAffiliateABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RestrictFromAffiliate *RestrictFromAffiliateRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _RestrictFromAffiliate.Contract.RestrictFromAffiliateCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RestrictFromAffiliate *RestrictFromAffiliateRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RestrictFromAffiliate.Contract.RestrictFromAffiliateTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RestrictFromAffiliate *RestrictFromAffiliateRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RestrictFromAffiliate.Contract.RestrictFromAffiliateTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RestrictFromAffiliate *RestrictFromAffiliateCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _RestrictFromAffiliate.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RestrictFromAffiliate *RestrictFromAffiliateTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RestrictFromAffiliate.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RestrictFromAffiliate *RestrictFromAffiliateTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RestrictFromAffiliate.Contract.contract.Transact(opts, method, params...)
}

// ZEROADDRESS is a free data retrieval call binding the contract method 0x538ba4f9.
//
// Solidity: function ZERO_ADDRESS() constant returns(address)
func (_RestrictFromAffiliate *RestrictFromAffiliateCaller) ZEROADDRESS(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _RestrictFromAffiliate.contract.Call(opts, out, "ZERO_ADDRESS")
	return *ret0, err
}

// ZEROADDRESS is a free data retrieval call binding the contract method 0x538ba4f9.
//
// Solidity: function ZERO_ADDRESS() constant returns(address)
func (_RestrictFromAffiliate *RestrictFromAffiliateSession) ZEROADDRESS() (common.Address, error) {
	return _RestrictFromAffiliate.Contract.ZEROADDRESS(&_RestrictFromAffiliate.CallOpts)
}

// ZEROADDRESS is a free data retrieval call binding the contract method 0x538ba4f9.
//
// Solidity: function ZERO_ADDRESS() constant returns(address)
func (_RestrictFromAffiliate *RestrictFromAffiliateCallerSession) ZEROADDRESS() (common.Address, error) {
	return _RestrictFromAffiliate.Contract.ZEROADDRESS(&_RestrictFromAffiliate.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_RestrictFromAffiliate *RestrictFromAffiliateCaller) Name(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _RestrictFromAffiliate.contract.Call(opts, out, "name")
	return *ret0, err
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_RestrictFromAffiliate *RestrictFromAffiliateSession) Name() (string, error) {
	return _RestrictFromAffiliate.Contract.Name(&_RestrictFromAffiliate.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_RestrictFromAffiliate *RestrictFromAffiliateCallerSession) Name() (string, error) {
	return _RestrictFromAffiliate.Contract.Name(&_RestrictFromAffiliate.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_RestrictFromAffiliate *RestrictFromAffiliateCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _RestrictFromAffiliate.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_RestrictFromAffiliate *RestrictFromAffiliateSession) Owner() (common.Address, error) {
	return _RestrictFromAffiliate.Contract.Owner(&_RestrictFromAffiliate.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_RestrictFromAffiliate *RestrictFromAffiliateCallerSession) Owner() (common.Address, error) {
	return _RestrictFromAffiliate.Contract.Owner(&_RestrictFromAffiliate.CallOpts)
}

// Test is a free data retrieval call binding the contract method 0x0b6dedfd.
//
// Solidity: function test(compliance address, token address, initiator address, from address, to address, tokens uint256) constant returns(s string)
func (_RestrictFromAffiliate *RestrictFromAffiliateCaller) Test(opts *bind.CallOpts, compliance common.Address, token common.Address, initiator common.Address, from common.Address, to common.Address, tokens *big.Int) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _RestrictFromAffiliate.contract.Call(opts, out, "test", compliance, token, initiator, from, to, tokens)
	return *ret0, err
}

// Test is a free data retrieval call binding the contract method 0x0b6dedfd.
//
// Solidity: function test(compliance address, token address, initiator address, from address, to address, tokens uint256) constant returns(s string)
func (_RestrictFromAffiliate *RestrictFromAffiliateSession) Test(compliance common.Address, token common.Address, initiator common.Address, from common.Address, to common.Address, tokens *big.Int) (string, error) {
	return _RestrictFromAffiliate.Contract.Test(&_RestrictFromAffiliate.CallOpts, compliance, token, initiator, from, to, tokens)
}

// Test is a free data retrieval call binding the contract method 0x0b6dedfd.
//
// Solidity: function test(compliance address, token address, initiator address, from address, to address, tokens uint256) constant returns(s string)
func (_RestrictFromAffiliate *RestrictFromAffiliateCallerSession) Test(compliance common.Address, token common.Address, initiator common.Address, from common.Address, to common.Address, tokens *big.Int) (string, error) {
	return _RestrictFromAffiliate.Contract.Test(&_RestrictFromAffiliate.CallOpts, compliance, token, initiator, from, to, tokens)
}

// Check is a paid mutator transaction binding the contract method 0x803fcd43.
//
// Solidity: function check(token address, initiator address, from address, to address, tokens uint256) returns()
func (_RestrictFromAffiliate *RestrictFromAffiliateTransactor) Check(opts *bind.TransactOpts, token common.Address, initiator common.Address, from common.Address, to common.Address, tokens *big.Int) (*types.Transaction, error) {
	return _RestrictFromAffiliate.contract.Transact(opts, "check", token, initiator, from, to, tokens)
}

// Check is a paid mutator transaction binding the contract method 0x803fcd43.
//
// Solidity: function check(token address, initiator address, from address, to address, tokens uint256) returns()
func (_RestrictFromAffiliate *RestrictFromAffiliateSession) Check(token common.Address, initiator common.Address, from common.Address, to common.Address, tokens *big.Int) (*types.Transaction, error) {
	return _RestrictFromAffiliate.Contract.Check(&_RestrictFromAffiliate.TransactOpts, token, initiator, from, to, tokens)
}

// Check is a paid mutator transaction binding the contract method 0x803fcd43.
//
// Solidity: function check(token address, initiator address, from address, to address, tokens uint256) returns()
func (_RestrictFromAffiliate *RestrictFromAffiliateTransactorSession) Check(token common.Address, initiator common.Address, from common.Address, to common.Address, tokens *big.Int) (*types.Transaction, error) {
	return _RestrictFromAffiliate.Contract.Check(&_RestrictFromAffiliate.TransactOpts, token, initiator, from, to, tokens)
}

// Kill is a paid mutator transaction binding the contract method 0x41c0e1b5.
//
// Solidity: function kill() returns()
func (_RestrictFromAffiliate *RestrictFromAffiliateTransactor) Kill(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RestrictFromAffiliate.contract.Transact(opts, "kill")
}

// Kill is a paid mutator transaction binding the contract method 0x41c0e1b5.
//
// Solidity: function kill() returns()
func (_RestrictFromAffiliate *RestrictFromAffiliateSession) Kill() (*types.Transaction, error) {
	return _RestrictFromAffiliate.Contract.Kill(&_RestrictFromAffiliate.TransactOpts)
}

// Kill is a paid mutator transaction binding the contract method 0x41c0e1b5.
//
// Solidity: function kill() returns()
func (_RestrictFromAffiliate *RestrictFromAffiliateTransactorSession) Kill() (*types.Transaction, error) {
	return _RestrictFromAffiliate.Contract.Kill(&_RestrictFromAffiliate.TransactOpts)
}

// TransferOwner is a paid mutator transaction binding the contract method 0x4fb2e45d.
//
// Solidity: function transferOwner(newOwner address) returns()
func (_RestrictFromAffiliate *RestrictFromAffiliateTransactor) TransferOwner(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _RestrictFromAffiliate.contract.Transact(opts, "transferOwner", newOwner)
}

// TransferOwner is a paid mutator transaction binding the contract method 0x4fb2e45d.
//
// Solidity: function transferOwner(newOwner address) returns()
func (_RestrictFromAffiliate *RestrictFromAffiliateSession) TransferOwner(newOwner common.Address) (*types.Transaction, error) {
	return _RestrictFromAffiliate.Contract.TransferOwner(&_RestrictFromAffiliate.TransactOpts, newOwner)
}

// TransferOwner is a paid mutator transaction binding the contract method 0x4fb2e45d.
//
// Solidity: function transferOwner(newOwner address) returns()
func (_RestrictFromAffiliate *RestrictFromAffiliateTransactorSession) TransferOwner(newOwner common.Address) (*types.Transaction, error) {
	return _RestrictFromAffiliate.Contract.TransferOwner(&_RestrictFromAffiliate.TransactOpts, newOwner)
}

// RestrictFromAffiliateOwnerTransferredIterator is returned from FilterOwnerTransferred and is used to iterate over the raw logs and unpacked data for OwnerTransferred events raised by the RestrictFromAffiliate contract.
type RestrictFromAffiliateOwnerTransferredIterator struct {
	Event *RestrictFromAffiliateOwnerTransferred // Event containing the contract specifics and raw log

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
func (it *RestrictFromAffiliateOwnerTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RestrictFromAffiliateOwnerTransferred)
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
		it.Event = new(RestrictFromAffiliateOwnerTransferred)
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
func (it *RestrictFromAffiliateOwnerTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RestrictFromAffiliateOwnerTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RestrictFromAffiliateOwnerTransferred represents a OwnerTransferred event raised by the RestrictFromAffiliate contract.
type RestrictFromAffiliateOwnerTransferred struct {
	OldOwner common.Address
	NewOwner common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOwnerTransferred is a free log retrieval operation binding the contract event 0x8934ce4adea8d9ce0d714d2c22b86790e41b7731c84b926fbbdc1d40ff6533c9.
//
// Solidity: e OwnerTransferred(oldOwner indexed address, newOwner indexed address)
func (_RestrictFromAffiliate *RestrictFromAffiliateFilterer) FilterOwnerTransferred(opts *bind.FilterOpts, oldOwner []common.Address, newOwner []common.Address) (*RestrictFromAffiliateOwnerTransferredIterator, error) {

	var oldOwnerRule []interface{}
	for _, oldOwnerItem := range oldOwner {
		oldOwnerRule = append(oldOwnerRule, oldOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _RestrictFromAffiliate.contract.FilterLogs(opts, "OwnerTransferred", oldOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &RestrictFromAffiliateOwnerTransferredIterator{contract: _RestrictFromAffiliate.contract, event: "OwnerTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnerTransferred is a free log subscription operation binding the contract event 0x8934ce4adea8d9ce0d714d2c22b86790e41b7731c84b926fbbdc1d40ff6533c9.
//
// Solidity: e OwnerTransferred(oldOwner indexed address, newOwner indexed address)
func (_RestrictFromAffiliate *RestrictFromAffiliateFilterer) WatchOwnerTransferred(opts *bind.WatchOpts, sink chan<- *RestrictFromAffiliateOwnerTransferred, oldOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var oldOwnerRule []interface{}
	for _, oldOwnerItem := range oldOwner {
		oldOwnerRule = append(oldOwnerRule, oldOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _RestrictFromAffiliate.contract.WatchLogs(opts, "OwnerTransferred", oldOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RestrictFromAffiliateOwnerTransferred)
				if err := _RestrictFromAffiliate.contract.UnpackLog(event, "OwnerTransferred", log); err != nil {
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
