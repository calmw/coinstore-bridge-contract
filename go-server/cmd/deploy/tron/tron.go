package tron

import (
	"coinstore/contract"
	"fmt"
	"os"
)

func InitTron(prvKey, adminAddress, feeAddress, serverAddress, realyerOneAddress, realyerTwoAddress, realyerThreeAddress string, fee uint64) {
	contract.InitTronEnv()
	err := os.Setenv("TB_KEY", prvKey)
	if err != nil {
		fmt.Println(err)
		return
	}
	bridge, err := contract.NewBridgeTron()
	if err != nil {
		fmt.Println(err)
		return
	}
	bridge.Init(adminAddress, fee)

	vote, err := contract.NewVoteTron()
	if err != nil {
		fmt.Println(err)
		return
	}
	vote.Init(realyerOneAddress, realyerTwoAddress, realyerThreeAddress)

	tantin, err := contract.NewTanTinTron()
	if err != nil {
		fmt.Println(err)
		return
	}
	tantin.Init()
}
