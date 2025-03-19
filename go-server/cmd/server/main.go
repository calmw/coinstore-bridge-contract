package main

import (
	"coinstore/cmd/server/service"
	"coinstore/db"
	log "github.com/calmw/blog"
	"github.com/didip/tollbooth/v7"
	"github.com/didip/tollbooth_gin"
	"github.com/gin-gonic/gin"
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
	_ = router.Run()
}
