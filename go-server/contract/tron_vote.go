package contract

import (
	"coinstore/tron_keystore"
	"fmt"
	"github.com/fbsobreira/gotron-sdk/pkg/client"
	"github.com/fbsobreira/gotron-sdk/pkg/client/transaction"
	"github.com/fbsobreira/gotron-sdk/pkg/common"
	"github.com/fbsobreira/gotron-sdk/pkg/keystore"
	"google.golang.org/grpc"
	"log"
	"math/big"
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
	ks, ka, err := tron_keystore.InitKeyStore()
	if err != nil {
		panic(fmt.Sprintf("private key conversion failed %v", err))
	}
	return &VoteTron{
		Ks:              ks,
		Ka:              ka,
		Cli:             cli,
		ContractAddress: ChainConfig.VoteContractAddress,
	}, nil
}

func (v *VoteTron) Init() {
	txHash, err := v.AdminSetEnv(big.NewInt(100000), big.NewInt(1))
	fmt.Println(txHash, err)
	txHash2, err2 := v.GrantBridgeRole("52ba824bfabc2bcfcdf7f0edbb486ebb05e1836c90e78047efeb949990f72e5f", ChainConfig.BridgeContractAddress)
	fmt.Println(txHash2, err2)
	txHash3, err3 := v.GrantBridgeRole("e2b7fb3b832174769106daebcfd6d1970523240dda11281102db9363b83b0dc4", OwnerAccount)
	fmt.Println(txHash3, err3)
}

func (v *VoteTron) AdminSetEnv(expiry *big.Int, relayerThreshold *big.Int) (string, error) {

	_ = v.Ks.Unlock(*v.Ka, tron_keystore.KeyStorePassphrase)
	defer v.Ks.Lock(v.Ka.Address)

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
	_ = v.Ks.Unlock(*v.Ka, tron_keystore.KeyStorePassphrase)
	defer v.Ks.Lock(v.Ka.Address)
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
