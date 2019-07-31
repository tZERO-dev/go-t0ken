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

// InvestorABI is the input ABI used to generate the binding from.
const InvestorABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"investor\",\"type\":\"address\"},{\"name\":\"accreditation\",\"type\":\"uint48\"}],\"name\":\"setAccreditation\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"investor\",\"type\":\"address\"}],\"name\":\"getHash\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"locked\",\"type\":\"bool\"}],\"name\":\"setLocked\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"investor\",\"type\":\"address\"},{\"name\":\"country\",\"type\":\"bytes2\"}],\"name\":\"setCountry\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"investor\",\"type\":\"address\"}],\"name\":\"remove\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"kill\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"investor\",\"type\":\"address\"},{\"name\":\"hash\",\"type\":\"bytes32\"},{\"name\":\"country\",\"type\":\"bytes2\"},{\"name\":\"accreditation\",\"type\":\"uint48\"}],\"name\":\"add\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwner\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ZERO_ADDRESS\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"investor\",\"type\":\"address\"}],\"name\":\"getAccreditation\",\"outputs\":[{\"name\":\"\",\"type\":\"uint48\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"registry\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isLocked\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"r\",\"type\":\"address\"}],\"name\":\"setRegistry\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"investor\",\"type\":\"address\"},{\"name\":\"frozen\",\"type\":\"bool\"}],\"name\":\"setFrozen\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"investor\",\"type\":\"address\"}],\"name\":\"getCountry\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes2\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"investor\",\"type\":\"address\"},{\"name\":\"hash\",\"type\":\"bytes32\"}],\"name\":\"setHash\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"r\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"oldOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnerTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"investor\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"InvestorAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"investor\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"InvestorUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"investor\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"InvestorRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"investor\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"frozen\",\"type\":\"bool\"},{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"InvestorFrozen\",\"type\":\"event\"}]"

