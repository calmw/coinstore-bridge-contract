package service

import (
	"coinstore/db"
	"coinstore/model"
	"coinstore/utils"
	"github.com/gin-gonic/gin"
	"strings"
	"time"
)

func BridgeTime(c *gin.Context) {
	var q Query
	if c.ShouldBindQuery(&q) != nil {
		c.JSON(200, gin.H{
			"code": 1,
			"msg":  "参数错误",
		})
		return
	}
	sourceChainId := c.Query("source_chain_id")
	destinationChainId := c.Query("destination_chain_id")
	depositHash := c.Query("deposit_hash")
	if len(sourceChainId) == 0 || len(destinationChainId) == 0 || len(depositHash) == 0 {
		c.JSON(200, gin.H{
			"code": 1,
			"msg":  "source_chain_id 、deposit_hash and destination_chain_id are required",
		})
		return
	}

	var seconds int64 = 300

	// 更新记录
	var record model.BridgeTx
	err := db.DB.Model(&model.BridgeTx{}).Where("source_chain_id=? and destination_chain_id=? and deposit_hash=?", sourceChainId, destinationChainId, strings.ToLower(depositHash)).Order("id desc").First(&record).Error
	if err == nil {
		if len(record.ExecuteHash) > 0 {
			seconds = 0
		} else {
			depositAt, _ := utils.DatetimeToUnix(record.DepositAt)
			diff := time.Now().Unix() - depositAt
			if diff <= seconds {
				seconds -= diff
			}
		}
	}

	c.JSON(200, gin.H{
		"code": 0,
		"msg":  "OK",
		"data": seconds,
	})
}
