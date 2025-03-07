package blockchain

import (
	erc20 "coinstore/pkg/binding/token"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
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

func (c Token) TokenInit() {

}
