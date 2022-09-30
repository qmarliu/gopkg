// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package FlashLoanerXtoY

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

// FlashLoanerXtoYABI is the input ABI used to generate the binding from.
// Deprecated: Use FlashLoanerXtoYMetaData.ABI instead.
var FlashLoanerXtoYABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_earnAddress\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_earnAddress\",\"type\":\"address\"}],\"name\":\"UpdateArgs\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"UpdateOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint160\",\"name\":\"_number\",\"type\":\"uint160\"}],\"name\":\"UpdateSalt\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount0Out\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_minProfit\",\"type\":\"uint256\"},{\"internalType\":\"uint160\",\"name\":\"_pair0\",\"type\":\"uint160\"},{\"internalType\":\"uint160\",\"name\":\"_pair1\",\"type\":\"uint160\"},{\"internalType\":\"uint160\",\"name\":\"_ts\",\"type\":\"uint160\"},{\"internalType\":\"uint160\",\"name\":\"_fee0\",\"type\":\"uint160\"},{\"internalType\":\"uint160\",\"name\":\"_fee1\",\"type\":\"uint160\"}],\"name\":\"Amount0CFlash\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount0Out\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_minProfit\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_pair0\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_pair1\",\"type\":\"address\"},{\"internalType\":\"uint160\",\"name\":\"_fee0\",\"type\":\"uint160\"},{\"internalType\":\"uint160\",\"name\":\"_fee1\",\"type\":\"uint160\"}],\"name\":\"Amount0Flash\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount1Out\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_minProfit\",\"type\":\"uint256\"},{\"internalType\":\"uint160\",\"name\":\"_pair0\",\"type\":\"uint160\"},{\"internalType\":\"uint160\",\"name\":\"_pair1\",\"type\":\"uint160\"},{\"internalType\":\"uint160\",\"name\":\"_ts\",\"type\":\"uint160\"},{\"internalType\":\"uint160\",\"name\":\"_fee0\",\"type\":\"uint160\"},{\"internalType\":\"uint160\",\"name\":\"_fee1\",\"type\":\"uint160\"}],\"name\":\"Amount1CFlash\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount1Out\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_minProfit\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_pair0\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_pair1\",\"type\":\"address\"},{\"internalType\":\"uint160\",\"name\":\"_fee0\",\"type\":\"uint160\"},{\"internalType\":\"uint160\",\"name\":\"_fee1\",\"type\":\"uint160\"}],\"name\":\"Amount1Flash\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount0Out\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_minProfit\",\"type\":\"uint256\"},{\"internalType\":\"uint160\",\"name\":\"_pair0\",\"type\":\"uint160\"},{\"internalType\":\"uint160\",\"name\":\"_pair1\",\"type\":\"uint160\"},{\"internalType\":\"uint160\",\"name\":\"_ts\",\"type\":\"uint160\"}],\"name\":\"testAmount0Flash\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"ts0r1In\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"ts1r1Out\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"s0r0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"s0r1\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"s1r0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"s1r1\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"win\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount1Out\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_minProfit\",\"type\":\"uint256\"},{\"internalType\":\"uint160\",\"name\":\"_pair0\",\"type\":\"uint160\"},{\"internalType\":\"uint160\",\"name\":\"_pair1\",\"type\":\"uint160\"},{\"internalType\":\"uint160\",\"name\":\"_ts\",\"type\":\"uint160\"}],\"name\":\"testAmount1Flash\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"ts0r0In\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"ts1r0Out\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"s0r0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"s0r1\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"s1r0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"s1r1\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"win\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_amount1\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"pancakeCall\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// FlashLoanerXtoY is an auto generated Go binding around an Ethereum contract.
type FlashLoanerXtoY struct {
	FlashLoanerXtoYCaller     // Read-only binding to the contract
	FlashLoanerXtoYTransactor // Write-only binding to the contract
	FlashLoanerXtoYFilterer   // Log filterer for contract events
}

