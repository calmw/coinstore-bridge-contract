package model

import (
	"coinstore/db"
	"fmt"
	log "github.com/calmw/clog"
	"testing"
)

func TestAddToken(t *testing.T) {
	logger := log.Root()
	logger.Debug("Starting CoinStore Bridge...")
	db.InitMysql(logger)
	//_ = AddToken(2, "", "../cmd/server/images/bsc.png")
	//_ = AddToken(3, "", "../cmd/server/images/tron.png")
	//_ = AddToken(4, "", "../cmd/server/images/eth.png")
	//_ = AddToken(1, "", "../cmd/server/images/tt.png")
	err := AddToken(1, "TETH", "0x9259F02Bf9C7457caf8944594e046CA77A18D86a", "../cmd/server/images/token/teth.png")

	fmt.Println(1, err)
	err = AddToken(1, "TUSDT", "0x50e39AddfC907710E53d9C242D9cbA2C1F2edCf2", "../cmd/server/images/token/tusdt.png")

	fmt.Println(2, err)
	err = AddToken(1, "TUSDC", "0x74E3DFDe1403b3F55924C1c92e39bEAA2A489965", "../cmd/server/images/token/tusdc.png")

	fmt.Println(3, err)
	err = AddToken(2, "WETH", "0x4259600CEEdC1fB51805cEAC1dD3A21c24f58388", "../cmd/server/images/token/weth.png")

	fmt.Println(4, err)
	err = AddToken(2, "USDT", "0xfBe1e02C25a04f6CD6b044F847697b48B3E99a16", "../cmd/server/images/token/tusdt.png")

	fmt.Println(5, err)
	err = AddToken(2, "USDC", "0x0Ac4dcf02969485132B765fa5D469111D6497c53", "../cmd/server/images/token/tusdc.png")

	fmt.Println(6, err)
	err = AddToken(3, "ETH", "TKUirRN2FEzQVYhdsPbTe7LqR7db7F6wpp", "../cmd/server/images/token/eth.png")

	fmt.Println(7, err)
	err = AddToken(3, "USDT", "TXYZopYRdj2D9XRtbG411XZZ3kM5VkAeBf", "../cmd/server/images/token/tusdt.png")

	fmt.Println(8, err)
	_ = AddToken(3, "USDC", "TE1zsjWWTFhsSzMaJKV45tgg1JqutBHJuZ", "../cmd/server/images/token/tusdc.png")

	fmt.Println(9, err)
	err = AddToken(4, "ETH", "0x0000000000000000000000000000000000000000", "../cmd/server/images/token/eth.png")

	fmt.Println(10, err)
	err = AddToken(4, "USDT", "0xd78CCb0C79489d8C50AfFE4E881B2C2E93706f8b", "../cmd/server/images/token/tusdt.png")

	fmt.Println(11, err)
	err = AddToken(4, "USDC", "0x424cF9Dc24c7c8F006421937d3E288Be84D2daa4", "../cmd/server/images/token/tusdc.png")
	fmt.Println(12, err)
	err = AddToken(41, "USDC", "0x424cF9Dc24c7c8F006421937d3E288Be84D2daa4", "../cmd/server/images/tt.png")
	fmt.Println(12, err)
	err = AddToken(42, "USDC", "0x424cF9Dc24c7c8F006421937d3E288Be84D2daa4", "../cmd/server/images/bsc.png")
	fmt.Println(12, err)
	err = AddToken(43, "USDC", "0x424cF9Dc24c7c8F006421937d3E288Be84D2daa4", "../cmd/server/images/tron.png")
	fmt.Println(12, err)
	err = AddToken(44, "USDC", "0x424cF9Dc24c7c8F006421937d3E288Be84D2daa4", "../cmd/server/images/eth.png")
	fmt.Println(12, err)
}
