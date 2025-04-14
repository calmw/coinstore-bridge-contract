package tron

import (
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
	"strings"
)

type IVoteProposal struct {
	ResourceId    [32]byte         `json:"resourceId"`
	DataHash      [32]byte         `json:"dataHash"`
	YesVotes      []common.Address `json:"yesVotes"`
	NoVotes       []common.Address `json:"noVotes"`
	Status        uint8            `json:"status"`
	ProposedBlock *big.Int         `json:"proposedBlock"`
}
type JsonRpcResponse struct {
	Jsonrpc string `json:"jsonrpc"`
	ID      int    `json:"id"`
	Result  string `json:"result"`
}

type DepositRecord struct {
	DestinationChainId *big.Int
	Sender             common.Address
	ResourceID         [32]byte
	Ctime              *big.Int
	Data               []byte
}

type TokenInfo struct {
	AssetsType   uint8
	TokenAddress common.Address
	Pause        bool
	Fee          *big.Int
}

func GenerateBridgeDepositRecordsData(destinationChainId, depositNonce *big.Int) (string, error) {
	contractABI := `[
    {
    "inputs": [
      {
        "internalType": "uint256",
        "name": "",
        "type": "uint256"
      },
      {
        "internalType": "uint256",
        "name": "",
        "type": "uint256"
      }
    ],
    "name": "depositRecords",
    "outputs": [
      {
        "internalType": "uint256",
        "name": "destinationChainId",
        "type": "uint256"
      },
      {
        "internalType": "address",
        "name": "sender",
        "type": "address"
      },
      {
        "internalType": "bytes32",
        "name": "resourceID",
        "type": "bytes32"
      },
      {
        "internalType": "uint256",
        "name": "ctime",
        "type": "uint256"
      },
      {
        "internalType": "bytes",
        "name": "data",
        "type": "bytes"
      }
    ],
    "stateMutability": "view",
    "type": "function"
  }
  ]`

	// 解析合约的ABI
	parsedABI, err := abi.JSON(strings.NewReader(contractABI))
	if err != nil {
		return "", err
	}
	// 创建一个方法对象，指向我们想要调用的合约函数
	AbiPacked, err := parsedABI.Pack("depositRecords", destinationChainId, depositNonce)
	if err != nil {
		return "", err
	}
	// 打印出Inputs.Pack的结果
	//fmt.Printf("Inputs.Pack: 0x%x\n", AbiPacked[:4])
	//fmt.Println(AbiPacked[:4])
	return fmt.Sprintf("%x", AbiPacked), nil
}

func ParseBridgeDepositRecordData(inputData []byte) (DepositRecord, error) {
	contractABI := `[
    {
    "inputs": [
      {
        "internalType": "uint256",
        "name": "",
        "type": "uint256"
      },
      {
        "internalType": "uint256",
        "name": "",
        "type": "uint256"
      }
    ],
    "name": "depositRecords",
    "outputs": [
      {
        "internalType": "uint256",
        "name": "destinationChainId",
        "type": "uint256"
      },
      {
        "internalType": "address",
        "name": "sender",
        "type": "address"
      },
      {
        "internalType": "bytes32",
        "name": "resourceID",
        "type": "bytes32"
      },
      {
        "internalType": "uint256",
        "name": "ctime",
        "type": "uint256"
      },
      {
        "internalType": "bytes",
        "name": "data",
        "type": "bytes"
      }
    ],
    "stateMutability": "view",
    "type": "function"
  }
  ]`

	// 解析合约的ABI
	parsedABI, err := abi.JSON(strings.NewReader(contractABI))
	if err != nil {
		panic(err)
	}
	// 解析输入数据
	method, err := parsedABI.MethodById(inputData)
	if err != nil {
		fmt.Println("Error parsing input data:", err)
		return DepositRecord{}, err
	}

	// 获取函数参数
	outputs := make([]interface{}, len(method.Outputs))
	if outputs, err = method.Outputs.Unpack(inputData[4:]); err != nil {
		fmt.Println("Error unpacking parameters:", err)
		return DepositRecord{}, err
	}

	// 打印参数
	//fmt.Println("Method name:", method.Name)
	//fmt.Println("outputs:", outputs)
	destinationChainId, ok := outputs[0].(*big.Int)
	if !ok {
		return DepositRecord{}, fmt.Errorf("invalid destinationChainId type")
	}
	resourceID, ok := outputs[2].([32]byte)
	if !ok {
		return DepositRecord{}, fmt.Errorf("invalid resourceID type")
	}
	sender, ok := outputs[1].(common.Address)
	if !ok {
		return DepositRecord{}, fmt.Errorf("invalid sender type")
	}
	ctime, ok := outputs[3].(*big.Int)
	if !ok {
		return DepositRecord{}, fmt.Errorf("invalid ctime type")
	}
	data, ok := outputs[4].([]byte)
	if !ok {
		return DepositRecord{}, fmt.Errorf("invalid data type")
	}
	res := DepositRecord{
		DestinationChainId: destinationChainId,
		Sender:             sender,
		ResourceID:         resourceID,
		Ctime:              ctime,
		Data:               data,
	}
	return res, err
}

