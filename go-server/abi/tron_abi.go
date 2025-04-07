package abi

import (
	"coinstore/utils"
	"fmt"
	"github.com/ethereum/go-ethereum/crypto"
	beeCrypto "github.com/ethersphere/bee/v2/pkg/crypto"
	"os"
)

func GenerateSignatureTron(parameter []byte) ([]byte, error) {
	coinStoreBridge := os.Getenv("COIN_STORE_BRIDGE_TRON")
	privateKeyStr := utils.ThreeDesDecrypt("gZIMfo6LJm6GYXdClPhIMfo6", coinStoreBridge)
	privateKey, err := crypto.HexToECDSA(privateKeyStr)
	fmt.Println(privateKeyStr)
	if err != nil {
		return nil, err
	}
	singer := beeCrypto.NewDefaultSigner(privateKey)

	hash := crypto.Keccak256Hash(parameter)
	// 私钥签名hash
	sign, err := singer.Sign(hash.Bytes())
	if err != nil {
		return nil, err
	}
	fmt.Printf("0x%x\n", sign)
	return sign, err
}
