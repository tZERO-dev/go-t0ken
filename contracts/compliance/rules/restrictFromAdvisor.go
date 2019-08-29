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

// RestrictFromAdvisorABI is the input ABI used to generate the binding from.
const RestrictFromAdvisorABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"compliance\",\"type\":\"address\"},{\"name\":\"token\",\"type\":\"address\"},{\"name\":\"initiator\",\"type\":\"address\"},{\"name\":\"from\",\"type\":\"address\"},{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"tokens\",\"type\":\"uint256\"}],\"name\":\"test\",\"outputs\":[{\"name\":\"s\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"kill\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwner\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ZERO_ADDRESS\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"token\",\"type\":\"address\"},{\"name\":\"initiator\",\"type\":\"address\"},{\"name\":\"from\",\"type\":\"address\"},{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"tokens\",\"type\":\"uint256\"}],\"name\":\"check\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"oldOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnerTransferred\",\"type\":\"event\"}]"

// RestrictFromAdvisorBin is the compiled bytecode used for deploying new contracts.
const RestrictFromAdvisorBin = `600180546001600160a01b031916905560c0604052601560808190527f52657374726963742046726f6d2041647669736f72000000000000000000000060a090815261004e9160029190610066565b50600080546001600160a01b03191633179055610101565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106100a757805160ff19168380011785556100d4565b828001600101855582156100d4579182015b828111156100d45782518255916020019190600101906100b9565b506100e09291506100e4565b5090565b6100fe91905b808211156100e057600081556001016100ea565b90565b610a62806101106000396000f3fe608060405234801561001057600080fd5b506004361061007d5760003560e01c80634fb2e45d1161005b5780634fb2e45d14610155578063538ba4f91461017b578063803fcd431461019f5780638da5cb5b146101e55761007d565b806306fdde03146100825780630b6dedfd146100ff57806341c0e1b51461014b575b600080fd5b61008a6101ed565b6040805160208082528351818301528351919283929083019185019080838360005b838110156100c45781810151838201526020016100ac565b50505050905090810190601f1680156100f15780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b61008a600480360360c081101561011557600080fd5b506001600160a01b0381358116916020810135821691604082013581169160608101358216916080820135169060a00135610278565b610153610512565b005b6101536004803603602081101561016b57600080fd5b50356001600160a01b031661059a565b610183610724565b604080516001600160a01b039092168252519081900360200190f35b610153600480360360a08110156101b557600080fd5b506001600160a01b0381358116916020810135821691604082013581169160608101359091169060800135610733565b610183610967565b6002805460408051602060018416156101000260001901909316849004601f810184900484028201840190925281815292918301828280156102705780601f1061024557610100808354040283529160200191610270565b820191906000526020600020905b81548152906001019060200180831161025357829003601f168201915b505050505081565b60606000866001600160a01b03166395d89b416040518163ffffffff1660e01b815260040160006040518083038186803b1580156102b557600080fd5b505afa1580156102c9573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f1916820160405260208110156102f257600080fd5b81019080805164010000000081111561030a57600080fd5b8201602081018481111561031d57600080fd5b815164010000000081118282018710171561033757600080fd5b50509291905050508560405160200180807f526573747269637446726f6d41647669736f722e697341647669736f72000000815250601d0183805190602001908083835b6020831061039a5780518252601f19909201916020918201910161037b565b6001836020036101000a038019825116818451168082178552505050505050905001826001600160a01b03166001600160a01b031660601b815260140192505050604051602081830303815290604052805190602001209050876001600160a01b031663975057e76040518163ffffffff1660e01b815260040160206040518083038186803b15801561042c57600080fd5b505afa158015610440573d6000803e3d6000fd5b505050506040513d602081101561045657600080fd5b5051604080517f7ae1cfca0000000000000000000000000000000000000000000000000000000081526004810184905290516001600160a01b0390921691637ae1cfca91602480820192602092909190829003018186803b1580156104ba57600080fd5b505afa1580156104ce573d6000803e3d6000fd5b505050506040513d60208110156104e457600080fd5b50511561050757604051806060016040528060268152602001610a086026913991505b509695505050505050565b6000546001600160a01b031633148061053b57506001546000546001600160a01b039081169116145b61058c576040805162461bcd60e51b815260206004820152601960248201527f4f776e6572206163636f756e7420697320726571756972656400000000000000604482015290519081900360640190fd5b6000546001600160a01b0316ff5b6000546001600160a01b03163314806105c357506001546000546001600160a01b039081169116145b610614576040805162461bcd60e51b815260206004820152601960248201527f4f776e6572206163636f756e7420697320726571756972656400000000000000604482015290519081900360640190fd5b6000546001600160a01b03828116911614156106615760405162461bcd60e51b81526004018080602001828103825260258152602001806109e36025913960400191505060405180910390fd5b6001600160a01b0381166106bc576040805162461bcd60e51b815260206004820181905260248201527f4e6577204f776e65722063616e6e6f74206265207a65726f2061646472657373604482015290519081900360640190fd5b600080546001600160a01b038381167fffffffffffffffffffffffff0000000000000000000000000000000000000000831681178455604051919092169283917f8934ce4adea8d9ce0d714d2c22b86790e41b7731c84b926fbbdc1d40ff6533c99190a35050565b6001546001600160a01b031681565b6000856001600160a01b03166395d89b416040518163ffffffff1660e01b815260040160006040518083038186803b15801561076e57600080fd5b505afa158015610782573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f1916820160405260208110156107ab57600080fd5b8101908080516401000000008111156107c357600080fd5b820160208101848111156107d657600080fd5b81516401000000008111828201871017156107f057600080fd5b50509291905050508460405160200180807f526573747269637446726f6d41647669736f722e697341647669736f72000000815250601d0183805190602001908083835b602083106108535780518252601f199092019160209182019101610834565b6001836020036101000a038019825116818451168082178552505050505050905001826001600160a01b03166001600160a01b031660601b8152601401925050506040516020818303038152906040528051906020012090506108b4610976565b6001600160a01b0316637ae1cfca826040518263ffffffff1660e01b81526004018082815260200191505060206040518083038186803b1580156108f757600080fd5b505afa15801561090b573d6000803e3d6000fd5b505050506040513d602081101561092157600080fd5b50511561095f5760405162461bcd60e51b8152600401808060200182810382526026815260200180610a086026913960400191505060405180910390fd5b505050505050565b6000546001600160a01b031681565b6000336001600160a01b031663975057e76040518163ffffffff1660e01b815260040160206040518083038186803b1580156109b157600080fd5b505afa1580156109c5573d6000803e3d6000fd5b505050506040513d60208110156109db57600080fd5b505190509056fe4e6577204f776e65722063616e6e6f74206265207468652063757272656e74206f776e65725472616e73666572732066726f6d2061647669736f7273206172652072657374726963746564a265627a7a72305820620034cb53ad98ae70b6bb7496eefda069c473253b26eb0b37ee1bde5a8a147664736f6c634300050a0032`

