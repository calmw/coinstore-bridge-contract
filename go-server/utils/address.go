package utils

import (
	"github.com/fbsobreira/gotron-sdk/pkg/address"
	"strings"
)

func TronToEth(addr string) (string, error) {
	toAddress, err := address.Base58ToAddress(addr)
	if err != nil {
		return "", err
	}
	return toAddress.Hex(), nil
}
func EthToTron(addr string) (string, error) {
	tokenAddress := "0x41" + strings.TrimPrefix(addr, "0x")
	tokenAddress = address.HexToAddress(tokenAddress).String()
	return tokenAddress, nil
}
