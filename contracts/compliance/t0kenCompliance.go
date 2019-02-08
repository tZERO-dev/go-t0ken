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

// T0kenComplianceABI is the input ABI used to generate the binding from.
const T0kenComplianceABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"admin\",\"type\":\"address\"}],\"name\":\"removeAdmin\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"locked\",\"type\":\"bool\"}],\"name\":\"setLocked\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"isAdmin\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"kill\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwner\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"admin\",\"type\":\"address\"},{\"name\":\"from\",\"type\":\"address\"},{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"tokens\",\"type\":\"uint256\"}],\"name\":\"canOverride\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"maxRules\",\"outputs\":[{\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"initiator\",\"type\":\"address\"},{\"name\":\"from\",\"type\":\"address\"},{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"tokens\",\"type\":\"uint256\"}],\"name\":\"canTransfer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"admin\",\"type\":\"address\"}],\"name\":\"addAdmin\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"addressFreezes\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"s\",\"type\":\"address\"}],\"name\":\"setStorage\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"store\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isLocked\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"admins\",\"outputs\":[{\"name\":\"count\",\"type\":\"int256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"},{\"name\":\"freeze\",\"type\":\"bool\"}],\"name\":\"setFrozen\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"kind\",\"type\":\"uint8\"}],\"name\":\"getRules\",\"outputs\":[{\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"kind\",\"type\":\"uint8\"},{\"name\":\"rules\",\"type\":\"address[]\"}],\"name\":\"setRules\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"limit\",\"type\":\"uint8\"}],\"name\":\"setMaxRules\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"issuer\",\"type\":\"address\"},{\"name\":\"from\",\"type\":\"address\"},{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"tokens\",\"type\":\"uint256\"}],\"name\":\"canIssue\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"admin\",\"type\":\"address\"}],\"name\":\"AdminAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"admin\",\"type\":\"address\"}],\"name\":\"AdminRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"oldOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnerTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"addr\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"isFrozen\",\"type\":\"bool\"},{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"AddressFrozen\",\"type\":\"event\"}]"

// T0kenComplianceBin is the compiled bytecode used for deploying new contracts.
const T0kenComplianceBin = `60806040526004805460008054600160a060020a0319163317905560ff1961190061ff00199092169190911716905561158b8061003d6000396000f3fe608060405234801561001057600080fd5b5060043610610189576000357c0100000000000000000000000000000000000000000000000000000000900480638da5cb5b116100ea578063ac869cd81161009e578063d179d77d11610083578063d179d77d14610431578063e03c6714146104ab578063fd8258bd146104cb57610189565b8063ac869cd814610393578063afc24bfb146103c157610189565b8063975057e7116100cf578063975057e714610369578063a4e2d63414610371578063a5de36191461037957610189565b80638da5cb5b1461031f5780639137c1a71461034357610189565b80635acba201116101415780636d62a4fe116101265780636d62a4fe1461029757806370480275146102d357806375dcfdbf146102f957610189565b80635acba2011461023d5780636ae728c61461027957610189565b806324d7806c1161017257806324d7806c146101d557806341c0e1b51461020f5780634fb2e45d1461021757610189565b80631785f53c1461018e578063211e28b6146101b6575b600080fd5b6101b4600480360360208110156101a457600080fd5b5035600160a060020a0316610507565b005b6101b4600480360360208110156101cc57600080fd5b50351515610611565b6101fb600480360360208110156101eb57600080fd5b5035600160a060020a03166106f4565b604080519115158252519081900360200190f35b6101b4610723565b6101b46004803603602081101561022d57600080fd5b5035600160a060020a0316610798565b6101fb6004803603608081101561025357600080fd5b50600160a060020a03813581169160208101358216916040820135169060600135610916565b610281610993565b6040805160ff9092168252519081900360200190f35b6101fb600480360360808110156102ad57600080fd5b50600160a060020a038135811691602081013582169160408201351690606001356109a1565b6101b4600480360360208110156102e957600080fd5b5035600160a060020a03166109fb565b6101fb6004803603602081101561030f57600080fd5b5035600160a060020a0316610b05565b610327610b1a565b60408051600160a060020a039092168252519081900360200190f35b6101b46004803603602081101561035957600080fd5b5035600160a060020a0316610b29565b610327610bd0565b6101fb610be5565b610381610bee565b60408051918252519081900360200190f35b6101b4600480360360408110156103a957600080fd5b50600160a060020a0381351690602001351515610bf4565b6103e1600480360360208110156103d757600080fd5b503560ff16610cc9565b60408051602080825283518183015283519192839290830191858101910280838360005b8381101561041d578181015183820152602001610405565b505050509050019250505060405180910390f35b6101b46004803603604081101561044757600080fd5b60ff823516919081019060408101602082013564010000000081111561046c57600080fd5b82018360208201111561047e57600080fd5b803590602001918460208302840111640100000000831117156104a057600080fd5b509092509050610d39565b6101b4600480360360208110156104c157600080fd5b503560ff16610e43565b6101fb600480360360808110156104e157600080fd5b50600160a060020a03813581169160208101358216916040820135169060600135610ec6565b600054600160a060020a0316331461056e57604080516000805160206114f0833981519152815260206004820152601960248201527f4f776e6572206163636f756e7420697320726571756972656400000000000000604482015290519081900360640190fd5b61057f60018263ffffffff610efd16565b15156105da57604080516000805160206114f0833981519152815260206004820152601660248201527f556e61626c6520746f2072656d6f76652061646d696e00000000000000000000604482015290519081900360640190fd5b604051600160a060020a038216907fa3b62bc36326052d97ea62d63c3d60308ed4c3ea8ac079dd8499f1e9c4f80c0f90600090a250565b600054600160a060020a0316331480610636575061063660013363ffffffff61100216565b151561069157604080516000805160206114f0833981519152815260206004820152601960248201527f41646d696e206163636f756e7420697320726571756972656400000000000000604482015290519081900360640190fd5b60045460ff16151581151514156106e1576040516000805160206114f083398151915281526004018080602001828103825260288152602001806115386028913960400191505060405180910390fd5b6004805460ff1916911515919091179055565b60008054600160a060020a038381169116148061071d575061071d60018363ffffffff61100216565b92915050565b600054600160a060020a0316331461078a57604080516000805160206114f0833981519152815260206004820152601960248201527f4f776e6572206163636f756e7420697320726571756972656400000000000000604482015290519081900360640190fd5b600054600160a060020a0316ff5b600054600160a060020a031633146107ff57604080516000805160206114f0833981519152815260206004820152601960248201527f4f776e6572206163636f756e7420697320726571756972656400000000000000604482015290519081900360640190fd5b600054600160a060020a0382811691161415610854576040516000805160206114f083398151915281526004018080602001828103825260258152602001806114a06025913960400191505060405180910390fd5b600160a060020a03811615156108b957604080516000805160206114f0833981519152815260206004820181905260248201527f4e6577204f776e65722063616e6e6f74206265207a65726f2061646472657373604482015290519081900360640190fd5b60008054600160a060020a0383811673ffffffffffffffffffffffffffffffffffffffff19831681178455604051919092169283917f8934ce4adea8d9ce0d714d2c22b86790e41b7731c84b926fbbdc1d40ff6533c99190a35050565b6000610921856106f4565b151561097c57604080516000805160206114f0833981519152815260206004820152601960248201527f41646d696e206163636f756e7420697320726571756972656400000000000000604482015290519081900360640190fd5b610987836000611039565b50600195945050505050565b600454610100900460ff1681565b6000806109af856000611039565b905060006109be856000611039565b9050600160a060020a03878116908716146109e0576109de876000611039565b505b6109ee87878785858961124a565b5060019695505050505050565b600054600160a060020a03163314610a6257604080516000805160206114f0833981519152815260206004820152601960248201527f4f776e6572206163636f756e7420697320726571756972656400000000000000604482015290519081900360640190fd5b610a7360018263ffffffff61134716565b1515610ace57604080516000805160206114f0833981519152815260206004820152601360248201527f556e61626c6520746f206164642061646d696e00000000000000000000000000604482015290519081900360640190fd5b604051600160a060020a038216907f44d6d25963f097ad14f29f06854a01f575648a1ef82f30e562ccd3889717e33990600090a250565b60056020526000908152604090205460ff1681565b600054600160a060020a031681565b600054600160a060020a03163314610b9057604080516000805160206114f0833981519152815260206004820152601960248201527f4f776e6572206163636f756e7420697320726571756972656400000000000000604482015290519081900360640190fd5b60048054600160a060020a0390921662010000027fffffffffffffffffffff0000000000000000000000000000000000000000ffff909216919091179055565b600454620100009004600160a060020a031681565b60045460ff1681565b60015481565b600054600160a060020a0316331480610c195750610c1960013363ffffffff61100216565b1515610c7457604080516000805160206114f0833981519152815260206004820152601960248201527f41646d696e206163636f756e7420697320726571756972656400000000000000604482015290519081900360640190fd5b600160a060020a038216600081815260056020526040808220805460ff19168515159081179091559051339391927f7fa523c84ab8d7fc5b72f08b9e46dbbf10c39e119a075b3e317002d14bc9f43691a45050565b60ff8116600090815260066020908152604091829020805483518184028101840190945280845260609392830182828015610d2d57602002820191906000526020600020905b8154600160a060020a03168152600190910190602001808311610d0f575b50505050509050919050565b600054600160a060020a0316331480610d5e5750610d5e60013363ffffffff61100216565b1515610db957604080516000805160206114f0833981519152815260206004820152601960248201527f41646d696e206163636f756e7420697320726571756972656400000000000000604482015290519081900360640190fd5b600454610100900460ff16811115610e2057604080516000805160206114f0833981519152815260206004820152600e60248201527f546f6f206d616e792072756c6573000000000000000000000000000000000000604482015290519081900360640190fd5b60ff83166000908152600660205260409020610e3d9083836113fb565b50505050565b600054600160a060020a03163314610eaa57604080516000805160206114f0833981519152815260206004820152601960248201527f4f776e6572206163636f756e7420697320726571756972656400000000000000604482015290519081900360640190fd5b6004805460ff9092166101000261ff0019909216919091179055565b6000610ed3836000611039565b50600160a060020a0385811690851614610ef257610987846000611039565b506001949350505050565b600160a060020a03811660009081526001808401602052604082205490811280610f275750835481135b15610f3657600091505061071d565b8354811215610faa5783546000908152600285016020818152604080842054600160a060020a031680855260018901835281852086905585855292909152808320805473ffffffffffffffffffffffffffffffffffffffff1990811690931790558654835290912080549091169055610fd6565b60008181526002850160205260409020805473ffffffffffffffffffffffffffffffffffffffff191690555b5050600160a060020a031660009081526001828101602052604082209190915581546000190190915590565b600160a060020a0381166000908152600183016020526040812054600019018181128015906110315750835481125b949350505050565b600160a060020a03821660009081526005602052604081205460ff1615611099576040516000805160206114f083398151915281526004018080602001828103825260288152602001806115106028913960400191505060405180910390fd5b6000806000600460029054906101000a9004600160a060020a0316600160a060020a031663137e37d9876040518263ffffffff167c01000000000000000000000000000000000000000000000000000000000281526004018082600160a060020a0316600160a060020a0316815260200191505060606040518083038186803b15801561112557600080fd5b505afa158015611139573d6000803e3d6000fd5b505050506040513d606081101561114f57600080fd5b5080516020820151604090920151909450909250905081156111aa576040516000805160206114f0833981519152815260040180806020018281038252602b8152602001806114c5602b913960400191505060405180910390fd5b60ff8516151561121957600560ff8416111561121557604080516000805160206114f0833981519152815260206004820152601460248201527f496e76616c6964206163636f756e74206b696e64000000000000000000000000604482015290519081900360640190fd5b8294505b600160ff841611801561122f5750600560ff8416105b156112405761123e8186611039565b505b5092949350505050565b60ff83166000908152600660205260408120905b815460ff8216101561133d57818160ff1681548110151561127b57fe5b600091825260208220015460048054604080517fb762c76d000000000000000000000000000000000000000000000000000000008152600160a060020a038e8116948201949094528c841660248201528b8416604482015260ff8a1660648201526084810189905262010000909204831660a483015251919092169263b762c76d9260c4808201939182900301818387803b15801561131957600080fd5b505af115801561132d573d6000803e3d6000fd5b50506001909201915061125e9050565b5050505050505050565b6000600160a060020a03821615156113615750600061071d565b600160a060020a0382166000908152600184016020526040812054600019019081128015906113905750835481125b1561139f57600091505061071d565b505081546001908101808455600160a060020a03831660008181528386016020908152604080832085905593825260028701905291909120805473ffffffffffffffffffffffffffffffffffffffff1916909117905592915050565b82805482825590600052602060002090810192821561145b579160200282015b8281111561145b57815473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0384351617825560209092019160019091019061141b565b5061146792915061146b565b5090565b61149c91905b8082111561146757805473ffffffffffffffffffffffffffffffffffffffff19168155600101611471565b9056fe4e6577204f776e65722063616e6e6f74206265207468652063757272656e74206f776e657241646472657373206f7220706172656e742069732066726f7a656e2061742074686520726567697374727908c379a00000000000000000000000000000000000000000000000000000000041646472657373206f7220706172656e742069732066726f7a656e2061742074686520746f6b656e436f6e747261637420616c726561647920696e20726571756573746564206c6f636b207374617465a165627a7a72305820f688bbd272e4e5567291e25d5ffd06fd52331943f124b8be1697ec20921f52620029`

// DeployT0kenCompliance deploys a new Ethereum contract, binding an instance of T0kenCompliance to it.
func DeployT0kenCompliance(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *T0kenCompliance, error) {
	parsed, err := abi.JSON(strings.NewReader(T0kenComplianceABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(T0kenComplianceBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &T0kenCompliance{T0kenComplianceCaller: T0kenComplianceCaller{contract: contract}, T0kenComplianceTransactor: T0kenComplianceTransactor{contract: contract}, T0kenComplianceFilterer: T0kenComplianceFilterer{contract: contract}}, nil
}

// T0kenCompliance is an auto generated Go binding around an Ethereum contract.
type T0kenCompliance struct {
	T0kenComplianceCaller     // Read-only binding to the contract
	T0kenComplianceTransactor // Write-only binding to the contract
	T0kenComplianceFilterer   // Log filterer for contract events
}

// T0kenComplianceCaller is an auto generated read-only Go binding around an Ethereum contract.
type T0kenComplianceCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// T0kenComplianceTransactor is an auto generated write-only Go binding around an Ethereum contract.
type T0kenComplianceTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// T0kenComplianceFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type T0kenComplianceFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// T0kenComplianceSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type T0kenComplianceSession struct {
	Contract     *T0kenCompliance  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// T0kenComplianceCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type T0kenComplianceCallerSession struct {
	Contract *T0kenComplianceCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// T0kenComplianceTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type T0kenComplianceTransactorSession struct {
	Contract     *T0kenComplianceTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// T0kenComplianceRaw is an auto generated low-level Go binding around an Ethereum contract.
type T0kenComplianceRaw struct {
	Contract *T0kenCompliance // Generic contract binding to access the raw methods on
}

// T0kenComplianceCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type T0kenComplianceCallerRaw struct {
	Contract *T0kenComplianceCaller // Generic read-only contract binding to access the raw methods on
}

// T0kenComplianceTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type T0kenComplianceTransactorRaw struct {
	Contract *T0kenComplianceTransactor // Generic write-only contract binding to access the raw methods on
}

// NewT0kenCompliance creates a new instance of T0kenCompliance, bound to a specific deployed contract.
func NewT0kenCompliance(address common.Address, backend bind.ContractBackend) (*T0kenCompliance, error) {
	contract, err := bindT0kenCompliance(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &T0kenCompliance{T0kenComplianceCaller: T0kenComplianceCaller{contract: contract}, T0kenComplianceTransactor: T0kenComplianceTransactor{contract: contract}, T0kenComplianceFilterer: T0kenComplianceFilterer{contract: contract}}, nil
}

// NewT0kenComplianceCaller creates a new read-only instance of T0kenCompliance, bound to a specific deployed contract.
func NewT0kenComplianceCaller(address common.Address, caller bind.ContractCaller) (*T0kenComplianceCaller, error) {
	contract, err := bindT0kenCompliance(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &T0kenComplianceCaller{contract: contract}, nil
}

// NewT0kenComplianceTransactor creates a new write-only instance of T0kenCompliance, bound to a specific deployed contract.
func NewT0kenComplianceTransactor(address common.Address, transactor bind.ContractTransactor) (*T0kenComplianceTransactor, error) {
	contract, err := bindT0kenCompliance(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &T0kenComplianceTransactor{contract: contract}, nil
}

// NewT0kenComplianceFilterer creates a new log filterer instance of T0kenCompliance, bound to a specific deployed contract.
func NewT0kenComplianceFilterer(address common.Address, filterer bind.ContractFilterer) (*T0kenComplianceFilterer, error) {
	contract, err := bindT0kenCompliance(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &T0kenComplianceFilterer{contract: contract}, nil
}

// bindT0kenCompliance binds a generic wrapper to an already deployed contract.
func bindT0kenCompliance(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(T0kenComplianceABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_T0kenCompliance *T0kenComplianceRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _T0kenCompliance.Contract.T0kenComplianceCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_T0kenCompliance *T0kenComplianceRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _T0kenCompliance.Contract.T0kenComplianceTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_T0kenCompliance *T0kenComplianceRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _T0kenCompliance.Contract.T0kenComplianceTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_T0kenCompliance *T0kenComplianceCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _T0kenCompliance.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_T0kenCompliance *T0kenComplianceTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _T0kenCompliance.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_T0kenCompliance *T0kenComplianceTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _T0kenCompliance.Contract.contract.Transact(opts, method, params...)
}

// AddressFreezes is a free data retrieval call binding the contract method 0x75dcfdbf.
//
// Solidity: function addressFreezes( address) constant returns(bool)
func (_T0kenCompliance *T0kenComplianceCaller) AddressFreezes(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _T0kenCompliance.contract.Call(opts, out, "addressFreezes", arg0)
	return *ret0, err
}

// AddressFreezes is a free data retrieval call binding the contract method 0x75dcfdbf.
//
// Solidity: function addressFreezes( address) constant returns(bool)
func (_T0kenCompliance *T0kenComplianceSession) AddressFreezes(arg0 common.Address) (bool, error) {
	return _T0kenCompliance.Contract.AddressFreezes(&_T0kenCompliance.CallOpts, arg0)
}

// AddressFreezes is a free data retrieval call binding the contract method 0x75dcfdbf.
//
// Solidity: function addressFreezes( address) constant returns(bool)
func (_T0kenCompliance *T0kenComplianceCallerSession) AddressFreezes(arg0 common.Address) (bool, error) {
	return _T0kenCompliance.Contract.AddressFreezes(&_T0kenCompliance.CallOpts, arg0)
}

// Admins is a free data retrieval call binding the contract method 0xa5de3619.
//
// Solidity: function admins() constant returns(count int256)
func (_T0kenCompliance *T0kenComplianceCaller) Admins(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _T0kenCompliance.contract.Call(opts, out, "admins")
	return *ret0, err
}

// Admins is a free data retrieval call binding the contract method 0xa5de3619.
//
// Solidity: function admins() constant returns(count int256)
func (_T0kenCompliance *T0kenComplianceSession) Admins() (*big.Int, error) {
	return _T0kenCompliance.Contract.Admins(&_T0kenCompliance.CallOpts)
}

// Admins is a free data retrieval call binding the contract method 0xa5de3619.
//
// Solidity: function admins() constant returns(count int256)
func (_T0kenCompliance *T0kenComplianceCallerSession) Admins() (*big.Int, error) {
	return _T0kenCompliance.Contract.Admins(&_T0kenCompliance.CallOpts)
}

// GetRules is a free data retrieval call binding the contract method 0xafc24bfb.
//
// Solidity: function getRules(kind uint8) constant returns(address[])
func (_T0kenCompliance *T0kenComplianceCaller) GetRules(opts *bind.CallOpts, kind uint8) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _T0kenCompliance.contract.Call(opts, out, "getRules", kind)
	return *ret0, err
}

// GetRules is a free data retrieval call binding the contract method 0xafc24bfb.
//
// Solidity: function getRules(kind uint8) constant returns(address[])
func (_T0kenCompliance *T0kenComplianceSession) GetRules(kind uint8) ([]common.Address, error) {
	return _T0kenCompliance.Contract.GetRules(&_T0kenCompliance.CallOpts, kind)
}

// GetRules is a free data retrieval call binding the contract method 0xafc24bfb.
//
// Solidity: function getRules(kind uint8) constant returns(address[])
func (_T0kenCompliance *T0kenComplianceCallerSession) GetRules(kind uint8) ([]common.Address, error) {
	return _T0kenCompliance.Contract.GetRules(&_T0kenCompliance.CallOpts, kind)
}

// IsAdmin is a free data retrieval call binding the contract method 0x24d7806c.
//
// Solidity: function isAdmin(addr address) constant returns(bool)
func (_T0kenCompliance *T0kenComplianceCaller) IsAdmin(opts *bind.CallOpts, addr common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _T0kenCompliance.contract.Call(opts, out, "isAdmin", addr)
	return *ret0, err
}

// IsAdmin is a free data retrieval call binding the contract method 0x24d7806c.
//
// Solidity: function isAdmin(addr address) constant returns(bool)
func (_T0kenCompliance *T0kenComplianceSession) IsAdmin(addr common.Address) (bool, error) {
	return _T0kenCompliance.Contract.IsAdmin(&_T0kenCompliance.CallOpts, addr)
}

// IsAdmin is a free data retrieval call binding the contract method 0x24d7806c.
//
// Solidity: function isAdmin(addr address) constant returns(bool)
func (_T0kenCompliance *T0kenComplianceCallerSession) IsAdmin(addr common.Address) (bool, error) {
	return _T0kenCompliance.Contract.IsAdmin(&_T0kenCompliance.CallOpts, addr)
}

// IsLocked is a free data retrieval call binding the contract method 0xa4e2d634.
//
// Solidity: function isLocked() constant returns(bool)
func (_T0kenCompliance *T0kenComplianceCaller) IsLocked(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _T0kenCompliance.contract.Call(opts, out, "isLocked")
	return *ret0, err
}

// IsLocked is a free data retrieval call binding the contract method 0xa4e2d634.
//
// Solidity: function isLocked() constant returns(bool)
func (_T0kenCompliance *T0kenComplianceSession) IsLocked() (bool, error) {
	return _T0kenCompliance.Contract.IsLocked(&_T0kenCompliance.CallOpts)
}

// IsLocked is a free data retrieval call binding the contract method 0xa4e2d634.
//
// Solidity: function isLocked() constant returns(bool)
func (_T0kenCompliance *T0kenComplianceCallerSession) IsLocked() (bool, error) {
	return _T0kenCompliance.Contract.IsLocked(&_T0kenCompliance.CallOpts)
}

// MaxRules is a free data retrieval call binding the contract method 0x6ae728c6.
//
// Solidity: function maxRules() constant returns(uint8)
func (_T0kenCompliance *T0kenComplianceCaller) MaxRules(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _T0kenCompliance.contract.Call(opts, out, "maxRules")
	return *ret0, err
}

// MaxRules is a free data retrieval call binding the contract method 0x6ae728c6.
//
// Solidity: function maxRules() constant returns(uint8)
func (_T0kenCompliance *T0kenComplianceSession) MaxRules() (uint8, error) {
	return _T0kenCompliance.Contract.MaxRules(&_T0kenCompliance.CallOpts)
}

// MaxRules is a free data retrieval call binding the contract method 0x6ae728c6.
//
// Solidity: function maxRules() constant returns(uint8)
func (_T0kenCompliance *T0kenComplianceCallerSession) MaxRules() (uint8, error) {
	return _T0kenCompliance.Contract.MaxRules(&_T0kenCompliance.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_T0kenCompliance *T0kenComplianceCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _T0kenCompliance.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_T0kenCompliance *T0kenComplianceSession) Owner() (common.Address, error) {
	return _T0kenCompliance.Contract.Owner(&_T0kenCompliance.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_T0kenCompliance *T0kenComplianceCallerSession) Owner() (common.Address, error) {
	return _T0kenCompliance.Contract.Owner(&_T0kenCompliance.CallOpts)
}

// Store is a free data retrieval call binding the contract method 0x975057e7.
//
// Solidity: function store() constant returns(address)
func (_T0kenCompliance *T0kenComplianceCaller) Store(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _T0kenCompliance.contract.Call(opts, out, "store")
	return *ret0, err
}

// Store is a free data retrieval call binding the contract method 0x975057e7.
//
// Solidity: function store() constant returns(address)
func (_T0kenCompliance *T0kenComplianceSession) Store() (common.Address, error) {
	return _T0kenCompliance.Contract.Store(&_T0kenCompliance.CallOpts)
}

// Store is a free data retrieval call binding the contract method 0x975057e7.
//
// Solidity: function store() constant returns(address)
func (_T0kenCompliance *T0kenComplianceCallerSession) Store() (common.Address, error) {
	return _T0kenCompliance.Contract.Store(&_T0kenCompliance.CallOpts)
}

// AddAdmin is a paid mutator transaction binding the contract method 0x70480275.
//
// Solidity: function addAdmin(admin address) returns()
func (_T0kenCompliance *T0kenComplianceTransactor) AddAdmin(opts *bind.TransactOpts, admin common.Address) (*types.Transaction, error) {
	return _T0kenCompliance.contract.Transact(opts, "addAdmin", admin)
}

// AddAdmin is a paid mutator transaction binding the contract method 0x70480275.
//
// Solidity: function addAdmin(admin address) returns()
func (_T0kenCompliance *T0kenComplianceSession) AddAdmin(admin common.Address) (*types.Transaction, error) {
	return _T0kenCompliance.Contract.AddAdmin(&_T0kenCompliance.TransactOpts, admin)
}

// AddAdmin is a paid mutator transaction binding the contract method 0x70480275.
//
// Solidity: function addAdmin(admin address) returns()
func (_T0kenCompliance *T0kenComplianceTransactorSession) AddAdmin(admin common.Address) (*types.Transaction, error) {
	return _T0kenCompliance.Contract.AddAdmin(&_T0kenCompliance.TransactOpts, admin)
}

// CanIssue is a paid mutator transaction binding the contract method 0xfd8258bd.
//
// Solidity: function canIssue(issuer address, from address, to address, tokens uint256) returns(bool)
func (_T0kenCompliance *T0kenComplianceTransactor) CanIssue(opts *bind.TransactOpts, issuer common.Address, from common.Address, to common.Address, tokens *big.Int) (*types.Transaction, error) {
	return _T0kenCompliance.contract.Transact(opts, "canIssue", issuer, from, to, tokens)
}

// CanIssue is a paid mutator transaction binding the contract method 0xfd8258bd.
//
// Solidity: function canIssue(issuer address, from address, to address, tokens uint256) returns(bool)
func (_T0kenCompliance *T0kenComplianceSession) CanIssue(issuer common.Address, from common.Address, to common.Address, tokens *big.Int) (*types.Transaction, error) {
	return _T0kenCompliance.Contract.CanIssue(&_T0kenCompliance.TransactOpts, issuer, from, to, tokens)
}

// CanIssue is a paid mutator transaction binding the contract method 0xfd8258bd.
//
// Solidity: function canIssue(issuer address, from address, to address, tokens uint256) returns(bool)
func (_T0kenCompliance *T0kenComplianceTransactorSession) CanIssue(issuer common.Address, from common.Address, to common.Address, tokens *big.Int) (*types.Transaction, error) {
	return _T0kenCompliance.Contract.CanIssue(&_T0kenCompliance.TransactOpts, issuer, from, to, tokens)
}

// CanOverride is a paid mutator transaction binding the contract method 0x5acba201.
//
// Solidity: function canOverride(admin address, from address, to address, tokens uint256) returns(bool)
func (_T0kenCompliance *T0kenComplianceTransactor) CanOverride(opts *bind.TransactOpts, admin common.Address, from common.Address, to common.Address, tokens *big.Int) (*types.Transaction, error) {
	return _T0kenCompliance.contract.Transact(opts, "canOverride", admin, from, to, tokens)
}

// CanOverride is a paid mutator transaction binding the contract method 0x5acba201.
//
// Solidity: function canOverride(admin address, from address, to address, tokens uint256) returns(bool)
func (_T0kenCompliance *T0kenComplianceSession) CanOverride(admin common.Address, from common.Address, to common.Address, tokens *big.Int) (*types.Transaction, error) {
	return _T0kenCompliance.Contract.CanOverride(&_T0kenCompliance.TransactOpts, admin, from, to, tokens)
}

// CanOverride is a paid mutator transaction binding the contract method 0x5acba201.
//
// Solidity: function canOverride(admin address, from address, to address, tokens uint256) returns(bool)
func (_T0kenCompliance *T0kenComplianceTransactorSession) CanOverride(admin common.Address, from common.Address, to common.Address, tokens *big.Int) (*types.Transaction, error) {
	return _T0kenCompliance.Contract.CanOverride(&_T0kenCompliance.TransactOpts, admin, from, to, tokens)
}

// CanTransfer is a paid mutator transaction binding the contract method 0x6d62a4fe.
//
// Solidity: function canTransfer(initiator address, from address, to address, tokens uint256) returns(bool)
func (_T0kenCompliance *T0kenComplianceTransactor) CanTransfer(opts *bind.TransactOpts, initiator common.Address, from common.Address, to common.Address, tokens *big.Int) (*types.Transaction, error) {
	return _T0kenCompliance.contract.Transact(opts, "canTransfer", initiator, from, to, tokens)
}

// CanTransfer is a paid mutator transaction binding the contract method 0x6d62a4fe.
//
// Solidity: function canTransfer(initiator address, from address, to address, tokens uint256) returns(bool)
func (_T0kenCompliance *T0kenComplianceSession) CanTransfer(initiator common.Address, from common.Address, to common.Address, tokens *big.Int) (*types.Transaction, error) {
	return _T0kenCompliance.Contract.CanTransfer(&_T0kenCompliance.TransactOpts, initiator, from, to, tokens)
}

// CanTransfer is a paid mutator transaction binding the contract method 0x6d62a4fe.
//
// Solidity: function canTransfer(initiator address, from address, to address, tokens uint256) returns(bool)
func (_T0kenCompliance *T0kenComplianceTransactorSession) CanTransfer(initiator common.Address, from common.Address, to common.Address, tokens *big.Int) (*types.Transaction, error) {
	return _T0kenCompliance.Contract.CanTransfer(&_T0kenCompliance.TransactOpts, initiator, from, to, tokens)
}

// Kill is a paid mutator transaction binding the contract method 0x41c0e1b5.
//
// Solidity: function kill() returns()
func (_T0kenCompliance *T0kenComplianceTransactor) Kill(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _T0kenCompliance.contract.Transact(opts, "kill")
}

// Kill is a paid mutator transaction binding the contract method 0x41c0e1b5.
//
// Solidity: function kill() returns()
func (_T0kenCompliance *T0kenComplianceSession) Kill() (*types.Transaction, error) {
	return _T0kenCompliance.Contract.Kill(&_T0kenCompliance.TransactOpts)
}

// Kill is a paid mutator transaction binding the contract method 0x41c0e1b5.
//
// Solidity: function kill() returns()
func (_T0kenCompliance *T0kenComplianceTransactorSession) Kill() (*types.Transaction, error) {
	return _T0kenCompliance.Contract.Kill(&_T0kenCompliance.TransactOpts)
}

// RemoveAdmin is a paid mutator transaction binding the contract method 0x1785f53c.
//
// Solidity: function removeAdmin(admin address) returns()
func (_T0kenCompliance *T0kenComplianceTransactor) RemoveAdmin(opts *bind.TransactOpts, admin common.Address) (*types.Transaction, error) {
	return _T0kenCompliance.contract.Transact(opts, "removeAdmin", admin)
}

// RemoveAdmin is a paid mutator transaction binding the contract method 0x1785f53c.
//
// Solidity: function removeAdmin(admin address) returns()
func (_T0kenCompliance *T0kenComplianceSession) RemoveAdmin(admin common.Address) (*types.Transaction, error) {
	return _T0kenCompliance.Contract.RemoveAdmin(&_T0kenCompliance.TransactOpts, admin)
}

// RemoveAdmin is a paid mutator transaction binding the contract method 0x1785f53c.
//
// Solidity: function removeAdmin(admin address) returns()
func (_T0kenCompliance *T0kenComplianceTransactorSession) RemoveAdmin(admin common.Address) (*types.Transaction, error) {
	return _T0kenCompliance.Contract.RemoveAdmin(&_T0kenCompliance.TransactOpts, admin)
}

// SetFrozen is a paid mutator transaction binding the contract method 0xac869cd8.
//
// Solidity: function setFrozen(addr address, freeze bool) returns()
func (_T0kenCompliance *T0kenComplianceTransactor) SetFrozen(opts *bind.TransactOpts, addr common.Address, freeze bool) (*types.Transaction, error) {
	return _T0kenCompliance.contract.Transact(opts, "setFrozen", addr, freeze)
}

// SetFrozen is a paid mutator transaction binding the contract method 0xac869cd8.
//
// Solidity: function setFrozen(addr address, freeze bool) returns()
func (_T0kenCompliance *T0kenComplianceSession) SetFrozen(addr common.Address, freeze bool) (*types.Transaction, error) {
	return _T0kenCompliance.Contract.SetFrozen(&_T0kenCompliance.TransactOpts, addr, freeze)
}

// SetFrozen is a paid mutator transaction binding the contract method 0xac869cd8.
//
// Solidity: function setFrozen(addr address, freeze bool) returns()
func (_T0kenCompliance *T0kenComplianceTransactorSession) SetFrozen(addr common.Address, freeze bool) (*types.Transaction, error) {
	return _T0kenCompliance.Contract.SetFrozen(&_T0kenCompliance.TransactOpts, addr, freeze)
}

// SetLocked is a paid mutator transaction binding the contract method 0x211e28b6.
//
// Solidity: function setLocked(locked bool) returns()
func (_T0kenCompliance *T0kenComplianceTransactor) SetLocked(opts *bind.TransactOpts, locked bool) (*types.Transaction, error) {
	return _T0kenCompliance.contract.Transact(opts, "setLocked", locked)
}

// SetLocked is a paid mutator transaction binding the contract method 0x211e28b6.
//
// Solidity: function setLocked(locked bool) returns()
func (_T0kenCompliance *T0kenComplianceSession) SetLocked(locked bool) (*types.Transaction, error) {
	return _T0kenCompliance.Contract.SetLocked(&_T0kenCompliance.TransactOpts, locked)
}

// SetLocked is a paid mutator transaction binding the contract method 0x211e28b6.
//
// Solidity: function setLocked(locked bool) returns()
func (_T0kenCompliance *T0kenComplianceTransactorSession) SetLocked(locked bool) (*types.Transaction, error) {
	return _T0kenCompliance.Contract.SetLocked(&_T0kenCompliance.TransactOpts, locked)
}

// SetMaxRules is a paid mutator transaction binding the contract method 0xe03c6714.
//
// Solidity: function setMaxRules(limit uint8) returns()
func (_T0kenCompliance *T0kenComplianceTransactor) SetMaxRules(opts *bind.TransactOpts, limit uint8) (*types.Transaction, error) {
	return _T0kenCompliance.contract.Transact(opts, "setMaxRules", limit)
}

// SetMaxRules is a paid mutator transaction binding the contract method 0xe03c6714.
//
// Solidity: function setMaxRules(limit uint8) returns()
func (_T0kenCompliance *T0kenComplianceSession) SetMaxRules(limit uint8) (*types.Transaction, error) {
	return _T0kenCompliance.Contract.SetMaxRules(&_T0kenCompliance.TransactOpts, limit)
}

// SetMaxRules is a paid mutator transaction binding the contract method 0xe03c6714.
//
// Solidity: function setMaxRules(limit uint8) returns()
func (_T0kenCompliance *T0kenComplianceTransactorSession) SetMaxRules(limit uint8) (*types.Transaction, error) {
	return _T0kenCompliance.Contract.SetMaxRules(&_T0kenCompliance.TransactOpts, limit)
}

// SetRules is a paid mutator transaction binding the contract method 0xd179d77d.
//
// Solidity: function setRules(kind uint8, rules address[]) returns()
func (_T0kenCompliance *T0kenComplianceTransactor) SetRules(opts *bind.TransactOpts, kind uint8, rules []common.Address) (*types.Transaction, error) {
	return _T0kenCompliance.contract.Transact(opts, "setRules", kind, rules)
}

// SetRules is a paid mutator transaction binding the contract method 0xd179d77d.
//
// Solidity: function setRules(kind uint8, rules address[]) returns()
func (_T0kenCompliance *T0kenComplianceSession) SetRules(kind uint8, rules []common.Address) (*types.Transaction, error) {
	return _T0kenCompliance.Contract.SetRules(&_T0kenCompliance.TransactOpts, kind, rules)
}

// SetRules is a paid mutator transaction binding the contract method 0xd179d77d.
//
// Solidity: function setRules(kind uint8, rules address[]) returns()
func (_T0kenCompliance *T0kenComplianceTransactorSession) SetRules(kind uint8, rules []common.Address) (*types.Transaction, error) {
	return _T0kenCompliance.Contract.SetRules(&_T0kenCompliance.TransactOpts, kind, rules)
}

// SetStorage is a paid mutator transaction binding the contract method 0x9137c1a7.
//
// Solidity: function setStorage(s address) returns()
func (_T0kenCompliance *T0kenComplianceTransactor) SetStorage(opts *bind.TransactOpts, s common.Address) (*types.Transaction, error) {
	return _T0kenCompliance.contract.Transact(opts, "setStorage", s)
}

// SetStorage is a paid mutator transaction binding the contract method 0x9137c1a7.
//
// Solidity: function setStorage(s address) returns()
func (_T0kenCompliance *T0kenComplianceSession) SetStorage(s common.Address) (*types.Transaction, error) {
	return _T0kenCompliance.Contract.SetStorage(&_T0kenCompliance.TransactOpts, s)
}

// SetStorage is a paid mutator transaction binding the contract method 0x9137c1a7.
//
// Solidity: function setStorage(s address) returns()
func (_T0kenCompliance *T0kenComplianceTransactorSession) SetStorage(s common.Address) (*types.Transaction, error) {
	return _T0kenCompliance.Contract.SetStorage(&_T0kenCompliance.TransactOpts, s)
}

// TransferOwner is a paid mutator transaction binding the contract method 0x4fb2e45d.
//
// Solidity: function transferOwner(newOwner address) returns()
func (_T0kenCompliance *T0kenComplianceTransactor) TransferOwner(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _T0kenCompliance.contract.Transact(opts, "transferOwner", newOwner)
}

// TransferOwner is a paid mutator transaction binding the contract method 0x4fb2e45d.
//
// Solidity: function transferOwner(newOwner address) returns()
func (_T0kenCompliance *T0kenComplianceSession) TransferOwner(newOwner common.Address) (*types.Transaction, error) {
	return _T0kenCompliance.Contract.TransferOwner(&_T0kenCompliance.TransactOpts, newOwner)
}

// TransferOwner is a paid mutator transaction binding the contract method 0x4fb2e45d.
//
// Solidity: function transferOwner(newOwner address) returns()
func (_T0kenCompliance *T0kenComplianceTransactorSession) TransferOwner(newOwner common.Address) (*types.Transaction, error) {
	return _T0kenCompliance.Contract.TransferOwner(&_T0kenCompliance.TransactOpts, newOwner)
}

// T0kenComplianceAddressFrozenIterator is returned from FilterAddressFrozen and is used to iterate over the raw logs and unpacked data for AddressFrozen events raised by the T0kenCompliance contract.
type T0kenComplianceAddressFrozenIterator struct {
	Event *T0kenComplianceAddressFrozen // Event containing the contract specifics and raw log

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
func (it *T0kenComplianceAddressFrozenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(T0kenComplianceAddressFrozen)
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
		it.Event = new(T0kenComplianceAddressFrozen)
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
func (it *T0kenComplianceAddressFrozenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *T0kenComplianceAddressFrozenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// T0kenComplianceAddressFrozen represents a AddressFrozen event raised by the T0kenCompliance contract.
type T0kenComplianceAddressFrozen struct {
	Addr     common.Address
	IsFrozen bool
	Owner    common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterAddressFrozen is a free log retrieval operation binding the contract event 0x7fa523c84ab8d7fc5b72f08b9e46dbbf10c39e119a075b3e317002d14bc9f436.
//
// Solidity: e AddressFrozen(addr indexed address, isFrozen indexed bool, owner indexed address)
func (_T0kenCompliance *T0kenComplianceFilterer) FilterAddressFrozen(opts *bind.FilterOpts, addr []common.Address, isFrozen []bool, owner []common.Address) (*T0kenComplianceAddressFrozenIterator, error) {

	var addrRule []interface{}
	for _, addrItem := range addr {
		addrRule = append(addrRule, addrItem)
	}
	var isFrozenRule []interface{}
	for _, isFrozenItem := range isFrozen {
		isFrozenRule = append(isFrozenRule, isFrozenItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _T0kenCompliance.contract.FilterLogs(opts, "AddressFrozen", addrRule, isFrozenRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return &T0kenComplianceAddressFrozenIterator{contract: _T0kenCompliance.contract, event: "AddressFrozen", logs: logs, sub: sub}, nil
}

// WatchAddressFrozen is a free log subscription operation binding the contract event 0x7fa523c84ab8d7fc5b72f08b9e46dbbf10c39e119a075b3e317002d14bc9f436.
//
// Solidity: e AddressFrozen(addr indexed address, isFrozen indexed bool, owner indexed address)
func (_T0kenCompliance *T0kenComplianceFilterer) WatchAddressFrozen(opts *bind.WatchOpts, sink chan<- *T0kenComplianceAddressFrozen, addr []common.Address, isFrozen []bool, owner []common.Address) (event.Subscription, error) {

	var addrRule []interface{}
	for _, addrItem := range addr {
		addrRule = append(addrRule, addrItem)
	}
	var isFrozenRule []interface{}
	for _, isFrozenItem := range isFrozen {
		isFrozenRule = append(isFrozenRule, isFrozenItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _T0kenCompliance.contract.WatchLogs(opts, "AddressFrozen", addrRule, isFrozenRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(T0kenComplianceAddressFrozen)
				if err := _T0kenCompliance.contract.UnpackLog(event, "AddressFrozen", log); err != nil {
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

// T0kenComplianceAdminAddedIterator is returned from FilterAdminAdded and is used to iterate over the raw logs and unpacked data for AdminAdded events raised by the T0kenCompliance contract.
type T0kenComplianceAdminAddedIterator struct {
	Event *T0kenComplianceAdminAdded // Event containing the contract specifics and raw log

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
func (it *T0kenComplianceAdminAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(T0kenComplianceAdminAdded)
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
		it.Event = new(T0kenComplianceAdminAdded)
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
func (it *T0kenComplianceAdminAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *T0kenComplianceAdminAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// T0kenComplianceAdminAdded represents a AdminAdded event raised by the T0kenCompliance contract.
type T0kenComplianceAdminAdded struct {
	Admin common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterAdminAdded is a free log retrieval operation binding the contract event 0x44d6d25963f097ad14f29f06854a01f575648a1ef82f30e562ccd3889717e339.
//
// Solidity: e AdminAdded(admin indexed address)
func (_T0kenCompliance *T0kenComplianceFilterer) FilterAdminAdded(opts *bind.FilterOpts, admin []common.Address) (*T0kenComplianceAdminAddedIterator, error) {

	var adminRule []interface{}
	for _, adminItem := range admin {
		adminRule = append(adminRule, adminItem)
	}

	logs, sub, err := _T0kenCompliance.contract.FilterLogs(opts, "AdminAdded", adminRule)
	if err != nil {
		return nil, err
	}
	return &T0kenComplianceAdminAddedIterator{contract: _T0kenCompliance.contract, event: "AdminAdded", logs: logs, sub: sub}, nil
}

// WatchAdminAdded is a free log subscription operation binding the contract event 0x44d6d25963f097ad14f29f06854a01f575648a1ef82f30e562ccd3889717e339.
//
// Solidity: e AdminAdded(admin indexed address)
func (_T0kenCompliance *T0kenComplianceFilterer) WatchAdminAdded(opts *bind.WatchOpts, sink chan<- *T0kenComplianceAdminAdded, admin []common.Address) (event.Subscription, error) {

	var adminRule []interface{}
	for _, adminItem := range admin {
		adminRule = append(adminRule, adminItem)
	}

	logs, sub, err := _T0kenCompliance.contract.WatchLogs(opts, "AdminAdded", adminRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(T0kenComplianceAdminAdded)
				if err := _T0kenCompliance.contract.UnpackLog(event, "AdminAdded", log); err != nil {
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

// T0kenComplianceAdminRemovedIterator is returned from FilterAdminRemoved and is used to iterate over the raw logs and unpacked data for AdminRemoved events raised by the T0kenCompliance contract.
type T0kenComplianceAdminRemovedIterator struct {
	Event *T0kenComplianceAdminRemoved // Event containing the contract specifics and raw log

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
func (it *T0kenComplianceAdminRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(T0kenComplianceAdminRemoved)
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
		it.Event = new(T0kenComplianceAdminRemoved)
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
func (it *T0kenComplianceAdminRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *T0kenComplianceAdminRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// T0kenComplianceAdminRemoved represents a AdminRemoved event raised by the T0kenCompliance contract.
type T0kenComplianceAdminRemoved struct {
	Admin common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterAdminRemoved is a free log retrieval operation binding the contract event 0xa3b62bc36326052d97ea62d63c3d60308ed4c3ea8ac079dd8499f1e9c4f80c0f.
//
// Solidity: e AdminRemoved(admin indexed address)
func (_T0kenCompliance *T0kenComplianceFilterer) FilterAdminRemoved(opts *bind.FilterOpts, admin []common.Address) (*T0kenComplianceAdminRemovedIterator, error) {

	var adminRule []interface{}
	for _, adminItem := range admin {
		adminRule = append(adminRule, adminItem)
	}

	logs, sub, err := _T0kenCompliance.contract.FilterLogs(opts, "AdminRemoved", adminRule)
	if err != nil {
		return nil, err
	}
	return &T0kenComplianceAdminRemovedIterator{contract: _T0kenCompliance.contract, event: "AdminRemoved", logs: logs, sub: sub}, nil
}

// WatchAdminRemoved is a free log subscription operation binding the contract event 0xa3b62bc36326052d97ea62d63c3d60308ed4c3ea8ac079dd8499f1e9c4f80c0f.
//
// Solidity: e AdminRemoved(admin indexed address)
func (_T0kenCompliance *T0kenComplianceFilterer) WatchAdminRemoved(opts *bind.WatchOpts, sink chan<- *T0kenComplianceAdminRemoved, admin []common.Address) (event.Subscription, error) {

	var adminRule []interface{}
	for _, adminItem := range admin {
		adminRule = append(adminRule, adminItem)
	}

	logs, sub, err := _T0kenCompliance.contract.WatchLogs(opts, "AdminRemoved", adminRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(T0kenComplianceAdminRemoved)
				if err := _T0kenCompliance.contract.UnpackLog(event, "AdminRemoved", log); err != nil {
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

// T0kenComplianceOwnerTransferredIterator is returned from FilterOwnerTransferred and is used to iterate over the raw logs and unpacked data for OwnerTransferred events raised by the T0kenCompliance contract.
type T0kenComplianceOwnerTransferredIterator struct {
	Event *T0kenComplianceOwnerTransferred // Event containing the contract specifics and raw log

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
func (it *T0kenComplianceOwnerTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(T0kenComplianceOwnerTransferred)
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
		it.Event = new(T0kenComplianceOwnerTransferred)
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
func (it *T0kenComplianceOwnerTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *T0kenComplianceOwnerTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// T0kenComplianceOwnerTransferred represents a OwnerTransferred event raised by the T0kenCompliance contract.
type T0kenComplianceOwnerTransferred struct {
	OldOwner common.Address
	NewOwner common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOwnerTransferred is a free log retrieval operation binding the contract event 0x8934ce4adea8d9ce0d714d2c22b86790e41b7731c84b926fbbdc1d40ff6533c9.
//
// Solidity: e OwnerTransferred(oldOwner indexed address, newOwner indexed address)
func (_T0kenCompliance *T0kenComplianceFilterer) FilterOwnerTransferred(opts *bind.FilterOpts, oldOwner []common.Address, newOwner []common.Address) (*T0kenComplianceOwnerTransferredIterator, error) {

	var oldOwnerRule []interface{}
	for _, oldOwnerItem := range oldOwner {
		oldOwnerRule = append(oldOwnerRule, oldOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _T0kenCompliance.contract.FilterLogs(opts, "OwnerTransferred", oldOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &T0kenComplianceOwnerTransferredIterator{contract: _T0kenCompliance.contract, event: "OwnerTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnerTransferred is a free log subscription operation binding the contract event 0x8934ce4adea8d9ce0d714d2c22b86790e41b7731c84b926fbbdc1d40ff6533c9.
//
// Solidity: e OwnerTransferred(oldOwner indexed address, newOwner indexed address)
func (_T0kenCompliance *T0kenComplianceFilterer) WatchOwnerTransferred(opts *bind.WatchOpts, sink chan<- *T0kenComplianceOwnerTransferred, oldOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var oldOwnerRule []interface{}
	for _, oldOwnerItem := range oldOwner {
		oldOwnerRule = append(oldOwnerRule, oldOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _T0kenCompliance.contract.WatchLogs(opts, "OwnerTransferred", oldOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(T0kenComplianceOwnerTransferred)
				if err := _T0kenCompliance.contract.UnpackLog(event, "OwnerTransferred", log); err != nil {
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
