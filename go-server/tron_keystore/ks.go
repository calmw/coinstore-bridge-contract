package tron_keystore

import (
	"encoding/hex"
	"github.com/btcsuite/btcd/btcec/v2"
	"github.com/calmw/tron-sdk/pkg/common"
	"github.com/calmw/tron-sdk/pkg/keystore"
	"os"
	"strings"
)

const (
	KeyStorePassphrase = "your_strong_password"
	KeyStoreDir        = "./tron_keystore/keystore"
)

func InitKeyStore(privateKeyStr string) (*keystore.KeyStore, *keystore.Account, error) {
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
