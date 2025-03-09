package utils

import (
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/status-im/keycard-go/hexutils"
	"strings"
)

func Keccak256(data string) string {
	return strings.ToLower("0x" + hexutils.BytesToHex(crypto.Keccak256([]byte(data))))
}

func Hash(data []byte) [32]byte {
	return crypto.Keccak256Hash(data)
}
