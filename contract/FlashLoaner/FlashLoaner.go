// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package FlashLoaner

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
)

// FlashLoanerABI is the input ABI used to generate the binding from.
// Deprecated: Use FlashLoanerMetaData.ABI instead.
var FlashLoanerABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_earnAddress\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_earnAddress\",\"type\":\"address\"}],\"name\":\"UpdateArgs\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"UpdateOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint160\",\"name\":\"_number\",\"type\":\"uint160\"}],\"name\":\"UpdateSalt\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount0Out\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_minProfit\",\"type\":\"uint256\"},{\"internalType\":\"uint160\",\"name\":\"_pair0\",\"type\":\"uint160\"},{\"internalType\":\"uint160\",\"name\":\"_pair1\",\"type\":\"uint160\"},{\"internalType\":\"uint160\",\"name\":\"_ts\",\"type\":\"uint160\"},{\"internalType\":\"uint160\",\"name\":\"_fee0\",\"type\":\"uint160\"},{\"internalType\":\"uint160\",\"name\":\"_fee1\",\"type\":\"uint160\"}],\"name\":\"Amount0CFlash\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount0Out\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_minProfit\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_pair0\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_pair1\",\"type\":\"address\"},{\"internalType\":\"uint160\",\"name\":\"_fee0\",\"type\":\"uint160\"},{\"internalType\":\"uint160\",\"name\":\"_fee1\",\"type\":\"uint160\"}],\"name\":\"Amount0Flash\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount1Out\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_minProfit\",\"type\":\"uint256\"},{\"internalType\":\"uint160\",\"name\":\"_pair0\",\"type\":\"uint160\"},{\"internalType\":\"uint160\",\"name\":\"_pair1\",\"type\":\"uint160\"},{\"internalType\":\"uint160\",\"name\":\"_ts\",\"type\":\"uint160\"},{\"internalType\":\"uint160\",\"name\":\"_fee0\",\"type\":\"uint160\"},{\"internalType\":\"uint160\",\"name\":\"_fee1\",\"type\":\"uint160\"}],\"name\":\"Amount1CFlash\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount1Out\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_minProfit\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_pair0\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_pair1\",\"type\":\"address\"},{\"internalType\":\"uint160\",\"name\":\"_fee0\",\"type\":\"uint160\"},{\"internalType\":\"uint160\",\"name\":\"_fee1\",\"type\":\"uint160\"}],\"name\":\"Amount1Flash\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount0Out\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_minProfit\",\"type\":\"uint256\"},{\"internalType\":\"uint160\",\"name\":\"_pair0\",\"type\":\"uint160\"},{\"internalType\":\"uint160\",\"name\":\"_pair1\",\"type\":\"uint160\"},{\"internalType\":\"uint160\",\"name\":\"_ts\",\"type\":\"uint160\"}],\"name\":\"testAmount0Flash\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"ts0r1In\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"ts1r0Out\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"s0r0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"s0r1\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"s1r0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"s1r1\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"win\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount1Out\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_minProfit\",\"type\":\"uint256\"},{\"internalType\":\"uint160\",\"name\":\"_pair0\",\"type\":\"uint160\"},{\"internalType\":\"uint160\",\"name\":\"_pair1\",\"type\":\"uint160\"},{\"internalType\":\"uint160\",\"name\":\"_ts\",\"type\":\"uint160\"}],\"name\":\"testAmount1Flash\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"ts0r0In\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"ts1r1Out\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"s0r0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"s0r1\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"s1r0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"s1r1\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"win\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_amount1\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"pancakeCall\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// FlashLoaner is an auto generated Go binding around an Ethereum contract.
type FlashLoaner struct {
	FlashLoanerCaller     // Read-only binding to the contract
	FlashLoanerTransactor // Write-only binding to the contract
	FlashLoanerFilterer   // Log filterer for contract events
}

