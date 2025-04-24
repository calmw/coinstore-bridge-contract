package model

import (
	"github.com/shopspring/decimal"
)

type RelayerRecharge struct {
	Id      uint64          `gorm:"primaryKey" json:"id"`
	ChainId ChainId         `gorm:"column:source_chain_id;comment:'链ID'" json:"chain_id"`
	Account string          `gorm:"column:resource_id;type:varchar(100);comment:'relayer账户'" json:"account"`
	Amount  decimal.Decimal `gorm:"column:amount;type:decimal(20,0);comment:'充币数额'" json:"amount"`
}

// 统计每日跨链桥Gas自消耗，就用浏览器ABI获取充值，然后每天获取一次余额，来计算
