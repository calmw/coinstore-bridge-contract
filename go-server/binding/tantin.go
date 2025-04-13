// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package binding

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

// TantinMetaData contains all meta data concerning the Tantin contract.
var TantinMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"AccessControlBadConfirmation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"neededRole\",\"type\":\"bytes32\"}],\"name\":\"AccessControlUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"AddressEmptyCode\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"AddressInsufficientBalance\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ECDSAInvalidSignature\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"length\",\"type\":\"uint256\"}],\"name\":\"ECDSAInvalidSignatureLength\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"ECDSAInvalidSignatureS\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"assetsType\",\"type\":\"uint8\"}],\"name\":\"ErrAssetsType\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedInnerCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"SafeERC20FailedOperation\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"AddBlacklist\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"depositer\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"depositNonce\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"destinationChainId\",\"type\":\"uint256\"}],\"name\":\"DepositEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"depositer\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"depositNonce\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"destinationChainId\",\"type\":\"uint256\"}],\"name\":\"DepositNftEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"depositer\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"originDepositNonce\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"originChainId\",\"type\":\"uint256\"}],\"name\":\"ExecuteNftEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"RemoveBlacklist\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"resourceID\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"enumITantinBridge.AssetsType\",\"name\":\"assetsType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"burnable\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"mintable\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"pause\",\"type\":\"bool\"}],\"name\":\"SetTokenEvent\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"BRIDGE_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"Bridge\",\"outputs\":[{\"internalType\":\"contractIBridge\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"adminAddBlacklist\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"adminRemoveBlacklist\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"feeAddress_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"bridgeAddress_\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"signature_\",\"type\":\"bytes\"}],\"name\":\"adminSetEnv\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"adminWithdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"blacklist\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"destinationChainId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"resourceId\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"depositRecord\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"destinationChainId\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"resourceId\",\"type\":\"bytes32\"}],\"name\":\"getFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"localNonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"callerConfirmation\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sigNonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// TantinABI is the input ABI used to generate the binding from.
// Deprecated: Use TantinMetaData.ABI instead.
var TantinABI = TantinMetaData.ABI

// Tantin is an auto generated Go binding around an Ethereum contract.
type Tantin struct {
	TantinCaller     // Read-only binding to the contract
	TantinTransactor // Write-only binding to the contract
	TantinFilterer   // Log filterer for contract events
}

