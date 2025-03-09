package model

import (
	"bytes"
	"coinstore/bridge/msg"
	"coinstore/db"
	"encoding/gob"
	"errors"
	"fmt"
	log "github.com/calmw/blog"
	"github.com/status-im/keycard-go/hexutils"
	"gorm.io/gorm"
)

type BridgeOrder struct {
	gorm.Model
	Data       []byte `gorm:"data"` // msg binary data
	Hash       string `gorm:"hash"`
	VoteStatus bool   `gorm:"vote_status"` //  vote 失败，成功
	Status     bool   `gorm:"status"`      //  execute 失败，成功
	Trash      bool   `gorm:"trash"`       //  放到回收站中
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

func SaveBridgeOrder(m msg.Message, log log.Logger) {
	log.Debug("🐧 检查订单是否存在, %d %d", m.Destination, m.DepositNonce)
	var bridgeOrder BridgeOrder
	resourceIdHex := "0x" + hexutils.BytesToHex(m.ResourceId[:])
	key := []byte(fmt.Sprintf("%s%d%d%d", resourceIdHex, m.Source, m.Destination, m.DepositNonce))

	// 记录不存在就插入
	err := db.DB.Model(BridgeOrder{}).Where("hash=?", string(key)).First(&bridgeOrder).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		orderData, err := MsgDataToBytes(m)
		if err != nil {
			log.Error(err.Error())
			return
		}
		bridgeOrder = BridgeOrder{
			Data: orderData,
			Hash: string(key),
		}
		log.Debug("🐧 插入订单数据, %d %d", m.Destination, m.DepositNonce)
		err = db.DB.Model(BridgeOrder{}).Create(&bridgeOrder).Error
		if err != nil {
			log.Error(err.Error())
		}
		return
	}
	if err == nil {
		log.Debug("🐧 订单已存在, %d %d", m.Destination, m.DepositNonce)
	}
}

func UpdateExecuteStatus(m msg.Message, status bool) {

	log.Debug("🐧 更新execute数据 %d %d", m.Destination, m.DepositNonce)
	resourceIdHex := "0x" + hexutils.BytesToHex(m.ResourceId[:])
	key := []byte(fmt.Sprintf("%s%d%d%d", resourceIdHex, m.Source, m.Destination, m.DepositNonce))
	// 更新记录
	err := db.DB.Model(&BridgeOrder{}).Where("hash=?", string(key)).Updates(map[string]interface{}{
		"status": status,
	}).Error
	if err != nil {
		log.Error(err.Error())
		return
	}
}

func UpdateVoteStatus(m msg.Message, voteStatus bool) {
	log.Debug("🐧 更新vote数据 %d %v", m.DepositNonce, voteStatus)
	resourceIdHex := "0x" + hexutils.BytesToHex(m.ResourceId[:])
	key := []byte(fmt.Sprintf("%s%d%d%d", resourceIdHex, m.Source, m.Destination, m.DepositNonce))
	// 更新记录
	err := db.DB.Model(&BridgeOrder{}).Where("hash=?", string(key)).Updates(map[string]interface{}{
		"vote_status": voteStatus,
	}).Error
	if err != nil {
		log.Error(err.Error())
		return
	}
}
