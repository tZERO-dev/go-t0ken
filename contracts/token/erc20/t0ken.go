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
const T0kenABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"spender\",\"type\":\"address\"},{\"name\":\"tokens\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"issuer\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"locked\",\"type\":\"bool\"}],\"name\":\"setLocked\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"from\",\"type\":\"address\"},{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"tokens\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"isSuperseded\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"cancellations\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"shareholders\",\"outputs\":[{\"name\":\"count\",\"type\":\"int256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"kill\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"issuingFinished\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwner\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newIssuer\",\"type\":\"address\"}],\"name\":\"setIssuer\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"compliance\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"original\",\"type\":\"address\"},{\"name\":\"replacement\",\"type\":\"address\"}],\"name\":\"cancelAndReissue\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"finishIssuing\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"from\",\"type\":\"address\"},{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"tokens\",\"type\":\"uint256\"}],\"name\":\"transferOverride\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"index\",\"type\":\"int256\"}],\"name\":\"holderAt\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isLocked\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"quantity\",\"type\":\"uint256\"}],\"name\":\"issueTokens\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"tokens\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"isHolder\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"addrOwner\",\"type\":\"address\"},{\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"getSuperseded\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newComplianceAddress\",\"type\":\"address\"}],\"name\":\"setCompliance\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"tokenName\",\"type\":\"string\"},{\"name\":\"tokenSymbol\",\"type\":\"string\"},{\"name\":\"tokenDecimals\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"original\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"replacement\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"VerifiedAddressSuperseded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousIssuer\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newIssuer\",\"type\":\"address\"}],\"name\":\"IssuerSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"tokens\",\"type\":\"uint256\"}],\"name\":\"Issue\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"IssueFinished\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"shareholder\",\"type\":\"address\"}],\"name\":\"ShareholderAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"shareholder\",\"type\":\"address\"}],\"name\":\"ShareholderRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"oldOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnerTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"}]"

