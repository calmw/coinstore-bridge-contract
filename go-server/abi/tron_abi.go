package abi

import (
	"fmt"
	"github.com/ethereum/go-ethereum/crypto"
	beeCrypto "github.com/ethersphere/bee/v2/pkg/crypto"
	"os"
)

func GenerateSignatureTron(parameter []byte) ([]byte, error) {
	privateKeyStr := os.Getenv("COINSTORE_BRIDGE_TRON_LOCAL")
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
