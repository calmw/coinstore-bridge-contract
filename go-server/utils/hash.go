package utils

import (
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	beeCrypto "github.com/ethersphere/bee/v2/pkg/crypto"
	"github.com/status-im/keycard-go/hexutils"
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
