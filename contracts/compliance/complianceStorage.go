// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package compliance

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
const StorageABI = "[{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"permissionExists\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"k\",\"type\":\"bytes32\"}],\"name\":\"deleteBytes32\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"k\",\"type\":\"bytes32\"}],\"name\":\"deleteAddress\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"getInt256\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bool\",\"name\":\"locked\",\"type\":\"bool\"}],\"name\":\"setLocked\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"getAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"k\",\"type\":\"bytes32\"}],\"name\":\"deleteBool\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"k\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"v\",\"type\":\"bytes\"}],\"name\":\"setBytes\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"getUint256\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"k\",\"type\":\"bytes32[]\"}],\"name\":\"getAddresses\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"o\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"permissionIndexOf\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"kill\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"k\",\"type\":\"bytes32\"}],\"name\":\"deleteInt256\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"k\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"v\",\"type\":\"bytes32\"}],\"name\":\"setBytes32\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"k\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"v\",\"type\":\"uint256\"}],\"name\":\"setUint256\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwner\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"k\",\"type\":\"bytes32\"}],\"name\":\"deleteUint256\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ZERO_ADDRESS\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"k\",\"type\":\"bytes32\"}],\"name\":\"deleteBytes\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"k\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"v\",\"type\":\"string\"}],\"name\":\"setString\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"k\",\"type\":\"bytes32[]\"}],\"name\":\"getBools\",\"outputs\":[{\"internalType\":\"bool[]\",\"name\":\"o\",\"type\":\"bool[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"getBool\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"getString\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isLocked\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"getBytes32\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"permissions\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"count\",\"type\":\"int256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"k\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"v\",\"type\":\"bool\"}],\"name\":\"setBool\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"getBytes\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"k\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"v\",\"type\":\"address\"}],\"name\":\"setAddress\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"k\",\"type\":\"bytes32[]\"}],\"name\":\"getUint256s\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"o\",\"type\":\"uint256[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"int256\",\"name\":\"index\",\"type\":\"int256\"}],\"name\":\"permissionAt\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"grant\",\"type\":\"bool\"}],\"name\":\"setPermission\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"k\",\"type\":\"bytes32[]\"}],\"name\":\"getInt256s\",\"outputs\":[{\"internalType\":\"int256[]\",\"name\":\"o\",\"type\":\"int256[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"k\",\"type\":\"bytes32\"},{\"internalType\":\"int256\",\"name\":\"v\",\"type\":\"int256\"}],\"name\":\"setInt256\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"k\",\"type\":\"bytes32\"}],\"name\":\"deleteString\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"k\",\"type\":\"bytes32[]\"}],\"name\":\"getBytes32s\",\"outputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"o\",\"type\":\"bytes32[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnerTransferred\",\"type\":\"event\"}]"

