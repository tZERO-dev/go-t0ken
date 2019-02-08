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

// BrokerDealerABI is the input ABI used to generate the binding from.
const BrokerDealerABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"brokerDealer\",\"type\":\"address\"}],\"name\":\"add\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"locked\",\"type\":\"bool\"}],\"name\":\"setLocked\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"brokerDealer\",\"type\":\"address\"}],\"name\":\"remove\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"kill\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"brokerDealer\",\"type\":\"address\"},{\"name\":\"account\",\"type\":\"address\"}],\"name\":\"addAccount\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwner\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"s\",\"type\":\"address\"}],\"name\":\"setStorage\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"store\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isLocked\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"brokerDealer\",\"type\":\"address\"},{\"name\":\"frozen\",\"type\":\"bool\"}],\"name\":\"setFrozen\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"account\",\"type\":\"address\"}],\"name\":\"removeAccount\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"oldOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnerTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"brokerDealer\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"BrokerDealerAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"brokerDealer\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"BrokerDealerRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"brokerDealer\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"frozen\",\"type\":\"bool\"},{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"BrokerDealerFrozen\",\"type\":\"event\"}]"

// BrokerDealerBin is the compiled bytecode used for deploying new contracts.
const BrokerDealerBin = `60806040526000805460a060020a60ff0219600160a060020a031990911633171690556118fa806100316000396000f3fe608060405234801561001057600080fd5b50600436106100f1576000357c0100000000000000000000000000000000000000000000000000000000900480638da5cb5b1161009e578063a4e2d63411610078578063a4e2d63414610211578063ac869cd81461022d578063c4740a951461025b576100f1565b80638da5cb5b146101bf5780639137c1a7146101e3578063975057e714610209576100f1565b806341c0e1b5116100cf57806341c0e1b51461016357806349b51b171461016b5780634fb2e45d14610199576100f1565b80630a3b0a4f146100f6578063211e28b61461011e57806329092d0e1461013d575b600080fd5b61011c6004803603602081101561010c57600080fd5b5035600160a060020a0316610281565b005b61011c6004803603602081101561013457600080fd5b50351515610527565b61011c6004803603602081101561015357600080fd5b5035600160a060020a0316610641565b61011c6109b8565b61011c6004803603604081101561018157600080fd5b50600160a060020a0381358116916020013516610a2d565b61011c600480360360208110156101af57600080fd5b5035600160a060020a0316610e73565b6101c7610ff1565b60408051600160a060020a039092168252519081900360200190f35b61011c600480360360208110156101f957600080fd5b5035600160a060020a0316611000565b6101c7611096565b6102196110a5565b604080519115158252519081900360200190f35b61011c6004803603604081101561024357600080fd5b50600160a060020a03813516906020013515156110c6565b61011c6004803603602081101561027157600080fd5b5035600160a060020a031661144c565b60018054604080517f0a22ee73000000000000000000000000000000000000000000000000000000008152336004820152602481019390935251600160a060020a0390911691630a22ee73916044808301926020929190829003018186803b1580156102ec57600080fd5b505afa158015610300573d6000803e3d6000fd5b505050506040513d602081101561031657600080fd5b505115156103735760408051600080516020611865833981519152815260206004820152601a60248201527f437573746f6469616e2061646472657373207265717569726564000000000000604482015290519081900360640190fd5b600154604080517f624522f90000000000000000000000000000000000000000000000000000000081523360048201529051600160a060020a039092169163624522f991602480820192602092909190829003018186803b1580156103d757600080fd5b505afa1580156103eb573d6000803e3d6000fd5b505050506040513d602081101561040157600080fd5b50511561045d5760408051600080516020611865833981519152815260206004820152601360248201527f437573746f6469616e2069732066726f7a656e00000000000000000000000000604482015290519081900360640190fd5b600154604080517f116c92b7000000000000000000000000000000000000000000000000000000008152600160a060020a03848116600483015260036024830152600060448301819052336064840152925193169263116c92b79260848084019391929182900301818387803b1580156104d657600080fd5b505af11580156104ea573d6000803e3d6000fd5b5050604051339250600160a060020a03841691507ff51b1c4e2953dbc0185cc885a05dcf6fe8de20cf5e8e0c32d7c32d76e45f114b90600090a350565b600054600160a060020a0316331461058e5760408051600080516020611865833981519152815260206004820152601960248201527f4f776e6572206163636f756e7420697320726571756972656400000000000000604482015290519081900360640190fd5b60005460ff7401000000000000000000000000000000000000000090910416151581151514156105f75760405160008051602061186583398151915281526004018080602001828103825260288152602001806118a76028913960400191505060405180910390fd5b6000805491151574010000000000000000000000000000000000000000027fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff909216919091179055565b60018054604080517f0a22ee730000000000000000000000000000000000000000000000000000000081523360048201526024810193909352518392600160a060020a0390921691630a22ee73916044808301926020929190829003018186803b1580156106ae57600080fd5b505afa1580156106c2573d6000803e3d6000fd5b505050506040513d60208110156106d857600080fd5b505115156107355760408051600080516020611865833981519152815260206004820152601260248201527f437573746f6469616e2072657175697265640000000000000000000000000000604482015290519081900360640190fd5b600154604080517f31aaa74a000000000000000000000000000000000000000000000000000000008152600160a060020a038481166004830152915191909216916331aaa74a916024808301926020929190829003018186803b15801561079b57600080fd5b505afa1580156107af573d6000803e3d6000fd5b505050506040513d60208110156107c557600080fd5b5051600160a060020a031633146108155760405160008051602061186583398151915281526004018080602001828103825260228152602001806118856022913960400191505060405180910390fd5b600154604080517f624522f90000000000000000000000000000000000000000000000000000000081523360048201529051600160a060020a039092169163624522f991602480820192602092909190829003018186803b15801561087957600080fd5b505afa15801561088d573d6000803e3d6000fd5b505050506040513d60208110156108a357600080fd5b5051156108ff5760408051600080516020611865833981519152815260206004820152601360248201527f437573746f6469616e2069732066726f7a656e00000000000000000000000000604482015290519081900360640190fd5b600154604080517fc4740a95000000000000000000000000000000000000000000000000000000008152600160a060020a0385811660048301529151919092169163c4740a9591602480830192600092919082900301818387803b15801561096657600080fd5b505af115801561097a573d6000803e3d6000fd5b5050604051339250600160a060020a03851691507fdb9237c04fdf8d1a83239f5f1d4e6663f5d36da3b652c728cffa1c21ae6d9f2290600090a35050565b600054600160a060020a03163314610a1f5760408051600080516020611865833981519152815260206004820152601960248201527f4f776e6572206163636f756e7420697320726571756972656400000000000000604482015290519081900360640190fd5b600054600160a060020a0316ff5b600154604080517f75cd51ed000000000000000000000000000000000000000000000000000000008152600160a060020a0380851660048301529151849392909216916375cd51ed91602480820192602092909190829003018186803b158015610a9657600080fd5b505afa158015610aaa573d6000803e3d6000fd5b505050506040513d6020811015610ac057600080fd5b505115610b1c5760408051600080516020611865833981519152815260206004820152601660248201527f4163636f756e7420616c72656164792065786973747300000000000000000000604482015290519081900360640190fd5b60018054604080517f0a22ee730000000000000000000000000000000000000000000000000000000081523360048201526024810193909352518592600160a060020a0390921691630a22ee73916044808301926020929190829003018186803b158015610b8957600080fd5b505afa158015610b9d573d6000803e3d6000fd5b505050506040513d6020811015610bb357600080fd5b50511515610c105760408051600080516020611865833981519152815260206004820152601260248201527f437573746f6469616e2072657175697265640000000000000000000000000000604482015290519081900360640190fd5b600154604080517f31aaa74a000000000000000000000000000000000000000000000000000000008152600160a060020a038481166004830152915191909216916331aaa74a916024808301926020929190829003018186803b158015610c7657600080fd5b505afa158015610c8a573d6000803e3d6000fd5b505050506040513d6020811015610ca057600080fd5b5051600160a060020a03163314610cf05760405160008051602061186583398151915281526004018080602001828103825260228152602001806118856022913960400191505060405180910390fd5b600154604080517f624522f90000000000000000000000000000000000000000000000000000000081523360048201529051600160a060020a039092169163624522f991602480820192602092909190829003018186803b158015610d5457600080fd5b505afa158015610d68573d6000803e3d6000fd5b505050506040513d6020811015610d7e57600080fd5b505115610dda5760408051600080516020611865833981519152815260206004820152601360248201527f437573746f6469616e2069732066726f7a656e00000000000000000000000000604482015290519081900360640190fd5b600154604080517f116c92b7000000000000000000000000000000000000000000000000000000008152600160a060020a038681166004830152600260248301526000604483018190528882166064840152925193169263116c92b79260848084019391929182900301818387803b158015610e5557600080fd5b505af1158015610e69573d6000803e3d6000fd5b5050505050505050565b600054600160a060020a03163314610eda5760408051600080516020611865833981519152815260206004820152601960248201527f4f776e6572206163636f756e7420697320726571756972656400000000000000604482015290519081900360640190fd5b600054600160a060020a0382811691161415610f2f5760405160008051602061186583398151915281526004018080602001828103825260258152602001806118406025913960400191505060405180910390fd5b600160a060020a0381161515610f945760408051600080516020611865833981519152815260206004820181905260248201527f4e6577204f776e65722063616e6e6f74206265207a65726f2061646472657373604482015290519081900360640190fd5b60008054600160a060020a0383811673ffffffffffffffffffffffffffffffffffffffff19831681178455604051919092169283917f8934ce4adea8d9ce0d714d2c22b86790e41b7731c84b926fbbdc1d40ff6533c99190a35050565b600054600160a060020a031681565b600054600160a060020a031633146110675760408051600080516020611865833981519152815260206004820152601960248201527f4f776e6572206163636f756e7420697320726571756972656400000000000000604482015290519081900360640190fd5b6001805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b600154600160a060020a031681565b60005474010000000000000000000000000000000000000000900460ff1681565b60018054604080517f0a22ee730000000000000000000000000000000000000000000000000000000081523360048201526024810193909352518492600160a060020a0390921691630a22ee73916044808301926020929190829003018186803b15801561113357600080fd5b505afa158015611147573d6000803e3d6000fd5b505050506040513d602081101561115d57600080fd5b505115156111ba5760408051600080516020611865833981519152815260206004820152601260248201527f437573746f6469616e2072657175697265640000000000000000000000000000604482015290519081900360640190fd5b600154604080517f31aaa74a000000000000000000000000000000000000000000000000000000008152600160a060020a038481166004830152915191909216916331aaa74a916024808301926020929190829003018186803b15801561122057600080fd5b505afa158015611234573d6000803e3d6000fd5b505050506040513d602081101561124a57600080fd5b5051600160a060020a0316331461129a5760405160008051602061186583398151915281526004018080602001828103825260228152602001806118856022913960400191505060405180910390fd5b600154604080517f624522f90000000000000000000000000000000000000000000000000000000081523360048201529051600160a060020a039092169163624522f991602480820192602092909190829003018186803b1580156112fe57600080fd5b505afa158015611312573d6000803e3d6000fd5b505050506040513d602081101561132857600080fd5b5051156113845760408051600080516020611865833981519152815260206004820152601360248201527f437573746f6469616e2069732066726f7a656e00000000000000000000000000604482015290519081900360640190fd5b60018054604080517fcbe5404f000000000000000000000000000000000000000000000000000000008152600160a060020a03878116600483015260248201949094529051929091169163cbe5404f9160448082019260009290919082900301818387803b1580156113f557600080fd5b505af1158015611409573d6000803e3d6000fd5b50506040513392508415159150600160a060020a038616907fbc170ec10277584572ef4c969f53616db496355c027a69325d121126add2a68190600090a4505050565b600154604080517f0a22ee73000000000000000000000000000000000000000000000000000000008152600160a060020a03808516600483015260026024830152915184939290921691630a22ee7391604480820192602092909190829003018186803b1580156114bc57600080fd5b505afa1580156114d0573d6000803e3d6000fd5b505050506040513d60208110156114e657600080fd5b505115156115435760408051600080516020611865833981519152815260206004820152601760248201527f4e6f74206120637573746f6469616c206163636f756e74000000000000000000604482015290519081900360640190fd5b600154604080517f31aaa74a000000000000000000000000000000000000000000000000000000008152600160a060020a038481166004830152915160009392909216916331aaa74a91602480820192602092909190829003018186803b1580156115ad57600080fd5b505afa1580156115c1573d6000803e3d6000fd5b505050506040513d60208110156115d757600080fd5b5051600154604080517f31aaa74a000000000000000000000000000000000000000000000000000000008152600160a060020a03808516600483015291519394509116916331aaa74a91602480820192602092909190829003018186803b15801561164157600080fd5b505afa158015611655573d6000803e3d6000fd5b505050506040513d602081101561166b57600080fd5b5051600160a060020a031633146116d15760408051600080516020611865833981519152815260206004820152601c60248201527f4163636f756e74277320637573746f6469616e20726571757269656400000000604482015290519081900360640190fd5b600154604080517f624522f90000000000000000000000000000000000000000000000000000000081523360048201529051600160a060020a039092169163624522f991602480820192602092909190829003018186803b15801561173557600080fd5b505afa158015611749573d6000803e3d6000fd5b505050506040513d602081101561175f57600080fd5b5051156117bb5760408051600080516020611865833981519152815260206004820152601360248201527f437573746f6469616e2069732066726f7a656e00000000000000000000000000604482015290519081900360640190fd5b600154604080517fc4740a95000000000000000000000000000000000000000000000000000000008152600160a060020a0386811660048301529151919092169163c4740a9591602480830192600092919082900301818387803b15801561182257600080fd5b505af1158015611836573d6000803e3d6000fd5b5050505050505056fe4e6577204f776e65722063616e6e6f74206265207468652063757272656e74206f776e657208c379a00000000000000000000000000000000000000000000000000000000042726f6b65722d4465616c6572277320637573746f6469616e207265717569726564436f6e747261637420616c726561647920696e20726571756573746564206c6f636b207374617465a165627a7a72305820de041f8c6a7617b048eaccd99607fa5a4e11daa2e2ad10e8dc0520cfa520179f0029`

