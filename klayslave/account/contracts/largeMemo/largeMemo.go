// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package largeMemo

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

// LargeMemoTCMetaData contains all meta data concerning the LargeMemoTC contract.
var LargeMemoTCMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"run\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_str\",\"type\":\"string\"}],\"name\":\"setName\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"str\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561000f575f80fd5b5060408051808201909152600c81526b12195b1b1bcb0815dbdc9b1960a21b60208201525f9061003f90826100dd565b5061019c565b634e487b7160e01b5f52604160045260245ffd5b600181811c9082168061006d57607f821691505b60208210810361008b57634e487b7160e01b5f52602260045260245ffd5b50919050565b601f8211156100d857805f5260205f20601f840160051c810160208510156100b65750805b601f840160051c820191505b818110156100d5575f81556001016100c2565b50505b505050565b81516001600160401b038111156100f6576100f6610045565b61010a816101048454610059565b84610091565b602080601f83116001811461013d575f84156101265750858301515b5f19600386901b1c1916600185901b178555610194565b5f85815260208120601f198616915b8281101561016b5788860151825594840194600190910190840161014c565b508582101561018857878501515f19600388901b60f8161c191681555b505060018460011b0185555b505050505050565b61042c806101a95f395ff3fe608060405234801561000f575f80fd5b506004361061003f575f3560e01c8063c040622614610043578063c15bae8414610061578063c47f002714610069575b5f80fd5b61004b61007e565b60405161005891906101a7565b60405180910390f35b61004b61010d565b61007c610077366004610207565b610198565b005b60605f805461008c906102b2565b80601f01602080910402602001604051908101604052809291908181526020018280546100b8906102b2565b80156101035780601f106100da57610100808354040283529160200191610103565b820191905f5260205f20905b8154815290600101906020018083116100e657829003601f168201915b5050505050905090565b5f8054610119906102b2565b80601f0160208091040260200160405190810160405280929190818152602001828054610145906102b2565b80156101905780601f1061016757610100808354040283529160200191610190565b820191905f5260205f20905b81548152906001019060200180831161017357829003601f168201915b505050505081565b5f6101a38282610336565b5050565b5f602080835283518060208501525f5b818110156101d3578581018301518582016040015282016101b7565b505f604082860101526040601f19601f8301168501019250505092915050565b634e487b7160e01b5f52604160045260245ffd5b5f60208284031215610217575f80fd5b813567ffffffffffffffff8082111561022e575f80fd5b818401915084601f830112610241575f80fd5b813581811115610253576102536101f3565b604051601f8201601f19908116603f0116810190838211818310171561027b5761027b6101f3565b81604052828152876020848701011115610293575f80fd5b826020860160208301375f928101602001929092525095945050505050565b600181811c908216806102c657607f821691505b6020821081036102e457634e487b7160e01b5f52602260045260245ffd5b50919050565b601f82111561033157805f5260205f20601f840160051c8101602085101561030f5750805b601f840160051c820191505b8181101561032e575f815560010161031b565b50505b505050565b815167ffffffffffffffff811115610350576103506101f3565b6103648161035e84546102b2565b846102ea565b602080601f831160018114610397575f84156103805750858301515b5f19600386901b1c1916600185901b1785556103ee565b5f85815260208120601f198616915b828110156103c5578886015182559484019460019091019084016103a6565b50858210156103e257878501515f19600388901b60f8161c191681555b505060018460011b0185555b50505050505056fea2646970667358221220c896c345552a36eea6b73606ff322b52d5dbea45843bad29a9b88b474afae6a564736f6c63430008180033",
}

// LargeMemoTCABI is the input ABI used to generate the binding from.
// Deprecated: Use LargeMemoTCMetaData.ABI instead.
var LargeMemoTCABI = LargeMemoTCMetaData.ABI

// LargeMemoTCBinRuntime is the compiled bytecode used for adding genesis block without deploying code.
const LargeMemoTCBinRuntime = ``

// LargeMemoTCBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use LargeMemoTCMetaData.Bin instead.
var LargeMemoTCBin = LargeMemoTCMetaData.Bin

// DeployLargeMemoTC deploys a new Kaia contract, binding an instance of LargeMemoTC to it.
func DeployLargeMemoTC(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *LargeMemoTC, error) {
	parsed, err := LargeMemoTCMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(LargeMemoTCBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &LargeMemoTC{LargeMemoTCCaller: LargeMemoTCCaller{contract: contract}, LargeMemoTCTransactor: LargeMemoTCTransactor{contract: contract}, LargeMemoTCFilterer: LargeMemoTCFilterer{contract: contract}}, nil
}

// LargeMemoTC is an auto generated Go binding around a Kaia contract.
type LargeMemoTC struct {
	LargeMemoTCCaller     // Read-only binding to the contract
	LargeMemoTCTransactor // Write-only binding to the contract
	LargeMemoTCFilterer   // Log filterer for contract events
}

// LargeMemoTCCaller is an auto generated read-only Go binding around a Kaia contract.
type LargeMemoTCCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LargeMemoTCTransactor is an auto generated write-only Go binding around a Kaia contract.
type LargeMemoTCTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LargeMemoTCFilterer is an auto generated log filtering Go binding around a Kaia contract events.
type LargeMemoTCFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LargeMemoTCSession is an auto generated Go binding around a Kaia contract,
// with pre-set call and transact options.
type LargeMemoTCSession struct {
	Contract     *LargeMemoTC      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// LargeMemoTCCallerSession is an auto generated read-only Go binding around a Kaia contract,
// with pre-set call options.
type LargeMemoTCCallerSession struct {
	Contract *LargeMemoTCCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// LargeMemoTCTransactorSession is an auto generated write-only Go binding around a Kaia contract,
// with pre-set transact options.
type LargeMemoTCTransactorSession struct {
	Contract     *LargeMemoTCTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// LargeMemoTCRaw is an auto generated low-level Go binding around a Kaia contract.
type LargeMemoTCRaw struct {
	Contract *LargeMemoTC // Generic contract binding to access the raw methods on
}

// LargeMemoTCCallerRaw is an auto generated low-level read-only Go binding around a Kaia contract.
type LargeMemoTCCallerRaw struct {
	Contract *LargeMemoTCCaller // Generic read-only contract binding to access the raw methods on
}

// LargeMemoTCTransactorRaw is an auto generated low-level write-only Go binding around a Kaia contract.
type LargeMemoTCTransactorRaw struct {
	Contract *LargeMemoTCTransactor // Generic write-only contract binding to access the raw methods on
}

// NewLargeMemoTC creates a new instance of LargeMemoTC, bound to a specific deployed contract.
func NewLargeMemoTC(address common.Address, backend bind.ContractBackend) (*LargeMemoTC, error) {
	contract, err := bindLargeMemoTC(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &LargeMemoTC{LargeMemoTCCaller: LargeMemoTCCaller{contract: contract}, LargeMemoTCTransactor: LargeMemoTCTransactor{contract: contract}, LargeMemoTCFilterer: LargeMemoTCFilterer{contract: contract}}, nil
}

// NewLargeMemoTCCaller creates a new read-only instance of LargeMemoTC, bound to a specific deployed contract.
func NewLargeMemoTCCaller(address common.Address, caller bind.ContractCaller) (*LargeMemoTCCaller, error) {
	contract, err := bindLargeMemoTC(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &LargeMemoTCCaller{contract: contract}, nil
}

// NewLargeMemoTCTransactor creates a new write-only instance of LargeMemoTC, bound to a specific deployed contract.
func NewLargeMemoTCTransactor(address common.Address, transactor bind.ContractTransactor) (*LargeMemoTCTransactor, error) {
	contract, err := bindLargeMemoTC(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &LargeMemoTCTransactor{contract: contract}, nil
}

// NewLargeMemoTCFilterer creates a new log filterer instance of LargeMemoTC, bound to a specific deployed contract.
func NewLargeMemoTCFilterer(address common.Address, filterer bind.ContractFilterer) (*LargeMemoTCFilterer, error) {
	contract, err := bindLargeMemoTC(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &LargeMemoTCFilterer{contract: contract}, nil
}

// bindLargeMemoTC binds a generic wrapper to an already deployed contract.
func bindLargeMemoTC(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := LargeMemoTCMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LargeMemoTC *LargeMemoTCRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LargeMemoTC.Contract.LargeMemoTCCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LargeMemoTC *LargeMemoTCRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LargeMemoTC.Contract.LargeMemoTCTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LargeMemoTC *LargeMemoTCRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LargeMemoTC.Contract.LargeMemoTCTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LargeMemoTC *LargeMemoTCCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LargeMemoTC.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LargeMemoTC *LargeMemoTCTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LargeMemoTC.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LargeMemoTC *LargeMemoTCTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LargeMemoTC.Contract.contract.Transact(opts, method, params...)
}

// Run is a free data retrieval call binding the contract method 0xc0406226.
//
// Solidity: function run() view returns(string)
func (_LargeMemoTC *LargeMemoTCCaller) Run(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _LargeMemoTC.contract.Call(opts, &out, "run")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Run is a free data retrieval call binding the contract method 0xc0406226.
//
// Solidity: function run() view returns(string)
func (_LargeMemoTC *LargeMemoTCSession) Run() (string, error) {
	return _LargeMemoTC.Contract.Run(&_LargeMemoTC.CallOpts)
}

// Run is a free data retrieval call binding the contract method 0xc0406226.
//
// Solidity: function run() view returns(string)
func (_LargeMemoTC *LargeMemoTCCallerSession) Run() (string, error) {
	return _LargeMemoTC.Contract.Run(&_LargeMemoTC.CallOpts)
}

// Str is a free data retrieval call binding the contract method 0xc15bae84.
//
// Solidity: function str() view returns(string)
func (_LargeMemoTC *LargeMemoTCCaller) Str(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _LargeMemoTC.contract.Call(opts, &out, "str")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Str is a free data retrieval call binding the contract method 0xc15bae84.
//
// Solidity: function str() view returns(string)
func (_LargeMemoTC *LargeMemoTCSession) Str() (string, error) {
	return _LargeMemoTC.Contract.Str(&_LargeMemoTC.CallOpts)
}

// Str is a free data retrieval call binding the contract method 0xc15bae84.
//
// Solidity: function str() view returns(string)
func (_LargeMemoTC *LargeMemoTCCallerSession) Str() (string, error) {
	return _LargeMemoTC.Contract.Str(&_LargeMemoTC.CallOpts)
}

// SetName is a paid mutator transaction binding the contract method 0xc47f0027.
//
// Solidity: function setName(string _str) returns()
func (_LargeMemoTC *LargeMemoTCTransactor) SetName(opts *bind.TransactOpts, _str string) (*types.Transaction, error) {
	return _LargeMemoTC.contract.Transact(opts, "setName", _str)
}

// SetName is a paid mutator transaction binding the contract method 0xc47f0027.
//
// Solidity: function setName(string _str) returns()
func (_LargeMemoTC *LargeMemoTCSession) SetName(_str string) (*types.Transaction, error) {
	return _LargeMemoTC.Contract.SetName(&_LargeMemoTC.TransactOpts, _str)
}

// SetName is a paid mutator transaction binding the contract method 0xc47f0027.
//
// Solidity: function setName(string _str) returns()
func (_LargeMemoTC *LargeMemoTCTransactorSession) SetName(_str string) (*types.Transaction, error) {
	return _LargeMemoTC.Contract.SetName(&_LargeMemoTC.TransactOpts, _str)
}