// StorageBin is the compiled bytecode used for deploying new contracts.
const StorageBin = `608060405260018054600080546001600160a01b031916331790556001600160a81b0319169055611ee2806100356000396000f3fe608060405234801561001057600080fd5b50600436106102775760003560e01c8063616b59f611610160578063abfdcced116100d8578063ec6263c01161008c578063f64f41db11610071578063f64f41db14610a13578063f6bb3cc414610a36578063ff957c4614610a5357610277565b8063ec6263c014610942578063f06e93261461097057610277565b8063ca446dd9116100bd578063ca446dd914610856578063d5f249e614610882578063ea11cd421461092557610277565b8063abfdcced14610814578063c031a1801461083957610277565b80638da5cb5b1161012f578063a4e2d63411610114578063a4e2d634146107e7578063a6ed563e146107ef578063ab8c71c01461080c57610277565b80638da5cb5b1461074d578063986e791a1461075557610277565b8063616b59f6146105f95780636e8995501461061657806374f432fb1461068d5780637ae1cfca1461073057610277565b806338bc01b5116101f35780634e91db08116101c25780634fb2e45d116101a75780634fb2e45d146105ae57806351f8e7dd146105d4578063538ba4f9146105f157610277565b80634e91db08146105685780634f3029c21461058b57610277565b806338bc01b51461042a5780633ad8f6871461051d57806341c0e1b5146105435780634869ecc91461054b57610277565b8063211e28b61161024a5780632c62ff2d1161022f5780632c62ff2d146103795780632e28d0841461039657806333598b001461040d57610277565b8063211e28b61461032157806321f8a7211461034057610277565b806301952ffe1461027c5780630b9adc57146102b65780630e14a376146102d557806316c7d1d5146102f2575b600080fd5b6102a26004803603602081101561029257600080fd5b50356001600160a01b0316610af6565b604080519115158252519081900360200190f35b6102d3600480360360208110156102cc57600080fd5b5035610b0f565b005b6102d3600480360360208110156102eb57600080fd5b5035610b85565b61030f6004803603602081101561030857600080fd5b5035610c08565b60408051918252519081900360200190f35b6102d36004803603602081101561033757600080fd5b50351515610c1a565b61035d6004803603602081101561035657600080fd5b5035610d1d565b604080516001600160a01b039092168252519081900360200190f35b6102d36004803603602081101561038f57600080fd5b5035610d38565b6102d3600480360360408110156103ac57600080fd5b813591908101906040810160208201356401000000008111156103ce57600080fd5b8201836020820111156103e057600080fd5b8035906020019184600183028401116401000000008311171561040257600080fd5b509092509050610db5565b61030f6004803603602081101561042357600080fd5b5035610e39565b6104cd6004803603602081101561044057600080fd5b81019060208101813564010000000081111561045b57600080fd5b82018360208201111561046d57600080fd5b8035906020019184602083028401116401000000008311171561048f57600080fd5b919080806020026020016040519081016040528093929190818152602001838360200280828437600092019190915250929550610e4b945050505050565b60408051602080825283518183015283519192839290830191858101910280838360005b838110156105095781810151838201526020016104f1565b505050509050019250505060405180910390f35b61030f6004803603602081101561053357600080fd5b50356001600160a01b0316610ef3565b6102d3610f06565b6102d36004803603602081101561056157600080fd5b5035610f8e565b6102d36004803603604081101561057e57600080fd5b5080359060200135611004565b6102d3600480360360408110156105a157600080fd5b508035906020013561107b565b6102d3600480360360208110156105c457600080fd5b50356001600160a01b03166110f2565b6102d3600480360360208110156105ea57600080fd5b5035611264565b61035d6112da565b6102d36004803603602081101561060f57600080fd5b50356112e9565b6102d36004803603604081101561062c57600080fd5b8135919081019060408101602082013564010000000081111561064e57600080fd5b82018360208201111561066057600080fd5b8035906020019184600183028401116401000000008311171561068257600080fd5b509092509050611368565b6104cd600480360360208110156106a357600080fd5b8101906020810181356401000000008111156106be57600080fd5b8201836020820111156106d057600080fd5b803590602001918460208302840111640100000000831117156106f257600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295506113e6945050505050565b6102a26004803603602081101561074657600080fd5b503561147a565b61035d61148f565b6107726004803603602081101561076b57600080fd5b503561149e565b6040805160208082528351818301528351919283929083019185019080838360005b838110156107ac578181015183820152602001610794565b50505050905090810190601f1680156107d95780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b6102a2611539565b61030f6004803603602081101561080557600080fd5b5035611549565b61030f61155b565b6102d36004803603604081101561082a57600080fd5b50803590602001351515611561565b6107726004803603602081101561084f57600080fd5b50356115e6565b6102d36004803603604081101561086c57600080fd5b50803590602001356001600160a01b031661164e565b6104cd6004803603602081101561089857600080fd5b8101906020810181356401000000008111156108b357600080fd5b8201836020820111156108c557600080fd5b803590602001918460208302840111640100000000831117156108e757600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295506116e1945050505050565b61035d6004803603602081101561093b57600080fd5b5035611763565b6102d36004803603604081101561095857600080fd5b506001600160a01b0381351690602001351515611776565b6104cd6004803603602081101561098657600080fd5b8101906020810181356401000000008111156109a157600080fd5b8201836020820111156109b357600080fd5b803590602001918460208302840111640100000000831117156109d557600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295506118c3945050505050565b6102d360048036036040811015610a2957600080fd5b5080359060200135611945565b6102d360048036036020811015610a4c57600080fd5b50356119bc565b6104cd60048036036020811015610a6957600080fd5b810190602081018135640100000000811115610a8457600080fd5b820183602082011115610a9657600080fd5b80359060200191846020830284011164010000000083111715610ab857600080fd5b919080806020026020016040519081016040528093929190818152602001838360200280828437600092019190915250929550611a38945050505050565b6000610b0960028363ffffffff611aba16565b92915050565b610b2060023363ffffffff611aba16565b80610b3557506000546001600160a01b031633145b610b74576040805162461bcd60e51b815260206004820152601a6024820152600080516020611e66833981519152604482015290519081900360640190fd5b600090815260076020526040812055565b610b9660023363ffffffff611aba16565b80610bab57506000546001600160a01b031633145b610bea576040805162461bcd60e51b815260206004820152601a6024820152600080516020611e66833981519152604482015290519081900360640190fd5b600090815260056020526040902080546001600160a01b0319169055565b60096020526000908152604090205481565b6000546001600160a01b0316331480610c4357506001546000546001600160a01b039081169116145b610c94576040805162461bcd60e51b815260206004820152601960248201527f4f776e6572206163636f756e7420697320726571756972656400000000000000604482015290519081900360640190fd5b60015460ff600160a01b9091041615158115151415610ce45760405162461bcd60e51b8152600401808060200182810382526028815260200180611e866028913960400191505060405180910390fd5b60018054911515600160a01b027fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff909216919091179055565b6005602052600090815260409020546001600160a01b031681565b610d4960023363ffffffff611aba16565b80610d5e57506000546001600160a01b031633145b610d9d576040805162461bcd60e51b815260206004820152601a6024820152600080516020611e66833981519152604482015290519081900360640190fd5b6000908152600660205260409020805460ff19169055565b610dc660023363ffffffff611aba16565b80610ddb57506000546001600160a01b031633145b610e1a576040805162461bcd60e51b815260206004820152601a6024820152600080516020611e66833981519152604482015290519081900360640190fd5b6000838152600860205260409020610e33908383611d65565b50505050565b600b6020526000908152604090205481565b60608151604051908082528060200260200182016040528015610e78578160200160208202803883390190505b50905060005b8251811015610eed5760056000848381518110610e9757fe5b6020026020010151815260200190815260200160002060009054906101000a90046001600160a01b0316828281518110610ecd57fe5b6001600160a01b0390921660209283029190910190910152600101610e7e565b50919050565b6000610b0960028363ffffffff611af116565b6000546001600160a01b0316331480610f2f57506001546000546001600160a01b039081169116145b610f80576040805162461bcd60e51b815260206004820152601960248201527f4f776e6572206163636f756e7420697320726571756972656400000000000000604482015290519081900360640190fd5b6000546001600160a01b0316ff5b610f9f60023363ffffffff611aba16565b80610fb457506000546001600160a01b031633145b610ff3576040805162461bcd60e51b815260206004820152601a6024820152600080516020611e66833981519152604482015290519081900360640190fd5b600090815260096020526040812055565b61101560023363ffffffff611aba16565b8061102a57506000546001600160a01b031633145b611069576040805162461bcd60e51b815260206004820152601a6024820152600080516020611e66833981519152604482015290519081900360640190fd5b60009182526007602052604090912055565b61108c60023363ffffffff611aba16565b806110a157506000546001600160a01b031633145b6110e0576040805162461bcd60e51b815260206004820152601a6024820152600080516020611e66833981519152604482015290519081900360640190fd5b6000918252600b602052604090912055565b6000546001600160a01b031633148061111b57506001546000546001600160a01b039081169116145b61116c576040805162461bcd60e51b815260206004820152601960248201527f4f776e6572206163636f756e7420697320726571756972656400000000000000604482015290519081900360640190fd5b6000546001600160a01b03828116911614156111b95760405162461bcd60e51b8152600401808060200182810382526025815260200180611e416025913960400191505060405180910390fd5b6001600160a01b038116611214576040805162461bcd60e51b815260206004820181905260248201527f4e6577204f776e65722063616e6e6f74206265207a65726f2061646472657373604482015290519081900360640190fd5b600080546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8934ce4adea8d9ce0d714d2c22b86790e41b7731c84b926fbbdc1d40ff6533c99190a35050565b61127560023363ffffffff611aba16565b8061128a57506000546001600160a01b031633145b6112c9576040805162461bcd60e51b815260206004820152601a6024820152600080516020611e66833981519152604482015290519081900360640190fd5b6000908152600b6020526040812055565b6001546001600160a01b031681565b6112fa60023363ffffffff611aba16565b8061130f57506000546001600160a01b031633145b61134e576040805162461bcd60e51b815260206004820152601a6024820152600080516020611e66833981519152604482015290519081900360640190fd5b600081815260086020526040812061136591611de3565b50565b61137960023363ffffffff611aba16565b8061138e57506000546001600160a01b031633145b6113cd576040805162461bcd60e51b815260206004820152601a6024820152600080516020611e66833981519152604482015290519081900360640190fd5b6000838152600a60205260409020610e33908383611d65565b60608151604051908082528060200260200182016040528015611413578160200160208202803883390190505b50905060005b8251811015610eed576006600084838151811061143257fe5b6020026020010151815260200190815260200160002060009054906101000a900460ff1682828151811061146257fe5b91151560209283029190910190910152600101611419565b60066020526000908152604090205460ff1681565b6000546001600160a01b031681565b600a6020908152600091825260409182902080548351601f6002600019610100600186161502019093169290920491820184900484028101840190945280845290918301828280156115315780601f1061150657610100808354040283529160200191611531565b820191906000526020600020905b81548152906001019060200180831161151457829003601f168201915b505050505081565b600154600160a01b900460ff1681565b60076020526000908152604090205481565b60025481565b61157260023363ffffffff611aba16565b8061158757506000546001600160a01b031633145b6115c6576040805162461bcd60e51b815260206004820152601a6024820152600080516020611e66833981519152604482015290519081900360640190fd5b600091825260066020526040909120805460ff1916911515919091179055565b60086020908152600091825260409182902080548351601f6002600019610100600186161502019093169290920491820184900484028101840190945280845290918301828280156115315780601f1061150657610100808354040283529160200191611531565b61165f60023363ffffffff611aba16565b8061167457506000546001600160a01b031633145b6116b3576040805162461bcd60e51b815260206004820152601a6024820152600080516020611e66833981519152604482015290519081900360640190fd5b60009182526005602052604090912080546001600160a01b0319166001600160a01b03909216919091179055565b6060815160405190808252806020026020018201604052801561170e578160200160208202803883390190505b50905060005b8251811015610eed57600b600084838151811061172d57fe5b602002602001015181526020019081526020016000205482828151811061175057fe5b6020908102919091010152600101611714565b6000610b0960028363ffffffff611b4f16565b6000546001600160a01b031633148061179f57506001546000546001600160a01b039081169116145b6117f0576040805162461bcd60e51b815260206004820152601960248201527f4f776e6572206163636f756e7420697320726571756972656400000000000000604482015290519081900360640190fd5b801561185d5761180760028363ffffffff611bd516565b611858576040805162461bcd60e51b815260206004820152601e60248201527f4164647265737320616c726561647920686173207065726d697373696f6e0000604482015290519081900360640190fd5b6118bf565b61186e60028363ffffffff611c7a16565b6118bf576040805162461bcd60e51b815260206004820152601e60248201527f41646472657373207065726d697373696f6e20646f6e27742065786973740000604482015290519081900360640190fd5b5050565b606081516040519080825280602002602001820160405280156118f0578160200160208202803883390190505b50905060005b8251811015610eed576009600084838151811061190f57fe5b602002602001015181526020019081526020016000205482828151811061193257fe5b60209081029190910101526001016118f6565b61195660023363ffffffff611aba16565b8061196b57506000546001600160a01b031633145b6119aa576040805162461bcd60e51b815260206004820152601a6024820152600080516020611e66833981519152604482015290519081900360640190fd5b60009182526009602052604090912055565b6119cd60023363ffffffff611aba16565b806119e257506000546001600160a01b031633145b611a21576040805162461bcd60e51b815260206004820152601a6024820152600080516020611e66833981519152604482015290519081900360640190fd5b6000818152600a6020526040812061136591611de3565b60608151604051908082528060200260200182016040528015611a65578160200160208202803883390190505b50905060005b8251811015610eed5760076000848381518110611a8457fe5b6020026020010151815260200190815260200160002054828281518110611aa757fe5b6020908102919091010152600101611a6b565b6001600160a01b038116600090815260018301602052604081205460001901818112801590611ae95750835481125b949350505050565b60006001600160a01b038216611b0a5750600019610b09565b6001600160a01b03821660009081526001840160205260408120546000190190811280611b38575083548112155b15611b4857600019915050610b09565b9392505050565b6000808212158015611b615750825482125b611bb2576040805162461bcd60e51b815260206004820152601860248201527f496e646578206f757473696465206f6620626f756e64732e0000000000000000604482015290519081900360640190fd5b50600101600090815260029190910160205260409020546001600160a01b031690565b60006001600160a01b038216611bed57506000610b09565b6001600160a01b038216600090815260018401602052604081205460001901908112801590611c1c5750835481125b15611c2b576000915050610b09565b5050815460019081018084556001600160a01b0383166000818152838601602090815260408083208590559382526002870190529190912080546001600160a01b031916909117905592915050565b6001600160a01b03811660009081526001808401602052604082205490811280611ca45750835481135b15611cb3576000915050610b09565b8354811215611d1a57835460009081526002850160208181526040808420546001600160a01b03168085526001890183528185208690558585529290915280832080546001600160a01b031990811690931790558654835290912080549091169055611d39565b6000818152600285016020526040902080546001600160a01b03191690555b50506001600160a01b031660009081526001828101602052604082209190915581546000190190915590565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10611da65782800160ff19823516178555611dd3565b82800160010185558215611dd3579182015b82811115611dd3578235825591602001919060010190611db8565b50611ddf929150611e23565b5090565b50805460018160011615610100020316600290046000825580601f10611e095750611365565b601f01602090049060005260206000209081019061136591905b611e3d91905b80821115611ddf5760008155600101611e29565b9056fe4e6577204f776e65722063616e6e6f74206265207468652063757272656e74206f776e65724d697373696e672073746f72616765207065726d697373696f6e000000000000436f6e747261637420616c726561647920696e20726571756573746564206c6f636b207374617465a265627a7a72315820804e27d37140a1f55b34b27b1504b5e601f969ea8404f0c2281c05b0e04912be64736f6c634300050b0032`

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

