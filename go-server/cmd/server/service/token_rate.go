package service

import (
	"coinstore/abi"
	"coinstore/cmd/server/token"
	"coinstore/db"
	"coinstore/model"
	"fmt"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
	"github.com/shopspring/decimal"
	"log"
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

	price, ok := token.ExTokenPriceData.Get(strings.ToUpper(tokenName))
	if !ok {
		log.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
		c.JSON(200, gin.H{
			"code": 1,
			"msg":  "failed",
		})
		return
	}
	timestamp := time.Now().Unix()
	signature := ""
	chainIdDeci, err1 := decimal.NewFromString(chainId)
	priceDeci, err2 := decimal.NewFromString(price)
	if err1 != nil || err2 != nil {
		c.JSON(200, gin.H{
			"code": 1,
			"msg":  "failed",
		})
		return
	}
	priceDeci = priceDeci.Mul(decimal.NewFromInt(1e6))
	if chainId == "3" {
		priceSignature, err := abi.TronPriceSignature(chainIdDeci.BigInt(), priceDeci.BigInt(), ethcommon.HexToAddress(tokenAddress))
		if err != nil {
			c.JSON(200, gin.H{
				"code": 1,
				"msg":  "failed",
			})
			return
		}
		signature = fmt.Sprintf("%x", priceSignature)
	} else {
		priceSignature, err := abi.EvmPriceSignature(chainIdDeci.BigInt(), priceDeci.BigInt(), ethcommon.HexToAddress(tokenAddress))
		if err != nil {
			c.JSON(200, gin.H{
				"code": 1,
				"msg":  "failed",
			})
			return
		}
		signature = fmt.Sprintf("%x", priceSignature)
	}

	c.JSON(200, gin.H{
		"code": 0,
		"msg":  "OK",
		"data": map[string]string{
			"price":     priceDeci.String(),
			"timestamp": fmt.Sprintf("%d", timestamp),
			"signature": signature,
		},
	})
}
