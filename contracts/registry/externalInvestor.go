// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package registry

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

// ExternalInvestorABI is the input ABI used to generate the binding from.
const ExternalInvestorABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"investor\",\"type\":\"address\"},{\"name\":\"accreditation\",\"type\":\"uint48\"}],\"name\":\"setAccreditation\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"admin\",\"type\":\"address\"}],\"name\":\"removeAdmin\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"getHash\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"locked\",\"type\":\"bool\"}],\"name\":\"setLocked\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"investor\",\"type\":\"address\"},{\"name\":\"country\",\"type\":\"bytes2\"}],\"name\":\"setCountry\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"isAdmin\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"investor\",\"type\":\"address\"}],\"name\":\"remove\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"kill\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"investor\",\"type\":\"address\"},{\"name\":\"hash\",\"type\":\"bytes32\"},{\"name\":\"country\",\"type\":\"bytes2\"},{\"name\":\"accreditation\",\"type\":\"uint48\"}],\"name\":\"add\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwner\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"getAccreditation\",\"outputs\":[{\"name\":\"\",\"type\":\"uint48\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"admin\",\"type\":\"address\"}],\"name\":\"addAdmin\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"s\",\"type\":\"address\"}],\"name\":\"setStorage\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"store\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isLocked\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"admins\",\"outputs\":[{\"name\":\"count\",\"type\":\"int256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"investor\",\"type\":\"address\"},{\"name\":\"frozen\",\"type\":\"bool\"}],\"name\":\"setFrozen\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"getCountry\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes2\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"investor\",\"type\":\"address\"},{\"name\":\"hash\",\"type\":\"bytes32\"}],\"name\":\"setHash\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"admin\",\"type\":\"address\"}],\"name\":\"AdminAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"admin\",\"type\":\"address\"}],\"name\":\"AdminRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"oldOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnerTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"investor\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"InvestorAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"investor\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"InvestorRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"investor\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"frozen\",\"type\":\"bool\"},{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"InvestorFrozen\",\"type\":\"event\"}]"

