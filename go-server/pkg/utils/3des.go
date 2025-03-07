package utils

import (
	"encoding/base64"
	"fmt"
	"github.com/forgoer/openssl"
)

func Ccc() {
	src := []byte("0,126")
	iv := []byte("1234567890123456")
	key := []byte("gZAnoptLJm6GYXdClPhIMfo6")
	//3DES-ECB:
	ecbEncrypt, err := openssl.Des3ECBEncrypt(src, key, openssl.PKCS7_PADDING)
	fmt.Println(ecbEncrypt, err, base64.StdEncoding.EncodeToString(ecbEncrypt))
	decrypt, err := openssl.Des3ECBDecrypt(ecbEncrypt, key, openssl.PKCS7_PADDING)

	fmt.Println(decrypt, string(decrypt), err, 88888)
	//3DES-CBC:

	encrypt, err := openssl.Des3CBCEncrypt(src, key, iv, openssl.PKCS7_PADDING)
	fmt.Println(encrypt, err)
	//openssl.Des3CBCDecrypt(src, key, iv, openssl.PKCS7_PADDING)
}
