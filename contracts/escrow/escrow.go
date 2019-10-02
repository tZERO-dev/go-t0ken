// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package escrow

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

// EscrowABI is the input ABI used to generate the binding from.
const EscrowABI = "[{\"constant\":false,\"inputs\":[{\"internalType\":\"contractIT0ken\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"holder\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"quantity\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"grantee\",\"type\":\"address\"}],\"name\":\"accept\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"contractIT0ken\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"holder\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"quantity\",\"type\":\"uint256\"}],\"name\":\"record\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"isAdmin\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"int256\",\"name\":\"index\",\"type\":\"int256\"}],\"name\":\"adminAt\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"add\",\"type\":\"bool\"}],\"name\":\"setAdmin\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwner\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ZERO_ADDRESS\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"registry\",\"outputs\":[{\"internalType\":\"contractIRegistry\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"contractIT0ken\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"holder\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"quantity\",\"type\":\"uint256\"}],\"name\":\"reject\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"holders\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"count\",\"type\":\"int256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"ledger\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"admins\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"count\",\"type\":\"int256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"contractIRegistry\",\"name\":\"r\",\"type\":\"address\"}],\"name\":\"setRegistry\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIRegistry\",\"name\":\"r\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractIT0ken\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"EscrowAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractIT0ken\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"EscrowRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"}],\"name\":\"AdminAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"}],\"name\":\"AdminRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnerTransferred\",\"type\":\"event\"}]"

