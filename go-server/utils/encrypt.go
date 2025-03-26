package utils

import (
	"bytes"
	"crypto/md5"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/hex"
	"encoding/pem"
	"errors"
	"fmt"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

// Sign ..
func Sign(rawBody interface{}, t int64, token string) string {
	s := fmt.Sprintf("%s%v%s", rawBody, t, token)
	return MD5(s)
}

// CalcMd5 ..
func MD5(text string) string {
	bytes := md5.Sum([]byte(text))
	return strings.ToLower(hex.EncodeToString(bytes[:]))
}

// EncryptPassword ..
func EncryptPassword(pass string) (encodePW string) {
	hash, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.MinCost)
	if err != nil {
		fmt.Println(err)
	}

	encodePW = string(hash)
	return
}

// ComparePassword ..
func ComparePassword(encryptedPassword, password string) bool {
	return bcrypt.CompareHashAndPassword([]byte(encryptedPassword), []byte(password)) == nil
}

// RsaEncryptToString ..
func RsaEncryptToString(origData []byte, publicKey *rsa.PublicKey) (string, error) {
	bytes, e := RsaEncrypt(origData, publicKey)
	if e != nil {
		return "", e
	}
	return base64.StdEncoding.EncodeToString(bytes), nil
}

// RsaDecryptFromString ..
func RsaDecryptFromString(base64String string, privateKey *rsa.PrivateKey) ([]byte, error) {
	bytes, e := base64.StdEncoding.DecodeString(base64String)
	if e != nil {
		return nil, e
	}
	return RsaDecrypt(bytes, privateKey)
}

// RsaEncrypt 加密
func RsaEncrypt(origData []byte, publicKey *rsa.PublicKey) ([]byte, error) {

	return rsa.EncryptPKCS1v15(rand.Reader, publicKey, origData)
}

// RsaDecrypt 解密
func RsaDecrypt(ciphertext []byte, privateKey *rsa.PrivateKey) ([]byte, error) {

	return rsa.DecryptPKCS1v15(rand.Reader, privateKey, ciphertext)
}

/*// RsaPrivateEncrypt
func RsaPrivateEncrypt(input string) string {
	//pubErr, priErr := codec.RSA.Init(pUBLICK_KEY, pRIVATE_KEY)
	str, _ := codec.RSA.String(input, codec.MODE_PRIKEY_ENCRYPT)
	return str
}

// RsaPublicDecrypt
func RsaPublicDecrypt(input string) string {
	str, _ := codec.RSA.String(input, codec.MODE_PUBKEY_DECRYPT)
	return str
}
*/

// ParsePrivateKey 解析私钥
func ParsePrivateKey(in string) (privateKey *rsa.PrivateKey, err error) {
	block, _ := pem.Decode([]byte(in))
	if block == nil {
		err = errors.New("private key error")
		return
	}
	privateKey, err = x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return
	}
	return
}

// ParsePublicKey 解析公钥
func ParsePublicKey(in string) (publicKey *rsa.PublicKey, err error) {
	block, _ := pem.Decode([]byte(in))
	if block == nil {
		err = errors.New("public key error")
		return
	}
	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	publicKey = pubInterface.(*rsa.PublicKey)
	return
}

func GenRsaKeyPair(bits int) (priKey, pubKey string, err error) {
	privateKey, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		return

	}
	derStream := x509.MarshalPKCS1PrivateKey(privateKey)
	block := &pem.Block{
		Type:  "rsa PRIVATE KEY",
		Bytes: derStream,
	}
	var buf = bytes.NewBufferString("")
	err = pem.Encode(buf, block)
	if err != nil {
		return
	}
	priKey = buf.String()
	// 生成公钥文件
	publicKey := &privateKey.PublicKey
	derPkix, err := x509.MarshalPKIXPublicKey(publicKey)
	if err != nil {
		return
	}
	block = &pem.Block{Type: "PUBLIC KEY",
		Bytes: derPkix,
	}
	buf.Reset()
	err = pem.Encode(buf, block)
	if err != nil {
		return
	}
	pubKey = buf.String()
	return
}
