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

type VoteEvm struct {
	Cli      *ethclient.Client
	Contract *binding.Vote
}

func NewVote() (*VoteEvm, error) {
	err, cli := Client(ChainConfig)
	if err != nil {
		return nil, err
	}
	contractObj, err := binding.NewVote(common.HexToAddress(ChainConfig.VoteContractAddress), cli)
	if err != nil {
		return nil, err
	}
	return &VoteEvm{
		Cli:      cli,
		Contract: contractObj,
	}, nil
}

func (c VoteEvm) Init() {
	c.AdminSetEnv(big.NewInt(100), big.NewInt(1))
	c.GrantBridgeRole(common.HexToAddress(ChainConfig.BridgeContractAddress))
	c.GrantRelayerRole(common.HexToAddress("0x80B27CDE65Fafb1f048405923fD4a624fEa2d1C6"))
}

func (c VoteEvm) AdminSetEnv(expiry *big.Int, relayerThreshold *big.Int) {
	var res *types.Transaction
	sigNonce, err := c.Contract.SigNonce(nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	signature, _ := abi.VoteAdminSetEnvSignature(
		sigNonce,
		common.HexToAddress(ChainConfig.BridgeContractAddress),
		expiry,
		relayerThreshold,
	)

	for {
		err, txOpts := GetAuth(c.Cli)
		if err != nil {
			log.Println(err)
			return
		}
		res, err = c.Contract.AdminSetEnv(
			txOpts,
			common.HexToAddress(ChainConfig.BridgeContractAddress),
			expiry,
			relayerThreshold,
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
		receipt, err := c.Cli.TransactionReceipt(context.Background(), res.Hash())
		if err == nil && receipt.Status == 1 {
			break
		}
		time.Sleep(time.Second * 2)
	}

	fmt.Println(fmt.Sprintf("AdminSetEnv 确认成功"))
}

func (c VoteEvm) GrantBridgeRole(addr common.Address) {
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

func (c VoteEvm) GrantRelayerRole(addr common.Address) {
	RelayerRole := "e2b7fb3b832174769106daebcfd6d1970523240dda11281102db9363b83b0dc4"
	BridgeRoleBytes := hexutils.HexToBytes(RelayerRole)

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