// ZEROADDRESS is a free data retrieval call binding the contract method 0x538ba4f9.
//
// Solidity: function ZERO_ADDRESS() constant returns(address)
func (_Storage *StorageCaller) ZEROADDRESS(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Storage.contract.Call(opts, out, "ZERO_ADDRESS")
	return *ret0, err
}

// ZEROADDRESS is a free data retrieval call binding the contract method 0x538ba4f9.
//
// Solidity: function ZERO_ADDRESS() constant returns(address)
func (_Storage *StorageSession) ZEROADDRESS() (common.Address, error) {
	return _Storage.Contract.ZEROADDRESS(&_Storage.CallOpts)
}

// ZEROADDRESS is a free data retrieval call binding the contract method 0x538ba4f9.
//
// Solidity: function ZERO_ADDRESS() constant returns(address)
func (_Storage *StorageCallerSession) ZEROADDRESS() (common.Address, error) {
	return _Storage.Contract.ZEROADDRESS(&_Storage.CallOpts)
}

// GetAddress is a free data retrieval call binding the contract method 0x21f8a721.
//
// Solidity: function getAddress( bytes32) constant returns(address)
func (_Storage *StorageCaller) GetAddress(opts *bind.CallOpts, arg0 [32]byte) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Storage.contract.Call(opts, out, "getAddress", arg0)
	return *ret0, err
}

