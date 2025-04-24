package task

import (
	"coinstore/db"
	"coinstore/model"
)

func RelayerRecharge() {
	//
	var accounts []model.RelayerAccount
	err := db.DB.Model(&model.RelayerAccount{}).Find(accounts).Error
	if err != nil {

	}
}
