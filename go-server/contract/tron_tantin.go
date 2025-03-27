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
	ks, ka, err := tron_keystore.InitKeyStore()
	if err != nil {
		panic(fmt.Sprintf("private key conversion failed %v", err))
	}
	return &TanTinTron{
		Ks:              ks,
		Ka:              ka,
		Cli:             cli,
		ContractAddress: ChainConfig.VoteContractAddress,
	}, nil
}

func (t *TanTinTron) Init() {
	// 手工
	txHash, err := t.AdminSetEnv()
	fmt.Println(txHash, err)
	// 手工
	txHash2, err2 := t.GrantBridgeRole("52ba824bfabc2bcfcdf7f0edbb486ebb05e1836c90e78047efeb949990f72e5f", ChainConfig.BridgeContractAddress)
	fmt.Println(txHash2, err2)
	// 手工
	txHash3, err3 := t.AdminSetToken(strings.TrimPrefix(ResourceIdUsdt, "0x"), 2, ChainConfig.UsdtAddress, false, false, false)
	fmt.Println(txHash3, err3)
}

func (t *TanTinTron) AdminSetEnv() (string, error) {
	_ = t.Ks.Unlock(*t.Ka, tron_keystore.KeyStorePassphrase)
	defer t.Ks.Lock(t.Ka.Address)

	triggerData := fmt.Sprintf("[{\"address\":\"%s\"}]", ChainConfig.BridgeContractAddress)
	fmt.Println(triggerData)
	tx, err := t.Cli.TriggerContract(OwnerAccount, t.ContractAddress, "adminSetEnv(address)", triggerData, 1500000000, 0, "", 0)
	if err != nil {
		return "", err
	}
	ctrlr := transaction.NewController(t.Cli, t.Ks, t.Ka, tx.Transaction)
	if err = ctrlr.ExecuteTransaction(); err != nil {
		return "", err
	}
	log.Println("tx hash: ", common.BytesToHexString(tx.GetTxid()))
	return common.BytesToHexString(tx.GetTxid()), nil
}

func (t *TanTinTron) GrantBridgeRole(role, addr string) (string, error) {
	_ = t.Ks.Unlock(*t.Ka, tron_keystore.KeyStorePassphrase)
	defer t.Ks.Lock(t.Ka.Address)

	triggerData := fmt.Sprintf("[{\"bytes32\":\"%s\"},{\"address\":\"%s\"}]", role, addr)
	tx, err := t.Cli.TriggerContract(OwnerAccount, t.ContractAddress, "grantRole(bytes32,address)", triggerData, 300000000, 0, "", 0)
	if err != nil {
		return "", err
	}
	ctrlr := transaction.NewController(t.Cli, t.Ks, t.Ka, tx.Transaction)
	if err = ctrlr.ExecuteTransaction(); err != nil {
		return "", err
	}
	log.Println("tx hash: ", common.BytesToHexString(tx.GetTxid()))
	return common.BytesToHexString(tx.GetTxid()), nil
}

func (t *TanTinTron) AdminSetToken(resourceID string, assetsType uint8, tokenAddress string, burnable, mintable, pause bool) (string, error) {
	_ = t.Ks.Unlock(*t.Ka, tron_keystore.KeyStorePassphrase)
	defer t.Ks.Lock(t.Ka.Address)

	triggerData := fmt.Sprintf("[{\"bytes32\":\"%s\"},{\"uint256\":\"%d\"},{\"address\":\"%s\"},{\"bool\":%v},{\"bool\":%v},{\"bool\":%v}]",
		resourceID,
		assetsType,
		tokenAddress,
		burnable,
		mintable,
		pause,
	)
	tx, err := t.Cli.TriggerContract(OwnerAccount, t.ContractAddress, "adminSetToken(bytes32,uint8,address,bool,bool,bool)", triggerData, 5000000000, 0, "", 0)
	if err != nil {
		return "", err
	}
	ctrlr := transaction.NewController(t.Cli, t.Ks, t.Ka, tx.Transaction)
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
	tx, err := t.Cli.TriggerContract(OwnerAccount, t.ContractAddress, "deposit(uint256,bytes32,address,uint256,bytes)", triggerData, 300000000, 0, "", 0)
	if err != nil {
		return "", err
	}
	ctrlr := transaction.NewController(t.Cli, t.Ks, t.Ka, tx.Transaction)
	if err = ctrlr.ExecuteTransaction(); err != nil {
		return "", err
	}
	log.Println("tx hash: ", common.BytesToHexString(tx.GetTxid()))
	return common.BytesToHexString(tx.GetTxid()), nil
}
