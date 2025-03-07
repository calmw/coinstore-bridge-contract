package blockchain

import (
	"betcorgi/pkg/binding/corgi"
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
	Contract *corgi.Token
}

func NewToken() (*Token, error) {
	err, cli := Client(ChainConfig)
	if err != nil {
		return nil, err
	}
	contractObj, err := corgi.NewToken(common.HexToAddress(ChainConfig.TokenContractAddress), cli)
	if err != nil {
		return nil, err
	}
	return &Token{
		Cli:      cli,
		Contract: contractObj,
	}, nil
}

func (c Token) TokenInit() {
	c.AdminSetToken()
	c.AdminSetTokenExchangeRate()
}

func (c Token) AdminSetToken() {
	var tokenInfos []TokenInfo
	b6 := big.NewInt(1e6)
	// ETH
	tokenInfos = append(tokenInfos, TokenInfo{
		GameId:       1,
		TokenId:      0,
		TokenAddress: "0xfBDec254E6D1c30c8051Bf53F336d1308e7c9D3e",
		MinBetAmount: big.NewInt(1),
		MaxBetAmount: b6.Mul(b6, big.NewInt(10000)),
	})
	tokenInfos = append(tokenInfos, TokenInfo{
		GameId:       2,
		TokenId:      0,
		TokenAddress: "0xfBDec254E6D1c30c8051Bf53F336d1308e7c9D3e",
		MinBetAmount: big.NewInt(1),
		MaxBetAmount: b6.Mul(b6, big.NewInt(10000)),
	})
	tokenInfos = append(tokenInfos, TokenInfo{
		GameId:       3,
		TokenId:      0,
		TokenAddress: "0xfBDec254E6D1c30c8051Bf53F336d1308e7c9D3e",
		MinBetAmount: big.NewInt(1),
		MaxBetAmount: b6.Mul(b6, big.NewInt(10000)),
	})
	tokenInfos = append(tokenInfos, TokenInfo{
		GameId:       4,
		TokenId:      0,
		TokenAddress: "0xfBDec254E6D1c30c8051Bf53F336d1308e7c9D3e",
		MinBetAmount: big.NewInt(1),
		MaxBetAmount: b6.Mul(b6, big.NewInt(10000)),
	})
	tokenInfos = append(tokenInfos, TokenInfo{
		GameId:       9,
		TokenId:      0,
		TokenAddress: "0xfBDec254E6D1c30c8051Bf53F336d1308e7c9D3e",
		MinBetAmount: big.NewInt(1),
		MaxBetAmount: b6.Mul(b6, big.NewInt(10000)),
	})
	tokenInfos = append(tokenInfos, TokenInfo{
		GameId:       16,
		TokenId:      0,
		TokenAddress: "0xfBDec254E6D1c30c8051Bf53F336d1308e7c9D3e",
		MinBetAmount: big.NewInt(1),
		MaxBetAmount: b6.Mul(b6, big.NewInt(10000)),
	})
	// USDT
	tokenInfos = append(tokenInfos, TokenInfo{
		GameId:       1,
		TokenId:      1,
		TokenAddress: "0xfBDec254E6D1c30c8051Bf53F336d1308e7c9D3e",
		MinBetAmount: big.NewInt(1),
		MaxBetAmount: b6.Mul(b6, big.NewInt(10000)),
	})
	tokenInfos = append(tokenInfos, TokenInfo{
		GameId:       2,
		TokenId:      1,
		TokenAddress: "0xfBDec254E6D1c30c8051Bf53F336d1308e7c9D3e",
		MinBetAmount: big.NewInt(1),
		MaxBetAmount: b6.Mul(b6, big.NewInt(10000)),
	})
	tokenInfos = append(tokenInfos, TokenInfo{
		GameId:       3,
		TokenId:      1,
		TokenAddress: "0xfBDec254E6D1c30c8051Bf53F336d1308e7c9D3e",
		MinBetAmount: big.NewInt(1),
		MaxBetAmount: b6.Mul(b6, big.NewInt(10000)),
	})
	tokenInfos = append(tokenInfos, TokenInfo{
		GameId:       4,
		TokenId:      1,
		TokenAddress: "0xfBDec254E6D1c30c8051Bf53F336d1308e7c9D3e",
		MinBetAmount: big.NewInt(1),
		MaxBetAmount: b6.Mul(b6, big.NewInt(10000)),
	})
	tokenInfos = append(tokenInfos, TokenInfo{
		GameId:       9,
		TokenId:      1,
		TokenAddress: "0xfBDec254E6D1c30c8051Bf53F336d1308e7c9D3e",
		MinBetAmount: big.NewInt(1),
		MaxBetAmount: b6.Mul(b6, big.NewInt(10000)),
	})
	tokenInfos = append(tokenInfos, TokenInfo{
		GameId:       16,
		TokenId:      1,
		TokenAddress: "0xfBDec254E6D1c30c8051Bf53F336d1308e7c9D3e",
		MinBetAmount: big.NewInt(1),
		MaxBetAmount: b6.Mul(b6, big.NewInt(10000)),
	})
	/// USDC
	tokenInfos = append(tokenInfos, TokenInfo{
		GameId:       1,
		TokenId:      2,
		TokenAddress: "0x70EF97cdFdafa043925225eA4EEeE595DAF22542",
		MinBetAmount: big.NewInt(1),
		MaxBetAmount: b6.Mul(b6, big.NewInt(10000)),
	})
	tokenInfos = append(tokenInfos, TokenInfo{
		GameId:       2,
		TokenId:      2,
		TokenAddress: "0x70EF97cdFdafa043925225eA4EEeE595DAF22542",
		MinBetAmount: big.NewInt(1),
		MaxBetAmount: b6.Mul(b6, big.NewInt(10000)),
	})
	tokenInfos = append(tokenInfos, TokenInfo{
		GameId:       3,
		TokenId:      2,
		TokenAddress: "0x70EF97cdFdafa043925225eA4EEeE595DAF22542",
		MinBetAmount: big.NewInt(1),
		MaxBetAmount: b6.Mul(b6, big.NewInt(10000)),
	})
	tokenInfos = append(tokenInfos, TokenInfo{
		GameId:       4,
		TokenId:      2,
		TokenAddress: "0x70EF97cdFdafa043925225eA4EEeE595DAF22542",
		MinBetAmount: big.NewInt(1),
		MaxBetAmount: b6.Mul(b6, big.NewInt(10000)),
	})
	tokenInfos = append(tokenInfos, TokenInfo{
		GameId:       9,
		TokenId:      2,
		TokenAddress: "0x70EF97cdFdafa043925225eA4EEeE595DAF22542",
		MinBetAmount: big.NewInt(1),
		MaxBetAmount: b6.Mul(b6, big.NewInt(10000)),
	})
	tokenInfos = append(tokenInfos, TokenInfo{
		GameId:       16,
		TokenId:      2,
		TokenAddress: "0x70EF97cdFdafa043925225eA4EEeE595DAF22542",
		MinBetAmount: big.NewInt(1),
		MaxBetAmount: b6.Mul(b6, big.NewInt(10000)),
	})
	for _, tokenInfo := range tokenInfos {
		c.SetToken(tokenInfo)
	}
}

