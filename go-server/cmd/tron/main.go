package main

import (
	"coinstore/cmd/deploy/tron"
	"coinstore/contract"
	"fmt"
	"math/big"
	"os"
	"time"
)

// https://api.trongrid.io
// https://api.trongrid.io/jsonrpc
// https://api.shasta.trongrid.io/jsonrpc

func main() {
	//contract.InitTronEnvProd()
	contract.InitTronEnv()
	//bridge, err := contract.NewBridgeTron()
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//bridge.Init()
	//
	//vote, err := contract.NewVoteTron()
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//vote.Init()
	prvKey := os.Getenv("COINSTORE_BRIDGE_TRON_LOCAL")
	fmt.Println(prvKey)
	ks, ka, err := tron.InitKeyStore(prvKey, "123456")
	if err != nil {
		fmt.Println(err)
		return
	}
	tantin, err := contract.NewTanTinTron(ka, ks, "123456")
	if err != nil {
		fmt.Println(err)
		return
	}
	//tantin.Init()

	for {
		//txHash, err := tantin.Deposit(big.NewInt(1), big.NewInt(3), contract.ResourceIdUsdt, "TQxhW4iv7BvT63qdnmx76GZK5FViy4qMfh")
		//txHash, err := tantin.Deposit(big.NewInt(1), big.NewInt(3), contract.ResourceIdUsdc, "TEkkeJsMAQD18HqodvYLZ91BJRv1kG1sN7")
		//txHash, err := tantin.Deposit(big.NewInt(1), big.NewInt(3), contract.ResourceIdEth, "TEkkeJsMAQD18HqodvYLZ91BJRv1kG1sN7")
		txHash, err := tantin.Deposit(
			big.NewInt(12302),
			big.NewInt(3),
			big.NewInt(1e6),
			big.NewInt(time.Now().Unix()),
			contract.ResourceIdUsdt,
			"TEkkeJsMAQD18HqodvYLZ91BJRv1kG1sN7",
		)
		fmt.Println(txHash, err)
		time.Sleep(time.Second * 60)
	}

}
