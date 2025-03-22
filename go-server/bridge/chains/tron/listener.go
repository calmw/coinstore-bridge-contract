package tron

import (
	"coinstore/binding"
	"coinstore/bridge/chains"
	"coinstore/bridge/chains/ethereum"
	"coinstore/bridge/config"
	"coinstore/bridge/core"
	"coinstore/bridge/msg"
	"coinstore/db"
	"coinstore/model"
	"coinstore/utils"
	"errors"
	"fmt"
	log "github.com/calmw/clog"
	"github.com/shopspring/decimal"
	"github.com/status-im/keycard-go/hexutils"
	"math/big"
	"strings"
	"time"
)

var BlockRetryInterval = time.Second * 5
var BlockRetryLimit = 5
var ErrFatalPolling = errors.New("bridge block polling failed")
var ListenersTron *Listener

type Listener struct {
	cfg         config.Config
	conn        *Connection
	Router      chains.Router
	log         log.Logger
	latestBlock core.LatestBlock
	stop        <-chan int
	sysErr      chan<- error
}

// NewListener creates and returns a Listener
func NewListener(conn *Connection, cfg *config.Config, log log.Logger, stop <-chan int, sysErr chan<- error) *Listener {
	listener := Listener{
		cfg:    *cfg,
		conn:   conn,
		log:    log,
		stop:   stop,
		sysErr: sysErr,
	}

	ListenersTron = &listener
	log.Debug("new listener", "id", cfg.ChainId)
	return &listener
}

func (l *Listener) start() error {
	l.log.Debug("Starting Listener...")

	go func() {
		err := l.pollBlocks()
		if err != nil {
			l.log.Error("Polling blocks failed", "err", err)
		}
	}()

	return nil
}

func (l *Listener) setRouter(r chains.Router) {
	l.Router = r
}

func (l *Listener) pollBlocks() error {
	var currentBlock = l.cfg.StartBlock
	l.log.Info("Polling Blocks...", "block", currentBlock)

	var retry = BlockRetryLimit
	for {
		select {
		case <-l.stop:
			return errors.New("polling terminated")
		default:
			// No more retries, goto next block
			if retry == 0 {
				l.log.Error("Polling failed, retries exceeded")
				l.sysErr <- ErrFatalPolling
				return nil
			}

			latestBlock, err := l.LatestBlock()
			if err != nil {
				l.log.Error("Unable to get latest block", "block", currentBlock, "err", err)
				retry--
				time.Sleep(BlockRetryInterval)
				continue
			}

			if big.NewInt(0).Sub(latestBlock, currentBlock).Cmp(l.cfg.BlockConfirmations) == -1 {
				l.log.Debug("Block not ready, will retry", "target", currentBlock, "latest", latestBlock)
				time.Sleep(BlockRetryInterval)
				continue
			}

			// Parse out events
			err = l.getDepositEventsForBlock(currentBlock)
			if err != nil {
				l.log.Error("Failed to get events for block", "block", currentBlock, "err", err)
				retry--
				continue
			}

			err = l.StoreBlock(currentBlock)
			if err != nil {
				l.log.Error("Failed to write latest block to blockstore", "block", currentBlock, "err", err)
			}

			currentBlock.Add(currentBlock, big.NewInt(1))
			retry = BlockRetryLimit
		}
	}
}

func (l *Listener) getDepositEventsForBlock(latestBlock *big.Int) error {
	l.log.Debug("Querying block for deposit events", "block", latestBlock)
	latestBlock = big.NewInt(55444496)
	data, err := GetEventData(latestBlock.Int64())
	if err != nil {
		return err
	}

	for _, log := range data {
		var m msg.Message
		l.log.Debug("get events:")
		l.log.Debug("ResourceID", log.ResourceID)
		l.log.Debug("DestinationChainId", log.DestinationChainId)
		//l.log.Debug("Sender", records.Sender)
		l.log.Debug("Data", log.Data)
		//var depositNonce big.Int
		var bigIntD big.Int
		var bigIntN big.Int
		//var success bool
		destinationChainId, success := bigIntD.SetString(log.DestinationChainId, 10)
		if !success || destinationChainId == nil {
			return errors.New("转换失败")
		}
		depositNonce, success := bigIntN.SetString(log.DepositNonce, 10)
		if !success || depositNonce == nil {
			return errors.New("转换失败")
		}
		recordsData, err := utils.GenerateBridgeDepositRecordsData(destinationChainId, depositNonce)
		if err != nil {
			return fmt.Errorf("generateBridgeDepositRecordsData error %v", err)
		}
		record, err := GetDepositRecord(binding.OwnerAccount, l.cfg.BridgeContractAddress, recordsData)
		if err != nil {
			return fmt.Errorf("getDepositRecord error %v", err)
		}
		m = msg.NewGenericTransfer(
			msg.ChainId(l.cfg.ChainId),
			msg.ChainId(destinationChainId.Int64()),
			msg.Nonce(depositNonce.Int64()),
			msg.ResourceId(hexutils.HexToBytes(log.ResourceID)),
			hexutils.HexToBytes(log.Data),
		)

		//// 获取目标链的信息
		dl, ok := ethereum.Listeners[int(destinationChainId.Int64())]
		if !ok {
			l.log.Error("destination listener not found", "chainId", destinationChainId)
			return errors.New(fmt.Sprintf("destination listener not found, chainId %d", destinationChainId))
		}
		_, t, _, err := dl.BridgeContract.GetToeknInfoByResourceId(nil, msg.ResourceId(hexutils.HexToBytes(log.ResourceID)))
		if err != nil {
			l.log.Error("destination token info not found", "chainId", destinationChainId)
		}
		fmt.Println("~~~~~~~~~t  ", t.String())
		fmt.Println("~~~~~~~~~t  ", m)
		//_, s, _, err := l.bridgeContract.GetTokenInfoByResourceId(nil, records.ResourceID)
		//if err != nil {
		//	l.log.Error("source token info not found", "chainId", records.DestinationChainId)
		//}
		//amount, caller, receiver, err := utils.ParseBridgeData(records.Data)
		// 保存到数据库
		model.SaveBridgeOrder(l.log, m, amount, fmt.Sprintf("%x", record.ResourceID), caller, receiver, strings.ToLower(s.String()), strings.ToLower(t.String()), log.TxHash, time.Unix(record.Ctime.Int64(), 0).Format("2006-01-02 15:04:05"))

		err = l.Router.Send(m)
		if err != nil {
			l.log.Error("subscription error: failed to route message", "err", err)
			return errors.New(fmt.Sprintf("subscription error: failed to route message,err %v", err))
		}
	}

	return nil
}

func (l *Listener) LatestBlock() (*big.Int, error) {
	header, err := l.conn.ClientTron().GetNowBlock()
	if err != nil {
		return nil, err
	}
	return big.NewInt(header.BlockHeader.RawData.Number), nil
}

func (l *Listener) StoreBlock(blockHeight *big.Int) error {
	return model.SetBlockHeight(db.DB, l.cfg.ChainId, l.cfg.From, decimal.NewFromBigInt(blockHeight, 0))
}