// FlashLoanerXtoYCaller is an auto generated read-only Go binding around an Ethereum contract.
type FlashLoanerXtoYCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FlashLoanerXtoYTransactor is an auto generated write-only Go binding around an Ethereum contract.
type FlashLoanerXtoYTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FlashLoanerXtoYFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type FlashLoanerXtoYFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FlashLoanerXtoYSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type FlashLoanerXtoYSession struct {
	Contract     *FlashLoanerXtoY  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// FlashLoanerXtoYCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type FlashLoanerXtoYCallerSession struct {
	Contract *FlashLoanerXtoYCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// FlashLoanerXtoYTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type FlashLoanerXtoYTransactorSession struct {
	Contract     *FlashLoanerXtoYTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// FlashLoanerXtoYRaw is an auto generated low-level Go binding around an Ethereum contract.
type FlashLoanerXtoYRaw struct {
	Contract *FlashLoanerXtoY // Generic contract binding to access the raw methods on
}

// FlashLoanerXtoYCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type FlashLoanerXtoYCallerRaw struct {
	Contract *FlashLoanerXtoYCaller // Generic read-only contract binding to access the raw methods on
}

// FlashLoanerXtoYTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type FlashLoanerXtoYTransactorRaw struct {
	Contract *FlashLoanerXtoYTransactor // Generic write-only contract binding to access the raw methods on
}

