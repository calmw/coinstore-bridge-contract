package ethereum

import (
	"context"
	"errors"
	"fmt"
	log "github.com/calmw/clog"
	"github.com/calmw/tron-sdk/pkg/client"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"math/big"
	"os"
	"sync"
	"time"
)

type Connection struct {
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
	/// trigger
	//keyStore   *keystore.KeyStore
	//keyAccount *keystore.Account
	connTron *client.GrpcClient
}

func NewConnection(endpoint, from string, http bool, log log.Logger, gasLimit, maxGasPrice, minGasPrice *big.Int) *Connection {
	return &Connection{
		endpoint:    endpoint,
		http:        http,
		from:        from,
		gasLimit:    gasLimit,
		maxGasPrice: maxGasPrice,
		minGasPrice: minGasPrice,
		optsLock:    &sync.Mutex{},
		log:         log,
	}
}

// Connect starts the ethereum WS connection
func (c *Connection) Connect() error {
	c.log.Info("Connecting to ethereum chain...", "rpc", c.endpoint)
	var rpcClient *rpc.Client
	var err error
	// Start http or ws client
	if c.http {
		rpcClient, err = rpc.DialHTTP(c.endpoint)
	} else {
		rpcClient, err = rpc.DialContext(context.Background(), c.endpoint)
	}
	if err != nil {
		return err
	}
	c.connEvm = ethclient.NewClient(rpcClient)
	c.callOpts = &bind.CallOpts{From: ethcommon.HexToAddress(c.from)}

	return nil
}

// newTransactOpts builds the TransactOpts for the connection's keypair.
//func (c *Connection) newTransactOpts() error {
//	nonce, err := c.connEvm.PendingNonceAt(context.Background(), ethcommon.HexToAddress(c.from))
//	if err != nil {
//		return err
//	}
//	gasPrice, err := c.connEvm.SuggestGasPrice(context.Background())
//	if err != nil {
//		return err
//	}
//	gasLimit, err := c.connEvm.EstimateGas(context.Background(), ethereum.CallMsg{})
//	if err != nil {
//		return err
//	}
//	gasTipCap, err := c.connEvm.SuggestGasTipCap(context.Background())
//	if err != nil {
//		return err
//	}
//	c.opts.Nonce = big.NewInt(int64(nonce))
//	c.opts.GasPrice = gasPrice
//	c.opts.GasLimit = gasLimit
//	c.opts.GasTipCap = gasTipCap
//
//	return nil
//}

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

//func (c *Connection) SafeEstimateGas() (*big.Int, error) {
//	var suggestedGasPrice *big.Int
//	nodePriceEstimate, err := c.connEvm.SuggestGasPrice(context.TODO())
//	if err != nil {
//		return nil, err
//	} else {
//		suggestedGasPrice = nodePriceEstimate
//	}
//
//	gasPrice := multiplyGasPrice(suggestedGasPrice, c.gasMultiplier)
//	if gasPrice.Cmp(c.minGasPrice) == -1 {
//		return c.minGasPrice, nil
//	} else if gasPrice.Cmp(c.maxGasPrice) == 1 {
//		return c.maxGasPrice, nil
//	} else {
//		return gasPrice, nil
//	}
//}

func (c *Connection) EstimateGasLondon(baseFee *big.Int) (*big.Int, *big.Int, error) {
	var maxPriorityFeePerGas *big.Int
	var maxFeePerGas *big.Int

	if c.maxGasPrice.Cmp(baseFee) < 0 {
		maxPriorityFeePerGas = big.NewInt(1000000000)
		maxFeePerGas = new(big.Int).Add(c.maxGasPrice, maxPriorityFeePerGas)
		return maxPriorityFeePerGas, maxFeePerGas, nil
	}

	maxPriorityFeePerGas, err := c.connEvm.SuggestGasTipCap(context.TODO())
	if err != nil {
		return nil, nil, err
	}

	maxFeePerGas = new(big.Int).Add(
		maxPriorityFeePerGas,
		new(big.Int).Mul(baseFee, big.NewInt(2)),
	)

	if maxFeePerGas.Cmp(maxPriorityFeePerGas) < 0 {
		return nil, nil, fmt.Errorf("maxFeePerGas (%v) < maxPriorityFeePerGas (%v)", maxFeePerGas, maxPriorityFeePerGas)
	}

	// Check we aren't exceeding our limit
	if maxFeePerGas.Cmp(c.maxGasPrice) == 1 {
		maxPriorityFeePerGas.Sub(c.maxGasPrice, baseFee)
		maxFeePerGas = c.maxGasPrice
	}
	return maxPriorityFeePerGas, maxFeePerGas, nil
}

//func multiplyGasPrice(gasEstimate *big.Int, gasMultiplier *big.Float) *big.Int {
//
//	gasEstimateFloat := new(big.Float).SetInt(gasEstimate)
//
//	result := gasEstimateFloat.Mul(gasEstimateFloat, gasMultiplier)
//
//	gasPrice := new(big.Int)
//
//	result.Int(gasPrice)
//
//	return gasPrice
//}

func (c *Connection) LockAndUpdateOpts() error {
	c.optsLock.Lock()
	sigAccount := os.Getenv("SIG_ACCOUNT_EVM")
	if len(sigAccount) <= 0 {
		sigAccount = "0x1933ccd14cafe561d862e5f35d0de75322a55412"
	}
	nonce, err := c.connEvm.PendingNonceAt(context.Background(), ethcommon.HexToAddress(sigAccount))
	if err != nil {
		return err
	}
	nodePriceEstimate, err := c.connEvm.SuggestGasPrice(context.TODO())
	if err != nil {
		return err
	}
	if c.opts == nil {
		c.opts = &bind.TransactOpts{}
	}
	c.opts.Nonce = big.NewInt(int64(nonce))
	c.opts.GasPrice = nodePriceEstimate
	c.opts.GasLimit = 21000 * 10
	c.opts.GasTipCap = nil

	return nil
}

func (c *Connection) UnlockOpts() {
	c.optsLock.Unlock()
}

func (c *Connection) LatestBlock() (*big.Int, error) {
	header, err := c.connEvm.HeaderByNumber(context.Background(), nil)
	if err != nil {
		return nil, err
	}
	return header.Number, nil
}

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

func (c *Connection) Close() {

	if c.connEvm != nil {
		c.connEvm.Close()
	}

}
