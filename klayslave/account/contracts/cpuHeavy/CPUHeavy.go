// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package cpuHeavy

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

// CpuHeavyTCMetaData contains all meta data concerning the CpuHeavyTC contract.
var CpuHeavyTCMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"size\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"signature\",\"type\":\"uint256\"}],\"name\":\"finish\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"checkResult\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"empty\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"size\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"signature\",\"type\":\"uint256\"}],\"name\":\"sort\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"signature\",\"type\":\"uint256\"}],\"name\":\"sortSingle\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561000f575f80fd5b506106e78061001d5f395ff3fe608060405234801561000f575f80fd5b506004361061004a575f3560e01c80637b395ec21461004e578063a21d942f14610063578063e71c6c821461007f578063f2a75fe414610092575b5f80fd5b61006161005c3660046105e5565b61009a565b005b61006b610171565b604051901515815260200160405180910390f35b61006161008d366004610605565b6101c2565b610061610249565b5f8267ffffffffffffffff8111156100b4576100b461061c565b6040519080825280602002602001820160405280156100dd578160200160208202803683370190505b5090505f5b815181101561011a576100f58185610644565b8282815181106101075761010761065d565b60209081029190910101526001016100e2565b50610133815f6001845161012e9190610644565b61025e565b60408051848152602081018490527fd596fdad182d29130ce218f4c1590c4b5ede105bee36690727baa6592bd2bfc8910160405180910390a1505050565b600180545f9182905b60148110156101b857600181601481106101965761019661065d565b01549150818311156101ab575f935050505090565b909150819060010161017a565b5060019250505090565b60145f5b60148110156101f6576101d98183610644565b600182601481106101ec576101ec61065d565b01556001016101c6565b5061020c5f61020760016014610644565b61043f565b60408051828152602081018490527fd596fdad182d29130ce218f4c1590c4b5ede105bee36690727baa6592bd2bfc8910160405180910390a15050565b5f8054908061025783610671565b9190505550565b5f805f838510156104375750839150829050815b8183101561039b575b85818151811061028d5761028d61065d565b60200260200101518684815181106102a7576102a761065d565b6020026020010151111580156102bc57508383105b156102d357826102cb81610671565b93505061027b565b8581815181106102e5576102e561065d565b60200260200101518683815181106102ff576102ff61065d565b6020026020010151111561031f578161031781610689565b9250506102d3565b81831015610396578582815181106103395761033961065d565b60200260200101518684815181106103535761035361065d565b602002602001015187858151811061036d5761036d61065d565b602002602001018885815181106103865761038661065d565b6020908102919091010191909152525b610272565b8582815181106103ad576103ad61065d565b60200260200101518682815181106103c7576103c761065d565b60200260200101518783815181106103e1576103e161065d565b602002602001018885815181106103fa576103fa61065d565b602090810291909101019190915252600182111561042257610422868661012e600186610644565b6104378661043184600161069e565b8661025e565b505050505050565b5f805f838510156105de5750839150829050815b81831015610556575b6001816014811061046f5761046f61065d565b0154600184601481106104845761048461065d565b01541115801561049357508383105b156104aa57826104a281610671565b93505061045c565b600181601481106104bd576104bd61065d565b0154600183601481106104d2576104d261065d565b015411156104ec57816104e481610689565b9250506104aa565b8183101561055157600182601481106105075761050761065d565b01546001846014811061051c5761051c61065d565b0154600185601481106105315761053161065d565b015f600186601481106105465761054661065d565b019290925591909155505b610453565b600182601481106105695761056961065d565b01546001826014811061057e5761057e61065d565b0154600183601481106105935761059361065d565b015f600186601481106105a8576105a861065d565b0192909255919091555060018211156105ca576105ca85610207600185610644565b6105de6105d883600161069e565b8561043f565b5050505050565b5f80604083850312156105f6575f80fd5b50508035926020909101359150565b5f60208284031215610615575f80fd5b5035919050565b634e487b7160e01b5f52604160045260245ffd5b634e487b7160e01b5f52601160045260245ffd5b8181038181111561065757610657610630565b92915050565b634e487b7160e01b5f52603260045260245ffd5b5f6001820161068257610682610630565b5060010190565b5f8161069757610697610630565b505f190190565b808201808211156106575761065761063056fea2646970667358221220c2afba695068cfa8270096d1ec1d5be5842a7db5d78cff7c762ae65a884d38fd64736f6c63430008180033",
}

