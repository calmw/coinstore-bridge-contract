package main

import (
	"coinstore/cmd/server/service"
	"coinstore/db"
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

	router := gin.Default()
	// 创建限速器,每秒5次
	limiter := tollbooth.NewLimiter(5, nil)
	// 使用限速中间件
	router.GET("/history", tollbooth_gin.LimitHandler(limiter), service.BridgeTx)
	router.GET("/check_address", tollbooth_gin.LimitHandler(limiter), service.CheckAddress)
	router.GET("/bridge_latest_time", tollbooth_gin.LimitHandler(limiter), service.BridgeLatestTime)
	router.GET("/convert_address", tollbooth_gin.LimitHandler(limiter), service.ConvertAddress)
	_ = router.Run(os.Getenv("LISTEN_ADDR"))
	//_ = router.Run()
}
