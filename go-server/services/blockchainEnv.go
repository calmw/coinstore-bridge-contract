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
		BridgeContractAddress: "0x7293d8256178331637557753790161f09899a118",
		VoteContractAddress:   "0x329552109874d5050420897035a828285f95f969",
		TantinContractAddress: "0x4F53242025F519072121B775D64D6409F6382D7D",
		UsdtAddress:           "0xF9A555844149DEDC69c3D1F4D3a3a6C083Eb5A4c",
		PrivateKey:            key,
	}
}

func InitMainnetEnv() {
	key := os.Getenv("COINSTORE_BRIDGE")
	blockchain.ChainConfig = blockchain.ChainConfigs{
		ChainId:               421614,
		RPC:                   "https://arbitrum-sepolia.gateway.tenderly.co",
		BridgeContractAddress: "0x7293d8256178331637557753790161f09899a118",
		VoteContractAddress:   "0x329552109874d5050420897035a828285f95f969",
		TantinContractAddress: "0x4F53242025F519072121B775D64D6409F6382D7D",
		UsdtAddress:           "0xF9A555844149DEDC69c3D1F4D3a3a6C083Eb5A4c",
		PrivateKey:            key,
	}
}
