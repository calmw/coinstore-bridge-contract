package abi

import (
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
	"strings"
)

func VoteAdminRemoveRelayerSignature(relayerAddress common.Address) ([]byte, error) {
	contractAbi, _ := abi.JSON(strings.NewReader(TantinSig))
	parameterBytes, _ := contractAbi.Pack("checkAdminRemoveRelayerSignature",
		relayerAddress,
	)
	return GenerateSignature(parameterBytes[4:])
}

func VoteAdminAddRelayerSignature(relayerAddress common.Address) ([]byte, error) {
	contractAbi, _ := abi.JSON(strings.NewReader(TantinSig))
	parameterBytes, _ := contractAbi.Pack("checkAdminAddRelayerSignature",
		relayerAddress,
	)
	return GenerateSignature(parameterBytes[4:])
}

func VoteAdminChangeRelayerThresholdSignature(newThreshold *big.Int) ([]byte, error) {
	contractAbi, _ := abi.JSON(strings.NewReader(TantinSig))
	parameterBytes, _ := contractAbi.Pack("checkAdminChangeRelayerThresholdSignature",
		newThreshold,
	)
	return GenerateSignature(parameterBytes[4:])
}

func VoteAdminSetEnvSignature(tantinBridgeAddress common.Address, expiry, relayerThreshold *big.Int) ([]byte, error) {
	contractAbi, _ := abi.JSON(strings.NewReader(TantinSig))
	parameterBytes, _ := contractAbi.Pack("checkAdminSetEnvSignature",
		tantinBridgeAddress,
		expiry,
		relayerThreshold,
	)
	return GenerateSignature(parameterBytes[4:])
}
