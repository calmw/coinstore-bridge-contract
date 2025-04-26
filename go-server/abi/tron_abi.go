package abi

import (
	"fmt"
	beeCrypto "github.com/calmw/bee-tron/pkg/crypto"
	"github.com/ethereum/go-ethereum/crypto"
	"os"
)

func GenerateSignatureTron(parameter []byte) ([]byte, error) {
	privateKeyStr := os.Getenv("COINSTORE_BRIDGE_TRON_LOCAL")
	//privateKeyStr := "3f9f4b92d709f167b8ba98b9f89a5ec5272973aeb8f1affd11d5d2c67c5acf62"
	privateKey, err := crypto.HexToECDSA(privateKeyStr)
	singer := beeCrypto.NewDefaultSigner(privateKey)
	hash := crypto.Keccak256Hash(parameter)
	sign, err := singer.Sign(hash.Bytes())
	if err != nil {
		return nil, err
	}
	fmt.Printf("0x%x\n", sign)
	return sign, err
}

// GetTronSignatureMachineFromSm 从签名机获取TRON链签名
func GetTronSignatureMachineFromSm(parameter []byte) ([]byte, error) {
	privateKeyStr := os.Getenv("COINSTORE_BRIDGE_TRON_LOCAL")
	//privateKeyStr := "3f9f4b92d709f167b8ba98b9f89a5ec5272973aeb8f1affd11d5d2c67c5acf62"
	privateKey, err := crypto.HexToECDSA(privateKeyStr)
	singer := beeCrypto.NewDefaultSigner(privateKey)
	hash := crypto.Keccak256Hash(parameter)
	sign, err := singer.Sign(hash.Bytes())
	if err != nil {
		return nil, err
	}
	fmt.Printf("0x%x\n", sign)
	return sign, err
}