// FlashLoanerCaller is an auto generated read-only Go binding around an Ethereum contract.
type FlashLoanerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FlashLoanerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type FlashLoanerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FlashLoanerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type FlashLoanerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FlashLoanerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type FlashLoanerSession struct {
	Contract     *FlashLoaner      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// FlashLoanerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type FlashLoanerCallerSession struct {
	Contract *FlashLoanerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// FlashLoanerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type FlashLoanerTransactorSession struct {
	Contract     *FlashLoanerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// FlashLoanerRaw is an auto generated low-level Go binding around an Ethereum contract.
type FlashLoanerRaw struct {
	Contract *FlashLoaner // Generic contract binding to access the raw methods on
}

// FlashLoanerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type FlashLoanerCallerRaw struct {
	Contract *FlashLoanerCaller // Generic read-only contract binding to access the raw methods on
}

// FlashLoanerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type FlashLoanerTransactorRaw struct {
	Contract *FlashLoanerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewFlashLoaner creates a new instance of FlashLoaner, bound to a specific deployed contract.
func NewFlashLoaner(address common.Address, backend bind.ContractBackend) (*FlashLoaner, error) {
	contract, err := bindFlashLoaner(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &FlashLoaner{FlashLoanerCaller: FlashLoanerCaller{contract: contract}, FlashLoanerTransactor: FlashLoanerTransactor{contract: contract}, FlashLoanerFilterer: FlashLoanerFilterer{contract: contract}}, nil
}

// NewFlashLoanerCaller creates a new read-only instance of FlashLoaner, bound to a specific deployed contract.
func NewFlashLoanerCaller(address common.Address, caller bind.ContractCaller) (*FlashLoanerCaller, error) {
	contract, err := bindFlashLoaner(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &FlashLoanerCaller{contract: contract}, nil
}

// NewFlashLoanerTransactor creates a new write-only instance of FlashLoaner, bound to a specific deployed contract.
func NewFlashLoanerTransactor(address common.Address, transactor bind.ContractTransactor) (*FlashLoanerTransactor, error) {
	contract, err := bindFlashLoaner(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &FlashLoanerTransactor{contract: contract}, nil
}

// NewFlashLoanerFilterer creates a new log filterer instance of FlashLoaner, bound to a specific deployed contract.
func NewFlashLoanerFilterer(address common.Address, filterer bind.ContractFilterer) (*FlashLoanerFilterer, error) {
	contract, err := bindFlashLoaner(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &FlashLoanerFilterer{contract: contract}, nil
}

// bindFlashLoaner binds a generic wrapper to an already deployed contract.
func bindFlashLoaner(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(FlashLoanerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FlashLoaner *FlashLoanerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FlashLoaner.Contract.FlashLoanerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FlashLoaner *FlashLoanerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FlashLoaner.Contract.FlashLoanerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FlashLoaner *FlashLoanerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FlashLoaner.Contract.FlashLoanerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FlashLoaner *FlashLoanerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FlashLoaner.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FlashLoaner *FlashLoanerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FlashLoaner.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FlashLoaner *FlashLoanerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FlashLoaner.Contract.contract.Transact(opts, method, params...)
}

// TestAmount0Flash is a free data retrieval call binding the contract method 0xad83a740.
//
// Solidity: function testAmount0Flash(uint256 _amount0Out, uint256 _minProfit, uint160 _pair0, uint160 _pair1, uint160 _ts) view returns(uint256 ts0r1In, uint256 ts1r0Out, uint256 s0r0, uint256 s0r1, uint256 s1r0, uint256 s1r1, bool win)
func (_FlashLoaner *FlashLoanerCaller) TestAmount0Flash(opts *bind.CallOpts, _amount0Out *big.Int, _minProfit *big.Int, _pair0 *big.Int, _pair1 *big.Int, _ts *big.Int) (struct {
	Ts0r1In  *big.Int
	Ts1r0Out *big.Int
	S0r0     *big.Int
	S0r1     *big.Int
	S1r0     *big.Int
	S1r1     *big.Int
	Win      bool
}, error) {
	var out []interface{}
	err := _FlashLoaner.contract.Call(opts, &out, "testAmount0Flash", _amount0Out, _minProfit, _pair0, _pair1, _ts)

	outstruct := new(struct {
		Ts0r1In  *big.Int
		Ts1r0Out *big.Int
		S0r0     *big.Int
		S0r1     *big.Int
		S1r0     *big.Int
		S1r1     *big.Int
		Win      bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Ts0r1In = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Ts1r0Out = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.S0r0 = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.S0r1 = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.S1r0 = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.S1r1 = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	outstruct.Win = *abi.ConvertType(out[6], new(bool)).(*bool)

	return *outstruct, err

}

// TestAmount0Flash is a free data retrieval call binding the contract method 0xad83a740.
//
// Solidity: function testAmount0Flash(uint256 _amount0Out, uint256 _minProfit, uint160 _pair0, uint160 _pair1, uint160 _ts) view returns(uint256 ts0r1In, uint256 ts1r0Out, uint256 s0r0, uint256 s0r1, uint256 s1r0, uint256 s1r1, bool win)
func (_FlashLoaner *FlashLoanerSession) TestAmount0Flash(_amount0Out *big.Int, _minProfit *big.Int, _pair0 *big.Int, _pair1 *big.Int, _ts *big.Int) (struct {
	Ts0r1In  *big.Int
	Ts1r0Out *big.Int
	S0r0     *big.Int
	S0r1     *big.Int
	S1r0     *big.Int
	S1r1     *big.Int
	Win      bool
}, error) {
	return _FlashLoaner.Contract.TestAmount0Flash(&_FlashLoaner.CallOpts, _amount0Out, _minProfit, _pair0, _pair1, _ts)
}

// TestAmount0Flash is a free data retrieval call binding the contract method 0xad83a740.
//
// Solidity: function testAmount0Flash(uint256 _amount0Out, uint256 _minProfit, uint160 _pair0, uint160 _pair1, uint160 _ts) view returns(uint256 ts0r1In, uint256 ts1r0Out, uint256 s0r0, uint256 s0r1, uint256 s1r0, uint256 s1r1, bool win)
func (_FlashLoaner *FlashLoanerCallerSession) TestAmount0Flash(_amount0Out *big.Int, _minProfit *big.Int, _pair0 *big.Int, _pair1 *big.Int, _ts *big.Int) (struct {
	Ts0r1In  *big.Int
	Ts1r0Out *big.Int
	S0r0     *big.Int
	S0r1     *big.Int
	S1r0     *big.Int
	S1r1     *big.Int
	Win      bool
}, error) {
	return _FlashLoaner.Contract.TestAmount0Flash(&_FlashLoaner.CallOpts, _amount0Out, _minProfit, _pair0, _pair1, _ts)
}

// TestAmount1Flash is a free data retrieval call binding the contract method 0xac41c2d0.
//
// Solidity: function testAmount1Flash(uint256 _amount1Out, uint256 _minProfit, uint160 _pair0, uint160 _pair1, uint160 _ts) view returns(uint256 ts0r0In, uint256 ts1r1Out, uint256 s0r0, uint256 s0r1, uint256 s1r0, uint256 s1r1, bool win)
func (_FlashLoaner *FlashLoanerCaller) TestAmount1Flash(opts *bind.CallOpts, _amount1Out *big.Int, _minProfit *big.Int, _pair0 *big.Int, _pair1 *big.Int, _ts *big.Int) (struct {
	Ts0r0In  *big.Int
	Ts1r1Out *big.Int
	S0r0     *big.Int
	S0r1     *big.Int
	S1r0     *big.Int
	S1r1     *big.Int
	Win      bool
}, error) {
	var out []interface{}
	err := _FlashLoaner.contract.Call(opts, &out, "testAmount1Flash", _amount1Out, _minProfit, _pair0, _pair1, _ts)

	outstruct := new(struct {
		Ts0r0In  *big.Int
		Ts1r1Out *big.Int
		S0r0     *big.Int
		S0r1     *big.Int
		S1r0     *big.Int
		S1r1     *big.Int
		Win      bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Ts0r0In = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Ts1r1Out = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.S0r0 = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.S0r1 = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.S1r0 = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.S1r1 = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	outstruct.Win = *abi.ConvertType(out[6], new(bool)).(*bool)

	return *outstruct, err

}

// TestAmount1Flash is a free data retrieval call binding the contract method 0xac41c2d0.
//
// Solidity: function testAmount1Flash(uint256 _amount1Out, uint256 _minProfit, uint160 _pair0, uint160 _pair1, uint160 _ts) view returns(uint256 ts0r0In, uint256 ts1r1Out, uint256 s0r0, uint256 s0r1, uint256 s1r0, uint256 s1r1, bool win)
func (_FlashLoaner *FlashLoanerSession) TestAmount1Flash(_amount1Out *big.Int, _minProfit *big.Int, _pair0 *big.Int, _pair1 *big.Int, _ts *big.Int) (struct {
	Ts0r0In  *big.Int
	Ts1r1Out *big.Int
	S0r0     *big.Int
	S0r1     *big.Int
	S1r0     *big.Int
	S1r1     *big.Int
	Win      bool
}, error) {
	return _FlashLoaner.Contract.TestAmount1Flash(&_FlashLoaner.CallOpts, _amount1Out, _minProfit, _pair0, _pair1, _ts)
}

// TestAmount1Flash is a free data retrieval call binding the contract method 0xac41c2d0.
//
// Solidity: function testAmount1Flash(uint256 _amount1Out, uint256 _minProfit, uint160 _pair0, uint160 _pair1, uint160 _ts) view returns(uint256 ts0r0In, uint256 ts1r1Out, uint256 s0r0, uint256 s0r1, uint256 s1r0, uint256 s1r1, bool win)
func (_FlashLoaner *FlashLoanerCallerSession) TestAmount1Flash(_amount1Out *big.Int, _minProfit *big.Int, _pair0 *big.Int, _pair1 *big.Int, _ts *big.Int) (struct {
	Ts0r0In  *big.Int
	Ts1r1Out *big.Int
	S0r0     *big.Int
	S0r1     *big.Int
	S1r0     *big.Int
	S1r1     *big.Int
	Win      bool
}, error) {
	return _FlashLoaner.Contract.TestAmount1Flash(&_FlashLoaner.CallOpts, _amount1Out, _minProfit, _pair0, _pair1, _ts)
}

// Amount0CFlash is a paid mutator transaction binding the contract method 0xa32a0803.
//
// Solidity: function Amount0CFlash(uint256 _amount0Out, uint256 _minProfit, uint160 _pair0, uint160 _pair1, uint160 _ts, uint160 _fee0, uint160 _fee1) returns()
func (_FlashLoaner *FlashLoanerTransactor) Amount0CFlash(opts *bind.TransactOpts, _amount0Out *big.Int, _minProfit *big.Int, _pair0 *big.Int, _pair1 *big.Int, _ts *big.Int, _fee0 *big.Int, _fee1 *big.Int) (*types.Transaction, error) {
	return _FlashLoaner.contract.Transact(opts, "Amount0CFlash", _amount0Out, _minProfit, _pair0, _pair1, _ts, _fee0, _fee1)
}

// Amount0CFlash is a paid mutator transaction binding the contract method 0xa32a0803.
//
// Solidity: function Amount0CFlash(uint256 _amount0Out, uint256 _minProfit, uint160 _pair0, uint160 _pair1, uint160 _ts, uint160 _fee0, uint160 _fee1) returns()
func (_FlashLoaner *FlashLoanerSession) Amount0CFlash(_amount0Out *big.Int, _minProfit *big.Int, _pair0 *big.Int, _pair1 *big.Int, _ts *big.Int, _fee0 *big.Int, _fee1 *big.Int) (*types.Transaction, error) {
	return _FlashLoaner.Contract.Amount0CFlash(&_FlashLoaner.TransactOpts, _amount0Out, _minProfit, _pair0, _pair1, _ts, _fee0, _fee1)
}

// Amount0CFlash is a paid mutator transaction binding the contract method 0xa32a0803.
//
// Solidity: function Amount0CFlash(uint256 _amount0Out, uint256 _minProfit, uint160 _pair0, uint160 _pair1, uint160 _ts, uint160 _fee0, uint160 _fee1) returns()
func (_FlashLoaner *FlashLoanerTransactorSession) Amount0CFlash(_amount0Out *big.Int, _minProfit *big.Int, _pair0 *big.Int, _pair1 *big.Int, _ts *big.Int, _fee0 *big.Int, _fee1 *big.Int) (*types.Transaction, error) {
	return _FlashLoaner.Contract.Amount0CFlash(&_FlashLoaner.TransactOpts, _amount0Out, _minProfit, _pair0, _pair1, _ts, _fee0, _fee1)
}

// Amount0Flash is a paid mutator transaction binding the contract method 0xcc34efa1.
//
// Solidity: function Amount0Flash(uint256 _amount0Out, uint256 _minProfit, address _pair0, address _pair1, uint160 _fee0, uint160 _fee1) returns()
func (_FlashLoaner *FlashLoanerTransactor) Amount0Flash(opts *bind.TransactOpts, _amount0Out *big.Int, _minProfit *big.Int, _pair0 common.Address, _pair1 common.Address, _fee0 *big.Int, _fee1 *big.Int) (*types.Transaction, error) {
	return _FlashLoaner.contract.Transact(opts, "Amount0Flash", _amount0Out, _minProfit, _pair0, _pair1, _fee0, _fee1)
}

// Amount0Flash is a paid mutator transaction binding the contract method 0xcc34efa1.
//
// Solidity: function Amount0Flash(uint256 _amount0Out, uint256 _minProfit, address _pair0, address _pair1, uint160 _fee0, uint160 _fee1) returns()
func (_FlashLoaner *FlashLoanerSession) Amount0Flash(_amount0Out *big.Int, _minProfit *big.Int, _pair0 common.Address, _pair1 common.Address, _fee0 *big.Int, _fee1 *big.Int) (*types.Transaction, error) {
	return _FlashLoaner.Contract.Amount0Flash(&_FlashLoaner.TransactOpts, _amount0Out, _minProfit, _pair0, _pair1, _fee0, _fee1)
}

// Amount0Flash is a paid mutator transaction binding the contract method 0xcc34efa1.
//
// Solidity: function Amount0Flash(uint256 _amount0Out, uint256 _minProfit, address _pair0, address _pair1, uint160 _fee0, uint160 _fee1) returns()
func (_FlashLoaner *FlashLoanerTransactorSession) Amount0Flash(_amount0Out *big.Int, _minProfit *big.Int, _pair0 common.Address, _pair1 common.Address, _fee0 *big.Int, _fee1 *big.Int) (*types.Transaction, error) {
	return _FlashLoaner.Contract.Amount0Flash(&_FlashLoaner.TransactOpts, _amount0Out, _minProfit, _pair0, _pair1, _fee0, _fee1)
}

// Amount1CFlash is a paid mutator transaction binding the contract method 0x4fec8116.
//
// Solidity: function Amount1CFlash(uint256 _amount1Out, uint256 _minProfit, uint160 _pair0, uint160 _pair1, uint160 _ts, uint160 _fee0, uint160 _fee1) returns()
func (_FlashLoaner *FlashLoanerTransactor) Amount1CFlash(opts *bind.TransactOpts, _amount1Out *big.Int, _minProfit *big.Int, _pair0 *big.Int, _pair1 *big.Int, _ts *big.Int, _fee0 *big.Int, _fee1 *big.Int) (*types.Transaction, error) {
	return _FlashLoaner.contract.Transact(opts, "Amount1CFlash", _amount1Out, _minProfit, _pair0, _pair1, _ts, _fee0, _fee1)
}

// Amount1CFlash is a paid mutator transaction binding the contract method 0x4fec8116.
//
// Solidity: function Amount1CFlash(uint256 _amount1Out, uint256 _minProfit, uint160 _pair0, uint160 _pair1, uint160 _ts, uint160 _fee0, uint160 _fee1) returns()
func (_FlashLoaner *FlashLoanerSession) Amount1CFlash(_amount1Out *big.Int, _minProfit *big.Int, _pair0 *big.Int, _pair1 *big.Int, _ts *big.Int, _fee0 *big.Int, _fee1 *big.Int) (*types.Transaction, error) {
	return _FlashLoaner.Contract.Amount1CFlash(&_FlashLoaner.TransactOpts, _amount1Out, _minProfit, _pair0, _pair1, _ts, _fee0, _fee1)
}

// Amount1CFlash is a paid mutator transaction binding the contract method 0x4fec8116.
//
// Solidity: function Amount1CFlash(uint256 _amount1Out, uint256 _minProfit, uint160 _pair0, uint160 _pair1, uint160 _ts, uint160 _fee0, uint160 _fee1) returns()
func (_FlashLoaner *FlashLoanerTransactorSession) Amount1CFlash(_amount1Out *big.Int, _minProfit *big.Int, _pair0 *big.Int, _pair1 *big.Int, _ts *big.Int, _fee0 *big.Int, _fee1 *big.Int) (*types.Transaction, error) {
	return _FlashLoaner.Contract.Amount1CFlash(&_FlashLoaner.TransactOpts, _amount1Out, _minProfit, _pair0, _pair1, _ts, _fee0, _fee1)
}

// Amount1Flash is a paid mutator transaction binding the contract method 0xef5affd7.
//
// Solidity: function Amount1Flash(uint256 _amount1Out, uint256 _minProfit, address _pair0, address _pair1, uint160 _fee0, uint160 _fee1) returns()
func (_FlashLoaner *FlashLoanerTransactor) Amount1Flash(opts *bind.TransactOpts, _amount1Out *big.Int, _minProfit *big.Int, _pair0 common.Address, _pair1 common.Address, _fee0 *big.Int, _fee1 *big.Int) (*types.Transaction, error) {
	return _FlashLoaner.contract.Transact(opts, "Amount1Flash", _amount1Out, _minProfit, _pair0, _pair1, _fee0, _fee1)
}

// Amount1Flash is a paid mutator transaction binding the contract method 0xef5affd7.
//
// Solidity: function Amount1Flash(uint256 _amount1Out, uint256 _minProfit, address _pair0, address _pair1, uint160 _fee0, uint160 _fee1) returns()
func (_FlashLoaner *FlashLoanerSession) Amount1Flash(_amount1Out *big.Int, _minProfit *big.Int, _pair0 common.Address, _pair1 common.Address, _fee0 *big.Int, _fee1 *big.Int) (*types.Transaction, error) {
	return _FlashLoaner.Contract.Amount1Flash(&_FlashLoaner.TransactOpts, _amount1Out, _minProfit, _pair0, _pair1, _fee0, _fee1)
}

// Amount1Flash is a paid mutator transaction binding the contract method 0xef5affd7.
//
// Solidity: function Amount1Flash(uint256 _amount1Out, uint256 _minProfit, address _pair0, address _pair1, uint160 _fee0, uint160 _fee1) returns()
func (_FlashLoaner *FlashLoanerTransactorSession) Amount1Flash(_amount1Out *big.Int, _minProfit *big.Int, _pair0 common.Address, _pair1 common.Address, _fee0 *big.Int, _fee1 *big.Int) (*types.Transaction, error) {
	return _FlashLoaner.Contract.Amount1Flash(&_FlashLoaner.TransactOpts, _amount1Out, _minProfit, _pair0, _pair1, _fee0, _fee1)
}

// UpdateArgs is a paid mutator transaction binding the contract method 0x4b118823.
//
// Solidity: function UpdateArgs(address _earnAddress) returns()
func (_FlashLoaner *FlashLoanerTransactor) UpdateArgs(opts *bind.TransactOpts, _earnAddress common.Address) (*types.Transaction, error) {
	return _FlashLoaner.contract.Transact(opts, "UpdateArgs", _earnAddress)
}

// UpdateArgs is a paid mutator transaction binding the contract method 0x4b118823.
//
// Solidity: function UpdateArgs(address _earnAddress) returns()
func (_FlashLoaner *FlashLoanerSession) UpdateArgs(_earnAddress common.Address) (*types.Transaction, error) {
	return _FlashLoaner.Contract.UpdateArgs(&_FlashLoaner.TransactOpts, _earnAddress)
}

// UpdateArgs is a paid mutator transaction binding the contract method 0x4b118823.
//
// Solidity: function UpdateArgs(address _earnAddress) returns()
func (_FlashLoaner *FlashLoanerTransactorSession) UpdateArgs(_earnAddress common.Address) (*types.Transaction, error) {
	return _FlashLoaner.Contract.UpdateArgs(&_FlashLoaner.TransactOpts, _earnAddress)
}

// UpdateOwner is a paid mutator transaction binding the contract method 0x84022644.
//
// Solidity: function UpdateOwner(address _owner) returns()
func (_FlashLoaner *FlashLoanerTransactor) UpdateOwner(opts *bind.TransactOpts, _owner common.Address) (*types.Transaction, error) {
	return _FlashLoaner.contract.Transact(opts, "UpdateOwner", _owner)
}

// UpdateOwner is a paid mutator transaction binding the contract method 0x84022644.
//
// Solidity: function UpdateOwner(address _owner) returns()
func (_FlashLoaner *FlashLoanerSession) UpdateOwner(_owner common.Address) (*types.Transaction, error) {
	return _FlashLoaner.Contract.UpdateOwner(&_FlashLoaner.TransactOpts, _owner)
}

// UpdateOwner is a paid mutator transaction binding the contract method 0x84022644.
//
// Solidity: function UpdateOwner(address _owner) returns()
func (_FlashLoaner *FlashLoanerTransactorSession) UpdateOwner(_owner common.Address) (*types.Transaction, error) {
	return _FlashLoaner.Contract.UpdateOwner(&_FlashLoaner.TransactOpts, _owner)
}

// UpdateSalt is a paid mutator transaction binding the contract method 0xa5af6d4b.
//
// Solidity: function UpdateSalt(uint160 _number) returns()
func (_FlashLoaner *FlashLoanerTransactor) UpdateSalt(opts *bind.TransactOpts, _number *big.Int) (*types.Transaction, error) {
	return _FlashLoaner.contract.Transact(opts, "UpdateSalt", _number)
}

// UpdateSalt is a paid mutator transaction binding the contract method 0xa5af6d4b.
//
// Solidity: function UpdateSalt(uint160 _number) returns()
func (_FlashLoaner *FlashLoanerSession) UpdateSalt(_number *big.Int) (*types.Transaction, error) {
	return _FlashLoaner.Contract.UpdateSalt(&_FlashLoaner.TransactOpts, _number)
}

// UpdateSalt is a paid mutator transaction binding the contract method 0xa5af6d4b.
//
// Solidity: function UpdateSalt(uint160 _number) returns()
func (_FlashLoaner *FlashLoanerTransactorSession) UpdateSalt(_number *big.Int) (*types.Transaction, error) {
	return _FlashLoaner.Contract.UpdateSalt(&_FlashLoaner.TransactOpts, _number)
}

// PancakeCall is a paid mutator transaction binding the contract method 0x84800812.
//
// Solidity: function pancakeCall(address _sender, uint256 _amount0, uint256 _amount1, bytes _data) returns()
func (_FlashLoaner *FlashLoanerTransactor) PancakeCall(opts *bind.TransactOpts, _sender common.Address, _amount0 *big.Int, _amount1 *big.Int, _data []byte) (*types.Transaction, error) {
	return _FlashLoaner.contract.Transact(opts, "pancakeCall", _sender, _amount0, _amount1, _data)
}

// PancakeCall is a paid mutator transaction binding the contract method 0x84800812.
//
// Solidity: function pancakeCall(address _sender, uint256 _amount0, uint256 _amount1, bytes _data) returns()
func (_FlashLoaner *FlashLoanerSession) PancakeCall(_sender common.Address, _amount0 *big.Int, _amount1 *big.Int, _data []byte) (*types.Transaction, error) {
	return _FlashLoaner.Contract.PancakeCall(&_FlashLoaner.TransactOpts, _sender, _amount0, _amount1, _data)
}

// PancakeCall is a paid mutator transaction binding the contract method 0x84800812.
//
// Solidity: function pancakeCall(address _sender, uint256 _amount0, uint256 _amount1, bytes _data) returns()
func (_FlashLoaner *FlashLoanerTransactorSession) PancakeCall(_sender common.Address, _amount0 *big.Int, _amount1 *big.Int, _data []byte) (*types.Transaction, error) {
	return _FlashLoaner.Contract.PancakeCall(&_FlashLoaner.TransactOpts, _sender, _amount0, _amount1, _data)
}
