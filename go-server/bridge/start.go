package bridge

import (
	"coinstore/bridge/chains/ethereum"
	"coinstore/bridge/config"
	"coinstore/bridge/core"
	"coinstore/db"
	"coinstore/model"
	"fmt"
	log "github.com/ChainSafe/log15"
)

func Run() error {
	logger := log.Root()
	logger.Debug("Starting CoinStore Bridge...")
	db.InitMysql()

	sysErr := make(chan error)
	c := core.NewCore(sysErr)

	cfgs, err := model.GetAllConfig()
	if err != nil {
		panic(fmt.Sprintf("get config error: %v", err))
	}
	for _, cfg := range cfgs {
		chainConfig := config.NewConfig(cfg)
		logger.Debug("chain config: ", chainConfig)

		var newChain core.Chain
		chainLogger := log.Root().New("chain", chainConfig.ChainName)

		// TODO: 根据不同链类型，使用不同的初始化方法，下面先做一种
		//if chainConfig
		newChain, err = ethereum.InitializeChain(&chainConfig, chainLogger, sysErr)
		if err != nil {
			logger.Error("chain config: ", chainConfig)
			return err
		}
		c.AddChain(newChain)
	}

	logger.Debug("Config on initialization... config: %v \n", cfgs)

	c.Start()
	return nil
}
