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

// RegistryABI is the input ABI used to generate the binding from.
const RegistryABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"index\",\"type\":\"int256\"}],\"name\":\"accountAt\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"uint8\"},{\"name\":\"\",\"type\":\"bool\"},{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"},{\"name\":\"kind\",\"type\":\"uint8\"}],\"name\":\"accountExists\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"accountGet\",\"outputs\":[{\"name\":\"\",\"type\":\"uint8\"},{\"name\":\"\",\"type\":\"bool\"},{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"},{\"name\":\"hash\",\"type\":\"bytes32\"}],\"name\":\"setAccountHash\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"kind\",\"type\":\"uint8\"},{\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"grantPermission\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"locked\",\"type\":\"bool\"}],\"name\":\"setLocked\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"accountParent\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"accountKind\",\"outputs\":[{\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"kill\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"},{\"name\":\"index\",\"type\":\"uint8\"},{\"name\":\"customData\",\"type\":\"bytes32\"}],\"name\":\"setAccountData\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwner\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ZERO_ADDRESS\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"uint8\"}],\"name\":\"data\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"kind\",\"type\":\"uint8\"},{\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"revokePermission\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"accountFrozen\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint8\"}],\"name\":\"permissions\",\"outputs\":[{\"name\":\"count\",\"type\":\"int256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"accounts\",\"outputs\":[{\"name\":\"\",\"type\":\"int256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"hash\",\"type\":\"bytes32\"},{\"name\":\"index\",\"type\":\"int256\"}],\"name\":\"accountAtHash\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"uint8\"},{\"name\":\"\",\"type\":\"bool\"},{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"accountKindAndFrozenAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"uint8\"},{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"accountExists\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"accountIndexOf\",\"outputs\":[{\"name\":\"\",\"type\":\"int256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"accountHash\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"kind\",\"type\":\"uint8\"},{\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"permissionExists\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isLocked\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"kind\",\"type\":\"uint8\"},{\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"permissionIndexOf\",\"outputs\":[{\"name\":\"\",\"type\":\"int256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"removeAccount\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"},{\"name\":\"frozen\",\"type\":\"bool\"}],\"name\":\"setAccountFrozen\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"hashes\",\"outputs\":[{\"name\":\"count\",\"type\":\"int256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"kind\",\"type\":\"uint8\"},{\"name\":\"index\",\"type\":\"int256\"}],\"name\":\"permissionAt\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"},{\"name\":\"kind\",\"type\":\"uint8\"},{\"name\":\"isFrozen\",\"type\":\"bool\"},{\"name\":\"parent\",\"type\":\"address\"},{\"name\":\"hash\",\"type\":\"bytes32\"}],\"name\":\"addAccount\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"hash\",\"type\":\"bytes32\"},{\"name\":\"index\",\"type\":\"int256\"}],\"name\":\"addressAtHash\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"oldOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnerTransferred\",\"type\":\"event\"}]"

// RegistryBin is the compiled bytecode used for deploying new contracts.
const RegistryBin = `608060405260018054600080546001600160a01b031916331790556001600160a81b0319169055612523806100356000396000f3fe608060405234801561001057600080fd5b50600436106101ef5760003560e01c806368cd03f61161010f578063a4e2d634116100a2578063d658d2e911610071578063d658d2e9146106ae578063d7379999146106cb578063e8d74d1f146106f1578063fcef15e614610738576101ef565b8063a4e2d63414610623578063a84694c81461062b578063c4740a951461065a578063cbe5404f14610680576101ef565b8063770850c8116100de578063770850c8146105a05780638da5cb5b146105c6578063960d27d3146105ce5780639d44ac4f146105f4576101ef565b806368cd03f6146105045780636a2a1bdc1461050c5780636f6072eb1461052f57806375cd51ed1461057a576101ef565b806341c0e1b511610187578063572bc6b111610156578063572bc6b11461044e578063612eb5e71461048f578063624522f9146104be5780636476b837146104e4576101ef565b806341c0e1b5146103e35780634f1c3b66146103eb5780634fb2e45d14610420578063538ba4f914610446576101ef565b80631a3722b2116101c35780631a3722b214610317578063211e28b61461034657806331aaa74a14610365578063351a97f8146103a7576101ef565b8062f17754146101f45780630a22ee731461024e578063137e37d9146102915780631706acfc146102e9575b600080fd5b6102116004803603602081101561020a57600080fd5b503561075b565b604080516001600160a01b03968716815260ff90951660208601529215158484015293166060830152608082019290925290519081900360a00190f35b61027d6004803603604081101561026457600080fd5b5080356001600160a01b0316906020013560ff1661084b565b604080519115158252519081900360200190f35b6102b7600480360360208110156102a757600080fd5b50356001600160a01b03166108aa565b6040805160ff909516855292151560208501526001600160a01b03909116838301526060830152519081900360800190f35b610315600480360360408110156102ff57600080fd5b506001600160a01b0381351690602001356108e6565b005b6103156004803603604081101561032d57600080fd5b50803560ff1690602001356001600160a01b0316610a7c565b6103156004803603602081101561035c57600080fd5b50351515610bc0565b61038b6004803603602081101561037b57600080fd5b50356001600160a01b0316610ca8565b604080516001600160a01b039092168252519081900360200190f35b6103cd600480360360208110156103bd57600080fd5b50356001600160a01b0316610cbf565b6040805160ff9092168252519081900360200190f35b610315610cd4565b6103156004803603606081101561040157600080fd5b506001600160a01b038135169060ff6020820135169060400135610d5c565b6103156004803603602081101561043657600080fd5b50356001600160a01b0316610f07565b61038b611079565b61047d6004803603604081101561046457600080fd5b5080356001600160a01b0316906020013560ff16611088565b60408051918252519081900360200190f35b610315600480360360408110156104a557600080fd5b50803560ff1690602001356001600160a01b03166110a5565b61027d600480360360208110156104d457600080fd5b50356001600160a01b03166111e3565b61047d600480360360208110156104fa57600080fd5b503560ff166111f8565b61047d61120a565b6102116004803603604081101561052257600080fd5b5080359060200135611210565b6105556004803603602081101561054557600080fd5b50356001600160a01b0316611271565b6040805160ff90931683526001600160a01b0390911660208301528051918290030190f35b61027d6004803603602081101561059057600080fd5b50356001600160a01b0316611455565b61047d600480360360208110156105b657600080fd5b50356001600160a01b0316611486565b61038b6114e5565b61047d600480360360208110156105e457600080fd5b50356001600160a01b03166114f4565b61027d6004803603604081101561060a57600080fd5b50803560ff1690602001356001600160a01b0316611509565b61027d61152b565b61047d6004803603604081101561064157600080fd5b50803560ff1690602001356001600160a01b031661153b565b6103156004803603602081101561067057600080fd5b50356001600160a01b031661155d565b6103156004803603604081101561069657600080fd5b506001600160a01b0381351690602001351515611884565b61047d600480360360208110156106c457600080fd5b50356119ef565b61038b600480360360408110156106e157600080fd5b5060ff8135169060200135611a01565b610315600480360360a081101561070757600080fd5b506001600160a01b03813581169160ff60208201351691604082013515159160608101359091169060800135611a23565b61038b6004803603604081101561074e57600080fd5b5080359060200135611e04565b6000806000806000808612158015610774575060045486125b6107c5576040805162461bcd60e51b815260206004820152601760248201527f496e646578206f757473696465206f6620626f756e6473000000000000000000604482015290519081900360640190fd5b6107cd612446565b505050506001928301600090815260036020908152604091829020825160a08101845281546001600160a01b0380821680845260ff600160a01b84048116968501879052600160a81b909304909216151595830186905297830154909716606082018190526002909201546080909101819052959691959294509250565b6001600160a01b0382166000908152600260205260408120546001811280610874575060045481135b156108835760009150506108a4565b60009081526003602052604090205460ff838116600160a01b909204161490505b92915050565b6000806000806108b8612446565b6108c186611e22565b6020810151604082015160608301516080909301519199909850919650945092505050565b600154600160a01b900460ff161561092f5760405162461bcd60e51b815260040180806020018281038252602d8152602001806124c2602d913960400191505060405180910390fd5b61093882611e22565b6020015160008160ff1611610994576040805162461bcd60e51b815260206004820152601e60248201527f496e76616c69642c206f72206d697373696e67207065726d697373696f6e0000604482015290519081900360640190fd5b6000546001600160a01b03163314610a0e5760ff811660009081526007602052604090206109c8903363ffffffff611f1116565b610a0e576040805162461bcd60e51b815260206004820152601260248201527126b4b9b9b4b733903832b936b4b9b9b4b7b760711b604482015290519081900360640190fd5b6001600160a01b0383166000908152600260208181526040808420548452600382528084209283015484526006909152909120610a51908563ffffffff611f4816565b506000838152600660205260409020610a70908563ffffffff61203316565b50600201919091555050565b600154600160a01b900460ff1615610ac55760405162461bcd60e51b815260040180806020018281038252602d8152602001806124c2602d913960400191505060405180910390fd5b8160008160ff1611610b1e576040805162461bcd60e51b815260206004820152601e60248201527f496e76616c69642c206f72206d697373696e67207065726d697373696f6e0000604482015290519081900360640190fd5b6000546001600160a01b03163314610b985760ff81166000908152600760205260409020610b52903363ffffffff611f1116565b610b98576040805162461bcd60e51b815260206004820152601260248201527126b4b9b9b4b733903832b936b4b9b9b4b7b760711b604482015290519081900360640190fd5b60ff83166000908152600760205260409020610bba908363ffffffff61203316565b50505050565b6000546001600160a01b0316331480610be957506001546000546001600160a01b039081169116145b610c3a576040805162461bcd60e51b815260206004820152601960248201527f4f776e6572206163636f756e7420697320726571756972656400000000000000604482015290519081900360640190fd5b60015460ff600160a01b9091041615158115151415610c8a5760405162461bcd60e51b815260040180806020018281038252602881526020018061249a6028913960400191505060405180910390fd5b60018054911515600160a01b0260ff60a01b19909216919091179055565b6000610cb382611e22565b6060015190505b919050565b6000610cca82611e22565b6020015192915050565b6000546001600160a01b0316331480610cfd57506001546000546001600160a01b039081169116145b610d4e576040805162461bcd60e51b815260206004820152601960248201527f4f776e6572206163636f756e7420697320726571756972656400000000000000604482015290519081900360640190fd5b6000546001600160a01b0316ff5b600154600160a01b900460ff1615610da55760405162461bcd60e51b815260040180806020018281038252602d8152602001806124c2602d913960400191505060405180910390fd5b610dae83611e22565b6020015160008160ff1611610e0a576040805162461bcd60e51b815260206004820152601e60248201527f496e76616c69642c206f72206d697373696e67207065726d697373696f6e0000604482015290519081900360640190fd5b6000546001600160a01b03163314610e845760ff81166000908152600760205260409020610e3e903363ffffffff611f1116565b610e84576040805162461bcd60e51b815260206004820152601260248201527126b4b9b9b4b733903832b936b4b9b9b4b7b760711b604482015290519081900360640190fd5b606460ff841610610edc576040805162461bcd60e51b815260206004820152601760248201527f696e646578206f757473696465206f6620626f756e6473000000000000000000604482015290519081900360640190fd5b506001600160a01b03909216600090815260056020908152604080832060ff90941683529290522055565b6000546001600160a01b0316331480610f3057506001546000546001600160a01b039081169116145b610f81576040805162461bcd60e51b815260206004820152601960248201527f4f776e6572206163636f756e7420697320726571756972656400000000000000604482015290519081900360640190fd5b6000546001600160a01b0382811691161415610fce5760405162461bcd60e51b81526004018080602001828103825260258152602001806124756025913960400191505060405180910390fd5b6001600160a01b038116611029576040805162461bcd60e51b815260206004820181905260248201527f4e6577204f776e65722063616e6e6f74206265207a65726f2061646472657373604482015290519081900360640190fd5b600080546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8934ce4adea8d9ce0d714d2c22b86790e41b7731c84b926fbbdc1d40ff6533c99190a35050565b6001546001600160a01b031681565b600560209081526000928352604080842090915290825290205481565b600154600160a01b900460ff16156110ee5760405162461bcd60e51b815260040180806020018281038252602d8152602001806124c2602d913960400191505060405180910390fd5b8160008160ff1611611147576040805162461bcd60e51b815260206004820152601e60248201527f496e76616c69642c206f72206d697373696e67207065726d697373696f6e0000604482015290519081900360640190fd5b6000546001600160a01b031633146111c15760ff8116600090815260076020526040902061117b903363ffffffff611f1116565b6111c1576040805162461bcd60e51b815260206004820152601260248201527126b4b9b9b4b733903832b936b4b9b9b4b7b760711b604482015290519081900360640190fd5b60ff83166000908152600760205260409020610bba908363ffffffff611f4816565b60006111ee82611e22565b6040015192915050565b60076020526000908152604090205481565b60045481565b6000806000806000611220612446565b600088815260066020526040902061124790611242908963ffffffff6120d816565b611e22565b8051602082015160408301516060840151608090940151929c919b50995091975095509350505050565b6001600160a01b0381166000908152600260205260408120548190818113801561129d57506004548113155b6112ee576040805162461bcd60e51b815260206004820152601560248201527f4163636f756e7420646f65736e27742065786973740000000000000000000000604482015290519081900360640190fd5b6112f6612446565b506000818152600360209081526040808320815160a08101835281546001600160a01b03808216835260ff600160a01b83048116968401879052600160a81b90920490911615801594830194909452600183015416606082015260029091015460808201529290611365575081515b6001546001600160a01b038281169116148015611395575060015460608401516001600160a01b03908116911614155b156114485760608301516001600160a01b0316600090815260026020526040902054935060018412806113c9575060045484135b156113d357611448565b600084815260036020908152604091829020825160a08101845281546001600160a01b03808216835260ff600160a01b8304811695840195909552600160a81b90910490931615801594820194909452600182015490921660608301526002015460808201529350611443575081515b611365565b909450925050505b915091565b6001600160a01b038116600090815260026020526040812054818113801561147f57506004548113155b9392505050565b6001546000906001600160a01b03838116911614156114a85750600019610cba565b6001600160a01b03821660009081526002602052604081205460001901908112806114d557506004548112155b156108a457600019915050610cba565b6000546001600160a01b031681565b60006114ff82611e22565b6080015192915050565b60ff8216600090815260076020526040812061147f908363ffffffff611f1116565b600154600160a01b900460ff1681565b60ff8216600090815260076020526040812061147f908363ffffffff61215e16565b600154600160a01b900460ff16156115a65760405162461bcd60e51b815260040180806020018281038252602d8152602001806124c2602d913960400191505060405180910390fd5b6115af81611e22565b6020015160008160ff161161160b576040805162461bcd60e51b815260206004820152601e60248201527f496e76616c69642c206f72206d697373696e67207065726d697373696f6e0000604482015290519081900360640190fd5b6000546001600160a01b031633146116855760ff8116600090815260076020526040902061163f903363ffffffff611f1116565b611685576040805162461bcd60e51b815260206004820152601260248201527126b4b9b9b4b733903832b936b4b9b9b4b7b760711b604482015290519081900360640190fd5b6001600160a01b0382166000908152600560205260408120815b606460ff821610156116e05760ff811660009081526020839052604090205483146116d85760ff81166000908152602083905260408120555b60010161169f565b506001600160a01b03841660009081526002602081815260408084205480855260038352818520909301548085526006909252909220909190611729908763ffffffff611f4816565b506004548212156118175760048054600090815260036020818152604080842080546001600160a01b03908116865260028085528387208a905589875294909352818520815481549085166001600160a01b031991821617808355835460ff60a01b19909116600160a01b9182900460ff90811690920217808455845460ff60a81b19909116600160a81b91829004909216151502178255600180840154818401805491909716908316179095559185015490850155945484528320805475ffffffffffffffffffffffffffffffffffffffffffff1916815590810180549094169093559190910155611859565b6000828152600360205260408120805475ffffffffffffffffffffffffffffffffffffffffffff191681556001810180546001600160a01b0319169055600201555b5050506001600160a01b03909216600090815260026020526040812055505060048054600019019055565b600154600160a01b900460ff16156118cd5760405162461bcd60e51b815260040180806020018281038252602d8152602001806124c2602d913960400191505060405180910390fd5b6118d682611e22565b6020015160008160ff1611611932576040805162461bcd60e51b815260206004820152601e60248201527f496e76616c69642c206f72206d697373696e67207065726d697373696f6e0000604482015290519081900360640190fd5b6000546001600160a01b031633146119ac5760ff81166000908152600760205260409020611966903363ffffffff611f1116565b6119ac576040805162461bcd60e51b815260206004820152601260248201527126b4b9b9b4b733903832b936b4b9b9b4b7b760711b604482015290519081900360640190fd5b506001600160a01b039091166000908152600260209081526040808320548352600390915290208054911515600160a81b0260ff60a81b19909216919091179055565b60066020526000908152604090205481565b60ff8216600090815260076020526040812061147f908363ffffffff6120d816565b600154600160a01b900460ff1615611a6c5760405162461bcd60e51b815260040180806020018281038252602d8152602001806124c2602d913960400191505060405180910390fd5b8360008160ff1611611ac5576040805162461bcd60e51b815260206004820152601e60248201527f496e76616c69642c206f72206d697373696e67207065726d697373696f6e0000604482015290519081900360640190fd5b6000546001600160a01b03163314611b3f5760ff81166000908152600760205260409020611af9903363ffffffff611f1116565b611b3f576040805162461bcd60e51b815260206004820152601260248201527126b4b9b9b4b733903832b936b4b9b9b4b7b760711b604482015290519081900360640190fd5b6001600160a01b0386166000908152600260205260409020546001811280611b68575060045481135b611bb9576040805162461bcd60e51b815260206004820152601660248201527f4163636f756e7420616c72656164792065786973747300000000000000000000604482015290519081900360640190fd5b611bc1612446565b6040518060a00160405280896001600160a01b031681526020018860ff1681526020018715158152602001866001600160a01b03168152602001858152509050600080611c0d836121b5565b915091508115611c25826001600160a01b03166122c8565b60405160200180807f4379636c6963616c206c696e6561676520617420616464726573733a20000000815250601d0182805190602001908083835b60208310611c7f5780518252601f199092019160209182019101611c60565b6001836020036101000a03801982511681845116808217855250505050505090500191505060405160208183030381529060405290611d3c5760405162461bcd60e51b81526004018080602001828103825283818151815260200191508051906020019080838360005b83811015611d01578181015183820152602001611ce9565b50505050905090810190601f168015611d2e5780820380516001836020036101000a031916815260200191505b509250505060405180910390fd5b50600480546001908101918290556001600160a01b03808d16600090815260026020818152604080842087905595835260038152858320895181548b8401518c8a01511515600160a81b0260ff60a81b1960ff909216600160a01b0260ff60a01b19948a166001600160a01b03199485161794909416939093171691909117825560608b015196820180549790961696169590951790935560808801519301929092558882526006905220611df7908b63ffffffff61203316565b5050505050505050505050565b600082815260066020526040812061147f908363ffffffff6120d816565b611e2a612446565b6001600160a01b0382166000908152600260205260408120549081138015611e5457506004548113155b611ea5576040805162461bcd60e51b815260206004820152601560248201527f4163636f756e7420646f65736e27742065786973740000000000000000000000604482015290519081900360640190fd5b600090815260036020908152604091829020825160a08101845281546001600160a01b03808216835260ff600160a01b8304811695840195909552600160a81b909104909316151593810193909352600181015490911660608301526002015460808201529050919050565b6001600160a01b038116600090815260018301602052604081205460001901818112801590611f405750835481125b949350505050565b6001600160a01b03811660009081526001808401602052604082205490811280611f725750835481135b15611f815760009150506108a4565b8354811215611fe857835460009081526002850160208181526040808420546001600160a01b03168085526001890183528185208690558585529290915280832080546001600160a01b031990811690931790558654835290912080549091169055612007565b6000818152600285016020526040902080546001600160a01b03191690555b50506001600160a01b031660009081526001828101602052604082209190915581546000190190915590565b60006001600160a01b03821661204b575060006108a4565b6001600160a01b03821660009081526001840160205260408120546000190190811280159061207a5750835481125b156120895760009150506108a4565b5050815460019081018084556001600160a01b0383166000818152838601602090815260408083208590559382526002870190529190912080546001600160a01b031916909117905592915050565b60008082121580156120ea5750825482125b61213b576040805162461bcd60e51b815260206004820152601860248201527f496e646578206f757473696465206f6620626f756e64732e0000000000000000604482015290519081900360640190fd5b50600101600090815260029190910160205260409020546001600160a01b031690565b60006001600160a01b03821661217757506000196108a4565b6001600160a01b038216600090815260018401602052604081205460001901908112806121a5575083548112155b1561147f576000199150506108a4565b60008060006121c2612446565b50835b60015460608201516001600160a01b039081169116146122af5760608101516001600160a01b0316600090815260026020526040902054600181128061220c575060045481135b1561221757506122af565b6020820151600160ff9091161b831780841415612241575050606001516001935091506114509050565b600091825260036020908152604092839020835160a08101855281546001600160a01b03808216835260ff600160a01b8304811695840195909552600160a81b90910490931615159481019490945260018101549091166060840152600201546080830152925090506121c5565b5050600154600092506001600160a01b03169050915091565b60408051602a8082526060828101909352829190602082018180388339019050509050603060f81b816000815181106122fd57fe5b60200101906001600160f81b031916908160001a905350607860f81b8160018151811061232657fe5b60200101906001600160f81b031916908160001a90535060005b601481101561243f5760008160130360080260020a856001600160a01b03168161236657fe5b0460f81b9050600060108260f81c60ff168161237e57fe5b0460f881811b9250601060ff90921691820284821c03901b90600a116123ad578160f81c60570160f81b6123b8565b8160f81c60300160f81b5b8585600202600201815181106123ca57fe5b60200101906001600160f81b031916908160001a905350600a60f882901c106123fc578060f81c60570160f81b612407565b8060f81c60300160f81b5b85856002026003018151811061241957fe5b60200101906001600160f81b031916908160001a90535050600190920191506123409050565b5092915050565b6040805160a0810182526000808252602082018190529181018290526060810182905260808101919091529056fe4e6577204f776e65722063616e6e6f74206265207468652063757272656e74206f776e6572436f6e747261637420616c726561647920696e20726571756573746564206c6f636b207374617465436f6e74726163742069732063757272656e746c79206c6f636b656420666f72206d6f64696669636174696f6ea265627a7a723058204f9b10cccb11fa0cb06cd2c77d139e150072ab93d236f0edf12b79545cce496264736f6c63430005090032`

// DeployRegistry deploys a new Ethereum contract, binding an instance of Registry to it.
func DeployRegistry(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Registry, error) {
	parsed, err := abi.JSON(strings.NewReader(RegistryABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(RegistryBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Registry{RegistryCaller: RegistryCaller{contract: contract}, RegistryTransactor: RegistryTransactor{contract: contract}, RegistryFilterer: RegistryFilterer{contract: contract}}, nil
}

// Registry is an auto generated Go binding around an Ethereum contract.
type Registry struct {
	RegistryCaller     // Read-only binding to the contract
	RegistryTransactor // Write-only binding to the contract
	RegistryFilterer   // Log filterer for contract events
}

// RegistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type RegistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RegistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RegistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RegistryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RegistryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RegistrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RegistrySession struct {
	Contract     *Registry         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RegistryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RegistryCallerSession struct {
	Contract *RegistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// RegistryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RegistryTransactorSession struct {
	Contract     *RegistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// RegistryRaw is an auto generated low-level Go binding around an Ethereum contract.
type RegistryRaw struct {
	Contract *Registry // Generic contract binding to access the raw methods on
}

// RegistryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RegistryCallerRaw struct {
	Contract *RegistryCaller // Generic read-only contract binding to access the raw methods on
}

// RegistryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RegistryTransactorRaw struct {
	Contract *RegistryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRegistry creates a new instance of Registry, bound to a specific deployed contract.
func NewRegistry(address common.Address, backend bind.ContractBackend) (*Registry, error) {
	contract, err := bindRegistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Registry{RegistryCaller: RegistryCaller{contract: contract}, RegistryTransactor: RegistryTransactor{contract: contract}, RegistryFilterer: RegistryFilterer{contract: contract}}, nil
}

// NewRegistryCaller creates a new read-only instance of Registry, bound to a specific deployed contract.
func NewRegistryCaller(address common.Address, caller bind.ContractCaller) (*RegistryCaller, error) {
	contract, err := bindRegistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RegistryCaller{contract: contract}, nil
}

// NewRegistryTransactor creates a new write-only instance of Registry, bound to a specific deployed contract.
func NewRegistryTransactor(address common.Address, transactor bind.ContractTransactor) (*RegistryTransactor, error) {
	contract, err := bindRegistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RegistryTransactor{contract: contract}, nil
}

// NewRegistryFilterer creates a new log filterer instance of Registry, bound to a specific deployed contract.
func NewRegistryFilterer(address common.Address, filterer bind.ContractFilterer) (*RegistryFilterer, error) {
	contract, err := bindRegistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RegistryFilterer{contract: contract}, nil
}

// bindRegistry binds a generic wrapper to an already deployed contract.
func bindRegistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(RegistryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Registry *RegistryRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Registry.Contract.RegistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Registry *RegistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Registry.Contract.RegistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Registry *RegistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Registry.Contract.RegistryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Registry *RegistryCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Registry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Registry *RegistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Registry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Registry *RegistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Registry.Contract.contract.Transact(opts, method, params...)
}

// ZEROADDRESS is a free data retrieval call binding the contract method 0x538ba4f9.
//
// Solidity: function ZERO_ADDRESS() constant returns(address)
func (_Registry *RegistryCaller) ZEROADDRESS(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Registry.contract.Call(opts, out, "ZERO_ADDRESS")
	return *ret0, err
}

// ZEROADDRESS is a free data retrieval call binding the contract method 0x538ba4f9.
//
// Solidity: function ZERO_ADDRESS() constant returns(address)
func (_Registry *RegistrySession) ZEROADDRESS() (common.Address, error) {
	return _Registry.Contract.ZEROADDRESS(&_Registry.CallOpts)
}

// ZEROADDRESS is a free data retrieval call binding the contract method 0x538ba4f9.
//
// Solidity: function ZERO_ADDRESS() constant returns(address)
func (_Registry *RegistryCallerSession) ZEROADDRESS() (common.Address, error) {
	return _Registry.Contract.ZEROADDRESS(&_Registry.CallOpts)
}

// AccountAt is a free data retrieval call binding the contract method 0x00f17754.
//
// Solidity: function accountAt(index int256) constant returns(address, uint8, bool, address, bytes32)
func (_Registry *RegistryCaller) AccountAt(opts *bind.CallOpts, index *big.Int) (common.Address, uint8, bool, common.Address, [32]byte, error) {
	var (
		ret0 = new(common.Address)
		ret1 = new(uint8)
		ret2 = new(bool)
		ret3 = new(common.Address)
		ret4 = new([32]byte)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
		ret3,
		ret4,
	}
	err := _Registry.contract.Call(opts, out, "accountAt", index)
	return *ret0, *ret1, *ret2, *ret3, *ret4, err
}

// AccountAt is a free data retrieval call binding the contract method 0x00f17754.
//
// Solidity: function accountAt(index int256) constant returns(address, uint8, bool, address, bytes32)
func (_Registry *RegistrySession) AccountAt(index *big.Int) (common.Address, uint8, bool, common.Address, [32]byte, error) {
	return _Registry.Contract.AccountAt(&_Registry.CallOpts, index)
}

// AccountAt is a free data retrieval call binding the contract method 0x00f17754.
//
// Solidity: function accountAt(index int256) constant returns(address, uint8, bool, address, bytes32)
func (_Registry *RegistryCallerSession) AccountAt(index *big.Int) (common.Address, uint8, bool, common.Address, [32]byte, error) {
	return _Registry.Contract.AccountAt(&_Registry.CallOpts, index)
}

// AccountAtHash is a free data retrieval call binding the contract method 0x6a2a1bdc.
//
// Solidity: function accountAtHash(hash bytes32, index int256) constant returns(address, uint8, bool, address, bytes32)
func (_Registry *RegistryCaller) AccountAtHash(opts *bind.CallOpts, hash [32]byte, index *big.Int) (common.Address, uint8, bool, common.Address, [32]byte, error) {
	var (
		ret0 = new(common.Address)
		ret1 = new(uint8)
		ret2 = new(bool)
		ret3 = new(common.Address)
		ret4 = new([32]byte)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
		ret3,
		ret4,
	}
	err := _Registry.contract.Call(opts, out, "accountAtHash", hash, index)
	return *ret0, *ret1, *ret2, *ret3, *ret4, err
}

// AccountAtHash is a free data retrieval call binding the contract method 0x6a2a1bdc.
//
// Solidity: function accountAtHash(hash bytes32, index int256) constant returns(address, uint8, bool, address, bytes32)
func (_Registry *RegistrySession) AccountAtHash(hash [32]byte, index *big.Int) (common.Address, uint8, bool, common.Address, [32]byte, error) {
	return _Registry.Contract.AccountAtHash(&_Registry.CallOpts, hash, index)
}

// AccountAtHash is a free data retrieval call binding the contract method 0x6a2a1bdc.
//
// Solidity: function accountAtHash(hash bytes32, index int256) constant returns(address, uint8, bool, address, bytes32)
func (_Registry *RegistryCallerSession) AccountAtHash(hash [32]byte, index *big.Int) (common.Address, uint8, bool, common.Address, [32]byte, error) {
	return _Registry.Contract.AccountAtHash(&_Registry.CallOpts, hash, index)
}

// AccountExists is a free data retrieval call binding the contract method 0x75cd51ed.
//
// Solidity: function accountExists(addr address) constant returns(bool)
func (_Registry *RegistryCaller) AccountExists(opts *bind.CallOpts, addr common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Registry.contract.Call(opts, out, "accountExists", addr)
	return *ret0, err
}

// AccountExists is a free data retrieval call binding the contract method 0x75cd51ed.
//
// Solidity: function accountExists(addr address) constant returns(bool)
func (_Registry *RegistrySession) AccountExists(addr common.Address) (bool, error) {
	return _Registry.Contract.AccountExists(&_Registry.CallOpts, addr)
}

// AccountExists is a free data retrieval call binding the contract method 0x75cd51ed.
//
// Solidity: function accountExists(addr address) constant returns(bool)
func (_Registry *RegistryCallerSession) AccountExists(addr common.Address) (bool, error) {
	return _Registry.Contract.AccountExists(&_Registry.CallOpts, addr)
}

// AccountFrozen is a free data retrieval call binding the contract method 0x624522f9.
//
// Solidity: function accountFrozen(addr address) constant returns(bool)
func (_Registry *RegistryCaller) AccountFrozen(opts *bind.CallOpts, addr common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Registry.contract.Call(opts, out, "accountFrozen", addr)
	return *ret0, err
}

// AccountFrozen is a free data retrieval call binding the contract method 0x624522f9.
//
// Solidity: function accountFrozen(addr address) constant returns(bool)
func (_Registry *RegistrySession) AccountFrozen(addr common.Address) (bool, error) {
	return _Registry.Contract.AccountFrozen(&_Registry.CallOpts, addr)
}

// AccountFrozen is a free data retrieval call binding the contract method 0x624522f9.
//
// Solidity: function accountFrozen(addr address) constant returns(bool)
func (_Registry *RegistryCallerSession) AccountFrozen(addr common.Address) (bool, error) {
	return _Registry.Contract.AccountFrozen(&_Registry.CallOpts, addr)
}

// AccountGet is a free data retrieval call binding the contract method 0x137e37d9.
//
// Solidity: function accountGet(addr address) constant returns(uint8, bool, address, bytes32)
func (_Registry *RegistryCaller) AccountGet(opts *bind.CallOpts, addr common.Address) (uint8, bool, common.Address, [32]byte, error) {
	var (
		ret0 = new(uint8)
		ret1 = new(bool)
		ret2 = new(common.Address)
		ret3 = new([32]byte)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
		ret3,
	}
	err := _Registry.contract.Call(opts, out, "accountGet", addr)
	return *ret0, *ret1, *ret2, *ret3, err
}

// AccountGet is a free data retrieval call binding the contract method 0x137e37d9.
//
// Solidity: function accountGet(addr address) constant returns(uint8, bool, address, bytes32)
func (_Registry *RegistrySession) AccountGet(addr common.Address) (uint8, bool, common.Address, [32]byte, error) {
	return _Registry.Contract.AccountGet(&_Registry.CallOpts, addr)
}

// AccountGet is a free data retrieval call binding the contract method 0x137e37d9.
//
// Solidity: function accountGet(addr address) constant returns(uint8, bool, address, bytes32)
func (_Registry *RegistryCallerSession) AccountGet(addr common.Address) (uint8, bool, common.Address, [32]byte, error) {
	return _Registry.Contract.AccountGet(&_Registry.CallOpts, addr)
}

// AccountHash is a free data retrieval call binding the contract method 0x960d27d3.
//
// Solidity: function accountHash(addr address) constant returns(bytes32)
func (_Registry *RegistryCaller) AccountHash(opts *bind.CallOpts, addr common.Address) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Registry.contract.Call(opts, out, "accountHash", addr)
	return *ret0, err
}

// AccountHash is a free data retrieval call binding the contract method 0x960d27d3.
//
// Solidity: function accountHash(addr address) constant returns(bytes32)
func (_Registry *RegistrySession) AccountHash(addr common.Address) ([32]byte, error) {
	return _Registry.Contract.AccountHash(&_Registry.CallOpts, addr)
}

// AccountHash is a free data retrieval call binding the contract method 0x960d27d3.
//
// Solidity: function accountHash(addr address) constant returns(bytes32)
func (_Registry *RegistryCallerSession) AccountHash(addr common.Address) ([32]byte, error) {
	return _Registry.Contract.AccountHash(&_Registry.CallOpts, addr)
}

// AccountIndexOf is a free data retrieval call binding the contract method 0x770850c8.
//
// Solidity: function accountIndexOf(addr address) constant returns(int256)
func (_Registry *RegistryCaller) AccountIndexOf(opts *bind.CallOpts, addr common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Registry.contract.Call(opts, out, "accountIndexOf", addr)
	return *ret0, err
}

// AccountIndexOf is a free data retrieval call binding the contract method 0x770850c8.
//
// Solidity: function accountIndexOf(addr address) constant returns(int256)
func (_Registry *RegistrySession) AccountIndexOf(addr common.Address) (*big.Int, error) {
	return _Registry.Contract.AccountIndexOf(&_Registry.CallOpts, addr)
}

// AccountIndexOf is a free data retrieval call binding the contract method 0x770850c8.
//
// Solidity: function accountIndexOf(addr address) constant returns(int256)
func (_Registry *RegistryCallerSession) AccountIndexOf(addr common.Address) (*big.Int, error) {
	return _Registry.Contract.AccountIndexOf(&_Registry.CallOpts, addr)
}

// AccountKind is a free data retrieval call binding the contract method 0x351a97f8.
//
// Solidity: function accountKind(addr address) constant returns(uint8)
func (_Registry *RegistryCaller) AccountKind(opts *bind.CallOpts, addr common.Address) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _Registry.contract.Call(opts, out, "accountKind", addr)
	return *ret0, err
}

// AccountKind is a free data retrieval call binding the contract method 0x351a97f8.
//
// Solidity: function accountKind(addr address) constant returns(uint8)
func (_Registry *RegistrySession) AccountKind(addr common.Address) (uint8, error) {
	return _Registry.Contract.AccountKind(&_Registry.CallOpts, addr)
}

// AccountKind is a free data retrieval call binding the contract method 0x351a97f8.
//
// Solidity: function accountKind(addr address) constant returns(uint8)
func (_Registry *RegistryCallerSession) AccountKind(addr common.Address) (uint8, error) {
	return _Registry.Contract.AccountKind(&_Registry.CallOpts, addr)
}

// AccountKindAndFrozenAddress is a free data retrieval call binding the contract method 0x6f6072eb.
//
// Solidity: function accountKindAndFrozenAddress(addr address) constant returns(uint8, address)
func (_Registry *RegistryCaller) AccountKindAndFrozenAddress(opts *bind.CallOpts, addr common.Address) (uint8, common.Address, error) {
	var (
		ret0 = new(uint8)
		ret1 = new(common.Address)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _Registry.contract.Call(opts, out, "accountKindAndFrozenAddress", addr)
	return *ret0, *ret1, err
}

// AccountKindAndFrozenAddress is a free data retrieval call binding the contract method 0x6f6072eb.
//
// Solidity: function accountKindAndFrozenAddress(addr address) constant returns(uint8, address)
func (_Registry *RegistrySession) AccountKindAndFrozenAddress(addr common.Address) (uint8, common.Address, error) {
	return _Registry.Contract.AccountKindAndFrozenAddress(&_Registry.CallOpts, addr)
}

// AccountKindAndFrozenAddress is a free data retrieval call binding the contract method 0x6f6072eb.
//
// Solidity: function accountKindAndFrozenAddress(addr address) constant returns(uint8, address)
func (_Registry *RegistryCallerSession) AccountKindAndFrozenAddress(addr common.Address) (uint8, common.Address, error) {
	return _Registry.Contract.AccountKindAndFrozenAddress(&_Registry.CallOpts, addr)
}

// AccountParent is a free data retrieval call binding the contract method 0x31aaa74a.
//
// Solidity: function accountParent(addr address) constant returns(address)
func (_Registry *RegistryCaller) AccountParent(opts *bind.CallOpts, addr common.Address) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Registry.contract.Call(opts, out, "accountParent", addr)
	return *ret0, err
}

// AccountParent is a free data retrieval call binding the contract method 0x31aaa74a.
//
// Solidity: function accountParent(addr address) constant returns(address)
func (_Registry *RegistrySession) AccountParent(addr common.Address) (common.Address, error) {
	return _Registry.Contract.AccountParent(&_Registry.CallOpts, addr)
}

// AccountParent is a free data retrieval call binding the contract method 0x31aaa74a.
//
// Solidity: function accountParent(addr address) constant returns(address)
func (_Registry *RegistryCallerSession) AccountParent(addr common.Address) (common.Address, error) {
	return _Registry.Contract.AccountParent(&_Registry.CallOpts, addr)
}

// Accounts is a free data retrieval call binding the contract method 0x68cd03f6.
//
// Solidity: function accounts() constant returns(int256)
func (_Registry *RegistryCaller) Accounts(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Registry.contract.Call(opts, out, "accounts")
	return *ret0, err
}

// Accounts is a free data retrieval call binding the contract method 0x68cd03f6.
//
// Solidity: function accounts() constant returns(int256)
func (_Registry *RegistrySession) Accounts() (*big.Int, error) {
	return _Registry.Contract.Accounts(&_Registry.CallOpts)
}

// Accounts is a free data retrieval call binding the contract method 0x68cd03f6.
//
// Solidity: function accounts() constant returns(int256)
func (_Registry *RegistryCallerSession) Accounts() (*big.Int, error) {
	return _Registry.Contract.Accounts(&_Registry.CallOpts)
}

// AddressAtHash is a free data retrieval call binding the contract method 0xfcef15e6.
//
// Solidity: function addressAtHash(hash bytes32, index int256) constant returns(address)
func (_Registry *RegistryCaller) AddressAtHash(opts *bind.CallOpts, hash [32]byte, index *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Registry.contract.Call(opts, out, "addressAtHash", hash, index)
	return *ret0, err
}

// AddressAtHash is a free data retrieval call binding the contract method 0xfcef15e6.
//
// Solidity: function addressAtHash(hash bytes32, index int256) constant returns(address)
func (_Registry *RegistrySession) AddressAtHash(hash [32]byte, index *big.Int) (common.Address, error) {
	return _Registry.Contract.AddressAtHash(&_Registry.CallOpts, hash, index)
}

// AddressAtHash is a free data retrieval call binding the contract method 0xfcef15e6.
//
// Solidity: function addressAtHash(hash bytes32, index int256) constant returns(address)
func (_Registry *RegistryCallerSession) AddressAtHash(hash [32]byte, index *big.Int) (common.Address, error) {
	return _Registry.Contract.AddressAtHash(&_Registry.CallOpts, hash, index)
}

// Data is a free data retrieval call binding the contract method 0x572bc6b1.
//
// Solidity: function data( address,  uint8) constant returns(bytes32)
func (_Registry *RegistryCaller) Data(opts *bind.CallOpts, arg0 common.Address, arg1 uint8) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Registry.contract.Call(opts, out, "data", arg0, arg1)
	return *ret0, err
}

// Data is a free data retrieval call binding the contract method 0x572bc6b1.
//
// Solidity: function data( address,  uint8) constant returns(bytes32)
func (_Registry *RegistrySession) Data(arg0 common.Address, arg1 uint8) ([32]byte, error) {
	return _Registry.Contract.Data(&_Registry.CallOpts, arg0, arg1)
}

// Data is a free data retrieval call binding the contract method 0x572bc6b1.
//
// Solidity: function data( address,  uint8) constant returns(bytes32)
func (_Registry *RegistryCallerSession) Data(arg0 common.Address, arg1 uint8) ([32]byte, error) {
	return _Registry.Contract.Data(&_Registry.CallOpts, arg0, arg1)
}

// Hashes is a free data retrieval call binding the contract method 0xd658d2e9.
//
// Solidity: function hashes( bytes32) constant returns(count int256)
func (_Registry *RegistryCaller) Hashes(opts *bind.CallOpts, arg0 [32]byte) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Registry.contract.Call(opts, out, "hashes", arg0)
	return *ret0, err
}

