package main

import (
	"coinstore/contract"
	"coinstore/services"
	"fmt"
)

func main() {
	services.InitTronEnv()
	//bridgeaa, err := contract.NewBridgeTron()
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//bridgeaa.Init()
	//
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
	tantin.Init()

	////for {
	//toAddress, err := address.Base58ToAddress("TFBymbm7LrbRreGtByMPRD2HUyneKabsqb")
	//fmt.Println(toAddress.String())
	//txHash, err := tantin.Deposit("2", strings.TrimPrefix(contract.ResourceIdUsdt, "0x"), toAddress.String(), "0x00", big.NewInt(1))
	//
	//fmt.Println(txHash, err)
	//time.Sleep(time.Second * 30)
	//}

}
