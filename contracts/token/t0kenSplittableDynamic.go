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

// T0kenSplittableDynamicABI is the input ABI used to generate the binding from.
const T0kenSplittableDynamicABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"spender\",\"type\":\"address\"},{\"name\":\"tokens\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"currentSplitIndex\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"issuer\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"locked\",\"type\":\"bool\"}],\"name\":\"setLocked\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"from\",\"type\":\"address\"},{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"tokens\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"balances\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"kill\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"issuanceFinished\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwner\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ZERO_ADDRESS\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newIssuer\",\"type\":\"address\"}],\"name\":\"setIssuer\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"compliance\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"from\",\"type\":\"address\"},{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"tokens\",\"type\":\"uint256\"}],\"name\":\"transferOverride\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"index\",\"type\":\"int256\"}],\"name\":\"holderAt\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"holders\",\"outputs\":[{\"name\":\"count\",\"type\":\"int256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"numerator\",\"type\":\"uint128\"},{\"name\":\"denominator\",\"type\":\"uint128\"}],\"name\":\"updateSplit\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"splits\",\"outputs\":[{\"name\":\"numerator\",\"type\":\"uint128\"},{\"name\":\"denominator\",\"type\":\"uint128\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"cashedOutTokenQuantity\",\"type\":\"uint256\"}],\"name\":\"splitTotalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"removeShareholder\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isLocked\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"quantity\",\"type\":\"uint256\"}],\"name\":\"issueTokens\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"tokens\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"holdersSplit\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"finishIssuance\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"isHolder\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"addrOwner\",\"type\":\"address\"},{\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newComplianceAddress\",\"type\":\"address\"}],\"name\":\"setCompliance\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"tokenName\",\"type\":\"string\"},{\"name\":\"tokenSymbol\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"oldOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnerTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousIssuer\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newIssuer\",\"type\":\"address\"}],\"name\":\"IssuerSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"tokens\",\"type\":\"uint256\"}],\"name\":\"Issuance\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"IssuanceFinished\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"shareholder\",\"type\":\"address\"}],\"name\":\"ShareholderAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"shareholder\",\"type\":\"address\"}],\"name\":\"ShareholderRemoved\",\"type\":\"event\"}]"

