package abi

import (
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"math/big"
	"strings"
)

func EvmPriceSignature(chainId, price, priceTimestamp *big.Int) ([]byte, error) {
	contractAbi, _ := abi.JSON(strings.NewReader(PriceSig))
	parameterBytes, _ := contractAbi.Pack("checkPriceSignature",
		chainId,
		price,
		priceTimestamp,
	)
	fmt.Println(chainId, price, priceTimestamp)
	return GeneratePriceSignature(parameterBytes[4:])
}
