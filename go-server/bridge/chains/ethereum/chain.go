package ethereum

import (
	"coinstore/binding"
	"coinstore/bridge/chains"
	"coinstore/bridge/config"
	"coinstore/bridge/core"
	"coinstore/bridge/msg"
	"fmt"
	log "github.com/calmw/clog"
	"github.com/ethereum/go-ethereum/common"
	"os"
)

var _ core.Chain = &Chain{}

type Chain struct {
	cfg      *config.Config    // The config of the chain
	conn     chains.Connection // THe chains connection
	listener *Listener         // The listener of this chain
	writer   *Writer           // The writer of the chain
	stop     chan<- int
}

func InitializeChain(cfg *config.Config, logger log.Logger, sysErr chan<- error) (*Chain, error) {
	stop := make(chan int)
	key := os.Getenv("COINSTORE_BRIDGE")
	//key:=utils2.ThreeDesDecrypt("",cfg.PrivateKey) // TODO 线上要改
	conn := NewConnection(cfg.ChainType, cfg.Endpoint, cfg.Http, key, logger, cfg.GasLimit, cfg.MaxGasPrice, cfg.MinGasPrice)
	err := conn.Connect()
	if err != nil {
		logger.Error("new connection", "error", err)
		return nil, err
	}

	bridgeContract, err := binding.NewBridge(common.HexToAddress(cfg.BridgeContractAddress), conn.ClientEvm())
	if err != nil {
		return nil, err
	}

	chainId, err := bridgeContract.ChainId(conn.CallOpts())
	if err != nil {
		return nil, err
	}

	if chainId.Int64() != int64(cfg.ChainId) {
		return nil, fmt.Errorf("chainId (%d) and configuration chainId (%d) do not match", chainId.Int64(), cfg.ChainId)
	}

	chainTypeId, err := bridgeContract.ChainType(conn.CallOpts())
	if err != nil {
		return nil, err
	}

	if chainTypeId.Int64() != int64(cfg.ChainType) {
		return nil, fmt.Errorf("chainTypeId (%d) and configuration chainTypeId (%d) do not match", chainTypeId.Int64(), cfg.ChainType)
	}

	if cfg.LatestBlock {
		curr, err := conn.LatestBlock()
		if err != nil {
			return nil, err
		}
		cfg.StartBlock = curr
	}

	listener := NewListener(conn, cfg, logger, stop, sysErr)

	// TODO: writer
	writer := NewWriter(conn, cfg, logger, stop, sysErr)

	return &Chain{
		cfg:      cfg,
		conn:     conn,
		writer:   writer,
		listener: listener,
		stop:     stop,
	}, nil
}

func (c *Chain) SetRouter(r *core.Router) {
	r.Listen(msg.ChainId(c.cfg.ChainId), c.writer)
	c.listener.setRouter(r)
}

func (c *Chain) Start() error {
	err := c.listener.start()
	if err != nil {
		return err
	}

	err = c.writer.start()
	if err != nil {
		return err
	}

	c.writer.log.Debug("Successfully started chain")
	return nil
}

func (c *Chain) Id() msg.ChainId {
	return msg.ChainId(c.cfg.ChainId)
}

func (c *Chain) Name() string {
	return c.cfg.ChainName
}

func (c *Chain) ChainType() config.ChainType {
	return c.cfg.ChainType
}

func (c *Chain) LatestBlock() core.LatestBlock {
	return c.listener.latestBlock
}

// Stop signals to any running routines to exit
func (c *Chain) Stop() {
	close(c.stop)
	if c.conn != nil {
		c.conn.Close()
	}
}
