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

// StorageABI is the input ABI used to generate the binding from.
const StorageABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"index\",\"type\":\"int256\"}],\"name\":\"accountAt\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"uint8\"},{\"name\":\"\",\"type\":\"bool\"},{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"},{\"name\":\"kind\",\"type\":\"uint8\"}],\"name\":\"accountExists\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"},{\"name\":\"kind\",\"type\":\"uint8\"},{\"name\":\"isFrozen\",\"type\":\"bool\"},{\"name\":\"parent\",\"type\":\"address\"}],\"name\":\"addAccount\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"accountGet\",\"outputs\":[{\"name\":\"\",\"type\":\"uint8\"},{\"name\":\"\",\"type\":\"bool\"},{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"kind\",\"type\":\"uint8\"},{\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"grantPermission\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"locked\",\"type\":\"bool\"}],\"name\":\"setLocked\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"accountParent\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"accountKind\",\"outputs\":[{\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"kill\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"},{\"name\":\"index\",\"type\":\"uint8\"},{\"name\":\"customData\",\"type\":\"bytes32\"}],\"name\":\"setAccountData\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwner\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"uint8\"}],\"name\":\"data\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"kind\",\"type\":\"uint8\"},{\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"revokePermission\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"accountFrozen\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint8\"}],\"name\":\"permissions\",\"outputs\":[{\"name\":\"count\",\"type\":\"int256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"accounts\",\"outputs\":[{\"name\":\"count\",\"type\":\"int256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"accountExists\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"accountIndexOf\",\"outputs\":[{\"name\":\"\",\"type\":\"int256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"kind\",\"type\":\"uint8\"},{\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"permissionExists\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isLocked\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"kind\",\"type\":\"uint8\"},{\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"permissionIndexOf\",\"outputs\":[{\"name\":\"\",\"type\":\"int256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"removeAccount\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"},{\"name\":\"frozen\",\"type\":\"bool\"}],\"name\":\"setAccountFrozen\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"kind\",\"type\":\"uint8\"},{\"name\":\"index\",\"type\":\"int256\"}],\"name\":\"permissionAt\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"oldOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnerTransferred\",\"type\":\"event\"}]"

// StorageBin is the compiled bytecode used for deploying new contracts.
const StorageBin = `60806040526000805460a060020a60ff0219600160a060020a03199091163317169055611d9e806100316000396000f3fe608060405234801561001057600080fd5b50600436106101bf576000357c010000000000000000000000000000000000000000000000000000000090048063612eb5e71161010b5780638da5cb5b116100b4578063a84694c81161008e578063a84694c814610564578063c4740a9514610593578063cbe5404f146105b9578063d7379999146105e7576101bf565b80638da5cb5b146105255780639d44ac4f1461052d578063a4e2d6341461055c576101bf565b806368cd03f6116100e557806368cd03f6146104d157806375cd51ed146104d9578063770850c8146104ff576101bf565b8063612eb5e71461045c578063624522f91461048b5780636476b837146104b1576101bf565b806331aaa74a1161016d5780634f1c3b66116101475780634f1c3b66146103c05780634fb2e45d146103f5578063572bc6b11461041b576101bf565b806331aaa74a1461033a578063351a97f81461037c57806341c0e1b5146103b8576101bf565b8063137e37d91161019e578063137e37d91461029b5780631a3722b2146102ec578063211e28b61461031b576101bf565b8062f17754146101c45780630a22ee7314610217578063116c92b71461025a575b600080fd5b6101e1600480360360208110156101da57600080fd5b503561060d565b60408051600160a060020a03958616815260ff909416602085015291151583830152909216606082015290519081900360800190f35b6102466004803603604081101561022d57600080fd5b508035600160a060020a0316906020013560ff1661064e565b604080519115158252519081900360200190f35b6102996004803603608081101561027057600080fd5b50600160a060020a03813581169160ff602082013516916040820135151591606001351661069e565b005b6102c1600480360360208110156102b157600080fd5b5035600160a060020a0316610866565b6040805160ff90941684529115156020840152600160a060020a031682820152519081900360600190f35b6102996004803603604081101561030257600080fd5b50803560ff169060200135600160a060020a03166108a1565b6102996004803603602081101561033157600080fd5b50351515610a1b565b6103606004803603602081101561035057600080fd5b5035600160a060020a0316610b2b565b60408051600160a060020a039092168252519081900360200190f35b6103a26004803603602081101561039257600080fd5b5035600160a060020a0316610b48565b6040805160ff9092168252519081900360200190f35b610299610b65565b610299600480360360608110156103d657600080fd5b50600160a060020a038135169060ff6020820135169060400135610bda565b6102996004803603602081101561040b57600080fd5b5035600160a060020a0316610dcb565b61044a6004803603604081101561043157600080fd5b508035600160a060020a0316906020013560ff16610f49565b60408051918252519081900360200190f35b6102996004803603604081101561047257600080fd5b50803560ff169060200135600160a060020a0316610f66565b610246600480360360208110156104a157600080fd5b5035600160a060020a03166110da565b61044a600480360360208110156104c757600080fd5b503560ff166110f7565b61044a611109565b610246600480360360208110156104ef57600080fd5b5035600160a060020a031661110f565b61044a6004803603602081101561051557600080fd5b5035600160a060020a0316611122565b610360611135565b6102466004803603604081101561054357600080fd5b50803560ff169060200135600160a060020a0316611144565b61024661116d565b61044a6004803603604081101561057a57600080fd5b50803560ff169060200135600160a060020a031661118e565b610299600480360360208110156105a957600080fd5b5035600160a060020a03166111b0565b610299600480360360408110156105cf57600080fd5b50600160a060020a0381351690602001351515611383565b610360600480360360408110156105fd57600080fd5b5060ff8135169060200135611551565b60008060008061061b611cb1565b61062c60018763ffffffff61157316565b8051602082015160408301516060909301519199909850919650945092505050565b60008061066260018563ffffffff61167016565b90506000811215610677576000915050610698565b60ff831661068c60018363ffffffff61157316565b6020015160ff16149150505b92915050565b60005474010000000000000000000000000000000000000000900460ff161561070057604051600080516020611cfe833981519152815260040180806020018281038252602d815260200180611d46602d913960400191505060405180910390fd5b82600060ff8216116107615760408051600080516020611cfe833981519152815260206004820152601e60248201527f496e76616c69642c206f72206d697373696e67207065726d697373696f6e0000604482015290519081900360640190fd5b600054600160a060020a031633146107f05760ff81166000908152600560205260409020610795903363ffffffff6116c916565b15156107f05760408051600080516020611cfe833981519152815260206004820152601260248201527f4d697373696e67207065726d697373696f6e0000000000000000000000000000604482015290519081900360640190fd5b61080460018686868663ffffffff61170016565b151561085f5760408051600080516020611cfe833981519152815260206004820152601660248201527f4163636f756e7420616c72656164792065786973747300000000000000000000604482015290519081900360640190fd5b5050505050565b6000806000610873611cb1565b61088460018663ffffffff61187016565b602081015160408201516060909201519097919650945092505050565b60005474010000000000000000000000000000000000000000900460ff161561090357604051600080516020611cfe833981519152815260040180806020018281038252602d815260200180611d46602d913960400191505060405180910390fd5b81600060ff8216116109645760408051600080516020611cfe833981519152815260206004820152601e60248201527f496e76616c69642c206f72206d697373696e67207065726d697373696f6e0000604482015290519081900360640190fd5b600054600160a060020a031633146109f35760ff81166000908152600560205260409020610998903363ffffffff6116c916565b15156109f35760408051600080516020611cfe833981519152815260206004820152601260248201527f4d697373696e67207065726d697373696f6e0000000000000000000000000000604482015290519081900360640190fd5b60ff83166000908152600560205260409020610a15908363ffffffff61188b16565b50505050565b600054600160a060020a03163314610a825760408051600080516020611cfe833981519152815260206004820152601960248201527f4f776e6572206163636f756e7420697320726571756972656400000000000000604482015290519081900360640190fd5b60005460ff740100000000000000000000000000000000000000009091041615158115151415610aeb57604051600080516020611cfe8339815191528152600401808060200182810382526028815260200180611d1e6028913960400191505060405180910390fd5b60008054911515740100000000000000000000000000000000000000000274ff000000000000000000000000000000000000000019909216919091179055565b6000610b3e60018363ffffffff61187016565b6060015192915050565b6000610b5b60018363ffffffff61187016565b6020015192915050565b600054600160a060020a03163314610bcc5760408051600080516020611cfe833981519152815260206004820152601960248201527f4f776e6572206163636f756e7420697320726571756972656400000000000000604482015290519081900360640190fd5b600054600160a060020a0316ff5b60005474010000000000000000000000000000000000000000900460ff1615610c3c57604051600080516020611cfe833981519152815260040180806020018281038252602d815260200180611d46602d913960400191505060405180910390fd5b610c4d60018463ffffffff61187016565b60200151600060ff821611610cb15760408051600080516020611cfe833981519152815260206004820152601e60248201527f496e76616c69642c206f72206d697373696e67207065726d697373696f6e0000604482015290519081900360640190fd5b600054600160a060020a03163314610d405760ff81166000908152600560205260409020610ce5903363ffffffff6116c916565b1515610d405760408051600080516020611cfe833981519152815260206004820152601260248201527f4d697373696e67207065726d697373696f6e0000000000000000000000000000604482015290519081900360640190fd5b601e60ff841610610da05760408051600080516020611cfe833981519152815260206004820152601760248201527f696e646578206f757473696465206f6620626f756e6473000000000000000000604482015290519081900360640190fd5b50600160a060020a03909216600090815260046020908152604080832060ff90941683529290522055565b600054600160a060020a03163314610e325760408051600080516020611cfe833981519152815260206004820152601960248201527f4f776e6572206163636f756e7420697320726571756972656400000000000000604482015290519081900360640190fd5b600054600160a060020a0382811691161415610e8757604051600080516020611cfe8339815191528152600401808060200182810382526025815260200180611cd96025913960400191505060405180910390fd5b600160a060020a0381161515610eec5760408051600080516020611cfe833981519152815260206004820181905260248201527f4e6577204f776e65722063616e6e6f74206265207a65726f2061646472657373604482015290519081900360640190fd5b60008054600160a060020a0383811673ffffffffffffffffffffffffffffffffffffffff19831681178455604051919092169283917f8934ce4adea8d9ce0d714d2c22b86790e41b7731c84b926fbbdc1d40ff6533c99190a35050565b600460209081526000928352604080842090915290825290205481565b60005474010000000000000000000000000000000000000000900460ff1615610fc857604051600080516020611cfe833981519152815260040180806020018281038252602d815260200180611d46602d913960400191505060405180910390fd5b81600060ff8216116110295760408051600080516020611cfe833981519152815260206004820152601e60248201527f496e76616c69642c206f72206d697373696e67207065726d697373696f6e0000604482015290519081900360640190fd5b600054600160a060020a031633146110b85760ff8116600090815260056020526040902061105d903363ffffffff6116c916565b15156110b85760408051600080516020611cfe833981519152815260206004820152601260248201527f4d697373696e67207065726d697373696f6e0000000000000000000000000000604482015290519081900360640190fd5b60ff83166000908152600560205260409020610a15908363ffffffff61193f16565b60006110ed60018363ffffffff61187016565b6040015192915050565b60056020526000908152604090205481565b60015481565b600061069860018363ffffffff6116c916565b600061069860018363ffffffff61167016565b600054600160a060020a031681565b60ff82166000908152600560205260408120611166908363ffffffff6116c916565b9392505050565b60005474010000000000000000000000000000000000000000900460ff1681565b60ff82166000908152600560205260408120611166908363ffffffff61167016565b60005474010000000000000000000000000000000000000000900460ff161561121257604051600080516020611cfe833981519152815260040180806020018281038252602d815260200180611d46602d913960400191505060405180910390fd5b61122360018263ffffffff61187016565b60200151600060ff8216116112875760408051600080516020611cfe833981519152815260206004820152601e60248201527f496e76616c69642c206f72206d697373696e67207065726d697373696f6e0000604482015290519081900360640190fd5b600054600160a060020a031633146113165760ff811660009081526005602052604090206112bb903363ffffffff6116c916565b15156113165760408051600080516020611cfe833981519152815260206004820152601260248201527f4d697373696e67207065726d697373696f6e0000000000000000000000000000604482015290519081900360640190fd5b600160a060020a0382166000908152600460205260408120815b601e60ff821610156113715760ff811660009081526020839052604090205483146113695760ff81166000908152602083905260408120555b600101611330565b5061085f60018563ffffffff611a4416565b60005474010000000000000000000000000000000000000000900460ff16156113e557604051600080516020611cfe833981519152815260040180806020018281038252602d815260200180611d46602d913960400191505060405180910390fd5b6113f660018363ffffffff61187016565b60200151600060ff82161161145a5760408051600080516020611cfe833981519152815260206004820152601e60248201527f496e76616c69642c206f72206d697373696e67207065726d697373696f6e0000604482015290519081900360640190fd5b600054600160a060020a031633146114e95760ff8116600090815260056020526040902061148e903363ffffffff6116c916565b15156114e95760408051600080516020611cfe833981519152815260206004820152601260248201527f4d697373696e67207065726d697373696f6e0000000000000000000000000000604482015290519081900360640190fd5b60006114fc60018563ffffffff61167016565b6001016000908152600360205260409020805493151575010000000000000000000000000000000000000000000275ff0000000000000000000000000000000000000000001990941693909317909255505050565b60ff82166000908152600560205260408120611166908363ffffffff611c2116565b61157b611cb1565b6000821215801561158c5750825482125b15156115e75760408051600080516020611cfe833981519152815260206004820152601860248201527f496e646578206f757473696465206f6620626f756e64732e0000000000000000604482015290519081900360640190fd5b5060019081016000908152600292909201602090815260409283902083516080810185528154600160a060020a03808216835260ff74010000000000000000000000000000000000000000830481169584019590955275010000000000000000000000000000000000000000009091049093161515948101949094529091015416606082015290565b6000600160a060020a038216151561168b5750600019610698565b600160a060020a038216600090815260018401602052604081205460001901908112806116b9575083548112155b1561116657600019915050610698565b600160a060020a0381166000908152600183016020526040812054600019018181128015906116f85750835481125b949350505050565b6000600160a060020a038516151561171a57506000611867565b600160a060020a0385166000908152600187016020526040812054600019019081128015906117495750865481125b15611758576000915050611867565b505084546001908101808755600160a060020a038087166000818152848a0160209081526040808320869055805160808101825293845260ff808b168584019081528a15158684019081528a88166060880190815298865260028f01909452919093209351845491519251151575010000000000000000000000000000000000000000000275ff0000000000000000000000000000000000000000001993909416740100000000000000000000000000000000000000000274ff00000000000000000000000000000000000000001991871673ffffffffffffffffffffffffffffffffffffffff1993841617919091161791909116919091178255925190840180549190921692169190911790555b95945050505050565b611878611cb1565b611166836118868585611670565b611573565b6000600160a060020a03821615156118a557506000610698565b600160a060020a0382166000908152600184016020526040812054600019019081128015906118d45750835481125b156118e3576000915050610698565b505081546001908101808455600160a060020a03831660008181528386016020908152604080832085905593825260028701905291909120805473ffffffffffffffffffffffffffffffffffffffff1916909117905592915050565b600160a060020a038116600090815260018084016020526040822054908112806119695750835481135b15611978576000915050610698565b83548112156119ec5783546000908152600285016020818152604080842054600160a060020a031680855260018901835281852086905585855292909152808320805473ffffffffffffffffffffffffffffffffffffffff1990811690931790558654835290912080549091169055611a18565b60008181526002850160205260409020805473ffffffffffffffffffffffffffffffffffffffff191690555b5050600160a060020a031660009081526001828101602052604082209190915581546000190190915590565b600160a060020a03811660009081526001808401602052604082205490811280611a6e5750835481135b15611a7d576000915050610698565b8354811215611bad578354600090815260028501602081815260408084208054600160a060020a0390811686526001808b018552838720889055878752949093528185208154815490851673ffffffffffffffffffffffffffffffffffffffff1991821617808355835474ff000000000000000000000000000000000000000019909116740100000000000000000000000000000000000000009182900460ff90811690920217808455845475ff00000000000000000000000000000000000000000019909116750100000000000000000000000000000000000000000091829004909216151502178255918501549085018054919094169082161790925587548452909220805475ffffffffffffffffffffffffffffffffffffffffffff191681550180549091169055611a18565b60009081526002840160209081526040808320805475ffffffffffffffffffffffffffffffffffffffffffff191681556001908101805473ffffffffffffffffffffffffffffffffffffffff19169055600160a060020a039590951683528486019091528120555081546000190190915590565b6000808212158015611c335750825482125b1515611c8e5760408051600080516020611cfe833981519152815260206004820152601860248201527f496e646578206f757473696465206f6620626f756e64732e0000000000000000604482015290519081900360640190fd5b5060010160009081526002919091016020526040902054600160a060020a031690565b6040805160808101825260008082526020820181905291810182905260608101919091529056fe4e6577204f776e65722063616e6e6f74206265207468652063757272656e74206f776e657208c379a000000000000000000000000000000000000000000000000000000000436f6e747261637420616c726561647920696e20726571756573746564206c6f636b207374617465436f6e74726163742069732063757272656e746c79206c6f636b656420666f72206d6f64696669636174696f6ea165627a7a7230582091e5737e6dc3160ae43070807f83d229c50cc4bb359d770797b5df7c076d47e60029`

// DeployStorage deploys a new Ethereum contract, binding an instance of Storage to it.
func DeployStorage(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Storage, error) {
	parsed, err := abi.JSON(strings.NewReader(StorageABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(StorageBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Storage{StorageCaller: StorageCaller{contract: contract}, StorageTransactor: StorageTransactor{contract: contract}, StorageFilterer: StorageFilterer{contract: contract}}, nil
}

// Storage is an auto generated Go binding around an Ethereum contract.
type Storage struct {
	StorageCaller     // Read-only binding to the contract
	StorageTransactor // Write-only binding to the contract
	StorageFilterer   // Log filterer for contract events
}

// StorageCaller is an auto generated read-only Go binding around an Ethereum contract.
type StorageCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StorageTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StorageTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StorageFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StorageFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StorageSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StorageSession struct {
	Contract     *Storage          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StorageCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StorageCallerSession struct {
	Contract *StorageCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// StorageTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StorageTransactorSession struct {
	Contract     *StorageTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// StorageRaw is an auto generated low-level Go binding around an Ethereum contract.
type StorageRaw struct {
	Contract *Storage // Generic contract binding to access the raw methods on
}

// StorageCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StorageCallerRaw struct {
	Contract *StorageCaller // Generic read-only contract binding to access the raw methods on
}

// StorageTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StorageTransactorRaw struct {
	Contract *StorageTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStorage creates a new instance of Storage, bound to a specific deployed contract.
func NewStorage(address common.Address, backend bind.ContractBackend) (*Storage, error) {
	contract, err := bindStorage(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Storage{StorageCaller: StorageCaller{contract: contract}, StorageTransactor: StorageTransactor{contract: contract}, StorageFilterer: StorageFilterer{contract: contract}}, nil
}

// NewStorageCaller creates a new read-only instance of Storage, bound to a specific deployed contract.
func NewStorageCaller(address common.Address, caller bind.ContractCaller) (*StorageCaller, error) {
	contract, err := bindStorage(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StorageCaller{contract: contract}, nil
}

// NewStorageTransactor creates a new write-only instance of Storage, bound to a specific deployed contract.
func NewStorageTransactor(address common.Address, transactor bind.ContractTransactor) (*StorageTransactor, error) {
	contract, err := bindStorage(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StorageTransactor{contract: contract}, nil
}

// NewStorageFilterer creates a new log filterer instance of Storage, bound to a specific deployed contract.
func NewStorageFilterer(address common.Address, filterer bind.ContractFilterer) (*StorageFilterer, error) {
	contract, err := bindStorage(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StorageFilterer{contract: contract}, nil
}

// bindStorage binds a generic wrapper to an already deployed contract.
func bindStorage(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(StorageABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Storage *StorageRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Storage.Contract.StorageCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Storage *StorageRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Storage.Contract.StorageTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Storage *StorageRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Storage.Contract.StorageTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Storage *StorageCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Storage.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Storage *StorageTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Storage.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Storage *StorageTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Storage.Contract.contract.Transact(opts, method, params...)
}

// AccountAt is a free data retrieval call binding the contract method 0x00f17754.
//
// Solidity: function accountAt(index int256) constant returns(address, uint8, bool, address)
func (_Storage *StorageCaller) AccountAt(opts *bind.CallOpts, index *big.Int) (common.Address, uint8, bool, common.Address, error) {
	var (
		ret0 = new(common.Address)
		ret1 = new(uint8)
		ret2 = new(bool)
		ret3 = new(common.Address)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
		ret3,
	}
	err := _Storage.contract.Call(opts, out, "accountAt", index)
	return *ret0, *ret1, *ret2, *ret3, err
}

// AccountAt is a free data retrieval call binding the contract method 0x00f17754.
//
// Solidity: function accountAt(index int256) constant returns(address, uint8, bool, address)
func (_Storage *StorageSession) AccountAt(index *big.Int) (common.Address, uint8, bool, common.Address, error) {
	return _Storage.Contract.AccountAt(&_Storage.CallOpts, index)
}

// AccountAt is a free data retrieval call binding the contract method 0x00f17754.
//
// Solidity: function accountAt(index int256) constant returns(address, uint8, bool, address)
func (_Storage *StorageCallerSession) AccountAt(index *big.Int) (common.Address, uint8, bool, common.Address, error) {
	return _Storage.Contract.AccountAt(&_Storage.CallOpts, index)
}

// AccountExists is a free data retrieval call binding the contract method 0x75cd51ed.
//
// Solidity: function accountExists(addr address) constant returns(bool)
func (_Storage *StorageCaller) AccountExists(opts *bind.CallOpts, addr common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Storage.contract.Call(opts, out, "accountExists", addr)
	return *ret0, err
}

// AccountExists is a free data retrieval call binding the contract method 0x75cd51ed.
//
// Solidity: function accountExists(addr address) constant returns(bool)
func (_Storage *StorageSession) AccountExists(addr common.Address) (bool, error) {
	return _Storage.Contract.AccountExists(&_Storage.CallOpts, addr)
}

// AccountExists is a free data retrieval call binding the contract method 0x75cd51ed.
//
// Solidity: function accountExists(addr address) constant returns(bool)
func (_Storage *StorageCallerSession) AccountExists(addr common.Address) (bool, error) {
	return _Storage.Contract.AccountExists(&_Storage.CallOpts, addr)
}

// AccountFrozen is a free data retrieval call binding the contract method 0x624522f9.
//
// Solidity: function accountFrozen(addr address) constant returns(bool)
func (_Storage *StorageCaller) AccountFrozen(opts *bind.CallOpts, addr common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Storage.contract.Call(opts, out, "accountFrozen", addr)
	return *ret0, err
}

// AccountFrozen is a free data retrieval call binding the contract method 0x624522f9.
//
// Solidity: function accountFrozen(addr address) constant returns(bool)
func (_Storage *StorageSession) AccountFrozen(addr common.Address) (bool, error) {
	return _Storage.Contract.AccountFrozen(&_Storage.CallOpts, addr)
}

// AccountFrozen is a free data retrieval call binding the contract method 0x624522f9.
//
// Solidity: function accountFrozen(addr address) constant returns(bool)
func (_Storage *StorageCallerSession) AccountFrozen(addr common.Address) (bool, error) {
	return _Storage.Contract.AccountFrozen(&_Storage.CallOpts, addr)
}

// AccountGet is a free data retrieval call binding the contract method 0x137e37d9.
//
// Solidity: function accountGet(addr address) constant returns(uint8, bool, address)
func (_Storage *StorageCaller) AccountGet(opts *bind.CallOpts, addr common.Address) (uint8, bool, common.Address, error) {
	var (
		ret0 = new(uint8)
		ret1 = new(bool)
		ret2 = new(common.Address)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
	}
	err := _Storage.contract.Call(opts, out, "accountGet", addr)
	return *ret0, *ret1, *ret2, err
}

// AccountGet is a free data retrieval call binding the contract method 0x137e37d9.
//
// Solidity: function accountGet(addr address) constant returns(uint8, bool, address)
func (_Storage *StorageSession) AccountGet(addr common.Address) (uint8, bool, common.Address, error) {
	return _Storage.Contract.AccountGet(&_Storage.CallOpts, addr)
}

// AccountGet is a free data retrieval call binding the contract method 0x137e37d9.
//
// Solidity: function accountGet(addr address) constant returns(uint8, bool, address)
func (_Storage *StorageCallerSession) AccountGet(addr common.Address) (uint8, bool, common.Address, error) {
	return _Storage.Contract.AccountGet(&_Storage.CallOpts, addr)
}

// AccountIndexOf is a free data retrieval call binding the contract method 0x770850c8.
//
// Solidity: function accountIndexOf(addr address) constant returns(int256)
func (_Storage *StorageCaller) AccountIndexOf(opts *bind.CallOpts, addr common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Storage.contract.Call(opts, out, "accountIndexOf", addr)
	return *ret0, err
}

// AccountIndexOf is a free data retrieval call binding the contract method 0x770850c8.
//
// Solidity: function accountIndexOf(addr address) constant returns(int256)
func (_Storage *StorageSession) AccountIndexOf(addr common.Address) (*big.Int, error) {
	return _Storage.Contract.AccountIndexOf(&_Storage.CallOpts, addr)
}

// AccountIndexOf is a free data retrieval call binding the contract method 0x770850c8.
//
// Solidity: function accountIndexOf(addr address) constant returns(int256)
func (_Storage *StorageCallerSession) AccountIndexOf(addr common.Address) (*big.Int, error) {
	return _Storage.Contract.AccountIndexOf(&_Storage.CallOpts, addr)
}

// AccountKind is a free data retrieval call binding the contract method 0x351a97f8.
//
// Solidity: function accountKind(addr address) constant returns(uint8)
func (_Storage *StorageCaller) AccountKind(opts *bind.CallOpts, addr common.Address) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _Storage.contract.Call(opts, out, "accountKind", addr)
	return *ret0, err
}

// AccountKind is a free data retrieval call binding the contract method 0x351a97f8.
//
// Solidity: function accountKind(addr address) constant returns(uint8)
func (_Storage *StorageSession) AccountKind(addr common.Address) (uint8, error) {
	return _Storage.Contract.AccountKind(&_Storage.CallOpts, addr)
}

// AccountKind is a free data retrieval call binding the contract method 0x351a97f8.
//
// Solidity: function accountKind(addr address) constant returns(uint8)
func (_Storage *StorageCallerSession) AccountKind(addr common.Address) (uint8, error) {
	return _Storage.Contract.AccountKind(&_Storage.CallOpts, addr)
}

// AccountParent is a free data retrieval call binding the contract method 0x31aaa74a.
//
// Solidity: function accountParent(addr address) constant returns(address)
func (_Storage *StorageCaller) AccountParent(opts *bind.CallOpts, addr common.Address) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Storage.contract.Call(opts, out, "accountParent", addr)
	return *ret0, err
}

// AccountParent is a free data retrieval call binding the contract method 0x31aaa74a.
//
// Solidity: function accountParent(addr address) constant returns(address)
func (_Storage *StorageSession) AccountParent(addr common.Address) (common.Address, error) {
	return _Storage.Contract.AccountParent(&_Storage.CallOpts, addr)
}

// AccountParent is a free data retrieval call binding the contract method 0x31aaa74a.
//
// Solidity: function accountParent(addr address) constant returns(address)
func (_Storage *StorageCallerSession) AccountParent(addr common.Address) (common.Address, error) {
	return _Storage.Contract.AccountParent(&_Storage.CallOpts, addr)
}

// Accounts is a free data retrieval call binding the contract method 0x68cd03f6.
//
// Solidity: function accounts() constant returns(count int256)
func (_Storage *StorageCaller) Accounts(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Storage.contract.Call(opts, out, "accounts")
	return *ret0, err
}

// Accounts is a free data retrieval call binding the contract method 0x68cd03f6.
//
// Solidity: function accounts() constant returns(count int256)
func (_Storage *StorageSession) Accounts() (*big.Int, error) {
	return _Storage.Contract.Accounts(&_Storage.CallOpts)
}

// Accounts is a free data retrieval call binding the contract method 0x68cd03f6.
//
// Solidity: function accounts() constant returns(count int256)
func (_Storage *StorageCallerSession) Accounts() (*big.Int, error) {
	return _Storage.Contract.Accounts(&_Storage.CallOpts)
}

// Data is a free data retrieval call binding the contract method 0x572bc6b1.
//
// Solidity: function data( address,  uint8) constant returns(bytes32)
func (_Storage *StorageCaller) Data(opts *bind.CallOpts, arg0 common.Address, arg1 uint8) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Storage.contract.Call(opts, out, "data", arg0, arg1)
	return *ret0, err
}

// Data is a free data retrieval call binding the contract method 0x572bc6b1.
//
// Solidity: function data( address,  uint8) constant returns(bytes32)
func (_Storage *StorageSession) Data(arg0 common.Address, arg1 uint8) ([32]byte, error) {
	return _Storage.Contract.Data(&_Storage.CallOpts, arg0, arg1)
}

// Data is a free data retrieval call binding the contract method 0x572bc6b1.
//
// Solidity: function data( address,  uint8) constant returns(bytes32)
func (_Storage *StorageCallerSession) Data(arg0 common.Address, arg1 uint8) ([32]byte, error) {
	return _Storage.Contract.Data(&_Storage.CallOpts, arg0, arg1)
}

// IsLocked is a free data retrieval call binding the contract method 0xa4e2d634.
//
// Solidity: function isLocked() constant returns(bool)
func (_Storage *StorageCaller) IsLocked(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Storage.contract.Call(opts, out, "isLocked")
	return *ret0, err
}

// IsLocked is a free data retrieval call binding the contract method 0xa4e2d634.
//
// Solidity: function isLocked() constant returns(bool)
func (_Storage *StorageSession) IsLocked() (bool, error) {
	return _Storage.Contract.IsLocked(&_Storage.CallOpts)
}

// IsLocked is a free data retrieval call binding the contract method 0xa4e2d634.
//
// Solidity: function isLocked() constant returns(bool)
func (_Storage *StorageCallerSession) IsLocked() (bool, error) {
	return _Storage.Contract.IsLocked(&_Storage.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Storage *StorageCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Storage.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Storage *StorageSession) Owner() (common.Address, error) {
	return _Storage.Contract.Owner(&_Storage.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Storage *StorageCallerSession) Owner() (common.Address, error) {
	return _Storage.Contract.Owner(&_Storage.CallOpts)
}

// PermissionAt is a free data retrieval call binding the contract method 0xd7379999.
//
// Solidity: function permissionAt(kind uint8, index int256) constant returns(address)
func (_Storage *StorageCaller) PermissionAt(opts *bind.CallOpts, kind uint8, index *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Storage.contract.Call(opts, out, "permissionAt", kind, index)
	return *ret0, err
}

// PermissionAt is a free data retrieval call binding the contract method 0xd7379999.
//
// Solidity: function permissionAt(kind uint8, index int256) constant returns(address)
func (_Storage *StorageSession) PermissionAt(kind uint8, index *big.Int) (common.Address, error) {
	return _Storage.Contract.PermissionAt(&_Storage.CallOpts, kind, index)
}

// PermissionAt is a free data retrieval call binding the contract method 0xd7379999.
//
// Solidity: function permissionAt(kind uint8, index int256) constant returns(address)
func (_Storage *StorageCallerSession) PermissionAt(kind uint8, index *big.Int) (common.Address, error) {
	return _Storage.Contract.PermissionAt(&_Storage.CallOpts, kind, index)
}

// PermissionExists is a free data retrieval call binding the contract method 0x9d44ac4f.
//
// Solidity: function permissionExists(kind uint8, addr address) constant returns(bool)
func (_Storage *StorageCaller) PermissionExists(opts *bind.CallOpts, kind uint8, addr common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Storage.contract.Call(opts, out, "permissionExists", kind, addr)
	return *ret0, err
}

// PermissionExists is a free data retrieval call binding the contract method 0x9d44ac4f.
//
// Solidity: function permissionExists(kind uint8, addr address) constant returns(bool)
func (_Storage *StorageSession) PermissionExists(kind uint8, addr common.Address) (bool, error) {
	return _Storage.Contract.PermissionExists(&_Storage.CallOpts, kind, addr)
}

// PermissionExists is a free data retrieval call binding the contract method 0x9d44ac4f.
//
// Solidity: function permissionExists(kind uint8, addr address) constant returns(bool)
func (_Storage *StorageCallerSession) PermissionExists(kind uint8, addr common.Address) (bool, error) {
	return _Storage.Contract.PermissionExists(&_Storage.CallOpts, kind, addr)
}

// PermissionIndexOf is a free data retrieval call binding the contract method 0xa84694c8.
//
// Solidity: function permissionIndexOf(kind uint8, addr address) constant returns(int256)
func (_Storage *StorageCaller) PermissionIndexOf(opts *bind.CallOpts, kind uint8, addr common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Storage.contract.Call(opts, out, "permissionIndexOf", kind, addr)
	return *ret0, err
}

// PermissionIndexOf is a free data retrieval call binding the contract method 0xa84694c8.
//
// Solidity: function permissionIndexOf(kind uint8, addr address) constant returns(int256)
func (_Storage *StorageSession) PermissionIndexOf(kind uint8, addr common.Address) (*big.Int, error) {
	return _Storage.Contract.PermissionIndexOf(&_Storage.CallOpts, kind, addr)
}

// PermissionIndexOf is a free data retrieval call binding the contract method 0xa84694c8.
//
// Solidity: function permissionIndexOf(kind uint8, addr address) constant returns(int256)
func (_Storage *StorageCallerSession) PermissionIndexOf(kind uint8, addr common.Address) (*big.Int, error) {
	return _Storage.Contract.PermissionIndexOf(&_Storage.CallOpts, kind, addr)
}

// Permissions is a free data retrieval call binding the contract method 0x6476b837.
//
// Solidity: function permissions( uint8) constant returns(count int256)
func (_Storage *StorageCaller) Permissions(opts *bind.CallOpts, arg0 uint8) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Storage.contract.Call(opts, out, "permissions", arg0)
	return *ret0, err
}

// Permissions is a free data retrieval call binding the contract method 0x6476b837.
//
// Solidity: function permissions( uint8) constant returns(count int256)
func (_Storage *StorageSession) Permissions(arg0 uint8) (*big.Int, error) {
	return _Storage.Contract.Permissions(&_Storage.CallOpts, arg0)
}

// Permissions is a free data retrieval call binding the contract method 0x6476b837.
//
// Solidity: function permissions( uint8) constant returns(count int256)
func (_Storage *StorageCallerSession) Permissions(arg0 uint8) (*big.Int, error) {
	return _Storage.Contract.Permissions(&_Storage.CallOpts, arg0)
}

// AddAccount is a paid mutator transaction binding the contract method 0x116c92b7.
//
// Solidity: function addAccount(addr address, kind uint8, isFrozen bool, parent address) returns()
func (_Storage *StorageTransactor) AddAccount(opts *bind.TransactOpts, addr common.Address, kind uint8, isFrozen bool, parent common.Address) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "addAccount", addr, kind, isFrozen, parent)
}

// AddAccount is a paid mutator transaction binding the contract method 0x116c92b7.
//
// Solidity: function addAccount(addr address, kind uint8, isFrozen bool, parent address) returns()
func (_Storage *StorageSession) AddAccount(addr common.Address, kind uint8, isFrozen bool, parent common.Address) (*types.Transaction, error) {
	return _Storage.Contract.AddAccount(&_Storage.TransactOpts, addr, kind, isFrozen, parent)
}

// AddAccount is a paid mutator transaction binding the contract method 0x116c92b7.
//
// Solidity: function addAccount(addr address, kind uint8, isFrozen bool, parent address) returns()
func (_Storage *StorageTransactorSession) AddAccount(addr common.Address, kind uint8, isFrozen bool, parent common.Address) (*types.Transaction, error) {
	return _Storage.Contract.AddAccount(&_Storage.TransactOpts, addr, kind, isFrozen, parent)
}

// GrantPermission is a paid mutator transaction binding the contract method 0x1a3722b2.
//
// Solidity: function grantPermission(kind uint8, addr address) returns()
func (_Storage *StorageTransactor) GrantPermission(opts *bind.TransactOpts, kind uint8, addr common.Address) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "grantPermission", kind, addr)
}

// GrantPermission is a paid mutator transaction binding the contract method 0x1a3722b2.
//
// Solidity: function grantPermission(kind uint8, addr address) returns()
func (_Storage *StorageSession) GrantPermission(kind uint8, addr common.Address) (*types.Transaction, error) {
	return _Storage.Contract.GrantPermission(&_Storage.TransactOpts, kind, addr)
}

// GrantPermission is a paid mutator transaction binding the contract method 0x1a3722b2.
//
// Solidity: function grantPermission(kind uint8, addr address) returns()
func (_Storage *StorageTransactorSession) GrantPermission(kind uint8, addr common.Address) (*types.Transaction, error) {
	return _Storage.Contract.GrantPermission(&_Storage.TransactOpts, kind, addr)
}

// Kill is a paid mutator transaction binding the contract method 0x41c0e1b5.
//
// Solidity: function kill() returns()
func (_Storage *StorageTransactor) Kill(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "kill")
}

// Kill is a paid mutator transaction binding the contract method 0x41c0e1b5.
//
// Solidity: function kill() returns()
func (_Storage *StorageSession) Kill() (*types.Transaction, error) {
	return _Storage.Contract.Kill(&_Storage.TransactOpts)
}

// Kill is a paid mutator transaction binding the contract method 0x41c0e1b5.
//
// Solidity: function kill() returns()
func (_Storage *StorageTransactorSession) Kill() (*types.Transaction, error) {
	return _Storage.Contract.Kill(&_Storage.TransactOpts)
}

// RemoveAccount is a paid mutator transaction binding the contract method 0xc4740a95.
//
// Solidity: function removeAccount(addr address) returns()
func (_Storage *StorageTransactor) RemoveAccount(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "removeAccount", addr)
}

// RemoveAccount is a paid mutator transaction binding the contract method 0xc4740a95.
//
// Solidity: function removeAccount(addr address) returns()
func (_Storage *StorageSession) RemoveAccount(addr common.Address) (*types.Transaction, error) {
	return _Storage.Contract.RemoveAccount(&_Storage.TransactOpts, addr)
}

// RemoveAccount is a paid mutator transaction binding the contract method 0xc4740a95.
//
// Solidity: function removeAccount(addr address) returns()
func (_Storage *StorageTransactorSession) RemoveAccount(addr common.Address) (*types.Transaction, error) {
	return _Storage.Contract.RemoveAccount(&_Storage.TransactOpts, addr)
}

// RevokePermission is a paid mutator transaction binding the contract method 0x612eb5e7.
//
// Solidity: function revokePermission(kind uint8, addr address) returns()
func (_Storage *StorageTransactor) RevokePermission(opts *bind.TransactOpts, kind uint8, addr common.Address) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "revokePermission", kind, addr)
}

