package services

import (
	"coinstore/blockchain"
	"fmt"
	"os"
)

func InitOpenBnbEnv() {
	key := os.Getenv("COINSTORE_BRIDGE")
	fmt.Println(key)

	blockchain.ChainConfig = blockchain.ChainConfigs{
		BridgeId:              1,
		ChainId:               5611,
		RPC:                   "https://opbnb-testnet-rpc.bnbchain.org",
		BridgeContractAddress: "0x73AB93aC16D86333a137D4bd9ba34dD310b9175a",
		VoteContractAddress:   "0x77bcb682e01D8763F6757c0D0Beaf577Afcfdf43",
		TantinContractAddress: "0xd78CCb0C79489d8C50AfFE4E881B2C2E93706f8b",
		UsdtAddress:           "0xfBe1e02C25a04f6CD6b044F847697b48B3E99a16",
		PrivateKey:            key,
	}
}

func InitTronEnv() {
	key := os.Getenv("COINSTORE_BRIDGE")
	fmt.Println(key)

	blockchain.ChainConfig = blockchain.ChainConfigs{
		BridgeId:              10,
		ChainId:               5611,
		RPC:                   "https://nile.trongrid.io/jsonrpc/",
		BridgeContractAddress: "0x73AB93aC16D86333a137D4bd9ba34dD310b9175a",
		VoteContractAddress:   "0x77bcb682e01D8763F6757c0D0Beaf577Afcfdf43",
		TantinContractAddress: "0xd78CCb0C79489d8C50AfFE4E881B2C2E93706f8b",
		UsdtAddress:           "0xfBe1e02C25a04f6CD6b044F847697b48B3E99a16",
		PrivateKey:            key,
	}
}
