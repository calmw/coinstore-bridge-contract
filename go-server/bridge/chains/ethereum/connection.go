package ethereum

import (
	"coinstore/bridge/chains/ethereum/egs"
	"coinstore/bridge/config"
	"coinstore/contract"
	"context"
	"crypto/ecdsa"
	"errors"
	"fmt"
	log "github.com/calmw/clog"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	ethcrypto "github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/fbsobreira/gotron-sdk/pkg/client"
	"github.com/fbsobreira/gotron-sdk/pkg/keystore"
	"github.com/fbsobreira/gotron-sdk/pkg/store"
	"google.golang.org/grpc"
	"math/big"
	"strings"
	"sync"
	"time"
)

type Connection struct {
	chainType     int
	endpoint      string
	http          bool
	prvKey        *ecdsa.PrivateKey
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
	// tron
	keyStore   *keystore.KeyStore
	keyAccount *keystore.Account
	connTron   *client.GrpcClient
}

func NewConnection(chainType int, endpoint string, http bool, prvKey string, log log.Logger, gasLimit, maxGasPrice, minGasPrice *big.Int) *Connection {
	if chainType == config.ChainTypeEvm {
		//key:=utils2.ThreeDesDecrypt("",cfg.PrivateKey) // TODO 线上要改
		privateKey, err := crypto.HexToECDSA(prvKey)
		if err != nil {
			panic("private key conversion failed")
		}
		return &Connection{
			chainType:   chainType,
			endpoint:    endpoint,
			http:        http,
			prvKey:      privateKey,
			gasLimit:    gasLimit,
			maxGasPrice: maxGasPrice,
			minGasPrice: minGasPrice,
			optsLock:    &sync.Mutex{},
			log:         log,
		}
	} else if chainType == config.ChainTypeTron {
		_, _, err := contract.GetKeyFromPrivateKey(prvKey, contract.AccountName, contract.Passphrase)
		if err != nil && !strings.Contains(err.Error(), "already exists") {
			panic("private key conversion failed")
		}
		ks, ka, err := store.UnlockedKeystore(contract.AccountName, contract.Passphrase)
		return &Connection{
			chainType:   chainType,
			endpoint:    endpoint,
			http:        http,
			prvKey:      nil,
			gasLimit:    gasLimit,
			maxGasPrice: maxGasPrice,
			minGasPrice: minGasPrice,
			log:         log,
			stop:        make(chan int),
			keyStore:    ks,
			keyAccount:  ka,
		}
	} else {
		panic("Unsupported chain type")
	}
}

// Connect starts the ethereum WS connection
func (c *Connection) Connect() error {
	c.log.Info("Connecting to ethereum chain...", "rpc", c.endpoint)
	if c.chainType == config.ChainTypeEvm {
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

		opts, _, err := c.newTransactOpts(big.NewInt(0), c.gasLimit, c.maxGasPrice)
		if err != nil {
			return err
		}
		c.opts = opts
		c.nonce = 0
		c.callOpts = &bind.CallOpts{From: crypto.PubkeyToAddress(c.prvKey.PublicKey)}
	} else if c.chainType == config.ChainTypeTron {
		var err error
		cli := client.NewGrpcClient(c.endpoint)
		err = cli.Start(grpc.WithInsecure())
		if err != nil {
			return err
		}
		c.connTron = cli
	}
	return nil
}

// newTransactOpts builds the TransactOpts for the connection's keypair.
func (c *Connection) newTransactOpts(value, gasLimit, gasPrice *big.Int) (*bind.TransactOpts, uint64, error) {
	privateKey := c.prvKey
	address := ethcrypto.PubkeyToAddress(privateKey.PublicKey)

	nonce, err := c.connEvm.PendingNonceAt(context.Background(), address)
	if err != nil {
		return nil, 0, err
	}

	id, err := c.connEvm.ChainID(context.Background())
	if err != nil {
		return nil, 0, err
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, id)
	if err != nil {
		return nil, 0, err
	}
	//gasPrice, err := c.conn.SuggestGasPrice(context.Background())
	//if err != nil {
	//	return nil, 0, err
	//}
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = value
	auth.GasLimit = uint64(gasLimit.Int64())
	//auth.GasLimit = 0
	auth.GasPrice = gasPrice
	//auth.GasPrice = nil
	auth.Context = context.Background()

	return auth, nonce, nil
}

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

