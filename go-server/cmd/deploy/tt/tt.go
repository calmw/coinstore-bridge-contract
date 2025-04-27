package tt

import (
	"coinstore/contract"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
	"os"
)

// 0xa47142f08f859aCeb2127C6Ab66eC8c8bc4FFBA9
// ./tb tt --admin_address '0xa47142f08f859aCeb2127C6Ab66eC8c8bc4FFBA9' --fee_address '0xa47142f08f859aCeb2127C6Ab66eC8c8bc4FFBA9' --server_address '0xa47142f08f859aCeb2127C6Ab66eC8c8bc4FFBA9' --key 'ee15b6b9e8f7483206fd56639a5443ca13c9e29d1ac84ada97464877687f98dc' --relayer_one_address  '0x1933ccd14cafe561d862e5f35d0de75322a55412'   --relayer_two_address  '0x0f2a804a069f9a240c617f5192de84ee1876c285' --relayer_three_address '0x2eb24f919c0428af3e564d94c4340044d8fd5376' --fee 4

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
		big.NewInt(1e18),
		big.NewInt(int64(fee)),
		false,
		false,
		false,
	)
	bridge.AdminSetResource(
		contract.ResourceIdWeth,
		2,
		common.HexToAddress(contract.ChainConfig.WEthAddress),
		big.NewInt(1e18),
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
