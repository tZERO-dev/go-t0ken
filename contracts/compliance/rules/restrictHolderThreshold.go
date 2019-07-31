// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package rules

import (
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// RestrictHolderThresholdABI is the input ABI used to generate the binding from.
const RestrictHolderThresholdABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"token\",\"type\":\"address\"},{\"name\":\"initiator\",\"type\":\"address\"},{\"name\":\"from\",\"type\":\"address\"},{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"toKind\",\"type\":\"uint8\"},{\"name\":\"tokens\",\"type\":\"uint256\"}],\"name\":\"check\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"c\",\"type\":\"address\"},{\"name\":\"symbol\",\"type\":\"string\"},{\"name\":\"limit\",\"type\":\"uint256\"}],\"name\":\"setNonAccreditedInvestorLimit\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"compliance\",\"type\":\"address\"},{\"name\":\"token\",\"type\":\"address\"},{\"name\":\"initiator\",\"type\":\"address\"},{\"name\":\"from\",\"type\":\"address\"},{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"toKind\",\"type\":\"uint8\"},{\"name\":\"tokens\",\"type\":\"uint256\"}],\"name\":\"test\",\"outputs\":[{\"name\":\"s\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"c\",\"type\":\"address\"},{\"name\":\"symbol\",\"type\":\"string\"},{\"name\":\"addr\",\"type\":\"address\"},{\"name\":\"adding\",\"type\":\"bool\"}],\"name\":\"updateNonAccreditedInvestor\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"c\",\"type\":\"address\"},{\"name\":\"symbol\",\"type\":\"string\"},{\"name\":\"limit\",\"type\":\"uint256\"}],\"name\":\"setTotalLimit\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// RestrictHolderThresholdBin is the compiled bytecode used for deploying new contracts.
const RestrictHolderThresholdBin = `60c0604052601960808190527f526573747269637420486f6c646572205468726573686f6c640000000000000060a090815262000040916000919062000055565b503480156200004e57600080fd5b50620000fa565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106200009857805160ff1916838001178555620000c8565b82800160010185558215620000c8579182015b82811115620000c8578251825591602001919060010190620000ab565b50620000d6929150620000da565b5090565b620000f791905b80821115620000d65760008155600101620000e1565b90565b6122d7806200010a6000396000f3fe608060405234801561001057600080fd5b50600436106100725760003560e01c8063a918956211610050578063a9189562146101fd578063aa993c8514610252578063d0453e331461031857610072565b806306fdde03146100775780635a47e1c7146100f457806386b9a0b614610145575b600080fd5b61007f6103d0565b6040805160208082528351818301528351919283929083019185019080838360005b838110156100b95781810151838201526020016100a1565b50505050905090810190601f1680156100e65780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b610143600480360360c081101561010a57600080fd5b506001600160a01b0381358116916020810135821691604082013581169160608101359091169060ff6080820135169060a0013561045e565b005b6101436004803603606081101561015b57600080fd5b6001600160a01b03823516919081019060408101602082013564010000000081111561018657600080fd5b82018360208201111561019857600080fd5b803590602001918460018302840111640100000000831117156101ba57600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295505091359250610fe5915050565b61007f600480360360e081101561021357600080fd5b506001600160a01b0381358116916020810135821691604082013581169160608101358216916080820135169060ff60a0820135169060c00135611187565b6101436004803603608081101561026857600080fd5b6001600160a01b03823516919081019060408101602082013564010000000081111561029357600080fd5b8201836020820111156102a557600080fd5b803590602001918460018302840111640100000000831117156102c757600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550506001600160a01b0383351693505050602001351515611a2e565b6101436004803603606081101561032e57600080fd5b6001600160a01b03823516919081019060408101602082013564010000000081111561035957600080fd5b82018360208201111561036b57600080fd5b8035906020019184600183028401116401000000008311171561038d57600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295505091359250611fe6915050565b6000805460408051602060026001851615610100026000190190941693909304601f810184900484028201840190925281815292918301828280156104565780601f1061042b57610100808354040283529160200191610456565b820191906000526020600020905b81548152906001019060200180831161043957829003601f168201915b505050505081565b336001600160a01b031663975057e76040518163ffffffff1660e01b815260040160206040518083038186803b15801561049757600080fd5b505afa1580156104ab573d6000803e3d6000fd5b505050506040513d60208110156104c157600080fd5b5051604080517f01952ffe00000000000000000000000000000000000000000000000000000000815233600482015290516001600160a01b03909216916301952ffe91602480820192602092909190829003018186803b15801561052457600080fd5b505afa158015610538573d6000803e3d6000fd5b505050506040513d602081101561054e57600080fd5b50516105a1576040805162461bcd60e51b815260206004820152601b60248201527f52657175697265732073746f72616765207065726d697373696f6e0000000000604482015290519081900360640190fd5b856001600160a01b031663d4d7b19a846040518263ffffffff1660e01b815260040180826001600160a01b03166001600160a01b0316815260200191505060206040518083038186803b1580156105f757600080fd5b505afa15801561060b573d6000803e3d6000fd5b505050506040513d602081101561062157600080fd5b50511561062d57610fdd565b60006106376120a5565b90506060876001600160a01b03166395d89b416040518163ffffffff1660e01b815260040160006040518083038186803b15801561067457600080fd5b505afa158015610688573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f1916820160405260208110156106b157600080fd5b8101908080516401000000008111156106c957600080fd5b820160208101848111156106dc57600080fd5b81516401000000008111828201871017156106f657600080fd5b505092919050505090506000826001600160a01b03166333598b006040518060400160405280601d81526020017f5265737472696374486f6c6465725468726573686f6c642e6c696d6974000000815250846040516020018083805190602001908083835b6020831061077a5780518252601f19909201916020918201910161075b565b51815160209384036101000a600019018019909216911617905285519190930192850191508083835b602083106107c25780518252601f1990920191602091820191016107a3565b6001836020036101000a03801982511681845116808217855250505050505090500192505050604051602081830303815290604052805190602001206040518263ffffffff1660e01b81526004018082815260200191505060206040518083038186803b15801561083257600080fd5b505afa158015610846573d6000803e3d6000fd5b505050506040513d602081101561085c57600080fd5b5051604080516320623dc760e21b815290519192506001600160a01b038b1691638188f71c91600480820192602092909190829003018186803b1580156108a257600080fd5b505afa1580156108b6573d6000803e3d6000fd5b505050506040513d60208110156108cc57600080fd5b50518111610921576040805162461bcd60e51b815260206004820152601d60248201527f5368617265686f6c646572207468726573686f6c642072656163686564000000604482015290519081900360640190fd5b6000836001600160a01b0316637ae1cfca60405180606001604052806025815260200161222a60259139858a6040516020018084805190602001908083835b6020831061097f5780518252601f199092019160209182019101610960565b51815160209384036101000a600019018019909216911617905286519190930192860191508083835b602083106109c75780518252601f1990920191602091820191016109a8565b6001836020036101000a038019825116818451168082178552505050505050905001826001600160a01b03166001600160a01b031660601b81526014019350505050604051602081830303815290604052805190602001206040518263ffffffff1660e01b81526004018082815260200191505060206040518083038186803b158015610a5357600080fd5b505afa158015610a67573d6000803e3d6000fd5b505050506040513d6020811015610a7d57600080fd5b5051905080158015610aad5750610aab87610a96612111565b6001600160a01b03169063ffffffff61214c16565b155b15610e20576000846001600160a01b03166333598b006040518060600160405280602a815260200161224f602a9139866040516020018083805190602001908083835b60208310610b0f5780518252601f199092019160209182019101610af0565b51815160209384036101000a600019018019909216911617905285519190930192850191508083835b60208310610b575780518252601f199092019160209182019101610b38565b6001836020036101000a03801982511681845116808217855250505050505090500192505050604051602081830303815290604052805190602001206040518263ffffffff1660e01b81526004018082815260200191505060206040518083038186803b158015610bc757600080fd5b505afa158015610bdb573d6000803e3d6000fd5b505050506040513d6020811015610bf157600080fd5b50516040805160608101909152602a8082529192506001600160a01b038716916333598b0091906122796020830139866040516020018083805190602001908083835b60208310610c535780518252601f199092019160209182019101610c34565b51815160209384036101000a600019018019909216911617905285519190930192850191508083835b60208310610c9b5780518252601f199092019160209182019101610c7c565b6001836020036101000a03801982511681845116808217855250505050505090500192505050604051602081830303815290604052805190602001206040518263ffffffff1660e01b81526004018082815260200191505060206040518083038186803b158015610d0b57600080fd5b505afa158015610d1f573d6000803e3d6000fd5b505050506040513d6020811015610d3557600080fd5b50519250808311610d775760405162461bcd60e51b815260040180806020018281038252602c8152602001806121fe602c913960400191505060405180910390fd5b610d8433858a6001611a2e565b858b6001600160a01b03166370a082318b6040518263ffffffff1660e01b815260040180826001600160a01b03166001600160a01b0316815260200191505060206040518083038186803b158015610ddb57600080fd5b505afa158015610def573d6000803e3d6000fd5b505050506040513d6020811015610e0557600080fd5b50511415610e1a57610e1a33858b6000611a2e565b50610fd8565b808015610e3a5750610e3487610a96612111565b15156001145b15610e5157610e4c3384896000611a2e565b610fd8565b836001600160a01b0316637ae1cfca60405180606001604052806025815260200161222a60259139858b6040516020018084805190602001908083835b60208310610ead5780518252601f199092019160209182019101610e8e565b51815160209384036101000a600019018019909216911617905286519190930192860191508083835b60208310610ef55780518252601f199092019160209182019101610ed6565b6001836020036101000a038019825116818451168082178552505050505050905001826001600160a01b03166001600160a01b031660601b81526014019350505050604051602081830303815290604052805190602001206040518263ffffffff1660e01b81526004018082815260200191505060206040518083038186803b158015610f8157600080fd5b505afa158015610f95573d6000803e3d6000fd5b505050506040513d6020811015610fab57600080fd5b50518015610fc65750610fc088610a96612111565b15156001145b15610fd857610fd833848a6000611a2e565b505050505b505050505050565b826001600160a01b031663975057e76040518163ffffffff1660e01b815260040160206040518083038186803b15801561101e57600080fd5b505afa158015611032573d6000803e3d6000fd5b505050506040513d602081101561104857600080fd5b50516040805160608101909152602a8082526001600160a01b0390921691634f3029c291906122796020830139846040516020018083805190602001908083835b602083106110a85780518252601f199092019160209182019101611089565b51815160209384036101000a600019018019909216911617905285519190930192850191508083835b602083106110f05780518252601f1990920191602091820191016110d1565b6001836020036101000a0380198251168184511680821785525050505050509050019250505060405160208183030381529060405280519060200120836040518363ffffffff1660e01b81526004018083815260200182815260200192505050600060405180830381600087803b15801561116a57600080fd5b505af115801561117e573d6000803e3d6000fd5b50505050505050565b6060866001600160a01b031663d4d7b19a856040518263ffffffff1660e01b815260040180826001600160a01b03166001600160a01b0316815260200191505060206040518083038186803b1580156111df57600080fd5b505afa1580156111f3573d6000803e3d6000fd5b505050506040513d602081101561120957600080fd5b50511561121557611a23565b6000886001600160a01b031663975057e76040518163ffffffff1660e01b815260040160206040518083038186803b15801561125057600080fd5b505afa158015611264573d6000803e3d6000fd5b505050506040513d602081101561127a57600080fd5b5051604080517f95d89b4100000000000000000000000000000000000000000000000000000000815290519192506060916001600160a01b038b16916395d89b41916004808301926000929190829003018186803b1580156112db57600080fd5b505afa1580156112ef573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052602081101561131857600080fd5b81019080805164010000000081111561133057600080fd5b8201602081018481111561134357600080fd5b815164010000000081118282018710171561135d57600080fd5b505092919050505090506000826001600160a01b03166333598b006040518060400160405280601d81526020017f5265737472696374486f6c6465725468726573686f6c642e6c696d6974000000815250846040516020018083805190602001908083835b602083106113e15780518252601f1990920191602091820191016113c2565b51815160209384036101000a600019018019909216911617905285519190930192850191508083835b602083106114295780518252601f19909201916020918201910161140a565b6001836020036101000a03801982511681845116808217855250505050505090500192505050604051602081830303815290604052805190602001206040518263ffffffff1660e01b81526004018082815260200191505060206040518083038186803b15801561149957600080fd5b505afa1580156114ad573d6000803e3d6000fd5b505050506040513d60208110156114c357600080fd5b5051604080516320623dc760e21b815290519192506001600160a01b038c1691638188f71c91600480820192602092909190829003018186803b15801561150957600080fd5b505afa15801561151d573d6000803e3d6000fd5b505050506040513d602081101561153357600080fd5b5051811161157b576040518060400160405280601d81526020017f5368617265686f6c646572207468726573686f6c6420726561636865640000008152509350505050611a23565b6000836001600160a01b0316637ae1cfca60405180606001604052806025815260200161222a60259139858b6040516020018084805190602001908083835b602083106115d95780518252601f1990920191602091820191016115ba565b51815160209384036101000a600019018019909216911617905286519190930192860191508083835b602083106116215780518252601f199092019160209182019101611602565b6001836020036101000a038019825116818451168082178552505050505050905001826001600160a01b03166001600160a01b031660601b81526014019350505050604051602081830303815290604052805190602001206040518263ffffffff1660e01b81526004018082815260200191505060206040518083038186803b1580156116ad57600080fd5b505afa1580156116c1573d6000803e3d6000fd5b505050506040513d60208110156116d757600080fd5b50519050801580156117645750611762888d6001600160a01b0316637b1039996040518163ffffffff1660e01b815260040160206040518083038186803b15801561172157600080fd5b505afa158015611735573d6000803e3d6000fd5b505050506040513d602081101561174b57600080fd5b50516001600160a01b03169063ffffffff61214c16565b155b15611a1e576000846001600160a01b03166333598b006040518060600160405280602a815260200161224f602a9139866040516020018083805190602001908083835b602083106117c65780518252601f1990920191602091820191016117a7565b51815160209384036101000a600019018019909216911617905285519190930192850191508083835b6020831061180e5780518252601f1990920191602091820191016117ef565b6001836020036101000a03801982511681845116808217855250505050505090500192505050604051602081830303815290604052805190602001206040518263ffffffff1660e01b81526004018082815260200191505060206040518083038186803b15801561187e57600080fd5b505afa158015611892573d6000803e3d6000fd5b505050506040513d60208110156118a857600080fd5b50516040805160608101909152602a8082529192506001600160a01b038716916333598b0091906122796020830139866040516020018083805190602001908083835b6020831061190a5780518252601f1990920191602091820191016118eb565b51815160209384036101000a600019018019909216911617905285519190930192850191508083835b602083106119525780518252601f199092019160209182019101611933565b6001836020036101000a03801982511681845116808217855250505050505090500192505050604051602081830303815290604052805190602001206040518263ffffffff1660e01b81526004018082815260200191505060206040518083038186803b1580156119c257600080fd5b505afa1580156119d6573d6000803e3d6000fd5b505050506040513d60208110156119ec57600080fd5b50519250808311611a1c576040518060600160405280602c81526020016121fe602c913995505050505050611a23565b505b505050505b979650505050505050565b6000846001600160a01b031663975057e76040518163ffffffff1660e01b815260040160206040518083038186803b158015611a6957600080fd5b505afa158015611a7d573d6000803e3d6000fd5b505050506040513d6020811015611a9357600080fd5b505160408051606081019091526025808252919250831515916001600160a01b03841691637ae1cfca9161222a602083013987876040516020018084805190602001908083835b60208310611af95780518252601f199092019160209182019101611ada565b51815160209384036101000a600019018019909216911617905286519190930192860191508083835b60208310611b415780518252601f199092019160209182019101611b22565b6001836020036101000a038019825116818451168082178552505050505050905001826001600160a01b03166001600160a01b031660601b81526014019350505050604051602081830303815290604052805190602001206040518263ffffffff1660e01b81526004018082815260200191505060206040518083038186803b158015611bcd57600080fd5b505afa158015611be1573d6000803e3d6000fd5b505050506040513d6020811015611bf757600080fd5b5051151514611fdf576000816001600160a01b03166333598b006040518060600160405280602a815260200161224f602a9139876040516020018083805190602001908083835b60208310611c5d5780518252601f199092019160209182019101611c3e565b51815160209384036101000a600019018019909216911617905285519190930192850191508083835b60208310611ca55780518252601f199092019160209182019101611c86565b6001836020036101000a03801982511681845116808217855250505050505090500192505050604051602081830303815290604052805190602001206040518263ffffffff1660e01b81526004018082815260200191505060206040518083038186803b158015611d1557600080fd5b505afa158015611d29573d6000803e3d6000fd5b505050506040513d6020811015611d3f57600080fd5b5051905082611d515760018103611d56565b806001015b9050816001600160a01b0316634f3029c26040518060600160405280602a815260200161224f602a9139876040516020018083805190602001908083835b60208310611db35780518252601f199092019160209182019101611d94565b51815160209384036101000a600019018019909216911617905285519190930192850191508083835b60208310611dfb5780518252601f199092019160209182019101611ddc565b6001836020036101000a0380198251168184511680821785525050505050509050019250505060405160208183030381529060405280519060200120836040518363ffffffff1660e01b81526004018083815260200182815260200192505050600060405180830381600087803b158015611e7557600080fd5b505af1158015611e89573d6000803e3d6000fd5b50505050816001600160a01b031663abfdcced60405180606001604052806025815260200161222a6025913987876040516020018084805190602001908083835b60208310611ee95780518252601f199092019160209182019101611eca565b51815160209384036101000a600019018019909216911617905286519190930192860191508083835b60208310611f315780518252601f199092019160209182019101611f12565b6001836020036101000a038019825116818451168082178552505050505050905001826001600160a01b03166001600160a01b031660601b8152601401935050505060405160208183030381529060405280519060200120856040518363ffffffff1660e01b8152600401808381526020018215151515815260200192505050600060405180830381600087803b158015611fcb57600080fd5b505af1158015610fd8573d6000803e3d6000fd5b5050505050565b826001600160a01b031663975057e76040518163ffffffff1660e01b815260040160206040518083038186803b15801561201f57600080fd5b505afa158015612033573d6000803e3d6000fd5b505050506040513d602081101561204957600080fd5b5051604080518082018252601d8082527f5265737472696374486f6c6465725468726573686f6c642e6c696d6974000000602080840191825293516001600160a01b0390951694634f3029c294889391019182918083836110a8565b6000336001600160a01b031663975057e76040518163ffffffff1660e01b815260040160206040518083038186803b1580156120e057600080fd5b505afa1580156120f4573d6000803e3d6000fd5b505050506040513d602081101561210a57600080fd5b5051905090565b6000336001600160a01b0316637b1039996040518163ffffffff1660e01b815260040160206040518083038186803b1580156120e057600080fd5b604080517f572bc6b10000000000000000000000000000000000000000000000000000000081526001600160a01b038381166004830152600060248301819052925183929186169163572bc6b1916044808301926020929190829003018186803b1580156121b957600080fd5b505afa1580156121cd573d6000803e3d6000fd5b505050506040513d60208110156121e357600080fd5b50514260109190911c65ffffffffffff161194935050505056fe4e6f6e2d61636372656469746564207368617265686f6c646572207468726573686f6c6420726561636865645265737472696374486f6c6465725468726573686f6c642e6e6f6e616363726564697465645265737472696374486f6c6465725468726573686f6c642e6e6f6e61636372656469746564436f756e745265737472696374486f6c6465725468726573686f6c642e6e6f6e616363726564697465644c696d6974a265627a7a7230582080db2f00a3825390c68cf655f20019994d333cd4f9a3c05566128bd483f1731164736f6c634300050a0032`

// DeployRestrictHolderThreshold deploys a new Ethereum contract, binding an instance of RestrictHolderThreshold to it.
func DeployRestrictHolderThreshold(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *RestrictHolderThreshold, error) {
	parsed, err := abi.JSON(strings.NewReader(RestrictHolderThresholdABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(RestrictHolderThresholdBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &RestrictHolderThreshold{RestrictHolderThresholdCaller: RestrictHolderThresholdCaller{contract: contract}, RestrictHolderThresholdTransactor: RestrictHolderThresholdTransactor{contract: contract}, RestrictHolderThresholdFilterer: RestrictHolderThresholdFilterer{contract: contract}}, nil
}

// RestrictHolderThreshold is an auto generated Go binding around an Ethereum contract.
type RestrictHolderThreshold struct {
	RestrictHolderThresholdCaller     // Read-only binding to the contract
	RestrictHolderThresholdTransactor // Write-only binding to the contract
	RestrictHolderThresholdFilterer   // Log filterer for contract events
}

// RestrictHolderThresholdCaller is an auto generated read-only Go binding around an Ethereum contract.
type RestrictHolderThresholdCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RestrictHolderThresholdTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RestrictHolderThresholdTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RestrictHolderThresholdFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RestrictHolderThresholdFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RestrictHolderThresholdSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RestrictHolderThresholdSession struct {
	Contract     *RestrictHolderThreshold // Generic contract binding to set the session for
	CallOpts     bind.CallOpts            // Call options to use throughout this session
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// RestrictHolderThresholdCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RestrictHolderThresholdCallerSession struct {
	Contract *RestrictHolderThresholdCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                  // Call options to use throughout this session
}

// RestrictHolderThresholdTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RestrictHolderThresholdTransactorSession struct {
	Contract     *RestrictHolderThresholdTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                  // Transaction auth options to use throughout this session
}

// RestrictHolderThresholdRaw is an auto generated low-level Go binding around an Ethereum contract.
type RestrictHolderThresholdRaw struct {
	Contract *RestrictHolderThreshold // Generic contract binding to access the raw methods on
}

// RestrictHolderThresholdCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RestrictHolderThresholdCallerRaw struct {
	Contract *RestrictHolderThresholdCaller // Generic read-only contract binding to access the raw methods on
}

// RestrictHolderThresholdTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RestrictHolderThresholdTransactorRaw struct {
	Contract *RestrictHolderThresholdTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRestrictHolderThreshold creates a new instance of RestrictHolderThreshold, bound to a specific deployed contract.
func NewRestrictHolderThreshold(address common.Address, backend bind.ContractBackend) (*RestrictHolderThreshold, error) {
	contract, err := bindRestrictHolderThreshold(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &RestrictHolderThreshold{RestrictHolderThresholdCaller: RestrictHolderThresholdCaller{contract: contract}, RestrictHolderThresholdTransactor: RestrictHolderThresholdTransactor{contract: contract}, RestrictHolderThresholdFilterer: RestrictHolderThresholdFilterer{contract: contract}}, nil
}

// NewRestrictHolderThresholdCaller creates a new read-only instance of RestrictHolderThreshold, bound to a specific deployed contract.
func NewRestrictHolderThresholdCaller(address common.Address, caller bind.ContractCaller) (*RestrictHolderThresholdCaller, error) {
	contract, err := bindRestrictHolderThreshold(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RestrictHolderThresholdCaller{contract: contract}, nil
}

// NewRestrictHolderThresholdTransactor creates a new write-only instance of RestrictHolderThreshold, bound to a specific deployed contract.
func NewRestrictHolderThresholdTransactor(address common.Address, transactor bind.ContractTransactor) (*RestrictHolderThresholdTransactor, error) {
	contract, err := bindRestrictHolderThreshold(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RestrictHolderThresholdTransactor{contract: contract}, nil
}

// NewRestrictHolderThresholdFilterer creates a new log filterer instance of RestrictHolderThreshold, bound to a specific deployed contract.
func NewRestrictHolderThresholdFilterer(address common.Address, filterer bind.ContractFilterer) (*RestrictHolderThresholdFilterer, error) {
	contract, err := bindRestrictHolderThreshold(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RestrictHolderThresholdFilterer{contract: contract}, nil
}

// bindRestrictHolderThreshold binds a generic wrapper to an already deployed contract.
func bindRestrictHolderThreshold(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(RestrictHolderThresholdABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RestrictHolderThreshold *RestrictHolderThresholdRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _RestrictHolderThreshold.Contract.RestrictHolderThresholdCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RestrictHolderThreshold *RestrictHolderThresholdRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RestrictHolderThreshold.Contract.RestrictHolderThresholdTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RestrictHolderThreshold *RestrictHolderThresholdRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RestrictHolderThreshold.Contract.RestrictHolderThresholdTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RestrictHolderThreshold *RestrictHolderThresholdCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _RestrictHolderThreshold.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RestrictHolderThreshold *RestrictHolderThresholdTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RestrictHolderThreshold.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RestrictHolderThreshold *RestrictHolderThresholdTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RestrictHolderThreshold.Contract.contract.Transact(opts, method, params...)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_RestrictHolderThreshold *RestrictHolderThresholdCaller) Name(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _RestrictHolderThreshold.contract.Call(opts, out, "name")
	return *ret0, err
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_RestrictHolderThreshold *RestrictHolderThresholdSession) Name() (string, error) {
	return _RestrictHolderThreshold.Contract.Name(&_RestrictHolderThreshold.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_RestrictHolderThreshold *RestrictHolderThresholdCallerSession) Name() (string, error) {
	return _RestrictHolderThreshold.Contract.Name(&_RestrictHolderThreshold.CallOpts)
}

// Test is a free data retrieval call binding the contract method 0xa9189562.
//
// Solidity: function test(compliance address, token address, initiator address, from address, to address, toKind uint8, tokens uint256) constant returns(s string)
func (_RestrictHolderThreshold *RestrictHolderThresholdCaller) Test(opts *bind.CallOpts, compliance common.Address, token common.Address, initiator common.Address, from common.Address, to common.Address, toKind uint8, tokens *big.Int) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _RestrictHolderThreshold.contract.Call(opts, out, "test", compliance, token, initiator, from, to, toKind, tokens)
	return *ret0, err
}

// Test is a free data retrieval call binding the contract method 0xa9189562.
//
// Solidity: function test(compliance address, token address, initiator address, from address, to address, toKind uint8, tokens uint256) constant returns(s string)
func (_RestrictHolderThreshold *RestrictHolderThresholdSession) Test(compliance common.Address, token common.Address, initiator common.Address, from common.Address, to common.Address, toKind uint8, tokens *big.Int) (string, error) {
	return _RestrictHolderThreshold.Contract.Test(&_RestrictHolderThreshold.CallOpts, compliance, token, initiator, from, to, toKind, tokens)
}

// Test is a free data retrieval call binding the contract method 0xa9189562.
//
// Solidity: function test(compliance address, token address, initiator address, from address, to address, toKind uint8, tokens uint256) constant returns(s string)
func (_RestrictHolderThreshold *RestrictHolderThresholdCallerSession) Test(compliance common.Address, token common.Address, initiator common.Address, from common.Address, to common.Address, toKind uint8, tokens *big.Int) (string, error) {
	return _RestrictHolderThreshold.Contract.Test(&_RestrictHolderThreshold.CallOpts, compliance, token, initiator, from, to, toKind, tokens)
}

// Check is a paid mutator transaction binding the contract method 0x5a47e1c7.
//
// Solidity: function check(token address, initiator address, from address, to address, toKind uint8, tokens uint256) returns()
func (_RestrictHolderThreshold *RestrictHolderThresholdTransactor) Check(opts *bind.TransactOpts, token common.Address, initiator common.Address, from common.Address, to common.Address, toKind uint8, tokens *big.Int) (*types.Transaction, error) {
	return _RestrictHolderThreshold.contract.Transact(opts, "check", token, initiator, from, to, toKind, tokens)
}

// Check is a paid mutator transaction binding the contract method 0x5a47e1c7.
//
// Solidity: function check(token address, initiator address, from address, to address, toKind uint8, tokens uint256) returns()
func (_RestrictHolderThreshold *RestrictHolderThresholdSession) Check(token common.Address, initiator common.Address, from common.Address, to common.Address, toKind uint8, tokens *big.Int) (*types.Transaction, error) {
	return _RestrictHolderThreshold.Contract.Check(&_RestrictHolderThreshold.TransactOpts, token, initiator, from, to, toKind, tokens)
}

// Check is a paid mutator transaction binding the contract method 0x5a47e1c7.
//
// Solidity: function check(token address, initiator address, from address, to address, toKind uint8, tokens uint256) returns()
func (_RestrictHolderThreshold *RestrictHolderThresholdTransactorSession) Check(token common.Address, initiator common.Address, from common.Address, to common.Address, toKind uint8, tokens *big.Int) (*types.Transaction, error) {
	return _RestrictHolderThreshold.Contract.Check(&_RestrictHolderThreshold.TransactOpts, token, initiator, from, to, toKind, tokens)
}

// SetNonAccreditedInvestorLimit is a paid mutator transaction binding the contract method 0x86b9a0b6.
//
// Solidity: function setNonAccreditedInvestorLimit(c address, symbol string, limit uint256) returns()
func (_RestrictHolderThreshold *RestrictHolderThresholdTransactor) SetNonAccreditedInvestorLimit(opts *bind.TransactOpts, c common.Address, symbol string, limit *big.Int) (*types.Transaction, error) {
	return _RestrictHolderThreshold.contract.Transact(opts, "setNonAccreditedInvestorLimit", c, symbol, limit)
}

// SetNonAccreditedInvestorLimit is a paid mutator transaction binding the contract method 0x86b9a0b6.
//
// Solidity: function setNonAccreditedInvestorLimit(c address, symbol string, limit uint256) returns()
func (_RestrictHolderThreshold *RestrictHolderThresholdSession) SetNonAccreditedInvestorLimit(c common.Address, symbol string, limit *big.Int) (*types.Transaction, error) {
	return _RestrictHolderThreshold.Contract.SetNonAccreditedInvestorLimit(&_RestrictHolderThreshold.TransactOpts, c, symbol, limit)
}

// SetNonAccreditedInvestorLimit is a paid mutator transaction binding the contract method 0x86b9a0b6.
//
// Solidity: function setNonAccreditedInvestorLimit(c address, symbol string, limit uint256) returns()
func (_RestrictHolderThreshold *RestrictHolderThresholdTransactorSession) SetNonAccreditedInvestorLimit(c common.Address, symbol string, limit *big.Int) (*types.Transaction, error) {
	return _RestrictHolderThreshold.Contract.SetNonAccreditedInvestorLimit(&_RestrictHolderThreshold.TransactOpts, c, symbol, limit)
}

// SetTotalLimit is a paid mutator transaction binding the contract method 0xd0453e33.
//
// Solidity: function setTotalLimit(c address, symbol string, limit uint256) returns()
func (_RestrictHolderThreshold *RestrictHolderThresholdTransactor) SetTotalLimit(opts *bind.TransactOpts, c common.Address, symbol string, limit *big.Int) (*types.Transaction, error) {
	return _RestrictHolderThreshold.contract.Transact(opts, "setTotalLimit", c, symbol, limit)
}

// SetTotalLimit is a paid mutator transaction binding the contract method 0xd0453e33.
//
// Solidity: function setTotalLimit(c address, symbol string, limit uint256) returns()
func (_RestrictHolderThreshold *RestrictHolderThresholdSession) SetTotalLimit(c common.Address, symbol string, limit *big.Int) (*types.Transaction, error) {
	return _RestrictHolderThreshold.Contract.SetTotalLimit(&_RestrictHolderThreshold.TransactOpts, c, symbol, limit)
}

// SetTotalLimit is a paid mutator transaction binding the contract method 0xd0453e33.
//
// Solidity: function setTotalLimit(c address, symbol string, limit uint256) returns()
func (_RestrictHolderThreshold *RestrictHolderThresholdTransactorSession) SetTotalLimit(c common.Address, symbol string, limit *big.Int) (*types.Transaction, error) {
	return _RestrictHolderThreshold.Contract.SetTotalLimit(&_RestrictHolderThreshold.TransactOpts, c, symbol, limit)
}

// UpdateNonAccreditedInvestor is a paid mutator transaction binding the contract method 0xaa993c85.
//
// Solidity: function updateNonAccreditedInvestor(c address, symbol string, addr address, adding bool) returns()
func (_RestrictHolderThreshold *RestrictHolderThresholdTransactor) UpdateNonAccreditedInvestor(opts *bind.TransactOpts, c common.Address, symbol string, addr common.Address, adding bool) (*types.Transaction, error) {
	return _RestrictHolderThreshold.contract.Transact(opts, "updateNonAccreditedInvestor", c, symbol, addr, adding)
}

// UpdateNonAccreditedInvestor is a paid mutator transaction binding the contract method 0xaa993c85.
//
// Solidity: function updateNonAccreditedInvestor(c address, symbol string, addr address, adding bool) returns()
func (_RestrictHolderThreshold *RestrictHolderThresholdSession) UpdateNonAccreditedInvestor(c common.Address, symbol string, addr common.Address, adding bool) (*types.Transaction, error) {
	return _RestrictHolderThreshold.Contract.UpdateNonAccreditedInvestor(&_RestrictHolderThreshold.TransactOpts, c, symbol, addr, adding)
}

// UpdateNonAccreditedInvestor is a paid mutator transaction binding the contract method 0xaa993c85.
//
// Solidity: function updateNonAccreditedInvestor(c address, symbol string, addr address, adding bool) returns()
func (_RestrictHolderThreshold *RestrictHolderThresholdTransactorSession) UpdateNonAccreditedInvestor(c common.Address, symbol string, addr common.Address, adding bool) (*types.Transaction, error) {
	return _RestrictHolderThreshold.Contract.UpdateNonAccreditedInvestor(&_RestrictHolderThreshold.TransactOpts, c, symbol, addr, adding)
}
