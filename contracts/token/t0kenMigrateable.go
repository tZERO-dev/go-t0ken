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

// T0kenMigrateableABI is the input ABI used to generate the binding from.
const T0kenMigrateableABI = "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"tokenName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"tokenSymbol\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"decimalPlaces\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokens\",\"type\":\"uint256\"}],\"name\":\"Issuance\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"IssuanceFinished\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousIssuer\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newIssuer\",\"type\":\"address\"}],\"name\":\"IssuerSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnerTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"predecessor\",\"type\":\"address\"}],\"name\":\"PrecededBy\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"shareholder\",\"type\":\"address\"}],\"name\":\"ShareholderAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"shareholder\",\"type\":\"address\"}],\"name\":\"ShareholderRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"successor\",\"type\":\"address\"}],\"name\":\"SucceededBy\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[],\"name\":\"ZERO_ADDRESS\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"addrOwner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokens\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"balances\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"compliance\",\"outputs\":[{\"internalType\":\"contractCompliance\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"finishIssuance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"int256\",\"name\":\"index\",\"type\":\"int256\"}],\"name\":\"holderAt\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"holders\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"count\",\"type\":\"int256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"isHolder\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isLocked\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"issuanceFinished\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"quantity\",\"type\":\"uint256\"}],\"name\":\"issueTokens\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"issuer\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"kill\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"predecessor\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"newComplianceAddress\",\"type\":\"address\"}],\"name\":\"setCompliance\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"newIssuer\",\"type\":\"address\"}],\"name\":\"setIssuer\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bool\",\"name\":\"locked\",\"type\":\"bool\"}],\"name\":\"setLocked\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"setPredecessor\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"setSuccessor\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"successor\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokens\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokens\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokens\",\"type\":\"uint256\"}],\"name\":\"transferOverride\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwner\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// T0kenMigrateableBin is the compiled bytecode used for deploying new contracts.
const T0kenMigrateableBin = `6080604052600180546001600160a01b03191690553480156200002157600080fd5b50604051620026c5380380620026c5833981810160405260608110156200004757600080fd5b81019080805160405193929190846401000000008211156200006857600080fd5b9083019060208201858111156200007e57600080fd5b82516401000000008111828201881017156200009957600080fd5b82525081516020918201929091019080838360005b83811015620000c8578181015183820152602001620000ae565b50505050905090810190601f168015620000f65780820380516001836020036101000a031916815260200191505b50604052602001805160405193929190846401000000008211156200011a57600080fd5b9083019060208201858111156200013057600080fd5b82516401000000008111828201881017156200014b57600080fd5b82525081516020918201929091019080838360005b838110156200017a57818101518382015260200162000160565b50505050905090810190601f168015620001a85780820380516001836020036101000a031916815260200191505b50604052602090810151600080546001600160a01b031916331790556001805460ff60a01b19169055855190935085925084918491620001ef916002919086019062000226565b5081516200020590600390602085019062000226565b506004805460ff191660ff9290921691909117905550620002cb9350505050565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106200026957805160ff191683800117855562000299565b8280016001018555821562000299579182015b82811115620002995782518255916020019190600101906200027c565b50620002a7929150620002ab565b5090565b620002c891905b80821115620002a75760008155600101620002b2565b90565b6123ea80620002db6000396000f3fe608060405234801561001057600080fd5b50600436106101e55760003560e01c80636ff968c31161010f578063a5820daa116100a2578063c422293b11610071578063c422293b1461052a578063d4d7b19a14610532578063dd62ed3e14610558578063f898178914610586576101e5565b8063a5820daa146104b3578063a9059cbb146104d0578063b3d724d6146104fc578063b719d03214610522576101e5565b80638188f71c116100de5780638188f71c146104935780638da5cb5b1461049b57806395d89b41146104a3578063a4e2d634146104ab576101e5565b80636ff968c31461041257806370a082311461041a57806380318be8146104405780638082a92914610476576101e5565b806327e235e3116101875780634fb2e45d116101565780634fb2e45d146103b6578063538ba4f9146103dc57806355cc4e57146103e45780636290865d1461040a576101e5565b806327e235e314610362578063313ce5671461038857806341c0e1b5146103a65780634662299a146103ae576101e5565b806318160ddd116101c357806318160ddd146102cf5780631d143848146102e9578063211e28b61461030d57806323b872dd1461032c576101e5565b806306fdde03146101ea578063095ea7b31461026757806310e5bff8146102a7575b600080fd5b6101f26105ac565b6040805160208082528351818301528351919283929083019185019080838360005b8381101561022c578181015183820152602001610214565b50505050905090810190601f1680156102595780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b6102936004803603604081101561027d57600080fd5b506001600160a01b038135169060200135610637565b604080519115158252519081900360200190f35b6102cd600480360360208110156102bd57600080fd5b50356001600160a01b031661074d565b005b6102d76108c4565b60408051918252519081900360200190f35b6102f16108ca565b604080516001600160a01b039092168252519081900360200190f35b6102cd6004803603602081101561032357600080fd5b503515156108d9565b6102936004803603606081101561034257600080fd5b506001600160a01b038135811691602081013590911690604001356109c1565b6102d76004803603602081101561037857600080fd5b50356001600160a01b0316610a1a565b610390610a2c565b6040805160ff9092168252519081900360200190f35b6102cd610a35565b610293610abd565b6102cd600480360360208110156103cc57600080fd5b50356001600160a01b0316610acd565b6102f1610c3f565b6102cd600480360360208110156103fa57600080fd5b50356001600160a01b0316610c4e565b6102f1610dd5565b6102f1610de4565b6102d76004803603602081101561043057600080fd5b50356001600160a01b0316610df3565b6102936004803603606081101561045657600080fd5b506001600160a01b03813581169160208101359091169060400135610e0e565b6102f16004803603602081101561048c57600080fd5b5035610eec565b6102d7610eff565b6102f1610f05565b6101f2610f14565b610293610f6f565b610293600480360360208110156104c957600080fd5b5035610f7f565b610293600480360360408110156104e657600080fd5b506001600160a01b03813516906020013561118d565b6102cd6004803603602081101561051257600080fd5b50356001600160a01b03166111ca565b6102f1611341565b610293611350565b6102936004803603602081101561054857600080fd5b50356001600160a01b031661149d565b6102d76004803603604081101561056e57600080fd5b506001600160a01b03813581169160200135166114b0565b6102cd6004803603602081101561059c57600080fd5b50356001600160a01b03166114db565b6002805460408051602060018416156101000260001901909316849004601f8101849004840282018401909252818152929183018282801561062f5780601f106106045761010080835404028352916020019161062f565b820191906000526020600020905b81548152906001019060200180831161061257829003601f168201915b505050505081565b600154600090600160a01b900460ff16156106835760405162461bcd60e51b815260040180806020018281038252602d815260200180612389602d913960400191505060405180910390fd5b61069460053363ffffffff6115c016565b6106e5576040805162461bcd60e51b815260206004820152601560248201527f4d7573742062652061207368617265686f6c6465720000000000000000000000604482015290519081900360640190fd5b336000818152600c602090815260408083206001600160a01b03881680855290835292819020869055805186815290519293927f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925929181900390910190a35060015b92915050565b600154600160a01b900460ff16156107965760405162461bcd60e51b815260040180806020018281038252602d815260200180612389602d913960400191505060405180910390fd5b6009546001600160a01b031633146107eb576040805162461bcd60e51b815260206004820152601360248201527213db9b1e481a5cdcdd595c88185b1b1bddd959606a1b604482015290519081900360640190fd5b600e546001600160a01b0316818115610839576040805162461bcd60e51b815260206004820152600b60248201526a105b1c9958591e481cd95d60aa1b604482015290519081900360640190fd5b806001600160a01b0316826001600160a01b031614156108a0576040805162461bcd60e51b815260206004820152601060248201527f4d69736d61746368656420746f6b656e00000000000000000000000000000000604482015290519081900360640190fd5b5050600e80546001600160a01b0319166001600160a01b0392909216919091179055565b600b5490565b6009546001600160a01b031681565b6000546001600160a01b031633148061090257506001546000546001600160a01b039081169116145b610953576040805162461bcd60e51b815260206004820152601960248201527f4f776e6572206163636f756e7420697320726571756972656400000000000000604482015290519081900360640190fd5b60015460ff600160a01b90910416151581151514156109a35760405162461bcd60e51b81526004018080602001828103825260288152602001806123616028913960400191505060405180910390fd5b60018054911515600160a01b0260ff60a01b19909216919091179055565b600d546000906001600160a01b031633148015906109ed5750600e546001600160a01b03848116911614155b15610a04576109fd8484846115f7565b9050610a13565b610a0f84848461179b565b5060015b9392505050565b600a6020526000908152604090205481565b60045460ff1681565b6000546001600160a01b0316331480610a5e57506001546000546001600160a01b039081169116145b610aaf576040805162461bcd60e51b815260206004820152601960248201527f4f776e6572206163636f756e7420697320726571756972656400000000000000604482015290519081900360640190fd5b6000546001600160a01b0316ff5b600954600160a01b900460ff1681565b6000546001600160a01b0316331480610af657506001546000546001600160a01b039081169116145b610b47576040805162461bcd60e51b815260206004820152601960248201527f4f776e6572206163636f756e7420697320726571756972656400000000000000604482015290519081900360640190fd5b6000546001600160a01b0382811691161415610b945760405162461bcd60e51b81526004018080602001828103825260258152602001806122f16025913960400191505060405180910390fd5b6001600160a01b038116610bef576040805162461bcd60e51b815260206004820181905260248201527f4e6577204f776e65722063616e6e6f74206265207a65726f2061646472657373604482015290519081900360640190fd5b600080546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8934ce4adea8d9ce0d714d2c22b86790e41b7731c84b926fbbdc1d40ff6533c99190a35050565b6001546001600160a01b031681565b600154600160a01b900460ff1615610c975760405162461bcd60e51b815260040180806020018281038252602d815260200180612389602d913960400191505060405180910390fd5b6000546001600160a01b0316331480610cc057506001546000546001600160a01b039081169116145b610d11576040805162461bcd60e51b815260206004820152601960248201527f4f776e6572206163636f756e7420697320726571756972656400000000000000604482015290519081900360640190fd5b610d2260058263ffffffff6115c016565b15610d5e5760405162461bcd60e51b81526004018080602001828103825260218152602001806122d06021913960400191505060405180910390fd5b6009546001600160a01b03166000818152600a6020526040902054610d8591908390611b63565b600980546001600160a01b0319166001600160a01b0383811691821792839055604051919216907ff7189b85d7899f5a32d733e6584c4f1dcdff0274f09d969d186c1797673ede1190600090a350565b6008546001600160a01b031681565b600e546001600160a01b031681565b6001600160a01b03166000908152600a602052604090205490565b600154600090600160a01b900460ff1615610e5a5760405162461bcd60e51b815260040180806020018281038252602d815260200180612389602d913960400191505060405180910390fd5b6001600160a01b0384166000908152600a602052604090205484908390811115610ec0576040805162461bcd60e51b8152602060048201526012602482015271496e73756666696369656e742066756e647360701b604482015290519081900360640190fd5b6000610ecf8787876001611d05565b90508015610ee257610ee2878787611b63565b9695505050505050565b600061074760058363ffffffff611e5316565b60055481565b6000546001600160a01b031681565b6003805460408051602060026001851615610100026000190190941693909304601f8101849004840282018401909252818152929183018282801561062f5780601f106106045761010080835404028352916020019161062f565b600154600160a01b900460ff1681565b600154600090600160a01b900460ff1615610fcb5760405162461bcd60e51b815260040180806020018281038252602d815260200180612389602d913960400191505060405180910390fd5b6009546001600160a01b03163314611020576040805162461bcd60e51b815260206004820152601360248201527213db9b1e481a5cdcdd595c88185b1b1bddd959606a1b604482015290519081900360640190fd5b600954600160a01b900460ff161561107f576040805162461bcd60e51b815260206004820152601960248201527f49737375616e636520616c72656164792066696e697368656400000000000000604482015290519081900360640190fd5b81156110fd57600b54611098908363ffffffff611ed916565b600b556009546001600160a01b03166000908152600a60205260409020546110c6908363ffffffff611ed916565b600980546001600160a01b039081166000908152600a6020526040902092909255546110fb916005911663ffffffff611f3316565b505b6009546040805184815290516001600160a01b03909216917f9cb9c14f7bc76e3a89b796b091850526236115352a198b1e472f00e91376bbcb9181900360200190a26009546040805184815290516001600160a01b03909216916000917fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef919081900360200190a3506001919050565b600e546000906001600160a01b038481169116146111b6576111af8383611fd8565b9050610747565b6111c133848461179b565b50600192915050565b600154600160a01b900460ff16156112135760405162461bcd60e51b815260040180806020018281038252602d815260200180612389602d913960400191505060405180910390fd5b6009546001600160a01b03163314611268576040805162461bcd60e51b815260206004820152601360248201527213db9b1e481a5cdcdd595c88185b1b1bddd959606a1b604482015290519081900360640190fd5b600d546001600160a01b03168181156112b6576040805162461bcd60e51b815260206004820152600b60248201526a105b1c9958591e481cd95d60aa1b604482015290519081900360640190fd5b806001600160a01b0316826001600160a01b0316141561131d576040805162461bcd60e51b815260206004820152601060248201527f4d69736d61746368656420746f6b656e00000000000000000000000000000000604482015290519081900360640190fd5b5050600d80546001600160a01b0319166001600160a01b0392909216919091179055565b600d546001600160a01b031681565b600154600090600160a01b900460ff161561139c5760405162461bcd60e51b815260040180806020018281038252602d815260200180612389602d913960400191505060405180910390fd5b6009546001600160a01b031633146113f1576040805162461bcd60e51b815260206004820152601360248201527213db9b1e481a5cdcdd595c88185b1b1bddd959606a1b604482015290519081900360640190fd5b600954600160a01b900460ff1615611450576040805162461bcd60e51b815260206004820152601960248201527f49737375616e636520616c72656164792066696e697368656400000000000000604482015290519081900360640190fd5b6009805460ff60a01b1916600160a01b1790556040517f29fe76cc5ca143e91eadf7242fda487fcef09318c1237900f958abe1e2c5beff90600090a150600954600160a01b900460ff1690565b600061074760058363ffffffff6115c016565b6001600160a01b039182166000908152600c6020908152604080832093909416825291909152205490565b600154600160a01b900460ff16156115245760405162461bcd60e51b815260040180806020018281038252602d815260200180612389602d913960400191505060405180910390fd5b6000546001600160a01b031633148061154d57506001546000546001600160a01b039081169116145b61159e576040805162461bcd60e51b815260206004820152601960248201527f4f776e6572206163636f756e7420697320726571756972656400000000000000604482015290519081900360640190fd5b600880546001600160a01b0319166001600160a01b0392909216919091179055565b6001600160a01b0381166000908152600183016020526040812054600019018181128015906115ef5750835481125b949350505050565b600154600090600160a01b900460ff16156116435760405162461bcd60e51b815260040180806020018281038252602d815260200180612389602d913960400191505060405180910390fd5b6001600160a01b0384166000908152600a6020526040902054849083908111156116a9576040805162461bcd60e51b8152602060048201526012602482015271496e73756666696369656e742066756e647360701b604482015290519081900360640190fd5b6001600160a01b0386166000908152600c60209081526040808320338452909152902054841115611721576040805162461bcd60e51b815260206004820152601a60248201527f5472616e73666572206578636565647320616c6c6f77616e6365000000000000604482015290519081900360640190fd5b60006117308787876000611d05565b90508015610ee2576001600160a01b0387166000908152600c6020908152604080832033845290915290205461176c908663ffffffff61218716565b6001600160a01b0388166000908152600c60209081526040808320338452909152902055610ee2878787611b63565b600154600160a01b900460ff16156117e45760405162461bcd60e51b815260040180806020018281038252602d815260200180612389602d913960400191505060405180910390fd5b600e546001600160a01b03838116911614156119a0578061180484610df3565b101561184c576040805162461bcd60e51b8152602060048201526012602482015271496e73756666696369656e742066756e647360701b604482015290519081900360640190fd5b336001600160a01b038416146118a5576009546001600160a01b031633146118a55760405162461bcd60e51b815260040180806020018281038252602681526020018061233b6026913960400191505060405180910390fd5b600e54604080517f23b872dd0000000000000000000000000000000000000000000000000000000081526001600160a01b038681166004830152858116602483015260448201859052915191909216916323b872dd9160648083019260209291908290030181600087803b15801561191c57600080fd5b505af1158015611930573d6000803e3d6000fd5b505050506040513d602081101561194657600080fd5b50516119835760405162461bcd60e51b81526004018080602001828103825260258152602001806123166025913960400191505060405180910390fd5b600e5461199b9084906001600160a01b031683611b63565b611b5e565b600d546001600160a01b0316331415611b115760095481906119ca906001600160a01b0316610df3565b1015611a1d576040805162461bcd60e51b815260206004820152601a60248201527f496e73756666696369656e7420746f6b656e7320697373756564000000000000604482015290519081900360640190fd5b600d54604080517f70a082310000000000000000000000000000000000000000000000000000000081526001600160a01b0386811660048301529151849392909216916370a0823191602480820192602092909190829003018186803b158015611a8657600080fd5b505afa158015611a9a573d6000803e3d6000fd5b505050506040513d6020811015611ab057600080fd5b50511015611afa576040805162461bcd60e51b8152602060048201526012602482015271496e73756666696369656e742066756e647360701b604482015290519081900360640190fd5b60095461199b906001600160a01b03168483611b63565b6040805162461bcd60e51b815260206004820152601160248201527f496e76616c6964206d6967726174696f6e000000000000000000000000000000604482015290519081900360640190fd5b505050565b6001600160a01b0383166000908152600a6020526040902054611b8c908263ffffffff61218716565b6001600160a01b038085166000908152600a60205260408082209390935590841681522054611bc1908263ffffffff611ed916565b6001600160a01b038084166000818152600a602090815260409182902094909455805185815290519193928716927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef92918290030190a36001600160a01b0382166000908152600a602052604090205415801590611c4b5750611c4b60058363ffffffff611f3316565b15611c8d57604080516001600160a01b038416815290517f3082c1c4977de80c4f67ee77b56b282610ec93a7ecbcc31b551e0ac2f7bd03959181900360200190a15b6001600160a01b0383166000908152600a6020526040902054158015611cbf5750611cbf60058463ffffffff6121e416565b15611b5e57604080516001600160a01b038516815290517f7ba114ff3d9844510a088eea94cd35562e7c97a2d36a418b37b2e61e5c77affe9181900360200190a1505050565b6008546000906001600160a01b0316611d20575080156115ef565b8115611dd657600854604080517f5acba2010000000000000000000000000000000000000000000000000000000081523360048201526001600160a01b03888116602483015287811660448301526064820187905291519190921691635acba2019160848083019260209291908290030181600087803b158015611da357600080fd5b505af1158015611db7573d6000803e3d6000fd5b505050506040513d6020811015611dcd57600080fd5b505190506115ef565b600854604080517f6d62a4fe0000000000000000000000000000000000000000000000000000000081523360048201526001600160a01b03888116602483015287811660448301526064820187905291519190921691636d62a4fe9160848083019260209291908290030181600087803b158015611da357600080fd5b6000808212158015611e655750825482125b611eb6576040805162461bcd60e51b815260206004820152601860248201527f496e646578206f757473696465206f6620626f756e64732e0000000000000000604482015290519081900360640190fd5b50600101600090815260029190910160205260409020546001600160a01b031690565b600082820183811015610a13576040805162461bcd60e51b815260206004820152601360248201527f526573756c747320696e206f766572666c6f7700000000000000000000000000604482015290519081900360640190fd5b60006001600160a01b038216611f4b57506000610747565b6001600160a01b038216600090815260018401602052604081205460001901908112801590611f7a5750835481125b15611f89576000915050610747565b5050815460019081018084556001600160a01b0383166000818152838601602090815260408083208590559382526002870190529190912080546001600160a01b031916909117905592915050565b600154600090600160a01b900460ff16156120245760405162461bcd60e51b815260040180806020018281038252602d815260200180612389602d913960400191505060405180910390fd5b336000818152600a6020526040902054839081111561207f576040805162461bcd60e51b8152602060048201526012602482015271496e73756666696369656e742066756e647360701b604482015290519081900360640190fd5b6009546000906001600160a01b031633141561215d57506008546001600160a01b0316158061215857600854600954604080517ffd8258bd0000000000000000000000000000000000000000000000000000000081526001600160a01b039283166004820181905260248201528983166044820152606481018990529051919092169163fd8258bd9160848083019260209291908290030181600087803b15801561212957600080fd5b505af115801561213d573d6000803e3d6000fd5b505050506040513d602081101561215357600080fd5b505190505b61216d565b61216a3387876000611d05565b90505b801561217e5761217e338787611b63565b95945050505050565b6000828211156121de576040805162461bcd60e51b815260206004820152601460248201527f526573756c747320696e20756e646572666c6f77000000000000000000000000604482015290519081900360640190fd5b50900390565b6001600160a01b0381166000908152600180840160205260408220549081128061220e5750835481135b1561221d576000915050610747565b835481121561228457835460009081526002850160208181526040808420546001600160a01b03168085526001890183528185208690558585529290915280832080546001600160a01b0319908116909317905586548352909120805490911690556122a3565b6000818152600285016020526040902080546001600160a01b03191690555b50506001600160a01b03166000908152600182810160205260408220919091558154600019019091559056fe4e6577206973737565722063616e27742062652061207368617265686f6c6465724e6577204f776e65722063616e6e6f74206265207468652063757272656e74206f776e65724661696c656420746f206d69677261746520746f6b656e7320746f20737563636573736f72466f72636564206d6967726174696f6e73207265737472696374656420746f20697373756572436f6e747261637420616c726561647920696e20726571756573746564206c6f636b207374617465436f6e74726163742069732063757272656e746c79206c6f636b656420666f72206d6f64696669636174696f6ea265627a7a723158202bc1455d7dbf7c85d88b23e197c51dae74ac97ecd204e6ead787adfec8fe925564736f6c634300050c0032`

