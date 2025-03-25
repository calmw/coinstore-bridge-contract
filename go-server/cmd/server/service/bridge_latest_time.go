package service

import (
	"github.com/gin-gonic/gin"
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
	if len(chainIdFrom) == 0 || len(chainIdTo) == 0 {
		c.JSON(200, gin.H{
			"code": 1,
			"msg":  "Address and address type are required",
		})
		return
	}

	c.JSON(200, gin.H{
		"code": 0,
		"data": 180,
		"msg":  "OK",
	})
}