// ExternalInvestorBin is the compiled bytecode used for deploying new contracts.
const ExternalInvestorBin = `608060405260008054600160a060020a031916331790556004805460ff1916905561194e8061002f6000396000f3fe608060405234801561001057600080fd5b5060043610610189576000357c010000000000000000000000000000000000000000000000000000000090048063701ae59f116100ea578063a4e2d6341161009e578063ac869cd811610083578063ac869cd814610448578063d821f92d14610476578063e84b8169146104d157610189565b8063a4e2d63414610438578063a5de36191461044057610189565b80638da5cb5b116100cf5780638da5cb5b146103e65780639137c1a71461040a578063975057e71461043057610189565b8063701ae59f1461037f57806370480275146103c057610189565b806324d7806c1161014157806341c0e1b51161012657806341c0e1b5146102ef57806345eb8a9b146102f75780634fb2e45d1461035957610189565b806324d7806c1461028f57806329092d0e146102c957610189565b80631da0b8fc116101725780631da0b8fc146101ea578063211e28b614610222578063242371a31461024157610189565b80631518c5bc1461018e5780631785f53c146101c4575b600080fd5b6101c2600480360360408110156101a457600080fd5b508035600160a060020a0316906020013565ffffffffffff166104fd565b005b6101c2600480360360208110156101da57600080fd5b5035600160a060020a03166106df565b6102106004803603602081101561020057600080fd5b5035600160a060020a03166107e9565b60408051918252519081900360200190f35b6101c26004803603602081101561023857600080fd5b50351515610894565b6101c26004803603604081101561025757600080fd5b508035600160a060020a031690602001357fffff0000000000000000000000000000000000000000000000000000000000001661095e565b6102b5600480360360208110156102a557600080fd5b5035600160a060020a0316610ab8565b604080519115158252519081900360200190f35b6101c2600480360360208110156102df57600080fd5b5035600160a060020a0316610ae7565b6101c2610c41565b6101c26004803603608081101561030d57600080fd5b508035600160a060020a03169060208101359060408101357fffff00000000000000000000000000000000000000000000000000000000000016906060013565ffffffffffff16610cb6565b6101c26004803603602081101561036f57600080fd5b5035600160a060020a0316610f6f565b6103a56004803603602081101561039557600080fd5b5035600160a060020a03166110ed565b6040805165ffffffffffff9092168252519081900360200190f35b6101c2600480360360208110156103d657600080fd5b5035600160a060020a031661119d565b6103ee6112a7565b60408051600160a060020a039092168252519081900360200190f35b6101c26004803603602081101561042057600080fd5b5035600160a060020a03166112b6565b6103ee61135c565b6102b5611370565b610210611379565b6101c26004803603604081101561045e57600080fd5b50600160a060020a038135169060200135151561137f565b61049c6004803603602081101561048c57600080fd5b5035600160a060020a03166114cc565b604080517fffff0000000000000000000000000000000000000000000000000000000000009092168252519081900360200190f35b6101c2600480360360408110156104e757600080fd5b50600160a060020a038135169060200135611598565b600054600160a060020a0316331480610522575061052260013363ffffffff6116af16565b151561057d57604080516000805160206118db833981519152815260206004820152601960248201527f41646d696e206163636f756e7420697320726571756972656400000000000000604482015290519081900360640190fd5b60048054604080517f572bc6b1000000000000000000000000000000000000000000000000000000008152600160a060020a03868116948201949094526001602482015290516000936101009093049092169163572bc6b191604480820192602092909190829003018186803b1580156105f657600080fd5b505afa15801561060a573d6000803e3d6000fd5b505050506040513d602081101561062057600080fd5b505190506106418165ffffffffffff841667ffffffffffff000060106116e6565b60048054604080517f4f1c3b66000000000000000000000000000000000000000000000000000000008152600160a060020a03888116948201949094526001602482015260448101859052905193945061010090910490911691634f1c3b669160648082019260009290919082900301818387803b1580156106c257600080fd5b505af11580156106d6573d6000803e3d6000fd5b50505050505050565b600054600160a060020a0316331461074657604080516000805160206118db833981519152815260206004820152601960248201527f4f776e6572206163636f756e7420697320726571756972656400000000000000604482015290519081900360640190fd5b61075760018263ffffffff6116fc16565b15156107b257604080516000805160206118db833981519152815260206004820152601660248201527f556e61626c6520746f2072656d6f76652061646d696e00000000000000000000604482015290519081900360640190fd5b604051600160a060020a038216907fa3b62bc36326052d97ea62d63c3d60308ed4c3ea8ac079dd8499f1e9c4f80c0f90600090a250565b60048054604080517f572bc6b1000000000000000000000000000000000000000000000000000000008152600160a060020a0385811694820194909452600060248201819052915191936101009093049092169163572bc6b1916044808301926020929190829003018186803b15801561086257600080fd5b505afa158015610876573d6000803e3d6000fd5b505050506040513d602081101561088c57600080fd5b505192915050565b600054600160a060020a031633146108fb57604080516000805160206118db833981519152815260206004820152601960248201527f4f776e6572206163636f756e7420697320726571756972656400000000000000604482015290519081900360640190fd5b60045460ff161515811515141561094b576040516000805160206118db83398151915281526004018080602001828103825260288152602001806118fb6028913960400191505060405180910390fd5b6004805460ff1916911515919091179055565b600054600160a060020a0316331480610983575061098360013363ffffffff6116af16565b15156109de57604080516000805160206118db833981519152815260206004820152601960248201527f41646d696e206163636f756e7420697320726571756972656400000000000000604482015290519081900360640190fd5b60048054604080517f572bc6b1000000000000000000000000000000000000000000000000000000008152600160a060020a03868116948201949094526001602482015290516000936101009093049092169163572bc6b191604480820192602092909190829003018186803b158015610a5757600080fd5b505afa158015610a6b573d6000803e3d6000fd5b505050506040513d6020811015610a8157600080fd5b505190506106418161ffff7e01000000000000000000000000000000000000000000000000000000000000850481169060006116e6565b60008054600160a060020a0383811691161480610ae15750610ae160018363ffffffff6116af16565b92915050565b600054600160a060020a0316331480610b0c5750610b0c60013363ffffffff6116af16565b1515610b6757604080516000805160206118db833981519152815260206004820152601960248201527f41646d696e206163636f756e7420697320726571756972656400000000000000604482015290519081900360640190fd5b600460019054906101000a9004600160a060020a0316600160a060020a031663c4740a95826040518263ffffffff167c01000000000000000000000000000000000000000000000000000000000281526004018082600160a060020a0316600160a060020a03168152602001915050600060405180830381600087803b158015610bf057600080fd5b505af1158015610c04573d6000803e3d6000fd5b5050604051339250600160a060020a03841691507fd8755221287ca1f6b28807977a086f5534d9e02ea27ebad003d7cb1a95659a4690600090a350565b600054600160a060020a03163314610ca857604080516000805160206118db833981519152815260206004820152601960248201527f4f776e6572206163636f756e7420697320726571756972656400000000000000604482015290519081900360640190fd5b600054600160a060020a0316ff5b600054600160a060020a0316331480610cdb5750610cdb60013363ffffffff6116af16565b1515610d3657604080516000805160206118db833981519152815260206004820152601960248201527f41646d696e206163636f756e7420697320726571756972656400000000000000604482015290519081900360640190fd5b60048054604080517f116c92b7000000000000000000000000000000000000000000000000000000008152600160a060020a0388811694820194909452600560248201526000604482018190526064820181905291517e01000000000000000000000000000000000000000000000000000000000000870461ffff1665ffffffffffff8716620100000217946101009094049093169263116c92b79260848084019391929182900301818387803b158015610df057600080fd5b505af1158015610e04573d6000803e3d6000fd5b505060048054604080517f4f1c3b66000000000000000000000000000000000000000000000000000000008152600160a060020a038b811694820194909452600060248201819052604482018b905291516101009093049093169450634f1c3b66935060648084019391929182900301818387803b158015610e8557600080fd5b505af1158015610e99573d6000803e3d6000fd5b505060048054604080517f4f1c3b66000000000000000000000000000000000000000000000000000000008152600160a060020a038b811694820194909452600160248201526044810187905290516101009092049092169350634f1c3b669250606480830192600092919082900301818387803b158015610f1a57600080fd5b505af1158015610f2e573d6000803e3d6000fd5b5050604051339250600160a060020a03881691507fe99183cc0b1657b54afa611991294ec1e4c458d7c36910518e2a5b76b2b6e73f90600090a35050505050565b600054600160a060020a03163314610fd657604080516000805160206118db833981519152815260206004820152601960248201527f4f776e6572206163636f756e7420697320726571756972656400000000000000604482015290519081900360640190fd5b600054600160a060020a038281169116141561102b576040516000805160206118db83398151915281526004018080602001828103825260258152602001806118b66025913960400191505060405180910390fd5b600160a060020a038116151561109057604080516000805160206118db833981519152815260206004820181905260248201527f4e6577204f776e65722063616e6e6f74206265207a65726f2061646472657373604482015290519081900360640190fd5b60008054600160a060020a0383811673ffffffffffffffffffffffffffffffffffffffff19831681178455604051919092169283917f8934ce4adea8d9ce0d714d2c22b86790e41b7731c84b926fbbdc1d40ff6533c99190a35050565b60048054604080517f572bc6b1000000000000000000000000000000000000000000000000000000008152600160a060020a038581169482019490945260016024820152905160009384936101009004169163572bc6b1916044808301926020929190829003018186803b15801561116457600080fd5b505afa158015611178573d6000803e3d6000fd5b505050506040513d602081101561118e57600080fd5b50516201000090049392505050565b600054600160a060020a0316331461120457604080516000805160206118db833981519152815260206004820152601960248201527f4f776e6572206163636f756e7420697320726571756972656400000000000000604482015290519081900360640190fd5b61121560018263ffffffff61180116565b151561127057604080516000805160206118db833981519152815260206004820152601360248201527f556e61626c6520746f206164642061646d696e00000000000000000000000000604482015290519081900360640190fd5b604051600160a060020a038216907f44d6d25963f097ad14f29f06854a01f575648a1ef82f30e562ccd3889717e33990600090a250565b600054600160a060020a031681565b600054600160a060020a0316331461131d57604080516000805160206118db833981519152815260206004820152601960248201527f4f776e6572206163636f756e7420697320726571756972656400000000000000604482015290519081900360640190fd5b60048054600160a060020a03909216610100027fffffffffffffffffffffff0000000000000000000000000000000000000000ff909216919091179055565b6004546101009004600160a060020a031681565b60045460ff1681565b60015481565b600054600160a060020a03163314806113a457506113a460013363ffffffff6116af16565b15156113ff57604080516000805160206118db833981519152815260206004820152601960248201527f41646d696e206163636f756e7420697320726571756972656400000000000000604482015290519081900360640190fd5b60048054604080517fcbe5404f000000000000000000000000000000000000000000000000000000008152600160a060020a0386811694820194909452841515602482015290516101009092049092169163cbe5404f91604480830192600092919082900301818387803b15801561147657600080fd5b505af115801561148a573d6000803e3d6000fd5b50506040513392508315159150600160a060020a038516907f6d30448ca28c66e149273ddd6d39fe9cb1982d4013649080aeb9d8251356d38190600090a45050565b60048054604080517f572bc6b1000000000000000000000000000000000000000000000000000000008152600160a060020a03858116948201949094526001602482015290516000936101009093049092169163572bc6b191604480820192602092909190829003018186803b15801561154557600080fd5b505afa158015611559573d6000803e3d6000fd5b505050506040513d602081101561156f57600080fd5b50517e010000000000000000000000000000000000000000000000000000000000000292915050565b600054600160a060020a03163314806115bd57506115bd60013363ffffffff6116af16565b151561161857604080516000805160206118db833981519152815260206004820152601960248201527f41646d696e206163636f756e7420697320726571756972656400000000000000604482015290519081900360640190fd5b60048054604080517f4f1c3b66000000000000000000000000000000000000000000000000000000008152600160a060020a038681169482019490945260006024820181905260448201869052915161010090930490931692634f1c3b66926064808301939282900301818387803b15801561169357600080fd5b505af11580156116a7573d6000803e3d6000fd5b505050505050565b600160a060020a0381166000908152600183016020526040812054600019018181128015906116de5750835481125b949350505050565b60ff1660020a9190910281169019919091161790565b600160a060020a038116600090815260018084016020526040822054908112806117265750835481135b15611735576000915050610ae1565b83548112156117a95783546000908152600285016020818152604080842054600160a060020a031680855260018901835281852086905585855292909152808320805473ffffffffffffffffffffffffffffffffffffffff19908116909317905586548352909120805490911690556117d5565b60008181526002850160205260409020805473ffffffffffffffffffffffffffffffffffffffff191690555b5050600160a060020a031660009081526001828101602052604082209190915581546000190190915590565b6000600160a060020a038216151561181b57506000610ae1565b600160a060020a03821660009081526001840160205260408120546000190190811280159061184a5750835481125b15611859576000915050610ae1565b505081546001908101808455600160a060020a03831660008181528386016020908152604080832085905593825260028701905291909120805473ffffffffffffffffffffffffffffffffffffffff191690911790559291505056fe4e6577204f776e65722063616e6e6f74206265207468652063757272656e74206f776e657208c379a000000000000000000000000000000000000000000000000000000000436f6e747261637420616c726561647920696e20726571756573746564206c6f636b207374617465a165627a7a7230582061f627bd5fd50bb5b8533e2502213ea0ffe6081d2349e85e581d0883b6338ee50029`

