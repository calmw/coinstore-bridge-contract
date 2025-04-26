package tt

import (
	"coinstore/contract"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
	"os"
)

func InitTt(prvKey, adminAddress, feeAddress, serverAddress, realyerOneAddress, realyerTwoAddress, realyerThreeAddress string, fee uint64) {
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
	bridge.Init(adminAddress, fee)
	bridge.AdminSetResource(
		contract.ResourceIdEth,
		2,
		common.HexToAddress(contract.ChainConfig.WEthAddress),
		big.NewInt(int64(fee)),
		false,
		false,
		false,
	)
	bridge.AdminSetResource(
		contract.ResourceIdWeth,
		2,
		common.HexToAddress(contract.ChainConfig.WEthAddress),
		big.NewInt(int64(fee)),
		false,
		false,
		false,
	)

	vote, err := contract.NewVote()
	if err != nil {
		fmt.Println(err)
		return
	}
	vote.Init(realyerOneAddress, realyerTwoAddress, realyerThreeAddress)

	tantin, err := contract.NewTanTin()
	if err != nil {
		fmt.Println(err)
		return
	}
	tantin.Init(adminAddress, feeAddress, serverAddress)
}
