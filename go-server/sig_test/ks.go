package main

import (
	"coinstore/cmd/deploy/tron"
	"fmt"
)

func KeyStore(key, passphrase string) {
	_, ka, err := tron.InitKeyStore(key, passphrase)
	fmt.Println(err)
	fmt.Println(fmt.Sprintf("账户地址:%s", ka.Address.String()))
}

//func InitKeyStore(privateKeyStr, passphrase string) (*keystore.KeyStore, *keystore.Account, error) {
//	keys.AddNewKey(passphrase)
//
//	return ks, &account, nil
//}
