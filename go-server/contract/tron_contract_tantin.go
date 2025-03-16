package contract

import (
	"coinstore/binding"
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/status-im/keycard-go/hexutils"
	"log"
	"math/big"
	"time"
)

const (
	ResourceIdUsdt = "ac589789ed8c9d2c61f17b13369864b5f181e58eba230a6ee4ec4c3e7750cd1d"
	ResourceIdCoin = "ac589789ed8c9d2c61f17b13369864b5f181e58eba230a6ee4ec4c3e7750cd1c"
)

type TanTin struct {
	Cli      *ethclient.Client
	Contract *binding.Tantin
}

func NewTanTin() (*TanTin, error) {
	err, cli := Client(ChainConfig)
	if err != nil {
		return nil, err
	}
	contractObj, err := binding.NewTantin(common.HexToAddress(ChainConfig.TantinContractAddress), cli)
	if err != nil {
		return nil, err
	}
	return &TanTin{
		Cli:      cli,
		Contract: contractObj,
	}, nil
}

func (c TanTin) Init() {
	c.AdminSetEnv()
	c.GrantBridgeRole(common.HexToAddress(ChainConfig.BridgeContractAddress))
}

func (c TanTin) AdminSetEnv() {
	var res *types.Transaction

	for {
		err, txOpts := GetAuth(c.Cli)
		if err != nil {
			log.Println(err)
			return
		}
		res, err = c.Contract.AdminSetEnv(txOpts, common.HexToAddress(ChainConfig.BridgeContractAddress))
		if err == nil {
			break
		}
		fmt.Println(fmt.Sprintf("AdminSetEnv error: %v", err))
		time.Sleep(3 * time.Second)
	}
	fmt.Println(fmt.Sprintf("AdminSetEnv 成功"))
	for {
		receipt, err := c.Cli.TransactionReceipt(context.Background(), res.Hash())
		if err == nil && receipt.Status == 1 {
			break
		}
		time.Sleep(time.Second * 2)
	}

	fmt.Println(fmt.Sprintf("AdminSetEnv 确认成功"))
}

func (c TanTin) GrantBridgeRole(addr common.Address) {
	BridgeRole := "52ba824bfabc2bcfcdf7f0edbb486ebb05e1836c90e78047efeb949990f72e5f"
	BridgeRoleBytes := hexutils.HexToBytes(BridgeRole)

	var res *types.Transaction

	for {
		err, txOpts := GetAuth(c.Cli)
		if err != nil {
			log.Println(err)
			return
		}
		res, err = c.Contract.GrantRole(txOpts, [32]byte(BridgeRoleBytes), addr)
		if err == nil {
			break
		}
		time.Sleep(3 * time.Second)
	}
	log.Println(fmt.Sprintf("GrantBridgeRole 成功"))
	for {
		receipt, err := c.Cli.TransactionReceipt(context.Background(), res.Hash())
		if err == nil && receipt.Status == 1 {
			break
		}
		time.Sleep(time.Second * 2)
	}

	log.Println(fmt.Sprintf("GrantBridgeRole 确认成功"))
}

func (c TanTin) AdminSetToken() {
	resourceIdBytes := hexutils.HexToBytes(ResourceIdUsdt)
	var res *types.Transaction
	for {
		err, txOpts := GetAuth(c.Cli)
		if err != nil {
			log.Println(err)
			return
		}
		res, err = c.Contract.AdminSetToken(
			txOpts,
			[32]byte(resourceIdBytes),
			uint8(2),
			common.HexToAddress(ChainConfig.UsdtAddress),
			false,
			false,
			false,
		)
		if err == nil {
			break
		}
		time.Sleep(3 * time.Second)
	}
	log.Println(fmt.Sprintf("AdminSetToken 成功"))
	for {
		receipt, err := c.Cli.TransactionReceipt(context.Background(), res.Hash())
		if err == nil && receipt.Status == 1 {
			break
		}
		time.Sleep(time.Second * 2)
	}

	log.Println(fmt.Sprintf("AdminSetToken 确认成功"))

	resourceIdBytes = hexutils.HexToBytes(ResourceIdCoin)
	for {
		err, txOpts := GetAuth(c.Cli)
		if err != nil {
			log.Println(err)
			return
		}
		res, err = c.Contract.AdminSetToken(
			txOpts,
			[32]byte(resourceIdBytes),
			uint8(1),
			common.HexToAddress("0x0000000000000000000000000000000000000000"),
			false,
			false,
			false,
		)
		if err == nil {
			break
		}
		time.Sleep(3 * time.Second)
	}
	log.Println(fmt.Sprintf("AdminSetToken 成功"))
	for {
		receipt, err := c.Cli.TransactionReceipt(context.Background(), res.Hash())
		if err == nil && receipt.Status == 1 {
			break
		}
		time.Sleep(time.Second * 2)
	}

	log.Println(fmt.Sprintf("AdminSetToken 确认成功"))
}

func (c TanTin) Deposit(receiver common.Address, resourceId [32]byte, destinationChainId, amount, fee *big.Int) {

	token, err := NewErc20(common.HexToAddress(ChainConfig.UsdtAddress))
	if err != nil {
		fmt.Println(err)
		return
	}
	token.Approve(amount)

	var res *types.Transaction
	var txOpts *bind.TransactOpts

	for {
		if fee.Int64() > 0 {
			err, txOpts = GetAuthWithValue(c.Cli, fee.Add(fee, amount))
			if err != nil {
				log.Println(err)
				return
			}
		} else {
			err, txOpts = GetAuth(c.Cli)
			if err != nil {
				log.Println(err)
				return
			}
		}

		res, err = c.Contract.Deposit(
			txOpts,
			destinationChainId,
			resourceId,
			receiver,
			amount,
		)
		if err == nil {
			break
		} else {
			log.Println(fmt.Sprintf("deposit error: %v", err))
		}
		time.Sleep(3 * time.Second)
	}
	log.Println(fmt.Sprintf("Deposit 成功"))
	for {
		receipt, err := c.Cli.TransactionReceipt(context.Background(), res.Hash())
		if err == nil && receipt.Status == 1 {
			break
		}
		time.Sleep(time.Second * 2)
	}

	log.Println(fmt.Sprintf("Deposit 确认成功 %s", res.Hash()))
}
