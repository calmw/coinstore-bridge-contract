package model

import (
	"github.com/shopspring/decimal"
)

type DailyReport struct {
	Id           uint64          `gorm:"primaryKey" json:"id"`
	ChainId      ChainId         `gorm:"column:source_chain_id;comment:'链ID'" json:"chain_id"`
	TokenName    string          `gorm:"column:resource_id;type:varchar(100);comment:'币种名称'" json:"token_name"`
	TokenAddress string          `gorm:"column:resource_id;type:varchar(100);comment:'币种地址'" json:"token_address"`
	Direct       int             `gorm:"column:vote_status;default:0;comment:'1入 2出'" json:"direct"`
	Amount       decimal.Decimal `gorm:"column:amount;type:decimal(20,0);comment:'今日累计跨链数额'" json:"amount"`
	Fee          decimal.Decimal `gorm:"column:amount;type:decimal(20,0);comment:'今日累计跨链费用'" json:"fee"`
	BridgeGasFee decimal.Decimal `gorm:"column:amount;type:decimal(20,0);comment:'跨链桥今日累计Gas费用(当前链Coin数量)'" json:"bridge_gas_fee"`
}

// 统计每日跨链桥Gas自消耗，就用浏览器ABI获取充值，然后每天获取一次余额，来计算
