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
const GivoABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"to\",\"type\":\"uint256\"},{\"name\":\"get_good_id\",\"type\":\"uint256\"},{\"name\":\"give_owner\",\"type\":\"uint256\"},{\"name\":\"give_good_id\",\"type\":\"uint256\"}],\"name\":\"cycle_formed\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"offers\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"node_id\",\"type\":\"uint256\"}],\"name\":\"get_offers\",\"outputs\":[{\"name\":\"my_offers\",\"type\":\"uint256[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"node_cntr\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"good_id\",\"type\":\"uint256\"}],\"name\":\"delete_offer\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"good_id\",\"type\":\"uint256\"}],\"name\":\"add_intrest\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"good_id\",\"type\":\"uint256\"}],\"name\":\"delete_intrest\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"name\",\"type\":\"string\"},{\"name\":\"ipfs_image\",\"type\":\"string\"},{\"name\":\"ipfs_details\",\"type\":\"string\"}],\"name\":\"create_offer\",\"outputs\":[{\"name\":\"good_id\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"good_id\",\"type\":\"uint256\"},{\"name\":\"refer_id\",\"type\":\"uint256\"}],\"name\":\"refer_intrest\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"all_goods\",\"outputs\":[{\"name\":\"node\",\"type\":\"uint256\"},{\"name\":\"name\",\"type\":\"string\"},{\"name\":\"ipfs_image\",\"type\":\"string\"},{\"name\":\"ipfs_details\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"to\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"good_id\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"from\",\"type\":\"uint256\"}],\"name\":\"intrested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"to\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"good_id\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"from\",\"type\":\"uint256\"}],\"name\":\"not_intrested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"to\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"good_id\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"from\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"refer_id\",\"type\":\"uint256\"}],\"name\":\"refer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"owner\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"good_id\",\"type\":\"uint256\"}],\"name\":\"deleted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"to\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"get_owner\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"get_good_id\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"give_owner\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"give_good_id\",\"type\":\"uint256\"}],\"name\":\"chained\",\"type\":\"event\"}]"

