package model

type TokenInfo struct {
	Id           uint64 `gorm:"primaryKey" json:"id"`
	TokenAddress string `gorm:"column:token_address;type:varchar(100);comment:' token地址，coin为0地址'" json:"token_address"`
	Icon         string `gorm:"column:icon;type:varchar(100);comment:' token地址，coin为0地址'" json:"icon"`
	ChainId      string `gorm:"column:chain_id;default:0;comment:'链ID'" json:"chain_id"`
}
