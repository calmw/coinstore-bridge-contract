package tron

import (
	"coinstore/contract"
	"fmt"
	"os"
)

func InitTron(prvKey, adminAddress, feeAddress, serverAddress, realyerOneAddress, realyerTwoAddress, realyerThreeAddress string, fee uint64, tronKeyStorePassphrase string) {
	err := os.Setenv("TB_KEY", prvKey)
	if err != nil {
		fmt.Println(err)
		return
	}
	//contract.InitTronEnv()
	contract.InitTronEnvProd()
	ks, ka, err := InitKeyStore(prvKey, tronKeyStorePassphrase)
	if err != nil {
		fmt.Println(err)
		return
	}
	bridge, err := contract.NewBridgeTron(ka, ks, tronKeyStorePassphrase)
	if err != nil {
		fmt.Println(err)
		return
	}
	bridge.Init(adminAddress, fee)

	vote, err := contract.NewVoteTron(ka, ks, tronKeyStorePassphrase)
	if err != nil {
		fmt.Println(err)
		return
	}
	vote.Init(realyerOneAddress, realyerTwoAddress, realyerThreeAddress)

	tantin, err := contract.NewTanTinTron(ka, ks, tronKeyStorePassphrase)
	if err != nil {
		fmt.Println(err)
		return
	}
	tantin.Init(adminAddress, feeAddress, serverAddress)
}
