package ethereum

import (
	"coinstore/binding"
	"coinstore/bridge/config"
	"coinstore/bridge/connections"
	"coinstore/bridge/core"
	"coinstore/bridge/msg"
	"crypto/ecdsa"
	"fmt"
	log "github.com/calmw/blog"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
)

var _ core.Chain = &Chain{}

var _ Connection = &connections.Connection{}

type Connection interface {
	Connect() error
	KeyPrv() *ecdsa.PrivateKey
	Opts() *bind.TransactOpts
	CallOpts() *bind.CallOpts
	LockAndUpdateOpts() error
	UnlockOpts()
	Client() *ethclient.Client
	EnsureHasBytecode(address common.Address) error
	LatestBlock() (*big.Int, error)
	WaitForBlock(block *big.Int, delay *big.Int) error
	Close()
}

type Chain struct {
	cfg      *config.Config // The config of the chain
	conn     Connection     // THe chains connection
	listener *Listener      // The listener of this chain
	writer   *Writer        // The writer of the chain
	stop     chan<- int
}

func InitializeChain(cfg *config.Config, logger log.Logger, sysErr chan<- error) (*Chain, error) {
	stop := make(chan int)
	conn := connections.NewConnection(cfg.Endpoint, cfg.Http, cfg.PrivateKey, logger, cfg.GasLimit, cfg.MaxGasPrice, cfg.MinGasPrice)
	err := conn.Connect()
	if err != nil {
		return nil, err
	}

	bridgeContract, err := binding.NewBridge(cfg.BridgeContractAddress, conn.Client())
	fmt.Println(cfg.BridgeContractAddress, err, cfg.ChainId, "~~~~~~~")
	if err != nil {
		return nil, err
	}

	chainId, err := bridgeContract.GetChainId(conn.CallOpts())
	if err != nil {
		return nil, err
	}

	if chainId.Int64() != int64(cfg.ChainId) {
		return nil, fmt.Errorf("chainId (%d) and configuration chainId (%d) do not match", chainId, cfg.ChainId)
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
	//writer.setContract(bridgeContract)

	//dislock.LockClientId = chainCfg.From // 设置锁的clientID

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