// T0kenBin is the compiled bytecode used for deploying new contracts.
const T0kenBin = `60806040526008805460a060020a60ff02191690553480156200002157600080fd5b50604051620025d5380380620025d5833981018060405260608110156200004757600080fd5b8101908080516401000000008111156200006057600080fd5b820160208101848111156200007457600080fd5b81516401000000008111828201871017156200008f57600080fd5b50509291906020018051640100000000811115620000ac57600080fd5b82016020810184811115620000c057600080fd5b8151640100000000811182820187101715620000db57600080fd5b50506020918201516000805460a060020a60ff0219600160a060020a03199091163317169055855191945092506200011a91600191908601906200014e565b508151620001309060029060208501906200014e565b506003805460ff191660ff9290921691909117905550620001f39050565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106200019157805160ff1916838001178555620001c1565b82800160010185558215620001c1579182015b82811115620001c1578251825591602001919060010190620001a4565b50620001cf929150620001d3565b5090565b620001f091905b80821115620001cf5760008155600101620001da565b90565b6123d280620002036000396000f3fe608060405234801561001057600080fd5b50600436106101ec576000357c0100000000000000000000000000000000000000000000000000000000900480636290865d1161012157806395d89b41116100bf578063d4d7b19a1161008e578063d4d7b19a14610529578063dd62ed3e1461054f578063e37ccac71461057d578063f8981789146105a3576101ec565b806395d89b41146104d0578063a4e2d634146104d8578063a5820daa146104e0578063a9059cbb146104fd576101ec565b80637e8d1a39116100fb5780637e8d1a391461046d57806380318be8146104755780638082a929146104ab5780638da5cb5b146104c8576101ec565b80636290865d1461041157806370a082311461041957806379f647201461043f576101ec565b8063313ce5671161018e57806341c0e1b51161016857806341c0e1b5146103b55780634ef05a71146103bd5780634fb2e45d146103c557806355cc4e57146103eb576101ec565b8063313ce5671461036957806334a84827146103875780633723bc0e146103ad576101ec565b80631d143848116101ca5780631d143848146102c8578063211e28b6146102ec57806323b872dd1461030d5780632da7293e14610343576101ec565b806306fdde03146101f1578063095ea7b31461026e57806318160ddd146102ae575b600080fd5b6101f96105c9565b6040805160208082528351818301528351919283929083019185019080838360005b8381101561023357818101518382015260200161021b565b50505050905090810190601f1680156102605780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b61029a6004803603604081101561028457600080fd5b50600160a060020a038135169060200135610656565b604080519115158252519081900360200190f35b6102b66107db565b60408051918252519081900360200190f35b6102d06107e1565b60408051600160a060020a039092168252519081900360200190f35b61030b6004803603602081101561030257600080fd5b503515156107f0565b005b61029a6004803603606081101561032357600080fd5b50600160a060020a038135811691602081013590911690604001356108f6565b61029a6004803603602081101561035957600080fd5b5035600160a060020a0316610b3a565b610371610b5d565b6040805160ff9092168252519081900360200190f35b6102d06004803603602081101561039d57600080fd5b5035600160a060020a0316610b66565b6102b6610b81565b61030b610b87565b61029a610bf7565b61030b600480360360208110156103db57600080fd5b5035600160a060020a0316610c18565b61030b6004803603602081101561040157600080fd5b5035600160a060020a0316610d87565b6102d0610ea3565b6102b66004803603602081101561042f57600080fd5b5035600160a060020a0316610eb2565b61030b6004803603604081101561045557600080fd5b50600160a060020a0381358116916020013516610ecd565b61029a61122d565b61029a6004803603606081101561048b57600080fd5b50600160a060020a038135811691602081013590911690604001356113e0565b6102d0600480360360208110156104c157600080fd5b5035611546565b6102d061155f565b6101f961156e565b61029a6115c6565b61029a600480360360208110156104f657600080fd5b50356115e7565b61029a6004803603604081101561051357600080fd5b50600160a060020a03813516906020013561182b565b61029a6004803603602081101561053f57600080fd5b5035600160a060020a0316611a6d565b6102b66004803603604081101561056557600080fd5b50600160a060020a0381358116916020013516611a80565b6102d06004803603602081101561059357600080fd5b5035600160a060020a0316611aab565b61030b600480360360208110156105b957600080fd5b5035600160a060020a0316611b39565b60018054604080516020600284861615610100026000190190941693909304601f8101849004840282018401909252818152929183018282801561064e5780601f106106235761010080835404028352916020019161064e565b820191906000526020600020905b81548152906001019060200180831161063157829003601f168201915b505050505081565b6000805474010000000000000000000000000000000000000000900460ff16156106b45760405160e560020a62461bcd02815260040180806020018281038252602d81526020018061237a602d913960400191505060405180910390fd5b33600081815260096020526040902054600160a060020a031615610722576040805160e560020a62461bcd02815260206004820152601a60248201527f4164647265737320686173206265656e2063616e63656c6c6564000000000000604482015290519081900360640190fd5b61073360043363ffffffff611c2716565b15156107735760405160e560020a62461bcd02815260040180806020018281038252602f8152602001806122fe602f913960400191505060405180910390fd5b336000818152600c60209081526040808320600160a060020a03891680855290835292819020879055805187815290519293927f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925929181900390910190a35060019392505050565b600b5490565b600854600160a060020a031681565b600054600160a060020a03163314610852576040805160e560020a62461bcd02815260206004820152601960248201527f4f776e6572206163636f756e7420697320726571756972656400000000000000604482015290519081900360640190fd5b60005460ff7401000000000000000000000000000000000000000090910416151581151514156108b65760405160e560020a62461bcd0281526004018080602001828103825260288152602001806123526028913960400191505060405180910390fd5b60008054911515740100000000000000000000000000000000000000000274ff000000000000000000000000000000000000000019909216919091179055565b6000805474010000000000000000000000000000000000000000900460ff16156109545760405160e560020a62461bcd02815260040180806020018281038252602d81526020018061237a602d913960400191505060405180910390fd5b600160a060020a03808416600090815260096020526040902054849116156109c6576040805160e560020a62461bcd02815260206004820152601a60248201527f4164647265737320686173206265656e2063616e63656c6c6564000000000000604482015290519081900360640190fd5b600160a060020a0385166000908152600a602052604090205485908490811115610a3a576040805160e560020a62461bcd02815260206004820152601260248201527f496e73756666696369656e742066756e64730000000000000000000000000000604482015290519081900360640190fd5b600160a060020a0387166000908152600c60209081526040808320338452909152902054851115610ab5576040805160e560020a62461bcd02815260206004820152601a60248201527f5472616e73666572206578636565647320616c6c6f77616e6365000000000000604482015290519081900360640190fd5b6000610ac48888886000611c5e565b90508015610b2f57600160a060020a0388166000908152600c60209081526040808320338452909152902054610b00908763ffffffff611e2916565b600160a060020a0389166000908152600c60209081526040808320338452909152902055610b2f888888611e89565b979650505050505050565b600160a060020a038181166000908152600960205260409020541615155b919050565b60035460ff1681565b600960205260009081526040902054600160a060020a031681565b60045481565b600054600160a060020a03163314610be9576040805160e560020a62461bcd02815260206004820152601960248201527f4f776e6572206163636f756e7420697320726571756972656400000000000000604482015290519081900360640190fd5b600054600160a060020a0316ff5b60085474010000000000000000000000000000000000000000900460ff1681565b600054600160a060020a03163314610c7a576040805160e560020a62461bcd02815260206004820152601960248201527f4f776e6572206163636f756e7420697320726571756972656400000000000000604482015290519081900360640190fd5b600054600160a060020a0382811691161415610cca5760405160e560020a62461bcd02815260040180806020018281038252602581526020018061232d6025913960400191505060405180910390fd5b600160a060020a0381161515610d2a576040805160e560020a62461bcd02815260206004820181905260248201527f4e6577204f776e65722063616e6e6f74206265207a65726f2061646472657373604482015290519081900360640190fd5b60008054600160a060020a0383811673ffffffffffffffffffffffffffffffffffffffff19831681178455604051919092169283917f8934ce4adea8d9ce0d714d2c22b86790e41b7731c84b926fbbdc1d40ff6533c99190a35050565b60005474010000000000000000000000000000000000000000900460ff1615610de45760405160e560020a62461bcd02815260040180806020018281038252602d81526020018061237a602d913960400191505060405180910390fd5b600054600160a060020a03163314610e46576040805160e560020a62461bcd02815260206004820152601960248201527f4f776e6572206163636f756e7420697320726571756972656400000000000000604482015290519081900360640190fd5b6008805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0383811691821792839055604051919216907ff7189b85d7899f5a32d733e6584c4f1dcdff0274f09d969d186c1797673ede1190600090a350565b600754600160a060020a031681565b600160a060020a03166000908152600a602052604090205490565b60005474010000000000000000000000000000000000000000900460ff1615610f2a5760405160e560020a62461bcd02815260040180806020018281038252602d81526020018061237a602d913960400191505060405180910390fd5b600854600160a060020a03163314610f8c576040805160e560020a62461bcd02815260206004820152601360248201527f4f6e6c792069737375657220616c6c6f77656400000000000000000000000000604482015290519081900360640190fd5b600160a060020a0380821660009081526009602052604090205482911615610ffe576040805160e560020a62461bcd02815260206004820152601a60248201527f4164647265737320686173206265656e2063616e63656c6c6564000000000000604482015290519081900360640190fd5b61100f60048463ffffffff611c2716565b8015611029575061102760048363ffffffff611c2716565b155b15156110695760405160e560020a62461bcd02815260040180806020018281038252602a8152602001806122d4602a913960400191505060405180910390fd5b600754600160a060020a03161561118857600754600160a060020a038481166000818152600a602090815260408083205481517ffd8258bd00000000000000000000000000000000000000000000000000000000815233600482015260248101959095528886166044860152606485015251939094169363fd8258bd93608480850194929391928390030190829087803b15801561110657600080fd5b505af115801561111a573d6000803e3d6000fd5b505050506040513d602081101561113057600080fd5b50511515611188576040805160e560020a62461bcd02815260206004820152601860248201527f4661696c6564202763616e49737375652720636865636b2e0000000000000000604482015290519081900360640190fd5b61119960048463ffffffff61202b16565b506111ab60048363ffffffff61213016565b50600160a060020a038381166000818152600960209081526040808320805473ffffffffffffffffffffffffffffffffffffffff19169588169586179055600a90915280822080548584528284205583835282905551339392917fb64971100522354f3d25283cb14e2eefcb0dd26a757482ccfe42479d0a68685791a4505050565b6000805474010000000000000000000000000000000000000000900460ff161561128b5760405160e560020a62461bcd02815260040180806020018281038252602d81526020018061237a602d913960400191505060405180910390fd5b600854600160a060020a031633146112ed576040805160e560020a62461bcd02815260206004820152601360248201527f4f6e6c792069737375657220616c6c6f77656400000000000000000000000000604482015290519081900360640190fd5b60085474010000000000000000000000000000000000000000900460ff1615611360576040805160e560020a62461bcd02815260206004820152601b60248201527f49737375696e6720697320616c72656164792066696e69736865640000000000604482015290519081900360640190fd5b6008805474ff00000000000000000000000000000000000000001916740100000000000000000000000000000000000000001790556040517f7bc15082cf539fecd9e12596dc8e4e77c17a81fccc2c267b626af583162232a290600090a15060085474010000000000000000000000000000000000000000900460ff1690565b6000805474010000000000000000000000000000000000000000900460ff161561143e5760405160e560020a62461bcd02815260040180806020018281038252602d81526020018061237a602d913960400191505060405180910390fd5b600160a060020a03808416600090815260096020526040902054849116156114b0576040805160e560020a62461bcd02815260206004820152601a60248201527f4164647265737320686173206265656e2063616e63656c6c6564000000000000604482015290519081900360640190fd5b600160a060020a0385166000908152600a602052604090205485908490811115611524576040805160e560020a62461bcd02815260206004820152601260248201527f496e73756666696369656e742066756e64730000000000000000000000000000604482015290519081900360640190fd5b60006115338888886001611c5e565b90508015610b2f57610b2f888888611e89565b600061155960048363ffffffff6121e416565b92915050565b600054600160a060020a031681565b6002805460408051602060018416156101000260001901909316849004601f8101849004840282018401909252818152929183018282801561064e5780601f106106235761010080835404028352916020019161064e565b60005474010000000000000000000000000000000000000000900460ff1681565b6000805474010000000000000000000000000000000000000000900460ff16156116455760405160e560020a62461bcd02815260040180806020018281038252602d81526020018061237a602d913960400191505060405180910390fd5b600854600160a060020a031633146116a7576040805160e560020a62461bcd02815260206004820152601360248201527f4f6e6c792069737375657220616c6c6f77656400000000000000000000000000604482015290519081900360640190fd5b60085474010000000000000000000000000000000000000000900460ff161561171a576040805160e560020a62461bcd02815260206004820152601b60248201527f49737375696e6720697320616c72656164792066696e69736865640000000000604482015290519081900360640190fd5b600082111561179b57600b54611736908363ffffffff61226f16565b600b55600854600160a060020a03166000908152600a6020526040902054611764908363ffffffff61226f16565b60088054600160a060020a039081166000908152600a602052604090209290925554611799916004911663ffffffff61213016565b505b600854604080518481529051600160a060020a03909216917fc65a3f767206d2fdcede0b094a4840e01c0dd0be1888b5ba800346eaa0123c169181900360200190a2600854604080518481529051600160a060020a03909216916000917fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef919081900360200190a3506001919050565b6000805474010000000000000000000000000000000000000000900460ff16156118895760405160e560020a62461bcd02815260040180806020018281038252602d81526020018061237a602d913960400191505060405180910390fd5b600160a060020a03808416600090815260096020526040902054849116156118fb576040805160e560020a62461bcd02815260206004820152601a60248201527f4164647265737320686173206265656e2063616e63656c6c6564000000000000604482015290519081900360640190fd5b336000818152600a60205260409020548490811115611964576040805160e560020a62461bcd02815260206004820152601260248201527f496e73756666696369656e742066756e64730000000000000000000000000000604482015290519081900360640190fd5b600854600090600160a060020a0316331415611a425750600754600160a060020a03161580611a3d57600754600854604080517ffd8258bd000000000000000000000000000000000000000000000000000000008152600160a060020a039283166004820181905260248201528a83166044820152606481018a90529051919092169163fd8258bd9160848083019260209291908290030181600087803b158015611a0e57600080fd5b505af1158015611a22573d6000803e3d6000fd5b505050506040513d6020811015611a3857600080fd5b505190505b611a52565b611a4f3388886000611c5e565b90505b8015611a6357611a63338888611e89565b9695505050505050565b600061155960048363ffffffff611c2716565b600160a060020a039182166000908152600c6020908152604080832093909416825291909152205490565b6000600160a060020a0382161515611b0d576040805160e560020a62461bcd02815260206004820152601960248201527f4e6f6e2d7a65726f206164647265737320726571756972656400000000000000604482015290519081900360640190fd5b600160a060020a0380831660009081526009602052604090205416801515611559576000915050610b58565b60005474010000000000000000000000000000000000000000900460ff1615611b965760405160e560020a62461bcd02815260040180806020018281038252602d81526020018061237a602d913960400191505060405180910390fd5b600054600160a060020a03163314611bf8576040805160e560020a62461bcd02815260206004820152601960248201527f4f776e6572206163636f756e7420697320726571756972656400000000000000604482015290519081900360640190fd5b6007805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b600160a060020a038116600090815260018301602052604081205460001901818112801590611c565750835481125b949350505050565b600160a060020a03808416600090815260096020526040812054909185911615611cd2576040805160e560020a62461bcd02815260206004820152601a60248201527f4164647265737320686173206265656e2063616e63656c6c6564000000000000604482015290519081900360640190fd5b600754600160a060020a03161515611ced5782159150611e20565b8215611da357600754604080517f5acba201000000000000000000000000000000000000000000000000000000008152336004820152600160a060020a03898116602483015288811660448301526064820188905291519190921691635acba2019160848083019260209291908290030181600087803b158015611d7057600080fd5b505af1158015611d84573d6000803e3d6000fd5b505050506040513d6020811015611d9a57600080fd5b50519150611e20565b600754604080517f6d62a4fe000000000000000000000000000000000000000000000000000000008152336004820152600160a060020a03898116602483015288811660448301526064820188905291519190921691636d62a4fe9160848083019260209291908290030181600087803b158015611d7057600080fd5b50949350505050565b600082821115611e83576040805160e560020a62461bcd02815260206004820152601460248201527f526573756c747320696e20756e646572666c6f77000000000000000000000000604482015290519081900360640190fd5b50900390565b600160a060020a0383166000908152600a6020526040902054611eb2908263ffffffff611e2916565b600160a060020a038085166000908152600a60205260408082209390935590841681522054611ee7908263ffffffff61226f16565b600160a060020a038084166000818152600a602090815260409182902094909455805185815290519193928716927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef92918290030190a3600160a060020a0382166000908152600a6020526040812054118015611f705750611f7060048363ffffffff61213016565b15611fb25760408051600160a060020a038416815290517f3082c1c4977de80c4f67ee77b56b282610ec93a7ecbcc31b551e0ac2f7bd03959181900360200190a15b600160a060020a0383166000908152600a6020526040902054158015611fe45750611fe460048463ffffffff61202b16565b156120265760408051600160a060020a038516815290517f7ba114ff3d9844510a088eea94cd35562e7c97a2d36a418b37b2e61e5c77affe9181900360200190a15b505050565b600160a060020a038116600090815260018084016020526040822054908112806120555750835481135b15612064576000915050611559565b83548112156120d85783546000908152600285016020818152604080842054600160a060020a031680855260018901835281852086905585855292909152808320805473ffffffffffffffffffffffffffffffffffffffff1990811690931790558654835290912080549091169055612104565b60008181526002850160205260409020805473ffffffffffffffffffffffffffffffffffffffff191690555b5050600160a060020a031660009081526001828101602052604082209190915581546000190190915590565b6000600160a060020a038216151561214a57506000611559565b600160a060020a0382166000908152600184016020526040812054600019019081128015906121795750835481125b15612188576000915050611559565b505081546001908101808455600160a060020a03831660008181528386016020908152604080832085905593825260028701905291909120805473ffffffffffffffffffffffffffffffffffffffff1916909117905592915050565b60008082121580156121f65750825482125b151561224c576040805160e560020a62461bcd02815260206004820152601860248201527f496e646578206f757473696465206f6620626f756e64732e0000000000000000604482015290519081900360640190fd5b5060010160009081526002919091016020526040902054600160a060020a031690565b6000828201838110156122cc576040805160e560020a62461bcd02815260206004820152601360248201527f526573756c747320696e206f766572666c6f7700000000000000000000000000604482015290519081900360640190fd5b939250505056fe4f726967696e616c20646f65736e2774206578697374206f72207265706c6163656d656e7420646f65734d7573742062652061207368617265686f6c64657220746f20617070726f766520746f6b656e207472616e736665724e6577204f776e65722063616e6e6f74206265207468652063757272656e74206f776e6572436f6e747261637420616c726561647920696e20726571756573746564206c6f636b207374617465436f6e74726163742069732063757272656e746c79206c6f636b656420666f72206d6f64696669636174696f6ea165627a7a723058205f11350b5acae2a80a8ff44499e561c81855829bb64776c29305a6763aa6afb90029`

