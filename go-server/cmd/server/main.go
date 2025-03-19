package main

import (
	"coinstore/cmd/server/service"
	"coinstore/db"
	log "github.com/calmw/blog"
	"github.com/gin-gonic/gin"
)

func main() {
	logger := log.Root()
	logger.Debug("Starting CoinStore Bridge Server...")
	db.InitMysql(logger)
	router := gin.Default()
	router.GET("/history", service.BridgeTx)
	_ = router.Run()
}
