package contract

import (
	"coinstore/utils"
	"encoding/hex"
	"fmt"
	"github.com/btcsuite/btcd/btcec/v2"
	"github.com/calmw/tron-sdk/pkg/common"
	"github.com/calmw/tron-sdk/pkg/keystore"
	"github.com/calmw/tron-sdk/pkg/store"
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
		//RPC: "grpc.trongrid.io:50051",
		//RPC:                   "3.225.171.164:50051",
		BridgeContractAddress: "TDv8tQFpajwxuyXasssppC8h9RwwKevXN9",
		VoteContractAddress:   "TLmg1WBwGJ3sSYbtevHP41n6bD1bvJ4xL6",
		//VoteContractAddress:   "TDVLRVqGpzAtoDBz16WbZvaa1eWF14ArYa",
		TantinContractAddress: "TPnLangi2RxS7vBKLa9uTBePELAdbowg5T",
		UsdtAddress:           "TXYZopYRdj2D9XRtbG411XZZ3kM5VkAeBf",
		UsdcAddress:           "TFXYQ93J5ptYcWaVFJxnadLZUB459uX2MK",
		EthAddress:            "0x0000000000000000000000000000000000000000",
		WEthAddress:           "TCAYwjRpHjFX4hgQqaJPbuNCFqxgqwJBns",
		PrivateKey:            privateKeyStr,
	}
}

func InitTronEnvProd() {
	coinStoreBridge := os.Getenv("COIN_STORE_BRIDGE_TRON")
	privateKeyStr := utils.ThreeDesDecrypt("gZIMfo6LJm6GYXdClPhIMfo6", coinStoreBridge)
	ChainConfig = ChainConfigs{
		BridgeId:    3,
		ChainId:     3448148188,
		ChainTypeId: 2,
		//RPC:         "grpc.nile.trongrid.io:50051",
		RPC: "grpc.trongrid.io:50051",
		//RPC:                   "3.225.171.164:50051",
		BridgeContractAddress: "TDv8tQFpajwxuyXasssppC8h9RwwKevXN9",
		VoteContractAddress:   "TLmg1WBwGJ3sSYbtevHP41n6bD1bvJ4xL6",
		TantinContractAddress: "TPnLangi2RxS7vBKLa9uTBePELAdbowg5T",
		UsdtAddress:           "TXYZopYRdj2D9XRtbG411XZZ3kM5VkAeBf",
		UsdcAddress:           "TFXYQ93J5ptYcWaVFJxnadLZUB459uX2MK",
		EthAddress:            "0x0000000000000000000000000000000000000000",
		WEthAddress:           "TCAYwjRpHjFX4hgQqaJPbuNCFqxgqwJBns",
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
