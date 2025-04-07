// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package paymaster

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// PackedUserOperation is an auto generated low-level Go binding around an user-defined struct.
type PackedUserOperation struct {
	Sender             common.Address
	Nonce              *big.Int
	InitCode           []byte
	CallData           []byte
	AccountGasLimits   [32]byte
	PreVerificationGas *big.Int
	GasFees            [32]byte
	PaymasterAndData   []byte
	Signature          []byte
}

// VerifyingPaymasterMetaData contains all meta data concerning the VerifyingPaymaster contract.
var VerifyingPaymasterMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractIEntryPoint\",\"name\":\"_entryPoint\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_verifyingSigner\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"ECDSAInvalidSignature\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"length\",\"type\":\"uint256\"}],\"name\":\"ECDSAInvalidSignatureLength\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"ECDSAInvalidSignatureS\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"unstakeDelaySec\",\"type\":\"uint32\"}],\"name\":\"addStake\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"entryPoint\",\"outputs\":[{\"internalType\":\"contractIEntryPoint\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getDeposit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"initCode\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"accountGasLimits\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"preVerificationGas\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"gasFees\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"paymasterAndData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structPackedUserOperation\",\"name\":\"userOp\",\"type\":\"tuple\"},{\"internalType\":\"uint48\",\"name\":\"validUntil\",\"type\":\"uint48\"},{\"internalType\":\"uint48\",\"name\":\"validAfter\",\"type\":\"uint48\"}],\"name\":\"getHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"paymasterAndData\",\"type\":\"bytes\"}],\"name\":\"parsePaymasterAndData\",\"outputs\":[{\"internalType\":\"uint48\",\"name\":\"validUntil\",\"type\":\"uint48\"},{\"internalType\":\"uint48\",\"name\":\"validAfter\",\"type\":\"uint48\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumIPaymaster.PostOpMode\",\"name\":\"mode\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"context\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"actualGasCost\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"actualUserOpFeePerGas\",\"type\":\"uint256\"}],\"name\":\"postOp\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unlockStake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"initCode\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"accountGasLimits\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"preVerificationGas\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"gasFees\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"paymasterAndData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structPackedUserOperation\",\"name\":\"userOp\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"userOpHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"maxCost\",\"type\":\"uint256\"}],\"name\":\"validatePaymasterUserOp\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"context\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"validationData\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"verifyingSigner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"withdrawAddress\",\"type\":\"address\"}],\"name\":\"withdrawStake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"withdrawAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdrawTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// VerifyingPaymasterABI is the input ABI used to generate the binding from.
// Deprecated: Use VerifyingPaymasterMetaData.ABI instead.
var VerifyingPaymasterABI = VerifyingPaymasterMetaData.ABI

// VerifyingPaymaster is an auto generated Go binding around an Ethereum contract.
type VerifyingPaymaster struct {
	VerifyingPaymasterCaller     // Read-only binding to the contract
	VerifyingPaymasterTransactor // Write-only binding to the contract
	VerifyingPaymasterFilterer   // Log filterer for contract events
}

// VerifyingPaymasterCaller is an auto generated read-only Go binding around an Ethereum contract.
type VerifyingPaymasterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VerifyingPaymasterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type VerifyingPaymasterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VerifyingPaymasterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type VerifyingPaymasterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VerifyingPaymasterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type VerifyingPaymasterSession struct {
	Contract     *VerifyingPaymaster // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// VerifyingPaymasterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type VerifyingPaymasterCallerSession struct {
	Contract *VerifyingPaymasterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// VerifyingPaymasterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type VerifyingPaymasterTransactorSession struct {
	Contract     *VerifyingPaymasterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// VerifyingPaymasterRaw is an auto generated low-level Go binding around an Ethereum contract.
type VerifyingPaymasterRaw struct {
	Contract *VerifyingPaymaster // Generic contract binding to access the raw methods on
}

// VerifyingPaymasterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type VerifyingPaymasterCallerRaw struct {
	Contract *VerifyingPaymasterCaller // Generic read-only contract binding to access the raw methods on
}

// VerifyingPaymasterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type VerifyingPaymasterTransactorRaw struct {
	Contract *VerifyingPaymasterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewVerifyingPaymaster creates a new instance of VerifyingPaymaster, bound to a specific deployed contract.