// RevokePermission is a paid mutator transaction binding the contract method 0x612eb5e7.
//
// Solidity: function revokePermission(kind uint8, addr address) returns()
func (_Storage *StorageSession) RevokePermission(kind uint8, addr common.Address) (*types.Transaction, error) {
	return _Storage.Contract.RevokePermission(&_Storage.TransactOpts, kind, addr)
}

// RevokePermission is a paid mutator transaction binding the contract method 0x612eb5e7.
//
// Solidity: function revokePermission(kind uint8, addr address) returns()
func (_Storage *StorageTransactorSession) RevokePermission(kind uint8, addr common.Address) (*types.Transaction, error) {
	return _Storage.Contract.RevokePermission(&_Storage.TransactOpts, kind, addr)
}

// SetAccountData is a paid mutator transaction binding the contract method 0x4f1c3b66.
//
// Solidity: function setAccountData(addr address, index uint8, customData bytes32) returns()
func (_Storage *StorageTransactor) SetAccountData(opts *bind.TransactOpts, addr common.Address, index uint8, customData [32]byte) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "setAccountData", addr, index, customData)
}

// SetAccountData is a paid mutator transaction binding the contract method 0x4f1c3b66.
//
// Solidity: function setAccountData(addr address, index uint8, customData bytes32) returns()
func (_Storage *StorageSession) SetAccountData(addr common.Address, index uint8, customData [32]byte) (*types.Transaction, error) {
	return _Storage.Contract.SetAccountData(&_Storage.TransactOpts, addr, index, customData)
}

