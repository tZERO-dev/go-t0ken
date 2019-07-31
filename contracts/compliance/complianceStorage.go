// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package compliance

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

// StorageABI is the input ABI used to generate the binding from.
const StorageABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"permissionExists\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"k\",\"type\":\"bytes32\"}],\"name\":\"deleteBytes32\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"k\",\"type\":\"bytes32\"}],\"name\":\"deleteAddress\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"k\",\"type\":\"bytes32\"}],\"name\":\"getInt256\",\"outputs\":[{\"name\":\"\",\"type\":\"int256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"locked\",\"type\":\"bool\"}],\"name\":\"setLocked\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"k\",\"type\":\"bytes32\"}],\"name\":\"getAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"k\",\"type\":\"bytes32\"}],\"name\":\"deleteBool\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"k\",\"type\":\"bytes32\"},{\"name\":\"v\",\"type\":\"bytes\"}],\"name\":\"setBytes\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"k\",\"type\":\"bytes32\"}],\"name\":\"getUint256\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"permissionIndexOf\",\"outputs\":[{\"name\":\"\",\"type\":\"int256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"kill\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"k\",\"type\":\"bytes32\"}],\"name\":\"deleteInt256\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"k\",\"type\":\"bytes32\"},{\"name\":\"v\",\"type\":\"bytes32\"}],\"name\":\"setBytes32\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"k\",\"type\":\"bytes32\"},{\"name\":\"v\",\"type\":\"uint256\"}],\"name\":\"setUint256\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwner\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"k\",\"type\":\"bytes32\"}],\"name\":\"deleteUint256\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ZERO_ADDRESS\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"k\",\"type\":\"bytes32\"}],\"name\":\"deleteBytes\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"k\",\"type\":\"bytes32\"},{\"name\":\"v\",\"type\":\"string\"}],\"name\":\"setString\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"k\",\"type\":\"bytes32\"}],\"name\":\"getBool\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"k\",\"type\":\"bytes32\"}],\"name\":\"getString\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isLocked\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"k\",\"type\":\"bytes32\"}],\"name\":\"getBytes32\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"permissions\",\"outputs\":[{\"name\":\"count\",\"type\":\"int256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"k\",\"type\":\"bytes32\"},{\"name\":\"v\",\"type\":\"bool\"}],\"name\":\"setBool\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"k\",\"type\":\"bytes32\"}],\"name\":\"getBytes\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"k\",\"type\":\"bytes32\"},{\"name\":\"v\",\"type\":\"address\"}],\"name\":\"setAddress\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"index\",\"type\":\"int256\"}],\"name\":\"permissionAt\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"},{\"name\":\"grant\",\"type\":\"bool\"}],\"name\":\"setPermission\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"k\",\"type\":\"bytes32\"},{\"name\":\"v\",\"type\":\"int256\"}],\"name\":\"setInt256\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"k\",\"type\":\"bytes32\"}],\"name\":\"deleteString\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"oldOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnerTransferred\",\"type\":\"event\"}]"

// StorageBin is the compiled bytecode used for deploying new contracts.
const StorageBin = `608060405260018054600080546001600160a01b031916331790556001600160a81b0319169055611822806100356000396000f3fe608060405234801561001057600080fd5b50600436106101f05760003560e01c8063538ba4f91161010f578063ab8c71c0116100a2578063ea11cd4211610071578063ea11cd4214610665578063ec6263c014610682578063f64f41db146106b0578063f6bb3cc4146106d3576101f0565b8063ab8c71c0146105ef578063abfdcced146105f7578063c031a1801461061c578063ca446dd914610639576101f0565b80638da5cb5b116100de5780638da5cb5b14610530578063986e791a14610538578063a4e2d634146105ca578063a6ed563e146105d2576101f0565b8063538ba4f914610477578063616b59f61461047f5780636e8995501461049c5780637ae1cfca14610513576101f0565b806333598b00116101875780634e91db08116101565780634e91db08146103ee5780634f3029c2146104115780634fb2e45d1461043457806351f8e7dd1461045a576101f0565b806333598b00146103865780633ad8f687146103a357806341c0e1b5146103c95780634869ecc9146103d1576101f0565b8063211e28b6116101c3578063211e28b61461029a57806321f8a721146102b95780632c62ff2d146102f25780632e28d0841461030f576101f0565b806301952ffe146101f55780630b9adc571461022f5780630e14a3761461024e57806316c7d1d51461026b575b600080fd5b61021b6004803603602081101561020b57600080fd5b50356001600160a01b03166106f0565b604080519115158252519081900360200190f35b61024c6004803603602081101561024557600080fd5b5035610709565b005b61024c6004803603602081101561026457600080fd5b503561077f565b6102886004803603602081101561028157600080fd5b5035610802565b60408051918252519081900360200190f35b61024c600480360360208110156102b057600080fd5b50351515610814565b6102d6600480360360208110156102cf57600080fd5b5035610917565b604080516001600160a01b039092168252519081900360200190f35b61024c6004803603602081101561030857600080fd5b5035610932565b61024c6004803603604081101561032557600080fd5b8135919081019060408101602082013564010000000081111561034757600080fd5b82018360208201111561035957600080fd5b8035906020019184600183028401116401000000008311171561037b57600080fd5b5090925090506109af565b6102886004803603602081101561039c57600080fd5b5035610a33565b610288600480360360208110156103b957600080fd5b50356001600160a01b0316610a45565b61024c610a58565b61024c600480360360208110156103e757600080fd5b5035610ae0565b61024c6004803603604081101561040457600080fd5b5080359060200135610b56565b61024c6004803603604081101561042757600080fd5b5080359060200135610bcd565b61024c6004803603602081101561044a57600080fd5b50356001600160a01b0316610c44565b61024c6004803603602081101561047057600080fd5b5035610db6565b6102d6610e2c565b61024c6004803603602081101561049557600080fd5b5035610e3b565b61024c600480360360408110156104b257600080fd5b813591908101906040810160208201356401000000008111156104d457600080fd5b8201836020820111156104e657600080fd5b8035906020019184600183028401116401000000008311171561050857600080fd5b509092509050610eba565b61021b6004803603602081101561052957600080fd5b5035610f38565b6102d6610f4d565b6105556004803603602081101561054e57600080fd5b5035610f5c565b6040805160208082528351818301528351919283929083019185019080838360005b8381101561058f578181015183820152602001610577565b50505050905090810190601f1680156105bc5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b61021b610ffd565b610288600480360360208110156105e857600080fd5b503561100d565b61028861101f565b61024c6004803603604081101561060d57600080fd5b50803590602001351515611025565b6105556004803603602081101561063257600080fd5b50356110aa565b61024c6004803603604081101561064f57600080fd5b50803590602001356001600160a01b0316611114565b6102d66004803603602081101561067b57600080fd5b50356111a7565b61024c6004803603604081101561069857600080fd5b506001600160a01b03813516906020013515156111ba565b61024c600480360360408110156106c657600080fd5b5080359060200135611307565b61024c600480360360208110156106e957600080fd5b503561137e565b600061070360028363ffffffff6113fa16565b92915050565b61071a60023363ffffffff6113fa16565b8061072f57506000546001600160a01b031633145b61076e576040805162461bcd60e51b815260206004820152601a60248201526000805160206117a6833981519152604482015290519081900360640190fd5b600090815260076020526040812055565b61079060023363ffffffff6113fa16565b806107a557506000546001600160a01b031633145b6107e4576040805162461bcd60e51b815260206004820152601a60248201526000805160206117a6833981519152604482015290519081900360640190fd5b600090815260056020526040902080546001600160a01b0319169055565b60009081526009602052604090205490565b6000546001600160a01b031633148061083d57506001546000546001600160a01b039081169116145b61088e576040805162461bcd60e51b815260206004820152601960248201527f4f776e6572206163636f756e7420697320726571756972656400000000000000604482015290519081900360640190fd5b60015460ff600160a01b90910416151581151514156108de5760405162461bcd60e51b81526004018080602001828103825260288152602001806117c66028913960400191505060405180910390fd5b60018054911515600160a01b027fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff909216919091179055565b6000908152600560205260409020546001600160a01b031690565b61094360023363ffffffff6113fa16565b8061095857506000546001600160a01b031633145b610997576040805162461bcd60e51b815260206004820152601a60248201526000805160206117a6833981519152604482015290519081900360640190fd5b6000908152600660205260409020805460ff19169055565b6109c060023363ffffffff6113fa16565b806109d557506000546001600160a01b031633145b610a14576040805162461bcd60e51b815260206004820152601a60248201526000805160206117a6833981519152604482015290519081900360640190fd5b6000838152600860205260409020610a2d9083836116a5565b50505050565b6000908152600b602052604090205490565b600061070360028363ffffffff61143116565b6000546001600160a01b0316331480610a8157506001546000546001600160a01b039081169116145b610ad2576040805162461bcd60e51b815260206004820152601960248201527f4f776e6572206163636f756e7420697320726571756972656400000000000000604482015290519081900360640190fd5b6000546001600160a01b0316ff5b610af160023363ffffffff6113fa16565b80610b0657506000546001600160a01b031633145b610b45576040805162461bcd60e51b815260206004820152601a60248201526000805160206117a6833981519152604482015290519081900360640190fd5b600090815260096020526040812055565b610b6760023363ffffffff6113fa16565b80610b7c57506000546001600160a01b031633145b610bbb576040805162461bcd60e51b815260206004820152601a60248201526000805160206117a6833981519152604482015290519081900360640190fd5b60009182526007602052604090912055565b610bde60023363ffffffff6113fa16565b80610bf357506000546001600160a01b031633145b610c32576040805162461bcd60e51b815260206004820152601a60248201526000805160206117a6833981519152604482015290519081900360640190fd5b6000918252600b602052604090912055565b6000546001600160a01b0316331480610c6d57506001546000546001600160a01b039081169116145b610cbe576040805162461bcd60e51b815260206004820152601960248201527f4f776e6572206163636f756e7420697320726571756972656400000000000000604482015290519081900360640190fd5b6000546001600160a01b0382811691161415610d0b5760405162461bcd60e51b81526004018080602001828103825260258152602001806117816025913960400191505060405180910390fd5b6001600160a01b038116610d66576040805162461bcd60e51b815260206004820181905260248201527f4e6577204f776e65722063616e6e6f74206265207a65726f2061646472657373604482015290519081900360640190fd5b600080546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8934ce4adea8d9ce0d714d2c22b86790e41b7731c84b926fbbdc1d40ff6533c99190a35050565b610dc760023363ffffffff6113fa16565b80610ddc57506000546001600160a01b031633145b610e1b576040805162461bcd60e51b815260206004820152601a60248201526000805160206117a6833981519152604482015290519081900360640190fd5b6000908152600b6020526040812055565b6001546001600160a01b031681565b610e4c60023363ffffffff6113fa16565b80610e6157506000546001600160a01b031633145b610ea0576040805162461bcd60e51b815260206004820152601a60248201526000805160206117a6833981519152604482015290519081900360640190fd5b6000818152600860205260408120610eb791611723565b50565b610ecb60023363ffffffff6113fa16565b80610ee057506000546001600160a01b031633145b610f1f576040805162461bcd60e51b815260206004820152601a60248201526000805160206117a6833981519152604482015290519081900360640190fd5b6000838152600a60205260409020610a2d9083836116a5565b60009081526006602052604090205460ff1690565b6000546001600160a01b031681565b6000818152600a602090815260409182902080548351601f6002600019610100600186161502019093169290920491820184900484028101840190945280845260609392830182828015610ff15780601f10610fc657610100808354040283529160200191610ff1565b820191906000526020600020905b815481529060010190602001808311610fd457829003601f168201915b50505050509050919050565b600154600160a01b900460ff1681565b60009081526007602052604090205490565b60025481565b61103660023363ffffffff6113fa16565b8061104b57506000546001600160a01b031633145b61108a576040805162461bcd60e51b815260206004820152601a60248201526000805160206117a6833981519152604482015290519081900360640190fd5b600091825260066020526040909120805460ff1916911515919091179055565b60008181526008602090815260409182902080548351601f6002600019610100600186161502019093169290920491820184900484028101840190945280845260609392830182828015610ff15780601f10610fc657610100808354040283529160200191610ff1565b61112560023363ffffffff6113fa16565b8061113a57506000546001600160a01b031633145b611179576040805162461bcd60e51b815260206004820152601a60248201526000805160206117a6833981519152604482015290519081900360640190fd5b60009182526005602052604090912080546001600160a01b0319166001600160a01b03909216919091179055565b600061070360028363ffffffff61148f16565b6000546001600160a01b03163314806111e357506001546000546001600160a01b039081169116145b611234576040805162461bcd60e51b815260206004820152601960248201527f4f776e6572206163636f756e7420697320726571756972656400000000000000604482015290519081900360640190fd5b80156112a15761124b60028363ffffffff61151516565b61129c576040805162461bcd60e51b815260206004820152601e60248201527f4164647265737320616c726561647920686173207065726d697373696f6e0000604482015290519081900360640190fd5b611303565b6112b260028363ffffffff6115ba16565b611303576040805162461bcd60e51b815260206004820152601e60248201527f41646472657373207065726d697373696f6e20646f6e27742065786973740000604482015290519081900360640190fd5b5050565b61131860023363ffffffff6113fa16565b8061132d57506000546001600160a01b031633145b61136c576040805162461bcd60e51b815260206004820152601a60248201526000805160206117a6833981519152604482015290519081900360640190fd5b60009182526009602052604090912055565b61138f60023363ffffffff6113fa16565b806113a457506000546001600160a01b031633145b6113e3576040805162461bcd60e51b815260206004820152601a60248201526000805160206117a6833981519152604482015290519081900360640190fd5b6000818152600a60205260408120610eb791611723565b6001600160a01b0381166000908152600183016020526040812054600019018181128015906114295750835481125b949350505050565b60006001600160a01b03821661144a5750600019610703565b6001600160a01b03821660009081526001840160205260408120546000190190811280611478575083548112155b1561148857600019915050610703565b9392505050565b60008082121580156114a15750825482125b6114f2576040805162461bcd60e51b815260206004820152601860248201527f496e646578206f757473696465206f6620626f756e64732e0000000000000000604482015290519081900360640190fd5b50600101600090815260029190910160205260409020546001600160a01b031690565b60006001600160a01b03821661152d57506000610703565b6001600160a01b03821660009081526001840160205260408120546000190190811280159061155c5750835481125b1561156b576000915050610703565b5050815460019081018084556001600160a01b0383166000818152838601602090815260408083208590559382526002870190529190912080546001600160a01b031916909117905592915050565b6001600160a01b038116600090815260018084016020526040822054908112806115e45750835481135b156115f3576000915050610703565b835481121561165a57835460009081526002850160208181526040808420546001600160a01b03168085526001890183528185208690558585529290915280832080546001600160a01b031990811690931790558654835290912080549091169055611679565b6000818152600285016020526040902080546001600160a01b03191690555b50506001600160a01b031660009081526001828101602052604082209190915581546000190190915590565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106116e65782800160ff19823516178555611713565b82800160010185558215611713579182015b828111156117135782358255916020019190600101906116f8565b5061171f929150611763565b5090565b50805460018160011615610100020316600290046000825580601f106117495750610eb7565b601f016020900490600052602060002090810190610eb791905b61177d91905b8082111561171f5760008155600101611769565b9056fe4e6577204f776e65722063616e6e6f74206265207468652063757272656e74206f776e65724d697373696e672073746f72616765207065726d697373696f6e000000000000436f6e747261637420616c726561647920696e20726571756573746564206c6f636b207374617465a265627a7a723058207c34e1af1c1a04740342d94971968283e606072973f9df80aa8ff3817ab3bfc564736f6c634300050a0032`

// DeployStorage deploys a new Ethereum contract, binding an instance of Storage to it.
func DeployStorage(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Storage, error) {
	parsed, err := abi.JSON(strings.NewReader(StorageABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(StorageBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Storage{StorageCaller: StorageCaller{contract: contract}, StorageTransactor: StorageTransactor{contract: contract}, StorageFilterer: StorageFilterer{contract: contract}}, nil
}

// Storage is an auto generated Go binding around an Ethereum contract.
type Storage struct {
	StorageCaller     // Read-only binding to the contract
	StorageTransactor // Write-only binding to the contract
	StorageFilterer   // Log filterer for contract events
}

// StorageCaller is an auto generated read-only Go binding around an Ethereum contract.
type StorageCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StorageTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StorageTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StorageFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StorageFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StorageSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StorageSession struct {
	Contract     *Storage          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StorageCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StorageCallerSession struct {
	Contract *StorageCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// StorageTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StorageTransactorSession struct {
	Contract     *StorageTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// StorageRaw is an auto generated low-level Go binding around an Ethereum contract.
type StorageRaw struct {
	Contract *Storage // Generic contract binding to access the raw methods on
}

// StorageCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StorageCallerRaw struct {
	Contract *StorageCaller // Generic read-only contract binding to access the raw methods on
}

// StorageTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StorageTransactorRaw struct {
	Contract *StorageTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStorage creates a new instance of Storage, bound to a specific deployed contract.
func NewStorage(address common.Address, backend bind.ContractBackend) (*Storage, error) {
	contract, err := bindStorage(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Storage{StorageCaller: StorageCaller{contract: contract}, StorageTransactor: StorageTransactor{contract: contract}, StorageFilterer: StorageFilterer{contract: contract}}, nil
}

// NewStorageCaller creates a new read-only instance of Storage, bound to a specific deployed contract.
func NewStorageCaller(address common.Address, caller bind.ContractCaller) (*StorageCaller, error) {
	contract, err := bindStorage(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StorageCaller{contract: contract}, nil
}

// NewStorageTransactor creates a new write-only instance of Storage, bound to a specific deployed contract.
func NewStorageTransactor(address common.Address, transactor bind.ContractTransactor) (*StorageTransactor, error) {
	contract, err := bindStorage(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StorageTransactor{contract: contract}, nil
}

// NewStorageFilterer creates a new log filterer instance of Storage, bound to a specific deployed contract.
func NewStorageFilterer(address common.Address, filterer bind.ContractFilterer) (*StorageFilterer, error) {
	contract, err := bindStorage(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StorageFilterer{contract: contract}, nil
}

// bindStorage binds a generic wrapper to an already deployed contract.
func bindStorage(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(StorageABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Storage *StorageRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Storage.Contract.StorageCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Storage *StorageRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Storage.Contract.StorageTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Storage *StorageRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Storage.Contract.StorageTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Storage *StorageCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Storage.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Storage *StorageTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Storage.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Storage *StorageTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Storage.Contract.contract.Transact(opts, method, params...)
}

// ZEROADDRESS is a free data retrieval call binding the contract method 0x538ba4f9.
//
// Solidity: function ZERO_ADDRESS() constant returns(address)
func (_Storage *StorageCaller) ZEROADDRESS(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Storage.contract.Call(opts, out, "ZERO_ADDRESS")
	return *ret0, err
}

// ZEROADDRESS is a free data retrieval call binding the contract method 0x538ba4f9.
//
// Solidity: function ZERO_ADDRESS() constant returns(address)
func (_Storage *StorageSession) ZEROADDRESS() (common.Address, error) {
	return _Storage.Contract.ZEROADDRESS(&_Storage.CallOpts)
}

// ZEROADDRESS is a free data retrieval call binding the contract method 0x538ba4f9.
//
// Solidity: function ZERO_ADDRESS() constant returns(address)
func (_Storage *StorageCallerSession) ZEROADDRESS() (common.Address, error) {
	return _Storage.Contract.ZEROADDRESS(&_Storage.CallOpts)
}

// GetAddress is a free data retrieval call binding the contract method 0x21f8a721.
//
// Solidity: function getAddress(k bytes32) constant returns(address)
func (_Storage *StorageCaller) GetAddress(opts *bind.CallOpts, k [32]byte) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Storage.contract.Call(opts, out, "getAddress", k)
	return *ret0, err
}

// GetAddress is a free data retrieval call binding the contract method 0x21f8a721.
//
// Solidity: function getAddress(k bytes32) constant returns(address)
func (_Storage *StorageSession) GetAddress(k [32]byte) (common.Address, error) {
	return _Storage.Contract.GetAddress(&_Storage.CallOpts, k)
}

// GetAddress is a free data retrieval call binding the contract method 0x21f8a721.
//
// Solidity: function getAddress(k bytes32) constant returns(address)
func (_Storage *StorageCallerSession) GetAddress(k [32]byte) (common.Address, error) {
	return _Storage.Contract.GetAddress(&_Storage.CallOpts, k)
}

// GetBool is a free data retrieval call binding the contract method 0x7ae1cfca.
//
// Solidity: function getBool(k bytes32) constant returns(bool)
func (_Storage *StorageCaller) GetBool(opts *bind.CallOpts, k [32]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Storage.contract.Call(opts, out, "getBool", k)
	return *ret0, err
}

// GetBool is a free data retrieval call binding the contract method 0x7ae1cfca.
//
// Solidity: function getBool(k bytes32) constant returns(bool)
func (_Storage *StorageSession) GetBool(k [32]byte) (bool, error) {
	return _Storage.Contract.GetBool(&_Storage.CallOpts, k)
}

// GetBool is a free data retrieval call binding the contract method 0x7ae1cfca.
//
// Solidity: function getBool(k bytes32) constant returns(bool)
func (_Storage *StorageCallerSession) GetBool(k [32]byte) (bool, error) {
	return _Storage.Contract.GetBool(&_Storage.CallOpts, k)
}

// GetBytes is a free data retrieval call binding the contract method 0xc031a180.
//
// Solidity: function getBytes(k bytes32) constant returns(bytes)
func (_Storage *StorageCaller) GetBytes(opts *bind.CallOpts, k [32]byte) ([]byte, error) {
	var (
		ret0 = new([]byte)
	)
	out := ret0
	err := _Storage.contract.Call(opts, out, "getBytes", k)
	return *ret0, err
}

// GetBytes is a free data retrieval call binding the contract method 0xc031a180.
//
// Solidity: function getBytes(k bytes32) constant returns(bytes)
func (_Storage *StorageSession) GetBytes(k [32]byte) ([]byte, error) {
	return _Storage.Contract.GetBytes(&_Storage.CallOpts, k)
}

// GetBytes is a free data retrieval call binding the contract method 0xc031a180.
//
// Solidity: function getBytes(k bytes32) constant returns(bytes)
func (_Storage *StorageCallerSession) GetBytes(k [32]byte) ([]byte, error) {
	return _Storage.Contract.GetBytes(&_Storage.CallOpts, k)
}

// GetBytes32 is a free data retrieval call binding the contract method 0xa6ed563e.
//
// Solidity: function getBytes32(k bytes32) constant returns(bytes32)
func (_Storage *StorageCaller) GetBytes32(opts *bind.CallOpts, k [32]byte) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Storage.contract.Call(opts, out, "getBytes32", k)
	return *ret0, err
}

// GetBytes32 is a free data retrieval call binding the contract method 0xa6ed563e.
//
// Solidity: function getBytes32(k bytes32) constant returns(bytes32)
func (_Storage *StorageSession) GetBytes32(k [32]byte) ([32]byte, error) {
	return _Storage.Contract.GetBytes32(&_Storage.CallOpts, k)
}

// GetBytes32 is a free data retrieval call binding the contract method 0xa6ed563e.
//
// Solidity: function getBytes32(k bytes32) constant returns(bytes32)
func (_Storage *StorageCallerSession) GetBytes32(k [32]byte) ([32]byte, error) {
	return _Storage.Contract.GetBytes32(&_Storage.CallOpts, k)
}

// GetInt256 is a free data retrieval call binding the contract method 0x16c7d1d5.
//
// Solidity: function getInt256(k bytes32) constant returns(int256)
func (_Storage *StorageCaller) GetInt256(opts *bind.CallOpts, k [32]byte) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Storage.contract.Call(opts, out, "getInt256", k)
	return *ret0, err
}

// GetInt256 is a free data retrieval call binding the contract method 0x16c7d1d5.
//
// Solidity: function getInt256(k bytes32) constant returns(int256)
func (_Storage *StorageSession) GetInt256(k [32]byte) (*big.Int, error) {
	return _Storage.Contract.GetInt256(&_Storage.CallOpts, k)
}

// GetInt256 is a free data retrieval call binding the contract method 0x16c7d1d5.
//
// Solidity: function getInt256(k bytes32) constant returns(int256)
func (_Storage *StorageCallerSession) GetInt256(k [32]byte) (*big.Int, error) {
	return _Storage.Contract.GetInt256(&_Storage.CallOpts, k)
}

// GetString is a free data retrieval call binding the contract method 0x986e791a.
//
// Solidity: function getString(k bytes32) constant returns(string)
func (_Storage *StorageCaller) GetString(opts *bind.CallOpts, k [32]byte) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Storage.contract.Call(opts, out, "getString", k)
	return *ret0, err
}

// GetString is a free data retrieval call binding the contract method 0x986e791a.
//
// Solidity: function getString(k bytes32) constant returns(string)
func (_Storage *StorageSession) GetString(k [32]byte) (string, error) {
	return _Storage.Contract.GetString(&_Storage.CallOpts, k)
}

// GetString is a free data retrieval call binding the contract method 0x986e791a.
//
// Solidity: function getString(k bytes32) constant returns(string)
func (_Storage *StorageCallerSession) GetString(k [32]byte) (string, error) {
	return _Storage.Contract.GetString(&_Storage.CallOpts, k)
}

// GetUint256 is a free data retrieval call binding the contract method 0x33598b00.
//
// Solidity: function getUint256(k bytes32) constant returns(uint256)
func (_Storage *StorageCaller) GetUint256(opts *bind.CallOpts, k [32]byte) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Storage.contract.Call(opts, out, "getUint256", k)
	return *ret0, err
}

// GetUint256 is a free data retrieval call binding the contract method 0x33598b00.
//
// Solidity: function getUint256(k bytes32) constant returns(uint256)
func (_Storage *StorageSession) GetUint256(k [32]byte) (*big.Int, error) {
	return _Storage.Contract.GetUint256(&_Storage.CallOpts, k)
}

// GetUint256 is a free data retrieval call binding the contract method 0x33598b00.
//
// Solidity: function getUint256(k bytes32) constant returns(uint256)
func (_Storage *StorageCallerSession) GetUint256(k [32]byte) (*big.Int, error) {
	return _Storage.Contract.GetUint256(&_Storage.CallOpts, k)
}

// IsLocked is a free data retrieval call binding the contract method 0xa4e2d634.
//
// Solidity: function isLocked() constant returns(bool)
func (_Storage *StorageCaller) IsLocked(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Storage.contract.Call(opts, out, "isLocked")
	return *ret0, err
}

// IsLocked is a free data retrieval call binding the contract method 0xa4e2d634.
//
// Solidity: function isLocked() constant returns(bool)
func (_Storage *StorageSession) IsLocked() (bool, error) {
	return _Storage.Contract.IsLocked(&_Storage.CallOpts)
}

// IsLocked is a free data retrieval call binding the contract method 0xa4e2d634.
//
// Solidity: function isLocked() constant returns(bool)
func (_Storage *StorageCallerSession) IsLocked() (bool, error) {
	return _Storage.Contract.IsLocked(&_Storage.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Storage *StorageCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Storage.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Storage *StorageSession) Owner() (common.Address, error) {
	return _Storage.Contract.Owner(&_Storage.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Storage *StorageCallerSession) Owner() (common.Address, error) {
	return _Storage.Contract.Owner(&_Storage.CallOpts)
}

// PermissionAt is a free data retrieval call binding the contract method 0xea11cd42.
//
// Solidity: function permissionAt(index int256) constant returns(address)
func (_Storage *StorageCaller) PermissionAt(opts *bind.CallOpts, index *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Storage.contract.Call(opts, out, "permissionAt", index)
	return *ret0, err
}

// PermissionAt is a free data retrieval call binding the contract method 0xea11cd42.
//
// Solidity: function permissionAt(index int256) constant returns(address)
func (_Storage *StorageSession) PermissionAt(index *big.Int) (common.Address, error) {
	return _Storage.Contract.PermissionAt(&_Storage.CallOpts, index)
}

// PermissionAt is a free data retrieval call binding the contract method 0xea11cd42.
//
// Solidity: function permissionAt(index int256) constant returns(address)
func (_Storage *StorageCallerSession) PermissionAt(index *big.Int) (common.Address, error) {
	return _Storage.Contract.PermissionAt(&_Storage.CallOpts, index)
}

// PermissionExists is a free data retrieval call binding the contract method 0x01952ffe.
//
// Solidity: function permissionExists(addr address) constant returns(bool)
func (_Storage *StorageCaller) PermissionExists(opts *bind.CallOpts, addr common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Storage.contract.Call(opts, out, "permissionExists", addr)
	return *ret0, err
}

// PermissionExists is a free data retrieval call binding the contract method 0x01952ffe.
//
// Solidity: function permissionExists(addr address) constant returns(bool)
func (_Storage *StorageSession) PermissionExists(addr common.Address) (bool, error) {
	return _Storage.Contract.PermissionExists(&_Storage.CallOpts, addr)
}

// PermissionExists is a free data retrieval call binding the contract method 0x01952ffe.
//
// Solidity: function permissionExists(addr address) constant returns(bool)
func (_Storage *StorageCallerSession) PermissionExists(addr common.Address) (bool, error) {
	return _Storage.Contract.PermissionExists(&_Storage.CallOpts, addr)
}

// PermissionIndexOf is a free data retrieval call binding the contract method 0x3ad8f687.
//
// Solidity: function permissionIndexOf(addr address) constant returns(int256)
func (_Storage *StorageCaller) PermissionIndexOf(opts *bind.CallOpts, addr common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Storage.contract.Call(opts, out, "permissionIndexOf", addr)
	return *ret0, err
}

// PermissionIndexOf is a free data retrieval call binding the contract method 0x3ad8f687.
//
// Solidity: function permissionIndexOf(addr address) constant returns(int256)
func (_Storage *StorageSession) PermissionIndexOf(addr common.Address) (*big.Int, error) {
	return _Storage.Contract.PermissionIndexOf(&_Storage.CallOpts, addr)
}

// PermissionIndexOf is a free data retrieval call binding the contract method 0x3ad8f687.
//
// Solidity: function permissionIndexOf(addr address) constant returns(int256)
func (_Storage *StorageCallerSession) PermissionIndexOf(addr common.Address) (*big.Int, error) {
	return _Storage.Contract.PermissionIndexOf(&_Storage.CallOpts, addr)
}

// Permissions is a free data retrieval call binding the contract method 0xab8c71c0.
//
// Solidity: function permissions() constant returns(count int256)
func (_Storage *StorageCaller) Permissions(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Storage.contract.Call(opts, out, "permissions")
	return *ret0, err
}

// Permissions is a free data retrieval call binding the contract method 0xab8c71c0.
//
// Solidity: function permissions() constant returns(count int256)
func (_Storage *StorageSession) Permissions() (*big.Int, error) {
	return _Storage.Contract.Permissions(&_Storage.CallOpts)
}

// Permissions is a free data retrieval call binding the contract method 0xab8c71c0.
//
// Solidity: function permissions() constant returns(count int256)
func (_Storage *StorageCallerSession) Permissions() (*big.Int, error) {
	return _Storage.Contract.Permissions(&_Storage.CallOpts)
}

// DeleteAddress is a paid mutator transaction binding the contract method 0x0e14a376.
//
// Solidity: function deleteAddress(k bytes32) returns()
func (_Storage *StorageTransactor) DeleteAddress(opts *bind.TransactOpts, k [32]byte) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "deleteAddress", k)
}

// DeleteAddress is a paid mutator transaction binding the contract method 0x0e14a376.
//
// Solidity: function deleteAddress(k bytes32) returns()
func (_Storage *StorageSession) DeleteAddress(k [32]byte) (*types.Transaction, error) {
	return _Storage.Contract.DeleteAddress(&_Storage.TransactOpts, k)
}

// DeleteAddress is a paid mutator transaction binding the contract method 0x0e14a376.
//
// Solidity: function deleteAddress(k bytes32) returns()
func (_Storage *StorageTransactorSession) DeleteAddress(k [32]byte) (*types.Transaction, error) {
	return _Storage.Contract.DeleteAddress(&_Storage.TransactOpts, k)
}

// DeleteBool is a paid mutator transaction binding the contract method 0x2c62ff2d.
//
// Solidity: function deleteBool(k bytes32) returns()
func (_Storage *StorageTransactor) DeleteBool(opts *bind.TransactOpts, k [32]byte) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "deleteBool", k)
}

// DeleteBool is a paid mutator transaction binding the contract method 0x2c62ff2d.
//
// Solidity: function deleteBool(k bytes32) returns()
func (_Storage *StorageSession) DeleteBool(k [32]byte) (*types.Transaction, error) {
	return _Storage.Contract.DeleteBool(&_Storage.TransactOpts, k)
}

// DeleteBool is a paid mutator transaction binding the contract method 0x2c62ff2d.
//
// Solidity: function deleteBool(k bytes32) returns()
func (_Storage *StorageTransactorSession) DeleteBool(k [32]byte) (*types.Transaction, error) {
	return _Storage.Contract.DeleteBool(&_Storage.TransactOpts, k)
}

// DeleteBytes is a paid mutator transaction binding the contract method 0x616b59f6.
//
// Solidity: function deleteBytes(k bytes32) returns()
func (_Storage *StorageTransactor) DeleteBytes(opts *bind.TransactOpts, k [32]byte) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "deleteBytes", k)
}

// DeleteBytes is a paid mutator transaction binding the contract method 0x616b59f6.
//
// Solidity: function deleteBytes(k bytes32) returns()
func (_Storage *StorageSession) DeleteBytes(k [32]byte) (*types.Transaction, error) {
	return _Storage.Contract.DeleteBytes(&_Storage.TransactOpts, k)
}

// DeleteBytes is a paid mutator transaction binding the contract method 0x616b59f6.
//
// Solidity: function deleteBytes(k bytes32) returns()
func (_Storage *StorageTransactorSession) DeleteBytes(k [32]byte) (*types.Transaction, error) {
	return _Storage.Contract.DeleteBytes(&_Storage.TransactOpts, k)
}

// DeleteBytes32 is a paid mutator transaction binding the contract method 0x0b9adc57.
//
// Solidity: function deleteBytes32(k bytes32) returns()
func (_Storage *StorageTransactor) DeleteBytes32(opts *bind.TransactOpts, k [32]byte) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "deleteBytes32", k)
}

