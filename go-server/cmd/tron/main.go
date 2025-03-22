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
		txHash, err := tantin.Deposit("2", strings.TrimPrefix(contract.ResourceIdUsdt, "0x"), "TFBymbm7LrbRreGtByMPRD2HUyneKabsqb", "0x00", big.NewInt(1))

		fmt.Println(txHash, err)
		time.Sleep(time.Second * 30)
	}

}
