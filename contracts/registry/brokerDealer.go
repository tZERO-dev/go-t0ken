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
const BrokerDealerABI = "[{\"constant\":false,\"inputs\":[{\"internalType\":\"bool\",\"name\":\"locked\",\"type\":\"bool\"}],\"name\":\"setLocked\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"brokerDealer\",\"type\":\"address\"}],\"name\":\"remove\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"brokerDealer\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"}],\"name\":\"add\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"kill\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"brokerDealer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"addAccount\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwner\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ZERO_ADDRESS\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"registry\",\"outputs\":[{\"internalType\":\"contractIRegistry\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isLocked\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"contractIRegistry\",\"name\":\"r\",\"type\":\"address\"}],\"name\":\"setRegistry\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"brokerDealer\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"frozen\",\"type\":\"bool\"}],\"name\":\"setFrozen\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"removeAccount\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIRegistry\",\"name\":\"r\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnerTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"brokerDealer\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"BrokerDealerAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"brokerDealer\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"BrokerDealerRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"brokerDealer\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bool\",\"name\":\"frozen\",\"type\":\"bool\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"BrokerDealerFrozen\",\"type\":\"event\"}]"

// BrokerDealerBin is the compiled bytecode used for deploying new contracts.
const BrokerDealerBin = `6080604052600180546001600160a01b031916905534801561002057600080fd5b5060405161176e38038061176e8339818101604052602081101561004357600080fd5b5051600080546001600160a01b031990811633179091556001805460ff60a01b19169055600280546001600160a01b03909316929091169190911790556116df8061008f6000396000f3fe608060405234801561001057600080fd5b50600436106100df5760003560e01c8063538ba4f91161008c578063a4e2d63411610066578063a4e2d634146101e7578063a91ee0dc14610203578063ac869cd814610229578063c4740a9514610257576100df565b8063538ba4f9146101b35780637b103999146101d75780638da5cb5b146101df576100df565b806341c0e1b5116100bd57806341c0e1b51461015757806349b51b171461015f5780634fb2e45d1461018d576100df565b8063211e28b6146100e457806329092d0e146101055780632d580ef61461012b575b600080fd5b610103600480360360208110156100fa57600080fd5b5035151561027d565b005b6101036004803603602081101561011b57600080fd5b50356001600160a01b0316610380565b6101036004803603604081101561014157600080fd5b506001600160a01b038135169060200135610661565b6101036108a8565b6101036004803603604081101561017557600080fd5b506001600160a01b0381358116916020013516610930565b610103600480360360208110156101a357600080fd5b50356001600160a01b0316610d7c565b6101bb610efb565b604080516001600160a01b039092168252519081900360200190f35b6101bb610f0a565b6101bb610f19565b6101ef610f28565b604080519115158252519081900360200190f35b6101036004803603602081101561021957600080fd5b50356001600160a01b0316610f38565b6101036004803603604081101561023f57600080fd5b506001600160a01b0381351690602001351515610fe1565b6101036004803603602081101561026d57600080fd5b50356001600160a01b03166112e8565b6000546001600160a01b03163314806102a657506001546000546001600160a01b039081169116145b6102f7576040805162461bcd60e51b815260206004820152601960248201527f4f776e6572206163636f756e7420697320726571756972656400000000000000604482015290519081900360640190fd5b60015460ff600160a01b90910416151581151514156103475760405162461bcd60e51b81526004018080602001828103825260288152602001806116836028913960400191505060405180910390fd5b60018054911515600160a01b027fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff909216919091179055565b60025460408051639ae83e3f60e01b815233600482015260016024820152905183926001600160a01b031691639ae83e3f916044808301926020929190829003018186803b1580156103d157600080fd5b505afa1580156103e5573d6000803e3d6000fd5b505050506040513d60208110156103fb57600080fd5b5051610443576040805162461bcd60e51b815260206004820152601260248201527110dd5cdd1bd91a585b881c995c5d5a5c995960721b604482015290519081900360640190fd5b600254604080516318d553a560e11b81526001600160a01b038481166004830152915191909216916331aaa74a916024808301926020929190829003018186803b15801561049057600080fd5b505afa1580156104a4573d6000803e3d6000fd5b505050506040513d60208110156104ba57600080fd5b50516001600160a01b031633146105025760405162461bcd60e51b81526004018080602001828103825260228152602001806116616022913960400191505060405180910390fd5b6002546040805163624522f960e01b815233600482015290516001600160a01b039092169163624522f991602480820192602092909190829003018186803b15801561054d57600080fd5b505afa158015610561573d6000803e3d6000fd5b505050506040513d602081101561057757600080fd5b5051156105c1576040805162461bcd60e51b815260206004820152601360248201527221bab9ba37b234b0b71034b990333937bd32b760691b604482015290519081900360640190fd5b6002546040805163c4740a9560e01b81526001600160a01b0385811660048301529151919092169163c4740a9591602480830192600092919082900301818387803b15801561060f57600080fd5b505af1158015610623573d6000803e3d6000fd5b50506040513392506001600160a01b03851691507fdb9237c04fdf8d1a83239f5f1d4e6663f5d36da3b652c728cffa1c21ae6d9f2290600090a35050565b60025460408051639ae83e3f60e01b81523360048201526001602482015290516001600160a01b0390921691639ae83e3f91604480820192602092909190829003018186803b1580156106b357600080fd5b505afa1580156106c7573d6000803e3d6000fd5b505050506040513d60208110156106dd57600080fd5b5051610730576040805162461bcd60e51b815260206004820152601a60248201527f437573746f6469616e2061646472657373207265717569726564000000000000604482015290519081900360640190fd5b6002546040805163624522f960e01b815233600482015290516001600160a01b039092169163624522f991602480820192602092909190829003018186803b15801561077b57600080fd5b505afa15801561078f573d6000803e3d6000fd5b505050506040513d60208110156107a557600080fd5b5051156107ef576040805162461bcd60e51b815260206004820152601360248201527221bab9ba37b234b0b71034b990333937bd32b760691b604482015290519081900360640190fd5b6002546040805163e8d74d1f60e01b81526001600160a01b0385811660048301526003602483015260006044830181905233606484015260848301869052925193169263e8d74d1f9260a48084019391929182900301818387803b15801561085657600080fd5b505af115801561086a573d6000803e3d6000fd5b50506040513392506001600160a01b03851691507ff51b1c4e2953dbc0185cc885a05dcf6fe8de20cf5e8e0c32d7c32d76e45f114b90600090a35050565b6000546001600160a01b03163314806108d157506001546000546001600160a01b039081169116145b610922576040805162461bcd60e51b815260206004820152601960248201527f4f776e6572206163636f756e7420697320726571756972656400000000000000604482015290519081900360640190fd5b6000546001600160a01b0316ff5b600254604080517f75cd51ed0000000000000000000000000000000000000000000000000000000081526001600160a01b0380851660048301529151849392909216916375cd51ed91602480820192602092909190829003018186803b15801561099957600080fd5b505afa1580156109ad573d6000803e3d6000fd5b505050506040513d60208110156109c357600080fd5b505115610a17576040805162461bcd60e51b815260206004820152601660248201527f4163636f756e7420616c72656164792065786973747300000000000000000000604482015290519081900360640190fd5b60025460408051639ae83e3f60e01b815233600482015260016024820152905185926001600160a01b031691639ae83e3f916044808301926020929190829003018186803b158015610a6857600080fd5b505afa158015610a7c573d6000803e3d6000fd5b505050506040513d6020811015610a9257600080fd5b5051610ada576040805162461bcd60e51b815260206004820152601260248201527110dd5cdd1bd91a585b881c995c5d5a5c995960721b604482015290519081900360640190fd5b600254604080516318d553a560e11b81526001600160a01b038481166004830152915191909216916331aaa74a916024808301926020929190829003018186803b158015610b2757600080fd5b505afa158015610b3b573d6000803e3d6000fd5b505050506040513d6020811015610b5157600080fd5b50516001600160a01b03163314610b995760405162461bcd60e51b81526004018080602001828103825260228152602001806116616022913960400191505060405180910390fd5b6002546040805163624522f960e01b815233600482015290516001600160a01b039092169163624522f991602480820192602092909190829003018186803b158015610be457600080fd5b505afa158015610bf8573d6000803e3d6000fd5b505050506040513d6020811015610c0e57600080fd5b505115610c58576040805162461bcd60e51b815260206004820152601360248201527221bab9ba37b234b0b71034b990333937bd32b760691b604482015290519081900360640190fd5b600254604080517f960d27d30000000000000000000000000000000000000000000000000000000081526001600160a01b0387811660048301529151600093929092169163960d27d391602480820192602092909190829003018186803b158015610cc257600080fd5b505afa158015610cd6573d6000803e3d6000fd5b505050506040513d6020811015610cec57600080fd5b5051600280546040805163e8d74d1f60e01b81526001600160a01b03898116600483015260248201949094526000604482018190528a85166064830152608482018690529151949550929091169263e8d74d1f9260a480820193929182900301818387803b158015610d5d57600080fd5b505af1158015610d71573d6000803e3d6000fd5b505050505050505050565b6000546001600160a01b0316331480610da557506001546000546001600160a01b039081169116145b610df6576040805162461bcd60e51b815260206004820152601960248201527f4f776e6572206163636f756e7420697320726571756972656400000000000000604482015290519081900360640190fd5b6000546001600160a01b0382811691161415610e435760405162461bcd60e51b815260040180806020018281038252602581526020018061163c6025913960400191505060405180910390fd5b6001600160a01b038116610e9e576040805162461bcd60e51b815260206004820181905260248201527f4e6577204f776e65722063616e6e6f74206265207a65726f2061646472657373604482015290519081900360640190fd5b600080546001600160a01b0383811673ffffffffffffffffffffffffffffffffffffffff19831681178455604051919092169283917f8934ce4adea8d9ce0d714d2c22b86790e41b7731c84b926fbbdc1d40ff6533c99190a35050565b6001546001600160a01b031681565b6002546001600160a01b031681565b6000546001600160a01b031681565b600154600160a01b900460ff1681565b6000546001600160a01b0316331480610f6157506001546000546001600160a01b039081169116145b610fb2576040805162461bcd60e51b815260206004820152601960248201527f4f776e6572206163636f756e7420697320726571756972656400000000000000604482015290519081900360640190fd5b6002805473ffffffffffffffffffffffffffffffffffffffff19166001600160a01b0392909216919091179055565b60025460408051639ae83e3f60e01b815233600482015260016024820152905184926001600160a01b031691639ae83e3f916044808301926020929190829003018186803b15801561103257600080fd5b505afa158015611046573d6000803e3d6000fd5b505050506040513d602081101561105c57600080fd5b50516110a4576040805162461bcd60e51b815260206004820152601260248201527110dd5cdd1bd91a585b881c995c5d5a5c995960721b604482015290519081900360640190fd5b600254604080516318d553a560e11b81526001600160a01b038481166004830152915191909216916331aaa74a916024808301926020929190829003018186803b1580156110f157600080fd5b505afa158015611105573d6000803e3d6000fd5b505050506040513d602081101561111b57600080fd5b50516001600160a01b031633146111635760405162461bcd60e51b81526004018080602001828103825260228152602001806116616022913960400191505060405180910390fd5b6002546040805163624522f960e01b815233600482015290516001600160a01b039092169163624522f991602480820192602092909190829003018186803b1580156111ae57600080fd5b505afa1580156111c2573d6000803e3d6000fd5b505050506040513d60208110156111d857600080fd5b505115611222576040805162461bcd60e51b815260206004820152601360248201527221bab9ba37b234b0b71034b990333937bd32b760691b604482015290519081900360640190fd5b600254604080517fcbe5404f0000000000000000000000000000000000000000000000000000000081526001600160a01b03868116600483015285151560248301529151919092169163cbe5404f91604480830192600092919082900301818387803b15801561129157600080fd5b505af11580156112a5573d6000803e3d6000fd5b505060405133925084151591506001600160a01b038616907fbc170ec10277584572ef4c969f53616db496355c027a69325d121126add2a68190600090a4505050565b6002805460408051639ae83e3f60e01b81526001600160a01b03808616600483015260248201949094529051849390921691639ae83e3f91604480820192602092909190829003018186803b15801561134057600080fd5b505afa158015611354573d6000803e3d6000fd5b505050506040513d602081101561136a57600080fd5b50516113bd576040805162461bcd60e51b815260206004820152601760248201527f4e6f74206120637573746f6469616c206163636f756e74000000000000000000604482015290519081900360640190fd5b600254604080516318d553a560e11b81526001600160a01b038481166004830152915160009392909216916331aaa74a91602480820192602092909190829003018186803b15801561140e57600080fd5b505afa158015611422573d6000803e3d6000fd5b505050506040513d602081101561143857600080fd5b5051600254604080516318d553a560e11b81526001600160a01b03808516600483015291519394509116916331aaa74a91602480820192602092909190829003018186803b15801561148957600080fd5b505afa15801561149d573d6000803e3d6000fd5b505050506040513d60208110156114b357600080fd5b50516001600160a01b03163314611511576040805162461bcd60e51b815260206004820152601c60248201527f4163636f756e74277320637573746f6469616e20726571757269656400000000604482015290519081900360640190fd5b6002546040805163624522f960e01b815233600482015290516001600160a01b039092169163624522f991602480820192602092909190829003018186803b15801561155c57600080fd5b505afa158015611570573d6000803e3d6000fd5b505050506040513d602081101561158657600080fd5b5051156115d0576040805162461bcd60e51b815260206004820152601360248201527221bab9ba37b234b0b71034b990333937bd32b760691b604482015290519081900360640190fd5b6002546040805163c4740a9560e01b81526001600160a01b0386811660048301529151919092169163c4740a9591602480830192600092919082900301818387803b15801561161e57600080fd5b505af1158015611632573d6000803e3d6000fd5b5050505050505056fe4e6577204f776e65722063616e6e6f74206265207468652063757272656e74206f776e657242726f6b65722d4465616c6572277320637573746f6469616e207265717569726564436f6e747261637420616c726561647920696e20726571756573746564206c6f636b207374617465a265627a7a723158205055da161d986d5d2c4f0a15e24184db10e43e2b83278cfa4546dc002c04543a64736f6c634300050b0032`

