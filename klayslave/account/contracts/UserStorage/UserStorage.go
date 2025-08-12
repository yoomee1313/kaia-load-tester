// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package userStorageTC

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

// UserStorageTCMetaData contains all meta data concerning the UserStorageTC contract.
var UserStorageTCMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"get\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"getUserData\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"x\",\"type\":\"uint256\"}],\"name\":\"set\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561000f575f80fd5b5061011a8061001d5f395ff3fe6080604052348015600e575f80fd5b5060043610603a575f3560e01c806360fe47b114603e5780636d4ce63c14605c578063ffc9896b14607e575b5f80fd5b605a604936600460a3565b335f90815260208190526040902055565b005b335f908152602081905260409020545b60405190815260200160405180910390f35b606c608936600460b9565b6001600160a01b03165f9081526020819052604090205490565b5f6020828403121560b2575f80fd5b5035919050565b5f6020828403121560c8575f80fd5b81356001600160a01b038116811460dd575f80fd5b939250505056fea2646970667358221220eef08bdea8d0cfec01948a092eac4a88e1c151d1cb0775449f3335f0d288d9d164736f6c63430008180033",
}

// UserStorageTCABI is the input ABI used to generate the binding from.
// Deprecated: Use UserStorageTCMetaData.ABI instead.
var UserStorageTCABI = UserStorageTCMetaData.ABI

// UserStorageTCBinRuntime is the compiled bytecode used for adding genesis block without deploying code.
const UserStorageTCBinRuntime = ``

// UserStorageTCBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use UserStorageTCMetaData.Bin instead.
var UserStorageTCBin = UserStorageTCMetaData.Bin

