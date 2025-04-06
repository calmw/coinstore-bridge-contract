package utils

import (
	"fmt"
	"testing"
)

func TestRecipientSignature(t *testing.T) {

	got, err := RecipientSignature()
	fmt.Println(got, err)
}

func TestTaa(t *testing.T) {
	Taa()
}