func (c *Connection) KeyPrv() *ecdsa.PrivateKey {
	return c.prvKey
}

func (c *Connection) SafeEstimateGas(ctx context.Context) (*big.Int, error) {
	var suggestedGasPrice *big.Int
	if c.egsApiKey != "" {
		price, err := egs.FetchGasPrice(c.egsApiKey, c.egsSpeed)
		if err != nil {
			c.log.Error("Couldn't fetch gasPrice from GSN", "err", err)
		} else {
			suggestedGasPrice = price
		}
	}
	if suggestedGasPrice == nil {
		c.log.Debug("Fetching gasPrice from node")
		nodePriceEstimate, err := c.connEvm.SuggestGasPrice(context.TODO())
		if err != nil {
			return nil, err
		} else {
			suggestedGasPrice = nodePriceEstimate
		}
	}
	gasPrice := multiplyGasPrice(suggestedGasPrice, c.gasMultiplier)
	if gasPrice.Cmp(c.minGasPrice) == -1 {
		return c.minGasPrice, nil
	} else if gasPrice.Cmp(c.maxGasPrice) == 1 {
		return c.maxGasPrice, nil
	} else {
		return gasPrice, nil
	}
}

func (c *Connection) EstimateGasLondon(ctx context.Context, baseFee *big.Int) (*big.Int, *big.Int, error) {
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

func multiplyGasPrice(gasEstimate *big.Int, gasMultiplier *big.Float) *big.Int {

	gasEstimateFloat := new(big.Float).SetInt(gasEstimate)

	result := gasEstimateFloat.Mul(gasEstimateFloat, gasMultiplier)

	gasPrice := new(big.Int)

	result.Int(gasPrice)

	return gasPrice
}

func (c *Connection) LockAndUpdateOpts() error {
	c.optsLock.Lock()

	head, err := c.connEvm.HeaderByNumber(context.TODO(), nil)
	if err != nil {
		c.UnlockOpts()
		return err
	}

	if head.BaseFee != nil {
		c.opts.GasTipCap, c.opts.GasFeeCap, err = c.EstimateGasLondon(context.TODO(), head.BaseFee)
		if err != nil {
			c.UnlockOpts()
			return err
		}

		// Both gasPrice and (maxFeePerGas or maxPriorityFeePerGas) cannot be specified: https://github.com/ethereum/go-ethereum/blob/95bbd46eabc5d95d9fb2108ec232dd62df2f44ab/accounts/abi/bind/base.go#L254
		c.opts.GasPrice = nil
	} else {
		var gasPrice *big.Int
		gasPrice, err = c.SafeEstimateGas(context.TODO())
		if err != nil {
			c.UnlockOpts()
			return err
		}
		c.opts.GasPrice = gasPrice
	}

	nonce, err := c.connEvm.PendingNonceAt(context.Background(), c.opts.From)
	if err != nil {
		c.optsLock.Unlock()
		return err
	}
	c.opts.Nonce.SetUint64(nonce)
	return nil
}

func (c *Connection) UnlockOpts() {
	c.optsLock.Unlock()
}

func (c *Connection) LatestBlock() (*big.Int, error) {

	if c.chainType == config.ChainTypeEvm {
		header, err := c.connEvm.HeaderByNumber(context.Background(), nil)
		if err != nil {
			return nil, err
		}
		return header.Number, nil
	} else if c.chainType == config.ChainTypeTron {
		block, err := c.connTron.GetNowBlock()
		if err != nil {
			return nil, fmt.Errorf("get block now: %v", err)
		}
		fmt.Println(block.String())
		return big.NewInt(5), nil
	}
	return nil, errors.New("unexpected")

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
