package abi

import (
	"fmt"
	"github.com/ethereum/go-ethereum/crypto"
	beeCrypto "github.com/ethersphere/bee/v2/pkg/crypto"
	"os"
)

func GenerateSignatureTron(parameter []byte) ([]byte, error) {
	privateKeyStr := os.Getenv("COINSTORE_BRIDGE_TRON_LOCAL")
	//privateKeyStr := "3f9f4b92d709f167b8ba98b9f89a5ec5272973aeb8f1affd11d5d2c67c5acf62"
	privateKey, err := crypto.HexToECDSA(privateKeyStr)
	singer := beeCrypto.NewDefaultSigner(privateKey)
	fmt.Println("1 ", fmt.Sprintf("%x", parameter))
	hash := crypto.Keccak256Hash(parameter)
	fmt.Println("2 ", fmt.Sprintf("%x", hash))
	sign, err := singer.Sign(hash.Bytes())
	if err != nil {
		return nil, err
	}
	fmt.Printf("0x%x\n", sign)
	return sign, err
}
