package abi

import (
	"fmt"
	"testing"
)

func TestTantinDepositSignature(t *testing.T) {
	signature, err := TantinDepositSignature("0x80B27CDE65Fafb1f048405923fD4a624fEa2d1C6")
	fmt.Println(fmt.Sprintf("0x%x", signature), err)
}
