package task

import (
	"coinstore/db"
	"coinstore/log"
	"coinstore/model"
	"context"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
	"time"
)

func UpdateRechargeRecord() {
	//
	var accounts []model.RelayerAccount
	err := db.DB.Model(&model.RelayerAccount{}).Find(accounts).Error
	if err != nil {
		log.Logger.Sugar().Errorf("get relayer account error %v", err)
	}
	for _, account := range accounts {

	}
}

func UpdateBalance() {
	var chains []model.ChainInfo
	err := db.DB.Model(&model.ChainInfo{}).Find(chains).Error
	if err != nil {
		log.Logger.Sugar().Errorf("get chain info error %v", err)
	}
	for _, chain := range chains {
		var relayerAccount model.RelayerAccount
		err = db.DB.Model(&model.RelayerAccount{}).Find(relayerAccount).Error
		if err != nil {
			log.Logger.Sugar().Errorf("get chain info error %v", err)
			continue
		}
		if chain.ChainType == 1 {
			balance, err := getEvmBalance(chain.Endpoint, relayerAccount.Account)
			if err == nil {
				return
			}
		} else if chain.ChainType == 2 {

		}
	}
}

func getEvmBalance(endpoint, account string) (*big.Int, error) {
	var err error
	var balance *big.Int
	var client *ethclient.Client
	for i := 0; i < 5; i++ {
		client, err = ethclient.Dial(endpoint)
		if err != nil || client == nil {
			log.Logger.Sugar().Errorf("dail %s failed", endpoint)
			time.Sleep(time.Second * 3)
			continue
		}
		balance, err = client.BalanceAt(context.Background(), common.HexToAddress(account), nil)
		if err != nil {
			log.Logger.Sugar().Errorf("get banlance error %v", err)
			time.Sleep(time.Second * 3)
			continue
		}
		break
	}

	return balance, err
}

func getTronBalance() {

}
