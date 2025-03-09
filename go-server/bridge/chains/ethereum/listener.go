package ethereum

import (
	"coinstore/binding"
	"coinstore/bridge/chains"
	"coinstore/bridge/config"
	"coinstore/bridge/core"
	"coinstore/bridge/event"
	"coinstore/bridge/msg"
	"coinstore/db"
	"coinstore/model"
	"context"
	"errors"
	"fmt"
	"github.com/calmw/blog"
	eth "github.com/ethereum/go-ethereum"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/shopspring/decimal"
	"math/big"
	"time"
)

var BlockRetryInterval = time.Second * 5
var BlockRetryLimit = 5
var ErrFatalPolling = errors.New("bridge block polling failed")
var Listeners = map[int]*Listener{}

type Listener struct {
	cfg            config.Config
	conn           Connection
	Router         chains.Router
	bridgeContract *binding.Bridge
	log            log15.Logger
	latestBlock    core.LatestBlock
	stop           <-chan int
	sysErr         chan<- error
}

// NewListener creates and returns a Listener
func NewListener(conn Connection, cfg *config.Config, log log15.Logger, stop <-chan int, sysErr chan<- error) *Listener {
	bridgeContract, err := binding.NewBridge(cfg.BridgeContractAddress, conn.Client())
	if err != nil {
		panic("new bridge contract failed")
	}
	listener := Listener{
		cfg:            *cfg,
		conn:           conn,
		bridgeContract: bridgeContract,
		log:            log,
		stop:           stop,
		sysErr:         sysErr,
	}

	Listeners[cfg.ChainId] = &listener
	log.Debug("new listener id", "id", cfg.ChainId)
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

			// Sleep if the difference is less than BlockDelay; (latest - current) < BlockDelay
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
	query := buildQuery(l.cfg.BridgeContractAddress, event.Deposit, latestBlock, latestBlock)

	// 获取日志
	logs, err := l.conn.Client().FilterLogs(context.Background(), query)
	if err != nil {
		return fmt.Errorf("unable to Filter Logs: %w", err)
	}

	for _, log := range logs {
		var m msg.Message
		destId := msg.ChainId(log.Topics[1].Big().Uint64())
		rId := msg.ResourceIdFromSlice(log.Topics[2].Bytes())
		nonce := msg.Nonce(log.Topics[3].Big().Uint64())

		records, err := l.bridgeContract.DepositRecords(nil, log.Topics[1].Big(), log.Topics[3].Big())
		if err != nil {
			return err
		}

		l.log.Debug("get events:")
		l.log.Debug("ResourceID", records.ResourceID)
		l.log.Debug("DestinationChainId", records.DestinationChainId)
		l.log.Debug("Sender", records.Sender)
		l.log.Debug("Data", records.Data)

		m = msg.NewGenericTransfer(
			msg.ChainId(l.cfg.ChainId),
			destId,
			nonce,
			rId,
			records.Data[:],
		)

		err = l.Router.Send(m)
		if err != nil {
			l.log.Error("subscription error: failed to route message", "err", err)
		}

		// 保存到数据库
		model.SaveBridgeOrder(m, l.log)
	}

	return nil
}

func (l *Listener) LatestBlock() (*big.Int, error) {
	header, err := l.conn.Client().HeaderByNumber(context.Background(), nil)
	if err != nil {
		return nil, err
	}
	return header.Number, nil
}

func (l *Listener) StoreBlock(blockHeight *big.Int) error {
	return model.SetBlockHeight(db.DB, l.cfg.ChainId, decimal.NewFromBigInt(blockHeight, 0))
}

func buildQuery(contract ethcommon.Address, sig event.Sig, startBlock *big.Int, endBlock *big.Int) eth.FilterQuery {
	query := eth.FilterQuery{
		FromBlock: startBlock,
		ToBlock:   endBlock,
		Addresses: []ethcommon.Address{contract},
		Topics: [][]ethcommon.Hash{
			{sig.GetTopic()},
		},
	}
	return query
}