// DeployUserStorageTC deploys a new Kaia contract, binding an instance of UserStorageTC to it.
func DeployUserStorageTC(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *UserStorageTC, error) {
	parsed, err := UserStorageTCMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(UserStorageTCBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &UserStorageTC{UserStorageTCCaller: UserStorageTCCaller{contract: contract}, UserStorageTCTransactor: UserStorageTCTransactor{contract: contract}, UserStorageTCFilterer: UserStorageTCFilterer{contract: contract}}, nil
}

// UserStorageTC is an auto generated Go binding around a Kaia contract.
type UserStorageTC struct {
	UserStorageTCCaller     // Read-only binding to the contract
	UserStorageTCTransactor // Write-only binding to the contract
	UserStorageTCFilterer   // Log filterer for contract events
}

// UserStorageTCCaller is an auto generated read-only Go binding around a Kaia contract.
type UserStorageTCCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UserStorageTCTransactor is an auto generated write-only Go binding around a Kaia contract.
type UserStorageTCTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UserStorageTCFilterer is an auto generated log filtering Go binding around a Kaia contract events.
type UserStorageTCFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UserStorageTCSession is an auto generated Go binding around a Kaia contract,
// with pre-set call and transact options.
type UserStorageTCSession struct {
	Contract     *UserStorageTC    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// UserStorageTCCallerSession is an auto generated read-only Go binding around a Kaia contract,
// with pre-set call options.
type UserStorageTCCallerSession struct {
	Contract *UserStorageTCCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// UserStorageTCTransactorSession is an auto generated write-only Go binding around a Kaia contract,
// with pre-set transact options.
type UserStorageTCTransactorSession struct {
	Contract     *UserStorageTCTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// UserStorageTCRaw is an auto generated low-level Go binding around a Kaia contract.
type UserStorageTCRaw struct {
	Contract *UserStorageTC // Generic contract binding to access the raw methods on
}

// UserStorageTCCallerRaw is an auto generated low-level read-only Go binding around a Kaia contract.
type UserStorageTCCallerRaw struct {
	Contract *UserStorageTCCaller // Generic read-only contract binding to access the raw methods on
}

// UserStorageTCTransactorRaw is an auto generated low-level write-only Go binding around a Kaia contract.
type UserStorageTCTransactorRaw struct {
	Contract *UserStorageTCTransactor // Generic write-only contract binding to access the raw methods on
}

// NewUserStorageTC creates a new instance of UserStorageTC, bound to a specific deployed contract.
func NewUserStorageTC(address common.Address, backend bind.ContractBackend) (*UserStorageTC, error) {
	contract, err := bindUserStorageTC(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &UserStorageTC{UserStorageTCCaller: UserStorageTCCaller{contract: contract}, UserStorageTCTransactor: UserStorageTCTransactor{contract: contract}, UserStorageTCFilterer: UserStorageTCFilterer{contract: contract}}, nil
}

// NewUserStorageTCCaller creates a new read-only instance of UserStorageTC, bound to a specific deployed contract.
func NewUserStorageTCCaller(address common.Address, caller bind.ContractCaller) (*UserStorageTCCaller, error) {
	contract, err := bindUserStorageTC(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &UserStorageTCCaller{contract: contract}, nil
}

// NewUserStorageTCTransactor creates a new write-only instance of UserStorageTC, bound to a specific deployed contract.
func NewUserStorageTCTransactor(address common.Address, transactor bind.ContractTransactor) (*UserStorageTCTransactor, error) {
	contract, err := bindUserStorageTC(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &UserStorageTCTransactor{contract: contract}, nil
}

// NewUserStorageTCFilterer creates a new log filterer instance of UserStorageTC, bound to a specific deployed contract.
func NewUserStorageTCFilterer(address common.Address, filterer bind.ContractFilterer) (*UserStorageTCFilterer, error) {
	contract, err := bindUserStorageTC(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &UserStorageTCFilterer{contract: contract}, nil
}

// bindUserStorageTC binds a generic wrapper to an already deployed contract.
func bindUserStorageTC(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := UserStorageTCMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_UserStorageTC *UserStorageTCRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _UserStorageTC.Contract.UserStorageTCCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_UserStorageTC *UserStorageTCRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UserStorageTC.Contract.UserStorageTCTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_UserStorageTC *UserStorageTCRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _UserStorageTC.Contract.UserStorageTCTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_UserStorageTC *UserStorageTCCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _UserStorageTC.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_UserStorageTC *UserStorageTCTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UserStorageTC.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_UserStorageTC *UserStorageTCTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _UserStorageTC.Contract.contract.Transact(opts, method, params...)
}

// Get is a free data retrieval call binding the contract method 0x6d4ce63c.
//
// Solidity: function get() view returns(uint256)
func (_UserStorageTC *UserStorageTCCaller) Get(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _UserStorageTC.contract.Call(opts, &out, "get")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Get is a free data retrieval call binding the contract method 0x6d4ce63c.
//
// Solidity: function get() view returns(uint256)
func (_UserStorageTC *UserStorageTCSession) Get() (*big.Int, error) {
	return _UserStorageTC.Contract.Get(&_UserStorageTC.CallOpts)
}

// Get is a free data retrieval call binding the contract method 0x6d4ce63c.
//
// Solidity: function get() view returns(uint256)
func (_UserStorageTC *UserStorageTCCallerSession) Get() (*big.Int, error) {
	return _UserStorageTC.Contract.Get(&_UserStorageTC.CallOpts)
}

// GetUserData is a free data retrieval call binding the contract method 0xffc9896b.
//
// Solidity: function getUserData(address user) view returns(uint256)
func (_UserStorageTC *UserStorageTCCaller) GetUserData(opts *bind.CallOpts, user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _UserStorageTC.contract.Call(opts, &out, "getUserData", user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetUserData is a free data retrieval call binding the contract method 0xffc9896b.
//
// Solidity: function getUserData(address user) view returns(uint256)
func (_UserStorageTC *UserStorageTCSession) GetUserData(user common.Address) (*big.Int, error) {
	return _UserStorageTC.Contract.GetUserData(&_UserStorageTC.CallOpts, user)
}

// GetUserData is a free data retrieval call binding the contract method 0xffc9896b.
//
// Solidity: function getUserData(address user) view returns(uint256)
func (_UserStorageTC *UserStorageTCCallerSession) GetUserData(user common.Address) (*big.Int, error) {
	return _UserStorageTC.Contract.GetUserData(&_UserStorageTC.CallOpts, user)
}

// Set is a paid mutator transaction binding the contract method 0x60fe47b1.
//
// Solidity: function set(uint256 x) returns()
func (_UserStorageTC *UserStorageTCTransactor) Set(opts *bind.TransactOpts, x *big.Int) (*types.Transaction, error) {
	return _UserStorageTC.contract.Transact(opts, "set", x)
}

// Set is a paid mutator transaction binding the contract method 0x60fe47b1.
//
// Solidity: function set(uint256 x) returns()
func (_UserStorageTC *UserStorageTCSession) Set(x *big.Int) (*types.Transaction, error) {
	return _UserStorageTC.Contract.Set(&_UserStorageTC.TransactOpts, x)
}

// Set is a paid mutator transaction binding the contract method 0x60fe47b1.
//
// Solidity: function set(uint256 x) returns()
func (_UserStorageTC *UserStorageTCTransactorSession) Set(x *big.Int) (*types.Transaction, error) {
	return _UserStorageTC.Contract.Set(&_UserStorageTC.TransactOpts, x)
}
