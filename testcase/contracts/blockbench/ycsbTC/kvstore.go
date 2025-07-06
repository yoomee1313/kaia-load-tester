// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ycsbTC

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

// YcsbTCMetaData contains all meta data concerning the YcsbTC contract.
var YcsbTCMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"}],\"name\":\"get\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"}],\"name\":\"set\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561000f575f80fd5b506104718061001d5f395ff3fe608060405234801561000f575f80fd5b5060043610610034575f3560e01c8063693ec85e14610038578063e942b51614610061575b5f80fd5b61004b6100463660046101ef565b610076565b604051610058919061024b565b60405180910390f35b61007461006f36600461027d565b610123565b005b60605f8260405161008791906102dd565b908152602001604051809103902080546100a0906102f8565b80601f01602080910402602001604051908101604052809291908181526020018280546100cc906102f8565b80156101175780601f106100ee57610100808354040283529160200191610117565b820191905f5260205f20905b8154815290600101906020018083116100fa57829003601f168201915b50505050509050919050565b805f8360405161013391906102dd565b9081526020016040518091039020908161014d919061037b565b505050565b634e487b7160e01b5f52604160045260245ffd5b5f82601f830112610175575f80fd5b813567ffffffffffffffff8082111561019057610190610152565b604051601f8301601f19908116603f011681019082821181831017156101b8576101b8610152565b816040528381528660208588010111156101d0575f80fd5b836020870160208301375f602085830101528094505050505092915050565b5f602082840312156101ff575f80fd5b813567ffffffffffffffff811115610215575f80fd5b61022184828501610166565b949350505050565b5f5b8381101561024357818101518382015260200161022b565b50505f910152565b602081525f8251806020840152610269816040850160208701610229565b601f01601f19169190910160400192915050565b5f806040838503121561028e575f80fd5b823567ffffffffffffffff808211156102a5575f80fd5b6102b186838701610166565b935060208501359150808211156102c6575f80fd5b506102d385828601610166565b9150509250929050565b5f82516102ee818460208701610229565b9190910192915050565b600181811c9082168061030c57607f821691505b60208210810361032a57634e487b7160e01b5f52602260045260245ffd5b50919050565b601f82111561014d57805f5260205f20601f840160051c810160208510156103555750805b601f840160051c820191505b81811015610374575f8155600101610361565b5050505050565b815167ffffffffffffffff81111561039557610395610152565b6103a9816103a384546102f8565b84610330565b602080601f8311600181146103dc575f84156103c55750858301515b5f19600386901b1c1916600185901b178555610433565b5f85815260208120601f198616915b8281101561040a578886015182559484019460019091019084016103eb565b508582101561042757878501515f19600388901b60f8161c191681555b505060018460011b0185555b50505050505056fea26469706673582212207c3f1de76836836d9802cd080bcdd65570a124eb74c7fcaeaca54c69042d90eb64736f6c63430008180033",
}

// YcsbTCABI is the input ABI used to generate the binding from.
// Deprecated: Use YcsbTCMetaData.ABI instead.
var YcsbTCABI = YcsbTCMetaData.ABI

// YcsbTCBinRuntime is the compiled bytecode used for adding genesis block without deploying code.
const YcsbTCBinRuntime = ``

// YcsbTCBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use YcsbTCMetaData.Bin instead.
var YcsbTCBin = YcsbTCMetaData.Bin

