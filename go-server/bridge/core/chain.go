package core

import (
	"coinstore/bridge/config"
	"coinstore/bridge/msg"
	"math/big"
	"time"
)

type Chain interface {
	Start() error // Start chain
	SetRouter(*Router)
	Id() msg.ChainId
	ChainType() config.ChainType
	Name() string
	LatestBlock() LatestBlock
	Stop()
}

type LatestBlock struct {
	Height      *big.Int
	LastUpdated time.Time
}

type ChainConfig struct {
	Name           string            // Human-readable chain name
	Id             msg.ChainId       // ChainID
	Endpoint       string            // url for rpc endpoint
	From           string            // address of key to use
	KeystorePath   string            // Location of key files
	Insecure       bool              // Indicated whether the test keyring should be used
	BlockstorePath string            // Location of blockstore
	FreshStart     bool              // If true, blockstore is ignored at start.
	LatestBlock    bool              // If true, overrides blockstore or latest block in config and starts from current block
	Opts           map[string]string // Per chain options
}

var ChainType = map[int]config.ChainType{}
