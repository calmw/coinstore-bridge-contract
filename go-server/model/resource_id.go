package model

type ResourceInfo struct {
	Id                      uint64  `gorm:"primaryKey" json:"id"`
	ResourceId              string  `gorm:"column:resource_id;type:varchar(100);comment:'resource ID'" json:"resource_id"`
	SourceChainId           ChainId `gorm:"column:source_chain_id;comment:'源链ID'" json:"source_chain_id"`
	SourceTokenAddress      string  `gorm:"column:source_token_address;comment:'源链token地址'" json:"source_token_address"`
	DestinationChainId      ChainId `gorm:"column:destination_chain_id;comment:'目标链ID'" json:"destination_chain_id"`
	DestinationTokenAddress string  `gorm:"column:destination_token_address;comment:'目标链token地址'" json:"destination_token_address"`
}