// Hashes is a free data retrieval call binding the contract method 0xd658d2e9.
//
// Solidity: function hashes( bytes32) constant returns(count int256)
func (_Registry *RegistrySession) Hashes(arg0 [32]byte) (*big.Int, error) {
	return _Registry.Contract.Hashes(&_Registry.CallOpts, arg0)
}

// Hashes is a free data retrieval call binding the contract method 0xd658d2e9.
//
// Solidity: function hashes( bytes32) constant returns(count int256)
func (_Registry *RegistryCallerSession) Hashes(arg0 [32]byte) (*big.Int, error) {
	return _Registry.Contract.Hashes(&_Registry.CallOpts, arg0)
}

// IsLocked is a free data retrieval call binding the contract method 0xa4e2d634.
//
// Solidity: function isLocked() constant returns(bool)
func (_Registry *RegistryCaller) IsLocked(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Registry.contract.Call(opts, out, "isLocked")
	return *ret0, err
}

// IsLocked is a free data retrieval call binding the contract method 0xa4e2d634.
//
// Solidity: function isLocked() constant returns(bool)
func (_Registry *RegistrySession) IsLocked() (bool, error) {
	return _Registry.Contract.IsLocked(&_Registry.CallOpts)
}

// IsLocked is a free data retrieval call binding the contract method 0xa4e2d634.
//
// Solidity: function isLocked() constant returns(bool)
func (_Registry *RegistryCallerSession) IsLocked() (bool, error) {
	return _Registry.Contract.IsLocked(&_Registry.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Registry *RegistryCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Registry.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Registry *RegistrySession) Owner() (common.Address, error) {
	return _Registry.Contract.Owner(&_Registry.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Registry *RegistryCallerSession) Owner() (common.Address, error) {
	return _Registry.Contract.Owner(&_Registry.CallOpts)
}

// PermissionAt is a free data retrieval call binding the contract method 0xd7379999.
//
// Solidity: function permissionAt(kind uint8, index int256) constant returns(address)
func (_Registry *RegistryCaller) PermissionAt(opts *bind.CallOpts, kind uint8, index *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Registry.contract.Call(opts, out, "permissionAt", kind, index)
	return *ret0, err
}

// PermissionAt is a free data retrieval call binding the contract method 0xd7379999.
//
// Solidity: function permissionAt(kind uint8, index int256) constant returns(address)
func (_Registry *RegistrySession) PermissionAt(kind uint8, index *big.Int) (common.Address, error) {
	return _Registry.Contract.PermissionAt(&_Registry.CallOpts, kind, index)
}

// PermissionAt is a free data retrieval call binding the contract method 0xd7379999.
//
// Solidity: function permissionAt(kind uint8, index int256) constant returns(address)
func (_Registry *RegistryCallerSession) PermissionAt(kind uint8, index *big.Int) (common.Address, error) {
	return _Registry.Contract.PermissionAt(&_Registry.CallOpts, kind, index)
}

// PermissionExists is a free data retrieval call binding the contract method 0x9d44ac4f.
//
// Solidity: function permissionExists(kind uint8, addr address) constant returns(bool)
func (_Registry *RegistryCaller) PermissionExists(opts *bind.CallOpts, kind uint8, addr common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Registry.contract.Call(opts, out, "permissionExists", kind, addr)
	return *ret0, err
}

// PermissionExists is a free data retrieval call binding the contract method 0x9d44ac4f.
//
// Solidity: function permissionExists(kind uint8, addr address) constant returns(bool)
func (_Registry *RegistrySession) PermissionExists(kind uint8, addr common.Address) (bool, error) {
	return _Registry.Contract.PermissionExists(&_Registry.CallOpts, kind, addr)
}

// PermissionExists is a free data retrieval call binding the contract method 0x9d44ac4f.
//
// Solidity: function permissionExists(kind uint8, addr address) constant returns(bool)
func (_Registry *RegistryCallerSession) PermissionExists(kind uint8, addr common.Address) (bool, error) {
	return _Registry.Contract.PermissionExists(&_Registry.CallOpts, kind, addr)
}

// PermissionIndexOf is a free data retrieval call binding the contract method 0xa84694c8.
//
// Solidity: function permissionIndexOf(kind uint8, addr address) constant returns(int256)
func (_Registry *RegistryCaller) PermissionIndexOf(opts *bind.CallOpts, kind uint8, addr common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Registry.contract.Call(opts, out, "permissionIndexOf", kind, addr)
	return *ret0, err
}

// PermissionIndexOf is a free data retrieval call binding the contract method 0xa84694c8.
//
// Solidity: function permissionIndexOf(kind uint8, addr address) constant returns(int256)
func (_Registry *RegistrySession) PermissionIndexOf(kind uint8, addr common.Address) (*big.Int, error) {
	return _Registry.Contract.PermissionIndexOf(&_Registry.CallOpts, kind, addr)
}

// PermissionIndexOf is a free data retrieval call binding the contract method 0xa84694c8.
//
// Solidity: function permissionIndexOf(kind uint8, addr address) constant returns(int256)
func (_Registry *RegistryCallerSession) PermissionIndexOf(kind uint8, addr common.Address) (*big.Int, error) {
	return _Registry.Contract.PermissionIndexOf(&_Registry.CallOpts, kind, addr)
}

// Permissions is a free data retrieval call binding the contract method 0x6476b837.
//
// Solidity: function permissions( uint8) constant returns(count int256)
func (_Registry *RegistryCaller) Permissions(opts *bind.CallOpts, arg0 uint8) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Registry.contract.Call(opts, out, "permissions", arg0)
	return *ret0, err
}

// Permissions is a free data retrieval call binding the contract method 0x6476b837.
//
// Solidity: function permissions( uint8) constant returns(count int256)
func (_Registry *RegistrySession) Permissions(arg0 uint8) (*big.Int, error) {
	return _Registry.Contract.Permissions(&_Registry.CallOpts, arg0)
}

// Permissions is a free data retrieval call binding the contract method 0x6476b837.
//
// Solidity: function permissions( uint8) constant returns(count int256)
func (_Registry *RegistryCallerSession) Permissions(arg0 uint8) (*big.Int, error) {
	return _Registry.Contract.Permissions(&_Registry.CallOpts, arg0)
}

// AddAccount is a paid mutator transaction binding the contract method 0xe8d74d1f.
//
// Solidity: function addAccount(addr address, kind uint8, isFrozen bool, parent address, hash bytes32) returns()
func (_Registry *RegistryTransactor) AddAccount(opts *bind.TransactOpts, addr common.Address, kind uint8, isFrozen bool, parent common.Address, hash [32]byte) (*types.Transaction, error) {
	return _Registry.contract.Transact(opts, "addAccount", addr, kind, isFrozen, parent, hash)
}

// AddAccount is a paid mutator transaction binding the contract method 0xe8d74d1f.
//
// Solidity: function addAccount(addr address, kind uint8, isFrozen bool, parent address, hash bytes32) returns()
func (_Registry *RegistrySession) AddAccount(addr common.Address, kind uint8, isFrozen bool, parent common.Address, hash [32]byte) (*types.Transaction, error) {
	return _Registry.Contract.AddAccount(&_Registry.TransactOpts, addr, kind, isFrozen, parent, hash)
}

// AddAccount is a paid mutator transaction binding the contract method 0xe8d74d1f.
//
// Solidity: function addAccount(addr address, kind uint8, isFrozen bool, parent address, hash bytes32) returns()
func (_Registry *RegistryTransactorSession) AddAccount(addr common.Address, kind uint8, isFrozen bool, parent common.Address, hash [32]byte) (*types.Transaction, error) {
	return _Registry.Contract.AddAccount(&_Registry.TransactOpts, addr, kind, isFrozen, parent, hash)
}

// GrantPermission is a paid mutator transaction binding the contract method 0x1a3722b2.
//
// Solidity: function grantPermission(kind uint8, addr address) returns()
func (_Registry *RegistryTransactor) GrantPermission(opts *bind.TransactOpts, kind uint8, addr common.Address) (*types.Transaction, error) {
	return _Registry.contract.Transact(opts, "grantPermission", kind, addr)
}

// GrantPermission is a paid mutator transaction binding the contract method 0x1a3722b2.
//
// Solidity: function grantPermission(kind uint8, addr address) returns()
func (_Registry *RegistrySession) GrantPermission(kind uint8, addr common.Address) (*types.Transaction, error) {
	return _Registry.Contract.GrantPermission(&_Registry.TransactOpts, kind, addr)
}

// GrantPermission is a paid mutator transaction binding the contract method 0x1a3722b2.
//
// Solidity: function grantPermission(kind uint8, addr address) returns()
func (_Registry *RegistryTransactorSession) GrantPermission(kind uint8, addr common.Address) (*types.Transaction, error) {
	return _Registry.Contract.GrantPermission(&_Registry.TransactOpts, kind, addr)
}

// Kill is a paid mutator transaction binding the contract method 0x41c0e1b5.
//
// Solidity: function kill() returns()
func (_Registry *RegistryTransactor) Kill(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Registry.contract.Transact(opts, "kill")
}

// Kill is a paid mutator transaction binding the contract method 0x41c0e1b5.
//
// Solidity: function kill() returns()
func (_Registry *RegistrySession) Kill() (*types.Transaction, error) {
	return _Registry.Contract.Kill(&_Registry.TransactOpts)
}

// Kill is a paid mutator transaction binding the contract method 0x41c0e1b5.
//
// Solidity: function kill() returns()
func (_Registry *RegistryTransactorSession) Kill() (*types.Transaction, error) {
	return _Registry.Contract.Kill(&_Registry.TransactOpts)
}

// RemoveAccount is a paid mutator transaction binding the contract method 0xc4740a95.
//
// Solidity: function removeAccount(addr address) returns()
func (_Registry *RegistryTransactor) RemoveAccount(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	return _Registry.contract.Transact(opts, "removeAccount", addr)
}

// RemoveAccount is a paid mutator transaction binding the contract method 0xc4740a95.
//
// Solidity: function removeAccount(addr address) returns()
func (_Registry *RegistrySession) RemoveAccount(addr common.Address) (*types.Transaction, error) {
	return _Registry.Contract.RemoveAccount(&_Registry.TransactOpts, addr)
}

// RemoveAccount is a paid mutator transaction binding the contract method 0xc4740a95.
//
// Solidity: function removeAccount(addr address) returns()
func (_Registry *RegistryTransactorSession) RemoveAccount(addr common.Address) (*types.Transaction, error) {
	return _Registry.Contract.RemoveAccount(&_Registry.TransactOpts, addr)
}

// RevokePermission is a paid mutator transaction binding the contract method 0x612eb5e7.
//
// Solidity: function revokePermission(kind uint8, addr address) returns()
func (_Registry *RegistryTransactor) RevokePermission(opts *bind.TransactOpts, kind uint8, addr common.Address) (*types.Transaction, error) {
	return _Registry.contract.Transact(opts, "revokePermission", kind, addr)
}

// RevokePermission is a paid mutator transaction binding the contract method 0x612eb5e7.
//
// Solidity: function revokePermission(kind uint8, addr address) returns()
func (_Registry *RegistrySession) RevokePermission(kind uint8, addr common.Address) (*types.Transaction, error) {
	return _Registry.Contract.RevokePermission(&_Registry.TransactOpts, kind, addr)
}

// RevokePermission is a paid mutator transaction binding the contract method 0x612eb5e7.
//
// Solidity: function revokePermission(kind uint8, addr address) returns()
func (_Registry *RegistryTransactorSession) RevokePermission(kind uint8, addr common.Address) (*types.Transaction, error) {
	return _Registry.Contract.RevokePermission(&_Registry.TransactOpts, kind, addr)
}

// SetAccountData is a paid mutator transaction binding the contract method 0x4f1c3b66.
//
// Solidity: function setAccountData(addr address, index uint8, customData bytes32) returns()
func (_Registry *RegistryTransactor) SetAccountData(opts *bind.TransactOpts, addr common.Address, index uint8, customData [32]byte) (*types.Transaction, error) {
	return _Registry.contract.Transact(opts, "setAccountData", addr, index, customData)
}

// SetAccountData is a paid mutator transaction binding the contract method 0x4f1c3b66.
//
// Solidity: function setAccountData(addr address, index uint8, customData bytes32) returns()
func (_Registry *RegistrySession) SetAccountData(addr common.Address, index uint8, customData [32]byte) (*types.Transaction, error) {
	return _Registry.Contract.SetAccountData(&_Registry.TransactOpts, addr, index, customData)
}

// SetAccountData is a paid mutator transaction binding the contract method 0x4f1c3b66.
//
// Solidity: function setAccountData(addr address, index uint8, customData bytes32) returns()
func (_Registry *RegistryTransactorSession) SetAccountData(addr common.Address, index uint8, customData [32]byte) (*types.Transaction, error) {
	return _Registry.Contract.SetAccountData(&_Registry.TransactOpts, addr, index, customData)
}

// SetAccountFrozen is a paid mutator transaction binding the contract method 0xcbe5404f.
//
// Solidity: function setAccountFrozen(addr address, frozen bool) returns()
func (_Registry *RegistryTransactor) SetAccountFrozen(opts *bind.TransactOpts, addr common.Address, frozen bool) (*types.Transaction, error) {
	return _Registry.contract.Transact(opts, "setAccountFrozen", addr, frozen)
}

// SetAccountFrozen is a paid mutator transaction binding the contract method 0xcbe5404f.
//
// Solidity: function setAccountFrozen(addr address, frozen bool) returns()
func (_Registry *RegistrySession) SetAccountFrozen(addr common.Address, frozen bool) (*types.Transaction, error) {
	return _Registry.Contract.SetAccountFrozen(&_Registry.TransactOpts, addr, frozen)
}

// SetAccountFrozen is a paid mutator transaction binding the contract method 0xcbe5404f.
//
// Solidity: function setAccountFrozen(addr address, frozen bool) returns()
func (_Registry *RegistryTransactorSession) SetAccountFrozen(addr common.Address, frozen bool) (*types.Transaction, error) {
	return _Registry.Contract.SetAccountFrozen(&_Registry.TransactOpts, addr, frozen)
}

// SetAccountHash is a paid mutator transaction binding the contract method 0x1706acfc.
//
// Solidity: function setAccountHash(addr address, hash bytes32) returns()
func (_Registry *RegistryTransactor) SetAccountHash(opts *bind.TransactOpts, addr common.Address, hash [32]byte) (*types.Transaction, error) {
	return _Registry.contract.Transact(opts, "setAccountHash", addr, hash)
}

// SetAccountHash is a paid mutator transaction binding the contract method 0x1706acfc.
//
// Solidity: function setAccountHash(addr address, hash bytes32) returns()
func (_Registry *RegistrySession) SetAccountHash(addr common.Address, hash [32]byte) (*types.Transaction, error) {
	return _Registry.Contract.SetAccountHash(&_Registry.TransactOpts, addr, hash)
}

// SetAccountHash is a paid mutator transaction binding the contract method 0x1706acfc.
//
// Solidity: function setAccountHash(addr address, hash bytes32) returns()
func (_Registry *RegistryTransactorSession) SetAccountHash(addr common.Address, hash [32]byte) (*types.Transaction, error) {
	return _Registry.Contract.SetAccountHash(&_Registry.TransactOpts, addr, hash)
}

// SetLocked is a paid mutator transaction binding the contract method 0x211e28b6.
//
// Solidity: function setLocked(locked bool) returns()
func (_Registry *RegistryTransactor) SetLocked(opts *bind.TransactOpts, locked bool) (*types.Transaction, error) {
	return _Registry.contract.Transact(opts, "setLocked", locked)
}

// SetLocked is a paid mutator transaction binding the contract method 0x211e28b6.
//
// Solidity: function setLocked(locked bool) returns()
func (_Registry *RegistrySession) SetLocked(locked bool) (*types.Transaction, error) {
	return _Registry.Contract.SetLocked(&_Registry.TransactOpts, locked)
}

// SetLocked is a paid mutator transaction binding the contract method 0x211e28b6.
//
// Solidity: function setLocked(locked bool) returns()
func (_Registry *RegistryTransactorSession) SetLocked(locked bool) (*types.Transaction, error) {
	return _Registry.Contract.SetLocked(&_Registry.TransactOpts, locked)
}

// TransferOwner is a paid mutator transaction binding the contract method 0x4fb2e45d.
//
// Solidity: function transferOwner(newOwner address) returns()
func (_Registry *RegistryTransactor) TransferOwner(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Registry.contract.Transact(opts, "transferOwner", newOwner)
}

// TransferOwner is a paid mutator transaction binding the contract method 0x4fb2e45d.
//
// Solidity: function transferOwner(newOwner address) returns()
func (_Registry *RegistrySession) TransferOwner(newOwner common.Address) (*types.Transaction, error) {
	return _Registry.Contract.TransferOwner(&_Registry.TransactOpts, newOwner)
}

// TransferOwner is a paid mutator transaction binding the contract method 0x4fb2e45d.
//
// Solidity: function transferOwner(newOwner address) returns()
func (_Registry *RegistryTransactorSession) TransferOwner(newOwner common.Address) (*types.Transaction, error) {
	return _Registry.Contract.TransferOwner(&_Registry.TransactOpts, newOwner)
}

// RegistryOwnerTransferredIterator is returned from FilterOwnerTransferred and is used to iterate over the raw logs and unpacked data for OwnerTransferred events raised by the Registry contract.
type RegistryOwnerTransferredIterator struct {
	Event *RegistryOwnerTransferred // Event containing the contract specifics and raw log

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
func (it *RegistryOwnerTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RegistryOwnerTransferred)
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
		it.Event = new(RegistryOwnerTransferred)
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
func (it *RegistryOwnerTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RegistryOwnerTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RegistryOwnerTransferred represents a OwnerTransferred event raised by the Registry contract.
type RegistryOwnerTransferred struct {
	OldOwner common.Address
	NewOwner common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOwnerTransferred is a free log retrieval operation binding the contract event 0x8934ce4adea8d9ce0d714d2c22b86790e41b7731c84b926fbbdc1d40ff6533c9.
//
// Solidity: e OwnerTransferred(oldOwner indexed address, newOwner indexed address)
func (_Registry *RegistryFilterer) FilterOwnerTransferred(opts *bind.FilterOpts, oldOwner []common.Address, newOwner []common.Address) (*RegistryOwnerTransferredIterator, error) {

	var oldOwnerRule []interface{}
	for _, oldOwnerItem := range oldOwner {
		oldOwnerRule = append(oldOwnerRule, oldOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Registry.contract.FilterLogs(opts, "OwnerTransferred", oldOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &RegistryOwnerTransferredIterator{contract: _Registry.contract, event: "OwnerTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnerTransferred is a free log subscription operation binding the contract event 0x8934ce4adea8d9ce0d714d2c22b86790e41b7731c84b926fbbdc1d40ff6533c9.
//
// Solidity: e OwnerTransferred(oldOwner indexed address, newOwner indexed address)
func (_Registry *RegistryFilterer) WatchOwnerTransferred(opts *bind.WatchOpts, sink chan<- *RegistryOwnerTransferred, oldOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var oldOwnerRule []interface{}
	for _, oldOwnerItem := range oldOwner {
		oldOwnerRule = append(oldOwnerRule, oldOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Registry.contract.WatchLogs(opts, "OwnerTransferred", oldOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RegistryOwnerTransferred)
				if err := _Registry.contract.UnpackLog(event, "OwnerTransferred", log); err != nil {
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
