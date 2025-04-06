package services

import (
	"coinstore/contract"
	"coinstore/utils"
	"os"
)

func InitTantinEnv() {
	coinStoreBridge := os.Getenv("COIN_STORE_BRIDGE")
	privateKeyStr := utils.ThreeDesDecrypt("gZIMfo6LJm6GYXdClPhIMfo6", coinStoreBridge)
	contract.ChainConfig = contract.ChainConfigs{
		BridgeId:              1,
		ChainId:               202502,
		ChainTypeId:           1,
		RPC:                   "https://rpc.tantin.com",
		BridgeContractAddress: "0x66bbc4d0916111aec07892B02d5330bdA7A800DD",
		VoteContractAddress:   "0x99Cb6a45BAB822912AE0519477221ad42C64FF36",
		TantinContractAddress: "0x22050578f91E9663A52D144A39740247FDbdb70A",
		UsdtAddress:           "0x2Bf013133aE838B6934B7F96fd43A10EE3FC3e18",
		UsdcAddress:           "0xe82F64b73D58C85803D78be00407fD44c6DeBe63",
		PrivateKey:            privateKeyStr,
	}
}

func InitBscEnv() {
	coinStoreBridge := os.Getenv("COIN_STORE_BRIDGE")
	privateKeyStr := utils.ThreeDesDecrypt("gZIMfo6LJm6GYXdClPhIMfo6", coinStoreBridge)
	contract.ChainConfig = contract.ChainConfigs{
		BridgeId:              2,
		ChainId:               5611,
		ChainTypeId:           1,
		RPC:                   "https://opbnb-testnet-rpc.bnbchain.org",
		BridgeContractAddress: "0xca540876A5c64eB1A0E51115CF7a5b2687F6e0d2",
		VoteContractAddress:   "0x7EC7dca61c29773466D33aEB9e4f7adbBA960Ca1",
		TantinContractAddress: "0x09125BB80eb099073b392637De2b6f3A42f7D1aC",
		UsdtAddress:           "0x671b21826BdFB241aCa2Dd49dD6C0B96A9309455",
		UsdcAddress:           "0x740b6892bFe90D8b5B926782761e5F9F3eaCC1A1",
		PrivateKey:            privateKeyStr,
	}
}

func InitTronEnv() {
	coinStoreBridge := os.Getenv("COIN_STORE_BRIDGE_TRON")
	privateKeyStr := utils.ThreeDesDecrypt("gZIMfo6LJm6GYXdClPhIMfo6", coinStoreBridge)
	contract.ChainConfig = contract.ChainConfigs{
		BridgeId:              3,
		ChainId:               3448148188,
		ChainTypeId:           2,
		RPC:                   "grpc.nile.trongrid.io:50051",
		BridgeContractAddress: "TWRj9YzaaDki5iUaji3D3jfoUwS3CBLJLq",
		VoteContractAddress:   "TV9ET14nSTmKZ88Dt15USBqKJHfaPsXbXH",
		TantinContractAddress: "TD4HbwLCW558wrBF3Qd5VgC8sG3poejKyS",
		UsdtAddress:           "TXYZopYRdj2D9XRtbG411XZZ3kM5VkAeBf",
		UsdcAddress:           "te1zsjwwtfhsszmajkv45tgg1jqutbhjuz",
		PrivateKey:            privateKeyStr,
	}
}

func InitEthEnv() {
	coinStoreBridge := os.Getenv("COIN_STORE_BRIDGE")
	privateKeyStr := utils.ThreeDesDecrypt("gZIMfo6LJm6GYXdClPhIMfo6", coinStoreBridge)
	contract.ChainConfig = contract.ChainConfigs{
		BridgeId:    4,
		ChainId:     11155111,
		ChainTypeId: 1,
		//RPC:         "https://sepolia.infura.io/v3/732f6502b35c486fb07e333b32e89c04",
		RPC:                   "https://sepolia.drpc.org",
		BridgeContractAddress: "0xC0E8a9C9872A6A7E7F5F2999731dec5d798D82B7",
		VoteContractAddress:   "0x6b66eBFA87AaC1dB355B0ec49ECab7F4b32b1b30",
		TantinContractAddress: "0x77bcb682e01D8763F6757c0D0Beaf577Afcfdf43",
		UsdtAddress:           "0xd78CCb0C79489d8C50AfFE4E881B2C2E93706f8b", // Tether USD
		UsdcAddress:           "0x424cF9Dc24c7c8F006421937d3E288Be84D2daa4", // USD Coin
		PrivateKey:            privateKeyStr,
	}
}
