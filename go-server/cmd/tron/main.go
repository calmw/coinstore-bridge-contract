package main

import (
	"coinstore/contract"
	"coinstore/services"
	"fmt"
)

func main() {
	services.InitTronEnv()
	bridge, err := contract.NewBridgeTron()
	if err != nil {
		fmt.Println(err)
		return
	}
	//bridge.Init()
	txHash, err := bridge.GrantVoteRole("0x0000000000000000000000000000000000000000000000000000000000000000", contract.OwnerAccount)
	fmt.Println(txHash, err)
	//someBytes := hexutils.HexToBytes("ac589789ed8c9d2c61f17b13369864b5f181e58eba230a6ee4ec4c3e7750cd1d")
	//txHash, err := bridge.AdminSetResource(big.NewInt(1), "0x"+hexutils.BytesToHex(someBytes))
	//fmt.Println(txHash, err)
}
