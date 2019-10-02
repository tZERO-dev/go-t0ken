// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package token

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

// T0kenSplittableABI is the input ABI used to generate the binding from.
const T0kenSplittableABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokens\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"numerator\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"denominator\",\"type\":\"uint128\"}],\"name\":\"setSplitMultiplier\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"issuer\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bool\",\"name\":\"locked\",\"type\":\"bool\"}],\"name\":\"setLocked\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokens\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"balances\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"splitBalance\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"kill\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"issuanceFinished\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwner\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ZERO_ADDRESS\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"newIssuer\",\"type\":\"address\"}],\"name\":\"setIssuer\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"compliance\",\"outputs\":[{\"internalType\":\"contractCompliance\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokens\",\"type\":\"uint256\"}],\"name\":\"transferOverride\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"int256\",\"name\":\"index\",\"type\":\"int256\"}],\"name\":\"holderAt\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"holders\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"count\",\"type\":\"int256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"cashedOutTokenQuantity\",\"type\":\"uint256\"}],\"name\":\"splitTotalSupply\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isLocked\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"quantity\",\"type\":\"uint256\"}],\"name\":\"issueTokens\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokens\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"finishIssuance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"isHolder\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"addrOwner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"split\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"numerator\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"denominator\",\"type\":\"uint128\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"newComplianceAddress\",\"type\":\"address\"}],\"name\":\"setCompliance\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"tokenName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"tokenSymbol\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnerTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousIssuer\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newIssuer\",\"type\":\"address\"}],\"name\":\"IssuerSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokens\",\"type\":\"uint256\"}],\"name\":\"Issuance\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"IssuanceFinished\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"shareholder\",\"type\":\"address\"}],\"name\":\"ShareholderAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"shareholder\",\"type\":\"address\"}],\"name\":\"ShareholderRemoved\",\"type\":\"event\"}]"

