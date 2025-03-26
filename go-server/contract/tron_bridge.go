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

const (
	AccountName  = "my_account"
	Passphrase   = "account_pwd"
	ShastaGrpc   = "grpc.shasta.trongrid.io:50051"
	NileGrpc     = "grpc.nile.trongrid.io:50051"
	OwnerAccount = "TFBymbm7LrbRreGtByMPRD2HUyneKabsqb"
)

type BridgeTron struct {
	ContractAddress string
	Ka              *keystore.Account
	Ks              *keystore.KeyStore
	Cli             *client.GrpcClient
}

func NewBridgeTron() (*BridgeTron, error) {
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
	return &BridgeTron{
		Ks:              ks,
		Ka:              ka,
		Cli:             cli,
		ContractAddress: ChainConfig.BridgeContractAddress,
	}, nil
}

func (b *BridgeTron) Init() {
	b.FreshPrk()
	txHash, err := b.AdminSetEnv()
	fmt.Println(txHash, err)
	b.FreshPrk()
	txHash2, err2 := b.GrantVoteRole("c65b6dc445843af69e7af2fc32667c7d3b98b02602373e2d0a7a047f274806f7", ChainConfig.VoteContractAddress)
	fmt.Println(txHash2, err2)
	b.FreshPrk()
	txHash3, err3 := b.AdminSetResource(big.NewInt(1))
	fmt.Println(txHash3, err3)
}

func (b *BridgeTron) FreshPrk() {
	_, _, _ = GetKeyFromPrivateKey(ChainConfig.PrivateKey, AccountName, Passphrase)
	ks, ka, _ := store.UnlockedKeystore(OwnerAccount, Passphrase)
	b.Ks = ks
	b.Ka = ka
}

func (b *BridgeTron) AdminSetEnv() (string, error) {
	triggerData := fmt.Sprintf("[{\"address\":\"%s\"},{\"uint256\":\"%d\"},{\"uint256\":\"%d\"}]", ChainConfig.VoteContractAddress, ChainConfig.BridgeId, ChainConfig.ChainTypeId)
	tx, err := b.Cli.TriggerContract(OwnerAccount, b.ContractAddress, "adminSetEnv(address,uint256,uint256)", triggerData, 300000000, 0, "", 0)
	if err != nil {
		return "", err
	}
	ctrlr := transaction.NewController(b.Cli, b.Ks, b.Ka, tx.Transaction)
	if err = ctrlr.ExecuteTransaction(); err != nil {
		return "", err
	}
	return common.BytesToHexString(tx.GetTxid()), nil
}

func (b *BridgeTron) GrantAdminRole(role, addr string) (string, error) {
	triggerData := fmt.Sprintf("[{\"bytes32\":\"%s\"},{\"address\":\"%s\"}]", role, addr)
	tx, err := b.Cli.TriggerContract(OwnerAccount, b.ContractAddress, "grantRole(bytes32,address)", triggerData, 300000000, 0, "", 0)
	if err != nil {
		return "", err
	}
	ctrlr := transaction.NewController(b.Cli, b.Ks, b.Ka, tx.Transaction)
	if err = ctrlr.ExecuteTransaction(); err != nil {
		return "", err
	}
	log.Println("tx hash: ", common.BytesToHexString(tx.GetTxid()))
	return common.BytesToHexString(tx.GetTxid()), nil
}

func (b *BridgeTron) GrantVoteRole(role, addr string) (string, error) {
	triggerData := fmt.Sprintf("[{\"bytes32\":\"%s\"},{\"address\":\"%s\"}]", role, addr)
	tx, err := b.Cli.TriggerContract(OwnerAccount, b.ContractAddress, "grantRole(bytes32,address)", triggerData, 9500000000, 0, "", 0)
	if err != nil {
		return "", err
	}
	ctrlr := transaction.NewController(b.Cli, b.Ks, b.Ka, tx.Transaction)
	if err = ctrlr.ExecuteTransaction(); err != nil {
		return "", err
	}
	log.Println("tx hash: ", common.BytesToHexString(tx.GetTxid()))
	return common.BytesToHexString(tx.GetTxid()), nil
}

func (b *BridgeTron) AdminSetResource(fee *big.Int) (string, error) {
	triggerData := fmt.Sprintf("[{\"bytes32\":\"%s\"},{\"uint8\":\"%d\"},{\"address\":\"%s\"},{\"uint256\":\"%s\"},{\"bool\":%v},{\"address\":\"%s\"}]",
		strings.TrimPrefix(ResourceIdUsdt, "0x"),
		uint8(2),
		ChainConfig.UsdtAddress,
		fee.String(),
		false,
		ChainConfig.TantinContractAddress,
	)
	tx, err := b.Cli.TriggerContract(OwnerAccount, b.ContractAddress, "adminSetResource(bytes32,uint8,address,uint256,bool,address)", triggerData, 300000000, 0, "", 0)
	if err != nil {
		return "", err
	}
	ctrlr := transaction.NewController(b.Cli, b.Ks, b.Ka, tx.Transaction)
	if err = ctrlr.ExecuteTransaction(); err != nil {
		return "", err
	}
	log.Println("tx hash: ", common.BytesToHexString(tx.GetTxid()))
	return common.BytesToHexString(tx.GetTxid()), nil
}
