package utils

import (
	"fmt"
	"testing"
)

func TestThreeDes(t *testing.T) {
	key := "gZAnoptLJm6GYXdClPhIMfo6"
	data := "some string"
	encrypt := ThreeDesEncrypt(key, data)
	res := ThreeDesDecrypt(key, encrypt)
	fmt.Println(res)
}
