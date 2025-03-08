package utils

import (
	"encoding/base64"
	"github.com/forgoer/openssl"
)

func ThreeDesEncrypt(key, data string) string {
	keyBytes := []byte(key)
	ecbEncrypt, _ := openssl.Des3ECBEncrypt([]byte(data), keyBytes, openssl.PKCS7_PADDING)
	return base64.StdEncoding.EncodeToString(ecbEncrypt)
}

func ThreeDesDecrypt(key, data string) string {
	keyBytes := []byte(key)
	decodeString, _ := base64.StdEncoding.DecodeString(data)
	decrypt, _ := openssl.Des3ECBDecrypt([]byte(decodeString), keyBytes, openssl.PKCS7_PADDING)
	return string(decrypt)
}
