// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

import (
	"encoding/json"

	"github.com/ethereum-optimism/optimism/op-bindings/solc"
)

const SystemConfigStorageLayoutJSON = "{\"storage\":[{\"astId\":29270,\"contract\":\"contracts/L1/SystemConfig.sol:SystemConfig\",\"label\":\"_initialized\",\"offset\":0,\"slot\":\"0\",\"type\":\"t_uint8\"},{\"astId\":29273,\"contract\":\"contracts/L1/SystemConfig.sol:SystemConfig\",\"label\":\"_initializing\",\"offset\":1,\"slot\":\"0\",\"type\":\"t_bool\"},{\"astId\":29884,\"contract\":\"contracts/L1/SystemConfig.sol:SystemConfig\",\"label\":\"__gap\",\"offset\":0,\"slot\":\"1\",\"type\":\"t_array(t_uint256)50_storage\"},{\"astId\":29142,\"contract\":\"contracts/L1/SystemConfig.sol:SystemConfig\",\"label\":\"_owner\",\"offset\":0,\"slot\":\"51\",\"type\":\"t_address\"},{\"astId\":29262,\"contract\":\"contracts/L1/SystemConfig.sol:SystemConfig\",\"label\":\"__gap\",\"offset\":0,\"slot\":\"52\",\"type\":\"t_array(t_uint256)49_storage\"},{\"astId\":1901,\"contract\":\"contracts/L1/SystemConfig.sol:SystemConfig\",\"label\":\"overhead\",\"offset\":0,\"slot\":\"101\",\"type\":\"t_uint256\"},{\"astId\":1904,\"contract\":\"contracts/L1/SystemConfig.sol:SystemConfig\",\"label\":\"scalar\",\"offset\":0,\"slot\":\"102\",\"type\":\"t_uint256\"},{\"astId\":1907,\"contract\":\"contracts/L1/SystemConfig.sol:SystemConfig\",\"label\":\"batcherHash\",\"offset\":0,\"slot\":\"103\",\"type\":\"t_bytes32\"},{\"astId\":1910,\"contract\":\"contracts/L1/SystemConfig.sol:SystemConfig\",\"label\":\"gasLimit\",\"offset\":0,\"slot\":\"104\",\"type\":\"t_uint64\"}],\"types\":{\"t_address\":{\"encoding\":\"inplace\",\"label\":\"address\",\"numberOfBytes\":\"20\"},\"t_array(t_uint256)49_storage\":{\"encoding\":\"inplace\",\"label\":\"uint256[49]\",\"numberOfBytes\":\"1568\"},\"t_array(t_uint256)50_storage\":{\"encoding\":\"inplace\",\"label\":\"uint256[50]\",\"numberOfBytes\":\"1600\"},\"t_bool\":{\"encoding\":\"inplace\",\"label\":\"bool\",\"numberOfBytes\":\"1\"},\"t_bytes32\":{\"encoding\":\"inplace\",\"label\":\"bytes32\",\"numberOfBytes\":\"32\"},\"t_uint256\":{\"encoding\":\"inplace\",\"label\":\"uint256\",\"numberOfBytes\":\"32\"},\"t_uint64\":{\"encoding\":\"inplace\",\"label\":\"uint64\",\"numberOfBytes\":\"8\"},\"t_uint8\":{\"encoding\":\"inplace\",\"label\":\"uint8\",\"numberOfBytes\":\"1\"}}}"

var SystemConfigStorageLayout = new(solc.StorageLayout)

