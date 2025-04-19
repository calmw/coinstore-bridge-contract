package utils

import (
	"github.com/calmw/tron-sdk/pkg/address"
	"strings"
)

func TronToEth(addr string) (string, error) {
	toAddress, err := address.Base58ToAddress(addr)
	if err != nil {
		return "", err
	}
	return "0x" + strings.TrimPrefix(toAddress.Hex(), "0x41"), nil
}
func EthToTron(addr string) (string, error) {
	toAddress := address.HexToAddress("0x41" + strings.TrimPrefix(addr, "0x"))
	return toAddress.String(), nil
}
