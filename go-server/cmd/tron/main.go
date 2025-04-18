package main

import (
	"coinstore/contract"
	"fmt"
	"math/big"
	"time"
)

// https://api.trongrid.io
// https://api.trongrid.io/jsonrpc
// https://api.shasta.trongrid.io/jsonrpc

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
		//txHash, err := tantin.Deposit(big.NewInt(1), big.NewInt(3), contract.ResourceIdUsdt, "TQxhW4iv7BvT63qdnmx76GZK5FViy4qMfh")
		txHash, err := tantin.Deposit(big.NewInt(1), big.NewInt(3), contract.ResourceIdUsdt, "TEkkeJsMAQD18HqodvYLZ91BJRv1kG1sN7")
		fmt.Println(txHash, err)
		time.Sleep(time.Second * 60)
	}

}
