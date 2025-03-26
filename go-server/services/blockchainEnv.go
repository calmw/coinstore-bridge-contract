package services

import (
	"coinstore/contract"
	"os"
)

func InitTantinEnv() {
	key := os.Getenv("COINSTORE_BRIDGE")
	contract.ChainConfig = contract.ChainConfigs{
		BridgeId:              1,
		ChainId:               202502,
		ChainTypeId:           1,
		RPC:                   "https://rpc.tantin.com",
		BridgeContractAddress: "0x66bbc4d0916111aec07892B02d5330bdA7A800DD",
		VoteContractAddress:   "0x99Cb6a45BAB822912AE0519477221ad42C64FF36",
		TantinContractAddress: "0x22050578f91E9663A52D144A39740247FDbdb70A",
		UsdtAddress:           "0xc7D34B0dC1742De46A346bee415Ad753e0e95370",
		PrivateKey:            key,
	}
}

func InitOpenBnbEnv() {
	key := os.Getenv("COINSTORE_BRIDGE")
	contract.ChainConfig = contract.ChainConfigs{
		BridgeId:              2,
		ChainId:               5611,
		ChainTypeId:           1,
		RPC:                   "https://opbnb-testnet-rpc.bnbchain.org",
		BridgeContractAddress: "0xca540876A5c64eB1A0E51115CF7a5b2687F6e0d2",
		VoteContractAddress:   "0x7EC7dca61c29773466D33aEB9e4f7adbBA960Ca1",
		TantinContractAddress: "0x09125BB80eb099073b392637De2b6f3A42f7D1aC",
		UsdtAddress:           "0xfBe1e02C25a04f6CD6b044F847697b48B3E99a16",
		PrivateKey:            key,
	}
}

func InitTronEnv() {
	key := os.Getenv("COINSTORE_BRIDGE_TRON")
	contract.ChainConfig = contract.ChainConfigs{
		BridgeId:              3,
		ChainId:               3448148188,
		ChainTypeId:           2,
		RPC:                   "grpc.nile.trongrid.io:50051",
		BridgeContractAddress: "TWRj9YzaaDki5iUaji3D3jfoUwS3CBLJLq",
		VoteContractAddress:   "TV9ET14nSTmKZ88Dt15USBqKJHfaPsXbXH",
		TantinContractAddress: "TRPag9uX6pgA98YSKiuWSVNnFZG1EBGiQW",
		UsdtAddress:           "TXYZopYRdj2D9XRtbG411XZZ3kM5VkAeBf",
		PrivateKey:            key,
	}
}