// DeleteBytes32 is a paid mutator transaction binding the contract method 0x0b9adc57.
//
// Solidity: function deleteBytes32(k bytes32) returns()
func (_Storage *StorageSession) DeleteBytes32(k [32]byte) (*types.Transaction, error) {
	return _Storage.Contract.DeleteBytes32(&_Storage.TransactOpts, k)
}

// DeleteBytes32 is a paid mutator transaction binding the contract method 0x0b9adc57.
//
// Solidity: function deleteBytes32(k bytes32) returns()
func (_Storage *StorageTransactorSession) DeleteBytes32(k [32]byte) (*types.Transaction, error) {
	return _Storage.Contract.DeleteBytes32(&_Storage.TransactOpts, k)
}

// DeleteInt256 is a paid mutator transaction binding the contract method 0x4869ecc9.
//
// Solidity: function deleteInt256(k bytes32) returns()
func (_Storage *StorageTransactor) DeleteInt256(opts *bind.TransactOpts, k [32]byte) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "deleteInt256", k)
}

// DeleteInt256 is a paid mutator transaction binding the contract method 0x4869ecc9.
//
// Solidity: function deleteInt256(k bytes32) returns()
func (_Storage *StorageSession) DeleteInt256(k [32]byte) (*types.Transaction, error) {
	return _Storage.Contract.DeleteInt256(&_Storage.TransactOpts, k)
}

