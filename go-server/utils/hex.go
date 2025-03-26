package utils

import (
	"errors"
	"math/big"
	"strings"
)

func HexToBigInt(hexStr string) (*big.Int, error) {
	hexStr = strings.Replace(hexStr, `"`, "", -1)
	if strings.HasPrefix(hexStr, "0x") {
		hexStr = strings.TrimPrefix(hexStr, "0x")
	}

	var bigInt big.Int
	_, ok := bigInt.SetString(hexStr, 16)
	if !ok {
		return nil, errors.New("转换失败")
	}

	return &bigInt, nil
}