// GetAddress is a free data retrieval call binding the contract method 0x21f8a721.
//
// Solidity: function getAddress( bytes32) constant returns(address)
func (_Storage *StorageSession) GetAddress(arg0 [32]byte) (common.Address, error) {
	return _Storage.Contract.GetAddress(&_Storage.CallOpts, arg0)
}

// GetAddress is a free data retrieval call binding the contract method 0x21f8a721.
//
// Solidity: function getAddress( bytes32) constant returns(address)
func (_Storage *StorageCallerSession) GetAddress(arg0 [32]byte) (common.Address, error) {
	return _Storage.Contract.GetAddress(&_Storage.CallOpts, arg0)
}

// GetAddresses is a free data retrieval call binding the contract method 0x38bc01b5.
//
// Solidity: function getAddresses(k bytes32[]) constant returns(o address[])
func (_Storage *StorageCaller) GetAddresses(opts *bind.CallOpts, k [][32]byte) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _Storage.contract.Call(opts, out, "getAddresses", k)
	return *ret0, err
}

// GetAddresses is a free data retrieval call binding the contract method 0x38bc01b5.
//
// Solidity: function getAddresses(k bytes32[]) constant returns(o address[])
func (_Storage *StorageSession) GetAddresses(k [][32]byte) ([]common.Address, error) {
	return _Storage.Contract.GetAddresses(&_Storage.CallOpts, k)
}

// GetAddresses is a free data retrieval call binding the contract method 0x38bc01b5.
//
// Solidity: function getAddresses(k bytes32[]) constant returns(o address[])
func (_Storage *StorageCallerSession) GetAddresses(k [][32]byte) ([]common.Address, error) {
	return _Storage.Contract.GetAddresses(&_Storage.CallOpts, k)
}

// GetBool is a free data retrieval call binding the contract method 0x7ae1cfca.
//
// Solidity: function getBool( bytes32) constant returns(bool)
func (_Storage *StorageCaller) GetBool(opts *bind.CallOpts, arg0 [32]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Storage.contract.Call(opts, out, "getBool", arg0)
	return *ret0, err
}

// GetBool is a free data retrieval call binding the contract method 0x7ae1cfca.
//
// Solidity: function getBool( bytes32) constant returns(bool)
func (_Storage *StorageSession) GetBool(arg0 [32]byte) (bool, error) {
	return _Storage.Contract.GetBool(&_Storage.CallOpts, arg0)
}

// GetBool is a free data retrieval call binding the contract method 0x7ae1cfca.
//
// Solidity: function getBool( bytes32) constant returns(bool)
func (_Storage *StorageCallerSession) GetBool(arg0 [32]byte) (bool, error) {
	return _Storage.Contract.GetBool(&_Storage.CallOpts, arg0)
}

// GetBools is a free data retrieval call binding the contract method 0x74f432fb.
//
// Solidity: function getBools(k bytes32[]) constant returns(o bool[])
func (_Storage *StorageCaller) GetBools(opts *bind.CallOpts, k [][32]byte) ([]bool, error) {
	var (
		ret0 = new([]bool)
	)
	out := ret0
	err := _Storage.contract.Call(opts, out, "getBools", k)
	return *ret0, err
}

// GetBools is a free data retrieval call binding the contract method 0x74f432fb.
//
// Solidity: function getBools(k bytes32[]) constant returns(o bool[])
func (_Storage *StorageSession) GetBools(k [][32]byte) ([]bool, error) {
	return _Storage.Contract.GetBools(&_Storage.CallOpts, k)
}

// GetBools is a free data retrieval call binding the contract method 0x74f432fb.
//
// Solidity: function getBools(k bytes32[]) constant returns(o bool[])
func (_Storage *StorageCallerSession) GetBools(k [][32]byte) ([]bool, error) {
	return _Storage.Contract.GetBools(&_Storage.CallOpts, k)
}

// GetBytes is a free data retrieval call binding the contract method 0xc031a180.
//
// Solidity: function getBytes( bytes32) constant returns(bytes)
func (_Storage *StorageCaller) GetBytes(opts *bind.CallOpts, arg0 [32]byte) ([]byte, error) {
	var (
		ret0 = new([]byte)
	)
	out := ret0
	err := _Storage.contract.Call(opts, out, "getBytes", arg0)
	return *ret0, err
}

// GetBytes is a free data retrieval call binding the contract method 0xc031a180.
//
// Solidity: function getBytes( bytes32) constant returns(bytes)
func (_Storage *StorageSession) GetBytes(arg0 [32]byte) ([]byte, error) {
	return _Storage.Contract.GetBytes(&_Storage.CallOpts, arg0)
}

// GetBytes is a free data retrieval call binding the contract method 0xc031a180.
//
// Solidity: function getBytes( bytes32) constant returns(bytes)
func (_Storage *StorageCallerSession) GetBytes(arg0 [32]byte) ([]byte, error) {
	return _Storage.Contract.GetBytes(&_Storage.CallOpts, arg0)
}

// GetBytes32 is a free data retrieval call binding the contract method 0xa6ed563e.
//
// Solidity: function getBytes32( bytes32) constant returns(bytes32)
func (_Storage *StorageCaller) GetBytes32(opts *bind.CallOpts, arg0 [32]byte) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Storage.contract.Call(opts, out, "getBytes32", arg0)
	return *ret0, err
}

// GetBytes32 is a free data retrieval call binding the contract method 0xa6ed563e.
//
// Solidity: function getBytes32( bytes32) constant returns(bytes32)
func (_Storage *StorageSession) GetBytes32(arg0 [32]byte) ([32]byte, error) {
	return _Storage.Contract.GetBytes32(&_Storage.CallOpts, arg0)
}

