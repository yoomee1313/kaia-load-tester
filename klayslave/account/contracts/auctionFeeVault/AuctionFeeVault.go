// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package auction

import (
	"errors"
	"math/big"
	"strings"

	"github.com/kaiachain/kaia"
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

// AuctionErrorMetaData contains all meta data concerning the AuctionError contract.
var AuctionErrorMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"EmptyDepositVault\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInput\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MinDepositNotOver\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyAuctionDepositVault\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyEntryPoint\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyProposer\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyStakingAdmin\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"WithdrawReservationExists\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"WithdrawalFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"WithdrawalNotAllowedYet\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroDepositAmount\",\"type\":\"error\"}]",
	Bin: "0x6080604052348015600e575f80fd5b50603e80601a5f395ff3fe60806040525f80fdfea26469706673582212207528d0d43ae476e30651a0be079c16a1af03532761e8a907f228cacf87ea7d1264736f6c63430008190033",
}

// AuctionErrorABI is the input ABI used to generate the binding from.
// Deprecated: Use AuctionErrorMetaData.ABI instead.
var AuctionErrorABI = AuctionErrorMetaData.ABI

// AuctionErrorBinRuntime is the compiled bytecode used for adding genesis block without deploying code.
const AuctionErrorBinRuntime = `60806040525f80fdfea26469706673582212207528d0d43ae476e30651a0be079c16a1af03532761e8a907f228cacf87ea7d1264736f6c63430008190033`

// AuctionErrorBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use AuctionErrorMetaData.Bin instead.
var AuctionErrorBin = AuctionErrorMetaData.Bin

