package utils

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"io"
)

var (
	RequestCont int64
)

// 加密
func AesEny(plaintext []byte, key, iv string) ([]byte, error) {
	var (
		block cipher.Block
		err   error
	)
	//创建aes
	if block, err = aes.NewCipher([]byte(key)); err != nil {
		return nil, err
	}
	//创建ctr
	stream := cipher.NewCTR(block, []byte(iv))
	//加密, src,dst 可以为同一个内存地址
	stream.XORKeyStream(plaintext, plaintext)
	return plaintext, nil
}

// 解密
func AesDecrypt(crpyted, key, iv []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockMode := cipher.NewCBCDecrypter(block, iv)
	origData := make([]byte, len(crpyted))
	blockMode.CryptBlocks(origData, crpyted)
	origData = PKCS5UnPadding(origData)
	return origData, nil
}

// 加密
func AesEncrypt(origData, key, iv []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()
	origData = PKCS5Padding(origData, blockSize)
	blockMode := cipher.NewCBCEncrypter(block, iv)
	dst := make([]byte, len(origData))
	blockMode.CryptBlocks(dst, origData)
	return dst, nil
}

func PKCS5UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}

func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func RandAes128KeyIv() (string, error) {
	keyBytes := [16]byte{}
	_, err := rand.Read(keyBytes[:])
	if err != nil {
		return "", err
	}
	h := hex.EncodeToString(keyBytes[:])

	return h, nil
}

func AESPKCS7Encrypt(plaintext []byte, key []byte) (string, error) {
	// 创建AES cipher
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	// 创建随机初始化向量IV（长度=16字节）
	iv := make([]byte, aes.BlockSize)
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return "", err
	}

	// 加密器配置
	mode := cipher.NewCBCEncrypter(block, iv)
	paddedText := PKCS7Pad(plaintext, aes.BlockSize) // PKCS#7填充

	// 执行加密
	ciphertext := make([]byte, len(paddedText))
	mode.CryptBlocks(ciphertext, paddedText)

	// 组合IV+密文并Base64编码
	combined := append(iv, ciphertext...)
	return base64.StdEncoding.EncodeToString(combined), nil
}

// PKCS#7填充函数
func PKCS7Pad(data []byte, blockSize int) []byte {
	padding := blockSize - (len(data) % blockSize)
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(data, padText...)
}
