package model

type ResourceIdInfo struct {
	Id               uint64 `gorm:"primaryKey" json:"id"`
	ResourceId       string `gorm:"column:resource_id;type:varchar(100);comment:'resource ID'" json:"resource_id"`
	ChainIdFrom      string `gorm:"column:chain_id_from;default:0;comment:'源链ID'" json:"chain_id_from"`
	ChainIdTo        string `gorm:"column:chain_id_to;default:0;comment:'目标链ID'" json:"chain_id_to"`
	TokenAddressFrom string `gorm:"column:token_address_from;default:0;comment:'源链token地址'" json:"token_address_from"`
	TokenAddressTo   string `gorm:"column:token_address_to;default:0;comment:'目标链token地址'" json:"token_address_to"`
}
