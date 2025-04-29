package bridge

import (
	"coinstore/bridge/chains/ethereum"
	"coinstore/bridge/chains/tron"
	"coinstore/bridge/config"
	"coinstore/bridge/core"
	"coinstore/db"
	"coinstore/model"
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

		if chainConfig.ChainId == 3448148188 || chainConfig.ChainId == 728126428 {
			fmt.Println("!!!!!!!!!!!! tron ", chainConfig.ChainId)
			config.TronCfg = chainConfig
			newChain, err = tron.InitializeChain(&chainConfig, chainLogger, sysErr)
			if err != nil {
				logger.Error("initialize chain", "error", err)
				return err
			}
		} else {
			fmt.Println("!!!!!!!!!!!! evm ", chainConfig.ChainId)
			newChain, err = ethereum.InitializeChain(chainConfig, chainLogger, sysErr)
			if err != nil {
				logger.Error("initialize chain", "error", err)
				return err
			}
		}
		core.ChainType[cfg.ChainId] = config.ChainType(cfg.ChainType)
		c.AddChain(newChain)
	}

	logger.Debug("ChainInfo on initialization... ")

	//go monitor.Start()
	c.Start()
	return nil
}
