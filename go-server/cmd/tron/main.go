package main

import (
	"coinstore/contract"
	"coinstore/services"
	"fmt"
	"math/big"
	"strings"
	"time"
)

func main() {
	services.InitTronEnv()
	//bridge, err := contract.NewBridgeTron()
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	////bridge.Init()
	//
	//someBytes := hexutils.HexToBytes("ac589789ed8c9d2c61f17b13369864b5f181e58eba230a6ee4ec4c3e7750cd1d")
	//txHash, err := bridge.AdminSetResource(big.NewInt(1), hexutils.BytesToHex(someBytes))
	//fmt.Println(txHash, err)

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
	//tantin.AdminSetToken()

	//signature, err := utils.RecipientSignature()
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	for {
		txHash, err := tantin.Deposit("1", strings.TrimPrefix(contract.ResourceIdUsdt, "0x"), "TFBymbm7LrbRreGtByMPRD2HUyneKabsqb", "0x6537", big.NewInt(1))

		fmt.Println(txHash, err)
		time.Sleep(time.Second * 30)
	}

}
