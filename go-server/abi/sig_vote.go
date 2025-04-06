package abi

import (
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
	"strings"
)

func VoteAdminRemoveRelayerSignature(sigNonce *big.Int, relayerAddress common.Address) ([]byte, error) {
	contractAbi, _ := abi.JSON(strings.NewReader(VoteSig))
	parameterBytes, _ := contractAbi.Pack("checkAdminRemoveRelayerSignature",
		sigNonce,
		relayerAddress,
	)
	return GenerateSignature(parameterBytes[4:])
}

func VoteAdminAddRelayerSignature(sigNonce *big.Int, relayerAddress common.Address) ([]byte, error) {
	contractAbi, _ := abi.JSON(strings.NewReader(VoteSig))
	parameterBytes, _ := contractAbi.Pack("checkAdminAddRelayerSignature",
		sigNonce,
		relayerAddress,
	)
	return GenerateSignature(parameterBytes[4:])
}

func VoteAdminChangeRelayerThresholdSignature(sigNonce *big.Int, newThreshold *big.Int) ([]byte, error) {
	contractAbi, _ := abi.JSON(strings.NewReader(VoteSig))
	parameterBytes, _ := contractAbi.Pack("checkAdminChangeRelayerThresholdSignature",
		sigNonce,
		newThreshold,
	)
	return GenerateSignature(parameterBytes[4:])
}

func VoteAdminSetEnvSignature(sigNonce *big.Int, tantinBridgeAddress common.Address, expiry, relayerThreshold *big.Int) ([]byte, error) {
	contractAbi, _ := abi.JSON(strings.NewReader(VoteSig))
	parameterBytes, _ := contractAbi.Pack("checkAdminSetEnvSignature",
		sigNonce,
		tantinBridgeAddress,
		expiry,
		relayerThreshold,
	)
	return GenerateSignature(parameterBytes[4:])
}
