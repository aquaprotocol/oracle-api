// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contract

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

// ContractABI is the input ABI used to generate the binding from.
const ContractABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"_amount\",\"type\":\"uint256\"},{\"name\":\"_bankAccountNumber\",\"type\":\"string\"}],\"name\":\"invokeFiatWithdraw\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_result\",\"type\":\"string\"},{\"name\":\"_callerAddress\",\"type\":\"address\"},{\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"SetFiatWithdrawResult\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_result\",\"type\":\"string\"},{\"name\":\"_callerAddress\",\"type\":\"address\"},{\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"SetFiatDepositResult\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_amount\",\"type\":\"uint256\"},{\"name\":\"_bankAccountNumber\",\"type\":\"string\"}],\"name\":\"invokeFiatDeposit\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"bankAccountNumber\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"InvokeFiatDepositEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"bankAccountNumber\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"InvokeFiatWithdrawEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_result\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"_callerAddress\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"SetFiatDepositResultEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_result\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"_callerAddress\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"SetFiatWithdrawResultEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"}]"

// Contract is an auto generated Go binding around an Ethereum contract.
type Contract struct {
	ContractCaller     // Read-only binding to the contract
	ContractTransactor // Write-only binding to the contract
	ContractFilterer   // Log filterer for contract events
}

// ContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractSession struct {
	Contract     *Contract         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractCallerSession struct {
	Contract *ContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// ContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractTransactorSession struct {
	Contract     *ContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// ContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractRaw struct {
	Contract *Contract // Generic contract binding to access the raw methods on
}

// ContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractCallerRaw struct {
	Contract *ContractCaller // Generic read-only contract binding to access the raw methods on
}

// ContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractTransactorRaw struct {
	Contract *ContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContract creates a new instance of Contract, bound to a specific deployed contract.
func NewContract(address common.Address, backend bind.ContractBackend) (*Contract, error) {
	contract, err := bindContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Contract{ContractCaller: ContractCaller{contract: contract}, ContractTransactor: ContractTransactor{contract: contract}, ContractFilterer: ContractFilterer{contract: contract}}, nil
}

// NewContractCaller creates a new read-only instance of Contract, bound to a specific deployed contract.
func NewContractCaller(address common.Address, caller bind.ContractCaller) (*ContractCaller, error) {
	contract, err := bindContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractCaller{contract: contract}, nil
}

// NewContractTransactor creates a new write-only instance of Contract, bound to a specific deployed contract.
func NewContractTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractTransactor, error) {
	contract, err := bindContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractTransactor{contract: contract}, nil
}

// NewContractFilterer creates a new log filterer instance of Contract, bound to a specific deployed contract.
func NewContractFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractFilterer, error) {
	contract, err := bindContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractFilterer{contract: contract}, nil
}