// CpuHeavyTCABI is the input ABI used to generate the binding from.
// Deprecated: Use CpuHeavyTCMetaData.ABI instead.
var CpuHeavyTCABI = CpuHeavyTCMetaData.ABI

// CpuHeavyTCBinRuntime is the compiled bytecode used for adding genesis block without deploying code.
const CpuHeavyTCBinRuntime = ``

// CpuHeavyTCBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use CpuHeavyTCMetaData.Bin instead.
var CpuHeavyTCBin = CpuHeavyTCMetaData.Bin

// DeployCpuHeavyTC deploys a new Kaia contract, binding an instance of CpuHeavyTC to it.
func DeployCpuHeavyTC(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *CpuHeavyTC, error) {
	parsed, err := CpuHeavyTCMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(CpuHeavyTCBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &CpuHeavyTC{CpuHeavyTCCaller: CpuHeavyTCCaller{contract: contract}, CpuHeavyTCTransactor: CpuHeavyTCTransactor{contract: contract}, CpuHeavyTCFilterer: CpuHeavyTCFilterer{contract: contract}}, nil
}

// CpuHeavyTC is an auto generated Go binding around a Kaia contract.
type CpuHeavyTC struct {
	CpuHeavyTCCaller     // Read-only binding to the contract
	CpuHeavyTCTransactor // Write-only binding to the contract
	CpuHeavyTCFilterer   // Log filterer for contract events
}

// CpuHeavyTCCaller is an auto generated read-only Go binding around a Kaia contract.
type CpuHeavyTCCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CpuHeavyTCTransactor is an auto generated write-only Go binding around a Kaia contract.
type CpuHeavyTCTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CpuHeavyTCFilterer is an auto generated log filtering Go binding around a Kaia contract events.
type CpuHeavyTCFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CpuHeavyTCSession is an auto generated Go binding around a Kaia contract,
// with pre-set call and transact options.
type CpuHeavyTCSession struct {
	Contract     *CpuHeavyTC       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CpuHeavyTCCallerSession is an auto generated read-only Go binding around a Kaia contract,
// with pre-set call options.
type CpuHeavyTCCallerSession struct {
	Contract *CpuHeavyTCCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// CpuHeavyTCTransactorSession is an auto generated write-only Go binding around a Kaia contract,
// with pre-set transact options.
type CpuHeavyTCTransactorSession struct {
	Contract     *CpuHeavyTCTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// CpuHeavyTCRaw is an auto generated low-level Go binding around a Kaia contract.
type CpuHeavyTCRaw struct {
	Contract *CpuHeavyTC // Generic contract binding to access the raw methods on
}

// CpuHeavyTCCallerRaw is an auto generated low-level read-only Go binding around a Kaia contract.
type CpuHeavyTCCallerRaw struct {
	Contract *CpuHeavyTCCaller // Generic read-only contract binding to access the raw methods on
}

// CpuHeavyTCTransactorRaw is an auto generated low-level write-only Go binding around a Kaia contract.
type CpuHeavyTCTransactorRaw struct {
	Contract *CpuHeavyTCTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCpuHeavyTC creates a new instance of CpuHeavyTC, bound to a specific deployed contract.
func NewCpuHeavyTC(address common.Address, backend bind.ContractBackend) (*CpuHeavyTC, error) {
	contract, err := bindCpuHeavyTC(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CpuHeavyTC{CpuHeavyTCCaller: CpuHeavyTCCaller{contract: contract}, CpuHeavyTCTransactor: CpuHeavyTCTransactor{contract: contract}, CpuHeavyTCFilterer: CpuHeavyTCFilterer{contract: contract}}, nil
}

// NewCpuHeavyTCCaller creates a new read-only instance of CpuHeavyTC, bound to a specific deployed contract.
func NewCpuHeavyTCCaller(address common.Address, caller bind.ContractCaller) (*CpuHeavyTCCaller, error) {
	contract, err := bindCpuHeavyTC(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CpuHeavyTCCaller{contract: contract}, nil
}

// NewCpuHeavyTCTransactor creates a new write-only instance of CpuHeavyTC, bound to a specific deployed contract.
func NewCpuHeavyTCTransactor(address common.Address, transactor bind.ContractTransactor) (*CpuHeavyTCTransactor, error) {
	contract, err := bindCpuHeavyTC(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CpuHeavyTCTransactor{contract: contract}, nil
}

// NewCpuHeavyTCFilterer creates a new log filterer instance of CpuHeavyTC, bound to a specific deployed contract.
func NewCpuHeavyTCFilterer(address common.Address, filterer bind.ContractFilterer) (*CpuHeavyTCFilterer, error) {
	contract, err := bindCpuHeavyTC(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CpuHeavyTCFilterer{contract: contract}, nil
}

// bindCpuHeavyTC binds a generic wrapper to an already deployed contract.
func bindCpuHeavyTC(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := CpuHeavyTCMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CpuHeavyTC *CpuHeavyTCRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CpuHeavyTC.Contract.CpuHeavyTCCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CpuHeavyTC *CpuHeavyTCRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CpuHeavyTC.Contract.CpuHeavyTCTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CpuHeavyTC *CpuHeavyTCRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CpuHeavyTC.Contract.CpuHeavyTCTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CpuHeavyTC *CpuHeavyTCCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CpuHeavyTC.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CpuHeavyTC *CpuHeavyTCTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CpuHeavyTC.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CpuHeavyTC *CpuHeavyTCTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CpuHeavyTC.Contract.contract.Transact(opts, method, params...)
}

// CheckResult is a free data retrieval call binding the contract method 0xa21d942f.
//
// Solidity: function checkResult() view returns(bool)
func (_CpuHeavyTC *CpuHeavyTCCaller) CheckResult(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _CpuHeavyTC.contract.Call(opts, &out, "checkResult")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CheckResult is a free data retrieval call binding the contract method 0xa21d942f.
//
// Solidity: function checkResult() view returns(bool)
func (_CpuHeavyTC *CpuHeavyTCSession) CheckResult() (bool, error) {
	return _CpuHeavyTC.Contract.CheckResult(&_CpuHeavyTC.CallOpts)
}

// CheckResult is a free data retrieval call binding the contract method 0xa21d942f.
//
// Solidity: function checkResult() view returns(bool)
func (_CpuHeavyTC *CpuHeavyTCCallerSession) CheckResult() (bool, error) {
	return _CpuHeavyTC.Contract.CheckResult(&_CpuHeavyTC.CallOpts)
}

// Empty is a paid mutator transaction binding the contract method 0xf2a75fe4.
//
// Solidity: function empty() returns()
func (_CpuHeavyTC *CpuHeavyTCTransactor) Empty(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CpuHeavyTC.contract.Transact(opts, "empty")
}

// Empty is a paid mutator transaction binding the contract method 0xf2a75fe4.
//
// Solidity: function empty() returns()
func (_CpuHeavyTC *CpuHeavyTCSession) Empty() (*types.Transaction, error) {
	return _CpuHeavyTC.Contract.Empty(&_CpuHeavyTC.TransactOpts)
}

// Empty is a paid mutator transaction binding the contract method 0xf2a75fe4.
//
// Solidity: function empty() returns()
func (_CpuHeavyTC *CpuHeavyTCTransactorSession) Empty() (*types.Transaction, error) {
	return _CpuHeavyTC.Contract.Empty(&_CpuHeavyTC.TransactOpts)
}

// Sort is a paid mutator transaction binding the contract method 0x7b395ec2.
//
// Solidity: function sort(uint256 size, uint256 signature) returns()
func (_CpuHeavyTC *CpuHeavyTCTransactor) Sort(opts *bind.TransactOpts, size *big.Int, signature *big.Int) (*types.Transaction, error) {
	return _CpuHeavyTC.contract.Transact(opts, "sort", size, signature)
}

// Sort is a paid mutator transaction binding the contract method 0x7b395ec2.
//
// Solidity: function sort(uint256 size, uint256 signature) returns()
func (_CpuHeavyTC *CpuHeavyTCSession) Sort(size *big.Int, signature *big.Int) (*types.Transaction, error) {
	return _CpuHeavyTC.Contract.Sort(&_CpuHeavyTC.TransactOpts, size, signature)
}

// Sort is a paid mutator transaction binding the contract method 0x7b395ec2.
//
// Solidity: function sort(uint256 size, uint256 signature) returns()
func (_CpuHeavyTC *CpuHeavyTCTransactorSession) Sort(size *big.Int, signature *big.Int) (*types.Transaction, error) {
	return _CpuHeavyTC.Contract.Sort(&_CpuHeavyTC.TransactOpts, size, signature)
}

// SortSingle is a paid mutator transaction binding the contract method 0xe71c6c82.
//
// Solidity: function sortSingle(uint256 signature) returns()
func (_CpuHeavyTC *CpuHeavyTCTransactor) SortSingle(opts *bind.TransactOpts, signature *big.Int) (*types.Transaction, error) {
	return _CpuHeavyTC.contract.Transact(opts, "sortSingle", signature)
}

// SortSingle is a paid mutator transaction binding the contract method 0xe71c6c82.
//
// Solidity: function sortSingle(uint256 signature) returns()
func (_CpuHeavyTC *CpuHeavyTCSession) SortSingle(signature *big.Int) (*types.Transaction, error) {
	return _CpuHeavyTC.Contract.SortSingle(&_CpuHeavyTC.TransactOpts, signature)
}

// SortSingle is a paid mutator transaction binding the contract method 0xe71c6c82.
//
// Solidity: function sortSingle(uint256 signature) returns()
func (_CpuHeavyTC *CpuHeavyTCTransactorSession) SortSingle(signature *big.Int) (*types.Transaction, error) {
	return _CpuHeavyTC.Contract.SortSingle(&_CpuHeavyTC.TransactOpts, signature)
}

// CpuHeavyTCFinishIterator is returned from FilterFinish and is used to iterate over the raw logs and unpacked data for Finish events raised by the CpuHeavyTC contract.
type CpuHeavyTCFinishIterator struct {
	Event *CpuHeavyTCFinish // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log    // Log channel receiving the found contract events
	sub  kaia.Subscription // Subscription for errors, completion and termination
	done bool              // Whether the subscription completed delivering logs
	fail error             // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CpuHeavyTCFinishIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CpuHeavyTCFinish)
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
		it.Event = new(CpuHeavyTCFinish)
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
func (it *CpuHeavyTCFinishIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CpuHeavyTCFinishIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CpuHeavyTCFinish represents a Finish event raised by the CpuHeavyTC contract.
type CpuHeavyTCFinish struct {
	Size      *big.Int
	Signature *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterFinish is a free log retrieval operation binding the contract event 0xd596fdad182d29130ce218f4c1590c4b5ede105bee36690727baa6592bd2bfc8.
//
// Solidity: event finish(uint256 size, uint256 signature)
func (_CpuHeavyTC *CpuHeavyTCFilterer) FilterFinish(opts *bind.FilterOpts) (*CpuHeavyTCFinishIterator, error) {

	logs, sub, err := _CpuHeavyTC.contract.FilterLogs(opts, "finish")
	if err != nil {
		return nil, err
	}
	return &CpuHeavyTCFinishIterator{contract: _CpuHeavyTC.contract, event: "finish", logs: logs, sub: sub}, nil
}

// WatchFinish is a free log subscription operation binding the contract event 0xd596fdad182d29130ce218f4c1590c4b5ede105bee36690727baa6592bd2bfc8.
//
// Solidity: event finish(uint256 size, uint256 signature)
func (_CpuHeavyTC *CpuHeavyTCFilterer) WatchFinish(opts *bind.WatchOpts, sink chan<- *CpuHeavyTCFinish) (event.Subscription, error) {

	logs, sub, err := _CpuHeavyTC.contract.WatchLogs(opts, "finish")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CpuHeavyTCFinish)
				if err := _CpuHeavyTC.contract.UnpackLog(event, "finish", log); err != nil {
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

// ParseFinish is a log parse operation binding the contract event 0xd596fdad182d29130ce218f4c1590c4b5ede105bee36690727baa6592bd2bfc8.
//
// Solidity: event finish(uint256 size, uint256 signature)
func (_CpuHeavyTC *CpuHeavyTCFilterer) ParseFinish(log types.Log) (*CpuHeavyTCFinish, error) {
	event := new(CpuHeavyTCFinish)
	if err := _CpuHeavyTC.contract.UnpackLog(event, "finish", log); err != nil {
		return nil, err
	}
	return event, nil
}