// EscrowBin is the compiled bytecode used for deploying new contracts.
const EscrowBin = `6080604052600180546001600160a01b031916905534801561002057600080fd5b506040516112e93803806112e98339818101604052602081101561004357600080fd5b5051600080546001600160a01b03199081163317909155600580546001600160a01b0390931692909116919091179055611267806100826000396000f3fe608060405234801561001057600080fd5b50600436106100ea5760003560e01c80637b1039991161008c5780638da5cb5b116100665780638da5cb5b1461028a57806391fe9b6a14610292578063a5de361914610343578063a91ee0dc1461034b576100ea565b80637b10399914610232578063813bedf21461023a5780638188f71c14610270576100ea565b80632bdbc56f116100c85780632bdbc56f1461019d5780634b0bddd2146101d65780634fb2e45d14610204578063538ba4f91461022a576100ea565b8063101e7075146100ef578063172a93fb1461012d57806324d7806c14610163575b600080fd5b61012b6004803603608081101561010557600080fd5b506001600160a01b03813581169160208101358216916040820135916060013516610371565b005b61012b6004803603606081101561014357600080fd5b506001600160a01b03813581169160208101359091169060400135610383565b6101896004803603602081101561017957600080fd5b50356001600160a01b0316610677565b604080519115158252519081900360200190f35b6101ba600480360360208110156101b357600080fd5b50356106a6565b604080516001600160a01b039092168252519081900360200190f35b61012b600480360360408110156101ec57600080fd5b506001600160a01b03813516906020013515156106b9565b61012b6004803603602081101561021a57600080fd5b50356001600160a01b031661086f565b6101ba6109e1565b6101ba6109f0565b61012b6004803603606081101561025057600080fd5b506001600160a01b038135811691602081013590911690604001356109ff565b610278610a10565b60408051918252519081900360200190f35b6101ba610a16565b610278600480360360408110156102a857600080fd5b8101906020810181356401000000008111156102c357600080fd5b8201836020820111156102d557600080fd5b803590602001918460018302840111640100000000831117156102f757600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550505090356001600160a01b03169150610a259050565b610278610a52565b61012b6004803603602081101561036157600080fd5b50356001600160a01b0316610a58565b61037d84848484610af0565b50505050565b61038c33610677565b6103dd576040805162461bcd60e51b815260206004820152601f60248201527f41646d696e20726571756972656420746f207265636f726420657363726f7700604482015290519081900360640190fd5b6060836001600160a01b03166395d89b416040518163ffffffff1660e01b815260040160006040518083038186803b15801561041857600080fd5b505afa15801561042c573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052602081101561045557600080fd5b810190808051604051939291908464010000000082111561047557600080fd5b90830190602082018581111561048a57600080fd5b82516401000000008111828201881017156104a457600080fd5b82525081516020918201929091019080838360005b838110156104d15781810151838201526020016104b9565b50505050905090810190601f1680156104fe5780820380516001836020036101000a031916815260200191505b506040525050509050600061058b836009846040518082805190602001908083835b6020831061053f5780518252601f199092019160209182019101610520565b51815160209384036101000a6000190180199092169116179052920194855250604080519485900382019094206001600160a01b038b1660009081529152929092205492915050610f02565b90506000811180156105a957506105a960068563ffffffff610f6316565b156105f457604080516001600160a01b0380881682528616602082015281517f77ec0a1f539783369bfc02c19452ee9903f7f59cb8f191a0f0937e5e61e8efe5929181900390910190a15b806009836040518082805190602001908083835b602083106106275780518252601f199092019160209182019101610608565b51815160209384036101000a6000190180199092169116179052920194855250604080519485900382019094206001600160a01b0399909916600090815298905250509094209390935550505050565b600080546001600160a01b03838116911614806106a057506106a060028363ffffffff61100816565b92915050565b60006106a060028363ffffffff61103f16565b6000546001600160a01b03163314806106e257506001546000546001600160a01b039081169116145b610733576040805162461bcd60e51b815260206004820152601960248201527f4f776e6572206163636f756e7420697320726571756972656400000000000000604482015290519081900360640190fd5b80156107d45761074a60028363ffffffff610f6316565b61079b576040805162461bcd60e51b815260206004820152601360248201527f556e61626c6520746f206164642061646d696e00000000000000000000000000604482015290519081900360640190fd5b6040516001600160a01b038316907f44d6d25963f097ad14f29f06854a01f575648a1ef82f30e562ccd3889717e33990600090a261086b565b6107e560028363ffffffff6110c516565b610836576040805162461bcd60e51b815260206004820152601660248201527f556e61626c6520746f2072656d6f76652061646d696e00000000000000000000604482015290519081900360640190fd5b6040516001600160a01b038316907fa3b62bc36326052d97ea62d63c3d60308ed4c3ea8ac079dd8499f1e9c4f80c0f90600090a25b5050565b6000546001600160a01b031633148061089857506001546000546001600160a01b039081169116145b6108e9576040805162461bcd60e51b815260206004820152601960248201527f4f776e6572206163636f756e7420697320726571756972656400000000000000604482015290519081900360640190fd5b6000546001600160a01b03828116911614156109365760405162461bcd60e51b815260040180806020018281038252602581526020018061120e6025913960400191505060405180910390fd5b6001600160a01b038116610991576040805162461bcd60e51b815260206004820181905260248201527f4e6577204f776e65722063616e6e6f74206265207a65726f2061646472657373604482015290519081900360640190fd5b600080546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8934ce4adea8d9ce0d714d2c22b86790e41b7731c84b926fbbdc1d40ff6533c99190a35050565b6001546001600160a01b031681565b6005546001600160a01b031681565b610a0b83838385610af0565b505050565b60065481565b6000546001600160a01b031681565b81516020818401810180516009825292820194820194909420919093529091526000908152604090205481565b60025481565b6000546001600160a01b0316331480610a7d5750610a7d60023363ffffffff61100816565b610ace576040805162461bcd60e51b815260206004820152601960248201527f41646d696e206163636f756e7420697320726571756972656400000000000000604482015290519081900360640190fd5b600580546001600160a01b0319166001600160a01b0392909216919091179055565b600554604080517f31aaa74a0000000000000000000000000000000000000000000000000000000081526001600160a01b03868116600483015291513393909216916331aaa74a91602480820192602092909190829003018186803b158015610b5857600080fd5b505afa158015610b6c573d6000803e3d6000fd5b505050506040513d6020811015610b8257600080fd5b50516001600160a01b031614610bdf576040805162461bcd60e51b815260206004820152601860248201527f486f6c646572277320706172656e742072657175697265640000000000000000604482015290519081900360640190fd5b6060846001600160a01b03166395d89b416040518163ffffffff1660e01b815260040160006040518083038186803b158015610c1a57600080fd5b505afa158015610c2e573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526020811015610c5757600080fd5b8101908080516040519392919084640100000000821115610c7757600080fd5b908301906020820185811115610c8c57600080fd5b8251640100000000811182820188101715610ca657600080fd5b82525081516020918201929091019080838360005b83811015610cd3578181015183820152602001610cbb565b50505050905090810190601f168015610d005780820380516001836020036101000a031916815260200191505b5060405250505090506000610d8d846009846040518082805190602001908083835b60208310610d415780518252601f199092019160209182019101610d22565b51815160209384036101000a6000190180199092169116179052920194855250604080519485900382019094206001600160a01b038c16600090815291529290922054929150506111b0565b905080158015610da95750610da960068663ffffffff6110c516565b15610df457604080516001600160a01b0380891682528716602082015281517fef10ec58b248e4840da7c5adaaf9061b779ca8cb5f402c6007c71bd772771e5e929181900390910190a15b806009836040518082805190602001908083835b60208310610e275780518252601f199092019160209182019101610e08565b51815160209384036101000a60001901801990921691161790529201948552506040805194859003820185206001600160a01b038c81166000908152918452828220979097557fa9059cbb0000000000000000000000000000000000000000000000000000000086528987166004870152602486018b90529051958c169563a9059cbb95604480820196509394509092908390030190829087803b158015610ece57600080fd5b505af1158015610ee2573d6000803e3d6000fd5b505050506040513d6020811015610ef857600080fd5b5050505050505050565b600082820183811015610f5c576040805162461bcd60e51b815260206004820152601360248201527f526573756c747320696e206f766572666c6f7700000000000000000000000000604482015290519081900360640190fd5b9392505050565b60006001600160a01b038216610f7b575060006106a0565b6001600160a01b038216600090815260018401602052604081205460001901908112801590610faa5750835481125b15610fb95760009150506106a0565b5050815460019081018084556001600160a01b0383166000818152838601602090815260408083208590559382526002870190529190912080546001600160a01b031916909117905592915050565b6001600160a01b0381166000908152600183016020526040812054600019018181128015906110375750835481125b949350505050565b60008082121580156110515750825482125b6110a2576040805162461bcd60e51b815260206004820152601860248201527f496e646578206f757473696465206f6620626f756e64732e0000000000000000604482015290519081900360640190fd5b50600101600090815260029190910160205260409020546001600160a01b031690565b6001600160a01b038116600090815260018084016020526040822054908112806110ef5750835481135b156110fe5760009150506106a0565b835481121561116557835460009081526002850160208181526040808420546001600160a01b03168085526001890183528185208690558585529290915280832080546001600160a01b031990811690931790558654835290912080549091169055611184565b6000818152600285016020526040902080546001600160a01b03191690555b50506001600160a01b031660009081526001828101602052604082209190915581546000190190915590565b600082821115611207576040805162461bcd60e51b815260206004820152601460248201527f526573756c747320696e20756e646572666c6f77000000000000000000000000604482015290519081900360640190fd5b5090039056fe4e6577204f776e65722063616e6e6f74206265207468652063757272656e74206f776e6572a265627a7a72315820dfdd50b90bcb8cd84a529735b975b8488857ea0579de5f1f98f47859ff2261f164736f6c634300050b0032`