// GetBytes32 is a free data retrieval call binding the contract method 0xa6ed563e.
//
// Solidity: function getBytes32( bytes32) constant returns(bytes32)
func (_Storage *StorageCallerSession) GetBytes32(arg0 [32]byte) ([32]byte, error) {
	return _Storage.Contract.GetBytes32(&_Storage.CallOpts, arg0)
}

// GetBytes32s is a free data retrieval call binding the contract method 0xff957c46.
//
// Solidity: function getBytes32s(k bytes32[]) constant returns(o bytes32[])
func (_Storage *StorageCaller) GetBytes32s(opts *bind.CallOpts, k [][32]byte) ([][32]byte, error) {
	var (
		ret0 = new([][32]byte)
	)
	out := ret0
	err := _Storage.contract.Call(opts, out, "getBytes32s", k)
	return *ret0, err
}

// GetBytes32s is a free data retrieval call binding the contract method 0xff957c46.
//
// Solidity: function getBytes32s(k bytes32[]) constant returns(o bytes32[])
func (_Storage *StorageSession) GetBytes32s(k [][32]byte) ([][32]byte, error) {
	return _Storage.Contract.GetBytes32s(&_Storage.CallOpts, k)
}

// GetBytes32s is a free data retrieval call binding the contract method 0xff957c46.
//
// Solidity: function getBytes32s(k bytes32[]) constant returns(o bytes32[])
func (_Storage *StorageCallerSession) GetBytes32s(k [][32]byte) ([][32]byte, error) {
	return _Storage.Contract.GetBytes32s(&_Storage.CallOpts, k)
}

// GetInt256 is a free data retrieval call binding the contract method 0x16c7d1d5.
//
// Solidity: function getInt256( bytes32) constant returns(int256)
func (_Storage *StorageCaller) GetInt256(opts *bind.CallOpts, arg0 [32]byte) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Storage.contract.Call(opts, out, "getInt256", arg0)
	return *ret0, err
}

// GetInt256 is a free data retrieval call binding the contract method 0x16c7d1d5.
//
// Solidity: function getInt256( bytes32) constant returns(int256)
func (_Storage *StorageSession) GetInt256(arg0 [32]byte) (*big.Int, error) {
	return _Storage.Contract.GetInt256(&_Storage.CallOpts, arg0)
}

// GetInt256 is a free data retrieval call binding the contract method 0x16c7d1d5.
//
// Solidity: function getInt256( bytes32) constant returns(int256)
func (_Storage *StorageCallerSession) GetInt256(arg0 [32]byte) (*big.Int, error) {
	return _Storage.Contract.GetInt256(&_Storage.CallOpts, arg0)
}

// GetInt256s is a free data retrieval call binding the contract method 0xf06e9326.
//
// Solidity: function getInt256s(k bytes32[]) constant returns(o int256[])
func (_Storage *StorageCaller) GetInt256s(opts *bind.CallOpts, k [][32]byte) ([]*big.Int, error) {
	var (
		ret0 = new([]*big.Int)
	)
	out := ret0
	err := _Storage.contract.Call(opts, out, "getInt256s", k)
	return *ret0, err
}

// GetInt256s is a free data retrieval call binding the contract method 0xf06e9326.
//
// Solidity: function getInt256s(k bytes32[]) constant returns(o int256[])
func (_Storage *StorageSession) GetInt256s(k [][32]byte) ([]*big.Int, error) {
	return _Storage.Contract.GetInt256s(&_Storage.CallOpts, k)
}

// GetInt256s is a free data retrieval call binding the contract method 0xf06e9326.
//
// Solidity: function getInt256s(k bytes32[]) constant returns(o int256[])
func (_Storage *StorageCallerSession) GetInt256s(k [][32]byte) ([]*big.Int, error) {
	return _Storage.Contract.GetInt256s(&_Storage.CallOpts, k)
}

// GetString is a free data retrieval call binding the contract method 0x986e791a.
//
// Solidity: function getString( bytes32) constant returns(string)
func (_Storage *StorageCaller) GetString(opts *bind.CallOpts, arg0 [32]byte) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Storage.contract.Call(opts, out, "getString", arg0)
	return *ret0, err
}

// GetString is a free data retrieval call binding the contract method 0x986e791a.
//
// Solidity: function getString( bytes32) constant returns(string)
func (_Storage *StorageSession) GetString(arg0 [32]byte) (string, error) {
	return _Storage.Contract.GetString(&_Storage.CallOpts, arg0)
}

// GetString is a free data retrieval call binding the contract method 0x986e791a.
//
// Solidity: function getString( bytes32) constant returns(string)
func (_Storage *StorageCallerSession) GetString(arg0 [32]byte) (string, error) {
	return _Storage.Contract.GetString(&_Storage.CallOpts, arg0)
}

// GetUint256 is a free data retrieval call binding the contract method 0x33598b00.
//
// Solidity: function getUint256( bytes32) constant returns(uint256)
func (_Storage *StorageCaller) GetUint256(opts *bind.CallOpts, arg0 [32]byte) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Storage.contract.Call(opts, out, "getUint256", arg0)
	return *ret0, err
}

// GetUint256 is a free data retrieval call binding the contract method 0x33598b00.
//
// Solidity: function getUint256( bytes32) constant returns(uint256)
func (_Storage *StorageSession) GetUint256(arg0 [32]byte) (*big.Int, error) {
	return _Storage.Contract.GetUint256(&_Storage.CallOpts, arg0)
}

// GetUint256 is a free data retrieval call binding the contract method 0x33598b00.
//
// Solidity: function getUint256( bytes32) constant returns(uint256)
func (_Storage *StorageCallerSession) GetUint256(arg0 [32]byte) (*big.Int, error) {
	return _Storage.Contract.GetUint256(&_Storage.CallOpts, arg0)
}

// GetUint256s is a free data retrieval call binding the contract method 0xd5f249e6.
//
// Solidity: function getUint256s(k bytes32[]) constant returns(o uint256[])
func (_Storage *StorageCaller) GetUint256s(opts *bind.CallOpts, k [][32]byte) ([]*big.Int, error) {
	var (
		ret0 = new([]*big.Int)
	)
	out := ret0
	err := _Storage.contract.Call(opts, out, "getUint256s", k)
	return *ret0, err
}

// GetUint256s is a free data retrieval call binding the contract method 0xd5f249e6.
//
// Solidity: function getUint256s(k bytes32[]) constant returns(o uint256[])
func (_Storage *StorageSession) GetUint256s(k [][32]byte) ([]*big.Int, error) {
	return _Storage.Contract.GetUint256s(&_Storage.CallOpts, k)
}