func NewVerifyingPaymaster(address common.Address, backend bind.ContractBackend) (*VerifyingPaymaster, error) {
	contract, err := bindVerifyingPaymaster(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &VerifyingPaymaster{VerifyingPaymasterCaller: VerifyingPaymasterCaller{contract: contract}, VerifyingPaymasterTransactor: VerifyingPaymasterTransactor{contract: contract}, VerifyingPaymasterFilterer: VerifyingPaymasterFilterer{contract: contract}}, nil
}

// NewVerifyingPaymasterCaller creates a new read-only instance of VerifyingPaymaster, bound to a specific deployed contract.
func NewVerifyingPaymasterCaller(address common.Address, caller bind.ContractCaller) (*VerifyingPaymasterCaller, error) {
	contract, err := bindVerifyingPaymaster(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &VerifyingPaymasterCaller{contract: contract}, nil
}

// NewVerifyingPaymasterTransactor creates a new write-only instance of VerifyingPaymaster, bound to a specific deployed contract.
func NewVerifyingPaymasterTransactor(address common.Address, transactor bind.ContractTransactor) (*VerifyingPaymasterTransactor, error) {
	contract, err := bindVerifyingPaymaster(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &VerifyingPaymasterTransactor{contract: contract}, nil
}

// NewVerifyingPaymasterFilterer creates a new log filterer instance of VerifyingPaymaster, bound to a specific deployed contract.
func NewVerifyingPaymasterFilterer(address common.Address, filterer bind.ContractFilterer) (*VerifyingPaymasterFilterer, error) {
	contract, err := bindVerifyingPaymaster(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &VerifyingPaymasterFilterer{contract: contract}, nil
}

// bindVerifyingPaymaster binds a generic wrapper to an already deployed contract.
func bindVerifyingPaymaster(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := VerifyingPaymasterMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VerifyingPaymaster *VerifyingPaymasterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VerifyingPaymaster.Contract.VerifyingPaymasterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VerifyingPaymaster *VerifyingPaymasterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VerifyingPaymaster.Contract.VerifyingPaymasterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VerifyingPaymaster *VerifyingPaymasterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VerifyingPaymaster.Contract.VerifyingPaymasterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VerifyingPaymaster *VerifyingPaymasterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VerifyingPaymaster.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VerifyingPaymaster *VerifyingPaymasterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VerifyingPaymaster.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VerifyingPaymaster *VerifyingPaymasterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VerifyingPaymaster.Contract.contract.Transact(opts, method, params...)
}

// EntryPoint is a free data retrieval call binding the contract method 0xb0d691fe.
//
// Solidity: function entryPoint() view returns(address)
func (_VerifyingPaymaster *VerifyingPaymasterCaller) EntryPoint(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VerifyingPaymaster.contract.Call(opts, &out, "entryPoint")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// EntryPoint is a free data retrieval call binding the contract method 0xb0d691fe.
//
// Solidity: function entryPoint() view returns(address)
func (_VerifyingPaymaster *VerifyingPaymasterSession) EntryPoint() (common.Address, error) {
	return _VerifyingPaymaster.Contract.EntryPoint(&_VerifyingPaymaster.CallOpts)
}

// EntryPoint is a free data retrieval call binding the contract method 0xb0d691fe.
//
// Solidity: function entryPoint() view returns(address)
func (_VerifyingPaymaster *VerifyingPaymasterCallerSession) EntryPoint() (common.Address, error) {
	return _VerifyingPaymaster.Contract.EntryPoint(&_VerifyingPaymaster.CallOpts)
}

// GetDeposit is a free data retrieval call binding the contract method 0xc399ec88.
//
// Solidity: function getDeposit() view returns(uint256)
func (_VerifyingPaymaster *VerifyingPaymasterCaller) GetDeposit(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VerifyingPaymaster.contract.Call(opts, &out, "getDeposit")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetDeposit is a free data retrieval call binding the contract method 0xc399ec88.
//
// Solidity: function getDeposit() view returns(uint256)
func (_VerifyingPaymaster *VerifyingPaymasterSession) GetDeposit() (*big.Int, error) {
	return _VerifyingPaymaster.Contract.GetDeposit(&_VerifyingPaymaster.CallOpts)
}

// GetDeposit is a free data retrieval call binding the contract method 0xc399ec88.
//
// Solidity: function getDeposit() view returns(uint256)
func (_VerifyingPaymaster *VerifyingPaymasterCallerSession) GetDeposit() (*big.Int, error) {
	return _VerifyingPaymaster.Contract.GetDeposit(&_VerifyingPaymaster.CallOpts)
}

// GetHash is a free data retrieval call binding the contract method 0x5829c5f5.
//
// Solidity: function getHash((address,uint256,bytes,bytes,bytes32,uint256,bytes32,bytes,bytes) userOp, uint48 validUntil, uint48 validAfter) view returns(bytes32)
func (_VerifyingPaymaster *VerifyingPaymasterCaller) GetHash(opts *bind.CallOpts, userOp PackedUserOperation, validUntil *big.Int, validAfter *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _VerifyingPaymaster.contract.Call(opts, &out, "getHash", userOp, validUntil, validAfter)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetHash is a free data retrieval call binding the contract method 0x5829c5f5.
//
// Solidity: function getHash((address,uint256,bytes,bytes,bytes32,uint256,bytes32,bytes,bytes) userOp, uint48 validUntil, uint48 validAfter) view returns(bytes32)
func (_VerifyingPaymaster *VerifyingPaymasterSession) GetHash(userOp PackedUserOperation, validUntil *big.Int, validAfter *big.Int) ([32]byte, error) {
	return _VerifyingPaymaster.Contract.GetHash(&_VerifyingPaymaster.CallOpts, userOp, validUntil, validAfter)
}

// GetHash is a free data retrieval call binding the contract method 0x5829c5f5.
//
// Solidity: function getHash((address,uint256,bytes,bytes,bytes32,uint256,bytes32,bytes,bytes) userOp, uint48 validUntil, uint48 validAfter) view returns(bytes32)
func (_VerifyingPaymaster *VerifyingPaymasterCallerSession) GetHash(userOp PackedUserOperation, validUntil *big.Int, validAfter *big.Int) ([32]byte, error) {
	return _VerifyingPaymaster.Contract.GetHash(&_VerifyingPaymaster.CallOpts, userOp, validUntil, validAfter)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_VerifyingPaymaster *VerifyingPaymasterCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VerifyingPaymaster.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_VerifyingPaymaster *VerifyingPaymasterSession) Owner() (common.Address, error) {
	return _VerifyingPaymaster.Contract.Owner(&_VerifyingPaymaster.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_VerifyingPaymaster *VerifyingPaymasterCallerSession) Owner() (common.Address, error) {
	return _VerifyingPaymaster.Contract.Owner(&_VerifyingPaymaster.CallOpts)
}

// ParsePaymasterAndData is a free data retrieval call binding the contract method 0x94d4ad60.
//
// Solidity: function parsePaymasterAndData(bytes paymasterAndData) pure returns(uint48 validUntil, uint48 validAfter, bytes signature)
func (_VerifyingPaymaster *VerifyingPaymasterCaller) ParsePaymasterAndData(opts *bind.CallOpts, paymasterAndData []byte) (struct {
	ValidUntil *big.Int
	ValidAfter *big.Int
	Signature  []byte
}, error) {
	var out []interface{}
	err := _VerifyingPaymaster.contract.Call(opts, &out, "parsePaymasterAndData", paymasterAndData)

	outstruct := new(struct {
		ValidUntil *big.Int
		ValidAfter *big.Int
		Signature  []byte
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.ValidUntil = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.ValidAfter = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Signature = *abi.ConvertType(out[2], new([]byte)).(*[]byte)

	return *outstruct, err

}

// ParsePaymasterAndData is a free data retrieval call binding the contract method 0x94d4ad60.
//
// Solidity: function parsePaymasterAndData(bytes paymasterAndData) pure returns(uint48 validUntil, uint48 validAfter, bytes signature)
func (_VerifyingPaymaster *VerifyingPaymasterSession) ParsePaymasterAndData(paymasterAndData []byte) (struct {
	ValidUntil *big.Int
	ValidAfter *big.Int
	Signature  []byte
}, error) {
	return _VerifyingPaymaster.Contract.ParsePaymasterAndData(&_VerifyingPaymaster.CallOpts, paymasterAndData)
}

// ParsePaymasterAndData is a free data retrieval call binding the contract method 0x94d4ad60.
//
// Solidity: function parsePaymasterAndData(bytes paymasterAndData) pure returns(uint48 validUntil, uint48 validAfter, bytes signature)
func (_VerifyingPaymaster *VerifyingPaymasterCallerSession) ParsePaymasterAndData(paymasterAndData []byte) (struct {
	ValidUntil *big.Int
	ValidAfter *big.Int
	Signature  []byte
}, error) {
	return _VerifyingPaymaster.Contract.ParsePaymasterAndData(&_VerifyingPaymaster.CallOpts, paymasterAndData)
}

// VerifyingSigner is a free data retrieval call binding the contract method 0x23d9ac9b.
//
// Solidity: function verifyingSigner() view returns(address)
func (_VerifyingPaymaster *VerifyingPaymasterCaller) VerifyingSigner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VerifyingPaymaster.contract.Call(opts, &out, "verifyingSigner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// VerifyingSigner is a free data retrieval call binding the contract method 0x23d9ac9b.
//
// Solidity: function verifyingSigner() view returns(address)
func (_VerifyingPaymaster *VerifyingPaymasterSession) VerifyingSigner() (common.Address, error) {
	return _VerifyingPaymaster.Contract.VerifyingSigner(&_VerifyingPaymaster.CallOpts)
}

// VerifyingSigner is a free data retrieval call binding the contract method 0x23d9ac9b.
//
// Solidity: function verifyingSigner() view returns(address)
func (_VerifyingPaymaster *VerifyingPaymasterCallerSession) VerifyingSigner() (common.Address, error) {
	return _VerifyingPaymaster.Contract.VerifyingSigner(&_VerifyingPaymaster.CallOpts)
}

// AddStake is a paid mutator transaction binding the contract method 0x0396cb60.
//
// Solidity: function addStake(uint32 unstakeDelaySec) payable returns()
func (_VerifyingPaymaster *VerifyingPaymasterTransactor) AddStake(opts *bind.TransactOpts, unstakeDelaySec uint32) (*types.Transaction, error) {
	return _VerifyingPaymaster.contract.Transact(opts, "addStake", unstakeDelaySec)
}

// AddStake is a paid mutator transaction binding the contract method 0x0396cb60.
//
// Solidity: function addStake(uint32 unstakeDelaySec) payable returns()
func (_VerifyingPaymaster *VerifyingPaymasterSession) AddStake(unstakeDelaySec uint32) (*types.Transaction, error) {
	return _VerifyingPaymaster.Contract.AddStake(&_VerifyingPaymaster.TransactOpts, unstakeDelaySec)
}

// AddStake is a paid mutator transaction binding the contract method 0x0396cb60.
//
// Solidity: function addStake(uint32 unstakeDelaySec) payable returns()
func (_VerifyingPaymaster *VerifyingPaymasterTransactorSession) AddStake(unstakeDelaySec uint32) (*types.Transaction, error) {
	return _VerifyingPaymaster.Contract.AddStake(&_VerifyingPaymaster.TransactOpts, unstakeDelaySec)
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns()
func (_VerifyingPaymaster *VerifyingPaymasterTransactor) Deposit(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VerifyingPaymaster.contract.Transact(opts, "deposit")
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns()
func (_VerifyingPaymaster *VerifyingPaymasterSession) Deposit() (*types.Transaction, error) {
	return _VerifyingPaymaster.Contract.Deposit(&_VerifyingPaymaster.TransactOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns()
func (_VerifyingPaymaster *VerifyingPaymasterTransactorSession) Deposit() (*types.Transaction, error) {
	return _VerifyingPaymaster.Contract.Deposit(&_VerifyingPaymaster.TransactOpts)
}

// PostOp is a paid mutator transaction binding the contract method 0x7c627b21.
//
// Solidity: function postOp(uint8 mode, bytes context, uint256 actualGasCost, uint256 actualUserOpFeePerGas) returns()
func (_VerifyingPaymaster *VerifyingPaymasterTransactor) PostOp(opts *bind.TransactOpts, mode uint8, context []byte, actualGasCost *big.Int, actualUserOpFeePerGas *big.Int) (*types.Transaction, error) {
	return _VerifyingPaymaster.contract.Transact(opts, "postOp", mode, context, actualGasCost, actualUserOpFeePerGas)
}

// PostOp is a paid mutator transaction binding the contract method 0x7c627b21.
//
// Solidity: function postOp(uint8 mode, bytes context, uint256 actualGasCost, uint256 actualUserOpFeePerGas) returns()
func (_VerifyingPaymaster *VerifyingPaymasterSession) PostOp(mode uint8, context []byte, actualGasCost *big.Int, actualUserOpFeePerGas *big.Int) (*types.Transaction, error) {
	return _VerifyingPaymaster.Contract.PostOp(&_VerifyingPaymaster.TransactOpts, mode, context, actualGasCost, actualUserOpFeePerGas)
}

// PostOp is a paid mutator transaction binding the contract method 0x7c627b21.
//
// Solidity: function postOp(uint8 mode, bytes context, uint256 actualGasCost, uint256 actualUserOpFeePerGas) returns()
func (_VerifyingPaymaster *VerifyingPaymasterTransactorSession) PostOp(mode uint8, context []byte, actualGasCost *big.Int, actualUserOpFeePerGas *big.Int) (*types.Transaction, error) {
	return _VerifyingPaymaster.Contract.PostOp(&_VerifyingPaymaster.TransactOpts, mode, context, actualGasCost, actualUserOpFeePerGas)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_VerifyingPaymaster *VerifyingPaymasterTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VerifyingPaymaster.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_VerifyingPaymaster *VerifyingPaymasterSession) RenounceOwnership() (*types.Transaction, error) {
	return _VerifyingPaymaster.Contract.RenounceOwnership(&_VerifyingPaymaster.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_VerifyingPaymaster *VerifyingPaymasterTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _VerifyingPaymaster.Contract.RenounceOwnership(&_VerifyingPaymaster.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_VerifyingPaymaster *VerifyingPaymasterTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _VerifyingPaymaster.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_VerifyingPaymaster *VerifyingPaymasterSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _VerifyingPaymaster.Contract.TransferOwnership(&_VerifyingPaymaster.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_VerifyingPaymaster *VerifyingPaymasterTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _VerifyingPaymaster.Contract.TransferOwnership(&_VerifyingPaymaster.TransactOpts, newOwner)
}

// UnlockStake is a paid mutator transaction binding the contract method 0xbb9fe6bf.
//
// Solidity: function unlockStake() returns()
func (_VerifyingPaymaster *VerifyingPaymasterTransactor) UnlockStake(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VerifyingPaymaster.contract.Transact(opts, "unlockStake")
}

// UnlockStake is a paid mutator transaction binding the contract method 0xbb9fe6bf.
//
// Solidity: function unlockStake() returns()
func (_VerifyingPaymaster *VerifyingPaymasterSession) UnlockStake() (*types.Transaction, error) {
	return _VerifyingPaymaster.Contract.UnlockStake(&_VerifyingPaymaster.TransactOpts)
}

// UnlockStake is a paid mutator transaction binding the contract method 0xbb9fe6bf.
//
// Solidity: function unlockStake() returns()
func (_VerifyingPaymaster *VerifyingPaymasterTransactorSession) UnlockStake() (*types.Transaction, error) {
	return _VerifyingPaymaster.Contract.UnlockStake(&_VerifyingPaymaster.TransactOpts)
}

// ValidatePaymasterUserOp is a paid mutator transaction binding the contract method 0x52b7512c.
//
// Solidity: function validatePaymasterUserOp((address,uint256,bytes,bytes,bytes32,uint256,bytes32,bytes,bytes) userOp, bytes32 userOpHash, uint256 maxCost) returns(bytes context, uint256 validationData)
func (_VerifyingPaymaster *VerifyingPaymasterTransactor) ValidatePaymasterUserOp(opts *bind.TransactOpts, userOp PackedUserOperation, userOpHash [32]byte, maxCost *big.Int) (*types.Transaction, error) {
	return _VerifyingPaymaster.contract.Transact(opts, "validatePaymasterUserOp", userOp, userOpHash, maxCost)
}

// ValidatePaymasterUserOp is a paid mutator transaction binding the contract method 0x52b7512c.
//
// Solidity: function validatePaymasterUserOp((address,uint256,bytes,bytes,bytes32,uint256,bytes32,bytes,bytes) userOp, bytes32 userOpHash, uint256 maxCost) returns(bytes context, uint256 validationData)
func (_VerifyingPaymaster *VerifyingPaymasterSession) ValidatePaymasterUserOp(userOp PackedUserOperation, userOpHash [32]byte, maxCost *big.Int) (*types.Transaction, error) {
	return _VerifyingPaymaster.Contract.ValidatePaymasterUserOp(&_VerifyingPaymaster.TransactOpts, userOp, userOpHash, maxCost)
}

// ValidatePaymasterUserOp is a paid mutator transaction binding the contract method 0x52b7512c.
//
// Solidity: function validatePaymasterUserOp((address,uint256,bytes,bytes,bytes32,uint256,bytes32,bytes,bytes) userOp, bytes32 userOpHash, uint256 maxCost) returns(bytes context, uint256 validationData)
func (_VerifyingPaymaster *VerifyingPaymasterTransactorSession) ValidatePaymasterUserOp(userOp PackedUserOperation, userOpHash [32]byte, maxCost *big.Int) (*types.Transaction, error) {
	return _VerifyingPaymaster.Contract.ValidatePaymasterUserOp(&_VerifyingPaymaster.TransactOpts, userOp, userOpHash, maxCost)
}

// WithdrawStake is a paid mutator transaction binding the contract method 0xc23a5cea.
//
// Solidity: function withdrawStake(address withdrawAddress) returns()
func (_VerifyingPaymaster *VerifyingPaymasterTransactor) WithdrawStake(opts *bind.TransactOpts, withdrawAddress common.Address) (*types.Transaction, error) {
	return _VerifyingPaymaster.contract.Transact(opts, "withdrawStake", withdrawAddress)
}

// WithdrawStake is a paid mutator transaction binding the contract method 0xc23a5cea.
//
// Solidity: function withdrawStake(address withdrawAddress) returns()
func (_VerifyingPaymaster *VerifyingPaymasterSession) WithdrawStake(withdrawAddress common.Address) (*types.Transaction, error) {
	return _VerifyingPaymaster.Contract.WithdrawStake(&_VerifyingPaymaster.TransactOpts, withdrawAddress)
}

// WithdrawStake is a paid mutator transaction binding the contract method 0xc23a5cea.
//
// Solidity: function withdrawStake(address withdrawAddress) returns()
func (_VerifyingPaymaster *VerifyingPaymasterTransactorSession) WithdrawStake(withdrawAddress common.Address) (*types.Transaction, error) {
	return _VerifyingPaymaster.Contract.WithdrawStake(&_VerifyingPaymaster.TransactOpts, withdrawAddress)
}

// WithdrawTo is a paid mutator transaction binding the contract method 0x205c2878.
//
// Solidity: function withdrawTo(address withdrawAddress, uint256 amount) returns()
func (_VerifyingPaymaster *VerifyingPaymasterTransactor) WithdrawTo(opts *bind.TransactOpts, withdrawAddress common.Address, amount *big.Int) (*types.Transaction, error) {
	return _VerifyingPaymaster.contract.Transact(opts, "withdrawTo", withdrawAddress, amount)
}

// WithdrawTo is a paid mutator transaction binding the contract method 0x205c2878.
//
// Solidity: function withdrawTo(address withdrawAddress, uint256 amount) returns()
func (_VerifyingPaymaster *VerifyingPaymasterSession) WithdrawTo(withdrawAddress common.Address, amount *big.Int) (*types.Transaction, error) {
	return _VerifyingPaymaster.Contract.WithdrawTo(&_VerifyingPaymaster.TransactOpts, withdrawAddress, amount)
}

// WithdrawTo is a paid mutator transaction binding the contract method 0x205c2878.
//
// Solidity: function withdrawTo(address withdrawAddress, uint256 amount) returns()
func (_VerifyingPaymaster *VerifyingPaymasterTransactorSession) WithdrawTo(withdrawAddress common.Address, amount *big.Int) (*types.Transaction, error) {
	return _VerifyingPaymaster.Contract.WithdrawTo(&_VerifyingPaymaster.TransactOpts, withdrawAddress, amount)
}

// VerifyingPaymasterOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the VerifyingPaymaster contract.
type VerifyingPaymasterOwnershipTransferredIterator struct {
	Event *VerifyingPaymasterOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *VerifyingPaymasterOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VerifyingPaymasterOwnershipTransferred)
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
		it.Event = new(VerifyingPaymasterOwnershipTransferred)
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
func (it *VerifyingPaymasterOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VerifyingPaymasterOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VerifyingPaymasterOwnershipTransferred represents a OwnershipTransferred event raised by the VerifyingPaymaster contract.
type VerifyingPaymasterOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_VerifyingPaymaster *VerifyingPaymasterFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*VerifyingPaymasterOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _VerifyingPaymaster.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &VerifyingPaymasterOwnershipTransferredIterator{contract: _VerifyingPaymaster.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_VerifyingPaymaster *VerifyingPaymasterFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *VerifyingPaymasterOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _VerifyingPaymaster.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VerifyingPaymasterOwnershipTransferred)
				if err := _VerifyingPaymaster.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_VerifyingPaymaster *VerifyingPaymasterFilterer) ParseOwnershipTransferred(log types.Log) (*VerifyingPaymasterOwnershipTransferred, error) {
	event := new(VerifyingPaymasterOwnershipTransferred)
	if err := _VerifyingPaymaster.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
