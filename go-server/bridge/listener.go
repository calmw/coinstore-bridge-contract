package bridge

import (
	"coinstore/binding"
	"coinstore/db"
	"coinstore/model"
	"context"
	"errors"
	"fmt"
	utils "github.com/ChainSafe/ChainBridge/shared/ethereum"
	"github.com/ChainSafe/chainbridge-utils/msg"
	"github.com/ChainSafe/log15"
	eth "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
	"time"
)

var BlockRetryInterval = time.Second * 5
var BlockRetryLimit = 5
var ErrFatalPolling = errors.New("listener block polling failed")
var Listeners = map[int]*Listener{}

type Listener struct {
	cfg            Config
	cli            *ethclient.Client
	bridgeContract *binding.Bridge
	voterContract  *binding.Vote
	log            log15.Logger
	latestBlock    *big.Int
	stop           <-chan int
	sysErr         chan<- error
}

// NewListener creates and returns a Listener
func NewListener(mCfg model.Config, log log15.Logger) *Listener {
	cfg := NewConfig(mCfg)
	client, err := ethclient.Dial(cfg.endpoint)
	if err != nil {
		panic("rpc dail failed")
	}
	bridgeContract, err := binding.NewBridge(cfg.bridgeContract, client)
	if err != nil {
		panic("new bridge contract failed")
	}
	voteContract, err := binding.NewVote(cfg.bridgeContract, client)
	if err != nil {
		panic("new vote contract failed")
	}
	listener := Listener{
		cfg:            cfg,
		cli:            client,
		bridgeContract: bridgeContract,
		voterContract:  voteContract,
		log:            log,
		latestBlock:    cfg.startBlock,
		stop:           make(<-chan int),
		sysErr:         make(chan<- error),
	}

	Listeners[cfg.chainId] = &listener
	log.Debug("new listener id", "id", cfg.chainId)
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

func (l *Listener) pollBlocks() error {
	var currentBlock = l.cfg.startBlock
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
			if big.NewInt(0).Sub(latestBlock, currentBlock).Cmp(l.cfg.blockConfirmations) == -1 {
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

			err = l.StoreBlock(*currentBlock)
			if err != nil {
				l.log.Error("Failed to write latest block to blockstore", "block", currentBlock, "err", err)
			}

			currentBlock.Add(currentBlock, big.NewInt(1))
			retry = BlockRetryLimit
		}
	}
}

// getDepositEventsForBlock looks for the deposit event in the latest block
func (l *Listener) getDepositEventsForBlock(latestBlock *big.Int) error {
	l.log.Debug("Querying block for deposit events", "block", latestBlock)
	query := buildQuery(l.cfg.bridgeContract, utils.Deposit, latestBlock, latestBlock)

	// 获取日志
	logs, err := l.cli.FilterLogs(context.Background(), query)
	if err != nil {
		return fmt.Errorf("unable to Filter Logs: %w", err)
	}

	for _, log := range logs {
		var m msg.Message
		destId := msg.ChainId(log.Topics[1].Big().Uint64())
		rId := msg.ResourceIdFromSlice(log.Topics[2].Bytes())
		nonce := msg.Nonce(log.Topics[3].Big().Uint64())

		addr, err := l.bridgeContract.ResourceIDToHandlerAddress(&bind.CallOpts{From: l.conn.Keypair().CommonAddress()}, rId)
		if err != nil {
			return fmt.Errorf("failed to get handler from resource ID %x", rId)
		}

		if addr == l.cfg.erc20HandlerContract {
			m, err = l.handleErc20DepositedEvent(destId, nonce)
		} else if addr == l.cfg.erc721HandlerContract {
			m, err = l.handleErc721DepositedEvent(destId, nonce)
		} else if addr == l.cfg.GenericHandlerContract {
			m, err = l.handleGenericDepositedEvent(destId, nonce)
		} else {
			l.log.Error("event has unrecognized handler", "handler", addr.Hex())
			return nil
		}

		if err != nil {
			return err
		}

		err = l.Router.Send(m)
		if err != nil {
			l.log.Error("subscription error: failed to route message", "err", err)
		}
		storage.Stg.Create(m)
	}

	return nil
}

func (l *Listener) LatestBlock() (*big.Int, error) {
	header, err := l.cli.HeaderByNumber(context.Background(), nil)
	if err != nil {
		return nil, err
	}
	return header.Number, nil
}

func (l *Listener) StoreBlock(blockHeight big.Int) error {
	return model.SetBlockHeight(db.DB, l.cfg.chainId, blockHeight)
}
func (l *Listener) GetCallOpts(blockHeight big.Int) error {
	return model.SetBlockHeight(db.DB, l.cfg.chainId, blockHeight)
}

func (l *Listener) GetDepositEvent(destId msg.ChainId, nonce msg.Nonce) (msg.Message, error) {
	l.log.Info("Handling generic deposit event")

	record, err := l.bridgeContract.DepositRecords(nil, destId, nonce)
	(&bind.CallOpts{From: l.conn.Keypair().CommonAddress()}, uint64(nonce), uint8(destId))
	if err != nil {
		l.log.Error("Error Unpacking Generic Deposit Record", "err", err)
		return msg.Message{}, nil
	}

	return msg.NewGenericTransfer(
		l.cfg.id,
		destId,
		nonce,
		record.ResourceID,
		record.MetaData[:],
	), nil
}

// buildQuery constructs a query for the bridgeContract by hashing sig to get the event topic
func buildQuery(contract ethcommon.Address, sig utils.EventSig, startBlock *big.Int, endBlock *big.Int) eth.FilterQuery {
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