// T0kenSplittableBin is the compiled bytecode used for deploying new contracts.
const T0kenSplittableBin = `600180546001600160a01b031916815560c0604052608081905260a0819052600d80546001600160801b0319169091176001600160801b03167001000000000000000000000000000000001790553480156200005a57600080fd5b506040516200237638038062002376833981810160405260408110156200008057600080fd5b8101908080516040519392919084640100000000821115620000a157600080fd5b908301906020820185811115620000b757600080fd5b8251640100000000811182820188101715620000d257600080fd5b82525081516020918201929091019080838360005b8381101562000101578181015183820152602001620000e7565b50505050905090810190601f1680156200012f5780820380516001836020036101000a031916815260200191505b50604052602001805160405193929190846401000000008211156200015357600080fd5b9083019060208201858111156200016957600080fd5b82516401000000008111828201881017156200018457600080fd5b82525081516020918201929091019080838360005b83811015620001b357818101518382015260200162000199565b50505050905090810190601f168015620001e15780820380516001836020036101000a031916815260200191505b506040525050600080546001600160a01b031916331781556001805460ff60a01b1916905583518492508391906200022190600290602086019062000257565b5081516200023790600390602085019062000257565b506004805460ff191660ff9290921691909117905550620002fc92505050565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106200029a57805160ff1916838001178555620002ca565b82800160010185558215620002ca579182015b82811115620002ca578251825591602001919060010190620002ad565b50620002d8929150620002dc565b5090565b620002f991905b80821115620002d85760008155600101620002e3565b90565b61206a806200030c6000396000f3fe608060405234801561001057600080fd5b50600436106101e55760003560e01c80636290865d1161010f578063a4e2d634116100a2578063d4d7b19a11610071578063d4d7b19a14610550578063dd62ed3e14610576578063f7654176146105a4578063f8981789146105db576101e5565b8063a4e2d634146104f7578063a5820daa146104ff578063a9059cbb1461051c578063c422293b14610548576101e5565b80638188f71c116100de5780638188f71c146104c25780638da5cb5b146104ca57806395d89b41146104d257806398d34b9b146104da576101e5565b80636290865d1461044157806370a082311461044957806380318be81461046f5780638082a929146104a5576101e5565b806327e235e3116101875780634662299a116101565780634662299a146103e55780634fb2e45d146103ed578063538ba4f91461041357806355cc4e571461041b576101e5565b806327e235e3146103735780632bbc501914610399578063313ce567146103bf57806341c0e1b5146103dd576101e5565b806318160ddd116101c357806318160ddd146102e05780631d143848146102fa578063211e28b61461031e57806323b872dd1461033d576101e5565b806306fdde03146101ea578063095ea7b3146102675780630c1554a1146102a7575b600080fd5b6101f2610601565b6040805160208082528351818301528351919283929083019185019080838360005b8381101561022c578181015183820152602001610214565b50505050905090810190601f1680156102595780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b6102936004803603604081101561027d57600080fd5b506001600160a01b03813516906020013561068c565b604080519115158252519081900360200190f35b6102de600480360360408110156102bd57600080fd5b506fffffffffffffffffffffffffffffffff813581169160200135166107a2565b005b6102e86108ae565b60408051918252519081900360200190f35b6103026108b4565b604080516001600160a01b039092168252519081900360200190f35b6102de6004803603602081101561033457600080fd5b503515156108c3565b6102936004803603606081101561035357600080fd5b506001600160a01b038135811691602081013590911690604001356109ab565b6102e86004803603602081101561038957600080fd5b50356001600160a01b0316610b59565b6102de600480360360208110156103af57600080fd5b50356001600160a01b0316610b6b565b6103c7610cda565b6040805160ff9092168252519081900360200190f35b6102de610ce3565b610293610d6b565b6102de6004803603602081101561040357600080fd5b50356001600160a01b0316610d7b565b610302610eed565b6102de6004803603602081101561043157600080fd5b50356001600160a01b0316610efc565b610302611083565b6102e86004803603602081101561045f57600080fd5b50356001600160a01b0316611092565b6102936004803603606081101561048557600080fd5b506001600160a01b038135811691602081013590911690604001356110ad565b610302600480360360208110156104bb57600080fd5b5035611181565b6102e8611194565b61030261119a565b6101f26111a9565b6102de600480360360208110156104f057600080fd5b5035611204565b6102936112f8565b6102936004803603602081101561051557600080fd5b5035611308565b6102936004803603604081101561053257600080fd5b506001600160a01b038135169060200135611516565b6102936116c5565b6102936004803603602081101561056657600080fd5b50356001600160a01b0316611812565b6102e86004803603604081101561058c57600080fd5b506001600160a01b0381358116916020013516611825565b6105ac611850565b604080516fffffffffffffffffffffffffffffffff938416815291909216602082015281519081900390910190f35b6102de600480360360208110156105f157600080fd5b50356001600160a01b0316611873565b6002805460408051602060018416156101000260001901909316849004601f810184900484028201840190925281815292918301828280156106845780601f1061065957610100808354040283529160200191610684565b820191906000526020600020905b81548152906001019060200180831161066757829003601f168201915b505050505081565b600154600090600160a01b900460ff16156106d85760405162461bcd60e51b815260040180806020018281038252602d815260200180612009602d913960400191505060405180910390fd5b6106e960053363ffffffff61195816565b61073a576040805162461bcd60e51b815260206004820152601560248201527f4d7573742062652061207368617265686f6c6465720000000000000000000000604482015290519081900360640190fd5b336000818152600c602090815260408083206001600160a01b03881680855290835292819020869055805186815290519293927f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925929181900390910190a35060015b92915050565b6009546001600160a01b031633146107f7576040805162461bcd60e51b815260206004820152601360248201527213db9b1e481a5cdcdd595c88185b1b1bddd959606a1b604482015290519081900360640190fd5b600154600160a01b900460ff16610846576040805162461bcd60e51b815260206004820152600e60248201526d135d5cdd081899481b1bd8dad95960921b604482015290519081900360640190fd5b604080518082019091526fffffffffffffffffffffffffffffffff9283168082529183166020909101819052600d80547fffffffffffffffffffffffffffffffff0000000000000000000000000000000016909217909216600160801b909202919091179055565b600b5490565b6009546001600160a01b031681565b6000546001600160a01b03163314806108ec57506001546000546001600160a01b039081169116145b61093d576040805162461bcd60e51b815260206004820152601960248201527f4f776e6572206163636f756e7420697320726571756972656400000000000000604482015290519081900360640190fd5b60015460ff600160a01b909104161515811515141561098d5760405162461bcd60e51b8152600401808060200182810382526028815260200180611fe16028913960400191505060405180910390fd5b60018054911515600160a01b0260ff60a01b19909216919091179055565b600154600090600160a01b900460ff16156109f75760405162461bcd60e51b815260040180806020018281038252602d815260200180612009602d913960400191505060405180910390fd5b6001600160a01b0384166000908152600a602052604090205484908390811115610a5d576040805162461bcd60e51b8152602060048201526012602482015271496e73756666696369656e742066756e647360701b604482015290519081900360640190fd5b6001600160a01b0386166000908152600c60209081526040808320338452909152902054841115610ad5576040805162461bcd60e51b815260206004820152601a60248201527f5472616e73666572206578636565647320616c6c6f77616e6365000000000000604482015290519081900360640190fd5b6000610ae4878787600061198f565b90508015610b4f576001600160a01b0387166000908152600c60209081526040808320338452909152902054610b20908663ffffffff611add16565b6001600160a01b0388166000908152600c60209081526040808320338452909152902055610b4f878787611b3a565b9695505050505050565b600a6020526000908152604090205481565b6009546001600160a01b03163314610bc0576040805162461bcd60e51b815260206004820152601360248201527213db9b1e481a5cdcdd595c88185b1b1bddd959606a1b604482015290519081900360640190fd5b600154600160a01b900460ff16610c0f576040805162461bcd60e51b815260206004820152600e60248201526d135d5cdd081899481b1bd8dad95960921b604482015290519081900360640190fd5b6001600160a01b0381166000908152600a602052604090205480610c7a576040805162461bcd60e51b815260206004820152601c60248201527f41646472657373206973206e6f742061207368617265686f6c64657200000000604482015290519081900360640190fd5b600d54610cba906fffffffffffffffffffffffffffffffff600160801b8204811691610cae9185911663ffffffff611cdd16565b9063ffffffff611d0b16565b6001600160a01b039092166000908152600a602052604090209190915550565b60045460ff1681565b6000546001600160a01b0316331480610d0c57506001546000546001600160a01b039081169116145b610d5d576040805162461bcd60e51b815260206004820152601960248201527f4f776e6572206163636f756e7420697320726571756972656400000000000000604482015290519081900360640190fd5b6000546001600160a01b0316ff5b600954600160a01b900460ff1681565b6000546001600160a01b0316331480610da457506001546000546001600160a01b039081169116145b610df5576040805162461bcd60e51b815260206004820152601960248201527f4f776e6572206163636f756e7420697320726571756972656400000000000000604482015290519081900360640190fd5b6000546001600160a01b0382811691161415610e425760405162461bcd60e51b8152600401808060200182810382526025815260200180611fbc6025913960400191505060405180910390fd5b6001600160a01b038116610e9d576040805162461bcd60e51b815260206004820181905260248201527f4e6577204f776e65722063616e6e6f74206265207a65726f2061646472657373604482015290519081900360640190fd5b600080546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8934ce4adea8d9ce0d714d2c22b86790e41b7731c84b926fbbdc1d40ff6533c99190a35050565b6001546001600160a01b031681565b600154600160a01b900460ff1615610f455760405162461bcd60e51b815260040180806020018281038252602d815260200180612009602d913960400191505060405180910390fd5b6000546001600160a01b0316331480610f6e57506001546000546001600160a01b039081169116145b610fbf576040805162461bcd60e51b815260206004820152601960248201527f4f776e6572206163636f756e7420697320726571756972656400000000000000604482015290519081900360640190fd5b610fd060058263ffffffff61195816565b1561100c5760405162461bcd60e51b8152600401808060200182810382526021815260200180611f9b6021913960400191505060405180910390fd5b6009546001600160a01b03166000818152600a602052604090205461103391908390611b3a565b600980546001600160a01b0319166001600160a01b0383811691821792839055604051919216907ff7189b85d7899f5a32d733e6584c4f1dcdff0274f09d969d186c1797673ede1190600090a350565b6008546001600160a01b031681565b6001600160a01b03166000908152600a602052604090205490565b600154600090600160a01b900460ff16156110f95760405162461bcd60e51b815260040180806020018281038252602d815260200180612009602d913960400191505060405180910390fd5b6001600160a01b0384166000908152600a60205260409020548490839081111561115f576040805162461bcd60e51b8152602060048201526012602482015271496e73756666696369656e742066756e647360701b604482015290519081900360640190fd5b600061116e878787600161198f565b90508015610b4f57610b4f878787611b3a565b600061079c60058363ffffffff611d2a16565b60055481565b6000546001600160a01b031681565b6003805460408051602060026001851615610100026000190190941693909304601f810184900484028201840190925281815292918301828280156106845780601f1061065957610100808354040283529160200191610684565b6009546001600160a01b03163314611259576040805162461bcd60e51b815260206004820152601360248201527213db9b1e481a5cdcdd595c88185b1b1bddd959606a1b604482015290519081900360640190fd5b600154600160a01b900460ff166112a8576040805162461bcd60e51b815260206004820152600e60248201526d135d5cdd081899481b1bd8dad95960921b604482015290519081900360640190fd5b600d54600b546112f2916fffffffffffffffffffffffffffffffff600160801b8204811692610cae9291909116906112e6908663ffffffff611add16565b9063ffffffff611cdd16565b600b5550565b600154600160a01b900460ff1681565b600154600090600160a01b900460ff16156113545760405162461bcd60e51b815260040180806020018281038252602d815260200180612009602d913960400191505060405180910390fd5b6009546001600160a01b031633146113a9576040805162461bcd60e51b815260206004820152601360248201527213db9b1e481a5cdcdd595c88185b1b1bddd959606a1b604482015290519081900360640190fd5b600954600160a01b900460ff1615611408576040805162461bcd60e51b815260206004820152601960248201527f49737375616e636520616c72656164792066696e697368656400000000000000604482015290519081900360640190fd5b811561148657600b54611421908363ffffffff611db016565b600b556009546001600160a01b03166000908152600a602052604090205461144f908363ffffffff611db016565b600980546001600160a01b039081166000908152600a602052604090209290925554611484916005911663ffffffff611e0a16565b505b6009546040805184815290516001600160a01b03909216917f9cb9c14f7bc76e3a89b796b091850526236115352a198b1e472f00e91376bbcb9181900360200190a26009546040805184815290516001600160a01b03909216916000917fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef919081900360200190a3506001919050565b600154600090600160a01b900460ff16156115625760405162461bcd60e51b815260040180806020018281038252602d815260200180612009602d913960400191505060405180910390fd5b336000818152600a602052604090205483908111156115bd576040805162461bcd60e51b8152602060048201526012602482015271496e73756666696369656e742066756e647360701b604482015290519081900360640190fd5b6009546000906001600160a01b031633141561169b57506008546001600160a01b0316158061169657600854600954604080517ffd8258bd0000000000000000000000000000000000000000000000000000000081526001600160a01b039283166004820181905260248201528983166044820152606481018990529051919092169163fd8258bd9160848083019260209291908290030181600087803b15801561166757600080fd5b505af115801561167b573d6000803e3d6000fd5b505050506040513d602081101561169157600080fd5b505190505b6116ab565b6116a8338787600061198f565b90505b80156116bc576116bc338787611b3a565b95945050505050565b600154600090600160a01b900460ff16156117115760405162461bcd60e51b815260040180806020018281038252602d815260200180612009602d913960400191505060405180910390fd5b6009546001600160a01b03163314611766576040805162461bcd60e51b815260206004820152601360248201527213db9b1e481a5cdcdd595c88185b1b1bddd959606a1b604482015290519081900360640190fd5b600954600160a01b900460ff16156117c5576040805162461bcd60e51b815260206004820152601960248201527f49737375616e636520616c72656164792066696e697368656400000000000000604482015290519081900360640190fd5b6009805460ff60a01b1916600160a01b1790556040517f29fe76cc5ca143e91eadf7242fda487fcef09318c1237900f958abe1e2c5beff90600090a150600954600160a01b900460ff1690565b600061079c60058363ffffffff61195816565b6001600160a01b039182166000908152600c6020908152604080832093909416825291909152205490565b600d546fffffffffffffffffffffffffffffffff80821691600160801b90041682565b600154600160a01b900460ff16156118bc5760405162461bcd60e51b815260040180806020018281038252602d815260200180612009602d913960400191505060405180910390fd5b6000546001600160a01b03163314806118e557506001546000546001600160a01b039081169116145b611936576040805162461bcd60e51b815260206004820152601960248201527f4f776e6572206163636f756e7420697320726571756972656400000000000000604482015290519081900360640190fd5b600880546001600160a01b0319166001600160a01b0392909216919091179055565b6001600160a01b0381166000908152600183016020526040812054600019018181128015906119875750835481125b949350505050565b6008546000906001600160a01b03166119aa57508015611987565b8115611a6057600854604080517f5acba2010000000000000000000000000000000000000000000000000000000081523360048201526001600160a01b03888116602483015287811660448301526064820187905291519190921691635acba2019160848083019260209291908290030181600087803b158015611a2d57600080fd5b505af1158015611a41573d6000803e3d6000fd5b505050506040513d6020811015611a5757600080fd5b50519050611987565b600854604080517f6d62a4fe0000000000000000000000000000000000000000000000000000000081523360048201526001600160a01b03888116602483015287811660448301526064820187905291519190921691636d62a4fe9160848083019260209291908290030181600087803b158015611a2d57600080fd5b600082821115611b34576040805162461bcd60e51b815260206004820152601460248201527f526573756c747320696e20756e646572666c6f77000000000000000000000000604482015290519081900360640190fd5b50900390565b6001600160a01b0383166000908152600a6020526040902054611b63908263ffffffff611add16565b6001600160a01b038085166000908152600a60205260408082209390935590841681522054611b98908263ffffffff611db016565b6001600160a01b038084166000818152600a602090815260409182902094909455805185815290519193928716927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef92918290030190a36001600160a01b0382166000908152600a602052604090205415801590611c225750611c2260058363ffffffff611e0a16565b15611c6457604080516001600160a01b038416815290517f3082c1c4977de80c4f67ee77b56b282610ec93a7ecbcc31b551e0ac2f7bd03959181900360200190a15b6001600160a01b0383166000908152600a6020526040902054158015611c965750611c9660058463ffffffff611eaf16565b15611cd857604080516001600160a01b038516815290517f7ba114ff3d9844510a088eea94cd35562e7c97a2d36a418b37b2e61e5c77affe9181900360200190a15b505050565b600082611cec5750600061079c565b82820282848281611cf957fe5b0414611d0457600080fd5b9392505050565b6000808211611d1957600080fd5b818381611d2257fe5b049392505050565b6000808212158015611d3c5750825482125b611d8d576040805162461bcd60e51b815260206004820152601860248201527f496e646578206f757473696465206f6620626f756e64732e0000000000000000604482015290519081900360640190fd5b50600101600090815260029190910160205260409020546001600160a01b031690565b600082820183811015611d04576040805162461bcd60e51b815260206004820152601360248201527f526573756c747320696e206f766572666c6f7700000000000000000000000000604482015290519081900360640190fd5b60006001600160a01b038216611e225750600061079c565b6001600160a01b038216600090815260018401602052604081205460001901908112801590611e515750835481125b15611e6057600091505061079c565b5050815460019081018084556001600160a01b0383166000818152838601602090815260408083208590559382526002870190529190912080546001600160a01b031916909117905592915050565b6001600160a01b03811660009081526001808401602052604082205490811280611ed95750835481135b15611ee857600091505061079c565b8354811215611f4f57835460009081526002850160208181526040808420546001600160a01b03168085526001890183528185208690558585529290915280832080546001600160a01b031990811690931790558654835290912080549091169055611f6e565b6000818152600285016020526040902080546001600160a01b03191690555b50506001600160a01b03166000908152600182810160205260408220919091558154600019019091559056fe4e6577206973737565722063616e27742062652061207368617265686f6c6465724e6577204f776e65722063616e6e6f74206265207468652063757272656e74206f776e6572436f6e747261637420616c726561647920696e20726571756573746564206c6f636b207374617465436f6e74726163742069732063757272656e746c79206c6f636b656420666f72206d6f64696669636174696f6ea265627a7a723158204d9822d8cf82a69b9859203ba0157cd4dc42f143ff22099dad344231b2aff27364736f6c634300050b0032`