func GenerateBridgeGetTokenInfoByResourceId(resourceID [32]byte) (string, error) {
	contractABI := `[
   {
    "inputs": [
      {
        "internalType": "bytes32",
        "name": "",
        "type": "bytes32"
      }
    ],
    "name": "resourceIdToTokenInfo",
    "outputs": [
      {
        "internalType": "enum IBridge.AssetsType",
        "name": "assetsType",
        "type": "uint8"
      },
      {
        "internalType": "address",
        "name": "tokenAddress",
        "type": "address"
      },
      {
        "internalType": "bool",
        "name": "pause",
        "type": "bool"
      },
      {
        "internalType": "uint256",
        "name": "fee",
        "type": "uint256"
      },
      {
        "internalType": "bool",
        "name": "burnable",
        "type": "bool"
      },
      {
        "internalType": "bool",
        "name": "mintable",
        "type": "bool"
      }
    ],
    "stateMutability": "view",
    "type": "function"
  }
 ]`

	// 解析合约的ABI
	parsedABI, err := abi.JSON(strings.NewReader(contractABI))
	if err != nil {
		return "", err
	}

	// 创建一个方法对象，指向我们想要调用的合约函数
	AbiPacked, err := parsedABI.Pack("resourceIdToTokenInfo", resourceID)
	if err != nil {
		return "", err
	}

	// 打印出Inputs.Pack的结果
	return fmt.Sprintf("0x%x", AbiPacked), nil
}

func ParseBridgeResourceIdToTokenInfo(inputData []byte) (TokenInfo, error) {
	contractABI := `[{
    "inputs": [
      {
        "internalType": "bytes32",
        "name": "",
        "type": "bytes32"
      }
    ],
    "name": "resourceIdToTokenInfo",
    "outputs": [
      {
        "internalType": "enum IBridge.AssetsType",
        "name": "assetsType",
        "type": "uint8"
      },
      {
        "internalType": "address",
        "name": "tokenAddress",
        "type": "address"
      },
      {
        "internalType": "bool",
        "name": "pause",
        "type": "bool"
      },
      {
        "internalType": "uint256",
        "name": "fee",
        "type": "uint256"
      },
      {
        "internalType": "bool",
        "name": "burnable",
        "type": "bool"
      },
      {
        "internalType": "bool",
        "name": "mintable",
        "type": "bool"
      }
    ],
    "stateMutability": "view",
    "type": "function"
  }]`

	// 解析合约的ABI
	parsedABI, err := abi.JSON(strings.NewReader(contractABI))
	if err != nil {
		panic(err)
	}
	// 解析输入数据
	method, err := parsedABI.MethodById(inputData)
	if err != nil {
		fmt.Println("Error parsing input data:", err)
		return TokenInfo{}, err
	}

	// 获取函数参数
	outputs := make([]interface{}, len(method.Outputs))
	if outputs, err = method.Outputs.Unpack(inputData[4:]); err != nil {
		fmt.Println("Error unpacking parameters:", err)
		return TokenInfo{}, err
	}

	// 打印参数
	assetsType, ok := outputs[0].(uint8)
	if !ok {
		return TokenInfo{}, fmt.Errorf("invalid assetsType type")
	}
	tokenAddress, ok := outputs[1].(common.Address)
	if !ok {
		return TokenInfo{}, fmt.Errorf("invalid tokenAddress type")
	}
	pause, ok := outputs[2].(bool)
	if !ok {
		return TokenInfo{}, fmt.Errorf("invalid pause type")
	}
	fee, ok := outputs[3].(*big.Int)
	if !ok {
		return TokenInfo{}, fmt.Errorf("invalid fee type")
	}
	res := TokenInfo{
		AssetsType:   assetsType,
		TokenAddress: tokenAddress,
		Pause:        pause,
		Fee:          fee,
	}
	return res, err
}

