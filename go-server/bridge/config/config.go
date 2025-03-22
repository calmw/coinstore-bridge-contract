package config

import (
	"coinstore/db"
	"coinstore/model"
	"github.com/shopspring/decimal"
	"math/big"
	"os"
)

const (
	ChainTypeEvm              = 1
	ChainTypeTron             = 2
	DefaultGasLimit           = 6721975
	DefaultGasPrice           = 20000000000
	DefaultMinGasPrice        = 0
	DefaultBlockConfirmations = 5
	TronApiHost               = "https://nile.trongrid.io"
)

type Config struct {
	ChainName string
	ChainId   int
	ChainType int
	Endpoint  string
	From      string
	//PrivateKey            *ecdsa.PrivateKey
	PrivateKey            string
	BridgeContractAddress string
	VoteContractAddress   string
	GasLimit              *big.Int
	MaxGasPrice           *big.Int
	MinGasPrice           *big.Int
	Http                  bool
	StartBlock            *big.Int
	BlockConfirmations    *big.Int
	FreshStart            bool // If true, blockstore is ignored at start.
	LatestBlock           bool // If true, overrides blockstore or latest block in config and starts from current block
}

func NewConfig(cfg model.Config) Config {
	key := os.Getenv("COINSTORE_BRIDGE")
	//key:=utils2.ThreeDesDecrypt("",cfg.PrivateKey) // TODO 线上要改
	//privateKey, err := crypto.HexToECDSA(key)
	//if err != nil {
	//	panic("private key conversion failed")
	//}
	gasLimit := big.NewInt(DefaultGasLimit)
	if cfg.GasLimit > 0 {
		gasLimit = big.NewInt(cfg.GasLimit)
	}
	maxGasPrice := big.NewInt(DefaultGasPrice)
	if cfg.MaxGasPrice > 0 {
		maxGasPrice = big.NewInt(cfg.MaxGasPrice)
	}
	minGasPrice := big.NewInt(DefaultMinGasPrice)
	if cfg.MaxGasPrice > 0 {
		minGasPrice = big.NewInt(cfg.MinGasPrice)
	}
	blockConfirmations := big.NewInt(DefaultBlockConfirmations)
	if cfg.BlockConfirmations > 0 {
		blockConfirmations = big.NewInt(cfg.BlockConfirmations)
	}
	startBlock := cfg.StartBlock
	if !cfg.FreshStart {
		height, err := model.GetBlockHeight(db.DB, cfg.ChainId, cfg.From)
		if err == nil {
			startBlock = decimal.NewFromBigInt(height, 0)
		}
	}
	http := false
	if cfg.Http > 0 {
		http = true
	}

	return Config{
		ChainName:             cfg.ChainName,
		ChainId:               cfg.ChainId,
		ChainType:             cfg.ChainType,
		Endpoint:              cfg.Endpoint,
		From:                  cfg.From,
		PrivateKey:            key,
		BridgeContractAddress: cfg.BridgeContract,
		VoteContractAddress:   cfg.VoteContract,
		GasLimit:              gasLimit,
		MaxGasPrice:           maxGasPrice,
		MinGasPrice:           minGasPrice,
		Http:                  http,
		StartBlock:            startBlock.BigInt(),
		BlockConfirmations:    blockConfirmations,
		FreshStart:            cfg.FreshStart,
		LatestBlock:           cfg.LatestBlock,
	}
}
