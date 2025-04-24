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
	b.GrantRole(AdminRole, common.HexToAddress(AdminAccount))
	b.GrantRole(VoteRole, common.HexToAddress(ChainConfig.VoteContractAddress))
	b.AdminSetEnv()
	b.AdminSetResource(
		ResourceIdUsdt,
		2,
		common.HexToAddress(ChainConfig.UsdtAddress),
		big.NewInt(100),
		false,
		false,
		false,
	)
	b.AdminSetResource(
		ResourceIdUsdc,
		2,
		common.HexToAddress(ChainConfig.UsdcAddress),
		big.NewInt(100),
		false,
		false,
		false,
	)
}

func (b *BridgeEvm) AdminSetEnv() {
	var res *types.Transaction
	sigNonce, err := b.Contract.SigNonce(nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	signature, _ := abi.BridgeAdminSetEnvSignature(
		sigNonce,
		common.HexToAddress(ChainConfig.VoteContractAddress),
		big.NewInt(ChainConfig.BridgeId),
		big.NewInt(ChainConfig.ChainTypeId),
	)
	fmt.Println(
		sigNonce,
		common.HexToAddress(ChainConfig.VoteContractAddress),
		big.NewInt(ChainConfig.BridgeId),
		big.NewInt(ChainConfig.ChainTypeId),
		fmt.Sprintf("%x", signature),
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
		fmt.Println(err, "~~~~")
		time.Sleep(3 * time.Second)
	}
	log.Println(fmt.Sprintf("GrantRole 成功"), res.Hash())
	for {
		receipt, err := b.Cli.TransactionReceipt(context.Background(), res.Hash())
		if err == nil && receipt.Status == 1 {
			break
		}
		time.Sleep(time.Second * 2)
	}

	log.Println(fmt.Sprintf("GrantRole 确认成功"))
}

// 	bridge.AdminSetResource(
//		contract.ResourceIdUsdt,
//		2,
//		common.HexToAddress(contract.ChainConfig.UsdtAddress),
//		big.NewInt(100),
//	)

func AdminSetResource(resourceId string, assetsType uint8, tokenAddress common.Address, fee *big.Int) {

}

// AdminSetResource 0x0000000000000000000000000000000000000000
func (b *BridgeEvm) AdminSetResource(resourceId string, assetsType uint8, tokenAddress common.Address, fee *big.Int, pause bool, burnable bool, mintable bool) {
	var res *types.Transaction
	resourceIdBytes := hexutils.HexToBytes(strings.TrimPrefix(resourceId, "0x"))

	for {
		sigNonce, err := b.Contract.SigNonce(nil)
		if err != nil {
			fmt.Println(err)
			return
		}
		chainId, err := b.Contract.ChainId(nil)
		if err != nil {
			fmt.Println(err)
			return
		}
		signature, _ := abi.BridgeAdminSetResourceSignature(
			sigNonce,
			chainId,
			[32]byte(resourceIdBytes),
			assetsType,
			tokenAddress,
			common.HexToAddress(ChainConfig.TantinContractAddress),
			fee,
			pause,
			burnable,
			mintable,
		)
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
			pause,
			burnable,
			mintable,
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
