package tt

import (
	"coinstore/contract"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
	"os"
)

func InitTt(prvKey, adminAddress, feeAddress, serverAddress string) {
	contract.InitTantinEnv()
	err := os.Setenv("TB_KEY", prvKey)
	if err != nil {
		fmt.Println(err)
		return
	}
	bridge, err := contract.NewBridge()
	if err != nil {
		fmt.Println(err)
		return
	}
	bridge.Init(adminAddress)
	bridge.AdminSetResource(
		contract.ResourceIdEth,
		2,
		common.HexToAddress(contract.ChainConfig.WEthAddress),
		big.NewInt(100),
		false,
		false,
		false,
	)
	bridge.AdminSetResource(
		contract.ResourceIdWeth,
		2,
		common.HexToAddress(contract.ChainConfig.WEthAddress),
		big.NewInt(100),
		false,
		false,
		false,
	)

	vote, err := contract.NewVote()
	if err != nil {
		fmt.Println(err)
		return
	}
	vote.Init()

	tantin, err := contract.NewTanTin()
	if err != nil {
		fmt.Println(err)
		return
	}
	tantin.Init(adminAddress, feeAddress, serverAddress)
}