func (c Token) SetToken(tokenInfo TokenInfo) {
	var res *types.Transaction

	for {
		err, txOpts := GetAuth(c.Cli)
		if err != nil {
			log.Println(err)
			return
		}
		res, err = c.Contract.AdminSetToken(
			txOpts, big.NewInt(tokenInfo.GameId),
			big.NewInt(tokenInfo.TokenId),
			common.HexToAddress(tokenInfo.TokenAddress),
			tokenInfo.MinBetAmount,
			tokenInfo.MaxBetAmount,
		)
		if err == nil {
			break
		}
		time.Sleep(3 * time.Second)
	}

	log.Println(fmt.Sprintf("设置成功"))
	for {
		receipt, err := c.Cli.TransactionReceipt(context.Background(), res.Hash())
		if err == nil && receipt.Status == 1 {
			break
		}
		time.Sleep(time.Second * 2)
	}
	log.Println("等待确认成功")
}

func (c Token) AdminSetTokenExchangeRate() {
	var res *types.Transaction

	// ETH
	for {
		err, txOpts := GetAuth(c.Cli)
		if err != nil {
			log.Println(err)
			return
		}
		res, err = c.Contract.AdminSetTokenExchangeRate(
			txOpts,
			big.NewInt(0),
			big.NewInt(33840000),
		)
		if err == nil {
			break
		}
		time.Sleep(3 * time.Second)
	}

	log.Println(fmt.Sprintf("设置成功"))
	for {
		receipt, err := c.Cli.TransactionReceipt(context.Background(), res.Hash())
		if err == nil && receipt.Status == 1 {
			break
		}
		time.Sleep(time.Second * 2)
	}
	log.Println("等待确认成功")

	// USDT
	for {
		err, txOpts := GetAuth(c.Cli)
		if err != nil {
			log.Println(err)
			return
		}
		res, err = c.Contract.AdminSetTokenExchangeRate(
			txOpts,
			big.NewInt(1),
			big.NewInt(10000),
		)
		if err == nil {
			break
		}
		time.Sleep(3 * time.Second)
	}

	log.Println(fmt.Sprintf("设置成功"))
	for {
		receipt, err := c.Cli.TransactionReceipt(context.Background(), res.Hash())
		if err == nil && receipt.Status == 1 {
			break
		}
		time.Sleep(time.Second * 2)
	}
	log.Println("等待确认成功")

	// USDC
	for {
		err, txOpts := GetAuth(c.Cli)
		if err != nil {
			log.Println(err)
			return
		}
		res, err = c.Contract.AdminSetTokenExchangeRate(
			txOpts,
			big.NewInt(2),
			big.NewInt(10014),
		)
		if err == nil {
			break
		}
		time.Sleep(3 * time.Second)
	}

	log.Println(fmt.Sprintf("设置成功"))
	for {
		receipt, err := c.Cli.TransactionReceipt(context.Background(), res.Hash())
		if err == nil && receipt.Status == 1 {
			break
		}
		time.Sleep(time.Second * 2)
	}
	log.Println("等待确认成功")
}
