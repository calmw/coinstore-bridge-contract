package service

import (
	"coinstore/db"
	"coinstore/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"strings"
	"time"
)

type Query struct {
	PageNum           int    `form:"page_num"`
	PageSize          int    `form:"page_size"`
	From              string `form:"from"`
	Recipient         string `form:"recipient"`
	FromToken         string `form:"from_token"`
	ToToken           string `form:"to_token"`
	Status            int    `form:"status"`
	DepositAtStart    int64  `form:"deposit_at_start"`
	DepositAtEnd      int64  `form:"deposit_at_end"`
	SourceTxHash      string `form:"source_tx_hash"`
	DestinationTxHash string `form:"destination_tx_hash"`
}

type Response struct {
	Id                      uint64 `json:"id"`
	ResourceId              string `json:"resource_id"`
	Amount                  string `json:"amount"`
	From                    string `json:"from"`
	Recipient               string `json:"recipient"`
	SourceChainId           int    `json:"source_chain_id"`
	SourceTokenAddress      string `json:"source_token_address"`
	SourceTxHash            string `json:"source_tx_hash"`
	DestinationChainId      int    `json:"destination_chain_id"`
	DestinationTokenAddress string `json:"destination_token_address"`
	DestinationTxHash       string `json:"destination_tx_hash"`
	BridgeStatus            int    `json:"bridge_status"`
	DepositAt               string `json:"deposit_at"`
	ReceiveAt               string `json:"receive_at"`
}

func BridgeTx(c *gin.Context) {
	var q Query
	if c.ShouldBindQuery(&q) != nil {
		c.JSON(200, gin.H{
			"code": 1,
			"msg":  "参数错误",
		})
		return
	}
	fmt.Println(q)
	fmt.Println(c.Query("source_tx_hash"))

	var total int64
	var page = 1
	var pageSize = 10
	if q.PageNum > 0 {
		page = q.PageNum
	}
	if q.PageSize > 0 {
		pageSize = q.PageSize
	}
	offset := (page - 1) * pageSize

	// 更新记录
	var records []model.BridgeTx
	var data []Response
	tx := db.DB.Model(&model.BridgeTx{})

	if len(q.From) > 0 {
		tx = tx.Where("caller=?", strings.ToLower(q.From))
	}
	if len(q.Recipient) > 0 {
		tx = tx.Where("receiver=?", strings.ToLower(q.Recipient))
	}
	if len(q.FromToken) > 0 {
		tx = tx.Where("source_token_address=?", strings.ToLower(q.FromToken))
	}
	if len(q.ToToken) > 0 {
		tx = tx.Where("destination_token_address=?", strings.ToLower(q.ToToken))
	}
	if q.Status > 0 {
		tx = tx.Where("bridge_status=?", q.Status)
	}
	if q.DepositAtStart > 0 {
		tx = tx.Where("deposit_at >=?", time.Unix(q.DepositAtStart, 0))
	}
	if q.DepositAtEnd > 0 {
		tx = tx.Where("deposit_at <=?", time.Unix(q.DepositAtEnd, 0))
	}
	if len(q.SourceTxHash) > 0 {
		tx = tx.Where("deposit_hash=?", strings.ToLower(q.SourceTxHash))
	}
	if len(q.DestinationTxHash) > 0 {
		tx = tx.Where("execute_hash=?", strings.ToLower(q.DestinationTxHash))
	}
	tx.Count(&total)
	err := tx.Offset(offset).Order("deposit_at DESC").Limit(pageSize).Find(&records).Error
	if err != nil {
		c.JSON(200, gin.H{
			"code": 1,
			"msg":  "系统繁忙",
		})
		return
	}

	for _, record := range records {
		data = append(data, Response{
			Id:                      record.Id,
			ResourceId:              record.ResourceId,
			Amount:                  record.Amount.String(),
			From:                    record.Caller,
			Recipient:               record.Receiver,
			SourceChainId:           record.SourceChainId,
			SourceTokenAddress:      record.SourceTokenAddress,
			SourceTxHash:            record.DepositHash,
			DestinationChainId:      record.DestinationChainId,
			DestinationTokenAddress: record.DestinationTokenAddress,
			DestinationTxHash:       record.ExecuteHash,
			BridgeStatus:            record.BridgeStatus,
			DepositAt:               record.DepositAt,
			ReceiveAt:               record.ReceiveAt,
		})
	}

	c.JSON(200, gin.H{
		"code":  0,
		"msg":   "OK",
		"total": total,
		"data":  data,
	})
}
