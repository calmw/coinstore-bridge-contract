package service

import (
	"github.com/gin-gonic/gin"
	"regexp"
	"strings"
)

func ConvertAddress(c *gin.Context) {
	var q Query
	if c.ShouldBindQuery(&q) != nil {
		c.JSON(200, gin.H{
			"code": 1,
			"msg":  "参数错误",
		})
		return
	}
	address := c.Query("address")
	chainIdFrom := c.Query("chain_id_from")
	chainIdTo := c.Query("chain_id_to")
	if len(chainIdFrom) == 0 || len(chainIdTo) == 0 {
		c.JSON(200, gin.H{
			"code": 1,
			"msg":  "chain_id_from and chain_id_to type are required",
		})
		return
	}
	if len(address) == 0 {
		c.JSON(200, gin.H{
			"code": 1,
			"msg":  "address is required",
		})
		return
	}
	if (chainIdFrom == "1" || chainIdFrom == "1" || chainIdFrom == "1") && (chainIdFrom == "1") {
		if isValidEthAddress(address) {
			c.JSON(200, gin.H{
				"code": 0,
				"msg":  "OK",
			})
			return
		} else {
			c.JSON(200, gin.H{
				"code": 1,
				"msg":  "Invalid address",
			})
			return
		}
	} else if addressType == "2" {
		if isValidTronAddress(address) {
			c.JSON(200, gin.H{
				"code": 0,
				"msg":  "OK",
			})
			return
		} else {
			c.JSON(200, gin.H{
				"code": 1,
				"msg":  "Invalid address",
			})
			return
		}
	} else {
		c.JSON(200, gin.H{
			"code": 1,
			"msg":  "Invalid addressType",
		})
	}
}
