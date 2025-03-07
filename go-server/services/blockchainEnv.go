package services

import (
	"coinstore/pkg/blockchain"
	"os"
)

func InitOpenBnbEnv() {
	key := os.Getenv("COINSTORE_BRIDGE")

	blockchain.ChainConfig = blockchain.ChainConfigs{
		ChainId:               5611,
		RPC:                   "https://opbnb-testnet-rpc.bnbchain.org",
		BridgeContractAddress: "0x6b66eBFA87AaC1dB355B0ec49ECab7F4b32b1b30",
		VoteContractAddress:   "0x77bcb682e01D8763F6757c0D0Beaf577Afcfdf43",
		TantinContractAddress: "0xd78CCb0C79489d8C50AfFE4E881B2C2E93706f8b",
		UsdtAddress:           "0xF9A555844149DEDC69c3D1F4D3a3a6C083Eb5A4c",
		PrivateKey:            key,
	}
}
