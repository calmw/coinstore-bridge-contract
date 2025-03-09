package main

import (
	"coinstore/blockchain"
	"coinstore/services"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/status-im/keycard-go/hexutils"
	"math/big"
)

func main() {
	services.InitOpenBnbEnv()
	bridge, err := blockchain.NewBridge()
	if err != nil {
		fmt.Println(err)
		return
	}
	bridge.Init()
	someBytes := hexutils.HexToBytes("ac589789ed8c9d2c61f17b13369864b5f181e58eba230a6ee4ec4c3e7750cd1d")
	bridge.AdminSetResource(big.NewInt(1), [4]byte(someBytes))

	vote, err := blockchain.NewVote()
	if err != nil {
		fmt.Println(err)
		return
	}
	vote.Init()

	tantin, err := blockchain.NewTanTin()
	if err != nil {
		fmt.Println(err)
		return
	}
	//tantin.Init()
	//tantin.AdminSetToken()

	resourceIdBytes := hexutils.HexToBytes(blockchain.ResourceIdUsdt)
	//resourceIdBytes := hexutils.HexToBytes(blockchain.ResourceIdCoin)
	tantin.Deposit(common.HexToAddress("0x80B27CDE65Fafb1f048405923fD4a624fEa2d1C6"), [32]byte(resourceIdBytes), big.NewInt(2), big.NewInt(0), big.NewInt(1))

}
