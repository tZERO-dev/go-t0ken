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
const EscrowABI = "[{\"inputs\":[{\"internalType\":\"contractIRegistry\",\"name\":\"r\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"}],\"name\":\"AdminAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"}],\"name\":\"AdminRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"contractIT0ken\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"grantor\",\"type\":\"address\"}],\"name\":\"EscrowAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"contractIT0ken\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"grantor\",\"type\":\"address\"}],\"name\":\"EscrowRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnerTransferred\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[],\"name\":\"ZERO_ADDRESS\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"contractIT0ken\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"grantor\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"quantity\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"grantee\",\"type\":\"address\"}],\"name\":\"accept\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"int256\",\"name\":\"index\",\"type\":\"int256\"}],\"name\":\"adminAt\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"admins\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"count\",\"type\":\"int256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"isAdmin\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"ledger\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"contractIT0ken\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"grantor\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"quantity\",\"type\":\"uint256\"}],\"name\":\"record\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"registry\",\"outputs\":[{\"internalType\":\"contractIRegistry\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"contractIT0ken\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"grantor\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"quantity\",\"type\":\"uint256\"}],\"name\":\"reject\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"add\",\"type\":\"bool\"}],\"name\":\"setAdmin\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"contractIRegistry\",\"name\":\"r\",\"type\":\"address\"}],\"name\":\"setRegistry\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwner\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// EscrowBin is the compiled bytecode used for deploying new contracts.
const EscrowBin = `6080604052600180546001600160a01b031916905534801561002057600080fd5b506040516111853803806111858339818101604052602081101561004357600080fd5b5051600080546001600160a01b03199081163317909155600580546001600160a01b0390931692909116919091179055611103806100826000396000f3fe608060405234801561001057600080fd5b50600436106100df5760003560e01c8063538ba4f91161008c5780638da5cb5b116100665780638da5cb5b1461026557806396dcdbf01461026d578063a5de3619146102ad578063a91ee0dc146102b5576100df565b8063538ba4f91461021f5780637b10399914610227578063813bedf21461022f576100df565b80632bdbc56f116100bd5780632bdbc56f146101925780634b0bddd2146101cb5780634fb2e45d146101f9576100df565b8063101e7075146100e4578063172a93fb1461012257806324d7806c14610158575b600080fd5b610120600480360360808110156100fa57600080fd5b506001600160a01b038135811691602081013582169160408201359160600135166102db565b005b6101206004803603606081101561013857600080fd5b506001600160a01b038135811691602081013590911690604001356102ed565b61017e6004803603602081101561016e57600080fd5b50356001600160a01b0316610583565b604080519115158252519081900360200190f35b6101af600480360360208110156101a857600080fd5b50356105b2565b604080516001600160a01b039092168252519081900360200190f35b610120600480360360408110156101e157600080fd5b506001600160a01b03813516906020013515156105c5565b6101206004803603602081101561020f57600080fd5b50356001600160a01b031661077b565b6101af6108ed565b6101af6108fc565b6101206004803603606081101561024557600080fd5b506001600160a01b0381358116916020810135909116906040013561090b565b6101af61091c565b61029b6004803603604081101561028357600080fd5b506001600160a01b038135811691602001351661092b565b60408051918252519081900360200190f35b61029b610948565b610120600480360360208110156102cb57600080fd5b50356001600160a01b031661094e565b6102e7848484846109e6565b50505050565b6102f633610583565b610347576040805162461bcd60e51b815260206004820152601f60248201527f41646d696e20726571756972656420746f207265636f726420657363726f7700604482015290519081900360640190fd5b6001600160a01b0380841660009081526006602090815260408083209386168352929052908120548491610381828563ffffffff610d9e16565b9050811580156103915750600081115b1561055357846001600160a01b0316866001600160a01b0316876001600160a01b03166395d89b416040518163ffffffff1660e01b815260040160006040518083038186803b1580156103e357600080fd5b505afa1580156103f7573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052602081101561042057600080fd5b810190808051604051939291908464010000000082111561044057600080fd5b90830190602082018581111561045557600080fd5b825164010000000081118282018810171561046f57600080fd5b82525081516020918201929091019080838360005b8381101561049c578181015183820152602001610484565b50505050905090810190601f1680156104c95780820380516001836020036101000a031916815260200191505b506040525050506040518082805190602001908083835b602083106104ff5780518252601f1990920191602091820191016104e0565b5181516020939093036101000a60001901801990911692169190911790526040519201829003822093507f9b64b4e41c9d5860b8271685f71d54a08315c0b04bf02aab5e4c75eaa74cc21c92506000919050a45b6001600160a01b039283166000908152600660209081526040808320979095168252959095529190932055505050565b600080546001600160a01b03838116911614806105ac57506105ac60028363ffffffff610dff16565b92915050565b60006105ac60028363ffffffff610e3616565b6000546001600160a01b03163314806105ee57506001546000546001600160a01b039081169116145b61063f576040805162461bcd60e51b815260206004820152601960248201527f4f776e6572206163636f756e7420697320726571756972656400000000000000604482015290519081900360640190fd5b80156106e05761065660028363ffffffff610ebc16565b6106a7576040805162461bcd60e51b815260206004820152601360248201527f556e61626c6520746f206164642061646d696e00000000000000000000000000604482015290519081900360640190fd5b6040516001600160a01b038316907f44d6d25963f097ad14f29f06854a01f575648a1ef82f30e562ccd3889717e33990600090a2610777565b6106f160028363ffffffff610f6116565b610742576040805162461bcd60e51b815260206004820152601660248201527f556e61626c6520746f2072656d6f76652061646d696e00000000000000000000604482015290519081900360640190fd5b6040516001600160a01b038316907fa3b62bc36326052d97ea62d63c3d60308ed4c3ea8ac079dd8499f1e9c4f80c0f90600090a25b5050565b6000546001600160a01b03163314806107a457506001546000546001600160a01b039081169116145b6107f5576040805162461bcd60e51b815260206004820152601960248201527f4f776e6572206163636f756e7420697320726571756972656400000000000000604482015290519081900360640190fd5b6000546001600160a01b03828116911614156108425760405162461bcd60e51b81526004018080602001828103825260258152602001806110aa6025913960400191505060405180910390fd5b6001600160a01b03811661089d576040805162461bcd60e51b815260206004820181905260248201527f4e6577204f776e65722063616e6e6f74206265207a65726f2061646472657373604482015290519081900360640190fd5b600080546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8934ce4adea8d9ce0d714d2c22b86790e41b7731c84b926fbbdc1d40ff6533c99190a35050565b6001546001600160a01b031681565b6005546001600160a01b031681565b610917838383856109e6565b505050565b6000546001600160a01b031681565b600660209081526000928352604080842090915290825290205481565b60025481565b6000546001600160a01b0316331480610973575061097360023363ffffffff610dff16565b6109c4576040805162461bcd60e51b815260206004820152601960248201527f41646d696e206163636f756e7420697320726571756972656400000000000000604482015290519081900360640190fd5b600580546001600160a01b0319166001600160a01b0392909216919091179055565b600554604080517f31aaa74a0000000000000000000000000000000000000000000000000000000081526001600160a01b03868116600483015291513393909216916331aaa74a91602480820192602092909190829003018186803b158015610a4e57600080fd5b505afa158015610a62573d6000803e3d6000fd5b505050506040513d6020811015610a7857600080fd5b50516001600160a01b031614610ad5576040805162461bcd60e51b815260206004820152601960248201527f4772616e746f72277320706172656e7420726571756972656400000000000000604482015290519081900360640190fd5b836001600160a01b031663a9059cbb82846040518363ffffffff1660e01b815260040180836001600160a01b03166001600160a01b0316815260200182815260200192505050602060405180830381600087803b158015610b3557600080fd5b505af1158015610b49573d6000803e3d6000fd5b505050506040513d6020811015610b5f57600080fd5b50506001600160a01b0380851660009081526006602090815260408083209387168352929052908120548591610b9b828663ffffffff61104c16565b9050600082118015610bab575080155b15610d6d57856001600160a01b0316876001600160a01b0316886001600160a01b03166395d89b416040518163ffffffff1660e01b815260040160006040518083038186803b158015610bfd57600080fd5b505afa158015610c11573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526020811015610c3a57600080fd5b8101908080516040519392919084640100000000821115610c5a57600080fd5b908301906020820185811115610c6f57600080fd5b8251640100000000811182820188101715610c8957600080fd5b82525081516020918201929091019080838360005b83811015610cb6578181015183820152602001610c9e565b50505050905090810190601f168015610ce35780820380516001836020036101000a031916815260200191505b506040525050506040518082805190602001908083835b60208310610d195780518252601f199092019160209182019101610cfa565b5181516020939093036101000a60001901801990911692169190911790526040519201829003822093507fd4903355a168bfba0a71e036316138be08ba1f745566878969512bb7554f2c3192506000919050a45b6001600160a01b03928316600090815260066020908152604080832098909516825296909652919094205550505050565b600082820183811015610df8576040805162461bcd60e51b815260206004820152601360248201527f526573756c747320696e206f766572666c6f7700000000000000000000000000604482015290519081900360640190fd5b9392505050565b6001600160a01b038116600090815260018301602052604081205460001901818112801590610e2e5750835481125b949350505050565b6000808212158015610e485750825482125b610e99576040805162461bcd60e51b815260206004820152601860248201527f496e646578206f757473696465206f6620626f756e64732e0000000000000000604482015290519081900360640190fd5b50600101600090815260029190910160205260409020546001600160a01b031690565b60006001600160a01b038216610ed4575060006105ac565b6001600160a01b038216600090815260018401602052604081205460001901908112801590610f035750835481125b15610f125760009150506105ac565b5050815460019081018084556001600160a01b0383166000818152838601602090815260408083208590559382526002870190529190912080546001600160a01b031916909117905592915050565b6001600160a01b03811660009081526001808401602052604082205490811280610f8b5750835481135b15610f9a5760009150506105ac565b835481121561100157835460009081526002850160208181526040808420546001600160a01b03168085526001890183528185208690558585529290915280832080546001600160a01b031990811690931790558654835290912080549091169055611020565b6000818152600285016020526040902080546001600160a01b03191690555b50506001600160a01b031660009081526001828101602052604082209190915581546000190190915590565b6000828211156110a3576040805162461bcd60e51b815260206004820152601460248201527f526573756c747320696e20756e646572666c6f77000000000000000000000000604482015290519081900360640190fd5b5090039056fe4e6577204f776e65722063616e6e6f74206265207468652063757272656e74206f776e6572a265627a7a7231582071de64da755edaba178c1552b39507b11d948d18ab9b7c6d45295fd2d5670a6064736f6c634300050c0032`

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