// DeleteInt256 is a paid mutator transaction binding the contract method 0x4869ecc9.
//
// Solidity: function deleteInt256(k bytes32) returns()
func (_Storage *StorageTransactorSession) DeleteInt256(k [32]byte) (*types.Transaction, error) {
	return _Storage.Contract.DeleteInt256(&_Storage.TransactOpts, k)
}

// DeleteString is a paid mutator transaction binding the contract method 0xf6bb3cc4.
//
// Solidity: function deleteString(k bytes32) returns()
func (_Storage *StorageTransactor) DeleteString(opts *bind.TransactOpts, k [32]byte) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "deleteString", k)
}

// DeleteString is a paid mutator transaction binding the contract method 0xf6bb3cc4.
//
// Solidity: function deleteString(k bytes32) returns()
func (_Storage *StorageSession) DeleteString(k [32]byte) (*types.Transaction, error) {
	return _Storage.Contract.DeleteString(&_Storage.TransactOpts, k)
}

// DeleteString is a paid mutator transaction binding the contract method 0xf6bb3cc4.
//
// Solidity: function deleteString(k bytes32) returns()
func (_Storage *StorageTransactorSession) DeleteString(k [32]byte) (*types.Transaction, error) {
	return _Storage.Contract.DeleteString(&_Storage.TransactOpts, k)
}