// GetUint256s is a free data retrieval call binding the contract method 0xd5f249e6.
//
// Solidity: function getUint256s(k bytes32[]) constant returns(o uint256[])
func (_Storage *StorageCallerSession) GetUint256s(k [][32]byte) ([]*big.Int, error) {
	return _Storage.Contract.GetUint256s(&_Storage.CallOpts, k)
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

// PermissionAt is a free data retrieval call binding the contract method 0xea11cd42.
//
// Solidity: function permissionAt(index int256) constant returns(address)
func (_Storage *StorageCaller) PermissionAt(opts *bind.CallOpts, index *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Storage.contract.Call(opts, out, "permissionAt", index)
	return *ret0, err
}

// PermissionAt is a free data retrieval call binding the contract method 0xea11cd42.
//
// Solidity: function permissionAt(index int256) constant returns(address)
func (_Storage *StorageSession) PermissionAt(index *big.Int) (common.Address, error) {
	return _Storage.Contract.PermissionAt(&_Storage.CallOpts, index)
}

// PermissionAt is a free data retrieval call binding the contract method 0xea11cd42.
//
// Solidity: function permissionAt(index int256) constant returns(address)
func (_Storage *StorageCallerSession) PermissionAt(index *big.Int) (common.Address, error) {
	return _Storage.Contract.PermissionAt(&_Storage.CallOpts, index)
}

// PermissionExists is a free data retrieval call binding the contract method 0x01952ffe.
//
// Solidity: function permissionExists(addr address) constant returns(bool)
func (_Storage *StorageCaller) PermissionExists(opts *bind.CallOpts, addr common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Storage.contract.Call(opts, out, "permissionExists", addr)
	return *ret0, err
}

// PermissionExists is a free data retrieval call binding the contract method 0x01952ffe.
//
// Solidity: function permissionExists(addr address) constant returns(bool)
func (_Storage *StorageSession) PermissionExists(addr common.Address) (bool, error) {
	return _Storage.Contract.PermissionExists(&_Storage.CallOpts, addr)
}

// PermissionExists is a free data retrieval call binding the contract method 0x01952ffe.
//
// Solidity: function permissionExists(addr address) constant returns(bool)
func (_Storage *StorageCallerSession) PermissionExists(addr common.Address) (bool, error) {
	return _Storage.Contract.PermissionExists(&_Storage.CallOpts, addr)
}

// PermissionIndexOf is a free data retrieval call binding the contract method 0x3ad8f687.
//
// Solidity: function permissionIndexOf(addr address) constant returns(int256)
func (_Storage *StorageCaller) PermissionIndexOf(opts *bind.CallOpts, addr common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Storage.contract.Call(opts, out, "permissionIndexOf", addr)
	return *ret0, err
}

// PermissionIndexOf is a free data retrieval call binding the contract method 0x3ad8f687.
//
// Solidity: function permissionIndexOf(addr address) constant returns(int256)
func (_Storage *StorageSession) PermissionIndexOf(addr common.Address) (*big.Int, error) {
	return _Storage.Contract.PermissionIndexOf(&_Storage.CallOpts, addr)
}

// PermissionIndexOf is a free data retrieval call binding the contract method 0x3ad8f687.
//
// Solidity: function permissionIndexOf(addr address) constant returns(int256)
func (_Storage *StorageCallerSession) PermissionIndexOf(addr common.Address) (*big.Int, error) {
	return _Storage.Contract.PermissionIndexOf(&_Storage.CallOpts, addr)
}

// Permissions is a free data retrieval call binding the contract method 0xab8c71c0.
//
// Solidity: function permissions() constant returns(count int256)
func (_Storage *StorageCaller) Permissions(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Storage.contract.Call(opts, out, "permissions")
	return *ret0, err
}

// Permissions is a free data retrieval call binding the contract method 0xab8c71c0.
//
// Solidity: function permissions() constant returns(count int256)
func (_Storage *StorageSession) Permissions() (*big.Int, error) {
	return _Storage.Contract.Permissions(&_Storage.CallOpts)
}

// Permissions is a free data retrieval call binding the contract method 0xab8c71c0.
//
// Solidity: function permissions() constant returns(count int256)
func (_Storage *StorageCallerSession) Permissions() (*big.Int, error) {
	return _Storage.Contract.Permissions(&_Storage.CallOpts)
}

// DeleteAddress is a paid mutator transaction binding the contract method 0x0e14a376.
//
// Solidity: function deleteAddress(k bytes32) returns()
func (_Storage *StorageTransactor) DeleteAddress(opts *bind.TransactOpts, k [32]byte) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "deleteAddress", k)
}

// DeleteAddress is a paid mutator transaction binding the contract method 0x0e14a376.
//
// Solidity: function deleteAddress(k bytes32) returns()
func (_Storage *StorageSession) DeleteAddress(k [32]byte) (*types.Transaction, error) {
	return _Storage.Contract.DeleteAddress(&_Storage.TransactOpts, k)
}

// DeleteAddress is a paid mutator transaction binding the contract method 0x0e14a376.
//
// Solidity: function deleteAddress(k bytes32) returns()
func (_Storage *StorageTransactorSession) DeleteAddress(k [32]byte) (*types.Transaction, error) {
	return _Storage.Contract.DeleteAddress(&_Storage.TransactOpts, k)
}

// DeleteBool is a paid mutator transaction binding the contract method 0x2c62ff2d.
//
// Solidity: function deleteBool(k bytes32) returns()
func (_Storage *StorageTransactor) DeleteBool(opts *bind.TransactOpts, k [32]byte) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "deleteBool", k)
}

// DeleteBool is a paid mutator transaction binding the contract method 0x2c62ff2d.
//
// Solidity: function deleteBool(k bytes32) returns()
func (_Storage *StorageSession) DeleteBool(k [32]byte) (*types.Transaction, error) {
	return _Storage.Contract.DeleteBool(&_Storage.TransactOpts, k)
}

// DeleteBool is a paid mutator transaction binding the contract method 0x2c62ff2d.
//
// Solidity: function deleteBool(k bytes32) returns()
func (_Storage *StorageTransactorSession) DeleteBool(k [32]byte) (*types.Transaction, error) {
	return _Storage.Contract.DeleteBool(&_Storage.TransactOpts, k)
}

