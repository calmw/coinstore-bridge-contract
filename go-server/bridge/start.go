package bridge

import (
	"coinstore/bridge/chains/ethereum"
	"coinstore/bridge/config"
	"coinstore/bridge/core"
	"coinstore/db"
	"coinstore/model"
	"fmt"
	log "github.com/calmw/blog"
)

func Run() error {
	logger := log.Root()
	logger.Debug("Starting CoinStore Bridge...")
	db.InitMysql(logger)

	//自动迁移为给定模型运行自动迁移，只会添加缺失的字段，不会删除/更改当前数据
	err := db.DB.AutoMigrate(&model.Config{}, &model.PollState{}, &model.BridgeOrder{})
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