// DeleteUint256 is a paid mutator transaction binding the contract method 0x51f8e7dd.
//
// Solidity: function deleteUint256(k bytes32) returns()
func (_Storage *StorageTransactor) DeleteUint256(opts *bind.TransactOpts, k [32]byte) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "deleteUint256", k)
}

// DeleteUint256 is a paid mutator transaction binding the contract method 0x51f8e7dd.
//
// Solidity: function deleteUint256(k bytes32) returns()
func (_Storage *StorageSession) DeleteUint256(k [32]byte) (*types.Transaction, error) {
	return _Storage.Contract.DeleteUint256(&_Storage.TransactOpts, k)
}

// DeleteUint256 is a paid mutator transaction binding the contract method 0x51f8e7dd.
//
// Solidity: function deleteUint256(k bytes32) returns()
func (_Storage *StorageTransactorSession) DeleteUint256(k [32]byte) (*types.Transaction, error) {
	return _Storage.Contract.DeleteUint256(&_Storage.TransactOpts, k)
}

// Kill is a paid mutator transaction binding the contract method 0x41c0e1b5.
//
// Solidity: function kill() returns()
func (_Storage *StorageTransactor) Kill(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "kill")
}

// Kill is a paid mutator transaction binding the contract method 0x41c0e1b5.
//
// Solidity: function kill() returns()
func (_Storage *StorageSession) Kill() (*types.Transaction, error) {
	return _Storage.Contract.Kill(&_Storage.TransactOpts)
}

