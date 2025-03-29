package model

import (
	"coinstore/db"
	"encoding/base64"
	"errors"
	"gorm.io/gorm"
	"os"
)

type TokenInfo struct {
	Id           uint64 `gorm:"primaryKey" json:"id"`
	TokenName    string `gorm:"column:token_name;type:varchar(100);comment:' token名称'" json:"token_name"`
	TokenAddress string `gorm:"column:token_address;type:varchar(100);comment:' token地址，coin为0地址'" json:"token_address"`
	Icon         string `gorm:"column:icon;type:longtext;comment:'icon'" json:"icon"`
	ChainId      string `gorm:"column:chain_id;default:0;comment:'链ID'" json:"chain_id"`
}

func AddToken(chainId, tokenAddress, iconFile string) error {
	srcByte, err := os.ReadFile(iconFile)
	if err != nil {
		return err
	}
	base64Str := "data:image/png;base64," + base64.StdEncoding.EncodeToString(srcByte)

	var token TokenInfo
	err = db.DB.Model(&TokenInfo{}).Where("chain_id=? and token_address=?", chainId, tokenAddress).First(&token).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return db.DB.Model(&TokenInfo{}).Where("chain_id=? and token_address=?", chainId, tokenAddress).Create(&TokenInfo{
			TokenAddress: tokenAddress,
			Icon:         base64Str,
			ChainId:      chainId,
		}).Error
	} else if err != nil {
		return err
	} else {
		return db.DB.Model(&TokenInfo{}).Where("chain_id=? and token_address=?", chainId, tokenAddress).Update(
			"icon", base64Str,
		).Error
	}
}
