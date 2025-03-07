// Copyright 2020 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package core

import (
	"coinstore/model"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

const DefaultGasLimit = 6721975
const DefaultGasPrice = 20000000000
const DefaultMinGasPrice = 0
const DefaultBlockConfirmations = 10
const DefaultGasMultiplier = 1

// Config encapsulates all necessary parameters in ethereum compatible forms
type Config struct {
	chainName          string
	chainId            int
	endpoint           string
	from               string
	privateKey         string
	freshStart         bool
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
	return Config{
		chainName:          cfg.ChainName,
		chainId:            cfg.ChainId,
		endpoint:           cfg.Endpoint,
		from:               cfg.From,
		privateKey:         cf,
		freshStart:         false,
		bridgeContract:     common.Address{},
		voteContract:       common.Address{},
		gasLimit:           nil,
		maxGasPrice:        nil,
		minGasPrice:        nil,
		http:               false,
		startBlock:         nil,
		blockConfirmations: nil,
	}
}