// DeployT0ken deploys a new Ethereum contract, binding an instance of T0ken to it.
func DeployT0ken(auth *bind.TransactOpts, backend bind.ContractBackend, tokenName string, tokenSymbol string, tokenDecimals uint8) (common.Address, *types.Transaction, *T0ken, error) {
	parsed, err := abi.JSON(strings.NewReader(T0kenABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(T0kenBin), backend, tokenName, tokenSymbol, tokenDecimals)
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

// IssuingFinished is a free data retrieval call binding the contract method 0x4ef05a71.
//
// Solidity: function issuingFinished() constant returns(bool)
func (_T0ken *T0kenCaller) IssuingFinished(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _T0ken.contract.Call(opts, out, "issuingFinished")
	return *ret0, err
}

// IssuingFinished is a free data retrieval call binding the contract method 0x4ef05a71.
//
// Solidity: function issuingFinished() constant returns(bool)
func (_T0ken *T0kenSession) IssuingFinished() (bool, error) {
	return _T0ken.Contract.IssuingFinished(&_T0ken.CallOpts)
}

// IssuingFinished is a free data retrieval call binding the contract method 0x4ef05a71.
//
// Solidity: function issuingFinished() constant returns(bool)
func (_T0ken *T0kenCallerSession) IssuingFinished() (bool, error) {
	return _T0ken.Contract.IssuingFinished(&_T0ken.CallOpts)
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

// Shareholders is a free data retrieval call binding the contract method 0x3723bc0e.
//
// Solidity: function shareholders() constant returns(count int256)
func (_T0ken *T0kenCaller) Shareholders(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _T0ken.contract.Call(opts, out, "shareholders")
	return *ret0, err
}

// Shareholders is a free data retrieval call binding the contract method 0x3723bc0e.
//
// Solidity: function shareholders() constant returns(count int256)
func (_T0ken *T0kenSession) Shareholders() (*big.Int, error) {
	return _T0ken.Contract.Shareholders(&_T0ken.CallOpts)
}

// Shareholders is a free data retrieval call binding the contract method 0x3723bc0e.
//
// Solidity: function shareholders() constant returns(count int256)
func (_T0ken *T0kenCallerSession) Shareholders() (*big.Int, error) {
	return _T0ken.Contract.Shareholders(&_T0ken.CallOpts)
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

// FinishIssuing is a paid mutator transaction binding the contract method 0x7e8d1a39.
//
// Solidity: function finishIssuing() returns(bool)
func (_T0ken *T0kenTransactor) FinishIssuing(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _T0ken.contract.Transact(opts, "finishIssuing")
}

// FinishIssuing is a paid mutator transaction binding the contract method 0x7e8d1a39.
//
// Solidity: function finishIssuing() returns(bool)
func (_T0ken *T0kenSession) FinishIssuing() (*types.Transaction, error) {
	return _T0ken.Contract.FinishIssuing(&_T0ken.TransactOpts)
}

// FinishIssuing is a paid mutator transaction binding the contract method 0x7e8d1a39.
//
// Solidity: function finishIssuing() returns(bool)
func (_T0ken *T0kenTransactorSession) FinishIssuing() (*types.Transaction, error) {
	return _T0ken.Contract.FinishIssuing(&_T0ken.TransactOpts)
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

// T0kenIssueIterator is returned from FilterIssue and is used to iterate over the raw logs and unpacked data for Issue events raised by the T0ken contract.
type T0kenIssueIterator struct {
	Event *T0kenIssue // Event containing the contract specifics and raw log

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
func (it *T0kenIssueIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(T0kenIssue)
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
		it.Event = new(T0kenIssue)
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
func (it *T0kenIssueIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *T0kenIssueIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// T0kenIssue represents a Issue event raised by the T0ken contract.
type T0kenIssue struct {
	To     common.Address
	Tokens *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterIssue is a free log retrieval operation binding the contract event 0xc65a3f767206d2fdcede0b094a4840e01c0dd0be1888b5ba800346eaa0123c16.
//
// Solidity: e Issue(to indexed address, tokens uint256)
func (_T0ken *T0kenFilterer) FilterIssue(opts *bind.FilterOpts, to []common.Address) (*T0kenIssueIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _T0ken.contract.FilterLogs(opts, "Issue", toRule)
	if err != nil {
		return nil, err
	}
	return &T0kenIssueIterator{contract: _T0ken.contract, event: "Issue", logs: logs, sub: sub}, nil
}

// WatchIssue is a free log subscription operation binding the contract event 0xc65a3f767206d2fdcede0b094a4840e01c0dd0be1888b5ba800346eaa0123c16.
//
// Solidity: e Issue(to indexed address, tokens uint256)
func (_T0ken *T0kenFilterer) WatchIssue(opts *bind.WatchOpts, sink chan<- *T0kenIssue, to []common.Address) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _T0ken.contract.WatchLogs(opts, "Issue", toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(T0kenIssue)
				if err := _T0ken.contract.UnpackLog(event, "Issue", log); err != nil {
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

// T0kenIssueFinishedIterator is returned from FilterIssueFinished and is used to iterate over the raw logs and unpacked data for IssueFinished events raised by the T0ken contract.
type T0kenIssueFinishedIterator struct {
	Event *T0kenIssueFinished // Event containing the contract specifics and raw log

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
func (it *T0kenIssueFinishedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(T0kenIssueFinished)
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
		it.Event = new(T0kenIssueFinished)
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
func (it *T0kenIssueFinishedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *T0kenIssueFinishedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// T0kenIssueFinished represents a IssueFinished event raised by the T0ken contract.
type T0kenIssueFinished struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterIssueFinished is a free log retrieval operation binding the contract event 0x7bc15082cf539fecd9e12596dc8e4e77c17a81fccc2c267b626af583162232a2.
//
// Solidity: e IssueFinished()
func (_T0ken *T0kenFilterer) FilterIssueFinished(opts *bind.FilterOpts) (*T0kenIssueFinishedIterator, error) {

	logs, sub, err := _T0ken.contract.FilterLogs(opts, "IssueFinished")
	if err != nil {
		return nil, err
	}
	return &T0kenIssueFinishedIterator{contract: _T0ken.contract, event: "IssueFinished", logs: logs, sub: sub}, nil
}

// WatchIssueFinished is a free log subscription operation binding the contract event 0x7bc15082cf539fecd9e12596dc8e4e77c17a81fccc2c267b626af583162232a2.
//
// Solidity: e IssueFinished()
func (_T0ken *T0kenFilterer) WatchIssueFinished(opts *bind.WatchOpts, sink chan<- *T0kenIssueFinished) (event.Subscription, error) {

	logs, sub, err := _T0ken.contract.WatchLogs(opts, "IssueFinished")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(T0kenIssueFinished)
				if err := _T0ken.contract.UnpackLog(event, "IssueFinished", log); err != nil {
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
