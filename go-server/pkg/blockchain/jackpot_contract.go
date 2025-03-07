package blockchain

import (
	"betcorgi/pkg/binding/corgi"
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/status-im/keycard-go/hexutils"
	"log"
	"math/big"
	"time"
)

type Jackpot struct {
	Cli      *ethclient.Client
	Contract *corgi.Jackpot
}

func NewJackpot() (*Jackpot, error) {
	err, cli := Client(ChainConfig)
	if err != nil {
		return nil, err
	}
	contractObj, err := corgi.NewJackpot(common.HexToAddress(ChainConfig.JackPotContractAddress), cli)
	if err != nil {
		return nil, err
	}
	return &Jackpot{
		Cli:      cli,
		Contract: contractObj,
	}, nil
}

func (c Jackpot) JackpotInit() {
	c.AdminSetEnv(big.NewInt(1e6), big.NewInt(10e6))
	c.AddAccess()
}

func (c Jackpot) Payout(user string, amount *big.Int) {
	var res *types.Transaction

	log.Println(111)
	poolId, err := c.Contract.GetLatestPoolId(nil)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println("poolId", poolId)

	for {
		err, txOpts := GetAuth(c.Cli)
		if err != nil {
			log.Println(err)
			return
		}
		res, err = c.Contract.Payout(txOpts, common.HexToAddress(user), poolId, amount)
		if err == nil {
			break
		}
		log.Println(err, 222)
		time.Sleep(3 * time.Second)
	}
	log.Println(fmt.Sprintf("成功: %s", res.Hash().String()))
	for {
		receipt, err := c.Cli.TransactionReceipt(context.Background(), res.Hash())
		if err == nil && receipt.Status == 1 {
			break
		}
		time.Sleep(time.Second * 2)
	}

	log.Println(fmt.Sprintf("确认成功"))
}

func (c Jackpot) AddAccess() {
	AdminRole := "a49807205ce4d355092ef5a8a18f56e8913cf4a201fbe287825b095693c21775"
	AdminRoleBytes := hexutils.HexToBytes(AdminRole)

	var res *types.Transaction

	for {
		err, txOpts := GetAuth(c.Cli)
		if err != nil {
			log.Println(err)
			return
		}
		res, err = c.Contract.GrantRole(txOpts, [32]byte(AdminRoleBytes), common.HexToAddress(ChainConfig.GameContractAddress))
		if err == nil {
			break
		}
		time.Sleep(3 * time.Second)
	}
	log.Println(fmt.Sprintf("成功"))
	for {
		receipt, err := c.Cli.TransactionReceipt(context.Background(), res.Hash())
		if err == nil && receipt.Status == 1 {
			break
		}
		time.Sleep(time.Second * 2)
	}

	log.Println(fmt.Sprintf("确认成功"))
}

func (c Jackpot) AdminSetEnv(addAmount, minOrderAmount *big.Int) {
	var res *types.Transaction
	for {
		err, txOpts := GetAuth(c.Cli)
		if err != nil {
			log.Println(err)
			return
		}
		res, err = c.Contract.AdminSetEnv(txOpts, addAmount, minOrderAmount, common.HexToAddress(ChainConfig.USDTAddress), common.HexToAddress(ChainConfig.TokenContractAddress))
		if err == nil {
			break
		}
		fmt.Println(err)
		time.Sleep(3 * time.Second)
	}
	log.Println(fmt.Sprintf("成功"))
	for {
		receipt, err := c.Cli.TransactionReceipt(context.Background(), res.Hash())
		if err == nil && receipt.Status == 1 {
			break
		}
		time.Sleep(time.Second * 2)
	}

	log.Println(fmt.Sprintf("确认成功"))
}