// bindContract binds a generic wrapper to an already deployed contract.
func bindContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ContractABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contract *ContractRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Contract.Contract.ContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contract *ContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.Contract.ContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contract *ContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contract.Contract.ContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contract *ContractCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Contract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contract *ContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contract *ContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contract.Contract.contract.Transact(opts, method, params...)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_Contract *ContractCaller) IsOwner(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Contract.contract.Call(opts, out, "isOwner")
	return *ret0, err
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_Contract *ContractSession) IsOwner() (bool, error) {
	return _Contract.Contract.IsOwner(&_Contract.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_Contract *ContractCallerSession) IsOwner() (bool, error) {
	return _Contract.Contract.IsOwner(&_Contract.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Contract *ContractCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Contract.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Contract *ContractSession) Owner() (common.Address, error) {
	return _Contract.Contract.Owner(&_Contract.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Contract *ContractCallerSession) Owner() (common.Address, error) {
	return _Contract.Contract.Owner(&_Contract.CallOpts)
}

// SetFiatDepositResult is a paid mutator transaction binding the contract method 0x72c2cb2f.
//
// Solidity: function SetFiatDepositResult(_result string, _callerAddress address, _id uint256) returns()
func (_Contract *ContractTransactor) SetFiatDepositResult(opts *bind.TransactOpts, _result string, _callerAddress common.Address, _id *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "SetFiatDepositResult", _result, _callerAddress, _id)
}

// SetFiatDepositResult is a paid mutator transaction binding the contract method 0x72c2cb2f.
//
// Solidity: function SetFiatDepositResult(_result string, _callerAddress address, _id uint256) returns()
func (_Contract *ContractSession) SetFiatDepositResult(_result string, _callerAddress common.Address, _id *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.SetFiatDepositResult(&_Contract.TransactOpts, _result, _callerAddress, _id)
}

// SetFiatDepositResult is a paid mutator transaction binding the contract method 0x72c2cb2f.
//
// Solidity: function SetFiatDepositResult(_result string, _callerAddress address, _id uint256) returns()
func (_Contract *ContractTransactorSession) SetFiatDepositResult(_result string, _callerAddress common.Address, _id *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.SetFiatDepositResult(&_Contract.TransactOpts, _result, _callerAddress, _id)
}

// SetFiatWithdrawResult is a paid mutator transaction binding the contract method 0x5aec883f.
//
// Solidity: function SetFiatWithdrawResult(_result string, _callerAddress address, _id uint256) returns()
func (_Contract *ContractTransactor) SetFiatWithdrawResult(opts *bind.TransactOpts, _result string, _callerAddress common.Address, _id *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "SetFiatWithdrawResult", _result, _callerAddress, _id)
}

// SetFiatWithdrawResult is a paid mutator transaction binding the contract method 0x5aec883f.
//
// Solidity: function SetFiatWithdrawResult(_result string, _callerAddress address, _id uint256) returns()
func (_Contract *ContractSession) SetFiatWithdrawResult(_result string, _callerAddress common.Address, _id *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.SetFiatWithdrawResult(&_Contract.TransactOpts, _result, _callerAddress, _id)
}

// SetFiatWithdrawResult is a paid mutator transaction binding the contract method 0x5aec883f.
//
// Solidity: function SetFiatWithdrawResult(_result string, _callerAddress address, _id uint256) returns()
func (_Contract *ContractTransactorSession) SetFiatWithdrawResult(_result string, _callerAddress common.Address, _id *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.SetFiatWithdrawResult(&_Contract.TransactOpts, _result, _callerAddress, _id)
}

// InvokeFiatDeposit is a paid mutator transaction binding the contract method 0xb70bf3b3.
//
// Solidity: function invokeFiatDeposit(_amount uint256, _bankAccountNumber string) returns(uint256)
func (_Contract *ContractTransactor) InvokeFiatDeposit(opts *bind.TransactOpts, _amount *big.Int, _bankAccountNumber string) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "invokeFiatDeposit", _amount, _bankAccountNumber)
}

// InvokeFiatDeposit is a paid mutator transaction binding the contract method 0xb70bf3b3.
//
// Solidity: function invokeFiatDeposit(_amount uint256, _bankAccountNumber string) returns(uint256)
func (_Contract *ContractSession) InvokeFiatDeposit(_amount *big.Int, _bankAccountNumber string) (*types.Transaction, error) {
	return _Contract.Contract.InvokeFiatDeposit(&_Contract.TransactOpts, _amount, _bankAccountNumber)
}

// InvokeFiatDeposit is a paid mutator transaction binding the contract method 0xb70bf3b3.
//
// Solidity: function invokeFiatDeposit(_amount uint256, _bankAccountNumber string) returns(uint256)
func (_Contract *ContractTransactorSession) InvokeFiatDeposit(_amount *big.Int, _bankAccountNumber string) (*types.Transaction, error) {
	return _Contract.Contract.InvokeFiatDeposit(&_Contract.TransactOpts, _amount, _bankAccountNumber)
}

// InvokeFiatWithdraw is a paid mutator transaction binding the contract method 0x0aea1c3d.
//
// Solidity: function invokeFiatWithdraw(_amount uint256, _bankAccountNumber string) returns(uint256)
func (_Contract *ContractTransactor) InvokeFiatWithdraw(opts *bind.TransactOpts, _amount *big.Int, _bankAccountNumber string) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "invokeFiatWithdraw", _amount, _bankAccountNumber)
}

