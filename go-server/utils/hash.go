package utils

import (
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	beeCrypto "github.com/ethersphere/bee/v2/pkg/crypto"
	"github.com/shopspring/decimal"
	"github.com/status-im/keycard-go/hexutils"
	"math/big"
	"os"
	"strings"
)

func Keccak256(data string) string {
	return strings.ToLower("0x" + hexutils.BytesToHex(crypto.Keccak256([]byte(data))))
}

func Hash(data []byte) [32]byte {
	return crypto.Keccak256Hash(data)
}

func RecipientSignature() ([]byte, error) {
	key := os.Getenv("COINSTORE_BRIDGE")
	abiString := `
[
	{
		"inputs": [{
        "internalType": "address",
        "name": "recipient",
        "type": "address"
      }],
		"name": "recipientSignature",
		"outputs": [],
		"stateMutability": "view",
		"type": "function"
	}
]
`
	contractAbi, _ := abi.JSON(strings.NewReader(abiString))
	// 生成hash
	bytes, _ := contractAbi.Pack("recipientSignature",
		common.HexToAddress("0x80B27CDE65Fafb1f048405923fD4a624fEa2d1C6"),
	)
	//fmt.Println(string(bytes))
	fmt.Printf("signature: %x\n", bytes)
	bytes = bytes[4:]
	fmt.Printf("signature: %x\n", bytes)
	hash := crypto.Keccak256Hash(bytes)
	//生成EIP191 data hash
	hash = crypto.Keccak256Hash(addEthereumEIP91Prefix(hash.Bytes()))
	// 私钥签名hash
	privateKey, err := crypto.HexToECDSA(key)
	if err != nil {
		return nil, err
	}
	singer := beeCrypto.NewDefaultSigner(privateKey)
	sign, err := singer.Sign(hash.Bytes())
	if err != nil {
		return nil, err
	}
	fmt.Printf("signature: %x\n", sign)
	keyR, err := beeCrypto.Recover(sign, hash.Bytes())
	if err != nil {
		return nil, err
	}

	address := crypto.PubkeyToAddress(*keyR)
	fmt.Println("wallet address: ", address.Hex())

	return sign, err
}

func addEthereumEIP91Prefix(data []byte) []byte {
	return []byte(fmt.Sprintf("\x19Ethereum Signed Message:\n%d%s", len(data), data))
}

func ParseBridgeData(depositData []byte) (decimal.Decimal, string, string, error) {
	abiStr := `[
    {
      "inputs": [
        {
          "internalType": "bytes32",
          "name": "resourceId",
          "type": "bytes32"
        },
        {
          "internalType": "uint256",
          "name": "chainId",
          "type": "uint256"
        },
        {
          "internalType": "address",
          "name": "sender",
          "type": "address"
        },
        {
          "internalType": "address",
          "name": "recipient",
          "type": "address"
        },
        {
          "internalType": "uint256",
          "name": "receiveAmount",
          "type": "uint256"
        },
        {
          "internalType": "uint256",
          "name": "localNonce",
          "type": "uint256"
        }
      ],
      "name": "depositData",
      "outputs": [],
      "stateMutability": "nonpayable",
      "type": "function"
    }
  ]`
	var inputData []byte
	inputData = append(hexutils.HexToBytes("0f25f052"), depositData...)
	contractAbi, _ := abi.JSON(strings.NewReader(abiStr))
	// 解析输入数据
	method, err := contractAbi.MethodById(inputData)
	if err != nil {
		return decimal.Zero, "", "", err
	}

	// 获取函数参数
	params := make([]interface{}, len(method.Inputs))
	if params, err = method.Inputs.Unpack(inputData[4:]); err != nil {
		return decimal.Zero, "", "", err
	}

	// 打印参数
	//fmt.Println("Method name:", method.Name)
	//fmt.Printf("Parameters:%x \n", params[0])
	//fmt.Println("Parameters:", params[1])
	//fmt.Println("Parameters:", params[2])
	//fmt.Println("Parameters:", params[3])
	//fmt.Println("Parameters:", params[4])

	amount := decimal.NewFromBigInt(params[4].(*big.Int), 0)
	caller := strings.ToLower(params[2].(common.Address).String())
	receiver := strings.ToLower(params[3].(common.Address).String())
	return amount, caller, receiver, nil
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
	//fmt.Printf("Inputs.Pack~~~~~~~~~~~~~~~~~~~~~~~~~~~~22222: 0x%x\n", AbiPacked[:4])
	//fmt.Println(AbiPacked[:4])
	return fmt.Sprintf("0x%x", AbiPacked), nil
}

type DepositRecord struct {
	DestinationChainId *big.Int
	Sender             common.Address
	ResourceID         [32]byte
	Ctime              *big.Int
	Data               []byte
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

type TokenInfo struct {
	AssetsType   uint8
	TokenAddress common.Address
	Pause        bool
	Fee          *big.Int
}

func ParseBridgeResourceIdToTokenInfo(inputData []byte) (TokenInfo, error) {
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
