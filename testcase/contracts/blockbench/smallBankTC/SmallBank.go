// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package smallBankTC

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

// SmallBankTCMetaData contains all meta data concerning the SmallBankTC contract.
var SmallBankTCMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"arg0\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"arg1\",\"type\":\"string\"}],\"name\":\"almagate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"arg0\",\"type\":\"string\"}],\"name\":\"getBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"arg0\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"arg1\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"arg2\",\"type\":\"uint256\"}],\"name\":\"sendPayment\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"arg0\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"arg1\",\"type\":\"uint256\"}],\"name\":\"updateBalance\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"arg0\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"arg1\",\"type\":\"uint256\"}],\"name\":\"updateSaving\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"arg0\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"arg1\",\"type\":\"uint256\"}],\"name\":\"writeCheck\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561000f575f80fd5b506106368061001d5f395ff3fe608060405234801561000f575f80fd5b5060043610610060575f3560e01c80630b488b37146100645780630be8374d146100795780633a51d2461461008c578063870187eb146100b1578063901d706f146100c4578063ca305435146100d7575b5f80fd5b610077610072366004610458565b6100ea565b005b610077610087366004610458565b61013c565b61009f61009a36600461049a565b6101ff565b60405190815260200160405180910390f35b6100776100bf366004610458565b610256565b6100776100d23660046104cc565b610294565b6100776100e536600461052c565b610311565b5f80836040516100fa9190610594565b9081526040519081900360200190205490508161011781836105d4565b5f856040516101269190610594565b9081526040519081900360200190205550505050565b5f60018360405161014d9190610594565b90815260200160405180910390205490505f808460405161016e9190610594565b9081526040519081900360200190205490508261018b82846105d4565b8110156101cd57600161019e82856105ed565b6101a891906105ed565b6001866040516101b89190610594565b908152604051908190036020019020556101f8565b6101d781846105ed565b6001866040516101e79190610594565b908152604051908190036020019020555b5050505050565b5f805f836040516102109190610594565b90815260200160405180910390205490505f6001846040516102329190610594565b90815260405190819003602001902054905061024e81836105d4565b949350505050565b5f6001836040516102679190610594565b9081526040519081900360200190205490508161028481836105d4565b6001856040516101269190610594565b5f80836040516102a49190610594565b90815260200160405180910390205490505f6001836040516102c69190610594565b90815260200160405180910390205490505f6001856040516102e89190610594565b9081526040519081900360200190205561030281836105d4565b5f846040516101269190610594565b5f6001846040516103229190610594565b90815260200160405180910390205490505f6001846040516103449190610594565b9081526040519081900360200190205490508261036181846105ed565b925061036d81836105d4565b9150826001876040516103809190610594565b908152602001604051809103902081905550816001866040516103a39190610594565b90815260405190819003602001902055505050505050565b634e487b7160e01b5f52604160045260245ffd5b5f82601f8301126103de575f80fd5b813567ffffffffffffffff808211156103f9576103f96103bb565b604051601f8301601f19908116603f01168101908282118183101715610421576104216103bb565b81604052838152866020858801011115610439575f80fd5b836020870160208301375f602085830101528094505050505092915050565b5f8060408385031215610469575f80fd5b823567ffffffffffffffff81111561047f575f80fd5b61048b858286016103cf565b95602094909401359450505050565b5f602082840312156104aa575f80fd5b813567ffffffffffffffff8111156104c0575f80fd5b61024e848285016103cf565b5f80604083850312156104dd575f80fd5b823567ffffffffffffffff808211156104f4575f80fd5b610500868387016103cf565b93506020850135915080821115610515575f80fd5b50610522858286016103cf565b9150509250929050565b5f805f6060848603121561053e575f80fd5b833567ffffffffffffffff80821115610555575f80fd5b610561878388016103cf565b94506020860135915080821115610576575f80fd5b50610583868287016103cf565b925050604084013590509250925092565b5f82515f5b818110156105b35760208186018101518583015201610599565b505f920191825250919050565b634e487b7160e01b5f52601160045260245ffd5b808201808211156105e7576105e76105c0565b92915050565b818103818111156105e7576105e76105c056fea264697066735822122086ca0cb5745e2990a6dfd1457f9c429d977d8cdba7d8ea7784f689277a669de164736f6c63430008180033",
}

