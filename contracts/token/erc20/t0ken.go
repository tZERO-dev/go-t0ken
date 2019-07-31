// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package erc20

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

// T0kenABI is the input ABI used to generate the binding from.
const T0kenABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"spender\",\"type\":\"address\"},{\"name\":\"tokens\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"issuer\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"locked\",\"type\":\"bool\"}],\"name\":\"setLocked\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"from\",\"type\":\"address\"},{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"tokens\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"balances\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"isSuperseded\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"cancellations\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"kill\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"issuanceFinished\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwner\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ZERO_ADDRESS\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newIssuer\",\"type\":\"address\"}],\"name\":\"setIssuer\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"compliance\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"original\",\"type\":\"address\"},{\"name\":\"replacement\",\"type\":\"address\"}],\"name\":\"cancelAndReissue\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"from\",\"type\":\"address\"},{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"tokens\",\"type\":\"uint256\"}],\"name\":\"transferOverride\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"index\",\"type\":\"int256\"}],\"name\":\"holderAt\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"holders\",\"outputs\":[{\"name\":\"count\",\"type\":\"int256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isLocked\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"quantity\",\"type\":\"uint256\"}],\"name\":\"issueTokens\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"tokens\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"finishIssuance\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"isHolder\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"addrOwner\",\"type\":\"address\"},{\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"getSuperseded\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newComplianceAddress\",\"type\":\"address\"}],\"name\":\"setCompliance\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"tokenName\",\"type\":\"string\"},{\"name\":\"tokenSymbol\",\"type\":\"string\"},{\"name\":\"decimalPlaces\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"oldOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnerTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"original\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"replacement\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"VerifiedAddressSuperseded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousIssuer\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newIssuer\",\"type\":\"address\"}],\"name\":\"IssuerSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"tokens\",\"type\":\"uint256\"}],\"name\":\"Issuance\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"IssuanceFinished\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"shareholder\",\"type\":\"address\"}],\"name\":\"ShareholderAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"shareholder\",\"type\":\"address\"}],\"name\":\"ShareholderRemoved\",\"type\":\"event\"}]"