// InvokeFiatWithdraw is a paid mutator transaction binding the contract method 0x0aea1c3d.
//
// Solidity: function invokeFiatWithdraw(_amount uint256, _bankAccountNumber string) returns(uint256)
func (_Contract *ContractSession) InvokeFiatWithdraw(_amount *big.Int, _bankAccountNumber string) (*types.Transaction, error) {
	return _Contract.Contract.InvokeFiatWithdraw(&_Contract.TransactOpts, _amount, _bankAccountNumber)
}

// InvokeFiatWithdraw is a paid mutator transaction binding the contract method 0x0aea1c3d.
//
// Solidity: function invokeFiatWithdraw(_amount uint256, _bankAccountNumber string) returns(uint256)
func (_Contract *ContractTransactorSession) InvokeFiatWithdraw(_amount *big.Int, _bankAccountNumber string) (*types.Transaction, error) {
	return _Contract.Contract.InvokeFiatWithdraw(&_Contract.TransactOpts, _amount, _bankAccountNumber)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Contract *ContractTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Contract *ContractSession) RenounceOwnership() (*types.Transaction, error) {
	return _Contract.Contract.RenounceOwnership(&_Contract.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Contract *ContractTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Contract.Contract.RenounceOwnership(&_Contract.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_Contract *ContractTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_Contract *ContractSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Contract.Contract.TransferOwnership(&_Contract.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_Contract *ContractTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Contract.Contract.TransferOwnership(&_Contract.TransactOpts, newOwner)
}

// ContractInvokeFiatDepositEventIterator is returned from FilterInvokeFiatDepositEvent and is used to iterate over the raw logs and unpacked data for InvokeFiatDepositEvent events raised by the Contract contract.
type ContractInvokeFiatDepositEventIterator struct {
	Event *ContractInvokeFiatDepositEvent // Event containing the contract specifics and raw log

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
func (it *ContractInvokeFiatDepositEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractInvokeFiatDepositEvent)
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
		it.Event = new(ContractInvokeFiatDepositEvent)
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
func (it *ContractInvokeFiatDepositEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractInvokeFiatDepositEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractInvokeFiatDepositEvent represents a InvokeFiatDepositEvent event raised by the Contract contract.
type ContractInvokeFiatDepositEvent struct {
	Amount            *big.Int
	BankAccountNumber string
	Id                *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterInvokeFiatDepositEvent is a free log retrieval operation binding the contract event 0x79376844064f450e7d613f4783a540ff751b80515f9b5fb535bda805e322e6d3.
//
// Solidity: e InvokeFiatDepositEvent(amount uint256, bankAccountNumber string, id uint256)
func (_Contract *ContractFilterer) FilterInvokeFiatDepositEvent(opts *bind.FilterOpts) (*ContractInvokeFiatDepositEventIterator, error) {

	logs, sub, err := _Contract.contract.FilterLogs(opts, "InvokeFiatDepositEvent")
	if err != nil {
		return nil, err
	}
	return &ContractInvokeFiatDepositEventIterator{contract: _Contract.contract, event: "InvokeFiatDepositEvent", logs: logs, sub: sub}, nil
}

// WatchInvokeFiatDepositEvent is a free log subscription operation binding the contract event 0x79376844064f450e7d613f4783a540ff751b80515f9b5fb535bda805e322e6d3.
//
// Solidity: e InvokeFiatDepositEvent(amount uint256, bankAccountNumber string, id uint256)
func (_Contract *ContractFilterer) WatchInvokeFiatDepositEvent(opts *bind.WatchOpts, sink chan<- *ContractInvokeFiatDepositEvent) (event.Subscription, error) {

	logs, sub, err := _Contract.contract.WatchLogs(opts, "InvokeFiatDepositEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractInvokeFiatDepositEvent)
				if err := _Contract.contract.UnpackLog(event, "InvokeFiatDepositEvent", log); err != nil {
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

// ContractInvokeFiatWithdrawEventIterator is returned from FilterInvokeFiatWithdrawEvent and is used to iterate over the raw logs and unpacked data for InvokeFiatWithdrawEvent events raised by the Contract contract.
type ContractInvokeFiatWithdrawEventIterator struct {
	Event *ContractInvokeFiatWithdrawEvent // Event containing the contract specifics and raw log

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
func (it *ContractInvokeFiatWithdrawEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractInvokeFiatWithdrawEvent)
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
		it.Event = new(ContractInvokeFiatWithdrawEvent)
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
func (it *ContractInvokeFiatWithdrawEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractInvokeFiatWithdrawEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractInvokeFiatWithdrawEvent represents a InvokeFiatWithdrawEvent event raised by the Contract contract.
type ContractInvokeFiatWithdrawEvent struct {
	Amount            *big.Int
	BankAccountNumber string
	Id                *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterInvokeFiatWithdrawEvent is a free log retrieval operation binding the contract event 0x571345f952d8dd1161c8a0334ad5d266bff37487e96bb7fed9f953dc33fd6a49.
//
// Solidity: e InvokeFiatWithdrawEvent(amount uint256, bankAccountNumber string, id uint256)
func (_Contract *ContractFilterer) FilterInvokeFiatWithdrawEvent(opts *bind.FilterOpts) (*ContractInvokeFiatWithdrawEventIterator, error) {

	logs, sub, err := _Contract.contract.FilterLogs(opts, "InvokeFiatWithdrawEvent")
	if err != nil {
		return nil, err
	}
	return &ContractInvokeFiatWithdrawEventIterator{contract: _Contract.contract, event: "InvokeFiatWithdrawEvent", logs: logs, sub: sub}, nil
}

// WatchInvokeFiatWithdrawEvent is a free log subscription operation binding the contract event 0x571345f952d8dd1161c8a0334ad5d266bff37487e96bb7fed9f953dc33fd6a49.
//
// Solidity: e InvokeFiatWithdrawEvent(amount uint256, bankAccountNumber string, id uint256)
func (_Contract *ContractFilterer) WatchInvokeFiatWithdrawEvent(opts *bind.WatchOpts, sink chan<- *ContractInvokeFiatWithdrawEvent) (event.Subscription, error) {

	logs, sub, err := _Contract.contract.WatchLogs(opts, "InvokeFiatWithdrawEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractInvokeFiatWithdrawEvent)
				if err := _Contract.contract.UnpackLog(event, "InvokeFiatWithdrawEvent", log); err != nil {
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

// ContractOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Contract contract.
type ContractOwnershipTransferredIterator struct {
	Event *ContractOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *ContractOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractOwnershipTransferred)
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
		it.Event = new(ContractOwnershipTransferred)
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
func (it *ContractOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractOwnershipTransferred represents a OwnershipTransferred event raised by the Contract contract.
type ContractOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: e OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_Contract *ContractFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ContractOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ContractOwnershipTransferredIterator{contract: _Contract.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: e OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_Contract *ContractFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ContractOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractOwnershipTransferred)
				if err := _Contract.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ContractSetFiatDepositResultEventIterator is returned from FilterSetFiatDepositResultEvent and is used to iterate over the raw logs and unpacked data for SetFiatDepositResultEvent events raised by the Contract contract.
type ContractSetFiatDepositResultEventIterator struct {
	Event *ContractSetFiatDepositResultEvent // Event containing the contract specifics and raw log

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
func (it *ContractSetFiatDepositResultEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractSetFiatDepositResultEvent)
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
		it.Event = new(ContractSetFiatDepositResultEvent)
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
func (it *ContractSetFiatDepositResultEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractSetFiatDepositResultEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractSetFiatDepositResultEvent represents a SetFiatDepositResultEvent event raised by the Contract contract.
type ContractSetFiatDepositResultEvent struct {
	Result        string
	CallerAddress common.Address
	Id            *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterSetFiatDepositResultEvent is a free log retrieval operation binding the contract event 0xe5cf4933cdee53a6cf07b8aa5e28c86cdcd310d1b7f0b9d7a02be893296e3cf3.
//
// Solidity: e SetFiatDepositResultEvent(_result string, _callerAddress address, _id uint256)
func (_Contract *ContractFilterer) FilterSetFiatDepositResultEvent(opts *bind.FilterOpts) (*ContractSetFiatDepositResultEventIterator, error) {

	logs, sub, err := _Contract.contract.FilterLogs(opts, "SetFiatDepositResultEvent")
	if err != nil {
		return nil, err
	}
	return &ContractSetFiatDepositResultEventIterator{contract: _Contract.contract, event: "SetFiatDepositResultEvent", logs: logs, sub: sub}, nil
}

// WatchSetFiatDepositResultEvent is a free log subscription operation binding the contract event 0xe5cf4933cdee53a6cf07b8aa5e28c86cdcd310d1b7f0b9d7a02be893296e3cf3.
//
// Solidity: e SetFiatDepositResultEvent(_result string, _callerAddress address, _id uint256)
func (_Contract *ContractFilterer) WatchSetFiatDepositResultEvent(opts *bind.WatchOpts, sink chan<- *ContractSetFiatDepositResultEvent) (event.Subscription, error) {

	logs, sub, err := _Contract.contract.WatchLogs(opts, "SetFiatDepositResultEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractSetFiatDepositResultEvent)
				if err := _Contract.contract.UnpackLog(event, "SetFiatDepositResultEvent", log); err != nil {
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

// ContractSetFiatWithdrawResultEventIterator is returned from FilterSetFiatWithdrawResultEvent and is used to iterate over the raw logs and unpacked data for SetFiatWithdrawResultEvent events raised by the Contract contract.
type ContractSetFiatWithdrawResultEventIterator struct {
	Event *ContractSetFiatWithdrawResultEvent // Event containing the contract specifics and raw log

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
func (it *ContractSetFiatWithdrawResultEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractSetFiatWithdrawResultEvent)
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
		it.Event = new(ContractSetFiatWithdrawResultEvent)
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
func (it *ContractSetFiatWithdrawResultEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractSetFiatWithdrawResultEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractSetFiatWithdrawResultEvent represents a SetFiatWithdrawResultEvent event raised by the Contract contract.
type ContractSetFiatWithdrawResultEvent struct {
	Result        string
	CallerAddress common.Address
	Id            *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterSetFiatWithdrawResultEvent is a free log retrieval operation binding the contract event 0xb7c121bbde1fe097a06d91286f43e0028796ba2fc98525c2eea856125ece6a8a.
//
// Solidity: e SetFiatWithdrawResultEvent(_result string, _callerAddress address, _id uint256)
func (_Contract *ContractFilterer) FilterSetFiatWithdrawResultEvent(opts *bind.FilterOpts) (*ContractSetFiatWithdrawResultEventIterator, error) {

	logs, sub, err := _Contract.contract.FilterLogs(opts, "SetFiatWithdrawResultEvent")
	if err != nil {
		return nil, err
	}
	return &ContractSetFiatWithdrawResultEventIterator{contract: _Contract.contract, event: "SetFiatWithdrawResultEvent", logs: logs, sub: sub}, nil
}

// WatchSetFiatWithdrawResultEvent is a free log subscription operation binding the contract event 0xb7c121bbde1fe097a06d91286f43e0028796ba2fc98525c2eea856125ece6a8a.
//
// Solidity: e SetFiatWithdrawResultEvent(_result string, _callerAddress address, _id uint256)
func (_Contract *ContractFilterer) WatchSetFiatWithdrawResultEvent(opts *bind.WatchOpts, sink chan<- *ContractSetFiatWithdrawResultEvent) (event.Subscription, error) {

	logs, sub, err := _Contract.contract.WatchLogs(opts, "SetFiatWithdrawResultEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractSetFiatWithdrawResultEvent)
				if err := _Contract.contract.UnpackLog(event, "SetFiatWithdrawResultEvent", log); err != nil {
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