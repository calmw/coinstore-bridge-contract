package contract

import (
	"coinstore/utils"
	"encoding/hex"
	"fmt"
	"github.com/btcsuite/btcd/btcec/v2"
	"github.com/fbsobreira/gotron-sdk/pkg/common"
	"github.com/fbsobreira/gotron-sdk/pkg/keystore"
	"github.com/fbsobreira/gotron-sdk/pkg/store"
	"os"
	"strings"
)

func InitTronEnv() {
	coinStoreBridge := os.Getenv("COIN_STORE_BRIDGE_TRON")
	privateKeyStr := utils.ThreeDesDecrypt("gZIMfo6LJm6GYXdClPhIMfo6", coinStoreBridge)
	ChainConfig = ChainConfigs{
		BridgeId:    3,
		ChainId:     3448148188,
		ChainTypeId: 2,
		RPC:         "grpc.nile.trongrid.io:50051",
		//RPC:                   "grpc.trongrid.io:50051",
		BridgeContractAddress: "TUAPRMUqZYaQgNrvyNbYs4ZzwjHRh8SnS6",
		VoteContractAddress:   "THy5A71b4EuAddwZJ5voncmCJw4uaWxTqX",
		TantinContractAddress: "TMY9sxRvwWHeUZmAuUCwAJxLimV3WUo8rL",
		UsdtAddress:           "TR7NHqjeKQxGTCi8q8ZY4pL8otSzgjLj6t",
		UsdcAddress:           "TEkxiTehnzSmSe2XqrBj4w32RUN966rdz8",
		EthAddress:            "0x0000000000000000000000000000000000000000",
		WEthAddress:           "THb4CqiFdwNHsWsQCs4JhzwjMWys4aqCbF",
		PrivateKey:            privateKeyStr,
	}
}

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