// DeployEscrow deploys a new Ethereum contract, binding an instance of Escrow to it.
func DeployEscrow(auth *bind.TransactOpts, backend bind.ContractBackend, r common.Address) (common.Address, *types.Transaction, *Escrow, error) {
	parsed, err := abi.JSON(strings.NewReader(EscrowABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(EscrowBin), backend, r)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Escrow{EscrowCaller: EscrowCaller{contract: contract}, EscrowTransactor: EscrowTransactor{contract: contract}, EscrowFilterer: EscrowFilterer{contract: contract}}, nil
}

// Escrow is an auto generated Go binding around an Ethereum contract.
type Escrow struct {
	EscrowCaller     // Read-only binding to the contract
	EscrowTransactor // Write-only binding to the contract
	EscrowFilterer   // Log filterer for contract events
}

// EscrowCaller is an auto generated read-only Go binding around an Ethereum contract.
type EscrowCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EscrowTransactor is an auto generated write-only Go binding around an Ethereum contract.
type EscrowTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EscrowFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type EscrowFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EscrowSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type EscrowSession struct {
	Contract     *Escrow           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// EscrowCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type EscrowCallerSession struct {
	Contract *EscrowCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// EscrowTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type EscrowTransactorSession struct {
	Contract     *EscrowTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// EscrowRaw is an auto generated low-level Go binding around an Ethereum contract.
type EscrowRaw struct {
	Contract *Escrow // Generic contract binding to access the raw methods on
}

// EscrowCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type EscrowCallerRaw struct {
	Contract *EscrowCaller // Generic read-only contract binding to access the raw methods on
}

// EscrowTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type EscrowTransactorRaw struct {
	Contract *EscrowTransactor // Generic write-only contract binding to access the raw methods on
}

