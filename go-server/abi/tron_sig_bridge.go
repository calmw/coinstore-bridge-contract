package abi

import (
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
	"strings"
)

func BridgeAdminSetResourceSignatureTron(sigNonce, chainId *big.Int, resourceID [32]byte, assetsType uint8, tokenAddress, tantinAddress common.Address, fee *big.Int, pause bool, burnable bool, mintable bool) ([]byte, error) {
	contractAbi, _ := abi.JSON(strings.NewReader(BridgeSig))
	parameterBytes, _ := contractAbi.Pack("checkAdminSetResourceSignature",
		sigNonce,
		chainId,
		resourceID,
		assetsType,
		tokenAddress,
		fee,
		pause,
		burnable,
		mintable,
		tantinAddress,
	)
	return GenerateSignatureTron(parameterBytes[4:])
}

func BridgeAdminSetEnvSignatureTron(sigNonce *big.Int, voteAddress common.Address, chainId, chainType *big.Int) ([]byte, error) {
	contractAbi, _ := abi.JSON(strings.NewReader(BridgeSig))
	parameterBytes, _ := contractAbi.Pack("checkAdminSetEnvSignature",
		sigNonce,
		chainId,
		voteAddress,
		chainId,
		chainType,
	)
	return GenerateSignatureTron(parameterBytes[4:])
}

func VoteAdminSetEnvSignatureTron(sigNonce *big.Int, bridgeAddress, tantinAddress common.Address, expiry, relayerThreshold *big.Int) ([]byte, error) {
	contractAbi, _ := abi.JSON(strings.NewReader(VoteSig))
	parameterBytes, _ := contractAbi.Pack("checkAdminSetEnvSignature",
		sigNonce,
		bridgeAddress,
		tantinAddress,
		expiry,
		relayerThreshold,
	)
	return GenerateSignatureTron(parameterBytes[4:])
}

func TantinAdminSetEnvSignatureTron(sigNonce *big.Int, feeAddress, bridgeAddress common.Address) ([]byte, error) {
	contractAbi, _ := abi.JSON(strings.NewReader(TantinSig))
	parameterBytes, _ := contractAbi.Pack("checkAdminSetEnvSignature",
		sigNonce,
		feeAddress,
		bridgeAddress,
	)
	return GenerateSignatureTron(parameterBytes[4:])
}

func TantinAdminSetTokenSignatureTron(sigNonce, chainId *big.Int, resourceID [32]byte, assetsType uint8, tokenAddress common.Address, burnable, mintable, pause bool) ([]byte, error) {
	contractAbi, _ := abi.JSON(strings.NewReader(TantinSig))
	parameterBytes, _ := contractAbi.Pack("checkAdminSetTokenSignature",
		sigNonce,
		chainId,
		resourceID,
		assetsType,
		tokenAddress,
		burnable,
		mintable,
		pause,
	)
	return GenerateSignatureTron(parameterBytes[4:])
}

func TantinDepositSignatureTron(recipientAddress common.Address) ([]byte, error) {
	contractAbi, _ := abi.JSON(strings.NewReader(TantinSig))
	parameterBytes, _ := contractAbi.Pack("checkDepositSignature",
		recipientAddress,
	)
	return GenerateSignatureTron(parameterBytes[4:])
}