// TantinCaller is an auto generated read-only Go binding around an Ethereum contract.
type TantinCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TantinTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TantinTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TantinFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TantinFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TantinSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TantinSession struct {
	Contract     *Tantin           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TantinCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TantinCallerSession struct {
	Contract *TantinCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// TantinTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TantinTransactorSession struct {
	Contract     *TantinTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TantinRaw is an auto generated low-level Go binding around an Ethereum contract.
type TantinRaw struct {
	Contract *Tantin // Generic contract binding to access the raw methods on
}

// TantinCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TantinCallerRaw struct {
	Contract *TantinCaller // Generic read-only contract binding to access the raw methods on
}

// TantinTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TantinTransactorRaw struct {
	Contract *TantinTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTantin creates a new instance of Tantin, bound to a specific deployed contract.
func NewTantin(address common.Address, backend bind.ContractBackend) (*Tantin, error) {
	contract, err := bindTantin(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Tantin{TantinCaller: TantinCaller{contract: contract}, TantinTransactor: TantinTransactor{contract: contract}, TantinFilterer: TantinFilterer{contract: contract}}, nil
}

// NewTantinCaller creates a new read-only instance of Tantin, bound to a specific deployed contract.
func NewTantinCaller(address common.Address, caller bind.ContractCaller) (*TantinCaller, error) {
	contract, err := bindTantin(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TantinCaller{contract: contract}, nil
}

// NewTantinTransactor creates a new write-only instance of Tantin, bound to a specific deployed contract.
func NewTantinTransactor(address common.Address, transactor bind.ContractTransactor) (*TantinTransactor, error) {
	contract, err := bindTantin(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TantinTransactor{contract: contract}, nil
}

// NewTantinFilterer creates a new log filterer instance of Tantin, bound to a specific deployed contract.
func NewTantinFilterer(address common.Address, filterer bind.ContractFilterer) (*TantinFilterer, error) {
	contract, err := bindTantin(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TantinFilterer{contract: contract}, nil
}

// bindTantin binds a generic wrapper to an already deployed contract.
func bindTantin(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := TantinMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Tantin *TantinRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Tantin.Contract.TantinCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Tantin *TantinRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Tantin.Contract.TantinTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Tantin *TantinRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Tantin.Contract.TantinTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Tantin *TantinCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Tantin.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Tantin *TantinTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Tantin.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Tantin *TantinTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Tantin.Contract.contract.Transact(opts, method, params...)
}

// ADMINROLE is a free data retrieval call binding the contract method 0x75b238fc.
//
// Solidity: function ADMIN_ROLE() view returns(bytes32)
func (_Tantin *TantinCaller) ADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Tantin.contract.Call(opts, &out, "ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ADMINROLE is a free data retrieval call binding the contract method 0x75b238fc.
//
// Solidity: function ADMIN_ROLE() view returns(bytes32)
func (_Tantin *TantinSession) ADMINROLE() ([32]byte, error) {
	return _Tantin.Contract.ADMINROLE(&_Tantin.CallOpts)
}

// ADMINROLE is a free data retrieval call binding the contract method 0x75b238fc.
//
// Solidity: function ADMIN_ROLE() view returns(bytes32)
func (_Tantin *TantinCallerSession) ADMINROLE() ([32]byte, error) {
	return _Tantin.Contract.ADMINROLE(&_Tantin.CallOpts)
}

// BRIDGEROLE is a free data retrieval call binding the contract method 0xb5bfddea.
//
// Solidity: function BRIDGE_ROLE() view returns(bytes32)
func (_Tantin *TantinCaller) BRIDGEROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Tantin.contract.Call(opts, &out, "BRIDGE_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// BRIDGEROLE is a free data retrieval call binding the contract method 0xb5bfddea.
//
// Solidity: function BRIDGE_ROLE() view returns(bytes32)
func (_Tantin *TantinSession) BRIDGEROLE() ([32]byte, error) {
	return _Tantin.Contract.BRIDGEROLE(&_Tantin.CallOpts)
}

// BRIDGEROLE is a free data retrieval call binding the contract method 0xb5bfddea.
//
// Solidity: function BRIDGE_ROLE() view returns(bytes32)
func (_Tantin *TantinCallerSession) BRIDGEROLE() ([32]byte, error) {
	return _Tantin.Contract.BRIDGEROLE(&_Tantin.CallOpts)
}

// Bridge is a free data retrieval call binding the contract method 0x5fa2ef10.
//
// Solidity: function Bridge() view returns(address)
func (_Tantin *TantinCaller) Bridge(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Tantin.contract.Call(opts, &out, "Bridge")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Bridge is a free data retrieval call binding the contract method 0x5fa2ef10.
//
// Solidity: function Bridge() view returns(address)
func (_Tantin *TantinSession) Bridge() (common.Address, error) {
	return _Tantin.Contract.Bridge(&_Tantin.CallOpts)
}

// Bridge is a free data retrieval call binding the contract method 0x5fa2ef10.
//
// Solidity: function Bridge() view returns(address)
func (_Tantin *TantinCallerSession) Bridge() (common.Address, error) {
	return _Tantin.Contract.Bridge(&_Tantin.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Tantin *TantinCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Tantin.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Tantin *TantinSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Tantin.Contract.DEFAULTADMINROLE(&_Tantin.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Tantin *TantinCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Tantin.Contract.DEFAULTADMINROLE(&_Tantin.CallOpts)
}

// Blacklist is a free data retrieval call binding the contract method 0xf9f92be4.
//
// Solidity: function blacklist(address ) view returns(bool)
func (_Tantin *TantinCaller) Blacklist(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Tantin.contract.Call(opts, &out, "blacklist", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Blacklist is a free data retrieval call binding the contract method 0xf9f92be4.
//
// Solidity: function blacklist(address ) view returns(bool)
func (_Tantin *TantinSession) Blacklist(arg0 common.Address) (bool, error) {
	return _Tantin.Contract.Blacklist(&_Tantin.CallOpts, arg0)
}

// Blacklist is a free data retrieval call binding the contract method 0xf9f92be4.
//
// Solidity: function blacklist(address ) view returns(bool)
func (_Tantin *TantinCallerSession) Blacklist(arg0 common.Address) (bool, error) {
	return _Tantin.Contract.Blacklist(&_Tantin.CallOpts, arg0)
}

// DepositRecord is a free data retrieval call binding the contract method 0x2f26cb6e.
//
// Solidity: function depositRecord(address , uint256 ) view returns(address tokenAddress, address sender, address recipient, uint256 amount, uint256 fee, uint256 destinationChainId)
func (_Tantin *TantinCaller) DepositRecord(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (struct {
	TokenAddress       common.Address
	Sender             common.Address
	Recipient          common.Address
	Amount             *big.Int
	Fee                *big.Int
	DestinationChainId *big.Int
}, error) {
	var out []interface{}
	err := _Tantin.contract.Call(opts, &out, "depositRecord", arg0, arg1)

	outstruct := new(struct {
		TokenAddress       common.Address
		Sender             common.Address
		Recipient          common.Address
		Amount             *big.Int
		Fee                *big.Int
		DestinationChainId *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.TokenAddress = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Sender = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.Recipient = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)
	outstruct.Amount = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.Fee = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.DestinationChainId = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// DepositRecord is a free data retrieval call binding the contract method 0x2f26cb6e.
//
// Solidity: function depositRecord(address , uint256 ) view returns(address tokenAddress, address sender, address recipient, uint256 amount, uint256 fee, uint256 destinationChainId)
func (_Tantin *TantinSession) DepositRecord(arg0 common.Address, arg1 *big.Int) (struct {
	TokenAddress       common.Address
	Sender             common.Address
	Recipient          common.Address
	Amount             *big.Int
	Fee                *big.Int
	DestinationChainId *big.Int
}, error) {
	return _Tantin.Contract.DepositRecord(&_Tantin.CallOpts, arg0, arg1)
}

// DepositRecord is a free data retrieval call binding the contract method 0x2f26cb6e.
//
// Solidity: function depositRecord(address , uint256 ) view returns(address tokenAddress, address sender, address recipient, uint256 amount, uint256 fee, uint256 destinationChainId)
func (_Tantin *TantinCallerSession) DepositRecord(arg0 common.Address, arg1 *big.Int) (struct {
	TokenAddress       common.Address
	Sender             common.Address
	Recipient          common.Address
	Amount             *big.Int
	Fee                *big.Int
	DestinationChainId *big.Int
}, error) {
	return _Tantin.Contract.DepositRecord(&_Tantin.CallOpts, arg0, arg1)
}

// GetFee is a free data retrieval call binding the contract method 0xe5f3d3a5.
//
// Solidity: function getFee(bytes32 resourceId) view returns(uint256)
func (_Tantin *TantinCaller) GetFee(opts *bind.CallOpts, resourceId [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _Tantin.contract.Call(opts, &out, "getFee", resourceId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetFee is a free data retrieval call binding the contract method 0xe5f3d3a5.
//
// Solidity: function getFee(bytes32 resourceId) view returns(uint256)
func (_Tantin *TantinSession) GetFee(resourceId [32]byte) (*big.Int, error) {
	return _Tantin.Contract.GetFee(&_Tantin.CallOpts, resourceId)
}

// GetFee is a free data retrieval call binding the contract method 0xe5f3d3a5.
//
// Solidity: function getFee(bytes32 resourceId) view returns(uint256)
func (_Tantin *TantinCallerSession) GetFee(resourceId [32]byte) (*big.Int, error) {
	return _Tantin.Contract.GetFee(&_Tantin.CallOpts, resourceId)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Tantin *TantinCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _Tantin.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Tantin *TantinSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Tantin.Contract.GetRoleAdmin(&_Tantin.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Tantin *TantinCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Tantin.Contract.GetRoleAdmin(&_Tantin.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Tantin *TantinCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _Tantin.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Tantin *TantinSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Tantin.Contract.HasRole(&_Tantin.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Tantin *TantinCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Tantin.Contract.HasRole(&_Tantin.CallOpts, role, account)
}

// LocalNonce is a free data retrieval call binding the contract method 0x13cb3591.
//
// Solidity: function localNonce() view returns(uint256)
func (_Tantin *TantinCaller) LocalNonce(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Tantin.contract.Call(opts, &out, "localNonce")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LocalNonce is a free data retrieval call binding the contract method 0x13cb3591.
//
// Solidity: function localNonce() view returns(uint256)
func (_Tantin *TantinSession) LocalNonce() (*big.Int, error) {
	return _Tantin.Contract.LocalNonce(&_Tantin.CallOpts)
}

// LocalNonce is a free data retrieval call binding the contract method 0x13cb3591.
//
// Solidity: function localNonce() view returns(uint256)
func (_Tantin *TantinCallerSession) LocalNonce() (*big.Int, error) {
	return _Tantin.Contract.LocalNonce(&_Tantin.CallOpts)
}

// SigNonce is a free data retrieval call binding the contract method 0xcd868c9c.
//
// Solidity: function sigNonce() view returns(uint256)
func (_Tantin *TantinCaller) SigNonce(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Tantin.contract.Call(opts, &out, "sigNonce")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SigNonce is a free data retrieval call binding the contract method 0xcd868c9c.
//
// Solidity: function sigNonce() view returns(uint256)
func (_Tantin *TantinSession) SigNonce() (*big.Int, error) {
	return _Tantin.Contract.SigNonce(&_Tantin.CallOpts)
}

// SigNonce is a free data retrieval call binding the contract method 0xcd868c9c.
//
// Solidity: function sigNonce() view returns(uint256)
func (_Tantin *TantinCallerSession) SigNonce() (*big.Int, error) {
	return _Tantin.Contract.SigNonce(&_Tantin.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Tantin *TantinCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Tantin.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Tantin *TantinSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Tantin.Contract.SupportsInterface(&_Tantin.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Tantin *TantinCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Tantin.Contract.SupportsInterface(&_Tantin.CallOpts, interfaceId)
}

// AdminAddBlacklist is a paid mutator transaction binding the contract method 0xd8954afd.
//
// Solidity: function adminAddBlacklist(address user, bytes signature) returns()
func (_Tantin *TantinTransactor) AdminAddBlacklist(opts *bind.TransactOpts, user common.Address, signature []byte) (*types.Transaction, error) {
	return _Tantin.contract.Transact(opts, "adminAddBlacklist", user, signature)
}

// AdminAddBlacklist is a paid mutator transaction binding the contract method 0xd8954afd.
//
// Solidity: function adminAddBlacklist(address user, bytes signature) returns()
func (_Tantin *TantinSession) AdminAddBlacklist(user common.Address, signature []byte) (*types.Transaction, error) {
	return _Tantin.Contract.AdminAddBlacklist(&_Tantin.TransactOpts, user, signature)
}

// AdminAddBlacklist is a paid mutator transaction binding the contract method 0xd8954afd.
//
// Solidity: function adminAddBlacklist(address user, bytes signature) returns()
func (_Tantin *TantinTransactorSession) AdminAddBlacklist(user common.Address, signature []byte) (*types.Transaction, error) {
	return _Tantin.Contract.AdminAddBlacklist(&_Tantin.TransactOpts, user, signature)
}

// AdminRemoveBlacklist is a paid mutator transaction binding the contract method 0x807a5222.
//
// Solidity: function adminRemoveBlacklist(address user, bytes signature) returns()
func (_Tantin *TantinTransactor) AdminRemoveBlacklist(opts *bind.TransactOpts, user common.Address, signature []byte) (*types.Transaction, error) {
	return _Tantin.contract.Transact(opts, "adminRemoveBlacklist", user, signature)
}

// AdminRemoveBlacklist is a paid mutator transaction binding the contract method 0x807a5222.
//
// Solidity: function adminRemoveBlacklist(address user, bytes signature) returns()
func (_Tantin *TantinSession) AdminRemoveBlacklist(user common.Address, signature []byte) (*types.Transaction, error) {
	return _Tantin.Contract.AdminRemoveBlacklist(&_Tantin.TransactOpts, user, signature)
}

// AdminRemoveBlacklist is a paid mutator transaction binding the contract method 0x807a5222.
//
// Solidity: function adminRemoveBlacklist(address user, bytes signature) returns()
func (_Tantin *TantinTransactorSession) AdminRemoveBlacklist(user common.Address, signature []byte) (*types.Transaction, error) {
	return _Tantin.Contract.AdminRemoveBlacklist(&_Tantin.TransactOpts, user, signature)
}

// AdminSetEnv is a paid mutator transaction binding the contract method 0x2f370527.
//
// Solidity: function adminSetEnv(address feeAddress_, address bridgeAddress_, bytes signature_) returns()
func (_Tantin *TantinTransactor) AdminSetEnv(opts *bind.TransactOpts, feeAddress_ common.Address, bridgeAddress_ common.Address, signature_ []byte) (*types.Transaction, error) {
	return _Tantin.contract.Transact(opts, "adminSetEnv", feeAddress_, bridgeAddress_, signature_)
}

// AdminSetEnv is a paid mutator transaction binding the contract method 0x2f370527.
//
// Solidity: function adminSetEnv(address feeAddress_, address bridgeAddress_, bytes signature_) returns()
func (_Tantin *TantinSession) AdminSetEnv(feeAddress_ common.Address, bridgeAddress_ common.Address, signature_ []byte) (*types.Transaction, error) {
	return _Tantin.Contract.AdminSetEnv(&_Tantin.TransactOpts, feeAddress_, bridgeAddress_, signature_)
}

// AdminSetEnv is a paid mutator transaction binding the contract method 0x2f370527.
//
// Solidity: function adminSetEnv(address feeAddress_, address bridgeAddress_, bytes signature_) returns()
func (_Tantin *TantinTransactorSession) AdminSetEnv(feeAddress_ common.Address, bridgeAddress_ common.Address, signature_ []byte) (*types.Transaction, error) {
	return _Tantin.Contract.AdminSetEnv(&_Tantin.TransactOpts, feeAddress_, bridgeAddress_, signature_)
}

// AdminWithdraw is a paid mutator transaction binding the contract method 0x3ebc4a0f.
//
// Solidity: function adminWithdraw(address tokenAddress, uint256 amount, bytes signature) returns()
func (_Tantin *TantinTransactor) AdminWithdraw(opts *bind.TransactOpts, tokenAddress common.Address, amount *big.Int, signature []byte) (*types.Transaction, error) {
	return _Tantin.contract.Transact(opts, "adminWithdraw", tokenAddress, amount, signature)
}

// AdminWithdraw is a paid mutator transaction binding the contract method 0x3ebc4a0f.
//
// Solidity: function adminWithdraw(address tokenAddress, uint256 amount, bytes signature) returns()
func (_Tantin *TantinSession) AdminWithdraw(tokenAddress common.Address, amount *big.Int, signature []byte) (*types.Transaction, error) {
	return _Tantin.Contract.AdminWithdraw(&_Tantin.TransactOpts, tokenAddress, amount, signature)
}

// AdminWithdraw is a paid mutator transaction binding the contract method 0x3ebc4a0f.
//
// Solidity: function adminWithdraw(address tokenAddress, uint256 amount, bytes signature) returns()
func (_Tantin *TantinTransactorSession) AdminWithdraw(tokenAddress common.Address, amount *big.Int, signature []byte) (*types.Transaction, error) {
	return _Tantin.Contract.AdminWithdraw(&_Tantin.TransactOpts, tokenAddress, amount, signature)
}

// Deposit is a paid mutator transaction binding the contract method 0xc78089b1.
//
// Solidity: function deposit(uint256 destinationChainId, bytes32 resourceId, address recipient, uint256 amount, bytes signature) payable returns()
func (_Tantin *TantinTransactor) Deposit(opts *bind.TransactOpts, destinationChainId *big.Int, resourceId [32]byte, recipient common.Address, amount *big.Int, signature []byte) (*types.Transaction, error) {
	return _Tantin.contract.Transact(opts, "deposit", destinationChainId, resourceId, recipient, amount, signature)
}

// Deposit is a paid mutator transaction binding the contract method 0xc78089b1.
//
// Solidity: function deposit(uint256 destinationChainId, bytes32 resourceId, address recipient, uint256 amount, bytes signature) payable returns()
func (_Tantin *TantinSession) Deposit(destinationChainId *big.Int, resourceId [32]byte, recipient common.Address, amount *big.Int, signature []byte) (*types.Transaction, error) {
	return _Tantin.Contract.Deposit(&_Tantin.TransactOpts, destinationChainId, resourceId, recipient, amount, signature)
}

// Deposit is a paid mutator transaction binding the contract method 0xc78089b1.
//
// Solidity: function deposit(uint256 destinationChainId, bytes32 resourceId, address recipient, uint256 amount, bytes signature) payable returns()
func (_Tantin *TantinTransactorSession) Deposit(destinationChainId *big.Int, resourceId [32]byte, recipient common.Address, amount *big.Int, signature []byte) (*types.Transaction, error) {
	return _Tantin.Contract.Deposit(&_Tantin.TransactOpts, destinationChainId, resourceId, recipient, amount, signature)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Tantin *TantinTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Tantin.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Tantin *TantinSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Tantin.Contract.GrantRole(&_Tantin.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Tantin *TantinTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Tantin.Contract.GrantRole(&_Tantin.TransactOpts, role, account)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_Tantin *TantinTransactor) Initialize(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Tantin.contract.Transact(opts, "initialize")
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_Tantin *TantinSession) Initialize() (*types.Transaction, error) {
	return _Tantin.Contract.Initialize(&_Tantin.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_Tantin *TantinTransactorSession) Initialize() (*types.Transaction, error) {
	return _Tantin.Contract.Initialize(&_Tantin.TransactOpts)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_Tantin *TantinTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _Tantin.contract.Transact(opts, "renounceRole", role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_Tantin *TantinSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _Tantin.Contract.RenounceRole(&_Tantin.TransactOpts, role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_Tantin *TantinTransactorSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _Tantin.Contract.RenounceRole(&_Tantin.TransactOpts, role, callerConfirmation)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Tantin *TantinTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Tantin.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Tantin *TantinSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Tantin.Contract.RevokeRole(&_Tantin.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Tantin *TantinTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Tantin.Contract.RevokeRole(&_Tantin.TransactOpts, role, account)
}

// TantinAddBlacklistIterator is returned from FilterAddBlacklist and is used to iterate over the raw logs and unpacked data for AddBlacklist events raised by the Tantin contract.
type TantinAddBlacklistIterator struct {
	Event *TantinAddBlacklist // Event containing the contract specifics and raw log

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
func (it *TantinAddBlacklistIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TantinAddBlacklist)
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
		it.Event = new(TantinAddBlacklist)
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
func (it *TantinAddBlacklistIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TantinAddBlacklistIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TantinAddBlacklist represents a AddBlacklist event raised by the Tantin contract.
type TantinAddBlacklist struct {
	User common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterAddBlacklist is a free log retrieval operation binding the contract event 0x7e239e822fa537514cc6b38d8350bde5ce06a8f9282c77161b926fc077a81026.
//
// Solidity: event AddBlacklist(address indexed user)
func (_Tantin *TantinFilterer) FilterAddBlacklist(opts *bind.FilterOpts, user []common.Address) (*TantinAddBlacklistIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Tantin.contract.FilterLogs(opts, "AddBlacklist", userRule)
	if err != nil {
		return nil, err
	}
	return &TantinAddBlacklistIterator{contract: _Tantin.contract, event: "AddBlacklist", logs: logs, sub: sub}, nil
}

// WatchAddBlacklist is a free log subscription operation binding the contract event 0x7e239e822fa537514cc6b38d8350bde5ce06a8f9282c77161b926fc077a81026.
//
// Solidity: event AddBlacklist(address indexed user)
func (_Tantin *TantinFilterer) WatchAddBlacklist(opts *bind.WatchOpts, sink chan<- *TantinAddBlacklist, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Tantin.contract.WatchLogs(opts, "AddBlacklist", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TantinAddBlacklist)
				if err := _Tantin.contract.UnpackLog(event, "AddBlacklist", log); err != nil {
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

// ParseAddBlacklist is a log parse operation binding the contract event 0x7e239e822fa537514cc6b38d8350bde5ce06a8f9282c77161b926fc077a81026.
//
// Solidity: event AddBlacklist(address indexed user)
func (_Tantin *TantinFilterer) ParseAddBlacklist(log types.Log) (*TantinAddBlacklist, error) {
	event := new(TantinAddBlacklist)
	if err := _Tantin.contract.UnpackLog(event, "AddBlacklist", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TantinDepositEventIterator is returned from FilterDepositEvent and is used to iterate over the raw logs and unpacked data for DepositEvent events raised by the Tantin contract.
type TantinDepositEventIterator struct {
	Event *TantinDepositEvent // Event containing the contract specifics and raw log

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
func (it *TantinDepositEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TantinDepositEvent)
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
		it.Event = new(TantinDepositEvent)
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
func (it *TantinDepositEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TantinDepositEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TantinDepositEvent represents a DepositEvent event raised by the Tantin contract.
type TantinDepositEvent struct {
	Depositer          common.Address
	Recipient          common.Address
	Amount             *big.Int
	TokenAddress       common.Address
	DepositNonce       *big.Int
	DestinationChainId *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterDepositEvent is a free log retrieval operation binding the contract event 0x6c243d2392e50eecacd348bcc6bffa9906438210df6f32b1ef5eed2089c21bad.
//
// Solidity: event DepositEvent(address indexed depositer, address indexed recipient, uint256 indexed amount, address tokenAddress, uint256 depositNonce, uint256 destinationChainId)
func (_Tantin *TantinFilterer) FilterDepositEvent(opts *bind.FilterOpts, depositer []common.Address, recipient []common.Address, amount []*big.Int) (*TantinDepositEventIterator, error) {

	var depositerRule []interface{}
	for _, depositerItem := range depositer {
		depositerRule = append(depositerRule, depositerItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}
	var amountRule []interface{}
	for _, amountItem := range amount {
		amountRule = append(amountRule, amountItem)
	}

	logs, sub, err := _Tantin.contract.FilterLogs(opts, "DepositEvent", depositerRule, recipientRule, amountRule)
	if err != nil {
		return nil, err
	}
	return &TantinDepositEventIterator{contract: _Tantin.contract, event: "DepositEvent", logs: logs, sub: sub}, nil
}

// WatchDepositEvent is a free log subscription operation binding the contract event 0x6c243d2392e50eecacd348bcc6bffa9906438210df6f32b1ef5eed2089c21bad.
//
// Solidity: event DepositEvent(address indexed depositer, address indexed recipient, uint256 indexed amount, address tokenAddress, uint256 depositNonce, uint256 destinationChainId)
func (_Tantin *TantinFilterer) WatchDepositEvent(opts *bind.WatchOpts, sink chan<- *TantinDepositEvent, depositer []common.Address, recipient []common.Address, amount []*big.Int) (event.Subscription, error) {

	var depositerRule []interface{}
	for _, depositerItem := range depositer {
		depositerRule = append(depositerRule, depositerItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}
	var amountRule []interface{}
	for _, amountItem := range amount {
		amountRule = append(amountRule, amountItem)
	}

	logs, sub, err := _Tantin.contract.WatchLogs(opts, "DepositEvent", depositerRule, recipientRule, amountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TantinDepositEvent)
				if err := _Tantin.contract.UnpackLog(event, "DepositEvent", log); err != nil {
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

// ParseDepositEvent is a log parse operation binding the contract event 0x6c243d2392e50eecacd348bcc6bffa9906438210df6f32b1ef5eed2089c21bad.
//
// Solidity: event DepositEvent(address indexed depositer, address indexed recipient, uint256 indexed amount, address tokenAddress, uint256 depositNonce, uint256 destinationChainId)
func (_Tantin *TantinFilterer) ParseDepositEvent(log types.Log) (*TantinDepositEvent, error) {
	event := new(TantinDepositEvent)
	if err := _Tantin.contract.UnpackLog(event, "DepositEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TantinDepositNftEventIterator is returned from FilterDepositNftEvent and is used to iterate over the raw logs and unpacked data for DepositNftEvent events raised by the Tantin contract.
type TantinDepositNftEventIterator struct {
	Event *TantinDepositNftEvent // Event containing the contract specifics and raw log

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
func (it *TantinDepositNftEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TantinDepositNftEvent)
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
		it.Event = new(TantinDepositNftEvent)
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
func (it *TantinDepositNftEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TantinDepositNftEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TantinDepositNftEvent represents a DepositNftEvent event raised by the Tantin contract.
type TantinDepositNftEvent struct {
	Depositer          common.Address
	Recipient          common.Address
	Amount             *big.Int
	TokenId            *big.Int
	TokenAddress       common.Address
	DepositNonce       *big.Int
	DestinationChainId *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterDepositNftEvent is a free log retrieval operation binding the contract event 0xfea593f6afb78012f4f829131473cb41b439bc2645c14d047c04cb115438158e.
//
// Solidity: event DepositNftEvent(address indexed depositer, address indexed recipient, uint256 indexed amount, uint256 tokenId, address tokenAddress, uint256 depositNonce, uint256 destinationChainId)
func (_Tantin *TantinFilterer) FilterDepositNftEvent(opts *bind.FilterOpts, depositer []common.Address, recipient []common.Address, amount []*big.Int) (*TantinDepositNftEventIterator, error) {

	var depositerRule []interface{}
	for _, depositerItem := range depositer {
		depositerRule = append(depositerRule, depositerItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}
	var amountRule []interface{}
	for _, amountItem := range amount {
		amountRule = append(amountRule, amountItem)
	}

	logs, sub, err := _Tantin.contract.FilterLogs(opts, "DepositNftEvent", depositerRule, recipientRule, amountRule)
	if err != nil {
		return nil, err
	}
	return &TantinDepositNftEventIterator{contract: _Tantin.contract, event: "DepositNftEvent", logs: logs, sub: sub}, nil
}

// WatchDepositNftEvent is a free log subscription operation binding the contract event 0xfea593f6afb78012f4f829131473cb41b439bc2645c14d047c04cb115438158e.
//
// Solidity: event DepositNftEvent(address indexed depositer, address indexed recipient, uint256 indexed amount, uint256 tokenId, address tokenAddress, uint256 depositNonce, uint256 destinationChainId)
func (_Tantin *TantinFilterer) WatchDepositNftEvent(opts *bind.WatchOpts, sink chan<- *TantinDepositNftEvent, depositer []common.Address, recipient []common.Address, amount []*big.Int) (event.Subscription, error) {

	var depositerRule []interface{}
	for _, depositerItem := range depositer {
		depositerRule = append(depositerRule, depositerItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}
	var amountRule []interface{}
	for _, amountItem := range amount {
		amountRule = append(amountRule, amountItem)
	}

	logs, sub, err := _Tantin.contract.WatchLogs(opts, "DepositNftEvent", depositerRule, recipientRule, amountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TantinDepositNftEvent)
				if err := _Tantin.contract.UnpackLog(event, "DepositNftEvent", log); err != nil {
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

// ParseDepositNftEvent is a log parse operation binding the contract event 0xfea593f6afb78012f4f829131473cb41b439bc2645c14d047c04cb115438158e.
//
// Solidity: event DepositNftEvent(address indexed depositer, address indexed recipient, uint256 indexed amount, uint256 tokenId, address tokenAddress, uint256 depositNonce, uint256 destinationChainId)
func (_Tantin *TantinFilterer) ParseDepositNftEvent(log types.Log) (*TantinDepositNftEvent, error) {
	event := new(TantinDepositNftEvent)
	if err := _Tantin.contract.UnpackLog(event, "DepositNftEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TantinExecuteNftEventIterator is returned from FilterExecuteNftEvent and is used to iterate over the raw logs and unpacked data for ExecuteNftEvent events raised by the Tantin contract.
type TantinExecuteNftEventIterator struct {
	Event *TantinExecuteNftEvent // Event containing the contract specifics and raw log

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
func (it *TantinExecuteNftEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TantinExecuteNftEvent)
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
		it.Event = new(TantinExecuteNftEvent)
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
func (it *TantinExecuteNftEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TantinExecuteNftEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TantinExecuteNftEvent represents a ExecuteNftEvent event raised by the Tantin contract.
type TantinExecuteNftEvent struct {
	Depositer          common.Address
	Recipient          common.Address
	Amount             *big.Int
	TokenId            *big.Int
	TokenAddress       common.Address
	OriginDepositNonce *big.Int
	OriginChainId      *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterExecuteNftEvent is a free log retrieval operation binding the contract event 0xfb81f733903da58465673369a80984f69ce0908d48c4a975d00ae726225c525b.
//
// Solidity: event ExecuteNftEvent(address indexed depositer, address indexed recipient, uint256 indexed amount, uint256 tokenId, address tokenAddress, uint256 originDepositNonce, uint256 originChainId)
func (_Tantin *TantinFilterer) FilterExecuteNftEvent(opts *bind.FilterOpts, depositer []common.Address, recipient []common.Address, amount []*big.Int) (*TantinExecuteNftEventIterator, error) {

	var depositerRule []interface{}
	for _, depositerItem := range depositer {
		depositerRule = append(depositerRule, depositerItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}
	var amountRule []interface{}
	for _, amountItem := range amount {
		amountRule = append(amountRule, amountItem)
	}

	logs, sub, err := _Tantin.contract.FilterLogs(opts, "ExecuteNftEvent", depositerRule, recipientRule, amountRule)
	if err != nil {
		return nil, err
	}
	return &TantinExecuteNftEventIterator{contract: _Tantin.contract, event: "ExecuteNftEvent", logs: logs, sub: sub}, nil
}

// WatchExecuteNftEvent is a free log subscription operation binding the contract event 0xfb81f733903da58465673369a80984f69ce0908d48c4a975d00ae726225c525b.
//
// Solidity: event ExecuteNftEvent(address indexed depositer, address indexed recipient, uint256 indexed amount, uint256 tokenId, address tokenAddress, uint256 originDepositNonce, uint256 originChainId)
func (_Tantin *TantinFilterer) WatchExecuteNftEvent(opts *bind.WatchOpts, sink chan<- *TantinExecuteNftEvent, depositer []common.Address, recipient []common.Address, amount []*big.Int) (event.Subscription, error) {

	var depositerRule []interface{}
	for _, depositerItem := range depositer {
		depositerRule = append(depositerRule, depositerItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}
	var amountRule []interface{}
	for _, amountItem := range amount {
		amountRule = append(amountRule, amountItem)
	}

	logs, sub, err := _Tantin.contract.WatchLogs(opts, "ExecuteNftEvent", depositerRule, recipientRule, amountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TantinExecuteNftEvent)
				if err := _Tantin.contract.UnpackLog(event, "ExecuteNftEvent", log); err != nil {
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

// ParseExecuteNftEvent is a log parse operation binding the contract event 0xfb81f733903da58465673369a80984f69ce0908d48c4a975d00ae726225c525b.
//
// Solidity: event ExecuteNftEvent(address indexed depositer, address indexed recipient, uint256 indexed amount, uint256 tokenId, address tokenAddress, uint256 originDepositNonce, uint256 originChainId)
func (_Tantin *TantinFilterer) ParseExecuteNftEvent(log types.Log) (*TantinExecuteNftEvent, error) {
	event := new(TantinExecuteNftEvent)
	if err := _Tantin.contract.UnpackLog(event, "ExecuteNftEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TantinInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Tantin contract.
type TantinInitializedIterator struct {
	Event *TantinInitialized // Event containing the contract specifics and raw log

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
func (it *TantinInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TantinInitialized)
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
		it.Event = new(TantinInitialized)
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
func (it *TantinInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TantinInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TantinInitialized represents a Initialized event raised by the Tantin contract.
type TantinInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_Tantin *TantinFilterer) FilterInitialized(opts *bind.FilterOpts) (*TantinInitializedIterator, error) {

	logs, sub, err := _Tantin.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &TantinInitializedIterator{contract: _Tantin.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_Tantin *TantinFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *TantinInitialized) (event.Subscription, error) {

	logs, sub, err := _Tantin.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TantinInitialized)
				if err := _Tantin.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_Tantin *TantinFilterer) ParseInitialized(log types.Log) (*TantinInitialized, error) {
	event := new(TantinInitialized)
	if err := _Tantin.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TantinRemoveBlacklistIterator is returned from FilterRemoveBlacklist and is used to iterate over the raw logs and unpacked data for RemoveBlacklist events raised by the Tantin contract.
type TantinRemoveBlacklistIterator struct {
	Event *TantinRemoveBlacklist // Event containing the contract specifics and raw log

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
func (it *TantinRemoveBlacklistIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TantinRemoveBlacklist)
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
		it.Event = new(TantinRemoveBlacklist)
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
func (it *TantinRemoveBlacklistIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TantinRemoveBlacklistIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TantinRemoveBlacklist represents a RemoveBlacklist event raised by the Tantin contract.
type TantinRemoveBlacklist struct {
	User common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterRemoveBlacklist is a free log retrieval operation binding the contract event 0x54646b2d47b5332deb93b310542f2c11bc9351e59950cdfb3ba518af28f13d29.
//
// Solidity: event RemoveBlacklist(address indexed user)
func (_Tantin *TantinFilterer) FilterRemoveBlacklist(opts *bind.FilterOpts, user []common.Address) (*TantinRemoveBlacklistIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Tantin.contract.FilterLogs(opts, "RemoveBlacklist", userRule)
	if err != nil {
		return nil, err
	}
	return &TantinRemoveBlacklistIterator{contract: _Tantin.contract, event: "RemoveBlacklist", logs: logs, sub: sub}, nil
}

// WatchRemoveBlacklist is a free log subscription operation binding the contract event 0x54646b2d47b5332deb93b310542f2c11bc9351e59950cdfb3ba518af28f13d29.
//
// Solidity: event RemoveBlacklist(address indexed user)
func (_Tantin *TantinFilterer) WatchRemoveBlacklist(opts *bind.WatchOpts, sink chan<- *TantinRemoveBlacklist, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Tantin.contract.WatchLogs(opts, "RemoveBlacklist", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TantinRemoveBlacklist)
				if err := _Tantin.contract.UnpackLog(event, "RemoveBlacklist", log); err != nil {
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

// ParseRemoveBlacklist is a log parse operation binding the contract event 0x54646b2d47b5332deb93b310542f2c11bc9351e59950cdfb3ba518af28f13d29.
//
// Solidity: event RemoveBlacklist(address indexed user)
func (_Tantin *TantinFilterer) ParseRemoveBlacklist(log types.Log) (*TantinRemoveBlacklist, error) {
	event := new(TantinRemoveBlacklist)
	if err := _Tantin.contract.UnpackLog(event, "RemoveBlacklist", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TantinRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the Tantin contract.
type TantinRoleAdminChangedIterator struct {
	Event *TantinRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *TantinRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TantinRoleAdminChanged)
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
		it.Event = new(TantinRoleAdminChanged)
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
func (it *TantinRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TantinRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TantinRoleAdminChanged represents a RoleAdminChanged event raised by the Tantin contract.
type TantinRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Tantin *TantinFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*TantinRoleAdminChangedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _Tantin.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &TantinRoleAdminChangedIterator{contract: _Tantin.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Tantin *TantinFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *TantinRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _Tantin.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TantinRoleAdminChanged)
				if err := _Tantin.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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

// ParseRoleAdminChanged is a log parse operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Tantin *TantinFilterer) ParseRoleAdminChanged(log types.Log) (*TantinRoleAdminChanged, error) {
	event := new(TantinRoleAdminChanged)
	if err := _Tantin.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TantinRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the Tantin contract.
type TantinRoleGrantedIterator struct {
	Event *TantinRoleGranted // Event containing the contract specifics and raw log

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
func (it *TantinRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TantinRoleGranted)
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
		it.Event = new(TantinRoleGranted)
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
func (it *TantinRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TantinRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TantinRoleGranted represents a RoleGranted event raised by the Tantin contract.
type TantinRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Tantin *TantinFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*TantinRoleGrantedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Tantin.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &TantinRoleGrantedIterator{contract: _Tantin.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Tantin *TantinFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *TantinRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Tantin.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TantinRoleGranted)
				if err := _Tantin.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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

// ParseRoleGranted is a log parse operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Tantin *TantinFilterer) ParseRoleGranted(log types.Log) (*TantinRoleGranted, error) {
	event := new(TantinRoleGranted)
	if err := _Tantin.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TantinRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the Tantin contract.
type TantinRoleRevokedIterator struct {
	Event *TantinRoleRevoked // Event containing the contract specifics and raw log

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
func (it *TantinRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TantinRoleRevoked)
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
		it.Event = new(TantinRoleRevoked)
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
func (it *TantinRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TantinRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TantinRoleRevoked represents a RoleRevoked event raised by the Tantin contract.
type TantinRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Tantin *TantinFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*TantinRoleRevokedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Tantin.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &TantinRoleRevokedIterator{contract: _Tantin.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Tantin *TantinFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *TantinRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Tantin.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TantinRoleRevoked)
				if err := _Tantin.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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

// ParseRoleRevoked is a log parse operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Tantin *TantinFilterer) ParseRoleRevoked(log types.Log) (*TantinRoleRevoked, error) {
	event := new(TantinRoleRevoked)
	if err := _Tantin.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TantinSetTokenEventIterator is returned from FilterSetTokenEvent and is used to iterate over the raw logs and unpacked data for SetTokenEvent events raised by the Tantin contract.
type TantinSetTokenEventIterator struct {
	Event *TantinSetTokenEvent // Event containing the contract specifics and raw log

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
func (it *TantinSetTokenEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TantinSetTokenEvent)
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
		it.Event = new(TantinSetTokenEvent)
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
func (it *TantinSetTokenEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TantinSetTokenEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TantinSetTokenEvent represents a SetTokenEvent event raised by the Tantin contract.
type TantinSetTokenEvent struct {
	ResourceID   [32]byte
	AssetsType   uint8
	TokenAddress common.Address
	Burnable     bool
	Mintable     bool
	Pause        bool
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterSetTokenEvent is a free log retrieval operation binding the contract event 0xc2f5fd768773c30110268117e11f879a4578df04d9cd8fcb91daed19a67a7a31.
//
// Solidity: event SetTokenEvent(bytes32 indexed resourceID, uint8 assetsType, address tokenAddress, bool burnable, bool mintable, bool pause)
func (_Tantin *TantinFilterer) FilterSetTokenEvent(opts *bind.FilterOpts, resourceID [][32]byte) (*TantinSetTokenEventIterator, error) {

	var resourceIDRule []interface{}
	for _, resourceIDItem := range resourceID {
		resourceIDRule = append(resourceIDRule, resourceIDItem)
	}

	logs, sub, err := _Tantin.contract.FilterLogs(opts, "SetTokenEvent", resourceIDRule)
	if err != nil {
		return nil, err
	}
	return &TantinSetTokenEventIterator{contract: _Tantin.contract, event: "SetTokenEvent", logs: logs, sub: sub}, nil
}

// WatchSetTokenEvent is a free log subscription operation binding the contract event 0xc2f5fd768773c30110268117e11f879a4578df04d9cd8fcb91daed19a67a7a31.
//
// Solidity: event SetTokenEvent(bytes32 indexed resourceID, uint8 assetsType, address tokenAddress, bool burnable, bool mintable, bool pause)
func (_Tantin *TantinFilterer) WatchSetTokenEvent(opts *bind.WatchOpts, sink chan<- *TantinSetTokenEvent, resourceID [][32]byte) (event.Subscription, error) {

	var resourceIDRule []interface{}
	for _, resourceIDItem := range resourceID {
		resourceIDRule = append(resourceIDRule, resourceIDItem)
	}

	logs, sub, err := _Tantin.contract.WatchLogs(opts, "SetTokenEvent", resourceIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TantinSetTokenEvent)
				if err := _Tantin.contract.UnpackLog(event, "SetTokenEvent", log); err != nil {
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

// ParseSetTokenEvent is a log parse operation binding the contract event 0xc2f5fd768773c30110268117e11f879a4578df04d9cd8fcb91daed19a67a7a31.
//
// Solidity: event SetTokenEvent(bytes32 indexed resourceID, uint8 assetsType, address tokenAddress, bool burnable, bool mintable, bool pause)
func (_Tantin *TantinFilterer) ParseSetTokenEvent(log types.Log) (*TantinSetTokenEvent, error) {
	event := new(TantinSetTokenEvent)
	if err := _Tantin.contract.UnpackLog(event, "SetTokenEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
