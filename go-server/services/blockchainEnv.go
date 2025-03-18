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
		BridgeContractAddress: "0xb24401133Dd5Ea36f057FF9c9fC325eaFA1C3905",
		VoteContractAddress:   "0xc7D34B0dC1742De46A346bee415Ad753e0e95370",
		TantinContractAddress: "0x2bcf35e0a8B731b1EeD1C94bEA298bD20c7f89E0",
		UsdtAddress:           "0x66bbc4d0916111aec07892B02d5330bdA7A800DD",
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
		BridgeContractAddress: "0x993FEF5848B8FFee9859cE6d7a2A7f07d9122cce",
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
		BridgeContractAddress: "TH9j26HtSzmRQC11xSMarhDFVnJsLhxRLr",
		VoteContractAddress:   "TEgAohYmMTz8sRRZSf5ht69Q3jPBA4vKSz",
		TantinContractAddress: "TBVtUdewEJJ1zYif9tmYUC2sgERDYfWUUS",
		UsdtAddress:           "TU6UjUJadm8TungBHvL4n9apv8Jns4wJiz",
		PrivateKey:            key,
	}
}