// T0kenSplittableDynamicBin is the compiled bytecode used for deploying new contracts.
const T0kenSplittableDynamicBin = `6080604052600180546001600160a01b03191690556000600f553480156200002657600080fd5b506040516200259438038062002594833981810160405260408110156200004c57600080fd5b8101908080516401000000008111156200006557600080fd5b820160208101848111156200007957600080fd5b81516401000000008111828201871017156200009457600080fd5b50509291906020018051640100000000811115620000b157600080fd5b82016020810184811115620000c557600080fd5b8151640100000000811182820187101715620000e057600080fd5b5050600080546001600160a01b031916331781556001805460ff60a01b1916905585519194508593508492509062000120906002906020860190620001d4565b50815162000136906003906020850190620001d4565b506004805460ff90921660ff199092169190911790555050604080518082019091526001808252602080830191825260008052600e905290517fe710864318d4a32f37d6ce54cb3fadbef648dd12d8dbdf53973564d56b7f881c805492516001600160801b03908116700100000000000000000000000000000000029281166001600160801b03199094169390931790921617905550620002799050565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106200021757805160ff191683800117855562000247565b8280016001018555821562000247579182015b82811115620002475782518255916020019190600101906200022a565b506200025592915062000259565b5090565b6200027691905b8082111562000255576000815560010162000260565b90565b61230b80620002896000396000f3fe608060405234801561001057600080fd5b506004361061020b5760003560e01c806380318be81161012a5780639babdad6116100bd578063b0b89d921161008c578063d4d7b19a11610071578063d4d7b19a146105de578063dd62ed3e14610604578063f8981789146106325761020b565b8063b0b89d92146105b0578063c422293b146105d65761020b565b80639babdad614610539578063a4e2d6341461055f578063a5820daa14610567578063a9059cbb146105845761020b565b8063884c3006116100f9578063884c3006146104c95780638da5cb5b1461050c57806395d89b411461051457806398d34b9b1461051c5761020b565b806380318be8146104405780638082a929146104765780638188f71c14610493578063872a72ea1461049b5761020b565b8063313ce567116101a2578063538ba4f911610171578063538ba4f9146103e457806355cc4e57146103ec5780636290865d1461041257806370a082311461041a5761020b565b8063313ce5671461039057806341c0e1b5146103ae5780634662299a146103b65780634fb2e45d146103be5761020b565b80631d143848116101de5780631d143848146102ef578063211e28b61461031357806323b872dd1461033457806327e235e31461036a5761020b565b806306fdde0314610210578063095ea7b31461028d57806317344a90146102cd57806318160ddd146102e7575b600080fd5b610218610658565b6040805160208082528351818301528351919283929083019185019080838360005b8381101561025257818101518382015260200161023a565b50505050905090810190601f16801561027f5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b6102b9600480360360408110156102a357600080fd5b506001600160a01b0381351690602001356106e3565b604080519115158252519081900360200190f35b6102d56107f9565b60408051918252519081900360200190f35b6102d56107ff565b6102f7610805565b604080516001600160a01b039092168252519081900360200190f35b6103326004803603602081101561032957600080fd5b50351515610814565b005b6102b96004803603606081101561034a57600080fd5b506001600160a01b038135811691602081013590911690604001356108fc565b6102d56004803603602081101561038057600080fd5b50356001600160a01b0316610925565b610398610937565b6040805160ff9092168252519081900360200190f35b610332610940565b6102b96109c8565b610332600480360360208110156103d457600080fd5b50356001600160a01b03166109d8565b6102f7610b4a565b6103326004803603602081101561040257600080fd5b50356001600160a01b0316610b59565b6102f7610ce0565b6102d56004803603602081101561043057600080fd5b50356001600160a01b0316610cef565b6102b96004803603606081101561045657600080fd5b506001600160a01b03813581169160208101359091169060400135610d74565b6102f76004803603602081101561048c57600080fd5b5035610d8b565b6102d5610d9e565b610332600480360360408110156104b157600080fd5b506001600160801b0381358116916020013516610da4565b6104e6600480360360208110156104df57600080fd5b5035610f43565b604080516001600160801b03938416815291909216602082015281519081900390910190f35b6102f7610f69565b610218610f78565b6102b96004803603602081101561053257600080fd5b5035610fd3565b6103326004803603602081101561054f57600080fd5b50356001600160a01b03166110e8565b6102b96111ef565b6102b96004803603602081101561057d57600080fd5b50356111ff565b6102b96004803603604081101561059a57600080fd5b506001600160a01b038135169060200135611229565b6102d5600480360360208110156105c657600080fd5b50356001600160a01b0316611246565b6102b9611258565b6102b9600480360360208110156105f457600080fd5b50356001600160a01b03166113a5565b6102d56004803603604081101561061a57600080fd5b506001600160a01b03813581169160200135166113b8565b6103326004803603602081101561064857600080fd5b50356001600160a01b0316611441565b6002805460408051602060018416156101000260001901909316849004601f810184900484028201840190925281815292918301828280156106db5780601f106106b0576101008083540402835291602001916106db565b820191906000526020600020905b8154815290600101906020018083116106be57829003601f168201915b505050505081565b600154600090600160a01b900460ff161561072f5760405162461bcd60e51b815260040180806020018281038252602d8152602001806122aa602d913960400191505060405180910390fd5b61074060053363ffffffff61152616565b610791576040805162461bcd60e51b815260206004820152601560248201527f4d7573742062652061207368617265686f6c6465720000000000000000000000604482015290519081900360640190fd5b336000818152600c602090815260408083206001600160a01b03881680855290835292819020869055805186815290519293927f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925929181900390910190a35060015b92915050565b600f5481565b600b5490565b6009546001600160a01b031681565b6000546001600160a01b031633148061083d57506001546000546001600160a01b039081169116145b61088e576040805162461bcd60e51b815260206004820152601960248201527f4f776e6572206163636f756e7420697320726571756972656400000000000000604482015290519081900360640190fd5b60015460ff600160a01b90910416151581151514156108de5760405162461bcd60e51b81526004018080602001828103825260288152602001806122826028913960400191505060405180910390fd5b60018054911515600160a01b0260ff60a01b19909216919091179055565b6000610908848461155b565b6109128484611598565b61091d8484846115f8565b949350505050565b600a6020526000908152604090205481565b60045460ff1681565b6000546001600160a01b031633148061096957506001546000546001600160a01b039081169116145b6109ba576040805162461bcd60e51b815260206004820152601960248201527f4f776e6572206163636f756e7420697320726571756972656400000000000000604482015290519081900360640190fd5b6000546001600160a01b0316ff5b600954600160a01b900460ff1681565b6000546001600160a01b0316331480610a0157506001546000546001600160a01b039081169116145b610a52576040805162461bcd60e51b815260206004820152601960248201527f4f776e6572206163636f756e7420697320726571756972656400000000000000604482015290519081900360640190fd5b6000546001600160a01b0382811691161415610a9f5760405162461bcd60e51b815260040180806020018281038252602581526020018061225d6025913960400191505060405180910390fd5b6001600160a01b038116610afa576040805162461bcd60e51b815260206004820181905260248201527f4e6577204f776e65722063616e6e6f74206265207a65726f2061646472657373604482015290519081900360640190fd5b600080546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8934ce4adea8d9ce0d714d2c22b86790e41b7731c84b926fbbdc1d40ff6533c99190a35050565b6001546001600160a01b031681565b600154600160a01b900460ff1615610ba25760405162461bcd60e51b815260040180806020018281038252602d8152602001806122aa602d913960400191505060405180910390fd5b6000546001600160a01b0316331480610bcb57506001546000546001600160a01b039081169116145b610c1c576040805162461bcd60e51b815260206004820152601960248201527f4f776e6572206163636f756e7420697320726571756972656400000000000000604482015290519081900360640190fd5b610c2d60058263ffffffff61152616565b15610c695760405162461bcd60e51b815260040180806020018281038252602181526020018061223c6021913960400191505060405180910390fd5b6009546001600160a01b03166000818152600a6020526040902054610c90919083906117a6565b600980546001600160a01b0319166001600160a01b0383811691821792839055604051919216907ff7189b85d7899f5a32d733e6584c4f1dcdff0274f09d969d186c1797673ede1190600090a350565b6008546001600160a01b031681565b6001600160a01b0381166000908152600a6020908152604080832054600d9092528220546001015b600f548111610d6d576000818152600e6020526040902054610d63906001600160801b03600160801b8204811691610d579186911663ffffffff61194916565b9063ffffffff61197016565b9150600101610d17565b5092915050565b6000610d808484611598565b61091d84848461198f565b60006107f360058363ffffffff611a6316565b60055481565b6009546001600160a01b03163314610df9576040805162461bcd60e51b815260206004820152601360248201527213db9b1e481a5cdcdd595c88185b1b1bddd959606a1b604482015290519081900360640190fd5b600154600160a01b900460ff16610e57576040805162461bcd60e51b815260206004820152600e60248201527f4d757374206265206c6f636b6564000000000000000000000000000000000000604482015290519081900360640190fd5b6000826001600160801b0316118015610e7957506000816001600160801b0316115b610eca576040805162461bcd60e51b815260206004820152601860248201527f4e756d62657273206d75737420626520706f7369746976650000000000000000604482015290519081900360640190fd5b600f8054600101908190556040805180820182526001600160801b03948516815292841660208085019182526000938452600e905291209151825491517fffffffffffffffffffffffffffffffff00000000000000000000000000000000909216908416178316600160801b9190931602919091179055565b600e602052600090815260409020546001600160801b0380821691600160801b90041682565b6000546001600160a01b031681565b6003805460408051602060026001851615610100026000190190941693909304601f810184900484028201840190925281815292918301828280156106db5780601f106106b0576101008083540402835291602001916106db565b6009546000906001600160a01b0316331461102b576040805162461bcd60e51b815260206004820152601360248201527213db9b1e481a5cdcdd595c88185b1b1bddd959606a1b604482015290519081900360640190fd5b600154600160a01b900460ff16611089576040805162461bcd60e51b815260206004820152600e60248201527f4d757374206265206c6f636b6564000000000000000000000000000000000000604482015290519081900360640190fd5b600f546000908152600e6020526040902054600b546001600160801b0380831692600160801b900416906110db908290610d579085906110cf908963ffffffff611ae916565b9063ffffffff61194916565b600b555060019392505050565b6009546001600160a01b0316331461113d576040805162461bcd60e51b815260206004820152601360248201527213db9b1e481a5cdcdd595c88185b1b1bddd959606a1b604482015290519081900360640190fd5b61114681610cef565b15801561115f575061115f60058263ffffffff611b4616565b6111b0576040805162461bcd60e51b815260206004820152601760248201527f5368617265686f6c646572206e6f742072656d6f766564000000000000000000604482015290519081900360640190fd5b604080516001600160a01b038316815290517f7ba114ff3d9844510a088eea94cd35562e7c97a2d36a418b37b2e61e5c77affe9181900360200190a150565b600154600160a01b900460ff1681565b600f546009546001600160a01b03166000908152600d60205260408120919091556107f382611c31565b60006112353384611598565b61123f8383611e3f565b9392505050565b600d6020526000908152604090205481565b600154600090600160a01b900460ff16156112a45760405162461bcd60e51b815260040180806020018281038252602d8152602001806122aa602d913960400191505060405180910390fd5b6009546001600160a01b031633146112f9576040805162461bcd60e51b815260206004820152601360248201527213db9b1e481a5cdcdd595c88185b1b1bddd959606a1b604482015290519081900360640190fd5b600954600160a01b900460ff1615611358576040805162461bcd60e51b815260206004820152601960248201527f49737375616e636520616c72656164792066696e697368656400000000000000604482015290519081900360640190fd5b6009805460ff60a01b1916600160a01b1790556040517f29fe76cc5ca143e91eadf7242fda487fcef09318c1237900f958abe1e2c5beff90600090a150600954600160a01b900460ff1690565b60006107f360058363ffffffff61152616565b6001600160a01b038083166000818152600c60209081526040808320948616835293815283822054928252600d9052918220546001015b600f548111611439576000818152600e602052604090205461142f906001600160801b03600160801b8204811691610d579186911663ffffffff61194916565b91506001016113ef565b509392505050565b600154600160a01b900460ff161561148a5760405162461bcd60e51b815260040180806020018281038252602d8152602001806122aa602d913960400191505060405180910390fd5b6000546001600160a01b03163314806114b357506001546000546001600160a01b039081169116145b611504576040805162461bcd60e51b815260206004820152601960248201527f4f776e6572206163636f756e7420697320726571756972656400000000000000604482015290519081900360640190fd5b600880546001600160a01b0319166001600160a01b0392909216919091179055565b6001600160a01b03811660009081526001830160205260408120546000190181811280159061091d5750925490921292915050565b600061156783836113b8565b6001600160a01b039384166000908152600c6020908152604080832095909616825293909352929091209190915550565b60006115a383610cef565b905060006115b083610cef565b600f546001600160a01b039586166000818152600d602090815260408083208590559790981680825287822093909355908152600a9096528486209390935591845250912055565b600154600090600160a01b900460ff16156116445760405162461bcd60e51b815260040180806020018281038252602d8152602001806122aa602d913960400191505060405180910390fd5b6001600160a01b0384166000908152600a6020526040902054849083908111156116aa576040805162461bcd60e51b8152602060048201526012602482015271496e73756666696369656e742066756e647360701b604482015290519081900360640190fd5b6001600160a01b0386166000908152600c60209081526040808320338452909152902054841115611722576040805162461bcd60e51b815260206004820152601a60248201527f5472616e73666572206578636565647320616c6c6f77616e6365000000000000604482015290519081900360640190fd5b60006117318787876000611fee565b9050801561179c576001600160a01b0387166000908152600c6020908152604080832033845290915290205461176d908663ffffffff611ae916565b6001600160a01b0388166000908152600c6020908152604080832033845290915290205561179c8787876117a6565b9695505050505050565b6001600160a01b0383166000908152600a60205260409020546117cf908263ffffffff611ae916565b6001600160a01b038085166000908152600a60205260408082209390935590841681522054611804908263ffffffff61213c16565b6001600160a01b038084166000818152600a602090815260409182902094909455805185815290519193928716927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef92918290030190a36001600160a01b0382166000908152600a60205260409020541580159061188e575061188e60058363ffffffff61219616565b156118d057604080516001600160a01b038416815290517f3082c1c4977de80c4f67ee77b56b282610ec93a7ecbcc31b551e0ac2f7bd03959181900360200190a15b6001600160a01b0383166000908152600a6020526040902054158015611902575061190260058463ffffffff611b4616565b1561194457604080516001600160a01b038516815290517f7ba114ff3d9844510a088eea94cd35562e7c97a2d36a418b37b2e61e5c77affe9181900360200190a15b505050565b600082611958575060006107f3565b8282028284828161196557fe5b041461123f57600080fd5b600080821161197e57600080fd5b81838161198757fe5b049392505050565b600154600090600160a01b900460ff16156119db5760405162461bcd60e51b815260040180806020018281038252602d8152602001806122aa602d913960400191505060405180910390fd5b6001600160a01b0384166000908152600a602052604090205484908390811115611a41576040805162461bcd60e51b8152602060048201526012602482015271496e73756666696369656e742066756e647360701b604482015290519081900360640190fd5b6000611a508787876001611fee565b9050801561179c5761179c8787876117a6565b6000808212158015611a755750825482125b611ac6576040805162461bcd60e51b815260206004820152601860248201527f496e646578206f757473696465206f6620626f756e64732e0000000000000000604482015290519081900360640190fd5b50600101600090815260029190910160205260409020546001600160a01b031690565b600082821115611b40576040805162461bcd60e51b815260206004820152601460248201527f526573756c747320696e20756e646572666c6f77000000000000000000000000604482015290519081900360640190fd5b50900390565b6001600160a01b03811660009081526001808401602052604082205490811280611b705750835481135b15611b7f5760009150506107f3565b8354811215611be657835460009081526002850160208181526040808420546001600160a01b03168085526001890183528185208690558585529290915280832080546001600160a01b031990811690931790558654835290912080549091169055611c05565b6000818152600285016020526040902080546001600160a01b03191690555b50506001600160a01b031660009081526001828101602052604082209190915581546000190190915590565b600154600090600160a01b900460ff1615611c7d5760405162461bcd60e51b815260040180806020018281038252602d8152602001806122aa602d913960400191505060405180910390fd5b6009546001600160a01b03163314611cd2576040805162461bcd60e51b815260206004820152601360248201527213db9b1e481a5cdcdd595c88185b1b1bddd959606a1b604482015290519081900360640190fd5b600954600160a01b900460ff1615611d31576040805162461bcd60e51b815260206004820152601960248201527f49737375616e636520616c72656164792066696e697368656400000000000000604482015290519081900360640190fd5b8115611daf57600b54611d4a908363ffffffff61213c16565b600b556009546001600160a01b03166000908152600a6020526040902054611d78908363ffffffff61213c16565b600980546001600160a01b039081166000908152600a602052604090209290925554611dad916005911663ffffffff61219616565b505b6009546040805184815290516001600160a01b03909216917f9cb9c14f7bc76e3a89b796b091850526236115352a198b1e472f00e91376bbcb9181900360200190a26009546040805184815290516001600160a01b03909216916000917fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef919081900360200190a3506001919050565b600154600090600160a01b900460ff1615611e8b5760405162461bcd60e51b815260040180806020018281038252602d8152602001806122aa602d913960400191505060405180910390fd5b336000818152600a60205260409020548390811115611ee6576040805162461bcd60e51b8152602060048201526012602482015271496e73756666696369656e742066756e647360701b604482015290519081900360640190fd5b6009546000906001600160a01b0316331415611fc457506008546001600160a01b03161580611fbf57600854600954604080517ffd8258bd0000000000000000000000000000000000000000000000000000000081526001600160a01b039283166004820181905260248201528983166044820152606481018990529051919092169163fd8258bd9160848083019260209291908290030181600087803b158015611f9057600080fd5b505af1158015611fa4573d6000803e3d6000fd5b505050506040513d6020811015611fba57600080fd5b505190505b611fd4565b611fd13387876000611fee565b90505b8015611fe557611fe53387876117a6565b95945050505050565b6008546000906001600160a01b03166120095750801561091d565b81156120bf57600854604080517f5acba2010000000000000000000000000000000000000000000000000000000081523360048201526001600160a01b03888116602483015287811660448301526064820187905291519190921691635acba2019160848083019260209291908290030181600087803b15801561208c57600080fd5b505af11580156120a0573d6000803e3d6000fd5b505050506040513d60208110156120b657600080fd5b5051905061091d565b600854604080517f6d62a4fe0000000000000000000000000000000000000000000000000000000081523360048201526001600160a01b03888116602483015287811660448301526064820187905291519190921691636d62a4fe9160848083019260209291908290030181600087803b15801561208c57600080fd5b60008282018381101561123f576040805162461bcd60e51b815260206004820152601360248201527f526573756c747320696e206f766572666c6f7700000000000000000000000000604482015290519081900360640190fd5b60006001600160a01b0382166121ae575060006107f3565b6001600160a01b0382166000908152600184016020526040812054600019019081128015906121dd5750835481125b156121ec5760009150506107f3565b5050815460019081018084556001600160a01b0383166000818152838601602090815260408083208590559382526002870190529190912080546001600160a01b03191690911790559291505056fe4e6577206973737565722063616e27742062652061207368617265686f6c6465724e6577204f776e65722063616e6e6f74206265207468652063757272656e74206f776e6572436f6e747261637420616c726561647920696e20726571756573746564206c6f636b207374617465436f6e74726163742069732063757272656e746c79206c6f636b656420666f72206d6f64696669636174696f6ea265627a7a72305820541aca03bb0247332b06510bc7ae6931d3e3fbf53e1ec65116c6de64a1a5fd5764736f6c634300050a0032`

