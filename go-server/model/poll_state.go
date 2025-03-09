package model

import (
	"errors"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
	"math/big"
)

type PollState struct {
	gorm.Model
	ChainId     int             `gorm:"chain_id;comment:'自定义链ID'" json:"chain_id"`
	BlockHeight decimal.Decimal `gorm:"block_height;type:bigint(30);comment:'扫块高度'" json:"block_height"`
}

func SetBlockHeight(tx *gorm.DB, chainId int, blockHeight decimal.Decimal) error {
	var ps PollState
	var err error
	err = tx.Model(&PollState{}).Where("chain_id=?", chainId).First(&ps).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return tx.Model(&PollState{}).Create(&PollState{
			ChainId:     chainId,
			BlockHeight: blockHeight,
		}).Error
	} else if err != nil {
		return err
	} else {
		return tx.Model(&PollState{}).Where("chain_id=?", chainId).Updates(map[string]interface{}{
			"block_height": blockHeight,
		}).Error
	}
}

func GetBlockHeight(tx *gorm.DB, chainId int) (*big.Int, error) {
	var ps PollState
	var err error
	err = tx.Model(&PollState{}).Where("chain_id=?", chainId).First(&ps).Error
	if err != nil {
		return nil, err
	}
	return ps.BlockHeight.BigInt(), nil
}
