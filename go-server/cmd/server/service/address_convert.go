package service

import (
	tronAddress "github.com/fbsobreira/gotron-sdk/pkg/address"
	"github.com/gin-gonic/gin"
	"strings"
)

func ConvertAddress(c *gin.Context) {
	var q Query
	var res string
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

	if !isValidEthAddress(address) && !isValidTronAddress(address) {
		c.JSON(200, gin.H{
			"code": 1,
			"msg":  "invalid address",
		})
		return
	}
	if (chainIdFrom == "1" || chainIdFrom == "2" || chainIdFrom == "4") && (chainIdTo == "3") {
		if !isValidEthAddress(address) {
			c.JSON(200, gin.H{
				"code": 1,
				"msg":  "invalid address",
			})
			return
		}
		toAddress := tronAddress.HexToAddress("0x41" + strings.TrimPrefix(address, "0x"))
		res = toAddress.String()
	} else if (chainIdFrom == "3") && (chainIdTo == "1" || chainIdTo == "2" || chainIdTo == "4") {
		if !isValidTronAddress(address) {
			c.JSON(200, gin.H{
				"code": 1,
				"msg":  "invalid address",
			})
			return
		}
		toAddress, err := tronAddress.Base58ToAddress(address)
		if err != nil {
			c.JSON(200, gin.H{
				"code": 1,
				"msg":  "invalid address",
			})
			return
		}
		res = "0x" + strings.TrimPrefix(toAddress.Hex(), "0x41")
	} else {
		res = address
	}
	c.JSON(200, gin.H{
		"code": 0,
		"msg":  "OK",
		"data": res,
	})
}