// DeleteBytes is a paid mutator transaction binding the contract method 0x616b59f6.
//
// Solidity: function deleteBytes(k bytes32) returns()
func (_Storage *StorageTransactor) DeleteBytes(opts *bind.TransactOpts, k [32]byte) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "deleteBytes", k)
}

// DeleteBytes is a paid mutator transaction binding the contract method 0x616b59f6.
//
// Solidity: function deleteBytes(k bytes32) returns()
func (_Storage *StorageSession) DeleteBytes(k [32]byte) (*types.Transaction, error) {
	return _Storage.Contract.DeleteBytes(&_Storage.TransactOpts, k)
}

// DeleteBytes is a paid mutator transaction binding the contract method 0x616b59f6.
//
// Solidity: function deleteBytes(k bytes32) returns()
func (_Storage *StorageTransactorSession) DeleteBytes(k [32]byte) (*types.Transaction, error) {
	return _Storage.Contract.DeleteBytes(&_Storage.TransactOpts, k)
}

// DeleteBytes32 is a paid mutator transaction binding the contract method 0x0b9adc57.
//
// Solidity: function deleteBytes32(k bytes32) returns()
func (_Storage *StorageTransactor) DeleteBytes32(opts *bind.TransactOpts, k [32]byte) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "deleteBytes32", k)
}

// DeleteBytes32 is a paid mutator transaction binding the contract method 0x0b9adc57.
//
// Solidity: function deleteBytes32(k bytes32) returns()
func (_Storage *StorageSession) DeleteBytes32(k [32]byte) (*types.Transaction, error) {
	return _Storage.Contract.DeleteBytes32(&_Storage.TransactOpts, k)
}

// DeleteBytes32 is a paid mutator transaction binding the contract method 0x0b9adc57.
//
// Solidity: function deleteBytes32(k bytes32) returns()
func (_Storage *StorageTransactorSession) DeleteBytes32(k [32]byte) (*types.Transaction, error) {
	return _Storage.Contract.DeleteBytes32(&_Storage.TransactOpts, k)
}

// DeleteInt256 is a paid mutator transaction binding the contract method 0x4869ecc9.
//
// Solidity: function deleteInt256(k bytes32) returns()
func (_Storage *StorageTransactor) DeleteInt256(opts *bind.TransactOpts, k [32]byte) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "deleteInt256", k)
}

// DeleteInt256 is a paid mutator transaction binding the contract method 0x4869ecc9.
//
// Solidity: function deleteInt256(k bytes32) returns()
func (_Storage *StorageSession) DeleteInt256(k [32]byte) (*types.Transaction, error) {
	return _Storage.Contract.DeleteInt256(&_Storage.TransactOpts, k)
}

// DeleteInt256 is a paid mutator transaction binding the contract method 0x4869ecc9.
//
// Solidity: function deleteInt256(k bytes32) returns()
func (_Storage *StorageTransactorSession) DeleteInt256(k [32]byte) (*types.Transaction, error) {
	return _Storage.Contract.DeleteInt256(&_Storage.TransactOpts, k)
}

// DeleteString is a paid mutator transaction binding the contract method 0xf6bb3cc4.
//
// Solidity: function deleteString(k bytes32) returns()
func (_Storage *StorageTransactor) DeleteString(opts *bind.TransactOpts, k [32]byte) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "deleteString", k)
}

// DeleteString is a paid mutator transaction binding the contract method 0xf6bb3cc4.
//
// Solidity: function deleteString(k bytes32) returns()
func (_Storage *StorageSession) DeleteString(k [32]byte) (*types.Transaction, error) {
	return _Storage.Contract.DeleteString(&_Storage.TransactOpts, k)
}

// DeleteString is a paid mutator transaction binding the contract method 0xf6bb3cc4.
//
// Solidity: function deleteString(k bytes32) returns()
func (_Storage *StorageTransactorSession) DeleteString(k [32]byte) (*types.Transaction, error) {
	return _Storage.Contract.DeleteString(&_Storage.TransactOpts, k)
}

// DeleteUint256 is a paid mutator transaction binding the contract method 0x51f8e7dd.
//
// Solidity: function deleteUint256(k bytes32) returns()
func (_Storage *StorageTransactor) DeleteUint256(opts *bind.TransactOpts, k [32]byte) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "deleteUint256", k)
}

// DeleteUint256 is a paid mutator transaction binding the contract method 0x51f8e7dd.
//
// Solidity: function deleteUint256(k bytes32) returns()
func (_Storage *StorageSession) DeleteUint256(k [32]byte) (*types.Transaction, error) {
	return _Storage.Contract.DeleteUint256(&_Storage.TransactOpts, k)
}

// DeleteUint256 is a paid mutator transaction binding the contract method 0x51f8e7dd.
//
// Solidity: function deleteUint256(k bytes32) returns()
func (_Storage *StorageTransactorSession) DeleteUint256(k [32]byte) (*types.Transaction, error) {
	return _Storage.Contract.DeleteUint256(&_Storage.TransactOpts, k)
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

// SetAddress is a paid mutator transaction binding the contract method 0xca446dd9.
//
// Solidity: function setAddress(k bytes32, v address) returns()
func (_Storage *StorageTransactor) SetAddress(opts *bind.TransactOpts, k [32]byte, v common.Address) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "setAddress", k, v)
}

// SetAddress is a paid mutator transaction binding the contract method 0xca446dd9.
//
// Solidity: function setAddress(k bytes32, v address) returns()
func (_Storage *StorageSession) SetAddress(k [32]byte, v common.Address) (*types.Transaction, error) {
	return _Storage.Contract.SetAddress(&_Storage.TransactOpts, k, v)
}

// SetAddress is a paid mutator transaction binding the contract method 0xca446dd9.
//
// Solidity: function setAddress(k bytes32, v address) returns()
func (_Storage *StorageTransactorSession) SetAddress(k [32]byte, v common.Address) (*types.Transaction, error) {
	return _Storage.Contract.SetAddress(&_Storage.TransactOpts, k, v)
}

// SetBool is a paid mutator transaction binding the contract method 0xabfdcced.
//
// Solidity: function setBool(k bytes32, v bool) returns()
func (_Storage *StorageTransactor) SetBool(opts *bind.TransactOpts, k [32]byte, v bool) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "setBool", k, v)
}

// SetBool is a paid mutator transaction binding the contract method 0xabfdcced.
//
// Solidity: function setBool(k bytes32, v bool) returns()
func (_Storage *StorageSession) SetBool(k [32]byte, v bool) (*types.Transaction, error) {
	return _Storage.Contract.SetBool(&_Storage.TransactOpts, k, v)
}

