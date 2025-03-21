package connections

import (
	"coinstore/contract"
	"errors"
	"fmt"
	log "github.com/calmw/blog"
	"github.com/fbsobreira/gotron-sdk/pkg/client"
	"github.com/fbsobreira/gotron-sdk/pkg/keystore"
	"github.com/fbsobreira/gotron-sdk/pkg/store"
	"google.golang.org/grpc"
	"math/big"
	"strings"
	"time"
)

type ConnectionTron struct {
	endpoint   string
	prvKey     string
	gasLimit   *big.Int // fee limit
	keyStore   *keystore.KeyStore
	keyAccount *keystore.Account
	conn       *client.GrpcClient
	log        log.Logger
	stop       chan int
}

func NewConnectionTron(endpoint string, prvKey string, log log.Logger, gasLimit *big.Int) (*ConnectionTron, error) {
	_, _, err := contract.GetKeyFromPrivateKey(prvKey, contract.AccountName, contract.Passphrase)
	if err != nil && !strings.Contains(err.Error(), "already exists") {
		return nil, err
	}
	ks, ka, err := store.UnlockedKeystore(contract.AccountName, contract.Passphrase)
	return &ConnectionTron{
		endpoint:   endpoint,
		prvKey:     prvKey,
		gasLimit:   gasLimit,
		keyStore:   ks,
		keyAccount: ka,
		log:        log,
		stop:       make(chan int),
	}, nil
}

// Connect starts the ethereum WS connection
func (c *ConnectionTron) Connect() error {
	c.log.Info("Connecting to tron chain...", "rpc", c.endpoint)
	var err error
	cli := client.NewGrpcClient(c.endpoint)
	err = cli.Start(grpc.WithInsecure())
	if err != nil {
		return err
	}
	c.conn = cli
	return nil
}

func (c *ConnectionTron) Client() *client.GrpcClient {
	return c.conn
}

func (c *ConnectionTron) KeyStore() *keystore.KeyStore {
	return c.keyStore
}

func (c *ConnectionTron) KeyAccount() *keystore.Account {
	return c.KeyAccount()
}

// LatestBlock returns the latest block from the current chain
func (c *ConnectionTron) LatestBlock() (*big.Int, error) {

	block, err := c.conn.GetNowBlock()
	if err != nil {
		return nil, fmt.Errorf("Get block now: %v", err)
	}
	fmt.Println(block.String())
	return big.NewInt(5), nil
}

// WaitForBlock will poll for the block number until the current block is equal or greater.
// If delay is provided it will wait until currBlock - delay = targetBlock
func (c *ConnectionTron) WaitForBlock(targetBlock *big.Int, delay *big.Int) error {
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

			// Equal or greater than target
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
func (c *ConnectionTron) Close() {
	if c.conn != nil {
		c.conn.Stop()
	}
	close(c.stop)
}