// Kill is a paid mutator transaction binding the contract method 0x41c0e1b5.
//
// Solidity: function kill() returns()
func (_Storage *StorageTransactorSession) Kill() (*types.Transaction, error) {
	return _Storage.Contract.Kill(&_Storage.TransactOpts)
}

// SetAddress is a paid mutator transaction binding the contract method 0xca446dd9.
//
// Solidity: function setAddress(k bytes32, v address) returns()
func (_Storage *StorageTransactor) SetAddress(opts *bind.TransactOpts, k [32]byte, v common.Address) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "setAddress", k, v)
}

// SetAddress is a paid mutator transaction binding the contract method 0xca446dd9.
//
// Solidity: function setAddress(k bytes32, v address) returns()
func (_Storage *StorageSession) SetAddress(k [32]byte, v common.Address) (*types.Transaction, error) {
	return _Storage.Contract.SetAddress(&_Storage.TransactOpts, k, v)
}

// SetAddress is a paid mutator transaction binding the contract method 0xca446dd9.
//
// Solidity: function setAddress(k bytes32, v address) returns()
func (_Storage *StorageTransactorSession) SetAddress(k [32]byte, v common.Address) (*types.Transaction, error) {
	return _Storage.Contract.SetAddress(&_Storage.TransactOpts, k, v)
}

