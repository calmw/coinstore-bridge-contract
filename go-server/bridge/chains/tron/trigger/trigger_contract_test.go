package trigger

import (
	"fmt"
	"testing"
)

func TestGetSigNonce(t *testing.T) {
	got, err := GetSigNonce("THqCU4BGnFRi6XGNa7vkM1BLaAFx6JQmYo", "TFBymbm7LrbRreGtByMPRD2HUyneKabsqb")
	fmt.Println(got, err)
}