// SetAccountData is a paid mutator transaction binding the contract method 0x4f1c3b66.
//
// Solidity: function setAccountData(addr address, index uint8, customData bytes32) returns()
func (_Storage *StorageTransactorSession) SetAccountData(addr common.Address, index uint8, customData [32]byte) (*types.Transaction, error) {
	return _Storage.Contract.SetAccountData(&_Storage.TransactOpts, addr, index, customData)
}

// SetAccountFrozen is a paid mutator transaction binding the contract method 0xcbe5404f.
//
// Solidity: function setAccountFrozen(addr address, frozen bool) returns()
func (_Storage *StorageTransactor) SetAccountFrozen(opts *bind.TransactOpts, addr common.Address, frozen bool) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "setAccountFrozen", addr, frozen)
}

// SetAccountFrozen is a paid mutator transaction binding the contract method 0xcbe5404f.
//
// Solidity: function setAccountFrozen(addr address, frozen bool) returns()
func (_Storage *StorageSession) SetAccountFrozen(addr common.Address, frozen bool) (*types.Transaction, error) {
	return _Storage.Contract.SetAccountFrozen(&_Storage.TransactOpts, addr, frozen)
}

// SetAccountFrozen is a paid mutator transaction binding the contract method 0xcbe5404f.
//
// Solidity: function setAccountFrozen(addr address, frozen bool) returns()
func (_Storage *StorageTransactorSession) SetAccountFrozen(addr common.Address, frozen bool) (*types.Transaction, error) {
	return _Storage.Contract.SetAccountFrozen(&_Storage.TransactOpts, addr, frozen)
}