// NewFlashLoanerXtoY creates a new instance of FlashLoanerXtoY, bound to a specific deployed contract.
func NewFlashLoanerXtoY(address common.Address, backend bind.ContractBackend) (*FlashLoanerXtoY, error) {
	contract, err := bindFlashLoanerXtoY(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &FlashLoanerXtoY{FlashLoanerXtoYCaller: FlashLoanerXtoYCaller{contract: contract}, FlashLoanerXtoYTransactor: FlashLoanerXtoYTransactor{contract: contract}, FlashLoanerXtoYFilterer: FlashLoanerXtoYFilterer{contract: contract}}, nil
}

// NewFlashLoanerXtoYCaller creates a new read-only instance of FlashLoanerXtoY, bound to a specific deployed contract.
func NewFlashLoanerXtoYCaller(address common.Address, caller bind.ContractCaller) (*FlashLoanerXtoYCaller, error) {
	contract, err := bindFlashLoanerXtoY(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &FlashLoanerXtoYCaller{contract: contract}, nil
}

// NewFlashLoanerXtoYTransactor creates a new write-only instance of FlashLoanerXtoY, bound to a specific deployed contract.
func NewFlashLoanerXtoYTransactor(address common.Address, transactor bind.ContractTransactor) (*FlashLoanerXtoYTransactor, error) {
	contract, err := bindFlashLoanerXtoY(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &FlashLoanerXtoYTransactor{contract: contract}, nil
}

// NewFlashLoanerXtoYFilterer creates a new log filterer instance of FlashLoanerXtoY, bound to a specific deployed contract.
func NewFlashLoanerXtoYFilterer(address common.Address, filterer bind.ContractFilterer) (*FlashLoanerXtoYFilterer, error) {
	contract, err := bindFlashLoanerXtoY(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &FlashLoanerXtoYFilterer{contract: contract}, nil
}

// bindFlashLoanerXtoY binds a generic wrapper to an already deployed contract.
func bindFlashLoanerXtoY(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(FlashLoanerXtoYABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FlashLoanerXtoY *FlashLoanerXtoYRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FlashLoanerXtoY.Contract.FlashLoanerXtoYCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FlashLoanerXtoY *FlashLoanerXtoYRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FlashLoanerXtoY.Contract.FlashLoanerXtoYTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FlashLoanerXtoY *FlashLoanerXtoYRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FlashLoanerXtoY.Contract.FlashLoanerXtoYTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FlashLoanerXtoY *FlashLoanerXtoYCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FlashLoanerXtoY.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FlashLoanerXtoY *FlashLoanerXtoYTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FlashLoanerXtoY.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FlashLoanerXtoY *FlashLoanerXtoYTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FlashLoanerXtoY.Contract.contract.Transact(opts, method, params...)
}

// TestAmount0Flash is a free data retrieval call binding the contract method 0xad83a740.
//
// Solidity: function testAmount0Flash(uint256 _amount0Out, uint256 _minProfit, uint160 _pair0, uint160 _pair1, uint160 _ts) view returns(uint256 ts0r1In, uint256 ts1r1Out, uint256 s0r0, uint256 s0r1, uint256 s1r0, uint256 s1r1, bool win)
func (_FlashLoanerXtoY *FlashLoanerXtoYCaller) TestAmount0Flash(opts *bind.CallOpts, _amount0Out *big.Int, _minProfit *big.Int, _pair0 *big.Int, _pair1 *big.Int, _ts *big.Int) (struct {
	Ts0r1In  *big.Int
	Ts1r1Out *big.Int
	S0r0     *big.Int
	S0r1     *big.Int
	S1r0     *big.Int
	S1r1     *big.Int
	Win      bool
}, error) {
	var out []interface{}
	err := _FlashLoanerXtoY.contract.Call(opts, &out, "testAmount0Flash", _amount0Out, _minProfit, _pair0, _pair1, _ts)

	outstruct := new(struct {
		Ts0r1In  *big.Int
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

	outstruct.Ts0r1In = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Ts1r1Out = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.S0r0 = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.S0r1 = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.S1r0 = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.S1r1 = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	outstruct.Win = *abi.ConvertType(out[6], new(bool)).(*bool)

	return *outstruct, err

}

// TestAmount0Flash is a free data retrieval call binding the contract method 0xad83a740.
//
// Solidity: function testAmount0Flash(uint256 _amount0Out, uint256 _minProfit, uint160 _pair0, uint160 _pair1, uint160 _ts) view returns(uint256 ts0r1In, uint256 ts1r1Out, uint256 s0r0, uint256 s0r1, uint256 s1r0, uint256 s1r1, bool win)
func (_FlashLoanerXtoY *FlashLoanerXtoYSession) TestAmount0Flash(_amount0Out *big.Int, _minProfit *big.Int, _pair0 *big.Int, _pair1 *big.Int, _ts *big.Int) (struct {
	Ts0r1In  *big.Int
	Ts1r1Out *big.Int
	S0r0     *big.Int
	S0r1     *big.Int
	S1r0     *big.Int
	S1r1     *big.Int
	Win      bool
}, error) {
	return _FlashLoanerXtoY.Contract.TestAmount0Flash(&_FlashLoanerXtoY.CallOpts, _amount0Out, _minProfit, _pair0, _pair1, _ts)
}

// TestAmount0Flash is a free data retrieval call binding the contract method 0xad83a740.
//
// Solidity: function testAmount0Flash(uint256 _amount0Out, uint256 _minProfit, uint160 _pair0, uint160 _pair1, uint160 _ts) view returns(uint256 ts0r1In, uint256 ts1r1Out, uint256 s0r0, uint256 s0r1, uint256 s1r0, uint256 s1r1, bool win)
func (_FlashLoanerXtoY *FlashLoanerXtoYCallerSession) TestAmount0Flash(_amount0Out *big.Int, _minProfit *big.Int, _pair0 *big.Int, _pair1 *big.Int, _ts *big.Int) (struct {
	Ts0r1In  *big.Int
	Ts1r1Out *big.Int
	S0r0     *big.Int
	S0r1     *big.Int
	S1r0     *big.Int
	S1r1     *big.Int
	Win      bool
}, error) {
	return _FlashLoanerXtoY.Contract.TestAmount0Flash(&_FlashLoanerXtoY.CallOpts, _amount0Out, _minProfit, _pair0, _pair1, _ts)
}

// TestAmount1Flash is a free data retrieval call binding the contract method 0xac41c2d0.
//
// Solidity: function testAmount1Flash(uint256 _amount1Out, uint256 _minProfit, uint160 _pair0, uint160 _pair1, uint160 _ts) view returns(uint256 ts0r0In, uint256 ts1r0Out, uint256 s0r0, uint256 s0r1, uint256 s1r0, uint256 s1r1, bool win)
func (_FlashLoanerXtoY *FlashLoanerXtoYCaller) TestAmount1Flash(opts *bind.CallOpts, _amount1Out *big.Int, _minProfit *big.Int, _pair0 *big.Int, _pair1 *big.Int, _ts *big.Int) (struct {
	Ts0r0In  *big.Int
	Ts1r0Out *big.Int
	S0r0     *big.Int
	S0r1     *big.Int
	S1r0     *big.Int
	S1r1     *big.Int
	Win      bool
}, error) {
	var out []interface{}
	err := _FlashLoanerXtoY.contract.Call(opts, &out, "testAmount1Flash", _amount1Out, _minProfit, _pair0, _pair1, _ts)

	outstruct := new(struct {
		Ts0r0In  *big.Int
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

	outstruct.Ts0r0In = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Ts1r0Out = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.S0r0 = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.S0r1 = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.S1r0 = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.S1r1 = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	outstruct.Win = *abi.ConvertType(out[6], new(bool)).(*bool)

	return *outstruct, err

}

// TestAmount1Flash is a free data retrieval call binding the contract method 0xac41c2d0.
//
// Solidity: function testAmount1Flash(uint256 _amount1Out, uint256 _minProfit, uint160 _pair0, uint160 _pair1, uint160 _ts) view returns(uint256 ts0r0In, uint256 ts1r0Out, uint256 s0r0, uint256 s0r1, uint256 s1r0, uint256 s1r1, bool win)
func (_FlashLoanerXtoY *FlashLoanerXtoYSession) TestAmount1Flash(_amount1Out *big.Int, _minProfit *big.Int, _pair0 *big.Int, _pair1 *big.Int, _ts *big.Int) (struct {
	Ts0r0In  *big.Int
	Ts1r0Out *big.Int
	S0r0     *big.Int
	S0r1     *big.Int
	S1r0     *big.Int
	S1r1     *big.Int
	Win      bool
}, error) {
	return _FlashLoanerXtoY.Contract.TestAmount1Flash(&_FlashLoanerXtoY.CallOpts, _amount1Out, _minProfit, _pair0, _pair1, _ts)
}

// TestAmount1Flash is a free data retrieval call binding the contract method 0xac41c2d0.
//
// Solidity: function testAmount1Flash(uint256 _amount1Out, uint256 _minProfit, uint160 _pair0, uint160 _pair1, uint160 _ts) view returns(uint256 ts0r0In, uint256 ts1r0Out, uint256 s0r0, uint256 s0r1, uint256 s1r0, uint256 s1r1, bool win)
func (_FlashLoanerXtoY *FlashLoanerXtoYCallerSession) TestAmount1Flash(_amount1Out *big.Int, _minProfit *big.Int, _pair0 *big.Int, _pair1 *big.Int, _ts *big.Int) (struct {
	Ts0r0In  *big.Int
	Ts1r0Out *big.Int
	S0r0     *big.Int
	S0r1     *big.Int
	S1r0     *big.Int
	S1r1     *big.Int
	Win      bool
}, error) {
	return _FlashLoanerXtoY.Contract.TestAmount1Flash(&_FlashLoanerXtoY.CallOpts, _amount1Out, _minProfit, _pair0, _pair1, _ts)
}

// Amount0CFlash is a paid mutator transaction binding the contract method 0xa32a0803.
//
// Solidity: function Amount0CFlash(uint256 _amount0Out, uint256 _minProfit, uint160 _pair0, uint160 _pair1, uint160 _ts, uint160 _fee0, uint160 _fee1) returns()
func (_FlashLoanerXtoY *FlashLoanerXtoYTransactor) Amount0CFlash(opts *bind.TransactOpts, _amount0Out *big.Int, _minProfit *big.Int, _pair0 *big.Int, _pair1 *big.Int, _ts *big.Int, _fee0 *big.Int, _fee1 *big.Int) (*types.Transaction, error) {
	return _FlashLoanerXtoY.contract.Transact(opts, "Amount0CFlash", _amount0Out, _minProfit, _pair0, _pair1, _ts, _fee0, _fee1)
}

// Amount0CFlash is a paid mutator transaction binding the contract method 0xa32a0803.
//
// Solidity: function Amount0CFlash(uint256 _amount0Out, uint256 _minProfit, uint160 _pair0, uint160 _pair1, uint160 _ts, uint160 _fee0, uint160 _fee1) returns()
func (_FlashLoanerXtoY *FlashLoanerXtoYSession) Amount0CFlash(_amount0Out *big.Int, _minProfit *big.Int, _pair0 *big.Int, _pair1 *big.Int, _ts *big.Int, _fee0 *big.Int, _fee1 *big.Int) (*types.Transaction, error) {
	return _FlashLoanerXtoY.Contract.Amount0CFlash(&_FlashLoanerXtoY.TransactOpts, _amount0Out, _minProfit, _pair0, _pair1, _ts, _fee0, _fee1)
}

// Amount0CFlash is a paid mutator transaction binding the contract method 0xa32a0803.
//
// Solidity: function Amount0CFlash(uint256 _amount0Out, uint256 _minProfit, uint160 _pair0, uint160 _pair1, uint160 _ts, uint160 _fee0, uint160 _fee1) returns()
func (_FlashLoanerXtoY *FlashLoanerXtoYTransactorSession) Amount0CFlash(_amount0Out *big.Int, _minProfit *big.Int, _pair0 *big.Int, _pair1 *big.Int, _ts *big.Int, _fee0 *big.Int, _fee1 *big.Int) (*types.Transaction, error) {
	return _FlashLoanerXtoY.Contract.Amount0CFlash(&_FlashLoanerXtoY.TransactOpts, _amount0Out, _minProfit, _pair0, _pair1, _ts, _fee0, _fee1)
}

// Amount0Flash is a paid mutator transaction binding the contract method 0xcc34efa1.
//
// Solidity: function Amount0Flash(uint256 _amount0Out, uint256 _minProfit, address _pair0, address _pair1, uint160 _fee0, uint160 _fee1) returns()
func (_FlashLoanerXtoY *FlashLoanerXtoYTransactor) Amount0Flash(opts *bind.TransactOpts, _amount0Out *big.Int, _minProfit *big.Int, _pair0 common.Address, _pair1 common.Address, _fee0 *big.Int, _fee1 *big.Int) (*types.Transaction, error) {
	return _FlashLoanerXtoY.contract.Transact(opts, "Amount0Flash", _amount0Out, _minProfit, _pair0, _pair1, _fee0, _fee1)
}

// Amount0Flash is a paid mutator transaction binding the contract method 0xcc34efa1.
//
// Solidity: function Amount0Flash(uint256 _amount0Out, uint256 _minProfit, address _pair0, address _pair1, uint160 _fee0, uint160 _fee1) returns()
func (_FlashLoanerXtoY *FlashLoanerXtoYSession) Amount0Flash(_amount0Out *big.Int, _minProfit *big.Int, _pair0 common.Address, _pair1 common.Address, _fee0 *big.Int, _fee1 *big.Int) (*types.Transaction, error) {
	return _FlashLoanerXtoY.Contract.Amount0Flash(&_FlashLoanerXtoY.TransactOpts, _amount0Out, _minProfit, _pair0, _pair1, _fee0, _fee1)
}

// Amount0Flash is a paid mutator transaction binding the contract method 0xcc34efa1.
//
// Solidity: function Amount0Flash(uint256 _amount0Out, uint256 _minProfit, address _pair0, address _pair1, uint160 _fee0, uint160 _fee1) returns()
func (_FlashLoanerXtoY *FlashLoanerXtoYTransactorSession) Amount0Flash(_amount0Out *big.Int, _minProfit *big.Int, _pair0 common.Address, _pair1 common.Address, _fee0 *big.Int, _fee1 *big.Int) (*types.Transaction, error) {
	return _FlashLoanerXtoY.Contract.Amount0Flash(&_FlashLoanerXtoY.TransactOpts, _amount0Out, _minProfit, _pair0, _pair1, _fee0, _fee1)
}

// Amount1CFlash is a paid mutator transaction binding the contract method 0x4fec8116.
//
// Solidity: function Amount1CFlash(uint256 _amount1Out, uint256 _minProfit, uint160 _pair0, uint160 _pair1, uint160 _ts, uint160 _fee0, uint160 _fee1) returns()
func (_FlashLoanerXtoY *FlashLoanerXtoYTransactor) Amount1CFlash(opts *bind.TransactOpts, _amount1Out *big.Int, _minProfit *big.Int, _pair0 *big.Int, _pair1 *big.Int, _ts *big.Int, _fee0 *big.Int, _fee1 *big.Int) (*types.Transaction, error) {
	return _FlashLoanerXtoY.contract.Transact(opts, "Amount1CFlash", _amount1Out, _minProfit, _pair0, _pair1, _ts, _fee0, _fee1)
}

// Amount1CFlash is a paid mutator transaction binding the contract method 0x4fec8116.
//
// Solidity: function Amount1CFlash(uint256 _amount1Out, uint256 _minProfit, uint160 _pair0, uint160 _pair1, uint160 _ts, uint160 _fee0, uint160 _fee1) returns()
func (_FlashLoanerXtoY *FlashLoanerXtoYSession) Amount1CFlash(_amount1Out *big.Int, _minProfit *big.Int, _pair0 *big.Int, _pair1 *big.Int, _ts *big.Int, _fee0 *big.Int, _fee1 *big.Int) (*types.Transaction, error) {
	return _FlashLoanerXtoY.Contract.Amount1CFlash(&_FlashLoanerXtoY.TransactOpts, _amount1Out, _minProfit, _pair0, _pair1, _ts, _fee0, _fee1)
}

// Amount1CFlash is a paid mutator transaction binding the contract method 0x4fec8116.
//
// Solidity: function Amount1CFlash(uint256 _amount1Out, uint256 _minProfit, uint160 _pair0, uint160 _pair1, uint160 _ts, uint160 _fee0, uint160 _fee1) returns()
func (_FlashLoanerXtoY *FlashLoanerXtoYTransactorSession) Amount1CFlash(_amount1Out *big.Int, _minProfit *big.Int, _pair0 *big.Int, _pair1 *big.Int, _ts *big.Int, _fee0 *big.Int, _fee1 *big.Int) (*types.Transaction, error) {
	return _FlashLoanerXtoY.Contract.Amount1CFlash(&_FlashLoanerXtoY.TransactOpts, _amount1Out, _minProfit, _pair0, _pair1, _ts, _fee0, _fee1)
}

// Amount1Flash is a paid mutator transaction binding the contract method 0xef5affd7.
//
// Solidity: function Amount1Flash(uint256 _amount1Out, uint256 _minProfit, address _pair0, address _pair1, uint160 _fee0, uint160 _fee1) returns()
func (_FlashLoanerXtoY *FlashLoanerXtoYTransactor) Amount1Flash(opts *bind.TransactOpts, _amount1Out *big.Int, _minProfit *big.Int, _pair0 common.Address, _pair1 common.Address, _fee0 *big.Int, _fee1 *big.Int) (*types.Transaction, error) {
	return _FlashLoanerXtoY.contract.Transact(opts, "Amount1Flash", _amount1Out, _minProfit, _pair0, _pair1, _fee0, _fee1)
}

// Amount1Flash is a paid mutator transaction binding the contract method 0xef5affd7.
//
// Solidity: function Amount1Flash(uint256 _amount1Out, uint256 _minProfit, address _pair0, address _pair1, uint160 _fee0, uint160 _fee1) returns()
func (_FlashLoanerXtoY *FlashLoanerXtoYSession) Amount1Flash(_amount1Out *big.Int, _minProfit *big.Int, _pair0 common.Address, _pair1 common.Address, _fee0 *big.Int, _fee1 *big.Int) (*types.Transaction, error) {
	return _FlashLoanerXtoY.Contract.Amount1Flash(&_FlashLoanerXtoY.TransactOpts, _amount1Out, _minProfit, _pair0, _pair1, _fee0, _fee1)
}

// Amount1Flash is a paid mutator transaction binding the contract method 0xef5affd7.
//
// Solidity: function Amount1Flash(uint256 _amount1Out, uint256 _minProfit, address _pair0, address _pair1, uint160 _fee0, uint160 _fee1) returns()
func (_FlashLoanerXtoY *FlashLoanerXtoYTransactorSession) Amount1Flash(_amount1Out *big.Int, _minProfit *big.Int, _pair0 common.Address, _pair1 common.Address, _fee0 *big.Int, _fee1 *big.Int) (*types.Transaction, error) {
	return _FlashLoanerXtoY.Contract.Amount1Flash(&_FlashLoanerXtoY.TransactOpts, _amount1Out, _minProfit, _pair0, _pair1, _fee0, _fee1)
}

// UpdateArgs is a paid mutator transaction binding the contract method 0x4b118823.
//
// Solidity: function UpdateArgs(address _earnAddress) returns()
func (_FlashLoanerXtoY *FlashLoanerXtoYTransactor) UpdateArgs(opts *bind.TransactOpts, _earnAddress common.Address) (*types.Transaction, error) {
	return _FlashLoanerXtoY.contract.Transact(opts, "UpdateArgs", _earnAddress)
}

// UpdateArgs is a paid mutator transaction binding the contract method 0x4b118823.
//
// Solidity: function UpdateArgs(address _earnAddress) returns()
func (_FlashLoanerXtoY *FlashLoanerXtoYSession) UpdateArgs(_earnAddress common.Address) (*types.Transaction, error) {
	return _FlashLoanerXtoY.Contract.UpdateArgs(&_FlashLoanerXtoY.TransactOpts, _earnAddress)
}

// UpdateArgs is a paid mutator transaction binding the contract method 0x4b118823.
//
// Solidity: function UpdateArgs(address _earnAddress) returns()
func (_FlashLoanerXtoY *FlashLoanerXtoYTransactorSession) UpdateArgs(_earnAddress common.Address) (*types.Transaction, error) {
	return _FlashLoanerXtoY.Contract.UpdateArgs(&_FlashLoanerXtoY.TransactOpts, _earnAddress)
}

// UpdateOwner is a paid mutator transaction binding the contract method 0x84022644.
//
// Solidity: function UpdateOwner(address _owner) returns()
func (_FlashLoanerXtoY *FlashLoanerXtoYTransactor) UpdateOwner(opts *bind.TransactOpts, _owner common.Address) (*types.Transaction, error) {
	return _FlashLoanerXtoY.contract.Transact(opts, "UpdateOwner", _owner)
}

// UpdateOwner is a paid mutator transaction binding the contract method 0x84022644.
//
// Solidity: function UpdateOwner(address _owner) returns()
func (_FlashLoanerXtoY *FlashLoanerXtoYSession) UpdateOwner(_owner common.Address) (*types.Transaction, error) {
	return _FlashLoanerXtoY.Contract.UpdateOwner(&_FlashLoanerXtoY.TransactOpts, _owner)
}

// UpdateOwner is a paid mutator transaction binding the contract method 0x84022644.
//
// Solidity: function UpdateOwner(address _owner) returns()
func (_FlashLoanerXtoY *FlashLoanerXtoYTransactorSession) UpdateOwner(_owner common.Address) (*types.Transaction, error) {
	return _FlashLoanerXtoY.Contract.UpdateOwner(&_FlashLoanerXtoY.TransactOpts, _owner)
}

// UpdateSalt is a paid mutator transaction binding the contract method 0xa5af6d4b.
//
// Solidity: function UpdateSalt(uint160 _number) returns()
func (_FlashLoanerXtoY *FlashLoanerXtoYTransactor) UpdateSalt(opts *bind.TransactOpts, _number *big.Int) (*types.Transaction, error) {
	return _FlashLoanerXtoY.contract.Transact(opts, "UpdateSalt", _number)
}

// UpdateSalt is a paid mutator transaction binding the contract method 0xa5af6d4b.
//
// Solidity: function UpdateSalt(uint160 _number) returns()
func (_FlashLoanerXtoY *FlashLoanerXtoYSession) UpdateSalt(_number *big.Int) (*types.Transaction, error) {
	return _FlashLoanerXtoY.Contract.UpdateSalt(&_FlashLoanerXtoY.TransactOpts, _number)
}

// UpdateSalt is a paid mutator transaction binding the contract method 0xa5af6d4b.
//
// Solidity: function UpdateSalt(uint160 _number) returns()
func (_FlashLoanerXtoY *FlashLoanerXtoYTransactorSession) UpdateSalt(_number *big.Int) (*types.Transaction, error) {
	return _FlashLoanerXtoY.Contract.UpdateSalt(&_FlashLoanerXtoY.TransactOpts, _number)
}

// PancakeCall is a paid mutator transaction binding the contract method 0x84800812.
//
// Solidity: function pancakeCall(address _sender, uint256 _amount0, uint256 _amount1, bytes _data) returns()
func (_FlashLoanerXtoY *FlashLoanerXtoYTransactor) PancakeCall(opts *bind.TransactOpts, _sender common.Address, _amount0 *big.Int, _amount1 *big.Int, _data []byte) (*types.Transaction, error) {
	return _FlashLoanerXtoY.contract.Transact(opts, "pancakeCall", _sender, _amount0, _amount1, _data)
}

// PancakeCall is a paid mutator transaction binding the contract method 0x84800812.
//
// Solidity: function pancakeCall(address _sender, uint256 _amount0, uint256 _amount1, bytes _data) returns()
func (_FlashLoanerXtoY *FlashLoanerXtoYSession) PancakeCall(_sender common.Address, _amount0 *big.Int, _amount1 *big.Int, _data []byte) (*types.Transaction, error) {
	return _FlashLoanerXtoY.Contract.PancakeCall(&_FlashLoanerXtoY.TransactOpts, _sender, _amount0, _amount1, _data)
}

// PancakeCall is a paid mutator transaction binding the contract method 0x84800812.
//
// Solidity: function pancakeCall(address _sender, uint256 _amount0, uint256 _amount1, bytes _data) returns()
func (_FlashLoanerXtoY *FlashLoanerXtoYTransactorSession) PancakeCall(_sender common.Address, _amount0 *big.Int, _amount1 *big.Int, _data []byte) (*types.Transaction, error) {
	return _FlashLoanerXtoY.Contract.PancakeCall(&_FlashLoanerXtoY.TransactOpts, _sender, _amount0, _amount1, _data)
}