// Ledger is a free data retrieval call binding the contract method 0x96dcdbf0.
//
// Solidity: function ledger( address,  address) constant returns(uint256)
func (_Escrow *EscrowCaller) Ledger(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Escrow.contract.Call(opts, out, "ledger", arg0, arg1)
	return *ret0, err
}

// Ledger is a free data retrieval call binding the contract method 0x96dcdbf0.
//
// Solidity: function ledger( address,  address) constant returns(uint256)
func (_Escrow *EscrowSession) Ledger(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _Escrow.Contract.Ledger(&_Escrow.CallOpts, arg0, arg1)
}

// Ledger is a free data retrieval call binding the contract method 0x96dcdbf0.
//
// Solidity: function ledger( address,  address) constant returns(uint256)
func (_Escrow *EscrowCallerSession) Ledger(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
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
// Solidity: function accept(token address, grantor address, quantity uint256, grantee address) returns()
func (_Escrow *EscrowTransactor) Accept(opts *bind.TransactOpts, token common.Address, grantor common.Address, quantity *big.Int, grantee common.Address) (*types.Transaction, error) {
	return _Escrow.contract.Transact(opts, "accept", token, grantor, quantity, grantee)
}

// Accept is a paid mutator transaction binding the contract method 0x101e7075.
//
// Solidity: function accept(token address, grantor address, quantity uint256, grantee address) returns()
func (_Escrow *EscrowSession) Accept(token common.Address, grantor common.Address, quantity *big.Int, grantee common.Address) (*types.Transaction, error) {
	return _Escrow.Contract.Accept(&_Escrow.TransactOpts, token, grantor, quantity, grantee)
}

// Accept is a paid mutator transaction binding the contract method 0x101e7075.
//
// Solidity: function accept(token address, grantor address, quantity uint256, grantee address) returns()
func (_Escrow *EscrowTransactorSession) Accept(token common.Address, grantor common.Address, quantity *big.Int, grantee common.Address) (*types.Transaction, error) {
	return _Escrow.Contract.Accept(&_Escrow.TransactOpts, token, grantor, quantity, grantee)
}

// Record is a paid mutator transaction binding the contract method 0x172a93fb.
//
// Solidity: function record(token address, grantor address, quantity uint256) returns()
func (_Escrow *EscrowTransactor) Record(opts *bind.TransactOpts, token common.Address, grantor common.Address, quantity *big.Int) (*types.Transaction, error) {
	return _Escrow.contract.Transact(opts, "record", token, grantor, quantity)
}

// Record is a paid mutator transaction binding the contract method 0x172a93fb.
//
// Solidity: function record(token address, grantor address, quantity uint256) returns()
func (_Escrow *EscrowSession) Record(token common.Address, grantor common.Address, quantity *big.Int) (*types.Transaction, error) {
	return _Escrow.Contract.Record(&_Escrow.TransactOpts, token, grantor, quantity)
}

// Record is a paid mutator transaction binding the contract method 0x172a93fb.
//
// Solidity: function record(token address, grantor address, quantity uint256) returns()
func (_Escrow *EscrowTransactorSession) Record(token common.Address, grantor common.Address, quantity *big.Int) (*types.Transaction, error) {
	return _Escrow.Contract.Record(&_Escrow.TransactOpts, token, grantor, quantity)
}

// Reject is a paid mutator transaction binding the contract method 0x813bedf2.
//
// Solidity: function reject(token address, grantor address, quantity uint256) returns()
func (_Escrow *EscrowTransactor) Reject(opts *bind.TransactOpts, token common.Address, grantor common.Address, quantity *big.Int) (*types.Transaction, error) {
	return _Escrow.contract.Transact(opts, "reject", token, grantor, quantity)
}

// Reject is a paid mutator transaction binding the contract method 0x813bedf2.
//
// Solidity: function reject(token address, grantor address, quantity uint256) returns()
func (_Escrow *EscrowSession) Reject(token common.Address, grantor common.Address, quantity *big.Int) (*types.Transaction, error) {
	return _Escrow.Contract.Reject(&_Escrow.TransactOpts, token, grantor, quantity)
}

// Reject is a paid mutator transaction binding the contract method 0x813bedf2.
//
// Solidity: function reject(token address, grantor address, quantity uint256) returns()
func (_Escrow *EscrowTransactorSession) Reject(token common.Address, grantor common.Address, quantity *big.Int) (*types.Transaction, error) {
	return _Escrow.Contract.Reject(&_Escrow.TransactOpts, token, grantor, quantity)
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
	Symbol  common.Hash
	Token   common.Address
	Grantor common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterEscrowAdded is a free log retrieval operation binding the contract event 0x9b64b4e41c9d5860b8271685f71d54a08315c0b04bf02aab5e4c75eaa74cc21c.
//
// Solidity: e EscrowAdded(symbol indexed string, token indexed address, grantor indexed address)
func (_Escrow *EscrowFilterer) FilterEscrowAdded(opts *bind.FilterOpts, symbol []string, token []common.Address, grantor []common.Address) (*EscrowEscrowAddedIterator, error) {

	var symbolRule []interface{}
	for _, symbolItem := range symbol {
		symbolRule = append(symbolRule, symbolItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var grantorRule []interface{}
	for _, grantorItem := range grantor {
		grantorRule = append(grantorRule, grantorItem)
	}

	logs, sub, err := _Escrow.contract.FilterLogs(opts, "EscrowAdded", symbolRule, tokenRule, grantorRule)
	if err != nil {
		return nil, err
	}
	return &EscrowEscrowAddedIterator{contract: _Escrow.contract, event: "EscrowAdded", logs: logs, sub: sub}, nil
}

// WatchEscrowAdded is a free log subscription operation binding the contract event 0x9b64b4e41c9d5860b8271685f71d54a08315c0b04bf02aab5e4c75eaa74cc21c.
//
// Solidity: e EscrowAdded(symbol indexed string, token indexed address, grantor indexed address)
func (_Escrow *EscrowFilterer) WatchEscrowAdded(opts *bind.WatchOpts, sink chan<- *EscrowEscrowAdded, symbol []string, token []common.Address, grantor []common.Address) (event.Subscription, error) {

	var symbolRule []interface{}
	for _, symbolItem := range symbol {
		symbolRule = append(symbolRule, symbolItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var grantorRule []interface{}
	for _, grantorItem := range grantor {
		grantorRule = append(grantorRule, grantorItem)
	}

	logs, sub, err := _Escrow.contract.WatchLogs(opts, "EscrowAdded", symbolRule, tokenRule, grantorRule)
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
	Symbol  common.Hash
	Token   common.Address
	Grantor common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterEscrowRemoved is a free log retrieval operation binding the contract event 0xd4903355a168bfba0a71e036316138be08ba1f745566878969512bb7554f2c31.
//
// Solidity: e EscrowRemoved(symbol indexed string, token indexed address, grantor indexed address)
func (_Escrow *EscrowFilterer) FilterEscrowRemoved(opts *bind.FilterOpts, symbol []string, token []common.Address, grantor []common.Address) (*EscrowEscrowRemovedIterator, error) {

	var symbolRule []interface{}
	for _, symbolItem := range symbol {
		symbolRule = append(symbolRule, symbolItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var grantorRule []interface{}
	for _, grantorItem := range grantor {
		grantorRule = append(grantorRule, grantorItem)
	}

	logs, sub, err := _Escrow.contract.FilterLogs(opts, "EscrowRemoved", symbolRule, tokenRule, grantorRule)
	if err != nil {
		return nil, err
	}
	return &EscrowEscrowRemovedIterator{contract: _Escrow.contract, event: "EscrowRemoved", logs: logs, sub: sub}, nil
}

// WatchEscrowRemoved is a free log subscription operation binding the contract event 0xd4903355a168bfba0a71e036316138be08ba1f745566878969512bb7554f2c31.
//
// Solidity: e EscrowRemoved(symbol indexed string, token indexed address, grantor indexed address)
func (_Escrow *EscrowFilterer) WatchEscrowRemoved(opts *bind.WatchOpts, sink chan<- *EscrowEscrowRemoved, symbol []string, token []common.Address, grantor []common.Address) (event.Subscription, error) {

	var symbolRule []interface{}
	for _, symbolItem := range symbol {
		symbolRule = append(symbolRule, symbolItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var grantorRule []interface{}
	for _, grantorItem := range grantor {
		grantorRule = append(grantorRule, grantorItem)
	}

	logs, sub, err := _Escrow.contract.WatchLogs(opts, "EscrowRemoved", symbolRule, tokenRule, grantorRule)
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
