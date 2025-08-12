// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package readApiCallContractTC

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

// ReadApiCallContractTCMetaData contains all meta data concerning the ReadApiCallContractTC contract.
var ReadApiCallContractTCMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"get\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"set\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405260045f553480156012575f80fd5b5060898061001f5f395ff3fe6080604052348015600e575f80fd5b50600436106030575f3560e01c80636d4ce63c146034578063b8e010de146048575b5f80fd5b5f5460405190815260200160405180910390f35b60516008600155565b00fea2646970667358221220df126a7401c0e4325514b30acabd5739aa3044200494e562de86408f9223952f64736f6c63430008180033",
}

// ReadApiCallContractTCABI is the input ABI used to generate the binding from.
// Deprecated: Use ReadApiCallContractTCMetaData.ABI instead.
var ReadApiCallContractTCABI = ReadApiCallContractTCMetaData.ABI

// ReadApiCallContractTCBinRuntime is the compiled bytecode used for adding genesis block without deploying code.
const ReadApiCallContractTCBinRuntime = ``

// ReadApiCallContractTCBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ReadApiCallContractTCMetaData.Bin instead.
var ReadApiCallContractTCBin = ReadApiCallContractTCMetaData.Bin

// DeployReadApiCallContractTC deploys a new Kaia contract, binding an instance of ReadApiCallContractTC to it.
func DeployReadApiCallContractTC(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ReadApiCallContractTC, error) {
	parsed, err := ReadApiCallContractTCMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ReadApiCallContractTCBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ReadApiCallContractTC{ReadApiCallContractTCCaller: ReadApiCallContractTCCaller{contract: contract}, ReadApiCallContractTCTransactor: ReadApiCallContractTCTransactor{contract: contract}, ReadApiCallContractTCFilterer: ReadApiCallContractTCFilterer{contract: contract}}, nil
}

// ReadApiCallContractTC is an auto generated Go binding around a Kaia contract.
type ReadApiCallContractTC struct {
	ReadApiCallContractTCCaller     // Read-only binding to the contract
	ReadApiCallContractTCTransactor // Write-only binding to the contract
	ReadApiCallContractTCFilterer   // Log filterer for contract events
}

