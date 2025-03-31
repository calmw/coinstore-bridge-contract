package main

import (
	"coinstore/cmd/server/service"
	"coinstore/db"
	"coinstore/model"
	log "github.com/calmw/clog"
	"github.com/didip/tollbooth/v7"
	"github.com/didip/tollbooth_gin"
	"github.com/gin-gonic/gin"
	"os"
)

func main() {
	logger := log.Root()
	logger.Debug("Starting CoinStore Bridge Server...")
	db.InitMysql(logger)

	//自动迁移
	err := db.DB.AutoMigrate(&model.ChainInfo{}, &model.BridgeTx{}, &model.ResourceIdInfo{}, &model.TokenInfo{})
	if err != nil {
		logger.Debug("db AutoMigrate err: ", err)
	}

	router := gin.Default()
	// 创建限速器,每秒5次
	limiter := tollbooth.NewLimiter(5, nil)
	// 使用限速中间件
	router.GET("/history", tollbooth_gin.LimitHandler(limiter), service.BridgeTx)
	router.GET("/check_address", tollbooth_gin.LimitHandler(limiter), service.CheckAddress)
	router.GET("/bridge_latest_time", tollbooth_gin.LimitHandler(limiter), service.BridgeLatestTime)
	router.GET("/convert_address", tollbooth_gin.LimitHandler(limiter), service.ConvertAddress)
	router.GET("/get_resource_id", tollbooth_gin.LimitHandler(limiter), service.GetResourceId)
	router.GET("/get_chain_list", tollbooth_gin.LimitHandler(limiter), service.GetChainList)
	router.GET("/get_token_list", tollbooth_gin.LimitHandler(limiter), service.GetTokenList)
	addr := os.Getenv("LISTEN_ADDR")
	if addr == "" {
		addr = "0.0.0.0:8080"
	}
	_ = router.Run(addr)
}