// T0kenBin is the compiled bytecode used for deploying new contracts.
const T0kenBin = `6080604052600180546001600160a01b03191690553480156200002157600080fd5b50604051620023ea380380620023ea833981810160405260608110156200004757600080fd5b8101908080516401000000008111156200006057600080fd5b820160208101848111156200007457600080fd5b81516401000000008111828201871017156200008f57600080fd5b50509291906020018051640100000000811115620000ac57600080fd5b82016020810184811115620000c057600080fd5b8151640100000000811182820187101715620000db57600080fd5b5050602091820151600080546001600160a01b031916331790556001805460ff60a01b19169055855191945092506200011b91600291908601906200014f565b508151620001319060039060208501906200014f565b506004805460ff191660ff9290921691909117905550620001f49050565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106200019257805160ff1916838001178555620001c2565b82800160010185558215620001c2579182015b82811115620001c2578251825591602001919060010190620001a5565b50620001d0929150620001d4565b5090565b620001f191905b80821115620001d05760008155600101620001db565b90565b6121e680620002046000396000f3fe608060405234801561001057600080fd5b50600436106101e55760003560e01c80636290865d1161010f578063a4e2d634116100a2578063d4d7b19a11610071578063d4d7b19a14610550578063dd62ed3e14610576578063e37ccac7146105a4578063f8981789146105ca576101e5565b8063a4e2d634146104f7578063a5820daa146104ff578063a9059cbb1461051c578063c422293b14610548576101e5565b80638082a929116100de5780638082a929146104c25780638188f71c146104df5780638da5cb5b146104e757806395d89b41146104ef576101e5565b80636290865d1461043057806370a082311461043857806379f647201461045e57806380318be81461048c576101e5565b80632da7293e116101875780634662299a116101565780634662299a146103d45780634fb2e45d146103dc578063538ba4f91461040257806355cc4e571461040a576101e5565b80632da7293e14610362578063313ce5671461038857806334a84827146103a657806341c0e1b5146103cc576101e5565b80631d143848116101c35780631d143848146102c1578063211e28b6146102e557806323b872dd1461030657806327e235e31461033c576101e5565b806306fdde03146101ea578063095ea7b31461026757806318160ddd146102a7575b600080fd5b6101f26105f0565b6040805160208082528351818301528351919283929083019185019080838360005b8381101561022c578181015183820152602001610214565b50505050905090810190601f1680156102595780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b6102936004803603604081101561027d57600080fd5b506001600160a01b03813516906020013561067b565b604080519115158252519081900360200190f35b6102af6107fc565b60408051918252519081900360200190f35b6102c9610802565b604080516001600160a01b039092168252519081900360200190f35b610304600480360360208110156102fb57600080fd5b50351515610811565b005b6102936004803603606081101561031c57600080fd5b506001600160a01b038135811691602081013590911690604001356108f9565b6102af6004803603602081101561035257600080fd5b50356001600160a01b0316610b17565b6102936004803603602081101561037857600080fd5b50356001600160a01b0316610b29565b610390610b4c565b6040805160ff9092168252519081900360200190f35b6102c9600480360360208110156103bc57600080fd5b50356001600160a01b0316610b55565b610304610b70565b610293610bf8565b610304600480360360208110156103f257600080fd5b50356001600160a01b0316610c08565b6102c9610d7a565b6103046004803603602081101561042057600080fd5b50356001600160a01b0316610d89565b6102c9610f10565b6102af6004803603602081101561044e57600080fd5b50356001600160a01b0316610f1f565b6103046004803603604081101561047457600080fd5b506001600160a01b0381358116916020013516610f3a565b610293600480360360608110156104a257600080fd5b506001600160a01b03813581169160208101359091169060400135611250565b6102c9600480360360208110156104d857600080fd5b5035611393565b6102af6113ac565b6102c96113b2565b6101f26113c1565b61029361141c565b6102936004803603602081101561051557600080fd5b503561142c565b6102936004803603604081101561053257600080fd5b506001600160a01b038135169060200135611644565b61029361184a565b6102936004803603602081101561056657600080fd5b50356001600160a01b03166119a1565b6102af6004803603604081101561058c57600080fd5b506001600160a01b03813581169160200135166119b4565b6102c9600480360360208110156105ba57600080fd5b50356001600160a01b03166119df565b610304600480360360208110156105e057600080fd5b50356001600160a01b0316611a0b565b6002805460408051602060018416156101000260001901909316849004601f810184900484028201840190925281815292918301828280156106735780601f1061064857610100808354040283529160200191610673565b820191906000526020600020905b81548152906001019060200180831161065657829003601f168201915b505050505081565b600154600090600160a01b900460ff16156106c75760405162461bcd60e51b815260040180806020018281038252602d815260200180612185602d913960400191505060405180910390fd5b336000818152600a60205260409020546001600160a01b031615610732576040805162461bcd60e51b815260206004820152601a60248201527f4164647265737320686173206265656e2063616e63656c6c6564000000000000604482015290519081900360640190fd5b61074360053363ffffffff611af016565b610794576040805162461bcd60e51b815260206004820152601560248201527f4d7573742062652061207368617265686f6c6465720000000000000000000000604482015290519081900360640190fd5b336000818152600d602090815260408083206001600160a01b03891680855290835292819020879055805187815290519293927f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925929181900390910190a35060019392505050565b600c5490565b6009546001600160a01b031681565b6000546001600160a01b031633148061083a57506001546000546001600160a01b039081169116145b61088b576040805162461bcd60e51b815260206004820152601960248201527f4f776e6572206163636f756e7420697320726571756972656400000000000000604482015290519081900360640190fd5b60015460ff600160a01b90910416151581151514156108db5760405162461bcd60e51b815260040180806020018281038252602881526020018061215d6028913960400191505060405180910390fd5b60018054911515600160a01b0260ff60a01b19909216919091179055565b600154600090600160a01b900460ff16156109455760405162461bcd60e51b815260040180806020018281038252602d815260200180612185602d913960400191505060405180910390fd5b6001600160a01b038084166000908152600a6020526040902054849116156109b4576040805162461bcd60e51b815260206004820152601a60248201527f4164647265737320686173206265656e2063616e63656c6c6564000000000000604482015290519081900360640190fd5b6001600160a01b0385166000908152600b602052604090205485908490811115610a1a576040805162461bcd60e51b8152602060048201526012602482015271496e73756666696369656e742066756e647360701b604482015290519081900360640190fd5b6001600160a01b0387166000908152600d60209081526040808320338452909152902054851115610a92576040805162461bcd60e51b815260206004820152601a60248201527f5472616e73666572206578636565647320616c6c6f77616e6365000000000000604482015290519081900360640190fd5b6000610aa18888886000611b27565b90508015610b0c576001600160a01b0388166000908152600d60209081526040808320338452909152902054610add908763ffffffff611c7516565b6001600160a01b0389166000908152600d60209081526040808320338452909152902055610b0c888888611cd2565b979650505050505050565b600b6020526000908152604090205481565b6001600160a01b038181166000908152600a60205260409020541615155b919050565b60045460ff1681565b600a602052600090815260409020546001600160a01b031681565b6000546001600160a01b0316331480610b9957506001546000546001600160a01b039081169116145b610bea576040805162461bcd60e51b815260206004820152601960248201527f4f776e6572206163636f756e7420697320726571756972656400000000000000604482015290519081900360640190fd5b6000546001600160a01b0316ff5b600954600160a01b900460ff1681565b6000546001600160a01b0316331480610c3157506001546000546001600160a01b039081169116145b610c82576040805162461bcd60e51b815260206004820152601960248201527f4f776e6572206163636f756e7420697320726571756972656400000000000000604482015290519081900360640190fd5b6000546001600160a01b0382811691161415610ccf5760405162461bcd60e51b81526004018080602001828103825260258152602001806121386025913960400191505060405180910390fd5b6001600160a01b038116610d2a576040805162461bcd60e51b815260206004820181905260248201527f4e6577204f776e65722063616e6e6f74206265207a65726f2061646472657373604482015290519081900360640190fd5b600080546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8934ce4adea8d9ce0d714d2c22b86790e41b7731c84b926fbbdc1d40ff6533c99190a35050565b6001546001600160a01b031681565b600154600160a01b900460ff1615610dd25760405162461bcd60e51b815260040180806020018281038252602d815260200180612185602d913960400191505060405180910390fd5b6000546001600160a01b0316331480610dfb57506001546000546001600160a01b039081169116145b610e4c576040805162461bcd60e51b815260206004820152601960248201527f4f776e6572206163636f756e7420697320726571756972656400000000000000604482015290519081900360640190fd5b610e5d60058263ffffffff611af016565b15610e995760405162461bcd60e51b81526004018080602001828103825260218152602001806120ed6021913960400191505060405180910390fd5b6009546001600160a01b03166000818152600b6020526040902054610ec091908390611cd2565b600980546001600160a01b0319166001600160a01b0383811691821792839055604051919216907ff7189b85d7899f5a32d733e6584c4f1dcdff0274f09d969d186c1797673ede1190600090a350565b6008546001600160a01b031681565b6001600160a01b03166000908152600b602052604090205490565b600154600160a01b900460ff1615610f835760405162461bcd60e51b815260040180806020018281038252602d815260200180612185602d913960400191505060405180910390fd5b6009546001600160a01b03163314610fe2576040805162461bcd60e51b815260206004820152601360248201527f4f6e6c792069737375657220616c6c6f77656400000000000000000000000000604482015290519081900360640190fd5b6001600160a01b038082166000908152600a602052604090205482911615611051576040805162461bcd60e51b815260206004820152601a60248201527f4164647265737320686173206265656e2063616e63656c6c6564000000000000604482015290519081900360640190fd5b61106260058463ffffffff611af016565b801561107c575061107a60058363ffffffff611af016565b155b6110b75760405162461bcd60e51b815260040180806020018281038252602a81526020018061210e602a913960400191505060405180910390fd5b6008546001600160a01b0316156111b8576008546001600160a01b038481166000818152600b6020908152604080832054815163fd8258bd60e01b815233600482015260248101959095528886166044860152606485015251939094169363fd8258bd93608480850194929391928390030190829087803b15801561113b57600080fd5b505af115801561114f573d6000803e3d6000fd5b505050506040513d602081101561116557600080fd5b50516111b8576040805162461bcd60e51b815260206004820152601760248201527f4661696c656420636f6d706c69616e636520636865636b000000000000000000604482015290519081900360640190fd5b6111c960058463ffffffff611e7516565b506111db60058363ffffffff611f6016565b506001600160a01b038381166000818152600a6020908152604080832080546001600160a01b0319169588169586179055600b90915280822080548584528284205583835282905551339392917fb64971100522354f3d25283cb14e2eefcb0dd26a757482ccfe42479d0a68685791a4505050565b600154600090600160a01b900460ff161561129c5760405162461bcd60e51b815260040180806020018281038252602d815260200180612185602d913960400191505060405180910390fd5b6001600160a01b038084166000908152600a60205260409020548491161561130b576040805162461bcd60e51b815260206004820152601a60248201527f4164647265737320686173206265656e2063616e63656c6c6564000000000000604482015290519081900360640190fd5b6001600160a01b0385166000908152600b602052604090205485908490811115611371576040805162461bcd60e51b8152602060048201526012602482015271496e73756666696369656e742066756e647360701b604482015290519081900360640190fd5b60006113808888886001611b27565b90508015610b0c57610b0c888888611cd2565b60006113a660058363ffffffff61200516565b92915050565b60055481565b6000546001600160a01b031681565b6003805460408051602060026001851615610100026000190190941693909304601f810184900484028201840190925281815292918301828280156106735780601f1061064857610100808354040283529160200191610673565b600154600160a01b900460ff1681565b600154600090600160a01b900460ff16156114785760405162461bcd60e51b815260040180806020018281038252602d815260200180612185602d913960400191505060405180910390fd5b6009546001600160a01b031633146114d7576040805162461bcd60e51b815260206004820152601360248201527f4f6e6c792069737375657220616c6c6f77656400000000000000000000000000604482015290519081900360640190fd5b600954600160a01b900460ff1615611536576040805162461bcd60e51b815260206004820152601960248201527f49737375616e636520616c72656164792066696e697368656400000000000000604482015290519081900360640190fd5b81156115b457600c5461154f908363ffffffff61208b16565b600c556009546001600160a01b03166000908152600b602052604090205461157d908363ffffffff61208b16565b600980546001600160a01b039081166000908152600b6020526040902092909255546115b2916005911663ffffffff611f6016565b505b6009546040805184815290516001600160a01b03909216917f9cb9c14f7bc76e3a89b796b091850526236115352a198b1e472f00e91376bbcb9181900360200190a26009546040805184815290516001600160a01b03909216916000917fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef919081900360200190a3506001919050565b600154600090600160a01b900460ff16156116905760405162461bcd60e51b815260040180806020018281038252602d815260200180612185602d913960400191505060405180910390fd5b6001600160a01b038084166000908152600a6020526040902054849116156116ff576040805162461bcd60e51b815260206004820152601a60248201527f4164647265737320686173206265656e2063616e63656c6c6564000000000000604482015290519081900360640190fd5b336000818152600b6020526040902054849081111561175a576040805162461bcd60e51b8152602060048201526012602482015271496e73756666696369656e742066756e647360701b604482015290519081900360640190fd5b6009546000906001600160a01b031633141561181f57506008546001600160a01b0316158061181a576008546009546040805163fd8258bd60e01b81526001600160a01b039283166004820181905260248201528a83166044820152606481018a90529051919092169163fd8258bd9160848083019260209291908290030181600087803b1580156117eb57600080fd5b505af11580156117ff573d6000803e3d6000fd5b505050506040513d602081101561181557600080fd5b505190505b61182f565b61182c3388886000611b27565b90505b801561184057611840338888611cd2565b9695505050505050565b600154600090600160a01b900460ff16156118965760405162461bcd60e51b815260040180806020018281038252602d815260200180612185602d913960400191505060405180910390fd5b6009546001600160a01b031633146118f5576040805162461bcd60e51b815260206004820152601360248201527f4f6e6c792069737375657220616c6c6f77656400000000000000000000000000604482015290519081900360640190fd5b600954600160a01b900460ff1615611954576040805162461bcd60e51b815260206004820152601960248201527f49737375616e636520616c72656164792066696e697368656400000000000000604482015290519081900360640190fd5b6009805460ff60a01b1916600160a01b1790556040517f29fe76cc5ca143e91eadf7242fda487fcef09318c1237900f958abe1e2c5beff90600090a150600954600160a01b900460ff1690565b60006113a660058363ffffffff611af016565b6001600160a01b039182166000908152600d6020908152604080832093909416825291909152205490565b6001600160a01b038082166000908152600a6020526040812054909116806113a6576000915050610b47565b600154600160a01b900460ff1615611a545760405162461bcd60e51b815260040180806020018281038252602d815260200180612185602d913960400191505060405180910390fd5b6000546001600160a01b0316331480611a7d57506001546000546001600160a01b039081169116145b611ace576040805162461bcd60e51b815260206004820152601960248201527f4f776e6572206163636f756e7420697320726571756972656400000000000000604482015290519081900360640190fd5b600880546001600160a01b0319166001600160a01b0392909216919091179055565b6001600160a01b038116600090815260018301602052604081205460001901818112801590611b1f5750835481125b949350505050565b6008546000906001600160a01b0316611b4257508015611b1f565b8115611bf857600854604080517f5acba2010000000000000000000000000000000000000000000000000000000081523360048201526001600160a01b03888116602483015287811660448301526064820187905291519190921691635acba2019160848083019260209291908290030181600087803b158015611bc557600080fd5b505af1158015611bd9573d6000803e3d6000fd5b505050506040513d6020811015611bef57600080fd5b50519050611b1f565b600854604080517f6d62a4fe0000000000000000000000000000000000000000000000000000000081523360048201526001600160a01b03888116602483015287811660448301526064820187905291519190921691636d62a4fe9160848083019260209291908290030181600087803b158015611bc557600080fd5b600082821115611ccc576040805162461bcd60e51b815260206004820152601460248201527f526573756c747320696e20756e646572666c6f77000000000000000000000000604482015290519081900360640190fd5b50900390565b6001600160a01b0383166000908152600b6020526040902054611cfb908263ffffffff611c7516565b6001600160a01b038085166000908152600b60205260408082209390935590841681522054611d30908263ffffffff61208b16565b6001600160a01b038084166000818152600b602090815260409182902094909455805185815290519193928716927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef92918290030190a36001600160a01b0382166000908152600b602052604090205415801590611dba5750611dba60058363ffffffff611f6016565b15611dfc57604080516001600160a01b038416815290517f3082c1c4977de80c4f67ee77b56b282610ec93a7ecbcc31b551e0ac2f7bd03959181900360200190a15b6001600160a01b0383166000908152600b6020526040902054158015611e2e5750611e2e60058463ffffffff611e7516565b15611e7057604080516001600160a01b038516815290517f7ba114ff3d9844510a088eea94cd35562e7c97a2d36a418b37b2e61e5c77affe9181900360200190a15b505050565b6001600160a01b03811660009081526001808401602052604082205490811280611e9f5750835481135b15611eae5760009150506113a6565b8354811215611f1557835460009081526002850160208181526040808420546001600160a01b03168085526001890183528185208690558585529290915280832080546001600160a01b031990811690931790558654835290912080549091169055611f34565b6000818152600285016020526040902080546001600160a01b03191690555b50506001600160a01b031660009081526001828101602052604082209190915581546000190190915590565b60006001600160a01b038216611f78575060006113a6565b6001600160a01b038216600090815260018401602052604081205460001901908112801590611fa75750835481125b15611fb65760009150506113a6565b5050815460019081018084556001600160a01b0383166000818152838601602090815260408083208590559382526002870190529190912080546001600160a01b031916909117905592915050565b60008082121580156120175750825482125b612068576040805162461bcd60e51b815260206004820152601860248201527f496e646578206f757473696465206f6620626f756e64732e0000000000000000604482015290519081900360640190fd5b50600101600090815260029190910160205260409020546001600160a01b031690565b6000828201838110156120e5576040805162461bcd60e51b815260206004820152601360248201527f526573756c747320696e206f766572666c6f7700000000000000000000000000604482015290519081900360640190fd5b939250505056fe4e6577206973737565722063616e27742062652061207368617265686f6c6465724f726967696e616c20646f65736e2774206578697374206f72207265706c6163656d656e7420646f65734e6577204f776e65722063616e6e6f74206265207468652063757272656e74206f776e6572436f6e747261637420616c726561647920696e20726571756573746564206c6f636b207374617465436f6e74726163742069732063757272656e746c79206c6f636b656420666f72206d6f64696669636174696f6ea265627a7a72305820bcce1c8f0d619c4f9be7adf8b086cc509349603d8960acc5ac10131f992ae1f564736f6c634300050a0032`

