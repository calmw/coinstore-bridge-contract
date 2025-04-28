package tron

import (
	"coinstore/bridge/config"
	"context"
	"errors"
	"fmt"
	log "github.com/calmw/clog"
	"github.com/calmw/tron-sdk/pkg/client"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"google.golang.org/grpc"
	"math/big"
	"sync"
	"time"
)

type Connection struct {
	chainType     config.ChainType
	endpoint      string
	http          bool
	from          string
	gasLimit      *big.Int
	maxGasPrice   *big.Int
	minGasPrice   *big.Int
	gasMultiplier *big.Float
	egsApiKey     string
	egsSpeed      string
	connEvm       *ethclient.Client
	opts          *bind.TransactOpts
	callOpts      *bind.CallOpts
	nonce         uint64
	optsLock      *sync.Mutex
	log           log.Logger
	stop          chan int // All routines should exit when this channel is closed
	// trigger
	//keyStore   *keystore.KeyStore
	//keyAccount *keystore.Account
	connTron *client.GrpcClient
}

func NewConnection(endpoint string, http bool, from string, log log.Logger, gasLimit, maxGasPrice, minGasPrice *big.Int) *Connection {
	return &Connection{
		//chainType:   chainType,
		endpoint:    endpoint,
		http:        http,
		from:        from,
		gasLimit:    gasLimit,
		maxGasPrice: maxGasPrice,
		minGasPrice: minGasPrice,
		log:         log,
		stop:        make(chan int),
		optsLock:    &sync.Mutex{},
	}
}

// Connect starts the ethereum WS connection
func (c *Connection) Connect() error {
	c.log.Info("Connecting to ethereum chain...", "rpc", c.endpoint)
	var err error
	cli := client.NewGrpcClient(c.endpoint)
	err = cli.Start(grpc.WithInsecure())
	if err != nil {
		return err
	}
	c.connTron = cli
	return nil
}

// newTransactOpts builds the TransactOpts for the connection's keypair.

func (c *Connection) ClientEvm() *ethclient.Client {
	return c.connEvm
}

func (c *Connection) ClientTron() *client.GrpcClient {
	return c.connTron
}

func (c *Connection) Opts() *bind.TransactOpts {
	return c.opts
}

func (c *Connection) CallOpts() *bind.CallOpts {
	return c.callOpts
}

// and gas price.
func (c *Connection) LockAndUpdateOpts() error {
	return nil
}

func (c *Connection) UnlockOpts() {
	c.optsLock.Unlock()
}

// LatestBlock returns the latest block from the current chain
func (c *Connection) LatestBlock() (*big.Int, error) {
	header, err := c.ClientTron().GetNowBlock()
	if err != nil {
		return nil, err
	}
	return big.NewInt(header.BlockHeader.RawData.Number), nil
}

// EnsureHasBytecode asserts if contract code exists at the specified address
func (c *Connection) EnsureHasBytecode(addr ethcommon.Address) error {
	code, err := c.connEvm.CodeAt(context.Background(), addr, nil)
	if err != nil {
		return err
	}

	if len(code) == 0 {
		return fmt.Errorf("no bytecode found at %s", addr.Hex())
	}
	return nil
}

// WaitForBlock will poll for the block number until the current block is equal or greater.
// If delay is provided it will wait until currBlock - delay = targetBlock
func (c *Connection) WaitForBlock(targetBlock *big.Int, delay *big.Int) error {
	for {
		select {
		case <-c.stop:
			return errors.New("connection terminated")
		default:
			currBlock, err := c.LatestBlock()
			if err != nil {
				return err
			}

			if delay != nil {
				currBlock.Sub(currBlock, delay)
			}

			if currBlock.Cmp(targetBlock) >= 0 {
				return nil
			}
			c.log.Trace("Block not ready, waiting", "target", targetBlock, "current", currBlock, "delay", delay)
			time.Sleep(BlockRetryInterval)
			continue
		}
	}
}

// Close terminates the client connection and stops any running routines
func (c *Connection) Close() {

	if c.chainType == config.ChainTypeEvm {
		if c.connEvm != nil {
			c.connEvm.Close()
		}
	} else if c.chainType == config.ChainTypeTron {
		if c.connTron != nil {
			c.connTron.Stop()
		}
	}

	close(c.stop)
}
