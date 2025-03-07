package utils

import (
	"encoding/base64"
	"fmt"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/forgoer/openssl"
	"github.com/status-im/keycard-go/hexutils"
	"strings"
)

func Keccak256(data string) string {
	return strings.ToLower("0x" + hexutils.BytesToHex(crypto.Keccak256([]byte(data))))
}

func ThreeDesEncrypt(data string) string {
	key := []byte("gZAnoptLJm6GYXdClPhIMfo6")
	//3DES-ECB:
	ecbEncrypt, _ := openssl.Des3ECBEncrypt([]byte(data), key, openssl.PKCS7_PADDING)
	return base64.StdEncoding.EncodeToString(ecbEncrypt)
}

func ThreeDesDecrypt(data string) string {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("ThreeDesDecrypt panic触发错误", err)
		}
	}()
	key := []byte("gZAnoptLJm6GYXdClPhIMfo6")
	//3DES-ECB:
	decodeString, _ := base64.StdEncoding.DecodeString(data)
	decrypt, _ := openssl.Des3ECBDecrypt([]byte(decodeString), key, openssl.PKCS7_PADDING)
	return string(decrypt)
}