// InvestorBin is the compiled bytecode used for deploying new contracts.
const InvestorBin = `6080604052600180546001600160a01b031916905534801561002057600080fd5b5060405161256b38038061256b8339818101604052602081101561004357600080fd5b5051600080546001600160a01b031990811633179091556001805460ff60a01b19169055600280546001600160a01b03909316929091169190911790556124dc8061008f6000396000f3fe608060405234801561001057600080fd5b506004361061011b5760003560e01c8063538ba4f9116100b2578063a4e2d63411610081578063ac869cd811610066578063ac869cd814610368578063d821f92d14610396578063e84b8169146103f15761011b565b8063a4e2d63414610326578063a91ee0dc146103425761011b565b8063538ba4f9146102b1578063701ae59f146102d55780637b103999146103165780638da5cb5b1461031e5761011b565b806329092d0e116100ee57806329092d0e146101fb57806341c0e1b51461022157806345eb8a9b146102295780634fb2e45d1461028b5761011b565b80631518c5bc146101205780631da0b8fc14610156578063211e28b61461018e578063242371a3146101ad575b600080fd5b6101546004803603604081101561013657600080fd5b5080356001600160a01b0316906020013565ffffffffffff1661041d565b005b61017c6004803603602081101561016c57600080fd5b50356001600160a01b0316610803565b60408051918252519081900360200190f35b610154600480360360208110156101a457600080fd5b5035151561089f565b610154600480360360408110156101c357600080fd5b5080356001600160a01b031690602001357fffff000000000000000000000000000000000000000000000000000000000000166109a2565b6101546004803603602081101561021157600080fd5b50356001600160a01b0316610d4d565b610154611194565b6101546004803603608081101561023f57600080fd5b5080356001600160a01b03169060208101359060408101357fffff00000000000000000000000000000000000000000000000000000000000016906060013565ffffffffffff1661121c565b610154600480360360208110156102a157600080fd5b50356001600160a01b03166115e1565b6102b9611760565b604080516001600160a01b039092168252519081900360200190f35b6102fb600480360360208110156102eb57600080fd5b50356001600160a01b031661176f565b6040805165ffffffffffff9092168252519081900360200190f35b6102b9611794565b6102b96117a3565b61032e6117b2565b604080519115158252519081900360200190f35b6101546004803603602081101561035857600080fd5b50356001600160a01b03166117c2565b6101546004803603604081101561037e57600080fd5b506001600160a01b038135169060200135151561186b565b6103bc600480360360208110156103ac57600080fd5b50356001600160a01b0316611c26565b604080517fffff0000000000000000000000000000000000000000000000000000000000009092168252519081900360200190f35b6101546004803603604081101561040757600080fd5b506001600160a01b038135169060200135611c45565b60025460408051630a22ee7360e01b815233600482015260036024820152905184926001600160a01b031691630a22ee73916044808301926020929190829003018186803b15801561046e57600080fd5b505afa158015610482573d6000803e3d6000fd5b505050506040513d602081101561049857600080fd5b50516104eb576040805162461bcd60e51b815260206004820152601660248201527f42726f6b65722d4465616c657220726571756972656400000000000000000000604482015290519081900360640190fd5b600254604080516318d553a560e11b81526001600160a01b038481166004830152915191909216916331aaa74a916024808301926020929190829003018186803b15801561053857600080fd5b505afa15801561054c573d6000803e3d6000fd5b505050506040513d602081101561056257600080fd5b50516001600160a01b031633146105aa5760405162461bcd60e51b815260040180806020018281038252602181526020018061245f6021913960400191505060405180910390fd5b6002546040805163624522f960e01b815233600482015290516001600160a01b039092169163624522f991602480820192602092909190829003018186803b1580156105f557600080fd5b505afa158015610609573d6000803e3d6000fd5b505050506040513d602081101561061f57600080fd5b505115610673576040805162461bcd60e51b815260206004820152601760248201527f42726f6b65722d6465616c65722069732066726f7a656e000000000000000000604482015290519081900360640190fd5b600254604080516318d553a560e11b815233600482015290516001600160a01b039092169163624522f99183916331aaa74a91602480820192602092909190829003018186803b1580156106c657600080fd5b505afa1580156106da573d6000803e3d6000fd5b505050506040513d60208110156106f057600080fd5b5051604080516001600160e01b031960e085901b1681526001600160a01b039092166004830152516024808301926020929190829003018186803b15801561073757600080fd5b505afa15801561074b573d6000803e3d6000fd5b505050506040513d602081101561076157600080fd5b5051156107ab576040805162461bcd60e51b815260206004820152601360248201527221bab9ba37b234b0b71034b990333937bd32b760691b604482015290519081900360640190fd5b6002546107c8906001600160a01b0316848463ffffffff61209416565b60405133906001600160a01b038516907f537d51d0d268a5585aebf683b7d15f87cac90b5755f0cedb4bf754c62cfdc6bf90600090a3505050565b600254604080517f960d27d30000000000000000000000000000000000000000000000000000000081526001600160a01b0384811660048301529151600093929092169163960d27d391602480820192602092909190829003018186803b15801561086d57600080fd5b505afa158015610881573d6000803e3d6000fd5b505050506040513d602081101561089757600080fd5b505192915050565b6000546001600160a01b03163314806108c857506001546000546001600160a01b039081169116145b610919576040805162461bcd60e51b815260206004820152601960248201527f4f776e6572206163636f756e7420697320726571756972656400000000000000604482015290519081900360640190fd5b60015460ff600160a01b90910416151581151514156109695760405162461bcd60e51b81526004018080602001828103825260288152602001806124806028913960400191505060405180910390fd5b60018054911515600160a01b027fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff909216919091179055565b60025460408051630a22ee7360e01b815233600482015260036024820152905184926001600160a01b031691630a22ee73916044808301926020929190829003018186803b1580156109f357600080fd5b505afa158015610a07573d6000803e3d6000fd5b505050506040513d6020811015610a1d57600080fd5b5051610a70576040805162461bcd60e51b815260206004820152601660248201527f42726f6b65722d4465616c657220726571756972656400000000000000000000604482015290519081900360640190fd5b600254604080516318d553a560e11b81526001600160a01b038481166004830152915191909216916331aaa74a916024808301926020929190829003018186803b158015610abd57600080fd5b505afa158015610ad1573d6000803e3d6000fd5b505050506040513d6020811015610ae757600080fd5b50516001600160a01b03163314610b2f5760405162461bcd60e51b815260040180806020018281038252602181526020018061245f6021913960400191505060405180910390fd5b6002546040805163624522f960e01b815233600482015290516001600160a01b039092169163624522f991602480820192602092909190829003018186803b158015610b7a57600080fd5b505afa158015610b8e573d6000803e3d6000fd5b505050506040513d6020811015610ba457600080fd5b505115610bf8576040805162461bcd60e51b815260206004820152601760248201527f42726f6b65722d6465616c65722069732066726f7a656e000000000000000000604482015290519081900360640190fd5b600254604080516318d553a560e11b815233600482015290516001600160a01b039092169163624522f99183916331aaa74a91602480820192602092909190829003018186803b158015610c4b57600080fd5b505afa158015610c5f573d6000803e3d6000fd5b505050506040513d6020811015610c7557600080fd5b5051604080516001600160e01b031960e085901b1681526001600160a01b039092166004830152516024808301926020929190829003018186803b158015610cbc57600080fd5b505afa158015610cd0573d6000803e3d6000fd5b505050506040513d6020811015610ce657600080fd5b505115610d30576040805162461bcd60e51b815260206004820152601360248201527221bab9ba37b234b0b71034b990333937bd32b760691b604482015290519081900360640190fd5b6002546107c8906001600160a01b0316848463ffffffff6121ac16565b60025460408051630a22ee7360e01b815233600482015260036024820152905183926001600160a01b031691630a22ee73916044808301926020929190829003018186803b158015610d9e57600080fd5b505afa158015610db2573d6000803e3d6000fd5b505050506040513d6020811015610dc857600080fd5b5051610e1b576040805162461bcd60e51b815260206004820152601660248201527f42726f6b65722d4465616c657220726571756972656400000000000000000000604482015290519081900360640190fd5b600254604080516318d553a560e11b81526001600160a01b038481166004830152915191909216916331aaa74a916024808301926020929190829003018186803b158015610e6857600080fd5b505afa158015610e7c573d6000803e3d6000fd5b505050506040513d6020811015610e9257600080fd5b50516001600160a01b03163314610eda5760405162461bcd60e51b815260040180806020018281038252602181526020018061245f6021913960400191505060405180910390fd5b6002546040805163624522f960e01b815233600482015290516001600160a01b039092169163624522f991602480820192602092909190829003018186803b158015610f2557600080fd5b505afa158015610f39573d6000803e3d6000fd5b505050506040513d6020811015610f4f57600080fd5b505115610fa3576040805162461bcd60e51b815260206004820152601760248201527f42726f6b65722d6465616c65722069732066726f7a656e000000000000000000604482015290519081900360640190fd5b600254604080516318d553a560e11b815233600482015290516001600160a01b039092169163624522f99183916331aaa74a91602480820192602092909190829003018186803b158015610ff657600080fd5b505afa15801561100a573d6000803e3d6000fd5b505050506040513d602081101561102057600080fd5b5051604080516001600160e01b031960e085901b1681526001600160a01b039092166004830152516024808301926020929190829003018186803b15801561106757600080fd5b505afa15801561107b573d6000803e3d6000fd5b505050506040513d602081101561109157600080fd5b5051156110db576040805162461bcd60e51b815260206004820152601360248201527221bab9ba37b234b0b71034b990333937bd32b760691b604482015290519081900360640190fd5b600254604080517fc4740a950000000000000000000000000000000000000000000000000000000081526001600160a01b0385811660048301529151919092169163c4740a9591602480830192600092919082900301818387803b15801561114257600080fd5b505af1158015611156573d6000803e3d6000fd5b50506040513392506001600160a01b03851691507fd8755221287ca1f6b28807977a086f5534d9e02ea27ebad003d7cb1a95659a4690600090a35050565b6000546001600160a01b03163314806111bd57506001546000546001600160a01b039081169116145b61120e576040805162461bcd60e51b815260206004820152601960248201527f4f776e6572206163636f756e7420697320726571756972656400000000000000604482015290519081900360640190fd5b6000546001600160a01b0316ff5b60025460408051630a22ee7360e01b81523360048201526003602482015290516001600160a01b0390921691630a22ee7391604480820192602092909190829003018186803b15801561126e57600080fd5b505afa158015611282573d6000803e3d6000fd5b505050506040513d602081101561129857600080fd5b50516112eb576040805162461bcd60e51b815260206004820152601660248201527f42726f6b65722d6465616c657220726571756972656400000000000000000000604482015290519081900360640190fd5b6002546040805163624522f960e01b815233600482015290516001600160a01b039092169163624522f991602480820192602092909190829003018186803b15801561133657600080fd5b505afa15801561134a573d6000803e3d6000fd5b505050506040513d602081101561136057600080fd5b5051156113b4576040805162461bcd60e51b815260206004820152601760248201527f42726f6b65722d6465616c65722069732066726f7a656e000000000000000000604482015290519081900360640190fd5b600254604080516318d553a560e11b815233600482015290516001600160a01b039092169163624522f99183916331aaa74a91602480820192602092909190829003018186803b15801561140757600080fd5b505afa15801561141b573d6000803e3d6000fd5b505050506040513d602081101561143157600080fd5b5051604080516001600160e01b031960e085901b1681526001600160a01b039092166004830152516024808301926020929190829003018186803b15801561147857600080fd5b505afa15801561148c573d6000803e3d6000fd5b505050506040513d60208110156114a257600080fd5b5051156114ec576040805162461bcd60e51b815260206004820152601360248201527221bab9ba37b234b0b71034b990333937bd32b760691b604482015290519081900360640190fd5b600254604080517fe8d74d1f0000000000000000000000000000000000000000000000000000000081526001600160a01b03878116600483810191909152602483015260006044830181905233606484015260848301889052925193169263e8d74d1f9260a48084019391929182900301818387803b15801561156e57600080fd5b505af1158015611582573d6000803e3d6000fd5b50506002546115a592506001600160a01b0316905085838563ffffffff61229516565b60405133906001600160a01b038616907fe99183cc0b1657b54afa611991294ec1e4c458d7c36910518e2a5b76b2b6e73f90600090a350505050565b6000546001600160a01b031633148061160a57506001546000546001600160a01b039081169116145b61165b576040805162461bcd60e51b815260206004820152601960248201527f4f776e6572206163636f756e7420697320726571756972656400000000000000604482015290519081900360640190fd5b6000546001600160a01b03828116911614156116a85760405162461bcd60e51b81526004018080602001828103825260258152602001806124176025913960400191505060405180910390fd5b6001600160a01b038116611703576040805162461bcd60e51b815260206004820181905260248201527f4e6577204f776e65722063616e6e6f74206265207a65726f2061646472657373604482015290519081900360640190fd5b600080546001600160a01b0383811673ffffffffffffffffffffffffffffffffffffffff19831681178455604051919092169283917f8934ce4adea8d9ce0d714d2c22b86790e41b7731c84b926fbbdc1d40ff6533c99190a35050565b6001546001600160a01b031681565b60025460009061178e906001600160a01b03168363ffffffff61230016565b92915050565b6002546001600160a01b031681565b6000546001600160a01b031681565b600154600160a01b900460ff1681565b6000546001600160a01b03163314806117eb57506001546000546001600160a01b039081169116145b61183c576040805162461bcd60e51b815260206004820152601960248201527f4f776e6572206163636f756e7420697320726571756972656400000000000000604482015290519081900360640190fd5b6002805473ffffffffffffffffffffffffffffffffffffffff19166001600160a01b0392909216919091179055565b600254604080516318d553a560e11b81526001600160a01b0380861660048301529151859360009316916331aaa74a916024808301926020929190829003018186803b1580156118ba57600080fd5b505afa1580156118ce573d6000803e3d6000fd5b505050506040513d60208110156118e457600080fd5b5051600254604080516318d553a560e11b81526001600160a01b038085166004830152915193945060009391909216916331aaa74a916024808301926020929190829003018186803b15801561193957600080fd5b505afa15801561194d573d6000803e3d6000fd5b505050506040513d602081101561196357600080fd5b50519050336001600160a01b03831614806119865750336001600160a01b038216145b6119c15760405162461bcd60e51b815260040180806020018281038252602381526020018061243c6023913960400191505060405180910390fd5b6002546040805163624522f960e01b81526001600160a01b0384811660048301529151919092169163624522f9916024808301926020929190829003018186803b158015611a0e57600080fd5b505afa158015611a22573d6000803e3d6000fd5b505050506040513d6020811015611a3857600080fd5b505115611a82576040805162461bcd60e51b815260206004820152601360248201527221bab9ba37b234b0b71034b990333937bd32b760691b604482015290519081900360640190fd5b336001600160a01b0383161415611b5e576002546040805163624522f960e01b81526001600160a01b0385811660048301529151919092169163624522f9916024808301926020929190829003018186803b158015611ae057600080fd5b505afa158015611af4573d6000803e3d6000fd5b505050506040513d6020811015611b0a57600080fd5b505115611b5e576040805162461bcd60e51b815260206004820152601760248201527f42726f6b65722d4465616c65722069732066726f7a656e000000000000000000604482015290519081900360640190fd5b600254604080517fcbe5404f0000000000000000000000000000000000000000000000000000000081526001600160a01b03888116600483015287151560248301529151919092169163cbe5404f91604480830192600092919082900301818387803b158015611bcd57600080fd5b505af1158015611be1573d6000803e3d6000fd5b505060405133925086151591506001600160a01b038816907f6d30448ca28c66e149273ddd6d39fe9cb1982d4013649080aeb9d8251356d38190600090a45050505050565b60025460009061178e906001600160a01b03168363ffffffff61238b16565b60025460408051630a22ee7360e01b815233600482015260036024820152905184926001600160a01b031691630a22ee73916044808301926020929190829003018186803b158015611c9657600080fd5b505afa158015611caa573d6000803e3d6000fd5b505050506040513d6020811015611cc057600080fd5b5051611d13576040805162461bcd60e51b815260206004820152601660248201527f42726f6b65722d4465616c657220726571756972656400000000000000000000604482015290519081900360640190fd5b600254604080516318d553a560e11b81526001600160a01b038481166004830152915191909216916331aaa74a916024808301926020929190829003018186803b158015611d6057600080fd5b505afa158015611d74573d6000803e3d6000fd5b505050506040513d6020811015611d8a57600080fd5b50516001600160a01b03163314611dd25760405162461bcd60e51b815260040180806020018281038252602181526020018061245f6021913960400191505060405180910390fd5b6002546040805163624522f960e01b815233600482015290516001600160a01b039092169163624522f991602480820192602092909190829003018186803b158015611e1d57600080fd5b505afa158015611e31573d6000803e3d6000fd5b505050506040513d6020811015611e4757600080fd5b505115611e9b576040805162461bcd60e51b815260206004820152601760248201527f42726f6b65722d6465616c65722069732066726f7a656e000000000000000000604482015290519081900360640190fd5b600254604080516318d553a560e11b815233600482015290516001600160a01b039092169163624522f99183916331aaa74a91602480820192602092909190829003018186803b158015611eee57600080fd5b505afa158015611f02573d6000803e3d6000fd5b505050506040513d6020811015611f1857600080fd5b5051604080516001600160e01b031960e085901b1681526001600160a01b039092166004830152516024808301926020929190829003018186803b158015611f5f57600080fd5b505afa158015611f73573d6000803e3d6000fd5b505050506040513d6020811015611f8957600080fd5b505115611fd3576040805162461bcd60e51b815260206004820152601360248201527221bab9ba37b234b0b71034b990333937bd32b760691b604482015290519081900360640190fd5b600254604080517f1706acfc0000000000000000000000000000000000000000000000000000000081526001600160a01b0386811660048301526024820186905291519190921691631706acfc91604480830192600092919082900301818387803b15801561204157600080fd5b505af1158015612055573d6000803e3d6000fd5b50506040513392506001600160a01b03861691507f537d51d0d268a5585aebf683b7d15f87cac90b5755f0cedb4bf754c62cfdc6bf90600090a3505050565b6040805163572bc6b160e01b81526001600160a01b03848116600483015260006024830181905292519086169163572bc6b1916044808301926020929190829003018186803b1580156120e657600080fd5b505afa1580156120fa573d6000803e3d6000fd5b505050506040513d602081101561211057600080fd5b50516040805163278e1db360e11b81526001600160a01b03868116600483015260006024830181905267ffffffffffff0000601088901b811667ffffffffffff0000199096169590951760448401819052935193955090881692634f1c3b6692606480820193929182900301818387803b15801561218d57600080fd5b505af11580156121a1573d6000803e3d6000fd5b505050505050505050565b6040805163572bc6b160e01b81526001600160a01b03848116600483015260006024830181905292519086169163572bc6b1916044808301926020929190829003018186803b1580156121fe57600080fd5b505afa158015612212573d6000803e3d6000fd5b505050506040513d602081101561222857600080fd5b50516040805163278e1db360e11b81526001600160a01b03868116600483015260006024830181905261ffff1990941660f087901c1760448301819052925192945061ffff9390881692634f1c3b6692606480820193929182900301818387803b15801561218d57600080fd5b6040805163278e1db360e11b81526001600160a01b03858116600483015260006024830181905260f085901c67ffffffffffff0000601088901b161760448401819052935191881692634f1c3b66926064808301939282900301818387803b15801561218d57600080fd5b6040805163572bc6b160e01b81526001600160a01b038381166004830152600060248301819052925183929186169163572bc6b1916044808301926020929190829003018186803b15801561235457600080fd5b505afa158015612368573d6000803e3d6000fd5b505050506040513d602081101561237e57600080fd5b505160101c949350505050565b6040805163572bc6b160e01b81526001600160a01b038381166004830152600060248301819052925183929186169163572bc6b1916044808301926020929190829003018186803b1580156123df57600080fd5b505afa1580156123f3573d6000803e3d6000fd5b505050506040513d602081101561240957600080fd5b505160f01b94935050505056fe4e6577204f776e65722063616e6e6f74206265207468652063757272656e74206f776e657242726f6b65722d4465616c6572206f7220437573746f6469616e207265717569726564496e766573746f7227732062726f6b65722d6465616c6572207265717569726564436f6e747261637420616c726561647920696e20726571756573746564206c6f636b207374617465a265627a7a723058206b3e0f2c26fb5d1e12970dc421619ffb058a3e7a8e215b6ea358e574cdc68aad64736f6c634300050a0032`