// DeployBrokerDealer deploys a new Ethereum contract, binding an instance of BrokerDealer to it.
func DeployBrokerDealer(auth *bind.TransactOpts, backend bind.ContractBackend, r common.Address) (common.Address, *types.Transaction, *BrokerDealer, error) {
	parsed, err := abi.JSON(strings.NewReader(BrokerDealerABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(BrokerDealerBin), backend, r)
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

// ZEROADDRESS is a free data retrieval call binding the contract method 0x538ba4f9.
//
// Solidity: function ZERO_ADDRESS() constant returns(address)
func (_BrokerDealer *BrokerDealerCaller) ZEROADDRESS(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _BrokerDealer.contract.Call(opts, out, "ZERO_ADDRESS")
	return *ret0, err
}

// ZEROADDRESS is a free data retrieval call binding the contract method 0x538ba4f9.
//
// Solidity: function ZERO_ADDRESS() constant returns(address)
func (_BrokerDealer *BrokerDealerSession) ZEROADDRESS() (common.Address, error) {
	return _BrokerDealer.Contract.ZEROADDRESS(&_BrokerDealer.CallOpts)
}

// ZEROADDRESS is a free data retrieval call binding the contract method 0x538ba4f9.
//
// Solidity: function ZERO_ADDRESS() constant returns(address)
func (_BrokerDealer *BrokerDealerCallerSession) ZEROADDRESS() (common.Address, error) {
	return _BrokerDealer.Contract.ZEROADDRESS(&_BrokerDealer.CallOpts)
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

// Registry is a free data retrieval call binding the contract method 0x7b103999.
//
// Solidity: function registry() constant returns(address)
func (_BrokerDealer *BrokerDealerCaller) Registry(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _BrokerDealer.contract.Call(opts, out, "registry")
	return *ret0, err
}

// Registry is a free data retrieval call binding the contract method 0x7b103999.
//
// Solidity: function registry() constant returns(address)
func (_BrokerDealer *BrokerDealerSession) Registry() (common.Address, error) {
	return _BrokerDealer.Contract.Registry(&_BrokerDealer.CallOpts)
}

// Registry is a free data retrieval call binding the contract method 0x7b103999.
//
// Solidity: function registry() constant returns(address)
func (_BrokerDealer *BrokerDealerCallerSession) Registry() (common.Address, error) {
	return _BrokerDealer.Contract.Registry(&_BrokerDealer.CallOpts)
}

// Add is a paid mutator transaction binding the contract method 0x2d580ef6.
//
// Solidity: function add(brokerDealer address, hash bytes32) returns()
func (_BrokerDealer *BrokerDealerTransactor) Add(opts *bind.TransactOpts, brokerDealer common.Address, hash [32]byte) (*types.Transaction, error) {
	return _BrokerDealer.contract.Transact(opts, "add", brokerDealer, hash)
}

// Add is a paid mutator transaction binding the contract method 0x2d580ef6.
//
// Solidity: function add(brokerDealer address, hash bytes32) returns()
func (_BrokerDealer *BrokerDealerSession) Add(brokerDealer common.Address, hash [32]byte) (*types.Transaction, error) {
	return _BrokerDealer.Contract.Add(&_BrokerDealer.TransactOpts, brokerDealer, hash)
}

// Add is a paid mutator transaction binding the contract method 0x2d580ef6.
//
// Solidity: function add(brokerDealer address, hash bytes32) returns()
func (_BrokerDealer *BrokerDealerTransactorSession) Add(brokerDealer common.Address, hash [32]byte) (*types.Transaction, error) {
	return _BrokerDealer.Contract.Add(&_BrokerDealer.TransactOpts, brokerDealer, hash)
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

// SetRegistry is a paid mutator transaction binding the contract method 0xa91ee0dc.
//
// Solidity: function setRegistry(r address) returns()
func (_BrokerDealer *BrokerDealerTransactor) SetRegistry(opts *bind.TransactOpts, r common.Address) (*types.Transaction, error) {
	return _BrokerDealer.contract.Transact(opts, "setRegistry", r)
}

// SetRegistry is a paid mutator transaction binding the contract method 0xa91ee0dc.
//
// Solidity: function setRegistry(r address) returns()
func (_BrokerDealer *BrokerDealerSession) SetRegistry(r common.Address) (*types.Transaction, error) {
	return _BrokerDealer.Contract.SetRegistry(&_BrokerDealer.TransactOpts, r)
}

// SetRegistry is a paid mutator transaction binding the contract method 0xa91ee0dc.
//
// Solidity: function setRegistry(r address) returns()
func (_BrokerDealer *BrokerDealerTransactorSession) SetRegistry(r common.Address) (*types.Transaction, error) {
	return _BrokerDealer.Contract.SetRegistry(&_BrokerDealer.TransactOpts, r)
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