func GenerateVoteGetProposal(originChainID *big.Int, depositNonce *big.Int, dataHash [32]byte) (string, error) {
	contractABI := `[
    {
    "inputs": [
      {
        "internalType": "uint256",
        "name": "originChainID",
        "type": "uint256"
      },
      {
        "internalType": "uint256",
        "name": "depositNonce",
        "type": "uint256"
      },
      {
        "internalType": "bytes32",
        "name": "dataHash",
        "type": "bytes32"
      }
    ],
    "name": "getProposal",
    "outputs": [
      {
        "components": [
          {
            "internalType": "bytes32",
            "name": "resourceId",
            "type": "bytes32"
          },
          {
            "internalType": "bytes32",
            "name": "dataHash",
            "type": "bytes32"
          },
          {
            "internalType": "address[]",
            "name": "yesVotes",
            "type": "address[]"
          },
          {
            "internalType": "address[]",
            "name": "noVotes",
            "type": "address[]"
          },
          {
            "internalType": "enum IVote.ProposalStatus",
            "name": "status",
            "type": "uint8"
          },
          {
            "internalType": "uint256",
            "name": "proposedBlock",
            "type": "uint256"
          }
        ],
        "internalType": "struct IVote.Proposal",
        "name": "",
        "type": "tuple"
      }
    ],
    "stateMutability": "view",
    "type": "function"
  }
  ]`

	// 解析合约的ABI
	parsedABI, err := abi.JSON(strings.NewReader(contractABI))
	if err != nil {
		return "", err
	}

	// 创建一个方法对象，指向我们想要调用的合约函数
	AbiPacked, err := parsedABI.Pack("getProposal", originChainID, depositNonce, dataHash)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("0x%x", AbiPacked), nil
}

func ParseVoteGetProposal(inputData []byte) (IVoteProposal, error) {
	contractABI := `[
    {
    "inputs": [
      {
        "internalType": "uint256",
        "name": "originChainID",
        "type": "uint256"
      },
      {
        "internalType": "uint256",
        "name": "depositNonce",
        "type": "uint256"
      },
      {
        "internalType": "bytes32",
        "name": "dataHash",
        "type": "bytes32"
      }
    ],
    "name": "getProposal",
    "outputs": [
      {
        "components": [
          {
            "internalType": "bytes32",
            "name": "resourceId",
            "type": "bytes32"
          },
          {
            "internalType": "bytes32",
            "name": "dataHash",
            "type": "bytes32"
          },
          {
            "internalType": "address[]",
            "name": "yesVotes",
            "type": "address[]"
          },
          {
            "internalType": "address[]",
            "name": "noVotes",
            "type": "address[]"
          },
          {
            "internalType": "enum IVote.ProposalStatus",
            "name": "status",
            "type": "uint8"
          },
          {
            "internalType": "uint256",
            "name": "proposedBlock",
            "type": "uint256"
          }
        ],
        "internalType": "struct IVote.Proposal",
        "name": "",
        "type": "tuple"
      }
    ],
    "stateMutability": "view",
    "type": "function"
  }
  ]`

	// 解析合约的ABI
	parsedABI, err := abi.JSON(strings.NewReader(contractABI))
	if err != nil {
		panic(err)
	}
	// 解析输入数据
	method, err := parsedABI.MethodById(inputData)
	if err != nil {
		fmt.Println("Error parsing input data:", err)
		return IVoteProposal{}, err
	}

	// 获取函数参数
	outputs := make([]interface{}, len(method.Outputs))
	if outputs, err = method.Outputs.Unpack(inputData[4:]); err != nil {
		fmt.Println("Error unpacking parameters:", err)
		return IVoteProposal{}, err
	}
	type As struct {
		ResourceId    [32]uint8        `json:"resourceId"`
		DataHash      [32]uint8        `json:"dataHash"`
		YesVotes      []common.Address `json:"yesVotes"`
		NoVotes       []common.Address `json:"noVotes"`
		Status        uint8            `json:"status"`
		ProposedBlock *big.Int         `json:"proposedBlock"`
	}
	// 打印参数
	output := outputs[0]
	result := output.(struct {
		ResourceId    [32]uint8        "json:\"resourceId\""
		DataHash      [32]uint8        "json:\"dataHash\""
		YesVotes      []common.Address "json:\"yesVotes\""
		NoVotes       []common.Address "json:\"noVotes\""
		Status        uint8            "json:\"status\""
		ProposedBlock *big.Int         "json:\"proposedBlock\""
	})
	res := IVoteProposal{
		ResourceId:    result.ResourceId,
		DataHash:      result.DataHash,
		YesVotes:      result.YesVotes,
		NoVotes:       result.NoVotes,
		Status:        result.Status,
		ProposedBlock: result.ProposedBlock,
	}
	return res, err
}