// DeployInvestor deploys a new Ethereum contract, binding an instance of Investor to it.
func DeployInvestor(auth *bind.TransactOpts, backend bind.ContractBackend, r common.Address) (common.Address, *types.Transaction, *Investor, error) {
	parsed, err := abi.JSON(strings.NewReader(InvestorABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(InvestorBin), backend, r)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Investor{InvestorCaller: InvestorCaller{contract: contract}, InvestorTransactor: InvestorTransactor{contract: contract}, InvestorFilterer: InvestorFilterer{contract: contract}}, nil
}

// Investor is an auto generated Go binding around an Ethereum contract.
type Investor struct {
	InvestorCaller     // Read-only binding to the contract
	InvestorTransactor // Write-only binding to the contract
	InvestorFilterer   // Log filterer for contract events
}

// InvestorCaller is an auto generated read-only Go binding around an Ethereum contract.
type InvestorCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InvestorTransactor is an auto generated write-only Go binding around an Ethereum contract.
type InvestorTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InvestorFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type InvestorFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InvestorSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type InvestorSession struct {
	Contract     *Investor         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// InvestorCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type InvestorCallerSession struct {
	Contract *InvestorCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// InvestorTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type InvestorTransactorSession struct {
	Contract     *InvestorTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// InvestorRaw is an auto generated low-level Go binding around an Ethereum contract.
type InvestorRaw struct {
	Contract *Investor // Generic contract binding to access the raw methods on
}

// InvestorCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type InvestorCallerRaw struct {
	Contract *InvestorCaller // Generic read-only contract binding to access the raw methods on
}

// InvestorTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type InvestorTransactorRaw struct {
	Contract *InvestorTransactor // Generic write-only contract binding to access the raw methods on
}

