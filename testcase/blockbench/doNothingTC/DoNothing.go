// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package doNothingTC

import (
	"errors"
	"math/big"
	"strings"

	kaia "github.com/kaiachain/kaia"
	"github.com/kaiachain/kaia/accounts/abi"
	"github.com/kaiachain/kaia/accounts/abi/bind"
	"github.com/kaiachain/kaia/blockchain/types"
	"github.com/kaiachain/kaia/common"
	"github.com/kaiachain/kaia/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = kaia.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// DoNothingTCMetaData contains all meta data concerning the DoNothingTC contract.
var DoNothingTCMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"nothing\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x6080604052348015600e575f80fd5b50606280601a5f395ff3fe6080604052348015600e575f80fd5b50600436106026575f3560e01c8063448f30a314602a575b5f80fd5b00fea264697066735822122006f3389fb4fffa31f676fd136ad9df79fc41b1fd1752529457d4106a71a0723964736f6c63430008180033",
}

// DoNothingTCABI is the input ABI used to generate the binding from.
// Deprecated: Use DoNothingTCMetaData.ABI instead.
var DoNothingTCABI = DoNothingTCMetaData.ABI

// DoNothingTCBinRuntime is the compiled bytecode used for adding genesis block without deploying code.
const DoNothingTCBinRuntime = ``

// DoNothingTCBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use DoNothingTCMetaData.Bin instead.
var DoNothingTCBin = DoNothingTCMetaData.Bin

// DeployDoNothingTC deploys a new Kaia contract, binding an instance of DoNothingTC to it.
func DeployDoNothingTC(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *DoNothingTC, error) {
	parsed, err := DoNothingTCMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(DoNothingTCBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &DoNothingTC{DoNothingTCCaller: DoNothingTCCaller{contract: contract}, DoNothingTCTransactor: DoNothingTCTransactor{contract: contract}, DoNothingTCFilterer: DoNothingTCFilterer{contract: contract}}, nil
}

// DoNothingTC is an auto generated Go binding around a Kaia contract.
type DoNothingTC struct {
	DoNothingTCCaller     // Read-only binding to the contract
	DoNothingTCTransactor // Write-only binding to the contract
	DoNothingTCFilterer   // Log filterer for contract events
}

// DoNothingTCCaller is an auto generated read-only Go binding around a Kaia contract.
type DoNothingTCCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DoNothingTCTransactor is an auto generated write-only Go binding around a Kaia contract.
type DoNothingTCTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DoNothingTCFilterer is an auto generated log filtering Go binding around a Kaia contract events.
type DoNothingTCFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DoNothingTCSession is an auto generated Go binding around a Kaia contract,
// with pre-set call and transact options.
type DoNothingTCSession struct {
	Contract     *DoNothingTC      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DoNothingTCCallerSession is an auto generated read-only Go binding around a Kaia contract,
// with pre-set call options.
type DoNothingTCCallerSession struct {
	Contract *DoNothingTCCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// DoNothingTCTransactorSession is an auto generated write-only Go binding around a Kaia contract,
// with pre-set transact options.
type DoNothingTCTransactorSession struct {
	Contract     *DoNothingTCTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// DoNothingTCRaw is an auto generated low-level Go binding around a Kaia contract.
type DoNothingTCRaw struct {
	Contract *DoNothingTC // Generic contract binding to access the raw methods on
}

// DoNothingTCCallerRaw is an auto generated low-level read-only Go binding around a Kaia contract.
type DoNothingTCCallerRaw struct {
	Contract *DoNothingTCCaller // Generic read-only contract binding to access the raw methods on
}

// DoNothingTCTransactorRaw is an auto generated low-level write-only Go binding around a Kaia contract.
type DoNothingTCTransactorRaw struct {
	Contract *DoNothingTCTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDoNothingTC creates a new instance of DoNothingTC, bound to a specific deployed contract.
func NewDoNothingTC(address common.Address, backend bind.ContractBackend) (*DoNothingTC, error) {
	contract, err := bindDoNothingTC(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DoNothingTC{DoNothingTCCaller: DoNothingTCCaller{contract: contract}, DoNothingTCTransactor: DoNothingTCTransactor{contract: contract}, DoNothingTCFilterer: DoNothingTCFilterer{contract: contract}}, nil
}

// NewDoNothingTCCaller creates a new read-only instance of DoNothingTC, bound to a specific deployed contract.
func NewDoNothingTCCaller(address common.Address, caller bind.ContractCaller) (*DoNothingTCCaller, error) {
	contract, err := bindDoNothingTC(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DoNothingTCCaller{contract: contract}, nil
}

// NewDoNothingTCTransactor creates a new write-only instance of DoNothingTC, bound to a specific deployed contract.
func NewDoNothingTCTransactor(address common.Address, transactor bind.ContractTransactor) (*DoNothingTCTransactor, error) {
	contract, err := bindDoNothingTC(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DoNothingTCTransactor{contract: contract}, nil
}

// NewDoNothingTCFilterer creates a new log filterer instance of DoNothingTC, bound to a specific deployed contract.
func NewDoNothingTCFilterer(address common.Address, filterer bind.ContractFilterer) (*DoNothingTCFilterer, error) {
	contract, err := bindDoNothingTC(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DoNothingTCFilterer{contract: contract}, nil
}

// bindDoNothingTC binds a generic wrapper to an already deployed contract.
func bindDoNothingTC(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := DoNothingTCMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DoNothingTC *DoNothingTCRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DoNothingTC.Contract.DoNothingTCCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DoNothingTC *DoNothingTCRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DoNothingTC.Contract.DoNothingTCTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DoNothingTC *DoNothingTCRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DoNothingTC.Contract.DoNothingTCTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DoNothingTC *DoNothingTCCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DoNothingTC.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DoNothingTC *DoNothingTCTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DoNothingTC.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DoNothingTC *DoNothingTCTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DoNothingTC.Contract.contract.Transact(opts, method, params...)
}

// Nothing is a paid mutator transaction binding the contract method 0x448f30a3.
//
// Solidity: function nothing() returns()
func (_DoNothingTC *DoNothingTCTransactor) Nothing(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DoNothingTC.contract.Transact(opts, "nothing")
}

// Nothing is a paid mutator transaction binding the contract method 0x448f30a3.
//
// Solidity: function nothing() returns()
func (_DoNothingTC *DoNothingTCSession) Nothing() (*types.Transaction, error) {
	return _DoNothingTC.Contract.Nothing(&_DoNothingTC.TransactOpts)
}

// Nothing is a paid mutator transaction binding the contract method 0x448f30a3.
//
// Solidity: function nothing() returns()
func (_DoNothingTC *DoNothingTCTransactorSession) Nothing() (*types.Transaction, error) {
	return _DoNothingTC.Contract.Nothing(&_DoNothingTC.TransactOpts)
}
