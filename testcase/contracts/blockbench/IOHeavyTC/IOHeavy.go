// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package IOHeavyTC

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

// IoHeavyTCMetaData contains all meta data concerning the IoHeavyTC contract.
var IoHeavyTCMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"size\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"signature\",\"type\":\"uint256\"}],\"name\":\"finishScan\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"size\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"signature\",\"type\":\"uint256\"}],\"name\":\"finishWrite\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes20\",\"name\":\"key\",\"type\":\"bytes20\"}],\"name\":\"get\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"start_key\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"size\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"signature\",\"type\":\"uint256\"}],\"name\":\"revert_scan\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"start_key\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"size\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"signature\",\"type\":\"uint256\"}],\"name\":\"scan\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes20\",\"name\":\"key\",\"type\":\"bytes20\"},{\"internalType\":\"bytes\",\"name\":\"value\",\"type\":\"bytes\"}],\"name\":\"set\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"start_key\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"size\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"signature\",\"type\":\"uint256\"}],\"name\":\"write\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561000f575f80fd5b506107a28061001d5f395ff3fe608060405234801561000f575f80fd5b5060043610610055575f3560e01c80635acecc78146100595780636531695d14610082578063c315d63e14610097578063d4cd8790146100aa578063d778e2da146100bd575b5f80fd5b61006c6100673660046103c7565b6100d0565b60405161007991906103e7565b60405180910390f35b610095610090366004610433565b61017a565b005b6100956100a5366004610433565b6101e5565b6100956100b8366004610433565b610222565b6100956100cb366004610470565b610295565b6001600160601b031981165f9081526020819052604090208054606091906100f79061052c565b80601f01602080910402602001604051908101604052809291908181526020018280546101239061052c565b801561016e5780601f106101455761010080835404028352916020019161016e565b820191905f5260205f20905b81548152906001019060200180831161015157829003601f168201915b50505050509050919050565b60605f5b838110156101a55761019b6100676101968388610572565b6102bc565b915060010161017e565b5060408051848152602081018490527f2e8128137e55a67bef5f6fa7e5c6722c5632e21b8c8bcf6df64bc32239dd6a3f910160405180910390a150505050565b60605f5b838110156101a557610218610067600183610204888a610572565b61020e9190610585565b6101969190610585565b91506001016101e9565b5f5b828110156102565761024e61023c6101968387610572565b6100cb6102498488610572565b6102cc565b600101610224565b5060408051838152602081018390527fe849f68c74be0ec2d162615e7bc539b752b8e3e7db7ccb69f93eb19c85597f7e910160405180910390a1505050565b6001600160601b031982165f9081526020819052604090206102b782826105e3565b505050565b5f6102c682610379565b92915050565b60408051606480825260a082019092526060916020820181803683370190505090505f5b6064811015610373576040518060c00160405280609681526020016106d7609691398161031e6032866106a3565b6103289190610572565b81518110610338576103386106c2565b602001015160f81c60f81b828281518110610355576103556106c2565b60200101906001600160f81b03191690815f1a9053506001016102f0565b50919050565b5f6020818152815b60148110156103a45760ff84826013036008021c16808284015350600101610381565b5050919050565b80356001600160601b0319811681146103c2575f80fd5b919050565b5f602082840312156103d7575f80fd5b6103e0826103ab565b9392505050565b5f602080835283518060208501525f5b81811015610413578581018301518582016040015282016103f7565b505f604082860101526040601f19601f8301168501019250505092915050565b5f805f60608486031215610445575f80fd5b505081359360208301359350604090920135919050565b634e487b7160e01b5f52604160045260245ffd5b5f8060408385031215610481575f80fd5b61048a836103ab565b9150602083013567ffffffffffffffff808211156104a6575f80fd5b818501915085601f8301126104b9575f80fd5b8135818111156104cb576104cb61045c565b604051601f8201601f19908116603f011681019083821181831017156104f3576104f361045c565b8160405282815288602084870101111561050b575f80fd5b826020860160208301375f6020848301015280955050505050509250929050565b600181811c9082168061054057607f821691505b60208210810361037357634e487b7160e01b5f52602260045260245ffd5b634e487b7160e01b5f52601160045260245ffd5b808201808211156102c6576102c661055e565b818103818111156102c6576102c661055e565b601f8211156102b757805f5260205f20601f840160051c810160208510156105bd5750805b601f840160051c820191505b818110156105dc575f81556001016105c9565b5050505050565b815167ffffffffffffffff8111156105fd576105fd61045c565b6106118161060b845461052c565b84610598565b602080601f831160018114610644575f841561062d5750858301515b5f19600386901b1c1916600185901b17855561069b565b5f85815260208120601f198616915b8281101561067257888601518255948401946001909101908401610653565b508582101561068f57878501515f19600388901b60f8161c191681555b505060018460011b0185555b505050505050565b5f826106bd57634e487b7160e01b5f52601260045260245ffd5b500690565b634e487b7160e01b5f52603260045260245ffdfe6162636465666768696a6b6c6d6e6f707172737475767778792324255e262a28295f2b5b5d7b7d7c3b3a2c2e2f3c3e3f607e6162636465666768696a6b6c6d6e6f707172737475767778792324255e262a28295f2b5b5d7b7d7c3b3a2c2e2f3c3e3f607e6162636465666768696a6b6c6d6e6f707172737475767778792324255e262a28295f2b5b5d7b7d7c3b3a2c2e2f3c3e3f607ea26469706673582212204eae2c718b3baaa920ce69505f212e4769090a30d987d96e454649ba1a92114a64736f6c63430008180033",
}