// SetLocked is a paid mutator transaction binding the contract method 0x211e28b6.
//
// Solidity: function setLocked(locked bool) returns()
func (_Storage *StorageTransactor) SetLocked(opts *bind.TransactOpts, locked bool) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "setLocked", locked)
}

// SetLocked is a paid mutator transaction binding the contract method 0x211e28b6.
//
// Solidity: function setLocked(locked bool) returns()
func (_Storage *StorageSession) SetLocked(locked bool) (*types.Transaction, error) {
	return _Storage.Contract.SetLocked(&_Storage.TransactOpts, locked)
}

// SetLocked is a paid mutator transaction binding the contract method 0x211e28b6.
//
// Solidity: function setLocked(locked bool) returns()
func (_Storage *StorageTransactorSession) SetLocked(locked bool) (*types.Transaction, error) {
	return _Storage.Contract.SetLocked(&_Storage.TransactOpts, locked)
}

// TransferOwner is a paid mutator transaction binding the contract method 0x4fb2e45d.
//
// Solidity: function transferOwner(newOwner address) returns()
func (_Storage *StorageTransactor) TransferOwner(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "transferOwner", newOwner)
}

// TransferOwner is a paid mutator transaction binding the contract method 0x4fb2e45d.
//
// Solidity: function transferOwner(newOwner address) returns()
func (_Storage *StorageSession) TransferOwner(newOwner common.Address) (*types.Transaction, error) {
	return _Storage.Contract.TransferOwner(&_Storage.TransactOpts, newOwner)
}

