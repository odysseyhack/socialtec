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
const GivoABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"if_owner\",\"type\":\"uint256\"},{\"name\":\"if_good\",\"type\":\"uint256\"},{\"name\":\"then_good\",\"type\":\"uint256\"},{\"name\":\"then_receiver\",\"type\":\"uint256\"}],\"name\":\"cycle_formed\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"offers\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"node_id\",\"type\":\"uint256\"}],\"name\":\"get_offers\",\"outputs\":[{\"name\":\"my_offers\",\"type\":\"uint256[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"address_to_id\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"node_cntr\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"good_id\",\"type\":\"uint256\"}],\"name\":\"delete_offer\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"good_id\",\"type\":\"uint256\"}],\"name\":\"add_intrest\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"good_id\",\"type\":\"uint256\"}],\"name\":\"delete_intrest\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"if_owner\",\"type\":\"uint256\"},{\"name\":\"if_good\",\"type\":\"uint256\"},{\"name\":\"if_receiver\",\"type\":\"uint256\"},{\"name\":\"then_good\",\"type\":\"uint256\"},{\"name\":\"then_receiver\",\"type\":\"uint256\"}],\"name\":\"cycle_propagate\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"name\",\"type\":\"string\"},{\"name\":\"ipfs_image\",\"type\":\"string\"},{\"name\":\"ipfs_details\",\"type\":\"string\"}],\"name\":\"create_offer\",\"outputs\":[{\"name\":\"good_id\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"max_all_good\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"good_id\",\"type\":\"uint256\"},{\"name\":\"refer_id\",\"type\":\"uint256\"}],\"name\":\"refer_intrest\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"get_id\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"all_goods\",\"outputs\":[{\"name\":\"node\",\"type\":\"uint256\"},{\"name\":\"index\",\"type\":\"uint256\"},{\"name\":\"available\",\"type\":\"bool\"},{\"name\":\"name\",\"type\":\"string\"},{\"name\":\"ipfs_image\",\"type\":\"string\"},{\"name\":\"ipfs_details\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"to\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"good_id\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"from\",\"type\":\"uint256\"}],\"name\":\"intrested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"to\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"good_id\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"from\",\"type\":\"uint256\"}],\"name\":\"not_intrested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"to\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"good_id\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"from\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"refer_id\",\"type\":\"uint256\"}],\"name\":\"refer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"owner\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"good_id\",\"type\":\"uint256\"}],\"name\":\"deleted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"if_owner\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"if_good\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"if_receiver\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"then_good\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"then_receiver\",\"type\":\"uint256\"}],\"name\":\"chained\",\"type\":\"event\"}]"

