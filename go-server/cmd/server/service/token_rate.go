package service

import (
	"coinstore/cmd/server/task"
	"coinstore/db"
	"coinstore/model"
	"github.com/gin-gonic/gin"
	"strings"
	"time"
)

func GetPrice(c *gin.Context) {
	var q Query
	if c.ShouldBindQuery(&q) != nil {
		c.JSON(200, gin.H{
			"code": 1,
			"msg":  "参数错误",
		})
		return
	}
	tokenName := c.Query("token_name")
	tokenAddress := c.Query("token_address")
	chainId := c.Query("chain_id")
	var tokenInfo model.TokenInfo
	err := db.DB.Model(&model.TokenInfo{}).Where("token_name=? and token_address=? ",
		tokenName,
		strings.ToLower(tokenAddress),
	).First(&tokenInfo).Error
	if err != nil {
		c.JSON(200, gin.H{
			"code": 1,
			"msg":  "token not found",
			"data": nil,
		})
		return
	}

	price := task.GetPrice(strings.ToUpper(tokenName))
	timestamp := time.Now().Unix()

	c.JSON(200, gin.H{
		"code": 0,
		"msg":  "OK",
		"data": resourceInfo.ResourceId,
	})
}
