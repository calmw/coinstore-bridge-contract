package model

import (
	"bytes"
	"coinstore/bridge/msg"
	"coinstore/db"
	"encoding/gob"
	"errors"
	"fmt"
	log "github.com/calmw/clog"
	"github.com/shopspring/decimal"
	"github.com/status-im/keycard-go/hexutils"
	"gorm.io/gorm"
	"strings"
)

type BridgeTx struct {
	Id                      uint64          `gorm:"primaryKey" json:"id"`
	BridgeData              string          `gorm:"column:bridge_data;type:varchar(1000);comment:'跨链数据'" json:"bridge_data"`
	BridgeMsg               []byte          `gorm:"column:bridge_msg;comment:'跨链nsg'" json:"bridge_msg"`
	ResourceId              ResourceId      `gorm:"column:resource_id;type:varchar(100);comment:'resource ID'" json:"resource_id"`
	Hash                    string          `gorm:"column:hash;unique;comment:'唯一索引'" json:"hash"`
	VoteStatus              int             `gorm:"column:vote_status;default:0;comment:'vote 0失败，1成功'" json:"vote_status"`
	ExecuteStatus           int             `gorm:"column:execute_status;default:0;comment:'execute 0失败，1成功'" json:"execute_status"`
	Amount                  decimal.Decimal `gorm:"column:amount;type:decimal(20,0);comment:'跨链数额'" json:"amount"`
	Fee                     decimal.Decimal `gorm:"column:fee;type:decimal(20,0);comment:'跨链费用，万分比'" json:"fee"`
	Caller                  string          `gorm:"column:caller;comment:'链链发起者地址'" json:"caller"`
	Receiver                string          `gorm:"column:receiver;comment:'目标链接受者地址'" json:"receiver"`
	SourceChainId           ChainId         `gorm:"column:source_chain_id;comment:'源链ID'" json:"source_chain_id"`
	SourceTokenAddress      string          `gorm:"column:source_token_address;comment:'源链token地址'" json:"source_token_address"`
	DestinationChainId      ChainId         `gorm:"column:destination_chain_id;comment:'目标链ID'" json:"destination_chain_id"`
	DestinationTokenAddress string          `gorm:"column:destination_token_address;comment:'目标链token地址'" json:"destination_token_address"`
	BridgeStatus            int             `gorm:"column:bridge_status;type:tinyint;comment:'跨链状态 1 源链deposit成功 2 目标链执行成功 3 失败';default:1" json:"bridge_status"`
	DepositHash             string          `gorm:"column:deposit_hash;comment:'deposit tx hash'" json:"deposit_hash"`
	ExecuteHash             string          `gorm:"column:execute_hash;comment:'execute tx hash'" json:"execute_hash"`
	DepositAt               string          `gorm:"column:deposit_at;comment:'跨链发起时间'" json:"deposit_at"`
	ReceiveAt               string          `gorm:"column:receive_at;comment:'跨链到账时间'" json:"receive_at"`
	DeletedAt               gorm.DeletedAt  `gorm:"index"`
	Version                 uint32          `gorm:"not null;default:0;comment:'乐观锁'" json:"version"`
	BridgeGasFee            decimal.Decimal `gorm:"column:amount;type:decimal(20,0);comment:'跨链桥今日累计Gas费用'" json:"bridge_gas_fee"`
}

type ChainId int64

func (c ChainId) String() string {
	switch c {
	case 112301:
		return "Tantin"
	case 12302:
		return "Tantin"
	case 97:
		return "BSC"
	case 56:
		return "BSC"
	case 3448148188:
		return "TRON"
	case 728126428:
		return "TRON"
	case 1:
		return "Ethereum"
	case 11155111:
		return "Ethereum"
	default:
		return ""
	}
}

type ResourceId string

