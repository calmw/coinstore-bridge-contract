package abi

import (
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
	"strings"
)

func TronPriceSignature(chainId, price *big.Int, tokenAddress common.Address) ([]byte, error) {
	contractAbi, _ := abi.JSON(strings.NewReader(PriceSig))
	parameterBytes, _ := contractAbi.Pack("checkPriceSignature",
		chainId,
		price,
		tokenAddress,
	)
	return GeneratePriceSignatureTron(parameterBytes[4:])
}
