package main

import (
	"coinstore/contract"
	"fmt"
	"github.com/fbsobreira/gotron-sdk/pkg/address"
	"math/big"
	"time"
)

func main() {
	contract.InitTronEnv()
	//bridge, err := contract.NewBridgeTron()
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//bridge.Init()

	//vote, err := contract.NewVoteTron()
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//vote.Init()

	tantin, err := contract.NewTanTinTron()
	if err != nil {
		fmt.Println(err)
		return
	}
	//tantin.Init()

	for {
		toAddress, err := address.Base58ToAddress("TFBymbm7LrbRreGtByMPRD2HUyneKabsqb")
		fmt.Println(toAddress.String())
		txHash, err := tantin.Deposit(big.NewInt(1), big.NewInt(1), contract.ResourceIdUsdt, "TFBymbm7LrbRreGtByMPRD2HUyneKabsqb")

		fmt.Println(txHash, err)
		time.Sleep(time.Second * 30)
	}

}
