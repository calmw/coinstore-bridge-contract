package binding

import (
	"fmt"
	"os"
	"testing"
)

func TestTransferCoin(t *testing.T) {
	privateKey := os.Getenv("COINSTORE_BRIDGE_TRON")
	fromAddress := "TFBymbm7LrbRreGtByMPRD2HUyneKabsqb"
	toAddress := "TEwX7WKNQqsRpxd6KyHHQPMMigLg9c258y"
	amount := int64(1)
	TransferCoin(privateKey, fromAddress, toAddress, amount)

}

func TestPrivateKeyToWalletAddress(t *testing.T) {
	privateKey := os.Getenv("COINSTORE_BRIDGE_TRON")
	address, err := PrivateKeyToWalletAddress(privateKey)
	fmt.Println(address, err)
}
