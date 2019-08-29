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
const ExternalInvestorABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"investor\",\"type\":\"address\"},{\"name\":\"accreditation\",\"type\":\"uint48\"}],\"name\":\"setAccreditation\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"investor\",\"type\":\"address\"}],\"name\":\"getHash\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"locked\",\"type\":\"bool\"}],\"name\":\"setLocked\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"investor\",\"type\":\"address\"},{\"name\":\"country\",\"type\":\"bytes2\"}],\"name\":\"setCountry\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"isAdmin\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"investor\",\"type\":\"address\"}],\"name\":\"remove\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"index\",\"type\":\"int256\"}],\"name\":\"adminAt\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"kill\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"investor\",\"type\":\"address\"},{\"name\":\"hash\",\"type\":\"bytes32\"},{\"name\":\"country\",\"type\":\"bytes2\"},{\"name\":\"accreditation\",\"type\":\"uint48\"}],\"name\":\"add\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"},{\"name\":\"add\",\"type\":\"bool\"}],\"name\":\"setAdmin\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwner\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ZERO_ADDRESS\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"investor\",\"type\":\"address\"}],\"name\":\"getAccreditation\",\"outputs\":[{\"name\":\"\",\"type\":\"uint48\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"registry\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isLocked\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"admins\",\"outputs\":[{\"name\":\"count\",\"type\":\"int256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"r\",\"type\":\"address\"}],\"name\":\"setRegistry\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"investor\",\"type\":\"address\"},{\"name\":\"frozen\",\"type\":\"bool\"}],\"name\":\"setFrozen\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"investor\",\"type\":\"address\"}],\"name\":\"getCountry\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes2\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"investor\",\"type\":\"address\"},{\"name\":\"hash\",\"type\":\"bytes32\"}],\"name\":\"setHash\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"r\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"admin\",\"type\":\"address\"}],\"name\":\"AdminAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"admin\",\"type\":\"address\"}],\"name\":\"AdminRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"oldOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnerTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"investor\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"InvestorAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"investor\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"InvestorUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"investor\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"InvestorRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"investor\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"frozen\",\"type\":\"bool\"},{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"InvestorFrozen\",\"type\":\"event\"}]"

// ExternalInvestorBin is the compiled bytecode used for deploying new contracts.
const ExternalInvestorBin = `6080604052600180546001600160a01b031916905534801561002057600080fd5b5060405161191d38038061191d8339818101604052602081101561004357600080fd5b5051600080546001600160a01b03191633179055600580546001600160a01b03909216610100026001600160a81b03199092169190911790556118928061008b6000396000f3fe608060405234801561001057600080fd5b50600436106101775760003560e01c80634fb2e45d116100d8578063a4e2d6341161008c578063ac869cd811610066578063ac869cd81461043d578063d821f92d1461046b578063e84b8169146104c657610177565b8063a4e2d63414610407578063a5de36191461040f578063a91ee0dc1461041757610177565b8063701ae59f116100bd578063701ae59f146103b65780637b103999146103f75780638da5cb5b146103ff57610177565b80634fb2e45d14610388578063538ba4f9146103ae57610177565b806329092d0e1161012f57806341c0e1b51161011457806341c0e1b5146102f057806345eb8a9b146102f85780634b0bddd21461035a57610177565b806329092d0e146102915780632bdbc56f146102b757610177565b8063211e28b611610160578063211e28b6146101ea578063242371a31461020957806324d7806c1461025757610177565b80631518c5bc1461017c5780631da0b8fc146101b2575b600080fd5b6101b06004803603604081101561019257600080fd5b5080356001600160a01b0316906020013565ffffffffffff166104f2565b005b6101d8600480360360208110156101c857600080fd5b50356001600160a01b03166105be565b60408051918252519081900360200190f35b6101b06004803603602081101561020057600080fd5b5035151561065e565b6101b06004803603604081101561021f57600080fd5b5080356001600160a01b031690602001357fffff00000000000000000000000000000000000000000000000000000000000016610733565b61027d6004803603602081101561026d57600080fd5b50356001600160a01b03166107c5565b604080519115158252519081900360200190f35b6101b0600480360360208110156102a757600080fd5b50356001600160a01b03166107f4565b6102d4600480360360208110156102cd57600080fd5b5035610928565b604080516001600160a01b039092168252519081900360200190f35b6101b061093b565b6101b06004803603608081101561030e57600080fd5b5080356001600160a01b03169060208101359060408101357fffff00000000000000000000000000000000000000000000000000000000000016906060013565ffffffffffff166109c3565b6101b06004803603604081101561037057600080fd5b506001600160a01b0381351690602001351515610b32565b6101b06004803603602081101561039e57600080fd5b50356001600160a01b0316610ce8565b6102d4610e67565b6103dc600480360360208110156103cc57600080fd5b50356001600160a01b0316610e76565b6040805165ffffffffffff9092168252519081900360200190f35b6102d4610e94565b6102d4610ea8565b61027d610eb7565b6101d8610ec0565b6101b06004803603602081101561042d57600080fd5b50356001600160a01b0316610ec6565b6101b06004803603604081101561045357600080fd5b506001600160a01b0381351690602001351515610f7f565b6104916004803603602081101561048157600080fd5b50356001600160a01b03166110c0565b604080517fffff0000000000000000000000000000000000000000000000000000000000009092168252519081900360200190f35b6101b0600480360360408110156104dc57600080fd5b506001600160a01b0381351690602001356110de565b6000546001600160a01b0316331480610517575061051760023363ffffffff61121a16565b610568576040805162461bcd60e51b815260206004820152601960248201527f41646d696e206163636f756e7420697320726571756972656400000000000000604482015290519081900360640190fd5b6005546105849061010090046001600160a01b03168383611251565b60405133906001600160a01b038416907f537d51d0d268a5585aebf683b7d15f87cac90b5755f0cedb4bf754c62cfdc6bf90600090a35050565b600554604080517f960d27d30000000000000000000000000000000000000000000000000000000081526001600160a01b038481166004830152915160009361010090049092169163960d27d391602480820192602092909190829003018186803b15801561062c57600080fd5b505afa158015610640573d6000803e3d6000fd5b505050506040513d602081101561065657600080fd5b505192915050565b6000546001600160a01b031633148061068757506001546000546001600160a01b039081169116145b6106d8576040805162461bcd60e51b815260206004820152601960248201527f4f776e6572206163636f756e7420697320726571756972656400000000000000604482015290519081900360640190fd5b60055460ff16151581151514156107205760405162461bcd60e51b81526004018080602001828103825260288152602001806118366028913960400191505060405180910390fd5b6005805460ff1916911515919091179055565b6000546001600160a01b0316331480610758575061075860023363ffffffff61121a16565b6107a9576040805162461bcd60e51b815260206004820152601960248201527f41646d696e206163636f756e7420697320726571756972656400000000000000604482015290519081900360640190fd5b6005546105849061010090046001600160a01b03168383611369565b600080546001600160a01b03838116911614806107ee57506107ee60028363ffffffff61121a16565b92915050565b6000546001600160a01b0316331480610819575061081960023363ffffffff61121a16565b61086a576040805162461bcd60e51b815260206004820152601960248201527f41646d696e206163636f756e7420697320726571756972656400000000000000604482015290519081900360640190fd5b600554604080517fc4740a950000000000000000000000000000000000000000000000000000000081526001600160a01b03848116600483015291516101009093049091169163c4740a959160248082019260009290919082900301818387803b1580156108d757600080fd5b505af11580156108eb573d6000803e3d6000fd5b50506040513392506001600160a01b03841691507fd8755221287ca1f6b28807977a086f5534d9e02ea27ebad003d7cb1a95659a4690600090a350565b60006107ee60028363ffffffff61145216565b6000546001600160a01b031633148061096457506001546000546001600160a01b039081169116145b6109b5576040805162461bcd60e51b815260206004820152601960248201527f4f776e6572206163636f756e7420697320726571756972656400000000000000604482015290519081900360640190fd5b6000546001600160a01b0316ff5b6000546001600160a01b03163314806109e857506109e860023363ffffffff61121a16565b610a39576040805162461bcd60e51b815260206004820152601960248201527f41646d696e206163636f756e7420697320726571756972656400000000000000604482015290519081900360640190fd5b60058054604080517fe8d74d1f0000000000000000000000000000000000000000000000000000000081526001600160a01b03888116600483015260248201949094526000604482018190523360648301526084820188905291516101009093049093169263e8d74d1f9260a4808301939282900301818387803b158015610ac057600080fd5b505af1158015610ad4573d6000803e3d6000fd5b5050600554610af6925061010090046001600160a01b031690508583856114d8565b60405133906001600160a01b038616907fe99183cc0b1657b54afa611991294ec1e4c458d7c36910518e2a5b76b2b6e73f90600090a350505050565b6000546001600160a01b0316331480610b5b57506001546000546001600160a01b039081169116145b610bac576040805162461bcd60e51b815260206004820152601960248201527f4f776e6572206163636f756e7420697320726571756972656400000000000000604482015290519081900360640190fd5b8015610c4d57610bc360028363ffffffff61154316565b610c14576040805162461bcd60e51b815260206004820152601360248201527f556e61626c6520746f206164642061646d696e00000000000000000000000000604482015290519081900360640190fd5b6040516001600160a01b038316907f44d6d25963f097ad14f29f06854a01f575648a1ef82f30e562ccd3889717e33990600090a2610ce4565b610c5e60028363ffffffff6115f516565b610caf576040805162461bcd60e51b815260206004820152601660248201527f556e61626c6520746f2072656d6f76652061646d696e00000000000000000000604482015290519081900360640190fd5b6040516001600160a01b038316907fa3b62bc36326052d97ea62d63c3d60308ed4c3ea8ac079dd8499f1e9c4f80c0f90600090a25b5050565b6000546001600160a01b0316331480610d1157506001546000546001600160a01b039081169116145b610d62576040805162461bcd60e51b815260206004820152601960248201527f4f776e6572206163636f756e7420697320726571756972656400000000000000604482015290519081900360640190fd5b6000546001600160a01b0382811691161415610daf5760405162461bcd60e51b81526004018080602001828103825260258152602001806118116025913960400191505060405180910390fd5b6001600160a01b038116610e0a576040805162461bcd60e51b815260206004820181905260248201527f4e6577204f776e65722063616e6e6f74206265207a65726f2061646472657373604482015290519081900360640190fd5b600080546001600160a01b0383811673ffffffffffffffffffffffffffffffffffffffff19831681178455604051919092169283917f8934ce4adea8d9ce0d714d2c22b86790e41b7731c84b926fbbdc1d40ff6533c99190a35050565b6001546001600160a01b031681565b6005546000906107ee9061010090046001600160a01b0316836116fa565b60055461010090046001600160a01b031681565b6000546001600160a01b031681565b60055460ff1681565b60025481565b6000546001600160a01b0316331480610eef57506001546000546001600160a01b039081169116145b610f40576040805162461bcd60e51b815260206004820152601960248201527f4f776e6572206163636f756e7420697320726571756972656400000000000000604482015290519081900360640190fd5b600580546001600160a01b03909216610100027fffffffffffffffffffffff0000000000000000000000000000000000000000ff909216919091179055565b6000546001600160a01b0316331480610fa45750610fa460023363ffffffff61121a16565b610ff5576040805162461bcd60e51b815260206004820152601960248201527f41646d696e206163636f756e7420697320726571756972656400000000000000604482015290519081900360640190fd5b600554604080517fcbe5404f0000000000000000000000000000000000000000000000000000000081526001600160a01b038581166004830152841515602483015291516101009093049091169163cbe5404f9160448082019260009290919082900301818387803b15801561106a57600080fd5b505af115801561107e573d6000803e3d6000fd5b505060405133925083151591506001600160a01b038516907f6d30448ca28c66e149273ddd6d39fe9cb1982d4013649080aeb9d8251356d38190600090a45050565b6005546000906107ee9061010090046001600160a01b031683611785565b6000546001600160a01b0316331480611103575061110360023363ffffffff61121a16565b611154576040805162461bcd60e51b815260206004820152601960248201527f41646d696e206163636f756e7420697320726571756972656400000000000000604482015290519081900360640190fd5b600554604080517f1706acfc0000000000000000000000000000000000000000000000000000000081526001600160a01b03858116600483015260248201859052915161010090930490911691631706acfc9160448082019260009290919082900301818387803b1580156111c857600080fd5b505af11580156111dc573d6000803e3d6000fd5b50506040513392506001600160a01b03851691507f537d51d0d268a5585aebf683b7d15f87cac90b5755f0cedb4bf754c62cfdc6bf90600090a35050565b6001600160a01b0381166000908152600183016020526040812054600019018181128015906112495750835481125b949350505050565b6040805163572bc6b160e01b81526001600160a01b03848116600483015260006024830181905292519086169163572bc6b1916044808301926020929190829003018186803b1580156112a357600080fd5b505afa1580156112b7573d6000803e3d6000fd5b505050506040513d60208110156112cd57600080fd5b50516040805163278e1db360e11b81526001600160a01b03868116600483015260006024830181905267ffffffffffff0000601088901b811667ffffffffffff0000199096169590951760448401819052935193955090881692634f1c3b6692606480820193929182900301818387803b15801561134a57600080fd5b505af115801561135e573d6000803e3d6000fd5b505050505050505050565b6040805163572bc6b160e01b81526001600160a01b03848116600483015260006024830181905292519086169163572bc6b1916044808301926020929190829003018186803b1580156113bb57600080fd5b505afa1580156113cf573d6000803e3d6000fd5b505050506040513d60208110156113e557600080fd5b50516040805163278e1db360e11b81526001600160a01b03868116600483015260006024830181905261ffff1990941660f087901c1760448301819052925192945061ffff9390881692634f1c3b6692606480820193929182900301818387803b15801561134a57600080fd5b60008082121580156114645750825482125b6114b5576040805162461bcd60e51b815260206004820152601860248201527f496e646578206f757473696465206f6620626f756e64732e0000000000000000604482015290519081900360640190fd5b50600101600090815260029190910160205260409020546001600160a01b031690565b6040805163278e1db360e11b81526001600160a01b03858116600483015260006024830181905260f085901c67ffffffffffff0000601088901b161760448401819052935191881692634f1c3b66926064808301939282900301818387803b15801561134a57600080fd5b60006001600160a01b03821661155b575060006107ee565b6001600160a01b03821660009081526001840160205260408120546000190190811280159061158a5750835481125b156115995760009150506107ee565b5050815460019081018084556001600160a01b03831660008181528386016020908152604080832085905593825260028701905291909120805473ffffffffffffffffffffffffffffffffffffffff1916909117905592915050565b6001600160a01b0381166000908152600180840160205260408220549081128061161f5750835481135b1561162e5760009150506107ee565b83548112156116a257835460009081526002850160208181526040808420546001600160a01b031680855260018901835281852086905585855292909152808320805473ffffffffffffffffffffffffffffffffffffffff19908116909317905586548352909120805490911690556116ce565b60008181526002850160205260409020805473ffffffffffffffffffffffffffffffffffffffff191690555b50506001600160a01b031660009081526001828101602052604082209190915581546000190190915590565b6040805163572bc6b160e01b81526001600160a01b038381166004830152600060248301819052925183929186169163572bc6b1916044808301926020929190829003018186803b15801561174e57600080fd5b505afa158015611762573d6000803e3d6000fd5b505050506040513d602081101561177857600080fd5b505160101c949350505050565b6040805163572bc6b160e01b81526001600160a01b038381166004830152600060248301819052925183929186169163572bc6b1916044808301926020929190829003018186803b1580156117d957600080fd5b505afa1580156117ed573d6000803e3d6000fd5b505050506040513d602081101561180357600080fd5b505160f01b94935050505056fe4e6577204f776e65722063616e6e6f74206265207468652063757272656e74206f776e6572436f6e747261637420616c726561647920696e20726571756573746564206c6f636b207374617465a265627a7a7230582058d4f4eb3c94f2504ad77ab4f5747a609008a4aca07c491510252c6a3e4fd11964736f6c634300050a0032`

// DeployExternalInvestor deploys a new Ethereum contract, binding an instance of ExternalInvestor to it.
func DeployExternalInvestor(auth *bind.TransactOpts, backend bind.ContractBackend, r common.Address) (common.Address, *types.Transaction, *ExternalInvestor, error) {
	parsed, err := abi.JSON(strings.NewReader(ExternalInvestorABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ExternalInvestorBin), backend, r)
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

// ZEROADDRESS is a free data retrieval call binding the contract method 0x538ba4f9.
//
// Solidity: function ZERO_ADDRESS() constant returns(address)
func (_ExternalInvestor *ExternalInvestorCaller) ZEROADDRESS(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ExternalInvestor.contract.Call(opts, out, "ZERO_ADDRESS")
	return *ret0, err
}

// ZEROADDRESS is a free data retrieval call binding the contract method 0x538ba4f9.
//
// Solidity: function ZERO_ADDRESS() constant returns(address)
func (_ExternalInvestor *ExternalInvestorSession) ZEROADDRESS() (common.Address, error) {
	return _ExternalInvestor.Contract.ZEROADDRESS(&_ExternalInvestor.CallOpts)
}

// ZEROADDRESS is a free data retrieval call binding the contract method 0x538ba4f9.
//
// Solidity: function ZERO_ADDRESS() constant returns(address)
func (_ExternalInvestor *ExternalInvestorCallerSession) ZEROADDRESS() (common.Address, error) {
	return _ExternalInvestor.Contract.ZEROADDRESS(&_ExternalInvestor.CallOpts)
}

// AdminAt is a free data retrieval call binding the contract method 0x2bdbc56f.
//
// Solidity: function adminAt(index int256) constant returns(address)
func (_ExternalInvestor *ExternalInvestorCaller) AdminAt(opts *bind.CallOpts, index *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ExternalInvestor.contract.Call(opts, out, "adminAt", index)
	return *ret0, err
}

// AdminAt is a free data retrieval call binding the contract method 0x2bdbc56f.
//
// Solidity: function adminAt(index int256) constant returns(address)
func (_ExternalInvestor *ExternalInvestorSession) AdminAt(index *big.Int) (common.Address, error) {
	return _ExternalInvestor.Contract.AdminAt(&_ExternalInvestor.CallOpts, index)
}

// AdminAt is a free data retrieval call binding the contract method 0x2bdbc56f.
//
// Solidity: function adminAt(index int256) constant returns(address)
func (_ExternalInvestor *ExternalInvestorCallerSession) AdminAt(index *big.Int) (common.Address, error) {
	return _ExternalInvestor.Contract.AdminAt(&_ExternalInvestor.CallOpts, index)
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
// Solidity: function getAccreditation(investor address) constant returns(uint48)
func (_ExternalInvestor *ExternalInvestorCaller) GetAccreditation(opts *bind.CallOpts, investor common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ExternalInvestor.contract.Call(opts, out, "getAccreditation", investor)
	return *ret0, err
}

// GetAccreditation is a free data retrieval call binding the contract method 0x701ae59f.
//
// Solidity: function getAccreditation(investor address) constant returns(uint48)
func (_ExternalInvestor *ExternalInvestorSession) GetAccreditation(investor common.Address) (*big.Int, error) {
	return _ExternalInvestor.Contract.GetAccreditation(&_ExternalInvestor.CallOpts, investor)
}

// GetAccreditation is a free data retrieval call binding the contract method 0x701ae59f.
//
// Solidity: function getAccreditation(investor address) constant returns(uint48)
func (_ExternalInvestor *ExternalInvestorCallerSession) GetAccreditation(investor common.Address) (*big.Int, error) {
	return _ExternalInvestor.Contract.GetAccreditation(&_ExternalInvestor.CallOpts, investor)
}

// GetCountry is a free data retrieval call binding the contract method 0xd821f92d.
//
// Solidity: function getCountry(investor address) constant returns(bytes2)
func (_ExternalInvestor *ExternalInvestorCaller) GetCountry(opts *bind.CallOpts, investor common.Address) ([2]byte, error) {
	var (
		ret0 = new([2]byte)
	)
	out := ret0
	err := _ExternalInvestor.contract.Call(opts, out, "getCountry", investor)
	return *ret0, err
}

// GetCountry is a free data retrieval call binding the contract method 0xd821f92d.
//
// Solidity: function getCountry(investor address) constant returns(bytes2)
func (_ExternalInvestor *ExternalInvestorSession) GetCountry(investor common.Address) ([2]byte, error) {
	return _ExternalInvestor.Contract.GetCountry(&_ExternalInvestor.CallOpts, investor)
}

// GetCountry is a free data retrieval call binding the contract method 0xd821f92d.
//
// Solidity: function getCountry(investor address) constant returns(bytes2)
func (_ExternalInvestor *ExternalInvestorCallerSession) GetCountry(investor common.Address) ([2]byte, error) {
	return _ExternalInvestor.Contract.GetCountry(&_ExternalInvestor.CallOpts, investor)
}

// GetHash is a free data retrieval call binding the contract method 0x1da0b8fc.
//
// Solidity: function getHash(investor address) constant returns(bytes32)
func (_ExternalInvestor *ExternalInvestorCaller) GetHash(opts *bind.CallOpts, investor common.Address) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _ExternalInvestor.contract.Call(opts, out, "getHash", investor)
	return *ret0, err
}

// GetHash is a free data retrieval call binding the contract method 0x1da0b8fc.
//
// Solidity: function getHash(investor address) constant returns(bytes32)
func (_ExternalInvestor *ExternalInvestorSession) GetHash(investor common.Address) ([32]byte, error) {
	return _ExternalInvestor.Contract.GetHash(&_ExternalInvestor.CallOpts, investor)
}

// GetHash is a free data retrieval call binding the contract method 0x1da0b8fc.
//
// Solidity: function getHash(investor address) constant returns(bytes32)
func (_ExternalInvestor *ExternalInvestorCallerSession) GetHash(investor common.Address) ([32]byte, error) {
	return _ExternalInvestor.Contract.GetHash(&_ExternalInvestor.CallOpts, investor)
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

// Registry is a free data retrieval call binding the contract method 0x7b103999.
//
// Solidity: function registry() constant returns(address)
func (_ExternalInvestor *ExternalInvestorCaller) Registry(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ExternalInvestor.contract.Call(opts, out, "registry")
	return *ret0, err
}

// Registry is a free data retrieval call binding the contract method 0x7b103999.
//
// Solidity: function registry() constant returns(address)
func (_ExternalInvestor *ExternalInvestorSession) Registry() (common.Address, error) {
	return _ExternalInvestor.Contract.Registry(&_ExternalInvestor.CallOpts)
}

// Registry is a free data retrieval call binding the contract method 0x7b103999.
//
// Solidity: function registry() constant returns(address)
func (_ExternalInvestor *ExternalInvestorCallerSession) Registry() (common.Address, error) {
	return _ExternalInvestor.Contract.Registry(&_ExternalInvestor.CallOpts)
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

// SetAdmin is a paid mutator transaction binding the contract method 0x4b0bddd2.
//
// Solidity: function setAdmin(addr address, add bool) returns()
func (_ExternalInvestor *ExternalInvestorTransactor) SetAdmin(opts *bind.TransactOpts, addr common.Address, add bool) (*types.Transaction, error) {
	return _ExternalInvestor.contract.Transact(opts, "setAdmin", addr, add)
}

// SetAdmin is a paid mutator transaction binding the contract method 0x4b0bddd2.
//
// Solidity: function setAdmin(addr address, add bool) returns()
func (_ExternalInvestor *ExternalInvestorSession) SetAdmin(addr common.Address, add bool) (*types.Transaction, error) {
	return _ExternalInvestor.Contract.SetAdmin(&_ExternalInvestor.TransactOpts, addr, add)
}

// SetAdmin is a paid mutator transaction binding the contract method 0x4b0bddd2.
//
// Solidity: function setAdmin(addr address, add bool) returns()
func (_ExternalInvestor *ExternalInvestorTransactorSession) SetAdmin(addr common.Address, add bool) (*types.Transaction, error) {
	return _ExternalInvestor.Contract.SetAdmin(&_ExternalInvestor.TransactOpts, addr, add)
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

// SetRegistry is a paid mutator transaction binding the contract method 0xa91ee0dc.
//
// Solidity: function setRegistry(r address) returns()
func (_ExternalInvestor *ExternalInvestorTransactor) SetRegistry(opts *bind.TransactOpts, r common.Address) (*types.Transaction, error) {
	return _ExternalInvestor.contract.Transact(opts, "setRegistry", r)
}

// SetRegistry is a paid mutator transaction binding the contract method 0xa91ee0dc.
//
// Solidity: function setRegistry(r address) returns()
func (_ExternalInvestor *ExternalInvestorSession) SetRegistry(r common.Address) (*types.Transaction, error) {
	return _ExternalInvestor.Contract.SetRegistry(&_ExternalInvestor.TransactOpts, r)
}

// SetRegistry is a paid mutator transaction binding the contract method 0xa91ee0dc.
//
// Solidity: function setRegistry(r address) returns()
func (_ExternalInvestor *ExternalInvestorTransactorSession) SetRegistry(r common.Address) (*types.Transaction, error) {
	return _ExternalInvestor.Contract.SetRegistry(&_ExternalInvestor.TransactOpts, r)
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

// ExternalInvestorInvestorUpdatedIterator is returned from FilterInvestorUpdated and is used to iterate over the raw logs and unpacked data for InvestorUpdated events raised by the ExternalInvestor contract.
type ExternalInvestorInvestorUpdatedIterator struct {
	Event *ExternalInvestorInvestorUpdated // Event containing the contract specifics and raw log

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
func (it *ExternalInvestorInvestorUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExternalInvestorInvestorUpdated)
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
		it.Event = new(ExternalInvestorInvestorUpdated)
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
func (it *ExternalInvestorInvestorUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExternalInvestorInvestorUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExternalInvestorInvestorUpdated represents a InvestorUpdated event raised by the ExternalInvestor contract.
type ExternalInvestorInvestorUpdated struct {
	Investor common.Address
	Owner    common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterInvestorUpdated is a free log retrieval operation binding the contract event 0x537d51d0d268a5585aebf683b7d15f87cac90b5755f0cedb4bf754c62cfdc6bf.
//
// Solidity: e InvestorUpdated(investor indexed address, owner indexed address)
func (_ExternalInvestor *ExternalInvestorFilterer) FilterInvestorUpdated(opts *bind.FilterOpts, investor []common.Address, owner []common.Address) (*ExternalInvestorInvestorUpdatedIterator, error) {

	var investorRule []interface{}
	for _, investorItem := range investor {
		investorRule = append(investorRule, investorItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _ExternalInvestor.contract.FilterLogs(opts, "InvestorUpdated", investorRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return &ExternalInvestorInvestorUpdatedIterator{contract: _ExternalInvestor.contract, event: "InvestorUpdated", logs: logs, sub: sub}, nil
}

// WatchInvestorUpdated is a free log subscription operation binding the contract event 0x537d51d0d268a5585aebf683b7d15f87cac90b5755f0cedb4bf754c62cfdc6bf.
//
// Solidity: e InvestorUpdated(investor indexed address, owner indexed address)
func (_ExternalInvestor *ExternalInvestorFilterer) WatchInvestorUpdated(opts *bind.WatchOpts, sink chan<- *ExternalInvestorInvestorUpdated, investor []common.Address, owner []common.Address) (event.Subscription, error) {

	var investorRule []interface{}
	for _, investorItem := range investor {
		investorRule = append(investorRule, investorItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _ExternalInvestor.contract.WatchLogs(opts, "InvestorUpdated", investorRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExternalInvestorInvestorUpdated)
				if err := _ExternalInvestor.contract.UnpackLog(event, "InvestorUpdated", log); err != nil {
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
