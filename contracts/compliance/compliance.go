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

// ComplianceABI is the input ABI used to generate the binding from.
const ComplianceABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"admin\",\"type\":\"address\"},{\"name\":\"from\",\"type\":\"address\"},{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"tokens\",\"type\":\"uint256\"}],\"name\":\"canOverride\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"initiator\",\"type\":\"address\"},{\"name\":\"from\",\"type\":\"address\"},{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"tokens\",\"type\":\"uint256\"}],\"name\":\"canTransfer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"},{\"name\":\"freeze\",\"type\":\"bool\"}],\"name\":\"setFrozen\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"kind\",\"type\":\"uint8\"}],\"name\":\"getRules\",\"outputs\":[{\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"kind\",\"type\":\"uint8\"},{\"name\":\"rules\",\"type\":\"address[]\"}],\"name\":\"setRules\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"issuer\",\"type\":\"address\"},{\"name\":\"from\",\"type\":\"address\"},{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"tokens\",\"type\":\"uint256\"}],\"name\":\"canIssue\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"addr\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"isFrozen\",\"type\":\"bool\"},{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"AddressFrozen\",\"type\":\"event\"}]"

// Compliance is an auto generated Go binding around an Ethereum contract.
type Compliance struct {
	ComplianceCaller     // Read-only binding to the contract
	ComplianceTransactor // Write-only binding to the contract
	ComplianceFilterer   // Log filterer for contract events
}

// ComplianceCaller is an auto generated read-only Go binding around an Ethereum contract.
type ComplianceCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ComplianceTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ComplianceTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ComplianceFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ComplianceFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ComplianceSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ComplianceSession struct {
	Contract     *Compliance       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ComplianceCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ComplianceCallerSession struct {
	Contract *ComplianceCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// ComplianceTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ComplianceTransactorSession struct {
	Contract     *ComplianceTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// ComplianceRaw is an auto generated low-level Go binding around an Ethereum contract.
type ComplianceRaw struct {
	Contract *Compliance // Generic contract binding to access the raw methods on
}

// ComplianceCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ComplianceCallerRaw struct {
	Contract *ComplianceCaller // Generic read-only contract binding to access the raw methods on
}

// ComplianceTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ComplianceTransactorRaw struct {
	Contract *ComplianceTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCompliance creates a new instance of Compliance, bound to a specific deployed contract.
func NewCompliance(address common.Address, backend bind.ContractBackend) (*Compliance, error) {
	contract, err := bindCompliance(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Compliance{ComplianceCaller: ComplianceCaller{contract: contract}, ComplianceTransactor: ComplianceTransactor{contract: contract}, ComplianceFilterer: ComplianceFilterer{contract: contract}}, nil
}

// NewComplianceCaller creates a new read-only instance of Compliance, bound to a specific deployed contract.
func NewComplianceCaller(address common.Address, caller bind.ContractCaller) (*ComplianceCaller, error) {
	contract, err := bindCompliance(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ComplianceCaller{contract: contract}, nil
}

// NewComplianceTransactor creates a new write-only instance of Compliance, bound to a specific deployed contract.
func NewComplianceTransactor(address common.Address, transactor bind.ContractTransactor) (*ComplianceTransactor, error) {
	contract, err := bindCompliance(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ComplianceTransactor{contract: contract}, nil
}

// NewComplianceFilterer creates a new log filterer instance of Compliance, bound to a specific deployed contract.
func NewComplianceFilterer(address common.Address, filterer bind.ContractFilterer) (*ComplianceFilterer, error) {
	contract, err := bindCompliance(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ComplianceFilterer{contract: contract}, nil
}

// bindCompliance binds a generic wrapper to an already deployed contract.
func bindCompliance(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ComplianceABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Compliance *ComplianceRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Compliance.Contract.ComplianceCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Compliance *ComplianceRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Compliance.Contract.ComplianceTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Compliance *ComplianceRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Compliance.Contract.ComplianceTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Compliance *ComplianceCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Compliance.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Compliance *ComplianceTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Compliance.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Compliance *ComplianceTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Compliance.Contract.contract.Transact(opts, method, params...)
}

// GetRules is a free data retrieval call binding the contract method 0xafc24bfb.
//
// Solidity: function getRules(kind uint8) constant returns(address[])
func (_Compliance *ComplianceCaller) GetRules(opts *bind.CallOpts, kind uint8) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _Compliance.contract.Call(opts, out, "getRules", kind)
	return *ret0, err
}

// GetRules is a free data retrieval call binding the contract method 0xafc24bfb.
//
// Solidity: function getRules(kind uint8) constant returns(address[])
func (_Compliance *ComplianceSession) GetRules(kind uint8) ([]common.Address, error) {
	return _Compliance.Contract.GetRules(&_Compliance.CallOpts, kind)
}

// GetRules is a free data retrieval call binding the contract method 0xafc24bfb.
//
// Solidity: function getRules(kind uint8) constant returns(address[])
func (_Compliance *ComplianceCallerSession) GetRules(kind uint8) ([]common.Address, error) {
	return _Compliance.Contract.GetRules(&_Compliance.CallOpts, kind)
}

// CanIssue is a paid mutator transaction binding the contract method 0xfd8258bd.
//
// Solidity: function canIssue(issuer address, from address, to address, tokens uint256) returns(bool)
func (_Compliance *ComplianceTransactor) CanIssue(opts *bind.TransactOpts, issuer common.Address, from common.Address, to common.Address, tokens *big.Int) (*types.Transaction, error) {
	return _Compliance.contract.Transact(opts, "canIssue", issuer, from, to, tokens)
}

// CanIssue is a paid mutator transaction binding the contract method 0xfd8258bd.
//
// Solidity: function canIssue(issuer address, from address, to address, tokens uint256) returns(bool)
func (_Compliance *ComplianceSession) CanIssue(issuer common.Address, from common.Address, to common.Address, tokens *big.Int) (*types.Transaction, error) {
	return _Compliance.Contract.CanIssue(&_Compliance.TransactOpts, issuer, from, to, tokens)
}

// CanIssue is a paid mutator transaction binding the contract method 0xfd8258bd.
//
// Solidity: function canIssue(issuer address, from address, to address, tokens uint256) returns(bool)
func (_Compliance *ComplianceTransactorSession) CanIssue(issuer common.Address, from common.Address, to common.Address, tokens *big.Int) (*types.Transaction, error) {
	return _Compliance.Contract.CanIssue(&_Compliance.TransactOpts, issuer, from, to, tokens)
}

// CanOverride is a paid mutator transaction binding the contract method 0x5acba201.
//
// Solidity: function canOverride(admin address, from address, to address, tokens uint256) returns(bool)
func (_Compliance *ComplianceTransactor) CanOverride(opts *bind.TransactOpts, admin common.Address, from common.Address, to common.Address, tokens *big.Int) (*types.Transaction, error) {
	return _Compliance.contract.Transact(opts, "canOverride", admin, from, to, tokens)
}

// CanOverride is a paid mutator transaction binding the contract method 0x5acba201.
//
// Solidity: function canOverride(admin address, from address, to address, tokens uint256) returns(bool)
func (_Compliance *ComplianceSession) CanOverride(admin common.Address, from common.Address, to common.Address, tokens *big.Int) (*types.Transaction, error) {
	return _Compliance.Contract.CanOverride(&_Compliance.TransactOpts, admin, from, to, tokens)
}

// CanOverride is a paid mutator transaction binding the contract method 0x5acba201.
//
// Solidity: function canOverride(admin address, from address, to address, tokens uint256) returns(bool)
func (_Compliance *ComplianceTransactorSession) CanOverride(admin common.Address, from common.Address, to common.Address, tokens *big.Int) (*types.Transaction, error) {
	return _Compliance.Contract.CanOverride(&_Compliance.TransactOpts, admin, from, to, tokens)
}

// CanTransfer is a paid mutator transaction binding the contract method 0x6d62a4fe.
//
// Solidity: function canTransfer(initiator address, from address, to address, tokens uint256) returns(bool)
func (_Compliance *ComplianceTransactor) CanTransfer(opts *bind.TransactOpts, initiator common.Address, from common.Address, to common.Address, tokens *big.Int) (*types.Transaction, error) {
	return _Compliance.contract.Transact(opts, "canTransfer", initiator, from, to, tokens)
}

// CanTransfer is a paid mutator transaction binding the contract method 0x6d62a4fe.
//
// Solidity: function canTransfer(initiator address, from address, to address, tokens uint256) returns(bool)
func (_Compliance *ComplianceSession) CanTransfer(initiator common.Address, from common.Address, to common.Address, tokens *big.Int) (*types.Transaction, error) {
	return _Compliance.Contract.CanTransfer(&_Compliance.TransactOpts, initiator, from, to, tokens)
}

// CanTransfer is a paid mutator transaction binding the contract method 0x6d62a4fe.
//
// Solidity: function canTransfer(initiator address, from address, to address, tokens uint256) returns(bool)
func (_Compliance *ComplianceTransactorSession) CanTransfer(initiator common.Address, from common.Address, to common.Address, tokens *big.Int) (*types.Transaction, error) {
	return _Compliance.Contract.CanTransfer(&_Compliance.TransactOpts, initiator, from, to, tokens)
}

// SetFrozen is a paid mutator transaction binding the contract method 0xac869cd8.
//
// Solidity: function setFrozen(addr address, freeze bool) returns()
func (_Compliance *ComplianceTransactor) SetFrozen(opts *bind.TransactOpts, addr common.Address, freeze bool) (*types.Transaction, error) {
	return _Compliance.contract.Transact(opts, "setFrozen", addr, freeze)
}

// SetFrozen is a paid mutator transaction binding the contract method 0xac869cd8.
//
// Solidity: function setFrozen(addr address, freeze bool) returns()
func (_Compliance *ComplianceSession) SetFrozen(addr common.Address, freeze bool) (*types.Transaction, error) {
	return _Compliance.Contract.SetFrozen(&_Compliance.TransactOpts, addr, freeze)
}

// SetFrozen is a paid mutator transaction binding the contract method 0xac869cd8.
//
// Solidity: function setFrozen(addr address, freeze bool) returns()
func (_Compliance *ComplianceTransactorSession) SetFrozen(addr common.Address, freeze bool) (*types.Transaction, error) {
	return _Compliance.Contract.SetFrozen(&_Compliance.TransactOpts, addr, freeze)
}

// SetRules is a paid mutator transaction binding the contract method 0xd179d77d.
//
// Solidity: function setRules(kind uint8, rules address[]) returns()
func (_Compliance *ComplianceTransactor) SetRules(opts *bind.TransactOpts, kind uint8, rules []common.Address) (*types.Transaction, error) {
	return _Compliance.contract.Transact(opts, "setRules", kind, rules)
}

// SetRules is a paid mutator transaction binding the contract method 0xd179d77d.
//
// Solidity: function setRules(kind uint8, rules address[]) returns()
func (_Compliance *ComplianceSession) SetRules(kind uint8, rules []common.Address) (*types.Transaction, error) {
	return _Compliance.Contract.SetRules(&_Compliance.TransactOpts, kind, rules)
}

// SetRules is a paid mutator transaction binding the contract method 0xd179d77d.
//
// Solidity: function setRules(kind uint8, rules address[]) returns()
func (_Compliance *ComplianceTransactorSession) SetRules(kind uint8, rules []common.Address) (*types.Transaction, error) {
	return _Compliance.Contract.SetRules(&_Compliance.TransactOpts, kind, rules)
}

// ComplianceAddressFrozenIterator is returned from FilterAddressFrozen and is used to iterate over the raw logs and unpacked data for AddressFrozen events raised by the Compliance contract.
type ComplianceAddressFrozenIterator struct {
	Event *ComplianceAddressFrozen // Event containing the contract specifics and raw log

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
func (it *ComplianceAddressFrozenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ComplianceAddressFrozen)
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
		it.Event = new(ComplianceAddressFrozen)
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
func (it *ComplianceAddressFrozenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ComplianceAddressFrozenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ComplianceAddressFrozen represents a AddressFrozen event raised by the Compliance contract.
type ComplianceAddressFrozen struct {
	Addr     common.Address
	IsFrozen bool
	Owner    common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterAddressFrozen is a free log retrieval operation binding the contract event 0x7fa523c84ab8d7fc5b72f08b9e46dbbf10c39e119a075b3e317002d14bc9f436.
//
// Solidity: e AddressFrozen(addr indexed address, isFrozen indexed bool, owner indexed address)
func (_Compliance *ComplianceFilterer) FilterAddressFrozen(opts *bind.FilterOpts, addr []common.Address, isFrozen []bool, owner []common.Address) (*ComplianceAddressFrozenIterator, error) {

	var addrRule []interface{}
	for _, addrItem := range addr {
		addrRule = append(addrRule, addrItem)
	}
	var isFrozenRule []interface{}
	for _, isFrozenItem := range isFrozen {
		isFrozenRule = append(isFrozenRule, isFrozenItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Compliance.contract.FilterLogs(opts, "AddressFrozen", addrRule, isFrozenRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return &ComplianceAddressFrozenIterator{contract: _Compliance.contract, event: "AddressFrozen", logs: logs, sub: sub}, nil
}

// WatchAddressFrozen is a free log subscription operation binding the contract event 0x7fa523c84ab8d7fc5b72f08b9e46dbbf10c39e119a075b3e317002d14bc9f436.
//
// Solidity: e AddressFrozen(addr indexed address, isFrozen indexed bool, owner indexed address)
func (_Compliance *ComplianceFilterer) WatchAddressFrozen(opts *bind.WatchOpts, sink chan<- *ComplianceAddressFrozen, addr []common.Address, isFrozen []bool, owner []common.Address) (event.Subscription, error) {

	var addrRule []interface{}
	for _, addrItem := range addr {
		addrRule = append(addrRule, addrItem)
	}
	var isFrozenRule []interface{}
	for _, isFrozenItem := range isFrozen {
		isFrozenRule = append(isFrozenRule, isFrozenItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Compliance.contract.WatchLogs(opts, "AddressFrozen", addrRule, isFrozenRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ComplianceAddressFrozen)
				if err := _Compliance.contract.UnpackLog(event, "AddressFrozen", log); err != nil {
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
