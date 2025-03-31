package model

import (
	"coinstore/db"
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
	_ = AddToken(1, "0x9259F02Bf9C7457caf8944594e046CA77A18D86a", "../cmd/server/images/token/teth.png")
	_ = AddToken(1, "0x50e39AddfC907710E53d9C242D9cbA2C1F2edCf2", "../cmd/server/images/token/tusdt.png")
	_ = AddToken(1, "0x74E3DFDe1403b3F55924C1c92e39bEAA2A489965", "../cmd/server/images/token/tusdc.png")
	_ = AddToken(2, "0x4259600CEEdC1fB51805cEAC1dD3A21c24f58388", "../cmd/server/images/token/weth.png")
	_ = AddToken(2, "0xfBe1e02C25a04f6CD6b044F847697b48B3E99a16", "../cmd/server/images/token/usdt.png")
	_ = AddToken(2, "0x0Ac4dcf02969485132B765fa5D469111D6497c53", "../cmd/server/images/token/usdc.png")
	_ = AddToken(3, "", "../cmd/server/images/token/eth.png")
	_ = AddToken(3, "TXYZopYRdj2D9XRtbG411XZZ3kM5VkAeBf", "../cmd/server/images/token/usdt.png")
	_ = AddToken(3, "", "../cmd/server/images/token/usdc.png")
	_ = AddToken(4, "", "../cmd/server/images/token/eth.png")
	_ = AddToken(4, "", "../cmd/server/images/token/usdt.png")
	_ = AddToken(4, "", "../cmd/server/images/token/usdc.png")
}
