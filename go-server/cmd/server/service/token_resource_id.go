package service

import (
	"coinstore/db"
	"coinstore/model"
	"github.com/gin-gonic/gin"
)

func GetResourceId(c *gin.Context) {
	var q Query
	if c.ShouldBindQuery(&q) != nil {
		c.JSON(200, gin.H{
			"code": 1,
			"msg":  "参数错误",
		})
		return
	}
	tokenAddressFrom := c.Query("token_address_from")
	tokenAddressTo := c.Query("token_address_to")
	chainIdFrom := c.Query("chain_id_from")
	chainIdTo := c.Query("chain_id_to")
	var resourceIdInfo model.ResourceIdInfo
	err := db.DB.Model(&model.ResourceIdInfo{}).Where("token_address_from=? and token_address_to=? and chain_id_from=? and chain_id_to=?",
		tokenAddressFrom,
		tokenAddressTo,
		chainIdFrom,
		chainIdTo,
	).First(&resourceIdInfo).Error
	if err != nil {
		c.JSON(200, gin.H{
			"code": 1,
			"msg":  "record not found",
			"data": nil,
		})
		return
	}

	c.JSON(200, gin.H{
		"code": 0,
		"msg":  "OK",
		"data": resourceIdInfo.ResourceId,
	})
}
