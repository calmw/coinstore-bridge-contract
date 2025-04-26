package abi

import (
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
	"strings"
)

func BridgeAdminSetResourceSignature(sigNonce, chainId *big.Int, resourceID [32]byte, assetsType uint8, tokenAddress, tantinAddress common.Address, fee *big.Int, pause bool, burnable bool, mintable bool) ([]byte, error) {
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
	return GenerateSignature(parameterBytes[4:])
}

func BridgeAdminSetEnvSignature(sigNonce *big.Int, voteAddress common.Address, chainId, chainType *big.Int) ([]byte, error) {
	contractAbi, _ := abi.JSON(strings.NewReader(BridgeSig))
	parameterBytes, _ := contractAbi.Pack("checkAdminSetEnvSignature",
		sigNonce,
		voteAddress,
		chainId,
		chainType,
	)
	return GenerateSignature(parameterBytes[4:])
}
