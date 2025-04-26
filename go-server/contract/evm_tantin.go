package contract

import (
	"coinstore/abi"
	"coinstore/binding"
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

type TanTinEvm struct {
	Cli      *ethclient.Client
	Contract *binding.Tantin
}

func NewTanTin() (*TanTinEvm, error) {
	err, cli := Client(ChainConfig)
	if err != nil {
		return nil, err
	}
	contractObj, err := binding.NewTantin(common.HexToAddress(ChainConfig.TantinContractAddress), cli)
	if err != nil {
		return nil, err
	}
	return &TanTinEvm{
		Cli:      cli,
		Contract: contractObj,
	}, nil
}

func (c TanTinEvm) Init(adminAddress, feeAddress, serverAddress string) {
	c.GrantRole(AdminRole, common.HexToAddress(adminAddress))
	c.GrantRole(BridgeRole, common.HexToAddress(ChainConfig.VoteContractAddress))
	c.AdminSetEnv(feeAddress, serverAddress)
}

func (c TanTinEvm) LatestBlock() {
	number, err := c.Cli.BlockNumber(context.Background())
	fmt.Println(number, err)
	time.Sleep(time.Second * 5)
	number, err = c.Cli.BlockNumber(context.Background())
	fmt.Println(number, err)
}

func (c TanTinEvm) AdminSetEnv(feeAddress, serverAddress string) {
	var res *types.Transaction
	sigNonce, err := c.Contract.SigNonce(nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	signature, _ := abi.TantinAdminSetEnvSignature(sigNonce, common.HexToAddress(feeAddress), common.HexToAddress(serverAddress), common.HexToAddress(ChainConfig.BridgeContractAddress))

	for {
		err, txOpts := GetAuth(c.Cli)
		if err != nil {
			log.Println(err)
			return
		}
		res, err = c.Contract.AdminSetEnv(txOpts, common.HexToAddress(feeAddress), common.HexToAddress(serverAddress), common.HexToAddress(ChainConfig.BridgeContractAddress), signature)
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

func (c TanTinEvm) GrantRole(role string, addr common.Address) {
	AdminRoleBytes := hexutils.HexToBytes(role)

	var res *types.Transaction

	for {
		err, txOpts := GetAuth(c.Cli)
		if err != nil {
			log.Println(err)
			return
		}
		res, err = c.Contract.GrantRole(txOpts, [32]byte(AdminRoleBytes), addr)
		if err == nil {
			break
		}
		time.Sleep(3 * time.Second)
	}
	log.Println(fmt.Sprintf("GrantRole 成功"))
	for {
		receipt, err := c.Cli.TransactionReceipt(context.Background(), res.Hash())
		if err == nil && receipt.Status == 1 {
			break
		}
		time.Sleep(time.Second * 2)
	}

	log.Println(fmt.Sprintf("GrantRole 确认成功"))
}

func (c TanTinEvm) Deposit(receiver common.Address, resourceId [32]byte, destinationChainId, amount, price, priceTimestamp *big.Int) {
	reSignature, _ := abi.TantinDepositSignature(receiver)
	prSignature, _ := abi.EvmPriceSignature(big.NewInt(ChainConfig.ChainTypeId), price, priceTimestamp)
	var res *types.Transaction

	fmt.Println("参数11")
	for {
		err, txOpts := GetAuth(c.Cli)
		if err != nil {
			log.Println(err)
			return
		}
		fmt.Println("参数")
		fmt.Println(
			destinationChainId,
			fmt.Sprintf("0x%x", resourceId),
			receiver,
			amount,
			price,
			priceTimestamp,
			fmt.Sprintf("0x%x", reSignature),
			fmt.Sprintf("0x%x", prSignature),
		)
		res, err = c.Contract.Deposit(
			txOpts,
			destinationChainId,
			resourceId,
			receiver,
			amount,
			price,
			priceTimestamp,
			prSignature,
			reSignature,
		)
		if err == nil {
			break
		} else {
			log.Println(fmt.Sprintf("deposit error: %v", err))
		}
		time.Sleep(3 * time.Second)
	}
	log.Println(fmt.Sprintf("Deposit 成功 %s", res.Hash().String()))
	for {
		receipt, err := c.Cli.TransactionReceipt(context.Background(), res.Hash())
		if err == nil && receipt.Status == 1 {
			break
		}
		time.Sleep(time.Second * 2)
	}

	log.Println(fmt.Sprintf("Deposit 确认成功 %s", res.Hash()))
}
