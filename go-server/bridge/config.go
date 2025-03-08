// Copyright 2020 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package bridge

import (
	"coinstore/db"
	"coinstore/model"
	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum/crypto"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/common"
)

const DefaultGasLimit = 6721975
const DefaultGasPrice = 20000000000
const DefaultMinGasPrice = 0
const DefaultBlockConfirmations = 5

// Config encapsulates all necessary parameters in ethereum compatible forms
type Config struct {
	chainName          string
	chainId            int
	endpoint           string
	from               string
	privateKey         *ecdsa.PrivateKey
	bridgeContract     common.Address
	voteContract       common.Address
	gasLimit           *big.Int
	maxGasPrice        *big.Int
	minGasPrice        *big.Int
	http               bool
	startBlock         *big.Int
	blockConfirmations *big.Int
}

func NewConfig(cfg model.Config) Config {
	key := os.Getenv("COINSTORE_BRIDGE")
	//key:=utils2.ThreeDesDecrypt("",mCfg.PrivateKey)
	privateKey, err := crypto.HexToECDSA(key)
	if err != nil {
		panic("private key conversion failed")
	}
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
	if cfg.MaxGasPrice > 0 {
		blockConfirmations = big.NewInt(cfg.BlockConfirmations)
	}
	startBlock := cfg.StartBlock
	if cfg.FreshStart == 0 {
		height, err := model.GetBlockHeight(db.DB, cfg.ChainId)
		if err == nil {
			startBlock = *height
		}
	}
	http := false
	if cfg.Http > 0 {
		http = true
	}

	return Config{
		chainName:          cfg.ChainName,
		chainId:            cfg.ChainId,
		endpoint:           cfg.Endpoint,
		from:               cfg.From,
		privateKey:         privateKey,
		bridgeContract:     common.Address{},
		voteContract:       common.Address{},
		gasLimit:           gasLimit,
		maxGasPrice:        maxGasPrice,
		minGasPrice:        minGasPrice,
		http:               http,
		startBlock:         &startBlock,
		blockConfirmations: blockConfirmations,
	}
}