// SetBool is a paid mutator transaction binding the contract method 0xabfdcced.
//
// Solidity: function setBool(k bytes32, v bool) returns()
func (_Storage *StorageTransactor) SetBool(opts *bind.TransactOpts, k [32]byte, v bool) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "setBool", k, v)
}

// SetBool is a paid mutator transaction binding the contract method 0xabfdcced.
//
// Solidity: function setBool(k bytes32, v bool) returns()
func (_Storage *StorageSession) SetBool(k [32]byte, v bool) (*types.Transaction, error) {
	return _Storage.Contract.SetBool(&_Storage.TransactOpts, k, v)
}

// SetBool is a paid mutator transaction binding the contract method 0xabfdcced.
//
// Solidity: function setBool(k bytes32, v bool) returns()
func (_Storage *StorageTransactorSession) SetBool(k [32]byte, v bool) (*types.Transaction, error) {
	return _Storage.Contract.SetBool(&_Storage.TransactOpts, k, v)
}

// SetBytes is a paid mutator transaction binding the contract method 0x2e28d084.
//
// Solidity: function setBytes(k bytes32, v bytes) returns()
func (_Storage *StorageTransactor) SetBytes(opts *bind.TransactOpts, k [32]byte, v []byte) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "setBytes", k, v)
}

