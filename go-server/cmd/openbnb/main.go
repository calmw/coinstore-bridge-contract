package main

import (
	"coinstore/pkg/blockchain"
	"coinstore/services"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
)

func main() {
	services.InitOpenBnbEnv()
	//bridge, err := blockchain.NewBridge()
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	////bridge.Init()
	//resourceId := "ac589789ed8c9d2c61f17b13369864b5f181e58eba230a6ee4ec4c3e7750cd1d"
	//resourceIdBytes := hexutils.HexToBytes(resourceId)
	//bridge.AdminSetResource(big.NewInt(0), [4]byte(resourceIdBytes))
	//
	//vote, err := blockchain.NewVote()
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//vote.Init()

	tantin, err := blockchain.NewTanTin()
	if err != nil {
		fmt.Println(err)
		return
	}
	//tantin.Init()
	//tantin.AdminSetToken()
	tantin.Deposit(common.HexToAddress("0x80B27CDE65Fafb1f048405923fD4a624fEa2d1C6"), big.NewInt(2), big.NewInt(2))
}
