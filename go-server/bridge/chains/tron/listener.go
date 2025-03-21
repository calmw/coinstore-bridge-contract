package tron

import (
	"bytes"
	"coinstore/bridge/chains"
	"coinstore/bridge/config"
	"coinstore/bridge/core"
	"coinstore/bridge/event"
	"coinstore/db"
	"coinstore/model"
	"coinstore/utils"
	"encoding/json"
	"errors"
	"fmt"
	log "github.com/calmw/clog"
	eth "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/shopspring/decimal"
	"io"
	"math/big"
	"net/http"
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
	TestA(latestBlock.Int64())
	//GetTronLog(l.cfg.BridgeContractAddress, latestBlock.Text(16), latestBlock.Text(16), "f8922d8955cfa0d76336adc31b6c0ba9255e8baf479e4ef06db6cabb8711806a")
	//query := buildQuery(common.HexToAddress(l.cfg.BridgeContractAddress), event.Deposit, latestBlock, latestBlock)
	//
	//// 获取日志
	//logs, err := l.conn.Client().FilterLogs(context.Background(), query)
	//if err != nil {
	//	return fmt.Errorf("unable to Filter Logs: %w", err)
	//}

	//for _, log := range logs {
	//	var m msg.Message
	//destId := msg.ChainId(log.Topics[1].Big().Uint64())
	//rId := msg.ResourceIdFromSlice(log.Topics[2].Bytes())
	//nonce := msg.Nonce(log.Topics[3].Big().Uint64())

	//records, err := l.bridgeContract.DepositRecords(nil, log.Topics[1].Big(), log.Topics[3].Big())
	//if err != nil {
	//	return err
	//}
	//
	//l.log.Debug("get events:")
	//l.log.Debug("ResourceID", records.ResourceID)
	//l.log.Debug("DestinationChainId", records.DestinationChainId)
	//l.log.Debug("Sender", records.Sender)
	//l.log.Debug("Data", records.Data)
	//
	//m = msg.NewGenericTransfer(
	//	msg.ChainId(l.cfg.ChainId),
	//	destId,
	//	nonce,
	//	rId,
	//	records.Data[:],
	//)

	//err = l.Router.Send(m)
	//if err != nil {
	//	l.log.Error("subscription error: failed to route message", "err", err)
	//}

	// 保存到数据库
	//model.SaveBridgeOrder(m, l.log)
	//}

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

type JSONRPCRequest struct {
	JSONRPC string      `json:"jsonrpc"`
	Method  string      `json:"method"`
	Params  []Param     `json:"params,omitempty"`
	ID      interface{} `json:"id"`
}

type Param struct {
	Address   []string `json:"address"`
	FromBlock string   `json:"fromBlock"`
	ToBlock   string   `json:"toBlock"`
	Topics    []string `json:"topics"`
}

func GetTronLog(bridgeAddress, fromBlock, toBlock, topic string) {
	param := Param{}
	param.FromBlock = "0x" + fromBlock
	param.ToBlock = "0x" + toBlock
	param.Topics = []string{"0x" + topic}
	param.Address = []string{bridgeAddress}
	fmt.Println(param, "!!!!!")
	//uri := "https://api.shasta.trongrid.io/jsonrpc"
	//uri := "https://event.nileex.io/jsonrpc"
	//uri := "http://47.252.19.181:8090/jsonrpc"
	//uri := "https://nile.trongrid.io/"
	//uri := "https://nile.trongrid.io/jsonrpc/"
	//uri := "https://api.nileex.io/jsonrpc"
	uri := "https://api.shasta.trongrid.io/jsonrpc"
	// 构造请求体
	request := []JSONRPCRequest{{
		JSONRPC: "2.0",
		Method:  "eth_getLogs",
		Params:  []Param{param},
		ID:      utils.RandInt(10, 1000),
	},
	}

	// 将请求结构体编码为JSON字节数组
	requestBytes, err := json.Marshal(request)
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		return
	}
	fmt.Println("~~~~", string(requestBytes))

	// 创建HTTP请求
	req, err := http.NewRequest("POST", uri, bytes.NewBuffer(requestBytes))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}
	req.Header.Set("Content-Type", "application/json")

	// 发送请求并获取响应
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}
	defer resp.Body.Close()

	// 读取响应体（根据需要处理响应）
	responseBody, _ := io.ReadAll(resp.Body)
	fmt.Println("Response:", string(responseBody))

}

func TestA(number int64) {
	url := fmt.Sprintf("https://nile.trongrid.io/v1/blocks/%d/events", number)
	//url := fmt.Sprintf("https://api.shasta.trongrid.io/v1/blocks/%d/events", number)

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("accept", "application/json")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

	var eventLog EventLog
	err := json.Unmarshal(body, &eventLog)
	fmt.Println(err, 66666)
	if err != nil {
		return
	}
	for _, d := range eventLog.Data {
		if d.EventName == "Deposit" {
			fmt.Println(d.BlockNumber, d.EventIndex)
			fmt.Println(d.Result.ResourceID)
			fmt.Println(d.Result.DepositNonce)
			fmt.Println(d.Result.DestinationChainId)
			fmt.Println(d.TransactionID)
		}
	}

	fmt.Println(string(body))
}
