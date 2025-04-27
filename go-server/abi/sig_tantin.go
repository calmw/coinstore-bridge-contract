package abi

import (
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
	"strings"
)

func TantinDepositSignature(recipient common.Address) ([]byte, error) {
	contractAbi, _ := abi.JSON(strings.NewReader(TantinSig))
	parameterBytes, _ := contractAbi.Pack("checkDepositSignature",
		recipient,
	)
	return GenerateSignature(parameterBytes[4:])
}

func TantinAdminSetTokenSignature(sigNonce, chainId *big.Int, resourceID [32]byte, assetsType uint8, tokenAddress common.Address, burnable, mintable, pause bool) ([]byte, error) {
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
	return GenerateSignature(parameterBytes[4:])
}

func TantinAdminSetEnvSignature(sigNonce, chainId *big.Int, feeAddress, serverAddress, bridgeAddress common.Address) ([]byte, error) {
	contractAbi, _ := abi.JSON(strings.NewReader(TantinSig))
	parameterBytes, err := contractAbi.Pack("checkAdminSetEnvSignature",
		sigNonce,
		chainId,
		feeAddress,
		serverAddress,
		bridgeAddress,
	)
	fmt.Println(err)
	return GenerateSignature(parameterBytes[4:])
}