// SetBytes is a paid mutator transaction binding the contract method 0x2e28d084.
//
// Solidity: function setBytes(k bytes32, v bytes) returns()
func (_Storage *StorageSession) SetBytes(k [32]byte, v []byte) (*types.Transaction, error) {
	return _Storage.Contract.SetBytes(&_Storage.TransactOpts, k, v)
}

// SetBytes is a paid mutator transaction binding the contract method 0x2e28d084.
//
// Solidity: function setBytes(k bytes32, v bytes) returns()
func (_Storage *StorageTransactorSession) SetBytes(k [32]byte, v []byte) (*types.Transaction, error) {
	return _Storage.Contract.SetBytes(&_Storage.TransactOpts, k, v)
}

// SetBytes32 is a paid mutator transaction binding the contract method 0x4e91db08.
//
// Solidity: function setBytes32(k bytes32, v bytes32) returns()
func (_Storage *StorageTransactor) SetBytes32(opts *bind.TransactOpts, k [32]byte, v [32]byte) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "setBytes32", k, v)
}

// SetBytes32 is a paid mutator transaction binding the contract method 0x4e91db08.
//
// Solidity: function setBytes32(k bytes32, v bytes32) returns()
func (_Storage *StorageSession) SetBytes32(k [32]byte, v [32]byte) (*types.Transaction, error) {
	return _Storage.Contract.SetBytes32(&_Storage.TransactOpts, k, v)
}

// SetBytes32 is a paid mutator transaction binding the contract method 0x4e91db08.
//
// Solidity: function setBytes32(k bytes32, v bytes32) returns()
func (_Storage *StorageTransactorSession) SetBytes32(k [32]byte, v [32]byte) (*types.Transaction, error) {
	return _Storage.Contract.SetBytes32(&_Storage.TransactOpts, k, v)
}

// SetInt256 is a paid mutator transaction binding the contract method 0xf64f41db.
//
// Solidity: function setInt256(k bytes32, v int256) returns()
func (_Storage *StorageTransactor) SetInt256(opts *bind.TransactOpts, k [32]byte, v *big.Int) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "setInt256", k, v)
}

// SetInt256 is a paid mutator transaction binding the contract method 0xf64f41db.
//
// Solidity: function setInt256(k bytes32, v int256) returns()
func (_Storage *StorageSession) SetInt256(k [32]byte, v *big.Int) (*types.Transaction, error) {
	return _Storage.Contract.SetInt256(&_Storage.TransactOpts, k, v)
}

// SetInt256 is a paid mutator transaction binding the contract method 0xf64f41db.
//
// Solidity: function setInt256(k bytes32, v int256) returns()
func (_Storage *StorageTransactorSession) SetInt256(k [32]byte, v *big.Int) (*types.Transaction, error) {
	return _Storage.Contract.SetInt256(&_Storage.TransactOpts, k, v)
}

// SetLocked is a paid mutator transaction binding the contract method 0x211e28b6.
//
// Solidity: function setLocked(locked bool) returns()
func (_Storage *StorageTransactor) SetLocked(opts *bind.TransactOpts, locked bool) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "setLocked", locked)
}

// SetLocked is a paid mutator transaction binding the contract method 0x211e28b6.
//
// Solidity: function setLocked(locked bool) returns()
func (_Storage *StorageSession) SetLocked(locked bool) (*types.Transaction, error) {
	return _Storage.Contract.SetLocked(&_Storage.TransactOpts, locked)
}

