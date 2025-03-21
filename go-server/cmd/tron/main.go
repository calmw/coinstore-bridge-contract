package main

import (
	"coinstore/contract"
	"coinstore/services"
	"fmt"
	"github.com/status-im/keycard-go/hexutils"
	"math/big"
)

func main() {
	services.InitTronEnv()
	bridge, err := contract.NewBridgeTron()
	if err != nil {
		fmt.Println(err)
		return
	}
	//bridge.Init()

	someBytes := hexutils.HexToBytes("ac589789ed8c9d2c61f17b13369864b5f181e58eba230a6ee4ec4c3e7750cd1d")
	txHash, err := bridge.AdminSetResource(big.NewInt(1), hexutils.BytesToHex(someBytes))
	fmt.Println(txHash, err)

	//vote, err := contract.NewVoteTron()
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//vote.Init()

	//tantin, err := contract.NewTanTinTron()
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	////tantin.Init()
	////tantin.AdminSetToken()
	//
	//resourceIdBytes := hexutils.HexToBytes(strings.TrimPrefix(contract.ResourceIdUsdt, "0x"))
	////resourceIdBytes := hexutils.HexToBytes(strings.TrimPrefix(contract.ResourceIdCoin, "0x"))
	//
	//signature, err := utils.RecipientSignature()
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//for {
	//	tantin.Deposit(common.HexToAddress("0x80B27CDE65Fafb1f048405923fD4a624fEa2d1C6"), [32]byte(resourceIdBytes), big.NewInt(1), big.NewInt(10), signature)
	//	time.Sleep(time.Second * 30)
	//}

}