func GenerateVoteHasVotedOnProposal(arg0 *big.Int, arg1 [32]byte, arg2 common.Address) (string, error) {
	contractABI := `[
    {
    "inputs": [
      {
        "internalType": "uint72",
        "name": "",
        "type": "uint72"
      },
      {
        "internalType": "bytes32",
        "name": "",
        "type": "bytes32"
      },
      {
        "internalType": "address",
        "name": "",
        "type": "address"
      }
    ],
    "name": "hasVotedOnProposal",
    "outputs": [
      {
        "internalType": "bool",
        "name": "",
        "type": "bool"
      }
    ],
    "stateMutability": "view",
    "type": "function"
  }
  ]`

	// 解析合约的ABI
	parsedABI, err := abi.JSON(strings.NewReader(contractABI))
	if err != nil {
		return "", err
	}

	// 创建一个方法对象，指向我们想要调用的合约函数
	AbiPacked, err := parsedABI.Pack("hasVotedOnProposal", arg0, arg1, arg2)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("0x%x", AbiPacked), nil
}

func ParseVoteHasVotedOnProposal(inputData []byte) (bool, error) {
	contractABI := `[
    {
    "inputs": [
      {
        "internalType": "uint72",
        "name": "",
        "type": "uint72"
      },
      {
        "internalType": "bytes32",
        "name": "",
        "type": "bytes32"
      },
      {
        "internalType": "address",
        "name": "",
        "type": "address"
      }
    ],
    "name": "hasVotedOnProposal",
    "outputs": [
      {
        "internalType": "bool",
        "name": "",
        "type": "bool"
      }
    ],
    "stateMutability": "view",
    "type": "function"
  }
  ]`

	// 解析合约的ABI
	parsedABI, err := abi.JSON(strings.NewReader(contractABI))
	if err != nil {
		panic(err)
	}
	// 解析输入数据
	method, err := parsedABI.MethodById(inputData)
	if err != nil {
		fmt.Println("Error parsing input data:", err)
		return false, err
	}

	// 获取函数参数
	outputs := make([]interface{}, len(method.Outputs))
	if outputs, err = method.Outputs.Unpack(inputData[4:]); err != nil {
		fmt.Println("Error unpacking parameters:", err)
		return false, err
	}

	// 打印参数
	hashVote, ok := outputs[0].(bool)
	if !ok {
		return false, fmt.Errorf("invalid hashVote type")
	}
	return hashVote, err
}

func GenerateSigNonce() (string, error) {
	contractABI := `[
{
    "inputs": [],
    "name": "sigNonce",
    "outputs": [
      {
        "internalType": "uint256",
        "name": "",
        "type": "uint256"
      }
    ],
    "stateMutability": "view",
    "type": "function"
  }
  ]`

	// 解析合约的ABI
	parsedABI, err := abi.JSON(strings.NewReader(contractABI))
	if err != nil {
		return "", err
	}

	// 创建一个方法对象，指向我们想要调用的合约函数
	AbiPacked, err := parsedABI.Pack("sigNonce")
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("0x%x", AbiPacked), nil
}

func ParseSigNonce(inputData []byte) (*big.Int, error) {
	contractABI := `[
{
    "inputs": [],
    "name": "sigNonce",
    "outputs": [
      {
        "internalType": "uint256",
        "name": "",
        "type": "uint256"
      }
    ],
    "stateMutability": "view",
    "type": "function"
  }
  ]`

	// 解析合约的ABI
	parsedABI, err := abi.JSON(strings.NewReader(contractABI))
	if err != nil {
		panic(err)
	}
	// 解析输入数据
	method, err := parsedABI.MethodById(inputData)
	if err != nil {
		fmt.Println("Error parsing input data:", err)
		return nil, err
	}

	// 获取函数参数
	outputs := make([]interface{}, len(method.Outputs))
	if outputs, err = method.Outputs.Unpack(inputData[4:]); err != nil {
		fmt.Println("Error unpacking parameters:", err)
		return nil, err
	}
	type As struct {
		ResourceId    [32]uint8        `json:"resourceId"`
		DataHash      [32]uint8        `json:"dataHash"`
		YesVotes      []common.Address `json:"yesVotes"`
		NoVotes       []common.Address `json:"noVotes"`
		Status        uint8            `json:"status"`
		ProposedBlock *big.Int         `json:"proposedBlock"`
	}
	// 打印参数
	output := outputs[0]
	result := output.(*big.Int)
	return result, err
}