// TransferOwner is a paid mutator transaction binding the contract method 0x4fb2e45d.
//
// Solidity: function transferOwner(newOwner address) returns()
func (_Storage *StorageTransactorSession) TransferOwner(newOwner common.Address) (*types.Transaction, error) {
	return _Storage.Contract.TransferOwner(&_Storage.TransactOpts, newOwner)
}

// StorageOwnerTransferredIterator is returned from FilterOwnerTransferred and is used to iterate over the raw logs and unpacked data for OwnerTransferred events raised by the Storage contract.
type StorageOwnerTransferredIterator struct {
	Event *StorageOwnerTransferred // Event containing the contract specifics and raw log

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
func (it *StorageOwnerTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StorageOwnerTransferred)
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
		it.Event = new(StorageOwnerTransferred)
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
func (it *StorageOwnerTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StorageOwnerTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StorageOwnerTransferred represents a OwnerTransferred event raised by the Storage contract.
type StorageOwnerTransferred struct {
	OldOwner common.Address
	NewOwner common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOwnerTransferred is a free log retrieval operation binding the contract event 0x8934ce4adea8d9ce0d714d2c22b86790e41b7731c84b926fbbdc1d40ff6533c9.
//
// Solidity: e OwnerTransferred(oldOwner indexed address, newOwner indexed address)
func (_Storage *StorageFilterer) FilterOwnerTransferred(opts *bind.FilterOpts, oldOwner []common.Address, newOwner []common.Address) (*StorageOwnerTransferredIterator, error) {

	var oldOwnerRule []interface{}
	for _, oldOwnerItem := range oldOwner {
		oldOwnerRule = append(oldOwnerRule, oldOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Storage.contract.FilterLogs(opts, "OwnerTransferred", oldOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &StorageOwnerTransferredIterator{contract: _Storage.contract, event: "OwnerTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnerTransferred is a free log subscription operation binding the contract event 0x8934ce4adea8d9ce0d714d2c22b86790e41b7731c84b926fbbdc1d40ff6533c9.
//
// Solidity: e OwnerTransferred(oldOwner indexed address, newOwner indexed address)
func (_Storage *StorageFilterer) WatchOwnerTransferred(opts *bind.WatchOpts, sink chan<- *StorageOwnerTransferred, oldOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var oldOwnerRule []interface{}
	for _, oldOwnerItem := range oldOwner {
		oldOwnerRule = append(oldOwnerRule, oldOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Storage.contract.WatchLogs(opts, "OwnerTransferred", oldOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StorageOwnerTransferred)
				if err := _Storage.contract.UnpackLog(event, "OwnerTransferred", log); err != nil {
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