// SetBool is a paid mutator transaction binding the contract method 0xabfdcced.
//
// Solidity: function setBool(k bytes32, v bool) returns()
func (_Storage *StorageTransactorSession) SetBool(k [32]byte, v bool) (*types.Transaction, error) {
	return _Storage.Contract.SetBool(&_Storage.TransactOpts, k, v)
}

// SetBytes is a paid mutator transaction binding the contract method 0x2e28d084.
//
// Solidity: function setBytes(k bytes32, v bytes) returns()
func (_Storage *StorageTransactor) SetBytes(opts *bind.TransactOpts, k [32]byte, v []byte) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "setBytes", k, v)
}

// SetBytes is a paid mutator transaction binding the contract method 0x2e28d084.
//
// Solidity: function setBytes(k bytes32, v bytes) returns()
func (_Storage *StorageSession) SetBytes(k [32]byte, v []byte) (*types.Transaction, error) {
	return _Storage.Contract.SetBytes(&_Storage.TransactOpts, k, v)
}

// SetBytes is a paid mutator transaction binding the contract method 0x2e28d084.
//
// Solidity: function setBytes(k bytes32, v bytes) returns()
func (_Storage *StorageTransactorSession) SetBytes(k [32]byte, v []byte) (*types.Transaction, error) {
	return _Storage.Contract.SetBytes(&_Storage.TransactOpts, k, v)
}

// SetBytes32 is a paid mutator transaction binding the contract method 0x4e91db08.
//
// Solidity: function setBytes32(k bytes32, v bytes32) returns()
func (_Storage *StorageTransactor) SetBytes32(opts *bind.TransactOpts, k [32]byte, v [32]byte) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "setBytes32", k, v)
}

// SetBytes32 is a paid mutator transaction binding the contract method 0x4e91db08.
//
// Solidity: function setBytes32(k bytes32, v bytes32) returns()
func (_Storage *StorageSession) SetBytes32(k [32]byte, v [32]byte) (*types.Transaction, error) {
	return _Storage.Contract.SetBytes32(&_Storage.TransactOpts, k, v)
}

// SetBytes32 is a paid mutator transaction binding the contract method 0x4e91db08.
//
// Solidity: function setBytes32(k bytes32, v bytes32) returns()
func (_Storage *StorageTransactorSession) SetBytes32(k [32]byte, v [32]byte) (*types.Transaction, error) {
	return _Storage.Contract.SetBytes32(&_Storage.TransactOpts, k, v)
}

// SetInt256 is a paid mutator transaction binding the contract method 0xf64f41db.
//
// Solidity: function setInt256(k bytes32, v int256) returns()
func (_Storage *StorageTransactor) SetInt256(opts *bind.TransactOpts, k [32]byte, v *big.Int) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "setInt256", k, v)
}

// SetInt256 is a paid mutator transaction binding the contract method 0xf64f41db.
//
// Solidity: function setInt256(k bytes32, v int256) returns()
func (_Storage *StorageSession) SetInt256(k [32]byte, v *big.Int) (*types.Transaction, error) {
	return _Storage.Contract.SetInt256(&_Storage.TransactOpts, k, v)
}

// SetInt256 is a paid mutator transaction binding the contract method 0xf64f41db.
//
// Solidity: function setInt256(k bytes32, v int256) returns()
func (_Storage *StorageTransactorSession) SetInt256(k [32]byte, v *big.Int) (*types.Transaction, error) {
	return _Storage.Contract.SetInt256(&_Storage.TransactOpts, k, v)
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

// SetPermission is a paid mutator transaction binding the contract method 0xec6263c0.
//
// Solidity: function setPermission(addr address, grant bool) returns()
func (_Storage *StorageTransactor) SetPermission(opts *bind.TransactOpts, addr common.Address, grant bool) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "setPermission", addr, grant)
}

// SetPermission is a paid mutator transaction binding the contract method 0xec6263c0.
//
// Solidity: function setPermission(addr address, grant bool) returns()
func (_Storage *StorageSession) SetPermission(addr common.Address, grant bool) (*types.Transaction, error) {
	return _Storage.Contract.SetPermission(&_Storage.TransactOpts, addr, grant)
}

// SetPermission is a paid mutator transaction binding the contract method 0xec6263c0.
//
// Solidity: function setPermission(addr address, grant bool) returns()
func (_Storage *StorageTransactorSession) SetPermission(addr common.Address, grant bool) (*types.Transaction, error) {
	return _Storage.Contract.SetPermission(&_Storage.TransactOpts, addr, grant)
}

// SetString is a paid mutator transaction binding the contract method 0x6e899550.
//
// Solidity: function setString(k bytes32, v string) returns()
func (_Storage *StorageTransactor) SetString(opts *bind.TransactOpts, k [32]byte, v string) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "setString", k, v)
}

// SetString is a paid mutator transaction binding the contract method 0x6e899550.
//
// Solidity: function setString(k bytes32, v string) returns()
func (_Storage *StorageSession) SetString(k [32]byte, v string) (*types.Transaction, error) {
	return _Storage.Contract.SetString(&_Storage.TransactOpts, k, v)
}

// SetString is a paid mutator transaction binding the contract method 0x6e899550.
//
// Solidity: function setString(k bytes32, v string) returns()
func (_Storage *StorageTransactorSession) SetString(k [32]byte, v string) (*types.Transaction, error) {
	return _Storage.Contract.SetString(&_Storage.TransactOpts, k, v)
}

// SetUint256 is a paid mutator transaction binding the contract method 0x4f3029c2.
//
// Solidity: function setUint256(k bytes32, v uint256) returns()
func (_Storage *StorageTransactor) SetUint256(opts *bind.TransactOpts, k [32]byte, v *big.Int) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "setUint256", k, v)
}

// SetUint256 is a paid mutator transaction binding the contract method 0x4f3029c2.
//
// Solidity: function setUint256(k bytes32, v uint256) returns()
func (_Storage *StorageSession) SetUint256(k [32]byte, v *big.Int) (*types.Transaction, error) {
	return _Storage.Contract.SetUint256(&_Storage.TransactOpts, k, v)
}

// SetUint256 is a paid mutator transaction binding the contract method 0x4f3029c2.
//
// Solidity: function setUint256(k bytes32, v uint256) returns()
func (_Storage *StorageTransactorSession) SetUint256(k [32]byte, v *big.Int) (*types.Transaction, error) {
	return _Storage.Contract.SetUint256(&_Storage.TransactOpts, k, v)
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
