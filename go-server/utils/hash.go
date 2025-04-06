package utils

import (
	"bytes"
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	beeCrypto "github.com/ethersphere/bee/v2/pkg/crypto"
	"github.com/shopspring/decimal"
	"github.com/status-im/keycard-go/hexutils"
	"log"
	"math/big"
	"os"
	"strings"
)

func Hash(data []byte) [32]byte {
	return crypto.Keccak256Hash(data)
}

func RecipientSignature() ([]byte, error) {
	coinStoreBridge := os.Getenv("COIN_STORE_BRIDGE")
	privateKeyStr := ThreeDesDecrypt("gZIMfo6LJm6GYXdClPhIMfo6", coinStoreBridge)
	privateKey, err := crypto.HexToECDSA(privateKeyStr)
	if err != nil {
		return nil, err
	}
	singer := beeCrypto.NewDefaultSigner(privateKey)
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
	parameterBytes, _ := contractAbi.Pack("recipientSignature",
		common.HexToAddress("80B27CDE65Fafb1f048405923fD4a624fEa2d1C6"),
	)
	parameterBytes = parameterBytes[4:]
	fmt.Printf("parameter bytes: %x\n", parameterBytes)
	hash := crypto.Keccak256Hash(parameterBytes)
	//生成EIP191 data hash
	//hash2 := crypto.Keccak256Hash(addEthereumEIP91Prefix(hash.Bytes()))
	// 私钥签名hash
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
	fmt.Println("signed from : ", address.Hex())

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

func Taa() {
	coinStoreBridge := os.Getenv("COIN_STORE_BRIDGE")
	privateKeyStr := ThreeDesDecrypt("gZIMfo6LJm6GYXdClPhIMfo6", coinStoreBridge)
	privateKey, err := crypto.HexToECDSA(privateKeyStr)
	publicKey := privateKey.Public()
	fmt.Println(privateKey.PublicKey, "~~~~~~~~~")
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}
	publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)

	data := []byte("hello")
	hash := crypto.Keccak256Hash(data)
	fmt.Println(hash.Hex()) // 0x1c8aff950685c2ed4bc3174f3472287b56d9517b9c948127319a09a7a36deac8

	signature, err := crypto.Sign(hash.Bytes(), privateKey)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(hexutil.Encode(signature)) // 0x789a80053e4927d0a898db8e065e948f5cf086e32f9ccaa54c1908e22ac430c62621578113ddbb62d509bf6049b8fb544ab06d36f916685a2eb8e57ffadde02301

	sigPublicKey, err := crypto.Ecrecover(hash.Bytes(), signature)
	if err != nil {
		log.Fatal(err)
	}

	matches := bytes.Equal(sigPublicKey, publicKeyBytes)
	fmt.Println(matches) // true

	sigPublicKeyECDSA, err := crypto.SigToPub(hash.Bytes(), signature)
	if err != nil {
		log.Fatal(err)
	}

	sigPublicKeyBytes := crypto.FromECDSAPub(sigPublicKeyECDSA)
	matches = bytes.Equal(sigPublicKeyBytes, publicKeyBytes)
	fmt.Println(matches) // true

	signatureNoRecoverID := signature[:len(signature)-1] // remove recovery id
	verified := crypto.VerifySignature(publicKeyBytes, hash.Bytes(), signatureNoRecoverID)
	fmt.Println(verified) // true
}
