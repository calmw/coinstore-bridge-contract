package service

import (
	"coinstore/abi"
	"coinstore/cmd/server/token"
	"coinstore/db"
	"coinstore/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/shopspring/decimal"
	"math/big"
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
	fmt.Println("~~~~~~~~~~~1 ")
	var ok bool
	price := "1"
	tokenName = strings.ToUpper(tokenName)
	if !strings.Contains(tokenName, "USDT") {
		price, ok = token.ExTokenPriceData.Get(tokenName)
		if !ok {
			c.JSON(200, gin.H{
				"code": 1,
				"msg":  "failed",
			})
			return
		}
	}

	fmt.Println("~~~~~~~~~~~2 ")
	timestamp := time.Now().Unix()
	signature := ""
	fmt.Println(price)
	priceDeci, err := decimal.NewFromString(price)
	fmt.Println(priceDeci.String())
	if err != nil {
		c.JSON(200, gin.H{
			"code": 1,
			"msg":  "failed",
		})
		return
	}
	priceDeci = priceDeci.Mul(decimal.NewFromInt(1e6))
	if tokenInfo.ChainId > 50000000 {
		priceSignature, err := abi.TronPriceSignature(big.NewInt(int64(tokenInfo.ChainId)), priceDeci.BigInt(), big.NewInt(timestamp))
		if err != nil {
			c.JSON(200, gin.H{
				"code": 1,
				"msg":  "failed",
			})
			return
		}
		signature = fmt.Sprintf("%x", priceSignature)
	} else {
		fmt.Println("~~~~~~~~~~~ 5 ")
		priceSignature, err := abi.EvmPriceSignature(big.NewInt(int64(tokenInfo.ChainId)), priceDeci.BigInt(), big.NewInt(timestamp))
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
