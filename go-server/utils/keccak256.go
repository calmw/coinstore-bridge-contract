package utils

import (
	"github.com/ethereum/go-ethereum/crypto"
)

func Keccak256(data []byte) [32]byte {
	return crypto.Keccak256Hash(data)
}
