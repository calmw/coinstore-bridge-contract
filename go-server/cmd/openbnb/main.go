package main

import (
	"coinstore/contract"
	"coinstore/services"
	"coinstore/utils"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/status-im/keycard-go/hexutils"
	"math/big"
	"strings"
	"time"
)

func main() {
	services.InitOpenBnbEnv()
	//bridge, err := contract.NewBridge()
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//bridge.Init()
	//someBytes := hexutils.HexToBytes("09c5eabe")
	//bridge.AdminSetResource(big.NewInt(1), [4]byte(someBytes))

	//vote, err := contract.NewVote()
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//vote.Init()

	tantin, err := contract.NewTanTin()
	if err != nil {
		fmt.Println(err)
		return
	}
	//tantin.Init()
	//tantin.AdminSetToken()

	resourceIdBytes := hexutils.HexToBytes(strings.TrimPrefix(contract.ResourceIdUsdt, "0x"))
	//resourceIdBytes := hexutils.HexToBytes(strings.TrimPrefix(contract.ResourceIdCoin, "0x"))

	signature, err := utils.RecipientSignature()
	if err != nil {
		fmt.Println(err)
		return
	}
	for {
		tantin.Deposit(common.HexToAddress("0x80B27CDE65Fafb1f048405923fD4a624fEa2d1C6"), [32]byte(resourceIdBytes), big.NewInt(1), big.NewInt(10), signature)
		time.Sleep(time.Second * 30)
	}
}
