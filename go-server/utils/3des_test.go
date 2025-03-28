package utils

import (
	"fmt"
	"testing"
)

func TestThreeDes(t *testing.T) {
	key := "gZIMfo6LJm6GYXdClPhIMfo6"
	//data := "a6e3eea0a20d89251acfcbd765d7ff6cda65f07854389a31c6f3ee0f97e45df3"
	data := "af056e353eaddc842029822cc4c7285ef379880eaae25a1f8d8a16da5f41ff2b"
	encrypt := ThreeDesEncrypt(key, data)
	fmt.Println(encrypt)
	res := ThreeDesDecrypt(key, encrypt)
	fmt.Println(res)
}
