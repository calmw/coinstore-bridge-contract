package contract

import (
	"coinstore/binding"
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
	b.GrantVoteRole(common.HexToAddress(ChainConfig.VoteContractAddress))
}

func (b *BridgeEvm) AdminSetEnv() {
	var res *types.Transaction

	for {
		err, txOpts := GetAuth(b.Cli)
		if err != nil {
			log.Println(err)
			return
		}
		res, err = b.Contract.AdminSetEnv(txOpts, common.HexToAddress(ChainConfig.VoteContractAddress), big.NewInt(ChainConfig.BridgeId), big.NewInt(ChainConfig.ChainTypeId))
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

	fmt.Println(fmt.Sprintf("AdminSetEnv 确认成功", res.Hash()))
}

func (b *BridgeEvm) GrantAdminRole(addr common.Address) {
	AdminRole := "a49807205ce4d355092ef5a8a18f56e8913cf4a201fbe287825b095693c21775"
	AdminRoleBytes := hexutils.HexToBytes(AdminRole)

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
	log.Println(fmt.Sprintf("GrantAdminRole 成功"))
	for {
		receipt, err := b.Cli.TransactionReceipt(context.Background(), res.Hash())
		if err == nil && receipt.Status == 1 {
			break
		}
		time.Sleep(time.Second * 2)
	}

	log.Println(fmt.Sprintf("GrantAdminRole 确认成功"))
}

func (b *BridgeEvm) GrantVoteRole(addr common.Address) {
	VoteRole := "c65b6dc445843af69e7af2fc32667c7d3b98b02602373e2d0a7a047f274806f7"
	VoteRoleBytes := hexutils.HexToBytes(VoteRole)

	var res *types.Transaction

	for {
		err, txOpts := GetAuth(b.Cli)
		if err != nil {
			log.Println(err)
			return
		}
		res, err = b.Contract.GrantRole(txOpts, [32]byte(VoteRoleBytes), addr)
		if err == nil {
			break
		}
		time.Sleep(3 * time.Second)
	}
	log.Println(fmt.Sprintf("GrantVoteRole 成功"))
	for {
		receipt, err := b.Cli.TransactionReceipt(context.Background(), res.Hash())
		if err == nil && receipt.Status == 1 {
			break
		}
		time.Sleep(time.Second * 2)
	}

	log.Println(fmt.Sprintf("GrantVoteRole 确认成功"))
}

func (b *BridgeEvm) AdminSetResource(fee *big.Int) {
	var res *types.Transaction
	resourceIdBytes := hexutils.HexToBytes(strings.TrimPrefix(ResourceIdUsdt, "0x"))
	for {
		err, txOpts := GetAuth(b.Cli)
		if err != nil {
			log.Println(err)
			return
		}
		res, err = b.Contract.AdminSetResource(
			txOpts,
			[32]byte(resourceIdBytes),
			uint8(2),
			common.HexToAddress(ChainConfig.UsdtAddress),
			fee,
			false,
			common.HexToAddress(ChainConfig.TantinContractAddress),
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

	resourceIdBytes = hexutils.HexToBytes(strings.TrimPrefix(ResourceIdCoin, "0x"))
	for {
		err, txOpts := GetAuth(b.Cli)
		if err != nil {
			log.Println(err)
			return
		}
		res, err = b.Contract.AdminSetResource(
			txOpts,
			[32]byte(resourceIdBytes),
			uint8(2),
			common.HexToAddress("0x0000000000000000000000000000000000000000"),
			fee,
			false,
			common.HexToAddress(ChainConfig.TantinContractAddress),
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
