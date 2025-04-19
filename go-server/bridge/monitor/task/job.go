package task

import (
	"coinstore/bridge/chains"
	"coinstore/bridge/chains/ethereum"
	"coinstore/bridge/chains/tron"
	"coinstore/db"
	"coinstore/model"
	"coinstore/utils"
	"errors"
	"github.com/calmw/clog"
	"github.com/calmw/tron-sdk/pkg/address"
	"github.com/ethereum/go-ethereum/common"
	"log"
	"strings"
	"time"
)

type Monitor struct {
	log                  clog.Logger
	ConcurrencyLimit     int
	FailedOrdersChain    chan model.BridgeTx
	ConcurrencyLimitChan chan struct{}
}

// NewMonitor concurrencyLimit 最多并发数,比如3～5
func NewMonitor(concurrencyLimit int) *Monitor {
	logger := clog.Root().New("module", "monitor")
	return &Monitor{
		log:                  logger,
		ConcurrencyLimit:     concurrencyLimit,
		FailedOrdersChain:    make(chan model.BridgeTx, 10000),
		ConcurrencyLimitChan: make(chan struct{}, concurrencyLimit),
	}
}

func (m *Monitor) ProcessFailedOrder() {
	for order := range m.FailedOrdersChain {
		m.ConcurrencyLimitChan <- struct{}{}
		go func(o model.BridgeTx) {
			defer func() {
				<-m.ConcurrencyLimitChan
				m.log.Debug("🍺 处理订单完成", "ID", order.Id)
			}()
			m.RetryFailedOrder(o)
		}(order)
	}
}

func (m *Monitor) RetryFailedOrder(order model.BridgeTx) {
	m.log.Debug("🍺 处理订单开始", "ID", order.Id)
	if order.ExecuteStatus > 0 {
		return
	}
	bridgeData, err := model.BytesToMsg(order.BridgeMsg)
	if err != nil {
		m.log.Error("处理订单 BytesToMsg ", "error", err)
		return
	}
	// 过滤异常，ResourceId为0的情况
	if bridgeData.ResourceId == [32]byte{} {
		m.log.Error("处理订单", "error", errors.New("异常数据"))
		model.UpdateVoteStatus(bridgeData, 1)
		model.UpdateExecuteStatus(bridgeData, 1, "-", time.Now().Format("2006-01-02 15:04:05"))
		return
	}
	if int64(bridgeData.Source) == 0 || int64(bridgeData.Destination) == 0 {
		m.DelFailedOrder(order.Hash)
		return
	}
	if order.VoteStatus > 0 { // 投票已经成功，执行execute
		if order.DestinationChainId == 3 {
			writer := ethereum.Writers[int(order.DestinationChainId)]
			m.log.Debug("🍺 重试execute", "sourceId", bridgeData.Source, "destinationId", bridgeData.Destination, "depositNonce", bridgeData.DepositNonce)
			metadata := bridgeData.Payload[0].([]byte)
			data := chains.ConstructGenericProposalData(metadata)
			toHash := append(common.HexToAddress(writer.Cfg.BridgeContractAddress).Bytes(), data...)
			dataHash := utils.Keccak256(toHash)
			writer.ExecuteProposal(bridgeData, data, dataHash)
		} else {
			writer := tron.WritersTron
			m.log.Debug("🍺 重试execute", "sourceId", bridgeData.Source, "destinationId", bridgeData.Destination, "depositNonce", bridgeData.DepositNonce)
			metadata := bridgeData.Payload[0].([]byte)
			data := chains.ConstructGenericProposalData(metadata)
			bridgeAddress, err := address.Base58ToAddress(writer.Cfg.BridgeContractAddress)
			if err != nil {
				m.log.Debug("🍺 重试execute", "地址转转错误", err)
				return
			}
			bridgeEthAddress := "0x" + strings.TrimPrefix(bridgeAddress.Hex(), "0x41")
			toHash := append(common.HexToAddress(bridgeEthAddress).Bytes(), data...)
			dataHash := utils.Keccak256(toHash)
			writer.ExecuteProposal(bridgeData, data, dataHash)
		}
	} else { // 投票未成功，vote + execute
		if order.DestinationChainId == 3 {
			writer := tron.WritersTron
			m.log.Debug("🍺 重试vote+execute", "sourceId", bridgeData.Source, "destinationId", bridgeData.Destination, "depositNonce", bridgeData.DepositNonce)
			writer.CreateProposal(bridgeData)
		} else {
			writer := ethereum.Writers[int(bridgeData.Destination)]
			m.log.Debug("🍺 重试vote+execute", "sourceId", bridgeData.Source, "destinationId", bridgeData.Destination, "depositNonce", bridgeData.DepositNonce)
			writer.CreateProposal(bridgeData)
		}
	}
}

func FailedTask() {
	monitor := NewMonitor(4)
	orders, err := monitor.FindFailedOrder()
	if err != nil {
		return
	}
	monitor.log.Debug("🍺 开始添加失败任务", "当前失败的总任务数量", len(orders))
	if len(monitor.FailedOrdersChain) > monitor.ConcurrencyLimit*150 {
		monitor.log.Error("跳过该轮添加")
		return
	}
	for i, order := range orders {
		if i >= monitor.ConcurrencyLimit*100 {
			return
		}
		monitor.FailedOrdersChain <- order
	}
}

func (m *Monitor) FindFailedOrder() ([]model.BridgeTx, error) {
	var orders []model.BridgeTx
	err := db.DB.Model(model.BridgeTx{}).Where("vote_status=0 or execute_status=0").Find(&orders).Error
	if err != nil {
		m.log.Error("处理订单 FindFailedOrder ", "error", err)
		return nil, err
	}

	return orders, nil
}

func (m *Monitor) CountFailedOrder() (int, error) {
	var orders []model.BridgeTx
	err := db.DB.Model(model.BridgeTx{}).Where("vote_status=0 or execute_status=0").Find(&orders).Error
	if err != nil {
		m.log.Error("处理订单 CountFailedOrder ", "error", err)
		return 0, err
	}

	return len(orders), nil
}

func (m *Monitor) DelFailedOrder(hash string) {
	var orders []model.BridgeTx

	err := db.DB.Model(model.BridgeTx{}).Where("`hash`=?", hash).Delete(&orders).Error
	if err != nil {
		log.Println(err)
	}
}