// DeployBrokerDealer deploys a new Ethereum contract, binding an instance of BrokerDealer to it.
func DeployBrokerDealer(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *BrokerDealer, error) {
	parsed, err := abi.JSON(strings.NewReader(BrokerDealerABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(BrokerDealerBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &BrokerDealer{BrokerDealerCaller: BrokerDealerCaller{contract: contract}, BrokerDealerTransactor: BrokerDealerTransactor{contract: contract}, BrokerDealerFilterer: BrokerDealerFilterer{contract: contract}}, nil
}

// BrokerDealer is an auto generated Go binding around an Ethereum contract.
type BrokerDealer struct {
	BrokerDealerCaller     // Read-only binding to the contract
	BrokerDealerTransactor // Write-only binding to the contract
	BrokerDealerFilterer   // Log filterer for contract events
}

// BrokerDealerCaller is an auto generated read-only Go binding around an Ethereum contract.
type BrokerDealerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BrokerDealerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BrokerDealerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BrokerDealerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BrokerDealerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BrokerDealerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BrokerDealerSession struct {
	Contract     *BrokerDealer     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BrokerDealerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BrokerDealerCallerSession struct {
	Contract *BrokerDealerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// BrokerDealerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BrokerDealerTransactorSession struct {
	Contract     *BrokerDealerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// BrokerDealerRaw is an auto generated low-level Go binding around an Ethereum contract.
type BrokerDealerRaw struct {
	Contract *BrokerDealer // Generic contract binding to access the raw methods on
}

// BrokerDealerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BrokerDealerCallerRaw struct {
	Contract *BrokerDealerCaller // Generic read-only contract binding to access the raw methods on
}

// BrokerDealerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BrokerDealerTransactorRaw struct {
	Contract *BrokerDealerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBrokerDealer creates a new instance of BrokerDealer, bound to a specific deployed contract.
func NewBrokerDealer(address common.Address, backend bind.ContractBackend) (*BrokerDealer, error) {
	contract, err := bindBrokerDealer(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BrokerDealer{BrokerDealerCaller: BrokerDealerCaller{contract: contract}, BrokerDealerTransactor: BrokerDealerTransactor{contract: contract}, BrokerDealerFilterer: BrokerDealerFilterer{contract: contract}}, nil
}

// NewBrokerDealerCaller creates a new read-only instance of BrokerDealer, bound to a specific deployed contract.
func NewBrokerDealerCaller(address common.Address, caller bind.ContractCaller) (*BrokerDealerCaller, error) {
	contract, err := bindBrokerDealer(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BrokerDealerCaller{contract: contract}, nil
}

// NewBrokerDealerTransactor creates a new write-only instance of BrokerDealer, bound to a specific deployed contract.
func NewBrokerDealerTransactor(address common.Address, transactor bind.ContractTransactor) (*BrokerDealerTransactor, error) {
	contract, err := bindBrokerDealer(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BrokerDealerTransactor{contract: contract}, nil
}

// NewBrokerDealerFilterer creates a new log filterer instance of BrokerDealer, bound to a specific deployed contract.
func NewBrokerDealerFilterer(address common.Address, filterer bind.ContractFilterer) (*BrokerDealerFilterer, error) {
	contract, err := bindBrokerDealer(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BrokerDealerFilterer{contract: contract}, nil
}

// bindBrokerDealer binds a generic wrapper to an already deployed contract.
func bindBrokerDealer(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BrokerDealerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BrokerDealer *BrokerDealerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _BrokerDealer.Contract.BrokerDealerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BrokerDealer *BrokerDealerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BrokerDealer.Contract.BrokerDealerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BrokerDealer *BrokerDealerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BrokerDealer.Contract.BrokerDealerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BrokerDealer *BrokerDealerCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _BrokerDealer.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BrokerDealer *BrokerDealerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BrokerDealer.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BrokerDealer *BrokerDealerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BrokerDealer.Contract.contract.Transact(opts, method, params...)
}

// IsLocked is a free data retrieval call binding the contract method 0xa4e2d634.
//
// Solidity: function isLocked() constant returns(bool)
func (_BrokerDealer *BrokerDealerCaller) IsLocked(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _BrokerDealer.contract.Call(opts, out, "isLocked")
	return *ret0, err
}

// IsLocked is a free data retrieval call binding the contract method 0xa4e2d634.
//
// Solidity: function isLocked() constant returns(bool)
func (_BrokerDealer *BrokerDealerSession) IsLocked() (bool, error) {
	return _BrokerDealer.Contract.IsLocked(&_BrokerDealer.CallOpts)
}

// IsLocked is a free data retrieval call binding the contract method 0xa4e2d634.
//
// Solidity: function isLocked() constant returns(bool)
func (_BrokerDealer *BrokerDealerCallerSession) IsLocked() (bool, error) {
	return _BrokerDealer.Contract.IsLocked(&_BrokerDealer.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_BrokerDealer *BrokerDealerCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _BrokerDealer.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_BrokerDealer *BrokerDealerSession) Owner() (common.Address, error) {
	return _BrokerDealer.Contract.Owner(&_BrokerDealer.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_BrokerDealer *BrokerDealerCallerSession) Owner() (common.Address, error) {
	return _BrokerDealer.Contract.Owner(&_BrokerDealer.CallOpts)
}

// Store is a free data retrieval call binding the contract method 0x975057e7.
//
// Solidity: function store() constant returns(address)
func (_BrokerDealer *BrokerDealerCaller) Store(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _BrokerDealer.contract.Call(opts, out, "store")
	return *ret0, err
}

// Store is a free data retrieval call binding the contract method 0x975057e7.
//
// Solidity: function store() constant returns(address)
func (_BrokerDealer *BrokerDealerSession) Store() (common.Address, error) {
	return _BrokerDealer.Contract.Store(&_BrokerDealer.CallOpts)
}

// Store is a free data retrieval call binding the contract method 0x975057e7.
//
// Solidity: function store() constant returns(address)
func (_BrokerDealer *BrokerDealerCallerSession) Store() (common.Address, error) {
	return _BrokerDealer.Contract.Store(&_BrokerDealer.CallOpts)
}

// Add is a paid mutator transaction binding the contract method 0x0a3b0a4f.
//
// Solidity: function add(brokerDealer address) returns()
func (_BrokerDealer *BrokerDealerTransactor) Add(opts *bind.TransactOpts, brokerDealer common.Address) (*types.Transaction, error) {
	return _BrokerDealer.contract.Transact(opts, "add", brokerDealer)
}

// Add is a paid mutator transaction binding the contract method 0x0a3b0a4f.
//
// Solidity: function add(brokerDealer address) returns()
func (_BrokerDealer *BrokerDealerSession) Add(brokerDealer common.Address) (*types.Transaction, error) {
	return _BrokerDealer.Contract.Add(&_BrokerDealer.TransactOpts, brokerDealer)
}

// Add is a paid mutator transaction binding the contract method 0x0a3b0a4f.
//
// Solidity: function add(brokerDealer address) returns()
func (_BrokerDealer *BrokerDealerTransactorSession) Add(brokerDealer common.Address) (*types.Transaction, error) {
	return _BrokerDealer.Contract.Add(&_BrokerDealer.TransactOpts, brokerDealer)
}

// AddAccount is a paid mutator transaction binding the contract method 0x49b51b17.
//
// Solidity: function addAccount(brokerDealer address, account address) returns()
func (_BrokerDealer *BrokerDealerTransactor) AddAccount(opts *bind.TransactOpts, brokerDealer common.Address, account common.Address) (*types.Transaction, error) {
	return _BrokerDealer.contract.Transact(opts, "addAccount", brokerDealer, account)
}

// AddAccount is a paid mutator transaction binding the contract method 0x49b51b17.
//
// Solidity: function addAccount(brokerDealer address, account address) returns()
func (_BrokerDealer *BrokerDealerSession) AddAccount(brokerDealer common.Address, account common.Address) (*types.Transaction, error) {
	return _BrokerDealer.Contract.AddAccount(&_BrokerDealer.TransactOpts, brokerDealer, account)
}

// AddAccount is a paid mutator transaction binding the contract method 0x49b51b17.
//
// Solidity: function addAccount(brokerDealer address, account address) returns()
func (_BrokerDealer *BrokerDealerTransactorSession) AddAccount(brokerDealer common.Address, account common.Address) (*types.Transaction, error) {
	return _BrokerDealer.Contract.AddAccount(&_BrokerDealer.TransactOpts, brokerDealer, account)
}

// Kill is a paid mutator transaction binding the contract method 0x41c0e1b5.
//
// Solidity: function kill() returns()
func (_BrokerDealer *BrokerDealerTransactor) Kill(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BrokerDealer.contract.Transact(opts, "kill")
}

// Kill is a paid mutator transaction binding the contract method 0x41c0e1b5.
//
// Solidity: function kill() returns()
func (_BrokerDealer *BrokerDealerSession) Kill() (*types.Transaction, error) {
	return _BrokerDealer.Contract.Kill(&_BrokerDealer.TransactOpts)
}

// Kill is a paid mutator transaction binding the contract method 0x41c0e1b5.
//
// Solidity: function kill() returns()
func (_BrokerDealer *BrokerDealerTransactorSession) Kill() (*types.Transaction, error) {
	return _BrokerDealer.Contract.Kill(&_BrokerDealer.TransactOpts)
}

// Remove is a paid mutator transaction binding the contract method 0x29092d0e.
//
// Solidity: function remove(brokerDealer address) returns()
func (_BrokerDealer *BrokerDealerTransactor) Remove(opts *bind.TransactOpts, brokerDealer common.Address) (*types.Transaction, error) {
	return _BrokerDealer.contract.Transact(opts, "remove", brokerDealer)
}

// Remove is a paid mutator transaction binding the contract method 0x29092d0e.
//
// Solidity: function remove(brokerDealer address) returns()
func (_BrokerDealer *BrokerDealerSession) Remove(brokerDealer common.Address) (*types.Transaction, error) {
	return _BrokerDealer.Contract.Remove(&_BrokerDealer.TransactOpts, brokerDealer)
}

// Remove is a paid mutator transaction binding the contract method 0x29092d0e.
//
// Solidity: function remove(brokerDealer address) returns()
func (_BrokerDealer *BrokerDealerTransactorSession) Remove(brokerDealer common.Address) (*types.Transaction, error) {
	return _BrokerDealer.Contract.Remove(&_BrokerDealer.TransactOpts, brokerDealer)
}

// RemoveAccount is a paid mutator transaction binding the contract method 0xc4740a95.
//
// Solidity: function removeAccount(account address) returns()
func (_BrokerDealer *BrokerDealerTransactor) RemoveAccount(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _BrokerDealer.contract.Transact(opts, "removeAccount", account)
}

// RemoveAccount is a paid mutator transaction binding the contract method 0xc4740a95.
//
// Solidity: function removeAccount(account address) returns()
func (_BrokerDealer *BrokerDealerSession) RemoveAccount(account common.Address) (*types.Transaction, error) {
	return _BrokerDealer.Contract.RemoveAccount(&_BrokerDealer.TransactOpts, account)
}

// RemoveAccount is a paid mutator transaction binding the contract method 0xc4740a95.
//
// Solidity: function removeAccount(account address) returns()
func (_BrokerDealer *BrokerDealerTransactorSession) RemoveAccount(account common.Address) (*types.Transaction, error) {
	return _BrokerDealer.Contract.RemoveAccount(&_BrokerDealer.TransactOpts, account)
}

// SetFrozen is a paid mutator transaction binding the contract method 0xac869cd8.
//
// Solidity: function setFrozen(brokerDealer address, frozen bool) returns()
func (_BrokerDealer *BrokerDealerTransactor) SetFrozen(opts *bind.TransactOpts, brokerDealer common.Address, frozen bool) (*types.Transaction, error) {
	return _BrokerDealer.contract.Transact(opts, "setFrozen", brokerDealer, frozen)
}

// SetFrozen is a paid mutator transaction binding the contract method 0xac869cd8.
//
// Solidity: function setFrozen(brokerDealer address, frozen bool) returns()
func (_BrokerDealer *BrokerDealerSession) SetFrozen(brokerDealer common.Address, frozen bool) (*types.Transaction, error) {
	return _BrokerDealer.Contract.SetFrozen(&_BrokerDealer.TransactOpts, brokerDealer, frozen)
}

// SetFrozen is a paid mutator transaction binding the contract method 0xac869cd8.
//
// Solidity: function setFrozen(brokerDealer address, frozen bool) returns()
func (_BrokerDealer *BrokerDealerTransactorSession) SetFrozen(brokerDealer common.Address, frozen bool) (*types.Transaction, error) {
	return _BrokerDealer.Contract.SetFrozen(&_BrokerDealer.TransactOpts, brokerDealer, frozen)
}

// SetLocked is a paid mutator transaction binding the contract method 0x211e28b6.
//
// Solidity: function setLocked(locked bool) returns()
func (_BrokerDealer *BrokerDealerTransactor) SetLocked(opts *bind.TransactOpts, locked bool) (*types.Transaction, error) {
	return _BrokerDealer.contract.Transact(opts, "setLocked", locked)
}

// SetLocked is a paid mutator transaction binding the contract method 0x211e28b6.
//
// Solidity: function setLocked(locked bool) returns()
func (_BrokerDealer *BrokerDealerSession) SetLocked(locked bool) (*types.Transaction, error) {
	return _BrokerDealer.Contract.SetLocked(&_BrokerDealer.TransactOpts, locked)
}

// SetLocked is a paid mutator transaction binding the contract method 0x211e28b6.
//
// Solidity: function setLocked(locked bool) returns()
func (_BrokerDealer *BrokerDealerTransactorSession) SetLocked(locked bool) (*types.Transaction, error) {
	return _BrokerDealer.Contract.SetLocked(&_BrokerDealer.TransactOpts, locked)
}

// SetStorage is a paid mutator transaction binding the contract method 0x9137c1a7.
//
// Solidity: function setStorage(s address) returns()
func (_BrokerDealer *BrokerDealerTransactor) SetStorage(opts *bind.TransactOpts, s common.Address) (*types.Transaction, error) {
	return _BrokerDealer.contract.Transact(opts, "setStorage", s)
}

// SetStorage is a paid mutator transaction binding the contract method 0x9137c1a7.
//
// Solidity: function setStorage(s address) returns()
func (_BrokerDealer *BrokerDealerSession) SetStorage(s common.Address) (*types.Transaction, error) {
	return _BrokerDealer.Contract.SetStorage(&_BrokerDealer.TransactOpts, s)
}

// SetStorage is a paid mutator transaction binding the contract method 0x9137c1a7.
//
// Solidity: function setStorage(s address) returns()
func (_BrokerDealer *BrokerDealerTransactorSession) SetStorage(s common.Address) (*types.Transaction, error) {
	return _BrokerDealer.Contract.SetStorage(&_BrokerDealer.TransactOpts, s)
}

// TransferOwner is a paid mutator transaction binding the contract method 0x4fb2e45d.
//
// Solidity: function transferOwner(newOwner address) returns()
func (_BrokerDealer *BrokerDealerTransactor) TransferOwner(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _BrokerDealer.contract.Transact(opts, "transferOwner", newOwner)
}

// TransferOwner is a paid mutator transaction binding the contract method 0x4fb2e45d.
//
// Solidity: function transferOwner(newOwner address) returns()
func (_BrokerDealer *BrokerDealerSession) TransferOwner(newOwner common.Address) (*types.Transaction, error) {
	return _BrokerDealer.Contract.TransferOwner(&_BrokerDealer.TransactOpts, newOwner)
}

// TransferOwner is a paid mutator transaction binding the contract method 0x4fb2e45d.
//
// Solidity: function transferOwner(newOwner address) returns()
func (_BrokerDealer *BrokerDealerTransactorSession) TransferOwner(newOwner common.Address) (*types.Transaction, error) {
	return _BrokerDealer.Contract.TransferOwner(&_BrokerDealer.TransactOpts, newOwner)
}

// BrokerDealerBrokerDealerAddedIterator is returned from FilterBrokerDealerAdded and is used to iterate over the raw logs and unpacked data for BrokerDealerAdded events raised by the BrokerDealer contract.
type BrokerDealerBrokerDealerAddedIterator struct {
	Event *BrokerDealerBrokerDealerAdded // Event containing the contract specifics and raw log

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
func (it *BrokerDealerBrokerDealerAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BrokerDealerBrokerDealerAdded)
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
		it.Event = new(BrokerDealerBrokerDealerAdded)
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
func (it *BrokerDealerBrokerDealerAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BrokerDealerBrokerDealerAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BrokerDealerBrokerDealerAdded represents a BrokerDealerAdded event raised by the BrokerDealer contract.
type BrokerDealerBrokerDealerAdded struct {
	BrokerDealer common.Address
	Owner        common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterBrokerDealerAdded is a free log retrieval operation binding the contract event 0xf51b1c4e2953dbc0185cc885a05dcf6fe8de20cf5e8e0c32d7c32d76e45f114b.
//
// Solidity: e BrokerDealerAdded(brokerDealer indexed address, owner indexed address)
func (_BrokerDealer *BrokerDealerFilterer) FilterBrokerDealerAdded(opts *bind.FilterOpts, brokerDealer []common.Address, owner []common.Address) (*BrokerDealerBrokerDealerAddedIterator, error) {

	var brokerDealerRule []interface{}
	for _, brokerDealerItem := range brokerDealer {
		brokerDealerRule = append(brokerDealerRule, brokerDealerItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _BrokerDealer.contract.FilterLogs(opts, "BrokerDealerAdded", brokerDealerRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return &BrokerDealerBrokerDealerAddedIterator{contract: _BrokerDealer.contract, event: "BrokerDealerAdded", logs: logs, sub: sub}, nil
}

// WatchBrokerDealerAdded is a free log subscription operation binding the contract event 0xf51b1c4e2953dbc0185cc885a05dcf6fe8de20cf5e8e0c32d7c32d76e45f114b.
//
// Solidity: e BrokerDealerAdded(brokerDealer indexed address, owner indexed address)
func (_BrokerDealer *BrokerDealerFilterer) WatchBrokerDealerAdded(opts *bind.WatchOpts, sink chan<- *BrokerDealerBrokerDealerAdded, brokerDealer []common.Address, owner []common.Address) (event.Subscription, error) {

	var brokerDealerRule []interface{}
	for _, brokerDealerItem := range brokerDealer {
		brokerDealerRule = append(brokerDealerRule, brokerDealerItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _BrokerDealer.contract.WatchLogs(opts, "BrokerDealerAdded", brokerDealerRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BrokerDealerBrokerDealerAdded)
				if err := _BrokerDealer.contract.UnpackLog(event, "BrokerDealerAdded", log); err != nil {
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

// BrokerDealerBrokerDealerFrozenIterator is returned from FilterBrokerDealerFrozen and is used to iterate over the raw logs and unpacked data for BrokerDealerFrozen events raised by the BrokerDealer contract.
type BrokerDealerBrokerDealerFrozenIterator struct {
	Event *BrokerDealerBrokerDealerFrozen // Event containing the contract specifics and raw log

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
func (it *BrokerDealerBrokerDealerFrozenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BrokerDealerBrokerDealerFrozen)
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
		it.Event = new(BrokerDealerBrokerDealerFrozen)
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
func (it *BrokerDealerBrokerDealerFrozenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BrokerDealerBrokerDealerFrozenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BrokerDealerBrokerDealerFrozen represents a BrokerDealerFrozen event raised by the BrokerDealer contract.
type BrokerDealerBrokerDealerFrozen struct {
	BrokerDealer common.Address
	Frozen       bool
	Owner        common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterBrokerDealerFrozen is a free log retrieval operation binding the contract event 0xbc170ec10277584572ef4c969f53616db496355c027a69325d121126add2a681.
//
// Solidity: e BrokerDealerFrozen(brokerDealer indexed address, frozen indexed bool, owner indexed address)
func (_BrokerDealer *BrokerDealerFilterer) FilterBrokerDealerFrozen(opts *bind.FilterOpts, brokerDealer []common.Address, frozen []bool, owner []common.Address) (*BrokerDealerBrokerDealerFrozenIterator, error) {

	var brokerDealerRule []interface{}
	for _, brokerDealerItem := range brokerDealer {
		brokerDealerRule = append(brokerDealerRule, brokerDealerItem)
	}
	var frozenRule []interface{}
	for _, frozenItem := range frozen {
		frozenRule = append(frozenRule, frozenItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _BrokerDealer.contract.FilterLogs(opts, "BrokerDealerFrozen", brokerDealerRule, frozenRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return &BrokerDealerBrokerDealerFrozenIterator{contract: _BrokerDealer.contract, event: "BrokerDealerFrozen", logs: logs, sub: sub}, nil
}

// WatchBrokerDealerFrozen is a free log subscription operation binding the contract event 0xbc170ec10277584572ef4c969f53616db496355c027a69325d121126add2a681.
//
// Solidity: e BrokerDealerFrozen(brokerDealer indexed address, frozen indexed bool, owner indexed address)
func (_BrokerDealer *BrokerDealerFilterer) WatchBrokerDealerFrozen(opts *bind.WatchOpts, sink chan<- *BrokerDealerBrokerDealerFrozen, brokerDealer []common.Address, frozen []bool, owner []common.Address) (event.Subscription, error) {

	var brokerDealerRule []interface{}
	for _, brokerDealerItem := range brokerDealer {
		brokerDealerRule = append(brokerDealerRule, brokerDealerItem)
	}
	var frozenRule []interface{}
	for _, frozenItem := range frozen {
		frozenRule = append(frozenRule, frozenItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _BrokerDealer.contract.WatchLogs(opts, "BrokerDealerFrozen", brokerDealerRule, frozenRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BrokerDealerBrokerDealerFrozen)
				if err := _BrokerDealer.contract.UnpackLog(event, "BrokerDealerFrozen", log); err != nil {
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

// BrokerDealerBrokerDealerRemovedIterator is returned from FilterBrokerDealerRemoved and is used to iterate over the raw logs and unpacked data for BrokerDealerRemoved events raised by the BrokerDealer contract.
type BrokerDealerBrokerDealerRemovedIterator struct {
	Event *BrokerDealerBrokerDealerRemoved // Event containing the contract specifics and raw log

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
func (it *BrokerDealerBrokerDealerRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BrokerDealerBrokerDealerRemoved)
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
		it.Event = new(BrokerDealerBrokerDealerRemoved)
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
func (it *BrokerDealerBrokerDealerRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BrokerDealerBrokerDealerRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BrokerDealerBrokerDealerRemoved represents a BrokerDealerRemoved event raised by the BrokerDealer contract.
type BrokerDealerBrokerDealerRemoved struct {
	BrokerDealer common.Address
	Owner        common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterBrokerDealerRemoved is a free log retrieval operation binding the contract event 0xdb9237c04fdf8d1a83239f5f1d4e6663f5d36da3b652c728cffa1c21ae6d9f22.
//
// Solidity: e BrokerDealerRemoved(brokerDealer indexed address, owner indexed address)
func (_BrokerDealer *BrokerDealerFilterer) FilterBrokerDealerRemoved(opts *bind.FilterOpts, brokerDealer []common.Address, owner []common.Address) (*BrokerDealerBrokerDealerRemovedIterator, error) {

	var brokerDealerRule []interface{}
	for _, brokerDealerItem := range brokerDealer {
		brokerDealerRule = append(brokerDealerRule, brokerDealerItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _BrokerDealer.contract.FilterLogs(opts, "BrokerDealerRemoved", brokerDealerRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return &BrokerDealerBrokerDealerRemovedIterator{contract: _BrokerDealer.contract, event: "BrokerDealerRemoved", logs: logs, sub: sub}, nil
}

// WatchBrokerDealerRemoved is a free log subscription operation binding the contract event 0xdb9237c04fdf8d1a83239f5f1d4e6663f5d36da3b652c728cffa1c21ae6d9f22.
//
// Solidity: e BrokerDealerRemoved(brokerDealer indexed address, owner indexed address)
func (_BrokerDealer *BrokerDealerFilterer) WatchBrokerDealerRemoved(opts *bind.WatchOpts, sink chan<- *BrokerDealerBrokerDealerRemoved, brokerDealer []common.Address, owner []common.Address) (event.Subscription, error) {

	var brokerDealerRule []interface{}
	for _, brokerDealerItem := range brokerDealer {
		brokerDealerRule = append(brokerDealerRule, brokerDealerItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _BrokerDealer.contract.WatchLogs(opts, "BrokerDealerRemoved", brokerDealerRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BrokerDealerBrokerDealerRemoved)
				if err := _BrokerDealer.contract.UnpackLog(event, "BrokerDealerRemoved", log); err != nil {
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

// BrokerDealerOwnerTransferredIterator is returned from FilterOwnerTransferred and is used to iterate over the raw logs and unpacked data for OwnerTransferred events raised by the BrokerDealer contract.
type BrokerDealerOwnerTransferredIterator struct {
	Event *BrokerDealerOwnerTransferred // Event containing the contract specifics and raw log

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
func (it *BrokerDealerOwnerTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BrokerDealerOwnerTransferred)
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
		it.Event = new(BrokerDealerOwnerTransferred)
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
func (it *BrokerDealerOwnerTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BrokerDealerOwnerTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BrokerDealerOwnerTransferred represents a OwnerTransferred event raised by the BrokerDealer contract.
type BrokerDealerOwnerTransferred struct {
	OldOwner common.Address
	NewOwner common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOwnerTransferred is a free log retrieval operation binding the contract event 0x8934ce4adea8d9ce0d714d2c22b86790e41b7731c84b926fbbdc1d40ff6533c9.
//
// Solidity: e OwnerTransferred(oldOwner indexed address, newOwner indexed address)
func (_BrokerDealer *BrokerDealerFilterer) FilterOwnerTransferred(opts *bind.FilterOpts, oldOwner []common.Address, newOwner []common.Address) (*BrokerDealerOwnerTransferredIterator, error) {

	var oldOwnerRule []interface{}
	for _, oldOwnerItem := range oldOwner {
		oldOwnerRule = append(oldOwnerRule, oldOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _BrokerDealer.contract.FilterLogs(opts, "OwnerTransferred", oldOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &BrokerDealerOwnerTransferredIterator{contract: _BrokerDealer.contract, event: "OwnerTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnerTransferred is a free log subscription operation binding the contract event 0x8934ce4adea8d9ce0d714d2c22b86790e41b7731c84b926fbbdc1d40ff6533c9.
//
// Solidity: e OwnerTransferred(oldOwner indexed address, newOwner indexed address)
func (_BrokerDealer *BrokerDealerFilterer) WatchOwnerTransferred(opts *bind.WatchOpts, sink chan<- *BrokerDealerOwnerTransferred, oldOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var oldOwnerRule []interface{}
	for _, oldOwnerItem := range oldOwner {
		oldOwnerRule = append(oldOwnerRule, oldOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _BrokerDealer.contract.WatchLogs(opts, "OwnerTransferred", oldOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BrokerDealerOwnerTransferred)
				if err := _BrokerDealer.contract.UnpackLog(event, "OwnerTransferred", log); err != nil {
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