// GivoBin is the compiled bytecode used for deploying new contracts.
const GivoBin = `0x60806040526000805534801561001457600080fd5b50610bd9806100246000396000f3fe608060405234801561001057600080fd5b506004361061009e5760003560e01c8063a2eeced911610066578063a2eeced91461019b578063b05538861461019b578063c175cdd3146101b8578063eeeb43791461036c578063f9833cc61461038f5761009e565b806343cef25e146100a35780634d03a9a5146100d45780634db5474c1461010957806386c1236214610176578063881d96371461017e575b600080fd5b6100d2600480360360808110156100b957600080fd5b50803590602081013590604081013590606001356104f7565b005b6100f7600480360360408110156100ea57600080fd5b508035906020013561055a565b60408051918252519081900360200190f35b6101266004803603602081101561011f57600080fd5b5035610588565b60408051602080825283518183015283519192839290830191858101910280838360005b8381101561016257818101518382015260200161014a565b505050509050019250505060405180910390f35b6100f76105ea565b6100d26004803603602081101561019457600080fd5b50356105f0565b6100d2600480360360208110156101b157600080fd5b50356106c0565b6100f7600480360360608110156101ce57600080fd5b8101906020810181356401000000008111156101e957600080fd5b8201836020820111156101fb57600080fd5b8035906020019184600183028401116401000000008311171561021d57600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929594936020810193503591505064010000000081111561027057600080fd5b82018360208201111561028257600080fd5b803590602001918460018302840111640100000000831117156102a457600080fd5b91908080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525092959493602081019350359150506401000000008111156102f757600080fd5b82018360208201111561030957600080fd5b8035906020019184600183028401116401000000008311171561032b57600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550610733945050505050565b6100d26004803603604081101561038257600080fd5b5080359060200135610884565b6103ac600480360360208110156103a557600080fd5b50356108ff565b60405180858152602001806020018060200180602001848103845287818151815260200191508051906020019080838360005b838110156103f75781810151838201526020016103df565b50505050905090810190601f1680156104245780820380516001836020036101000a031916815260200191505b50848103835286518152865160209182019188019080838360005b8381101561045757818101518382015260200161043f565b50505050905090810190601f1680156104845780820380516001836020036101000a031916815260200191505b50848103825285518152855160209182019187019080838360005b838110156104b757818101518382015260200161049f565b50505050905090810190601f1680156104e45780820380516001836020036101000a031916815260200191505b5097505050505050505060405180910390f35b3360009081526002602090815260409182902054825187815291820152808201859052606081018490526080810183905290517f17d3c3f8c4f0fa8d1c6e1bb5edde9fda0537a5f56cafed38fea7b389ea9116879181900360a00190a150505050565b6001602052816000526040600020818154811061057357fe5b90600052602060002001600091509150505481565b6000818152600160209081526040918290208054835181840281018401909452808452606093928301828280156105de57602002820191906000526020600020905b8154815260200190600101908083116105ca575b50939695505050505050565b60005481565b33600090815260026020908152604080832054835260019091529020805482111561061a57600080fd5b80548190600019810190811061062c57fe5b906000526020600020015481838154811061064357fe5b90600052602060002001819055508080548061065b57fe5b6000828152602080822083016000199081018390559092019092553382526002815260409182902054825190815290810184905281517fa7510b9f06bca625a2290e8871362d0ccc23a2b155fc82b80c66cf8c285c7323929181900390910190a15050565b6000600382815481106106cf57fe5b60009182526020808320600490920290910154338352600282526040928390205483518281529283018690528284015291519192507fcfeec790cb492f73558ea2f8bbac696dcb49ec548daa4ee42f15fe4908815976919081900360600190a15050565b33600090815260026020908152604080832054835260019091528120805461076e576000805433825260026020526040822081905560010190555b610776610aea565b506040805160808101825233600090815260026020908152838220548352808301898152938301889052606083018790526003805460018082018084559290945284517fc2575a0e9e593c00f959f8c92f12db2869c3395a3b0502d05e2516446f71f85b60049092029182019081559551805195969495929487949093610823937fc2575a0e9e593c00f959f8c92f12db2869c3395a3b0502d05e2516446f71f85c019290910190610b12565b506040820151805161083f916002840191602090910190610b12565b506060820151805161085b916003840191602090910190610b12565b505084546001810186556000958652602090952092909103919093018190559695505050505050565b60006003838154811061089357fe5b6000918252602080832060049092029091015433835260028252604092839020548351828152928301879052828401526060820185905291519192507fcc99f51473cc3bf43367db4f64d4efcdfd1acd4c70b4155cb7ffade4f2f6d362919081900360800190a1505050565b6003818154811061090c57fe5b9060005260206000209060040201600091509050806000015490806001018054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156109be5780601f10610993576101008083540402835291602001916109be565b820191906000526020600020905b8154815290600101906020018083116109a157829003601f168201915b50505060028085018054604080516020601f6000196101006001871615020190941695909504928301859004850281018501909152818152959695945090925090830182828015610a505780601f10610a2557610100808354040283529160200191610a50565b820191906000526020600020905b815481529060010190602001808311610a3357829003601f168201915b5050505060038301805460408051602060026001851615610100026000190190941693909304601f8101849004840282018401909252818152949594935090830182828015610ae05780601f10610ab557610100808354040283529160200191610ae0565b820191906000526020600020905b815481529060010190602001808311610ac357829003601f168201915b5050505050905084565b6040518060800160405280600081526020016060815260200160608152602001606081525090565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10610b5357805160ff1916838001178555610b80565b82800160010185558215610b80579182015b82811115610b80578251825591602001919060010190610b65565b50610b8c929150610b90565b5090565b610baa91905b80821115610b8c5760008155600101610b96565b9056fea165627a7a72305820d332364343708b70036d7ce13fbc885b6e42905fdf9475493d64ff222f8650730029`

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

