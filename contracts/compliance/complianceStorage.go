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

// ComplianceStorageABI is the input ABI used to generate the binding from.
const ComplianceStorageABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"permissionExists\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"k\",\"type\":\"bytes32\"}],\"name\":\"getInt256\",\"outputs\":[{\"name\":\"\",\"type\":\"int256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"k\",\"type\":\"bytes32\"}],\"name\":\"getAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"k\",\"type\":\"bytes32\"},{\"name\":\"v\",\"type\":\"bytes\"}],\"name\":\"setBytes\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"k\",\"type\":\"bytes32\"}],\"name\":\"getUint256\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"permissionIndexOf\",\"outputs\":[{\"name\":\"\",\"type\":\"int256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"kill\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"k\",\"type\":\"bytes32\"},{\"name\":\"v\",\"type\":\"bytes32\"}],\"name\":\"setBytes32\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"k\",\"type\":\"bytes32\"},{\"name\":\"v\",\"type\":\"uint256\"}],\"name\":\"setUint256\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwner\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ZERO_ADDRESS\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"k\",\"type\":\"bytes32\"}],\"name\":\"deleteBytes\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"k\",\"type\":\"bytes32\"},{\"name\":\"v\",\"type\":\"string\"}],\"name\":\"setString\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"k\",\"type\":\"bytes32\"}],\"name\":\"getBool\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"k\",\"type\":\"bytes32\"}],\"name\":\"getString\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"k\",\"type\":\"bytes32\"}],\"name\":\"getBytes32\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"permissions\",\"outputs\":[{\"name\":\"count\",\"type\":\"int256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"k\",\"type\":\"bytes32\"},{\"name\":\"v\",\"type\":\"bool\"}],\"name\":\"setBool\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"k\",\"type\":\"bytes32\"}],\"name\":\"getBytes\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"k\",\"type\":\"bytes32\"},{\"name\":\"v\",\"type\":\"address\"}],\"name\":\"setAddress\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"k\",\"type\":\"bytes32\"}],\"name\":\"deleteValue\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"index\",\"type\":\"int256\"}],\"name\":\"permissionAt\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"},{\"name\":\"grant\",\"type\":\"bool\"}],\"name\":\"setPermission\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"k\",\"type\":\"bytes32\"},{\"name\":\"v\",\"type\":\"int256\"}],\"name\":\"setInt256\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"k\",\"type\":\"bytes32\"}],\"name\":\"deleteString\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"oldOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnerTransferred\",\"type\":\"event\"}]"

// ComplianceStorageBin is the compiled bytecode used for deploying new contracts.
const ComplianceStorageBin = `6080604052600180546001600160a01b03199081169091556000805490911633179055611307806100316000396000f3fe608060405234801561001057600080fd5b50600436106101ae5760003560e01c80637ae1cfca116100ee578063c031a18011610097578063ea11cd4211610071578063ea11cd4214610524578063ec6263c014610541578063f64f41db146102f5578063f6bb3cc41461056f576101ae565b8063c031a180146104be578063ca446dd9146104db578063ddd875d814610507576101ae565b8063a6ed563e116100c8578063a6ed563e146101ed578063ab8c71c014610491578063abfdcced14610499576101ae565b80637ae1cfca146103da5780638da5cb5b146103f7578063986e791a146103ff576101ae565b806341c0e1b51161015b5780634fb2e45d116101355780634fb2e45d14610318578063538ba4f91461033e578063616b59f6146103465780636e89955014610363576101ae565b806341c0e1b5146102ed5780634e91db08146102f55780634f3029c2146102f5576101ae565b80632e28d0841161018c5780632e28d0841461024e57806333598b00146101ed5780633ad8f687146102c7576101ae565b806301952ffe146101b357806316c7d1d5146101ed57806321f8a7211461021c575b600080fd5b6101d9600480360360208110156101c957600080fd5b50356001600160a01b031661058c565b604080519115158252519081900360200190f35b61020a6004803603602081101561020357600080fd5b50356105a5565b60408051918252519081900360200190f35b6102326004803603602081101561020357600080fd5b604080516001600160a01b039092168252519081900360200190f35b6102c56004803603604081101561026457600080fd5b8135919081019060408101602082013564010000000081111561028657600080fd5b82018360208201111561029857600080fd5b803590602001918460018302840111640100000000831117156102ba57600080fd5b5090925090506105b7565b005b61020a600480360360208110156102dd57600080fd5b50356001600160a01b031661064d565b6102c5610660565b6102c56004803603604081101561030b57600080fd5b50803590602001356106e8565b6102c56004803603602081101561032e57600080fd5b50356001600160a01b0316610771565b6102326108f0565b6102c56004803603602081101561035c57600080fd5b50356108ff565b6102c56004803603604081101561037957600080fd5b8135919081019060408101602082013564010000000081111561039b57600080fd5b8201836020820111156103ad57600080fd5b803590602001918460018302840111640100000000831117156103cf57600080fd5b509092509050610990565b6101d9600480360360208110156103f057600080fd5b5035610a20565b610232610a35565b61041c6004803603602081101561041557600080fd5b5035610a44565b6040805160208082528351818301528351919283929083019185019080838360005b8381101561045657818101518382015260200161043e565b50505050905090810190601f1680156104835780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b61020a610ae5565b6102c5600480360360408110156104af57600080fd5b50803590602001351515610aeb565b61041c600480360360208110156104d457600080fd5b5035610b8b565b6102c5600480360360408110156104f157600080fd5b50803590602001356001600160a01b0316610bf5565b6102c56004803603602081101561051d57600080fd5b5035610c8a565b6102326004803603602081101561053a57600080fd5b5035610d12565b6102c56004803603604081101561055757600080fd5b506001600160a01b0381351690602001351515610d25565b6102c56004803603602081101561058557600080fd5b5035610e72565b600061059f60028363ffffffff610f0016565b92915050565b60009081526007602052604090205490565b6105c860023363ffffffff610f0016565b806105dd57506000546001600160a01b031633145b61062e576040805162461bcd60e51b815260206004820152601a60248201527f4d697373696e672073746f72616765207065726d697373696f6e000000000000604482015290519081900360640190fd5b60008381526005602052604090206106479083836111d2565b50505050565b600061059f60028363ffffffff610f3716565b6000546001600160a01b031633148061068957506001546000546001600160a01b039081169116145b6106da576040805162461bcd60e51b815260206004820152601960248201527f4f776e6572206163636f756e7420697320726571756972656400000000000000604482015290519081900360640190fd5b6000546001600160a01b0316ff5b6106f960023363ffffffff610f0016565b8061070e57506000546001600160a01b031633145b61075f576040805162461bcd60e51b815260206004820152601a60248201527f4d697373696e672073746f72616765207065726d697373696f6e000000000000604482015290519081900360640190fd5b60009182526007602052604090912055565b6000546001600160a01b031633148061079a57506001546000546001600160a01b039081169116145b6107eb576040805162461bcd60e51b815260206004820152601960248201527f4f776e6572206163636f756e7420697320726571756972656400000000000000604482015290519081900360640190fd5b6000546001600160a01b03828116911614156108385760405162461bcd60e51b81526004018080602001828103825260258152602001806112ae6025913960400191505060405180910390fd5b6001600160a01b038116610893576040805162461bcd60e51b815260206004820181905260248201527f4e6577204f776e65722063616e6e6f74206265207a65726f2061646472657373604482015290519081900360640190fd5b600080546001600160a01b0383811673ffffffffffffffffffffffffffffffffffffffff19831681178455604051919092169283917f8934ce4adea8d9ce0d714d2c22b86790e41b7731c84b926fbbdc1d40ff6533c99190a35050565b6001546001600160a01b031681565b61091060023363ffffffff610f0016565b8061092557506000546001600160a01b031633145b610976576040805162461bcd60e51b815260206004820152601a60248201527f4d697373696e672073746f72616765207065726d697373696f6e000000000000604482015290519081900360640190fd5b600081815260056020526040812061098d91611250565b50565b6109a160023363ffffffff610f0016565b806109b657506000546001600160a01b031633145b610a07576040805162461bcd60e51b815260206004820152601a60248201527f4d697373696e672073746f72616765207065726d697373696f6e000000000000604482015290519081900360640190fd5b60008381526006602052604090206106479083836111d2565b60009081526007602052604090205460011490565b6000546001600160a01b031681565b60008181526006602090815260409182902080548351601f6002600019610100600186161502019093169290920491820184900484028101840190945280845260609392830182828015610ad95780601f10610aae57610100808354040283529160200191610ad9565b820191906000526020600020905b815481529060010190602001808311610abc57829003601f168201915b50505050509050919050565b60025481565b610afc60023363ffffffff610f0016565b80610b1157506000546001600160a01b031633145b610b62576040805162461bcd60e51b815260206004820152601a60248201527f4d697373696e672073746f72616765207065726d697373696f6e000000000000604482015290519081900360640190fd5b80610b6e576000610b71565b60015b60009283526007602052604090922060ff90921690915550565b60008181526005602090815260409182902080548351601f6002600019610100600186161502019093169290920491820184900484028101840190945280845260609392830182828015610ad95780601f10610aae57610100808354040283529160200191610ad9565b610c0660023363ffffffff610f0016565b80610c1b57506000546001600160a01b031633145b610c6c576040805162461bcd60e51b815260206004820152601a60248201527f4d697373696e672073746f72616765207065726d697373696f6e000000000000604482015290519081900360640190fd5b6000918252600760205260409091206001600160a01b039091169055565b610c9b60023363ffffffff610f0016565b80610cb057506000546001600160a01b031633145b610d01576040805162461bcd60e51b815260206004820152601a60248201527f4d697373696e672073746f72616765207065726d697373696f6e000000000000604482015290519081900360640190fd5b600090815260076020526040812055565b600061059f60028363ffffffff610f9516565b6000546001600160a01b0316331480610d4e57506001546000546001600160a01b039081169116145b610d9f576040805162461bcd60e51b815260206004820152601960248201527f4f776e6572206163636f756e7420697320726571756972656400000000000000604482015290519081900360640190fd5b8015610e0c57610db660028363ffffffff61101b16565b610e07576040805162461bcd60e51b815260206004820152601e60248201527f4164647265737320616c726561647920686173207065726d697373696f6e0000604482015290519081900360640190fd5b610e6e565b610e1d60028363ffffffff6110cd16565b610e6e576040805162461bcd60e51b815260206004820152601560248201527f4164647265737320646f65736e27742065786973740000000000000000000000604482015290519081900360640190fd5b5050565b610e8360023363ffffffff610f0016565b80610e9857506000546001600160a01b031633145b610ee9576040805162461bcd60e51b815260206004820152601a60248201527f4d697373696e672073746f72616765207065726d697373696f6e000000000000604482015290519081900360640190fd5b600081815260066020526040812061098d91611250565b6001600160a01b038116600090815260018301602052604081205460001901818112801590610f2f5750835481125b949350505050565b60006001600160a01b038216610f50575060001961059f565b6001600160a01b03821660009081526001840160205260408120546000190190811280610f7e575083548112155b15610f8e5760001991505061059f565b9392505050565b6000808212158015610fa75750825482125b610ff8576040805162461bcd60e51b815260206004820152601860248201527f496e646578206f757473696465206f6620626f756e64732e0000000000000000604482015290519081900360640190fd5b50600101600090815260029190910160205260409020546001600160a01b031690565b60006001600160a01b0382166110335750600061059f565b6001600160a01b0382166000908152600184016020526040812054600019019081128015906110625750835481125b1561107157600091505061059f565b5050815460019081018084556001600160a01b03831660008181528386016020908152604080832085905593825260028701905291909120805473ffffffffffffffffffffffffffffffffffffffff1916909117905592915050565b6001600160a01b038116600090815260018084016020526040822054908112806110f75750835481135b1561110657600091505061059f565b835481121561117a57835460009081526002850160208181526040808420546001600160a01b031680855260018901835281852086905585855292909152808320805473ffffffffffffffffffffffffffffffffffffffff19908116909317905586548352909120805490911690556111a6565b60008181526002850160205260409020805473ffffffffffffffffffffffffffffffffffffffff191690555b50506001600160a01b031660009081526001828101602052604082209190915581546000190190915590565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106112135782800160ff19823516178555611240565b82800160010185558215611240579182015b82811115611240578235825591602001919060010190611225565b5061124c929150611290565b5090565b50805460018160011615610100020316600290046000825580601f10611276575061098d565b601f01602090049060005260206000209081019061098d91905b6112aa91905b8082111561124c5760008155600101611296565b9056fe4e6577204f776e65722063616e6e6f74206265207468652063757272656e74206f776e6572a265627a7a72305820864c639d1f68290ccd2296df47a6f7a10f6f08e351a7f4e243242ab5c012f40b64736f6c63430005090032`

// DeployComplianceStorage deploys a new Ethereum contract, binding an instance of ComplianceStorage to it.
func DeployComplianceStorage(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ComplianceStorage, error) {
	parsed, err := abi.JSON(strings.NewReader(ComplianceStorageABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ComplianceStorageBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ComplianceStorage{ComplianceStorageCaller: ComplianceStorageCaller{contract: contract}, ComplianceStorageTransactor: ComplianceStorageTransactor{contract: contract}, ComplianceStorageFilterer: ComplianceStorageFilterer{contract: contract}}, nil
}

// ComplianceStorage is an auto generated Go binding around an Ethereum contract.
type ComplianceStorage struct {
	ComplianceStorageCaller     // Read-only binding to the contract
	ComplianceStorageTransactor // Write-only binding to the contract
	ComplianceStorageFilterer   // Log filterer for contract events
}

// ComplianceStorageCaller is an auto generated read-only Go binding around an Ethereum contract.
type ComplianceStorageCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ComplianceStorageTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ComplianceStorageTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ComplianceStorageFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ComplianceStorageFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ComplianceStorageSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ComplianceStorageSession struct {
	Contract     *ComplianceStorage // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// ComplianceStorageCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ComplianceStorageCallerSession struct {
	Contract *ComplianceStorageCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// ComplianceStorageTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ComplianceStorageTransactorSession struct {
	Contract     *ComplianceStorageTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// ComplianceStorageRaw is an auto generated low-level Go binding around an Ethereum contract.
type ComplianceStorageRaw struct {
	Contract *ComplianceStorage // Generic contract binding to access the raw methods on
}

// ComplianceStorageCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ComplianceStorageCallerRaw struct {
	Contract *ComplianceStorageCaller // Generic read-only contract binding to access the raw methods on
}

// ComplianceStorageTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ComplianceStorageTransactorRaw struct {
	Contract *ComplianceStorageTransactor // Generic write-only contract binding to access the raw methods on
}

// NewComplianceStorage creates a new instance of ComplianceStorage, bound to a specific deployed contract.
func NewComplianceStorage(address common.Address, backend bind.ContractBackend) (*ComplianceStorage, error) {
	contract, err := bindComplianceStorage(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ComplianceStorage{ComplianceStorageCaller: ComplianceStorageCaller{contract: contract}, ComplianceStorageTransactor: ComplianceStorageTransactor{contract: contract}, ComplianceStorageFilterer: ComplianceStorageFilterer{contract: contract}}, nil
}

// NewComplianceStorageCaller creates a new read-only instance of ComplianceStorage, bound to a specific deployed contract.
func NewComplianceStorageCaller(address common.Address, caller bind.ContractCaller) (*ComplianceStorageCaller, error) {
	contract, err := bindComplianceStorage(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ComplianceStorageCaller{contract: contract}, nil
}

// NewComplianceStorageTransactor creates a new write-only instance of ComplianceStorage, bound to a specific deployed contract.
func NewComplianceStorageTransactor(address common.Address, transactor bind.ContractTransactor) (*ComplianceStorageTransactor, error) {
	contract, err := bindComplianceStorage(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ComplianceStorageTransactor{contract: contract}, nil
}

// NewComplianceStorageFilterer creates a new log filterer instance of ComplianceStorage, bound to a specific deployed contract.
func NewComplianceStorageFilterer(address common.Address, filterer bind.ContractFilterer) (*ComplianceStorageFilterer, error) {
	contract, err := bindComplianceStorage(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ComplianceStorageFilterer{contract: contract}, nil
}

// bindComplianceStorage binds a generic wrapper to an already deployed contract.
func bindComplianceStorage(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ComplianceStorageABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ComplianceStorage *ComplianceStorageRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ComplianceStorage.Contract.ComplianceStorageCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ComplianceStorage *ComplianceStorageRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ComplianceStorage.Contract.ComplianceStorageTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ComplianceStorage *ComplianceStorageRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ComplianceStorage.Contract.ComplianceStorageTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ComplianceStorage *ComplianceStorageCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ComplianceStorage.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ComplianceStorage *ComplianceStorageTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ComplianceStorage.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ComplianceStorage *ComplianceStorageTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ComplianceStorage.Contract.contract.Transact(opts, method, params...)
}

// ZEROADDRESS is a free data retrieval call binding the contract method 0x538ba4f9.
//
// Solidity: function ZERO_ADDRESS() constant returns(address)
func (_ComplianceStorage *ComplianceStorageCaller) ZEROADDRESS(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ComplianceStorage.contract.Call(opts, out, "ZERO_ADDRESS")
	return *ret0, err
}

// ZEROADDRESS is a free data retrieval call binding the contract method 0x538ba4f9.
//
// Solidity: function ZERO_ADDRESS() constant returns(address)
func (_ComplianceStorage *ComplianceStorageSession) ZEROADDRESS() (common.Address, error) {
	return _ComplianceStorage.Contract.ZEROADDRESS(&_ComplianceStorage.CallOpts)
}

// ZEROADDRESS is a free data retrieval call binding the contract method 0x538ba4f9.
//
// Solidity: function ZERO_ADDRESS() constant returns(address)
func (_ComplianceStorage *ComplianceStorageCallerSession) ZEROADDRESS() (common.Address, error) {
	return _ComplianceStorage.Contract.ZEROADDRESS(&_ComplianceStorage.CallOpts)
}

// GetAddress is a free data retrieval call binding the contract method 0x21f8a721.
//
// Solidity: function getAddress(k bytes32) constant returns(address)
func (_ComplianceStorage *ComplianceStorageCaller) GetAddress(opts *bind.CallOpts, k [32]byte) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ComplianceStorage.contract.Call(opts, out, "getAddress", k)
	return *ret0, err
}

// GetAddress is a free data retrieval call binding the contract method 0x21f8a721.
//
// Solidity: function getAddress(k bytes32) constant returns(address)
func (_ComplianceStorage *ComplianceStorageSession) GetAddress(k [32]byte) (common.Address, error) {
	return _ComplianceStorage.Contract.GetAddress(&_ComplianceStorage.CallOpts, k)
}

// GetAddress is a free data retrieval call binding the contract method 0x21f8a721.
//
// Solidity: function getAddress(k bytes32) constant returns(address)
func (_ComplianceStorage *ComplianceStorageCallerSession) GetAddress(k [32]byte) (common.Address, error) {
	return _ComplianceStorage.Contract.GetAddress(&_ComplianceStorage.CallOpts, k)
}

// GetBool is a free data retrieval call binding the contract method 0x7ae1cfca.
//
// Solidity: function getBool(k bytes32) constant returns(bool)
func (_ComplianceStorage *ComplianceStorageCaller) GetBool(opts *bind.CallOpts, k [32]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _ComplianceStorage.contract.Call(opts, out, "getBool", k)
	return *ret0, err
}

// GetBool is a free data retrieval call binding the contract method 0x7ae1cfca.
//
// Solidity: function getBool(k bytes32) constant returns(bool)
func (_ComplianceStorage *ComplianceStorageSession) GetBool(k [32]byte) (bool, error) {
	return _ComplianceStorage.Contract.GetBool(&_ComplianceStorage.CallOpts, k)
}

// GetBool is a free data retrieval call binding the contract method 0x7ae1cfca.
//
// Solidity: function getBool(k bytes32) constant returns(bool)
func (_ComplianceStorage *ComplianceStorageCallerSession) GetBool(k [32]byte) (bool, error) {
	return _ComplianceStorage.Contract.GetBool(&_ComplianceStorage.CallOpts, k)
}

// GetBytes is a free data retrieval call binding the contract method 0xc031a180.
//
// Solidity: function getBytes(k bytes32) constant returns(bytes)
func (_ComplianceStorage *ComplianceStorageCaller) GetBytes(opts *bind.CallOpts, k [32]byte) ([]byte, error) {
	var (
		ret0 = new([]byte)
	)
	out := ret0
	err := _ComplianceStorage.contract.Call(opts, out, "getBytes", k)
	return *ret0, err
}

// GetBytes is a free data retrieval call binding the contract method 0xc031a180.
//
// Solidity: function getBytes(k bytes32) constant returns(bytes)
func (_ComplianceStorage *ComplianceStorageSession) GetBytes(k [32]byte) ([]byte, error) {
	return _ComplianceStorage.Contract.GetBytes(&_ComplianceStorage.CallOpts, k)
}

// GetBytes is a free data retrieval call binding the contract method 0xc031a180.
//
// Solidity: function getBytes(k bytes32) constant returns(bytes)
func (_ComplianceStorage *ComplianceStorageCallerSession) GetBytes(k [32]byte) ([]byte, error) {
	return _ComplianceStorage.Contract.GetBytes(&_ComplianceStorage.CallOpts, k)
}

// GetBytes32 is a free data retrieval call binding the contract method 0xa6ed563e.
//
// Solidity: function getBytes32(k bytes32) constant returns(bytes32)
func (_ComplianceStorage *ComplianceStorageCaller) GetBytes32(opts *bind.CallOpts, k [32]byte) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _ComplianceStorage.contract.Call(opts, out, "getBytes32", k)
	return *ret0, err
}

// GetBytes32 is a free data retrieval call binding the contract method 0xa6ed563e.
//
// Solidity: function getBytes32(k bytes32) constant returns(bytes32)
func (_ComplianceStorage *ComplianceStorageSession) GetBytes32(k [32]byte) ([32]byte, error) {
	return _ComplianceStorage.Contract.GetBytes32(&_ComplianceStorage.CallOpts, k)
}

// GetBytes32 is a free data retrieval call binding the contract method 0xa6ed563e.
//
// Solidity: function getBytes32(k bytes32) constant returns(bytes32)
func (_ComplianceStorage *ComplianceStorageCallerSession) GetBytes32(k [32]byte) ([32]byte, error) {
	return _ComplianceStorage.Contract.GetBytes32(&_ComplianceStorage.CallOpts, k)
}

// GetInt256 is a free data retrieval call binding the contract method 0x16c7d1d5.
//
// Solidity: function getInt256(k bytes32) constant returns(int256)
func (_ComplianceStorage *ComplianceStorageCaller) GetInt256(opts *bind.CallOpts, k [32]byte) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ComplianceStorage.contract.Call(opts, out, "getInt256", k)
	return *ret0, err
}

// GetInt256 is a free data retrieval call binding the contract method 0x16c7d1d5.
//
// Solidity: function getInt256(k bytes32) constant returns(int256)
func (_ComplianceStorage *ComplianceStorageSession) GetInt256(k [32]byte) (*big.Int, error) {
	return _ComplianceStorage.Contract.GetInt256(&_ComplianceStorage.CallOpts, k)
}

// GetInt256 is a free data retrieval call binding the contract method 0x16c7d1d5.
//
// Solidity: function getInt256(k bytes32) constant returns(int256)
func (_ComplianceStorage *ComplianceStorageCallerSession) GetInt256(k [32]byte) (*big.Int, error) {
	return _ComplianceStorage.Contract.GetInt256(&_ComplianceStorage.CallOpts, k)
}

// GetString is a free data retrieval call binding the contract method 0x986e791a.
//
// Solidity: function getString(k bytes32) constant returns(string)
func (_ComplianceStorage *ComplianceStorageCaller) GetString(opts *bind.CallOpts, k [32]byte) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _ComplianceStorage.contract.Call(opts, out, "getString", k)
	return *ret0, err
}

// GetString is a free data retrieval call binding the contract method 0x986e791a.
//
// Solidity: function getString(k bytes32) constant returns(string)
func (_ComplianceStorage *ComplianceStorageSession) GetString(k [32]byte) (string, error) {
	return _ComplianceStorage.Contract.GetString(&_ComplianceStorage.CallOpts, k)
}

// GetString is a free data retrieval call binding the contract method 0x986e791a.
//
// Solidity: function getString(k bytes32) constant returns(string)
func (_ComplianceStorage *ComplianceStorageCallerSession) GetString(k [32]byte) (string, error) {
	return _ComplianceStorage.Contract.GetString(&_ComplianceStorage.CallOpts, k)
}

// GetUint256 is a free data retrieval call binding the contract method 0x33598b00.
//
// Solidity: function getUint256(k bytes32) constant returns(uint256)
func (_ComplianceStorage *ComplianceStorageCaller) GetUint256(opts *bind.CallOpts, k [32]byte) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ComplianceStorage.contract.Call(opts, out, "getUint256", k)
	return *ret0, err
}

// GetUint256 is a free data retrieval call binding the contract method 0x33598b00.
//
// Solidity: function getUint256(k bytes32) constant returns(uint256)
func (_ComplianceStorage *ComplianceStorageSession) GetUint256(k [32]byte) (*big.Int, error) {
	return _ComplianceStorage.Contract.GetUint256(&_ComplianceStorage.CallOpts, k)
}

// GetUint256 is a free data retrieval call binding the contract method 0x33598b00.
//
// Solidity: function getUint256(k bytes32) constant returns(uint256)
func (_ComplianceStorage *ComplianceStorageCallerSession) GetUint256(k [32]byte) (*big.Int, error) {
	return _ComplianceStorage.Contract.GetUint256(&_ComplianceStorage.CallOpts, k)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_ComplianceStorage *ComplianceStorageCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ComplianceStorage.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_ComplianceStorage *ComplianceStorageSession) Owner() (common.Address, error) {
	return _ComplianceStorage.Contract.Owner(&_ComplianceStorage.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_ComplianceStorage *ComplianceStorageCallerSession) Owner() (common.Address, error) {
	return _ComplianceStorage.Contract.Owner(&_ComplianceStorage.CallOpts)
}

// PermissionAt is a free data retrieval call binding the contract method 0xea11cd42.
//
// Solidity: function permissionAt(index int256) constant returns(address)
func (_ComplianceStorage *ComplianceStorageCaller) PermissionAt(opts *bind.CallOpts, index *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ComplianceStorage.contract.Call(opts, out, "permissionAt", index)
	return *ret0, err
}

// PermissionAt is a free data retrieval call binding the contract method 0xea11cd42.
//
// Solidity: function permissionAt(index int256) constant returns(address)
func (_ComplianceStorage *ComplianceStorageSession) PermissionAt(index *big.Int) (common.Address, error) {
	return _ComplianceStorage.Contract.PermissionAt(&_ComplianceStorage.CallOpts, index)
}

// PermissionAt is a free data retrieval call binding the contract method 0xea11cd42.
//
// Solidity: function permissionAt(index int256) constant returns(address)
func (_ComplianceStorage *ComplianceStorageCallerSession) PermissionAt(index *big.Int) (common.Address, error) {
	return _ComplianceStorage.Contract.PermissionAt(&_ComplianceStorage.CallOpts, index)
}

// PermissionExists is a free data retrieval call binding the contract method 0x01952ffe.
//
// Solidity: function permissionExists(addr address) constant returns(bool)
func (_ComplianceStorage *ComplianceStorageCaller) PermissionExists(opts *bind.CallOpts, addr common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _ComplianceStorage.contract.Call(opts, out, "permissionExists", addr)
	return *ret0, err
}

// PermissionExists is a free data retrieval call binding the contract method 0x01952ffe.
//
// Solidity: function permissionExists(addr address) constant returns(bool)
func (_ComplianceStorage *ComplianceStorageSession) PermissionExists(addr common.Address) (bool, error) {
	return _ComplianceStorage.Contract.PermissionExists(&_ComplianceStorage.CallOpts, addr)
}

// PermissionExists is a free data retrieval call binding the contract method 0x01952ffe.
//
// Solidity: function permissionExists(addr address) constant returns(bool)
func (_ComplianceStorage *ComplianceStorageCallerSession) PermissionExists(addr common.Address) (bool, error) {
	return _ComplianceStorage.Contract.PermissionExists(&_ComplianceStorage.CallOpts, addr)
}

// PermissionIndexOf is a free data retrieval call binding the contract method 0x3ad8f687.
//
// Solidity: function permissionIndexOf(addr address) constant returns(int256)
func (_ComplianceStorage *ComplianceStorageCaller) PermissionIndexOf(opts *bind.CallOpts, addr common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ComplianceStorage.contract.Call(opts, out, "permissionIndexOf", addr)
	return *ret0, err
}

// PermissionIndexOf is a free data retrieval call binding the contract method 0x3ad8f687.
//
// Solidity: function permissionIndexOf(addr address) constant returns(int256)
func (_ComplianceStorage *ComplianceStorageSession) PermissionIndexOf(addr common.Address) (*big.Int, error) {
	return _ComplianceStorage.Contract.PermissionIndexOf(&_ComplianceStorage.CallOpts, addr)
}

// PermissionIndexOf is a free data retrieval call binding the contract method 0x3ad8f687.
//
// Solidity: function permissionIndexOf(addr address) constant returns(int256)
func (_ComplianceStorage *ComplianceStorageCallerSession) PermissionIndexOf(addr common.Address) (*big.Int, error) {
	return _ComplianceStorage.Contract.PermissionIndexOf(&_ComplianceStorage.CallOpts, addr)
}

// Permissions is a free data retrieval call binding the contract method 0xab8c71c0.
//
// Solidity: function permissions() constant returns(count int256)
func (_ComplianceStorage *ComplianceStorageCaller) Permissions(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ComplianceStorage.contract.Call(opts, out, "permissions")
	return *ret0, err
}

// Permissions is a free data retrieval call binding the contract method 0xab8c71c0.
//
// Solidity: function permissions() constant returns(count int256)
func (_ComplianceStorage *ComplianceStorageSession) Permissions() (*big.Int, error) {
	return _ComplianceStorage.Contract.Permissions(&_ComplianceStorage.CallOpts)
}

// Permissions is a free data retrieval call binding the contract method 0xab8c71c0.
//
// Solidity: function permissions() constant returns(count int256)
func (_ComplianceStorage *ComplianceStorageCallerSession) Permissions() (*big.Int, error) {
	return _ComplianceStorage.Contract.Permissions(&_ComplianceStorage.CallOpts)
}

// DeleteBytes is a paid mutator transaction binding the contract method 0x616b59f6.
//
// Solidity: function deleteBytes(k bytes32) returns()
func (_ComplianceStorage *ComplianceStorageTransactor) DeleteBytes(opts *bind.TransactOpts, k [32]byte) (*types.Transaction, error) {
	return _ComplianceStorage.contract.Transact(opts, "deleteBytes", k)
}

// DeleteBytes is a paid mutator transaction binding the contract method 0x616b59f6.
//
// Solidity: function deleteBytes(k bytes32) returns()
func (_ComplianceStorage *ComplianceStorageSession) DeleteBytes(k [32]byte) (*types.Transaction, error) {
	return _ComplianceStorage.Contract.DeleteBytes(&_ComplianceStorage.TransactOpts, k)
}

// DeleteBytes is a paid mutator transaction binding the contract method 0x616b59f6.
//
// Solidity: function deleteBytes(k bytes32) returns()
func (_ComplianceStorage *ComplianceStorageTransactorSession) DeleteBytes(k [32]byte) (*types.Transaction, error) {
	return _ComplianceStorage.Contract.DeleteBytes(&_ComplianceStorage.TransactOpts, k)
}

// DeleteString is a paid mutator transaction binding the contract method 0xf6bb3cc4.
//
// Solidity: function deleteString(k bytes32) returns()
func (_ComplianceStorage *ComplianceStorageTransactor) DeleteString(opts *bind.TransactOpts, k [32]byte) (*types.Transaction, error) {
	return _ComplianceStorage.contract.Transact(opts, "deleteString", k)
}

// DeleteString is a paid mutator transaction binding the contract method 0xf6bb3cc4.
//
// Solidity: function deleteString(k bytes32) returns()
func (_ComplianceStorage *ComplianceStorageSession) DeleteString(k [32]byte) (*types.Transaction, error) {
	return _ComplianceStorage.Contract.DeleteString(&_ComplianceStorage.TransactOpts, k)
}

// DeleteString is a paid mutator transaction binding the contract method 0xf6bb3cc4.
//
// Solidity: function deleteString(k bytes32) returns()
func (_ComplianceStorage *ComplianceStorageTransactorSession) DeleteString(k [32]byte) (*types.Transaction, error) {
	return _ComplianceStorage.Contract.DeleteString(&_ComplianceStorage.TransactOpts, k)
}

// DeleteValue is a paid mutator transaction binding the contract method 0xddd875d8.
//
// Solidity: function deleteValue(k bytes32) returns()
func (_ComplianceStorage *ComplianceStorageTransactor) DeleteValue(opts *bind.TransactOpts, k [32]byte) (*types.Transaction, error) {
	return _ComplianceStorage.contract.Transact(opts, "deleteValue", k)
}

// DeleteValue is a paid mutator transaction binding the contract method 0xddd875d8.
//
// Solidity: function deleteValue(k bytes32) returns()
func (_ComplianceStorage *ComplianceStorageSession) DeleteValue(k [32]byte) (*types.Transaction, error) {
	return _ComplianceStorage.Contract.DeleteValue(&_ComplianceStorage.TransactOpts, k)
}

// DeleteValue is a paid mutator transaction binding the contract method 0xddd875d8.
//
// Solidity: function deleteValue(k bytes32) returns()
func (_ComplianceStorage *ComplianceStorageTransactorSession) DeleteValue(k [32]byte) (*types.Transaction, error) {
	return _ComplianceStorage.Contract.DeleteValue(&_ComplianceStorage.TransactOpts, k)
}

// Kill is a paid mutator transaction binding the contract method 0x41c0e1b5.
//
// Solidity: function kill() returns()
func (_ComplianceStorage *ComplianceStorageTransactor) Kill(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ComplianceStorage.contract.Transact(opts, "kill")
}

// Kill is a paid mutator transaction binding the contract method 0x41c0e1b5.
//
// Solidity: function kill() returns()
func (_ComplianceStorage *ComplianceStorageSession) Kill() (*types.Transaction, error) {
	return _ComplianceStorage.Contract.Kill(&_ComplianceStorage.TransactOpts)
}

// Kill is a paid mutator transaction binding the contract method 0x41c0e1b5.
//
// Solidity: function kill() returns()
func (_ComplianceStorage *ComplianceStorageTransactorSession) Kill() (*types.Transaction, error) {
	return _ComplianceStorage.Contract.Kill(&_ComplianceStorage.TransactOpts)
}

// SetAddress is a paid mutator transaction binding the contract method 0xca446dd9.
//
// Solidity: function setAddress(k bytes32, v address) returns()
func (_ComplianceStorage *ComplianceStorageTransactor) SetAddress(opts *bind.TransactOpts, k [32]byte, v common.Address) (*types.Transaction, error) {
	return _ComplianceStorage.contract.Transact(opts, "setAddress", k, v)
}

// SetAddress is a paid mutator transaction binding the contract method 0xca446dd9.
//
// Solidity: function setAddress(k bytes32, v address) returns()
func (_ComplianceStorage *ComplianceStorageSession) SetAddress(k [32]byte, v common.Address) (*types.Transaction, error) {
	return _ComplianceStorage.Contract.SetAddress(&_ComplianceStorage.TransactOpts, k, v)
}

// SetAddress is a paid mutator transaction binding the contract method 0xca446dd9.
//
// Solidity: function setAddress(k bytes32, v address) returns()
func (_ComplianceStorage *ComplianceStorageTransactorSession) SetAddress(k [32]byte, v common.Address) (*types.Transaction, error) {
	return _ComplianceStorage.Contract.SetAddress(&_ComplianceStorage.TransactOpts, k, v)
}

// SetBool is a paid mutator transaction binding the contract method 0xabfdcced.
//
// Solidity: function setBool(k bytes32, v bool) returns()
func (_ComplianceStorage *ComplianceStorageTransactor) SetBool(opts *bind.TransactOpts, k [32]byte, v bool) (*types.Transaction, error) {
	return _ComplianceStorage.contract.Transact(opts, "setBool", k, v)
}

// SetBool is a paid mutator transaction binding the contract method 0xabfdcced.
//
// Solidity: function setBool(k bytes32, v bool) returns()
func (_ComplianceStorage *ComplianceStorageSession) SetBool(k [32]byte, v bool) (*types.Transaction, error) {
	return _ComplianceStorage.Contract.SetBool(&_ComplianceStorage.TransactOpts, k, v)
}

// SetBool is a paid mutator transaction binding the contract method 0xabfdcced.
//
// Solidity: function setBool(k bytes32, v bool) returns()
func (_ComplianceStorage *ComplianceStorageTransactorSession) SetBool(k [32]byte, v bool) (*types.Transaction, error) {
	return _ComplianceStorage.Contract.SetBool(&_ComplianceStorage.TransactOpts, k, v)
}

// SetBytes is a paid mutator transaction binding the contract method 0x2e28d084.
//
// Solidity: function setBytes(k bytes32, v bytes) returns()
func (_ComplianceStorage *ComplianceStorageTransactor) SetBytes(opts *bind.TransactOpts, k [32]byte, v []byte) (*types.Transaction, error) {
	return _ComplianceStorage.contract.Transact(opts, "setBytes", k, v)
}

// SetBytes is a paid mutator transaction binding the contract method 0x2e28d084.
//
// Solidity: function setBytes(k bytes32, v bytes) returns()
func (_ComplianceStorage *ComplianceStorageSession) SetBytes(k [32]byte, v []byte) (*types.Transaction, error) {
	return _ComplianceStorage.Contract.SetBytes(&_ComplianceStorage.TransactOpts, k, v)
}

// SetBytes is a paid mutator transaction binding the contract method 0x2e28d084.
//
// Solidity: function setBytes(k bytes32, v bytes) returns()
func (_ComplianceStorage *ComplianceStorageTransactorSession) SetBytes(k [32]byte, v []byte) (*types.Transaction, error) {
	return _ComplianceStorage.Contract.SetBytes(&_ComplianceStorage.TransactOpts, k, v)
}

// SetBytes32 is a paid mutator transaction binding the contract method 0x4e91db08.
//
// Solidity: function setBytes32(k bytes32, v bytes32) returns()
func (_ComplianceStorage *ComplianceStorageTransactor) SetBytes32(opts *bind.TransactOpts, k [32]byte, v [32]byte) (*types.Transaction, error) {
	return _ComplianceStorage.contract.Transact(opts, "setBytes32", k, v)
}

// SetBytes32 is a paid mutator transaction binding the contract method 0x4e91db08.
//
// Solidity: function setBytes32(k bytes32, v bytes32) returns()
func (_ComplianceStorage *ComplianceStorageSession) SetBytes32(k [32]byte, v [32]byte) (*types.Transaction, error) {
	return _ComplianceStorage.Contract.SetBytes32(&_ComplianceStorage.TransactOpts, k, v)
}

// SetBytes32 is a paid mutator transaction binding the contract method 0x4e91db08.
//
// Solidity: function setBytes32(k bytes32, v bytes32) returns()
func (_ComplianceStorage *ComplianceStorageTransactorSession) SetBytes32(k [32]byte, v [32]byte) (*types.Transaction, error) {
	return _ComplianceStorage.Contract.SetBytes32(&_ComplianceStorage.TransactOpts, k, v)
}

// SetInt256 is a paid mutator transaction binding the contract method 0xf64f41db.
//
// Solidity: function setInt256(k bytes32, v int256) returns()
func (_ComplianceStorage *ComplianceStorageTransactor) SetInt256(opts *bind.TransactOpts, k [32]byte, v *big.Int) (*types.Transaction, error) {
	return _ComplianceStorage.contract.Transact(opts, "setInt256", k, v)
}

// SetInt256 is a paid mutator transaction binding the contract method 0xf64f41db.
//
// Solidity: function setInt256(k bytes32, v int256) returns()
func (_ComplianceStorage *ComplianceStorageSession) SetInt256(k [32]byte, v *big.Int) (*types.Transaction, error) {
	return _ComplianceStorage.Contract.SetInt256(&_ComplianceStorage.TransactOpts, k, v)
}

// SetInt256 is a paid mutator transaction binding the contract method 0xf64f41db.
//
// Solidity: function setInt256(k bytes32, v int256) returns()
func (_ComplianceStorage *ComplianceStorageTransactorSession) SetInt256(k [32]byte, v *big.Int) (*types.Transaction, error) {
	return _ComplianceStorage.Contract.SetInt256(&_ComplianceStorage.TransactOpts, k, v)
}

// SetPermission is a paid mutator transaction binding the contract method 0xec6263c0.
//
// Solidity: function setPermission(addr address, grant bool) returns()
func (_ComplianceStorage *ComplianceStorageTransactor) SetPermission(opts *bind.TransactOpts, addr common.Address, grant bool) (*types.Transaction, error) {
	return _ComplianceStorage.contract.Transact(opts, "setPermission", addr, grant)
}

// SetPermission is a paid mutator transaction binding the contract method 0xec6263c0.
//
// Solidity: function setPermission(addr address, grant bool) returns()
func (_ComplianceStorage *ComplianceStorageSession) SetPermission(addr common.Address, grant bool) (*types.Transaction, error) {
	return _ComplianceStorage.Contract.SetPermission(&_ComplianceStorage.TransactOpts, addr, grant)
}

// SetPermission is a paid mutator transaction binding the contract method 0xec6263c0.
//
// Solidity: function setPermission(addr address, grant bool) returns()
func (_ComplianceStorage *ComplianceStorageTransactorSession) SetPermission(addr common.Address, grant bool) (*types.Transaction, error) {
	return _ComplianceStorage.Contract.SetPermission(&_ComplianceStorage.TransactOpts, addr, grant)
}

// SetString is a paid mutator transaction binding the contract method 0x6e899550.
//
// Solidity: function setString(k bytes32, v string) returns()
func (_ComplianceStorage *ComplianceStorageTransactor) SetString(opts *bind.TransactOpts, k [32]byte, v string) (*types.Transaction, error) {
	return _ComplianceStorage.contract.Transact(opts, "setString", k, v)
}

// SetString is a paid mutator transaction binding the contract method 0x6e899550.
//
// Solidity: function setString(k bytes32, v string) returns()
func (_ComplianceStorage *ComplianceStorageSession) SetString(k [32]byte, v string) (*types.Transaction, error) {
	return _ComplianceStorage.Contract.SetString(&_ComplianceStorage.TransactOpts, k, v)
}

// SetString is a paid mutator transaction binding the contract method 0x6e899550.
//
// Solidity: function setString(k bytes32, v string) returns()
func (_ComplianceStorage *ComplianceStorageTransactorSession) SetString(k [32]byte, v string) (*types.Transaction, error) {
	return _ComplianceStorage.Contract.SetString(&_ComplianceStorage.TransactOpts, k, v)
}

// SetUint256 is a paid mutator transaction binding the contract method 0x4f3029c2.
//
// Solidity: function setUint256(k bytes32, v uint256) returns()
func (_ComplianceStorage *ComplianceStorageTransactor) SetUint256(opts *bind.TransactOpts, k [32]byte, v *big.Int) (*types.Transaction, error) {
	return _ComplianceStorage.contract.Transact(opts, "setUint256", k, v)
}

// SetUint256 is a paid mutator transaction binding the contract method 0x4f3029c2.
//
// Solidity: function setUint256(k bytes32, v uint256) returns()
func (_ComplianceStorage *ComplianceStorageSession) SetUint256(k [32]byte, v *big.Int) (*types.Transaction, error) {
	return _ComplianceStorage.Contract.SetUint256(&_ComplianceStorage.TransactOpts, k, v)
}

// SetUint256 is a paid mutator transaction binding the contract method 0x4f3029c2.
//
// Solidity: function setUint256(k bytes32, v uint256) returns()
func (_ComplianceStorage *ComplianceStorageTransactorSession) SetUint256(k [32]byte, v *big.Int) (*types.Transaction, error) {
	return _ComplianceStorage.Contract.SetUint256(&_ComplianceStorage.TransactOpts, k, v)
}

// TransferOwner is a paid mutator transaction binding the contract method 0x4fb2e45d.
//
// Solidity: function transferOwner(newOwner address) returns()
func (_ComplianceStorage *ComplianceStorageTransactor) TransferOwner(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _ComplianceStorage.contract.Transact(opts, "transferOwner", newOwner)
}

// TransferOwner is a paid mutator transaction binding the contract method 0x4fb2e45d.
//
// Solidity: function transferOwner(newOwner address) returns()
func (_ComplianceStorage *ComplianceStorageSession) TransferOwner(newOwner common.Address) (*types.Transaction, error) {
	return _ComplianceStorage.Contract.TransferOwner(&_ComplianceStorage.TransactOpts, newOwner)
}

// TransferOwner is a paid mutator transaction binding the contract method 0x4fb2e45d.
//
// Solidity: function transferOwner(newOwner address) returns()
func (_ComplianceStorage *ComplianceStorageTransactorSession) TransferOwner(newOwner common.Address) (*types.Transaction, error) {
	return _ComplianceStorage.Contract.TransferOwner(&_ComplianceStorage.TransactOpts, newOwner)
}

// ComplianceStorageOwnerTransferredIterator is returned from FilterOwnerTransferred and is used to iterate over the raw logs and unpacked data for OwnerTransferred events raised by the ComplianceStorage contract.
type ComplianceStorageOwnerTransferredIterator struct {
	Event *ComplianceStorageOwnerTransferred // Event containing the contract specifics and raw log

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
func (it *ComplianceStorageOwnerTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ComplianceStorageOwnerTransferred)
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
		it.Event = new(ComplianceStorageOwnerTransferred)
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
func (it *ComplianceStorageOwnerTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ComplianceStorageOwnerTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ComplianceStorageOwnerTransferred represents a OwnerTransferred event raised by the ComplianceStorage contract.
type ComplianceStorageOwnerTransferred struct {
	OldOwner common.Address
	NewOwner common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOwnerTransferred is a free log retrieval operation binding the contract event 0x8934ce4adea8d9ce0d714d2c22b86790e41b7731c84b926fbbdc1d40ff6533c9.
//
// Solidity: e OwnerTransferred(oldOwner indexed address, newOwner indexed address)
func (_ComplianceStorage *ComplianceStorageFilterer) FilterOwnerTransferred(opts *bind.FilterOpts, oldOwner []common.Address, newOwner []common.Address) (*ComplianceStorageOwnerTransferredIterator, error) {

	var oldOwnerRule []interface{}
	for _, oldOwnerItem := range oldOwner {
		oldOwnerRule = append(oldOwnerRule, oldOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ComplianceStorage.contract.FilterLogs(opts, "OwnerTransferred", oldOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ComplianceStorageOwnerTransferredIterator{contract: _ComplianceStorage.contract, event: "OwnerTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnerTransferred is a free log subscription operation binding the contract event 0x8934ce4adea8d9ce0d714d2c22b86790e41b7731c84b926fbbdc1d40ff6533c9.
//
// Solidity: e OwnerTransferred(oldOwner indexed address, newOwner indexed address)
func (_ComplianceStorage *ComplianceStorageFilterer) WatchOwnerTransferred(opts *bind.WatchOpts, sink chan<- *ComplianceStorageOwnerTransferred, oldOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var oldOwnerRule []interface{}
	for _, oldOwnerItem := range oldOwner {
		oldOwnerRule = append(oldOwnerRule, oldOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ComplianceStorage.contract.WatchLogs(opts, "OwnerTransferred", oldOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ComplianceStorageOwnerTransferred)
				if err := _ComplianceStorage.contract.UnpackLog(event, "OwnerTransferred", log); err != nil {
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
