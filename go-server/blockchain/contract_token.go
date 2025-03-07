package blockchain

import (
	"coinstore/binding/token"
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math/big"
	"time"
)

type TokenInfo struct {
	GameId       int64    `json:"game_id"`
	TokenId      int64    `json:"token_id"`
	TokenAddress string   `json:"token_address"`
	MinBetAmount *big.Int `json:"min_bet_amount"`
	MaxBetAmount *big.Int `json:"max_bet_amount"`
}

type Token struct {
	Cli      *ethclient.Client
	Contract *erc20.Erc20
}

func NewToken(addr common.Address) (*Token, error) {
	err, cli := Client(ChainConfig)
	if err != nil {
		return nil, err
	}
	contractObj, err := erc20.NewErc20(addr, cli)
	if err != nil {
		return nil, err
	}
	return &Token{
		Cli:      cli,
		Contract: contractObj,
	}, nil
}

func (c Token) Approve(amount *big.Int) {
	var res *types.Transaction
	for {
		err, txOpts := GetAuth(c.Cli)
		if err != nil {
			log.Println(err)
			return
		}
		res, err = c.Contract.Approve(
			txOpts,
			common.HexToAddress(ChainConfig.TantinContractAddress),
			amount,
		)
		if err == nil {
			break
		} else {
			log.Println(fmt.Sprintf("Approve error: %v", err))
		}
		time.Sleep(3 * time.Second)
	}
	log.Println(fmt.Sprintf("Approve 成功"))
	for {
		receipt, err := c.Cli.TransactionReceipt(context.Background(), res.Hash())
		if err == nil && receipt.Status == 1 {
			break
		}
		time.Sleep(time.Second * 2)
	}

	log.Println(fmt.Sprintf("Approve 确认成功"))
}
