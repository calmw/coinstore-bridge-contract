package main

import (
	"coinstore/contract"
	"fmt"
	"math/big"
	"time"
)

func main() {
	contract.InitTronEnv()
	bridge, err := contract.NewBridgeTron()
	if err != nil {
		fmt.Println(err)
		return
	}
	bridge.Init()

	vote, err := contract.NewVoteTron()
	if err != nil {
		fmt.Println(err)
		return
	}
	vote.Init()

	tantin, err := contract.NewTanTinTron()
	if err != nil {
		fmt.Println(err)
		return
	}
	//tantin.Init()

	for {
		txHash, err := tantin.Deposit(big.NewInt(2), big.NewInt(1), contract.ResourceIdUsdt, "TFBymbm7LrbRreGtByMPRD2HUyneKabsqb")
		fmt.Println(txHash, err)
		time.Sleep(time.Second * 60)
	}

}