// DeployT0kenSplittableDynamic deploys a new Ethereum contract, binding an instance of T0kenSplittableDynamic to it.
func DeployT0kenSplittableDynamic(auth *bind.TransactOpts, backend bind.ContractBackend, tokenName string, tokenSymbol string) (common.Address, *types.Transaction, *T0kenSplittableDynamic, error) {
	parsed, err := abi.JSON(strings.NewReader(T0kenSplittableDynamicABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(T0kenSplittableDynamicBin), backend, tokenName, tokenSymbol)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &T0kenSplittableDynamic{T0kenSplittableDynamicCaller: T0kenSplittableDynamicCaller{contract: contract}, T0kenSplittableDynamicTransactor: T0kenSplittableDynamicTransactor{contract: contract}, T0kenSplittableDynamicFilterer: T0kenSplittableDynamicFilterer{contract: contract}}, nil
}

// T0kenSplittableDynamic is an auto generated Go binding around an Ethereum contract.
type T0kenSplittableDynamic struct {
	T0kenSplittableDynamicCaller     // Read-only binding to the contract
	T0kenSplittableDynamicTransactor // Write-only binding to the contract
	T0kenSplittableDynamicFilterer   // Log filterer for contract events
}

// T0kenSplittableDynamicCaller is an auto generated read-only Go binding around an Ethereum contract.
type T0kenSplittableDynamicCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// T0kenSplittableDynamicTransactor is an auto generated write-only Go binding around an Ethereum contract.
type T0kenSplittableDynamicTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// T0kenSplittableDynamicFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type T0kenSplittableDynamicFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// T0kenSplittableDynamicSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type T0kenSplittableDynamicSession struct {
	Contract     *T0kenSplittableDynamic // Generic contract binding to set the session for
	CallOpts     bind.CallOpts           // Call options to use throughout this session
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// T0kenSplittableDynamicCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type T0kenSplittableDynamicCallerSession struct {
	Contract *T0kenSplittableDynamicCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                 // Call options to use throughout this session
}

// T0kenSplittableDynamicTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type T0kenSplittableDynamicTransactorSession struct {
	Contract     *T0kenSplittableDynamicTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                 // Transaction auth options to use throughout this session
}

// T0kenSplittableDynamicRaw is an auto generated low-level Go binding around an Ethereum contract.
type T0kenSplittableDynamicRaw struct {
	Contract *T0kenSplittableDynamic // Generic contract binding to access the raw methods on
}

// T0kenSplittableDynamicCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type T0kenSplittableDynamicCallerRaw struct {
	Contract *T0kenSplittableDynamicCaller // Generic read-only contract binding to access the raw methods on
}

// T0kenSplittableDynamicTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type T0kenSplittableDynamicTransactorRaw struct {
	Contract *T0kenSplittableDynamicTransactor // Generic write-only contract binding to access the raw methods on
}

