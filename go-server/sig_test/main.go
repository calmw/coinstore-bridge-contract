package main

import (
	"coinstore/utils"
	"fmt"
)

func main() {
	//tx, err := TransferCoin()
	//fmt.Println(tx, err)
	//GenerateEvmAccount()
	//err := TronTest()
	//fmt.Println(err)

	Aa()

	//KeyStore("4d69dca2ebc32e9d448a1c6d6fa3da61fc537077279c3ab6c6702adfe70dd8a2", "123456")
}

func Aa() {
	eth, err := utils.TronToEth("TTDXsZRNQP2JajbdUBvUHV8sSqDApypAgH")
	fmt.Println(eth, err)
}