// IoHeavyTCABI is the input ABI used to generate the binding from.
// Deprecated: Use IoHeavyTCMetaData.ABI instead.
var IoHeavyTCABI = IoHeavyTCMetaData.ABI

// IoHeavyTCBinRuntime is the compiled bytecode used for adding genesis block without deploying code.
const IoHeavyTCBinRuntime = ``

// IoHeavyTCBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use IoHeavyTCMetaData.Bin instead.
var IoHeavyTCBin = IoHeavyTCMetaData.Bin

// DeployIoHeavyTC deploys a new Kaia contract, binding an instance of IoHeavyTC to it.
func DeployIoHeavyTC(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *IoHeavyTC, error) {
	parsed, err := IoHeavyTCMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(IoHeavyTCBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &IoHeavyTC{IoHeavyTCCaller: IoHeavyTCCaller{contract: contract}, IoHeavyTCTransactor: IoHeavyTCTransactor{contract: contract}, IoHeavyTCFilterer: IoHeavyTCFilterer{contract: contract}}, nil
}

// IoHeavyTC is an auto generated Go binding around a Kaia contract.
type IoHeavyTC struct {
	IoHeavyTCCaller     // Read-only binding to the contract
	IoHeavyTCTransactor // Write-only binding to the contract
	IoHeavyTCFilterer   // Log filterer for contract events
}

// IoHeavyTCCaller is an auto generated read-only Go binding around a Kaia contract.
type IoHeavyTCCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IoHeavyTCTransactor is an auto generated write-only Go binding around a Kaia contract.
type IoHeavyTCTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IoHeavyTCFilterer is an auto generated log filtering Go binding around a Kaia contract events.
type IoHeavyTCFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IoHeavyTCSession is an auto generated Go binding around a Kaia contract,
// with pre-set call and transact options.
type IoHeavyTCSession struct {
	Contract     *IoHeavyTC        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IoHeavyTCCallerSession is an auto generated read-only Go binding around a Kaia contract,
// with pre-set call options.
type IoHeavyTCCallerSession struct {
	Contract *IoHeavyTCCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// IoHeavyTCTransactorSession is an auto generated write-only Go binding around a Kaia contract,
// with pre-set transact options.
type IoHeavyTCTransactorSession struct {
	Contract     *IoHeavyTCTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// IoHeavyTCRaw is an auto generated low-level Go binding around a Kaia contract.
type IoHeavyTCRaw struct {
	Contract *IoHeavyTC // Generic contract binding to access the raw methods on
}

// IoHeavyTCCallerRaw is an auto generated low-level read-only Go binding around a Kaia contract.
type IoHeavyTCCallerRaw struct {
	Contract *IoHeavyTCCaller // Generic read-only contract binding to access the raw methods on
}

// IoHeavyTCTransactorRaw is an auto generated low-level write-only Go binding around a Kaia contract.
type IoHeavyTCTransactorRaw struct {
	Contract *IoHeavyTCTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIoHeavyTC creates a new instance of IoHeavyTC, bound to a specific deployed contract.
func NewIoHeavyTC(address common.Address, backend bind.ContractBackend) (*IoHeavyTC, error) {
	contract, err := bindIoHeavyTC(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IoHeavyTC{IoHeavyTCCaller: IoHeavyTCCaller{contract: contract}, IoHeavyTCTransactor: IoHeavyTCTransactor{contract: contract}, IoHeavyTCFilterer: IoHeavyTCFilterer{contract: contract}}, nil
}

// NewIoHeavyTCCaller creates a new read-only instance of IoHeavyTC, bound to a specific deployed contract.
func NewIoHeavyTCCaller(address common.Address, caller bind.ContractCaller) (*IoHeavyTCCaller, error) {
	contract, err := bindIoHeavyTC(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IoHeavyTCCaller{contract: contract}, nil
}

// NewIoHeavyTCTransactor creates a new write-only instance of IoHeavyTC, bound to a specific deployed contract.
func NewIoHeavyTCTransactor(address common.Address, transactor bind.ContractTransactor) (*IoHeavyTCTransactor, error) {
	contract, err := bindIoHeavyTC(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IoHeavyTCTransactor{contract: contract}, nil
}

// NewIoHeavyTCFilterer creates a new log filterer instance of IoHeavyTC, bound to a specific deployed contract.
func NewIoHeavyTCFilterer(address common.Address, filterer bind.ContractFilterer) (*IoHeavyTCFilterer, error) {
	contract, err := bindIoHeavyTC(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IoHeavyTCFilterer{contract: contract}, nil
}

// bindIoHeavyTC binds a generic wrapper to an already deployed contract.
func bindIoHeavyTC(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IoHeavyTCMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IoHeavyTC *IoHeavyTCRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IoHeavyTC.Contract.IoHeavyTCCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IoHeavyTC *IoHeavyTCRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IoHeavyTC.Contract.IoHeavyTCTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IoHeavyTC *IoHeavyTCRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IoHeavyTC.Contract.IoHeavyTCTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IoHeavyTC *IoHeavyTCCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IoHeavyTC.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IoHeavyTC *IoHeavyTCTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IoHeavyTC.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IoHeavyTC *IoHeavyTCTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IoHeavyTC.Contract.contract.Transact(opts, method, params...)
}

// Get is a free data retrieval call binding the contract method 0x5acecc78.
//
// Solidity: function get(bytes20 key) view returns(bytes)
func (_IoHeavyTC *IoHeavyTCCaller) Get(opts *bind.CallOpts, key [20]byte) ([]byte, error) {
	var out []interface{}
	err := _IoHeavyTC.contract.Call(opts, &out, "get", key)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// Get is a free data retrieval call binding the contract method 0x5acecc78.
//
// Solidity: function get(bytes20 key) view returns(bytes)
func (_IoHeavyTC *IoHeavyTCSession) Get(key [20]byte) ([]byte, error) {
	return _IoHeavyTC.Contract.Get(&_IoHeavyTC.CallOpts, key)
}

// Get is a free data retrieval call binding the contract method 0x5acecc78.
//
// Solidity: function get(bytes20 key) view returns(bytes)
func (_IoHeavyTC *IoHeavyTCCallerSession) Get(key [20]byte) ([]byte, error) {
	return _IoHeavyTC.Contract.Get(&_IoHeavyTC.CallOpts, key)
}

// RevertScan is a paid mutator transaction binding the contract method 0xc315d63e.
//
// Solidity: function revert_scan(uint256 start_key, uint256 size, uint256 signature) returns()
func (_IoHeavyTC *IoHeavyTCTransactor) RevertScan(opts *bind.TransactOpts, start_key *big.Int, size *big.Int, signature *big.Int) (*types.Transaction, error) {
	return _IoHeavyTC.contract.Transact(opts, "revert_scan", start_key, size, signature)
}

// RevertScan is a paid mutator transaction binding the contract method 0xc315d63e.
//
// Solidity: function revert_scan(uint256 start_key, uint256 size, uint256 signature) returns()
func (_IoHeavyTC *IoHeavyTCSession) RevertScan(start_key *big.Int, size *big.Int, signature *big.Int) (*types.Transaction, error) {
	return _IoHeavyTC.Contract.RevertScan(&_IoHeavyTC.TransactOpts, start_key, size, signature)
}

// RevertScan is a paid mutator transaction binding the contract method 0xc315d63e.
//
// Solidity: function revert_scan(uint256 start_key, uint256 size, uint256 signature) returns()
func (_IoHeavyTC *IoHeavyTCTransactorSession) RevertScan(start_key *big.Int, size *big.Int, signature *big.Int) (*types.Transaction, error) {
	return _IoHeavyTC.Contract.RevertScan(&_IoHeavyTC.TransactOpts, start_key, size, signature)
}

// Scan is a paid mutator transaction binding the contract method 0x6531695d.
//
// Solidity: function scan(uint256 start_key, uint256 size, uint256 signature) returns()
func (_IoHeavyTC *IoHeavyTCTransactor) Scan(opts *bind.TransactOpts, start_key *big.Int, size *big.Int, signature *big.Int) (*types.Transaction, error) {
	return _IoHeavyTC.contract.Transact(opts, "scan", start_key, size, signature)
}

// Scan is a paid mutator transaction binding the contract method 0x6531695d.
//
// Solidity: function scan(uint256 start_key, uint256 size, uint256 signature) returns()
func (_IoHeavyTC *IoHeavyTCSession) Scan(start_key *big.Int, size *big.Int, signature *big.Int) (*types.Transaction, error) {
	return _IoHeavyTC.Contract.Scan(&_IoHeavyTC.TransactOpts, start_key, size, signature)
}

// Scan is a paid mutator transaction binding the contract method 0x6531695d.
//
// Solidity: function scan(uint256 start_key, uint256 size, uint256 signature) returns()
func (_IoHeavyTC *IoHeavyTCTransactorSession) Scan(start_key *big.Int, size *big.Int, signature *big.Int) (*types.Transaction, error) {
	return _IoHeavyTC.Contract.Scan(&_IoHeavyTC.TransactOpts, start_key, size, signature)
}

// Set is a paid mutator transaction binding the contract method 0xd778e2da.
//
// Solidity: function set(bytes20 key, bytes value) returns()
func (_IoHeavyTC *IoHeavyTCTransactor) Set(opts *bind.TransactOpts, key [20]byte, value []byte) (*types.Transaction, error) {
	return _IoHeavyTC.contract.Transact(opts, "set", key, value)
}

// Set is a paid mutator transaction binding the contract method 0xd778e2da.
//
// Solidity: function set(bytes20 key, bytes value) returns()
func (_IoHeavyTC *IoHeavyTCSession) Set(key [20]byte, value []byte) (*types.Transaction, error) {
	return _IoHeavyTC.Contract.Set(&_IoHeavyTC.TransactOpts, key, value)
}

// Set is a paid mutator transaction binding the contract method 0xd778e2da.
//
// Solidity: function set(bytes20 key, bytes value) returns()
func (_IoHeavyTC *IoHeavyTCTransactorSession) Set(key [20]byte, value []byte) (*types.Transaction, error) {
	return _IoHeavyTC.Contract.Set(&_IoHeavyTC.TransactOpts, key, value)
}

// Write is a paid mutator transaction binding the contract method 0xd4cd8790.
//
// Solidity: function write(uint256 start_key, uint256 size, uint256 signature) returns()
func (_IoHeavyTC *IoHeavyTCTransactor) Write(opts *bind.TransactOpts, start_key *big.Int, size *big.Int, signature *big.Int) (*types.Transaction, error) {
	return _IoHeavyTC.contract.Transact(opts, "write", start_key, size, signature)
}

// Write is a paid mutator transaction binding the contract method 0xd4cd8790.
//
// Solidity: function write(uint256 start_key, uint256 size, uint256 signature) returns()
func (_IoHeavyTC *IoHeavyTCSession) Write(start_key *big.Int, size *big.Int, signature *big.Int) (*types.Transaction, error) {
	return _IoHeavyTC.Contract.Write(&_IoHeavyTC.TransactOpts, start_key, size, signature)
}

// Write is a paid mutator transaction binding the contract method 0xd4cd8790.
//
// Solidity: function write(uint256 start_key, uint256 size, uint256 signature) returns()
func (_IoHeavyTC *IoHeavyTCTransactorSession) Write(start_key *big.Int, size *big.Int, signature *big.Int) (*types.Transaction, error) {
	return _IoHeavyTC.Contract.Write(&_IoHeavyTC.TransactOpts, start_key, size, signature)
}

// IoHeavyTCFinishScanIterator is returned from FilterFinishScan and is used to iterate over the raw logs and unpacked data for FinishScan events raised by the IoHeavyTC contract.
type IoHeavyTCFinishScanIterator struct {
	Event *IoHeavyTCFinishScan // Event containing the contract specifics and raw log

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
func (it *IoHeavyTCFinishScanIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IoHeavyTCFinishScan)
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
		it.Event = new(IoHeavyTCFinishScan)
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
func (it *IoHeavyTCFinishScanIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IoHeavyTCFinishScanIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IoHeavyTCFinishScan represents a FinishScan event raised by the IoHeavyTC contract.
type IoHeavyTCFinishScan struct {
	Size      *big.Int
	Signature *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterFinishScan is a free log retrieval operation binding the contract event 0x2e8128137e55a67bef5f6fa7e5c6722c5632e21b8c8bcf6df64bc32239dd6a3f.
//
// Solidity: event finishScan(uint256 size, uint256 signature)
func (_IoHeavyTC *IoHeavyTCFilterer) FilterFinishScan(opts *bind.FilterOpts) (*IoHeavyTCFinishScanIterator, error) {

	logs, sub, err := _IoHeavyTC.contract.FilterLogs(opts, "finishScan")
	if err != nil {
		return nil, err
	}
	return &IoHeavyTCFinishScanIterator{contract: _IoHeavyTC.contract, event: "finishScan", logs: logs, sub: sub}, nil
}

// WatchFinishScan is a free log subscription operation binding the contract event 0x2e8128137e55a67bef5f6fa7e5c6722c5632e21b8c8bcf6df64bc32239dd6a3f.
//
// Solidity: event finishScan(uint256 size, uint256 signature)
func (_IoHeavyTC *IoHeavyTCFilterer) WatchFinishScan(opts *bind.WatchOpts, sink chan<- *IoHeavyTCFinishScan) (event.Subscription, error) {

	logs, sub, err := _IoHeavyTC.contract.WatchLogs(opts, "finishScan")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IoHeavyTCFinishScan)
				if err := _IoHeavyTC.contract.UnpackLog(event, "finishScan", log); err != nil {
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

// ParseFinishScan is a log parse operation binding the contract event 0x2e8128137e55a67bef5f6fa7e5c6722c5632e21b8c8bcf6df64bc32239dd6a3f.
//
// Solidity: event finishScan(uint256 size, uint256 signature)
func (_IoHeavyTC *IoHeavyTCFilterer) ParseFinishScan(log types.Log) (*IoHeavyTCFinishScan, error) {
	event := new(IoHeavyTCFinishScan)
	if err := _IoHeavyTC.contract.UnpackLog(event, "finishScan", log); err != nil {
		return nil, err
	}
	return event, nil
}

// IoHeavyTCFinishWriteIterator is returned from FilterFinishWrite and is used to iterate over the raw logs and unpacked data for FinishWrite events raised by the IoHeavyTC contract.
type IoHeavyTCFinishWriteIterator struct {
	Event *IoHeavyTCFinishWrite // Event containing the contract specifics and raw log

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
func (it *IoHeavyTCFinishWriteIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IoHeavyTCFinishWrite)
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
		it.Event = new(IoHeavyTCFinishWrite)
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
func (it *IoHeavyTCFinishWriteIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IoHeavyTCFinishWriteIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IoHeavyTCFinishWrite represents a FinishWrite event raised by the IoHeavyTC contract.
type IoHeavyTCFinishWrite struct {
	Size      *big.Int
	Signature *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterFinishWrite is a free log retrieval operation binding the contract event 0xe849f68c74be0ec2d162615e7bc539b752b8e3e7db7ccb69f93eb19c85597f7e.
//
// Solidity: event finishWrite(uint256 size, uint256 signature)
func (_IoHeavyTC *IoHeavyTCFilterer) FilterFinishWrite(opts *bind.FilterOpts) (*IoHeavyTCFinishWriteIterator, error) {

	logs, sub, err := _IoHeavyTC.contract.FilterLogs(opts, "finishWrite")
	if err != nil {
		return nil, err
	}
	return &IoHeavyTCFinishWriteIterator{contract: _IoHeavyTC.contract, event: "finishWrite", logs: logs, sub: sub}, nil
}

// WatchFinishWrite is a free log subscription operation binding the contract event 0xe849f68c74be0ec2d162615e7bc539b752b8e3e7db7ccb69f93eb19c85597f7e.
//
// Solidity: event finishWrite(uint256 size, uint256 signature)
func (_IoHeavyTC *IoHeavyTCFilterer) WatchFinishWrite(opts *bind.WatchOpts, sink chan<- *IoHeavyTCFinishWrite) (event.Subscription, error) {

	logs, sub, err := _IoHeavyTC.contract.WatchLogs(opts, "finishWrite")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IoHeavyTCFinishWrite)
				if err := _IoHeavyTC.contract.UnpackLog(event, "finishWrite", log); err != nil {
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

// ParseFinishWrite is a log parse operation binding the contract event 0xe849f68c74be0ec2d162615e7bc539b752b8e3e7db7ccb69f93eb19c85597f7e.
//
// Solidity: event finishWrite(uint256 size, uint256 signature)
func (_IoHeavyTC *IoHeavyTCFilterer) ParseFinishWrite(log types.Log) (*IoHeavyTCFinishWrite, error) {
	event := new(IoHeavyTCFinishWrite)
	if err := _IoHeavyTC.contract.UnpackLog(event, "finishWrite", log); err != nil {
		return nil, err
	}
	return event, nil
}
