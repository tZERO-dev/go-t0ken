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

// RestrictABI is the input ABI used to generate the binding from.
const RestrictABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"contractICompliance\",\"name\":\"compliance\",\"type\":\"address\"},{\"internalType\":\"contractIT0ken\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"initiator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokens\",\"type\":\"uint256\"}],\"name\":\"test\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"s\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"kill\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwner\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ZERO_ADDRESS\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"contractIT0ken\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"initiator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokens\",\"type\":\"uint256\"}],\"name\":\"check\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnerTransferred\",\"type\":\"event\"}]"

// RestrictBin is the compiled bytecode used for deploying new contracts.
const RestrictBin = `600180546001600160a01b031916905560c0604052600860808190527f526573747269637400000000000000000000000000000000000000000000000060a090815261004e9160029190610066565b50600080546001600160a01b03191633179055610101565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106100a757805160ff19168380011785556100d4565b828001600101855582156100d4579182015b828111156100d45782518255916020019190600101906100b9565b506100e09291506100e4565b5090565b6100fe91905b808211156100e057600081556001016100ea565b90565b610afb806101106000396000f3fe608060405234801561001057600080fd5b506004361061007d5760003560e01c80634fb2e45d1161005b5780634fb2e45d14610155578063538ba4f91461017b578063803fcd431461019f5780638da5cb5b146101e55761007d565b806306fdde03146100825780630b6dedfd146100ff57806341c0e1b51461014b575b600080fd5b61008a6101ed565b6040805160208082528351818301528351919283929083019185019080838360005b838110156100c45781810151838201526020016100ac565b50505050905090810190601f1680156100f15780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b61008a600480360360c081101561011557600080fd5b506001600160a01b0381358116916020810135821691604082013581169160608101358216916080820135169060a00135610278565b610153610575565b005b6101536004803603602081101561016b57600080fd5b50356001600160a01b03166105fd565b610183610787565b604080516001600160a01b039092168252519081900360200190f35b610153600480360360a08110156101b557600080fd5b506001600160a01b0381358116916020810135821691604082013581169160608101359091169060800135610796565b610183610a26565b6002805460408051602060018416156101000260001901909316849004601f810184900484028201840190925281815292918301828280156102705780601f1061024557610100808354040283529160200191610270565b820191906000526020600020905b81548152906001019060200180831161025357829003601f168201915b505050505081565b60606000866001600160a01b03166395d89b416040518163ffffffff1660e01b815260040160006040518083038186803b1580156102b557600080fd5b505afa1580156102c9573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f1916820160405260208110156102f257600080fd5b810190808051604051939291908464010000000082111561031257600080fd5b90830190602082018581111561032757600080fd5b825164010000000081118282018810171561034157600080fd5b82525081516020918201929091019080838360005b8381101561036e578181015183820152602001610356565b50505050905090810190601f16801561039b5780820380516001836020036101000a031916815260200191505b5060405250505060405160200180807f52657374726963742e697352657374726963746564000000000000000000000081525060150182805190602001908083835b602083106103fc5780518252601f1990920191602091820191016103dd565b6001836020036101000a038019825116818451168082178552505050505050905001915050604051602081830303815290604052805190602001209050876001600160a01b031663975057e76040518163ffffffff1660e01b815260040160206040518083038186803b15801561047257600080fd5b505afa158015610486573d6000803e3d6000fd5b505050506040513d602081101561049c57600080fd5b5051604080517f7ae1cfca0000000000000000000000000000000000000000000000000000000081526004810184905290516001600160a01b0390921691637ae1cfca91602480820192602092909190829003018186803b15801561050057600080fd5b505afa158015610514573d6000803e3d6000fd5b505050506040513d602081101561052a57600080fd5b50511561056a576040518060400160405280602081526020017f5265737472696374696f6e2069732063757272656e746c7920656e61626c656481525091505b509695505050505050565b6000546001600160a01b031633148061059e57506001546000546001600160a01b039081169116145b6105ef576040805162461bcd60e51b815260206004820152601960248201527f4f776e6572206163636f756e7420697320726571756972656400000000000000604482015290519081900360640190fd5b6000546001600160a01b0316ff5b6000546001600160a01b031633148061062657506001546000546001600160a01b039081169116145b610677576040805162461bcd60e51b815260206004820152601960248201527f4f776e6572206163636f756e7420697320726571756972656400000000000000604482015290519081900360640190fd5b6000546001600160a01b03828116911614156106c45760405162461bcd60e51b8152600401808060200182810382526025815260200180610aa26025913960400191505060405180910390fd5b6001600160a01b03811661071f576040805162461bcd60e51b815260206004820181905260248201527f4e6577204f776e65722063616e6e6f74206265207a65726f2061646472657373604482015290519081900360640190fd5b600080546001600160a01b038381167fffffffffffffffffffffffff0000000000000000000000000000000000000000831681178455604051919092169283917f8934ce4adea8d9ce0d714d2c22b86790e41b7731c84b926fbbdc1d40ff6533c99190a35050565b6001546001600160a01b031681565b6000856001600160a01b03166395d89b416040518163ffffffff1660e01b815260040160006040518083038186803b1580156107d157600080fd5b505afa1580156107e5573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052602081101561080e57600080fd5b810190808051604051939291908464010000000082111561082e57600080fd5b90830190602082018581111561084357600080fd5b825164010000000081118282018810171561085d57600080fd5b82525081516020918201929091019080838360005b8381101561088a578181015183820152602001610872565b50505050905090810190601f1680156108b75780820380516001836020036101000a031916815260200191505b5060405250505060405160200180807f52657374726963742e697352657374726963746564000000000000000000000081525060150182805190602001908083835b602083106109185780518252601f1990920191602091820191016108f9565b6001836020036101000a03801982511681845116808217855250505050505090500191505060405160208183030381529060405280519060200120905061095d610a35565b6001600160a01b0316637ae1cfca826040518263ffffffff1660e01b81526004018082815260200191505060206040518083038186803b1580156109a057600080fd5b505afa1580156109b4573d6000803e3d6000fd5b505050506040513d60208110156109ca57600080fd5b505115610a1e576040805162461bcd60e51b815260206004820181905260248201527f5265737472696374696f6e2069732063757272656e746c7920656e61626c6564604482015290519081900360640190fd5b505050505050565b6000546001600160a01b031681565b6000336001600160a01b031663975057e76040518163ffffffff1660e01b815260040160206040518083038186803b158015610a7057600080fd5b505afa158015610a84573d6000803e3d6000fd5b505050506040513d6020811015610a9a57600080fd5b505190509056fe4e6577204f776e65722063616e6e6f74206265207468652063757272656e74206f776e6572a265627a7a723158204fc46831be39f8f5a15256eb42842eb6d1995aec48d17a712406da1e2f171a5164736f6c634300050b0032`

