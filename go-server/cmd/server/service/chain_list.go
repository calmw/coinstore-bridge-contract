package service

import (
	"coinstore/db"
	"coinstore/model"
	"github.com/gin-gonic/gin"
)

func GetChainList(c *gin.Context) {
	var q Query
	if c.ShouldBindQuery(&q) != nil {
		c.JSON(200, gin.H{
			"code": 1,
			"msg":  "参数错误",
		})
		return
	}
	tx := db.DB.Model(&model.ChainInfo{})
	chainId := c.Query("chain_id")
	if len(chainId) > 0 {
		tx = tx.Where("chain_id=?", chainId)
	}
	var chainInfo []model.ChainInfo
	tx.Find(&chainInfo)
	c.JSON(200, gin.H{
		"code": 0,
		"msg":  "OK",
		"data": chainInfo,
	})
}