// SetLocked is a paid mutator transaction binding the contract method 0x211e28b6.
//
// Solidity: function setLocked(locked bool) returns()
func (_Storage *StorageTransactorSession) SetLocked(locked bool) (*types.Transaction, error) {
	return _Storage.Contract.SetLocked(&_Storage.TransactOpts, locked)
}

// SetPermission is a paid mutator transaction binding the contract method 0xec6263c0.
//
// Solidity: function setPermission(addr address, grant bool) returns()
func (_Storage *StorageTransactor) SetPermission(opts *bind.TransactOpts, addr common.Address, grant bool) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "setPermission", addr, grant)
}

// SetPermission is a paid mutator transaction binding the contract method 0xec6263c0.
//
// Solidity: function setPermission(addr address, grant bool) returns()
func (_Storage *StorageSession) SetPermission(addr common.Address, grant bool) (*types.Transaction, error) {
	return _Storage.Contract.SetPermission(&_Storage.TransactOpts, addr, grant)
}

// SetPermission is a paid mutator transaction binding the contract method 0xec6263c0.
//
// Solidity: function setPermission(addr address, grant bool) returns()
func (_Storage *StorageTransactorSession) SetPermission(addr common.Address, grant bool) (*types.Transaction, error) {
	return _Storage.Contract.SetPermission(&_Storage.TransactOpts, addr, grant)
}

// SetString is a paid mutator transaction binding the contract method 0x6e899550.
//
// Solidity: function setString(k bytes32, v string) returns()
func (_Storage *StorageTransactor) SetString(opts *bind.TransactOpts, k [32]byte, v string) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "setString", k, v)
}

// SetString is a paid mutator transaction binding the contract method 0x6e899550.
//
// Solidity: function setString(k bytes32, v string) returns()
func (_Storage *StorageSession) SetString(k [32]byte, v string) (*types.Transaction, error) {
	return _Storage.Contract.SetString(&_Storage.TransactOpts, k, v)
}

// SetString is a paid mutator transaction binding the contract method 0x6e899550.
//
// Solidity: function setString(k bytes32, v string) returns()
func (_Storage *StorageTransactorSession) SetString(k [32]byte, v string) (*types.Transaction, error) {
	return _Storage.Contract.SetString(&_Storage.TransactOpts, k, v)
}

// SetUint256 is a paid mutator transaction binding the contract method 0x4f3029c2.
//
// Solidity: function setUint256(k bytes32, v uint256) returns()
func (_Storage *StorageTransactor) SetUint256(opts *bind.TransactOpts, k [32]byte, v *big.Int) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "setUint256", k, v)
}

// SetUint256 is a paid mutator transaction binding the contract method 0x4f3029c2.
//
// Solidity: function setUint256(k bytes32, v uint256) returns()
func (_Storage *StorageSession) SetUint256(k [32]byte, v *big.Int) (*types.Transaction, error) {
	return _Storage.Contract.SetUint256(&_Storage.TransactOpts, k, v)
}

// SetUint256 is a paid mutator transaction binding the contract method 0x4f3029c2.
//
// Solidity: function setUint256(k bytes32, v uint256) returns()
func (_Storage *StorageTransactorSession) SetUint256(k [32]byte, v *big.Int) (*types.Transaction, error) {
	return _Storage.Contract.SetUint256(&_Storage.TransactOpts, k, v)
}

// TransferOwner is a paid mutator transaction binding the contract method 0x4fb2e45d.
//
// Solidity: function transferOwner(newOwner address) returns()
func (_Storage *StorageTransactor) TransferOwner(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "transferOwner", newOwner)
}

// TransferOwner is a paid mutator transaction binding the contract method 0x4fb2e45d.
//
// Solidity: function transferOwner(newOwner address) returns()
func (_Storage *StorageSession) TransferOwner(newOwner common.Address) (*types.Transaction, error) {
	return _Storage.Contract.TransferOwner(&_Storage.TransactOpts, newOwner)
}

// TransferOwner is a paid mutator transaction binding the contract method 0x4fb2e45d.
//
// Solidity: function transferOwner(newOwner address) returns()
func (_Storage *StorageTransactorSession) TransferOwner(newOwner common.Address) (*types.Transaction, error) {
	return _Storage.Contract.TransferOwner(&_Storage.TransactOpts, newOwner)
}

// StorageOwnerTransferredIterator is returned from FilterOwnerTransferred and is used to iterate over the raw logs and unpacked data for OwnerTransferred events raised by the Storage contract.
type StorageOwnerTransferredIterator struct {
	Event *StorageOwnerTransferred // Event containing the contract specifics and raw log

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
func (it *StorageOwnerTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StorageOwnerTransferred)
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
		it.Event = new(StorageOwnerTransferred)
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
func (it *StorageOwnerTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StorageOwnerTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StorageOwnerTransferred represents a OwnerTransferred event raised by the Storage contract.
type StorageOwnerTransferred struct {
	OldOwner common.Address
	NewOwner common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOwnerTransferred is a free log retrieval operation binding the contract event 0x8934ce4adea8d9ce0d714d2c22b86790e41b7731c84b926fbbdc1d40ff6533c9.
//
// Solidity: e OwnerTransferred(oldOwner indexed address, newOwner indexed address)
func (_Storage *StorageFilterer) FilterOwnerTransferred(opts *bind.FilterOpts, oldOwner []common.Address, newOwner []common.Address) (*StorageOwnerTransferredIterator, error) {

	var oldOwnerRule []interface{}
	for _, oldOwnerItem := range oldOwner {
		oldOwnerRule = append(oldOwnerRule, oldOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Storage.contract.FilterLogs(opts, "OwnerTransferred", oldOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &StorageOwnerTransferredIterator{contract: _Storage.contract, event: "OwnerTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnerTransferred is a free log subscription operation binding the contract event 0x8934ce4adea8d9ce0d714d2c22b86790e41b7731c84b926fbbdc1d40ff6533c9.
//
// Solidity: e OwnerTransferred(oldOwner indexed address, newOwner indexed address)
func (_Storage *StorageFilterer) WatchOwnerTransferred(opts *bind.WatchOpts, sink chan<- *StorageOwnerTransferred, oldOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var oldOwnerRule []interface{}
	for _, oldOwnerItem := range oldOwner {
		oldOwnerRule = append(oldOwnerRule, oldOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Storage.contract.WatchLogs(opts, "OwnerTransferred", oldOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StorageOwnerTransferred)
				if err := _Storage.contract.UnpackLog(event, "OwnerTransferred", log); err != nil {
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