// AllGoods is a free data retrieval call binding the contract method 0xf9833cc6.
//
// Solidity: function all_goods(uint256 ) constant returns(uint256 node, string name, string ipfs_image, string ipfs_details)
func (_Givo *GivoCaller) AllGoods(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Node        *big.Int
	Name        string
	IpfsImage   string
	IpfsDetails string
}, error) {
	ret := new(struct {
		Node        *big.Int
		Name        string
		IpfsImage   string
		IpfsDetails string
	})
	out := ret
	err := _Givo.contract.Call(opts, out, "all_goods", arg0)
	return *ret, err
}

// AllGoods is a free data retrieval call binding the contract method 0xf9833cc6.
//
// Solidity: function all_goods(uint256 ) constant returns(uint256 node, string name, string ipfs_image, string ipfs_details)
func (_Givo *GivoSession) AllGoods(arg0 *big.Int) (struct {
	Node        *big.Int
	Name        string
	IpfsImage   string
	IpfsDetails string
}, error) {
	return _Givo.Contract.AllGoods(&_Givo.CallOpts, arg0)
}

// AllGoods is a free data retrieval call binding the contract method 0xf9833cc6.
//
// Solidity: function all_goods(uint256 ) constant returns(uint256 node, string name, string ipfs_image, string ipfs_details)
func (_Givo *GivoCallerSession) AllGoods(arg0 *big.Int) (struct {
	Node        *big.Int
	Name        string
	IpfsImage   string
	IpfsDetails string
}, error) {
	return _Givo.Contract.AllGoods(&_Givo.CallOpts, arg0)
}

// GetOffers is a free data retrieval call binding the contract method 0x4db5474c.
//
// Solidity: function get_offers(uint256 node_id) constant returns(uint256[] my_offers)
func (_Givo *GivoCaller) GetOffers(opts *bind.CallOpts, node_id *big.Int) ([]*big.Int, error) {
	var (
		ret0 = new([]*big.Int)
	)
	out := ret0
	err := _Givo.contract.Call(opts, out, "get_offers", node_id)
	return *ret0, err
}

// GetOffers is a free data retrieval call binding the contract method 0x4db5474c.
//
// Solidity: function get_offers(uint256 node_id) constant returns(uint256[] my_offers)
func (_Givo *GivoSession) GetOffers(node_id *big.Int) ([]*big.Int, error) {
	return _Givo.Contract.GetOffers(&_Givo.CallOpts, node_id)
}

