package model

type RelayerAccount struct {
	Id      uint64  `gorm:"primaryKey" json:"id"`
	ChainId ChainId `gorm:"column:source_chain_id;comment:'链ID'" json:"chain_id"`
	Account string  `gorm:"column:resource_id;type:varchar(100);comment:'relayer账户'" json:"account"`
}

// 统计每日跨链桥Gas自消耗，就用浏览器ABI获取充值，然后每天获取一次余额，来计算