// DeployT0kenSplittable deploys a new Ethereum contract, binding an instance of T0kenSplittable to it.
func DeployT0kenSplittable(auth *bind.TransactOpts, backend bind.ContractBackend, tokenName string, tokenSymbol string) (common.Address, *types.Transaction, *T0kenSplittable, error) {
	parsed, err := abi.JSON(strings.NewReader(T0kenSplittableABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(T0kenSplittableBin), backend, tokenName, tokenSymbol)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &T0kenSplittable{T0kenSplittableCaller: T0kenSplittableCaller{contract: contract}, T0kenSplittableTransactor: T0kenSplittableTransactor{contract: contract}, T0kenSplittableFilterer: T0kenSplittableFilterer{contract: contract}}, nil
}

// T0kenSplittable is an auto generated Go binding around an Ethereum contract.
type T0kenSplittable struct {
	T0kenSplittableCaller     // Read-only binding to the contract
	T0kenSplittableTransactor // Write-only binding to the contract
	T0kenSplittableFilterer   // Log filterer for contract events
}

// T0kenSplittableCaller is an auto generated read-only Go binding around an Ethereum contract.
type T0kenSplittableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// T0kenSplittableTransactor is an auto generated write-only Go binding around an Ethereum contract.
type T0kenSplittableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// T0kenSplittableFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type T0kenSplittableFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// T0kenSplittableSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type T0kenSplittableSession struct {
	Contract     *T0kenSplittable  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// T0kenSplittableCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type T0kenSplittableCallerSession struct {
	Contract *T0kenSplittableCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// T0kenSplittableTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type T0kenSplittableTransactorSession struct {
	Contract     *T0kenSplittableTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// T0kenSplittableRaw is an auto generated low-level Go binding around an Ethereum contract.
type T0kenSplittableRaw struct {
	Contract *T0kenSplittable // Generic contract binding to access the raw methods on
}

// T0kenSplittableCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type T0kenSplittableCallerRaw struct {
	Contract *T0kenSplittableCaller // Generic read-only contract binding to access the raw methods on
}

// T0kenSplittableTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type T0kenSplittableTransactorRaw struct {
	Contract *T0kenSplittableTransactor // Generic write-only contract binding to access the raw methods on
}

