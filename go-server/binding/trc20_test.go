package binding

import (
	"fmt"
	"github.com/fbsobreira/gotron-sdk/pkg/store"
	"math/big"
	"os"
	"testing"
)

func TestTransferCoin(t *testing.T) {
	privateKey := os.Getenv("COINSTORE_BRIDGE_TRON")
	fromAddress := "TFBymbm7LrbRreGtByMPRD2HUyneKabsqb"
	toAddress := "TEwX7WKNQqsRpxd6KyHHQPMMigLg9c258y"
	amount := int64(1)
	coin, err := TransferCoin(privateKey, fromAddress, toAddress, amount)
	fmt.Println(coin, err)

}

func TestPrivateKeyToWalletAddress(t *testing.T) {
	privateKey := os.Getenv("COINSTORE_BRIDGE_TRON")
	address, err := PrivateKeyToWalletAddress(privateKey)
	fmt.Println(address, err)
}

func TestImportFromPrivateKey(t *testing.T) {
	privateKey := os.Getenv("COINSTORE_BRIDGE_TRON")
	ks, ka, err := GetKeyFromPrivateKey(privateKey, AccountName, Passphrase)
	fmt.Println(ks, ka, err)
	ks, ka, err = store.UnlockedKeystore(OwnerAccount, Passphrase)
	fmt.Println("~~~~~~~", ks, ka, err)
	// /Users/cisco/.tronctl/account-keys/my_account/UTC--2025-03-15T15-30-58.603647000Z--413942fda93c573e2ce9e85b0bb00ba98a144f27f6
}

func TestTransferToken(t *testing.T) {
	txHash, err := TransferToken()
	fmt.Println(txHash, err)
}

func TestTrc20_Approve(t1 *testing.T) {
	erc20, err := NewTrc20(UsdtAddress)
	if err != nil {
		fmt.Println(err)
		return
	}
	txHash, err := erc20.Approve("TEy2BtGxixqhbcM7w65rJiotAerSBFR48W", big.NewInt(1))
	fmt.Println(txHash, err)
}

func TestTrc20_Transfer(t1 *testing.T) {
	erc20, err := NewTrc20(UsdtAddress)
	if err != nil {
		fmt.Println(err)
		return
	}
	txHash, err := erc20.Transfer("TEy2BtGxixqhbcM7w65rJiotAerSBFR48W", big.NewInt(1))
	fmt.Println(txHash, err)
}

func TestAA(t *testing.T) {
	AA()
}
