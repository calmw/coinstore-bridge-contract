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

// 波场地址格式检测
func isValidTronAddress(address string) bool {
	// 有效的Base58字符集，除去容易混淆的数字和字母
	validChars := "123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz"
	for _, char := range address {
		if !strings.ContainsRune(validChars, char) {
			return false // 包含无效字符
		}
	}
	// 检查长度（通常是34个字符）但这不是绝对标准，因为长度也可能根据版本字节变化
	if len(address) != 34 { // 示例长度，实际应根据具体版本调整
		return false // 长度不符合预期（可根据需要调整）
	}
	return true
}

// 以太坊地址格式检测
func isValidEthAddress(address string) bool {
	re := regexp.MustCompile("^0x[0-9a-fA-F]{40}$")
	return re.MatchString(address)
}
