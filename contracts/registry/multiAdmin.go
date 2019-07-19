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

// MultiAdminABI is the input ABI used to generate the binding from.
const MultiAdminABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"creator\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"electionResults\",\"outputs\":[{\"name\":\"\",\"type\":\"uint8\"},{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"uint8\"},{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"voter\",\"type\":\"address\"}],\"name\":\"hasVoted\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"incumbent\",\"type\":\"address\"},{\"name\":\"challenger\",\"type\":\"address\"},{\"name\":\"yea\",\"type\":\"bool\"}],\"name\":\"vote\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"salt\",\"type\":\"uint256\"}],\"name\":\"deterministicAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"isAdmin\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"index\",\"type\":\"int256\"}],\"name\":\"adminAt\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"registryAddress\",\"type\":\"address\"},{\"name\":\"adminAddresses\",\"type\":\"address[]\"}],\"name\":\"init\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"voteThreshold\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"incumbent\",\"type\":\"address\"},{\"name\":\"challenger\",\"type\":\"address\"}],\"name\":\"createElection\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"registry\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"admins\",\"outputs\":[{\"name\":\"count\",\"type\":\"int256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"r\",\"type\":\"address\"}],\"name\":\"setRegistry\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"salt\",\"type\":\"uint256\"},{\"name\":\"addresses\",\"type\":\"address[]\"}],\"name\":\"clone\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ballotDuration\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"salt\",\"type\":\"uint256\"}],\"name\":\"clone\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"fallback\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"incumbent\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"challenger\",\"type\":\"address\"}],\"name\":\"ElectionCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"incumbent\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"challenger\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"passed\",\"type\":\"bool\"},{\"indexed\":false,\"name\":\"margin\",\"type\":\"uint256\"}],\"name\":\"ElectionClosed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"incumbent\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"challenger\",\"type\":\"address\"}],\"name\":\"VoteCast\",\"type\":\"event\"}]"

