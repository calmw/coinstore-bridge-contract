package main

import (
	"coinstore/pkg/blockchain"
	"coinstore/services"
	"log"
	"math/big"
)

func main() {
	services.InitOpenBnbEnv()

	//fee, err := blockchain.NewFee()
	//if err != nil {
	//	log.Println(err)
	//	return
	//}
	//fee.FeeInit()
	//
	//token, err := blockchain.NewToken()
	//if err != nil {
	//	log.Println(err)
	//	return
	//}
	//token.TokenInit()

	//gameCategory, err := blockchain.NewGameCategory()
	//if err != nil {
	//	log.Println(err)
	//	return
	//}
	//gameCategory.GameCategoryInit()
	//
	//order, err := blockchain.NewOrder()
	//if err != nil {
	//	log.Println(err)
	//	return
	//}
	//order.OrderInit()
	////order.GetLog(big.NewInt(103315800), big.NewInt(103316000))
	//
	//autoBet, err := blockchain.NewAutoBet()
	//if err != nil {
	//	log.Println(err)
	//	return
	//}
	//autoBet.AutoBetInit()
	//
	//game, err := blockchain.NewGame()
	//if err != nil {
	//	log.Println(err)
	//	return
	//}
	//game.AutoBetTest()
	//game.GameInit()
	//game.BetTest()

	jackpot, err := blockchain.NewJackpot()
	if err != nil {
		log.Println(err)
		return
	}
	jackpot.Payout("0x37225753153de1241Bc8846A1c816453B0Bfa3f1", big.NewInt(2))
	//jackpot.JackpotInit()

	//erc20, err := blockchain.NewErc20(blockchain.ChainConfig.USDTAddress)
	//if err != nil {
	//	log.Println(err)
	//	return
	//}
	//erc20.Approve(blockchain.ChainConfig.GameContractAddress, 1000000000000000000, 6)
	//erc20.Transfer(blockchain.ChainConfig.GameContractAddress, 1000000, 6)
	//erc20, err = blockchain.NewErc20(blockchain.ChainConfig.USDCAddress)
	//if err != nil {
	//	log.Println(err)
	//	return
	//}
	//erc20.Approve(blockchain.ChainConfig.GameContractAddress, 1000000000000000000, 6)
	//erc20.Transfer(blockchain.ChainConfig.GameContractAddress, 1000000, 6)
	//
	//erc20, err = blockchain.NewErc20(blockchain.ChainConfig.USDTAddress)
	//if err != nil {
	//	log.Println(err)
	//	return
	//}
	//erc20.Approve(blockchain.ChainConfig.JackPotContractAddress, 1000000000000000000, 6)
	//erc20.Transfer(blockchain.ChainConfig.JackPotContractAddress, 1000000, 6)
	//erc20, err = blockchain.NewErc20(blockchain.ChainConfig.USDCAddress)
	//if err != nil {
	//	log.Println(err)
	//	return
	//}
	//erc20.Approve(blockchain.ChainConfig.JackPotContractAddress, 1000000000000000000, 6)
	//erc20.Transfer(blockchain.ChainConfig.JackPotContractAddress, 1000000, 6)
}
