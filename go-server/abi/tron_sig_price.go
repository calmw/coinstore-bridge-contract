package abi

import (
	"github.com/ethereum/go-ethereum/accounts/abi"
	"math/big"
	"strings"
)

func TronPriceSignature(chainId, price, priceTimestamp *big.Int) ([]byte, error) {
	contractAbi, _ := abi.JSON(strings.NewReader(PriceSig))
	parameterBytes, _ := contractAbi.Pack("checkPriceSignature",
		chainId,
		price,
		priceTimestamp,
	)
	return GeneratePriceSignatureTron(parameterBytes[4:])
}