func (r ResourceId) TokenName() string {
	switch strings.ToLower(string(r)) {
	case "0xac589789ed8c9d2c61f17b13369864b5f181e58eba230a6ee4ec4c3e7750cd1b":
		return "ETH"
	case "0xac589789ed8c9d2c61f17b13369864b5f181e58eba230a6ee4ec4c3e7750cd1c":
		return "ETH"
	case "0xac589789ed8c9d2c61f17b13369864b5f181e58eba230a6ee4ec4c3e7750cd1d":
		return "USDT"
	case "0xac589789ed8c9d2c61f17b13369864b5f181e58eba230a6ee4ec4c3e7750cd1e":
		return "USDC"
	default:
		return ""
	}
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

func SaveBridgeOrder(log log.Logger, m msg.Message, amount decimal.Decimal, resourceId, caller, receiver, sourceTokenAddress, destinationTokenAddress, depositTxHash, dateTime string, fee decimal.Decimal) {
	log.Debug("检查订单是否存在", "Destination", m.Destination, "DepositNonce", m.DepositNonce)
	var bridgeOrder BridgeTx
	resourceIdHex := "0x" + hexutils.BytesToHex(m.ResourceId[:])
	key := strings.ToLower(fmt.Sprintf("%s%s%s%s", resourceIdHex, m.Source.Big().String(), m.Destination.Big().String(), m.DepositNonce.Big().String()))

	// 记录不存在就插入
	err := db.DB.Model(BridgeTx{}).Where("hash=?", key).First(&bridgeOrder).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		orderData, err := MsgDataToBytes(m)
		if err != nil {
			log.Error(err.Error())
			log.Debug("SaveBridgeOrder", "error", err.Error())
			return
		}
		if !strings.HasPrefix(depositTxHash, "0x") {
			depositTxHash = "0x" + depositTxHash
		}
		bridgeOrder = BridgeTx{
			BridgeData:              fmt.Sprintf("%x", m.Payload[0].([]byte)),
			BridgeMsg:               orderData,
			Hash:                    key,
			Amount:                  amount,
			Fee:                     fee,
			ResourceId:              ResourceId(resourceId),
			Caller:                  caller,
			Receiver:                receiver,
			SourceChainId:           ChainId(m.Source),
			SourceTokenAddress:      sourceTokenAddress,
			DestinationChainId:      ChainId(m.Destination),
			DestinationTokenAddress: destinationTokenAddress,
			DepositAt:               dateTime,
			DepositHash:             depositTxHash,
		}
		log.Debug("插入订单数据", "Destination", m.Destination, "DepositNonce", m.DepositNonce)
		err = db.DB.Model(BridgeTx{}).Create(&bridgeOrder).Error
		if err != nil {
			log.Debug("SaveBridgeOrder", "error", err.Error())
		}
		return
	}
	if err == nil {
		log.Debug("订单已存在", "Destination", m.Destination, "DepositNonce", m.DepositNonce)
	}
}

func UpdateExecuteStatus(m msg.Message, status int, executeTxHash, dateTime string) {

	log.Debug("更新execute数据", "Destination", m.Destination, "DepositNonce", m.DepositNonce)
	resourceIdHex := "0x" + hexutils.BytesToHex(m.ResourceId[:])
	key := strings.ToLower(fmt.Sprintf("%s%s%s%s", resourceIdHex, m.Source.Big().String(), m.Destination.Big().String(), m.DepositNonce.Big().String()))
	// 更新记录
	var record BridgeTx
	err := db.DB.Model(&BridgeTx{}).Where("hash=?", key).First(&record).Error
	if err != nil {
		log.Debug("更新execute数据", "error", err)
		return
	}
	if record.ExecuteStatus == 1 {
		return
	}
	err = db.DB.Model(&BridgeTx{}).Where("hash=?", key).Updates(map[string]interface{}{
		"execute_hash":   executeTxHash,
		"execute_status": status,
		"bridge_status":  2,
		"receive_at":     dateTime,
	}).Error
	if err != nil {
		log.Debug("更新execute数据", "error", err)
		return
	}
}

func UpdateVoteStatus(m msg.Message, voteStatus int) {
	log.Debug("更新vote数据", "Destination", m.Destination, "DepositNonce", m.DepositNonce)
	resourceIdHex := "0x" + hexutils.BytesToHex(m.ResourceId[:])
	key := strings.ToLower(fmt.Sprintf("%s%s%s%s", resourceIdHex, m.Source.Big().String(), m.Destination.Big().String(), m.DepositNonce.Big().String()))
	// 更新记录
	var record BridgeTx
	err := db.DB.Model(&BridgeTx{}).Where("hash=?", key).First(&record).Error
	if err != nil {
		log.Debug("更新vote数据", "error", err)
		return
	}
	if record.VoteStatus == 1 {
		return
	}
	err = db.DB.Model(&BridgeTx{}).Where("hash=?", key).Updates(map[string]interface{}{
		"vote_status": voteStatus,
	}).Error
	if err != nil {
		log.Debug("更新vote数据", "error", err)
		return
	}
}

func GetBridgeTxStatus(m msg.Message) (int, int, error) {
	resourceIdHex := "0x" + hexutils.BytesToHex(m.ResourceId[:])
	key := strings.ToLower(fmt.Sprintf("%s%s%s%s", resourceIdHex, m.Source.Big().String(), m.Destination.Big().String(), m.DepositNonce.Big().String()))
	// 更新记录
	var record BridgeTx
	err := db.DB.Model(&BridgeTx{}).Where("hash=?", key).First(&record).Error
	if err != nil {
		return 0, 0, err
	}
	return record.VoteStatus, record.ExecuteStatus, nil
}
