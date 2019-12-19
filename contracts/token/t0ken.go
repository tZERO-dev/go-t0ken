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

// T0kenABI is the input ABI used to generate the binding from.
const T0kenABI = "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"tokenName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"tokenSymbol\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"decimalPlaces\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokens\",\"type\":\"uint256\"}],\"name\":\"Issuance\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"IssuanceFinished\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousIssuer\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newIssuer\",\"type\":\"address\"}],\"name\":\"IssuerSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnerTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"shareholder\",\"type\":\"address\"}],\"name\":\"ShareholderAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"shareholder\",\"type\":\"address\"}],\"name\":\"ShareholderRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[],\"name\":\"ZERO_ADDRESS\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"addrOwner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokens\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"balances\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"compliance\",\"outputs\":[{\"internalType\":\"contractCompliance\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"finishIssuance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"int256\",\"name\":\"index\",\"type\":\"int256\"}],\"name\":\"holderAt\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"holders\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"count\",\"type\":\"int256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"isHolder\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isLocked\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"issuanceFinished\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"quantity\",\"type\":\"uint256\"}],\"name\":\"issueTokens\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"issuer\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"kill\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"newComplianceAddress\",\"type\":\"address\"}],\"name\":\"setCompliance\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"newIssuer\",\"type\":\"address\"}],\"name\":\"setIssuer\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bool\",\"name\":\"locked\",\"type\":\"bool\"}],\"name\":\"setLocked\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokens\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokens\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokens\",\"type\":\"uint256\"}],\"name\":\"transferOverride\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwner\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// T0kenBin is the compiled bytecode used for deploying new contracts.
const T0kenBin = `6080604052600180546001600160a01b03191690553480156200002157600080fd5b5060405162001e9a38038062001e9a833981810160405260608110156200004757600080fd5b81019080805160405193929190846401000000008211156200006857600080fd5b9083019060208201858111156200007e57600080fd5b82516401000000008111828201881017156200009957600080fd5b82525081516020918201929091019080838360005b83811015620000c8578181015183820152602001620000ae565b50505050905090810190601f168015620000f65780820380516001836020036101000a031916815260200191505b50604052602001805160405193929190846401000000008211156200011a57600080fd5b9083019060208201858111156200013057600080fd5b82516401000000008111828201881017156200014b57600080fd5b82525081516020918201929091019080838360005b838110156200017a57818101518382015260200162000160565b50505050905090810190601f168015620001a85780820380516001836020036101000a031916815260200191505b50604052602090810151600080546001600160a01b031916331790556001805460ff60a01b191690558551909350620001e892506002918601906200021c565b508151620001fe9060039060208501906200021c565b506004805460ff191660ff9290921691909117905550620002c19050565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106200025f57805160ff19168380011785556200028f565b828001600101855582156200028f579182015b828111156200028f57825182559160200191906001019062000272565b506200029d929150620002a1565b5090565b620002be91905b808211156200029d5760008155600101620002a8565b90565b611bc980620002d16000396000f3fe608060405234801561001057600080fd5b50600436106101b95760003560e01c80636290865d116100f9578063a4e2d63411610097578063c422293b11610071578063c422293b146104a2578063d4d7b19a146104aa578063dd62ed3e146104d0578063f8981789146104fe576101b9565b8063a4e2d63414610451578063a5820daa14610459578063a9059cbb14610476576101b9565b80638082a929116100d35780638082a9291461041c5780638188f71c146104395780638da5cb5b1461044157806395d89b4114610449576101b9565b80636290865d146103b857806370a08231146103c057806380318be8146103e6576101b9565b806327e235e3116101665780634662299a116101405780634662299a1461035c5780634fb2e45d14610364578063538ba4f91461038a57806355cc4e5714610392576101b9565b806327e235e314610310578063313ce5671461033657806341c0e1b514610354576101b9565b80631d143848116101975780631d14384814610295578063211e28b6146102b957806323b872dd146102da576101b9565b806306fdde03146101be578063095ea7b31461023b57806318160ddd1461027b575b600080fd5b6101c6610524565b6040805160208082528351818301528351919283929083019185019080838360005b838110156102005781810151838201526020016101e8565b50505050905090810190601f16801561022d5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b6102676004803603604081101561025157600080fd5b506001600160a01b0381351690602001356105af565b604080519115158252519081900360200190f35b6102836106c5565b60408051918252519081900360200190f35b61029d6106cb565b604080516001600160a01b039092168252519081900360200190f35b6102d8600480360360208110156102cf57600080fd5b503515156106da565b005b610267600480360360608110156102f057600080fd5b506001600160a01b038135811691602081013590911690604001356107c2565b6102836004803603602081101561032657600080fd5b50356001600160a01b0316610970565b61033e610982565b6040805160ff9092168252519081900360200190f35b6102d861098b565b610267610a13565b6102d86004803603602081101561037a57600080fd5b50356001600160a01b0316610a23565b61029d610b95565b6102d8600480360360208110156103a857600080fd5b50356001600160a01b0316610ba4565b61029d610d2b565b610283600480360360208110156103d657600080fd5b50356001600160a01b0316610d3a565b610267600480360360608110156103fc57600080fd5b506001600160a01b03813581169160208101359091169060400135610d55565b61029d6004803603602081101561043257600080fd5b5035610e29565b610283610e3c565b61029d610e42565b6101c6610e51565b610267610eac565b6102676004803603602081101561046f57600080fd5b5035610ebc565b6102676004803603604081101561048c57600080fd5b506001600160a01b0381351690602001356110d4565b610267611283565b610267600480360360208110156104c057600080fd5b50356001600160a01b03166113da565b610283600480360360408110156104e657600080fd5b506001600160a01b03813581169160200135166113ed565b6102d86004803603602081101561051457600080fd5b50356001600160a01b0316611418565b6002805460408051602060018416156101000260001901909316849004601f810184900484028201840190925281815292918301828280156105a75780601f1061057c576101008083540402835291602001916105a7565b820191906000526020600020905b81548152906001019060200180831161058a57829003601f168201915b505050505081565b600154600090600160a01b900460ff16156105fb5760405162461bcd60e51b815260040180806020018281038252602d815260200180611b68602d913960400191505060405180910390fd5b61060c60053363ffffffff6114fd16565b61065d576040805162461bcd60e51b815260206004820152601560248201527f4d7573742062652061207368617265686f6c6465720000000000000000000000604482015290519081900360640190fd5b336000818152600c602090815260408083206001600160a01b03881680855290835292819020869055805186815290519293927f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925929181900390910190a35060015b92915050565b600b5490565b6009546001600160a01b031681565b6000546001600160a01b031633148061070357506001546000546001600160a01b039081169116145b610754576040805162461bcd60e51b815260206004820152601960248201527f4f776e6572206163636f756e7420697320726571756972656400000000000000604482015290519081900360640190fd5b60015460ff600160a01b90910416151581151514156107a45760405162461bcd60e51b8152600401808060200182810382526028815260200180611b406028913960400191505060405180910390fd5b60018054911515600160a01b0260ff60a01b19909216919091179055565b600154600090600160a01b900460ff161561080e5760405162461bcd60e51b815260040180806020018281038252602d815260200180611b68602d913960400191505060405180910390fd5b6001600160a01b0384166000908152600a602052604090205484908390811115610874576040805162461bcd60e51b8152602060048201526012602482015271496e73756666696369656e742066756e647360701b604482015290519081900360640190fd5b6001600160a01b0386166000908152600c602090815260408083203384529091529020548411156108ec576040805162461bcd60e51b815260206004820152601a60248201527f5472616e73666572206578636565647320616c6c6f77616e6365000000000000604482015290519081900360640190fd5b60006108fb8787876000611534565b90508015610966576001600160a01b0387166000908152600c60209081526040808320338452909152902054610937908663ffffffff61168216565b6001600160a01b0388166000908152600c602090815260408083203384529091529020556109668787876116df565b9695505050505050565b600a6020526000908152604090205481565b60045460ff1681565b6000546001600160a01b03163314806109b457506001546000546001600160a01b039081169116145b610a05576040805162461bcd60e51b815260206004820152601960248201527f4f776e6572206163636f756e7420697320726571756972656400000000000000604482015290519081900360640190fd5b6000546001600160a01b0316ff5b600954600160a01b900460ff1681565b6000546001600160a01b0316331480610a4c57506001546000546001600160a01b039081169116145b610a9d576040805162461bcd60e51b815260206004820152601960248201527f4f776e6572206163636f756e7420697320726571756972656400000000000000604482015290519081900360640190fd5b6000546001600160a01b0382811691161415610aea5760405162461bcd60e51b8152600401808060200182810382526025815260200180611b1b6025913960400191505060405180910390fd5b6001600160a01b038116610b45576040805162461bcd60e51b815260206004820181905260248201527f4e6577204f776e65722063616e6e6f74206265207a65726f2061646472657373604482015290519081900360640190fd5b600080546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8934ce4adea8d9ce0d714d2c22b86790e41b7731c84b926fbbdc1d40ff6533c99190a35050565b6001546001600160a01b031681565b600154600160a01b900460ff1615610bed5760405162461bcd60e51b815260040180806020018281038252602d815260200180611b68602d913960400191505060405180910390fd5b6000546001600160a01b0316331480610c1657506001546000546001600160a01b039081169116145b610c67576040805162461bcd60e51b815260206004820152601960248201527f4f776e6572206163636f756e7420697320726571756972656400000000000000604482015290519081900360640190fd5b610c7860058263ffffffff6114fd16565b15610cb45760405162461bcd60e51b8152600401808060200182810382526021815260200180611afa6021913960400191505060405180910390fd5b6009546001600160a01b03166000818152600a6020526040902054610cdb919083906116df565b600980546001600160a01b0319166001600160a01b0383811691821792839055604051919216907ff7189b85d7899f5a32d733e6584c4f1dcdff0274f09d969d186c1797673ede1190600090a350565b6008546001600160a01b031681565b6001600160a01b03166000908152600a602052604090205490565b600154600090600160a01b900460ff1615610da15760405162461bcd60e51b815260040180806020018281038252602d815260200180611b68602d913960400191505060405180910390fd5b6001600160a01b0384166000908152600a602052604090205484908390811115610e07576040805162461bcd60e51b8152602060048201526012602482015271496e73756666696369656e742066756e647360701b604482015290519081900360640190fd5b6000610e168787876001611534565b90508015610966576109668787876116df565b60006106bf60058363ffffffff61188216565b60055481565b6000546001600160a01b031681565b6003805460408051602060026001851615610100026000190190941693909304601f810184900484028201840190925281815292918301828280156105a75780601f1061057c576101008083540402835291602001916105a7565b600154600160a01b900460ff1681565b600154600090600160a01b900460ff1615610f085760405162461bcd60e51b815260040180806020018281038252602d815260200180611b68602d913960400191505060405180910390fd5b6009546001600160a01b03163314610f67576040805162461bcd60e51b815260206004820152601360248201527f4f6e6c792069737375657220616c6c6f77656400000000000000000000000000604482015290519081900360640190fd5b600954600160a01b900460ff1615610fc6576040805162461bcd60e51b815260206004820152601960248201527f49737375616e636520616c72656164792066696e697368656400000000000000604482015290519081900360640190fd5b811561104457600b54610fdf908363ffffffff61190816565b600b556009546001600160a01b03166000908152600a602052604090205461100d908363ffffffff61190816565b600980546001600160a01b039081166000908152600a602052604090209290925554611042916005911663ffffffff61196916565b505b6009546040805184815290516001600160a01b03909216917f9cb9c14f7bc76e3a89b796b091850526236115352a198b1e472f00e91376bbcb9181900360200190a26009546040805184815290516001600160a01b03909216916000917fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef919081900360200190a3506001919050565b600154600090600160a01b900460ff16156111205760405162461bcd60e51b815260040180806020018281038252602d815260200180611b68602d913960400191505060405180910390fd5b336000818152600a6020526040902054839081111561117b576040805162461bcd60e51b8152602060048201526012602482015271496e73756666696369656e742066756e647360701b604482015290519081900360640190fd5b6009546000906001600160a01b031633141561125957506008546001600160a01b0316158061125457600854600954604080517ffd8258bd0000000000000000000000000000000000000000000000000000000081526001600160a01b039283166004820181905260248201528983166044820152606481018990529051919092169163fd8258bd9160848083019260209291908290030181600087803b15801561122557600080fd5b505af1158015611239573d6000803e3d6000fd5b505050506040513d602081101561124f57600080fd5b505190505b611269565b6112663387876000611534565b90505b801561127a5761127a3387876116df565b95945050505050565b600154600090600160a01b900460ff16156112cf5760405162461bcd60e51b815260040180806020018281038252602d815260200180611b68602d913960400191505060405180910390fd5b6009546001600160a01b0316331461132e576040805162461bcd60e51b815260206004820152601360248201527f4f6e6c792069737375657220616c6c6f77656400000000000000000000000000604482015290519081900360640190fd5b600954600160a01b900460ff161561138d576040805162461bcd60e51b815260206004820152601960248201527f49737375616e636520616c72656164792066696e697368656400000000000000604482015290519081900360640190fd5b6009805460ff60a01b1916600160a01b1790556040517f29fe76cc5ca143e91eadf7242fda487fcef09318c1237900f958abe1e2c5beff90600090a150600954600160a01b900460ff1690565b60006106bf60058363ffffffff6114fd16565b6001600160a01b039182166000908152600c6020908152604080832093909416825291909152205490565b600154600160a01b900460ff16156114615760405162461bcd60e51b815260040180806020018281038252602d815260200180611b68602d913960400191505060405180910390fd5b6000546001600160a01b031633148061148a57506001546000546001600160a01b039081169116145b6114db576040805162461bcd60e51b815260206004820152601960248201527f4f776e6572206163636f756e7420697320726571756972656400000000000000604482015290519081900360640190fd5b600880546001600160a01b0319166001600160a01b0392909216919091179055565b6001600160a01b03811660009081526001830160205260408120546000190181811280159061152c5750835481125b949350505050565b6008546000906001600160a01b031661154f5750801561152c565b811561160557600854604080517f5acba2010000000000000000000000000000000000000000000000000000000081523360048201526001600160a01b03888116602483015287811660448301526064820187905291519190921691635acba2019160848083019260209291908290030181600087803b1580156115d257600080fd5b505af11580156115e6573d6000803e3d6000fd5b505050506040513d60208110156115fc57600080fd5b5051905061152c565b600854604080517f6d62a4fe0000000000000000000000000000000000000000000000000000000081523360048201526001600160a01b03888116602483015287811660448301526064820187905291519190921691636d62a4fe9160848083019260209291908290030181600087803b1580156115d257600080fd5b6000828211156116d9576040805162461bcd60e51b815260206004820152601460248201527f526573756c747320696e20756e646572666c6f77000000000000000000000000604482015290519081900360640190fd5b50900390565b6001600160a01b0383166000908152600a6020526040902054611708908263ffffffff61168216565b6001600160a01b038085166000908152600a6020526040808220939093559084168152205461173d908263ffffffff61190816565b6001600160a01b038084166000818152600a602090815260409182902094909455805185815290519193928716927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef92918290030190a36001600160a01b0382166000908152600a6020526040902054158015906117c757506117c760058363ffffffff61196916565b1561180957604080516001600160a01b038416815290517f3082c1c4977de80c4f67ee77b56b282610ec93a7ecbcc31b551e0ac2f7bd03959181900360200190a15b6001600160a01b0383166000908152600a602052604090205415801561183b575061183b60058463ffffffff611a0e16565b1561187d57604080516001600160a01b038516815290517f7ba114ff3d9844510a088eea94cd35562e7c97a2d36a418b37b2e61e5c77affe9181900360200190a15b505050565b60008082121580156118945750825482125b6118e5576040805162461bcd60e51b815260206004820152601860248201527f496e646578206f757473696465206f6620626f756e64732e0000000000000000604482015290519081900360640190fd5b50600101600090815260029190910160205260409020546001600160a01b031690565b600082820183811015611962576040805162461bcd60e51b815260206004820152601360248201527f526573756c747320696e206f766572666c6f7700000000000000000000000000604482015290519081900360640190fd5b9392505050565b60006001600160a01b038216611981575060006106bf565b6001600160a01b0382166000908152600184016020526040812054600019019081128015906119b05750835481125b156119bf5760009150506106bf565b5050815460019081018084556001600160a01b0383166000818152838601602090815260408083208590559382526002870190529190912080546001600160a01b031916909117905592915050565b6001600160a01b03811660009081526001808401602052604082205490811280611a385750835481135b15611a475760009150506106bf565b8354811215611aae57835460009081526002850160208181526040808420546001600160a01b03168085526001890183528185208690558585529290915280832080546001600160a01b031990811690931790558654835290912080549091169055611acd565b6000818152600285016020526040902080546001600160a01b03191690555b50506001600160a01b03166000908152600182810160205260408220919091558154600019019091559056fe4e6577206973737565722063616e27742062652061207368617265686f6c6465724e6577204f776e65722063616e6e6f74206265207468652063757272656e74206f776e6572436f6e747261637420616c726561647920696e20726571756573746564206c6f636b207374617465436f6e74726163742069732063757272656e746c79206c6f636b656420666f72206d6f64696669636174696f6ea265627a7a7231582051b72e7674b4ff2b1210cded3109b432af0eb37b772b360ec33e1f9b68de66b764736f6c634300050c0032`

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
