package tron

import (
	"encoding/hex"
	"fmt"
	"github.com/btcsuite/btcd/btcec/v2"
	"github.com/calmw/tron-sdk/pkg/common"
	"github.com/calmw/tron-sdk/pkg/keystore"
	"os"
	"strings"
)

const (
	KeyStoreDir = "./tron_keystore/keystore"
)

func InitKeyStore(privateKeyStr, passphrase string) (*keystore.KeyStore, *keystore.Account, error) {
	if err := Mkdir(KeyStoreDir); err != nil {
		return nil, nil, err
	}
	privateKeyBytes, err := hex.DecodeString(privateKeyStr)
	if err != nil {
		return nil, nil, err
	}
	if len(privateKeyBytes) != common.Secp256k1PrivateKeyBytesLength {
		return nil, nil, common.ErrBadKeyLength
	}

	sk, _ := btcec.PrivKeyFromBytes(privateKeyBytes)
	ks := keystore.NewKeyStore(KeyStoreDir, keystore.StandardScryptN, keystore.StandardScryptP)
	account, err := ks.ImportECDSA(sk.ToECDSA(), passphrase)
	if err != nil && !strings.HasPrefix(err.Error(), "account already exists") {
		return nil, nil, err
	}

	return ks, &account, nil
}

func Mkdir(dirPath string) error {
	if _, err := os.Stat(dirPath); os.IsNotExist(err) {
		err = os.MkdirAll(dirPath, 0755) // 0755是权限设置，你可以根据需要修改
		if err != nil {
			fmt.Println("创建文件夹失败:", err)
			return err
		}
		fmt.Println("文件夹创建成功:", dirPath)
	}
	return nil
}
