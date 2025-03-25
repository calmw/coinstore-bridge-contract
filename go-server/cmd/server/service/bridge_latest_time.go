package service

import (
	"coinstore/db"
	"coinstore/model"
	"github.com/gin-gonic/gin"
	"time"
)

func BridgeLatestTime(c *gin.Context) {
	var q Query
	if c.ShouldBindQuery(&q) != nil {
		c.JSON(200, gin.H{
			"code": 1,
			"msg":  "参数错误",
		})
		return
	}
	chainIdFrom := c.Query("chain_id_from")
	chainIdTo := c.Query("chain_id_to")
	var bridgeTx model.BridgeTx
	if err := db.DB.Model(model.BridgeTx{}).Where("source_chain_id=? and destination_chain_id=? and bridge_status=2", chainIdFrom, chainIdTo).Order("id desc").First(&bridgeTx).Error; err != nil {
		c.JSON(200, gin.H{
			"code": 1,
			"msg":  "no record",
		})
		return
	}
	depositAt, _ := time.ParseInLocation("2006-01-02 15:04:05", bridgeTx.DepositAt, time.Local)
	receiveAt, _ := time.ParseInLocation("2006-01-02 15:04:05", bridgeTx.ReceiveAt, time.Local)

	c.JSON(200, gin.H{
		"code": 0,
		"data": receiveAt.Unix() - depositAt.Unix(),
		"msg":  "OK",
	})
}
