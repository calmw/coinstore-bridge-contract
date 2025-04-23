package bridge

import (
	"coinstore/bridge/chains/ethereum"
	"coinstore/bridge/chains/tron"
	"coinstore/bridge/config"
	"coinstore/bridge/core"
	"coinstore/bridge/monitor"
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

	//自动迁移
	err := db.DB.AutoMigrate(&model.ChainInfo{}, &model.PollState{}, &model.BridgeTx{})
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
			newChain, err = ethereum.InitializeChain(chainConfig, chainLogger, sysErr)
			if err != nil {
				logger.Error("initialize chain", "error", err)
				return err
			}
		} else if chainConfig.ChainType == config.ChainTypeTron {
			config.TronCfg = chainConfig
			newChain, err = tron.InitializeChain(&chainConfig, chainLogger, sysErr)
			if err != nil {
				logger.Error("initialize chain", "error", err)
				return err
			}
		} else {
			logger.Error("chain type", "error", err)
			return errors.New("chain type not supported")
		}
		core.ChainType[cfg.ChainId] = config.ChainType(cfg.ChainType)
		c.AddChain(newChain)
	}

	logger.Debug("ChainInfo on initialization... ")

	go monitor.Start()
	c.Start()
	return nil
}