// DeployYcsbTC deploys a new Kaia contract, binding an instance of YcsbTC to it.
func DeployYcsbTC(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *YcsbTC, error) {
	parsed, err := YcsbTCMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(YcsbTCBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &YcsbTC{YcsbTCCaller: YcsbTCCaller{contract: contract}, YcsbTCTransactor: YcsbTCTransactor{contract: contract}, YcsbTCFilterer: YcsbTCFilterer{contract: contract}}, nil
}

// YcsbTC is an auto generated Go binding around a Kaia contract.
type YcsbTC struct {
	YcsbTCCaller     // Read-only binding to the contract
	YcsbTCTransactor // Write-only binding to the contract
	YcsbTCFilterer   // Log filterer for contract events
}

// YcsbTCCaller is an auto generated read-only Go binding around a Kaia contract.
type YcsbTCCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// YcsbTCTransactor is an auto generated write-only Go binding around a Kaia contract.
type YcsbTCTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// YcsbTCFilterer is an auto generated log filtering Go binding around a Kaia contract events.
type YcsbTCFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// YcsbTCSession is an auto generated Go binding around a Kaia contract,
// with pre-set call and transact options.
type YcsbTCSession struct {
	Contract     *YcsbTC           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// YcsbTCCallerSession is an auto generated read-only Go binding around a Kaia contract,
// with pre-set call options.
type YcsbTCCallerSession struct {
	Contract *YcsbTCCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// YcsbTCTransactorSession is an auto generated write-only Go binding around a Kaia contract,
// with pre-set transact options.
type YcsbTCTransactorSession struct {
	Contract     *YcsbTCTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// YcsbTCRaw is an auto generated low-level Go binding around a Kaia contract.
type YcsbTCRaw struct {
	Contract *YcsbTC // Generic contract binding to access the raw methods on
}

// YcsbTCCallerRaw is an auto generated low-level read-only Go binding around a Kaia contract.
type YcsbTCCallerRaw struct {
	Contract *YcsbTCCaller // Generic read-only contract binding to access the raw methods on
}

// YcsbTCTransactorRaw is an auto generated low-level write-only Go binding around a Kaia contract.
type YcsbTCTransactorRaw struct {
	Contract *YcsbTCTransactor // Generic write-only contract binding to access the raw methods on
}

// NewYcsbTC creates a new instance of YcsbTC, bound to a specific deployed contract.
func NewYcsbTC(address common.Address, backend bind.ContractBackend) (*YcsbTC, error) {
	contract, err := bindYcsbTC(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &YcsbTC{YcsbTCCaller: YcsbTCCaller{contract: contract}, YcsbTCTransactor: YcsbTCTransactor{contract: contract}, YcsbTCFilterer: YcsbTCFilterer{contract: contract}}, nil
}

// NewYcsbTCCaller creates a new read-only instance of YcsbTC, bound to a specific deployed contract.
func NewYcsbTCCaller(address common.Address, caller bind.ContractCaller) (*YcsbTCCaller, error) {
	contract, err := bindYcsbTC(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &YcsbTCCaller{contract: contract}, nil
}

// NewYcsbTCTransactor creates a new write-only instance of YcsbTC, bound to a specific deployed contract.
func NewYcsbTCTransactor(address common.Address, transactor bind.ContractTransactor) (*YcsbTCTransactor, error) {
	contract, err := bindYcsbTC(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &YcsbTCTransactor{contract: contract}, nil
}

// NewYcsbTCFilterer creates a new log filterer instance of YcsbTC, bound to a specific deployed contract.
func NewYcsbTCFilterer(address common.Address, filterer bind.ContractFilterer) (*YcsbTCFilterer, error) {
	contract, err := bindYcsbTC(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &YcsbTCFilterer{contract: contract}, nil
}

// bindYcsbTC binds a generic wrapper to an already deployed contract.
func bindYcsbTC(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := YcsbTCMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_YcsbTC *YcsbTCRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _YcsbTC.Contract.YcsbTCCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_YcsbTC *YcsbTCRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _YcsbTC.Contract.YcsbTCTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_YcsbTC *YcsbTCRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _YcsbTC.Contract.YcsbTCTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_YcsbTC *YcsbTCCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _YcsbTC.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_YcsbTC *YcsbTCTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _YcsbTC.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_YcsbTC *YcsbTCTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _YcsbTC.Contract.contract.Transact(opts, method, params...)
}

// Get is a free data retrieval call binding the contract method 0x693ec85e.
//
// Solidity: function get(string key) view returns(string)
func (_YcsbTC *YcsbTCCaller) Get(opts *bind.CallOpts, key string) (string, error) {
	var out []interface{}
	err := _YcsbTC.contract.Call(opts, &out, "get", key)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Get is a free data retrieval call binding the contract method 0x693ec85e.
//
// Solidity: function get(string key) view returns(string)
func (_YcsbTC *YcsbTCSession) Get(key string) (string, error) {
	return _YcsbTC.Contract.Get(&_YcsbTC.CallOpts, key)
}

// Get is a free data retrieval call binding the contract method 0x693ec85e.
//
// Solidity: function get(string key) view returns(string)
func (_YcsbTC *YcsbTCCallerSession) Get(key string) (string, error) {
	return _YcsbTC.Contract.Get(&_YcsbTC.CallOpts, key)
}

// Set is a paid mutator transaction binding the contract method 0xe942b516.
//
// Solidity: function set(string key, string value) returns()
func (_YcsbTC *YcsbTCTransactor) Set(opts *bind.TransactOpts, key string, value string) (*types.Transaction, error) {
	return _YcsbTC.contract.Transact(opts, "set", key, value)
}

// Set is a paid mutator transaction binding the contract method 0xe942b516.
//
// Solidity: function set(string key, string value) returns()
func (_YcsbTC *YcsbTCSession) Set(key string, value string) (*types.Transaction, error) {
	return _YcsbTC.Contract.Set(&_YcsbTC.TransactOpts, key, value)
}

// Set is a paid mutator transaction binding the contract method 0xe942b516.
//
// Solidity: function set(string key, string value) returns()
func (_YcsbTC *YcsbTCTransactorSession) Set(key string, value string) (*types.Transaction, error) {
	return _YcsbTC.Contract.Set(&_YcsbTC.TransactOpts, key, value)
}