// DeployExternalInvestor deploys a new Ethereum contract, binding an instance of ExternalInvestor to it.
func DeployExternalInvestor(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ExternalInvestor, error) {
	parsed, err := abi.JSON(strings.NewReader(ExternalInvestorABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ExternalInvestorBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ExternalInvestor{ExternalInvestorCaller: ExternalInvestorCaller{contract: contract}, ExternalInvestorTransactor: ExternalInvestorTransactor{contract: contract}, ExternalInvestorFilterer: ExternalInvestorFilterer{contract: contract}}, nil
}

// ExternalInvestor is an auto generated Go binding around an Ethereum contract.
type ExternalInvestor struct {
	ExternalInvestorCaller     // Read-only binding to the contract
	ExternalInvestorTransactor // Write-only binding to the contract
	ExternalInvestorFilterer   // Log filterer for contract events
}

// ExternalInvestorCaller is an auto generated read-only Go binding around an Ethereum contract.
type ExternalInvestorCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ExternalInvestorTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ExternalInvestorTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ExternalInvestorFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ExternalInvestorFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ExternalInvestorSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ExternalInvestorSession struct {
	Contract     *ExternalInvestor // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ExternalInvestorCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ExternalInvestorCallerSession struct {
	Contract *ExternalInvestorCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// ExternalInvestorTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ExternalInvestorTransactorSession struct {
	Contract     *ExternalInvestorTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// ExternalInvestorRaw is an auto generated low-level Go binding around an Ethereum contract.
type ExternalInvestorRaw struct {
	Contract *ExternalInvestor // Generic contract binding to access the raw methods on
}

// ExternalInvestorCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ExternalInvestorCallerRaw struct {
	Contract *ExternalInvestorCaller // Generic read-only contract binding to access the raw methods on
}

// ExternalInvestorTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ExternalInvestorTransactorRaw struct {
	Contract *ExternalInvestorTransactor // Generic write-only contract binding to access the raw methods on
}

// NewExternalInvestor creates a new instance of ExternalInvestor, bound to a specific deployed contract.
func NewExternalInvestor(address common.Address, backend bind.ContractBackend) (*ExternalInvestor, error) {
	contract, err := bindExternalInvestor(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ExternalInvestor{ExternalInvestorCaller: ExternalInvestorCaller{contract: contract}, ExternalInvestorTransactor: ExternalInvestorTransactor{contract: contract}, ExternalInvestorFilterer: ExternalInvestorFilterer{contract: contract}}, nil
}

// NewExternalInvestorCaller creates a new read-only instance of ExternalInvestor, bound to a specific deployed contract.
func NewExternalInvestorCaller(address common.Address, caller bind.ContractCaller) (*ExternalInvestorCaller, error) {
	contract, err := bindExternalInvestor(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ExternalInvestorCaller{contract: contract}, nil
}

// NewExternalInvestorTransactor creates a new write-only instance of ExternalInvestor, bound to a specific deployed contract.
func NewExternalInvestorTransactor(address common.Address, transactor bind.ContractTransactor) (*ExternalInvestorTransactor, error) {
	contract, err := bindExternalInvestor(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ExternalInvestorTransactor{contract: contract}, nil
}

// NewExternalInvestorFilterer creates a new log filterer instance of ExternalInvestor, bound to a specific deployed contract.
func NewExternalInvestorFilterer(address common.Address, filterer bind.ContractFilterer) (*ExternalInvestorFilterer, error) {
	contract, err := bindExternalInvestor(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ExternalInvestorFilterer{contract: contract}, nil
}

// bindExternalInvestor binds a generic wrapper to an already deployed contract.
func bindExternalInvestor(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ExternalInvestorABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ExternalInvestor *ExternalInvestorRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ExternalInvestor.Contract.ExternalInvestorCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ExternalInvestor *ExternalInvestorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ExternalInvestor.Contract.ExternalInvestorTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ExternalInvestor *ExternalInvestorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ExternalInvestor.Contract.ExternalInvestorTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ExternalInvestor *ExternalInvestorCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ExternalInvestor.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ExternalInvestor *ExternalInvestorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ExternalInvestor.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ExternalInvestor *ExternalInvestorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ExternalInvestor.Contract.contract.Transact(opts, method, params...)
}

// Admins is a free data retrieval call binding the contract method 0xa5de3619.
//
// Solidity: function admins() constant returns(count int256)
func (_ExternalInvestor *ExternalInvestorCaller) Admins(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ExternalInvestor.contract.Call(opts, out, "admins")
	return *ret0, err
}

// Admins is a free data retrieval call binding the contract method 0xa5de3619.
//
// Solidity: function admins() constant returns(count int256)
func (_ExternalInvestor *ExternalInvestorSession) Admins() (*big.Int, error) {
	return _ExternalInvestor.Contract.Admins(&_ExternalInvestor.CallOpts)
}

// Admins is a free data retrieval call binding the contract method 0xa5de3619.
//
// Solidity: function admins() constant returns(count int256)
func (_ExternalInvestor *ExternalInvestorCallerSession) Admins() (*big.Int, error) {
	return _ExternalInvestor.Contract.Admins(&_ExternalInvestor.CallOpts)
}

// GetAccreditation is a free data retrieval call binding the contract method 0x701ae59f.
//
// Solidity: function getAccreditation(addr address) constant returns(uint48)
func (_ExternalInvestor *ExternalInvestorCaller) GetAccreditation(opts *bind.CallOpts, addr common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ExternalInvestor.contract.Call(opts, out, "getAccreditation", addr)
	return *ret0, err
}

// GetAccreditation is a free data retrieval call binding the contract method 0x701ae59f.
//
// Solidity: function getAccreditation(addr address) constant returns(uint48)
func (_ExternalInvestor *ExternalInvestorSession) GetAccreditation(addr common.Address) (*big.Int, error) {
	return _ExternalInvestor.Contract.GetAccreditation(&_ExternalInvestor.CallOpts, addr)
}

// GetAccreditation is a free data retrieval call binding the contract method 0x701ae59f.
//
// Solidity: function getAccreditation(addr address) constant returns(uint48)
func (_ExternalInvestor *ExternalInvestorCallerSession) GetAccreditation(addr common.Address) (*big.Int, error) {
	return _ExternalInvestor.Contract.GetAccreditation(&_ExternalInvestor.CallOpts, addr)
}

// GetCountry is a free data retrieval call binding the contract method 0xd821f92d.
//
// Solidity: function getCountry(addr address) constant returns(bytes2)
func (_ExternalInvestor *ExternalInvestorCaller) GetCountry(opts *bind.CallOpts, addr common.Address) ([2]byte, error) {
	var (
		ret0 = new([2]byte)
	)
	out := ret0
	err := _ExternalInvestor.contract.Call(opts, out, "getCountry", addr)
	return *ret0, err
}

// GetCountry is a free data retrieval call binding the contract method 0xd821f92d.
//
// Solidity: function getCountry(addr address) constant returns(bytes2)
func (_ExternalInvestor *ExternalInvestorSession) GetCountry(addr common.Address) ([2]byte, error) {
	return _ExternalInvestor.Contract.GetCountry(&_ExternalInvestor.CallOpts, addr)
}

// GetCountry is a free data retrieval call binding the contract method 0xd821f92d.
//
// Solidity: function getCountry(addr address) constant returns(bytes2)
func (_ExternalInvestor *ExternalInvestorCallerSession) GetCountry(addr common.Address) ([2]byte, error) {
	return _ExternalInvestor.Contract.GetCountry(&_ExternalInvestor.CallOpts, addr)
}

// GetHash is a free data retrieval call binding the contract method 0x1da0b8fc.
//
// Solidity: function getHash(addr address) constant returns(bytes32)
func (_ExternalInvestor *ExternalInvestorCaller) GetHash(opts *bind.CallOpts, addr common.Address) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _ExternalInvestor.contract.Call(opts, out, "getHash", addr)
	return *ret0, err
}

// GetHash is a free data retrieval call binding the contract method 0x1da0b8fc.
//
// Solidity: function getHash(addr address) constant returns(bytes32)
func (_ExternalInvestor *ExternalInvestorSession) GetHash(addr common.Address) ([32]byte, error) {
	return _ExternalInvestor.Contract.GetHash(&_ExternalInvestor.CallOpts, addr)
}

// GetHash is a free data retrieval call binding the contract method 0x1da0b8fc.
//
// Solidity: function getHash(addr address) constant returns(bytes32)
func (_ExternalInvestor *ExternalInvestorCallerSession) GetHash(addr common.Address) ([32]byte, error) {
	return _ExternalInvestor.Contract.GetHash(&_ExternalInvestor.CallOpts, addr)
}

// IsAdmin is a free data retrieval call binding the contract method 0x24d7806c.
//
// Solidity: function isAdmin(addr address) constant returns(bool)
func (_ExternalInvestor *ExternalInvestorCaller) IsAdmin(opts *bind.CallOpts, addr common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _ExternalInvestor.contract.Call(opts, out, "isAdmin", addr)
	return *ret0, err
}

// IsAdmin is a free data retrieval call binding the contract method 0x24d7806c.
//
// Solidity: function isAdmin(addr address) constant returns(bool)
func (_ExternalInvestor *ExternalInvestorSession) IsAdmin(addr common.Address) (bool, error) {
	return _ExternalInvestor.Contract.IsAdmin(&_ExternalInvestor.CallOpts, addr)
}

// IsAdmin is a free data retrieval call binding the contract method 0x24d7806c.
//
// Solidity: function isAdmin(addr address) constant returns(bool)
func (_ExternalInvestor *ExternalInvestorCallerSession) IsAdmin(addr common.Address) (bool, error) {
	return _ExternalInvestor.Contract.IsAdmin(&_ExternalInvestor.CallOpts, addr)
}

// IsLocked is a free data retrieval call binding the contract method 0xa4e2d634.
//
// Solidity: function isLocked() constant returns(bool)
func (_ExternalInvestor *ExternalInvestorCaller) IsLocked(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _ExternalInvestor.contract.Call(opts, out, "isLocked")
	return *ret0, err
}

// IsLocked is a free data retrieval call binding the contract method 0xa4e2d634.
//
// Solidity: function isLocked() constant returns(bool)
func (_ExternalInvestor *ExternalInvestorSession) IsLocked() (bool, error) {
	return _ExternalInvestor.Contract.IsLocked(&_ExternalInvestor.CallOpts)
}

// IsLocked is a free data retrieval call binding the contract method 0xa4e2d634.
//
// Solidity: function isLocked() constant returns(bool)
func (_ExternalInvestor *ExternalInvestorCallerSession) IsLocked() (bool, error) {
	return _ExternalInvestor.Contract.IsLocked(&_ExternalInvestor.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_ExternalInvestor *ExternalInvestorCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ExternalInvestor.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_ExternalInvestor *ExternalInvestorSession) Owner() (common.Address, error) {
	return _ExternalInvestor.Contract.Owner(&_ExternalInvestor.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_ExternalInvestor *ExternalInvestorCallerSession) Owner() (common.Address, error) {
	return _ExternalInvestor.Contract.Owner(&_ExternalInvestor.CallOpts)
}

// Store is a free data retrieval call binding the contract method 0x975057e7.
//
// Solidity: function store() constant returns(address)
func (_ExternalInvestor *ExternalInvestorCaller) Store(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ExternalInvestor.contract.Call(opts, out, "store")
	return *ret0, err
}

// Store is a free data retrieval call binding the contract method 0x975057e7.
//
// Solidity: function store() constant returns(address)
func (_ExternalInvestor *ExternalInvestorSession) Store() (common.Address, error) {
	return _ExternalInvestor.Contract.Store(&_ExternalInvestor.CallOpts)
}

// Store is a free data retrieval call binding the contract method 0x975057e7.
//
// Solidity: function store() constant returns(address)
func (_ExternalInvestor *ExternalInvestorCallerSession) Store() (common.Address, error) {
	return _ExternalInvestor.Contract.Store(&_ExternalInvestor.CallOpts)
}

// Add is a paid mutator transaction binding the contract method 0x45eb8a9b.
//
// Solidity: function add(investor address, hash bytes32, country bytes2, accreditation uint48) returns()
func (_ExternalInvestor *ExternalInvestorTransactor) Add(opts *bind.TransactOpts, investor common.Address, hash [32]byte, country [2]byte, accreditation *big.Int) (*types.Transaction, error) {
	return _ExternalInvestor.contract.Transact(opts, "add", investor, hash, country, accreditation)
}

// Add is a paid mutator transaction binding the contract method 0x45eb8a9b.
//
// Solidity: function add(investor address, hash bytes32, country bytes2, accreditation uint48) returns()
func (_ExternalInvestor *ExternalInvestorSession) Add(investor common.Address, hash [32]byte, country [2]byte, accreditation *big.Int) (*types.Transaction, error) {
	return _ExternalInvestor.Contract.Add(&_ExternalInvestor.TransactOpts, investor, hash, country, accreditation)
}

// Add is a paid mutator transaction binding the contract method 0x45eb8a9b.
//
// Solidity: function add(investor address, hash bytes32, country bytes2, accreditation uint48) returns()
func (_ExternalInvestor *ExternalInvestorTransactorSession) Add(investor common.Address, hash [32]byte, country [2]byte, accreditation *big.Int) (*types.Transaction, error) {
	return _ExternalInvestor.Contract.Add(&_ExternalInvestor.TransactOpts, investor, hash, country, accreditation)
}

// AddAdmin is a paid mutator transaction binding the contract method 0x70480275.
//
// Solidity: function addAdmin(admin address) returns()
func (_ExternalInvestor *ExternalInvestorTransactor) AddAdmin(opts *bind.TransactOpts, admin common.Address) (*types.Transaction, error) {
	return _ExternalInvestor.contract.Transact(opts, "addAdmin", admin)
}

// AddAdmin is a paid mutator transaction binding the contract method 0x70480275.
//
// Solidity: function addAdmin(admin address) returns()
func (_ExternalInvestor *ExternalInvestorSession) AddAdmin(admin common.Address) (*types.Transaction, error) {
	return _ExternalInvestor.Contract.AddAdmin(&_ExternalInvestor.TransactOpts, admin)
}

// AddAdmin is a paid mutator transaction binding the contract method 0x70480275.
//
// Solidity: function addAdmin(admin address) returns()
func (_ExternalInvestor *ExternalInvestorTransactorSession) AddAdmin(admin common.Address) (*types.Transaction, error) {
	return _ExternalInvestor.Contract.AddAdmin(&_ExternalInvestor.TransactOpts, admin)
}

// Kill is a paid mutator transaction binding the contract method 0x41c0e1b5.
//
// Solidity: function kill() returns()
func (_ExternalInvestor *ExternalInvestorTransactor) Kill(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ExternalInvestor.contract.Transact(opts, "kill")
}

// Kill is a paid mutator transaction binding the contract method 0x41c0e1b5.
//
// Solidity: function kill() returns()
func (_ExternalInvestor *ExternalInvestorSession) Kill() (*types.Transaction, error) {
	return _ExternalInvestor.Contract.Kill(&_ExternalInvestor.TransactOpts)
}

// Kill is a paid mutator transaction binding the contract method 0x41c0e1b5.
//
// Solidity: function kill() returns()
func (_ExternalInvestor *ExternalInvestorTransactorSession) Kill() (*types.Transaction, error) {
	return _ExternalInvestor.Contract.Kill(&_ExternalInvestor.TransactOpts)
}

// Remove is a paid mutator transaction binding the contract method 0x29092d0e.
//
// Solidity: function remove(investor address) returns()
func (_ExternalInvestor *ExternalInvestorTransactor) Remove(opts *bind.TransactOpts, investor common.Address) (*types.Transaction, error) {
	return _ExternalInvestor.contract.Transact(opts, "remove", investor)
}

// Remove is a paid mutator transaction binding the contract method 0x29092d0e.
//
// Solidity: function remove(investor address) returns()
func (_ExternalInvestor *ExternalInvestorSession) Remove(investor common.Address) (*types.Transaction, error) {
	return _ExternalInvestor.Contract.Remove(&_ExternalInvestor.TransactOpts, investor)
}

// Remove is a paid mutator transaction binding the contract method 0x29092d0e.
//
// Solidity: function remove(investor address) returns()
func (_ExternalInvestor *ExternalInvestorTransactorSession) Remove(investor common.Address) (*types.Transaction, error) {
	return _ExternalInvestor.Contract.Remove(&_ExternalInvestor.TransactOpts, investor)
}

// RemoveAdmin is a paid mutator transaction binding the contract method 0x1785f53c.
//
// Solidity: function removeAdmin(admin address) returns()
func (_ExternalInvestor *ExternalInvestorTransactor) RemoveAdmin(opts *bind.TransactOpts, admin common.Address) (*types.Transaction, error) {
	return _ExternalInvestor.contract.Transact(opts, "removeAdmin", admin)
}

// RemoveAdmin is a paid mutator transaction binding the contract method 0x1785f53c.
//
// Solidity: function removeAdmin(admin address) returns()
func (_ExternalInvestor *ExternalInvestorSession) RemoveAdmin(admin common.Address) (*types.Transaction, error) {
	return _ExternalInvestor.Contract.RemoveAdmin(&_ExternalInvestor.TransactOpts, admin)
}

// RemoveAdmin is a paid mutator transaction binding the contract method 0x1785f53c.
//
// Solidity: function removeAdmin(admin address) returns()
func (_ExternalInvestor *ExternalInvestorTransactorSession) RemoveAdmin(admin common.Address) (*types.Transaction, error) {
	return _ExternalInvestor.Contract.RemoveAdmin(&_ExternalInvestor.TransactOpts, admin)
}

// SetAccreditation is a paid mutator transaction binding the contract method 0x1518c5bc.
//
// Solidity: function setAccreditation(investor address, accreditation uint48) returns()
func (_ExternalInvestor *ExternalInvestorTransactor) SetAccreditation(opts *bind.TransactOpts, investor common.Address, accreditation *big.Int) (*types.Transaction, error) {
	return _ExternalInvestor.contract.Transact(opts, "setAccreditation", investor, accreditation)
}

// SetAccreditation is a paid mutator transaction binding the contract method 0x1518c5bc.
//
// Solidity: function setAccreditation(investor address, accreditation uint48) returns()
func (_ExternalInvestor *ExternalInvestorSession) SetAccreditation(investor common.Address, accreditation *big.Int) (*types.Transaction, error) {
	return _ExternalInvestor.Contract.SetAccreditation(&_ExternalInvestor.TransactOpts, investor, accreditation)
}

// SetAccreditation is a paid mutator transaction binding the contract method 0x1518c5bc.
//
// Solidity: function setAccreditation(investor address, accreditation uint48) returns()
func (_ExternalInvestor *ExternalInvestorTransactorSession) SetAccreditation(investor common.Address, accreditation *big.Int) (*types.Transaction, error) {
	return _ExternalInvestor.Contract.SetAccreditation(&_ExternalInvestor.TransactOpts, investor, accreditation)
}

// SetCountry is a paid mutator transaction binding the contract method 0x242371a3.
//
// Solidity: function setCountry(investor address, country bytes2) returns()
func (_ExternalInvestor *ExternalInvestorTransactor) SetCountry(opts *bind.TransactOpts, investor common.Address, country [2]byte) (*types.Transaction, error) {
	return _ExternalInvestor.contract.Transact(opts, "setCountry", investor, country)
}

// SetCountry is a paid mutator transaction binding the contract method 0x242371a3.
//
// Solidity: function setCountry(investor address, country bytes2) returns()
func (_ExternalInvestor *ExternalInvestorSession) SetCountry(investor common.Address, country [2]byte) (*types.Transaction, error) {
	return _ExternalInvestor.Contract.SetCountry(&_ExternalInvestor.TransactOpts, investor, country)
}

// SetCountry is a paid mutator transaction binding the contract method 0x242371a3.
//
// Solidity: function setCountry(investor address, country bytes2) returns()
func (_ExternalInvestor *ExternalInvestorTransactorSession) SetCountry(investor common.Address, country [2]byte) (*types.Transaction, error) {
	return _ExternalInvestor.Contract.SetCountry(&_ExternalInvestor.TransactOpts, investor, country)
}

// SetFrozen is a paid mutator transaction binding the contract method 0xac869cd8.
//
// Solidity: function setFrozen(investor address, frozen bool) returns()
func (_ExternalInvestor *ExternalInvestorTransactor) SetFrozen(opts *bind.TransactOpts, investor common.Address, frozen bool) (*types.Transaction, error) {
	return _ExternalInvestor.contract.Transact(opts, "setFrozen", investor, frozen)
}

// SetFrozen is a paid mutator transaction binding the contract method 0xac869cd8.
//
// Solidity: function setFrozen(investor address, frozen bool) returns()
func (_ExternalInvestor *ExternalInvestorSession) SetFrozen(investor common.Address, frozen bool) (*types.Transaction, error) {
	return _ExternalInvestor.Contract.SetFrozen(&_ExternalInvestor.TransactOpts, investor, frozen)
}

// SetFrozen is a paid mutator transaction binding the contract method 0xac869cd8.
//
// Solidity: function setFrozen(investor address, frozen bool) returns()
func (_ExternalInvestor *ExternalInvestorTransactorSession) SetFrozen(investor common.Address, frozen bool) (*types.Transaction, error) {
	return _ExternalInvestor.Contract.SetFrozen(&_ExternalInvestor.TransactOpts, investor, frozen)
}

// SetHash is a paid mutator transaction binding the contract method 0xe84b8169.
//
// Solidity: function setHash(investor address, hash bytes32) returns()
func (_ExternalInvestor *ExternalInvestorTransactor) SetHash(opts *bind.TransactOpts, investor common.Address, hash [32]byte) (*types.Transaction, error) {
	return _ExternalInvestor.contract.Transact(opts, "setHash", investor, hash)
}

// SetHash is a paid mutator transaction binding the contract method 0xe84b8169.
//
// Solidity: function setHash(investor address, hash bytes32) returns()
func (_ExternalInvestor *ExternalInvestorSession) SetHash(investor common.Address, hash [32]byte) (*types.Transaction, error) {
	return _ExternalInvestor.Contract.SetHash(&_ExternalInvestor.TransactOpts, investor, hash)
}

// SetHash is a paid mutator transaction binding the contract method 0xe84b8169.
//
// Solidity: function setHash(investor address, hash bytes32) returns()
func (_ExternalInvestor *ExternalInvestorTransactorSession) SetHash(investor common.Address, hash [32]byte) (*types.Transaction, error) {
	return _ExternalInvestor.Contract.SetHash(&_ExternalInvestor.TransactOpts, investor, hash)
}

// SetLocked is a paid mutator transaction binding the contract method 0x211e28b6.
//
// Solidity: function setLocked(locked bool) returns()
func (_ExternalInvestor *ExternalInvestorTransactor) SetLocked(opts *bind.TransactOpts, locked bool) (*types.Transaction, error) {
	return _ExternalInvestor.contract.Transact(opts, "setLocked", locked)
}

// SetLocked is a paid mutator transaction binding the contract method 0x211e28b6.
//
// Solidity: function setLocked(locked bool) returns()
func (_ExternalInvestor *ExternalInvestorSession) SetLocked(locked bool) (*types.Transaction, error) {
	return _ExternalInvestor.Contract.SetLocked(&_ExternalInvestor.TransactOpts, locked)
}

// SetLocked is a paid mutator transaction binding the contract method 0x211e28b6.
//
// Solidity: function setLocked(locked bool) returns()
func (_ExternalInvestor *ExternalInvestorTransactorSession) SetLocked(locked bool) (*types.Transaction, error) {
	return _ExternalInvestor.Contract.SetLocked(&_ExternalInvestor.TransactOpts, locked)
}

// SetStorage is a paid mutator transaction binding the contract method 0x9137c1a7.
//
// Solidity: function setStorage(s address) returns()
func (_ExternalInvestor *ExternalInvestorTransactor) SetStorage(opts *bind.TransactOpts, s common.Address) (*types.Transaction, error) {
	return _ExternalInvestor.contract.Transact(opts, "setStorage", s)
}

// SetStorage is a paid mutator transaction binding the contract method 0x9137c1a7.
//
// Solidity: function setStorage(s address) returns()
func (_ExternalInvestor *ExternalInvestorSession) SetStorage(s common.Address) (*types.Transaction, error) {
	return _ExternalInvestor.Contract.SetStorage(&_ExternalInvestor.TransactOpts, s)
}

// SetStorage is a paid mutator transaction binding the contract method 0x9137c1a7.
//
// Solidity: function setStorage(s address) returns()
func (_ExternalInvestor *ExternalInvestorTransactorSession) SetStorage(s common.Address) (*types.Transaction, error) {
	return _ExternalInvestor.Contract.SetStorage(&_ExternalInvestor.TransactOpts, s)
}

// TransferOwner is a paid mutator transaction binding the contract method 0x4fb2e45d.
//
// Solidity: function transferOwner(newOwner address) returns()
func (_ExternalInvestor *ExternalInvestorTransactor) TransferOwner(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _ExternalInvestor.contract.Transact(opts, "transferOwner", newOwner)
}

// TransferOwner is a paid mutator transaction binding the contract method 0x4fb2e45d.
//
// Solidity: function transferOwner(newOwner address) returns()
func (_ExternalInvestor *ExternalInvestorSession) TransferOwner(newOwner common.Address) (*types.Transaction, error) {
	return _ExternalInvestor.Contract.TransferOwner(&_ExternalInvestor.TransactOpts, newOwner)
}

// TransferOwner is a paid mutator transaction binding the contract method 0x4fb2e45d.
//
// Solidity: function transferOwner(newOwner address) returns()
func (_ExternalInvestor *ExternalInvestorTransactorSession) TransferOwner(newOwner common.Address) (*types.Transaction, error) {
	return _ExternalInvestor.Contract.TransferOwner(&_ExternalInvestor.TransactOpts, newOwner)
}

// ExternalInvestorAdminAddedIterator is returned from FilterAdminAdded and is used to iterate over the raw logs and unpacked data for AdminAdded events raised by the ExternalInvestor contract.
type ExternalInvestorAdminAddedIterator struct {
	Event *ExternalInvestorAdminAdded // Event containing the contract specifics and raw log

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
func (it *ExternalInvestorAdminAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExternalInvestorAdminAdded)
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
		it.Event = new(ExternalInvestorAdminAdded)
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
func (it *ExternalInvestorAdminAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExternalInvestorAdminAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExternalInvestorAdminAdded represents a AdminAdded event raised by the ExternalInvestor contract.
type ExternalInvestorAdminAdded struct {
	Admin common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterAdminAdded is a free log retrieval operation binding the contract event 0x44d6d25963f097ad14f29f06854a01f575648a1ef82f30e562ccd3889717e339.
//
// Solidity: e AdminAdded(admin indexed address)
func (_ExternalInvestor *ExternalInvestorFilterer) FilterAdminAdded(opts *bind.FilterOpts, admin []common.Address) (*ExternalInvestorAdminAddedIterator, error) {

	var adminRule []interface{}
	for _, adminItem := range admin {
		adminRule = append(adminRule, adminItem)
	}

	logs, sub, err := _ExternalInvestor.contract.FilterLogs(opts, "AdminAdded", adminRule)
	if err != nil {
		return nil, err
	}
	return &ExternalInvestorAdminAddedIterator{contract: _ExternalInvestor.contract, event: "AdminAdded", logs: logs, sub: sub}, nil
}

// WatchAdminAdded is a free log subscription operation binding the contract event 0x44d6d25963f097ad14f29f06854a01f575648a1ef82f30e562ccd3889717e339.
//
// Solidity: e AdminAdded(admin indexed address)
func (_ExternalInvestor *ExternalInvestorFilterer) WatchAdminAdded(opts *bind.WatchOpts, sink chan<- *ExternalInvestorAdminAdded, admin []common.Address) (event.Subscription, error) {

	var adminRule []interface{}
	for _, adminItem := range admin {
		adminRule = append(adminRule, adminItem)
	}

	logs, sub, err := _ExternalInvestor.contract.WatchLogs(opts, "AdminAdded", adminRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExternalInvestorAdminAdded)
				if err := _ExternalInvestor.contract.UnpackLog(event, "AdminAdded", log); err != nil {
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

// ExternalInvestorAdminRemovedIterator is returned from FilterAdminRemoved and is used to iterate over the raw logs and unpacked data for AdminRemoved events raised by the ExternalInvestor contract.
type ExternalInvestorAdminRemovedIterator struct {
	Event *ExternalInvestorAdminRemoved // Event containing the contract specifics and raw log

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
func (it *ExternalInvestorAdminRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExternalInvestorAdminRemoved)
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
		it.Event = new(ExternalInvestorAdminRemoved)
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
func (it *ExternalInvestorAdminRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExternalInvestorAdminRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExternalInvestorAdminRemoved represents a AdminRemoved event raised by the ExternalInvestor contract.
type ExternalInvestorAdminRemoved struct {
	Admin common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterAdminRemoved is a free log retrieval operation binding the contract event 0xa3b62bc36326052d97ea62d63c3d60308ed4c3ea8ac079dd8499f1e9c4f80c0f.
//
// Solidity: e AdminRemoved(admin indexed address)
func (_ExternalInvestor *ExternalInvestorFilterer) FilterAdminRemoved(opts *bind.FilterOpts, admin []common.Address) (*ExternalInvestorAdminRemovedIterator, error) {

	var adminRule []interface{}
	for _, adminItem := range admin {
		adminRule = append(adminRule, adminItem)
	}

	logs, sub, err := _ExternalInvestor.contract.FilterLogs(opts, "AdminRemoved", adminRule)
	if err != nil {
		return nil, err
	}
	return &ExternalInvestorAdminRemovedIterator{contract: _ExternalInvestor.contract, event: "AdminRemoved", logs: logs, sub: sub}, nil
}

// WatchAdminRemoved is a free log subscription operation binding the contract event 0xa3b62bc36326052d97ea62d63c3d60308ed4c3ea8ac079dd8499f1e9c4f80c0f.
//
// Solidity: e AdminRemoved(admin indexed address)
func (_ExternalInvestor *ExternalInvestorFilterer) WatchAdminRemoved(opts *bind.WatchOpts, sink chan<- *ExternalInvestorAdminRemoved, admin []common.Address) (event.Subscription, error) {

	var adminRule []interface{}
	for _, adminItem := range admin {
		adminRule = append(adminRule, adminItem)
	}

	logs, sub, err := _ExternalInvestor.contract.WatchLogs(opts, "AdminRemoved", adminRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExternalInvestorAdminRemoved)
				if err := _ExternalInvestor.contract.UnpackLog(event, "AdminRemoved", log); err != nil {
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

// ExternalInvestorInvestorAddedIterator is returned from FilterInvestorAdded and is used to iterate over the raw logs and unpacked data for InvestorAdded events raised by the ExternalInvestor contract.
type ExternalInvestorInvestorAddedIterator struct {
	Event *ExternalInvestorInvestorAdded // Event containing the contract specifics and raw log

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
func (it *ExternalInvestorInvestorAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExternalInvestorInvestorAdded)
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
		it.Event = new(ExternalInvestorInvestorAdded)
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
func (it *ExternalInvestorInvestorAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExternalInvestorInvestorAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExternalInvestorInvestorAdded represents a InvestorAdded event raised by the ExternalInvestor contract.
type ExternalInvestorInvestorAdded struct {
	Investor common.Address
	Owner    common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterInvestorAdded is a free log retrieval operation binding the contract event 0xe99183cc0b1657b54afa611991294ec1e4c458d7c36910518e2a5b76b2b6e73f.
//
// Solidity: e InvestorAdded(investor indexed address, owner indexed address)
func (_ExternalInvestor *ExternalInvestorFilterer) FilterInvestorAdded(opts *bind.FilterOpts, investor []common.Address, owner []common.Address) (*ExternalInvestorInvestorAddedIterator, error) {

	var investorRule []interface{}
	for _, investorItem := range investor {
		investorRule = append(investorRule, investorItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _ExternalInvestor.contract.FilterLogs(opts, "InvestorAdded", investorRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return &ExternalInvestorInvestorAddedIterator{contract: _ExternalInvestor.contract, event: "InvestorAdded", logs: logs, sub: sub}, nil
}

// WatchInvestorAdded is a free log subscription operation binding the contract event 0xe99183cc0b1657b54afa611991294ec1e4c458d7c36910518e2a5b76b2b6e73f.
//
// Solidity: e InvestorAdded(investor indexed address, owner indexed address)
func (_ExternalInvestor *ExternalInvestorFilterer) WatchInvestorAdded(opts *bind.WatchOpts, sink chan<- *ExternalInvestorInvestorAdded, investor []common.Address, owner []common.Address) (event.Subscription, error) {

	var investorRule []interface{}
	for _, investorItem := range investor {
		investorRule = append(investorRule, investorItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _ExternalInvestor.contract.WatchLogs(opts, "InvestorAdded", investorRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExternalInvestorInvestorAdded)
				if err := _ExternalInvestor.contract.UnpackLog(event, "InvestorAdded", log); err != nil {
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

// ExternalInvestorInvestorFrozenIterator is returned from FilterInvestorFrozen and is used to iterate over the raw logs and unpacked data for InvestorFrozen events raised by the ExternalInvestor contract.
type ExternalInvestorInvestorFrozenIterator struct {
	Event *ExternalInvestorInvestorFrozen // Event containing the contract specifics and raw log

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
func (it *ExternalInvestorInvestorFrozenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExternalInvestorInvestorFrozen)
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
		it.Event = new(ExternalInvestorInvestorFrozen)
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
func (it *ExternalInvestorInvestorFrozenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExternalInvestorInvestorFrozenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExternalInvestorInvestorFrozen represents a InvestorFrozen event raised by the ExternalInvestor contract.
type ExternalInvestorInvestorFrozen struct {
	Investor common.Address
	Frozen   bool
	Owner    common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterInvestorFrozen is a free log retrieval operation binding the contract event 0x6d30448ca28c66e149273ddd6d39fe9cb1982d4013649080aeb9d8251356d381.
//
// Solidity: e InvestorFrozen(investor indexed address, frozen indexed bool, owner indexed address)
func (_ExternalInvestor *ExternalInvestorFilterer) FilterInvestorFrozen(opts *bind.FilterOpts, investor []common.Address, frozen []bool, owner []common.Address) (*ExternalInvestorInvestorFrozenIterator, error) {

	var investorRule []interface{}
	for _, investorItem := range investor {
		investorRule = append(investorRule, investorItem)
	}
	var frozenRule []interface{}
	for _, frozenItem := range frozen {
		frozenRule = append(frozenRule, frozenItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _ExternalInvestor.contract.FilterLogs(opts, "InvestorFrozen", investorRule, frozenRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return &ExternalInvestorInvestorFrozenIterator{contract: _ExternalInvestor.contract, event: "InvestorFrozen", logs: logs, sub: sub}, nil
}

// WatchInvestorFrozen is a free log subscription operation binding the contract event 0x6d30448ca28c66e149273ddd6d39fe9cb1982d4013649080aeb9d8251356d381.
//
// Solidity: e InvestorFrozen(investor indexed address, frozen indexed bool, owner indexed address)
func (_ExternalInvestor *ExternalInvestorFilterer) WatchInvestorFrozen(opts *bind.WatchOpts, sink chan<- *ExternalInvestorInvestorFrozen, investor []common.Address, frozen []bool, owner []common.Address) (event.Subscription, error) {

	var investorRule []interface{}
	for _, investorItem := range investor {
		investorRule = append(investorRule, investorItem)
	}
	var frozenRule []interface{}
	for _, frozenItem := range frozen {
		frozenRule = append(frozenRule, frozenItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _ExternalInvestor.contract.WatchLogs(opts, "InvestorFrozen", investorRule, frozenRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExternalInvestorInvestorFrozen)
				if err := _ExternalInvestor.contract.UnpackLog(event, "InvestorFrozen", log); err != nil {
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

// ExternalInvestorInvestorRemovedIterator is returned from FilterInvestorRemoved and is used to iterate over the raw logs and unpacked data for InvestorRemoved events raised by the ExternalInvestor contract.
type ExternalInvestorInvestorRemovedIterator struct {
	Event *ExternalInvestorInvestorRemoved // Event containing the contract specifics and raw log

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
func (it *ExternalInvestorInvestorRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExternalInvestorInvestorRemoved)
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
		it.Event = new(ExternalInvestorInvestorRemoved)
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
func (it *ExternalInvestorInvestorRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExternalInvestorInvestorRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExternalInvestorInvestorRemoved represents a InvestorRemoved event raised by the ExternalInvestor contract.
type ExternalInvestorInvestorRemoved struct {
	Investor common.Address
	Owner    common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterInvestorRemoved is a free log retrieval operation binding the contract event 0xd8755221287ca1f6b28807977a086f5534d9e02ea27ebad003d7cb1a95659a46.
//
// Solidity: e InvestorRemoved(investor indexed address, owner indexed address)
func (_ExternalInvestor *ExternalInvestorFilterer) FilterInvestorRemoved(opts *bind.FilterOpts, investor []common.Address, owner []common.Address) (*ExternalInvestorInvestorRemovedIterator, error) {

	var investorRule []interface{}
	for _, investorItem := range investor {
		investorRule = append(investorRule, investorItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _ExternalInvestor.contract.FilterLogs(opts, "InvestorRemoved", investorRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return &ExternalInvestorInvestorRemovedIterator{contract: _ExternalInvestor.contract, event: "InvestorRemoved", logs: logs, sub: sub}, nil
}

// WatchInvestorRemoved is a free log subscription operation binding the contract event 0xd8755221287ca1f6b28807977a086f5534d9e02ea27ebad003d7cb1a95659a46.
//
// Solidity: e InvestorRemoved(investor indexed address, owner indexed address)
func (_ExternalInvestor *ExternalInvestorFilterer) WatchInvestorRemoved(opts *bind.WatchOpts, sink chan<- *ExternalInvestorInvestorRemoved, investor []common.Address, owner []common.Address) (event.Subscription, error) {

	var investorRule []interface{}
	for _, investorItem := range investor {
		investorRule = append(investorRule, investorItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _ExternalInvestor.contract.WatchLogs(opts, "InvestorRemoved", investorRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExternalInvestorInvestorRemoved)
				if err := _ExternalInvestor.contract.UnpackLog(event, "InvestorRemoved", log); err != nil {
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

// ExternalInvestorOwnerTransferredIterator is returned from FilterOwnerTransferred and is used to iterate over the raw logs and unpacked data for OwnerTransferred events raised by the ExternalInvestor contract.
type ExternalInvestorOwnerTransferredIterator struct {
	Event *ExternalInvestorOwnerTransferred // Event containing the contract specifics and raw log

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
func (it *ExternalInvestorOwnerTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExternalInvestorOwnerTransferred)
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
		it.Event = new(ExternalInvestorOwnerTransferred)
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
func (it *ExternalInvestorOwnerTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExternalInvestorOwnerTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExternalInvestorOwnerTransferred represents a OwnerTransferred event raised by the ExternalInvestor contract.
type ExternalInvestorOwnerTransferred struct {
	OldOwner common.Address
	NewOwner common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOwnerTransferred is a free log retrieval operation binding the contract event 0x8934ce4adea8d9ce0d714d2c22b86790e41b7731c84b926fbbdc1d40ff6533c9.
//
// Solidity: e OwnerTransferred(oldOwner indexed address, newOwner indexed address)
func (_ExternalInvestor *ExternalInvestorFilterer) FilterOwnerTransferred(opts *bind.FilterOpts, oldOwner []common.Address, newOwner []common.Address) (*ExternalInvestorOwnerTransferredIterator, error) {

	var oldOwnerRule []interface{}
	for _, oldOwnerItem := range oldOwner {
		oldOwnerRule = append(oldOwnerRule, oldOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ExternalInvestor.contract.FilterLogs(opts, "OwnerTransferred", oldOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ExternalInvestorOwnerTransferredIterator{contract: _ExternalInvestor.contract, event: "OwnerTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnerTransferred is a free log subscription operation binding the contract event 0x8934ce4adea8d9ce0d714d2c22b86790e41b7731c84b926fbbdc1d40ff6533c9.
//
// Solidity: e OwnerTransferred(oldOwner indexed address, newOwner indexed address)
func (_ExternalInvestor *ExternalInvestorFilterer) WatchOwnerTransferred(opts *bind.WatchOpts, sink chan<- *ExternalInvestorOwnerTransferred, oldOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var oldOwnerRule []interface{}
	for _, oldOwnerItem := range oldOwner {
		oldOwnerRule = append(oldOwnerRule, oldOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ExternalInvestor.contract.WatchLogs(opts, "OwnerTransferred", oldOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExternalInvestorOwnerTransferred)
				if err := _ExternalInvestor.contract.UnpackLog(event, "OwnerTransferred", log); err != nil {
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