// NewInvestor creates a new instance of Investor, bound to a specific deployed contract.
func NewInvestor(address common.Address, backend bind.ContractBackend) (*Investor, error) {
	contract, err := bindInvestor(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Investor{InvestorCaller: InvestorCaller{contract: contract}, InvestorTransactor: InvestorTransactor{contract: contract}, InvestorFilterer: InvestorFilterer{contract: contract}}, nil
}

// NewInvestorCaller creates a new read-only instance of Investor, bound to a specific deployed contract.
func NewInvestorCaller(address common.Address, caller bind.ContractCaller) (*InvestorCaller, error) {
	contract, err := bindInvestor(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &InvestorCaller{contract: contract}, nil
}

// NewInvestorTransactor creates a new write-only instance of Investor, bound to a specific deployed contract.
func NewInvestorTransactor(address common.Address, transactor bind.ContractTransactor) (*InvestorTransactor, error) {
	contract, err := bindInvestor(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &InvestorTransactor{contract: contract}, nil
}

// NewInvestorFilterer creates a new log filterer instance of Investor, bound to a specific deployed contract.
func NewInvestorFilterer(address common.Address, filterer bind.ContractFilterer) (*InvestorFilterer, error) {
	contract, err := bindInvestor(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &InvestorFilterer{contract: contract}, nil
}

// bindInvestor binds a generic wrapper to an already deployed contract.
func bindInvestor(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(InvestorABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Investor *InvestorRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Investor.Contract.InvestorCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Investor *InvestorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Investor.Contract.InvestorTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Investor *InvestorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Investor.Contract.InvestorTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Investor *InvestorCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Investor.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Investor *InvestorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Investor.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Investor *InvestorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Investor.Contract.contract.Transact(opts, method, params...)
}

// ZEROADDRESS is a free data retrieval call binding the contract method 0x538ba4f9.
//
// Solidity: function ZERO_ADDRESS() constant returns(address)
func (_Investor *InvestorCaller) ZEROADDRESS(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Investor.contract.Call(opts, out, "ZERO_ADDRESS")
	return *ret0, err
}

// ZEROADDRESS is a free data retrieval call binding the contract method 0x538ba4f9.
//
// Solidity: function ZERO_ADDRESS() constant returns(address)
func (_Investor *InvestorSession) ZEROADDRESS() (common.Address, error) {
	return _Investor.Contract.ZEROADDRESS(&_Investor.CallOpts)
}

// ZEROADDRESS is a free data retrieval call binding the contract method 0x538ba4f9.
//
// Solidity: function ZERO_ADDRESS() constant returns(address)
func (_Investor *InvestorCallerSession) ZEROADDRESS() (common.Address, error) {
	return _Investor.Contract.ZEROADDRESS(&_Investor.CallOpts)
}

// GetAccreditation is a free data retrieval call binding the contract method 0x701ae59f.
//
// Solidity: function getAccreditation(investor address) constant returns(uint48)
func (_Investor *InvestorCaller) GetAccreditation(opts *bind.CallOpts, investor common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Investor.contract.Call(opts, out, "getAccreditation", investor)
	return *ret0, err
}

// GetAccreditation is a free data retrieval call binding the contract method 0x701ae59f.
//
// Solidity: function getAccreditation(investor address) constant returns(uint48)
func (_Investor *InvestorSession) GetAccreditation(investor common.Address) (*big.Int, error) {
	return _Investor.Contract.GetAccreditation(&_Investor.CallOpts, investor)
}

// GetAccreditation is a free data retrieval call binding the contract method 0x701ae59f.
//
// Solidity: function getAccreditation(investor address) constant returns(uint48)
func (_Investor *InvestorCallerSession) GetAccreditation(investor common.Address) (*big.Int, error) {
	return _Investor.Contract.GetAccreditation(&_Investor.CallOpts, investor)
}

// GetCountry is a free data retrieval call binding the contract method 0xd821f92d.
//
// Solidity: function getCountry(investor address) constant returns(bytes2)
func (_Investor *InvestorCaller) GetCountry(opts *bind.CallOpts, investor common.Address) ([2]byte, error) {
	var (
		ret0 = new([2]byte)
	)
	out := ret0
	err := _Investor.contract.Call(opts, out, "getCountry", investor)
	return *ret0, err
}

// GetCountry is a free data retrieval call binding the contract method 0xd821f92d.
//
// Solidity: function getCountry(investor address) constant returns(bytes2)
func (_Investor *InvestorSession) GetCountry(investor common.Address) ([2]byte, error) {
	return _Investor.Contract.GetCountry(&_Investor.CallOpts, investor)
}

// GetCountry is a free data retrieval call binding the contract method 0xd821f92d.
//
// Solidity: function getCountry(investor address) constant returns(bytes2)
func (_Investor *InvestorCallerSession) GetCountry(investor common.Address) ([2]byte, error) {
	return _Investor.Contract.GetCountry(&_Investor.CallOpts, investor)
}

// GetHash is a free data retrieval call binding the contract method 0x1da0b8fc.
//
// Solidity: function getHash(investor address) constant returns(bytes32)
func (_Investor *InvestorCaller) GetHash(opts *bind.CallOpts, investor common.Address) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Investor.contract.Call(opts, out, "getHash", investor)
	return *ret0, err
}

// GetHash is a free data retrieval call binding the contract method 0x1da0b8fc.
//
// Solidity: function getHash(investor address) constant returns(bytes32)
func (_Investor *InvestorSession) GetHash(investor common.Address) ([32]byte, error) {
	return _Investor.Contract.GetHash(&_Investor.CallOpts, investor)
}

// GetHash is a free data retrieval call binding the contract method 0x1da0b8fc.
//
// Solidity: function getHash(investor address) constant returns(bytes32)
func (_Investor *InvestorCallerSession) GetHash(investor common.Address) ([32]byte, error) {
	return _Investor.Contract.GetHash(&_Investor.CallOpts, investor)
}

// IsLocked is a free data retrieval call binding the contract method 0xa4e2d634.
//
// Solidity: function isLocked() constant returns(bool)
func (_Investor *InvestorCaller) IsLocked(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Investor.contract.Call(opts, out, "isLocked")
	return *ret0, err
}

// IsLocked is a free data retrieval call binding the contract method 0xa4e2d634.
//
// Solidity: function isLocked() constant returns(bool)
func (_Investor *InvestorSession) IsLocked() (bool, error) {
	return _Investor.Contract.IsLocked(&_Investor.CallOpts)
}

// IsLocked is a free data retrieval call binding the contract method 0xa4e2d634.
//
// Solidity: function isLocked() constant returns(bool)
func (_Investor *InvestorCallerSession) IsLocked() (bool, error) {
	return _Investor.Contract.IsLocked(&_Investor.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Investor *InvestorCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Investor.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Investor *InvestorSession) Owner() (common.Address, error) {
	return _Investor.Contract.Owner(&_Investor.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Investor *InvestorCallerSession) Owner() (common.Address, error) {
	return _Investor.Contract.Owner(&_Investor.CallOpts)
}

// Registry is a free data retrieval call binding the contract method 0x7b103999.
//
// Solidity: function registry() constant returns(address)
func (_Investor *InvestorCaller) Registry(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Investor.contract.Call(opts, out, "registry")
	return *ret0, err
}

// Registry is a free data retrieval call binding the contract method 0x7b103999.
//
// Solidity: function registry() constant returns(address)
func (_Investor *InvestorSession) Registry() (common.Address, error) {
	return _Investor.Contract.Registry(&_Investor.CallOpts)
}

// Registry is a free data retrieval call binding the contract method 0x7b103999.
//
// Solidity: function registry() constant returns(address)
func (_Investor *InvestorCallerSession) Registry() (common.Address, error) {
	return _Investor.Contract.Registry(&_Investor.CallOpts)
}

// Add is a paid mutator transaction binding the contract method 0x45eb8a9b.
//
// Solidity: function add(investor address, hash bytes32, country bytes2, accreditation uint48) returns()
func (_Investor *InvestorTransactor) Add(opts *bind.TransactOpts, investor common.Address, hash [32]byte, country [2]byte, accreditation *big.Int) (*types.Transaction, error) {
	return _Investor.contract.Transact(opts, "add", investor, hash, country, accreditation)
}

// Add is a paid mutator transaction binding the contract method 0x45eb8a9b.
//
// Solidity: function add(investor address, hash bytes32, country bytes2, accreditation uint48) returns()
func (_Investor *InvestorSession) Add(investor common.Address, hash [32]byte, country [2]byte, accreditation *big.Int) (*types.Transaction, error) {
	return _Investor.Contract.Add(&_Investor.TransactOpts, investor, hash, country, accreditation)
}

// Add is a paid mutator transaction binding the contract method 0x45eb8a9b.
//
// Solidity: function add(investor address, hash bytes32, country bytes2, accreditation uint48) returns()
func (_Investor *InvestorTransactorSession) Add(investor common.Address, hash [32]byte, country [2]byte, accreditation *big.Int) (*types.Transaction, error) {
	return _Investor.Contract.Add(&_Investor.TransactOpts, investor, hash, country, accreditation)
}

// Kill is a paid mutator transaction binding the contract method 0x41c0e1b5.
//
// Solidity: function kill() returns()
func (_Investor *InvestorTransactor) Kill(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Investor.contract.Transact(opts, "kill")
}

// Kill is a paid mutator transaction binding the contract method 0x41c0e1b5.
//
// Solidity: function kill() returns()
func (_Investor *InvestorSession) Kill() (*types.Transaction, error) {
	return _Investor.Contract.Kill(&_Investor.TransactOpts)
}

// Kill is a paid mutator transaction binding the contract method 0x41c0e1b5.
//
// Solidity: function kill() returns()
func (_Investor *InvestorTransactorSession) Kill() (*types.Transaction, error) {
	return _Investor.Contract.Kill(&_Investor.TransactOpts)
}

// Remove is a paid mutator transaction binding the contract method 0x29092d0e.
//
// Solidity: function remove(investor address) returns()
func (_Investor *InvestorTransactor) Remove(opts *bind.TransactOpts, investor common.Address) (*types.Transaction, error) {
	return _Investor.contract.Transact(opts, "remove", investor)
}

// Remove is a paid mutator transaction binding the contract method 0x29092d0e.
//
// Solidity: function remove(investor address) returns()
func (_Investor *InvestorSession) Remove(investor common.Address) (*types.Transaction, error) {
	return _Investor.Contract.Remove(&_Investor.TransactOpts, investor)
}

// Remove is a paid mutator transaction binding the contract method 0x29092d0e.
//
// Solidity: function remove(investor address) returns()
func (_Investor *InvestorTransactorSession) Remove(investor common.Address) (*types.Transaction, error) {
	return _Investor.Contract.Remove(&_Investor.TransactOpts, investor)
}

// SetAccreditation is a paid mutator transaction binding the contract method 0x1518c5bc.
//
// Solidity: function setAccreditation(investor address, accreditation uint48) returns()
func (_Investor *InvestorTransactor) SetAccreditation(opts *bind.TransactOpts, investor common.Address, accreditation *big.Int) (*types.Transaction, error) {
	return _Investor.contract.Transact(opts, "setAccreditation", investor, accreditation)
}

// SetAccreditation is a paid mutator transaction binding the contract method 0x1518c5bc.
//
// Solidity: function setAccreditation(investor address, accreditation uint48) returns()
func (_Investor *InvestorSession) SetAccreditation(investor common.Address, accreditation *big.Int) (*types.Transaction, error) {
	return _Investor.Contract.SetAccreditation(&_Investor.TransactOpts, investor, accreditation)
}

// SetAccreditation is a paid mutator transaction binding the contract method 0x1518c5bc.
//
// Solidity: function setAccreditation(investor address, accreditation uint48) returns()
func (_Investor *InvestorTransactorSession) SetAccreditation(investor common.Address, accreditation *big.Int) (*types.Transaction, error) {
	return _Investor.Contract.SetAccreditation(&_Investor.TransactOpts, investor, accreditation)
}

// SetCountry is a paid mutator transaction binding the contract method 0x242371a3.
//
// Solidity: function setCountry(investor address, country bytes2) returns()
func (_Investor *InvestorTransactor) SetCountry(opts *bind.TransactOpts, investor common.Address, country [2]byte) (*types.Transaction, error) {
	return _Investor.contract.Transact(opts, "setCountry", investor, country)
}

// SetCountry is a paid mutator transaction binding the contract method 0x242371a3.
//
// Solidity: function setCountry(investor address, country bytes2) returns()
func (_Investor *InvestorSession) SetCountry(investor common.Address, country [2]byte) (*types.Transaction, error) {
	return _Investor.Contract.SetCountry(&_Investor.TransactOpts, investor, country)
}

// SetCountry is a paid mutator transaction binding the contract method 0x242371a3.
//
// Solidity: function setCountry(investor address, country bytes2) returns()
func (_Investor *InvestorTransactorSession) SetCountry(investor common.Address, country [2]byte) (*types.Transaction, error) {
	return _Investor.Contract.SetCountry(&_Investor.TransactOpts, investor, country)
}

// SetFrozen is a paid mutator transaction binding the contract method 0xac869cd8.
//
// Solidity: function setFrozen(investor address, frozen bool) returns()
func (_Investor *InvestorTransactor) SetFrozen(opts *bind.TransactOpts, investor common.Address, frozen bool) (*types.Transaction, error) {
	return _Investor.contract.Transact(opts, "setFrozen", investor, frozen)
}

// SetFrozen is a paid mutator transaction binding the contract method 0xac869cd8.
//
// Solidity: function setFrozen(investor address, frozen bool) returns()
func (_Investor *InvestorSession) SetFrozen(investor common.Address, frozen bool) (*types.Transaction, error) {
	return _Investor.Contract.SetFrozen(&_Investor.TransactOpts, investor, frozen)
}

// SetFrozen is a paid mutator transaction binding the contract method 0xac869cd8.
//
// Solidity: function setFrozen(investor address, frozen bool) returns()
func (_Investor *InvestorTransactorSession) SetFrozen(investor common.Address, frozen bool) (*types.Transaction, error) {
	return _Investor.Contract.SetFrozen(&_Investor.TransactOpts, investor, frozen)
}

// SetHash is a paid mutator transaction binding the contract method 0xe84b8169.
//
// Solidity: function setHash(investor address, hash bytes32) returns()
func (_Investor *InvestorTransactor) SetHash(opts *bind.TransactOpts, investor common.Address, hash [32]byte) (*types.Transaction, error) {
	return _Investor.contract.Transact(opts, "setHash", investor, hash)
}

// SetHash is a paid mutator transaction binding the contract method 0xe84b8169.
//
// Solidity: function setHash(investor address, hash bytes32) returns()
func (_Investor *InvestorSession) SetHash(investor common.Address, hash [32]byte) (*types.Transaction, error) {
	return _Investor.Contract.SetHash(&_Investor.TransactOpts, investor, hash)
}

// SetHash is a paid mutator transaction binding the contract method 0xe84b8169.
//
// Solidity: function setHash(investor address, hash bytes32) returns()
func (_Investor *InvestorTransactorSession) SetHash(investor common.Address, hash [32]byte) (*types.Transaction, error) {
	return _Investor.Contract.SetHash(&_Investor.TransactOpts, investor, hash)
}

// SetLocked is a paid mutator transaction binding the contract method 0x211e28b6.
//
// Solidity: function setLocked(locked bool) returns()
func (_Investor *InvestorTransactor) SetLocked(opts *bind.TransactOpts, locked bool) (*types.Transaction, error) {
	return _Investor.contract.Transact(opts, "setLocked", locked)
}

// SetLocked is a paid mutator transaction binding the contract method 0x211e28b6.
//
// Solidity: function setLocked(locked bool) returns()
func (_Investor *InvestorSession) SetLocked(locked bool) (*types.Transaction, error) {
	return _Investor.Contract.SetLocked(&_Investor.TransactOpts, locked)
}

// SetLocked is a paid mutator transaction binding the contract method 0x211e28b6.
//
// Solidity: function setLocked(locked bool) returns()
func (_Investor *InvestorTransactorSession) SetLocked(locked bool) (*types.Transaction, error) {
	return _Investor.Contract.SetLocked(&_Investor.TransactOpts, locked)
}

// SetRegistry is a paid mutator transaction binding the contract method 0xa91ee0dc.
//
// Solidity: function setRegistry(r address) returns()
func (_Investor *InvestorTransactor) SetRegistry(opts *bind.TransactOpts, r common.Address) (*types.Transaction, error) {
	return _Investor.contract.Transact(opts, "setRegistry", r)
}

// SetRegistry is a paid mutator transaction binding the contract method 0xa91ee0dc.
//
// Solidity: function setRegistry(r address) returns()
func (_Investor *InvestorSession) SetRegistry(r common.Address) (*types.Transaction, error) {
	return _Investor.Contract.SetRegistry(&_Investor.TransactOpts, r)
}

// SetRegistry is a paid mutator transaction binding the contract method 0xa91ee0dc.
//
// Solidity: function setRegistry(r address) returns()
func (_Investor *InvestorTransactorSession) SetRegistry(r common.Address) (*types.Transaction, error) {
	return _Investor.Contract.SetRegistry(&_Investor.TransactOpts, r)
}

// TransferOwner is a paid mutator transaction binding the contract method 0x4fb2e45d.
//
// Solidity: function transferOwner(newOwner address) returns()
func (_Investor *InvestorTransactor) TransferOwner(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Investor.contract.Transact(opts, "transferOwner", newOwner)
}

// TransferOwner is a paid mutator transaction binding the contract method 0x4fb2e45d.
//
// Solidity: function transferOwner(newOwner address) returns()
func (_Investor *InvestorSession) TransferOwner(newOwner common.Address) (*types.Transaction, error) {
	return _Investor.Contract.TransferOwner(&_Investor.TransactOpts, newOwner)
}

// TransferOwner is a paid mutator transaction binding the contract method 0x4fb2e45d.
//
// Solidity: function transferOwner(newOwner address) returns()
func (_Investor *InvestorTransactorSession) TransferOwner(newOwner common.Address) (*types.Transaction, error) {
	return _Investor.Contract.TransferOwner(&_Investor.TransactOpts, newOwner)
}

// InvestorInvestorAddedIterator is returned from FilterInvestorAdded and is used to iterate over the raw logs and unpacked data for InvestorAdded events raised by the Investor contract.
type InvestorInvestorAddedIterator struct {
	Event *InvestorInvestorAdded // Event containing the contract specifics and raw log

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
func (it *InvestorInvestorAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InvestorInvestorAdded)
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
		it.Event = new(InvestorInvestorAdded)
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
func (it *InvestorInvestorAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InvestorInvestorAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InvestorInvestorAdded represents a InvestorAdded event raised by the Investor contract.
type InvestorInvestorAdded struct {
	Investor common.Address
	Owner    common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterInvestorAdded is a free log retrieval operation binding the contract event 0xe99183cc0b1657b54afa611991294ec1e4c458d7c36910518e2a5b76b2b6e73f.
//
// Solidity: e InvestorAdded(investor indexed address, owner indexed address)
func (_Investor *InvestorFilterer) FilterInvestorAdded(opts *bind.FilterOpts, investor []common.Address, owner []common.Address) (*InvestorInvestorAddedIterator, error) {

	var investorRule []interface{}
	for _, investorItem := range investor {
		investorRule = append(investorRule, investorItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Investor.contract.FilterLogs(opts, "InvestorAdded", investorRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return &InvestorInvestorAddedIterator{contract: _Investor.contract, event: "InvestorAdded", logs: logs, sub: sub}, nil
}

// WatchInvestorAdded is a free log subscription operation binding the contract event 0xe99183cc0b1657b54afa611991294ec1e4c458d7c36910518e2a5b76b2b6e73f.
//
// Solidity: e InvestorAdded(investor indexed address, owner indexed address)
func (_Investor *InvestorFilterer) WatchInvestorAdded(opts *bind.WatchOpts, sink chan<- *InvestorInvestorAdded, investor []common.Address, owner []common.Address) (event.Subscription, error) {

	var investorRule []interface{}
	for _, investorItem := range investor {
		investorRule = append(investorRule, investorItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Investor.contract.WatchLogs(opts, "InvestorAdded", investorRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InvestorInvestorAdded)
				if err := _Investor.contract.UnpackLog(event, "InvestorAdded", log); err != nil {
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

// InvestorInvestorFrozenIterator is returned from FilterInvestorFrozen and is used to iterate over the raw logs and unpacked data for InvestorFrozen events raised by the Investor contract.
type InvestorInvestorFrozenIterator struct {
	Event *InvestorInvestorFrozen // Event containing the contract specifics and raw log

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
func (it *InvestorInvestorFrozenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InvestorInvestorFrozen)
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
		it.Event = new(InvestorInvestorFrozen)
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
func (it *InvestorInvestorFrozenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InvestorInvestorFrozenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InvestorInvestorFrozen represents a InvestorFrozen event raised by the Investor contract.
type InvestorInvestorFrozen struct {
	Investor common.Address
	Frozen   bool
	Owner    common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterInvestorFrozen is a free log retrieval operation binding the contract event 0x6d30448ca28c66e149273ddd6d39fe9cb1982d4013649080aeb9d8251356d381.
//
// Solidity: e InvestorFrozen(investor indexed address, frozen indexed bool, owner indexed address)
func (_Investor *InvestorFilterer) FilterInvestorFrozen(opts *bind.FilterOpts, investor []common.Address, frozen []bool, owner []common.Address) (*InvestorInvestorFrozenIterator, error) {

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

	logs, sub, err := _Investor.contract.FilterLogs(opts, "InvestorFrozen", investorRule, frozenRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return &InvestorInvestorFrozenIterator{contract: _Investor.contract, event: "InvestorFrozen", logs: logs, sub: sub}, nil
}

// WatchInvestorFrozen is a free log subscription operation binding the contract event 0x6d30448ca28c66e149273ddd6d39fe9cb1982d4013649080aeb9d8251356d381.
//
// Solidity: e InvestorFrozen(investor indexed address, frozen indexed bool, owner indexed address)
func (_Investor *InvestorFilterer) WatchInvestorFrozen(opts *bind.WatchOpts, sink chan<- *InvestorInvestorFrozen, investor []common.Address, frozen []bool, owner []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _Investor.contract.WatchLogs(opts, "InvestorFrozen", investorRule, frozenRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InvestorInvestorFrozen)
				if err := _Investor.contract.UnpackLog(event, "InvestorFrozen", log); err != nil {
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

// InvestorInvestorRemovedIterator is returned from FilterInvestorRemoved and is used to iterate over the raw logs and unpacked data for InvestorRemoved events raised by the Investor contract.
type InvestorInvestorRemovedIterator struct {
	Event *InvestorInvestorRemoved // Event containing the contract specifics and raw log

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
func (it *InvestorInvestorRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InvestorInvestorRemoved)
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
		it.Event = new(InvestorInvestorRemoved)
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
func (it *InvestorInvestorRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InvestorInvestorRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InvestorInvestorRemoved represents a InvestorRemoved event raised by the Investor contract.
type InvestorInvestorRemoved struct {
	Investor common.Address
	Owner    common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterInvestorRemoved is a free log retrieval operation binding the contract event 0xd8755221287ca1f6b28807977a086f5534d9e02ea27ebad003d7cb1a95659a46.
//
// Solidity: e InvestorRemoved(investor indexed address, owner indexed address)
func (_Investor *InvestorFilterer) FilterInvestorRemoved(opts *bind.FilterOpts, investor []common.Address, owner []common.Address) (*InvestorInvestorRemovedIterator, error) {

	var investorRule []interface{}
	for _, investorItem := range investor {
		investorRule = append(investorRule, investorItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Investor.contract.FilterLogs(opts, "InvestorRemoved", investorRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return &InvestorInvestorRemovedIterator{contract: _Investor.contract, event: "InvestorRemoved", logs: logs, sub: sub}, nil
}

// WatchInvestorRemoved is a free log subscription operation binding the contract event 0xd8755221287ca1f6b28807977a086f5534d9e02ea27ebad003d7cb1a95659a46.
//
// Solidity: e InvestorRemoved(investor indexed address, owner indexed address)
func (_Investor *InvestorFilterer) WatchInvestorRemoved(opts *bind.WatchOpts, sink chan<- *InvestorInvestorRemoved, investor []common.Address, owner []common.Address) (event.Subscription, error) {

	var investorRule []interface{}
	for _, investorItem := range investor {
		investorRule = append(investorRule, investorItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Investor.contract.WatchLogs(opts, "InvestorRemoved", investorRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InvestorInvestorRemoved)
				if err := _Investor.contract.UnpackLog(event, "InvestorRemoved", log); err != nil {
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

// InvestorInvestorUpdatedIterator is returned from FilterInvestorUpdated and is used to iterate over the raw logs and unpacked data for InvestorUpdated events raised by the Investor contract.
type InvestorInvestorUpdatedIterator struct {
	Event *InvestorInvestorUpdated // Event containing the contract specifics and raw log

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
func (it *InvestorInvestorUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InvestorInvestorUpdated)
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
		it.Event = new(InvestorInvestorUpdated)
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
func (it *InvestorInvestorUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InvestorInvestorUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InvestorInvestorUpdated represents a InvestorUpdated event raised by the Investor contract.
type InvestorInvestorUpdated struct {
	Investor common.Address
	Owner    common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterInvestorUpdated is a free log retrieval operation binding the contract event 0x537d51d0d268a5585aebf683b7d15f87cac90b5755f0cedb4bf754c62cfdc6bf.
//
// Solidity: e InvestorUpdated(investor indexed address, owner indexed address)
func (_Investor *InvestorFilterer) FilterInvestorUpdated(opts *bind.FilterOpts, investor []common.Address, owner []common.Address) (*InvestorInvestorUpdatedIterator, error) {

	var investorRule []interface{}
	for _, investorItem := range investor {
		investorRule = append(investorRule, investorItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Investor.contract.FilterLogs(opts, "InvestorUpdated", investorRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return &InvestorInvestorUpdatedIterator{contract: _Investor.contract, event: "InvestorUpdated", logs: logs, sub: sub}, nil
}

// WatchInvestorUpdated is a free log subscription operation binding the contract event 0x537d51d0d268a5585aebf683b7d15f87cac90b5755f0cedb4bf754c62cfdc6bf.
//
// Solidity: e InvestorUpdated(investor indexed address, owner indexed address)
func (_Investor *InvestorFilterer) WatchInvestorUpdated(opts *bind.WatchOpts, sink chan<- *InvestorInvestorUpdated, investor []common.Address, owner []common.Address) (event.Subscription, error) {

	var investorRule []interface{}
	for _, investorItem := range investor {
		investorRule = append(investorRule, investorItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Investor.contract.WatchLogs(opts, "InvestorUpdated", investorRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InvestorInvestorUpdated)
				if err := _Investor.contract.UnpackLog(event, "InvestorUpdated", log); err != nil {
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

// InvestorOwnerTransferredIterator is returned from FilterOwnerTransferred and is used to iterate over the raw logs and unpacked data for OwnerTransferred events raised by the Investor contract.
type InvestorOwnerTransferredIterator struct {
	Event *InvestorOwnerTransferred // Event containing the contract specifics and raw log

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
func (it *InvestorOwnerTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InvestorOwnerTransferred)
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
		it.Event = new(InvestorOwnerTransferred)
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
func (it *InvestorOwnerTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InvestorOwnerTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InvestorOwnerTransferred represents a OwnerTransferred event raised by the Investor contract.
type InvestorOwnerTransferred struct {
	OldOwner common.Address
	NewOwner common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOwnerTransferred is a free log retrieval operation binding the contract event 0x8934ce4adea8d9ce0d714d2c22b86790e41b7731c84b926fbbdc1d40ff6533c9.
//
// Solidity: e OwnerTransferred(oldOwner indexed address, newOwner indexed address)
func (_Investor *InvestorFilterer) FilterOwnerTransferred(opts *bind.FilterOpts, oldOwner []common.Address, newOwner []common.Address) (*InvestorOwnerTransferredIterator, error) {

	var oldOwnerRule []interface{}
	for _, oldOwnerItem := range oldOwner {
		oldOwnerRule = append(oldOwnerRule, oldOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Investor.contract.FilterLogs(opts, "OwnerTransferred", oldOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &InvestorOwnerTransferredIterator{contract: _Investor.contract, event: "OwnerTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnerTransferred is a free log subscription operation binding the contract event 0x8934ce4adea8d9ce0d714d2c22b86790e41b7731c84b926fbbdc1d40ff6533c9.
//
// Solidity: e OwnerTransferred(oldOwner indexed address, newOwner indexed address)
func (_Investor *InvestorFilterer) WatchOwnerTransferred(opts *bind.WatchOpts, sink chan<- *InvestorOwnerTransferred, oldOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var oldOwnerRule []interface{}
	for _, oldOwnerItem := range oldOwner {
		oldOwnerRule = append(oldOwnerRule, oldOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Investor.contract.WatchLogs(opts, "OwnerTransferred", oldOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InvestorOwnerTransferred)
				if err := _Investor.contract.UnpackLog(event, "OwnerTransferred", log); err != nil {
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