// NewT0kenSplittableDynamic creates a new instance of T0kenSplittableDynamic, bound to a specific deployed contract.
func NewT0kenSplittableDynamic(address common.Address, backend bind.ContractBackend) (*T0kenSplittableDynamic, error) {
	contract, err := bindT0kenSplittableDynamic(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &T0kenSplittableDynamic{T0kenSplittableDynamicCaller: T0kenSplittableDynamicCaller{contract: contract}, T0kenSplittableDynamicTransactor: T0kenSplittableDynamicTransactor{contract: contract}, T0kenSplittableDynamicFilterer: T0kenSplittableDynamicFilterer{contract: contract}}, nil
}

// NewT0kenSplittableDynamicCaller creates a new read-only instance of T0kenSplittableDynamic, bound to a specific deployed contract.
func NewT0kenSplittableDynamicCaller(address common.Address, caller bind.ContractCaller) (*T0kenSplittableDynamicCaller, error) {
	contract, err := bindT0kenSplittableDynamic(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &T0kenSplittableDynamicCaller{contract: contract}, nil
}

// NewT0kenSplittableDynamicTransactor creates a new write-only instance of T0kenSplittableDynamic, bound to a specific deployed contract.
func NewT0kenSplittableDynamicTransactor(address common.Address, transactor bind.ContractTransactor) (*T0kenSplittableDynamicTransactor, error) {
	contract, err := bindT0kenSplittableDynamic(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &T0kenSplittableDynamicTransactor{contract: contract}, nil
}

// NewT0kenSplittableDynamicFilterer creates a new log filterer instance of T0kenSplittableDynamic, bound to a specific deployed contract.
func NewT0kenSplittableDynamicFilterer(address common.Address, filterer bind.ContractFilterer) (*T0kenSplittableDynamicFilterer, error) {
	contract, err := bindT0kenSplittableDynamic(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &T0kenSplittableDynamicFilterer{contract: contract}, nil
}

// bindT0kenSplittableDynamic binds a generic wrapper to an already deployed contract.
func bindT0kenSplittableDynamic(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(T0kenSplittableDynamicABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_T0kenSplittableDynamic *T0kenSplittableDynamicRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _T0kenSplittableDynamic.Contract.T0kenSplittableDynamicCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_T0kenSplittableDynamic *T0kenSplittableDynamicRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _T0kenSplittableDynamic.Contract.T0kenSplittableDynamicTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_T0kenSplittableDynamic *T0kenSplittableDynamicRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _T0kenSplittableDynamic.Contract.T0kenSplittableDynamicTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_T0kenSplittableDynamic *T0kenSplittableDynamicCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _T0kenSplittableDynamic.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_T0kenSplittableDynamic *T0kenSplittableDynamicTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _T0kenSplittableDynamic.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_T0kenSplittableDynamic *T0kenSplittableDynamicTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _T0kenSplittableDynamic.Contract.contract.Transact(opts, method, params...)
}

// ZEROADDRESS is a free data retrieval call binding the contract method 0x538ba4f9.
//
// Solidity: function ZERO_ADDRESS() constant returns(address)
func (_T0kenSplittableDynamic *T0kenSplittableDynamicCaller) ZEROADDRESS(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _T0kenSplittableDynamic.contract.Call(opts, out, "ZERO_ADDRESS")
	return *ret0, err
}

// ZEROADDRESS is a free data retrieval call binding the contract method 0x538ba4f9.
//
// Solidity: function ZERO_ADDRESS() constant returns(address)
func (_T0kenSplittableDynamic *T0kenSplittableDynamicSession) ZEROADDRESS() (common.Address, error) {
	return _T0kenSplittableDynamic.Contract.ZEROADDRESS(&_T0kenSplittableDynamic.CallOpts)
}

// ZEROADDRESS is a free data retrieval call binding the contract method 0x538ba4f9.
//
// Solidity: function ZERO_ADDRESS() constant returns(address)
func (_T0kenSplittableDynamic *T0kenSplittableDynamicCallerSession) ZEROADDRESS() (common.Address, error) {
	return _T0kenSplittableDynamic.Contract.ZEROADDRESS(&_T0kenSplittableDynamic.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(addrOwner address, spender address) constant returns(uint256)
func (_T0kenSplittableDynamic *T0kenSplittableDynamicCaller) Allowance(opts *bind.CallOpts, addrOwner common.Address, spender common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _T0kenSplittableDynamic.contract.Call(opts, out, "allowance", addrOwner, spender)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(addrOwner address, spender address) constant returns(uint256)
func (_T0kenSplittableDynamic *T0kenSplittableDynamicSession) Allowance(addrOwner common.Address, spender common.Address) (*big.Int, error) {
	return _T0kenSplittableDynamic.Contract.Allowance(&_T0kenSplittableDynamic.CallOpts, addrOwner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(addrOwner address, spender address) constant returns(uint256)
func (_T0kenSplittableDynamic *T0kenSplittableDynamicCallerSession) Allowance(addrOwner common.Address, spender common.Address) (*big.Int, error) {
	return _T0kenSplittableDynamic.Contract.Allowance(&_T0kenSplittableDynamic.CallOpts, addrOwner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(addr address) constant returns(uint256)
func (_T0kenSplittableDynamic *T0kenSplittableDynamicCaller) BalanceOf(opts *bind.CallOpts, addr common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _T0kenSplittableDynamic.contract.Call(opts, out, "balanceOf", addr)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(addr address) constant returns(uint256)
func (_T0kenSplittableDynamic *T0kenSplittableDynamicSession) BalanceOf(addr common.Address) (*big.Int, error) {
	return _T0kenSplittableDynamic.Contract.BalanceOf(&_T0kenSplittableDynamic.CallOpts, addr)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(addr address) constant returns(uint256)
func (_T0kenSplittableDynamic *T0kenSplittableDynamicCallerSession) BalanceOf(addr common.Address) (*big.Int, error) {
	return _T0kenSplittableDynamic.Contract.BalanceOf(&_T0kenSplittableDynamic.CallOpts, addr)
}

// Balances is a free data retrieval call binding the contract method 0x27e235e3.
//
// Solidity: function balances( address) constant returns(uint256)
func (_T0kenSplittableDynamic *T0kenSplittableDynamicCaller) Balances(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _T0kenSplittableDynamic.contract.Call(opts, out, "balances", arg0)
	return *ret0, err
}

// Balances is a free data retrieval call binding the contract method 0x27e235e3.
//
// Solidity: function balances( address) constant returns(uint256)
func (_T0kenSplittableDynamic *T0kenSplittableDynamicSession) Balances(arg0 common.Address) (*big.Int, error) {
	return _T0kenSplittableDynamic.Contract.Balances(&_T0kenSplittableDynamic.CallOpts, arg0)
}

// Balances is a free data retrieval call binding the contract method 0x27e235e3.
//
// Solidity: function balances( address) constant returns(uint256)
func (_T0kenSplittableDynamic *T0kenSplittableDynamicCallerSession) Balances(arg0 common.Address) (*big.Int, error) {
	return _T0kenSplittableDynamic.Contract.Balances(&_T0kenSplittableDynamic.CallOpts, arg0)
}

// Compliance is a free data retrieval call binding the contract method 0x6290865d.
//
// Solidity: function compliance() constant returns(address)
func (_T0kenSplittableDynamic *T0kenSplittableDynamicCaller) Compliance(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _T0kenSplittableDynamic.contract.Call(opts, out, "compliance")
	return *ret0, err
}

// Compliance is a free data retrieval call binding the contract method 0x6290865d.
//
// Solidity: function compliance() constant returns(address)
func (_T0kenSplittableDynamic *T0kenSplittableDynamicSession) Compliance() (common.Address, error) {
	return _T0kenSplittableDynamic.Contract.Compliance(&_T0kenSplittableDynamic.CallOpts)
}

// Compliance is a free data retrieval call binding the contract method 0x6290865d.
//
// Solidity: function compliance() constant returns(address)
func (_T0kenSplittableDynamic *T0kenSplittableDynamicCallerSession) Compliance() (common.Address, error) {
	return _T0kenSplittableDynamic.Contract.Compliance(&_T0kenSplittableDynamic.CallOpts)
}

// CurrentSplitIndex is a free data retrieval call binding the contract method 0x17344a90.
//
// Solidity: function currentSplitIndex() constant returns(uint256)
func (_T0kenSplittableDynamic *T0kenSplittableDynamicCaller) CurrentSplitIndex(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _T0kenSplittableDynamic.contract.Call(opts, out, "currentSplitIndex")
	return *ret0, err
}

// CurrentSplitIndex is a free data retrieval call binding the contract method 0x17344a90.
//
// Solidity: function currentSplitIndex() constant returns(uint256)
func (_T0kenSplittableDynamic *T0kenSplittableDynamicSession) CurrentSplitIndex() (*big.Int, error) {
	return _T0kenSplittableDynamic.Contract.CurrentSplitIndex(&_T0kenSplittableDynamic.CallOpts)
}

// CurrentSplitIndex is a free data retrieval call binding the contract method 0x17344a90.
//
// Solidity: function currentSplitIndex() constant returns(uint256)
func (_T0kenSplittableDynamic *T0kenSplittableDynamicCallerSession) CurrentSplitIndex() (*big.Int, error) {
	return _T0kenSplittableDynamic.Contract.CurrentSplitIndex(&_T0kenSplittableDynamic.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_T0kenSplittableDynamic *T0kenSplittableDynamicCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _T0kenSplittableDynamic.contract.Call(opts, out, "decimals")
	return *ret0, err
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_T0kenSplittableDynamic *T0kenSplittableDynamicSession) Decimals() (uint8, error) {
	return _T0kenSplittableDynamic.Contract.Decimals(&_T0kenSplittableDynamic.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_T0kenSplittableDynamic *T0kenSplittableDynamicCallerSession) Decimals() (uint8, error) {
	return _T0kenSplittableDynamic.Contract.Decimals(&_T0kenSplittableDynamic.CallOpts)
}

// HolderAt is a free data retrieval call binding the contract method 0x8082a929.
//
// Solidity: function holderAt(index int256) constant returns(address)
func (_T0kenSplittableDynamic *T0kenSplittableDynamicCaller) HolderAt(opts *bind.CallOpts, index *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _T0kenSplittableDynamic.contract.Call(opts, out, "holderAt", index)
	return *ret0, err
}

// HolderAt is a free data retrieval call binding the contract method 0x8082a929.
//
// Solidity: function holderAt(index int256) constant returns(address)
func (_T0kenSplittableDynamic *T0kenSplittableDynamicSession) HolderAt(index *big.Int) (common.Address, error) {
	return _T0kenSplittableDynamic.Contract.HolderAt(&_T0kenSplittableDynamic.CallOpts, index)
}

// HolderAt is a free data retrieval call binding the contract method 0x8082a929.
//
// Solidity: function holderAt(index int256) constant returns(address)
func (_T0kenSplittableDynamic *T0kenSplittableDynamicCallerSession) HolderAt(index *big.Int) (common.Address, error) {
	return _T0kenSplittableDynamic.Contract.HolderAt(&_T0kenSplittableDynamic.CallOpts, index)
}

// Holders is a free data retrieval call binding the contract method 0x8188f71c.
//
// Solidity: function holders() constant returns(count int256)
func (_T0kenSplittableDynamic *T0kenSplittableDynamicCaller) Holders(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _T0kenSplittableDynamic.contract.Call(opts, out, "holders")
	return *ret0, err
}

// Holders is a free data retrieval call binding the contract method 0x8188f71c.
//
// Solidity: function holders() constant returns(count int256)
func (_T0kenSplittableDynamic *T0kenSplittableDynamicSession) Holders() (*big.Int, error) {
	return _T0kenSplittableDynamic.Contract.Holders(&_T0kenSplittableDynamic.CallOpts)
}

// Holders is a free data retrieval call binding the contract method 0x8188f71c.
//
// Solidity: function holders() constant returns(count int256)
func (_T0kenSplittableDynamic *T0kenSplittableDynamicCallerSession) Holders() (*big.Int, error) {
	return _T0kenSplittableDynamic.Contract.Holders(&_T0kenSplittableDynamic.CallOpts)
}

// HoldersSplit is a free data retrieval call binding the contract method 0xb0b89d92.
//
// Solidity: function holdersSplit( address) constant returns(uint256)
func (_T0kenSplittableDynamic *T0kenSplittableDynamicCaller) HoldersSplit(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _T0kenSplittableDynamic.contract.Call(opts, out, "holdersSplit", arg0)
	return *ret0, err
}

// HoldersSplit is a free data retrieval call binding the contract method 0xb0b89d92.
//
// Solidity: function holdersSplit( address) constant returns(uint256)
func (_T0kenSplittableDynamic *T0kenSplittableDynamicSession) HoldersSplit(arg0 common.Address) (*big.Int, error) {
	return _T0kenSplittableDynamic.Contract.HoldersSplit(&_T0kenSplittableDynamic.CallOpts, arg0)
}

// HoldersSplit is a free data retrieval call binding the contract method 0xb0b89d92.
//
// Solidity: function holdersSplit( address) constant returns(uint256)
func (_T0kenSplittableDynamic *T0kenSplittableDynamicCallerSession) HoldersSplit(arg0 common.Address) (*big.Int, error) {
	return _T0kenSplittableDynamic.Contract.HoldersSplit(&_T0kenSplittableDynamic.CallOpts, arg0)
}

// IsHolder is a free data retrieval call binding the contract method 0xd4d7b19a.
//
// Solidity: function isHolder(addr address) constant returns(bool)
func (_T0kenSplittableDynamic *T0kenSplittableDynamicCaller) IsHolder(opts *bind.CallOpts, addr common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _T0kenSplittableDynamic.contract.Call(opts, out, "isHolder", addr)
	return *ret0, err
}

// IsHolder is a free data retrieval call binding the contract method 0xd4d7b19a.
//
// Solidity: function isHolder(addr address) constant returns(bool)
func (_T0kenSplittableDynamic *T0kenSplittableDynamicSession) IsHolder(addr common.Address) (bool, error) {
	return _T0kenSplittableDynamic.Contract.IsHolder(&_T0kenSplittableDynamic.CallOpts, addr)
}

// IsHolder is a free data retrieval call binding the contract method 0xd4d7b19a.
//
// Solidity: function isHolder(addr address) constant returns(bool)
func (_T0kenSplittableDynamic *T0kenSplittableDynamicCallerSession) IsHolder(addr common.Address) (bool, error) {
	return _T0kenSplittableDynamic.Contract.IsHolder(&_T0kenSplittableDynamic.CallOpts, addr)
}

// IsLocked is a free data retrieval call binding the contract method 0xa4e2d634.
//
// Solidity: function isLocked() constant returns(bool)
func (_T0kenSplittableDynamic *T0kenSplittableDynamicCaller) IsLocked(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _T0kenSplittableDynamic.contract.Call(opts, out, "isLocked")
	return *ret0, err
}

// IsLocked is a free data retrieval call binding the contract method 0xa4e2d634.
//
// Solidity: function isLocked() constant returns(bool)
func (_T0kenSplittableDynamic *T0kenSplittableDynamicSession) IsLocked() (bool, error) {
	return _T0kenSplittableDynamic.Contract.IsLocked(&_T0kenSplittableDynamic.CallOpts)
}

// IsLocked is a free data retrieval call binding the contract method 0xa4e2d634.
//
// Solidity: function isLocked() constant returns(bool)
func (_T0kenSplittableDynamic *T0kenSplittableDynamicCallerSession) IsLocked() (bool, error) {
	return _T0kenSplittableDynamic.Contract.IsLocked(&_T0kenSplittableDynamic.CallOpts)
}

// IssuanceFinished is a free data retrieval call binding the contract method 0x4662299a.
//
// Solidity: function issuanceFinished() constant returns(bool)
func (_T0kenSplittableDynamic *T0kenSplittableDynamicCaller) IssuanceFinished(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _T0kenSplittableDynamic.contract.Call(opts, out, "issuanceFinished")
	return *ret0, err
}

// IssuanceFinished is a free data retrieval call binding the contract method 0x4662299a.
//
// Solidity: function issuanceFinished() constant returns(bool)
func (_T0kenSplittableDynamic *T0kenSplittableDynamicSession) IssuanceFinished() (bool, error) {
	return _T0kenSplittableDynamic.Contract.IssuanceFinished(&_T0kenSplittableDynamic.CallOpts)
}

// IssuanceFinished is a free data retrieval call binding the contract method 0x4662299a.
//
// Solidity: function issuanceFinished() constant returns(bool)
func (_T0kenSplittableDynamic *T0kenSplittableDynamicCallerSession) IssuanceFinished() (bool, error) {
	return _T0kenSplittableDynamic.Contract.IssuanceFinished(&_T0kenSplittableDynamic.CallOpts)
}

// Issuer is a free data retrieval call binding the contract method 0x1d143848.
//
// Solidity: function issuer() constant returns(address)
func (_T0kenSplittableDynamic *T0kenSplittableDynamicCaller) Issuer(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _T0kenSplittableDynamic.contract.Call(opts, out, "issuer")
	return *ret0, err
}

// Issuer is a free data retrieval call binding the contract method 0x1d143848.
//
// Solidity: function issuer() constant returns(address)
func (_T0kenSplittableDynamic *T0kenSplittableDynamicSession) Issuer() (common.Address, error) {
	return _T0kenSplittableDynamic.Contract.Issuer(&_T0kenSplittableDynamic.CallOpts)
}

// Issuer is a free data retrieval call binding the contract method 0x1d143848.
//
// Solidity: function issuer() constant returns(address)
func (_T0kenSplittableDynamic *T0kenSplittableDynamicCallerSession) Issuer() (common.Address, error) {
	return _T0kenSplittableDynamic.Contract.Issuer(&_T0kenSplittableDynamic.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_T0kenSplittableDynamic *T0kenSplittableDynamicCaller) Name(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _T0kenSplittableDynamic.contract.Call(opts, out, "name")
	return *ret0, err
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_T0kenSplittableDynamic *T0kenSplittableDynamicSession) Name() (string, error) {
	return _T0kenSplittableDynamic.Contract.Name(&_T0kenSplittableDynamic.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_T0kenSplittableDynamic *T0kenSplittableDynamicCallerSession) Name() (string, error) {
	return _T0kenSplittableDynamic.Contract.Name(&_T0kenSplittableDynamic.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_T0kenSplittableDynamic *T0kenSplittableDynamicCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _T0kenSplittableDynamic.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_T0kenSplittableDynamic *T0kenSplittableDynamicSession) Owner() (common.Address, error) {
	return _T0kenSplittableDynamic.Contract.Owner(&_T0kenSplittableDynamic.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_T0kenSplittableDynamic *T0kenSplittableDynamicCallerSession) Owner() (common.Address, error) {
	return _T0kenSplittableDynamic.Contract.Owner(&_T0kenSplittableDynamic.CallOpts)
}

// Splits is a free data retrieval call binding the contract method 0x884c3006.
//
// Solidity: function splits( uint256) constant returns(numerator uint128, denominator uint128)
func (_T0kenSplittableDynamic *T0kenSplittableDynamicCaller) Splits(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Numerator   *big.Int
	Denominator *big.Int
}, error) {
	ret := new(struct {
		Numerator   *big.Int
		Denominator *big.Int
	})
	out := ret
	err := _T0kenSplittableDynamic.contract.Call(opts, out, "splits", arg0)
	return *ret, err
}

// Splits is a free data retrieval call binding the contract method 0x884c3006.
//
// Solidity: function splits( uint256) constant returns(numerator uint128, denominator uint128)
func (_T0kenSplittableDynamic *T0kenSplittableDynamicSession) Splits(arg0 *big.Int) (struct {
	Numerator   *big.Int
	Denominator *big.Int
}, error) {
	return _T0kenSplittableDynamic.Contract.Splits(&_T0kenSplittableDynamic.CallOpts, arg0)
}

// Splits is a free data retrieval call binding the contract method 0x884c3006.
//
// Solidity: function splits( uint256) constant returns(numerator uint128, denominator uint128)
func (_T0kenSplittableDynamic *T0kenSplittableDynamicCallerSession) Splits(arg0 *big.Int) (struct {
	Numerator   *big.Int
	Denominator *big.Int
}, error) {
	return _T0kenSplittableDynamic.Contract.Splits(&_T0kenSplittableDynamic.CallOpts, arg0)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_T0kenSplittableDynamic *T0kenSplittableDynamicCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _T0kenSplittableDynamic.contract.Call(opts, out, "symbol")
	return *ret0, err
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_T0kenSplittableDynamic *T0kenSplittableDynamicSession) Symbol() (string, error) {
	return _T0kenSplittableDynamic.Contract.Symbol(&_T0kenSplittableDynamic.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_T0kenSplittableDynamic *T0kenSplittableDynamicCallerSession) Symbol() (string, error) {
	return _T0kenSplittableDynamic.Contract.Symbol(&_T0kenSplittableDynamic.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_T0kenSplittableDynamic *T0kenSplittableDynamicCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _T0kenSplittableDynamic.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_T0kenSplittableDynamic *T0kenSplittableDynamicSession) TotalSupply() (*big.Int, error) {
	return _T0kenSplittableDynamic.Contract.TotalSupply(&_T0kenSplittableDynamic.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_T0kenSplittableDynamic *T0kenSplittableDynamicCallerSession) TotalSupply() (*big.Int, error) {
	return _T0kenSplittableDynamic.Contract.TotalSupply(&_T0kenSplittableDynamic.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(spender address, tokens uint256) returns(bool)
func (_T0kenSplittableDynamic *T0kenSplittableDynamicTransactor) Approve(opts *bind.TransactOpts, spender common.Address, tokens *big.Int) (*types.Transaction, error) {
	return _T0kenSplittableDynamic.contract.Transact(opts, "approve", spender, tokens)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(spender address, tokens uint256) returns(bool)
func (_T0kenSplittableDynamic *T0kenSplittableDynamicSession) Approve(spender common.Address, tokens *big.Int) (*types.Transaction, error) {
	return _T0kenSplittableDynamic.Contract.Approve(&_T0kenSplittableDynamic.TransactOpts, spender, tokens)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(spender address, tokens uint256) returns(bool)
func (_T0kenSplittableDynamic *T0kenSplittableDynamicTransactorSession) Approve(spender common.Address, tokens *big.Int) (*types.Transaction, error) {
	return _T0kenSplittableDynamic.Contract.Approve(&_T0kenSplittableDynamic.TransactOpts, spender, tokens)
}

// FinishIssuance is a paid mutator transaction binding the contract method 0xc422293b.
//
// Solidity: function finishIssuance() returns(bool)
func (_T0kenSplittableDynamic *T0kenSplittableDynamicTransactor) FinishIssuance(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _T0kenSplittableDynamic.contract.Transact(opts, "finishIssuance")
}

// FinishIssuance is a paid mutator transaction binding the contract method 0xc422293b.
//
// Solidity: function finishIssuance() returns(bool)
func (_T0kenSplittableDynamic *T0kenSplittableDynamicSession) FinishIssuance() (*types.Transaction, error) {
	return _T0kenSplittableDynamic.Contract.FinishIssuance(&_T0kenSplittableDynamic.TransactOpts)
}

// FinishIssuance is a paid mutator transaction binding the contract method 0xc422293b.
//
// Solidity: function finishIssuance() returns(bool)
func (_T0kenSplittableDynamic *T0kenSplittableDynamicTransactorSession) FinishIssuance() (*types.Transaction, error) {
	return _T0kenSplittableDynamic.Contract.FinishIssuance(&_T0kenSplittableDynamic.TransactOpts)
}

// IssueTokens is a paid mutator transaction binding the contract method 0xa5820daa.
//
// Solidity: function issueTokens(quantity uint256) returns(bool)
func (_T0kenSplittableDynamic *T0kenSplittableDynamicTransactor) IssueTokens(opts *bind.TransactOpts, quantity *big.Int) (*types.Transaction, error) {
	return _T0kenSplittableDynamic.contract.Transact(opts, "issueTokens", quantity)
}

// IssueTokens is a paid mutator transaction binding the contract method 0xa5820daa.
//
// Solidity: function issueTokens(quantity uint256) returns(bool)
func (_T0kenSplittableDynamic *T0kenSplittableDynamicSession) IssueTokens(quantity *big.Int) (*types.Transaction, error) {
	return _T0kenSplittableDynamic.Contract.IssueTokens(&_T0kenSplittableDynamic.TransactOpts, quantity)
}

// IssueTokens is a paid mutator transaction binding the contract method 0xa5820daa.
//
// Solidity: function issueTokens(quantity uint256) returns(bool)
func (_T0kenSplittableDynamic *T0kenSplittableDynamicTransactorSession) IssueTokens(quantity *big.Int) (*types.Transaction, error) {
	return _T0kenSplittableDynamic.Contract.IssueTokens(&_T0kenSplittableDynamic.TransactOpts, quantity)
}

// Kill is a paid mutator transaction binding the contract method 0x41c0e1b5.
//
// Solidity: function kill() returns()
func (_T0kenSplittableDynamic *T0kenSplittableDynamicTransactor) Kill(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _T0kenSplittableDynamic.contract.Transact(opts, "kill")
}

// Kill is a paid mutator transaction binding the contract method 0x41c0e1b5.
//
// Solidity: function kill() returns()
func (_T0kenSplittableDynamic *T0kenSplittableDynamicSession) Kill() (*types.Transaction, error) {
	return _T0kenSplittableDynamic.Contract.Kill(&_T0kenSplittableDynamic.TransactOpts)
}

// Kill is a paid mutator transaction binding the contract method 0x41c0e1b5.
//
// Solidity: function kill() returns()
func (_T0kenSplittableDynamic *T0kenSplittableDynamicTransactorSession) Kill() (*types.Transaction, error) {
	return _T0kenSplittableDynamic.Contract.Kill(&_T0kenSplittableDynamic.TransactOpts)
}

// RemoveShareholder is a paid mutator transaction binding the contract method 0x9babdad6.
//
// Solidity: function removeShareholder(addr address) returns()
func (_T0kenSplittableDynamic *T0kenSplittableDynamicTransactor) RemoveShareholder(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	return _T0kenSplittableDynamic.contract.Transact(opts, "removeShareholder", addr)
}

// RemoveShareholder is a paid mutator transaction binding the contract method 0x9babdad6.
//
// Solidity: function removeShareholder(addr address) returns()
func (_T0kenSplittableDynamic *T0kenSplittableDynamicSession) RemoveShareholder(addr common.Address) (*types.Transaction, error) {
	return _T0kenSplittableDynamic.Contract.RemoveShareholder(&_T0kenSplittableDynamic.TransactOpts, addr)
}

// RemoveShareholder is a paid mutator transaction binding the contract method 0x9babdad6.
//
// Solidity: function removeShareholder(addr address) returns()
func (_T0kenSplittableDynamic *T0kenSplittableDynamicTransactorSession) RemoveShareholder(addr common.Address) (*types.Transaction, error) {
	return _T0kenSplittableDynamic.Contract.RemoveShareholder(&_T0kenSplittableDynamic.TransactOpts, addr)
}

// SetCompliance is a paid mutator transaction binding the contract method 0xf8981789.
//
// Solidity: function setCompliance(newComplianceAddress address) returns()
func (_T0kenSplittableDynamic *T0kenSplittableDynamicTransactor) SetCompliance(opts *bind.TransactOpts, newComplianceAddress common.Address) (*types.Transaction, error) {
	return _T0kenSplittableDynamic.contract.Transact(opts, "setCompliance", newComplianceAddress)
}

// SetCompliance is a paid mutator transaction binding the contract method 0xf8981789.
//
// Solidity: function setCompliance(newComplianceAddress address) returns()
func (_T0kenSplittableDynamic *T0kenSplittableDynamicSession) SetCompliance(newComplianceAddress common.Address) (*types.Transaction, error) {
	return _T0kenSplittableDynamic.Contract.SetCompliance(&_T0kenSplittableDynamic.TransactOpts, newComplianceAddress)
}

// SetCompliance is a paid mutator transaction binding the contract method 0xf8981789.
//
// Solidity: function setCompliance(newComplianceAddress address) returns()
func (_T0kenSplittableDynamic *T0kenSplittableDynamicTransactorSession) SetCompliance(newComplianceAddress common.Address) (*types.Transaction, error) {
	return _T0kenSplittableDynamic.Contract.SetCompliance(&_T0kenSplittableDynamic.TransactOpts, newComplianceAddress)
}

// SetIssuer is a paid mutator transaction binding the contract method 0x55cc4e57.
//
// Solidity: function setIssuer(newIssuer address) returns()
func (_T0kenSplittableDynamic *T0kenSplittableDynamicTransactor) SetIssuer(opts *bind.TransactOpts, newIssuer common.Address) (*types.Transaction, error) {
	return _T0kenSplittableDynamic.contract.Transact(opts, "setIssuer", newIssuer)
}

// SetIssuer is a paid mutator transaction binding the contract method 0x55cc4e57.
//
// Solidity: function setIssuer(newIssuer address) returns()
func (_T0kenSplittableDynamic *T0kenSplittableDynamicSession) SetIssuer(newIssuer common.Address) (*types.Transaction, error) {
	return _T0kenSplittableDynamic.Contract.SetIssuer(&_T0kenSplittableDynamic.TransactOpts, newIssuer)
}

// SetIssuer is a paid mutator transaction binding the contract method 0x55cc4e57.
//
// Solidity: function setIssuer(newIssuer address) returns()
func (_T0kenSplittableDynamic *T0kenSplittableDynamicTransactorSession) SetIssuer(newIssuer common.Address) (*types.Transaction, error) {
	return _T0kenSplittableDynamic.Contract.SetIssuer(&_T0kenSplittableDynamic.TransactOpts, newIssuer)
}

// SetLocked is a paid mutator transaction binding the contract method 0x211e28b6.
//
// Solidity: function setLocked(locked bool) returns()
func (_T0kenSplittableDynamic *T0kenSplittableDynamicTransactor) SetLocked(opts *bind.TransactOpts, locked bool) (*types.Transaction, error) {
	return _T0kenSplittableDynamic.contract.Transact(opts, "setLocked", locked)
}

// SetLocked is a paid mutator transaction binding the contract method 0x211e28b6.
//
// Solidity: function setLocked(locked bool) returns()
func (_T0kenSplittableDynamic *T0kenSplittableDynamicSession) SetLocked(locked bool) (*types.Transaction, error) {
	return _T0kenSplittableDynamic.Contract.SetLocked(&_T0kenSplittableDynamic.TransactOpts, locked)
}

// SetLocked is a paid mutator transaction binding the contract method 0x211e28b6.
//
// Solidity: function setLocked(locked bool) returns()
func (_T0kenSplittableDynamic *T0kenSplittableDynamicTransactorSession) SetLocked(locked bool) (*types.Transaction, error) {
	return _T0kenSplittableDynamic.Contract.SetLocked(&_T0kenSplittableDynamic.TransactOpts, locked)
}

// SplitTotalSupply is a paid mutator transaction binding the contract method 0x98d34b9b.
//
// Solidity: function splitTotalSupply(cashedOutTokenQuantity uint256) returns(bool)
func (_T0kenSplittableDynamic *T0kenSplittableDynamicTransactor) SplitTotalSupply(opts *bind.TransactOpts, cashedOutTokenQuantity *big.Int) (*types.Transaction, error) {
	return _T0kenSplittableDynamic.contract.Transact(opts, "splitTotalSupply", cashedOutTokenQuantity)
}

// SplitTotalSupply is a paid mutator transaction binding the contract method 0x98d34b9b.
//
// Solidity: function splitTotalSupply(cashedOutTokenQuantity uint256) returns(bool)
func (_T0kenSplittableDynamic *T0kenSplittableDynamicSession) SplitTotalSupply(cashedOutTokenQuantity *big.Int) (*types.Transaction, error) {
	return _T0kenSplittableDynamic.Contract.SplitTotalSupply(&_T0kenSplittableDynamic.TransactOpts, cashedOutTokenQuantity)
}

// SplitTotalSupply is a paid mutator transaction binding the contract method 0x98d34b9b.
//
// Solidity: function splitTotalSupply(cashedOutTokenQuantity uint256) returns(bool)
func (_T0kenSplittableDynamic *T0kenSplittableDynamicTransactorSession) SplitTotalSupply(cashedOutTokenQuantity *big.Int) (*types.Transaction, error) {
	return _T0kenSplittableDynamic.Contract.SplitTotalSupply(&_T0kenSplittableDynamic.TransactOpts, cashedOutTokenQuantity)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(to address, tokens uint256) returns(bool)
func (_T0kenSplittableDynamic *T0kenSplittableDynamicTransactor) Transfer(opts *bind.TransactOpts, to common.Address, tokens *big.Int) (*types.Transaction, error) {
	return _T0kenSplittableDynamic.contract.Transact(opts, "transfer", to, tokens)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(to address, tokens uint256) returns(bool)
func (_T0kenSplittableDynamic *T0kenSplittableDynamicSession) Transfer(to common.Address, tokens *big.Int) (*types.Transaction, error) {
	return _T0kenSplittableDynamic.Contract.Transfer(&_T0kenSplittableDynamic.TransactOpts, to, tokens)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(to address, tokens uint256) returns(bool)
func (_T0kenSplittableDynamic *T0kenSplittableDynamicTransactorSession) Transfer(to common.Address, tokens *big.Int) (*types.Transaction, error) {
	return _T0kenSplittableDynamic.Contract.Transfer(&_T0kenSplittableDynamic.TransactOpts, to, tokens)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(from address, to address, tokens uint256) returns(bool)
func (_T0kenSplittableDynamic *T0kenSplittableDynamicTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokens *big.Int) (*types.Transaction, error) {
	return _T0kenSplittableDynamic.contract.Transact(opts, "transferFrom", from, to, tokens)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(from address, to address, tokens uint256) returns(bool)
func (_T0kenSplittableDynamic *T0kenSplittableDynamicSession) TransferFrom(from common.Address, to common.Address, tokens *big.Int) (*types.Transaction, error) {
	return _T0kenSplittableDynamic.Contract.TransferFrom(&_T0kenSplittableDynamic.TransactOpts, from, to, tokens)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(from address, to address, tokens uint256) returns(bool)
func (_T0kenSplittableDynamic *T0kenSplittableDynamicTransactorSession) TransferFrom(from common.Address, to common.Address, tokens *big.Int) (*types.Transaction, error) {
	return _T0kenSplittableDynamic.Contract.TransferFrom(&_T0kenSplittableDynamic.TransactOpts, from, to, tokens)
}

// TransferOverride is a paid mutator transaction binding the contract method 0x80318be8.
//
// Solidity: function transferOverride(from address, to address, tokens uint256) returns(bool)
func (_T0kenSplittableDynamic *T0kenSplittableDynamicTransactor) TransferOverride(opts *bind.TransactOpts, from common.Address, to common.Address, tokens *big.Int) (*types.Transaction, error) {
	return _T0kenSplittableDynamic.contract.Transact(opts, "transferOverride", from, to, tokens)
}

// TransferOverride is a paid mutator transaction binding the contract method 0x80318be8.
//
// Solidity: function transferOverride(from address, to address, tokens uint256) returns(bool)
func (_T0kenSplittableDynamic *T0kenSplittableDynamicSession) TransferOverride(from common.Address, to common.Address, tokens *big.Int) (*types.Transaction, error) {
	return _T0kenSplittableDynamic.Contract.TransferOverride(&_T0kenSplittableDynamic.TransactOpts, from, to, tokens)
}

// TransferOverride is a paid mutator transaction binding the contract method 0x80318be8.
//
// Solidity: function transferOverride(from address, to address, tokens uint256) returns(bool)
func (_T0kenSplittableDynamic *T0kenSplittableDynamicTransactorSession) TransferOverride(from common.Address, to common.Address, tokens *big.Int) (*types.Transaction, error) {
	return _T0kenSplittableDynamic.Contract.TransferOverride(&_T0kenSplittableDynamic.TransactOpts, from, to, tokens)
}

// TransferOwner is a paid mutator transaction binding the contract method 0x4fb2e45d.
//
// Solidity: function transferOwner(newOwner address) returns()
func (_T0kenSplittableDynamic *T0kenSplittableDynamicTransactor) TransferOwner(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _T0kenSplittableDynamic.contract.Transact(opts, "transferOwner", newOwner)
}

// TransferOwner is a paid mutator transaction binding the contract method 0x4fb2e45d.
//
// Solidity: function transferOwner(newOwner address) returns()
func (_T0kenSplittableDynamic *T0kenSplittableDynamicSession) TransferOwner(newOwner common.Address) (*types.Transaction, error) {
	return _T0kenSplittableDynamic.Contract.TransferOwner(&_T0kenSplittableDynamic.TransactOpts, newOwner)
}

// TransferOwner is a paid mutator transaction binding the contract method 0x4fb2e45d.
//
// Solidity: function transferOwner(newOwner address) returns()
func (_T0kenSplittableDynamic *T0kenSplittableDynamicTransactorSession) TransferOwner(newOwner common.Address) (*types.Transaction, error) {
	return _T0kenSplittableDynamic.Contract.TransferOwner(&_T0kenSplittableDynamic.TransactOpts, newOwner)
}

// UpdateSplit is a paid mutator transaction binding the contract method 0x872a72ea.
//
// Solidity: function updateSplit(numerator uint128, denominator uint128) returns()
func (_T0kenSplittableDynamic *T0kenSplittableDynamicTransactor) UpdateSplit(opts *bind.TransactOpts, numerator *big.Int, denominator *big.Int) (*types.Transaction, error) {
	return _T0kenSplittableDynamic.contract.Transact(opts, "updateSplit", numerator, denominator)
}

// UpdateSplit is a paid mutator transaction binding the contract method 0x872a72ea.
//
// Solidity: function updateSplit(numerator uint128, denominator uint128) returns()
func (_T0kenSplittableDynamic *T0kenSplittableDynamicSession) UpdateSplit(numerator *big.Int, denominator *big.Int) (*types.Transaction, error) {
	return _T0kenSplittableDynamic.Contract.UpdateSplit(&_T0kenSplittableDynamic.TransactOpts, numerator, denominator)
}

// UpdateSplit is a paid mutator transaction binding the contract method 0x872a72ea.
//
// Solidity: function updateSplit(numerator uint128, denominator uint128) returns()
func (_T0kenSplittableDynamic *T0kenSplittableDynamicTransactorSession) UpdateSplit(numerator *big.Int, denominator *big.Int) (*types.Transaction, error) {
	return _T0kenSplittableDynamic.Contract.UpdateSplit(&_T0kenSplittableDynamic.TransactOpts, numerator, denominator)
}

// T0kenSplittableDynamicApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the T0kenSplittableDynamic contract.
type T0kenSplittableDynamicApprovalIterator struct {
	Event *T0kenSplittableDynamicApproval // Event containing the contract specifics and raw log

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
func (it *T0kenSplittableDynamicApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(T0kenSplittableDynamicApproval)
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
		it.Event = new(T0kenSplittableDynamicApproval)
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
func (it *T0kenSplittableDynamicApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *T0kenSplittableDynamicApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// T0kenSplittableDynamicApproval represents a Approval event raised by the T0kenSplittableDynamic contract.
type T0kenSplittableDynamicApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: e Approval(owner indexed address, spender indexed address, value uint256)
func (_T0kenSplittableDynamic *T0kenSplittableDynamicFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*T0kenSplittableDynamicApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _T0kenSplittableDynamic.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &T0kenSplittableDynamicApprovalIterator{contract: _T0kenSplittableDynamic.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: e Approval(owner indexed address, spender indexed address, value uint256)
func (_T0kenSplittableDynamic *T0kenSplittableDynamicFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *T0kenSplittableDynamicApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _T0kenSplittableDynamic.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(T0kenSplittableDynamicApproval)
				if err := _T0kenSplittableDynamic.contract.UnpackLog(event, "Approval", log); err != nil {
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

// T0kenSplittableDynamicIssuanceIterator is returned from FilterIssuance and is used to iterate over the raw logs and unpacked data for Issuance events raised by the T0kenSplittableDynamic contract.
type T0kenSplittableDynamicIssuanceIterator struct {
	Event *T0kenSplittableDynamicIssuance // Event containing the contract specifics and raw log

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
func (it *T0kenSplittableDynamicIssuanceIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(T0kenSplittableDynamicIssuance)
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
		it.Event = new(T0kenSplittableDynamicIssuance)
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
func (it *T0kenSplittableDynamicIssuanceIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *T0kenSplittableDynamicIssuanceIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// T0kenSplittableDynamicIssuance represents a Issuance event raised by the T0kenSplittableDynamic contract.
type T0kenSplittableDynamicIssuance struct {
	To     common.Address
	Tokens *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterIssuance is a free log retrieval operation binding the contract event 0x9cb9c14f7bc76e3a89b796b091850526236115352a198b1e472f00e91376bbcb.
//
// Solidity: e Issuance(to indexed address, tokens uint256)
func (_T0kenSplittableDynamic *T0kenSplittableDynamicFilterer) FilterIssuance(opts *bind.FilterOpts, to []common.Address) (*T0kenSplittableDynamicIssuanceIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _T0kenSplittableDynamic.contract.FilterLogs(opts, "Issuance", toRule)
	if err != nil {
		return nil, err
	}
	return &T0kenSplittableDynamicIssuanceIterator{contract: _T0kenSplittableDynamic.contract, event: "Issuance", logs: logs, sub: sub}, nil
}

// WatchIssuance is a free log subscription operation binding the contract event 0x9cb9c14f7bc76e3a89b796b091850526236115352a198b1e472f00e91376bbcb.
//
// Solidity: e Issuance(to indexed address, tokens uint256)
func (_T0kenSplittableDynamic *T0kenSplittableDynamicFilterer) WatchIssuance(opts *bind.WatchOpts, sink chan<- *T0kenSplittableDynamicIssuance, to []common.Address) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _T0kenSplittableDynamic.contract.WatchLogs(opts, "Issuance", toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(T0kenSplittableDynamicIssuance)
				if err := _T0kenSplittableDynamic.contract.UnpackLog(event, "Issuance", log); err != nil {
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

// T0kenSplittableDynamicIssuanceFinishedIterator is returned from FilterIssuanceFinished and is used to iterate over the raw logs and unpacked data for IssuanceFinished events raised by the T0kenSplittableDynamic contract.
type T0kenSplittableDynamicIssuanceFinishedIterator struct {
	Event *T0kenSplittableDynamicIssuanceFinished // Event containing the contract specifics and raw log

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
func (it *T0kenSplittableDynamicIssuanceFinishedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(T0kenSplittableDynamicIssuanceFinished)
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
		it.Event = new(T0kenSplittableDynamicIssuanceFinished)
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
func (it *T0kenSplittableDynamicIssuanceFinishedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *T0kenSplittableDynamicIssuanceFinishedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// T0kenSplittableDynamicIssuanceFinished represents a IssuanceFinished event raised by the T0kenSplittableDynamic contract.
type T0kenSplittableDynamicIssuanceFinished struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterIssuanceFinished is a free log retrieval operation binding the contract event 0x29fe76cc5ca143e91eadf7242fda487fcef09318c1237900f958abe1e2c5beff.
//
// Solidity: e IssuanceFinished()
func (_T0kenSplittableDynamic *T0kenSplittableDynamicFilterer) FilterIssuanceFinished(opts *bind.FilterOpts) (*T0kenSplittableDynamicIssuanceFinishedIterator, error) {

	logs, sub, err := _T0kenSplittableDynamic.contract.FilterLogs(opts, "IssuanceFinished")
	if err != nil {
		return nil, err
	}
	return &T0kenSplittableDynamicIssuanceFinishedIterator{contract: _T0kenSplittableDynamic.contract, event: "IssuanceFinished", logs: logs, sub: sub}, nil
}

// WatchIssuanceFinished is a free log subscription operation binding the contract event 0x29fe76cc5ca143e91eadf7242fda487fcef09318c1237900f958abe1e2c5beff.
//
// Solidity: e IssuanceFinished()
func (_T0kenSplittableDynamic *T0kenSplittableDynamicFilterer) WatchIssuanceFinished(opts *bind.WatchOpts, sink chan<- *T0kenSplittableDynamicIssuanceFinished) (event.Subscription, error) {

	logs, sub, err := _T0kenSplittableDynamic.contract.WatchLogs(opts, "IssuanceFinished")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(T0kenSplittableDynamicIssuanceFinished)
				if err := _T0kenSplittableDynamic.contract.UnpackLog(event, "IssuanceFinished", log); err != nil {
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

// T0kenSplittableDynamicIssuerSetIterator is returned from FilterIssuerSet and is used to iterate over the raw logs and unpacked data for IssuerSet events raised by the T0kenSplittableDynamic contract.
type T0kenSplittableDynamicIssuerSetIterator struct {
	Event *T0kenSplittableDynamicIssuerSet // Event containing the contract specifics and raw log

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
func (it *T0kenSplittableDynamicIssuerSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(T0kenSplittableDynamicIssuerSet)
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
		it.Event = new(T0kenSplittableDynamicIssuerSet)
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
func (it *T0kenSplittableDynamicIssuerSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *T0kenSplittableDynamicIssuerSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// T0kenSplittableDynamicIssuerSet represents a IssuerSet event raised by the T0kenSplittableDynamic contract.
type T0kenSplittableDynamicIssuerSet struct {
	PreviousIssuer common.Address
	NewIssuer      common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterIssuerSet is a free log retrieval operation binding the contract event 0xf7189b85d7899f5a32d733e6584c4f1dcdff0274f09d969d186c1797673ede11.
//
// Solidity: e IssuerSet(previousIssuer indexed address, newIssuer indexed address)
func (_T0kenSplittableDynamic *T0kenSplittableDynamicFilterer) FilterIssuerSet(opts *bind.FilterOpts, previousIssuer []common.Address, newIssuer []common.Address) (*T0kenSplittableDynamicIssuerSetIterator, error) {

	var previousIssuerRule []interface{}
	for _, previousIssuerItem := range previousIssuer {
		previousIssuerRule = append(previousIssuerRule, previousIssuerItem)
	}
	var newIssuerRule []interface{}
	for _, newIssuerItem := range newIssuer {
		newIssuerRule = append(newIssuerRule, newIssuerItem)
	}

	logs, sub, err := _T0kenSplittableDynamic.contract.FilterLogs(opts, "IssuerSet", previousIssuerRule, newIssuerRule)
	if err != nil {
		return nil, err
	}
	return &T0kenSplittableDynamicIssuerSetIterator{contract: _T0kenSplittableDynamic.contract, event: "IssuerSet", logs: logs, sub: sub}, nil
}

// WatchIssuerSet is a free log subscription operation binding the contract event 0xf7189b85d7899f5a32d733e6584c4f1dcdff0274f09d969d186c1797673ede11.
//
// Solidity: e IssuerSet(previousIssuer indexed address, newIssuer indexed address)
func (_T0kenSplittableDynamic *T0kenSplittableDynamicFilterer) WatchIssuerSet(opts *bind.WatchOpts, sink chan<- *T0kenSplittableDynamicIssuerSet, previousIssuer []common.Address, newIssuer []common.Address) (event.Subscription, error) {

	var previousIssuerRule []interface{}
	for _, previousIssuerItem := range previousIssuer {
		previousIssuerRule = append(previousIssuerRule, previousIssuerItem)
	}
	var newIssuerRule []interface{}
	for _, newIssuerItem := range newIssuer {
		newIssuerRule = append(newIssuerRule, newIssuerItem)
	}

	logs, sub, err := _T0kenSplittableDynamic.contract.WatchLogs(opts, "IssuerSet", previousIssuerRule, newIssuerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(T0kenSplittableDynamicIssuerSet)
				if err := _T0kenSplittableDynamic.contract.UnpackLog(event, "IssuerSet", log); err != nil {
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

// T0kenSplittableDynamicOwnerTransferredIterator is returned from FilterOwnerTransferred and is used to iterate over the raw logs and unpacked data for OwnerTransferred events raised by the T0kenSplittableDynamic contract.
type T0kenSplittableDynamicOwnerTransferredIterator struct {
	Event *T0kenSplittableDynamicOwnerTransferred // Event containing the contract specifics and raw log

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
func (it *T0kenSplittableDynamicOwnerTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(T0kenSplittableDynamicOwnerTransferred)
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
		it.Event = new(T0kenSplittableDynamicOwnerTransferred)
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
func (it *T0kenSplittableDynamicOwnerTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *T0kenSplittableDynamicOwnerTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// T0kenSplittableDynamicOwnerTransferred represents a OwnerTransferred event raised by the T0kenSplittableDynamic contract.
type T0kenSplittableDynamicOwnerTransferred struct {
	OldOwner common.Address
	NewOwner common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOwnerTransferred is a free log retrieval operation binding the contract event 0x8934ce4adea8d9ce0d714d2c22b86790e41b7731c84b926fbbdc1d40ff6533c9.
//
// Solidity: e OwnerTransferred(oldOwner indexed address, newOwner indexed address)
func (_T0kenSplittableDynamic *T0kenSplittableDynamicFilterer) FilterOwnerTransferred(opts *bind.FilterOpts, oldOwner []common.Address, newOwner []common.Address) (*T0kenSplittableDynamicOwnerTransferredIterator, error) {

	var oldOwnerRule []interface{}
	for _, oldOwnerItem := range oldOwner {
		oldOwnerRule = append(oldOwnerRule, oldOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _T0kenSplittableDynamic.contract.FilterLogs(opts, "OwnerTransferred", oldOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &T0kenSplittableDynamicOwnerTransferredIterator{contract: _T0kenSplittableDynamic.contract, event: "OwnerTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnerTransferred is a free log subscription operation binding the contract event 0x8934ce4adea8d9ce0d714d2c22b86790e41b7731c84b926fbbdc1d40ff6533c9.
//
// Solidity: e OwnerTransferred(oldOwner indexed address, newOwner indexed address)
func (_T0kenSplittableDynamic *T0kenSplittableDynamicFilterer) WatchOwnerTransferred(opts *bind.WatchOpts, sink chan<- *T0kenSplittableDynamicOwnerTransferred, oldOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var oldOwnerRule []interface{}
	for _, oldOwnerItem := range oldOwner {
		oldOwnerRule = append(oldOwnerRule, oldOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _T0kenSplittableDynamic.contract.WatchLogs(opts, "OwnerTransferred", oldOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(T0kenSplittableDynamicOwnerTransferred)
				if err := _T0kenSplittableDynamic.contract.UnpackLog(event, "OwnerTransferred", log); err != nil {
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

// T0kenSplittableDynamicShareholderAddedIterator is returned from FilterShareholderAdded and is used to iterate over the raw logs and unpacked data for ShareholderAdded events raised by the T0kenSplittableDynamic contract.
type T0kenSplittableDynamicShareholderAddedIterator struct {
	Event *T0kenSplittableDynamicShareholderAdded // Event containing the contract specifics and raw log

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
func (it *T0kenSplittableDynamicShareholderAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(T0kenSplittableDynamicShareholderAdded)
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
		it.Event = new(T0kenSplittableDynamicShareholderAdded)
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
func (it *T0kenSplittableDynamicShareholderAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *T0kenSplittableDynamicShareholderAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// T0kenSplittableDynamicShareholderAdded represents a ShareholderAdded event raised by the T0kenSplittableDynamic contract.
type T0kenSplittableDynamicShareholderAdded struct {
	Shareholder common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterShareholderAdded is a free log retrieval operation binding the contract event 0x3082c1c4977de80c4f67ee77b56b282610ec93a7ecbcc31b551e0ac2f7bd0395.
//
// Solidity: e ShareholderAdded(shareholder address)
func (_T0kenSplittableDynamic *T0kenSplittableDynamicFilterer) FilterShareholderAdded(opts *bind.FilterOpts) (*T0kenSplittableDynamicShareholderAddedIterator, error) {

	logs, sub, err := _T0kenSplittableDynamic.contract.FilterLogs(opts, "ShareholderAdded")
	if err != nil {
		return nil, err
	}
	return &T0kenSplittableDynamicShareholderAddedIterator{contract: _T0kenSplittableDynamic.contract, event: "ShareholderAdded", logs: logs, sub: sub}, nil
}

// WatchShareholderAdded is a free log subscription operation binding the contract event 0x3082c1c4977de80c4f67ee77b56b282610ec93a7ecbcc31b551e0ac2f7bd0395.
//
// Solidity: e ShareholderAdded(shareholder address)
func (_T0kenSplittableDynamic *T0kenSplittableDynamicFilterer) WatchShareholderAdded(opts *bind.WatchOpts, sink chan<- *T0kenSplittableDynamicShareholderAdded) (event.Subscription, error) {

	logs, sub, err := _T0kenSplittableDynamic.contract.WatchLogs(opts, "ShareholderAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(T0kenSplittableDynamicShareholderAdded)
				if err := _T0kenSplittableDynamic.contract.UnpackLog(event, "ShareholderAdded", log); err != nil {
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

// T0kenSplittableDynamicShareholderRemovedIterator is returned from FilterShareholderRemoved and is used to iterate over the raw logs and unpacked data for ShareholderRemoved events raised by the T0kenSplittableDynamic contract.
type T0kenSplittableDynamicShareholderRemovedIterator struct {
	Event *T0kenSplittableDynamicShareholderRemoved // Event containing the contract specifics and raw log

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
func (it *T0kenSplittableDynamicShareholderRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(T0kenSplittableDynamicShareholderRemoved)
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
		it.Event = new(T0kenSplittableDynamicShareholderRemoved)
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
func (it *T0kenSplittableDynamicShareholderRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *T0kenSplittableDynamicShareholderRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// T0kenSplittableDynamicShareholderRemoved represents a ShareholderRemoved event raised by the T0kenSplittableDynamic contract.
type T0kenSplittableDynamicShareholderRemoved struct {
	Shareholder common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterShareholderRemoved is a free log retrieval operation binding the contract event 0x7ba114ff3d9844510a088eea94cd35562e7c97a2d36a418b37b2e61e5c77affe.
//
// Solidity: e ShareholderRemoved(shareholder address)
func (_T0kenSplittableDynamic *T0kenSplittableDynamicFilterer) FilterShareholderRemoved(opts *bind.FilterOpts) (*T0kenSplittableDynamicShareholderRemovedIterator, error) {

	logs, sub, err := _T0kenSplittableDynamic.contract.FilterLogs(opts, "ShareholderRemoved")
	if err != nil {
		return nil, err
	}
	return &T0kenSplittableDynamicShareholderRemovedIterator{contract: _T0kenSplittableDynamic.contract, event: "ShareholderRemoved", logs: logs, sub: sub}, nil
}

// WatchShareholderRemoved is a free log subscription operation binding the contract event 0x7ba114ff3d9844510a088eea94cd35562e7c97a2d36a418b37b2e61e5c77affe.
//
// Solidity: e ShareholderRemoved(shareholder address)
func (_T0kenSplittableDynamic *T0kenSplittableDynamicFilterer) WatchShareholderRemoved(opts *bind.WatchOpts, sink chan<- *T0kenSplittableDynamicShareholderRemoved) (event.Subscription, error) {

	logs, sub, err := _T0kenSplittableDynamic.contract.WatchLogs(opts, "ShareholderRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(T0kenSplittableDynamicShareholderRemoved)
				if err := _T0kenSplittableDynamic.contract.UnpackLog(event, "ShareholderRemoved", log); err != nil {
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

// T0kenSplittableDynamicTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the T0kenSplittableDynamic contract.
type T0kenSplittableDynamicTransferIterator struct {
	Event *T0kenSplittableDynamicTransfer // Event containing the contract specifics and raw log

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
func (it *T0kenSplittableDynamicTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(T0kenSplittableDynamicTransfer)
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
		it.Event = new(T0kenSplittableDynamicTransfer)
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
func (it *T0kenSplittableDynamicTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *T0kenSplittableDynamicTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// T0kenSplittableDynamicTransfer represents a Transfer event raised by the T0kenSplittableDynamic contract.
type T0kenSplittableDynamicTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: e Transfer(from indexed address, to indexed address, value uint256)
func (_T0kenSplittableDynamic *T0kenSplittableDynamicFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*T0kenSplittableDynamicTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _T0kenSplittableDynamic.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &T0kenSplittableDynamicTransferIterator{contract: _T0kenSplittableDynamic.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: e Transfer(from indexed address, to indexed address, value uint256)
func (_T0kenSplittableDynamic *T0kenSplittableDynamicFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *T0kenSplittableDynamicTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _T0kenSplittableDynamic.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(T0kenSplittableDynamicTransfer)
				if err := _T0kenSplittableDynamic.contract.UnpackLog(event, "Transfer", log); err != nil {
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
