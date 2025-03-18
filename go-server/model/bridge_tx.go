package model

import (
	"bytes"
	"coinstore/bridge/msg"
	"coinstore/db"
	"encoding/gob"
	"errors"
	"fmt"
	log "github.com/calmw/blog"
	"github.com/shopspring/decimal"
	"github.com/status-im/keycard-go/hexutils"
	"gorm.io/gorm"
)

type BridgeTx struct {
	gorm.Model
	Data                    []byte          `gorm:"data"` // msg binary data
	Hash                    string          `gorm:"hash"`
	VoteStatus              bool            `gorm:"vote_status"`               //  vote 失败，成功
	Status                  bool            `gorm:"status"`                    //  execute 失败，成功
	Trash                   bool            `gorm:"trash"`                     //  放到回收站中
	Amount                  decimal.Decimal `gorm:"type:bigint(30);amount"`    //  跨链数额
	SourceChainId           int             `gorm:"source_chain_id"`           //  源链ID
	SourceTokenAddress      string          `gorm:"source_token_address"`      //  源链token地址
	DestinationChainId      int             `gorm:"destination_chain_id"`      //  目标链ID
	DestinationTokenAddress string          `gorm:"destination_token_address"` //  目标链ID
	BridgeStatus            int             `gorm:"default:1;bridge_status"`   //  跨链状态 1 源链deposit成功 2 目标链执行成功
}

func MsgDataToBytes(el msg.Message) ([]byte, error) {
	buf := new(bytes.Buffer)
	enc := gob.NewEncoder(buf)
	if err := enc.Encode(el); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func BytesToMsg(b []byte) (msg.Message, error) {
	buf := bytes.NewBuffer(b)
	dec := gob.NewDecoder(buf)
	var m *msg.Message
	if err := dec.Decode(&m); err != nil {
		return msg.Message{}, err
	}

	return *m, nil
}

func SaveBridgeOrder(log log.Logger, m msg.Message, amount decimal.Decimal, sourceTokenAddress, destinationTokenAddress string) {
	log.Debug("🐧 检查订单是否存在", "Destination", m.Destination, "DepositNonce", m.DepositNonce)
	var bridgeOrder BridgeTx
	resourceIdHex := "0x" + hexutils.BytesToHex(m.ResourceId[:])
	key := []byte(fmt.Sprintf("%s%d%d%d", resourceIdHex, m.Source, m.Destination, m.DepositNonce))

	// 记录不存在就插入
	err := db.DB.Model(BridgeTx{}).Where("hash=?", string(key)).First(&bridgeOrder).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		orderData, err := MsgDataToBytes(m)
		if err != nil {
			log.Error(err.Error())
			log.Debug("🐧 SaveBridgeOrder", "error", err.Error())
			return
		}
		bridgeOrder = BridgeTx{
			Data:                    orderData,
			Hash:                    string(key),
			Amount:                  amount,
			SourceChainId:           int(m.Source),
			SourceTokenAddress:      sourceTokenAddress,
			DestinationChainId:      int(m.Destination),
			DestinationTokenAddress: destinationTokenAddress,
		}
		log.Debug("🐧 插入订单数据", "Destination", m.Destination, "DepositNonce", m.DepositNonce)
		err = db.DB.Model(BridgeTx{}).Create(&bridgeOrder).Error
		if err != nil {
			log.Debug("🐧 SaveBridgeOrder", "error", err.Error())
		}
		return
	}
	if err == nil {
		log.Debug("🐧 订单已存在", "Destination", m.Destination, "DepositNonce", m.DepositNonce)
	}
}

func UpdateExecuteStatus(m msg.Message, status bool) {

	log.Debug("🐧 更新execute数据", "Destination", m.Destination, "DepositNonce", m.DepositNonce)
	resourceIdHex := "0x" + hexutils.BytesToHex(m.ResourceId[:])
	key := []byte(fmt.Sprintf("%s%d%d%d", resourceIdHex, m.Source, m.Destination, m.DepositNonce))
	// 更新记录
	err := db.DB.Model(&BridgeTx{}).Where("hash=?", string(key)).Updates(map[string]interface{}{
		"status": status,
	}).Error
	if err != nil {
		log.Error(err.Error())
		return
	}
}

func UpdateVoteStatus(m msg.Message, voteStatus bool) {
	log.Debug("🐧 更新vote数据", "Destination", m.Destination, "DepositNonce", m.DepositNonce)
	resourceIdHex := "0x" + hexutils.BytesToHex(m.ResourceId[:])
	key := []byte(fmt.Sprintf("%s%d%d%d", resourceIdHex, m.Source, m.Destination, m.DepositNonce))
	// 更新记录
	err := db.DB.Model(&BridgeTx{}).Where("hash=?", string(key)).Updates(map[string]interface{}{
		"vote_status": voteStatus,
	}).Error
	if err != nil {
		log.Error(err.Error())
		return
	}
}