// MultiAdminBin is the compiled bytecode used for deploying new contracts.
const MultiAdminBin = `608060405234801561001057600080fd5b50600780546001600160a01b0319163317905561157b806100326000396000f3fe6080604052600436106100f35760003560e01c80634fe437d51161008a578063a91ee0dc11610059578063a91ee0dc1461046a578063a9fd834a1461049d578063ab8f953614610547578063cdc19adf1461055c576100f3565b80634fe437d5146103de578063727dba5d146104055780637b10399914610440578063a5de361914610455576100f3565b806313590940116100c657806313590940146102ca57806324d7806c146102f45780632bdbc56f146103275780633c5a3cea14610351576100f3565b806302d05d3f146101b857806308670f4f146101e957806309eef43e1461023c578063128d871114610283575b3480156100ff57600080fd5b5061011160043363ffffffff61057916565b610153576040805162461bcd60e51b815260206004820152600e60248201526d10591b5a5b881c995c5d5a5c995960921b604482015290519081900360640190fd5b60606000368080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201829052506008549495509384935036925060208601915083906127105a03f16040513d6000823e816101b4573d81fd5b3d81f35b3480156101c457600080fd5b506101cd6105b2565b604080516001600160a01b039092168252519081900360200190f35b3480156101f557600080fd5b506101fe6105c1565b6040805160ff96871681526001600160a01b039586166020820152939095168386015292166060820152608081019190915290519081900360a00190f35b34801561024857600080fd5b5061026f6004803603602081101561025f57600080fd5b50356001600160a01b0316610670565b604080519115158252519081900360200190f35b34801561028f57600080fd5b506102c8600480360360608110156102a657600080fd5b506001600160a01b0381358116916020810135909116906040013515156106b9565b005b3480156102d657600080fd5b506101cd600480360360208110156102ed57600080fd5b5035610a95565b34801561030057600080fd5b5061026f6004803603602081101561031757600080fd5b50356001600160a01b0316610b03565b34801561033357600080fd5b506101cd6004803603602081101561034a57600080fd5b5035610b16565b34801561035d57600080fd5b506102c86004803603604081101561037457600080fd5b6001600160a01b03823516919081019060408101602082013564010000000081111561039f57600080fd5b8201836020820111156103b157600080fd5b803590602001918460208302840111640100000000831117156103d357600080fd5b509092509050610b29565b3480156103ea57600080fd5b506103f3610c0c565b60408051918252519081900360200190f35b34801561041157600080fd5b506102c86004803603604081101561042857600080fd5b506001600160a01b0381358116916020013516610c12565b34801561044c57600080fd5b506101cd610e24565b34801561046157600080fd5b506103f3610e33565b34801561047657600080fd5b506102c86004803603602081101561048d57600080fd5b50356001600160a01b0316610e39565b6101cd600480360360408110156104b357600080fd5b813591908101906040810160208201356401000000008111156104d557600080fd5b8201836020820111156104e757600080fd5b8035906020019184602083028401116401000000008311171561050957600080fd5b919080806020026020016040519081016040528093929190818152602001838360200280828437600092019190915250929550610f09945050505050565b34801561055357600080fd5b506103f3611039565b6101cd6004803603602081101561057257600080fd5b503561103f565b6001600160a01b0381166000908152600183016020526040812054600019018181128015906105a85750835481125b9150505b92915050565b6007546001600160a01b031681565b6000805481908190819081906301000000900463ffffffff16428190039015806105ec575060025481115b15610607575060009450849350839250829150819050610669565b6201518081600254038161061757fe5b600054919004915060ff6101008204811691610642916004916201000090041663ffffffff6110d116565b600054919750955060ff8116945067010000000000000090046001600160a01b0316925090505b9091929394565b60008061068460048463ffffffff61115716565b905060008112156106995760009150506106b4565b600154819060008212156106a957fe5b6001911c8116149150505b919050565b6106ca60043363ffffffff61057916565b61070c576040805162461bcd60e51b815260206004820152600e60248201526d10591b5a5b881c995c5d5a5c995960921b604482015290519081900360640190fd5b6002546000546301000000900463ffffffff1642031115610774576040805162461bcd60e51b815260206004820152600960248201527f4e6f2062616c6c6f740000000000000000000000000000000000000000000000604482015290519081900360640190fd5b6000546001600160a01b03838116670100000000000000909204161480156107c557506000546001600160a01b038416906107ba9060049062010000900460ff166110d1565b6001600160a01b0316145b610816576040805162461bcd60e51b815260206004820152601260248201527f496e76616c69642063616e646964617465730000000000000000000000000000604482015290519081900360640190fd5b600061082960043363ffffffff61115716565b6001805491925090821c1615610886576040805162461bcd60e51b815260206004820152600d60248201527f416c726561647920766f74656400000000000000000000000000000000000000604482015290519081900360640190fd5b81156108a7576000805460ff8082166001011660ff199091161790556108c7565b6000805460ff61010080830482166001019091160261ff00199091161790555b6001805481831b1790556040516001600160a01b0380851691908616907f113ed09b18bb8fe2635af16f21d67464502d920ff1015abfe08dbe9a2971a7e290600090a360035460005460ff1610158061092c5750600354600054610100900460ff1610155b8061094c575060045460005460ff80821661010090920481169190910116145b15610a8f57600080546040805160c08101825283815260208101849052908101839052606081018390526080810183905260a0018290527fffffffffff00000000000000000000000000000000000000000000000000000081168255600182905560ff6101008204811691160390811315610a39576109d260048663ffffffff6111ae16565b506109e460048563ffffffff61129916565b5060011515846001600160a01b0316866001600160a01b03167f81a623b5156020c64c45d3fb56c628333f904077134ff40c1983fd1b23de573e846040518082815260200191505060405180910390a4610a8d565b60001515846001600160a01b0316866001600160a01b03167f81a623b5156020c64c45d3fb56c628333f904077134ff40c1983fd1b23de573e84196001016040518082815260200191505060405180910390a45b505b50505050565b60003060581b7fff0000000000000000000000000000000000000000000000000000000000000017600052816015523060481b7f5880730000000000000000000000000000000000000000803b80938091923cf3176035526020603520603552605560002060005260206000f35b60006105ac60048363ffffffff61057916565b60006105ac60048363ffffffff6110d116565b6007546001600160a01b0316331480610b4b57506007546001600160a01b0316155b610b9c576040805162461bcd60e51b815260206004820152601b60248201527f496e766f636174696f6e2072657175697265732063726561746f720000000000604482015290519081900360640190fd5b610be882828080602002602001604051908101604052809392919081815260200183836020028082843760009201919091525060029250859150610bdd9050565b04600101600f61133e565b5050600880546001600160a01b0319166001600160a01b0392909216919091179055565b60035481565b610c2360043363ffffffff61057916565b610c65576040805162461bcd60e51b815260206004820152600e60248201526d10591b5a5b881c995c5d5a5c995960921b604482015290519081900360640190fd5b6002546000546301000000900463ffffffff16420311610ccc576040805162461bcd60e51b815260206004820152601360248201527f456c656374696f6e20696e2070726f6365737300000000000000000000000000604482015290519081900360640190fd5b6000610cdf60048463ffffffff61115716565b905060008112158015610d005750610cfe60048363ffffffff61057916565b155b610d51576040805162461bcd60e51b815260206004820152601260248201527f496e76616c69642063616e646964617465730000000000000000000000000000604482015290519081900360640190fd5b6040805160c08101825260008082526020820181905260ff84168284018190524263ffffffff16606084018190526001600160a01b038781166080860181905260a0909501849052835462ffffff1916620100009093029290921766ffffffff00000019166301000000909102177fffffffffff0000000000000000000000000000000000000000ffffffffffffff1667010000000000000084021782556001829055925191928616917f2931cee8058f0a9299aff15bc55251858ff1a2c223d96c6d55ebe9c67ee0351b9190a3505050565b6008546001600160a01b031681565b60045481565b610e4a60043363ffffffff61057916565b610e8c576040805162461bcd60e51b815260206004820152600e60248201526d10591b5a5b881c995c5d5a5c995960921b604482015290519081900360640190fd5b6001600160a01b038116610ee7576040805162461bcd60e51b815260206004820152601660248201527f56616c6964206164647265737320726571756972656400000000000000000000604482015290519081900360640190fd5b600880546001600160a01b0319166001600160a01b0392909216919091179055565b6000610f1c60043363ffffffff61057916565b610f5e576040805162461bcd60e51b815260206004820152600e60248201526d10591b5a5b881c995c5d5a5c995960921b604482015290519081900360640190fd5b6000610f6984611509565b600854604080517f3c5a3cea0000000000000000000000000000000000000000000000000000000081526001600160a01b03928316600482018181526024830193845288516044840152885195965093861694633c5a3cea949193899390916064909101906020808601910280838360005b83811015610ff3578181015183820152602001610fdb565b505050509050019350505050600060405180830381600087803b15801561101957600080fd5b505af115801561102d573d6000803e3d6000fd5b50929695505050505050565b60025481565b60006060600460000154604051908082528060200260200182016040528015611072578160200160208202803883390190505b50905060005b6004548112156110bf5761109360048263ffffffff6110d116565b82828151811061109f57fe5b6001600160a01b0390921660209283029190910190910152600101611078565b506110ca8382610f09565b9392505050565b60008082121580156110e35750825482125b611134576040805162461bcd60e51b815260206004820152601860248201527f496e646578206f757473696465206f6620626f756e64732e0000000000000000604482015290519081900360640190fd5b50600101600090815260029190910160205260409020546001600160a01b031690565b60006001600160a01b03821661117057506000196105ac565b6001600160a01b0382166000908152600184016020526040812054600019019081128061119e575083548112155b156110ca576000199150506105ac565b6001600160a01b038116600090815260018084016020526040822054908112806111d85750835481135b156111e75760009150506105ac565b835481121561124e57835460009081526002850160208181526040808420546001600160a01b03168085526001890183528185208690558585529290915280832080546001600160a01b03199081169093179055865483529091208054909116905561126d565b6000818152600285016020526040902080546001600160a01b03191690555b50506001600160a01b031660009081526001828101602052604082209190915581546000190190915590565b60006001600160a01b0382166112b1575060006105ac565b6001600160a01b0382166000908152600184016020526040812054600019019081128015906112e05750835481125b156112ef5760009150506105ac565b5050815460019081018084556001600160a01b0383166000818152838601602090815260408083208590559382526002870190529190912080546001600160a01b031916909117905592915050565b60045415611393576040805162461bcd60e51b815260206004820152601360248201527f416c726561647920696e697469616c697a656400000000000000000000000000604482015290519081900360640190fd5b600183511180156113a657506101008351105b6113f7576040805162461bcd60e51b815260206004820152601860248201527f41646d696e73206f757473696465206f6620626f756e64730000000000000000604482015290519081900360640190fd5b600182118015611408575082518211155b611459576040805162461bcd60e51b815260206004820152601160248201527f496e76616c6964207468726573686f6c64000000000000000000000000000000604482015290519081900360640190fd5b60008111801561146a575061016e81105b6114bb576040805162461bcd60e51b815260206004820152601060248201527f496e76616c6964206475726174696f6e00000000000000000000000000000000604482015290519081900360640190fd5b60005b83518110156114f7576114ee8482815181106114d657fe5b6020026020010151600461129990919063ffffffff16565b506001016114be565b50600391909155620151800260025550565b6000803060481b7f5880730000000000000000000000000000000000000000803b80938091923cf317600052826020600034f59050803b6105ac57fefea265627a7a723058205e85d26b4d484d0d8aa47fbab50222fccf529c66d2639e4441e7a7e6a06d99e564736f6c63430005090032`

