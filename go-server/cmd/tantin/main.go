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
	contract.InitTantinEnv()
	//adminAddress := "0xa47142f08f859aCeb2127C6Ab66eC8c8bc4FFBA9"
	//feeAddress := "0xa47142f08f859aCeb2127C6Ab66eC8c8bc4FFBA9"
	//serverAddress := "0xa47142f08f859aCeb2127C6Ab66eC8c8bc4FFBA9"
	//bridge, err := contract.NewBridge()
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//bridge.Init(adminAddress)
	//bridge.AdminSetResource(
	//	contract.ResourceIdEth,
	//	2,
	//	common.HexToAddress(contract.ChainConfig.WEthAddress),
	//	big.NewInt(1e18),
	//	big.NewInt(100),
	//	false,
	//	false,
	//	false,
	//)
	//bridge.AdminSetResource(
	//	contract.ResourceIdWeth,
	//	2,
	//	common.HexToAddress(contract.ChainConfig.WEthAddress),
	//	big.NewInt(1e18),
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

	tantin, err := contract.NewTanTin()
	if err != nil {
		fmt.Println(err)
		return
	}
	//tantin.Init(adminAddress, feeAddress, serverAddress)
	//tantin.LatestBlock()

	amount := big.NewInt(6000000)
	Usdt, err := contract.NewErc20(common.HexToAddress(contract.ChainConfig.UsdtAddress))
	if err != nil {
		fmt.Println(err)
		return
	}
	am := big.NewInt(1).Mul(amount, big.NewInt(1e18))
	Usdt.Approve(am, contract.ChainConfig.TantinContractAddress)
	//Usdc, err := contract.NewErc20(common.HexToAddress(contract.ChainConfig.UsdcAddress))
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//Usdc.Approve(amount, contract.ChainConfig.TantinContractAddress)
	resourceIdBytes := hexutils.HexToBytes(strings.TrimPrefix(contract.ResourceIdUsdt, "0x"))
	//resourceIdBytes := hexutils.HexToBytes(strings.TrimPrefix(contract.ResourceIdUsdc, "0x"))
	//resourceIdBytes := hexutils.HexToBytes(strings.TrimPrefix(contract.ResourceIdEth, "0x"))
	for {
		tantin.Deposit(
			common.HexToAddress("0x80B27CDE65Fafb1f048405923fD4a624fEa2d1C6"),
			[32]byte(resourceIdBytes),
			big.NewInt(3448148188),
			//big.NewInt(97),
			amount,
			big.NewInt(1e6),
			big.NewInt(time.Now().Unix()),
		)
		time.Sleep(time.Second * 30)
	}
}
