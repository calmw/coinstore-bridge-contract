package contract

import (
	"coinstore/abi"
	"coinstore/binding"
	"coinstore/services"
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/status-im/keycard-go/hexutils"
	"log"
	"math/big"
	"strings"
	"time"
)

type BridgeEvm struct {
	Cli      *ethclient.Client
	Contract *binding.Bridge
}

func NewBridge() (*BridgeEvm, error) {
	err, cli := Client(ChainConfig)
	if err != nil {
		return nil, err
	}
	contractObj, err := binding.NewBridge(common.HexToAddress(ChainConfig.BridgeContractAddress), cli)
	if err != nil {
		return nil, err
	}
	return &BridgeEvm{
		Cli:      cli,
		Contract: contractObj,
	}, nil
}

func (b *BridgeEvm) Init() {
	b.AdminSetEnv()
	b.GrantRole(services.VoteRole, common.HexToAddress(ChainConfig.VoteContractAddress))
	b.GrantRole(services.AdminRole, common.HexToAddress(services.AdminAccount))
}

func (b *BridgeEvm) AdminSetEnv() {
	var res *types.Transaction
	signature, _ := abi.BridgeAdminSetEnvSignature(
		common.HexToAddress(ChainConfig.VoteContractAddress),
		big.NewInt(ChainConfig.BridgeId),
		big.NewInt(ChainConfig.ChainTypeId),
	)

	for {
		err, txOpts := GetAuth(b.Cli)
		if err != nil {
			log.Println(err)
			return
		}
		res, err = b.Contract.AdminSetEnv(
			txOpts,
			common.HexToAddress(ChainConfig.VoteContractAddress),
			big.NewInt(ChainConfig.BridgeId),
			big.NewInt(ChainConfig.ChainTypeId),
			signature,
		)
		if err == nil {
			break
		}
		fmt.Println(fmt.Sprintf("AdminSetEnv error: %v", err))
		time.Sleep(3 * time.Second)
	}
	fmt.Println(fmt.Sprintf("AdminSetEnv 成功"))
	for {
		receipt, err := b.Cli.TransactionReceipt(context.Background(), res.Hash())
		if err == nil && receipt.Status == 1 {
			break
		}
		time.Sleep(time.Second * 2)
	}

	fmt.Println(fmt.Sprintf("AdminSetEnv 确认成功 %v", res.Hash()))
}

func (b *BridgeEvm) GrantRole(role string, addr common.Address) {
	AdminRoleBytes := hexutils.HexToBytes(role)

	var res *types.Transaction

	for {
		err, txOpts := GetAuth(b.Cli)
		if err != nil {
			log.Println(err)
			return
		}
		res, err = b.Contract.GrantRole(txOpts, [32]byte(AdminRoleBytes), addr)
		if err == nil {
			break
		}
		time.Sleep(3 * time.Second)
	}
	log.Println(fmt.Sprintf("GrantRole 成功"))
	for {
		receipt, err := b.Cli.TransactionReceipt(context.Background(), res.Hash())
		if err == nil && receipt.Status == 1 {
			break
		}
		time.Sleep(time.Second * 2)
	}

	log.Println(fmt.Sprintf("GrantRole 确认成功"))
}

// AdminSetResource 0x0000000000000000000000000000000000000000
func (b *BridgeEvm) AdminSetResource(resourceId string, assetsType uint8, tokenAddress common.Address, fee *big.Int) {
	var res *types.Transaction
	resourceIdBytes := hexutils.HexToBytes(strings.TrimPrefix(resourceId, "0x"))
	signature, _ := abi.BridgeAdminSetResourceSignature(
		[32]byte(resourceIdBytes),
		assetsType, //uint8(2),
		common.HexToAddress(ChainConfig.UsdtAddress),
		tokenAddress,
		fee,
		false,
	)
	for {
		err, txOpts := GetAuth(b.Cli)
		if err != nil {
			log.Println(err)
			return
		}
		res, err = b.Contract.AdminSetResource(
			txOpts,
			[32]byte(resourceIdBytes),
			assetsType, //uint8(2),
			tokenAddress,
			fee,
			false,
			common.HexToAddress(ChainConfig.TantinContractAddress),
			signature,
		)
		if err == nil {
			break
		} else {
			fmt.Println(fmt.Sprintf("AdminSetResource error: %v", err))
		}
		time.Sleep(3 * time.Second)
	}
	fmt.Println(fmt.Sprintf("AdminSetResource 成功"))
	for {
		receipt, err := b.Cli.TransactionReceipt(context.Background(), res.Hash())
		if err == nil && receipt.Status == 1 {
			break
		}
		time.Sleep(time.Second * 2)
	}

	fmt.Println(fmt.Sprintf("AdminSetResource 确认成功"))
}