// DeployT0ken deploys a new Ethereum contract, binding an instance of T0ken to it.
func DeployT0ken(auth *bind.TransactOpts, backend bind.ContractBackend, tokenName string, tokenSymbol string, decimalPlaces uint8) (common.Address, *types.Transaction, *T0ken, error) {
	parsed, err := abi.JSON(strings.NewReader(T0kenABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(T0kenBin), backend, tokenName, tokenSymbol, decimalPlaces)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &T0ken{T0kenCaller: T0kenCaller{contract: contract}, T0kenTransactor: T0kenTransactor{contract: contract}, T0kenFilterer: T0kenFilterer{contract: contract}}, nil
}

// T0ken is an auto generated Go binding around an Ethereum contract.
type T0ken struct {
	T0kenCaller     // Read-only binding to the contract
	T0kenTransactor // Write-only binding to the contract
	T0kenFilterer   // Log filterer for contract events
}

// T0kenCaller is an auto generated read-only Go binding around an Ethereum contract.
type T0kenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// T0kenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type T0kenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// T0kenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type T0kenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// T0kenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type T0kenSession struct {
	Contract     *T0ken            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// T0kenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type T0kenCallerSession struct {
	Contract *T0kenCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// T0kenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type T0kenTransactorSession struct {
	Contract     *T0kenTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// T0kenRaw is an auto generated low-level Go binding around an Ethereum contract.
type T0kenRaw struct {
	Contract *T0ken // Generic contract binding to access the raw methods on
}

// T0kenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type T0kenCallerRaw struct {
	Contract *T0kenCaller // Generic read-only contract binding to access the raw methods on
}

// T0kenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type T0kenTransactorRaw struct {
	Contract *T0kenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewT0ken creates a new instance of T0ken, bound to a specific deployed contract.
func NewT0ken(address common.Address, backend bind.ContractBackend) (*T0ken, error) {
	contract, err := bindT0ken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &T0ken{T0kenCaller: T0kenCaller{contract: contract}, T0kenTransactor: T0kenTransactor{contract: contract}, T0kenFilterer: T0kenFilterer{contract: contract}}, nil
}

// NewT0kenCaller creates a new read-only instance of T0ken, bound to a specific deployed contract.
func NewT0kenCaller(address common.Address, caller bind.ContractCaller) (*T0kenCaller, error) {
	contract, err := bindT0ken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &T0kenCaller{contract: contract}, nil
}

// NewT0kenTransactor creates a new write-only instance of T0ken, bound to a specific deployed contract.
func NewT0kenTransactor(address common.Address, transactor bind.ContractTransactor) (*T0kenTransactor, error) {
	contract, err := bindT0ken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &T0kenTransactor{contract: contract}, nil
}

// NewT0kenFilterer creates a new log filterer instance of T0ken, bound to a specific deployed contract.
func NewT0kenFilterer(address common.Address, filterer bind.ContractFilterer) (*T0kenFilterer, error) {
	contract, err := bindT0ken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &T0kenFilterer{contract: contract}, nil
}

// bindT0ken binds a generic wrapper to an already deployed contract.
func bindT0ken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(T0kenABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_T0ken *T0kenRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _T0ken.Contract.T0kenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_T0ken *T0kenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _T0ken.Contract.T0kenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_T0ken *T0kenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _T0ken.Contract.T0kenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_T0ken *T0kenCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _T0ken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_T0ken *T0kenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _T0ken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_T0ken *T0kenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _T0ken.Contract.contract.Transact(opts, method, params...)
}

// ZEROADDRESS is a free data retrieval call binding the contract method 0x538ba4f9.
//
// Solidity: function ZERO_ADDRESS() constant returns(address)
func (_T0ken *T0kenCaller) ZEROADDRESS(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _T0ken.contract.Call(opts, out, "ZERO_ADDRESS")
	return *ret0, err
}

// ZEROADDRESS is a free data retrieval call binding the contract method 0x538ba4f9.
//
// Solidity: function ZERO_ADDRESS() constant returns(address)
func (_T0ken *T0kenSession) ZEROADDRESS() (common.Address, error) {
	return _T0ken.Contract.ZEROADDRESS(&_T0ken.CallOpts)
}

