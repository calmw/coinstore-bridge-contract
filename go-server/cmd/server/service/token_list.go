package service

import (
	"coinstore/db"
	"coinstore/model"
	"github.com/gin-gonic/gin"
)

func GetTokenList(c *gin.Context) {
	var q Query
	if c.ShouldBindQuery(&q) != nil {
		c.JSON(200, gin.H{
			"code": 1,
			"msg":  "参数错误",
		})
		return
	}
	chainId := c.Query("chain_id")
	if len(chainId) == 0 {
		c.JSON(200, gin.H{
			"code": 1,
			"msg":  "chain_id is required",
		})
		return
	}
	var tokenInfo []model.TokenInfo
	db.DB.Model(&model.TokenInfo{}).Where("chain_id=?", chainId).Find(&tokenInfo)
	c.JSON(200, gin.H{
		"code": 0,
		"msg":  "OK",
		"data": tokenInfo,
	})
}
