package main

import (
	"coinstore/contract"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/status-im/keycard-go/hexutils"
	"math/big"
	"strings"
	"time"
)

func main() {
	contract.InitBscEnv()
	//contract.InitBscProdEnv()
	//bridge, err := contract.NewBridge()
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//bridge.Init()

	//bridge.AdminSetResource(
	//	contract.ResourceIdEth,
	//	2,
	//	common.HexToAddress(contract.ChainConfig.WEthAddress),
	//	big.NewInt(100),
	//	false,
	//	false,
	//	false,
	//)
	//
	//vote, err := contract.NewVote()
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//vote.Init()
	//
	tantin, err := contract.NewTanTin()
	if err != nil {
		fmt.Println(err)
		return
	}
	//tantin.Init()

	amount := big.NewInt(2)
	Usdt, err := contract.NewErc20(common.HexToAddress(contract.ChainConfig.UsdtAddress))
	if err != nil {
		fmt.Println(err)
		return
	}
	Usdt.Approve(amount, contract.ChainConfig.TantinContractAddress)
	resourceIdBytes := hexutils.HexToBytes(strings.TrimPrefix(contract.ResourceIdUsdt, "0x"))
	//resourceIdBytes := hexutils.HexToBytes(strings.TrimPrefix(contract.ResourceIdCoin, "0x"))

	for {
		tantin.Deposit(common.HexToAddress("0x80B27CDE65Fafb1f048405923fD4a624fEa2d1C6"), [32]byte(resourceIdBytes), big.NewInt(3), amount)
		time.Sleep(time.Second * 30)
	}
}