// SmallBankTCABI is the input ABI used to generate the binding from.
// Deprecated: Use SmallBankTCMetaData.ABI instead.
var SmallBankTCABI = SmallBankTCMetaData.ABI

// SmallBankTCBinRuntime is the compiled bytecode used for adding genesis block without deploying code.
const SmallBankTCBinRuntime = ``

// SmallBankTCBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use SmallBankTCMetaData.Bin instead.
var SmallBankTCBin = SmallBankTCMetaData.Bin

// DeploySmallBankTC deploys a new Kaia contract, binding an instance of SmallBankTC to it.
func DeploySmallBankTC(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SmallBankTC, error) {
	parsed, err := SmallBankTCMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(SmallBankTCBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SmallBankTC{SmallBankTCCaller: SmallBankTCCaller{contract: contract}, SmallBankTCTransactor: SmallBankTCTransactor{contract: contract}, SmallBankTCFilterer: SmallBankTCFilterer{contract: contract}}, nil
}

// SmallBankTC is an auto generated Go binding around a Kaia contract.
type SmallBankTC struct {
	SmallBankTCCaller     // Read-only binding to the contract
	SmallBankTCTransactor // Write-only binding to the contract
	SmallBankTCFilterer   // Log filterer for contract events
}

// SmallBankTCCaller is an auto generated read-only Go binding around a Kaia contract.
type SmallBankTCCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SmallBankTCTransactor is an auto generated write-only Go binding around a Kaia contract.
type SmallBankTCTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SmallBankTCFilterer is an auto generated log filtering Go binding around a Kaia contract events.
type SmallBankTCFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SmallBankTCSession is an auto generated Go binding around a Kaia contract,
// with pre-set call and transact options.
type SmallBankTCSession struct {
	Contract     *SmallBankTC      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SmallBankTCCallerSession is an auto generated read-only Go binding around a Kaia contract,
// with pre-set call options.
type SmallBankTCCallerSession struct {
	Contract *SmallBankTCCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// SmallBankTCTransactorSession is an auto generated write-only Go binding around a Kaia contract,
// with pre-set transact options.
type SmallBankTCTransactorSession struct {
	Contract     *SmallBankTCTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// SmallBankTCRaw is an auto generated low-level Go binding around a Kaia contract.
type SmallBankTCRaw struct {
	Contract *SmallBankTC // Generic contract binding to access the raw methods on
}

// SmallBankTCCallerRaw is an auto generated low-level read-only Go binding around a Kaia contract.
type SmallBankTCCallerRaw struct {
	Contract *SmallBankTCCaller // Generic read-only contract binding to access the raw methods on
}

// SmallBankTCTransactorRaw is an auto generated low-level write-only Go binding around a Kaia contract.
type SmallBankTCTransactorRaw struct {
	Contract *SmallBankTCTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSmallBankTC creates a new instance of SmallBankTC, bound to a specific deployed contract.
func NewSmallBankTC(address common.Address, backend bind.ContractBackend) (*SmallBankTC, error) {
	contract, err := bindSmallBankTC(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SmallBankTC{SmallBankTCCaller: SmallBankTCCaller{contract: contract}, SmallBankTCTransactor: SmallBankTCTransactor{contract: contract}, SmallBankTCFilterer: SmallBankTCFilterer{contract: contract}}, nil
}

// NewSmallBankTCCaller creates a new read-only instance of SmallBankTC, bound to a specific deployed contract.
func NewSmallBankTCCaller(address common.Address, caller bind.ContractCaller) (*SmallBankTCCaller, error) {
	contract, err := bindSmallBankTC(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SmallBankTCCaller{contract: contract}, nil
}

// NewSmallBankTCTransactor creates a new write-only instance of SmallBankTC, bound to a specific deployed contract.
func NewSmallBankTCTransactor(address common.Address, transactor bind.ContractTransactor) (*SmallBankTCTransactor, error) {
	contract, err := bindSmallBankTC(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SmallBankTCTransactor{contract: contract}, nil
}

// NewSmallBankTCFilterer creates a new log filterer instance of SmallBankTC, bound to a specific deployed contract.
func NewSmallBankTCFilterer(address common.Address, filterer bind.ContractFilterer) (*SmallBankTCFilterer, error) {
	contract, err := bindSmallBankTC(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SmallBankTCFilterer{contract: contract}, nil
}

// bindSmallBankTC binds a generic wrapper to an already deployed contract.
func bindSmallBankTC(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := SmallBankTCMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SmallBankTC *SmallBankTCRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SmallBankTC.Contract.SmallBankTCCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SmallBankTC *SmallBankTCRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SmallBankTC.Contract.SmallBankTCTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SmallBankTC *SmallBankTCRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SmallBankTC.Contract.SmallBankTCTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SmallBankTC *SmallBankTCCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SmallBankTC.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SmallBankTC *SmallBankTCTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SmallBankTC.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SmallBankTC *SmallBankTCTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SmallBankTC.Contract.contract.Transact(opts, method, params...)
}

// GetBalance is a free data retrieval call binding the contract method 0x3a51d246.
//
// Solidity: function getBalance(string arg0) view returns(uint256 balance)
func (_SmallBankTC *SmallBankTCCaller) GetBalance(opts *bind.CallOpts, arg0 string) (*big.Int, error) {
	var out []interface{}
	err := _SmallBankTC.contract.Call(opts, &out, "getBalance", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBalance is a free data retrieval call binding the contract method 0x3a51d246.
//
// Solidity: function getBalance(string arg0) view returns(uint256 balance)
func (_SmallBankTC *SmallBankTCSession) GetBalance(arg0 string) (*big.Int, error) {
	return _SmallBankTC.Contract.GetBalance(&_SmallBankTC.CallOpts, arg0)
}

// GetBalance is a free data retrieval call binding the contract method 0x3a51d246.
//
// Solidity: function getBalance(string arg0) view returns(uint256 balance)
func (_SmallBankTC *SmallBankTCCallerSession) GetBalance(arg0 string) (*big.Int, error) {
	return _SmallBankTC.Contract.GetBalance(&_SmallBankTC.CallOpts, arg0)
}

// Almagate is a paid mutator transaction binding the contract method 0x901d706f.
//
// Solidity: function almagate(string arg0, string arg1) returns()
func (_SmallBankTC *SmallBankTCTransactor) Almagate(opts *bind.TransactOpts, arg0 string, arg1 string) (*types.Transaction, error) {
	return _SmallBankTC.contract.Transact(opts, "almagate", arg0, arg1)
}

// Almagate is a paid mutator transaction binding the contract method 0x901d706f.
//
// Solidity: function almagate(string arg0, string arg1) returns()
func (_SmallBankTC *SmallBankTCSession) Almagate(arg0 string, arg1 string) (*types.Transaction, error) {
	return _SmallBankTC.Contract.Almagate(&_SmallBankTC.TransactOpts, arg0, arg1)
}

// Almagate is a paid mutator transaction binding the contract method 0x901d706f.
//
// Solidity: function almagate(string arg0, string arg1) returns()
func (_SmallBankTC *SmallBankTCTransactorSession) Almagate(arg0 string, arg1 string) (*types.Transaction, error) {
	return _SmallBankTC.Contract.Almagate(&_SmallBankTC.TransactOpts, arg0, arg1)
}

// SendPayment is a paid mutator transaction binding the contract method 0xca305435.
//
// Solidity: function sendPayment(string arg0, string arg1, uint256 arg2) returns()
func (_SmallBankTC *SmallBankTCTransactor) SendPayment(opts *bind.TransactOpts, arg0 string, arg1 string, arg2 *big.Int) (*types.Transaction, error) {
	return _SmallBankTC.contract.Transact(opts, "sendPayment", arg0, arg1, arg2)
}

// SendPayment is a paid mutator transaction binding the contract method 0xca305435.
//
// Solidity: function sendPayment(string arg0, string arg1, uint256 arg2) returns()
func (_SmallBankTC *SmallBankTCSession) SendPayment(arg0 string, arg1 string, arg2 *big.Int) (*types.Transaction, error) {
	return _SmallBankTC.Contract.SendPayment(&_SmallBankTC.TransactOpts, arg0, arg1, arg2)
}

// SendPayment is a paid mutator transaction binding the contract method 0xca305435.
//
// Solidity: function sendPayment(string arg0, string arg1, uint256 arg2) returns()
func (_SmallBankTC *SmallBankTCTransactorSession) SendPayment(arg0 string, arg1 string, arg2 *big.Int) (*types.Transaction, error) {
	return _SmallBankTC.Contract.SendPayment(&_SmallBankTC.TransactOpts, arg0, arg1, arg2)
}

// UpdateBalance is a paid mutator transaction binding the contract method 0x870187eb.
//
// Solidity: function updateBalance(string arg0, uint256 arg1) returns()
func (_SmallBankTC *SmallBankTCTransactor) UpdateBalance(opts *bind.TransactOpts, arg0 string, arg1 *big.Int) (*types.Transaction, error) {
	return _SmallBankTC.contract.Transact(opts, "updateBalance", arg0, arg1)
}

// UpdateBalance is a paid mutator transaction binding the contract method 0x870187eb.
//
// Solidity: function updateBalance(string arg0, uint256 arg1) returns()
func (_SmallBankTC *SmallBankTCSession) UpdateBalance(arg0 string, arg1 *big.Int) (*types.Transaction, error) {
	return _SmallBankTC.Contract.UpdateBalance(&_SmallBankTC.TransactOpts, arg0, arg1)
}

// UpdateBalance is a paid mutator transaction binding the contract method 0x870187eb.
//
// Solidity: function updateBalance(string arg0, uint256 arg1) returns()
func (_SmallBankTC *SmallBankTCTransactorSession) UpdateBalance(arg0 string, arg1 *big.Int) (*types.Transaction, error) {
	return _SmallBankTC.Contract.UpdateBalance(&_SmallBankTC.TransactOpts, arg0, arg1)
}

// UpdateSaving is a paid mutator transaction binding the contract method 0x0b488b37.
//
// Solidity: function updateSaving(string arg0, uint256 arg1) returns()
func (_SmallBankTC *SmallBankTCTransactor) UpdateSaving(opts *bind.TransactOpts, arg0 string, arg1 *big.Int) (*types.Transaction, error) {
	return _SmallBankTC.contract.Transact(opts, "updateSaving", arg0, arg1)
}

// UpdateSaving is a paid mutator transaction binding the contract method 0x0b488b37.
//
// Solidity: function updateSaving(string arg0, uint256 arg1) returns()
func (_SmallBankTC *SmallBankTCSession) UpdateSaving(arg0 string, arg1 *big.Int) (*types.Transaction, error) {
	return _SmallBankTC.Contract.UpdateSaving(&_SmallBankTC.TransactOpts, arg0, arg1)
}

// UpdateSaving is a paid mutator transaction binding the contract method 0x0b488b37.
//
// Solidity: function updateSaving(string arg0, uint256 arg1) returns()
func (_SmallBankTC *SmallBankTCTransactorSession) UpdateSaving(arg0 string, arg1 *big.Int) (*types.Transaction, error) {
	return _SmallBankTC.Contract.UpdateSaving(&_SmallBankTC.TransactOpts, arg0, arg1)
}

// WriteCheck is a paid mutator transaction binding the contract method 0x0be8374d.
//
// Solidity: function writeCheck(string arg0, uint256 arg1) returns()
func (_SmallBankTC *SmallBankTCTransactor) WriteCheck(opts *bind.TransactOpts, arg0 string, arg1 *big.Int) (*types.Transaction, error) {
	return _SmallBankTC.contract.Transact(opts, "writeCheck", arg0, arg1)
}

// WriteCheck is a paid mutator transaction binding the contract method 0x0be8374d.
//
// Solidity: function writeCheck(string arg0, uint256 arg1) returns()
func (_SmallBankTC *SmallBankTCSession) WriteCheck(arg0 string, arg1 *big.Int) (*types.Transaction, error) {
	return _SmallBankTC.Contract.WriteCheck(&_SmallBankTC.TransactOpts, arg0, arg1)
}

// WriteCheck is a paid mutator transaction binding the contract method 0x0be8374d.
//
// Solidity: function writeCheck(string arg0, uint256 arg1) returns()
func (_SmallBankTC *SmallBankTCTransactorSession) WriteCheck(arg0 string, arg1 *big.Int) (*types.Transaction, error) {
	return _SmallBankTC.Contract.WriteCheck(&_SmallBankTC.TransactOpts, arg0, arg1)
}