// GetOffers is a free data retrieval call binding the contract method 0x4db5474c.
//
// Solidity: function get_offers(uint256 node_id) constant returns(uint256[] my_offers)
func (_Givo *GivoCallerSession) GetOffers(node_id *big.Int) ([]*big.Int, error) {
	return _Givo.Contract.GetOffers(&_Givo.CallOpts, node_id)
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
// Solidity: function offers(uint256 , uint256 ) constant returns(uint256)
func (_Givo *GivoCaller) Offers(opts *bind.CallOpts, arg0 *big.Int, arg1 *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Givo.contract.Call(opts, out, "offers", arg0, arg1)
	return *ret0, err
}

// Offers is a free data retrieval call binding the contract method 0x4d03a9a5.
//
// Solidity: function offers(uint256 , uint256 ) constant returns(uint256)
func (_Givo *GivoSession) Offers(arg0 *big.Int, arg1 *big.Int) (*big.Int, error) {
	return _Givo.Contract.Offers(&_Givo.CallOpts, arg0, arg1)
}

// Offers is a free data retrieval call binding the contract method 0x4d03a9a5.
//
// Solidity: function offers(uint256 , uint256 ) constant returns(uint256)
func (_Givo *GivoCallerSession) Offers(arg0 *big.Int, arg1 *big.Int) (*big.Int, error) {
	return _Givo.Contract.Offers(&_Givo.CallOpts, arg0, arg1)
}

// AddIntrest is a paid mutator transaction binding the contract method 0xa2eeced9.
//
// Solidity: function add_intrest(uint256 good_id) returns()
func (_Givo *GivoTransactor) AddIntrest(opts *bind.TransactOpts, good_id *big.Int) (*types.Transaction, error) {
	return _Givo.contract.Transact(opts, "add_intrest", good_id)
}

// AddIntrest is a paid mutator transaction binding the contract method 0xa2eeced9.
//
// Solidity: function add_intrest(uint256 good_id) returns()
func (_Givo *GivoSession) AddIntrest(good_id *big.Int) (*types.Transaction, error) {
	return _Givo.Contract.AddIntrest(&_Givo.TransactOpts, good_id)
}

// AddIntrest is a paid mutator transaction binding the contract method 0xa2eeced9.
//
// Solidity: function add_intrest(uint256 good_id) returns()
func (_Givo *GivoTransactorSession) AddIntrest(good_id *big.Int) (*types.Transaction, error) {
	return _Givo.Contract.AddIntrest(&_Givo.TransactOpts, good_id)
}

// CreateOffer is a paid mutator transaction binding the contract method 0xc175cdd3.
//
// Solidity: function create_offer(string name, string ipfs_image, string ipfs_details) returns(uint256 good_id)
func (_Givo *GivoTransactor) CreateOffer(opts *bind.TransactOpts, name string, ipfs_image string, ipfs_details string) (*types.Transaction, error) {
	return _Givo.contract.Transact(opts, "create_offer", name, ipfs_image, ipfs_details)
}

// CreateOffer is a paid mutator transaction binding the contract method 0xc175cdd3.
//
// Solidity: function create_offer(string name, string ipfs_image, string ipfs_details) returns(uint256 good_id)
func (_Givo *GivoSession) CreateOffer(name string, ipfs_image string, ipfs_details string) (*types.Transaction, error) {
	return _Givo.Contract.CreateOffer(&_Givo.TransactOpts, name, ipfs_image, ipfs_details)
}

// CreateOffer is a paid mutator transaction binding the contract method 0xc175cdd3.
//
// Solidity: function create_offer(string name, string ipfs_image, string ipfs_details) returns(uint256 good_id)
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

// DeleteIntrest is a paid mutator transaction binding the contract method 0xb0553886.
//
// Solidity: function delete_intrest(uint256 good_id) returns()
func (_Givo *GivoTransactor) DeleteIntrest(opts *bind.TransactOpts, good_id *big.Int) (*types.Transaction, error) {
	return _Givo.contract.Transact(opts, "delete_intrest", good_id)
}

// DeleteIntrest is a paid mutator transaction binding the contract method 0xb0553886.
//
// Solidity: function delete_intrest(uint256 good_id) returns()
func (_Givo *GivoSession) DeleteIntrest(good_id *big.Int) (*types.Transaction, error) {
	return _Givo.Contract.DeleteIntrest(&_Givo.TransactOpts, good_id)
}

// DeleteIntrest is a paid mutator transaction binding the contract method 0xb0553886.
//
// Solidity: function delete_intrest(uint256 good_id) returns()
func (_Givo *GivoTransactorSession) DeleteIntrest(good_id *big.Int) (*types.Transaction, error) {
	return _Givo.Contract.DeleteIntrest(&_Givo.TransactOpts, good_id)
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

// ReferIntrest is a paid mutator transaction binding the contract method 0xeeeb4379.
//
// Solidity: function refer_intrest(uint256 good_id, uint256 refer_id) returns()
func (_Givo *GivoTransactor) ReferIntrest(opts *bind.TransactOpts, good_id *big.Int, refer_id *big.Int) (*types.Transaction, error) {
	return _Givo.contract.Transact(opts, "refer_intrest", good_id, refer_id)
}

// ReferIntrest is a paid mutator transaction binding the contract method 0xeeeb4379.
//
// Solidity: function refer_intrest(uint256 good_id, uint256 refer_id) returns()
func (_Givo *GivoSession) ReferIntrest(good_id *big.Int, refer_id *big.Int) (*types.Transaction, error) {
	return _Givo.Contract.ReferIntrest(&_Givo.TransactOpts, good_id, refer_id)
}

// ReferIntrest is a paid mutator transaction binding the contract method 0xeeeb4379.
//
// Solidity: function refer_intrest(uint256 good_id, uint256 refer_id) returns()
func (_Givo *GivoTransactorSession) ReferIntrest(good_id *big.Int, refer_id *big.Int) (*types.Transaction, error) {
	return _Givo.Contract.ReferIntrest(&_Givo.TransactOpts, good_id, refer_id)
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

