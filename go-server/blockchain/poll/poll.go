package poll

import (
	"context"
	log "github.com/ChainSafe/log15"
	"github.com/ethereum/go-ethereum/ethclient"
	"time"
)

const (
	ChainId            = 1         // 自定义链ID
	ChainType          = 1         // 链类型 1 evm 2 tron
	ChainName          = "OpenBNB" // 链类型 1 evm 2 tron
	StartBlock         = 54443202  // 开始块高
	BlockConfirmations = 2         // 待确认块高
	Bridge             = "0x6b66eBFA87AaC1dB355B0ec49ECab7F4b32b1b30"
	Vote               = "0x77bcb682e01D8763F6757c0D0Beaf577Afcfdf43"
	Endpoint           = "https://opbnb-testnet-rpc.bnbchain.org"
	From               = ""   // relayer 账户
	Http               = true // endpoint 类型 true http; false ws
	GasLimit           = 1000000000
	MaxGasPrice        = 1000000000
)

var LatestBlockNumber uint64 = 0

var Log log.Logger

type Listener struct {
	Cli         *ethclient.Client
	BlockHeight uint64
}

func NewListener() *Listener {
	err, cli := Client(Endpoint)
	if err != nil {
		panic(err)
	}
	return &Listener{
		Cli:         cli,
		BlockHeight: StartBlock,
	}
}

func Init() {
	pLog := log.New("chain", ChainName)
	Log = pLog
}

func Client(RPC string) (error, *ethclient.Client) {
	client, err := ethclient.Dial(RPC)
	if err != nil {
		return err, nil
	}
	return nil, client
}

func (p *Listener) Run() {
	go p.GetLatestNumber()
}

func (p *Listener) GetLatestNumber() {
	for {
		BlockNumber(p.Cli)
	}
}

func (p *Listener) Poll() {

}

func BlockNumber(client *ethclient.Client) {
	var blockNumber uint64
	var err error
	for i := 0; i < 10; i++ {
		header, err := c.conn.HeaderByNumber(context.Background(), nil)
		if err != nil {
			return nil, err
		}
		if err != nil {
			time.Sleep(time.Second)
			continue
		} else {
			LatestBlockNumber = blockNumber
			Log.Debug("latest block height", "height", blockNumber)
			return
		}
	}
	panic("get block number error")
}