var SystemConfigDeployedBin = "0x608060405234801561001057600080fd5b50600436106100ea5760003560e01c8063b40a817c1161008c578063f2fde38b11610066578063f2fde38b146101ca578063f45e65d8146101dd578063f68016b7146101e6578063ffa1ad74146101fa57600080fd5b8063b40a817c1461019b578063c9b26f61146101ae578063e81b2c6d146101c157600080fd5b806370bde19c116100c857806370bde19c14610143578063715018a6146101585780638da5cb5b14610160578063935f029e1461018857600080fd5b80630c18c162146100ef57806329477e861461010b57806354fd4d501461012e575b600080fd5b6100f860655481565b6040519081526020015b60405180910390f35b610115627a120081565b60405167ffffffffffffffff9091168152602001610102565b610136610202565b6040516101029190610bb3565b610156610151366004610c0e565b6102a5565b005b610156610504565b60335460405173ffffffffffffffffffffffffffffffffffffffff9091168152602001610102565b610156610196366004610c5e565b610518565b6101566101a9366004610c80565b6105b1565b6101566101bc366004610c9b565b6106c2565b6100f860675481565b6101566101d8366004610cb4565b6106f2565b6100f860665481565b6068546101159067ffffffffffffffff1681565b6100f8600081565b606061022d7f00000000000000000000000000000000000000000000000000000000000000006107c5565b6102567f00000000000000000000000000000000000000000000000000000000000000006107c5565b61027f7f00000000000000000000000000000000000000000000000000000000000000006107c5565b60405160200161029193929190610ccf565b604051602081830303815290604052905090565b600054610100900460ff16158080156102c55750600054600160ff909116105b806102df5750303b1580156102df575060005460ff166001145b610370576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201527f647920696e697469616c697a656400000000000000000000000000000000000060648201526084015b60405180910390fd5b600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016600117905580156103ce57600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff166101001790555b627a120067ffffffffffffffff83161015610445576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601f60248201527f53797374656d436f6e6669673a20676173206c696d697420746f6f206c6f77006044820152606401610367565b61044d610902565b610456866106f2565b606585905560668490556067839055606880547fffffffffffffffffffffffffffffffffffffffffffffffff00000000000000001667ffffffffffffffff841617905580156104fc57600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b505050505050565b61050c6109a1565b6105166000610a22565b565b6105206109a1565b606582905560668190556040805160208101849052908101829052600090606001604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529190529050600160007f1d2b0bda21d56b8bd12d4f94ebacffdfb35f5e226f84b461103bb8beab6353be836040516105a49190610bb3565b60405180910390a3505050565b6105b96109a1565b627a120067ffffffffffffffff82161015610630576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601f60248201527f53797374656d436f6e6669673a20676173206c696d697420746f6f206c6f77006044820152606401610367565b606880547fffffffffffffffffffffffffffffffffffffffffffffffff00000000000000001667ffffffffffffffff831690811790915560408051602080820193909352815180820390930183528101905260025b60007f1d2b0bda21d56b8bd12d4f94ebacffdfb35f5e226f84b461103bb8beab6353be836040516106b69190610bb3565b60405180910390a35050565b6106ca6109a1565b6067819055604080516020808201849052825180830390910181529082019091526000610685565b6106fa6109a1565b73ffffffffffffffffffffffffffffffffffffffff811661079d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201527f64647265737300000000000000000000000000000000000000000000000000006064820152608401610367565b6107a681610a22565b50565b73ffffffffffffffffffffffffffffffffffffffff163b151590565b60608160000361080857505060408051808201909152600181527f3000000000000000000000000000000000000000000000000000000000000000602082015290565b8160005b8115610832578061081c81610d74565b915061082b9050600a83610ddb565b915061080c565b60008167ffffffffffffffff81111561084d5761084d610def565b6040519080825280601f01601f191660200182016040528015610877576020820181803683370190505b5090505b84156108fa5761088c600183610e1e565b9150610899600a86610e35565b6108a4906030610e49565b60f81b8183815181106108b9576108b9610e61565b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a9053506108f3600a86610ddb565b945061087b565b949350505050565b600054610100900460ff16610999576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e670000000000000000000000000000000000000000006064820152608401610367565b610516610a99565b60335473ffffffffffffffffffffffffffffffffffffffff163314610516576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401610367565b6033805473ffffffffffffffffffffffffffffffffffffffff8381167fffffffffffffffffffffffff0000000000000000000000000000000000000000831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b600054610100900460ff16610b30576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e670000000000000000000000000000000000000000006064820152608401610367565b61051633610a22565b60005b83811015610b54578181015183820152602001610b3c565b83811115610b63576000848401525b50505050565b60008151808452610b81816020860160208601610b39565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b602081526000610bc66020830184610b69565b9392505050565b803573ffffffffffffffffffffffffffffffffffffffff81168114610bf157600080fd5b919050565b803567ffffffffffffffff81168114610bf157600080fd5b600080600080600060a08688031215610c2657600080fd5b610c2f86610bcd565b9450602086013593506040860135925060608601359150610c5260808701610bf6565b90509295509295909350565b60008060408385031215610c7157600080fd5b50508035926020909101359150565b600060208284031215610c9257600080fd5b610bc682610bf6565b600060208284031215610cad57600080fd5b5035919050565b600060208284031215610cc657600080fd5b610bc682610bcd565b60008451610ce1818460208901610b39565b80830190507f2e000000000000000000000000000000000000000000000000000000000000008082528551610d1d816001850160208a01610b39565b60019201918201528351610d38816002840160208801610b39565b0160020195945050505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203610da557610da5610d45565b5060010190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b600082610dea57610dea610dac565b500490565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b600082821015610e3057610e30610d45565b500390565b600082610e4457610e44610dac565b500690565b60008219821115610e5c57610e5c610d45565b500190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fdfea164736f6c634300080f000a"

func init() {
	if err := json.Unmarshal([]byte(SystemConfigStorageLayoutJSON), SystemConfigStorageLayout); err != nil {
		panic(err)
	}

	layouts["SystemConfig"] = SystemConfigStorageLayout
	deployedBytecodes["SystemConfig"] = SystemConfigDeployedBin
}