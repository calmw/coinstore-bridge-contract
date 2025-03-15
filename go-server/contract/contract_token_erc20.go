package contract

import (
	"coinstore/binding"
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math/big"
	"time"
)

type Erc20 struct {
	Cli      *ethclient.Client
	Contract *binding.Erc20
}

func NewErc20(addr common.Address) (*Erc20, error) {
	err, cli := Client(ChainConfig)
	if err != nil {
		return nil, err
	}
	contractObj, err := binding.NewErc20(addr, cli)
	if err != nil {
		return nil, err
	}
	return &Erc20{
		Cli:      cli,
		Contract: contractObj,
	}, nil
}

func (c Erc20) Approve(amount *big.Int) {
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
