package contract

import (
	"fmt"
	"github.com/fbsobreira/gotron-sdk/pkg/client"
	"github.com/fbsobreira/gotron-sdk/pkg/client/transaction"
	"github.com/fbsobreira/gotron-sdk/pkg/common"
	"github.com/fbsobreira/gotron-sdk/pkg/store"
	"google.golang.org/grpc"
	"log"
	"math/big"
	"os"
	"strings"
)

type Trc20 struct {
	TokenAddress string
	Client       *client.GrpcClient
}

func NewTrc20(address string) (*Trc20, error) {
	cli := client.NewGrpcClient(NileGrpc)
	err := cli.Start(grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	return &Trc20{
		TokenAddress: address,
		Client:       cli,
	}, nil
}

func (t *Trc20) Approve(spender string, value *big.Int) (string, error) {
	triggerData := fmt.Sprintf("[{\"address\":\"%s\"},{\"uint256\":\"%s\"}]", spender, value.String())
	cli := client.NewGrpcClient(NileGrpc)
	err := cli.Start(grpc.WithInsecure())
	if err != nil {
		return "", err
	}
	tx, err := cli.TriggerContract(OwnerAccount, t.TokenAddress, "approve(address,uint256)", triggerData, 300000000, 0, "", 0)
	if err != nil {
		return "", err
	}
	privateKey := os.Getenv("COINSTORE_BRIDGE_TRON")
	_, _, err = GetKeyFromPrivateKey(privateKey, AccountName, Passphrase)
	if err != nil && !strings.Contains(err.Error(), "already exists") {
		return "", err
	}
	ks, acct, err := store.UnlockedKeystore(OwnerAccount, Passphrase)
	if err != nil {
		return "", err
	}
	ctrlr := transaction.NewController(cli, ks, acct, tx.Transaction)
	if err = ctrlr.ExecuteTransaction(); err != nil {
		return "", err
	}
	log.Println("tx hash: ", common.BytesToHexString(tx.GetTxid()))
	return common.BytesToHexString(tx.GetTxid()), nil
}

func (t *Trc20) Transfer(to string, value *big.Int) (string, error) {
	triggerData := fmt.Sprintf("[{\"address\":\"%s\"},{\"uint256\":\"%s\"}]", to, value.String())
	cli := client.NewGrpcClient(NileGrpc)
	err := cli.Start(grpc.WithInsecure())
	if err != nil {
		return "", err
	}
	tx, err := cli.TriggerContract(OwnerAccount, t.TokenAddress, "transfer(address,uint256)", triggerData, 300000000, 0, "", 0)
	if err != nil {
		return "", err
	}
	privateKey := os.Getenv("COINSTORE_BRIDGE_TRON")
	_, _, err = GetKeyFromPrivateKey(privateKey, AccountName, Passphrase)
	if err != nil && !strings.Contains(err.Error(), "already exists") {
		return "", err
	}
	ks, acct, err := store.UnlockedKeystore(OwnerAccount, Passphrase)
	if err != nil {
		return "", err
	}
	ctrlr := transaction.NewController(cli, ks, acct, tx.Transaction)
	if err = ctrlr.ExecuteTransaction(); err != nil {
		return "", err
	}
	log.Println("tx hash: ", common.BytesToHexString(tx.GetTxid()))
	return common.BytesToHexString(tx.GetTxid()), nil
}
