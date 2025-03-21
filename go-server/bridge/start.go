package bridge

import (
	"coinstore/bridge/chains/ethereum"
	"coinstore/bridge/chains/tron"
	"coinstore/bridge/config"
	"coinstore/bridge/core"
	"coinstore/db"
	"coinstore/model"
	"errors"
	"fmt"
	log "github.com/calmw/clog"
)

func Run() error {
	logger := log.Root()
	logger.Debug("Starting CoinStore Bridge...")
	db.InitMysql(logger)

	//自动迁移为给定模型运行自动迁移，只会添加缺失的字段，不会删除/更改当前数据
	err := db.DB.AutoMigrate(&model.Config{}, &model.PollState{}, &model.BridgeTx{})
	if err != nil {
		logger.Debug("db AutoMigrate err: ", err)
	}

	sysErr := make(chan error)
	c := core.NewCore(sysErr)

	cfgs, err := model.GetAllConfig()
	if err != nil {
		panic(fmt.Sprintf("get config error: %v", err))
	}
	for _, cfg := range cfgs {
		chainConfig := config.NewConfig(cfg)
		logger.Debug("chain config: ", "config=", chainConfig)

		var newChain core.Chain
		chainLogger := log.Root().New("chain", chainConfig.ChainName)

		if chainConfig.ChainType == config.ChainTypeEvm {
			newChain, err = ethereum.InitializeChain(&chainConfig, chainLogger, sysErr)
			if err != nil {
				logger.Error("initialize chain", "error", err)
				return err
			}
		} else if chainConfig.ChainType == config.ChainTypeTron {
			newChain, err = tron.InitializeChain(&chainConfig, chainLogger, sysErr)
			if err != nil {
				logger.Error("initialize chain", "error", err)
				return err
			}
		} else {
			logger.Error("chain type", "error", err)
			return errors.New("chain type not supported")
		}

		c.AddChain(newChain)
	}

	logger.Debug("Config on initialization... ")

	c.Start()
	return nil
}
