package model

import (
	"coinstore/db"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"io"
	"os"
)

type TokenInfo struct {
	Id           uint64 `gorm:"primaryKey" json:"id"`
	TokenName    string `gorm:"column:token_name;type:varchar(100);comment:' token名称'" json:"token_name"`
	TokenAddress string `gorm:"column:token_address;type:varchar(100);comment:' token地址，coin为0地址'" json:"token_address"`
	Icon         []byte `gorm:"column:icon;comment:'icon'" json:"icon"`
	ChainId      string `gorm:"column:chain_id;default:0;comment:'链ID'" json:"chain_id"`
}

func AddToken(chainId, tokenAddress, iconFile string) error {
	file, err := os.Open(iconFile)
	if err != nil {
		return err
	}
	defer file.Close()
	imageData, err := io.ReadAll(file)
	if err != nil {
		return err
	}
	var token TokenInfo
	err = db.DB.Model(&TokenInfo{}).Where("chain_id=? and token_address=?", chainId, tokenAddress).First(&token).Error
	fmt.Println(err)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return db.DB.Model(&TokenInfo{}).Where("chain_id=? and token_address=?", chainId, tokenAddress).Create(&TokenInfo{
			TokenAddress: tokenAddress,
			Icon:         imageData,
			ChainId:      chainId,
		}).Error
	} else if err != nil {
		return err
	} else {
		return db.DB.Model(&TokenInfo{}).Where("chain_id=? and token_address=?", chainId, tokenAddress).Update(
			"icon", imageData,
		).Error
	}
}
