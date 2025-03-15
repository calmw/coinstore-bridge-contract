package services

import (
	"coinstore/contract"
	"os"
)

func InitOpenBnbEnv() {
	key := os.Getenv("COINSTORE_BRIDGE")
	contract.ChainConfig = contract.ChainConfigs{
		BridgeId:              1,
		ChainId:               5611,
		ChainTypeId:           1,
		RPC:                   "https://opbnb-testnet-rpc.bnbchain.org",
		BridgeContractAddress: "0xEA2c0B226670cd66f44560deA091DfA860C892eF",
		VoteContractAddress:   "0x77bcb682e01D8763F6757c0D0Beaf577Afcfdf43",
		TantinContractAddress: "0xd78CCb0C79489d8C50AfFE4E881B2C2E93706f8b",
		UsdtAddress:           "0xfBe1e02C25a04f6CD6b044F847697b48B3E99a16",
		PrivateKey:            key,
	}
}

func InitTronEnv() {
	key := os.Getenv("COINSTORE_BRIDGE_TRON")
	contract.ChainConfig = contract.ChainConfigs{
		BridgeId:              2,
		ChainId:               3448148188,
		ChainTypeId:           2,
		RPC:                   "grpc.nile.trongrid.io:50051",
		BridgeContractAddress: "TBbjCKYeR2gDXvTtUNbRgTppWLv2W2XRTH",
		VoteContractAddress:   "TEgAohYmMTz8sRRZSf5ht69Q3jPBA4vKSz",
		TantinContractAddress: "TBVtUdewEJJ1zYif9tmYUC2sgERDYfWUUS",
		UsdtAddress:           "TU6UjUJadm8TungBHvL4n9apv8Jns4wJiz",
		PrivateKey:            key,
	}
}
