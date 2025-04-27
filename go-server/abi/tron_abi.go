package abi

import (
	"coinstore/utils"
	"fmt"
	beeCrypto "github.com/calmw/bee-tron/pkg/crypto"
	"github.com/ethereum/go-ethereum/crypto"
	"os"
)

func GenerateSignatureTron(parameter []byte) ([]byte, error) {
	privateKeyStr := os.Getenv("TB_KEY")
	if len(privateKeyStr) < 0 {
		privateKeyStr = os.Getenv("COINSTORE_BRIDGE_TRON_LOCAL")
	}
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

func GeneratePriceSignatureTron(parameter []byte) ([]byte, error) {
	privateKeyStr := os.Getenv("PRICE_SIG_ACCOUNT_TRON")
	privateKeyStr = utils.ThreeDesDecrypt("gZIMfo6LJm6GYXdClPhIMfo6", privateKeyStr)
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
