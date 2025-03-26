package contract

import (
	"fmt"
	"github.com/fbsobreira/gotron-sdk/pkg/client"
	"github.com/fbsobreira/gotron-sdk/pkg/client/transaction"
	"github.com/fbsobreira/gotron-sdk/pkg/common"
	"github.com/fbsobreira/gotron-sdk/pkg/keystore"
	"github.com/fbsobreira/gotron-sdk/pkg/store"
	"google.golang.org/grpc"
	"log"
	"math/big"
	"strings"
)

type VoteTron struct {
	ContractAddress string
	Ka              *keystore.Account
	Ks              *keystore.KeyStore
	Cli             *client.GrpcClient
}

func NewVoteTron() (*VoteTron, error) {
	cli := client.NewGrpcClient(NileGrpc)
	err := cli.Start(grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	_, _, err = GetKeyFromPrivateKey(ChainConfig.PrivateKey, AccountName, Passphrase)
	if err != nil && !strings.Contains(err.Error(), "already exists") {
		return nil, err
	}
	ks, ka, err := store.UnlockedKeystore(OwnerAccount, Passphrase)
	if err != nil {
		return nil, err
	}
	return &VoteTron{
		Ks:              ks,
		Ka:              ka,
		Cli:             cli,
		ContractAddress: ChainConfig.VoteContractAddress,
	}, nil
}

func (v *VoteTron) Init() {
	v.FreshPrk()
	txHash, err := v.AdminSetEnv(big.NewInt(100000), big.NewInt(1))
	fmt.Println(txHash, err)
	v.FreshPrk()
	txHash2, err2 := v.GrantBridgeRole("52ba824bfabc2bcfcdf7f0edbb486ebb05e1836c90e78047efeb949990f72e5f", ChainConfig.BridgeContractAddress)
	fmt.Println(txHash2, err2)
	v.FreshPrk()
	txHash3, err3 := v.GrantBridgeRole("e2b7fb3b832174769106daebcfd6d1970523240dda11281102db9363b83b0dc4", OwnerAccount)
	fmt.Println(txHash3, err3)
}

func (v *VoteTron) FreshPrk() {
	_, _, _ = GetKeyFromPrivateKey(ChainConfig.PrivateKey, AccountName, Passphrase)
	ks, ka, _ := store.UnlockedKeystore(OwnerAccount, Passphrase)
	v.Ks = ks
	v.Ka = ka
}

func (v *VoteTron) AdminSetEnv(expiry *big.Int, relayerThreshold *big.Int) (string, error) {
	triggerData := fmt.Sprintf("[{\"address\":\"%s\"},{\"uint256\":\"%s\"},{\"uint256\":\"%s\"}]", ChainConfig.BridgeContractAddress, expiry.String(), relayerThreshold.String())
	cli := client.NewGrpcClient(NileGrpc)
	err := cli.Start(grpc.WithInsecure())
	if err != nil {
		return "", err
	}
	tx, err := cli.TriggerContract(OwnerAccount, v.ContractAddress, "adminSetEnv(address,uint256,uint256)", triggerData, 300000000, 0, "", 0)
	if err != nil {
		return "", err
	}
	ctrlr := transaction.NewController(cli, v.Ks, v.Ka, tx.Transaction)
	if err = ctrlr.ExecuteTransaction(); err != nil {
		return "", err
	}
	log.Println("tx hash: ", common.BytesToHexString(tx.GetTxid()))
	return common.BytesToHexString(tx.GetTxid()), nil
}

func (v *VoteTron) GrantBridgeRole(role, addr string) (string, error) {
	triggerData := fmt.Sprintf("[{\"bytes32\":\"%s\"},{\"address\":\"%s\"}]", role, addr)
	cli := client.NewGrpcClient(NileGrpc)
	err := cli.Start(grpc.WithInsecure())
	if err != nil {
		return "", err
	}
	tx, err := cli.TriggerContract(OwnerAccount, v.ContractAddress, "grantRole(bytes32,address)", triggerData, 300000000, 0, "", 0)
	if err != nil {
		return "", err
	}
	ctrlr := transaction.NewController(cli, v.Ks, v.Ka, tx.Transaction)
	if err = ctrlr.ExecuteTransaction(); err != nil {
		return "", err
	}
	log.Println("tx hash: ", common.BytesToHexString(tx.GetTxid()))
	return common.BytesToHexString(tx.GetTxid()), nil
}