// DeployRestrictFromAdvisor deploys a new Ethereum contract, binding an instance of RestrictFromAdvisor to it.
func DeployRestrictFromAdvisor(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *RestrictFromAdvisor, error) {
	parsed, err := abi.JSON(strings.NewReader(RestrictFromAdvisorABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(RestrictFromAdvisorBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &RestrictFromAdvisor{RestrictFromAdvisorCaller: RestrictFromAdvisorCaller{contract: contract}, RestrictFromAdvisorTransactor: RestrictFromAdvisorTransactor{contract: contract}, RestrictFromAdvisorFilterer: RestrictFromAdvisorFilterer{contract: contract}}, nil
}

// RestrictFromAdvisor is an auto generated Go binding around an Ethereum contract.
type RestrictFromAdvisor struct {
	RestrictFromAdvisorCaller     // Read-only binding to the contract
	RestrictFromAdvisorTransactor // Write-only binding to the contract
	RestrictFromAdvisorFilterer   // Log filterer for contract events
}

// RestrictFromAdvisorCaller is an auto generated read-only Go binding around an Ethereum contract.
type RestrictFromAdvisorCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RestrictFromAdvisorTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RestrictFromAdvisorTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RestrictFromAdvisorFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RestrictFromAdvisorFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RestrictFromAdvisorSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RestrictFromAdvisorSession struct {
	Contract     *RestrictFromAdvisor // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// RestrictFromAdvisorCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RestrictFromAdvisorCallerSession struct {
	Contract *RestrictFromAdvisorCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// RestrictFromAdvisorTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RestrictFromAdvisorTransactorSession struct {
	Contract     *RestrictFromAdvisorTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// RestrictFromAdvisorRaw is an auto generated low-level Go binding around an Ethereum contract.
type RestrictFromAdvisorRaw struct {
	Contract *RestrictFromAdvisor // Generic contract binding to access the raw methods on
}

// RestrictFromAdvisorCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RestrictFromAdvisorCallerRaw struct {
	Contract *RestrictFromAdvisorCaller // Generic read-only contract binding to access the raw methods on
}

// RestrictFromAdvisorTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RestrictFromAdvisorTransactorRaw struct {
	Contract *RestrictFromAdvisorTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRestrictFromAdvisor creates a new instance of RestrictFromAdvisor, bound to a specific deployed contract.
func NewRestrictFromAdvisor(address common.Address, backend bind.ContractBackend) (*RestrictFromAdvisor, error) {
	contract, err := bindRestrictFromAdvisor(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &RestrictFromAdvisor{RestrictFromAdvisorCaller: RestrictFromAdvisorCaller{contract: contract}, RestrictFromAdvisorTransactor: RestrictFromAdvisorTransactor{contract: contract}, RestrictFromAdvisorFilterer: RestrictFromAdvisorFilterer{contract: contract}}, nil
}

// NewRestrictFromAdvisorCaller creates a new read-only instance of RestrictFromAdvisor, bound to a specific deployed contract.
func NewRestrictFromAdvisorCaller(address common.Address, caller bind.ContractCaller) (*RestrictFromAdvisorCaller, error) {
	contract, err := bindRestrictFromAdvisor(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RestrictFromAdvisorCaller{contract: contract}, nil
}

// NewRestrictFromAdvisorTransactor creates a new write-only instance of RestrictFromAdvisor, bound to a specific deployed contract.
func NewRestrictFromAdvisorTransactor(address common.Address, transactor bind.ContractTransactor) (*RestrictFromAdvisorTransactor, error) {
	contract, err := bindRestrictFromAdvisor(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RestrictFromAdvisorTransactor{contract: contract}, nil
}

// NewRestrictFromAdvisorFilterer creates a new log filterer instance of RestrictFromAdvisor, bound to a specific deployed contract.
func NewRestrictFromAdvisorFilterer(address common.Address, filterer bind.ContractFilterer) (*RestrictFromAdvisorFilterer, error) {
	contract, err := bindRestrictFromAdvisor(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RestrictFromAdvisorFilterer{contract: contract}, nil
}

// bindRestrictFromAdvisor binds a generic wrapper to an already deployed contract.
func bindRestrictFromAdvisor(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(RestrictFromAdvisorABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RestrictFromAdvisor *RestrictFromAdvisorRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _RestrictFromAdvisor.Contract.RestrictFromAdvisorCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RestrictFromAdvisor *RestrictFromAdvisorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RestrictFromAdvisor.Contract.RestrictFromAdvisorTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RestrictFromAdvisor *RestrictFromAdvisorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RestrictFromAdvisor.Contract.RestrictFromAdvisorTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RestrictFromAdvisor *RestrictFromAdvisorCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _RestrictFromAdvisor.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RestrictFromAdvisor *RestrictFromAdvisorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RestrictFromAdvisor.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RestrictFromAdvisor *RestrictFromAdvisorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RestrictFromAdvisor.Contract.contract.Transact(opts, method, params...)
}

// ZEROADDRESS is a free data retrieval call binding the contract method 0x538ba4f9.
//
// Solidity: function ZERO_ADDRESS() constant returns(address)
func (_RestrictFromAdvisor *RestrictFromAdvisorCaller) ZEROADDRESS(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _RestrictFromAdvisor.contract.Call(opts, out, "ZERO_ADDRESS")
	return *ret0, err
}

// ZEROADDRESS is a free data retrieval call binding the contract method 0x538ba4f9.
//
// Solidity: function ZERO_ADDRESS() constant returns(address)
func (_RestrictFromAdvisor *RestrictFromAdvisorSession) ZEROADDRESS() (common.Address, error) {
	return _RestrictFromAdvisor.Contract.ZEROADDRESS(&_RestrictFromAdvisor.CallOpts)
}

// ZEROADDRESS is a free data retrieval call binding the contract method 0x538ba4f9.
//
// Solidity: function ZERO_ADDRESS() constant returns(address)
func (_RestrictFromAdvisor *RestrictFromAdvisorCallerSession) ZEROADDRESS() (common.Address, error) {
	return _RestrictFromAdvisor.Contract.ZEROADDRESS(&_RestrictFromAdvisor.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_RestrictFromAdvisor *RestrictFromAdvisorCaller) Name(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _RestrictFromAdvisor.contract.Call(opts, out, "name")
	return *ret0, err
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_RestrictFromAdvisor *RestrictFromAdvisorSession) Name() (string, error) {
	return _RestrictFromAdvisor.Contract.Name(&_RestrictFromAdvisor.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_RestrictFromAdvisor *RestrictFromAdvisorCallerSession) Name() (string, error) {
	return _RestrictFromAdvisor.Contract.Name(&_RestrictFromAdvisor.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_RestrictFromAdvisor *RestrictFromAdvisorCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _RestrictFromAdvisor.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_RestrictFromAdvisor *RestrictFromAdvisorSession) Owner() (common.Address, error) {
	return _RestrictFromAdvisor.Contract.Owner(&_RestrictFromAdvisor.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_RestrictFromAdvisor *RestrictFromAdvisorCallerSession) Owner() (common.Address, error) {
	return _RestrictFromAdvisor.Contract.Owner(&_RestrictFromAdvisor.CallOpts)
}

// Test is a free data retrieval call binding the contract method 0x0b6dedfd.
//
// Solidity: function test(compliance address, token address, initiator address, from address, to address, tokens uint256) constant returns(s string)
func (_RestrictFromAdvisor *RestrictFromAdvisorCaller) Test(opts *bind.CallOpts, compliance common.Address, token common.Address, initiator common.Address, from common.Address, to common.Address, tokens *big.Int) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _RestrictFromAdvisor.contract.Call(opts, out, "test", compliance, token, initiator, from, to, tokens)
	return *ret0, err
}

// Test is a free data retrieval call binding the contract method 0x0b6dedfd.
//
// Solidity: function test(compliance address, token address, initiator address, from address, to address, tokens uint256) constant returns(s string)
func (_RestrictFromAdvisor *RestrictFromAdvisorSession) Test(compliance common.Address, token common.Address, initiator common.Address, from common.Address, to common.Address, tokens *big.Int) (string, error) {
	return _RestrictFromAdvisor.Contract.Test(&_RestrictFromAdvisor.CallOpts, compliance, token, initiator, from, to, tokens)
}

// Test is a free data retrieval call binding the contract method 0x0b6dedfd.
//
// Solidity: function test(compliance address, token address, initiator address, from address, to address, tokens uint256) constant returns(s string)
func (_RestrictFromAdvisor *RestrictFromAdvisorCallerSession) Test(compliance common.Address, token common.Address, initiator common.Address, from common.Address, to common.Address, tokens *big.Int) (string, error) {
	return _RestrictFromAdvisor.Contract.Test(&_RestrictFromAdvisor.CallOpts, compliance, token, initiator, from, to, tokens)
}

// Check is a paid mutator transaction binding the contract method 0x803fcd43.
//
// Solidity: function check(token address, initiator address, from address, to address, tokens uint256) returns()
func (_RestrictFromAdvisor *RestrictFromAdvisorTransactor) Check(opts *bind.TransactOpts, token common.Address, initiator common.Address, from common.Address, to common.Address, tokens *big.Int) (*types.Transaction, error) {
	return _RestrictFromAdvisor.contract.Transact(opts, "check", token, initiator, from, to, tokens)
}

// Check is a paid mutator transaction binding the contract method 0x803fcd43.
//
// Solidity: function check(token address, initiator address, from address, to address, tokens uint256) returns()
func (_RestrictFromAdvisor *RestrictFromAdvisorSession) Check(token common.Address, initiator common.Address, from common.Address, to common.Address, tokens *big.Int) (*types.Transaction, error) {
	return _RestrictFromAdvisor.Contract.Check(&_RestrictFromAdvisor.TransactOpts, token, initiator, from, to, tokens)
}

// Check is a paid mutator transaction binding the contract method 0x803fcd43.
//
// Solidity: function check(token address, initiator address, from address, to address, tokens uint256) returns()
func (_RestrictFromAdvisor *RestrictFromAdvisorTransactorSession) Check(token common.Address, initiator common.Address, from common.Address, to common.Address, tokens *big.Int) (*types.Transaction, error) {
	return _RestrictFromAdvisor.Contract.Check(&_RestrictFromAdvisor.TransactOpts, token, initiator, from, to, tokens)
}

// Kill is a paid mutator transaction binding the contract method 0x41c0e1b5.
//
// Solidity: function kill() returns()
func (_RestrictFromAdvisor *RestrictFromAdvisorTransactor) Kill(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RestrictFromAdvisor.contract.Transact(opts, "kill")
}

// Kill is a paid mutator transaction binding the contract method 0x41c0e1b5.
//
// Solidity: function kill() returns()
func (_RestrictFromAdvisor *RestrictFromAdvisorSession) Kill() (*types.Transaction, error) {
	return _RestrictFromAdvisor.Contract.Kill(&_RestrictFromAdvisor.TransactOpts)
}

// Kill is a paid mutator transaction binding the contract method 0x41c0e1b5.
//
// Solidity: function kill() returns()
func (_RestrictFromAdvisor *RestrictFromAdvisorTransactorSession) Kill() (*types.Transaction, error) {
	return _RestrictFromAdvisor.Contract.Kill(&_RestrictFromAdvisor.TransactOpts)
}

// TransferOwner is a paid mutator transaction binding the contract method 0x4fb2e45d.
//
// Solidity: function transferOwner(newOwner address) returns()
func (_RestrictFromAdvisor *RestrictFromAdvisorTransactor) TransferOwner(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _RestrictFromAdvisor.contract.Transact(opts, "transferOwner", newOwner)
}

// TransferOwner is a paid mutator transaction binding the contract method 0x4fb2e45d.
//
// Solidity: function transferOwner(newOwner address) returns()
func (_RestrictFromAdvisor *RestrictFromAdvisorSession) TransferOwner(newOwner common.Address) (*types.Transaction, error) {
	return _RestrictFromAdvisor.Contract.TransferOwner(&_RestrictFromAdvisor.TransactOpts, newOwner)
}

// TransferOwner is a paid mutator transaction binding the contract method 0x4fb2e45d.
//
// Solidity: function transferOwner(newOwner address) returns()
func (_RestrictFromAdvisor *RestrictFromAdvisorTransactorSession) TransferOwner(newOwner common.Address) (*types.Transaction, error) {
	return _RestrictFromAdvisor.Contract.TransferOwner(&_RestrictFromAdvisor.TransactOpts, newOwner)
}

// RestrictFromAdvisorOwnerTransferredIterator is returned from FilterOwnerTransferred and is used to iterate over the raw logs and unpacked data for OwnerTransferred events raised by the RestrictFromAdvisor contract.
type RestrictFromAdvisorOwnerTransferredIterator struct {
	Event *RestrictFromAdvisorOwnerTransferred // Event containing the contract specifics and raw log

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
func (it *RestrictFromAdvisorOwnerTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RestrictFromAdvisorOwnerTransferred)
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
		it.Event = new(RestrictFromAdvisorOwnerTransferred)
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
func (it *RestrictFromAdvisorOwnerTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RestrictFromAdvisorOwnerTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RestrictFromAdvisorOwnerTransferred represents a OwnerTransferred event raised by the RestrictFromAdvisor contract.
type RestrictFromAdvisorOwnerTransferred struct {
	OldOwner common.Address
	NewOwner common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOwnerTransferred is a free log retrieval operation binding the contract event 0x8934ce4adea8d9ce0d714d2c22b86790e41b7731c84b926fbbdc1d40ff6533c9.
//
// Solidity: e OwnerTransferred(oldOwner indexed address, newOwner indexed address)
func (_RestrictFromAdvisor *RestrictFromAdvisorFilterer) FilterOwnerTransferred(opts *bind.FilterOpts, oldOwner []common.Address, newOwner []common.Address) (*RestrictFromAdvisorOwnerTransferredIterator, error) {

	var oldOwnerRule []interface{}
	for _, oldOwnerItem := range oldOwner {
		oldOwnerRule = append(oldOwnerRule, oldOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _RestrictFromAdvisor.contract.FilterLogs(opts, "OwnerTransferred", oldOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &RestrictFromAdvisorOwnerTransferredIterator{contract: _RestrictFromAdvisor.contract, event: "OwnerTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnerTransferred is a free log subscription operation binding the contract event 0x8934ce4adea8d9ce0d714d2c22b86790e41b7731c84b926fbbdc1d40ff6533c9.
//
// Solidity: e OwnerTransferred(oldOwner indexed address, newOwner indexed address)
func (_RestrictFromAdvisor *RestrictFromAdvisorFilterer) WatchOwnerTransferred(opts *bind.WatchOpts, sink chan<- *RestrictFromAdvisorOwnerTransferred, oldOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var oldOwnerRule []interface{}
	for _, oldOwnerItem := range oldOwner {
		oldOwnerRule = append(oldOwnerRule, oldOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _RestrictFromAdvisor.contract.WatchLogs(opts, "OwnerTransferred", oldOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RestrictFromAdvisorOwnerTransferred)
				if err := _RestrictFromAdvisor.contract.UnpackLog(event, "OwnerTransferred", log); err != nil {
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
