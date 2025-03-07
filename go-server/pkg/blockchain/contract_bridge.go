package blockchain

import (
	"coinstore/pkg/binding/bridge"
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

type Bridge struct {
	Cli      *ethclient.Client
	Contract *bridge.Bridge
}

func NewBridge() (*Bridge, error) {
	err, cli := Client(ChainConfig)
	if err != nil {
		return nil, err
	}
	contractObj, err := bridge.NewBridge(common.HexToAddress(ChainConfig.BridgeContractAddress), cli)
	if err != nil {
		return nil, err
	}
	return &Bridge{
		Cli:      cli,
		Contract: contractObj,
	}, nil
}

func (c Bridge) AutoBetInit() {
	c.AdminSetEnv()
}
func (c Bridge) AdminSetEnv() {
	c.Contract.AdminSetEnv()
}

func (c AutoBet) AddAccess() {
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
	for {
		err, txOpts := GetAuth(c.Cli)
		if err != nil {
			log.Println(err)
			return
		}
		res, err = c.Contract.GrantRole(txOpts, [32]byte(AdminRoleBytes), common.HexToAddress(ChainConfig.OrderContractAddress))
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

func (c AutoBet) AdminSetEnv() {
	var res *types.Transaction

	for {
		err, txOpts := GetAuth(c.Cli)
		if err != nil {
			log.Println(err)
			return
		}
		res, err = c.Contract.AdminSetEnv(txOpts, common.HexToAddress(ChainConfig.GameContractAddress))
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

func (c AutoBet) GetNewAutoId() *big.Int {
	var orderId *big.Int
	var err error

	for {
		orderId, err = c.Contract.GetNewAutoId(nil)
		if err == nil {
			log.Println("autoId:", orderId.String())
			break
		}
		log.Println("GetNewAutoId error:", err)
		time.Sleep(3 * time.Second)
	}
	return orderId
}