// DeployAuctionError deploys a new Kaia contract, binding an instance of AuctionError to it.
func DeployAuctionError(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *AuctionError, error) {
	parsed, err := AuctionErrorMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(AuctionErrorBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &AuctionError{AuctionErrorCaller: AuctionErrorCaller{contract: contract}, AuctionErrorTransactor: AuctionErrorTransactor{contract: contract}, AuctionErrorFilterer: AuctionErrorFilterer{contract: contract}}, nil
}

// AuctionError is an auto generated Go binding around a Kaia contract.
type AuctionError struct {
	AuctionErrorCaller     // Read-only binding to the contract
	AuctionErrorTransactor // Write-only binding to the contract
	AuctionErrorFilterer   // Log filterer for contract events
}

// AuctionErrorCaller is an auto generated read-only Go binding around a Kaia contract.
type AuctionErrorCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AuctionErrorTransactor is an auto generated write-only Go binding around a Kaia contract.
type AuctionErrorTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AuctionErrorFilterer is an auto generated log filtering Go binding around a Kaia contract events.
type AuctionErrorFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AuctionErrorSession is an auto generated Go binding around a Kaia contract,
// with pre-set call and transact options.
type AuctionErrorSession struct {
	Contract     *AuctionError     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AuctionErrorCallerSession is an auto generated read-only Go binding around a Kaia contract,
// with pre-set call options.
type AuctionErrorCallerSession struct {
	Contract *AuctionErrorCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// AuctionErrorTransactorSession is an auto generated write-only Go binding around a Kaia contract,
// with pre-set transact options.
type AuctionErrorTransactorSession struct {
	Contract     *AuctionErrorTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// AuctionErrorRaw is an auto generated low-level Go binding around a Kaia contract.
type AuctionErrorRaw struct {
	Contract *AuctionError // Generic contract binding to access the raw methods on
}

// AuctionErrorCallerRaw is an auto generated low-level read-only Go binding around a Kaia contract.
type AuctionErrorCallerRaw struct {
	Contract *AuctionErrorCaller // Generic read-only contract binding to access the raw methods on
}

// AuctionErrorTransactorRaw is an auto generated low-level write-only Go binding around a Kaia contract.
type AuctionErrorTransactorRaw struct {
	Contract *AuctionErrorTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAuctionError creates a new instance of AuctionError, bound to a specific deployed contract.
func NewAuctionError(address common.Address, backend bind.ContractBackend) (*AuctionError, error) {
	contract, err := bindAuctionError(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AuctionError{AuctionErrorCaller: AuctionErrorCaller{contract: contract}, AuctionErrorTransactor: AuctionErrorTransactor{contract: contract}, AuctionErrorFilterer: AuctionErrorFilterer{contract: contract}}, nil
}

// NewAuctionErrorCaller creates a new read-only instance of AuctionError, bound to a specific deployed contract.
func NewAuctionErrorCaller(address common.Address, caller bind.ContractCaller) (*AuctionErrorCaller, error) {
	contract, err := bindAuctionError(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AuctionErrorCaller{contract: contract}, nil
}

// NewAuctionErrorTransactor creates a new write-only instance of AuctionError, bound to a specific deployed contract.
func NewAuctionErrorTransactor(address common.Address, transactor bind.ContractTransactor) (*AuctionErrorTransactor, error) {
	contract, err := bindAuctionError(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AuctionErrorTransactor{contract: contract}, nil
}

// NewAuctionErrorFilterer creates a new log filterer instance of AuctionError, bound to a specific deployed contract.
func NewAuctionErrorFilterer(address common.Address, filterer bind.ContractFilterer) (*AuctionErrorFilterer, error) {
	contract, err := bindAuctionError(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AuctionErrorFilterer{contract: contract}, nil
}

// bindAuctionError binds a generic wrapper to an already deployed contract.
func bindAuctionError(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := AuctionErrorMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AuctionError *AuctionErrorRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AuctionError.Contract.AuctionErrorCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AuctionError *AuctionErrorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AuctionError.Contract.AuctionErrorTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AuctionError *AuctionErrorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AuctionError.Contract.AuctionErrorTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AuctionError *AuctionErrorCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AuctionError.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AuctionError *AuctionErrorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AuctionError.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AuctionError *AuctionErrorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AuctionError.Contract.contract.Transact(opts, method, params...)
}

// AuctionFeeVaultMetaData contains all meta data concerning the AuctionFeeVault contract.
var AuctionFeeVaultMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"initialOwner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_searcherPaybackRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_validatorPaybackRate\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"EmptyDepositVault\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInput\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MinDepositNotOver\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyAuctionDepositVault\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyEntryPoint\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyProposer\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyStakingAdmin\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"WithdrawReservationExists\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"WithdrawalFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"WithdrawalNotAllowedYet\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroDepositAmount\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"paybackAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"validatorPaybackAmount\",\"type\":\"uint256\"}],\"name\":\"FeeDeposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"FeeWithdrawal\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"nodeId\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"reward\",\"type\":\"address\"}],\"name\":\"RewardAddressRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"searcherPaybackRate\",\"type\":\"uint256\"}],\"name\":\"SearcherPaybackRateUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"validatorPaybackRate\",\"type\":\"uint256\"}],\"name\":\"ValidatorPaybackRateUpdated\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"ADDRESS_BOOK\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_PAYBACK_RATE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"accumulatedBids\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"nodeId\",\"type\":\"address\"}],\"name\":\"getRewardAddr\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"nodeId\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"rewardAddr\",\"type\":\"address\"}],\"name\":\"registerRewardAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"searcherPaybackRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_searcherPaybackRate\",\"type\":\"uint256\"}],\"name\":\"setSearcherPaybackRate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_validatorPaybackRate\",\"type\":\"uint256\"}],\"name\":\"setValidatorPaybackRate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"searcher\",\"type\":\"address\"}],\"name\":\"takeBid\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"validatorPaybackRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"0ccfe3e2": "ADDRESS_BOOK()",
		"c5c1fc04": "MAX_PAYBACK_RATE()",
		"49723142": "accumulatedBids()",
		"27a50f72": "getRewardAddr(address)",
		"8da5cb5b": "owner()",
		"363d5183": "registerRewardAddress(address,address)",
		"715018a6": "renounceOwnership()",
		"3c702fbd": "searcherPaybackRate()",
		"36cf2c63": "setSearcherPaybackRate(uint256)",
		"11062696": "setValidatorPaybackRate(uint256)",
		"8573e2ff": "takeBid(address)",
		"f2fde38b": "transferOwnership(address)",
		"89b703aa": "validatorPaybackRate()",
		"51cff8d9": "withdraw(address)",
	},
	Bin: "0x608060405234801561000f575f80fd5b50604051610ac5380380610ac583398101604081905261002e91610104565b826001600160a01b03811661005c57604051631e4fbdf760e01b81525f600482015260240160405180910390fd5b6100658161009c565b5061007082826100eb565b61008d5760405163b4fa3fb360e01b815260040160405180910390fd5b60029190915560015550610162565b5f80546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b5f6127106100f98385610143565b111590505b92915050565b5f805f60608486031215610116575f80fd5b83516001600160a01b038116811461012c575f80fd5b602085015160409095015190969495509392505050565b808201808211156100fe57634e487b7160e01b5f52601160045260245ffd5b6109568061016f5f395ff3fe6080604052600436106100d9575f3560e01c806351cff8d91161007c57806389b703aa1161005757806389b703aa146102235780638da5cb5b14610238578063c5c1fc0414610254578063f2fde38b14610269575f80fd5b806351cff8d9146101dd578063715018a6146101fc5780638573e2ff14610210575f80fd5b8063363d5183116100b7578063363d51831461016757806336cf2c63146101865780633c702fbd146101a557806349723142146101c8575f80fd5b80630ccfe3e2146100dd578063110626961461010f57806327a50f7214610130575b5f80fd5b3480156100e8575f80fd5b506100f261040081565b6040516001600160a01b0390911681526020015b60405180910390f35b34801561011a575f80fd5b5061012e6101293660046107d6565b610288565b005b34801561013b575f80fd5b506100f261014a366004610801565b6001600160a01b039081165f908152600460205260409020541690565b348015610172575f80fd5b5061012e610181366004610823565b6102f5565b348015610191575f80fd5b5061012e6101a03660046107d6565b61043e565b3480156101b0575f80fd5b506101ba60025481565b604051908152602001610106565b3480156101d3575f80fd5b506101ba60035481565b3480156101e8575f80fd5b5061012e6101f7366004610801565b6104a4565b348015610207575f80fd5b5061012e610555565b61012e61021e366004610801565b610568565b34801561022e575f80fd5b506101ba60015481565b348015610243575f80fd5b505f546001600160a01b03166100f2565b34801561025f575f80fd5b506101ba61271081565b348015610274575f80fd5b5061012e610283366004610801565b610700565b610290610742565b61029c8160025461076e565b6102b95760405163b4fa3fb360e01b815260040160405180910390fd5b60018190556040518181527f5309d48fe743a67ce32d8f66af9e2388d65bfc8cc026a4e1fbed3a4612a0af98906020015b60405180910390a150565b604051630aabaead60e11b81526001600160a01b03831660048201525f90610400906315575d5a90602401606060405180830381865afa15801561033b573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061035f919061085a565b50604051630935e01b60e21b81523360048201529092506001600160a01b03831691506324d7806c90602401602060405180830381865afa1580156103a6573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906103ca91906108a4565b6103e757604051632b14b52960e01b815260040160405180910390fd5b6001600160a01b038381165f8181526004602052604080822080546001600160a01b0319169487169485179055517fe608476cb01b1d04f944f0fdb25841b1f483d26965d42c4a1fab67b8b1488b3b9190a3505050565b610446610742565b6104528160015461076e565b61046f5760405163b4fa3fb360e01b815260040160405180910390fd5b60028190556040518181527f71745430318b073bd776904f2432cb283ce3d2ded537bafe2640cf4d6e4bc64f906020016102ea565b6104ac610742565b60405147905f906001600160a01b0384169083908381818185875af1925050503d805f81146104f6576040519150601f19603f3d011682016040523d82523d5f602084013e6104fb565b606091505b505090508061051d576040516327fcd9d160e01b815260040160405180910390fd5b6040518281527f706d7f48c702007c2fb0881cea5759732e64f52faee427d5ab030787cfb7d7879060200160405180910390a1505050565b61055d610742565b6105665f610787565b565b345f036105725750565b5f3490508060035f82825461058791906108d7565b90915550506002545f906127109061059f90846108ea565b6105a99190610901565b9050801561060e575f836001600160a01b0316826040515f6040518083038185875af1925050503d805f81146105fa576040519150601f19603f3d011682016040523d82523d5f602084013e6105ff565b606091505b505090508061060c575f91505b505b5f6127106001548461062091906108ea565b61062a9190610901565b905080156106b757415f908152600460205260409020546001600160a01b031680156106b1575f816001600160a01b0316836040515f6040518083038185875af1925050503d805f8114610699576040519150601f19603f3d011682016040523d82523d5f602084013e61069e565b606091505b50509050806106ab575f92505b506106b5565b5f91505b505b604080518481526020810184905290810182905241907fa34c9ef6ada915fef21639b2d5c085580cf79046cca66c2c2e8b87e2f3bd85679060600160405180910390a250505050565b610708610742565b6001600160a01b03811661073657604051631e4fbdf760e01b81525f60048201526024015b60405180910390fd5b61073f81610787565b50565b5f546001600160a01b031633146105665760405163118cdaa760e01b815233600482015260240161072d565b5f61271061077c83856108d7565b111590505b92915050565b5f80546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b5f602082840312156107e6575f80fd5b5035919050565b6001600160a01b038116811461073f575f80fd5b5f60208284031215610811575f80fd5b813561081c816107ed565b9392505050565b5f8060408385031215610834575f80fd5b823561083f816107ed565b9150602083013561084f816107ed565b809150509250929050565b5f805f6060848603121561086c575f80fd5b8351610877816107ed565b6020850151909350610888816107ed565b6040850151909250610899816107ed565b809150509250925092565b5f602082840312156108b4575f80fd5b8151801515811461081c575f80fd5b634e487b7160e01b5f52601160045260245ffd5b80820180821115610781576107816108c3565b8082028115828204841417610781576107816108c3565b5f8261091b57634e487b7160e01b5f52601260045260245ffd5b50049056fea2646970667358221220822abca0096ba9dda29726159d7bff58484e125bfa9067c91396d99c519e807964736f6c63430008190033",
}

// AuctionFeeVaultABI is the input ABI used to generate the binding from.
// Deprecated: Use AuctionFeeVaultMetaData.ABI instead.
var AuctionFeeVaultABI = AuctionFeeVaultMetaData.ABI

// AuctionFeeVaultBinRuntime is the compiled bytecode used for adding genesis block without deploying code.
const AuctionFeeVaultBinRuntime = `6080604052600436106100d9575f3560e01c806351cff8d91161007c57806389b703aa1161005757806389b703aa146102235780638da5cb5b14610238578063c5c1fc0414610254578063f2fde38b14610269575f80fd5b806351cff8d9146101dd578063715018a6146101fc5780638573e2ff14610210575f80fd5b8063363d5183116100b7578063363d51831461016757806336cf2c63146101865780633c702fbd146101a557806349723142146101c8575f80fd5b80630ccfe3e2146100dd578063110626961461010f57806327a50f7214610130575b5f80fd5b3480156100e8575f80fd5b506100f261040081565b6040516001600160a01b0390911681526020015b60405180910390f35b34801561011a575f80fd5b5061012e6101293660046107d6565b610288565b005b34801561013b575f80fd5b506100f261014a366004610801565b6001600160a01b039081165f908152600460205260409020541690565b348015610172575f80fd5b5061012e610181366004610823565b6102f5565b348015610191575f80fd5b5061012e6101a03660046107d6565b61043e565b3480156101b0575f80fd5b506101ba60025481565b604051908152602001610106565b3480156101d3575f80fd5b506101ba60035481565b3480156101e8575f80fd5b5061012e6101f7366004610801565b6104a4565b348015610207575f80fd5b5061012e610555565b61012e61021e366004610801565b610568565b34801561022e575f80fd5b506101ba60015481565b348015610243575f80fd5b505f546001600160a01b03166100f2565b34801561025f575f80fd5b506101ba61271081565b348015610274575f80fd5b5061012e610283366004610801565b610700565b610290610742565b61029c8160025461076e565b6102b95760405163b4fa3fb360e01b815260040160405180910390fd5b60018190556040518181527f5309d48fe743a67ce32d8f66af9e2388d65bfc8cc026a4e1fbed3a4612a0af98906020015b60405180910390a150565b604051630aabaead60e11b81526001600160a01b03831660048201525f90610400906315575d5a90602401606060405180830381865afa15801561033b573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061035f919061085a565b50604051630935e01b60e21b81523360048201529092506001600160a01b03831691506324d7806c90602401602060405180830381865afa1580156103a6573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906103ca91906108a4565b6103e757604051632b14b52960e01b815260040160405180910390fd5b6001600160a01b038381165f8181526004602052604080822080546001600160a01b0319169487169485179055517fe608476cb01b1d04f944f0fdb25841b1f483d26965d42c4a1fab67b8b1488b3b9190a3505050565b610446610742565b6104528160015461076e565b61046f5760405163b4fa3fb360e01b815260040160405180910390fd5b60028190556040518181527f71745430318b073bd776904f2432cb283ce3d2ded537bafe2640cf4d6e4bc64f906020016102ea565b6104ac610742565b60405147905f906001600160a01b0384169083908381818185875af1925050503d805f81146104f6576040519150601f19603f3d011682016040523d82523d5f602084013e6104fb565b606091505b505090508061051d576040516327fcd9d160e01b815260040160405180910390fd5b6040518281527f706d7f48c702007c2fb0881cea5759732e64f52faee427d5ab030787cfb7d7879060200160405180910390a1505050565b61055d610742565b6105665f610787565b565b345f036105725750565b5f3490508060035f82825461058791906108d7565b90915550506002545f906127109061059f90846108ea565b6105a99190610901565b9050801561060e575f836001600160a01b0316826040515f6040518083038185875af1925050503d805f81146105fa576040519150601f19603f3d011682016040523d82523d5f602084013e6105ff565b606091505b505090508061060c575f91505b505b5f6127106001548461062091906108ea565b61062a9190610901565b905080156106b757415f908152600460205260409020546001600160a01b031680156106b1575f816001600160a01b0316836040515f6040518083038185875af1925050503d805f8114610699576040519150601f19603f3d011682016040523d82523d5f602084013e61069e565b606091505b50509050806106ab575f92505b506106b5565b5f91505b505b604080518481526020810184905290810182905241907fa34c9ef6ada915fef21639b2d5c085580cf79046cca66c2c2e8b87e2f3bd85679060600160405180910390a250505050565b610708610742565b6001600160a01b03811661073657604051631e4fbdf760e01b81525f60048201526024015b60405180910390fd5b61073f81610787565b50565b5f546001600160a01b031633146105665760405163118cdaa760e01b815233600482015260240161072d565b5f61271061077c83856108d7565b111590505b92915050565b5f80546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b5f602082840312156107e6575f80fd5b5035919050565b6001600160a01b038116811461073f575f80fd5b5f60208284031215610811575f80fd5b813561081c816107ed565b9392505050565b5f8060408385031215610834575f80fd5b823561083f816107ed565b9150602083013561084f816107ed565b809150509250929050565b5f805f6060848603121561086c575f80fd5b8351610877816107ed565b6020850151909350610888816107ed565b6040850151909250610899816107ed565b809150509250925092565b5f602082840312156108b4575f80fd5b8151801515811461081c575f80fd5b634e487b7160e01b5f52601160045260245ffd5b80820180821115610781576107816108c3565b8082028115828204841417610781576107816108c3565b5f8261091b57634e487b7160e01b5f52601260045260245ffd5b50049056fea2646970667358221220822abca0096ba9dda29726159d7bff58484e125bfa9067c91396d99c519e807964736f6c63430008190033`

// Deprecated: Use AuctionFeeVaultMetaData.Sigs instead.
// AuctionFeeVaultFuncSigs maps the 4-byte function signature to its string representation.
var AuctionFeeVaultFuncSigs = AuctionFeeVaultMetaData.Sigs

// AuctionFeeVaultBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use AuctionFeeVaultMetaData.Bin instead.
var AuctionFeeVaultBin = AuctionFeeVaultMetaData.Bin

// DeployAuctionFeeVault deploys a new Kaia contract, binding an instance of AuctionFeeVault to it.
func DeployAuctionFeeVault(auth *bind.TransactOpts, backend bind.ContractBackend, initialOwner common.Address, _searcherPaybackRate *big.Int, _validatorPaybackRate *big.Int) (common.Address, *types.Transaction, *AuctionFeeVault, error) {
	parsed, err := AuctionFeeVaultMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(AuctionFeeVaultBin), backend, initialOwner, _searcherPaybackRate, _validatorPaybackRate)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &AuctionFeeVault{AuctionFeeVaultCaller: AuctionFeeVaultCaller{contract: contract}, AuctionFeeVaultTransactor: AuctionFeeVaultTransactor{contract: contract}, AuctionFeeVaultFilterer: AuctionFeeVaultFilterer{contract: contract}}, nil
}

// AuctionFeeVault is an auto generated Go binding around a Kaia contract.
type AuctionFeeVault struct {
	AuctionFeeVaultCaller     // Read-only binding to the contract
	AuctionFeeVaultTransactor // Write-only binding to the contract
	AuctionFeeVaultFilterer   // Log filterer for contract events
}

// AuctionFeeVaultCaller is an auto generated read-only Go binding around a Kaia contract.
type AuctionFeeVaultCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AuctionFeeVaultTransactor is an auto generated write-only Go binding around a Kaia contract.
type AuctionFeeVaultTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AuctionFeeVaultFilterer is an auto generated log filtering Go binding around a Kaia contract events.
type AuctionFeeVaultFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AuctionFeeVaultSession is an auto generated Go binding around a Kaia contract,
// with pre-set call and transact options.
type AuctionFeeVaultSession struct {
	Contract     *AuctionFeeVault  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AuctionFeeVaultCallerSession is an auto generated read-only Go binding around a Kaia contract,
// with pre-set call options.
type AuctionFeeVaultCallerSession struct {
	Contract *AuctionFeeVaultCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// AuctionFeeVaultTransactorSession is an auto generated write-only Go binding around a Kaia contract,
// with pre-set transact options.
type AuctionFeeVaultTransactorSession struct {
	Contract     *AuctionFeeVaultTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// AuctionFeeVaultRaw is an auto generated low-level Go binding around a Kaia contract.
type AuctionFeeVaultRaw struct {
	Contract *AuctionFeeVault // Generic contract binding to access the raw methods on
}

// AuctionFeeVaultCallerRaw is an auto generated low-level read-only Go binding around a Kaia contract.
type AuctionFeeVaultCallerRaw struct {
	Contract *AuctionFeeVaultCaller // Generic read-only contract binding to access the raw methods on
}

// AuctionFeeVaultTransactorRaw is an auto generated low-level write-only Go binding around a Kaia contract.
type AuctionFeeVaultTransactorRaw struct {
	Contract *AuctionFeeVaultTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAuctionFeeVault creates a new instance of AuctionFeeVault, bound to a specific deployed contract.
func NewAuctionFeeVault(address common.Address, backend bind.ContractBackend) (*AuctionFeeVault, error) {
	contract, err := bindAuctionFeeVault(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AuctionFeeVault{AuctionFeeVaultCaller: AuctionFeeVaultCaller{contract: contract}, AuctionFeeVaultTransactor: AuctionFeeVaultTransactor{contract: contract}, AuctionFeeVaultFilterer: AuctionFeeVaultFilterer{contract: contract}}, nil
}

// NewAuctionFeeVaultCaller creates a new read-only instance of AuctionFeeVault, bound to a specific deployed contract.
func NewAuctionFeeVaultCaller(address common.Address, caller bind.ContractCaller) (*AuctionFeeVaultCaller, error) {
	contract, err := bindAuctionFeeVault(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AuctionFeeVaultCaller{contract: contract}, nil
}

// NewAuctionFeeVaultTransactor creates a new write-only instance of AuctionFeeVault, bound to a specific deployed contract.
func NewAuctionFeeVaultTransactor(address common.Address, transactor bind.ContractTransactor) (*AuctionFeeVaultTransactor, error) {
	contract, err := bindAuctionFeeVault(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AuctionFeeVaultTransactor{contract: contract}, nil
}

// NewAuctionFeeVaultFilterer creates a new log filterer instance of AuctionFeeVault, bound to a specific deployed contract.
func NewAuctionFeeVaultFilterer(address common.Address, filterer bind.ContractFilterer) (*AuctionFeeVaultFilterer, error) {
	contract, err := bindAuctionFeeVault(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AuctionFeeVaultFilterer{contract: contract}, nil
}

// bindAuctionFeeVault binds a generic wrapper to an already deployed contract.
func bindAuctionFeeVault(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := AuctionFeeVaultMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AuctionFeeVault *AuctionFeeVaultRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AuctionFeeVault.Contract.AuctionFeeVaultCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AuctionFeeVault *AuctionFeeVaultRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AuctionFeeVault.Contract.AuctionFeeVaultTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AuctionFeeVault *AuctionFeeVaultRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AuctionFeeVault.Contract.AuctionFeeVaultTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AuctionFeeVault *AuctionFeeVaultCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AuctionFeeVault.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AuctionFeeVault *AuctionFeeVaultTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AuctionFeeVault.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AuctionFeeVault *AuctionFeeVaultTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AuctionFeeVault.Contract.contract.Transact(opts, method, params...)
}

// ADDRESSBOOK is a free data retrieval call binding the contract method 0x0ccfe3e2.
//
// Solidity: function ADDRESS_BOOK() view returns(address)
func (_AuctionFeeVault *AuctionFeeVaultCaller) ADDRESSBOOK(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AuctionFeeVault.contract.Call(opts, &out, "ADDRESS_BOOK")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ADDRESSBOOK is a free data retrieval call binding the contract method 0x0ccfe3e2.
//
// Solidity: function ADDRESS_BOOK() view returns(address)
func (_AuctionFeeVault *AuctionFeeVaultSession) ADDRESSBOOK() (common.Address, error) {
	return _AuctionFeeVault.Contract.ADDRESSBOOK(&_AuctionFeeVault.CallOpts)
}

// ADDRESSBOOK is a free data retrieval call binding the contract method 0x0ccfe3e2.
//
// Solidity: function ADDRESS_BOOK() view returns(address)
func (_AuctionFeeVault *AuctionFeeVaultCallerSession) ADDRESSBOOK() (common.Address, error) {
	return _AuctionFeeVault.Contract.ADDRESSBOOK(&_AuctionFeeVault.CallOpts)
}

// MAXPAYBACKRATE is a free data retrieval call binding the contract method 0xc5c1fc04.
//
// Solidity: function MAX_PAYBACK_RATE() view returns(uint256)
func (_AuctionFeeVault *AuctionFeeVaultCaller) MAXPAYBACKRATE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AuctionFeeVault.contract.Call(opts, &out, "MAX_PAYBACK_RATE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXPAYBACKRATE is a free data retrieval call binding the contract method 0xc5c1fc04.
//
// Solidity: function MAX_PAYBACK_RATE() view returns(uint256)
func (_AuctionFeeVault *AuctionFeeVaultSession) MAXPAYBACKRATE() (*big.Int, error) {
	return _AuctionFeeVault.Contract.MAXPAYBACKRATE(&_AuctionFeeVault.CallOpts)
}

// MAXPAYBACKRATE is a free data retrieval call binding the contract method 0xc5c1fc04.
//
// Solidity: function MAX_PAYBACK_RATE() view returns(uint256)
func (_AuctionFeeVault *AuctionFeeVaultCallerSession) MAXPAYBACKRATE() (*big.Int, error) {
	return _AuctionFeeVault.Contract.MAXPAYBACKRATE(&_AuctionFeeVault.CallOpts)
}

// AccumulatedBids is a free data retrieval call binding the contract method 0x49723142.
//
// Solidity: function accumulatedBids() view returns(uint256)
func (_AuctionFeeVault *AuctionFeeVaultCaller) AccumulatedBids(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AuctionFeeVault.contract.Call(opts, &out, "accumulatedBids")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AccumulatedBids is a free data retrieval call binding the contract method 0x49723142.
//
// Solidity: function accumulatedBids() view returns(uint256)
func (_AuctionFeeVault *AuctionFeeVaultSession) AccumulatedBids() (*big.Int, error) {
	return _AuctionFeeVault.Contract.AccumulatedBids(&_AuctionFeeVault.CallOpts)
}

// AccumulatedBids is a free data retrieval call binding the contract method 0x49723142.
//
// Solidity: function accumulatedBids() view returns(uint256)
func (_AuctionFeeVault *AuctionFeeVaultCallerSession) AccumulatedBids() (*big.Int, error) {
	return _AuctionFeeVault.Contract.AccumulatedBids(&_AuctionFeeVault.CallOpts)
}

// GetRewardAddr is a free data retrieval call binding the contract method 0x27a50f72.
//
// Solidity: function getRewardAddr(address nodeId) view returns(address)
func (_AuctionFeeVault *AuctionFeeVaultCaller) GetRewardAddr(opts *bind.CallOpts, nodeId common.Address) (common.Address, error) {
	var out []interface{}
	err := _AuctionFeeVault.contract.Call(opts, &out, "getRewardAddr", nodeId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetRewardAddr is a free data retrieval call binding the contract method 0x27a50f72.
//
// Solidity: function getRewardAddr(address nodeId) view returns(address)
func (_AuctionFeeVault *AuctionFeeVaultSession) GetRewardAddr(nodeId common.Address) (common.Address, error) {
	return _AuctionFeeVault.Contract.GetRewardAddr(&_AuctionFeeVault.CallOpts, nodeId)
}

// GetRewardAddr is a free data retrieval call binding the contract method 0x27a50f72.
//
// Solidity: function getRewardAddr(address nodeId) view returns(address)
func (_AuctionFeeVault *AuctionFeeVaultCallerSession) GetRewardAddr(nodeId common.Address) (common.Address, error) {
	return _AuctionFeeVault.Contract.GetRewardAddr(&_AuctionFeeVault.CallOpts, nodeId)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_AuctionFeeVault *AuctionFeeVaultCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AuctionFeeVault.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_AuctionFeeVault *AuctionFeeVaultSession) Owner() (common.Address, error) {
	return _AuctionFeeVault.Contract.Owner(&_AuctionFeeVault.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_AuctionFeeVault *AuctionFeeVaultCallerSession) Owner() (common.Address, error) {
	return _AuctionFeeVault.Contract.Owner(&_AuctionFeeVault.CallOpts)
}

// SearcherPaybackRate is a free data retrieval call binding the contract method 0x3c702fbd.
//
// Solidity: function searcherPaybackRate() view returns(uint256)
func (_AuctionFeeVault *AuctionFeeVaultCaller) SearcherPaybackRate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AuctionFeeVault.contract.Call(opts, &out, "searcherPaybackRate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SearcherPaybackRate is a free data retrieval call binding the contract method 0x3c702fbd.
//
// Solidity: function searcherPaybackRate() view returns(uint256)
func (_AuctionFeeVault *AuctionFeeVaultSession) SearcherPaybackRate() (*big.Int, error) {
	return _AuctionFeeVault.Contract.SearcherPaybackRate(&_AuctionFeeVault.CallOpts)
}

// SearcherPaybackRate is a free data retrieval call binding the contract method 0x3c702fbd.
//
// Solidity: function searcherPaybackRate() view returns(uint256)
func (_AuctionFeeVault *AuctionFeeVaultCallerSession) SearcherPaybackRate() (*big.Int, error) {
	return _AuctionFeeVault.Contract.SearcherPaybackRate(&_AuctionFeeVault.CallOpts)
}

// ValidatorPaybackRate is a free data retrieval call binding the contract method 0x89b703aa.
//
// Solidity: function validatorPaybackRate() view returns(uint256)
func (_AuctionFeeVault *AuctionFeeVaultCaller) ValidatorPaybackRate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AuctionFeeVault.contract.Call(opts, &out, "validatorPaybackRate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ValidatorPaybackRate is a free data retrieval call binding the contract method 0x89b703aa.
//
// Solidity: function validatorPaybackRate() view returns(uint256)
func (_AuctionFeeVault *AuctionFeeVaultSession) ValidatorPaybackRate() (*big.Int, error) {
	return _AuctionFeeVault.Contract.ValidatorPaybackRate(&_AuctionFeeVault.CallOpts)
}

// ValidatorPaybackRate is a free data retrieval call binding the contract method 0x89b703aa.
//
// Solidity: function validatorPaybackRate() view returns(uint256)
func (_AuctionFeeVault *AuctionFeeVaultCallerSession) ValidatorPaybackRate() (*big.Int, error) {
	return _AuctionFeeVault.Contract.ValidatorPaybackRate(&_AuctionFeeVault.CallOpts)
}

// RegisterRewardAddress is a paid mutator transaction binding the contract method 0x363d5183.
//
// Solidity: function registerRewardAddress(address nodeId, address rewardAddr) returns()
func (_AuctionFeeVault *AuctionFeeVaultTransactor) RegisterRewardAddress(opts *bind.TransactOpts, nodeId common.Address, rewardAddr common.Address) (*types.Transaction, error) {
	return _AuctionFeeVault.contract.Transact(opts, "registerRewardAddress", nodeId, rewardAddr)
}

// RegisterRewardAddress is a paid mutator transaction binding the contract method 0x363d5183.
//
// Solidity: function registerRewardAddress(address nodeId, address rewardAddr) returns()
func (_AuctionFeeVault *AuctionFeeVaultSession) RegisterRewardAddress(nodeId common.Address, rewardAddr common.Address) (*types.Transaction, error) {
	return _AuctionFeeVault.Contract.RegisterRewardAddress(&_AuctionFeeVault.TransactOpts, nodeId, rewardAddr)
}

// RegisterRewardAddress is a paid mutator transaction binding the contract method 0x363d5183.
//
// Solidity: function registerRewardAddress(address nodeId, address rewardAddr) returns()
func (_AuctionFeeVault *AuctionFeeVaultTransactorSession) RegisterRewardAddress(nodeId common.Address, rewardAddr common.Address) (*types.Transaction, error) {
	return _AuctionFeeVault.Contract.RegisterRewardAddress(&_AuctionFeeVault.TransactOpts, nodeId, rewardAddr)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_AuctionFeeVault *AuctionFeeVaultTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AuctionFeeVault.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_AuctionFeeVault *AuctionFeeVaultSession) RenounceOwnership() (*types.Transaction, error) {
	return _AuctionFeeVault.Contract.RenounceOwnership(&_AuctionFeeVault.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_AuctionFeeVault *AuctionFeeVaultTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _AuctionFeeVault.Contract.RenounceOwnership(&_AuctionFeeVault.TransactOpts)
}

// SetSearcherPaybackRate is a paid mutator transaction binding the contract method 0x36cf2c63.
//
// Solidity: function setSearcherPaybackRate(uint256 _searcherPaybackRate) returns()
func (_AuctionFeeVault *AuctionFeeVaultTransactor) SetSearcherPaybackRate(opts *bind.TransactOpts, _searcherPaybackRate *big.Int) (*types.Transaction, error) {
	return _AuctionFeeVault.contract.Transact(opts, "setSearcherPaybackRate", _searcherPaybackRate)
}

// SetSearcherPaybackRate is a paid mutator transaction binding the contract method 0x36cf2c63.
//
// Solidity: function setSearcherPaybackRate(uint256 _searcherPaybackRate) returns()
func (_AuctionFeeVault *AuctionFeeVaultSession) SetSearcherPaybackRate(_searcherPaybackRate *big.Int) (*types.Transaction, error) {
	return _AuctionFeeVault.Contract.SetSearcherPaybackRate(&_AuctionFeeVault.TransactOpts, _searcherPaybackRate)
}

// SetSearcherPaybackRate is a paid mutator transaction binding the contract method 0x36cf2c63.
//
// Solidity: function setSearcherPaybackRate(uint256 _searcherPaybackRate) returns()
func (_AuctionFeeVault *AuctionFeeVaultTransactorSession) SetSearcherPaybackRate(_searcherPaybackRate *big.Int) (*types.Transaction, error) {
	return _AuctionFeeVault.Contract.SetSearcherPaybackRate(&_AuctionFeeVault.TransactOpts, _searcherPaybackRate)
}

// SetValidatorPaybackRate is a paid mutator transaction binding the contract method 0x11062696.
//
// Solidity: function setValidatorPaybackRate(uint256 _validatorPaybackRate) returns()
func (_AuctionFeeVault *AuctionFeeVaultTransactor) SetValidatorPaybackRate(opts *bind.TransactOpts, _validatorPaybackRate *big.Int) (*types.Transaction, error) {
	return _AuctionFeeVault.contract.Transact(opts, "setValidatorPaybackRate", _validatorPaybackRate)
}

// SetValidatorPaybackRate is a paid mutator transaction binding the contract method 0x11062696.
//
// Solidity: function setValidatorPaybackRate(uint256 _validatorPaybackRate) returns()
func (_AuctionFeeVault *AuctionFeeVaultSession) SetValidatorPaybackRate(_validatorPaybackRate *big.Int) (*types.Transaction, error) {
	return _AuctionFeeVault.Contract.SetValidatorPaybackRate(&_AuctionFeeVault.TransactOpts, _validatorPaybackRate)
}

// SetValidatorPaybackRate is a paid mutator transaction binding the contract method 0x11062696.
//
// Solidity: function setValidatorPaybackRate(uint256 _validatorPaybackRate) returns()
func (_AuctionFeeVault *AuctionFeeVaultTransactorSession) SetValidatorPaybackRate(_validatorPaybackRate *big.Int) (*types.Transaction, error) {
	return _AuctionFeeVault.Contract.SetValidatorPaybackRate(&_AuctionFeeVault.TransactOpts, _validatorPaybackRate)
}

// TakeBid is a paid mutator transaction binding the contract method 0x8573e2ff.
//
// Solidity: function takeBid(address searcher) payable returns()
func (_AuctionFeeVault *AuctionFeeVaultTransactor) TakeBid(opts *bind.TransactOpts, searcher common.Address) (*types.Transaction, error) {
	return _AuctionFeeVault.contract.Transact(opts, "takeBid", searcher)
}

// TakeBid is a paid mutator transaction binding the contract method 0x8573e2ff.
//
// Solidity: function takeBid(address searcher) payable returns()
func (_AuctionFeeVault *AuctionFeeVaultSession) TakeBid(searcher common.Address) (*types.Transaction, error) {
	return _AuctionFeeVault.Contract.TakeBid(&_AuctionFeeVault.TransactOpts, searcher)
}

// TakeBid is a paid mutator transaction binding the contract method 0x8573e2ff.
//
// Solidity: function takeBid(address searcher) payable returns()
func (_AuctionFeeVault *AuctionFeeVaultTransactorSession) TakeBid(searcher common.Address) (*types.Transaction, error) {
	return _AuctionFeeVault.Contract.TakeBid(&_AuctionFeeVault.TransactOpts, searcher)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_AuctionFeeVault *AuctionFeeVaultTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _AuctionFeeVault.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_AuctionFeeVault *AuctionFeeVaultSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _AuctionFeeVault.Contract.TransferOwnership(&_AuctionFeeVault.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_AuctionFeeVault *AuctionFeeVaultTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _AuctionFeeVault.Contract.TransferOwnership(&_AuctionFeeVault.TransactOpts, newOwner)
}

// Withdraw is a paid mutator transaction binding the contract method 0x51cff8d9.
//
// Solidity: function withdraw(address to) returns()
func (_AuctionFeeVault *AuctionFeeVaultTransactor) Withdraw(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _AuctionFeeVault.contract.Transact(opts, "withdraw", to)
}

// Withdraw is a paid mutator transaction binding the contract method 0x51cff8d9.
//
// Solidity: function withdraw(address to) returns()
func (_AuctionFeeVault *AuctionFeeVaultSession) Withdraw(to common.Address) (*types.Transaction, error) {
	return _AuctionFeeVault.Contract.Withdraw(&_AuctionFeeVault.TransactOpts, to)
}

// Withdraw is a paid mutator transaction binding the contract method 0x51cff8d9.
//
// Solidity: function withdraw(address to) returns()
func (_AuctionFeeVault *AuctionFeeVaultTransactorSession) Withdraw(to common.Address) (*types.Transaction, error) {
	return _AuctionFeeVault.Contract.Withdraw(&_AuctionFeeVault.TransactOpts, to)
}

// AuctionFeeVaultFeeDepositIterator is returned from FilterFeeDeposit and is used to iterate over the raw logs and unpacked data for FeeDeposit events raised by the AuctionFeeVault contract.
type AuctionFeeVaultFeeDepositIterator struct {
	Event *AuctionFeeVaultFeeDeposit // Event containing the contract specifics and raw log

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
func (it *AuctionFeeVaultFeeDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AuctionFeeVaultFeeDeposit)
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
		it.Event = new(AuctionFeeVaultFeeDeposit)
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
func (it *AuctionFeeVaultFeeDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AuctionFeeVaultFeeDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AuctionFeeVaultFeeDeposit represents a FeeDeposit event raised by the AuctionFeeVault contract.
type AuctionFeeVaultFeeDeposit struct {
	Sender                 common.Address
	Amount                 *big.Int
	PaybackAmount          *big.Int
	ValidatorPaybackAmount *big.Int
	Raw                    types.Log // Blockchain specific contextual infos
}

// FilterFeeDeposit is a free log retrieval operation binding the contract event 0xa34c9ef6ada915fef21639b2d5c085580cf79046cca66c2c2e8b87e2f3bd8567.
//
// Solidity: event FeeDeposit(address indexed sender, uint256 amount, uint256 paybackAmount, uint256 validatorPaybackAmount)
func (_AuctionFeeVault *AuctionFeeVaultFilterer) FilterFeeDeposit(opts *bind.FilterOpts, sender []common.Address) (*AuctionFeeVaultFeeDepositIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _AuctionFeeVault.contract.FilterLogs(opts, "FeeDeposit", senderRule)
	if err != nil {
		return nil, err
	}
	return &AuctionFeeVaultFeeDepositIterator{contract: _AuctionFeeVault.contract, event: "FeeDeposit", logs: logs, sub: sub}, nil
}

// WatchFeeDeposit is a free log subscription operation binding the contract event 0xa34c9ef6ada915fef21639b2d5c085580cf79046cca66c2c2e8b87e2f3bd8567.
//
// Solidity: event FeeDeposit(address indexed sender, uint256 amount, uint256 paybackAmount, uint256 validatorPaybackAmount)
func (_AuctionFeeVault *AuctionFeeVaultFilterer) WatchFeeDeposit(opts *bind.WatchOpts, sink chan<- *AuctionFeeVaultFeeDeposit, sender []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _AuctionFeeVault.contract.WatchLogs(opts, "FeeDeposit", senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AuctionFeeVaultFeeDeposit)
				if err := _AuctionFeeVault.contract.UnpackLog(event, "FeeDeposit", log); err != nil {
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

// ParseFeeDeposit is a log parse operation binding the contract event 0xa34c9ef6ada915fef21639b2d5c085580cf79046cca66c2c2e8b87e2f3bd8567.
//
// Solidity: event FeeDeposit(address indexed sender, uint256 amount, uint256 paybackAmount, uint256 validatorPaybackAmount)
func (_AuctionFeeVault *AuctionFeeVaultFilterer) ParseFeeDeposit(log types.Log) (*AuctionFeeVaultFeeDeposit, error) {
	event := new(AuctionFeeVaultFeeDeposit)
	if err := _AuctionFeeVault.contract.UnpackLog(event, "FeeDeposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AuctionFeeVaultFeeWithdrawalIterator is returned from FilterFeeWithdrawal and is used to iterate over the raw logs and unpacked data for FeeWithdrawal events raised by the AuctionFeeVault contract.
type AuctionFeeVaultFeeWithdrawalIterator struct {
	Event *AuctionFeeVaultFeeWithdrawal // Event containing the contract specifics and raw log

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
func (it *AuctionFeeVaultFeeWithdrawalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AuctionFeeVaultFeeWithdrawal)
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
		it.Event = new(AuctionFeeVaultFeeWithdrawal)
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
func (it *AuctionFeeVaultFeeWithdrawalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AuctionFeeVaultFeeWithdrawalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AuctionFeeVaultFeeWithdrawal represents a FeeWithdrawal event raised by the AuctionFeeVault contract.
type AuctionFeeVaultFeeWithdrawal struct {
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterFeeWithdrawal is a free log retrieval operation binding the contract event 0x706d7f48c702007c2fb0881cea5759732e64f52faee427d5ab030787cfb7d787.
//
// Solidity: event FeeWithdrawal(uint256 amount)
func (_AuctionFeeVault *AuctionFeeVaultFilterer) FilterFeeWithdrawal(opts *bind.FilterOpts) (*AuctionFeeVaultFeeWithdrawalIterator, error) {

	logs, sub, err := _AuctionFeeVault.contract.FilterLogs(opts, "FeeWithdrawal")
	if err != nil {
		return nil, err
	}
	return &AuctionFeeVaultFeeWithdrawalIterator{contract: _AuctionFeeVault.contract, event: "FeeWithdrawal", logs: logs, sub: sub}, nil
}

// WatchFeeWithdrawal is a free log subscription operation binding the contract event 0x706d7f48c702007c2fb0881cea5759732e64f52faee427d5ab030787cfb7d787.
//
// Solidity: event FeeWithdrawal(uint256 amount)
func (_AuctionFeeVault *AuctionFeeVaultFilterer) WatchFeeWithdrawal(opts *bind.WatchOpts, sink chan<- *AuctionFeeVaultFeeWithdrawal) (event.Subscription, error) {

	logs, sub, err := _AuctionFeeVault.contract.WatchLogs(opts, "FeeWithdrawal")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AuctionFeeVaultFeeWithdrawal)
				if err := _AuctionFeeVault.contract.UnpackLog(event, "FeeWithdrawal", log); err != nil {
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

// ParseFeeWithdrawal is a log parse operation binding the contract event 0x706d7f48c702007c2fb0881cea5759732e64f52faee427d5ab030787cfb7d787.
//
// Solidity: event FeeWithdrawal(uint256 amount)
func (_AuctionFeeVault *AuctionFeeVaultFilterer) ParseFeeWithdrawal(log types.Log) (*AuctionFeeVaultFeeWithdrawal, error) {
	event := new(AuctionFeeVaultFeeWithdrawal)
	if err := _AuctionFeeVault.contract.UnpackLog(event, "FeeWithdrawal", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AuctionFeeVaultOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the AuctionFeeVault contract.
type AuctionFeeVaultOwnershipTransferredIterator struct {
	Event *AuctionFeeVaultOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *AuctionFeeVaultOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AuctionFeeVaultOwnershipTransferred)
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
		it.Event = new(AuctionFeeVaultOwnershipTransferred)
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
func (it *AuctionFeeVaultOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AuctionFeeVaultOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AuctionFeeVaultOwnershipTransferred represents a OwnershipTransferred event raised by the AuctionFeeVault contract.
type AuctionFeeVaultOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_AuctionFeeVault *AuctionFeeVaultFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*AuctionFeeVaultOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _AuctionFeeVault.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &AuctionFeeVaultOwnershipTransferredIterator{contract: _AuctionFeeVault.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_AuctionFeeVault *AuctionFeeVaultFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *AuctionFeeVaultOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _AuctionFeeVault.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AuctionFeeVaultOwnershipTransferred)
				if err := _AuctionFeeVault.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_AuctionFeeVault *AuctionFeeVaultFilterer) ParseOwnershipTransferred(log types.Log) (*AuctionFeeVaultOwnershipTransferred, error) {
	event := new(AuctionFeeVaultOwnershipTransferred)
	if err := _AuctionFeeVault.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AuctionFeeVaultRewardAddressRegisteredIterator is returned from FilterRewardAddressRegistered and is used to iterate over the raw logs and unpacked data for RewardAddressRegistered events raised by the AuctionFeeVault contract.
type AuctionFeeVaultRewardAddressRegisteredIterator struct {
	Event *AuctionFeeVaultRewardAddressRegistered // Event containing the contract specifics and raw log

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
func (it *AuctionFeeVaultRewardAddressRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AuctionFeeVaultRewardAddressRegistered)
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
		it.Event = new(AuctionFeeVaultRewardAddressRegistered)
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
func (it *AuctionFeeVaultRewardAddressRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AuctionFeeVaultRewardAddressRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AuctionFeeVaultRewardAddressRegistered represents a RewardAddressRegistered event raised by the AuctionFeeVault contract.
type AuctionFeeVaultRewardAddressRegistered struct {
	NodeId common.Address
	Reward common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRewardAddressRegistered is a free log retrieval operation binding the contract event 0xe608476cb01b1d04f944f0fdb25841b1f483d26965d42c4a1fab67b8b1488b3b.
//
// Solidity: event RewardAddressRegistered(address indexed nodeId, address indexed reward)
func (_AuctionFeeVault *AuctionFeeVaultFilterer) FilterRewardAddressRegistered(opts *bind.FilterOpts, nodeId []common.Address, reward []common.Address) (*AuctionFeeVaultRewardAddressRegisteredIterator, error) {

	var nodeIdRule []interface{}
	for _, nodeIdItem := range nodeId {
		nodeIdRule = append(nodeIdRule, nodeIdItem)
	}
	var rewardRule []interface{}
	for _, rewardItem := range reward {
		rewardRule = append(rewardRule, rewardItem)
	}

	logs, sub, err := _AuctionFeeVault.contract.FilterLogs(opts, "RewardAddressRegistered", nodeIdRule, rewardRule)
	if err != nil {
		return nil, err
	}
	return &AuctionFeeVaultRewardAddressRegisteredIterator{contract: _AuctionFeeVault.contract, event: "RewardAddressRegistered", logs: logs, sub: sub}, nil
}

// WatchRewardAddressRegistered is a free log subscription operation binding the contract event 0xe608476cb01b1d04f944f0fdb25841b1f483d26965d42c4a1fab67b8b1488b3b.
//
// Solidity: event RewardAddressRegistered(address indexed nodeId, address indexed reward)
func (_AuctionFeeVault *AuctionFeeVaultFilterer) WatchRewardAddressRegistered(opts *bind.WatchOpts, sink chan<- *AuctionFeeVaultRewardAddressRegistered, nodeId []common.Address, reward []common.Address) (event.Subscription, error) {

	var nodeIdRule []interface{}
	for _, nodeIdItem := range nodeId {
		nodeIdRule = append(nodeIdRule, nodeIdItem)
	}
	var rewardRule []interface{}
	for _, rewardItem := range reward {
		rewardRule = append(rewardRule, rewardItem)
	}

	logs, sub, err := _AuctionFeeVault.contract.WatchLogs(opts, "RewardAddressRegistered", nodeIdRule, rewardRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AuctionFeeVaultRewardAddressRegistered)
				if err := _AuctionFeeVault.contract.UnpackLog(event, "RewardAddressRegistered", log); err != nil {
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

// ParseRewardAddressRegistered is a log parse operation binding the contract event 0xe608476cb01b1d04f944f0fdb25841b1f483d26965d42c4a1fab67b8b1488b3b.
//
// Solidity: event RewardAddressRegistered(address indexed nodeId, address indexed reward)
func (_AuctionFeeVault *AuctionFeeVaultFilterer) ParseRewardAddressRegistered(log types.Log) (*AuctionFeeVaultRewardAddressRegistered, error) {
	event := new(AuctionFeeVaultRewardAddressRegistered)
	if err := _AuctionFeeVault.contract.UnpackLog(event, "RewardAddressRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AuctionFeeVaultSearcherPaybackRateUpdatedIterator is returned from FilterSearcherPaybackRateUpdated and is used to iterate over the raw logs and unpacked data for SearcherPaybackRateUpdated events raised by the AuctionFeeVault contract.
type AuctionFeeVaultSearcherPaybackRateUpdatedIterator struct {
	Event *AuctionFeeVaultSearcherPaybackRateUpdated // Event containing the contract specifics and raw log

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
func (it *AuctionFeeVaultSearcherPaybackRateUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AuctionFeeVaultSearcherPaybackRateUpdated)
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
		it.Event = new(AuctionFeeVaultSearcherPaybackRateUpdated)
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
func (it *AuctionFeeVaultSearcherPaybackRateUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AuctionFeeVaultSearcherPaybackRateUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AuctionFeeVaultSearcherPaybackRateUpdated represents a SearcherPaybackRateUpdated event raised by the AuctionFeeVault contract.
type AuctionFeeVaultSearcherPaybackRateUpdated struct {
	SearcherPaybackRate *big.Int
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterSearcherPaybackRateUpdated is a free log retrieval operation binding the contract event 0x71745430318b073bd776904f2432cb283ce3d2ded537bafe2640cf4d6e4bc64f.
//
// Solidity: event SearcherPaybackRateUpdated(uint256 searcherPaybackRate)
func (_AuctionFeeVault *AuctionFeeVaultFilterer) FilterSearcherPaybackRateUpdated(opts *bind.FilterOpts) (*AuctionFeeVaultSearcherPaybackRateUpdatedIterator, error) {

	logs, sub, err := _AuctionFeeVault.contract.FilterLogs(opts, "SearcherPaybackRateUpdated")
	if err != nil {
		return nil, err
	}
	return &AuctionFeeVaultSearcherPaybackRateUpdatedIterator{contract: _AuctionFeeVault.contract, event: "SearcherPaybackRateUpdated", logs: logs, sub: sub}, nil
}

// WatchSearcherPaybackRateUpdated is a free log subscription operation binding the contract event 0x71745430318b073bd776904f2432cb283ce3d2ded537bafe2640cf4d6e4bc64f.
//
// Solidity: event SearcherPaybackRateUpdated(uint256 searcherPaybackRate)
func (_AuctionFeeVault *AuctionFeeVaultFilterer) WatchSearcherPaybackRateUpdated(opts *bind.WatchOpts, sink chan<- *AuctionFeeVaultSearcherPaybackRateUpdated) (event.Subscription, error) {

	logs, sub, err := _AuctionFeeVault.contract.WatchLogs(opts, "SearcherPaybackRateUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AuctionFeeVaultSearcherPaybackRateUpdated)
				if err := _AuctionFeeVault.contract.UnpackLog(event, "SearcherPaybackRateUpdated", log); err != nil {
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

// ParseSearcherPaybackRateUpdated is a log parse operation binding the contract event 0x71745430318b073bd776904f2432cb283ce3d2ded537bafe2640cf4d6e4bc64f.
//
// Solidity: event SearcherPaybackRateUpdated(uint256 searcherPaybackRate)
func (_AuctionFeeVault *AuctionFeeVaultFilterer) ParseSearcherPaybackRateUpdated(log types.Log) (*AuctionFeeVaultSearcherPaybackRateUpdated, error) {
	event := new(AuctionFeeVaultSearcherPaybackRateUpdated)
	if err := _AuctionFeeVault.contract.UnpackLog(event, "SearcherPaybackRateUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AuctionFeeVaultValidatorPaybackRateUpdatedIterator is returned from FilterValidatorPaybackRateUpdated and is used to iterate over the raw logs and unpacked data for ValidatorPaybackRateUpdated events raised by the AuctionFeeVault contract.
type AuctionFeeVaultValidatorPaybackRateUpdatedIterator struct {
	Event *AuctionFeeVaultValidatorPaybackRateUpdated // Event containing the contract specifics and raw log

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
func (it *AuctionFeeVaultValidatorPaybackRateUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AuctionFeeVaultValidatorPaybackRateUpdated)
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
		it.Event = new(AuctionFeeVaultValidatorPaybackRateUpdated)
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
func (it *AuctionFeeVaultValidatorPaybackRateUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AuctionFeeVaultValidatorPaybackRateUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AuctionFeeVaultValidatorPaybackRateUpdated represents a ValidatorPaybackRateUpdated event raised by the AuctionFeeVault contract.
type AuctionFeeVaultValidatorPaybackRateUpdated struct {
	ValidatorPaybackRate *big.Int
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterValidatorPaybackRateUpdated is a free log retrieval operation binding the contract event 0x5309d48fe743a67ce32d8f66af9e2388d65bfc8cc026a4e1fbed3a4612a0af98.
//
// Solidity: event ValidatorPaybackRateUpdated(uint256 validatorPaybackRate)
func (_AuctionFeeVault *AuctionFeeVaultFilterer) FilterValidatorPaybackRateUpdated(opts *bind.FilterOpts) (*AuctionFeeVaultValidatorPaybackRateUpdatedIterator, error) {

	logs, sub, err := _AuctionFeeVault.contract.FilterLogs(opts, "ValidatorPaybackRateUpdated")
	if err != nil {
		return nil, err
	}
	return &AuctionFeeVaultValidatorPaybackRateUpdatedIterator{contract: _AuctionFeeVault.contract, event: "ValidatorPaybackRateUpdated", logs: logs, sub: sub}, nil
}

// WatchValidatorPaybackRateUpdated is a free log subscription operation binding the contract event 0x5309d48fe743a67ce32d8f66af9e2388d65bfc8cc026a4e1fbed3a4612a0af98.
//
// Solidity: event ValidatorPaybackRateUpdated(uint256 validatorPaybackRate)
func (_AuctionFeeVault *AuctionFeeVaultFilterer) WatchValidatorPaybackRateUpdated(opts *bind.WatchOpts, sink chan<- *AuctionFeeVaultValidatorPaybackRateUpdated) (event.Subscription, error) {

	logs, sub, err := _AuctionFeeVault.contract.WatchLogs(opts, "ValidatorPaybackRateUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AuctionFeeVaultValidatorPaybackRateUpdated)
				if err := _AuctionFeeVault.contract.UnpackLog(event, "ValidatorPaybackRateUpdated", log); err != nil {
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

// ParseValidatorPaybackRateUpdated is a log parse operation binding the contract event 0x5309d48fe743a67ce32d8f66af9e2388d65bfc8cc026a4e1fbed3a4612a0af98.
//
// Solidity: event ValidatorPaybackRateUpdated(uint256 validatorPaybackRate)
func (_AuctionFeeVault *AuctionFeeVaultFilterer) ParseValidatorPaybackRateUpdated(log types.Log) (*AuctionFeeVaultValidatorPaybackRateUpdated, error) {
	event := new(AuctionFeeVaultValidatorPaybackRateUpdated)
	if err := _AuctionFeeVault.contract.UnpackLog(event, "ValidatorPaybackRateUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContextMetaData contains all meta data concerning the Context contract.
var ContextMetaData = &bind.MetaData{
	ABI: "[]",
}

// ContextABI is the input ABI used to generate the binding from.
// Deprecated: Use ContextMetaData.ABI instead.
var ContextABI = ContextMetaData.ABI

// ContextBinRuntime is the compiled bytecode used for adding genesis block without deploying code.
const ContextBinRuntime = ``

// Context is an auto generated Go binding around a Kaia contract.
type Context struct {
	ContextCaller     // Read-only binding to the contract
	ContextTransactor // Write-only binding to the contract
	ContextFilterer   // Log filterer for contract events
}

// ContextCaller is an auto generated read-only Go binding around a Kaia contract.
type ContextCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContextTransactor is an auto generated write-only Go binding around a Kaia contract.
type ContextTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContextFilterer is an auto generated log filtering Go binding around a Kaia contract events.
type ContextFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContextSession is an auto generated Go binding around a Kaia contract,
// with pre-set call and transact options.
type ContextSession struct {
	Contract     *Context          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ContextCallerSession is an auto generated read-only Go binding around a Kaia contract,
// with pre-set call options.
type ContextCallerSession struct {
	Contract *ContextCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// ContextTransactorSession is an auto generated write-only Go binding around a Kaia contract,
// with pre-set transact options.
type ContextTransactorSession struct {
	Contract     *ContextTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// ContextRaw is an auto generated low-level Go binding around a Kaia contract.
type ContextRaw struct {
	Contract *Context // Generic contract binding to access the raw methods on
}

// ContextCallerRaw is an auto generated low-level read-only Go binding around a Kaia contract.
type ContextCallerRaw struct {
	Contract *ContextCaller // Generic read-only contract binding to access the raw methods on
}

// ContextTransactorRaw is an auto generated low-level write-only Go binding around a Kaia contract.
type ContextTransactorRaw struct {
	Contract *ContextTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContext creates a new instance of Context, bound to a specific deployed contract.
func NewContext(address common.Address, backend bind.ContractBackend) (*Context, error) {
	contract, err := bindContext(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Context{ContextCaller: ContextCaller{contract: contract}, ContextTransactor: ContextTransactor{contract: contract}, ContextFilterer: ContextFilterer{contract: contract}}, nil
}

// NewContextCaller creates a new read-only instance of Context, bound to a specific deployed contract.
func NewContextCaller(address common.Address, caller bind.ContractCaller) (*ContextCaller, error) {
	contract, err := bindContext(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContextCaller{contract: contract}, nil
}

// NewContextTransactor creates a new write-only instance of Context, bound to a specific deployed contract.
func NewContextTransactor(address common.Address, transactor bind.ContractTransactor) (*ContextTransactor, error) {
	contract, err := bindContext(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContextTransactor{contract: contract}, nil
}

// NewContextFilterer creates a new log filterer instance of Context, bound to a specific deployed contract.
func NewContextFilterer(address common.Address, filterer bind.ContractFilterer) (*ContextFilterer, error) {
	contract, err := bindContext(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContextFilterer{contract: contract}, nil
}

// bindContext binds a generic wrapper to an already deployed contract.
func bindContext(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ContextMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Context *ContextRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Context.Contract.ContextCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Context *ContextRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Context.Contract.ContextTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Context *ContextRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Context.Contract.ContextTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Context *ContextCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Context.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Context *ContextTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Context.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Context *ContextTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Context.Contract.contract.Transact(opts, method, params...)
}

// EnumerableSetMetaData contains all meta data concerning the EnumerableSet contract.
var EnumerableSetMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60556032600b8282823980515f1a607314602657634e487b7160e01b5f525f60045260245ffd5b305f52607381538281f3fe730000000000000000000000000000000000000000301460806040525f80fdfea2646970667358221220ab7ec786af90a9f4a47555f20c6e9ec9c8cbd49047b90c75fa05c2991f6f865664736f6c63430008190033",
}

// EnumerableSetABI is the input ABI used to generate the binding from.
// Deprecated: Use EnumerableSetMetaData.ABI instead.
var EnumerableSetABI = EnumerableSetMetaData.ABI

// EnumerableSetBinRuntime is the compiled bytecode used for adding genesis block without deploying code.
const EnumerableSetBinRuntime = `730000000000000000000000000000000000000000301460806040525f80fdfea2646970667358221220ab7ec786af90a9f4a47555f20c6e9ec9c8cbd49047b90c75fa05c2991f6f865664736f6c63430008190033`

// EnumerableSetBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use EnumerableSetMetaData.Bin instead.
var EnumerableSetBin = EnumerableSetMetaData.Bin

// DeployEnumerableSet deploys a new Kaia contract, binding an instance of EnumerableSet to it.
func DeployEnumerableSet(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *EnumerableSet, error) {
	parsed, err := EnumerableSetMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(EnumerableSetBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &EnumerableSet{EnumerableSetCaller: EnumerableSetCaller{contract: contract}, EnumerableSetTransactor: EnumerableSetTransactor{contract: contract}, EnumerableSetFilterer: EnumerableSetFilterer{contract: contract}}, nil
}

// EnumerableSet is an auto generated Go binding around a Kaia contract.
type EnumerableSet struct {
	EnumerableSetCaller     // Read-only binding to the contract
	EnumerableSetTransactor // Write-only binding to the contract
	EnumerableSetFilterer   // Log filterer for contract events
}

// EnumerableSetCaller is an auto generated read-only Go binding around a Kaia contract.
type EnumerableSetCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EnumerableSetTransactor is an auto generated write-only Go binding around a Kaia contract.
type EnumerableSetTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EnumerableSetFilterer is an auto generated log filtering Go binding around a Kaia contract events.
type EnumerableSetFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EnumerableSetSession is an auto generated Go binding around a Kaia contract,
// with pre-set call and transact options.
type EnumerableSetSession struct {
	Contract     *EnumerableSet    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// EnumerableSetCallerSession is an auto generated read-only Go binding around a Kaia contract,
// with pre-set call options.
type EnumerableSetCallerSession struct {
	Contract *EnumerableSetCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// EnumerableSetTransactorSession is an auto generated write-only Go binding around a Kaia contract,
// with pre-set transact options.
type EnumerableSetTransactorSession struct {
	Contract     *EnumerableSetTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// EnumerableSetRaw is an auto generated low-level Go binding around a Kaia contract.
type EnumerableSetRaw struct {
	Contract *EnumerableSet // Generic contract binding to access the raw methods on
}

// EnumerableSetCallerRaw is an auto generated low-level read-only Go binding around a Kaia contract.
type EnumerableSetCallerRaw struct {
	Contract *EnumerableSetCaller // Generic read-only contract binding to access the raw methods on
}

// EnumerableSetTransactorRaw is an auto generated low-level write-only Go binding around a Kaia contract.
type EnumerableSetTransactorRaw struct {
	Contract *EnumerableSetTransactor // Generic write-only contract binding to access the raw methods on
}

// NewEnumerableSet creates a new instance of EnumerableSet, bound to a specific deployed contract.
func NewEnumerableSet(address common.Address, backend bind.ContractBackend) (*EnumerableSet, error) {
	contract, err := bindEnumerableSet(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &EnumerableSet{EnumerableSetCaller: EnumerableSetCaller{contract: contract}, EnumerableSetTransactor: EnumerableSetTransactor{contract: contract}, EnumerableSetFilterer: EnumerableSetFilterer{contract: contract}}, nil
}

// NewEnumerableSetCaller creates a new read-only instance of EnumerableSet, bound to a specific deployed contract.
func NewEnumerableSetCaller(address common.Address, caller bind.ContractCaller) (*EnumerableSetCaller, error) {
	contract, err := bindEnumerableSet(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &EnumerableSetCaller{contract: contract}, nil
}

// NewEnumerableSetTransactor creates a new write-only instance of EnumerableSet, bound to a specific deployed contract.
func NewEnumerableSetTransactor(address common.Address, transactor bind.ContractTransactor) (*EnumerableSetTransactor, error) {
	contract, err := bindEnumerableSet(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &EnumerableSetTransactor{contract: contract}, nil
}

// NewEnumerableSetFilterer creates a new log filterer instance of EnumerableSet, bound to a specific deployed contract.
func NewEnumerableSetFilterer(address common.Address, filterer bind.ContractFilterer) (*EnumerableSetFilterer, error) {
	contract, err := bindEnumerableSet(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &EnumerableSetFilterer{contract: contract}, nil
}

// bindEnumerableSet binds a generic wrapper to an already deployed contract.
func bindEnumerableSet(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := EnumerableSetMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EnumerableSet *EnumerableSetRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EnumerableSet.Contract.EnumerableSetCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EnumerableSet *EnumerableSetRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EnumerableSet.Contract.EnumerableSetTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EnumerableSet *EnumerableSetRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EnumerableSet.Contract.EnumerableSetTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EnumerableSet *EnumerableSetCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EnumerableSet.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EnumerableSet *EnumerableSetTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EnumerableSet.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EnumerableSet *EnumerableSetTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EnumerableSet.Contract.contract.Transact(opts, method, params...)
}

// IAddressBookMetaData contains all meta data concerning the IAddressBook contract.
var IAddressBookMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_cnNodeId\",\"type\":\"address\"}],\"name\":\"getCnInfo\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"15575d5a": "getCnInfo(address)",
	},
}

// IAddressBookABI is the input ABI used to generate the binding from.
// Deprecated: Use IAddressBookMetaData.ABI instead.
var IAddressBookABI = IAddressBookMetaData.ABI

// IAddressBookBinRuntime is the compiled bytecode used for adding genesis block without deploying code.
const IAddressBookBinRuntime = ``

// Deprecated: Use IAddressBookMetaData.Sigs instead.
// IAddressBookFuncSigs maps the 4-byte function signature to its string representation.
var IAddressBookFuncSigs = IAddressBookMetaData.Sigs

// IAddressBook is an auto generated Go binding around a Kaia contract.
type IAddressBook struct {
	IAddressBookCaller     // Read-only binding to the contract
	IAddressBookTransactor // Write-only binding to the contract
	IAddressBookFilterer   // Log filterer for contract events
}

// IAddressBookCaller is an auto generated read-only Go binding around a Kaia contract.
type IAddressBookCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IAddressBookTransactor is an auto generated write-only Go binding around a Kaia contract.
type IAddressBookTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IAddressBookFilterer is an auto generated log filtering Go binding around a Kaia contract events.
type IAddressBookFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IAddressBookSession is an auto generated Go binding around a Kaia contract,
// with pre-set call and transact options.
type IAddressBookSession struct {
	Contract     *IAddressBook     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IAddressBookCallerSession is an auto generated read-only Go binding around a Kaia contract,
// with pre-set call options.
type IAddressBookCallerSession struct {
	Contract *IAddressBookCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// IAddressBookTransactorSession is an auto generated write-only Go binding around a Kaia contract,
// with pre-set transact options.
type IAddressBookTransactorSession struct {
	Contract     *IAddressBookTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// IAddressBookRaw is an auto generated low-level Go binding around a Kaia contract.
type IAddressBookRaw struct {
	Contract *IAddressBook // Generic contract binding to access the raw methods on
}

// IAddressBookCallerRaw is an auto generated low-level read-only Go binding around a Kaia contract.
type IAddressBookCallerRaw struct {
	Contract *IAddressBookCaller // Generic read-only contract binding to access the raw methods on
}

// IAddressBookTransactorRaw is an auto generated low-level write-only Go binding around a Kaia contract.
type IAddressBookTransactorRaw struct {
	Contract *IAddressBookTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIAddressBook creates a new instance of IAddressBook, bound to a specific deployed contract.
func NewIAddressBook(address common.Address, backend bind.ContractBackend) (*IAddressBook, error) {
	contract, err := bindIAddressBook(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IAddressBook{IAddressBookCaller: IAddressBookCaller{contract: contract}, IAddressBookTransactor: IAddressBookTransactor{contract: contract}, IAddressBookFilterer: IAddressBookFilterer{contract: contract}}, nil
}

// NewIAddressBookCaller creates a new read-only instance of IAddressBook, bound to a specific deployed contract.
func NewIAddressBookCaller(address common.Address, caller bind.ContractCaller) (*IAddressBookCaller, error) {
	contract, err := bindIAddressBook(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IAddressBookCaller{contract: contract}, nil
}

// NewIAddressBookTransactor creates a new write-only instance of IAddressBook, bound to a specific deployed contract.
func NewIAddressBookTransactor(address common.Address, transactor bind.ContractTransactor) (*IAddressBookTransactor, error) {
	contract, err := bindIAddressBook(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IAddressBookTransactor{contract: contract}, nil
}

// NewIAddressBookFilterer creates a new log filterer instance of IAddressBook, bound to a specific deployed contract.
func NewIAddressBookFilterer(address common.Address, filterer bind.ContractFilterer) (*IAddressBookFilterer, error) {
	contract, err := bindIAddressBook(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IAddressBookFilterer{contract: contract}, nil
}

// bindIAddressBook binds a generic wrapper to an already deployed contract.
func bindIAddressBook(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IAddressBookMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IAddressBook *IAddressBookRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IAddressBook.Contract.IAddressBookCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IAddressBook *IAddressBookRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IAddressBook.Contract.IAddressBookTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IAddressBook *IAddressBookRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IAddressBook.Contract.IAddressBookTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IAddressBook *IAddressBookCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IAddressBook.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IAddressBook *IAddressBookTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IAddressBook.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IAddressBook *IAddressBookTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IAddressBook.Contract.contract.Transact(opts, method, params...)
}

// GetCnInfo is a free data retrieval call binding the contract method 0x15575d5a.
//
// Solidity: function getCnInfo(address _cnNodeId) view returns(address, address, address)
func (_IAddressBook *IAddressBookCaller) GetCnInfo(opts *bind.CallOpts, _cnNodeId common.Address) (common.Address, common.Address, common.Address, error) {
	var out []interface{}
	err := _IAddressBook.contract.Call(opts, &out, "getCnInfo", _cnNodeId)

	if err != nil {
		return *new(common.Address), *new(common.Address), *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	out1 := *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	out2 := *abi.ConvertType(out[2], new(common.Address)).(*common.Address)

	return out0, out1, out2, err

}

// GetCnInfo is a free data retrieval call binding the contract method 0x15575d5a.
//
// Solidity: function getCnInfo(address _cnNodeId) view returns(address, address, address)
func (_IAddressBook *IAddressBookSession) GetCnInfo(_cnNodeId common.Address) (common.Address, common.Address, common.Address, error) {
	return _IAddressBook.Contract.GetCnInfo(&_IAddressBook.CallOpts, _cnNodeId)
}

// GetCnInfo is a free data retrieval call binding the contract method 0x15575d5a.
//
// Solidity: function getCnInfo(address _cnNodeId) view returns(address, address, address)
func (_IAddressBook *IAddressBookCallerSession) GetCnInfo(_cnNodeId common.Address) (common.Address, common.Address, common.Address, error) {
	return _IAddressBook.Contract.GetCnInfo(&_IAddressBook.CallOpts, _cnNodeId)
}

// IAuctionFeeVaultMetaData contains all meta data concerning the IAuctionFeeVault contract.
var IAuctionFeeVaultMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"paybackAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"validatorPaybackAmount\",\"type\":\"uint256\"}],\"name\":\"FeeDeposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"FeeWithdrawal\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"nodeId\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"reward\",\"type\":\"address\"}],\"name\":\"RewardAddressRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"searcherPaybackRate\",\"type\":\"uint256\"}],\"name\":\"SearcherPaybackRateUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"validatorPaybackRate\",\"type\":\"uint256\"}],\"name\":\"ValidatorPaybackRateUpdated\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"nodeId\",\"type\":\"address\"}],\"name\":\"getRewardAddr\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"nodeId\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"rewardAddr\",\"type\":\"address\"}],\"name\":\"registerRewardAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_searcherPaybackRate\",\"type\":\"uint256\"}],\"name\":\"setSearcherPaybackRate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_validatorPaybackRate\",\"type\":\"uint256\"}],\"name\":\"setValidatorPaybackRate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"searcher\",\"type\":\"address\"}],\"name\":\"takeBid\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"27a50f72": "getRewardAddr(address)",
		"363d5183": "registerRewardAddress(address,address)",
		"36cf2c63": "setSearcherPaybackRate(uint256)",
		"11062696": "setValidatorPaybackRate(uint256)",
		"8573e2ff": "takeBid(address)",
		"51cff8d9": "withdraw(address)",
	},
}

// IAuctionFeeVaultABI is the input ABI used to generate the binding from.
// Deprecated: Use IAuctionFeeVaultMetaData.ABI instead.
var IAuctionFeeVaultABI = IAuctionFeeVaultMetaData.ABI

// IAuctionFeeVaultBinRuntime is the compiled bytecode used for adding genesis block without deploying code.
const IAuctionFeeVaultBinRuntime = ``

// Deprecated: Use IAuctionFeeVaultMetaData.Sigs instead.
// IAuctionFeeVaultFuncSigs maps the 4-byte function signature to its string representation.
var IAuctionFeeVaultFuncSigs = IAuctionFeeVaultMetaData.Sigs

// IAuctionFeeVault is an auto generated Go binding around a Kaia contract.
type IAuctionFeeVault struct {
	IAuctionFeeVaultCaller     // Read-only binding to the contract
	IAuctionFeeVaultTransactor // Write-only binding to the contract
	IAuctionFeeVaultFilterer   // Log filterer for contract events
}

// IAuctionFeeVaultCaller is an auto generated read-only Go binding around a Kaia contract.
type IAuctionFeeVaultCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IAuctionFeeVaultTransactor is an auto generated write-only Go binding around a Kaia contract.
type IAuctionFeeVaultTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IAuctionFeeVaultFilterer is an auto generated log filtering Go binding around a Kaia contract events.
type IAuctionFeeVaultFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IAuctionFeeVaultSession is an auto generated Go binding around a Kaia contract,
// with pre-set call and transact options.
type IAuctionFeeVaultSession struct {
	Contract     *IAuctionFeeVault // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IAuctionFeeVaultCallerSession is an auto generated read-only Go binding around a Kaia contract,
// with pre-set call options.
type IAuctionFeeVaultCallerSession struct {
	Contract *IAuctionFeeVaultCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// IAuctionFeeVaultTransactorSession is an auto generated write-only Go binding around a Kaia contract,
// with pre-set transact options.
type IAuctionFeeVaultTransactorSession struct {
	Contract     *IAuctionFeeVaultTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// IAuctionFeeVaultRaw is an auto generated low-level Go binding around a Kaia contract.
type IAuctionFeeVaultRaw struct {
	Contract *IAuctionFeeVault // Generic contract binding to access the raw methods on
}

// IAuctionFeeVaultCallerRaw is an auto generated low-level read-only Go binding around a Kaia contract.
type IAuctionFeeVaultCallerRaw struct {
	Contract *IAuctionFeeVaultCaller // Generic read-only contract binding to access the raw methods on
}

// IAuctionFeeVaultTransactorRaw is an auto generated low-level write-only Go binding around a Kaia contract.
type IAuctionFeeVaultTransactorRaw struct {
	Contract *IAuctionFeeVaultTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIAuctionFeeVault creates a new instance of IAuctionFeeVault, bound to a specific deployed contract.
func NewIAuctionFeeVault(address common.Address, backend bind.ContractBackend) (*IAuctionFeeVault, error) {
	contract, err := bindIAuctionFeeVault(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IAuctionFeeVault{IAuctionFeeVaultCaller: IAuctionFeeVaultCaller{contract: contract}, IAuctionFeeVaultTransactor: IAuctionFeeVaultTransactor{contract: contract}, IAuctionFeeVaultFilterer: IAuctionFeeVaultFilterer{contract: contract}}, nil
}

// NewIAuctionFeeVaultCaller creates a new read-only instance of IAuctionFeeVault, bound to a specific deployed contract.
func NewIAuctionFeeVaultCaller(address common.Address, caller bind.ContractCaller) (*IAuctionFeeVaultCaller, error) {
	contract, err := bindIAuctionFeeVault(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IAuctionFeeVaultCaller{contract: contract}, nil
}

// NewIAuctionFeeVaultTransactor creates a new write-only instance of IAuctionFeeVault, bound to a specific deployed contract.
func NewIAuctionFeeVaultTransactor(address common.Address, transactor bind.ContractTransactor) (*IAuctionFeeVaultTransactor, error) {
	contract, err := bindIAuctionFeeVault(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IAuctionFeeVaultTransactor{contract: contract}, nil
}

// NewIAuctionFeeVaultFilterer creates a new log filterer instance of IAuctionFeeVault, bound to a specific deployed contract.
func NewIAuctionFeeVaultFilterer(address common.Address, filterer bind.ContractFilterer) (*IAuctionFeeVaultFilterer, error) {
	contract, err := bindIAuctionFeeVault(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IAuctionFeeVaultFilterer{contract: contract}, nil
}

// bindIAuctionFeeVault binds a generic wrapper to an already deployed contract.
func bindIAuctionFeeVault(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IAuctionFeeVaultMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IAuctionFeeVault *IAuctionFeeVaultRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IAuctionFeeVault.Contract.IAuctionFeeVaultCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IAuctionFeeVault *IAuctionFeeVaultRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IAuctionFeeVault.Contract.IAuctionFeeVaultTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IAuctionFeeVault *IAuctionFeeVaultRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IAuctionFeeVault.Contract.IAuctionFeeVaultTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IAuctionFeeVault *IAuctionFeeVaultCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IAuctionFeeVault.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IAuctionFeeVault *IAuctionFeeVaultTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IAuctionFeeVault.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IAuctionFeeVault *IAuctionFeeVaultTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IAuctionFeeVault.Contract.contract.Transact(opts, method, params...)
}

// GetRewardAddr is a free data retrieval call binding the contract method 0x27a50f72.
//
// Solidity: function getRewardAddr(address nodeId) view returns(address)
func (_IAuctionFeeVault *IAuctionFeeVaultCaller) GetRewardAddr(opts *bind.CallOpts, nodeId common.Address) (common.Address, error) {
	var out []interface{}
	err := _IAuctionFeeVault.contract.Call(opts, &out, "getRewardAddr", nodeId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetRewardAddr is a free data retrieval call binding the contract method 0x27a50f72.
//
// Solidity: function getRewardAddr(address nodeId) view returns(address)
func (_IAuctionFeeVault *IAuctionFeeVaultSession) GetRewardAddr(nodeId common.Address) (common.Address, error) {
	return _IAuctionFeeVault.Contract.GetRewardAddr(&_IAuctionFeeVault.CallOpts, nodeId)
}

// GetRewardAddr is a free data retrieval call binding the contract method 0x27a50f72.
//
// Solidity: function getRewardAddr(address nodeId) view returns(address)
func (_IAuctionFeeVault *IAuctionFeeVaultCallerSession) GetRewardAddr(nodeId common.Address) (common.Address, error) {
	return _IAuctionFeeVault.Contract.GetRewardAddr(&_IAuctionFeeVault.CallOpts, nodeId)
}

// RegisterRewardAddress is a paid mutator transaction binding the contract method 0x363d5183.
//
// Solidity: function registerRewardAddress(address nodeId, address rewardAddr) returns()
func (_IAuctionFeeVault *IAuctionFeeVaultTransactor) RegisterRewardAddress(opts *bind.TransactOpts, nodeId common.Address, rewardAddr common.Address) (*types.Transaction, error) {
	return _IAuctionFeeVault.contract.Transact(opts, "registerRewardAddress", nodeId, rewardAddr)
}

// RegisterRewardAddress is a paid mutator transaction binding the contract method 0x363d5183.
//
// Solidity: function registerRewardAddress(address nodeId, address rewardAddr) returns()
func (_IAuctionFeeVault *IAuctionFeeVaultSession) RegisterRewardAddress(nodeId common.Address, rewardAddr common.Address) (*types.Transaction, error) {
	return _IAuctionFeeVault.Contract.RegisterRewardAddress(&_IAuctionFeeVault.TransactOpts, nodeId, rewardAddr)
}

// RegisterRewardAddress is a paid mutator transaction binding the contract method 0x363d5183.
//
// Solidity: function registerRewardAddress(address nodeId, address rewardAddr) returns()
func (_IAuctionFeeVault *IAuctionFeeVaultTransactorSession) RegisterRewardAddress(nodeId common.Address, rewardAddr common.Address) (*types.Transaction, error) {
	return _IAuctionFeeVault.Contract.RegisterRewardAddress(&_IAuctionFeeVault.TransactOpts, nodeId, rewardAddr)
}

// SetSearcherPaybackRate is a paid mutator transaction binding the contract method 0x36cf2c63.
//
// Solidity: function setSearcherPaybackRate(uint256 _searcherPaybackRate) returns()
func (_IAuctionFeeVault *IAuctionFeeVaultTransactor) SetSearcherPaybackRate(opts *bind.TransactOpts, _searcherPaybackRate *big.Int) (*types.Transaction, error) {
	return _IAuctionFeeVault.contract.Transact(opts, "setSearcherPaybackRate", _searcherPaybackRate)
}

// SetSearcherPaybackRate is a paid mutator transaction binding the contract method 0x36cf2c63.
//
// Solidity: function setSearcherPaybackRate(uint256 _searcherPaybackRate) returns()
func (_IAuctionFeeVault *IAuctionFeeVaultSession) SetSearcherPaybackRate(_searcherPaybackRate *big.Int) (*types.Transaction, error) {
	return _IAuctionFeeVault.Contract.SetSearcherPaybackRate(&_IAuctionFeeVault.TransactOpts, _searcherPaybackRate)
}

// SetSearcherPaybackRate is a paid mutator transaction binding the contract method 0x36cf2c63.
//
// Solidity: function setSearcherPaybackRate(uint256 _searcherPaybackRate) returns()
func (_IAuctionFeeVault *IAuctionFeeVaultTransactorSession) SetSearcherPaybackRate(_searcherPaybackRate *big.Int) (*types.Transaction, error) {
	return _IAuctionFeeVault.Contract.SetSearcherPaybackRate(&_IAuctionFeeVault.TransactOpts, _searcherPaybackRate)
}

// SetValidatorPaybackRate is a paid mutator transaction binding the contract method 0x11062696.
//
// Solidity: function setValidatorPaybackRate(uint256 _validatorPaybackRate) returns()
func (_IAuctionFeeVault *IAuctionFeeVaultTransactor) SetValidatorPaybackRate(opts *bind.TransactOpts, _validatorPaybackRate *big.Int) (*types.Transaction, error) {
	return _IAuctionFeeVault.contract.Transact(opts, "setValidatorPaybackRate", _validatorPaybackRate)
}

// SetValidatorPaybackRate is a paid mutator transaction binding the contract method 0x11062696.
//
// Solidity: function setValidatorPaybackRate(uint256 _validatorPaybackRate) returns()
func (_IAuctionFeeVault *IAuctionFeeVaultSession) SetValidatorPaybackRate(_validatorPaybackRate *big.Int) (*types.Transaction, error) {
	return _IAuctionFeeVault.Contract.SetValidatorPaybackRate(&_IAuctionFeeVault.TransactOpts, _validatorPaybackRate)
}

// SetValidatorPaybackRate is a paid mutator transaction binding the contract method 0x11062696.
//
// Solidity: function setValidatorPaybackRate(uint256 _validatorPaybackRate) returns()
func (_IAuctionFeeVault *IAuctionFeeVaultTransactorSession) SetValidatorPaybackRate(_validatorPaybackRate *big.Int) (*types.Transaction, error) {
	return _IAuctionFeeVault.Contract.SetValidatorPaybackRate(&_IAuctionFeeVault.TransactOpts, _validatorPaybackRate)
}

// TakeBid is a paid mutator transaction binding the contract method 0x8573e2ff.
//
// Solidity: function takeBid(address searcher) payable returns()
func (_IAuctionFeeVault *IAuctionFeeVaultTransactor) TakeBid(opts *bind.TransactOpts, searcher common.Address) (*types.Transaction, error) {
	return _IAuctionFeeVault.contract.Transact(opts, "takeBid", searcher)
}

// TakeBid is a paid mutator transaction binding the contract method 0x8573e2ff.
//
// Solidity: function takeBid(address searcher) payable returns()
func (_IAuctionFeeVault *IAuctionFeeVaultSession) TakeBid(searcher common.Address) (*types.Transaction, error) {
	return _IAuctionFeeVault.Contract.TakeBid(&_IAuctionFeeVault.TransactOpts, searcher)
}

// TakeBid is a paid mutator transaction binding the contract method 0x8573e2ff.
//
// Solidity: function takeBid(address searcher) payable returns()
func (_IAuctionFeeVault *IAuctionFeeVaultTransactorSession) TakeBid(searcher common.Address) (*types.Transaction, error) {
	return _IAuctionFeeVault.Contract.TakeBid(&_IAuctionFeeVault.TransactOpts, searcher)
}

// Withdraw is a paid mutator transaction binding the contract method 0x51cff8d9.
//
// Solidity: function withdraw(address to) returns()
func (_IAuctionFeeVault *IAuctionFeeVaultTransactor) Withdraw(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _IAuctionFeeVault.contract.Transact(opts, "withdraw", to)
}

// Withdraw is a paid mutator transaction binding the contract method 0x51cff8d9.
//
// Solidity: function withdraw(address to) returns()
func (_IAuctionFeeVault *IAuctionFeeVaultSession) Withdraw(to common.Address) (*types.Transaction, error) {
	return _IAuctionFeeVault.Contract.Withdraw(&_IAuctionFeeVault.TransactOpts, to)
}

// Withdraw is a paid mutator transaction binding the contract method 0x51cff8d9.
//
// Solidity: function withdraw(address to) returns()
func (_IAuctionFeeVault *IAuctionFeeVaultTransactorSession) Withdraw(to common.Address) (*types.Transaction, error) {
	return _IAuctionFeeVault.Contract.Withdraw(&_IAuctionFeeVault.TransactOpts, to)
}

// IAuctionFeeVaultFeeDepositIterator is returned from FilterFeeDeposit and is used to iterate over the raw logs and unpacked data for FeeDeposit events raised by the IAuctionFeeVault contract.
type IAuctionFeeVaultFeeDepositIterator struct {
	Event *IAuctionFeeVaultFeeDeposit // Event containing the contract specifics and raw log

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
func (it *IAuctionFeeVaultFeeDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IAuctionFeeVaultFeeDeposit)
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
		it.Event = new(IAuctionFeeVaultFeeDeposit)
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
func (it *IAuctionFeeVaultFeeDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IAuctionFeeVaultFeeDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IAuctionFeeVaultFeeDeposit represents a FeeDeposit event raised by the IAuctionFeeVault contract.
type IAuctionFeeVaultFeeDeposit struct {
	Sender                 common.Address
	Amount                 *big.Int
	PaybackAmount          *big.Int
	ValidatorPaybackAmount *big.Int
	Raw                    types.Log // Blockchain specific contextual infos
}

// FilterFeeDeposit is a free log retrieval operation binding the contract event 0xa34c9ef6ada915fef21639b2d5c085580cf79046cca66c2c2e8b87e2f3bd8567.
//
// Solidity: event FeeDeposit(address indexed sender, uint256 amount, uint256 paybackAmount, uint256 validatorPaybackAmount)
func (_IAuctionFeeVault *IAuctionFeeVaultFilterer) FilterFeeDeposit(opts *bind.FilterOpts, sender []common.Address) (*IAuctionFeeVaultFeeDepositIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _IAuctionFeeVault.contract.FilterLogs(opts, "FeeDeposit", senderRule)
	if err != nil {
		return nil, err
	}
	return &IAuctionFeeVaultFeeDepositIterator{contract: _IAuctionFeeVault.contract, event: "FeeDeposit", logs: logs, sub: sub}, nil
}

// WatchFeeDeposit is a free log subscription operation binding the contract event 0xa34c9ef6ada915fef21639b2d5c085580cf79046cca66c2c2e8b87e2f3bd8567.
//
// Solidity: event FeeDeposit(address indexed sender, uint256 amount, uint256 paybackAmount, uint256 validatorPaybackAmount)
func (_IAuctionFeeVault *IAuctionFeeVaultFilterer) WatchFeeDeposit(opts *bind.WatchOpts, sink chan<- *IAuctionFeeVaultFeeDeposit, sender []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _IAuctionFeeVault.contract.WatchLogs(opts, "FeeDeposit", senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IAuctionFeeVaultFeeDeposit)
				if err := _IAuctionFeeVault.contract.UnpackLog(event, "FeeDeposit", log); err != nil {
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

// ParseFeeDeposit is a log parse operation binding the contract event 0xa34c9ef6ada915fef21639b2d5c085580cf79046cca66c2c2e8b87e2f3bd8567.
//
// Solidity: event FeeDeposit(address indexed sender, uint256 amount, uint256 paybackAmount, uint256 validatorPaybackAmount)
func (_IAuctionFeeVault *IAuctionFeeVaultFilterer) ParseFeeDeposit(log types.Log) (*IAuctionFeeVaultFeeDeposit, error) {
	event := new(IAuctionFeeVaultFeeDeposit)
	if err := _IAuctionFeeVault.contract.UnpackLog(event, "FeeDeposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IAuctionFeeVaultFeeWithdrawalIterator is returned from FilterFeeWithdrawal and is used to iterate over the raw logs and unpacked data for FeeWithdrawal events raised by the IAuctionFeeVault contract.
type IAuctionFeeVaultFeeWithdrawalIterator struct {
	Event *IAuctionFeeVaultFeeWithdrawal // Event containing the contract specifics and raw log

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
func (it *IAuctionFeeVaultFeeWithdrawalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IAuctionFeeVaultFeeWithdrawal)
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
		it.Event = new(IAuctionFeeVaultFeeWithdrawal)
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
func (it *IAuctionFeeVaultFeeWithdrawalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IAuctionFeeVaultFeeWithdrawalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IAuctionFeeVaultFeeWithdrawal represents a FeeWithdrawal event raised by the IAuctionFeeVault contract.
type IAuctionFeeVaultFeeWithdrawal struct {
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterFeeWithdrawal is a free log retrieval operation binding the contract event 0x706d7f48c702007c2fb0881cea5759732e64f52faee427d5ab030787cfb7d787.
//
// Solidity: event FeeWithdrawal(uint256 amount)
func (_IAuctionFeeVault *IAuctionFeeVaultFilterer) FilterFeeWithdrawal(opts *bind.FilterOpts) (*IAuctionFeeVaultFeeWithdrawalIterator, error) {

	logs, sub, err := _IAuctionFeeVault.contract.FilterLogs(opts, "FeeWithdrawal")
	if err != nil {
		return nil, err
	}
	return &IAuctionFeeVaultFeeWithdrawalIterator{contract: _IAuctionFeeVault.contract, event: "FeeWithdrawal", logs: logs, sub: sub}, nil
}

// WatchFeeWithdrawal is a free log subscription operation binding the contract event 0x706d7f48c702007c2fb0881cea5759732e64f52faee427d5ab030787cfb7d787.
//
// Solidity: event FeeWithdrawal(uint256 amount)
func (_IAuctionFeeVault *IAuctionFeeVaultFilterer) WatchFeeWithdrawal(opts *bind.WatchOpts, sink chan<- *IAuctionFeeVaultFeeWithdrawal) (event.Subscription, error) {

	logs, sub, err := _IAuctionFeeVault.contract.WatchLogs(opts, "FeeWithdrawal")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IAuctionFeeVaultFeeWithdrawal)
				if err := _IAuctionFeeVault.contract.UnpackLog(event, "FeeWithdrawal", log); err != nil {
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

// ParseFeeWithdrawal is a log parse operation binding the contract event 0x706d7f48c702007c2fb0881cea5759732e64f52faee427d5ab030787cfb7d787.
//
// Solidity: event FeeWithdrawal(uint256 amount)
func (_IAuctionFeeVault *IAuctionFeeVaultFilterer) ParseFeeWithdrawal(log types.Log) (*IAuctionFeeVaultFeeWithdrawal, error) {
	event := new(IAuctionFeeVaultFeeWithdrawal)
	if err := _IAuctionFeeVault.contract.UnpackLog(event, "FeeWithdrawal", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IAuctionFeeVaultRewardAddressRegisteredIterator is returned from FilterRewardAddressRegistered and is used to iterate over the raw logs and unpacked data for RewardAddressRegistered events raised by the IAuctionFeeVault contract.
type IAuctionFeeVaultRewardAddressRegisteredIterator struct {
	Event *IAuctionFeeVaultRewardAddressRegistered // Event containing the contract specifics and raw log

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
func (it *IAuctionFeeVaultRewardAddressRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IAuctionFeeVaultRewardAddressRegistered)
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
		it.Event = new(IAuctionFeeVaultRewardAddressRegistered)
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
func (it *IAuctionFeeVaultRewardAddressRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IAuctionFeeVaultRewardAddressRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IAuctionFeeVaultRewardAddressRegistered represents a RewardAddressRegistered event raised by the IAuctionFeeVault contract.
type IAuctionFeeVaultRewardAddressRegistered struct {
	NodeId common.Address
	Reward common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRewardAddressRegistered is a free log retrieval operation binding the contract event 0xe608476cb01b1d04f944f0fdb25841b1f483d26965d42c4a1fab67b8b1488b3b.
//
// Solidity: event RewardAddressRegistered(address indexed nodeId, address indexed reward)
func (_IAuctionFeeVault *IAuctionFeeVaultFilterer) FilterRewardAddressRegistered(opts *bind.FilterOpts, nodeId []common.Address, reward []common.Address) (*IAuctionFeeVaultRewardAddressRegisteredIterator, error) {

	var nodeIdRule []interface{}
	for _, nodeIdItem := range nodeId {
		nodeIdRule = append(nodeIdRule, nodeIdItem)
	}
	var rewardRule []interface{}
	for _, rewardItem := range reward {
		rewardRule = append(rewardRule, rewardItem)
	}

	logs, sub, err := _IAuctionFeeVault.contract.FilterLogs(opts, "RewardAddressRegistered", nodeIdRule, rewardRule)
	if err != nil {
		return nil, err
	}
	return &IAuctionFeeVaultRewardAddressRegisteredIterator{contract: _IAuctionFeeVault.contract, event: "RewardAddressRegistered", logs: logs, sub: sub}, nil
}

// WatchRewardAddressRegistered is a free log subscription operation binding the contract event 0xe608476cb01b1d04f944f0fdb25841b1f483d26965d42c4a1fab67b8b1488b3b.
//
// Solidity: event RewardAddressRegistered(address indexed nodeId, address indexed reward)
func (_IAuctionFeeVault *IAuctionFeeVaultFilterer) WatchRewardAddressRegistered(opts *bind.WatchOpts, sink chan<- *IAuctionFeeVaultRewardAddressRegistered, nodeId []common.Address, reward []common.Address) (event.Subscription, error) {

	var nodeIdRule []interface{}
	for _, nodeIdItem := range nodeId {
		nodeIdRule = append(nodeIdRule, nodeIdItem)
	}
	var rewardRule []interface{}
	for _, rewardItem := range reward {
		rewardRule = append(rewardRule, rewardItem)
	}

	logs, sub, err := _IAuctionFeeVault.contract.WatchLogs(opts, "RewardAddressRegistered", nodeIdRule, rewardRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IAuctionFeeVaultRewardAddressRegistered)
				if err := _IAuctionFeeVault.contract.UnpackLog(event, "RewardAddressRegistered", log); err != nil {
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

// ParseRewardAddressRegistered is a log parse operation binding the contract event 0xe608476cb01b1d04f944f0fdb25841b1f483d26965d42c4a1fab67b8b1488b3b.
//
// Solidity: event RewardAddressRegistered(address indexed nodeId, address indexed reward)
func (_IAuctionFeeVault *IAuctionFeeVaultFilterer) ParseRewardAddressRegistered(log types.Log) (*IAuctionFeeVaultRewardAddressRegistered, error) {
	event := new(IAuctionFeeVaultRewardAddressRegistered)
	if err := _IAuctionFeeVault.contract.UnpackLog(event, "RewardAddressRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IAuctionFeeVaultSearcherPaybackRateUpdatedIterator is returned from FilterSearcherPaybackRateUpdated and is used to iterate over the raw logs and unpacked data for SearcherPaybackRateUpdated events raised by the IAuctionFeeVault contract.
type IAuctionFeeVaultSearcherPaybackRateUpdatedIterator struct {
	Event *IAuctionFeeVaultSearcherPaybackRateUpdated // Event containing the contract specifics and raw log

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
func (it *IAuctionFeeVaultSearcherPaybackRateUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IAuctionFeeVaultSearcherPaybackRateUpdated)
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
		it.Event = new(IAuctionFeeVaultSearcherPaybackRateUpdated)
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
func (it *IAuctionFeeVaultSearcherPaybackRateUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IAuctionFeeVaultSearcherPaybackRateUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IAuctionFeeVaultSearcherPaybackRateUpdated represents a SearcherPaybackRateUpdated event raised by the IAuctionFeeVault contract.
type IAuctionFeeVaultSearcherPaybackRateUpdated struct {
	SearcherPaybackRate *big.Int
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterSearcherPaybackRateUpdated is a free log retrieval operation binding the contract event 0x71745430318b073bd776904f2432cb283ce3d2ded537bafe2640cf4d6e4bc64f.
//
// Solidity: event SearcherPaybackRateUpdated(uint256 searcherPaybackRate)
func (_IAuctionFeeVault *IAuctionFeeVaultFilterer) FilterSearcherPaybackRateUpdated(opts *bind.FilterOpts) (*IAuctionFeeVaultSearcherPaybackRateUpdatedIterator, error) {

	logs, sub, err := _IAuctionFeeVault.contract.FilterLogs(opts, "SearcherPaybackRateUpdated")
	if err != nil {
		return nil, err
	}
	return &IAuctionFeeVaultSearcherPaybackRateUpdatedIterator{contract: _IAuctionFeeVault.contract, event: "SearcherPaybackRateUpdated", logs: logs, sub: sub}, nil
}

// WatchSearcherPaybackRateUpdated is a free log subscription operation binding the contract event 0x71745430318b073bd776904f2432cb283ce3d2ded537bafe2640cf4d6e4bc64f.
//
// Solidity: event SearcherPaybackRateUpdated(uint256 searcherPaybackRate)
func (_IAuctionFeeVault *IAuctionFeeVaultFilterer) WatchSearcherPaybackRateUpdated(opts *bind.WatchOpts, sink chan<- *IAuctionFeeVaultSearcherPaybackRateUpdated) (event.Subscription, error) {

	logs, sub, err := _IAuctionFeeVault.contract.WatchLogs(opts, "SearcherPaybackRateUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IAuctionFeeVaultSearcherPaybackRateUpdated)
				if err := _IAuctionFeeVault.contract.UnpackLog(event, "SearcherPaybackRateUpdated", log); err != nil {
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

// ParseSearcherPaybackRateUpdated is a log parse operation binding the contract event 0x71745430318b073bd776904f2432cb283ce3d2ded537bafe2640cf4d6e4bc64f.
//
// Solidity: event SearcherPaybackRateUpdated(uint256 searcherPaybackRate)
func (_IAuctionFeeVault *IAuctionFeeVaultFilterer) ParseSearcherPaybackRateUpdated(log types.Log) (*IAuctionFeeVaultSearcherPaybackRateUpdated, error) {
	event := new(IAuctionFeeVaultSearcherPaybackRateUpdated)
	if err := _IAuctionFeeVault.contract.UnpackLog(event, "SearcherPaybackRateUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IAuctionFeeVaultValidatorPaybackRateUpdatedIterator is returned from FilterValidatorPaybackRateUpdated and is used to iterate over the raw logs and unpacked data for ValidatorPaybackRateUpdated events raised by the IAuctionFeeVault contract.
type IAuctionFeeVaultValidatorPaybackRateUpdatedIterator struct {
	Event *IAuctionFeeVaultValidatorPaybackRateUpdated // Event containing the contract specifics and raw log

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
func (it *IAuctionFeeVaultValidatorPaybackRateUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IAuctionFeeVaultValidatorPaybackRateUpdated)
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
		it.Event = new(IAuctionFeeVaultValidatorPaybackRateUpdated)
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
func (it *IAuctionFeeVaultValidatorPaybackRateUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IAuctionFeeVaultValidatorPaybackRateUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IAuctionFeeVaultValidatorPaybackRateUpdated represents a ValidatorPaybackRateUpdated event raised by the IAuctionFeeVault contract.
type IAuctionFeeVaultValidatorPaybackRateUpdated struct {
	ValidatorPaybackRate *big.Int
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterValidatorPaybackRateUpdated is a free log retrieval operation binding the contract event 0x5309d48fe743a67ce32d8f66af9e2388d65bfc8cc026a4e1fbed3a4612a0af98.
//
// Solidity: event ValidatorPaybackRateUpdated(uint256 validatorPaybackRate)
func (_IAuctionFeeVault *IAuctionFeeVaultFilterer) FilterValidatorPaybackRateUpdated(opts *bind.FilterOpts) (*IAuctionFeeVaultValidatorPaybackRateUpdatedIterator, error) {

	logs, sub, err := _IAuctionFeeVault.contract.FilterLogs(opts, "ValidatorPaybackRateUpdated")
	if err != nil {
		return nil, err
	}
	return &IAuctionFeeVaultValidatorPaybackRateUpdatedIterator{contract: _IAuctionFeeVault.contract, event: "ValidatorPaybackRateUpdated", logs: logs, sub: sub}, nil
}

// WatchValidatorPaybackRateUpdated is a free log subscription operation binding the contract event 0x5309d48fe743a67ce32d8f66af9e2388d65bfc8cc026a4e1fbed3a4612a0af98.
//
// Solidity: event ValidatorPaybackRateUpdated(uint256 validatorPaybackRate)
func (_IAuctionFeeVault *IAuctionFeeVaultFilterer) WatchValidatorPaybackRateUpdated(opts *bind.WatchOpts, sink chan<- *IAuctionFeeVaultValidatorPaybackRateUpdated) (event.Subscription, error) {

	logs, sub, err := _IAuctionFeeVault.contract.WatchLogs(opts, "ValidatorPaybackRateUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IAuctionFeeVaultValidatorPaybackRateUpdated)
				if err := _IAuctionFeeVault.contract.UnpackLog(event, "ValidatorPaybackRateUpdated", log); err != nil {
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

// ParseValidatorPaybackRateUpdated is a log parse operation binding the contract event 0x5309d48fe743a67ce32d8f66af9e2388d65bfc8cc026a4e1fbed3a4612a0af98.
//
// Solidity: event ValidatorPaybackRateUpdated(uint256 validatorPaybackRate)
func (_IAuctionFeeVault *IAuctionFeeVaultFilterer) ParseValidatorPaybackRateUpdated(log types.Log) (*IAuctionFeeVaultValidatorPaybackRateUpdated, error) {
	event := new(IAuctionFeeVaultValidatorPaybackRateUpdated)
	if err := _IAuctionFeeVault.contract.UnpackLog(event, "ValidatorPaybackRateUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IStakingMetaData contains all meta data concerning the IStaking contract.
var IStakingMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_admin\",\"type\":\"address\"}],\"name\":\"isAdmin\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"24d7806c": "isAdmin(address)",
	},
}

// IStakingABI is the input ABI used to generate the binding from.
// Deprecated: Use IStakingMetaData.ABI instead.
var IStakingABI = IStakingMetaData.ABI

// IStakingBinRuntime is the compiled bytecode used for adding genesis block without deploying code.
const IStakingBinRuntime = ``

// Deprecated: Use IStakingMetaData.Sigs instead.
// IStakingFuncSigs maps the 4-byte function signature to its string representation.
var IStakingFuncSigs = IStakingMetaData.Sigs

// IStaking is an auto generated Go binding around a Kaia contract.
type IStaking struct {
	IStakingCaller     // Read-only binding to the contract
	IStakingTransactor // Write-only binding to the contract
	IStakingFilterer   // Log filterer for contract events
}

// IStakingCaller is an auto generated read-only Go binding around a Kaia contract.
type IStakingCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IStakingTransactor is an auto generated write-only Go binding around a Kaia contract.
type IStakingTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IStakingFilterer is an auto generated log filtering Go binding around a Kaia contract events.
type IStakingFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IStakingSession is an auto generated Go binding around a Kaia contract,
// with pre-set call and transact options.
type IStakingSession struct {
	Contract     *IStaking         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IStakingCallerSession is an auto generated read-only Go binding around a Kaia contract,
// with pre-set call options.
type IStakingCallerSession struct {
	Contract *IStakingCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// IStakingTransactorSession is an auto generated write-only Go binding around a Kaia contract,
// with pre-set transact options.
type IStakingTransactorSession struct {
	Contract     *IStakingTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// IStakingRaw is an auto generated low-level Go binding around a Kaia contract.
type IStakingRaw struct {
	Contract *IStaking // Generic contract binding to access the raw methods on
}

// IStakingCallerRaw is an auto generated low-level read-only Go binding around a Kaia contract.
type IStakingCallerRaw struct {
	Contract *IStakingCaller // Generic read-only contract binding to access the raw methods on
}

// IStakingTransactorRaw is an auto generated low-level write-only Go binding around a Kaia contract.
type IStakingTransactorRaw struct {
	Contract *IStakingTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIStaking creates a new instance of IStaking, bound to a specific deployed contract.
func NewIStaking(address common.Address, backend bind.ContractBackend) (*IStaking, error) {
	contract, err := bindIStaking(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IStaking{IStakingCaller: IStakingCaller{contract: contract}, IStakingTransactor: IStakingTransactor{contract: contract}, IStakingFilterer: IStakingFilterer{contract: contract}}, nil
}

// NewIStakingCaller creates a new read-only instance of IStaking, bound to a specific deployed contract.
func NewIStakingCaller(address common.Address, caller bind.ContractCaller) (*IStakingCaller, error) {
	contract, err := bindIStaking(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IStakingCaller{contract: contract}, nil
}

// NewIStakingTransactor creates a new write-only instance of IStaking, bound to a specific deployed contract.
func NewIStakingTransactor(address common.Address, transactor bind.ContractTransactor) (*IStakingTransactor, error) {
	contract, err := bindIStaking(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IStakingTransactor{contract: contract}, nil
}

// NewIStakingFilterer creates a new log filterer instance of IStaking, bound to a specific deployed contract.
func NewIStakingFilterer(address common.Address, filterer bind.ContractFilterer) (*IStakingFilterer, error) {
	contract, err := bindIStaking(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IStakingFilterer{contract: contract}, nil
}

// bindIStaking binds a generic wrapper to an already deployed contract.
func bindIStaking(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IStakingMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IStaking *IStakingRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IStaking.Contract.IStakingCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IStaking *IStakingRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IStaking.Contract.IStakingTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IStaking *IStakingRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IStaking.Contract.IStakingTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IStaking *IStakingCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IStaking.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IStaking *IStakingTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IStaking.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IStaking *IStakingTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IStaking.Contract.contract.Transact(opts, method, params...)
}

// IsAdmin is a free data retrieval call binding the contract method 0x24d7806c.
//
// Solidity: function isAdmin(address _admin) view returns(bool)
func (_IStaking *IStakingCaller) IsAdmin(opts *bind.CallOpts, _admin common.Address) (bool, error) {
	var out []interface{}
	err := _IStaking.contract.Call(opts, &out, "isAdmin", _admin)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsAdmin is a free data retrieval call binding the contract method 0x24d7806c.
//
// Solidity: function isAdmin(address _admin) view returns(bool)
func (_IStaking *IStakingSession) IsAdmin(_admin common.Address) (bool, error) {
	return _IStaking.Contract.IsAdmin(&_IStaking.CallOpts, _admin)
}

// IsAdmin is a free data retrieval call binding the contract method 0x24d7806c.
//
// Solidity: function isAdmin(address _admin) view returns(bool)
func (_IStaking *IStakingCallerSession) IsAdmin(_admin common.Address) (bool, error) {
	return _IStaking.Contract.IsAdmin(&_IStaking.CallOpts, _admin)
}

// OwnableMetaData contains all meta data concerning the Ownable contract.
var OwnableMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"8da5cb5b": "owner()",
		"715018a6": "renounceOwnership()",
		"f2fde38b": "transferOwnership(address)",
	},
}

// OwnableABI is the input ABI used to generate the binding from.
// Deprecated: Use OwnableMetaData.ABI instead.
var OwnableABI = OwnableMetaData.ABI

// OwnableBinRuntime is the compiled bytecode used for adding genesis block without deploying code.
const OwnableBinRuntime = ``

// Deprecated: Use OwnableMetaData.Sigs instead.
// OwnableFuncSigs maps the 4-byte function signature to its string representation.
var OwnableFuncSigs = OwnableMetaData.Sigs

// Ownable is an auto generated Go binding around a Kaia contract.
type Ownable struct {
	OwnableCaller     // Read-only binding to the contract
	OwnableTransactor // Write-only binding to the contract
	OwnableFilterer   // Log filterer for contract events
}

// OwnableCaller is an auto generated read-only Go binding around a Kaia contract.
type OwnableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnableTransactor is an auto generated write-only Go binding around a Kaia contract.
type OwnableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnableFilterer is an auto generated log filtering Go binding around a Kaia contract events.
type OwnableFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnableSession is an auto generated Go binding around a Kaia contract,
// with pre-set call and transact options.
type OwnableSession struct {
	Contract     *Ownable          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OwnableCallerSession is an auto generated read-only Go binding around a Kaia contract,
// with pre-set call options.
type OwnableCallerSession struct {
	Contract *OwnableCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// OwnableTransactorSession is an auto generated write-only Go binding around a Kaia contract,
// with pre-set transact options.
type OwnableTransactorSession struct {
	Contract     *OwnableTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// OwnableRaw is an auto generated low-level Go binding around a Kaia contract.
type OwnableRaw struct {
	Contract *Ownable // Generic contract binding to access the raw methods on
}

// OwnableCallerRaw is an auto generated low-level read-only Go binding around a Kaia contract.
type OwnableCallerRaw struct {
	Contract *OwnableCaller // Generic read-only contract binding to access the raw methods on
}

// OwnableTransactorRaw is an auto generated low-level write-only Go binding around a Kaia contract.
type OwnableTransactorRaw struct {
	Contract *OwnableTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOwnable creates a new instance of Ownable, bound to a specific deployed contract.
func NewOwnable(address common.Address, backend bind.ContractBackend) (*Ownable, error) {
	contract, err := bindOwnable(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Ownable{OwnableCaller: OwnableCaller{contract: contract}, OwnableTransactor: OwnableTransactor{contract: contract}, OwnableFilterer: OwnableFilterer{contract: contract}}, nil
}

// NewOwnableCaller creates a new read-only instance of Ownable, bound to a specific deployed contract.
func NewOwnableCaller(address common.Address, caller bind.ContractCaller) (*OwnableCaller, error) {
	contract, err := bindOwnable(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OwnableCaller{contract: contract}, nil
}

// NewOwnableTransactor creates a new write-only instance of Ownable, bound to a specific deployed contract.
func NewOwnableTransactor(address common.Address, transactor bind.ContractTransactor) (*OwnableTransactor, error) {
	contract, err := bindOwnable(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OwnableTransactor{contract: contract}, nil
}

// NewOwnableFilterer creates a new log filterer instance of Ownable, bound to a specific deployed contract.
func NewOwnableFilterer(address common.Address, filterer bind.ContractFilterer) (*OwnableFilterer, error) {
	contract, err := bindOwnable(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OwnableFilterer{contract: contract}, nil
}

// bindOwnable binds a generic wrapper to an already deployed contract.
func bindOwnable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := OwnableMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ownable *OwnableRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Ownable.Contract.OwnableCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ownable *OwnableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ownable.Contract.OwnableTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ownable *OwnableRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ownable.Contract.OwnableTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ownable *OwnableCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Ownable.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ownable *OwnableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ownable.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ownable *OwnableTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ownable.Contract.contract.Transact(opts, method, params...)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Ownable *OwnableCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Ownable.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Ownable *OwnableSession) Owner() (common.Address, error) {
	return _Ownable.Contract.Owner(&_Ownable.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Ownable *OwnableCallerSession) Owner() (common.Address, error) {
	return _Ownable.Contract.Owner(&_Ownable.CallOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Ownable *OwnableTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ownable.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Ownable *OwnableSession) RenounceOwnership() (*types.Transaction, error) {
	return _Ownable.Contract.RenounceOwnership(&_Ownable.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Ownable *OwnableTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Ownable.Contract.RenounceOwnership(&_Ownable.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Ownable *OwnableTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Ownable.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Ownable *OwnableSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Ownable.Contract.TransferOwnership(&_Ownable.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Ownable *OwnableTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Ownable.Contract.TransferOwnership(&_Ownable.TransactOpts, newOwner)
}

// OwnableOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Ownable contract.
type OwnableOwnershipTransferredIterator struct {
	Event *OwnableOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *OwnableOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OwnableOwnershipTransferred)
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
		it.Event = new(OwnableOwnershipTransferred)
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
func (it *OwnableOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OwnableOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OwnableOwnershipTransferred represents a OwnershipTransferred event raised by the Ownable contract.
type OwnableOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Ownable *OwnableFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*OwnableOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Ownable.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &OwnableOwnershipTransferredIterator{contract: _Ownable.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Ownable *OwnableFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *OwnableOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Ownable.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OwnableOwnershipTransferred)
				if err := _Ownable.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Ownable *OwnableFilterer) ParseOwnershipTransferred(log types.Log) (*OwnableOwnershipTransferred, error) {
	event := new(OwnableOwnershipTransferred)
	if err := _Ownable.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