// DeployT0kenMigrateable deploys a new Ethereum contract, binding an instance of T0kenMigrateable to it.
func DeployT0kenMigrateable(auth *bind.TransactOpts, backend bind.ContractBackend, tokenName string, tokenSymbol string, decimalPlaces uint8) (common.Address, *types.Transaction, *T0kenMigrateable, error) {
	parsed, err := abi.JSON(strings.NewReader(T0kenMigrateableABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(T0kenMigrateableBin), backend, tokenName, tokenSymbol, decimalPlaces)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &T0kenMigrateable{T0kenMigrateableCaller: T0kenMigrateableCaller{contract: contract}, T0kenMigrateableTransactor: T0kenMigrateableTransactor{contract: contract}, T0kenMigrateableFilterer: T0kenMigrateableFilterer{contract: contract}}, nil
}

// T0kenMigrateable is an auto generated Go binding around an Ethereum contract.
type T0kenMigrateable struct {
	T0kenMigrateableCaller     // Read-only binding to the contract
	T0kenMigrateableTransactor // Write-only binding to the contract
	T0kenMigrateableFilterer   // Log filterer for contract events
}

// T0kenMigrateableCaller is an auto generated read-only Go binding around an Ethereum contract.
type T0kenMigrateableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// T0kenMigrateableTransactor is an auto generated write-only Go binding around an Ethereum contract.
type T0kenMigrateableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// T0kenMigrateableFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type T0kenMigrateableFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// T0kenMigrateableSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type T0kenMigrateableSession struct {
	Contract     *T0kenMigrateable // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// T0kenMigrateableCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type T0kenMigrateableCallerSession struct {
	Contract *T0kenMigrateableCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// T0kenMigrateableTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type T0kenMigrateableTransactorSession struct {
	Contract     *T0kenMigrateableTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// T0kenMigrateableRaw is an auto generated low-level Go binding around an Ethereum contract.
type T0kenMigrateableRaw struct {
	Contract *T0kenMigrateable // Generic contract binding to access the raw methods on
}

// T0kenMigrateableCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type T0kenMigrateableCallerRaw struct {
	Contract *T0kenMigrateableCaller // Generic read-only contract binding to access the raw methods on
}