// DeployMultiAdmin deploys a new Ethereum contract, binding an instance of MultiAdmin to it.
func DeployMultiAdmin(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *MultiAdmin, error) {
	parsed, err := abi.JSON(strings.NewReader(MultiAdminABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(MultiAdminBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &MultiAdmin{MultiAdminCaller: MultiAdminCaller{contract: contract}, MultiAdminTransactor: MultiAdminTransactor{contract: contract}, MultiAdminFilterer: MultiAdminFilterer{contract: contract}}, nil
}

// MultiAdmin is an auto generated Go binding around an Ethereum contract.
type MultiAdmin struct {
	MultiAdminCaller     // Read-only binding to the contract
	MultiAdminTransactor // Write-only binding to the contract
	MultiAdminFilterer   // Log filterer for contract events
}

// MultiAdminCaller is an auto generated read-only Go binding around an Ethereum contract.
type MultiAdminCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MultiAdminTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MultiAdminTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MultiAdminFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MultiAdminFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MultiAdminSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MultiAdminSession struct {
	Contract     *MultiAdmin       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MultiAdminCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MultiAdminCallerSession struct {
	Contract *MultiAdminCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// MultiAdminTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MultiAdminTransactorSession struct {
	Contract     *MultiAdminTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// MultiAdminRaw is an auto generated low-level Go binding around an Ethereum contract.
type MultiAdminRaw struct {
	Contract *MultiAdmin // Generic contract binding to access the raw methods on
}

// MultiAdminCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MultiAdminCallerRaw struct {
	Contract *MultiAdminCaller // Generic read-only contract binding to access the raw methods on
}

// MultiAdminTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MultiAdminTransactorRaw struct {
	Contract *MultiAdminTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMultiAdmin creates a new instance of MultiAdmin, bound to a specific deployed contract.
func NewMultiAdmin(address common.Address, backend bind.ContractBackend) (*MultiAdmin, error) {
	contract, err := bindMultiAdmin(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MultiAdmin{MultiAdminCaller: MultiAdminCaller{contract: contract}, MultiAdminTransactor: MultiAdminTransactor{contract: contract}, MultiAdminFilterer: MultiAdminFilterer{contract: contract}}, nil
}

// NewMultiAdminCaller creates a new read-only instance of MultiAdmin, bound to a specific deployed contract.
func NewMultiAdminCaller(address common.Address, caller bind.ContractCaller) (*MultiAdminCaller, error) {
	contract, err := bindMultiAdmin(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MultiAdminCaller{contract: contract}, nil
}

// NewMultiAdminTransactor creates a new write-only instance of MultiAdmin, bound to a specific deployed contract.
func NewMultiAdminTransactor(address common.Address, transactor bind.ContractTransactor) (*MultiAdminTransactor, error) {
	contract, err := bindMultiAdmin(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MultiAdminTransactor{contract: contract}, nil
}

// NewMultiAdminFilterer creates a new log filterer instance of MultiAdmin, bound to a specific deployed contract.
func NewMultiAdminFilterer(address common.Address, filterer bind.ContractFilterer) (*MultiAdminFilterer, error) {
	contract, err := bindMultiAdmin(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MultiAdminFilterer{contract: contract}, nil
}

// bindMultiAdmin binds a generic wrapper to an already deployed contract.
func bindMultiAdmin(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MultiAdminABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MultiAdmin *MultiAdminRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _MultiAdmin.Contract.MultiAdminCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MultiAdmin *MultiAdminRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MultiAdmin.Contract.MultiAdminTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MultiAdmin *MultiAdminRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MultiAdmin.Contract.MultiAdminTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MultiAdmin *MultiAdminCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _MultiAdmin.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MultiAdmin *MultiAdminTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MultiAdmin.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MultiAdmin *MultiAdminTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MultiAdmin.Contract.contract.Transact(opts, method, params...)
}

// AdminAt is a free data retrieval call binding the contract method 0x2bdbc56f.
//
// Solidity: function adminAt(index int256) constant returns(address)
func (_MultiAdmin *MultiAdminCaller) AdminAt(opts *bind.CallOpts, index *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _MultiAdmin.contract.Call(opts, out, "adminAt", index)
	return *ret0, err
}

// AdminAt is a free data retrieval call binding the contract method 0x2bdbc56f.
//
// Solidity: function adminAt(index int256) constant returns(address)
func (_MultiAdmin *MultiAdminSession) AdminAt(index *big.Int) (common.Address, error) {
	return _MultiAdmin.Contract.AdminAt(&_MultiAdmin.CallOpts, index)
}

// AdminAt is a free data retrieval call binding the contract method 0x2bdbc56f.
//
// Solidity: function adminAt(index int256) constant returns(address)
func (_MultiAdmin *MultiAdminCallerSession) AdminAt(index *big.Int) (common.Address, error) {
	return _MultiAdmin.Contract.AdminAt(&_MultiAdmin.CallOpts, index)
}

// Admins is a free data retrieval call binding the contract method 0xa5de3619.
//
// Solidity: function admins() constant returns(count int256)
func (_MultiAdmin *MultiAdminCaller) Admins(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _MultiAdmin.contract.Call(opts, out, "admins")
	return *ret0, err
}

// Admins is a free data retrieval call binding the contract method 0xa5de3619.
//
// Solidity: function admins() constant returns(count int256)
func (_MultiAdmin *MultiAdminSession) Admins() (*big.Int, error) {
	return _MultiAdmin.Contract.Admins(&_MultiAdmin.CallOpts)
}

// Admins is a free data retrieval call binding the contract method 0xa5de3619.
//
// Solidity: function admins() constant returns(count int256)
func (_MultiAdmin *MultiAdminCallerSession) Admins() (*big.Int, error) {
	return _MultiAdmin.Contract.Admins(&_MultiAdmin.CallOpts)
}

// BallotDuration is a free data retrieval call binding the contract method 0xab8f9536.
//
// Solidity: function ballotDuration() constant returns(uint256)
func (_MultiAdmin *MultiAdminCaller) BallotDuration(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _MultiAdmin.contract.Call(opts, out, "ballotDuration")
	return *ret0, err
}

// BallotDuration is a free data retrieval call binding the contract method 0xab8f9536.
//
// Solidity: function ballotDuration() constant returns(uint256)
func (_MultiAdmin *MultiAdminSession) BallotDuration() (*big.Int, error) {
	return _MultiAdmin.Contract.BallotDuration(&_MultiAdmin.CallOpts)
}

// BallotDuration is a free data retrieval call binding the contract method 0xab8f9536.
//
// Solidity: function ballotDuration() constant returns(uint256)
func (_MultiAdmin *MultiAdminCallerSession) BallotDuration() (*big.Int, error) {
	return _MultiAdmin.Contract.BallotDuration(&_MultiAdmin.CallOpts)
}

// Creator is a free data retrieval call binding the contract method 0x02d05d3f.
//
// Solidity: function creator() constant returns(address)
func (_MultiAdmin *MultiAdminCaller) Creator(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _MultiAdmin.contract.Call(opts, out, "creator")
	return *ret0, err
}

// Creator is a free data retrieval call binding the contract method 0x02d05d3f.
//
// Solidity: function creator() constant returns(address)
func (_MultiAdmin *MultiAdminSession) Creator() (common.Address, error) {
	return _MultiAdmin.Contract.Creator(&_MultiAdmin.CallOpts)
}

// Creator is a free data retrieval call binding the contract method 0x02d05d3f.
//
// Solidity: function creator() constant returns(address)
func (_MultiAdmin *MultiAdminCallerSession) Creator() (common.Address, error) {
	return _MultiAdmin.Contract.Creator(&_MultiAdmin.CallOpts)
}

// DeterministicAddress is a free data retrieval call binding the contract method 0x13590940.
//
// Solidity: function deterministicAddress(salt uint256) constant returns(address)
func (_MultiAdmin *MultiAdminCaller) DeterministicAddress(opts *bind.CallOpts, salt *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _MultiAdmin.contract.Call(opts, out, "deterministicAddress", salt)
	return *ret0, err
}

// DeterministicAddress is a free data retrieval call binding the contract method 0x13590940.
//
// Solidity: function deterministicAddress(salt uint256) constant returns(address)
func (_MultiAdmin *MultiAdminSession) DeterministicAddress(salt *big.Int) (common.Address, error) {
	return _MultiAdmin.Contract.DeterministicAddress(&_MultiAdmin.CallOpts, salt)
}

// DeterministicAddress is a free data retrieval call binding the contract method 0x13590940.
//
// Solidity: function deterministicAddress(salt uint256) constant returns(address)
func (_MultiAdmin *MultiAdminCallerSession) DeterministicAddress(salt *big.Int) (common.Address, error) {
	return _MultiAdmin.Contract.DeterministicAddress(&_MultiAdmin.CallOpts, salt)
}

// ElectionResults is a free data retrieval call binding the contract method 0x08670f4f.
//
// Solidity: function electionResults() constant returns(uint8, address, uint8, address, uint256)
func (_MultiAdmin *MultiAdminCaller) ElectionResults(opts *bind.CallOpts) (uint8, common.Address, uint8, common.Address, *big.Int, error) {
	var (
		ret0 = new(uint8)
		ret1 = new(common.Address)
		ret2 = new(uint8)
		ret3 = new(common.Address)
		ret4 = new(*big.Int)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
		ret3,
		ret4,
	}
	err := _MultiAdmin.contract.Call(opts, out, "electionResults")
	return *ret0, *ret1, *ret2, *ret3, *ret4, err
}

// ElectionResults is a free data retrieval call binding the contract method 0x08670f4f.
//
// Solidity: function electionResults() constant returns(uint8, address, uint8, address, uint256)
func (_MultiAdmin *MultiAdminSession) ElectionResults() (uint8, common.Address, uint8, common.Address, *big.Int, error) {
	return _MultiAdmin.Contract.ElectionResults(&_MultiAdmin.CallOpts)
}

// ElectionResults is a free data retrieval call binding the contract method 0x08670f4f.
//
// Solidity: function electionResults() constant returns(uint8, address, uint8, address, uint256)
func (_MultiAdmin *MultiAdminCallerSession) ElectionResults() (uint8, common.Address, uint8, common.Address, *big.Int, error) {
	return _MultiAdmin.Contract.ElectionResults(&_MultiAdmin.CallOpts)
}

// HasVoted is a free data retrieval call binding the contract method 0x09eef43e.
//
// Solidity: function hasVoted(voter address) constant returns(bool)
func (_MultiAdmin *MultiAdminCaller) HasVoted(opts *bind.CallOpts, voter common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _MultiAdmin.contract.Call(opts, out, "hasVoted", voter)
	return *ret0, err
}

// HasVoted is a free data retrieval call binding the contract method 0x09eef43e.
//
// Solidity: function hasVoted(voter address) constant returns(bool)
func (_MultiAdmin *MultiAdminSession) HasVoted(voter common.Address) (bool, error) {
	return _MultiAdmin.Contract.HasVoted(&_MultiAdmin.CallOpts, voter)
}

// HasVoted is a free data retrieval call binding the contract method 0x09eef43e.
//
// Solidity: function hasVoted(voter address) constant returns(bool)
func (_MultiAdmin *MultiAdminCallerSession) HasVoted(voter common.Address) (bool, error) {
	return _MultiAdmin.Contract.HasVoted(&_MultiAdmin.CallOpts, voter)
}

// IsAdmin is a free data retrieval call binding the contract method 0x24d7806c.
//
// Solidity: function isAdmin(addr address) constant returns(bool)
func (_MultiAdmin *MultiAdminCaller) IsAdmin(opts *bind.CallOpts, addr common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _MultiAdmin.contract.Call(opts, out, "isAdmin", addr)
	return *ret0, err
}

// IsAdmin is a free data retrieval call binding the contract method 0x24d7806c.
//
// Solidity: function isAdmin(addr address) constant returns(bool)
func (_MultiAdmin *MultiAdminSession) IsAdmin(addr common.Address) (bool, error) {
	return _MultiAdmin.Contract.IsAdmin(&_MultiAdmin.CallOpts, addr)
}

// IsAdmin is a free data retrieval call binding the contract method 0x24d7806c.
//
// Solidity: function isAdmin(addr address) constant returns(bool)
func (_MultiAdmin *MultiAdminCallerSession) IsAdmin(addr common.Address) (bool, error) {
	return _MultiAdmin.Contract.IsAdmin(&_MultiAdmin.CallOpts, addr)
}

// Registry is a free data retrieval call binding the contract method 0x7b103999.
//
// Solidity: function registry() constant returns(address)
func (_MultiAdmin *MultiAdminCaller) Registry(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _MultiAdmin.contract.Call(opts, out, "registry")
	return *ret0, err
}

// Registry is a free data retrieval call binding the contract method 0x7b103999.
//
// Solidity: function registry() constant returns(address)
func (_MultiAdmin *MultiAdminSession) Registry() (common.Address, error) {
	return _MultiAdmin.Contract.Registry(&_MultiAdmin.CallOpts)
}

// Registry is a free data retrieval call binding the contract method 0x7b103999.
//
// Solidity: function registry() constant returns(address)
func (_MultiAdmin *MultiAdminCallerSession) Registry() (common.Address, error) {
	return _MultiAdmin.Contract.Registry(&_MultiAdmin.CallOpts)
}

// VoteThreshold is a free data retrieval call binding the contract method 0x4fe437d5.
//
// Solidity: function voteThreshold() constant returns(uint256)
func (_MultiAdmin *MultiAdminCaller) VoteThreshold(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _MultiAdmin.contract.Call(opts, out, "voteThreshold")
	return *ret0, err
}

// VoteThreshold is a free data retrieval call binding the contract method 0x4fe437d5.
//
// Solidity: function voteThreshold() constant returns(uint256)
func (_MultiAdmin *MultiAdminSession) VoteThreshold() (*big.Int, error) {
	return _MultiAdmin.Contract.VoteThreshold(&_MultiAdmin.CallOpts)
}

// VoteThreshold is a free data retrieval call binding the contract method 0x4fe437d5.
//
// Solidity: function voteThreshold() constant returns(uint256)
func (_MultiAdmin *MultiAdminCallerSession) VoteThreshold() (*big.Int, error) {
	return _MultiAdmin.Contract.VoteThreshold(&_MultiAdmin.CallOpts)
}

// Clone is a paid mutator transaction binding the contract method 0xcdc19adf.
//
// Solidity: function clone(salt uint256) returns(address)
func (_MultiAdmin *MultiAdminTransactor) Clone(opts *bind.TransactOpts, salt *big.Int) (*types.Transaction, error) {
	return _MultiAdmin.contract.Transact(opts, "clone", salt)
}

// Clone is a paid mutator transaction binding the contract method 0xcdc19adf.
//
// Solidity: function clone(salt uint256) returns(address)
func (_MultiAdmin *MultiAdminSession) Clone(salt *big.Int) (*types.Transaction, error) {
	return _MultiAdmin.Contract.Clone(&_MultiAdmin.TransactOpts, salt)
}

// Clone is a paid mutator transaction binding the contract method 0xcdc19adf.
//
// Solidity: function clone(salt uint256) returns(address)
func (_MultiAdmin *MultiAdminTransactorSession) Clone(salt *big.Int) (*types.Transaction, error) {
	return _MultiAdmin.Contract.Clone(&_MultiAdmin.TransactOpts, salt)
}

// CreateElection is a paid mutator transaction binding the contract method 0x727dba5d.
//
// Solidity: function createElection(incumbent address, challenger address) returns()
func (_MultiAdmin *MultiAdminTransactor) CreateElection(opts *bind.TransactOpts, incumbent common.Address, challenger common.Address) (*types.Transaction, error) {
	return _MultiAdmin.contract.Transact(opts, "createElection", incumbent, challenger)
}

// CreateElection is a paid mutator transaction binding the contract method 0x727dba5d.
//
// Solidity: function createElection(incumbent address, challenger address) returns()
func (_MultiAdmin *MultiAdminSession) CreateElection(incumbent common.Address, challenger common.Address) (*types.Transaction, error) {
	return _MultiAdmin.Contract.CreateElection(&_MultiAdmin.TransactOpts, incumbent, challenger)
}

// CreateElection is a paid mutator transaction binding the contract method 0x727dba5d.
//
// Solidity: function createElection(incumbent address, challenger address) returns()
func (_MultiAdmin *MultiAdminTransactorSession) CreateElection(incumbent common.Address, challenger common.Address) (*types.Transaction, error) {
	return _MultiAdmin.Contract.CreateElection(&_MultiAdmin.TransactOpts, incumbent, challenger)
}

// Init is a paid mutator transaction binding the contract method 0x3c5a3cea.
//
// Solidity: function init(registryAddress address, adminAddresses address[]) returns()
func (_MultiAdmin *MultiAdminTransactor) Init(opts *bind.TransactOpts, registryAddress common.Address, adminAddresses []common.Address) (*types.Transaction, error) {
	return _MultiAdmin.contract.Transact(opts, "init", registryAddress, adminAddresses)
}

// Init is a paid mutator transaction binding the contract method 0x3c5a3cea.
//
// Solidity: function init(registryAddress address, adminAddresses address[]) returns()
func (_MultiAdmin *MultiAdminSession) Init(registryAddress common.Address, adminAddresses []common.Address) (*types.Transaction, error) {
	return _MultiAdmin.Contract.Init(&_MultiAdmin.TransactOpts, registryAddress, adminAddresses)
}

// Init is a paid mutator transaction binding the contract method 0x3c5a3cea.
//
// Solidity: function init(registryAddress address, adminAddresses address[]) returns()
func (_MultiAdmin *MultiAdminTransactorSession) Init(registryAddress common.Address, adminAddresses []common.Address) (*types.Transaction, error) {
	return _MultiAdmin.Contract.Init(&_MultiAdmin.TransactOpts, registryAddress, adminAddresses)
}

// SetRegistry is a paid mutator transaction binding the contract method 0xa91ee0dc.
//
// Solidity: function setRegistry(r address) returns()
func (_MultiAdmin *MultiAdminTransactor) SetRegistry(opts *bind.TransactOpts, r common.Address) (*types.Transaction, error) {
	return _MultiAdmin.contract.Transact(opts, "setRegistry", r)
}

// SetRegistry is a paid mutator transaction binding the contract method 0xa91ee0dc.
//
// Solidity: function setRegistry(r address) returns()
func (_MultiAdmin *MultiAdminSession) SetRegistry(r common.Address) (*types.Transaction, error) {
	return _MultiAdmin.Contract.SetRegistry(&_MultiAdmin.TransactOpts, r)
}

// SetRegistry is a paid mutator transaction binding the contract method 0xa91ee0dc.
//
// Solidity: function setRegistry(r address) returns()
func (_MultiAdmin *MultiAdminTransactorSession) SetRegistry(r common.Address) (*types.Transaction, error) {
	return _MultiAdmin.Contract.SetRegistry(&_MultiAdmin.TransactOpts, r)
}

// Vote is a paid mutator transaction binding the contract method 0x128d8711.
//
// Solidity: function vote(incumbent address, challenger address, yea bool) returns()
func (_MultiAdmin *MultiAdminTransactor) Vote(opts *bind.TransactOpts, incumbent common.Address, challenger common.Address, yea bool) (*types.Transaction, error) {
	return _MultiAdmin.contract.Transact(opts, "vote", incumbent, challenger, yea)
}

// Vote is a paid mutator transaction binding the contract method 0x128d8711.
//
// Solidity: function vote(incumbent address, challenger address, yea bool) returns()
func (_MultiAdmin *MultiAdminSession) Vote(incumbent common.Address, challenger common.Address, yea bool) (*types.Transaction, error) {
	return _MultiAdmin.Contract.Vote(&_MultiAdmin.TransactOpts, incumbent, challenger, yea)
}

// Vote is a paid mutator transaction binding the contract method 0x128d8711.
//
// Solidity: function vote(incumbent address, challenger address, yea bool) returns()
func (_MultiAdmin *MultiAdminTransactorSession) Vote(incumbent common.Address, challenger common.Address, yea bool) (*types.Transaction, error) {
	return _MultiAdmin.Contract.Vote(&_MultiAdmin.TransactOpts, incumbent, challenger, yea)
}

// MultiAdminElectionClosedIterator is returned from FilterElectionClosed and is used to iterate over the raw logs and unpacked data for ElectionClosed events raised by the MultiAdmin contract.
type MultiAdminElectionClosedIterator struct {
	Event *MultiAdminElectionClosed // Event containing the contract specifics and raw log

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
func (it *MultiAdminElectionClosedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MultiAdminElectionClosed)
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
		it.Event = new(MultiAdminElectionClosed)
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
func (it *MultiAdminElectionClosedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MultiAdminElectionClosedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MultiAdminElectionClosed represents a ElectionClosed event raised by the MultiAdmin contract.
type MultiAdminElectionClosed struct {
	Incumbent  common.Address
	Challenger common.Address
	Passed     bool
	Margin     *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterElectionClosed is a free log retrieval operation binding the contract event 0x81a623b5156020c64c45d3fb56c628333f904077134ff40c1983fd1b23de573e.
//
// Solidity: e ElectionClosed(incumbent indexed address, challenger indexed address, passed indexed bool, margin uint256)
func (_MultiAdmin *MultiAdminFilterer) FilterElectionClosed(opts *bind.FilterOpts, incumbent []common.Address, challenger []common.Address, passed []bool) (*MultiAdminElectionClosedIterator, error) {

	var incumbentRule []interface{}
	for _, incumbentItem := range incumbent {
		incumbentRule = append(incumbentRule, incumbentItem)
	}
	var challengerRule []interface{}
	for _, challengerItem := range challenger {
		challengerRule = append(challengerRule, challengerItem)
	}
	var passedRule []interface{}
	for _, passedItem := range passed {
		passedRule = append(passedRule, passedItem)
	}

	logs, sub, err := _MultiAdmin.contract.FilterLogs(opts, "ElectionClosed", incumbentRule, challengerRule, passedRule)
	if err != nil {
		return nil, err
	}
	return &MultiAdminElectionClosedIterator{contract: _MultiAdmin.contract, event: "ElectionClosed", logs: logs, sub: sub}, nil
}

// WatchElectionClosed is a free log subscription operation binding the contract event 0x81a623b5156020c64c45d3fb56c628333f904077134ff40c1983fd1b23de573e.
//
// Solidity: e ElectionClosed(incumbent indexed address, challenger indexed address, passed indexed bool, margin uint256)
func (_MultiAdmin *MultiAdminFilterer) WatchElectionClosed(opts *bind.WatchOpts, sink chan<- *MultiAdminElectionClosed, incumbent []common.Address, challenger []common.Address, passed []bool) (event.Subscription, error) {

	var incumbentRule []interface{}
	for _, incumbentItem := range incumbent {
		incumbentRule = append(incumbentRule, incumbentItem)
	}
	var challengerRule []interface{}
	for _, challengerItem := range challenger {
		challengerRule = append(challengerRule, challengerItem)
	}
	var passedRule []interface{}
	for _, passedItem := range passed {
		passedRule = append(passedRule, passedItem)
	}

	logs, sub, err := _MultiAdmin.contract.WatchLogs(opts, "ElectionClosed", incumbentRule, challengerRule, passedRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MultiAdminElectionClosed)
				if err := _MultiAdmin.contract.UnpackLog(event, "ElectionClosed", log); err != nil {
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

// MultiAdminElectionCreatedIterator is returned from FilterElectionCreated and is used to iterate over the raw logs and unpacked data for ElectionCreated events raised by the MultiAdmin contract.
type MultiAdminElectionCreatedIterator struct {
	Event *MultiAdminElectionCreated // Event containing the contract specifics and raw log

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
func (it *MultiAdminElectionCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MultiAdminElectionCreated)
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
		it.Event = new(MultiAdminElectionCreated)
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
func (it *MultiAdminElectionCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MultiAdminElectionCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MultiAdminElectionCreated represents a ElectionCreated event raised by the MultiAdmin contract.
type MultiAdminElectionCreated struct {
	Incumbent  common.Address
	Challenger common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterElectionCreated is a free log retrieval operation binding the contract event 0x2931cee8058f0a9299aff15bc55251858ff1a2c223d96c6d55ebe9c67ee0351b.
//
// Solidity: e ElectionCreated(incumbent indexed address, challenger indexed address)
func (_MultiAdmin *MultiAdminFilterer) FilterElectionCreated(opts *bind.FilterOpts, incumbent []common.Address, challenger []common.Address) (*MultiAdminElectionCreatedIterator, error) {

	var incumbentRule []interface{}
	for _, incumbentItem := range incumbent {
		incumbentRule = append(incumbentRule, incumbentItem)
	}
	var challengerRule []interface{}
	for _, challengerItem := range challenger {
		challengerRule = append(challengerRule, challengerItem)
	}

	logs, sub, err := _MultiAdmin.contract.FilterLogs(opts, "ElectionCreated", incumbentRule, challengerRule)
	if err != nil {
		return nil, err
	}
	return &MultiAdminElectionCreatedIterator{contract: _MultiAdmin.contract, event: "ElectionCreated", logs: logs, sub: sub}, nil
}

// WatchElectionCreated is a free log subscription operation binding the contract event 0x2931cee8058f0a9299aff15bc55251858ff1a2c223d96c6d55ebe9c67ee0351b.
//
// Solidity: e ElectionCreated(incumbent indexed address, challenger indexed address)
func (_MultiAdmin *MultiAdminFilterer) WatchElectionCreated(opts *bind.WatchOpts, sink chan<- *MultiAdminElectionCreated, incumbent []common.Address, challenger []common.Address) (event.Subscription, error) {

	var incumbentRule []interface{}
	for _, incumbentItem := range incumbent {
		incumbentRule = append(incumbentRule, incumbentItem)
	}
	var challengerRule []interface{}
	for _, challengerItem := range challenger {
		challengerRule = append(challengerRule, challengerItem)
	}

	logs, sub, err := _MultiAdmin.contract.WatchLogs(opts, "ElectionCreated", incumbentRule, challengerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MultiAdminElectionCreated)
				if err := _MultiAdmin.contract.UnpackLog(event, "ElectionCreated", log); err != nil {
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

// MultiAdminVoteCastIterator is returned from FilterVoteCast and is used to iterate over the raw logs and unpacked data for VoteCast events raised by the MultiAdmin contract.
type MultiAdminVoteCastIterator struct {
	Event *MultiAdminVoteCast // Event containing the contract specifics and raw log

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
func (it *MultiAdminVoteCastIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MultiAdminVoteCast)
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
		it.Event = new(MultiAdminVoteCast)
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
func (it *MultiAdminVoteCastIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MultiAdminVoteCastIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MultiAdminVoteCast represents a VoteCast event raised by the MultiAdmin contract.
type MultiAdminVoteCast struct {
	Incumbent  common.Address
	Challenger common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterVoteCast is a free log retrieval operation binding the contract event 0x113ed09b18bb8fe2635af16f21d67464502d920ff1015abfe08dbe9a2971a7e2.
//
// Solidity: e VoteCast(incumbent indexed address, challenger indexed address)
func (_MultiAdmin *MultiAdminFilterer) FilterVoteCast(opts *bind.FilterOpts, incumbent []common.Address, challenger []common.Address) (*MultiAdminVoteCastIterator, error) {

	var incumbentRule []interface{}
	for _, incumbentItem := range incumbent {
		incumbentRule = append(incumbentRule, incumbentItem)
	}
	var challengerRule []interface{}
	for _, challengerItem := range challenger {
		challengerRule = append(challengerRule, challengerItem)
	}

	logs, sub, err := _MultiAdmin.contract.FilterLogs(opts, "VoteCast", incumbentRule, challengerRule)
	if err != nil {
		return nil, err
	}
	return &MultiAdminVoteCastIterator{contract: _MultiAdmin.contract, event: "VoteCast", logs: logs, sub: sub}, nil
}

// WatchVoteCast is a free log subscription operation binding the contract event 0x113ed09b18bb8fe2635af16f21d67464502d920ff1015abfe08dbe9a2971a7e2.
//
// Solidity: e VoteCast(incumbent indexed address, challenger indexed address)
func (_MultiAdmin *MultiAdminFilterer) WatchVoteCast(opts *bind.WatchOpts, sink chan<- *MultiAdminVoteCast, incumbent []common.Address, challenger []common.Address) (event.Subscription, error) {

	var incumbentRule []interface{}
	for _, incumbentItem := range incumbent {
		incumbentRule = append(incumbentRule, incumbentItem)
	}
	var challengerRule []interface{}
	for _, challengerItem := range challenger {
		challengerRule = append(challengerRule, challengerItem)
	}

	logs, sub, err := _MultiAdmin.contract.WatchLogs(opts, "VoteCast", incumbentRule, challengerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MultiAdminVoteCast)
				if err := _MultiAdmin.contract.UnpackLog(event, "VoteCast", log); err != nil {
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