// ZEROADDRESS is a free data retrieval call binding the contract method 0x538ba4f9.
//
// Solidity: function ZERO_ADDRESS() constant returns(address)
func (_T0ken *T0kenCallerSession) ZEROADDRESS() (common.Address, error) {
	return _T0ken.Contract.ZEROADDRESS(&_T0ken.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(addrOwner address, spender address) constant returns(uint256)
func (_T0ken *T0kenCaller) Allowance(opts *bind.CallOpts, addrOwner common.Address, spender common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _T0ken.contract.Call(opts, out, "allowance", addrOwner, spender)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(addrOwner address, spender address) constant returns(uint256)
func (_T0ken *T0kenSession) Allowance(addrOwner common.Address, spender common.Address) (*big.Int, error) {
	return _T0ken.Contract.Allowance(&_T0ken.CallOpts, addrOwner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(addrOwner address, spender address) constant returns(uint256)
func (_T0ken *T0kenCallerSession) Allowance(addrOwner common.Address, spender common.Address) (*big.Int, error) {
	return _T0ken.Contract.Allowance(&_T0ken.CallOpts, addrOwner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(addr address) constant returns(uint256)
func (_T0ken *T0kenCaller) BalanceOf(opts *bind.CallOpts, addr common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _T0ken.contract.Call(opts, out, "balanceOf", addr)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(addr address) constant returns(uint256)
func (_T0ken *T0kenSession) BalanceOf(addr common.Address) (*big.Int, error) {
	return _T0ken.Contract.BalanceOf(&_T0ken.CallOpts, addr)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(addr address) constant returns(uint256)
func (_T0ken *T0kenCallerSession) BalanceOf(addr common.Address) (*big.Int, error) {
	return _T0ken.Contract.BalanceOf(&_T0ken.CallOpts, addr)
}

// Balances is a free data retrieval call binding the contract method 0x27e235e3.
//
// Solidity: function balances( address) constant returns(uint256)
func (_T0ken *T0kenCaller) Balances(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _T0ken.contract.Call(opts, out, "balances", arg0)
	return *ret0, err
}

// Balances is a free data retrieval call binding the contract method 0x27e235e3.
//
// Solidity: function balances( address) constant returns(uint256)
func (_T0ken *T0kenSession) Balances(arg0 common.Address) (*big.Int, error) {
	return _T0ken.Contract.Balances(&_T0ken.CallOpts, arg0)
}

// Balances is a free data retrieval call binding the contract method 0x27e235e3.
//
// Solidity: function balances( address) constant returns(uint256)
func (_T0ken *T0kenCallerSession) Balances(arg0 common.Address) (*big.Int, error) {
	return _T0ken.Contract.Balances(&_T0ken.CallOpts, arg0)
}

// Cancellations is a free data retrieval call binding the contract method 0x34a84827.
//
// Solidity: function cancellations( address) constant returns(address)
func (_T0ken *T0kenCaller) Cancellations(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _T0ken.contract.Call(opts, out, "cancellations", arg0)
	return *ret0, err
}

// Cancellations is a free data retrieval call binding the contract method 0x34a84827.
//
// Solidity: function cancellations( address) constant returns(address)
func (_T0ken *T0kenSession) Cancellations(arg0 common.Address) (common.Address, error) {
	return _T0ken.Contract.Cancellations(&_T0ken.CallOpts, arg0)
}

// Cancellations is a free data retrieval call binding the contract method 0x34a84827.
//
// Solidity: function cancellations( address) constant returns(address)
func (_T0ken *T0kenCallerSession) Cancellations(arg0 common.Address) (common.Address, error) {
	return _T0ken.Contract.Cancellations(&_T0ken.CallOpts, arg0)
}

// Compliance is a free data retrieval call binding the contract method 0x6290865d.
//
// Solidity: function compliance() constant returns(address)
func (_T0ken *T0kenCaller) Compliance(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _T0ken.contract.Call(opts, out, "compliance")
	return *ret0, err
}

// Compliance is a free data retrieval call binding the contract method 0x6290865d.
//
// Solidity: function compliance() constant returns(address)
func (_T0ken *T0kenSession) Compliance() (common.Address, error) {
	return _T0ken.Contract.Compliance(&_T0ken.CallOpts)
}

// Compliance is a free data retrieval call binding the contract method 0x6290865d.
//
// Solidity: function compliance() constant returns(address)
func (_T0ken *T0kenCallerSession) Compliance() (common.Address, error) {
	return _T0ken.Contract.Compliance(&_T0ken.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_T0ken *T0kenCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _T0ken.contract.Call(opts, out, "decimals")
	return *ret0, err
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_T0ken *T0kenSession) Decimals() (uint8, error) {
	return _T0ken.Contract.Decimals(&_T0ken.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_T0ken *T0kenCallerSession) Decimals() (uint8, error) {
	return _T0ken.Contract.Decimals(&_T0ken.CallOpts)
}

// GetSuperseded is a free data retrieval call binding the contract method 0xe37ccac7.
//
// Solidity: function getSuperseded(addr address) constant returns(address)
func (_T0ken *T0kenCaller) GetSuperseded(opts *bind.CallOpts, addr common.Address) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _T0ken.contract.Call(opts, out, "getSuperseded", addr)
	return *ret0, err
}

// GetSuperseded is a free data retrieval call binding the contract method 0xe37ccac7.
//
// Solidity: function getSuperseded(addr address) constant returns(address)
func (_T0ken *T0kenSession) GetSuperseded(addr common.Address) (common.Address, error) {
	return _T0ken.Contract.GetSuperseded(&_T0ken.CallOpts, addr)
}

// GetSuperseded is a free data retrieval call binding the contract method 0xe37ccac7.
//
// Solidity: function getSuperseded(addr address) constant returns(address)
func (_T0ken *T0kenCallerSession) GetSuperseded(addr common.Address) (common.Address, error) {
	return _T0ken.Contract.GetSuperseded(&_T0ken.CallOpts, addr)
}

// HolderAt is a free data retrieval call binding the contract method 0x8082a929.
//
// Solidity: function holderAt(index int256) constant returns(address)
func (_T0ken *T0kenCaller) HolderAt(opts *bind.CallOpts, index *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _T0ken.contract.Call(opts, out, "holderAt", index)
	return *ret0, err
}

// HolderAt is a free data retrieval call binding the contract method 0x8082a929.
//
// Solidity: function holderAt(index int256) constant returns(address)
func (_T0ken *T0kenSession) HolderAt(index *big.Int) (common.Address, error) {
	return _T0ken.Contract.HolderAt(&_T0ken.CallOpts, index)
}

// HolderAt is a free data retrieval call binding the contract method 0x8082a929.
//
// Solidity: function holderAt(index int256) constant returns(address)
func (_T0ken *T0kenCallerSession) HolderAt(index *big.Int) (common.Address, error) {
	return _T0ken.Contract.HolderAt(&_T0ken.CallOpts, index)
}

// Holders is a free data retrieval call binding the contract method 0x8188f71c.
//
// Solidity: function holders() constant returns(count int256)
func (_T0ken *T0kenCaller) Holders(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _T0ken.contract.Call(opts, out, "holders")
	return *ret0, err
}

// Holders is a free data retrieval call binding the contract method 0x8188f71c.
//
// Solidity: function holders() constant returns(count int256)
func (_T0ken *T0kenSession) Holders() (*big.Int, error) {
	return _T0ken.Contract.Holders(&_T0ken.CallOpts)
}

// Holders is a free data retrieval call binding the contract method 0x8188f71c.
//
// Solidity: function holders() constant returns(count int256)
func (_T0ken *T0kenCallerSession) Holders() (*big.Int, error) {
	return _T0ken.Contract.Holders(&_T0ken.CallOpts)
}

// IsHolder is a free data retrieval call binding the contract method 0xd4d7b19a.
//
// Solidity: function isHolder(addr address) constant returns(bool)
func (_T0ken *T0kenCaller) IsHolder(opts *bind.CallOpts, addr common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _T0ken.contract.Call(opts, out, "isHolder", addr)
	return *ret0, err
}

// IsHolder is a free data retrieval call binding the contract method 0xd4d7b19a.
//
// Solidity: function isHolder(addr address) constant returns(bool)
func (_T0ken *T0kenSession) IsHolder(addr common.Address) (bool, error) {
	return _T0ken.Contract.IsHolder(&_T0ken.CallOpts, addr)
}

// IsHolder is a free data retrieval call binding the contract method 0xd4d7b19a.
//
// Solidity: function isHolder(addr address) constant returns(bool)
func (_T0ken *T0kenCallerSession) IsHolder(addr common.Address) (bool, error) {
	return _T0ken.Contract.IsHolder(&_T0ken.CallOpts, addr)
}

// IsLocked is a free data retrieval call binding the contract method 0xa4e2d634.
//
// Solidity: function isLocked() constant returns(bool)
func (_T0ken *T0kenCaller) IsLocked(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _T0ken.contract.Call(opts, out, "isLocked")
	return *ret0, err
}

// IsLocked is a free data retrieval call binding the contract method 0xa4e2d634.
//
// Solidity: function isLocked() constant returns(bool)
func (_T0ken *T0kenSession) IsLocked() (bool, error) {
	return _T0ken.Contract.IsLocked(&_T0ken.CallOpts)
}

// IsLocked is a free data retrieval call binding the contract method 0xa4e2d634.
//
// Solidity: function isLocked() constant returns(bool)
func (_T0ken *T0kenCallerSession) IsLocked() (bool, error) {
	return _T0ken.Contract.IsLocked(&_T0ken.CallOpts)
}

// IsSuperseded is a free data retrieval call binding the contract method 0x2da7293e.
//
// Solidity: function isSuperseded(addr address) constant returns(bool)
func (_T0ken *T0kenCaller) IsSuperseded(opts *bind.CallOpts, addr common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _T0ken.contract.Call(opts, out, "isSuperseded", addr)
	return *ret0, err
}

// IsSuperseded is a free data retrieval call binding the contract method 0x2da7293e.
//
// Solidity: function isSuperseded(addr address) constant returns(bool)
func (_T0ken *T0kenSession) IsSuperseded(addr common.Address) (bool, error) {
	return _T0ken.Contract.IsSuperseded(&_T0ken.CallOpts, addr)
}

// IsSuperseded is a free data retrieval call binding the contract method 0x2da7293e.
//
// Solidity: function isSuperseded(addr address) constant returns(bool)
func (_T0ken *T0kenCallerSession) IsSuperseded(addr common.Address) (bool, error) {
	return _T0ken.Contract.IsSuperseded(&_T0ken.CallOpts, addr)
}

// IssuanceFinished is a free data retrieval call binding the contract method 0x4662299a.
//
// Solidity: function issuanceFinished() constant returns(bool)
func (_T0ken *T0kenCaller) IssuanceFinished(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _T0ken.contract.Call(opts, out, "issuanceFinished")
	return *ret0, err
}

// IssuanceFinished is a free data retrieval call binding the contract method 0x4662299a.
//
// Solidity: function issuanceFinished() constant returns(bool)
func (_T0ken *T0kenSession) IssuanceFinished() (bool, error) {
	return _T0ken.Contract.IssuanceFinished(&_T0ken.CallOpts)
}

// IssuanceFinished is a free data retrieval call binding the contract method 0x4662299a.
//
// Solidity: function issuanceFinished() constant returns(bool)
func (_T0ken *T0kenCallerSession) IssuanceFinished() (bool, error) {
	return _T0ken.Contract.IssuanceFinished(&_T0ken.CallOpts)
}

// Issuer is a free data retrieval call binding the contract method 0x1d143848.
//
// Solidity: function issuer() constant returns(address)
func (_T0ken *T0kenCaller) Issuer(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _T0ken.contract.Call(opts, out, "issuer")
	return *ret0, err
}

// Issuer is a free data retrieval call binding the contract method 0x1d143848.
//
// Solidity: function issuer() constant returns(address)
func (_T0ken *T0kenSession) Issuer() (common.Address, error) {
	return _T0ken.Contract.Issuer(&_T0ken.CallOpts)
}

// Issuer is a free data retrieval call binding the contract method 0x1d143848.
//
// Solidity: function issuer() constant returns(address)
func (_T0ken *T0kenCallerSession) Issuer() (common.Address, error) {
	return _T0ken.Contract.Issuer(&_T0ken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_T0ken *T0kenCaller) Name(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _T0ken.contract.Call(opts, out, "name")
	return *ret0, err
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_T0ken *T0kenSession) Name() (string, error) {
	return _T0ken.Contract.Name(&_T0ken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_T0ken *T0kenCallerSession) Name() (string, error) {
	return _T0ken.Contract.Name(&_T0ken.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_T0ken *T0kenCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _T0ken.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_T0ken *T0kenSession) Owner() (common.Address, error) {
	return _T0ken.Contract.Owner(&_T0ken.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_T0ken *T0kenCallerSession) Owner() (common.Address, error) {
	return _T0ken.Contract.Owner(&_T0ken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_T0ken *T0kenCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _T0ken.contract.Call(opts, out, "symbol")
	return *ret0, err
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_T0ken *T0kenSession) Symbol() (string, error) {
	return _T0ken.Contract.Symbol(&_T0ken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_T0ken *T0kenCallerSession) Symbol() (string, error) {
	return _T0ken.Contract.Symbol(&_T0ken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_T0ken *T0kenCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _T0ken.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_T0ken *T0kenSession) TotalSupply() (*big.Int, error) {
	return _T0ken.Contract.TotalSupply(&_T0ken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_T0ken *T0kenCallerSession) TotalSupply() (*big.Int, error) {
	return _T0ken.Contract.TotalSupply(&_T0ken.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(spender address, tokens uint256) returns(bool)
func (_T0ken *T0kenTransactor) Approve(opts *bind.TransactOpts, spender common.Address, tokens *big.Int) (*types.Transaction, error) {
	return _T0ken.contract.Transact(opts, "approve", spender, tokens)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(spender address, tokens uint256) returns(bool)
func (_T0ken *T0kenSession) Approve(spender common.Address, tokens *big.Int) (*types.Transaction, error) {
	return _T0ken.Contract.Approve(&_T0ken.TransactOpts, spender, tokens)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(spender address, tokens uint256) returns(bool)
func (_T0ken *T0kenTransactorSession) Approve(spender common.Address, tokens *big.Int) (*types.Transaction, error) {
	return _T0ken.Contract.Approve(&_T0ken.TransactOpts, spender, tokens)
}

// CancelAndReissue is a paid mutator transaction binding the contract method 0x79f64720.
//
// Solidity: function cancelAndReissue(original address, replacement address) returns()
func (_T0ken *T0kenTransactor) CancelAndReissue(opts *bind.TransactOpts, original common.Address, replacement common.Address) (*types.Transaction, error) {
	return _T0ken.contract.Transact(opts, "cancelAndReissue", original, replacement)
}

// CancelAndReissue is a paid mutator transaction binding the contract method 0x79f64720.
//
// Solidity: function cancelAndReissue(original address, replacement address) returns()
func (_T0ken *T0kenSession) CancelAndReissue(original common.Address, replacement common.Address) (*types.Transaction, error) {
	return _T0ken.Contract.CancelAndReissue(&_T0ken.TransactOpts, original, replacement)
}

// CancelAndReissue is a paid mutator transaction binding the contract method 0x79f64720.
//
// Solidity: function cancelAndReissue(original address, replacement address) returns()
func (_T0ken *T0kenTransactorSession) CancelAndReissue(original common.Address, replacement common.Address) (*types.Transaction, error) {
	return _T0ken.Contract.CancelAndReissue(&_T0ken.TransactOpts, original, replacement)
}

// FinishIssuance is a paid mutator transaction binding the contract method 0xc422293b.
//
// Solidity: function finishIssuance() returns(bool)
func (_T0ken *T0kenTransactor) FinishIssuance(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _T0ken.contract.Transact(opts, "finishIssuance")
}

// FinishIssuance is a paid mutator transaction binding the contract method 0xc422293b.
//
// Solidity: function finishIssuance() returns(bool)
func (_T0ken *T0kenSession) FinishIssuance() (*types.Transaction, error) {
	return _T0ken.Contract.FinishIssuance(&_T0ken.TransactOpts)
}

// FinishIssuance is a paid mutator transaction binding the contract method 0xc422293b.
//
// Solidity: function finishIssuance() returns(bool)
func (_T0ken *T0kenTransactorSession) FinishIssuance() (*types.Transaction, error) {
	return _T0ken.Contract.FinishIssuance(&_T0ken.TransactOpts)
}

// IssueTokens is a paid mutator transaction binding the contract method 0xa5820daa.
//
// Solidity: function issueTokens(quantity uint256) returns(bool)
func (_T0ken *T0kenTransactor) IssueTokens(opts *bind.TransactOpts, quantity *big.Int) (*types.Transaction, error) {
	return _T0ken.contract.Transact(opts, "issueTokens", quantity)
}

// IssueTokens is a paid mutator transaction binding the contract method 0xa5820daa.
//
// Solidity: function issueTokens(quantity uint256) returns(bool)
func (_T0ken *T0kenSession) IssueTokens(quantity *big.Int) (*types.Transaction, error) {
	return _T0ken.Contract.IssueTokens(&_T0ken.TransactOpts, quantity)
}

// IssueTokens is a paid mutator transaction binding the contract method 0xa5820daa.
//
// Solidity: function issueTokens(quantity uint256) returns(bool)
func (_T0ken *T0kenTransactorSession) IssueTokens(quantity *big.Int) (*types.Transaction, error) {
	return _T0ken.Contract.IssueTokens(&_T0ken.TransactOpts, quantity)
}

// Kill is a paid mutator transaction binding the contract method 0x41c0e1b5.
//
// Solidity: function kill() returns()
func (_T0ken *T0kenTransactor) Kill(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _T0ken.contract.Transact(opts, "kill")
}

// Kill is a paid mutator transaction binding the contract method 0x41c0e1b5.
//
// Solidity: function kill() returns()
func (_T0ken *T0kenSession) Kill() (*types.Transaction, error) {
	return _T0ken.Contract.Kill(&_T0ken.TransactOpts)
}

// Kill is a paid mutator transaction binding the contract method 0x41c0e1b5.
//
// Solidity: function kill() returns()
func (_T0ken *T0kenTransactorSession) Kill() (*types.Transaction, error) {
	return _T0ken.Contract.Kill(&_T0ken.TransactOpts)
}

// SetCompliance is a paid mutator transaction binding the contract method 0xf8981789.
//
// Solidity: function setCompliance(newComplianceAddress address) returns()
func (_T0ken *T0kenTransactor) SetCompliance(opts *bind.TransactOpts, newComplianceAddress common.Address) (*types.Transaction, error) {
	return _T0ken.contract.Transact(opts, "setCompliance", newComplianceAddress)
}

// SetCompliance is a paid mutator transaction binding the contract method 0xf8981789.
//
// Solidity: function setCompliance(newComplianceAddress address) returns()
func (_T0ken *T0kenSession) SetCompliance(newComplianceAddress common.Address) (*types.Transaction, error) {
	return _T0ken.Contract.SetCompliance(&_T0ken.TransactOpts, newComplianceAddress)
}

// SetCompliance is a paid mutator transaction binding the contract method 0xf8981789.
//
// Solidity: function setCompliance(newComplianceAddress address) returns()
func (_T0ken *T0kenTransactorSession) SetCompliance(newComplianceAddress common.Address) (*types.Transaction, error) {
	return _T0ken.Contract.SetCompliance(&_T0ken.TransactOpts, newComplianceAddress)
}

// SetIssuer is a paid mutator transaction binding the contract method 0x55cc4e57.
//
// Solidity: function setIssuer(newIssuer address) returns()
func (_T0ken *T0kenTransactor) SetIssuer(opts *bind.TransactOpts, newIssuer common.Address) (*types.Transaction, error) {
	return _T0ken.contract.Transact(opts, "setIssuer", newIssuer)
}

// SetIssuer is a paid mutator transaction binding the contract method 0x55cc4e57.
//
// Solidity: function setIssuer(newIssuer address) returns()
func (_T0ken *T0kenSession) SetIssuer(newIssuer common.Address) (*types.Transaction, error) {
	return _T0ken.Contract.SetIssuer(&_T0ken.TransactOpts, newIssuer)
}

// SetIssuer is a paid mutator transaction binding the contract method 0x55cc4e57.
//
// Solidity: function setIssuer(newIssuer address) returns()
func (_T0ken *T0kenTransactorSession) SetIssuer(newIssuer common.Address) (*types.Transaction, error) {
	return _T0ken.Contract.SetIssuer(&_T0ken.TransactOpts, newIssuer)
}

// SetLocked is a paid mutator transaction binding the contract method 0x211e28b6.
//
// Solidity: function setLocked(locked bool) returns()
func (_T0ken *T0kenTransactor) SetLocked(opts *bind.TransactOpts, locked bool) (*types.Transaction, error) {
	return _T0ken.contract.Transact(opts, "setLocked", locked)
}

// SetLocked is a paid mutator transaction binding the contract method 0x211e28b6.
//
// Solidity: function setLocked(locked bool) returns()
func (_T0ken *T0kenSession) SetLocked(locked bool) (*types.Transaction, error) {
	return _T0ken.Contract.SetLocked(&_T0ken.TransactOpts, locked)
}

// SetLocked is a paid mutator transaction binding the contract method 0x211e28b6.
//
// Solidity: function setLocked(locked bool) returns()
func (_T0ken *T0kenTransactorSession) SetLocked(locked bool) (*types.Transaction, error) {
	return _T0ken.Contract.SetLocked(&_T0ken.TransactOpts, locked)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(to address, tokens uint256) returns(bool)
func (_T0ken *T0kenTransactor) Transfer(opts *bind.TransactOpts, to common.Address, tokens *big.Int) (*types.Transaction, error) {
	return _T0ken.contract.Transact(opts, "transfer", to, tokens)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(to address, tokens uint256) returns(bool)
func (_T0ken *T0kenSession) Transfer(to common.Address, tokens *big.Int) (*types.Transaction, error) {
	return _T0ken.Contract.Transfer(&_T0ken.TransactOpts, to, tokens)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(to address, tokens uint256) returns(bool)
func (_T0ken *T0kenTransactorSession) Transfer(to common.Address, tokens *big.Int) (*types.Transaction, error) {
	return _T0ken.Contract.Transfer(&_T0ken.TransactOpts, to, tokens)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(from address, to address, tokens uint256) returns(bool)
func (_T0ken *T0kenTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokens *big.Int) (*types.Transaction, error) {
	return _T0ken.contract.Transact(opts, "transferFrom", from, to, tokens)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(from address, to address, tokens uint256) returns(bool)
func (_T0ken *T0kenSession) TransferFrom(from common.Address, to common.Address, tokens *big.Int) (*types.Transaction, error) {
	return _T0ken.Contract.TransferFrom(&_T0ken.TransactOpts, from, to, tokens)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(from address, to address, tokens uint256) returns(bool)
func (_T0ken *T0kenTransactorSession) TransferFrom(from common.Address, to common.Address, tokens *big.Int) (*types.Transaction, error) {
	return _T0ken.Contract.TransferFrom(&_T0ken.TransactOpts, from, to, tokens)
}

// TransferOverride is a paid mutator transaction binding the contract method 0x80318be8.
//
// Solidity: function transferOverride(from address, to address, tokens uint256) returns(bool)
func (_T0ken *T0kenTransactor) TransferOverride(opts *bind.TransactOpts, from common.Address, to common.Address, tokens *big.Int) (*types.Transaction, error) {
	return _T0ken.contract.Transact(opts, "transferOverride", from, to, tokens)
}

// TransferOverride is a paid mutator transaction binding the contract method 0x80318be8.
//
// Solidity: function transferOverride(from address, to address, tokens uint256) returns(bool)
func (_T0ken *T0kenSession) TransferOverride(from common.Address, to common.Address, tokens *big.Int) (*types.Transaction, error) {
	return _T0ken.Contract.TransferOverride(&_T0ken.TransactOpts, from, to, tokens)
}

// TransferOverride is a paid mutator transaction binding the contract method 0x80318be8.
//
// Solidity: function transferOverride(from address, to address, tokens uint256) returns(bool)
func (_T0ken *T0kenTransactorSession) TransferOverride(from common.Address, to common.Address, tokens *big.Int) (*types.Transaction, error) {
	return _T0ken.Contract.TransferOverride(&_T0ken.TransactOpts, from, to, tokens)
}

// TransferOwner is a paid mutator transaction binding the contract method 0x4fb2e45d.
//
// Solidity: function transferOwner(newOwner address) returns()
func (_T0ken *T0kenTransactor) TransferOwner(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _T0ken.contract.Transact(opts, "transferOwner", newOwner)
}

// TransferOwner is a paid mutator transaction binding the contract method 0x4fb2e45d.
//
// Solidity: function transferOwner(newOwner address) returns()
func (_T0ken *T0kenSession) TransferOwner(newOwner common.Address) (*types.Transaction, error) {
	return _T0ken.Contract.TransferOwner(&_T0ken.TransactOpts, newOwner)
}

// TransferOwner is a paid mutator transaction binding the contract method 0x4fb2e45d.
//
// Solidity: function transferOwner(newOwner address) returns()
func (_T0ken *T0kenTransactorSession) TransferOwner(newOwner common.Address) (*types.Transaction, error) {
	return _T0ken.Contract.TransferOwner(&_T0ken.TransactOpts, newOwner)
}

// T0kenApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the T0ken contract.
type T0kenApprovalIterator struct {
	Event *T0kenApproval // Event containing the contract specifics and raw log

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
func (it *T0kenApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(T0kenApproval)
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
		it.Event = new(T0kenApproval)
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
func (it *T0kenApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *T0kenApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// T0kenApproval represents a Approval event raised by the T0ken contract.
type T0kenApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: e Approval(owner indexed address, spender indexed address, value uint256)
func (_T0ken *T0kenFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*T0kenApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _T0ken.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &T0kenApprovalIterator{contract: _T0ken.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: e Approval(owner indexed address, spender indexed address, value uint256)
func (_T0ken *T0kenFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *T0kenApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _T0ken.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(T0kenApproval)
				if err := _T0ken.contract.UnpackLog(event, "Approval", log); err != nil {
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

// T0kenIssuanceIterator is returned from FilterIssuance and is used to iterate over the raw logs and unpacked data for Issuance events raised by the T0ken contract.
type T0kenIssuanceIterator struct {
	Event *T0kenIssuance // Event containing the contract specifics and raw log

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
func (it *T0kenIssuanceIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(T0kenIssuance)
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
		it.Event = new(T0kenIssuance)
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
func (it *T0kenIssuanceIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *T0kenIssuanceIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// T0kenIssuance represents a Issuance event raised by the T0ken contract.
type T0kenIssuance struct {
	To     common.Address
	Tokens *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterIssuance is a free log retrieval operation binding the contract event 0x9cb9c14f7bc76e3a89b796b091850526236115352a198b1e472f00e91376bbcb.
//
// Solidity: e Issuance(to indexed address, tokens uint256)
func (_T0ken *T0kenFilterer) FilterIssuance(opts *bind.FilterOpts, to []common.Address) (*T0kenIssuanceIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _T0ken.contract.FilterLogs(opts, "Issuance", toRule)
	if err != nil {
		return nil, err
	}
	return &T0kenIssuanceIterator{contract: _T0ken.contract, event: "Issuance", logs: logs, sub: sub}, nil
}

// WatchIssuance is a free log subscription operation binding the contract event 0x9cb9c14f7bc76e3a89b796b091850526236115352a198b1e472f00e91376bbcb.
//
// Solidity: e Issuance(to indexed address, tokens uint256)
func (_T0ken *T0kenFilterer) WatchIssuance(opts *bind.WatchOpts, sink chan<- *T0kenIssuance, to []common.Address) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _T0ken.contract.WatchLogs(opts, "Issuance", toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(T0kenIssuance)
				if err := _T0ken.contract.UnpackLog(event, "Issuance", log); err != nil {
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

// T0kenIssuanceFinishedIterator is returned from FilterIssuanceFinished and is used to iterate over the raw logs and unpacked data for IssuanceFinished events raised by the T0ken contract.
type T0kenIssuanceFinishedIterator struct {
	Event *T0kenIssuanceFinished // Event containing the contract specifics and raw log

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
func (it *T0kenIssuanceFinishedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(T0kenIssuanceFinished)
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
		it.Event = new(T0kenIssuanceFinished)
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
func (it *T0kenIssuanceFinishedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *T0kenIssuanceFinishedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// T0kenIssuanceFinished represents a IssuanceFinished event raised by the T0ken contract.
type T0kenIssuanceFinished struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterIssuanceFinished is a free log retrieval operation binding the contract event 0x29fe76cc5ca143e91eadf7242fda487fcef09318c1237900f958abe1e2c5beff.
//
// Solidity: e IssuanceFinished()
func (_T0ken *T0kenFilterer) FilterIssuanceFinished(opts *bind.FilterOpts) (*T0kenIssuanceFinishedIterator, error) {

	logs, sub, err := _T0ken.contract.FilterLogs(opts, "IssuanceFinished")
	if err != nil {
		return nil, err
	}
	return &T0kenIssuanceFinishedIterator{contract: _T0ken.contract, event: "IssuanceFinished", logs: logs, sub: sub}, nil
}

// WatchIssuanceFinished is a free log subscription operation binding the contract event 0x29fe76cc5ca143e91eadf7242fda487fcef09318c1237900f958abe1e2c5beff.
//
// Solidity: e IssuanceFinished()
func (_T0ken *T0kenFilterer) WatchIssuanceFinished(opts *bind.WatchOpts, sink chan<- *T0kenIssuanceFinished) (event.Subscription, error) {

	logs, sub, err := _T0ken.contract.WatchLogs(opts, "IssuanceFinished")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(T0kenIssuanceFinished)
				if err := _T0ken.contract.UnpackLog(event, "IssuanceFinished", log); err != nil {
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

// T0kenIssuerSetIterator is returned from FilterIssuerSet and is used to iterate over the raw logs and unpacked data for IssuerSet events raised by the T0ken contract.
type T0kenIssuerSetIterator struct {
	Event *T0kenIssuerSet // Event containing the contract specifics and raw log

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
func (it *T0kenIssuerSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(T0kenIssuerSet)
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
		it.Event = new(T0kenIssuerSet)
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
func (it *T0kenIssuerSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *T0kenIssuerSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// T0kenIssuerSet represents a IssuerSet event raised by the T0ken contract.
type T0kenIssuerSet struct {
	PreviousIssuer common.Address
	NewIssuer      common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterIssuerSet is a free log retrieval operation binding the contract event 0xf7189b85d7899f5a32d733e6584c4f1dcdff0274f09d969d186c1797673ede11.
//
// Solidity: e IssuerSet(previousIssuer indexed address, newIssuer indexed address)
func (_T0ken *T0kenFilterer) FilterIssuerSet(opts *bind.FilterOpts, previousIssuer []common.Address, newIssuer []common.Address) (*T0kenIssuerSetIterator, error) {

	var previousIssuerRule []interface{}
	for _, previousIssuerItem := range previousIssuer {
		previousIssuerRule = append(previousIssuerRule, previousIssuerItem)
	}
	var newIssuerRule []interface{}
	for _, newIssuerItem := range newIssuer {
		newIssuerRule = append(newIssuerRule, newIssuerItem)
	}

	logs, sub, err := _T0ken.contract.FilterLogs(opts, "IssuerSet", previousIssuerRule, newIssuerRule)
	if err != nil {
		return nil, err
	}
	return &T0kenIssuerSetIterator{contract: _T0ken.contract, event: "IssuerSet", logs: logs, sub: sub}, nil
}

// WatchIssuerSet is a free log subscription operation binding the contract event 0xf7189b85d7899f5a32d733e6584c4f1dcdff0274f09d969d186c1797673ede11.
//
// Solidity: e IssuerSet(previousIssuer indexed address, newIssuer indexed address)
func (_T0ken *T0kenFilterer) WatchIssuerSet(opts *bind.WatchOpts, sink chan<- *T0kenIssuerSet, previousIssuer []common.Address, newIssuer []common.Address) (event.Subscription, error) {

	var previousIssuerRule []interface{}
	for _, previousIssuerItem := range previousIssuer {
		previousIssuerRule = append(previousIssuerRule, previousIssuerItem)
	}
	var newIssuerRule []interface{}
	for _, newIssuerItem := range newIssuer {
		newIssuerRule = append(newIssuerRule, newIssuerItem)
	}

	logs, sub, err := _T0ken.contract.WatchLogs(opts, "IssuerSet", previousIssuerRule, newIssuerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(T0kenIssuerSet)
				if err := _T0ken.contract.UnpackLog(event, "IssuerSet", log); err != nil {
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

// T0kenOwnerTransferredIterator is returned from FilterOwnerTransferred and is used to iterate over the raw logs and unpacked data for OwnerTransferred events raised by the T0ken contract.
type T0kenOwnerTransferredIterator struct {
	Event *T0kenOwnerTransferred // Event containing the contract specifics and raw log

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
func (it *T0kenOwnerTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(T0kenOwnerTransferred)
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
		it.Event = new(T0kenOwnerTransferred)
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
func (it *T0kenOwnerTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *T0kenOwnerTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// T0kenOwnerTransferred represents a OwnerTransferred event raised by the T0ken contract.
type T0kenOwnerTransferred struct {
	OldOwner common.Address
	NewOwner common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOwnerTransferred is a free log retrieval operation binding the contract event 0x8934ce4adea8d9ce0d714d2c22b86790e41b7731c84b926fbbdc1d40ff6533c9.
//
// Solidity: e OwnerTransferred(oldOwner indexed address, newOwner indexed address)
func (_T0ken *T0kenFilterer) FilterOwnerTransferred(opts *bind.FilterOpts, oldOwner []common.Address, newOwner []common.Address) (*T0kenOwnerTransferredIterator, error) {

	var oldOwnerRule []interface{}
	for _, oldOwnerItem := range oldOwner {
		oldOwnerRule = append(oldOwnerRule, oldOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _T0ken.contract.FilterLogs(opts, "OwnerTransferred", oldOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &T0kenOwnerTransferredIterator{contract: _T0ken.contract, event: "OwnerTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnerTransferred is a free log subscription operation binding the contract event 0x8934ce4adea8d9ce0d714d2c22b86790e41b7731c84b926fbbdc1d40ff6533c9.
//
// Solidity: e OwnerTransferred(oldOwner indexed address, newOwner indexed address)
func (_T0ken *T0kenFilterer) WatchOwnerTransferred(opts *bind.WatchOpts, sink chan<- *T0kenOwnerTransferred, oldOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var oldOwnerRule []interface{}
	for _, oldOwnerItem := range oldOwner {
		oldOwnerRule = append(oldOwnerRule, oldOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _T0ken.contract.WatchLogs(opts, "OwnerTransferred", oldOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(T0kenOwnerTransferred)
				if err := _T0ken.contract.UnpackLog(event, "OwnerTransferred", log); err != nil {
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

// T0kenShareholderAddedIterator is returned from FilterShareholderAdded and is used to iterate over the raw logs and unpacked data for ShareholderAdded events raised by the T0ken contract.
type T0kenShareholderAddedIterator struct {
	Event *T0kenShareholderAdded // Event containing the contract specifics and raw log

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
func (it *T0kenShareholderAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(T0kenShareholderAdded)
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
		it.Event = new(T0kenShareholderAdded)
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
func (it *T0kenShareholderAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *T0kenShareholderAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// T0kenShareholderAdded represents a ShareholderAdded event raised by the T0ken contract.
type T0kenShareholderAdded struct {
	Shareholder common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterShareholderAdded is a free log retrieval operation binding the contract event 0x3082c1c4977de80c4f67ee77b56b282610ec93a7ecbcc31b551e0ac2f7bd0395.
//
// Solidity: e ShareholderAdded(shareholder address)
func (_T0ken *T0kenFilterer) FilterShareholderAdded(opts *bind.FilterOpts) (*T0kenShareholderAddedIterator, error) {

	logs, sub, err := _T0ken.contract.FilterLogs(opts, "ShareholderAdded")
	if err != nil {
		return nil, err
	}
	return &T0kenShareholderAddedIterator{contract: _T0ken.contract, event: "ShareholderAdded", logs: logs, sub: sub}, nil
}

// WatchShareholderAdded is a free log subscription operation binding the contract event 0x3082c1c4977de80c4f67ee77b56b282610ec93a7ecbcc31b551e0ac2f7bd0395.
//
// Solidity: e ShareholderAdded(shareholder address)
func (_T0ken *T0kenFilterer) WatchShareholderAdded(opts *bind.WatchOpts, sink chan<- *T0kenShareholderAdded) (event.Subscription, error) {

	logs, sub, err := _T0ken.contract.WatchLogs(opts, "ShareholderAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(T0kenShareholderAdded)
				if err := _T0ken.contract.UnpackLog(event, "ShareholderAdded", log); err != nil {
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

// T0kenShareholderRemovedIterator is returned from FilterShareholderRemoved and is used to iterate over the raw logs and unpacked data for ShareholderRemoved events raised by the T0ken contract.
type T0kenShareholderRemovedIterator struct {
	Event *T0kenShareholderRemoved // Event containing the contract specifics and raw log

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
func (it *T0kenShareholderRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(T0kenShareholderRemoved)
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
		it.Event = new(T0kenShareholderRemoved)
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
func (it *T0kenShareholderRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *T0kenShareholderRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// T0kenShareholderRemoved represents a ShareholderRemoved event raised by the T0ken contract.
type T0kenShareholderRemoved struct {
	Shareholder common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterShareholderRemoved is a free log retrieval operation binding the contract event 0x7ba114ff3d9844510a088eea94cd35562e7c97a2d36a418b37b2e61e5c77affe.
//
// Solidity: e ShareholderRemoved(shareholder address)
func (_T0ken *T0kenFilterer) FilterShareholderRemoved(opts *bind.FilterOpts) (*T0kenShareholderRemovedIterator, error) {

	logs, sub, err := _T0ken.contract.FilterLogs(opts, "ShareholderRemoved")
	if err != nil {
		return nil, err
	}
	return &T0kenShareholderRemovedIterator{contract: _T0ken.contract, event: "ShareholderRemoved", logs: logs, sub: sub}, nil
}

// WatchShareholderRemoved is a free log subscription operation binding the contract event 0x7ba114ff3d9844510a088eea94cd35562e7c97a2d36a418b37b2e61e5c77affe.
//
// Solidity: e ShareholderRemoved(shareholder address)
func (_T0ken *T0kenFilterer) WatchShareholderRemoved(opts *bind.WatchOpts, sink chan<- *T0kenShareholderRemoved) (event.Subscription, error) {

	logs, sub, err := _T0ken.contract.WatchLogs(opts, "ShareholderRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(T0kenShareholderRemoved)
				if err := _T0ken.contract.UnpackLog(event, "ShareholderRemoved", log); err != nil {
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

// T0kenTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the T0ken contract.
type T0kenTransferIterator struct {
	Event *T0kenTransfer // Event containing the contract specifics and raw log

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
func (it *T0kenTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(T0kenTransfer)
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
		it.Event = new(T0kenTransfer)
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
func (it *T0kenTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *T0kenTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// T0kenTransfer represents a Transfer event raised by the T0ken contract.
type T0kenTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: e Transfer(from indexed address, to indexed address, value uint256)
func (_T0ken *T0kenFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*T0kenTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _T0ken.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &T0kenTransferIterator{contract: _T0ken.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: e Transfer(from indexed address, to indexed address, value uint256)
func (_T0ken *T0kenFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *T0kenTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _T0ken.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(T0kenTransfer)
				if err := _T0ken.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// T0kenVerifiedAddressSupersededIterator is returned from FilterVerifiedAddressSuperseded and is used to iterate over the raw logs and unpacked data for VerifiedAddressSuperseded events raised by the T0ken contract.
type T0kenVerifiedAddressSupersededIterator struct {
	Event *T0kenVerifiedAddressSuperseded // Event containing the contract specifics and raw log

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
func (it *T0kenVerifiedAddressSupersededIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(T0kenVerifiedAddressSuperseded)
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
		it.Event = new(T0kenVerifiedAddressSuperseded)
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
func (it *T0kenVerifiedAddressSupersededIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *T0kenVerifiedAddressSupersededIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// T0kenVerifiedAddressSuperseded represents a VerifiedAddressSuperseded event raised by the T0ken contract.
type T0kenVerifiedAddressSuperseded struct {
	Original    common.Address
	Replacement common.Address
	Sender      common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterVerifiedAddressSuperseded is a free log retrieval operation binding the contract event 0xb64971100522354f3d25283cb14e2eefcb0dd26a757482ccfe42479d0a686857.
//
// Solidity: e VerifiedAddressSuperseded(original indexed address, replacement indexed address, sender indexed address)
func (_T0ken *T0kenFilterer) FilterVerifiedAddressSuperseded(opts *bind.FilterOpts, original []common.Address, replacement []common.Address, sender []common.Address) (*T0kenVerifiedAddressSupersededIterator, error) {

	var originalRule []interface{}
	for _, originalItem := range original {
		originalRule = append(originalRule, originalItem)
	}
	var replacementRule []interface{}
	for _, replacementItem := range replacement {
		replacementRule = append(replacementRule, replacementItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _T0ken.contract.FilterLogs(opts, "VerifiedAddressSuperseded", originalRule, replacementRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &T0kenVerifiedAddressSupersededIterator{contract: _T0ken.contract, event: "VerifiedAddressSuperseded", logs: logs, sub: sub}, nil
}

// WatchVerifiedAddressSuperseded is a free log subscription operation binding the contract event 0xb64971100522354f3d25283cb14e2eefcb0dd26a757482ccfe42479d0a686857.
//
// Solidity: e VerifiedAddressSuperseded(original indexed address, replacement indexed address, sender indexed address)
func (_T0ken *T0kenFilterer) WatchVerifiedAddressSuperseded(opts *bind.WatchOpts, sink chan<- *T0kenVerifiedAddressSuperseded, original []common.Address, replacement []common.Address, sender []common.Address) (event.Subscription, error) {

	var originalRule []interface{}
	for _, originalItem := range original {
		originalRule = append(originalRule, originalItem)
	}
	var replacementRule []interface{}
	for _, replacementItem := range replacement {
		replacementRule = append(replacementRule, replacementItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _T0ken.contract.WatchLogs(opts, "VerifiedAddressSuperseded", originalRule, replacementRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(T0kenVerifiedAddressSuperseded)
				if err := _T0ken.contract.UnpackLog(event, "VerifiedAddressSuperseded", log); err != nil {
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
