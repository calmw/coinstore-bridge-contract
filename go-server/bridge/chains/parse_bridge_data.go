package chains

import (
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/shopspring/decimal"
	"github.com/status-im/keycard-go/hexutils"
	"math/big"
	"strings"
)

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
