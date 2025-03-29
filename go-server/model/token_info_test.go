package model

import (
	"coinstore/db"
	"fmt"
	log "github.com/calmw/clog"
	"testing"
)

func TestAddToken(t *testing.T) {
	logger := log.Root()
	logger.Debug("Starting CoinStore Bridge...")
	db.InitMysql(logger)
	err := AddToken("3", "", "../cmd/server/images/BSC.png")
	fmt.Println(err)
}
