package utils

import (
	"fmt"
	"testing"
)

func TestThreeDes(t *testing.T) {
	key := "gZIMfo6LJm6GYXdClPhIMfo5"
	data := ""
	encrypt := ThreeDesEncrypt(key, data)
	fmt.Println(encrypt)
	res := ThreeDesDecrypt(key, encrypt)
	fmt.Println(res)
}
