package abi

import (
	"fmt"
	beeCrypto "github.com/calmw/bee-tron/pkg/crypto"
	"github.com/ethereum/go-ethereum/crypto"
	"os"
)

func GenerateSignatureTron(parameter []byte) ([]byte, error) {
	//privateKeyStr := os.Getenv("COINSTORE_BRIDGE_TRON_LOCAL")
	//privateKeyStr := "3f9f4b92d709f167b8ba98b9f89a5ec5272973aeb8f1affd11d5d2c67c5acf62"
	privateKeyStr := os.Getenv("TB_KEY")
	if len(privateKeyStr) <= 0 {
		privateKeyStr = os.Getenv("COINSTORE_BRIDGE_TRON_LOCAL")
		//privateKeyStr = utils.ThreeDesDecrypt("gZIMfo6LJm6GYXdClPhIMfo6", privateKeyStr)
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
	//privateKeyStr := os.Getenv("PRICE_SIG_ACCOUNT_TRON")
	//privateKeyStr = utils.ThreeDesDecrypt("gZIMfo6LJm6GYXdClPhIMfo6", privateKeyStr)
	privateKeyStr := os.Getenv("TB_KEY")
	if len(privateKeyStr) <= 0 {
		privateKeyStr = os.Getenv("COINSTORE_BRIDGE_TRON_LOCAL")
		//privateKeyStr = utils.ThreeDesDecrypt("gZIMfo6LJm6GYXdClPhIMfo6", privateKeyStr)
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