// DeployRestrict deploys a new Ethereum contract, binding an instance of Restrict to it.
func DeployRestrict(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Restrict, error) {
	parsed, err := abi.JSON(strings.NewReader(RestrictABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(RestrictBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Restrict{RestrictCaller: RestrictCaller{contract: contract}, RestrictTransactor: RestrictTransactor{contract: contract}, RestrictFilterer: RestrictFilterer{contract: contract}}, nil
}

// Restrict is an auto generated Go binding around an Ethereum contract.
type Restrict struct {
	RestrictCaller     // Read-only binding to the contract
	RestrictTransactor // Write-only binding to the contract
	RestrictFilterer   // Log filterer for contract events
}

// RestrictCaller is an auto generated read-only Go binding around an Ethereum contract.
type RestrictCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RestrictTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RestrictTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RestrictFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RestrictFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RestrictSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RestrictSession struct {
	Contract     *Restrict         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RestrictCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RestrictCallerSession struct {
	Contract *RestrictCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// RestrictTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RestrictTransactorSession struct {
	Contract     *RestrictTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// RestrictRaw is an auto generated low-level Go binding around an Ethereum contract.
type RestrictRaw struct {
	Contract *Restrict // Generic contract binding to access the raw methods on
}

// RestrictCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RestrictCallerRaw struct {
	Contract *RestrictCaller // Generic read-only contract binding to access the raw methods on
}

// RestrictTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RestrictTransactorRaw struct {
	Contract *RestrictTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRestrict creates a new instance of Restrict, bound to a specific deployed contract.
func NewRestrict(address common.Address, backend bind.ContractBackend) (*Restrict, error) {
	contract, err := bindRestrict(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Restrict{RestrictCaller: RestrictCaller{contract: contract}, RestrictTransactor: RestrictTransactor{contract: contract}, RestrictFilterer: RestrictFilterer{contract: contract}}, nil
}

// NewRestrictCaller creates a new read-only instance of Restrict, bound to a specific deployed contract.
func NewRestrictCaller(address common.Address, caller bind.ContractCaller) (*RestrictCaller, error) {
	contract, err := bindRestrict(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RestrictCaller{contract: contract}, nil
}

// NewRestrictTransactor creates a new write-only instance of Restrict, bound to a specific deployed contract.
func NewRestrictTransactor(address common.Address, transactor bind.ContractTransactor) (*RestrictTransactor, error) {
	contract, err := bindRestrict(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RestrictTransactor{contract: contract}, nil
}

// NewRestrictFilterer creates a new log filterer instance of Restrict, bound to a specific deployed contract.
func NewRestrictFilterer(address common.Address, filterer bind.ContractFilterer) (*RestrictFilterer, error) {
	contract, err := bindRestrict(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RestrictFilterer{contract: contract}, nil
}

// bindRestrict binds a generic wrapper to an already deployed contract.
func bindRestrict(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(RestrictABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Restrict *RestrictRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Restrict.Contract.RestrictCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Restrict *RestrictRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Restrict.Contract.RestrictTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Restrict *RestrictRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Restrict.Contract.RestrictTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Restrict *RestrictCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Restrict.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Restrict *RestrictTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Restrict.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Restrict *RestrictTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Restrict.Contract.contract.Transact(opts, method, params...)
}

// ZEROADDRESS is a free data retrieval call binding the contract method 0x538ba4f9.
//
// Solidity: function ZERO_ADDRESS() constant returns(address)
func (_Restrict *RestrictCaller) ZEROADDRESS(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Restrict.contract.Call(opts, out, "ZERO_ADDRESS")
	return *ret0, err
}

// ZEROADDRESS is a free data retrieval call binding the contract method 0x538ba4f9.
//
// Solidity: function ZERO_ADDRESS() constant returns(address)
func (_Restrict *RestrictSession) ZEROADDRESS() (common.Address, error) {
	return _Restrict.Contract.ZEROADDRESS(&_Restrict.CallOpts)
}

// ZEROADDRESS is a free data retrieval call binding the contract method 0x538ba4f9.
//
// Solidity: function ZERO_ADDRESS() constant returns(address)
func (_Restrict *RestrictCallerSession) ZEROADDRESS() (common.Address, error) {
	return _Restrict.Contract.ZEROADDRESS(&_Restrict.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_Restrict *RestrictCaller) Name(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Restrict.contract.Call(opts, out, "name")
	return *ret0, err
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_Restrict *RestrictSession) Name() (string, error) {
	return _Restrict.Contract.Name(&_Restrict.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_Restrict *RestrictCallerSession) Name() (string, error) {
	return _Restrict.Contract.Name(&_Restrict.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Restrict *RestrictCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Restrict.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Restrict *RestrictSession) Owner() (common.Address, error) {
	return _Restrict.Contract.Owner(&_Restrict.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Restrict *RestrictCallerSession) Owner() (common.Address, error) {
	return _Restrict.Contract.Owner(&_Restrict.CallOpts)
}

// Test is a free data retrieval call binding the contract method 0x0b6dedfd.
//
// Solidity: function test(compliance address, token address, initiator address, from address, to address, tokens uint256) constant returns(s string)
func (_Restrict *RestrictCaller) Test(opts *bind.CallOpts, compliance common.Address, token common.Address, initiator common.Address, from common.Address, to common.Address, tokens *big.Int) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Restrict.contract.Call(opts, out, "test", compliance, token, initiator, from, to, tokens)
	return *ret0, err
}

// Test is a free data retrieval call binding the contract method 0x0b6dedfd.
//
// Solidity: function test(compliance address, token address, initiator address, from address, to address, tokens uint256) constant returns(s string)
func (_Restrict *RestrictSession) Test(compliance common.Address, token common.Address, initiator common.Address, from common.Address, to common.Address, tokens *big.Int) (string, error) {
	return _Restrict.Contract.Test(&_Restrict.CallOpts, compliance, token, initiator, from, to, tokens)
}

// Test is a free data retrieval call binding the contract method 0x0b6dedfd.
//
// Solidity: function test(compliance address, token address, initiator address, from address, to address, tokens uint256) constant returns(s string)
func (_Restrict *RestrictCallerSession) Test(compliance common.Address, token common.Address, initiator common.Address, from common.Address, to common.Address, tokens *big.Int) (string, error) {
	return _Restrict.Contract.Test(&_Restrict.CallOpts, compliance, token, initiator, from, to, tokens)
}

// Check is a paid mutator transaction binding the contract method 0x803fcd43.
//
// Solidity: function check(token address, initiator address, from address, to address, tokens uint256) returns()
func (_Restrict *RestrictTransactor) Check(opts *bind.TransactOpts, token common.Address, initiator common.Address, from common.Address, to common.Address, tokens *big.Int) (*types.Transaction, error) {
	return _Restrict.contract.Transact(opts, "check", token, initiator, from, to, tokens)
}

// Check is a paid mutator transaction binding the contract method 0x803fcd43.
//
// Solidity: function check(token address, initiator address, from address, to address, tokens uint256) returns()
func (_Restrict *RestrictSession) Check(token common.Address, initiator common.Address, from common.Address, to common.Address, tokens *big.Int) (*types.Transaction, error) {
	return _Restrict.Contract.Check(&_Restrict.TransactOpts, token, initiator, from, to, tokens)
}

// Check is a paid mutator transaction binding the contract method 0x803fcd43.
//
// Solidity: function check(token address, initiator address, from address, to address, tokens uint256) returns()
func (_Restrict *RestrictTransactorSession) Check(token common.Address, initiator common.Address, from common.Address, to common.Address, tokens *big.Int) (*types.Transaction, error) {
	return _Restrict.Contract.Check(&_Restrict.TransactOpts, token, initiator, from, to, tokens)
}

// Kill is a paid mutator transaction binding the contract method 0x41c0e1b5.
//
// Solidity: function kill() returns()
func (_Restrict *RestrictTransactor) Kill(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Restrict.contract.Transact(opts, "kill")
}

// Kill is a paid mutator transaction binding the contract method 0x41c0e1b5.
//
// Solidity: function kill() returns()
func (_Restrict *RestrictSession) Kill() (*types.Transaction, error) {
	return _Restrict.Contract.Kill(&_Restrict.TransactOpts)
}

// Kill is a paid mutator transaction binding the contract method 0x41c0e1b5.
//
// Solidity: function kill() returns()
func (_Restrict *RestrictTransactorSession) Kill() (*types.Transaction, error) {
	return _Restrict.Contract.Kill(&_Restrict.TransactOpts)
}

// TransferOwner is a paid mutator transaction binding the contract method 0x4fb2e45d.
//
// Solidity: function transferOwner(newOwner address) returns()
func (_Restrict *RestrictTransactor) TransferOwner(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Restrict.contract.Transact(opts, "transferOwner", newOwner)
}

// TransferOwner is a paid mutator transaction binding the contract method 0x4fb2e45d.
//
// Solidity: function transferOwner(newOwner address) returns()
func (_Restrict *RestrictSession) TransferOwner(newOwner common.Address) (*types.Transaction, error) {
	return _Restrict.Contract.TransferOwner(&_Restrict.TransactOpts, newOwner)
}

// TransferOwner is a paid mutator transaction binding the contract method 0x4fb2e45d.
//
// Solidity: function transferOwner(newOwner address) returns()
func (_Restrict *RestrictTransactorSession) TransferOwner(newOwner common.Address) (*types.Transaction, error) {
	return _Restrict.Contract.TransferOwner(&_Restrict.TransactOpts, newOwner)
}

// RestrictOwnerTransferredIterator is returned from FilterOwnerTransferred and is used to iterate over the raw logs and unpacked data for OwnerTransferred events raised by the Restrict contract.
type RestrictOwnerTransferredIterator struct {
	Event *RestrictOwnerTransferred // Event containing the contract specifics and raw log

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
func (it *RestrictOwnerTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RestrictOwnerTransferred)
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
		it.Event = new(RestrictOwnerTransferred)
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
func (it *RestrictOwnerTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RestrictOwnerTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RestrictOwnerTransferred represents a OwnerTransferred event raised by the Restrict contract.
type RestrictOwnerTransferred struct {
	OldOwner common.Address
	NewOwner common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOwnerTransferred is a free log retrieval operation binding the contract event 0x8934ce4adea8d9ce0d714d2c22b86790e41b7731c84b926fbbdc1d40ff6533c9.
//
// Solidity: e OwnerTransferred(oldOwner indexed address, newOwner indexed address)
func (_Restrict *RestrictFilterer) FilterOwnerTransferred(opts *bind.FilterOpts, oldOwner []common.Address, newOwner []common.Address) (*RestrictOwnerTransferredIterator, error) {

	var oldOwnerRule []interface{}
	for _, oldOwnerItem := range oldOwner {
		oldOwnerRule = append(oldOwnerRule, oldOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Restrict.contract.FilterLogs(opts, "OwnerTransferred", oldOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &RestrictOwnerTransferredIterator{contract: _Restrict.contract, event: "OwnerTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnerTransferred is a free log subscription operation binding the contract event 0x8934ce4adea8d9ce0d714d2c22b86790e41b7731c84b926fbbdc1d40ff6533c9.
//
// Solidity: e OwnerTransferred(oldOwner indexed address, newOwner indexed address)
func (_Restrict *RestrictFilterer) WatchOwnerTransferred(opts *bind.WatchOpts, sink chan<- *RestrictOwnerTransferred, oldOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var oldOwnerRule []interface{}
	for _, oldOwnerItem := range oldOwner {
		oldOwnerRule = append(oldOwnerRule, oldOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Restrict.contract.WatchLogs(opts, "OwnerTransferred", oldOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RestrictOwnerTransferred)
				if err := _Restrict.contract.UnpackLog(event, "OwnerTransferred", log); err != nil {
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