// T0kenMigrateableTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type T0kenMigrateableTransactorRaw struct {
	Contract *T0kenMigrateableTransactor // Generic write-only contract binding to access the raw methods on
}

// NewT0kenMigrateable creates a new instance of T0kenMigrateable, bound to a specific deployed contract.
func NewT0kenMigrateable(address common.Address, backend bind.ContractBackend) (*T0kenMigrateable, error) {
	contract, err := bindT0kenMigrateable(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &T0kenMigrateable{T0kenMigrateableCaller: T0kenMigrateableCaller{contract: contract}, T0kenMigrateableTransactor: T0kenMigrateableTransactor{contract: contract}, T0kenMigrateableFilterer: T0kenMigrateableFilterer{contract: contract}}, nil
}

// NewT0kenMigrateableCaller creates a new read-only instance of T0kenMigrateable, bound to a specific deployed contract.
func NewT0kenMigrateableCaller(address common.Address, caller bind.ContractCaller) (*T0kenMigrateableCaller, error) {
	contract, err := bindT0kenMigrateable(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &T0kenMigrateableCaller{contract: contract}, nil
}

// NewT0kenMigrateableTransactor creates a new write-only instance of T0kenMigrateable, bound to a specific deployed contract.
func NewT0kenMigrateableTransactor(address common.Address, transactor bind.ContractTransactor) (*T0kenMigrateableTransactor, error) {
	contract, err := bindT0kenMigrateable(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &T0kenMigrateableTransactor{contract: contract}, nil
}

// NewT0kenMigrateableFilterer creates a new log filterer instance of T0kenMigrateable, bound to a specific deployed contract.
func NewT0kenMigrateableFilterer(address common.Address, filterer bind.ContractFilterer) (*T0kenMigrateableFilterer, error) {
	contract, err := bindT0kenMigrateable(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &T0kenMigrateableFilterer{contract: contract}, nil
}

// bindT0kenMigrateable binds a generic wrapper to an already deployed contract.
func bindT0kenMigrateable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(T0kenMigrateableABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_T0kenMigrateable *T0kenMigrateableRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _T0kenMigrateable.Contract.T0kenMigrateableCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_T0kenMigrateable *T0kenMigrateableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _T0kenMigrateable.Contract.T0kenMigrateableTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_T0kenMigrateable *T0kenMigrateableRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _T0kenMigrateable.Contract.T0kenMigrateableTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_T0kenMigrateable *T0kenMigrateableCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _T0kenMigrateable.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_T0kenMigrateable *T0kenMigrateableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _T0kenMigrateable.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_T0kenMigrateable *T0kenMigrateableTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _T0kenMigrateable.Contract.contract.Transact(opts, method, params...)
}

// ZEROADDRESS is a free data retrieval call binding the contract method 0x538ba4f9.
//
// Solidity: function ZERO_ADDRESS() constant returns(address)
func (_T0kenMigrateable *T0kenMigrateableCaller) ZEROADDRESS(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _T0kenMigrateable.contract.Call(opts, out, "ZERO_ADDRESS")
	return *ret0, err
}

// ZEROADDRESS is a free data retrieval call binding the contract method 0x538ba4f9.
//
// Solidity: function ZERO_ADDRESS() constant returns(address)
func (_T0kenMigrateable *T0kenMigrateableSession) ZEROADDRESS() (common.Address, error) {
	return _T0kenMigrateable.Contract.ZEROADDRESS(&_T0kenMigrateable.CallOpts)
}

// ZEROADDRESS is a free data retrieval call binding the contract method 0x538ba4f9.
//
// Solidity: function ZERO_ADDRESS() constant returns(address)
func (_T0kenMigrateable *T0kenMigrateableCallerSession) ZEROADDRESS() (common.Address, error) {
	return _T0kenMigrateable.Contract.ZEROADDRESS(&_T0kenMigrateable.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(addrOwner address, spender address) constant returns(uint256)
func (_T0kenMigrateable *T0kenMigrateableCaller) Allowance(opts *bind.CallOpts, addrOwner common.Address, spender common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _T0kenMigrateable.contract.Call(opts, out, "allowance", addrOwner, spender)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(addrOwner address, spender address) constant returns(uint256)
func (_T0kenMigrateable *T0kenMigrateableSession) Allowance(addrOwner common.Address, spender common.Address) (*big.Int, error) {
	return _T0kenMigrateable.Contract.Allowance(&_T0kenMigrateable.CallOpts, addrOwner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(addrOwner address, spender address) constant returns(uint256)
func (_T0kenMigrateable *T0kenMigrateableCallerSession) Allowance(addrOwner common.Address, spender common.Address) (*big.Int, error) {
	return _T0kenMigrateable.Contract.Allowance(&_T0kenMigrateable.CallOpts, addrOwner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(addr address) constant returns(uint256)
func (_T0kenMigrateable *T0kenMigrateableCaller) BalanceOf(opts *bind.CallOpts, addr common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _T0kenMigrateable.contract.Call(opts, out, "balanceOf", addr)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(addr address) constant returns(uint256)
func (_T0kenMigrateable *T0kenMigrateableSession) BalanceOf(addr common.Address) (*big.Int, error) {
	return _T0kenMigrateable.Contract.BalanceOf(&_T0kenMigrateable.CallOpts, addr)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(addr address) constant returns(uint256)
func (_T0kenMigrateable *T0kenMigrateableCallerSession) BalanceOf(addr common.Address) (*big.Int, error) {
	return _T0kenMigrateable.Contract.BalanceOf(&_T0kenMigrateable.CallOpts, addr)
}

// Balances is a free data retrieval call binding the contract method 0x27e235e3.
//
// Solidity: function balances( address) constant returns(uint256)
func (_T0kenMigrateable *T0kenMigrateableCaller) Balances(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _T0kenMigrateable.contract.Call(opts, out, "balances", arg0)
	return *ret0, err
}

// Balances is a free data retrieval call binding the contract method 0x27e235e3.
//
// Solidity: function balances( address) constant returns(uint256)
func (_T0kenMigrateable *T0kenMigrateableSession) Balances(arg0 common.Address) (*big.Int, error) {
	return _T0kenMigrateable.Contract.Balances(&_T0kenMigrateable.CallOpts, arg0)
}

// Balances is a free data retrieval call binding the contract method 0x27e235e3.
//
// Solidity: function balances( address) constant returns(uint256)
func (_T0kenMigrateable *T0kenMigrateableCallerSession) Balances(arg0 common.Address) (*big.Int, error) {
	return _T0kenMigrateable.Contract.Balances(&_T0kenMigrateable.CallOpts, arg0)
}

// Compliance is a free data retrieval call binding the contract method 0x6290865d.
//
// Solidity: function compliance() constant returns(address)
func (_T0kenMigrateable *T0kenMigrateableCaller) Compliance(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _T0kenMigrateable.contract.Call(opts, out, "compliance")
	return *ret0, err
}

// Compliance is a free data retrieval call binding the contract method 0x6290865d.
//
// Solidity: function compliance() constant returns(address)
func (_T0kenMigrateable *T0kenMigrateableSession) Compliance() (common.Address, error) {
	return _T0kenMigrateable.Contract.Compliance(&_T0kenMigrateable.CallOpts)
}

// Compliance is a free data retrieval call binding the contract method 0x6290865d.
//
// Solidity: function compliance() constant returns(address)
func (_T0kenMigrateable *T0kenMigrateableCallerSession) Compliance() (common.Address, error) {
	return _T0kenMigrateable.Contract.Compliance(&_T0kenMigrateable.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_T0kenMigrateable *T0kenMigrateableCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _T0kenMigrateable.contract.Call(opts, out, "decimals")
	return *ret0, err
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_T0kenMigrateable *T0kenMigrateableSession) Decimals() (uint8, error) {
	return _T0kenMigrateable.Contract.Decimals(&_T0kenMigrateable.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_T0kenMigrateable *T0kenMigrateableCallerSession) Decimals() (uint8, error) {
	return _T0kenMigrateable.Contract.Decimals(&_T0kenMigrateable.CallOpts)
}

// HolderAt is a free data retrieval call binding the contract method 0x8082a929.
//
// Solidity: function holderAt(index int256) constant returns(address)
func (_T0kenMigrateable *T0kenMigrateableCaller) HolderAt(opts *bind.CallOpts, index *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _T0kenMigrateable.contract.Call(opts, out, "holderAt", index)
	return *ret0, err
}

// HolderAt is a free data retrieval call binding the contract method 0x8082a929.
//
// Solidity: function holderAt(index int256) constant returns(address)
func (_T0kenMigrateable *T0kenMigrateableSession) HolderAt(index *big.Int) (common.Address, error) {
	return _T0kenMigrateable.Contract.HolderAt(&_T0kenMigrateable.CallOpts, index)
}

// HolderAt is a free data retrieval call binding the contract method 0x8082a929.
//
// Solidity: function holderAt(index int256) constant returns(address)
func (_T0kenMigrateable *T0kenMigrateableCallerSession) HolderAt(index *big.Int) (common.Address, error) {
	return _T0kenMigrateable.Contract.HolderAt(&_T0kenMigrateable.CallOpts, index)
}

// Holders is a free data retrieval call binding the contract method 0x8188f71c.
//
// Solidity: function holders() constant returns(count int256)
func (_T0kenMigrateable *T0kenMigrateableCaller) Holders(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _T0kenMigrateable.contract.Call(opts, out, "holders")
	return *ret0, err
}

// Holders is a free data retrieval call binding the contract method 0x8188f71c.
//
// Solidity: function holders() constant returns(count int256)
func (_T0kenMigrateable *T0kenMigrateableSession) Holders() (*big.Int, error) {
	return _T0kenMigrateable.Contract.Holders(&_T0kenMigrateable.CallOpts)
}

// Holders is a free data retrieval call binding the contract method 0x8188f71c.
//
// Solidity: function holders() constant returns(count int256)
func (_T0kenMigrateable *T0kenMigrateableCallerSession) Holders() (*big.Int, error) {
	return _T0kenMigrateable.Contract.Holders(&_T0kenMigrateable.CallOpts)
}

// IsHolder is a free data retrieval call binding the contract method 0xd4d7b19a.
//
// Solidity: function isHolder(addr address) constant returns(bool)
func (_T0kenMigrateable *T0kenMigrateableCaller) IsHolder(opts *bind.CallOpts, addr common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _T0kenMigrateable.contract.Call(opts, out, "isHolder", addr)
	return *ret0, err
}

// IsHolder is a free data retrieval call binding the contract method 0xd4d7b19a.
//
// Solidity: function isHolder(addr address) constant returns(bool)
func (_T0kenMigrateable *T0kenMigrateableSession) IsHolder(addr common.Address) (bool, error) {
	return _T0kenMigrateable.Contract.IsHolder(&_T0kenMigrateable.CallOpts, addr)
}

// IsHolder is a free data retrieval call binding the contract method 0xd4d7b19a.
//
// Solidity: function isHolder(addr address) constant returns(bool)
func (_T0kenMigrateable *T0kenMigrateableCallerSession) IsHolder(addr common.Address) (bool, error) {
	return _T0kenMigrateable.Contract.IsHolder(&_T0kenMigrateable.CallOpts, addr)
}

// IsLocked is a free data retrieval call binding the contract method 0xa4e2d634.
//
// Solidity: function isLocked() constant returns(bool)
func (_T0kenMigrateable *T0kenMigrateableCaller) IsLocked(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _T0kenMigrateable.contract.Call(opts, out, "isLocked")
	return *ret0, err
}

// IsLocked is a free data retrieval call binding the contract method 0xa4e2d634.
//
// Solidity: function isLocked() constant returns(bool)
func (_T0kenMigrateable *T0kenMigrateableSession) IsLocked() (bool, error) {
	return _T0kenMigrateable.Contract.IsLocked(&_T0kenMigrateable.CallOpts)
}

// IsLocked is a free data retrieval call binding the contract method 0xa4e2d634.
//
// Solidity: function isLocked() constant returns(bool)
func (_T0kenMigrateable *T0kenMigrateableCallerSession) IsLocked() (bool, error) {
	return _T0kenMigrateable.Contract.IsLocked(&_T0kenMigrateable.CallOpts)
}

// IssuanceFinished is a free data retrieval call binding the contract method 0x4662299a.
//
// Solidity: function issuanceFinished() constant returns(bool)
func (_T0kenMigrateable *T0kenMigrateableCaller) IssuanceFinished(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _T0kenMigrateable.contract.Call(opts, out, "issuanceFinished")
	return *ret0, err
}

// IssuanceFinished is a free data retrieval call binding the contract method 0x4662299a.
//
// Solidity: function issuanceFinished() constant returns(bool)
func (_T0kenMigrateable *T0kenMigrateableSession) IssuanceFinished() (bool, error) {
	return _T0kenMigrateable.Contract.IssuanceFinished(&_T0kenMigrateable.CallOpts)
}

// IssuanceFinished is a free data retrieval call binding the contract method 0x4662299a.
//
// Solidity: function issuanceFinished() constant returns(bool)
func (_T0kenMigrateable *T0kenMigrateableCallerSession) IssuanceFinished() (bool, error) {
	return _T0kenMigrateable.Contract.IssuanceFinished(&_T0kenMigrateable.CallOpts)
}

// Issuer is a free data retrieval call binding the contract method 0x1d143848.
//
// Solidity: function issuer() constant returns(address)
func (_T0kenMigrateable *T0kenMigrateableCaller) Issuer(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _T0kenMigrateable.contract.Call(opts, out, "issuer")
	return *ret0, err
}

// Issuer is a free data retrieval call binding the contract method 0x1d143848.
//
// Solidity: function issuer() constant returns(address)
func (_T0kenMigrateable *T0kenMigrateableSession) Issuer() (common.Address, error) {
	return _T0kenMigrateable.Contract.Issuer(&_T0kenMigrateable.CallOpts)
}

// Issuer is a free data retrieval call binding the contract method 0x1d143848.
//
// Solidity: function issuer() constant returns(address)
func (_T0kenMigrateable *T0kenMigrateableCallerSession) Issuer() (common.Address, error) {
	return _T0kenMigrateable.Contract.Issuer(&_T0kenMigrateable.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_T0kenMigrateable *T0kenMigrateableCaller) Name(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _T0kenMigrateable.contract.Call(opts, out, "name")
	return *ret0, err
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_T0kenMigrateable *T0kenMigrateableSession) Name() (string, error) {
	return _T0kenMigrateable.Contract.Name(&_T0kenMigrateable.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_T0kenMigrateable *T0kenMigrateableCallerSession) Name() (string, error) {
	return _T0kenMigrateable.Contract.Name(&_T0kenMigrateable.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_T0kenMigrateable *T0kenMigrateableCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _T0kenMigrateable.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_T0kenMigrateable *T0kenMigrateableSession) Owner() (common.Address, error) {
	return _T0kenMigrateable.Contract.Owner(&_T0kenMigrateable.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_T0kenMigrateable *T0kenMigrateableCallerSession) Owner() (common.Address, error) {
	return _T0kenMigrateable.Contract.Owner(&_T0kenMigrateable.CallOpts)
}

// Predecessor is a free data retrieval call binding the contract method 0xb719d032.
//
// Solidity: function predecessor() constant returns(address)
func (_T0kenMigrateable *T0kenMigrateableCaller) Predecessor(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _T0kenMigrateable.contract.Call(opts, out, "predecessor")
	return *ret0, err
}

// Predecessor is a free data retrieval call binding the contract method 0xb719d032.
//
// Solidity: function predecessor() constant returns(address)
func (_T0kenMigrateable *T0kenMigrateableSession) Predecessor() (common.Address, error) {
	return _T0kenMigrateable.Contract.Predecessor(&_T0kenMigrateable.CallOpts)
}

// Predecessor is a free data retrieval call binding the contract method 0xb719d032.
//
// Solidity: function predecessor() constant returns(address)
func (_T0kenMigrateable *T0kenMigrateableCallerSession) Predecessor() (common.Address, error) {
	return _T0kenMigrateable.Contract.Predecessor(&_T0kenMigrateable.CallOpts)
}

// Successor is a free data retrieval call binding the contract method 0x6ff968c3.
//
// Solidity: function successor() constant returns(address)
func (_T0kenMigrateable *T0kenMigrateableCaller) Successor(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _T0kenMigrateable.contract.Call(opts, out, "successor")
	return *ret0, err
}

// Successor is a free data retrieval call binding the contract method 0x6ff968c3.
//
// Solidity: function successor() constant returns(address)
func (_T0kenMigrateable *T0kenMigrateableSession) Successor() (common.Address, error) {
	return _T0kenMigrateable.Contract.Successor(&_T0kenMigrateable.CallOpts)
}

// Successor is a free data retrieval call binding the contract method 0x6ff968c3.
//
// Solidity: function successor() constant returns(address)
func (_T0kenMigrateable *T0kenMigrateableCallerSession) Successor() (common.Address, error) {
	return _T0kenMigrateable.Contract.Successor(&_T0kenMigrateable.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_T0kenMigrateable *T0kenMigrateableCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _T0kenMigrateable.contract.Call(opts, out, "symbol")
	return *ret0, err
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_T0kenMigrateable *T0kenMigrateableSession) Symbol() (string, error) {
	return _T0kenMigrateable.Contract.Symbol(&_T0kenMigrateable.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_T0kenMigrateable *T0kenMigrateableCallerSession) Symbol() (string, error) {
	return _T0kenMigrateable.Contract.Symbol(&_T0kenMigrateable.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_T0kenMigrateable *T0kenMigrateableCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _T0kenMigrateable.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_T0kenMigrateable *T0kenMigrateableSession) TotalSupply() (*big.Int, error) {
	return _T0kenMigrateable.Contract.TotalSupply(&_T0kenMigrateable.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_T0kenMigrateable *T0kenMigrateableCallerSession) TotalSupply() (*big.Int, error) {
	return _T0kenMigrateable.Contract.TotalSupply(&_T0kenMigrateable.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(spender address, tokens uint256) returns(bool)
func (_T0kenMigrateable *T0kenMigrateableTransactor) Approve(opts *bind.TransactOpts, spender common.Address, tokens *big.Int) (*types.Transaction, error) {
	return _T0kenMigrateable.contract.Transact(opts, "approve", spender, tokens)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(spender address, tokens uint256) returns(bool)
func (_T0kenMigrateable *T0kenMigrateableSession) Approve(spender common.Address, tokens *big.Int) (*types.Transaction, error) {
	return _T0kenMigrateable.Contract.Approve(&_T0kenMigrateable.TransactOpts, spender, tokens)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(spender address, tokens uint256) returns(bool)
func (_T0kenMigrateable *T0kenMigrateableTransactorSession) Approve(spender common.Address, tokens *big.Int) (*types.Transaction, error) {
	return _T0kenMigrateable.Contract.Approve(&_T0kenMigrateable.TransactOpts, spender, tokens)
}

// FinishIssuance is a paid mutator transaction binding the contract method 0xc422293b.
//
// Solidity: function finishIssuance() returns(bool)
func (_T0kenMigrateable *T0kenMigrateableTransactor) FinishIssuance(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _T0kenMigrateable.contract.Transact(opts, "finishIssuance")
}

// FinishIssuance is a paid mutator transaction binding the contract method 0xc422293b.
//
// Solidity: function finishIssuance() returns(bool)
func (_T0kenMigrateable *T0kenMigrateableSession) FinishIssuance() (*types.Transaction, error) {
	return _T0kenMigrateable.Contract.FinishIssuance(&_T0kenMigrateable.TransactOpts)
}

// FinishIssuance is a paid mutator transaction binding the contract method 0xc422293b.
//
// Solidity: function finishIssuance() returns(bool)
func (_T0kenMigrateable *T0kenMigrateableTransactorSession) FinishIssuance() (*types.Transaction, error) {
	return _T0kenMigrateable.Contract.FinishIssuance(&_T0kenMigrateable.TransactOpts)
}

// IssueTokens is a paid mutator transaction binding the contract method 0xa5820daa.
//
// Solidity: function issueTokens(quantity uint256) returns(bool)
func (_T0kenMigrateable *T0kenMigrateableTransactor) IssueTokens(opts *bind.TransactOpts, quantity *big.Int) (*types.Transaction, error) {
	return _T0kenMigrateable.contract.Transact(opts, "issueTokens", quantity)
}

// IssueTokens is a paid mutator transaction binding the contract method 0xa5820daa.
//
// Solidity: function issueTokens(quantity uint256) returns(bool)
func (_T0kenMigrateable *T0kenMigrateableSession) IssueTokens(quantity *big.Int) (*types.Transaction, error) {
	return _T0kenMigrateable.Contract.IssueTokens(&_T0kenMigrateable.TransactOpts, quantity)
}

// IssueTokens is a paid mutator transaction binding the contract method 0xa5820daa.
//
// Solidity: function issueTokens(quantity uint256) returns(bool)
func (_T0kenMigrateable *T0kenMigrateableTransactorSession) IssueTokens(quantity *big.Int) (*types.Transaction, error) {
	return _T0kenMigrateable.Contract.IssueTokens(&_T0kenMigrateable.TransactOpts, quantity)
}

// Kill is a paid mutator transaction binding the contract method 0x41c0e1b5.
//
// Solidity: function kill() returns()
func (_T0kenMigrateable *T0kenMigrateableTransactor) Kill(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _T0kenMigrateable.contract.Transact(opts, "kill")
}

// Kill is a paid mutator transaction binding the contract method 0x41c0e1b5.
//
// Solidity: function kill() returns()
func (_T0kenMigrateable *T0kenMigrateableSession) Kill() (*types.Transaction, error) {
	return _T0kenMigrateable.Contract.Kill(&_T0kenMigrateable.TransactOpts)
}

// Kill is a paid mutator transaction binding the contract method 0x41c0e1b5.
//
// Solidity: function kill() returns()
func (_T0kenMigrateable *T0kenMigrateableTransactorSession) Kill() (*types.Transaction, error) {
	return _T0kenMigrateable.Contract.Kill(&_T0kenMigrateable.TransactOpts)
}

// SetCompliance is a paid mutator transaction binding the contract method 0xf8981789.
//
// Solidity: function setCompliance(newComplianceAddress address) returns()
func (_T0kenMigrateable *T0kenMigrateableTransactor) SetCompliance(opts *bind.TransactOpts, newComplianceAddress common.Address) (*types.Transaction, error) {
	return _T0kenMigrateable.contract.Transact(opts, "setCompliance", newComplianceAddress)
}

// SetCompliance is a paid mutator transaction binding the contract method 0xf8981789.
//
// Solidity: function setCompliance(newComplianceAddress address) returns()
func (_T0kenMigrateable *T0kenMigrateableSession) SetCompliance(newComplianceAddress common.Address) (*types.Transaction, error) {
	return _T0kenMigrateable.Contract.SetCompliance(&_T0kenMigrateable.TransactOpts, newComplianceAddress)
}

// SetCompliance is a paid mutator transaction binding the contract method 0xf8981789.
//
// Solidity: function setCompliance(newComplianceAddress address) returns()
func (_T0kenMigrateable *T0kenMigrateableTransactorSession) SetCompliance(newComplianceAddress common.Address) (*types.Transaction, error) {
	return _T0kenMigrateable.Contract.SetCompliance(&_T0kenMigrateable.TransactOpts, newComplianceAddress)
}

// SetIssuer is a paid mutator transaction binding the contract method 0x55cc4e57.
//
// Solidity: function setIssuer(newIssuer address) returns()
func (_T0kenMigrateable *T0kenMigrateableTransactor) SetIssuer(opts *bind.TransactOpts, newIssuer common.Address) (*types.Transaction, error) {
	return _T0kenMigrateable.contract.Transact(opts, "setIssuer", newIssuer)
}

// SetIssuer is a paid mutator transaction binding the contract method 0x55cc4e57.
//
// Solidity: function setIssuer(newIssuer address) returns()
func (_T0kenMigrateable *T0kenMigrateableSession) SetIssuer(newIssuer common.Address) (*types.Transaction, error) {
	return _T0kenMigrateable.Contract.SetIssuer(&_T0kenMigrateable.TransactOpts, newIssuer)
}

// SetIssuer is a paid mutator transaction binding the contract method 0x55cc4e57.
//
// Solidity: function setIssuer(newIssuer address) returns()
func (_T0kenMigrateable *T0kenMigrateableTransactorSession) SetIssuer(newIssuer common.Address) (*types.Transaction, error) {
	return _T0kenMigrateable.Contract.SetIssuer(&_T0kenMigrateable.TransactOpts, newIssuer)
}

// SetLocked is a paid mutator transaction binding the contract method 0x211e28b6.
//
// Solidity: function setLocked(locked bool) returns()
func (_T0kenMigrateable *T0kenMigrateableTransactor) SetLocked(opts *bind.TransactOpts, locked bool) (*types.Transaction, error) {
	return _T0kenMigrateable.contract.Transact(opts, "setLocked", locked)
}

// SetLocked is a paid mutator transaction binding the contract method 0x211e28b6.
//
// Solidity: function setLocked(locked bool) returns()
func (_T0kenMigrateable *T0kenMigrateableSession) SetLocked(locked bool) (*types.Transaction, error) {
	return _T0kenMigrateable.Contract.SetLocked(&_T0kenMigrateable.TransactOpts, locked)
}

// SetLocked is a paid mutator transaction binding the contract method 0x211e28b6.
//
// Solidity: function setLocked(locked bool) returns()
func (_T0kenMigrateable *T0kenMigrateableTransactorSession) SetLocked(locked bool) (*types.Transaction, error) {
	return _T0kenMigrateable.Contract.SetLocked(&_T0kenMigrateable.TransactOpts, locked)
}

// SetPredecessor is a paid mutator transaction binding the contract method 0xb3d724d6.
//
// Solidity: function setPredecessor(token address) returns()
func (_T0kenMigrateable *T0kenMigrateableTransactor) SetPredecessor(opts *bind.TransactOpts, token common.Address) (*types.Transaction, error) {
	return _T0kenMigrateable.contract.Transact(opts, "setPredecessor", token)
}

// SetPredecessor is a paid mutator transaction binding the contract method 0xb3d724d6.
//
// Solidity: function setPredecessor(token address) returns()
func (_T0kenMigrateable *T0kenMigrateableSession) SetPredecessor(token common.Address) (*types.Transaction, error) {
	return _T0kenMigrateable.Contract.SetPredecessor(&_T0kenMigrateable.TransactOpts, token)
}

// SetPredecessor is a paid mutator transaction binding the contract method 0xb3d724d6.
//
// Solidity: function setPredecessor(token address) returns()
func (_T0kenMigrateable *T0kenMigrateableTransactorSession) SetPredecessor(token common.Address) (*types.Transaction, error) {
	return _T0kenMigrateable.Contract.SetPredecessor(&_T0kenMigrateable.TransactOpts, token)
}

// SetSuccessor is a paid mutator transaction binding the contract method 0x10e5bff8.
//
// Solidity: function setSuccessor(token address) returns()
func (_T0kenMigrateable *T0kenMigrateableTransactor) SetSuccessor(opts *bind.TransactOpts, token common.Address) (*types.Transaction, error) {
	return _T0kenMigrateable.contract.Transact(opts, "setSuccessor", token)
}

// SetSuccessor is a paid mutator transaction binding the contract method 0x10e5bff8.
//
// Solidity: function setSuccessor(token address) returns()
func (_T0kenMigrateable *T0kenMigrateableSession) SetSuccessor(token common.Address) (*types.Transaction, error) {
	return _T0kenMigrateable.Contract.SetSuccessor(&_T0kenMigrateable.TransactOpts, token)
}

// SetSuccessor is a paid mutator transaction binding the contract method 0x10e5bff8.
//
// Solidity: function setSuccessor(token address) returns()
func (_T0kenMigrateable *T0kenMigrateableTransactorSession) SetSuccessor(token common.Address) (*types.Transaction, error) {
	return _T0kenMigrateable.Contract.SetSuccessor(&_T0kenMigrateable.TransactOpts, token)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(to address, tokens uint256) returns(bool)
func (_T0kenMigrateable *T0kenMigrateableTransactor) Transfer(opts *bind.TransactOpts, to common.Address, tokens *big.Int) (*types.Transaction, error) {
	return _T0kenMigrateable.contract.Transact(opts, "transfer", to, tokens)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(to address, tokens uint256) returns(bool)
func (_T0kenMigrateable *T0kenMigrateableSession) Transfer(to common.Address, tokens *big.Int) (*types.Transaction, error) {
	return _T0kenMigrateable.Contract.Transfer(&_T0kenMigrateable.TransactOpts, to, tokens)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(to address, tokens uint256) returns(bool)
func (_T0kenMigrateable *T0kenMigrateableTransactorSession) Transfer(to common.Address, tokens *big.Int) (*types.Transaction, error) {
	return _T0kenMigrateable.Contract.Transfer(&_T0kenMigrateable.TransactOpts, to, tokens)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(from address, to address, tokens uint256) returns(bool)
func (_T0kenMigrateable *T0kenMigrateableTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokens *big.Int) (*types.Transaction, error) {
	return _T0kenMigrateable.contract.Transact(opts, "transferFrom", from, to, tokens)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(from address, to address, tokens uint256) returns(bool)
func (_T0kenMigrateable *T0kenMigrateableSession) TransferFrom(from common.Address, to common.Address, tokens *big.Int) (*types.Transaction, error) {
	return _T0kenMigrateable.Contract.TransferFrom(&_T0kenMigrateable.TransactOpts, from, to, tokens)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(from address, to address, tokens uint256) returns(bool)
func (_T0kenMigrateable *T0kenMigrateableTransactorSession) TransferFrom(from common.Address, to common.Address, tokens *big.Int) (*types.Transaction, error) {
	return _T0kenMigrateable.Contract.TransferFrom(&_T0kenMigrateable.TransactOpts, from, to, tokens)
}

// TransferOverride is a paid mutator transaction binding the contract method 0x80318be8.
//
// Solidity: function transferOverride(from address, to address, tokens uint256) returns(bool)
func (_T0kenMigrateable *T0kenMigrateableTransactor) TransferOverride(opts *bind.TransactOpts, from common.Address, to common.Address, tokens *big.Int) (*types.Transaction, error) {
	return _T0kenMigrateable.contract.Transact(opts, "transferOverride", from, to, tokens)
}

// TransferOverride is a paid mutator transaction binding the contract method 0x80318be8.
//
// Solidity: function transferOverride(from address, to address, tokens uint256) returns(bool)
func (_T0kenMigrateable *T0kenMigrateableSession) TransferOverride(from common.Address, to common.Address, tokens *big.Int) (*types.Transaction, error) {
	return _T0kenMigrateable.Contract.TransferOverride(&_T0kenMigrateable.TransactOpts, from, to, tokens)
}

// TransferOverride is a paid mutator transaction binding the contract method 0x80318be8.
//
// Solidity: function transferOverride(from address, to address, tokens uint256) returns(bool)
func (_T0kenMigrateable *T0kenMigrateableTransactorSession) TransferOverride(from common.Address, to common.Address, tokens *big.Int) (*types.Transaction, error) {
	return _T0kenMigrateable.Contract.TransferOverride(&_T0kenMigrateable.TransactOpts, from, to, tokens)
}

// TransferOwner is a paid mutator transaction binding the contract method 0x4fb2e45d.
//
// Solidity: function transferOwner(newOwner address) returns()
func (_T0kenMigrateable *T0kenMigrateableTransactor) TransferOwner(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _T0kenMigrateable.contract.Transact(opts, "transferOwner", newOwner)
}

// TransferOwner is a paid mutator transaction binding the contract method 0x4fb2e45d.
//
// Solidity: function transferOwner(newOwner address) returns()
func (_T0kenMigrateable *T0kenMigrateableSession) TransferOwner(newOwner common.Address) (*types.Transaction, error) {
	return _T0kenMigrateable.Contract.TransferOwner(&_T0kenMigrateable.TransactOpts, newOwner)
}

// TransferOwner is a paid mutator transaction binding the contract method 0x4fb2e45d.
//
// Solidity: function transferOwner(newOwner address) returns()
func (_T0kenMigrateable *T0kenMigrateableTransactorSession) TransferOwner(newOwner common.Address) (*types.Transaction, error) {
	return _T0kenMigrateable.Contract.TransferOwner(&_T0kenMigrateable.TransactOpts, newOwner)
}

// T0kenMigrateableApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the T0kenMigrateable contract.
type T0kenMigrateableApprovalIterator struct {
	Event *T0kenMigrateableApproval // Event containing the contract specifics and raw log

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
func (it *T0kenMigrateableApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(T0kenMigrateableApproval)
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
		it.Event = new(T0kenMigrateableApproval)
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
func (it *T0kenMigrateableApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *T0kenMigrateableApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// T0kenMigrateableApproval represents a Approval event raised by the T0kenMigrateable contract.
type T0kenMigrateableApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: e Approval(owner indexed address, spender indexed address, value uint256)
func (_T0kenMigrateable *T0kenMigrateableFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*T0kenMigrateableApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _T0kenMigrateable.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &T0kenMigrateableApprovalIterator{contract: _T0kenMigrateable.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: e Approval(owner indexed address, spender indexed address, value uint256)
func (_T0kenMigrateable *T0kenMigrateableFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *T0kenMigrateableApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _T0kenMigrateable.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(T0kenMigrateableApproval)
				if err := _T0kenMigrateable.contract.UnpackLog(event, "Approval", log); err != nil {
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

// T0kenMigrateableIssuanceIterator is returned from FilterIssuance and is used to iterate over the raw logs and unpacked data for Issuance events raised by the T0kenMigrateable contract.
type T0kenMigrateableIssuanceIterator struct {
	Event *T0kenMigrateableIssuance // Event containing the contract specifics and raw log

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
func (it *T0kenMigrateableIssuanceIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(T0kenMigrateableIssuance)
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
		it.Event = new(T0kenMigrateableIssuance)
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
func (it *T0kenMigrateableIssuanceIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *T0kenMigrateableIssuanceIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// T0kenMigrateableIssuance represents a Issuance event raised by the T0kenMigrateable contract.
type T0kenMigrateableIssuance struct {
	To     common.Address
	Tokens *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterIssuance is a free log retrieval operation binding the contract event 0x9cb9c14f7bc76e3a89b796b091850526236115352a198b1e472f00e91376bbcb.
//
// Solidity: e Issuance(to indexed address, tokens uint256)
func (_T0kenMigrateable *T0kenMigrateableFilterer) FilterIssuance(opts *bind.FilterOpts, to []common.Address) (*T0kenMigrateableIssuanceIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _T0kenMigrateable.contract.FilterLogs(opts, "Issuance", toRule)
	if err != nil {
		return nil, err
	}
	return &T0kenMigrateableIssuanceIterator{contract: _T0kenMigrateable.contract, event: "Issuance", logs: logs, sub: sub}, nil
}

// WatchIssuance is a free log subscription operation binding the contract event 0x9cb9c14f7bc76e3a89b796b091850526236115352a198b1e472f00e91376bbcb.
//
// Solidity: e Issuance(to indexed address, tokens uint256)
func (_T0kenMigrateable *T0kenMigrateableFilterer) WatchIssuance(opts *bind.WatchOpts, sink chan<- *T0kenMigrateableIssuance, to []common.Address) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _T0kenMigrateable.contract.WatchLogs(opts, "Issuance", toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(T0kenMigrateableIssuance)
				if err := _T0kenMigrateable.contract.UnpackLog(event, "Issuance", log); err != nil {
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

// T0kenMigrateableIssuanceFinishedIterator is returned from FilterIssuanceFinished and is used to iterate over the raw logs and unpacked data for IssuanceFinished events raised by the T0kenMigrateable contract.
type T0kenMigrateableIssuanceFinishedIterator struct {
	Event *T0kenMigrateableIssuanceFinished // Event containing the contract specifics and raw log

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
func (it *T0kenMigrateableIssuanceFinishedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(T0kenMigrateableIssuanceFinished)
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
		it.Event = new(T0kenMigrateableIssuanceFinished)
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
func (it *T0kenMigrateableIssuanceFinishedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *T0kenMigrateableIssuanceFinishedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// T0kenMigrateableIssuanceFinished represents a IssuanceFinished event raised by the T0kenMigrateable contract.
type T0kenMigrateableIssuanceFinished struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterIssuanceFinished is a free log retrieval operation binding the contract event 0x29fe76cc5ca143e91eadf7242fda487fcef09318c1237900f958abe1e2c5beff.
//
// Solidity: e IssuanceFinished()
func (_T0kenMigrateable *T0kenMigrateableFilterer) FilterIssuanceFinished(opts *bind.FilterOpts) (*T0kenMigrateableIssuanceFinishedIterator, error) {

	logs, sub, err := _T0kenMigrateable.contract.FilterLogs(opts, "IssuanceFinished")
	if err != nil {
		return nil, err
	}
	return &T0kenMigrateableIssuanceFinishedIterator{contract: _T0kenMigrateable.contract, event: "IssuanceFinished", logs: logs, sub: sub}, nil
}

// WatchIssuanceFinished is a free log subscription operation binding the contract event 0x29fe76cc5ca143e91eadf7242fda487fcef09318c1237900f958abe1e2c5beff.
//
// Solidity: e IssuanceFinished()
func (_T0kenMigrateable *T0kenMigrateableFilterer) WatchIssuanceFinished(opts *bind.WatchOpts, sink chan<- *T0kenMigrateableIssuanceFinished) (event.Subscription, error) {

	logs, sub, err := _T0kenMigrateable.contract.WatchLogs(opts, "IssuanceFinished")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(T0kenMigrateableIssuanceFinished)
				if err := _T0kenMigrateable.contract.UnpackLog(event, "IssuanceFinished", log); err != nil {
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

// T0kenMigrateableIssuerSetIterator is returned from FilterIssuerSet and is used to iterate over the raw logs and unpacked data for IssuerSet events raised by the T0kenMigrateable contract.
type T0kenMigrateableIssuerSetIterator struct {
	Event *T0kenMigrateableIssuerSet // Event containing the contract specifics and raw log

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
func (it *T0kenMigrateableIssuerSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(T0kenMigrateableIssuerSet)
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
		it.Event = new(T0kenMigrateableIssuerSet)
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
func (it *T0kenMigrateableIssuerSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *T0kenMigrateableIssuerSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// T0kenMigrateableIssuerSet represents a IssuerSet event raised by the T0kenMigrateable contract.
type T0kenMigrateableIssuerSet struct {
	PreviousIssuer common.Address
	NewIssuer      common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterIssuerSet is a free log retrieval operation binding the contract event 0xf7189b85d7899f5a32d733e6584c4f1dcdff0274f09d969d186c1797673ede11.
//
// Solidity: e IssuerSet(previousIssuer indexed address, newIssuer indexed address)
func (_T0kenMigrateable *T0kenMigrateableFilterer) FilterIssuerSet(opts *bind.FilterOpts, previousIssuer []common.Address, newIssuer []common.Address) (*T0kenMigrateableIssuerSetIterator, error) {

	var previousIssuerRule []interface{}
	for _, previousIssuerItem := range previousIssuer {
		previousIssuerRule = append(previousIssuerRule, previousIssuerItem)
	}
	var newIssuerRule []interface{}
	for _, newIssuerItem := range newIssuer {
		newIssuerRule = append(newIssuerRule, newIssuerItem)
	}

	logs, sub, err := _T0kenMigrateable.contract.FilterLogs(opts, "IssuerSet", previousIssuerRule, newIssuerRule)
	if err != nil {
		return nil, err
	}
	return &T0kenMigrateableIssuerSetIterator{contract: _T0kenMigrateable.contract, event: "IssuerSet", logs: logs, sub: sub}, nil
}

// WatchIssuerSet is a free log subscription operation binding the contract event 0xf7189b85d7899f5a32d733e6584c4f1dcdff0274f09d969d186c1797673ede11.
//
// Solidity: e IssuerSet(previousIssuer indexed address, newIssuer indexed address)
func (_T0kenMigrateable *T0kenMigrateableFilterer) WatchIssuerSet(opts *bind.WatchOpts, sink chan<- *T0kenMigrateableIssuerSet, previousIssuer []common.Address, newIssuer []common.Address) (event.Subscription, error) {

	var previousIssuerRule []interface{}
	for _, previousIssuerItem := range previousIssuer {
		previousIssuerRule = append(previousIssuerRule, previousIssuerItem)
	}
	var newIssuerRule []interface{}
	for _, newIssuerItem := range newIssuer {
		newIssuerRule = append(newIssuerRule, newIssuerItem)
	}

	logs, sub, err := _T0kenMigrateable.contract.WatchLogs(opts, "IssuerSet", previousIssuerRule, newIssuerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(T0kenMigrateableIssuerSet)
				if err := _T0kenMigrateable.contract.UnpackLog(event, "IssuerSet", log); err != nil {
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

// T0kenMigrateableOwnerTransferredIterator is returned from FilterOwnerTransferred and is used to iterate over the raw logs and unpacked data for OwnerTransferred events raised by the T0kenMigrateable contract.
type T0kenMigrateableOwnerTransferredIterator struct {
	Event *T0kenMigrateableOwnerTransferred // Event containing the contract specifics and raw log

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
func (it *T0kenMigrateableOwnerTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(T0kenMigrateableOwnerTransferred)
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
		it.Event = new(T0kenMigrateableOwnerTransferred)
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
func (it *T0kenMigrateableOwnerTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *T0kenMigrateableOwnerTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// T0kenMigrateableOwnerTransferred represents a OwnerTransferred event raised by the T0kenMigrateable contract.
type T0kenMigrateableOwnerTransferred struct {
	OldOwner common.Address
	NewOwner common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOwnerTransferred is a free log retrieval operation binding the contract event 0x8934ce4adea8d9ce0d714d2c22b86790e41b7731c84b926fbbdc1d40ff6533c9.
//
// Solidity: e OwnerTransferred(oldOwner indexed address, newOwner indexed address)
func (_T0kenMigrateable *T0kenMigrateableFilterer) FilterOwnerTransferred(opts *bind.FilterOpts, oldOwner []common.Address, newOwner []common.Address) (*T0kenMigrateableOwnerTransferredIterator, error) {

	var oldOwnerRule []interface{}
	for _, oldOwnerItem := range oldOwner {
		oldOwnerRule = append(oldOwnerRule, oldOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _T0kenMigrateable.contract.FilterLogs(opts, "OwnerTransferred", oldOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &T0kenMigrateableOwnerTransferredIterator{contract: _T0kenMigrateable.contract, event: "OwnerTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnerTransferred is a free log subscription operation binding the contract event 0x8934ce4adea8d9ce0d714d2c22b86790e41b7731c84b926fbbdc1d40ff6533c9.
//
// Solidity: e OwnerTransferred(oldOwner indexed address, newOwner indexed address)
func (_T0kenMigrateable *T0kenMigrateableFilterer) WatchOwnerTransferred(opts *bind.WatchOpts, sink chan<- *T0kenMigrateableOwnerTransferred, oldOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var oldOwnerRule []interface{}
	for _, oldOwnerItem := range oldOwner {
		oldOwnerRule = append(oldOwnerRule, oldOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _T0kenMigrateable.contract.WatchLogs(opts, "OwnerTransferred", oldOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(T0kenMigrateableOwnerTransferred)
				if err := _T0kenMigrateable.contract.UnpackLog(event, "OwnerTransferred", log); err != nil {
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

// T0kenMigrateablePrecededByIterator is returned from FilterPrecededBy and is used to iterate over the raw logs and unpacked data for PrecededBy events raised by the T0kenMigrateable contract.
type T0kenMigrateablePrecededByIterator struct {
	Event *T0kenMigrateablePrecededBy // Event containing the contract specifics and raw log

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
func (it *T0kenMigrateablePrecededByIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(T0kenMigrateablePrecededBy)
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
		it.Event = new(T0kenMigrateablePrecededBy)
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
func (it *T0kenMigrateablePrecededByIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *T0kenMigrateablePrecededByIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// T0kenMigrateablePrecededBy represents a PrecededBy event raised by the T0kenMigrateable contract.
type T0kenMigrateablePrecededBy struct {
	Predecessor common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterPrecededBy is a free log retrieval operation binding the contract event 0x4779b83fa355cab77653ce138f51eb963e2d7c4632c698e9c77ee1ddb920d56d.
//
// Solidity: e PrecededBy(predecessor indexed address)
func (_T0kenMigrateable *T0kenMigrateableFilterer) FilterPrecededBy(opts *bind.FilterOpts, predecessor []common.Address) (*T0kenMigrateablePrecededByIterator, error) {

	var predecessorRule []interface{}
	for _, predecessorItem := range predecessor {
		predecessorRule = append(predecessorRule, predecessorItem)
	}

	logs, sub, err := _T0kenMigrateable.contract.FilterLogs(opts, "PrecededBy", predecessorRule)
	if err != nil {
		return nil, err
	}
	return &T0kenMigrateablePrecededByIterator{contract: _T0kenMigrateable.contract, event: "PrecededBy", logs: logs, sub: sub}, nil
}

// WatchPrecededBy is a free log subscription operation binding the contract event 0x4779b83fa355cab77653ce138f51eb963e2d7c4632c698e9c77ee1ddb920d56d.
//
// Solidity: e PrecededBy(predecessor indexed address)
func (_T0kenMigrateable *T0kenMigrateableFilterer) WatchPrecededBy(opts *bind.WatchOpts, sink chan<- *T0kenMigrateablePrecededBy, predecessor []common.Address) (event.Subscription, error) {

	var predecessorRule []interface{}
	for _, predecessorItem := range predecessor {
		predecessorRule = append(predecessorRule, predecessorItem)
	}

	logs, sub, err := _T0kenMigrateable.contract.WatchLogs(opts, "PrecededBy", predecessorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(T0kenMigrateablePrecededBy)
				if err := _T0kenMigrateable.contract.UnpackLog(event, "PrecededBy", log); err != nil {
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

// T0kenMigrateableShareholderAddedIterator is returned from FilterShareholderAdded and is used to iterate over the raw logs and unpacked data for ShareholderAdded events raised by the T0kenMigrateable contract.
type T0kenMigrateableShareholderAddedIterator struct {
	Event *T0kenMigrateableShareholderAdded // Event containing the contract specifics and raw log

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
func (it *T0kenMigrateableShareholderAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(T0kenMigrateableShareholderAdded)
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
		it.Event = new(T0kenMigrateableShareholderAdded)
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
func (it *T0kenMigrateableShareholderAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *T0kenMigrateableShareholderAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// T0kenMigrateableShareholderAdded represents a ShareholderAdded event raised by the T0kenMigrateable contract.
type T0kenMigrateableShareholderAdded struct {
	Shareholder common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterShareholderAdded is a free log retrieval operation binding the contract event 0x3082c1c4977de80c4f67ee77b56b282610ec93a7ecbcc31b551e0ac2f7bd0395.
//
// Solidity: e ShareholderAdded(shareholder address)
func (_T0kenMigrateable *T0kenMigrateableFilterer) FilterShareholderAdded(opts *bind.FilterOpts) (*T0kenMigrateableShareholderAddedIterator, error) {

	logs, sub, err := _T0kenMigrateable.contract.FilterLogs(opts, "ShareholderAdded")
	if err != nil {
		return nil, err
	}
	return &T0kenMigrateableShareholderAddedIterator{contract: _T0kenMigrateable.contract, event: "ShareholderAdded", logs: logs, sub: sub}, nil
}

// WatchShareholderAdded is a free log subscription operation binding the contract event 0x3082c1c4977de80c4f67ee77b56b282610ec93a7ecbcc31b551e0ac2f7bd0395.
//
// Solidity: e ShareholderAdded(shareholder address)
func (_T0kenMigrateable *T0kenMigrateableFilterer) WatchShareholderAdded(opts *bind.WatchOpts, sink chan<- *T0kenMigrateableShareholderAdded) (event.Subscription, error) {

	logs, sub, err := _T0kenMigrateable.contract.WatchLogs(opts, "ShareholderAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(T0kenMigrateableShareholderAdded)
				if err := _T0kenMigrateable.contract.UnpackLog(event, "ShareholderAdded", log); err != nil {
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

// T0kenMigrateableShareholderRemovedIterator is returned from FilterShareholderRemoved and is used to iterate over the raw logs and unpacked data for ShareholderRemoved events raised by the T0kenMigrateable contract.
type T0kenMigrateableShareholderRemovedIterator struct {
	Event *T0kenMigrateableShareholderRemoved // Event containing the contract specifics and raw log

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
func (it *T0kenMigrateableShareholderRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(T0kenMigrateableShareholderRemoved)
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
		it.Event = new(T0kenMigrateableShareholderRemoved)
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
func (it *T0kenMigrateableShareholderRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *T0kenMigrateableShareholderRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// T0kenMigrateableShareholderRemoved represents a ShareholderRemoved event raised by the T0kenMigrateable contract.
type T0kenMigrateableShareholderRemoved struct {
	Shareholder common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterShareholderRemoved is a free log retrieval operation binding the contract event 0x7ba114ff3d9844510a088eea94cd35562e7c97a2d36a418b37b2e61e5c77affe.
//
// Solidity: e ShareholderRemoved(shareholder address)
func (_T0kenMigrateable *T0kenMigrateableFilterer) FilterShareholderRemoved(opts *bind.FilterOpts) (*T0kenMigrateableShareholderRemovedIterator, error) {

	logs, sub, err := _T0kenMigrateable.contract.FilterLogs(opts, "ShareholderRemoved")
	if err != nil {
		return nil, err
	}
	return &T0kenMigrateableShareholderRemovedIterator{contract: _T0kenMigrateable.contract, event: "ShareholderRemoved", logs: logs, sub: sub}, nil
}

// WatchShareholderRemoved is a free log subscription operation binding the contract event 0x7ba114ff3d9844510a088eea94cd35562e7c97a2d36a418b37b2e61e5c77affe.
//
// Solidity: e ShareholderRemoved(shareholder address)
func (_T0kenMigrateable *T0kenMigrateableFilterer) WatchShareholderRemoved(opts *bind.WatchOpts, sink chan<- *T0kenMigrateableShareholderRemoved) (event.Subscription, error) {

	logs, sub, err := _T0kenMigrateable.contract.WatchLogs(opts, "ShareholderRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(T0kenMigrateableShareholderRemoved)
				if err := _T0kenMigrateable.contract.UnpackLog(event, "ShareholderRemoved", log); err != nil {
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

// T0kenMigrateableSucceededByIterator is returned from FilterSucceededBy and is used to iterate over the raw logs and unpacked data for SucceededBy events raised by the T0kenMigrateable contract.
type T0kenMigrateableSucceededByIterator struct {
	Event *T0kenMigrateableSucceededBy // Event containing the contract specifics and raw log

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
func (it *T0kenMigrateableSucceededByIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(T0kenMigrateableSucceededBy)
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
		it.Event = new(T0kenMigrateableSucceededBy)
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
func (it *T0kenMigrateableSucceededByIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *T0kenMigrateableSucceededByIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// T0kenMigrateableSucceededBy represents a SucceededBy event raised by the T0kenMigrateable contract.
type T0kenMigrateableSucceededBy struct {
	Successor common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterSucceededBy is a free log retrieval operation binding the contract event 0x86640a0f4d220948cb9c4dcba2cd1269adf1171c504757c96a10448c9285d424.
//
// Solidity: e SucceededBy(successor indexed address)
func (_T0kenMigrateable *T0kenMigrateableFilterer) FilterSucceededBy(opts *bind.FilterOpts, successor []common.Address) (*T0kenMigrateableSucceededByIterator, error) {

	var successorRule []interface{}
	for _, successorItem := range successor {
		successorRule = append(successorRule, successorItem)
	}

	logs, sub, err := _T0kenMigrateable.contract.FilterLogs(opts, "SucceededBy", successorRule)
	if err != nil {
		return nil, err
	}
	return &T0kenMigrateableSucceededByIterator{contract: _T0kenMigrateable.contract, event: "SucceededBy", logs: logs, sub: sub}, nil
}

// WatchSucceededBy is a free log subscription operation binding the contract event 0x86640a0f4d220948cb9c4dcba2cd1269adf1171c504757c96a10448c9285d424.
//
// Solidity: e SucceededBy(successor indexed address)
func (_T0kenMigrateable *T0kenMigrateableFilterer) WatchSucceededBy(opts *bind.WatchOpts, sink chan<- *T0kenMigrateableSucceededBy, successor []common.Address) (event.Subscription, error) {

	var successorRule []interface{}
	for _, successorItem := range successor {
		successorRule = append(successorRule, successorItem)
	}

	logs, sub, err := _T0kenMigrateable.contract.WatchLogs(opts, "SucceededBy", successorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(T0kenMigrateableSucceededBy)
				if err := _T0kenMigrateable.contract.UnpackLog(event, "SucceededBy", log); err != nil {
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

// T0kenMigrateableTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the T0kenMigrateable contract.
type T0kenMigrateableTransferIterator struct {
	Event *T0kenMigrateableTransfer // Event containing the contract specifics and raw log

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
func (it *T0kenMigrateableTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(T0kenMigrateableTransfer)
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
		it.Event = new(T0kenMigrateableTransfer)
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
func (it *T0kenMigrateableTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *T0kenMigrateableTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// T0kenMigrateableTransfer represents a Transfer event raised by the T0kenMigrateable contract.
type T0kenMigrateableTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: e Transfer(from indexed address, to indexed address, value uint256)
func (_T0kenMigrateable *T0kenMigrateableFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*T0kenMigrateableTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _T0kenMigrateable.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &T0kenMigrateableTransferIterator{contract: _T0kenMigrateable.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: e Transfer(from indexed address, to indexed address, value uint256)
func (_T0kenMigrateable *T0kenMigrateableFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *T0kenMigrateableTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _T0kenMigrateable.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(T0kenMigrateableTransfer)
				if err := _T0kenMigrateable.contract.UnpackLog(event, "Transfer", log); err != nil {
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
