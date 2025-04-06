package abi

import (
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"strings"
)

func TantinDepositSignature(recipient string) ([]byte, error) {
	contractAbi, _ := abi.JSON(strings.NewReader(TantinSig))
	parameterBytes, _ := contractAbi.Pack("checkDepositSignature",
		common.HexToAddress(recipient),
	)
	return GenerateSignature(parameterBytes[4:])
}

func TantinAdminSetTokenSignature(resourceID [32]byte, assetsType uint8, tokenAddress common.Address, burnable, mintable, pause bool) ([]byte, error) {
	contractAbi, _ := abi.JSON(strings.NewReader(TantinSig))
	parameterBytes, _ := contractAbi.Pack("checkAdminSetTokenSignature",
		resourceID,
		assetsType,
		tokenAddress,
		burnable,
		mintable,
		pause,
	)
	return GenerateSignature(parameterBytes[4:])
}

func TantinAdminSetEnvSignature(bridgeAddress common.Address) ([]byte, error) {
	contractAbi, _ := abi.JSON(strings.NewReader(TantinSig))
	parameterBytes, _ := contractAbi.Pack("checkAdminSetEnvSignature",
		bridgeAddress,
	)
	return GenerateSignature(parameterBytes[4:])
}