// NewEscrow creates a new instance of Escrow, bound to a specific deployed contract.
func NewEscrow(address common.Address, backend bind.ContractBackend) (*Escrow, error) {
	contract, err := bindEscrow(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Escrow{EscrowCaller: EscrowCaller{contract: contract}, EscrowTransactor: EscrowTransactor{contract: contract}, EscrowFilterer: EscrowFilterer{contract: contract}}, nil
}

// NewEscrowCaller creates a new read-only instance of Escrow, bound to a specific deployed contract.
func NewEscrowCaller(address common.Address, caller bind.ContractCaller) (*EscrowCaller, error) {
	contract, err := bindEscrow(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &EscrowCaller{contract: contract}, nil
}

// NewEscrowTransactor creates a new write-only instance of Escrow, bound to a specific deployed contract.
func NewEscrowTransactor(address common.Address, transactor bind.ContractTransactor) (*EscrowTransactor, error) {
	contract, err := bindEscrow(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &EscrowTransactor{contract: contract}, nil
}

// NewEscrowFilterer creates a new log filterer instance of Escrow, bound to a specific deployed contract.
func NewEscrowFilterer(address common.Address, filterer bind.ContractFilterer) (*EscrowFilterer, error) {
	contract, err := bindEscrow(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &EscrowFilterer{contract: contract}, nil
}

// bindEscrow binds a generic wrapper to an already deployed contract.
func bindEscrow(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(EscrowABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Escrow *EscrowRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Escrow.Contract.EscrowCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Escrow *EscrowRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Escrow.Contract.EscrowTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Escrow *EscrowRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Escrow.Contract.EscrowTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Escrow *EscrowCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Escrow.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Escrow *EscrowTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Escrow.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Escrow *EscrowTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Escrow.Contract.contract.Transact(opts, method, params...)
}

// ZEROADDRESS is a free data retrieval call binding the contract method 0x538ba4f9.
//
// Solidity: function ZERO_ADDRESS() constant returns(address)
func (_Escrow *EscrowCaller) ZEROADDRESS(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Escrow.contract.Call(opts, out, "ZERO_ADDRESS")
	return *ret0, err
}

// ZEROADDRESS is a free data retrieval call binding the contract method 0x538ba4f9.
//
// Solidity: function ZERO_ADDRESS() constant returns(address)
func (_Escrow *EscrowSession) ZEROADDRESS() (common.Address, error) {
	return _Escrow.Contract.ZEROADDRESS(&_Escrow.CallOpts)
}

// ZEROADDRESS is a free data retrieval call binding the contract method 0x538ba4f9.
//
// Solidity: function ZERO_ADDRESS() constant returns(address)
func (_Escrow *EscrowCallerSession) ZEROADDRESS() (common.Address, error) {
	return _Escrow.Contract.ZEROADDRESS(&_Escrow.CallOpts)
}

// AdminAt is a free data retrieval call binding the contract method 0x2bdbc56f.
//
// Solidity: function adminAt(index int256) constant returns(address)
func (_Escrow *EscrowCaller) AdminAt(opts *bind.CallOpts, index *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Escrow.contract.Call(opts, out, "adminAt", index)
	return *ret0, err
}

// AdminAt is a free data retrieval call binding the contract method 0x2bdbc56f.
//
// Solidity: function adminAt(index int256) constant returns(address)
func (_Escrow *EscrowSession) AdminAt(index *big.Int) (common.Address, error) {
	return _Escrow.Contract.AdminAt(&_Escrow.CallOpts, index)
}

// AdminAt is a free data retrieval call binding the contract method 0x2bdbc56f.
//
// Solidity: function adminAt(index int256) constant returns(address)
func (_Escrow *EscrowCallerSession) AdminAt(index *big.Int) (common.Address, error) {
	return _Escrow.Contract.AdminAt(&_Escrow.CallOpts, index)
}

// Admins is a free data retrieval call binding the contract method 0xa5de3619.
//
// Solidity: function admins() constant returns(count int256)
func (_Escrow *EscrowCaller) Admins(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Escrow.contract.Call(opts, out, "admins")
	return *ret0, err
}

// Admins is a free data retrieval call binding the contract method 0xa5de3619.
//
// Solidity: function admins() constant returns(count int256)
func (_Escrow *EscrowSession) Admins() (*big.Int, error) {
	return _Escrow.Contract.Admins(&_Escrow.CallOpts)
}

// Admins is a free data retrieval call binding the contract method 0xa5de3619.
//
// Solidity: function admins() constant returns(count int256)
func (_Escrow *EscrowCallerSession) Admins() (*big.Int, error) {
	return _Escrow.Contract.Admins(&_Escrow.CallOpts)
}

// Holders is a free data retrieval call binding the contract method 0x8188f71c.
//
// Solidity: function holders() constant returns(count int256)
func (_Escrow *EscrowCaller) Holders(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Escrow.contract.Call(opts, out, "holders")
	return *ret0, err
}

// Holders is a free data retrieval call binding the contract method 0x8188f71c.
//
// Solidity: function holders() constant returns(count int256)
func (_Escrow *EscrowSession) Holders() (*big.Int, error) {
	return _Escrow.Contract.Holders(&_Escrow.CallOpts)
}

// Holders is a free data retrieval call binding the contract method 0x8188f71c.
//
// Solidity: function holders() constant returns(count int256)
func (_Escrow *EscrowCallerSession) Holders() (*big.Int, error) {
	return _Escrow.Contract.Holders(&_Escrow.CallOpts)
}

// IsAdmin is a free data retrieval call binding the contract method 0x24d7806c.
//
// Solidity: function isAdmin(addr address) constant returns(bool)
func (_Escrow *EscrowCaller) IsAdmin(opts *bind.CallOpts, addr common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Escrow.contract.Call(opts, out, "isAdmin", addr)
	return *ret0, err
}

// IsAdmin is a free data retrieval call binding the contract method 0x24d7806c.
//
// Solidity: function isAdmin(addr address) constant returns(bool)
func (_Escrow *EscrowSession) IsAdmin(addr common.Address) (bool, error) {
	return _Escrow.Contract.IsAdmin(&_Escrow.CallOpts, addr)
}

// IsAdmin is a free data retrieval call binding the contract method 0x24d7806c.
//
// Solidity: function isAdmin(addr address) constant returns(bool)
func (_Escrow *EscrowCallerSession) IsAdmin(addr common.Address) (bool, error) {
	return _Escrow.Contract.IsAdmin(&_Escrow.CallOpts, addr)
}

// Ledger is a free data retrieval call binding the contract method 0x91fe9b6a.
//
// Solidity: function ledger( string,  address) constant returns(uint256)
func (_Escrow *EscrowCaller) Ledger(opts *bind.CallOpts, arg0 string, arg1 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Escrow.contract.Call(opts, out, "ledger", arg0, arg1)
	return *ret0, err
}

// Ledger is a free data retrieval call binding the contract method 0x91fe9b6a.
//
// Solidity: function ledger( string,  address) constant returns(uint256)
func (_Escrow *EscrowSession) Ledger(arg0 string, arg1 common.Address) (*big.Int, error) {
	return _Escrow.Contract.Ledger(&_Escrow.CallOpts, arg0, arg1)
}

// Ledger is a free data retrieval call binding the contract method 0x91fe9b6a.
//
// Solidity: function ledger( string,  address) constant returns(uint256)
func (_Escrow *EscrowCallerSession) Ledger(arg0 string, arg1 common.Address) (*big.Int, error) {
	return _Escrow.Contract.Ledger(&_Escrow.CallOpts, arg0, arg1)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Escrow *EscrowCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Escrow.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Escrow *EscrowSession) Owner() (common.Address, error) {
	return _Escrow.Contract.Owner(&_Escrow.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Escrow *EscrowCallerSession) Owner() (common.Address, error) {
	return _Escrow.Contract.Owner(&_Escrow.CallOpts)
}

// Registry is a free data retrieval call binding the contract method 0x7b103999.
//
// Solidity: function registry() constant returns(address)
func (_Escrow *EscrowCaller) Registry(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Escrow.contract.Call(opts, out, "registry")
	return *ret0, err
}

// Registry is a free data retrieval call binding the contract method 0x7b103999.
//
// Solidity: function registry() constant returns(address)
func (_Escrow *EscrowSession) Registry() (common.Address, error) {
	return _Escrow.Contract.Registry(&_Escrow.CallOpts)
}

// Registry is a free data retrieval call binding the contract method 0x7b103999.
//
// Solidity: function registry() constant returns(address)
func (_Escrow *EscrowCallerSession) Registry() (common.Address, error) {
	return _Escrow.Contract.Registry(&_Escrow.CallOpts)
}

// Accept is a paid mutator transaction binding the contract method 0x101e7075.
//
// Solidity: function accept(token address, holder address, quantity uint256, grantee address) returns()
func (_Escrow *EscrowTransactor) Accept(opts *bind.TransactOpts, token common.Address, holder common.Address, quantity *big.Int, grantee common.Address) (*types.Transaction, error) {
	return _Escrow.contract.Transact(opts, "accept", token, holder, quantity, grantee)
}

// Accept is a paid mutator transaction binding the contract method 0x101e7075.
//
// Solidity: function accept(token address, holder address, quantity uint256, grantee address) returns()
func (_Escrow *EscrowSession) Accept(token common.Address, holder common.Address, quantity *big.Int, grantee common.Address) (*types.Transaction, error) {
	return _Escrow.Contract.Accept(&_Escrow.TransactOpts, token, holder, quantity, grantee)
}

// Accept is a paid mutator transaction binding the contract method 0x101e7075.
//
// Solidity: function accept(token address, holder address, quantity uint256, grantee address) returns()
func (_Escrow *EscrowTransactorSession) Accept(token common.Address, holder common.Address, quantity *big.Int, grantee common.Address) (*types.Transaction, error) {
	return _Escrow.Contract.Accept(&_Escrow.TransactOpts, token, holder, quantity, grantee)
}

// Record is a paid mutator transaction binding the contract method 0x172a93fb.
//
// Solidity: function record(token address, holder address, quantity uint256) returns()
func (_Escrow *EscrowTransactor) Record(opts *bind.TransactOpts, token common.Address, holder common.Address, quantity *big.Int) (*types.Transaction, error) {
	return _Escrow.contract.Transact(opts, "record", token, holder, quantity)
}

// Record is a paid mutator transaction binding the contract method 0x172a93fb.
//
// Solidity: function record(token address, holder address, quantity uint256) returns()
func (_Escrow *EscrowSession) Record(token common.Address, holder common.Address, quantity *big.Int) (*types.Transaction, error) {
	return _Escrow.Contract.Record(&_Escrow.TransactOpts, token, holder, quantity)
}

// Record is a paid mutator transaction binding the contract method 0x172a93fb.
//
// Solidity: function record(token address, holder address, quantity uint256) returns()
func (_Escrow *EscrowTransactorSession) Record(token common.Address, holder common.Address, quantity *big.Int) (*types.Transaction, error) {
	return _Escrow.Contract.Record(&_Escrow.TransactOpts, token, holder, quantity)
}

// Reject is a paid mutator transaction binding the contract method 0x813bedf2.
//
// Solidity: function reject(token address, holder address, quantity uint256) returns()
func (_Escrow *EscrowTransactor) Reject(opts *bind.TransactOpts, token common.Address, holder common.Address, quantity *big.Int) (*types.Transaction, error) {
	return _Escrow.contract.Transact(opts, "reject", token, holder, quantity)
}

// Reject is a paid mutator transaction binding the contract method 0x813bedf2.
//
// Solidity: function reject(token address, holder address, quantity uint256) returns()
func (_Escrow *EscrowSession) Reject(token common.Address, holder common.Address, quantity *big.Int) (*types.Transaction, error) {
	return _Escrow.Contract.Reject(&_Escrow.TransactOpts, token, holder, quantity)
}

// Reject is a paid mutator transaction binding the contract method 0x813bedf2.
//
// Solidity: function reject(token address, holder address, quantity uint256) returns()
func (_Escrow *EscrowTransactorSession) Reject(token common.Address, holder common.Address, quantity *big.Int) (*types.Transaction, error) {
	return _Escrow.Contract.Reject(&_Escrow.TransactOpts, token, holder, quantity)
}

// SetAdmin is a paid mutator transaction binding the contract method 0x4b0bddd2.
//
// Solidity: function setAdmin(addr address, add bool) returns()
func (_Escrow *EscrowTransactor) SetAdmin(opts *bind.TransactOpts, addr common.Address, add bool) (*types.Transaction, error) {
	return _Escrow.contract.Transact(opts, "setAdmin", addr, add)
}

// SetAdmin is a paid mutator transaction binding the contract method 0x4b0bddd2.
//
// Solidity: function setAdmin(addr address, add bool) returns()
func (_Escrow *EscrowSession) SetAdmin(addr common.Address, add bool) (*types.Transaction, error) {
	return _Escrow.Contract.SetAdmin(&_Escrow.TransactOpts, addr, add)
}

// SetAdmin is a paid mutator transaction binding the contract method 0x4b0bddd2.
//
// Solidity: function setAdmin(addr address, add bool) returns()
func (_Escrow *EscrowTransactorSession) SetAdmin(addr common.Address, add bool) (*types.Transaction, error) {
	return _Escrow.Contract.SetAdmin(&_Escrow.TransactOpts, addr, add)
}

// SetRegistry is a paid mutator transaction binding the contract method 0xa91ee0dc.
//
// Solidity: function setRegistry(r address) returns()
func (_Escrow *EscrowTransactor) SetRegistry(opts *bind.TransactOpts, r common.Address) (*types.Transaction, error) {
	return _Escrow.contract.Transact(opts, "setRegistry", r)
}

// SetRegistry is a paid mutator transaction binding the contract method 0xa91ee0dc.
//
// Solidity: function setRegistry(r address) returns()
func (_Escrow *EscrowSession) SetRegistry(r common.Address) (*types.Transaction, error) {
	return _Escrow.Contract.SetRegistry(&_Escrow.TransactOpts, r)
}

// SetRegistry is a paid mutator transaction binding the contract method 0xa91ee0dc.
//
// Solidity: function setRegistry(r address) returns()
func (_Escrow *EscrowTransactorSession) SetRegistry(r common.Address) (*types.Transaction, error) {
	return _Escrow.Contract.SetRegistry(&_Escrow.TransactOpts, r)
}

// TransferOwner is a paid mutator transaction binding the contract method 0x4fb2e45d.
//
// Solidity: function transferOwner(newOwner address) returns()
func (_Escrow *EscrowTransactor) TransferOwner(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Escrow.contract.Transact(opts, "transferOwner", newOwner)
}

// TransferOwner is a paid mutator transaction binding the contract method 0x4fb2e45d.
//
// Solidity: function transferOwner(newOwner address) returns()
func (_Escrow *EscrowSession) TransferOwner(newOwner common.Address) (*types.Transaction, error) {
	return _Escrow.Contract.TransferOwner(&_Escrow.TransactOpts, newOwner)
}

// TransferOwner is a paid mutator transaction binding the contract method 0x4fb2e45d.
//
// Solidity: function transferOwner(newOwner address) returns()
func (_Escrow *EscrowTransactorSession) TransferOwner(newOwner common.Address) (*types.Transaction, error) {
	return _Escrow.Contract.TransferOwner(&_Escrow.TransactOpts, newOwner)
}

// EscrowAdminAddedIterator is returned from FilterAdminAdded and is used to iterate over the raw logs and unpacked data for AdminAdded events raised by the Escrow contract.
type EscrowAdminAddedIterator struct {
	Event *EscrowAdminAdded // Event containing the contract specifics and raw log

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
func (it *EscrowAdminAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EscrowAdminAdded)
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
		it.Event = new(EscrowAdminAdded)
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
func (it *EscrowAdminAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EscrowAdminAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EscrowAdminAdded represents a AdminAdded event raised by the Escrow contract.
type EscrowAdminAdded struct {
	Admin common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterAdminAdded is a free log retrieval operation binding the contract event 0x44d6d25963f097ad14f29f06854a01f575648a1ef82f30e562ccd3889717e339.
//
// Solidity: e AdminAdded(admin indexed address)
func (_Escrow *EscrowFilterer) FilterAdminAdded(opts *bind.FilterOpts, admin []common.Address) (*EscrowAdminAddedIterator, error) {

	var adminRule []interface{}
	for _, adminItem := range admin {
		adminRule = append(adminRule, adminItem)
	}

	logs, sub, err := _Escrow.contract.FilterLogs(opts, "AdminAdded", adminRule)
	if err != nil {
		return nil, err
	}
	return &EscrowAdminAddedIterator{contract: _Escrow.contract, event: "AdminAdded", logs: logs, sub: sub}, nil
}

// WatchAdminAdded is a free log subscription operation binding the contract event 0x44d6d25963f097ad14f29f06854a01f575648a1ef82f30e562ccd3889717e339.
//
// Solidity: e AdminAdded(admin indexed address)
func (_Escrow *EscrowFilterer) WatchAdminAdded(opts *bind.WatchOpts, sink chan<- *EscrowAdminAdded, admin []common.Address) (event.Subscription, error) {

	var adminRule []interface{}
	for _, adminItem := range admin {
		adminRule = append(adminRule, adminItem)
	}

	logs, sub, err := _Escrow.contract.WatchLogs(opts, "AdminAdded", adminRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EscrowAdminAdded)
				if err := _Escrow.contract.UnpackLog(event, "AdminAdded", log); err != nil {
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

// EscrowAdminRemovedIterator is returned from FilterAdminRemoved and is used to iterate over the raw logs and unpacked data for AdminRemoved events raised by the Escrow contract.
type EscrowAdminRemovedIterator struct {
	Event *EscrowAdminRemoved // Event containing the contract specifics and raw log

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
func (it *EscrowAdminRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EscrowAdminRemoved)
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
		it.Event = new(EscrowAdminRemoved)
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
func (it *EscrowAdminRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EscrowAdminRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EscrowAdminRemoved represents a AdminRemoved event raised by the Escrow contract.
type EscrowAdminRemoved struct {
	Admin common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterAdminRemoved is a free log retrieval operation binding the contract event 0xa3b62bc36326052d97ea62d63c3d60308ed4c3ea8ac079dd8499f1e9c4f80c0f.
//
// Solidity: e AdminRemoved(admin indexed address)
func (_Escrow *EscrowFilterer) FilterAdminRemoved(opts *bind.FilterOpts, admin []common.Address) (*EscrowAdminRemovedIterator, error) {

	var adminRule []interface{}
	for _, adminItem := range admin {
		adminRule = append(adminRule, adminItem)
	}

	logs, sub, err := _Escrow.contract.FilterLogs(opts, "AdminRemoved", adminRule)
	if err != nil {
		return nil, err
	}
	return &EscrowAdminRemovedIterator{contract: _Escrow.contract, event: "AdminRemoved", logs: logs, sub: sub}, nil
}

// WatchAdminRemoved is a free log subscription operation binding the contract event 0xa3b62bc36326052d97ea62d63c3d60308ed4c3ea8ac079dd8499f1e9c4f80c0f.
//
// Solidity: e AdminRemoved(admin indexed address)
func (_Escrow *EscrowFilterer) WatchAdminRemoved(opts *bind.WatchOpts, sink chan<- *EscrowAdminRemoved, admin []common.Address) (event.Subscription, error) {

	var adminRule []interface{}
	for _, adminItem := range admin {
		adminRule = append(adminRule, adminItem)
	}

	logs, sub, err := _Escrow.contract.WatchLogs(opts, "AdminRemoved", adminRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EscrowAdminRemoved)
				if err := _Escrow.contract.UnpackLog(event, "AdminRemoved", log); err != nil {
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

// EscrowEscrowAddedIterator is returned from FilterEscrowAdded and is used to iterate over the raw logs and unpacked data for EscrowAdded events raised by the Escrow contract.
type EscrowEscrowAddedIterator struct {
	Event *EscrowEscrowAdded // Event containing the contract specifics and raw log

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
func (it *EscrowEscrowAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EscrowEscrowAdded)
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
		it.Event = new(EscrowEscrowAdded)
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
func (it *EscrowEscrowAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EscrowEscrowAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EscrowEscrowAdded represents a EscrowAdded event raised by the Escrow contract.
type EscrowEscrowAdded struct {
	Token common.Address
	Addr  common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterEscrowAdded is a free log retrieval operation binding the contract event 0x77ec0a1f539783369bfc02c19452ee9903f7f59cb8f191a0f0937e5e61e8efe5.
//
// Solidity: e EscrowAdded(token address, addr address)
func (_Escrow *EscrowFilterer) FilterEscrowAdded(opts *bind.FilterOpts) (*EscrowEscrowAddedIterator, error) {

	logs, sub, err := _Escrow.contract.FilterLogs(opts, "EscrowAdded")
	if err != nil {
		return nil, err
	}
	return &EscrowEscrowAddedIterator{contract: _Escrow.contract, event: "EscrowAdded", logs: logs, sub: sub}, nil
}

// WatchEscrowAdded is a free log subscription operation binding the contract event 0x77ec0a1f539783369bfc02c19452ee9903f7f59cb8f191a0f0937e5e61e8efe5.
//
// Solidity: e EscrowAdded(token address, addr address)
func (_Escrow *EscrowFilterer) WatchEscrowAdded(opts *bind.WatchOpts, sink chan<- *EscrowEscrowAdded) (event.Subscription, error) {

	logs, sub, err := _Escrow.contract.WatchLogs(opts, "EscrowAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EscrowEscrowAdded)
				if err := _Escrow.contract.UnpackLog(event, "EscrowAdded", log); err != nil {
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

// EscrowEscrowRemovedIterator is returned from FilterEscrowRemoved and is used to iterate over the raw logs and unpacked data for EscrowRemoved events raised by the Escrow contract.
type EscrowEscrowRemovedIterator struct {
	Event *EscrowEscrowRemoved // Event containing the contract specifics and raw log

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
func (it *EscrowEscrowRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EscrowEscrowRemoved)
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
		it.Event = new(EscrowEscrowRemoved)
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
func (it *EscrowEscrowRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EscrowEscrowRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EscrowEscrowRemoved represents a EscrowRemoved event raised by the Escrow contract.
type EscrowEscrowRemoved struct {
	Token common.Address
	Addr  common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterEscrowRemoved is a free log retrieval operation binding the contract event 0xef10ec58b248e4840da7c5adaaf9061b779ca8cb5f402c6007c71bd772771e5e.
//
// Solidity: e EscrowRemoved(token address, addr address)
func (_Escrow *EscrowFilterer) FilterEscrowRemoved(opts *bind.FilterOpts) (*EscrowEscrowRemovedIterator, error) {

	logs, sub, err := _Escrow.contract.FilterLogs(opts, "EscrowRemoved")
	if err != nil {
		return nil, err
	}
	return &EscrowEscrowRemovedIterator{contract: _Escrow.contract, event: "EscrowRemoved", logs: logs, sub: sub}, nil
}

// WatchEscrowRemoved is a free log subscription operation binding the contract event 0xef10ec58b248e4840da7c5adaaf9061b779ca8cb5f402c6007c71bd772771e5e.
//
// Solidity: e EscrowRemoved(token address, addr address)
func (_Escrow *EscrowFilterer) WatchEscrowRemoved(opts *bind.WatchOpts, sink chan<- *EscrowEscrowRemoved) (event.Subscription, error) {

	logs, sub, err := _Escrow.contract.WatchLogs(opts, "EscrowRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EscrowEscrowRemoved)
				if err := _Escrow.contract.UnpackLog(event, "EscrowRemoved", log); err != nil {
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

// EscrowOwnerTransferredIterator is returned from FilterOwnerTransferred and is used to iterate over the raw logs and unpacked data for OwnerTransferred events raised by the Escrow contract.
type EscrowOwnerTransferredIterator struct {
	Event *EscrowOwnerTransferred // Event containing the contract specifics and raw log

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
func (it *EscrowOwnerTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EscrowOwnerTransferred)
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
		it.Event = new(EscrowOwnerTransferred)
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
func (it *EscrowOwnerTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EscrowOwnerTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EscrowOwnerTransferred represents a OwnerTransferred event raised by the Escrow contract.
type EscrowOwnerTransferred struct {
	OldOwner common.Address
	NewOwner common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOwnerTransferred is a free log retrieval operation binding the contract event 0x8934ce4adea8d9ce0d714d2c22b86790e41b7731c84b926fbbdc1d40ff6533c9.
//
// Solidity: e OwnerTransferred(oldOwner indexed address, newOwner indexed address)
func (_Escrow *EscrowFilterer) FilterOwnerTransferred(opts *bind.FilterOpts, oldOwner []common.Address, newOwner []common.Address) (*EscrowOwnerTransferredIterator, error) {

	var oldOwnerRule []interface{}
	for _, oldOwnerItem := range oldOwner {
		oldOwnerRule = append(oldOwnerRule, oldOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Escrow.contract.FilterLogs(opts, "OwnerTransferred", oldOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &EscrowOwnerTransferredIterator{contract: _Escrow.contract, event: "OwnerTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnerTransferred is a free log subscription operation binding the contract event 0x8934ce4adea8d9ce0d714d2c22b86790e41b7731c84b926fbbdc1d40ff6533c9.
//
// Solidity: e OwnerTransferred(oldOwner indexed address, newOwner indexed address)
func (_Escrow *EscrowFilterer) WatchOwnerTransferred(opts *bind.WatchOpts, sink chan<- *EscrowOwnerTransferred, oldOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var oldOwnerRule []interface{}
	for _, oldOwnerItem := range oldOwner {
		oldOwnerRule = append(oldOwnerRule, oldOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Escrow.contract.WatchLogs(opts, "OwnerTransferred", oldOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EscrowOwnerTransferred)
				if err := _Escrow.contract.UnpackLog(event, "OwnerTransferred", log); err != nil {
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
