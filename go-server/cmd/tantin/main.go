package main

import (
	"coinstore/contract"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
)

func main() {
	contract.InitTantinEnv()

	bridge, err := contract.NewBridge()
	if err != nil {
		fmt.Println(err)
		return
	}
	bridge.Init()
	bridge.AdminSetResource(
		contract.ResourceIdEth,
		2,
		common.HexToAddress(contract.ZeroAddress),
		big.NewInt(100),
	)
	bridge.AdminSetResource(
		contract.ResourceIdWeth,
		2,
		common.HexToAddress(contract.ChainConfig.UsdtAddress),
		big.NewInt(100),
	)

	//vote, err := contract.NewVote()
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//vote.Init()

	//tantin, err := contract.NewTanTin()
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	////tantin.Init()
	////tantin.AdminSetToken(contract.ResourceIdUsdt, 2, common.HexToAddress(contract.ChainConfig.UsdtAddress), false, false, false)
	//
	//resourceIdBytes := hexutils.HexToBytes(strings.TrimPrefix(contract.ResourceIdUsdt, "0x"))
	////resourceIdBytes := hexutils.HexToBytes(strings.TrimPrefix(contract.ResourceIdCoin, "0x"))
	//
	//for {
	//	tantin.Deposit(common.HexToAddress("0x80B27CDE65Fafb1f048405923fD4a624fEa2d1C6"), [32]byte(resourceIdBytes), big.NewInt(3), big.NewInt(2))
	//	time.Sleep(time.Second * 30)
	//}
}
