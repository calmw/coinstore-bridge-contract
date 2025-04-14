package ethereum

import (
	"coinstore/binding"
	"coinstore/bridge/chains"
	"coinstore/bridge/config"
	"coinstore/bridge/core"
	"coinstore/bridge/event"
	"coinstore/bridge/msg"
	"coinstore/bridge/tron"
	"coinstore/db"
	"coinstore/model"
	"coinstore/utils"
	"context"
	"errors"
	"fmt"
	log "github.com/calmw/clog"
	eth "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/fbsobreira/gotron-sdk/pkg/address"
	"github.com/shopspring/decimal"
	"math/big"
	"strings"
	"time"
)

var BlockRetryInterval = time.Second * 5
var BlockRetryLimit = 5
var ErrFatalPolling = errors.New("bridge block polling failed")
var Listeners = map[int]*Listener{}

type Listener struct {
	cfg            config.Config
	conn           *Connection
	Router         chains.Router
	BridgeContract *binding.Bridge
	log            log.Logger
	latestBlock    core.LatestBlock
	stop           <-chan int
	sysErr         chan<- error
}

// NewListener creates and returns a Listener
func NewListener(conn *Connection, cfg config.Config, log log.Logger, stop <-chan int, sysErr chan<- error) *Listener {
	bridgeContract, err := binding.NewBridge(common.HexToAddress(cfg.BridgeContractAddress), conn.ClientEvm())
	if err != nil {
		panic("new bridge contract failed")
	}
	listener := Listener{
		cfg:            cfg,
		conn:           conn,
		BridgeContract: bridgeContract,
		log:            log,
		stop:           stop,
		sysErr:         sysErr,
	}

	Listeners[cfg.ChainId] = &listener
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
	query := buildQuery(common.HexToAddress(l.cfg.BridgeContractAddress), event.Deposit, latestBlock, latestBlock)

	// 获取日志
	logs, err := l.conn.ClientEvm().FilterLogs(context.Background(), query)
	if err != nil {
		return fmt.Errorf("unable to Filter Logs: %w", err)
	}

	for _, logE := range logs {
		var m msg.Message
		destId := msg.ChainId(logE.Topics[1].Big().Uint64())
		rId := msg.ResourceIdFromSlice(logE.Topics[2].Bytes())
		nonce := msg.Nonce(logE.Topics[3].Big().Uint64())

		record, err := l.BridgeContract.DepositRecords(nil, logE.Topics[1].Big(), logE.Topics[3].Big())
		if err != nil {
			return err
		}

		m = msg.NewGenericTransfer(
			msg.ChainId(l.cfg.ChainId),
			destId,
			nonce,
			rId,
			record.Data[:],
		)

		// 获取目标链的信息
		var toAddr string
		destChainType := core.ChainType[int(record.DestinationChainId.Int64())]
		if destChainType == config.ChainTypeEvm {
			dl, ok := Listeners[int(record.DestinationChainId.Int64())]
			if !ok {
				l.log.Error("destination listener not found", "chainId", record.DestinationChainId)
				return errors.New(fmt.Sprintf("destination listener not found, chainId %d", record.DestinationChainId))
			}
			token, err := dl.BridgeContract.ResourceIdToTokenInfo(nil, record.ResourceID)
			if err != nil {
				l.log.Error("destination token info not found", "chainId", record.DestinationChainId)
			}
			toAddr = strings.ToLower(token.TokenAddress.String())
		} else if destChainType == config.ChainTypeTron {
			tCfg := config.TronCfg
			tokenInfo, err := tron.ResourceIdToTokenInfo(binding.OwnerAccount, tCfg.BridgeContractAddress, record.ResourceID)
			if err != nil {
				return err
			}
			tokenAddress := "0x41" + strings.TrimPrefix(tokenInfo.TokenAddress.String(), "0x")
			toAddr = address.HexToAddress(tokenAddress).String()
		}

		tokenInfo, err := l.BridgeContract.ResourceIdToTokenInfo(nil, record.ResourceID)
		fmt.Println("000 ", l.cfg.BridgeContractAddress)
		fmt.Println(fmt.Sprintf("0x%x", record.ResourceID), " ", record.DestinationChainId)
		fmt.Println(tokenInfo)
		fmt.Println(err, "666666666666666666666666666")
		if err != nil {
			l.log.Error("source token info not found", "chainId", record.DestinationChainId)
		}
		amount, caller, receiver, err := utils.ParseBridgeData(record.Data)
		if destChainType == config.ChainTypeTron {
			receiver, _ = utils.EthToTron(receiver)
		}
		fee := decimal.NewFromBigInt(tokenInfo.Fee, 0)
		// 保存到数据库
		model.SaveBridgeOrder(l.log, m, amount, fmt.Sprintf("%x", record.ResourceID), caller, receiver, strings.ToLower(tokenInfo.TokenAddress.String()), toAddr, logE.TxHash.String(), time.Unix(record.Ctime.Int64(), 0).Format("2006-01-02 15:04:05"), fee)

		err = l.Router.Send(m)
		if err != nil {
			l.log.Error("subscription error: failed to route message", "err", err)
			return errors.New(fmt.Sprintf("subscription error: failed to route message,err %v", err))
		}
	}

	return nil
}

func (l *Listener) LatestBlock() (*big.Int, error) {
	header, err := l.conn.ClientEvm().HeaderByNumber(context.Background(), nil)
	if err != nil {
		return nil, err
	}
	l.latestBlock = core.LatestBlock{
		Height:      header.Number,
		LastUpdated: time.Unix(int64(header.Time), 0),
	}
	return header.Number, nil
}

func (l *Listener) StoreBlock(blockHeight *big.Int) error {
	return model.SetBlockHeight(db.DB, l.cfg.ChainId, l.cfg.From, decimal.NewFromBigInt(blockHeight, 0))
}

func buildQuery(contract common.Address, sig event.Sig, startBlock *big.Int, endBlock *big.Int) eth.FilterQuery {
	query := eth.FilterQuery{
		FromBlock: startBlock,
		ToBlock:   endBlock,
		Addresses: []common.Address{contract},
		Topics: [][]common.Hash{
			{sig.GetTopic()},
		},
	}
	return query
}
