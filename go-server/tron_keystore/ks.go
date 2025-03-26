package tron_keystore

import (
	"encoding/hex"
	"github.com/btcsuite/btcd/btcec/v2"
	"github.com/fbsobreira/gotron-sdk/pkg/common"
	"github.com/fbsobreira/gotron-sdk/pkg/keystore"
	"os"
	"strings"
)

const (
	KeyStorePassphrase = "your_strong_password"
	KeyStoreDir        = "./tron_keystore/keystore"
)

func InitKeyStore() (*keystore.KeyStore, *keystore.Account, error) {
	privateKeyStr := os.Getenv("COINSTORE_BRIDGE_TRON")
	if err := os.MkdirAll(KeyStoreDir, 0700); err != nil {
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
	account, err := ks.ImportECDSA(sk.ToECDSA(), KeyStorePassphrase)
	if err != nil && !strings.HasPrefix(err.Error(), "account already exists") {
		return nil, nil, err
	}

	return ks, &account, nil
}