// GivoBin is the compiled bytecode used for deploying new contracts.
const GivoBin = `0x60806040526001600055600060015534801561001a57600080fd5b50610e658061002a6000396000f3fe608060405234801561001057600080fd5b50600436106100ea5760003560e01c8063b05538861161008c578063dee7afcf11610066578063dee7afcf14610413578063eeeb43791461041b578063f43fa8051461043e578063f9833cc614610446576100ea565b8063b05538861461020d578063bfb9c4531461022a578063c175cdd31461025f576100ea565b80637a52f30e116100c85780637a52f30e146101c257806386c12362146101e8578063881d9637146101f0578063a2eeced91461020d576100ea565b806343cef25e146100ef5780634d03a9a5146101205780634db5474c14610155575b600080fd5b61011e6004803603608081101561010557600080fd5b50803590602081013590604081013590606001356105c0565b005b6101436004803603604081101561013657600080fd5b5080359060200135610623565b60408051918252519081900360200190f35b6101726004803603602081101561016b57600080fd5b5035610651565b60408051602080825283518183015283519192839290830191858101910280838360005b838110156101ae578181015183820152602001610196565b505050509050019250505060405180910390f35b610143600480360360208110156101d857600080fd5b50356001600160a01b03166106b3565b6101436106c5565b61011e6004803603602081101561020657600080fd5b50356106cb565b61011e6004803603602081101561022357600080fd5b5035610830565b61011e600480360360a081101561024057600080fd5b50803590602081013590604081013590606081013590608001356108a3565b6101436004803603606081101561027557600080fd5b81019060208101813564010000000081111561029057600080fd5b8201836020820111156102a257600080fd5b803590602001918460018302840111640100000000831117156102c457600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929594936020810193503591505064010000000081111561031757600080fd5b82018360208201111561032957600080fd5b8035906020019184600183028401116401000000008311171561034b57600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929594936020810193503591505064010000000081111561039e57600080fd5b8201836020820111156103b057600080fd5b803590602001918460018302840111640100000000831117156103d257600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295506108f8945050505050565b610143610ac9565b61011e6004803603604081101561043157600080fd5b5080359060200135610acf565b610143610b4a565b6104636004803603602081101561045c57600080fd5b5035610b81565b6040518087815260200186815260200185151515158152602001806020018060200180602001848103845287818151815260200191508051906020019080838360005b838110156104be5781810151838201526020016104a6565b50505050905090810190601f1680156104eb5780820380516001836020036101000a031916815260200191505b50848103835286518152865160209182019188019080838360005b8381101561051e578181015183820152602001610506565b50505050905090810190601f16801561054b5780820380516001836020036101000a031916815260200191505b50848103825285518152855160209182019187019080838360005b8381101561057e578181015183820152602001610566565b50505050905090810190601f1680156105ab5780820380516001836020036101000a031916815260200191505b50995050505050505050505060405180910390f35b3360009081526003602090815260409182902054825187815291820186905281830152606081018490526080810183905290517f17d3c3f8c4f0fa8d1c6e1bb5edde9fda0537a5f56cafed38fea7b389ea9116879181900360a00190a150505050565b6002602052816000526040600020818154811061063c57fe5b90600052602060002001600091509150505481565b6000818152600260209081526040918290208054835181840281018401909452808452606093928301828280156106a757602002820191906000526020600020905b815481526020019060010190808311610693575b50939695505050505050565b60036020526000908152604090205481565b60005481565b3360009081526003602090815260408083205483526002909152812060048054919291849081106106f857fe5b90600052602060002090600602019050826001541161071657600080fd5b3360009081526003602052604090205481541461073257600080fd5b60028101805460ff1916905581548290600019810190811061075057fe5b90600052602060002001548282600101548154811061076b57fe5b9060005260206000200181905550600060048383600101548154811061078d57fe5b9060005260206000200154815481106107a257fe5b9060005260206000209060060201905081600101548160010181905550828054806107c957fe5b6000828152602080822083016000199081018390559092019092553382526003815260409182902054825190815290810186905281517fa7510b9f06bca625a2290e8871362d0ccc23a2b155fc82b80c66cf8c285c7323929181900390910190a150505050565b60006004828154811061083f57fe5b60009182526020808320600690920290910154338352600382526040928390205483518281529283018690528284015291519192507fcfeec790cb492f73558ea2f8bbac696dcb49ec548daa4ee42f15fe4908815976919081900360600190a15050565b6040805186815260208101869052808201859052606081018490526080810183905290517f17d3c3f8c4f0fa8d1c6e1bb5edde9fda0537a5f56cafed38fea7b389ea9116879181900360a00190a15050505050565b3360009081526003602052604081205480610926576000805433825260036020526040822081905560010190555b3360009081526003602090815260408083205483526002909152902061094a610d69565b506040805160c081018252336000908152600360209081528382205483528083018281526001948401858152606085018c8152608086018c905260a086018b9052600480548089018083559190965286517f8a35acfbc15ff81a39ae7d344fd709f28e8600b4aa8c65c6b64bfe7fe36bd19b600690970296870190815593517f8a35acfbc15ff81a39ae7d344fd709f28e8600b4aa8c65c6b64bfe7fe36bd19c87015591517f8a35acfbc15ff81a39ae7d344fd709f28e8600b4aa8c65c6b64bfe7fe36bd19d8601805460ff191691151591909117905551805195969591948794610a5b937f8a35acfbc15ff81a39ae7d344fd709f28e8600b4aa8c65c6b64bfe7fe36bd19e909201920190610da1565b5060808201518051610a77916004840191602090910190610da1565b5060a08201518051610a93916005840191602090910190610da1565b50508454600180820187556000968752602096879020949093039381018490559390940192909252908190559695505050505050565b60015481565b600060048381548110610ade57fe5b6000918252602080832060069092029091015433835260038252604092839020548351828152928301879052828401526060820185905291519192507fcc99f51473cc3bf43367db4f64d4efcdfd1acd4c70b4155cb7ffade4f2f6d362919081900360800190a1505050565b3360009081526003602052604081205480610b7b575060008054338252600360205260408220819055600181019091555b90505b90565b60048181548110610b8e57fe5b60009182526020918290206006919091020180546001808301546002808501546003860180546040805161010097831615979097026000190190911693909304601f8101899004890286018901909352828552949750919560ff909216949391830182828015610c3f5780601f10610c1457610100808354040283529160200191610c3f565b820191906000526020600020905b815481529060010190602001808311610c2257829003601f168201915b5050505060048301805460408051602060026001851615610100026000190190941693909304601f8101849004840282018401909252818152949594935090830182828015610ccf5780601f10610ca457610100808354040283529160200191610ccf565b820191906000526020600020905b815481529060010190602001808311610cb257829003601f168201915b5050505060058301805460408051602060026001851615610100026000190190941693909304601f8101849004840282018401909252818152949594935090830182828015610d5f5780601f10610d3457610100808354040283529160200191610d5f565b820191906000526020600020905b815481529060010190602001808311610d4257829003601f168201915b5050505050905086565b6040518060c0016040528060008152602001600081526020016000151581526020016060815260200160608152602001606081525090565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10610de257805160ff1916838001178555610e0f565b82800160010185558215610e0f579182015b82811115610e0f578251825591602001919060010190610df4565b50610e1b929150610e1f565b5090565b610b7e91905b80821115610e1b5760008155600101610e2556fea165627a7a72305820900a59f57c1ad1bcd5d94c89c5bdff68a507db5e8899925f403dd17d65f173cc0029`

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