// NewT0kenSplittable creates a new instance of T0kenSplittable, bound to a specific deployed contract.
func NewT0kenSplittable(address common.Address, backend bind.ContractBackend) (*T0kenSplittable, error) {
	contract, err := bindT0kenSplittable(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &T0kenSplittable{T0kenSplittableCaller: T0kenSplittableCaller{contract: contract}, T0kenSplittableTransactor: T0kenSplittableTransactor{contract: contract}, T0kenSplittableFilterer: T0kenSplittableFilterer{contract: contract}}, nil
}

// NewT0kenSplittableCaller creates a new read-only instance of T0kenSplittable, bound to a specific deployed contract.
func NewT0kenSplittableCaller(address common.Address, caller bind.ContractCaller) (*T0kenSplittableCaller, error) {
	contract, err := bindT0kenSplittable(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &T0kenSplittableCaller{contract: contract}, nil
}

// NewT0kenSplittableTransactor creates a new write-only instance of T0kenSplittable, bound to a specific deployed contract.
func NewT0kenSplittableTransactor(address common.Address, transactor bind.ContractTransactor) (*T0kenSplittableTransactor, error) {
	contract, err := bindT0kenSplittable(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &T0kenSplittableTransactor{contract: contract}, nil
}

// NewT0kenSplittableFilterer creates a new log filterer instance of T0kenSplittable, bound to a specific deployed contract.
func NewT0kenSplittableFilterer(address common.Address, filterer bind.ContractFilterer) (*T0kenSplittableFilterer, error) {
	contract, err := bindT0kenSplittable(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &T0kenSplittableFilterer{contract: contract}, nil
}

// bindT0kenSplittable binds a generic wrapper to an already deployed contract.
func bindT0kenSplittable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(T0kenSplittableABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_T0kenSplittable *T0kenSplittableRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _T0kenSplittable.Contract.T0kenSplittableCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_T0kenSplittable *T0kenSplittableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _T0kenSplittable.Contract.T0kenSplittableTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_T0kenSplittable *T0kenSplittableRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _T0kenSplittable.Contract.T0kenSplittableTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_T0kenSplittable *T0kenSplittableCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _T0kenSplittable.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_T0kenSplittable *T0kenSplittableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _T0kenSplittable.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_T0kenSplittable *T0kenSplittableTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _T0kenSplittable.Contract.contract.Transact(opts, method, params...)
}

// ZEROADDRESS is a free data retrieval call binding the contract method 0x538ba4f9.
//
// Solidity: function ZERO_ADDRESS() constant returns(address)
func (_T0kenSplittable *T0kenSplittableCaller) ZEROADDRESS(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _T0kenSplittable.contract.Call(opts, out, "ZERO_ADDRESS")
	return *ret0, err
}

// ZEROADDRESS is a free data retrieval call binding the contract method 0x538ba4f9.
//
// Solidity: function ZERO_ADDRESS() constant returns(address)
func (_T0kenSplittable *T0kenSplittableSession) ZEROADDRESS() (common.Address, error) {
	return _T0kenSplittable.Contract.ZEROADDRESS(&_T0kenSplittable.CallOpts)
}

// ZEROADDRESS is a free data retrieval call binding the contract method 0x538ba4f9.
//
// Solidity: function ZERO_ADDRESS() constant returns(address)
func (_T0kenSplittable *T0kenSplittableCallerSession) ZEROADDRESS() (common.Address, error) {
	return _T0kenSplittable.Contract.ZEROADDRESS(&_T0kenSplittable.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(addrOwner address, spender address) constant returns(uint256)
func (_T0kenSplittable *T0kenSplittableCaller) Allowance(opts *bind.CallOpts, addrOwner common.Address, spender common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _T0kenSplittable.contract.Call(opts, out, "allowance", addrOwner, spender)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(addrOwner address, spender address) constant returns(uint256)
func (_T0kenSplittable *T0kenSplittableSession) Allowance(addrOwner common.Address, spender common.Address) (*big.Int, error) {
	return _T0kenSplittable.Contract.Allowance(&_T0kenSplittable.CallOpts, addrOwner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(addrOwner address, spender address) constant returns(uint256)
func (_T0kenSplittable *T0kenSplittableCallerSession) Allowance(addrOwner common.Address, spender common.Address) (*big.Int, error) {
	return _T0kenSplittable.Contract.Allowance(&_T0kenSplittable.CallOpts, addrOwner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(addr address) constant returns(uint256)
func (_T0kenSplittable *T0kenSplittableCaller) BalanceOf(opts *bind.CallOpts, addr common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _T0kenSplittable.contract.Call(opts, out, "balanceOf", addr)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(addr address) constant returns(uint256)
func (_T0kenSplittable *T0kenSplittableSession) BalanceOf(addr common.Address) (*big.Int, error) {
	return _T0kenSplittable.Contract.BalanceOf(&_T0kenSplittable.CallOpts, addr)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(addr address) constant returns(uint256)
func (_T0kenSplittable *T0kenSplittableCallerSession) BalanceOf(addr common.Address) (*big.Int, error) {
	return _T0kenSplittable.Contract.BalanceOf(&_T0kenSplittable.CallOpts, addr)
}

// Balances is a free data retrieval call binding the contract method 0x27e235e3.
//
// Solidity: function balances( address) constant returns(uint256)
func (_T0kenSplittable *T0kenSplittableCaller) Balances(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _T0kenSplittable.contract.Call(opts, out, "balances", arg0)
	return *ret0, err
}

// Balances is a free data retrieval call binding the contract method 0x27e235e3.
//
// Solidity: function balances( address) constant returns(uint256)
func (_T0kenSplittable *T0kenSplittableSession) Balances(arg0 common.Address) (*big.Int, error) {
	return _T0kenSplittable.Contract.Balances(&_T0kenSplittable.CallOpts, arg0)
}

// Balances is a free data retrieval call binding the contract method 0x27e235e3.
//
// Solidity: function balances( address) constant returns(uint256)
func (_T0kenSplittable *T0kenSplittableCallerSession) Balances(arg0 common.Address) (*big.Int, error) {
	return _T0kenSplittable.Contract.Balances(&_T0kenSplittable.CallOpts, arg0)
}

// Compliance is a free data retrieval call binding the contract method 0x6290865d.
//
// Solidity: function compliance() constant returns(address)
func (_T0kenSplittable *T0kenSplittableCaller) Compliance(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _T0kenSplittable.contract.Call(opts, out, "compliance")
	return *ret0, err
}

// Compliance is a free data retrieval call binding the contract method 0x6290865d.
//
// Solidity: function compliance() constant returns(address)
func (_T0kenSplittable *T0kenSplittableSession) Compliance() (common.Address, error) {
	return _T0kenSplittable.Contract.Compliance(&_T0kenSplittable.CallOpts)
}

// Compliance is a free data retrieval call binding the contract method 0x6290865d.
//
// Solidity: function compliance() constant returns(address)
func (_T0kenSplittable *T0kenSplittableCallerSession) Compliance() (common.Address, error) {
	return _T0kenSplittable.Contract.Compliance(&_T0kenSplittable.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_T0kenSplittable *T0kenSplittableCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _T0kenSplittable.contract.Call(opts, out, "decimals")
	return *ret0, err
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_T0kenSplittable *T0kenSplittableSession) Decimals() (uint8, error) {
	return _T0kenSplittable.Contract.Decimals(&_T0kenSplittable.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_T0kenSplittable *T0kenSplittableCallerSession) Decimals() (uint8, error) {
	return _T0kenSplittable.Contract.Decimals(&_T0kenSplittable.CallOpts)
}

// HolderAt is a free data retrieval call binding the contract method 0x8082a929.
//
// Solidity: function holderAt(index int256) constant returns(address)
func (_T0kenSplittable *T0kenSplittableCaller) HolderAt(opts *bind.CallOpts, index *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _T0kenSplittable.contract.Call(opts, out, "holderAt", index)
	return *ret0, err
}

// HolderAt is a free data retrieval call binding the contract method 0x8082a929.
//
// Solidity: function holderAt(index int256) constant returns(address)
func (_T0kenSplittable *T0kenSplittableSession) HolderAt(index *big.Int) (common.Address, error) {
	return _T0kenSplittable.Contract.HolderAt(&_T0kenSplittable.CallOpts, index)
}

// HolderAt is a free data retrieval call binding the contract method 0x8082a929.
//
// Solidity: function holderAt(index int256) constant returns(address)
func (_T0kenSplittable *T0kenSplittableCallerSession) HolderAt(index *big.Int) (common.Address, error) {
	return _T0kenSplittable.Contract.HolderAt(&_T0kenSplittable.CallOpts, index)
}

// Holders is a free data retrieval call binding the contract method 0x8188f71c.
//
// Solidity: function holders() constant returns(count int256)
func (_T0kenSplittable *T0kenSplittableCaller) Holders(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _T0kenSplittable.contract.Call(opts, out, "holders")
	return *ret0, err
}

// Holders is a free data retrieval call binding the contract method 0x8188f71c.
//
// Solidity: function holders() constant returns(count int256)
func (_T0kenSplittable *T0kenSplittableSession) Holders() (*big.Int, error) {
	return _T0kenSplittable.Contract.Holders(&_T0kenSplittable.CallOpts)
}

// Holders is a free data retrieval call binding the contract method 0x8188f71c.
//
// Solidity: function holders() constant returns(count int256)
func (_T0kenSplittable *T0kenSplittableCallerSession) Holders() (*big.Int, error) {
	return _T0kenSplittable.Contract.Holders(&_T0kenSplittable.CallOpts)
}

// IsHolder is a free data retrieval call binding the contract method 0xd4d7b19a.
//
// Solidity: function isHolder(addr address) constant returns(bool)
func (_T0kenSplittable *T0kenSplittableCaller) IsHolder(opts *bind.CallOpts, addr common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _T0kenSplittable.contract.Call(opts, out, "isHolder", addr)
	return *ret0, err
}

// IsHolder is a free data retrieval call binding the contract method 0xd4d7b19a.
//
// Solidity: function isHolder(addr address) constant returns(bool)
func (_T0kenSplittable *T0kenSplittableSession) IsHolder(addr common.Address) (bool, error) {
	return _T0kenSplittable.Contract.IsHolder(&_T0kenSplittable.CallOpts, addr)
}

// IsHolder is a free data retrieval call binding the contract method 0xd4d7b19a.
//
// Solidity: function isHolder(addr address) constant returns(bool)
func (_T0kenSplittable *T0kenSplittableCallerSession) IsHolder(addr common.Address) (bool, error) {
	return _T0kenSplittable.Contract.IsHolder(&_T0kenSplittable.CallOpts, addr)
}

// IsLocked is a free data retrieval call binding the contract method 0xa4e2d634.
//
// Solidity: function isLocked() constant returns(bool)
func (_T0kenSplittable *T0kenSplittableCaller) IsLocked(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _T0kenSplittable.contract.Call(opts, out, "isLocked")
	return *ret0, err
}

// IsLocked is a free data retrieval call binding the contract method 0xa4e2d634.
//
// Solidity: function isLocked() constant returns(bool)
func (_T0kenSplittable *T0kenSplittableSession) IsLocked() (bool, error) {
	return _T0kenSplittable.Contract.IsLocked(&_T0kenSplittable.CallOpts)
}

// IsLocked is a free data retrieval call binding the contract method 0xa4e2d634.
//
// Solidity: function isLocked() constant returns(bool)
func (_T0kenSplittable *T0kenSplittableCallerSession) IsLocked() (bool, error) {
	return _T0kenSplittable.Contract.IsLocked(&_T0kenSplittable.CallOpts)
}

// IssuanceFinished is a free data retrieval call binding the contract method 0x4662299a.
//
// Solidity: function issuanceFinished() constant returns(bool)
func (_T0kenSplittable *T0kenSplittableCaller) IssuanceFinished(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _T0kenSplittable.contract.Call(opts, out, "issuanceFinished")
	return *ret0, err
}

// IssuanceFinished is a free data retrieval call binding the contract method 0x4662299a.
//
// Solidity: function issuanceFinished() constant returns(bool)
func (_T0kenSplittable *T0kenSplittableSession) IssuanceFinished() (bool, error) {
	return _T0kenSplittable.Contract.IssuanceFinished(&_T0kenSplittable.CallOpts)
}

// IssuanceFinished is a free data retrieval call binding the contract method 0x4662299a.
//
// Solidity: function issuanceFinished() constant returns(bool)
func (_T0kenSplittable *T0kenSplittableCallerSession) IssuanceFinished() (bool, error) {
	return _T0kenSplittable.Contract.IssuanceFinished(&_T0kenSplittable.CallOpts)
}

// Issuer is a free data retrieval call binding the contract method 0x1d143848.
//
// Solidity: function issuer() constant returns(address)
func (_T0kenSplittable *T0kenSplittableCaller) Issuer(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _T0kenSplittable.contract.Call(opts, out, "issuer")
	return *ret0, err
}

// Issuer is a free data retrieval call binding the contract method 0x1d143848.
//
// Solidity: function issuer() constant returns(address)
func (_T0kenSplittable *T0kenSplittableSession) Issuer() (common.Address, error) {
	return _T0kenSplittable.Contract.Issuer(&_T0kenSplittable.CallOpts)
}

// Issuer is a free data retrieval call binding the contract method 0x1d143848.
//
// Solidity: function issuer() constant returns(address)
func (_T0kenSplittable *T0kenSplittableCallerSession) Issuer() (common.Address, error) {
	return _T0kenSplittable.Contract.Issuer(&_T0kenSplittable.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_T0kenSplittable *T0kenSplittableCaller) Name(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _T0kenSplittable.contract.Call(opts, out, "name")
	return *ret0, err
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_T0kenSplittable *T0kenSplittableSession) Name() (string, error) {
	return _T0kenSplittable.Contract.Name(&_T0kenSplittable.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_T0kenSplittable *T0kenSplittableCallerSession) Name() (string, error) {
	return _T0kenSplittable.Contract.Name(&_T0kenSplittable.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_T0kenSplittable *T0kenSplittableCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _T0kenSplittable.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_T0kenSplittable *T0kenSplittableSession) Owner() (common.Address, error) {
	return _T0kenSplittable.Contract.Owner(&_T0kenSplittable.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_T0kenSplittable *T0kenSplittableCallerSession) Owner() (common.Address, error) {
	return _T0kenSplittable.Contract.Owner(&_T0kenSplittable.CallOpts)
}

// Split is a free data retrieval call binding the contract method 0xf7654176.
//
// Solidity: function split() constant returns(numerator uint128, denominator uint128)
func (_T0kenSplittable *T0kenSplittableCaller) Split(opts *bind.CallOpts) (struct {
	Numerator   *big.Int
	Denominator *big.Int
}, error) {
	ret := new(struct {
		Numerator   *big.Int
		Denominator *big.Int
	})
	out := ret
	err := _T0kenSplittable.contract.Call(opts, out, "split")
	return *ret, err
}

// Split is a free data retrieval call binding the contract method 0xf7654176.
//
// Solidity: function split() constant returns(numerator uint128, denominator uint128)
func (_T0kenSplittable *T0kenSplittableSession) Split() (struct {
	Numerator   *big.Int
	Denominator *big.Int
}, error) {
	return _T0kenSplittable.Contract.Split(&_T0kenSplittable.CallOpts)
}

// Split is a free data retrieval call binding the contract method 0xf7654176.
//
// Solidity: function split() constant returns(numerator uint128, denominator uint128)
func (_T0kenSplittable *T0kenSplittableCallerSession) Split() (struct {
	Numerator   *big.Int
	Denominator *big.Int
}, error) {
	return _T0kenSplittable.Contract.Split(&_T0kenSplittable.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_T0kenSplittable *T0kenSplittableCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _T0kenSplittable.contract.Call(opts, out, "symbol")
	return *ret0, err
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_T0kenSplittable *T0kenSplittableSession) Symbol() (string, error) {
	return _T0kenSplittable.Contract.Symbol(&_T0kenSplittable.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_T0kenSplittable *T0kenSplittableCallerSession) Symbol() (string, error) {
	return _T0kenSplittable.Contract.Symbol(&_T0kenSplittable.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_T0kenSplittable *T0kenSplittableCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _T0kenSplittable.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_T0kenSplittable *T0kenSplittableSession) TotalSupply() (*big.Int, error) {
	return _T0kenSplittable.Contract.TotalSupply(&_T0kenSplittable.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_T0kenSplittable *T0kenSplittableCallerSession) TotalSupply() (*big.Int, error) {
	return _T0kenSplittable.Contract.TotalSupply(&_T0kenSplittable.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(spender address, tokens uint256) returns(bool)
func (_T0kenSplittable *T0kenSplittableTransactor) Approve(opts *bind.TransactOpts, spender common.Address, tokens *big.Int) (*types.Transaction, error) {
	return _T0kenSplittable.contract.Transact(opts, "approve", spender, tokens)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(spender address, tokens uint256) returns(bool)
func (_T0kenSplittable *T0kenSplittableSession) Approve(spender common.Address, tokens *big.Int) (*types.Transaction, error) {
	return _T0kenSplittable.Contract.Approve(&_T0kenSplittable.TransactOpts, spender, tokens)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(spender address, tokens uint256) returns(bool)
func (_T0kenSplittable *T0kenSplittableTransactorSession) Approve(spender common.Address, tokens *big.Int) (*types.Transaction, error) {
	return _T0kenSplittable.Contract.Approve(&_T0kenSplittable.TransactOpts, spender, tokens)
}

// FinishIssuance is a paid mutator transaction binding the contract method 0xc422293b.
//
// Solidity: function finishIssuance() returns(bool)
func (_T0kenSplittable *T0kenSplittableTransactor) FinishIssuance(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _T0kenSplittable.contract.Transact(opts, "finishIssuance")
}

// FinishIssuance is a paid mutator transaction binding the contract method 0xc422293b.
//
// Solidity: function finishIssuance() returns(bool)
func (_T0kenSplittable *T0kenSplittableSession) FinishIssuance() (*types.Transaction, error) {
	return _T0kenSplittable.Contract.FinishIssuance(&_T0kenSplittable.TransactOpts)
}

// FinishIssuance is a paid mutator transaction binding the contract method 0xc422293b.
//
// Solidity: function finishIssuance() returns(bool)
func (_T0kenSplittable *T0kenSplittableTransactorSession) FinishIssuance() (*types.Transaction, error) {
	return _T0kenSplittable.Contract.FinishIssuance(&_T0kenSplittable.TransactOpts)
}

// IssueTokens is a paid mutator transaction binding the contract method 0xa5820daa.
//
// Solidity: function issueTokens(quantity uint256) returns(bool)
func (_T0kenSplittable *T0kenSplittableTransactor) IssueTokens(opts *bind.TransactOpts, quantity *big.Int) (*types.Transaction, error) {
	return _T0kenSplittable.contract.Transact(opts, "issueTokens", quantity)
}

// IssueTokens is a paid mutator transaction binding the contract method 0xa5820daa.
//
// Solidity: function issueTokens(quantity uint256) returns(bool)
func (_T0kenSplittable *T0kenSplittableSession) IssueTokens(quantity *big.Int) (*types.Transaction, error) {
	return _T0kenSplittable.Contract.IssueTokens(&_T0kenSplittable.TransactOpts, quantity)
}

// IssueTokens is a paid mutator transaction binding the contract method 0xa5820daa.
//
// Solidity: function issueTokens(quantity uint256) returns(bool)
func (_T0kenSplittable *T0kenSplittableTransactorSession) IssueTokens(quantity *big.Int) (*types.Transaction, error) {
	return _T0kenSplittable.Contract.IssueTokens(&_T0kenSplittable.TransactOpts, quantity)
}

// Kill is a paid mutator transaction binding the contract method 0x41c0e1b5.
//
// Solidity: function kill() returns()
func (_T0kenSplittable *T0kenSplittableTransactor) Kill(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _T0kenSplittable.contract.Transact(opts, "kill")
}

// Kill is a paid mutator transaction binding the contract method 0x41c0e1b5.
//
// Solidity: function kill() returns()
func (_T0kenSplittable *T0kenSplittableSession) Kill() (*types.Transaction, error) {
	return _T0kenSplittable.Contract.Kill(&_T0kenSplittable.TransactOpts)
}

// Kill is a paid mutator transaction binding the contract method 0x41c0e1b5.
//
// Solidity: function kill() returns()
func (_T0kenSplittable *T0kenSplittableTransactorSession) Kill() (*types.Transaction, error) {
	return _T0kenSplittable.Contract.Kill(&_T0kenSplittable.TransactOpts)
}

// SetCompliance is a paid mutator transaction binding the contract method 0xf8981789.
//
// Solidity: function setCompliance(newComplianceAddress address) returns()
func (_T0kenSplittable *T0kenSplittableTransactor) SetCompliance(opts *bind.TransactOpts, newComplianceAddress common.Address) (*types.Transaction, error) {
	return _T0kenSplittable.contract.Transact(opts, "setCompliance", newComplianceAddress)
}

// SetCompliance is a paid mutator transaction binding the contract method 0xf8981789.
//
// Solidity: function setCompliance(newComplianceAddress address) returns()
func (_T0kenSplittable *T0kenSplittableSession) SetCompliance(newComplianceAddress common.Address) (*types.Transaction, error) {
	return _T0kenSplittable.Contract.SetCompliance(&_T0kenSplittable.TransactOpts, newComplianceAddress)
}

// SetCompliance is a paid mutator transaction binding the contract method 0xf8981789.
//
// Solidity: function setCompliance(newComplianceAddress address) returns()
func (_T0kenSplittable *T0kenSplittableTransactorSession) SetCompliance(newComplianceAddress common.Address) (*types.Transaction, error) {
	return _T0kenSplittable.Contract.SetCompliance(&_T0kenSplittable.TransactOpts, newComplianceAddress)
}

// SetIssuer is a paid mutator transaction binding the contract method 0x55cc4e57.
//
// Solidity: function setIssuer(newIssuer address) returns()
func (_T0kenSplittable *T0kenSplittableTransactor) SetIssuer(opts *bind.TransactOpts, newIssuer common.Address) (*types.Transaction, error) {
	return _T0kenSplittable.contract.Transact(opts, "setIssuer", newIssuer)
}

// SetIssuer is a paid mutator transaction binding the contract method 0x55cc4e57.
//
// Solidity: function setIssuer(newIssuer address) returns()
func (_T0kenSplittable *T0kenSplittableSession) SetIssuer(newIssuer common.Address) (*types.Transaction, error) {
	return _T0kenSplittable.Contract.SetIssuer(&_T0kenSplittable.TransactOpts, newIssuer)
}

// SetIssuer is a paid mutator transaction binding the contract method 0x55cc4e57.
//
// Solidity: function setIssuer(newIssuer address) returns()
func (_T0kenSplittable *T0kenSplittableTransactorSession) SetIssuer(newIssuer common.Address) (*types.Transaction, error) {
	return _T0kenSplittable.Contract.SetIssuer(&_T0kenSplittable.TransactOpts, newIssuer)
}

// SetLocked is a paid mutator transaction binding the contract method 0x211e28b6.
//
// Solidity: function setLocked(locked bool) returns()
func (_T0kenSplittable *T0kenSplittableTransactor) SetLocked(opts *bind.TransactOpts, locked bool) (*types.Transaction, error) {
	return _T0kenSplittable.contract.Transact(opts, "setLocked", locked)
}

// SetLocked is a paid mutator transaction binding the contract method 0x211e28b6.
//
// Solidity: function setLocked(locked bool) returns()
func (_T0kenSplittable *T0kenSplittableSession) SetLocked(locked bool) (*types.Transaction, error) {
	return _T0kenSplittable.Contract.SetLocked(&_T0kenSplittable.TransactOpts, locked)
}

// SetLocked is a paid mutator transaction binding the contract method 0x211e28b6.
//
// Solidity: function setLocked(locked bool) returns()
func (_T0kenSplittable *T0kenSplittableTransactorSession) SetLocked(locked bool) (*types.Transaction, error) {
	return _T0kenSplittable.Contract.SetLocked(&_T0kenSplittable.TransactOpts, locked)
}

// SetSplitMultiplier is a paid mutator transaction binding the contract method 0x0c1554a1.
//
// Solidity: function setSplitMultiplier(numerator uint128, denominator uint128) returns()
func (_T0kenSplittable *T0kenSplittableTransactor) SetSplitMultiplier(opts *bind.TransactOpts, numerator *big.Int, denominator *big.Int) (*types.Transaction, error) {
	return _T0kenSplittable.contract.Transact(opts, "setSplitMultiplier", numerator, denominator)
}

// SetSplitMultiplier is a paid mutator transaction binding the contract method 0x0c1554a1.
//
// Solidity: function setSplitMultiplier(numerator uint128, denominator uint128) returns()
func (_T0kenSplittable *T0kenSplittableSession) SetSplitMultiplier(numerator *big.Int, denominator *big.Int) (*types.Transaction, error) {
	return _T0kenSplittable.Contract.SetSplitMultiplier(&_T0kenSplittable.TransactOpts, numerator, denominator)
}

// SetSplitMultiplier is a paid mutator transaction binding the contract method 0x0c1554a1.
//
// Solidity: function setSplitMultiplier(numerator uint128, denominator uint128) returns()
func (_T0kenSplittable *T0kenSplittableTransactorSession) SetSplitMultiplier(numerator *big.Int, denominator *big.Int) (*types.Transaction, error) {
	return _T0kenSplittable.Contract.SetSplitMultiplier(&_T0kenSplittable.TransactOpts, numerator, denominator)
}

// SplitBalance is a paid mutator transaction binding the contract method 0x2bbc5019.
//
// Solidity: function splitBalance(addr address) returns()
func (_T0kenSplittable *T0kenSplittableTransactor) SplitBalance(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	return _T0kenSplittable.contract.Transact(opts, "splitBalance", addr)
}

// SplitBalance is a paid mutator transaction binding the contract method 0x2bbc5019.
//
// Solidity: function splitBalance(addr address) returns()
func (_T0kenSplittable *T0kenSplittableSession) SplitBalance(addr common.Address) (*types.Transaction, error) {
	return _T0kenSplittable.Contract.SplitBalance(&_T0kenSplittable.TransactOpts, addr)
}

// SplitBalance is a paid mutator transaction binding the contract method 0x2bbc5019.
//
// Solidity: function splitBalance(addr address) returns()
func (_T0kenSplittable *T0kenSplittableTransactorSession) SplitBalance(addr common.Address) (*types.Transaction, error) {
	return _T0kenSplittable.Contract.SplitBalance(&_T0kenSplittable.TransactOpts, addr)
}

// SplitTotalSupply is a paid mutator transaction binding the contract method 0x98d34b9b.
//
// Solidity: function splitTotalSupply(cashedOutTokenQuantity uint256) returns()
func (_T0kenSplittable *T0kenSplittableTransactor) SplitTotalSupply(opts *bind.TransactOpts, cashedOutTokenQuantity *big.Int) (*types.Transaction, error) {
	return _T0kenSplittable.contract.Transact(opts, "splitTotalSupply", cashedOutTokenQuantity)
}

// SplitTotalSupply is a paid mutator transaction binding the contract method 0x98d34b9b.
//
// Solidity: function splitTotalSupply(cashedOutTokenQuantity uint256) returns()
func (_T0kenSplittable *T0kenSplittableSession) SplitTotalSupply(cashedOutTokenQuantity *big.Int) (*types.Transaction, error) {
	return _T0kenSplittable.Contract.SplitTotalSupply(&_T0kenSplittable.TransactOpts, cashedOutTokenQuantity)
}

// SplitTotalSupply is a paid mutator transaction binding the contract method 0x98d34b9b.
//
// Solidity: function splitTotalSupply(cashedOutTokenQuantity uint256) returns()
func (_T0kenSplittable *T0kenSplittableTransactorSession) SplitTotalSupply(cashedOutTokenQuantity *big.Int) (*types.Transaction, error) {
	return _T0kenSplittable.Contract.SplitTotalSupply(&_T0kenSplittable.TransactOpts, cashedOutTokenQuantity)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(to address, tokens uint256) returns(bool)
func (_T0kenSplittable *T0kenSplittableTransactor) Transfer(opts *bind.TransactOpts, to common.Address, tokens *big.Int) (*types.Transaction, error) {
	return _T0kenSplittable.contract.Transact(opts, "transfer", to, tokens)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(to address, tokens uint256) returns(bool)
func (_T0kenSplittable *T0kenSplittableSession) Transfer(to common.Address, tokens *big.Int) (*types.Transaction, error) {
	return _T0kenSplittable.Contract.Transfer(&_T0kenSplittable.TransactOpts, to, tokens)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(to address, tokens uint256) returns(bool)
func (_T0kenSplittable *T0kenSplittableTransactorSession) Transfer(to common.Address, tokens *big.Int) (*types.Transaction, error) {
	return _T0kenSplittable.Contract.Transfer(&_T0kenSplittable.TransactOpts, to, tokens)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(from address, to address, tokens uint256) returns(bool)
func (_T0kenSplittable *T0kenSplittableTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokens *big.Int) (*types.Transaction, error) {
	return _T0kenSplittable.contract.Transact(opts, "transferFrom", from, to, tokens)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(from address, to address, tokens uint256) returns(bool)
func (_T0kenSplittable *T0kenSplittableSession) TransferFrom(from common.Address, to common.Address, tokens *big.Int) (*types.Transaction, error) {
	return _T0kenSplittable.Contract.TransferFrom(&_T0kenSplittable.TransactOpts, from, to, tokens)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(from address, to address, tokens uint256) returns(bool)
func (_T0kenSplittable *T0kenSplittableTransactorSession) TransferFrom(from common.Address, to common.Address, tokens *big.Int) (*types.Transaction, error) {
	return _T0kenSplittable.Contract.TransferFrom(&_T0kenSplittable.TransactOpts, from, to, tokens)
}

// TransferOverride is a paid mutator transaction binding the contract method 0x80318be8.
//
// Solidity: function transferOverride(from address, to address, tokens uint256) returns(bool)
func (_T0kenSplittable *T0kenSplittableTransactor) TransferOverride(opts *bind.TransactOpts, from common.Address, to common.Address, tokens *big.Int) (*types.Transaction, error) {
	return _T0kenSplittable.contract.Transact(opts, "transferOverride", from, to, tokens)
}

// TransferOverride is a paid mutator transaction binding the contract method 0x80318be8.
//
// Solidity: function transferOverride(from address, to address, tokens uint256) returns(bool)
func (_T0kenSplittable *T0kenSplittableSession) TransferOverride(from common.Address, to common.Address, tokens *big.Int) (*types.Transaction, error) {
	return _T0kenSplittable.Contract.TransferOverride(&_T0kenSplittable.TransactOpts, from, to, tokens)
}

// TransferOverride is a paid mutator transaction binding the contract method 0x80318be8.
//
// Solidity: function transferOverride(from address, to address, tokens uint256) returns(bool)
func (_T0kenSplittable *T0kenSplittableTransactorSession) TransferOverride(from common.Address, to common.Address, tokens *big.Int) (*types.Transaction, error) {
	return _T0kenSplittable.Contract.TransferOverride(&_T0kenSplittable.TransactOpts, from, to, tokens)
}

// TransferOwner is a paid mutator transaction binding the contract method 0x4fb2e45d.
//
// Solidity: function transferOwner(newOwner address) returns()
func (_T0kenSplittable *T0kenSplittableTransactor) TransferOwner(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _T0kenSplittable.contract.Transact(opts, "transferOwner", newOwner)
}

// TransferOwner is a paid mutator transaction binding the contract method 0x4fb2e45d.
//
// Solidity: function transferOwner(newOwner address) returns()
func (_T0kenSplittable *T0kenSplittableSession) TransferOwner(newOwner common.Address) (*types.Transaction, error) {
	return _T0kenSplittable.Contract.TransferOwner(&_T0kenSplittable.TransactOpts, newOwner)
}

// TransferOwner is a paid mutator transaction binding the contract method 0x4fb2e45d.
//
// Solidity: function transferOwner(newOwner address) returns()
func (_T0kenSplittable *T0kenSplittableTransactorSession) TransferOwner(newOwner common.Address) (*types.Transaction, error) {
	return _T0kenSplittable.Contract.TransferOwner(&_T0kenSplittable.TransactOpts, newOwner)
}

// T0kenSplittableApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the T0kenSplittable contract.
type T0kenSplittableApprovalIterator struct {
	Event *T0kenSplittableApproval // Event containing the contract specifics and raw log

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
func (it *T0kenSplittableApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(T0kenSplittableApproval)
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
		it.Event = new(T0kenSplittableApproval)
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
func (it *T0kenSplittableApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *T0kenSplittableApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// T0kenSplittableApproval represents a Approval event raised by the T0kenSplittable contract.
type T0kenSplittableApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: e Approval(owner indexed address, spender indexed address, value uint256)
func (_T0kenSplittable *T0kenSplittableFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*T0kenSplittableApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _T0kenSplittable.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &T0kenSplittableApprovalIterator{contract: _T0kenSplittable.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: e Approval(owner indexed address, spender indexed address, value uint256)
func (_T0kenSplittable *T0kenSplittableFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *T0kenSplittableApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _T0kenSplittable.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(T0kenSplittableApproval)
				if err := _T0kenSplittable.contract.UnpackLog(event, "Approval", log); err != nil {
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

// T0kenSplittableIssuanceIterator is returned from FilterIssuance and is used to iterate over the raw logs and unpacked data for Issuance events raised by the T0kenSplittable contract.
type T0kenSplittableIssuanceIterator struct {
	Event *T0kenSplittableIssuance // Event containing the contract specifics and raw log

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
func (it *T0kenSplittableIssuanceIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(T0kenSplittableIssuance)
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
		it.Event = new(T0kenSplittableIssuance)
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
func (it *T0kenSplittableIssuanceIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *T0kenSplittableIssuanceIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// T0kenSplittableIssuance represents a Issuance event raised by the T0kenSplittable contract.
type T0kenSplittableIssuance struct {
	To     common.Address
	Tokens *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterIssuance is a free log retrieval operation binding the contract event 0x9cb9c14f7bc76e3a89b796b091850526236115352a198b1e472f00e91376bbcb.
//
// Solidity: e Issuance(to indexed address, tokens uint256)
func (_T0kenSplittable *T0kenSplittableFilterer) FilterIssuance(opts *bind.FilterOpts, to []common.Address) (*T0kenSplittableIssuanceIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _T0kenSplittable.contract.FilterLogs(opts, "Issuance", toRule)
	if err != nil {
		return nil, err
	}
	return &T0kenSplittableIssuanceIterator{contract: _T0kenSplittable.contract, event: "Issuance", logs: logs, sub: sub}, nil
}

// WatchIssuance is a free log subscription operation binding the contract event 0x9cb9c14f7bc76e3a89b796b091850526236115352a198b1e472f00e91376bbcb.
//
// Solidity: e Issuance(to indexed address, tokens uint256)
func (_T0kenSplittable *T0kenSplittableFilterer) WatchIssuance(opts *bind.WatchOpts, sink chan<- *T0kenSplittableIssuance, to []common.Address) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _T0kenSplittable.contract.WatchLogs(opts, "Issuance", toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(T0kenSplittableIssuance)
				if err := _T0kenSplittable.contract.UnpackLog(event, "Issuance", log); err != nil {
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

// T0kenSplittableIssuanceFinishedIterator is returned from FilterIssuanceFinished and is used to iterate over the raw logs and unpacked data for IssuanceFinished events raised by the T0kenSplittable contract.
type T0kenSplittableIssuanceFinishedIterator struct {
	Event *T0kenSplittableIssuanceFinished // Event containing the contract specifics and raw log

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
func (it *T0kenSplittableIssuanceFinishedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(T0kenSplittableIssuanceFinished)
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
		it.Event = new(T0kenSplittableIssuanceFinished)
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
func (it *T0kenSplittableIssuanceFinishedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *T0kenSplittableIssuanceFinishedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// T0kenSplittableIssuanceFinished represents a IssuanceFinished event raised by the T0kenSplittable contract.
type T0kenSplittableIssuanceFinished struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterIssuanceFinished is a free log retrieval operation binding the contract event 0x29fe76cc5ca143e91eadf7242fda487fcef09318c1237900f958abe1e2c5beff.
//
// Solidity: e IssuanceFinished()
func (_T0kenSplittable *T0kenSplittableFilterer) FilterIssuanceFinished(opts *bind.FilterOpts) (*T0kenSplittableIssuanceFinishedIterator, error) {

	logs, sub, err := _T0kenSplittable.contract.FilterLogs(opts, "IssuanceFinished")
	if err != nil {
		return nil, err
	}
	return &T0kenSplittableIssuanceFinishedIterator{contract: _T0kenSplittable.contract, event: "IssuanceFinished", logs: logs, sub: sub}, nil
}

// WatchIssuanceFinished is a free log subscription operation binding the contract event 0x29fe76cc5ca143e91eadf7242fda487fcef09318c1237900f958abe1e2c5beff.
//
// Solidity: e IssuanceFinished()
func (_T0kenSplittable *T0kenSplittableFilterer) WatchIssuanceFinished(opts *bind.WatchOpts, sink chan<- *T0kenSplittableIssuanceFinished) (event.Subscription, error) {

	logs, sub, err := _T0kenSplittable.contract.WatchLogs(opts, "IssuanceFinished")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(T0kenSplittableIssuanceFinished)
				if err := _T0kenSplittable.contract.UnpackLog(event, "IssuanceFinished", log); err != nil {
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

// T0kenSplittableIssuerSetIterator is returned from FilterIssuerSet and is used to iterate over the raw logs and unpacked data for IssuerSet events raised by the T0kenSplittable contract.
type T0kenSplittableIssuerSetIterator struct {
	Event *T0kenSplittableIssuerSet // Event containing the contract specifics and raw log

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
func (it *T0kenSplittableIssuerSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(T0kenSplittableIssuerSet)
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
		it.Event = new(T0kenSplittableIssuerSet)
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
func (it *T0kenSplittableIssuerSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *T0kenSplittableIssuerSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// T0kenSplittableIssuerSet represents a IssuerSet event raised by the T0kenSplittable contract.
type T0kenSplittableIssuerSet struct {
	PreviousIssuer common.Address
	NewIssuer      common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterIssuerSet is a free log retrieval operation binding the contract event 0xf7189b85d7899f5a32d733e6584c4f1dcdff0274f09d969d186c1797673ede11.
//
// Solidity: e IssuerSet(previousIssuer indexed address, newIssuer indexed address)
func (_T0kenSplittable *T0kenSplittableFilterer) FilterIssuerSet(opts *bind.FilterOpts, previousIssuer []common.Address, newIssuer []common.Address) (*T0kenSplittableIssuerSetIterator, error) {

	var previousIssuerRule []interface{}
	for _, previousIssuerItem := range previousIssuer {
		previousIssuerRule = append(previousIssuerRule, previousIssuerItem)
	}
	var newIssuerRule []interface{}
	for _, newIssuerItem := range newIssuer {
		newIssuerRule = append(newIssuerRule, newIssuerItem)
	}

	logs, sub, err := _T0kenSplittable.contract.FilterLogs(opts, "IssuerSet", previousIssuerRule, newIssuerRule)
	if err != nil {
		return nil, err
	}
	return &T0kenSplittableIssuerSetIterator{contract: _T0kenSplittable.contract, event: "IssuerSet", logs: logs, sub: sub}, nil
}

// WatchIssuerSet is a free log subscription operation binding the contract event 0xf7189b85d7899f5a32d733e6584c4f1dcdff0274f09d969d186c1797673ede11.
//
// Solidity: e IssuerSet(previousIssuer indexed address, newIssuer indexed address)
func (_T0kenSplittable *T0kenSplittableFilterer) WatchIssuerSet(opts *bind.WatchOpts, sink chan<- *T0kenSplittableIssuerSet, previousIssuer []common.Address, newIssuer []common.Address) (event.Subscription, error) {

	var previousIssuerRule []interface{}
	for _, previousIssuerItem := range previousIssuer {
		previousIssuerRule = append(previousIssuerRule, previousIssuerItem)
	}
	var newIssuerRule []interface{}
	for _, newIssuerItem := range newIssuer {
		newIssuerRule = append(newIssuerRule, newIssuerItem)
	}

	logs, sub, err := _T0kenSplittable.contract.WatchLogs(opts, "IssuerSet", previousIssuerRule, newIssuerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(T0kenSplittableIssuerSet)
				if err := _T0kenSplittable.contract.UnpackLog(event, "IssuerSet", log); err != nil {
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

// T0kenSplittableOwnerTransferredIterator is returned from FilterOwnerTransferred and is used to iterate over the raw logs and unpacked data for OwnerTransferred events raised by the T0kenSplittable contract.
type T0kenSplittableOwnerTransferredIterator struct {
	Event *T0kenSplittableOwnerTransferred // Event containing the contract specifics and raw log

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
func (it *T0kenSplittableOwnerTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(T0kenSplittableOwnerTransferred)
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
		it.Event = new(T0kenSplittableOwnerTransferred)
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
func (it *T0kenSplittableOwnerTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *T0kenSplittableOwnerTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// T0kenSplittableOwnerTransferred represents a OwnerTransferred event raised by the T0kenSplittable contract.
type T0kenSplittableOwnerTransferred struct {
	OldOwner common.Address
	NewOwner common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOwnerTransferred is a free log retrieval operation binding the contract event 0x8934ce4adea8d9ce0d714d2c22b86790e41b7731c84b926fbbdc1d40ff6533c9.
//
// Solidity: e OwnerTransferred(oldOwner indexed address, newOwner indexed address)
func (_T0kenSplittable *T0kenSplittableFilterer) FilterOwnerTransferred(opts *bind.FilterOpts, oldOwner []common.Address, newOwner []common.Address) (*T0kenSplittableOwnerTransferredIterator, error) {

	var oldOwnerRule []interface{}
	for _, oldOwnerItem := range oldOwner {
		oldOwnerRule = append(oldOwnerRule, oldOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _T0kenSplittable.contract.FilterLogs(opts, "OwnerTransferred", oldOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &T0kenSplittableOwnerTransferredIterator{contract: _T0kenSplittable.contract, event: "OwnerTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnerTransferred is a free log subscription operation binding the contract event 0x8934ce4adea8d9ce0d714d2c22b86790e41b7731c84b926fbbdc1d40ff6533c9.
//
// Solidity: e OwnerTransferred(oldOwner indexed address, newOwner indexed address)
func (_T0kenSplittable *T0kenSplittableFilterer) WatchOwnerTransferred(opts *bind.WatchOpts, sink chan<- *T0kenSplittableOwnerTransferred, oldOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var oldOwnerRule []interface{}
	for _, oldOwnerItem := range oldOwner {
		oldOwnerRule = append(oldOwnerRule, oldOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _T0kenSplittable.contract.WatchLogs(opts, "OwnerTransferred", oldOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(T0kenSplittableOwnerTransferred)
				if err := _T0kenSplittable.contract.UnpackLog(event, "OwnerTransferred", log); err != nil {
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

// T0kenSplittableShareholderAddedIterator is returned from FilterShareholderAdded and is used to iterate over the raw logs and unpacked data for ShareholderAdded events raised by the T0kenSplittable contract.
type T0kenSplittableShareholderAddedIterator struct {
	Event *T0kenSplittableShareholderAdded // Event containing the contract specifics and raw log

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
func (it *T0kenSplittableShareholderAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(T0kenSplittableShareholderAdded)
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
		it.Event = new(T0kenSplittableShareholderAdded)
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
func (it *T0kenSplittableShareholderAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *T0kenSplittableShareholderAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// T0kenSplittableShareholderAdded represents a ShareholderAdded event raised by the T0kenSplittable contract.
type T0kenSplittableShareholderAdded struct {
	Shareholder common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterShareholderAdded is a free log retrieval operation binding the contract event 0x3082c1c4977de80c4f67ee77b56b282610ec93a7ecbcc31b551e0ac2f7bd0395.
//
// Solidity: e ShareholderAdded(shareholder address)
func (_T0kenSplittable *T0kenSplittableFilterer) FilterShareholderAdded(opts *bind.FilterOpts) (*T0kenSplittableShareholderAddedIterator, error) {

	logs, sub, err := _T0kenSplittable.contract.FilterLogs(opts, "ShareholderAdded")
	if err != nil {
		return nil, err
	}
	return &T0kenSplittableShareholderAddedIterator{contract: _T0kenSplittable.contract, event: "ShareholderAdded", logs: logs, sub: sub}, nil
}

// WatchShareholderAdded is a free log subscription operation binding the contract event 0x3082c1c4977de80c4f67ee77b56b282610ec93a7ecbcc31b551e0ac2f7bd0395.
//
// Solidity: e ShareholderAdded(shareholder address)
func (_T0kenSplittable *T0kenSplittableFilterer) WatchShareholderAdded(opts *bind.WatchOpts, sink chan<- *T0kenSplittableShareholderAdded) (event.Subscription, error) {

	logs, sub, err := _T0kenSplittable.contract.WatchLogs(opts, "ShareholderAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(T0kenSplittableShareholderAdded)
				if err := _T0kenSplittable.contract.UnpackLog(event, "ShareholderAdded", log); err != nil {
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

// T0kenSplittableShareholderRemovedIterator is returned from FilterShareholderRemoved and is used to iterate over the raw logs and unpacked data for ShareholderRemoved events raised by the T0kenSplittable contract.
type T0kenSplittableShareholderRemovedIterator struct {
	Event *T0kenSplittableShareholderRemoved // Event containing the contract specifics and raw log

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
func (it *T0kenSplittableShareholderRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(T0kenSplittableShareholderRemoved)
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
		it.Event = new(T0kenSplittableShareholderRemoved)
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
func (it *T0kenSplittableShareholderRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *T0kenSplittableShareholderRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// T0kenSplittableShareholderRemoved represents a ShareholderRemoved event raised by the T0kenSplittable contract.
type T0kenSplittableShareholderRemoved struct {
	Shareholder common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterShareholderRemoved is a free log retrieval operation binding the contract event 0x7ba114ff3d9844510a088eea94cd35562e7c97a2d36a418b37b2e61e5c77affe.
//
// Solidity: e ShareholderRemoved(shareholder address)
func (_T0kenSplittable *T0kenSplittableFilterer) FilterShareholderRemoved(opts *bind.FilterOpts) (*T0kenSplittableShareholderRemovedIterator, error) {

	logs, sub, err := _T0kenSplittable.contract.FilterLogs(opts, "ShareholderRemoved")
	if err != nil {
		return nil, err
	}
	return &T0kenSplittableShareholderRemovedIterator{contract: _T0kenSplittable.contract, event: "ShareholderRemoved", logs: logs, sub: sub}, nil
}

// WatchShareholderRemoved is a free log subscription operation binding the contract event 0x7ba114ff3d9844510a088eea94cd35562e7c97a2d36a418b37b2e61e5c77affe.
//
// Solidity: e ShareholderRemoved(shareholder address)
func (_T0kenSplittable *T0kenSplittableFilterer) WatchShareholderRemoved(opts *bind.WatchOpts, sink chan<- *T0kenSplittableShareholderRemoved) (event.Subscription, error) {

	logs, sub, err := _T0kenSplittable.contract.WatchLogs(opts, "ShareholderRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(T0kenSplittableShareholderRemoved)
				if err := _T0kenSplittable.contract.UnpackLog(event, "ShareholderRemoved", log); err != nil {
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

// T0kenSplittableTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the T0kenSplittable contract.
type T0kenSplittableTransferIterator struct {
	Event *T0kenSplittableTransfer // Event containing the contract specifics and raw log

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
func (it *T0kenSplittableTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(T0kenSplittableTransfer)
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
		it.Event = new(T0kenSplittableTransfer)
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
func (it *T0kenSplittableTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *T0kenSplittableTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// T0kenSplittableTransfer represents a Transfer event raised by the T0kenSplittable contract.
type T0kenSplittableTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: e Transfer(from indexed address, to indexed address, value uint256)
func (_T0kenSplittable *T0kenSplittableFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*T0kenSplittableTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _T0kenSplittable.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &T0kenSplittableTransferIterator{contract: _T0kenSplittable.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: e Transfer(from indexed address, to indexed address, value uint256)
func (_T0kenSplittable *T0kenSplittableFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *T0kenSplittableTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _T0kenSplittable.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(T0kenSplittableTransfer)
				if err := _T0kenSplittable.contract.UnpackLog(event, "Transfer", log); err != nil {
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
