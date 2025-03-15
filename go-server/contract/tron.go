package contract

import (
	"encoding/hex"
	"fmt"
	"github.com/btcsuite/btcd/btcec/v2"
	"github.com/fbsobreira/gotron-sdk/pkg/common"
	"github.com/fbsobreira/gotron-sdk/pkg/keystore"
	"github.com/fbsobreira/gotron-sdk/pkg/store"
	"strings"
)

func GetKeyFromPrivateKey(privateKey, name, passphrase string) (*keystore.KeyStore, *keystore.Account, error) {
	privateKey = strings.TrimPrefix(privateKey, "0x")

	if store.DoesNamedAccountExist(name) {
		return nil, nil, fmt.Errorf("account %s already exists", name)
	}

	privateKeyBytes, err := hex.DecodeString(privateKey)
	if err != nil {
		return nil, nil, err
	}
	if len(privateKeyBytes) != common.Secp256k1PrivateKeyBytesLength {
		return nil, nil, common.ErrBadKeyLength
	}

	sk, _ := btcec.PrivKeyFromBytes(privateKeyBytes)
	ks := store.FromAccountName(name)
	account, err := ks.ImportECDSA(sk.ToECDSA(), passphrase)
	if err != nil {
		return nil, nil, err
	}
	return store.FromAccountName(name), &account, err
}
