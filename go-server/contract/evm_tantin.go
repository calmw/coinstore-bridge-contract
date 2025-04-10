package contract

import (
	"coinstore/abi"
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
	"strings"
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

func (c TanTinEvm) Init() {
	c.GrantRole(AdminRole, common.HexToAddress(AdminAccount))
	c.GrantRole(BridgeRole, common.HexToAddress(ChainConfig.VoteContractAddress))
	c.AdminSetEnv()
	c.AdminSetToken(ResourceIdUsdt, 2, common.HexToAddress(ChainConfig.UsdtAddress), false, false, false)
	c.AdminSetToken(ResourceIdUsdc, 2, common.HexToAddress(ChainConfig.UsdcAddress), false, false, false)
}

func (c TanTinEvm) AdminSetEnv() {
	var res *types.Transaction
	sigNonce, err := c.Contract.SigNonce(nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	signature, _ := abi.TantinAdminSetEnvSignature(sigNonce, common.HexToAddress(ChainConfig.BridgeContractAddress))

	for {
		err, txOpts := GetAuth(c.Cli)
		if err != nil {
			log.Println(err)
			return
		}
		res, err = c.Contract.AdminSetEnv(txOpts, common.HexToAddress(ChainConfig.BridgeContractAddress), signature)
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
		fmt.Println(err)
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

func (c TanTinEvm) AdminSetToken(resourceId string, assetsType uint8, tokenAddress common.Address, burnable, mintable, pause bool) {
	resourceIdBytes := hexutils.HexToBytes(strings.TrimPrefix(resourceId, "0x"))
	sigNonce, err := c.Contract.SigNonce(nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	signature, _ := abi.TantinAdminSetTokenSignature(
		sigNonce,
		big.NewInt(ChainConfig.BridgeId),
		[32]byte(resourceIdBytes),
		assetsType,
		tokenAddress,
		burnable,
		mintable,
		pause,
	)
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
			assetsType,
			tokenAddress,
			burnable,
			mintable,
			pause,
			signature,
		)
		if err == nil {
			break
		}
		fmt.Println(err)
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

func (c TanTinEvm) Deposit(receiver common.Address, resourceId [32]byte, destinationChainId, amount *big.Int) {
	signature, _ := abi.TantinDepositSignature(receiver)
	token, err := NewErc20(common.HexToAddress(ChainConfig.UsdtAddress))
	if err != nil {
		fmt.Println(err)
		return
	}
	token.Approve(amount)

	var res *types.Transaction
	var txOpts *bind.TransactOpts

	for {
		err, txOpts = GetAuth(c.Cli)
		if err != nil {
			log.Println(err)
			return
		}

		res, err = c.Contract.Deposit(
			txOpts,
			destinationChainId,
			resourceId,
			receiver,
			amount,
			signature,
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