// ReadApiCallContractTCCaller is an auto generated read-only Go binding around a Kaia contract.
type ReadApiCallContractTCCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ReadApiCallContractTCTransactor is an auto generated write-only Go binding around a Kaia contract.
type ReadApiCallContractTCTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ReadApiCallContractTCFilterer is an auto generated log filtering Go binding around a Kaia contract events.
type ReadApiCallContractTCFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ReadApiCallContractTCSession is an auto generated Go binding around a Kaia contract,
// with pre-set call and transact options.
type ReadApiCallContractTCSession struct {
	Contract     *ReadApiCallContractTC // Generic contract binding to set the session for
	CallOpts     bind.CallOpts          // Call options to use throughout this session
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// ReadApiCallContractTCCallerSession is an auto generated read-only Go binding around a Kaia contract,
// with pre-set call options.
type ReadApiCallContractTCCallerSession struct {
	Contract *ReadApiCallContractTCCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                // Call options to use throughout this session
}

// ReadApiCallContractTCTransactorSession is an auto generated write-only Go binding around a Kaia contract,
// with pre-set transact options.
type ReadApiCallContractTCTransactorSession struct {
	Contract     *ReadApiCallContractTCTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                // Transaction auth options to use throughout this session
}

// ReadApiCallContractTCRaw is an auto generated low-level Go binding around a Kaia contract.
type ReadApiCallContractTCRaw struct {
	Contract *ReadApiCallContractTC // Generic contract binding to access the raw methods on
}

// ReadApiCallContractTCCallerRaw is an auto generated low-level read-only Go binding around a Kaia contract.
type ReadApiCallContractTCCallerRaw struct {
	Contract *ReadApiCallContractTCCaller // Generic read-only contract binding to access the raw methods on
}

// ReadApiCallContractTCTransactorRaw is an auto generated low-level write-only Go binding around a Kaia contract.
type ReadApiCallContractTCTransactorRaw struct {
	Contract *ReadApiCallContractTCTransactor // Generic write-only contract binding to access the raw methods on
}

// NewReadApiCallContractTC creates a new instance of ReadApiCallContractTC, bound to a specific deployed contract.
func NewReadApiCallContractTC(address common.Address, backend bind.ContractBackend) (*ReadApiCallContractTC, error) {
	contract, err := bindReadApiCallContractTC(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ReadApiCallContractTC{ReadApiCallContractTCCaller: ReadApiCallContractTCCaller{contract: contract}, ReadApiCallContractTCTransactor: ReadApiCallContractTCTransactor{contract: contract}, ReadApiCallContractTCFilterer: ReadApiCallContractTCFilterer{contract: contract}}, nil
}

// NewReadApiCallContractTCCaller creates a new read-only instance of ReadApiCallContractTC, bound to a specific deployed contract.
func NewReadApiCallContractTCCaller(address common.Address, caller bind.ContractCaller) (*ReadApiCallContractTCCaller, error) {
	contract, err := bindReadApiCallContractTC(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ReadApiCallContractTCCaller{contract: contract}, nil
}

// NewReadApiCallContractTCTransactor creates a new write-only instance of ReadApiCallContractTC, bound to a specific deployed contract.
func NewReadApiCallContractTCTransactor(address common.Address, transactor bind.ContractTransactor) (*ReadApiCallContractTCTransactor, error) {
	contract, err := bindReadApiCallContractTC(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ReadApiCallContractTCTransactor{contract: contract}, nil
}

// NewReadApiCallContractTCFilterer creates a new log filterer instance of ReadApiCallContractTC, bound to a specific deployed contract.
func NewReadApiCallContractTCFilterer(address common.Address, filterer bind.ContractFilterer) (*ReadApiCallContractTCFilterer, error) {
	contract, err := bindReadApiCallContractTC(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ReadApiCallContractTCFilterer{contract: contract}, nil
}

// bindReadApiCallContractTC binds a generic wrapper to an already deployed contract.
func bindReadApiCallContractTC(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ReadApiCallContractTCMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ReadApiCallContractTC *ReadApiCallContractTCRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ReadApiCallContractTC.Contract.ReadApiCallContractTCCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ReadApiCallContractTC *ReadApiCallContractTCRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ReadApiCallContractTC.Contract.ReadApiCallContractTCTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ReadApiCallContractTC *ReadApiCallContractTCRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ReadApiCallContractTC.Contract.ReadApiCallContractTCTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ReadApiCallContractTC *ReadApiCallContractTCCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ReadApiCallContractTC.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ReadApiCallContractTC *ReadApiCallContractTCTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ReadApiCallContractTC.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ReadApiCallContractTC *ReadApiCallContractTCTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ReadApiCallContractTC.Contract.contract.Transact(opts, method, params...)
}

// Get is a free data retrieval call binding the contract method 0x6d4ce63c.
//
// Solidity: function get() view returns(uint256)
func (_ReadApiCallContractTC *ReadApiCallContractTCCaller) Get(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ReadApiCallContractTC.contract.Call(opts, &out, "get")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Get is a free data retrieval call binding the contract method 0x6d4ce63c.
//
// Solidity: function get() view returns(uint256)
func (_ReadApiCallContractTC *ReadApiCallContractTCSession) Get() (*big.Int, error) {
	return _ReadApiCallContractTC.Contract.Get(&_ReadApiCallContractTC.CallOpts)
}

// Get is a free data retrieval call binding the contract method 0x6d4ce63c.
//
// Solidity: function get() view returns(uint256)
func (_ReadApiCallContractTC *ReadApiCallContractTCCallerSession) Get() (*big.Int, error) {
	return _ReadApiCallContractTC.Contract.Get(&_ReadApiCallContractTC.CallOpts)
}

// Set is a paid mutator transaction binding the contract method 0xb8e010de.
//
// Solidity: function set() returns()
func (_ReadApiCallContractTC *ReadApiCallContractTCTransactor) Set(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ReadApiCallContractTC.contract.Transact(opts, "set")
}

// Set is a paid mutator transaction binding the contract method 0xb8e010de.
//
// Solidity: function set() returns()
func (_ReadApiCallContractTC *ReadApiCallContractTCSession) Set() (*types.Transaction, error) {
	return _ReadApiCallContractTC.Contract.Set(&_ReadApiCallContractTC.TransactOpts)
}

// Set is a paid mutator transaction binding the contract method 0xb8e010de.
//
// Solidity: function set() returns()
func (_ReadApiCallContractTC *ReadApiCallContractTCTransactorSession) Set() (*types.Transaction, error) {
	return _ReadApiCallContractTC.Contract.Set(&_ReadApiCallContractTC.TransactOpts)
}
