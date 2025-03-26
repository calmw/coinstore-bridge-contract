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

type TanTinTron struct {
	ContractAddress string
	Ka              *keystore.Account
	Ks              *keystore.KeyStore
	Cli             *client.GrpcClient
}

func NewTanTinTron() (*TanTinTron, error) {
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
	return &TanTinTron{
		Ks:              ks,
		Ka:              ka,
		Cli:             cli,
		ContractAddress: ChainConfig.VoteContractAddress,
	}, nil
}

func (t *TanTinTron) Init() {
	txHash, err := t.AdminSetEnv()
	fmt.Println(txHash, err)
	t.FreshPrk()
	txHash2, err2 := t.GrantBridgeRole("52ba824bfabc2bcfcdf7f0edbb486ebb05e1836c90e78047efeb949990f72e5f", ChainConfig.BridgeContractAddress)
	fmt.Println(txHash2, err2)
	t.FreshPrk()
	txHash3, err3 := t.AdminSetToken(strings.TrimPrefix(ResourceIdUsdt, "0x"), "2", ChainConfig.UsdtAddress, false, false, false)
	fmt.Println(txHash3, err3)
}

func (t *TanTinTron) FreshPrk() {
	_, _, _ = GetKeyFromPrivateKey(ChainConfig.PrivateKey, AccountName, Passphrase)
	ks, ka, _ := store.UnlockedKeystore(OwnerAccount, Passphrase)
	t.Ks = ks
	t.Ka = ka
}

func (t *TanTinTron) AdminSetEnv() (string, error) {
	triggerData := fmt.Sprintf("[{\"address\":\"%s\"}]", ChainConfig.BridgeContractAddress)
	fmt.Println(triggerData)
	cli := client.NewGrpcClient(NileGrpc)
	err := cli.Start(grpc.WithInsecure())
	if err != nil {
		return "", err
	}
	tx, err := cli.TriggerContract(OwnerAccount, t.ContractAddress, "adminSetEnv(address)", triggerData, 300000000, 0, "", 0)
	if err != nil {
		return "", err
	}
	ctrlr := transaction.NewController(cli, t.Ks, t.Ka, tx.Transaction)
	if err = ctrlr.ExecuteTransaction(); err != nil {
		return "", err
	}
	log.Println("tx hash: ", common.BytesToHexString(tx.GetTxid()))
	return common.BytesToHexString(tx.GetTxid()), nil
}

func (t *TanTinTron) GrantBridgeRole(role, addr string) (string, error) {
	triggerData := fmt.Sprintf("[{\"bytes32\":\"%s\"},{\"address\":\"%s\"}]", role, addr)
	cli := client.NewGrpcClient(NileGrpc)
	err := cli.Start(grpc.WithInsecure())
	if err != nil {
		return "", err
	}
	tx, err := cli.TriggerContract(OwnerAccount, t.ContractAddress, "grantRole(bytes32,address)", triggerData, 300000000, 0, "", 0)
	if err != nil {
		return "", err
	}
	ctrlr := transaction.NewController(cli, t.Ks, t.Ka, tx.Transaction)
	if err = ctrlr.ExecuteTransaction(); err != nil {
		return "", err
	}
	log.Println("tx hash: ", common.BytesToHexString(tx.GetTxid()))
	return common.BytesToHexString(tx.GetTxid()), nil
}

func (t *TanTinTron) AdminSetToken(resourceID, assetsType, tokenAddress string, burnable, mintable, pause bool) (string, error) {
	triggerData := fmt.Sprintf("[{\"bytes32\":\"%s\"},{\"uint8\":\"%s\"},{\"address\":\"%s\"},{\"bool\":%v},{\"bool\":%v},{\"bool\":%v}]",
		resourceID, assetsType, tokenAddress, burnable, mintable, pause,
	)
	cli := client.NewGrpcClient(NileGrpc)
	err := cli.Start(grpc.WithInsecure())
	if err != nil {
		return "", err
	}
	tx, err := cli.TriggerContract(OwnerAccount, t.ContractAddress, "adminSetToken(bytes32,uint8,address,bool,bool,bool)", triggerData, 300000000, 0, "", 0)
	if err != nil {
		return "", err
	}
	ctrlr := transaction.NewController(cli, t.Ks, t.Ka, tx.Transaction)
	if err = ctrlr.ExecuteTransaction(); err != nil {
		return "", err
	}
	log.Println("tx hash: ", common.BytesToHexString(tx.GetTxid()))
	return common.BytesToHexString(tx.GetTxid()), nil
}

func (t *TanTinTron) Deposit(destinationChainId, resourceId, recipient, signature string, amount *big.Int) (string, error) {
	triggerData := fmt.Sprintf("[{\"uint256\":\"%s\"},{\"bytes32\":\"%s\"},{\"address\":\"%s\"},{\"uint256\":\"%s\"},{\"bytes\":\"%s\"}]",
		destinationChainId, resourceId, recipient, amount.String(), signature,
	)
	fmt.Println(triggerData)
	cli := client.NewGrpcClient(NileGrpc)
	err := cli.Start(grpc.WithInsecure())
	if err != nil {
		return "", err
	}
	tx, err := cli.TriggerContract(OwnerAccount, t.ContractAddress, "deposit(uint256,bytes32,address,uint256,bytes)", triggerData, 300000000, 0, "", 0)
	if err != nil {
		return "", err
	}
	ctrlr := transaction.NewController(cli, t.Ks, t.Ka, tx.Transaction)
	if err = ctrlr.ExecuteTransaction(); err != nil {
		return "", err
	}
	log.Println("tx hash: ", common.BytesToHexString(tx.GetTxid()))
	return common.BytesToHexString(tx.GetTxid()), nil
}
