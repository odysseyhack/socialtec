// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package market

import (
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
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = abi.U256
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// GivoABI is the input ABI used to generate the binding from.
const GivoABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"node_id\",\"type\":\"uint256\"},{\"name\":\"good_id\",\"type\":\"uint256\"},{\"name\":\"refer_id\",\"type\":\"uint256\"}],\"name\":\"refer_intrest\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"to\",\"type\":\"uint256\"},{\"name\":\"get_good_id\",\"type\":\"uint256\"},{\"name\":\"give_owner\",\"type\":\"uint256\"},{\"name\":\"give_good_id\",\"type\":\"uint256\"}],\"name\":\"cycle_formed\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"offers\",\"outputs\":[{\"name\":\"node\",\"type\":\"address\"},{\"name\":\"name\",\"type\":\"string\"},{\"name\":\"ipfs_image\",\"type\":\"string\"},{\"name\":\"ipfs_details\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"node_id\",\"type\":\"uint256\"},{\"name\":\"good_id\",\"type\":\"uint256\"}],\"name\":\"add_intrest\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"node_cntr\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"good_id\",\"type\":\"uint256\"}],\"name\":\"delete_offer\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"node_id\",\"type\":\"uint256\"},{\"name\":\"good_id\",\"type\":\"uint256\"}],\"name\":\"delete_intrest\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"name\",\"type\":\"string\"},{\"name\":\"ipfs_image\",\"type\":\"string\"},{\"name\":\"ipfs_details\",\"type\":\"string\"}],\"name\":\"create_offer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"to\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"good_id\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"from\",\"type\":\"uint256\"}],\"name\":\"intrested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"to\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"good_id\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"from\",\"type\":\"uint256\"}],\"name\":\"not_intrested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"to\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"good_id\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"from\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"refer_id\",\"type\":\"uint256\"}],\"name\":\"refer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"owner\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"good_id\",\"type\":\"uint256\"}],\"name\":\"deleted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"to\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"get_owner\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"get_good_id\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"give_owner\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"give_good_id\",\"type\":\"uint256\"}],\"name\":\"chained\",\"type\":\"event\"}]"

// GivoBin is the compiled bytecode used for deploying new contracts.
const GivoBin = `0x60806040526000805534801561001457600080fd5b50610be9806100246000396000f3fe608060405234801561001057600080fd5b50600436106100885760003560e01c806386c123621161005b57806386c123621461028a578063881d9637146102a4578063a7fb194b14610267578063c175cdd3146102c157610088565b8063394db5741461008d57806343cef25e146100b85780634d03a9a5146100e757806356679f0c14610267575b600080fd5b6100b6600480360360608110156100a357600080fd5b5080359060208101359060400135610489565b005b6100b6600480360360808110156100ce57600080fd5b50803590602081013590604081013590606001356104e4565b61010a600480360360408110156100fd57600080fd5b5080359060200135610547565b60405180856001600160a01b03166001600160a01b03168152602001806020018060200180602001848103845287818151815260200191508051906020019080838360005b8381101561016757818101518382015260200161014f565b50505050905090810190601f1680156101945780820380516001836020036101000a031916815260200191505b50848103835286518152865160209182019188019080838360005b838110156101c75781810151838201526020016101af565b50505050905090810190601f1680156101f45780820380516001836020036101000a031916815260200191505b50848103825285518152855160209182019187019080838360005b8381101561022757818101518382015260200161020f565b50505050905090810190601f1680156102545780820380516001836020036101000a031916815260200191505b5097505050505050505060405180910390f35b6100b66004803603604081101561027d57600080fd5b5080359060200135610737565b61029261078a565b60408051918252519081900360200190f35b6100b6600480360360208110156102ba57600080fd5b5035610790565b610475600480360360608110156102d757600080fd5b8101906020810181356401000000008111156102f257600080fd5b82018360208201111561030457600080fd5b8035906020019184600183028401116401000000008311171561032657600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929594936020810193503591505064010000000081111561037957600080fd5b82018360208201111561038b57600080fd5b803590602001918460018302840111640100000000831117156103ad57600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929594936020810193503591505064010000000081111561040057600080fd5b82018360208201111561041257600080fd5b8035906020019184600183028401116401000000008311171561043457600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550610939945050505050565b604080519115158252519081900360200190f35b33600090815260026020908152604091829020548251868152918201859052818301526060810183905290517fcc99f51473cc3bf43367db4f64d4efcdfd1acd4c70b4155cb7ffade4f2f6d3629181900360800190a1505050565b3360009081526002602090815260409182902054825187815291820152808201859052606081018490526080810183905290517f17d3c3f8c4f0fa8d1c6e1bb5edde9fda0537a5f56cafed38fea7b389ea9116879181900360a00190a150505050565b6001602052816000526040600020818154811061056057fe5b600091825260209182902060049091020180546001808301805460408051601f60026000199685161561010002969096019093169490940491820187900487028401870190528083526001600160a01b0390931696509294509283018282801561060b5780601f106105e05761010080835404028352916020019161060b565b820191906000526020600020905b8154815290600101906020018083116105ee57829003601f168201915b50505060028085018054604080516020601f600019610100600187161502019094169590950492830185900485028101850190915281815295969594509092509083018282801561069d5780601f106106725761010080835404028352916020019161069d565b820191906000526020600020905b81548152906001019060200180831161068057829003601f168201915b5050505060038301805460408051602060026001851615610100026000190190941693909304601f810184900484028201840190925281815294959493509083018282801561072d5780601f106107025761010080835404028352916020019161072d565b820191906000526020600020905b81548152906001019060200180831161071057829003601f168201915b5050505050905084565b336000908152600260209081526040918290205482518581529182018490528183015290517fcfeec790cb492f73558ea2f8bbac696dcb49ec548daa4ee42f15fe49088159769181900360600190a15050565b60005481565b3360009081526002602090815260408083205483526001909152902080548211156107ba57600080fd5b8054819060001981019081106107cc57fe5b90600052602060002090600402018183815481106107e657fe5b60009182526020909120825460049092020180546001600160a01b0319166001600160a01b039092169190911781556001808301805461083d92808501929160026000199282161561010002929092011604610a35565b5060028281018054610862928481019291600019610100600183161502011604610a35565b506003820181600301908054600181600116156101000203166002900461088a929190610a35565b509050508080548061089857fe5b60008281526020812060046000199093019283020180546001600160a01b0319168155906108c96001830182610aba565b6108d7600283016000610aba565b6108e5600383016000610aba565b505090553360009081526002602090815260409182902054825190815290810184905281517fa7510b9f06bca625a2290e8871362d0ccc23a2b155fc82b80c66cf8c285c7323929181900390910190a15050565b336000908152600260209081526040808320548352600190915281208054610974576000805433825260026020526040822081905560010190555b61097c610b01565b5060408051608081018252338152602080820188815292820187905260608201869052835460018082018087556000878152849020855160049094020180546001600160a01b0319166001600160a01b039094169390931783559451805194959486946109ed938501920190610b32565b5060408201518051610a09916002840191602090910190610b32565b5060608201518051610a25916003840191602090910190610b32565b5060019998505050505050505050565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10610a6e5780548555610aaa565b82800160010185558215610aaa57600052602060002091601f016020900482015b82811115610aaa578254825591600101919060010190610a8f565b50610ab6929150610ba0565b5090565b50805460018160011615610100020316600290046000825580601f10610ae05750610afe565b601f016020900490600052602060002090810190610afe9190610ba0565b50565b604051806080016040528060006001600160a01b031681526020016060815260200160608152602001606081525090565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10610b7357805160ff1916838001178555610aaa565b82800160010185558215610aaa579182015b82811115610aaa578251825591602001919060010190610b85565b610bba91905b80821115610ab65760008155600101610ba6565b9056fea165627a7a72305820bc40bceffcd7f23907969dc213e82a8fe3f4697d8d1ba72843f973cba10605460029`

// DeployGivo deploys a new Ethereum contract, binding an instance of Givo to it.
func DeployGivo(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Givo, error) {
	parsed, err := abi.JSON(strings.NewReader(GivoABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(GivoBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Givo{GivoCaller: GivoCaller{contract: contract}, GivoTransactor: GivoTransactor{contract: contract}, GivoFilterer: GivoFilterer{contract: contract}}, nil
}

// Givo is an auto generated Go binding around an Ethereum contract.
type Givo struct {
	GivoCaller     // Read-only binding to the contract
	GivoTransactor // Write-only binding to the contract
	GivoFilterer   // Log filterer for contract events
}

// GivoCaller is an auto generated read-only Go binding around an Ethereum contract.
type GivoCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GivoTransactor is an auto generated write-only Go binding around an Ethereum contract.
type GivoTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GivoFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type GivoFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GivoSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type GivoSession struct {
	Contract     *Givo             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// GivoCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type GivoCallerSession struct {
	Contract *GivoCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// GivoTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type GivoTransactorSession struct {
	Contract     *GivoTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// GivoRaw is an auto generated low-level Go binding around an Ethereum contract.
type GivoRaw struct {
	Contract *Givo // Generic contract binding to access the raw methods on
}

// GivoCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type GivoCallerRaw struct {
	Contract *GivoCaller // Generic read-only contract binding to access the raw methods on
}

// GivoTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type GivoTransactorRaw struct {
	Contract *GivoTransactor // Generic write-only contract binding to access the raw methods on
}

// NewGivo creates a new instance of Givo, bound to a specific deployed contract.
func NewGivo(address common.Address, backend bind.ContractBackend) (*Givo, error) {
	contract, err := bindGivo(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Givo{GivoCaller: GivoCaller{contract: contract}, GivoTransactor: GivoTransactor{contract: contract}, GivoFilterer: GivoFilterer{contract: contract}}, nil
}

// NewGivoCaller creates a new read-only instance of Givo, bound to a specific deployed contract.
func NewGivoCaller(address common.Address, caller bind.ContractCaller) (*GivoCaller, error) {
	contract, err := bindGivo(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &GivoCaller{contract: contract}, nil
}

// NewGivoTransactor creates a new write-only instance of Givo, bound to a specific deployed contract.
func NewGivoTransactor(address common.Address, transactor bind.ContractTransactor) (*GivoTransactor, error) {
	contract, err := bindGivo(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &GivoTransactor{contract: contract}, nil
}

// NewGivoFilterer creates a new log filterer instance of Givo, bound to a specific deployed contract.
func NewGivoFilterer(address common.Address, filterer bind.ContractFilterer) (*GivoFilterer, error) {
	contract, err := bindGivo(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &GivoFilterer{contract: contract}, nil
}

// bindGivo binds a generic wrapper to an already deployed contract.
func bindGivo(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(GivoABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Givo *GivoRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Givo.Contract.GivoCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Givo *GivoRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Givo.Contract.GivoTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Givo *GivoRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Givo.Contract.GivoTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Givo *GivoCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Givo.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Givo *GivoTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Givo.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Givo *GivoTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Givo.Contract.contract.Transact(opts, method, params...)
}

// NodeCntr is a free data retrieval call binding the contract method 0x86c12362.
//
// Solidity: function node_cntr() constant returns(uint256)
func (_Givo *GivoCaller) NodeCntr(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Givo.contract.Call(opts, out, "node_cntr")
	return *ret0, err
}

// NodeCntr is a free data retrieval call binding the contract method 0x86c12362.
//
// Solidity: function node_cntr() constant returns(uint256)
func (_Givo *GivoSession) NodeCntr() (*big.Int, error) {
	return _Givo.Contract.NodeCntr(&_Givo.CallOpts)
}

// NodeCntr is a free data retrieval call binding the contract method 0x86c12362.
//
// Solidity: function node_cntr() constant returns(uint256)
func (_Givo *GivoCallerSession) NodeCntr() (*big.Int, error) {
	return _Givo.Contract.NodeCntr(&_Givo.CallOpts)
}

// Offers is a free data retrieval call binding the contract method 0x4d03a9a5.
//
// Solidity: function offers(uint256 , uint256 ) constant returns(address node, string name, string ipfs_image, string ipfs_details)
func (_Givo *GivoCaller) Offers(opts *bind.CallOpts, arg0 *big.Int, arg1 *big.Int) (struct {
	Node        common.Address
	Name        string
	IpfsImage   string
	IpfsDetails string
}, error) {
	ret := new(struct {
		Node        common.Address
		Name        string
		IpfsImage   string
		IpfsDetails string
	})
	out := ret
	err := _Givo.contract.Call(opts, out, "offers", arg0, arg1)
	return *ret, err
}

// Offers is a free data retrieval call binding the contract method 0x4d03a9a5.
//
// Solidity: function offers(uint256 , uint256 ) constant returns(address node, string name, string ipfs_image, string ipfs_details)
func (_Givo *GivoSession) Offers(arg0 *big.Int, arg1 *big.Int) (struct {
	Node        common.Address
	Name        string
	IpfsImage   string
	IpfsDetails string
}, error) {
	return _Givo.Contract.Offers(&_Givo.CallOpts, arg0, arg1)
}

// Offers is a free data retrieval call binding the contract method 0x4d03a9a5.
//
// Solidity: function offers(uint256 , uint256 ) constant returns(address node, string name, string ipfs_image, string ipfs_details)
func (_Givo *GivoCallerSession) Offers(arg0 *big.Int, arg1 *big.Int) (struct {
	Node        common.Address
	Name        string
	IpfsImage   string
	IpfsDetails string
}, error) {
	return _Givo.Contract.Offers(&_Givo.CallOpts, arg0, arg1)
}

// AddIntrest is a paid mutator transaction binding the contract method 0x56679f0c.
//
// Solidity: function add_intrest(uint256 node_id, uint256 good_id) returns()
func (_Givo *GivoTransactor) AddIntrest(opts *bind.TransactOpts, node_id *big.Int, good_id *big.Int) (*types.Transaction, error) {
	return _Givo.contract.Transact(opts, "add_intrest", node_id, good_id)
}

// AddIntrest is a paid mutator transaction binding the contract method 0x56679f0c.
//
// Solidity: function add_intrest(uint256 node_id, uint256 good_id) returns()
func (_Givo *GivoSession) AddIntrest(node_id *big.Int, good_id *big.Int) (*types.Transaction, error) {
	return _Givo.Contract.AddIntrest(&_Givo.TransactOpts, node_id, good_id)
}

// AddIntrest is a paid mutator transaction binding the contract method 0x56679f0c.
//
// Solidity: function add_intrest(uint256 node_id, uint256 good_id) returns()
func (_Givo *GivoTransactorSession) AddIntrest(node_id *big.Int, good_id *big.Int) (*types.Transaction, error) {
	return _Givo.Contract.AddIntrest(&_Givo.TransactOpts, node_id, good_id)
}

// CreateOffer is a paid mutator transaction binding the contract method 0xc175cdd3.
//
// Solidity: function create_offer(string name, string ipfs_image, string ipfs_details) returns(bool)
func (_Givo *GivoTransactor) CreateOffer(opts *bind.TransactOpts, name string, ipfs_image string, ipfs_details string) (*types.Transaction, error) {
	return _Givo.contract.Transact(opts, "create_offer", name, ipfs_image, ipfs_details)
}

// CreateOffer is a paid mutator transaction binding the contract method 0xc175cdd3.
//
// Solidity: function create_offer(string name, string ipfs_image, string ipfs_details) returns(bool)
func (_Givo *GivoSession) CreateOffer(name string, ipfs_image string, ipfs_details string) (*types.Transaction, error) {
	return _Givo.Contract.CreateOffer(&_Givo.TransactOpts, name, ipfs_image, ipfs_details)
}

// CreateOffer is a paid mutator transaction binding the contract method 0xc175cdd3.
//
// Solidity: function create_offer(string name, string ipfs_image, string ipfs_details) returns(bool)
func (_Givo *GivoTransactorSession) CreateOffer(name string, ipfs_image string, ipfs_details string) (*types.Transaction, error) {
	return _Givo.Contract.CreateOffer(&_Givo.TransactOpts, name, ipfs_image, ipfs_details)
}

// CycleFormed is a paid mutator transaction binding the contract method 0x43cef25e.
//
// Solidity: function cycle_formed(uint256 to, uint256 get_good_id, uint256 give_owner, uint256 give_good_id) returns()
func (_Givo *GivoTransactor) CycleFormed(opts *bind.TransactOpts, to *big.Int, get_good_id *big.Int, give_owner *big.Int, give_good_id *big.Int) (*types.Transaction, error) {
	return _Givo.contract.Transact(opts, "cycle_formed", to, get_good_id, give_owner, give_good_id)
}

// CycleFormed is a paid mutator transaction binding the contract method 0x43cef25e.
//
// Solidity: function cycle_formed(uint256 to, uint256 get_good_id, uint256 give_owner, uint256 give_good_id) returns()
func (_Givo *GivoSession) CycleFormed(to *big.Int, get_good_id *big.Int, give_owner *big.Int, give_good_id *big.Int) (*types.Transaction, error) {
	return _Givo.Contract.CycleFormed(&_Givo.TransactOpts, to, get_good_id, give_owner, give_good_id)
}

// CycleFormed is a paid mutator transaction binding the contract method 0x43cef25e.
//
// Solidity: function cycle_formed(uint256 to, uint256 get_good_id, uint256 give_owner, uint256 give_good_id) returns()
func (_Givo *GivoTransactorSession) CycleFormed(to *big.Int, get_good_id *big.Int, give_owner *big.Int, give_good_id *big.Int) (*types.Transaction, error) {
	return _Givo.Contract.CycleFormed(&_Givo.TransactOpts, to, get_good_id, give_owner, give_good_id)
}

// DeleteIntrest is a paid mutator transaction binding the contract method 0xa7fb194b.
//
// Solidity: function delete_intrest(uint256 node_id, uint256 good_id) returns()
func (_Givo *GivoTransactor) DeleteIntrest(opts *bind.TransactOpts, node_id *big.Int, good_id *big.Int) (*types.Transaction, error) {
	return _Givo.contract.Transact(opts, "delete_intrest", node_id, good_id)
}

// DeleteIntrest is a paid mutator transaction binding the contract method 0xa7fb194b.
//
// Solidity: function delete_intrest(uint256 node_id, uint256 good_id) returns()
func (_Givo *GivoSession) DeleteIntrest(node_id *big.Int, good_id *big.Int) (*types.Transaction, error) {
	return _Givo.Contract.DeleteIntrest(&_Givo.TransactOpts, node_id, good_id)
}

// DeleteIntrest is a paid mutator transaction binding the contract method 0xa7fb194b.
//
// Solidity: function delete_intrest(uint256 node_id, uint256 good_id) returns()
func (_Givo *GivoTransactorSession) DeleteIntrest(node_id *big.Int, good_id *big.Int) (*types.Transaction, error) {
	return _Givo.Contract.DeleteIntrest(&_Givo.TransactOpts, node_id, good_id)
}

// DeleteOffer is a paid mutator transaction binding the contract method 0x881d9637.
//
// Solidity: function delete_offer(uint256 good_id) returns()
func (_Givo *GivoTransactor) DeleteOffer(opts *bind.TransactOpts, good_id *big.Int) (*types.Transaction, error) {
	return _Givo.contract.Transact(opts, "delete_offer", good_id)
}

// DeleteOffer is a paid mutator transaction binding the contract method 0x881d9637.
//
// Solidity: function delete_offer(uint256 good_id) returns()
func (_Givo *GivoSession) DeleteOffer(good_id *big.Int) (*types.Transaction, error) {
	return _Givo.Contract.DeleteOffer(&_Givo.TransactOpts, good_id)
}

// DeleteOffer is a paid mutator transaction binding the contract method 0x881d9637.
//
// Solidity: function delete_offer(uint256 good_id) returns()
func (_Givo *GivoTransactorSession) DeleteOffer(good_id *big.Int) (*types.Transaction, error) {
	return _Givo.Contract.DeleteOffer(&_Givo.TransactOpts, good_id)
}

// ReferIntrest is a paid mutator transaction binding the contract method 0x394db574.
//
// Solidity: function refer_intrest(uint256 node_id, uint256 good_id, uint256 refer_id) returns()
func (_Givo *GivoTransactor) ReferIntrest(opts *bind.TransactOpts, node_id *big.Int, good_id *big.Int, refer_id *big.Int) (*types.Transaction, error) {
	return _Givo.contract.Transact(opts, "refer_intrest", node_id, good_id, refer_id)
}

// ReferIntrest is a paid mutator transaction binding the contract method 0x394db574.
//
// Solidity: function refer_intrest(uint256 node_id, uint256 good_id, uint256 refer_id) returns()
func (_Givo *GivoSession) ReferIntrest(node_id *big.Int, good_id *big.Int, refer_id *big.Int) (*types.Transaction, error) {
	return _Givo.Contract.ReferIntrest(&_Givo.TransactOpts, node_id, good_id, refer_id)
}

// ReferIntrest is a paid mutator transaction binding the contract method 0x394db574.
//
// Solidity: function refer_intrest(uint256 node_id, uint256 good_id, uint256 refer_id) returns()
func (_Givo *GivoTransactorSession) ReferIntrest(node_id *big.Int, good_id *big.Int, refer_id *big.Int) (*types.Transaction, error) {
	return _Givo.Contract.ReferIntrest(&_Givo.TransactOpts, node_id, good_id, refer_id)
}

// GivoChainedIterator is returned from FilterChained and is used to iterate over the raw logs and unpacked data for Chained events raised by the Givo contract.
type GivoChainedIterator struct {
	Event *GivoChained // Event containing the contract specifics and raw log

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
func (it *GivoChainedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GivoChained)
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
		it.Event = new(GivoChained)
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
func (it *GivoChainedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GivoChainedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GivoChained represents a Chained event raised by the Givo contract.
type GivoChained struct {
	To         *big.Int
	GetOwner   *big.Int
	GetGoodId  *big.Int
	GiveOwner  *big.Int
	GiveGoodId *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterChained is a free log retrieval operation binding the contract event 0x17d3c3f8c4f0fa8d1c6e1bb5edde9fda0537a5f56cafed38fea7b389ea911687.
//
// Solidity: event chained(uint256 to, uint256 get_owner, uint256 get_good_id, uint256 give_owner, uint256 give_good_id)
func (_Givo *GivoFilterer) FilterChained(opts *bind.FilterOpts) (*GivoChainedIterator, error) {

	logs, sub, err := _Givo.contract.FilterLogs(opts, "chained")
	if err != nil {
		return nil, err
	}
	return &GivoChainedIterator{contract: _Givo.contract, event: "chained", logs: logs, sub: sub}, nil
}

// WatchChained is a free log subscription operation binding the contract event 0x17d3c3f8c4f0fa8d1c6e1bb5edde9fda0537a5f56cafed38fea7b389ea911687.
//
// Solidity: event chained(uint256 to, uint256 get_owner, uint256 get_good_id, uint256 give_owner, uint256 give_good_id)
func (_Givo *GivoFilterer) WatchChained(opts *bind.WatchOpts, sink chan<- *GivoChained) (event.Subscription, error) {

	logs, sub, err := _Givo.contract.WatchLogs(opts, "chained")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GivoChained)
				if err := _Givo.contract.UnpackLog(event, "chained", log); err != nil {
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

// GivoDeletedIterator is returned from FilterDeleted and is used to iterate over the raw logs and unpacked data for Deleted events raised by the Givo contract.
type GivoDeletedIterator struct {
	Event *GivoDeleted // Event containing the contract specifics and raw log

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
func (it *GivoDeletedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GivoDeleted)
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
		it.Event = new(GivoDeleted)
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
func (it *GivoDeletedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GivoDeletedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GivoDeleted represents a Deleted event raised by the Givo contract.
type GivoDeleted struct {
	Owner  *big.Int
	GoodId *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDeleted is a free log retrieval operation binding the contract event 0xa7510b9f06bca625a2290e8871362d0ccc23a2b155fc82b80c66cf8c285c7323.
//
// Solidity: event deleted(uint256 owner, uint256 good_id)
func (_Givo *GivoFilterer) FilterDeleted(opts *bind.FilterOpts) (*GivoDeletedIterator, error) {

	logs, sub, err := _Givo.contract.FilterLogs(opts, "deleted")
	if err != nil {
		return nil, err
	}
	return &GivoDeletedIterator{contract: _Givo.contract, event: "deleted", logs: logs, sub: sub}, nil
}

// WatchDeleted is a free log subscription operation binding the contract event 0xa7510b9f06bca625a2290e8871362d0ccc23a2b155fc82b80c66cf8c285c7323.
//
// Solidity: event deleted(uint256 owner, uint256 good_id)
func (_Givo *GivoFilterer) WatchDeleted(opts *bind.WatchOpts, sink chan<- *GivoDeleted) (event.Subscription, error) {

	logs, sub, err := _Givo.contract.WatchLogs(opts, "deleted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GivoDeleted)
				if err := _Givo.contract.UnpackLog(event, "deleted", log); err != nil {
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

// GivoIntrestedIterator is returned from FilterIntrested and is used to iterate over the raw logs and unpacked data for Intrested events raised by the Givo contract.
type GivoIntrestedIterator struct {
	Event *GivoIntrested // Event containing the contract specifics and raw log

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
func (it *GivoIntrestedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GivoIntrested)
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
		it.Event = new(GivoIntrested)
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
func (it *GivoIntrestedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GivoIntrestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GivoIntrested represents a Intrested event raised by the Givo contract.
type GivoIntrested struct {
	To     *big.Int
	GoodId *big.Int
	From   *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterIntrested is a free log retrieval operation binding the contract event 0xcfeec790cb492f73558ea2f8bbac696dcb49ec548daa4ee42f15fe4908815976.
//
// Solidity: event intrested(uint256 to, uint256 good_id, uint256 from)
func (_Givo *GivoFilterer) FilterIntrested(opts *bind.FilterOpts) (*GivoIntrestedIterator, error) {

	logs, sub, err := _Givo.contract.FilterLogs(opts, "intrested")
	if err != nil {
		return nil, err
	}
	return &GivoIntrestedIterator{contract: _Givo.contract, event: "intrested", logs: logs, sub: sub}, nil
}

// WatchIntrested is a free log subscription operation binding the contract event 0xcfeec790cb492f73558ea2f8bbac696dcb49ec548daa4ee42f15fe4908815976.
//
// Solidity: event intrested(uint256 to, uint256 good_id, uint256 from)
func (_Givo *GivoFilterer) WatchIntrested(opts *bind.WatchOpts, sink chan<- *GivoIntrested) (event.Subscription, error) {

	logs, sub, err := _Givo.contract.WatchLogs(opts, "intrested")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GivoIntrested)
				if err := _Givo.contract.UnpackLog(event, "intrested", log); err != nil {
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

// GivoNotIntrestedIterator is returned from FilterNotIntrested and is used to iterate over the raw logs and unpacked data for NotIntrested events raised by the Givo contract.
type GivoNotIntrestedIterator struct {
	Event *GivoNotIntrested // Event containing the contract specifics and raw log

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
func (it *GivoNotIntrestedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GivoNotIntrested)
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
		it.Event = new(GivoNotIntrested)
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
func (it *GivoNotIntrestedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GivoNotIntrestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GivoNotIntrested represents a NotIntrested event raised by the Givo contract.
type GivoNotIntrested struct {
	To     *big.Int
	GoodId *big.Int
	From   *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterNotIntrested is a free log retrieval operation binding the contract event 0x9aaa23475fee2c5febd67beaa2b49f4104a85d040109fed3569b1cd73c45b84a.
//
// Solidity: event not_intrested(uint256 to, uint256 good_id, uint256 from)
func (_Givo *GivoFilterer) FilterNotIntrested(opts *bind.FilterOpts) (*GivoNotIntrestedIterator, error) {

	logs, sub, err := _Givo.contract.FilterLogs(opts, "not_intrested")
	if err != nil {
		return nil, err
	}
	return &GivoNotIntrestedIterator{contract: _Givo.contract, event: "not_intrested", logs: logs, sub: sub}, nil
}

// WatchNotIntrested is a free log subscription operation binding the contract event 0x9aaa23475fee2c5febd67beaa2b49f4104a85d040109fed3569b1cd73c45b84a.
//
// Solidity: event not_intrested(uint256 to, uint256 good_id, uint256 from)
func (_Givo *GivoFilterer) WatchNotIntrested(opts *bind.WatchOpts, sink chan<- *GivoNotIntrested) (event.Subscription, error) {

	logs, sub, err := _Givo.contract.WatchLogs(opts, "not_intrested")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GivoNotIntrested)
				if err := _Givo.contract.UnpackLog(event, "not_intrested", log); err != nil {
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

// GivoReferIterator is returned from FilterRefer and is used to iterate over the raw logs and unpacked data for Refer events raised by the Givo contract.
type GivoReferIterator struct {
	Event *GivoRefer // Event containing the contract specifics and raw log

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
func (it *GivoReferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GivoRefer)
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
		it.Event = new(GivoRefer)
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
func (it *GivoReferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GivoReferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GivoRefer represents a Refer event raised by the Givo contract.
type GivoRefer struct {
	To      *big.Int
	GoodId  *big.Int
	From    *big.Int
	ReferId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRefer is a free log retrieval operation binding the contract event 0xcc99f51473cc3bf43367db4f64d4efcdfd1acd4c70b4155cb7ffade4f2f6d362.
//
// Solidity: event refer(uint256 to, uint256 good_id, uint256 from, uint256 refer_id)
func (_Givo *GivoFilterer) FilterRefer(opts *bind.FilterOpts) (*GivoReferIterator, error) {

	logs, sub, err := _Givo.contract.FilterLogs(opts, "refer")
	if err != nil {
		return nil, err
	}
	return &GivoReferIterator{contract: _Givo.contract, event: "refer", logs: logs, sub: sub}, nil
}

// WatchRefer is a free log subscription operation binding the contract event 0xcc99f51473cc3bf43367db4f64d4efcdfd1acd4c70b4155cb7ffade4f2f6d362.
//
// Solidity: event refer(uint256 to, uint256 good_id, uint256 from, uint256 refer_id)
func (_Givo *GivoFilterer) WatchRefer(opts *bind.WatchOpts, sink chan<- *GivoRefer) (event.Subscription, error) {

	logs, sub, err := _Givo.contract.WatchLogs(opts, "refer")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GivoRefer)
				if err := _Givo.contract.UnpackLog(event, "refer", log); err != nil {
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