// AddressToId is a free data retrieval call binding the contract method 0x7a52f30e.
//
// Solidity: function address_to_id(address ) constant returns(uint256)
func (_Givo *GivoCaller) AddressToId(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Givo.contract.Call(opts, out, "address_to_id", arg0)
	return *ret0, err
}

// AddressToId is a free data retrieval call binding the contract method 0x7a52f30e.
//
// Solidity: function address_to_id(address ) constant returns(uint256)
func (_Givo *GivoSession) AddressToId(arg0 common.Address) (*big.Int, error) {
	return _Givo.Contract.AddressToId(&_Givo.CallOpts, arg0)
}

// AddressToId is a free data retrieval call binding the contract method 0x7a52f30e.
//
// Solidity: function address_to_id(address ) constant returns(uint256)
func (_Givo *GivoCallerSession) AddressToId(arg0 common.Address) (*big.Int, error) {
	return _Givo.Contract.AddressToId(&_Givo.CallOpts, arg0)
}

// AllGoods is a free data retrieval call binding the contract method 0xf9833cc6.
//
// Solidity: function all_goods(uint256 ) constant returns(uint256 node, uint256 index, bool available, string name, string ipfs_image, string ipfs_details)
func (_Givo *GivoCaller) AllGoods(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Node        *big.Int
	Index       *big.Int
	Available   bool
	Name        string
	IpfsImage   string
	IpfsDetails string
}, error) {
	ret := new(struct {
		Node        *big.Int
		Index       *big.Int
		Available   bool
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
// Solidity: function all_goods(uint256 ) constant returns(uint256 node, uint256 index, bool available, string name, string ipfs_image, string ipfs_details)
func (_Givo *GivoSession) AllGoods(arg0 *big.Int) (struct {
	Node        *big.Int
	Index       *big.Int
	Available   bool
	Name        string
	IpfsImage   string
	IpfsDetails string
}, error) {
	return _Givo.Contract.AllGoods(&_Givo.CallOpts, arg0)
}

// AllGoods is a free data retrieval call binding the contract method 0xf9833cc6.
//
// Solidity: function all_goods(uint256 ) constant returns(uint256 node, uint256 index, bool available, string name, string ipfs_image, string ipfs_details)
func (_Givo *GivoCallerSession) AllGoods(arg0 *big.Int) (struct {
	Node        *big.Int
	Index       *big.Int
	Available   bool
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

// MaxAllGood is a free data retrieval call binding the contract method 0xdee7afcf.
//
// Solidity: function max_all_good() constant returns(uint256)
func (_Givo *GivoCaller) MaxAllGood(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Givo.contract.Call(opts, out, "max_all_good")
	return *ret0, err
}

// MaxAllGood is a free data retrieval call binding the contract method 0xdee7afcf.
//
// Solidity: function max_all_good() constant returns(uint256)
func (_Givo *GivoSession) MaxAllGood() (*big.Int, error) {
	return _Givo.Contract.MaxAllGood(&_Givo.CallOpts)
}

// MaxAllGood is a free data retrieval call binding the contract method 0xdee7afcf.
//
// Solidity: function max_all_good() constant returns(uint256)
func (_Givo *GivoCallerSession) MaxAllGood() (*big.Int, error) {
	return _Givo.Contract.MaxAllGood(&_Givo.CallOpts)
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
// Solidity: function cycle_formed(uint256 if_owner, uint256 if_good, uint256 then_good, uint256 then_receiver) returns()
func (_Givo *GivoTransactor) CycleFormed(opts *bind.TransactOpts, if_owner *big.Int, if_good *big.Int, then_good *big.Int, then_receiver *big.Int) (*types.Transaction, error) {
	return _Givo.contract.Transact(opts, "cycle_formed", if_owner, if_good, then_good, then_receiver)
}

// CycleFormed is a paid mutator transaction binding the contract method 0x43cef25e.
//
// Solidity: function cycle_formed(uint256 if_owner, uint256 if_good, uint256 then_good, uint256 then_receiver) returns()
func (_Givo *GivoSession) CycleFormed(if_owner *big.Int, if_good *big.Int, then_good *big.Int, then_receiver *big.Int) (*types.Transaction, error) {
	return _Givo.Contract.CycleFormed(&_Givo.TransactOpts, if_owner, if_good, then_good, then_receiver)
}

// CycleFormed is a paid mutator transaction binding the contract method 0x43cef25e.
//
// Solidity: function cycle_formed(uint256 if_owner, uint256 if_good, uint256 then_good, uint256 then_receiver) returns()
func (_Givo *GivoTransactorSession) CycleFormed(if_owner *big.Int, if_good *big.Int, then_good *big.Int, then_receiver *big.Int) (*types.Transaction, error) {
	return _Givo.Contract.CycleFormed(&_Givo.TransactOpts, if_owner, if_good, then_good, then_receiver)
}

// CyclePropagate is a paid mutator transaction binding the contract method 0xbfb9c453.
//
// Solidity: function cycle_propagate(uint256 if_owner, uint256 if_good, uint256 if_receiver, uint256 then_good, uint256 then_receiver) returns()
func (_Givo *GivoTransactor) CyclePropagate(opts *bind.TransactOpts, if_owner *big.Int, if_good *big.Int, if_receiver *big.Int, then_good *big.Int, then_receiver *big.Int) (*types.Transaction, error) {
	return _Givo.contract.Transact(opts, "cycle_propagate", if_owner, if_good, if_receiver, then_good, then_receiver)
}

// CyclePropagate is a paid mutator transaction binding the contract method 0xbfb9c453.
//
// Solidity: function cycle_propagate(uint256 if_owner, uint256 if_good, uint256 if_receiver, uint256 then_good, uint256 then_receiver) returns()
func (_Givo *GivoSession) CyclePropagate(if_owner *big.Int, if_good *big.Int, if_receiver *big.Int, then_good *big.Int, then_receiver *big.Int) (*types.Transaction, error) {
	return _Givo.Contract.CyclePropagate(&_Givo.TransactOpts, if_owner, if_good, if_receiver, then_good, then_receiver)
}

// CyclePropagate is a paid mutator transaction binding the contract method 0xbfb9c453.
//
// Solidity: function cycle_propagate(uint256 if_owner, uint256 if_good, uint256 if_receiver, uint256 then_good, uint256 then_receiver) returns()
func (_Givo *GivoTransactorSession) CyclePropagate(if_owner *big.Int, if_good *big.Int, if_receiver *big.Int, then_good *big.Int, then_receiver *big.Int) (*types.Transaction, error) {
	return _Givo.Contract.CyclePropagate(&_Givo.TransactOpts, if_owner, if_good, if_receiver, then_good, then_receiver)
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

// GetId is a paid mutator transaction binding the contract method 0xf43fa805.
//
// Solidity: function get_id() returns(uint256)
func (_Givo *GivoTransactor) GetId(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Givo.contract.Transact(opts, "get_id")
}

// GetId is a paid mutator transaction binding the contract method 0xf43fa805.
//
// Solidity: function get_id() returns(uint256)
func (_Givo *GivoSession) GetId() (*types.Transaction, error) {
	return _Givo.Contract.GetId(&_Givo.TransactOpts)
}

// GetId is a paid mutator transaction binding the contract method 0xf43fa805.
//
// Solidity: function get_id() returns(uint256)
func (_Givo *GivoTransactorSession) GetId() (*types.Transaction, error) {
	return _Givo.Contract.GetId(&_Givo.TransactOpts)
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
	IfOwner      *big.Int
	IfGood       *big.Int
	IfReceiver   *big.Int
	ThenGood     *big.Int
	ThenReceiver *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterChained is a free log retrieval operation binding the contract event 0x17d3c3f8c4f0fa8d1c6e1bb5edde9fda0537a5f56cafed38fea7b389ea911687.
//
// Solidity: event chained(uint256 if_owner, uint256 if_good, uint256 if_receiver, uint256 then_good, uint256 then_receiver)
func (_Givo *GivoFilterer) FilterChained(opts *bind.FilterOpts) (*GivoChainedIterator, error) {

	logs, sub, err := _Givo.contract.FilterLogs(opts, "chained")
	if err != nil {
		return nil, err
	}
	return &GivoChainedIterator{contract: _Givo.contract, event: "chained", logs: logs, sub: sub}, nil
}

// WatchChained is a free log subscription operation binding the contract event 0x17d3c3f8c4f0fa8d1c6e1bb5edde9fda0537a5f56cafed38fea7b389ea911687.
//
// Solidity: event chained(uint256 if_owner, uint256 if_good, uint256 if_receiver, uint256 then_good, uint256 then_receiver)
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

