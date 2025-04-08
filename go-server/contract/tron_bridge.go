package contract

import (
	"coinstore/abi"
	"coinstore/bridge/tron"
	"coinstore/tron_keystore"
	"coinstore/utils"
	"fmt"
	ethCommon "github.com/ethereum/go-ethereum/common"
	"github.com/fbsobreira/gotron-sdk/pkg/client"
	"github.com/fbsobreira/gotron-sdk/pkg/client/transaction"
	"github.com/fbsobreira/gotron-sdk/pkg/common"
	"github.com/fbsobreira/gotron-sdk/pkg/keystore"
	"google.golang.org/grpc"
	"log"
	"math/big"
	"os"
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
	prvKey := utils.ThreeDesDecrypt("gZIMfo6LJm6GYXdClPhIMfo6", os.Getenv("COIN_STORE_BRIDGE_TRON"))
	ks, ka, err := tron_keystore.InitKeyStore(prvKey)
	if err != nil {
		panic(fmt.Sprintf("private key conversion failed %v", err))
	}
	return &BridgeTron{
		Ks:              ks,
		Ka:              ka,
		Cli:             cli,
		ContractAddress: ChainConfig.BridgeContractAddress,
	}, nil
}

func (b *BridgeTron) Init() {
	//txHash2, err2 := b.GrantRole(AdminRole, OwnerAccount)
	//fmt.Println(txHash2, err2)
	//txHash3, err3 := b.GrantRole(VoteRole, ChainConfig.VoteContractAddress)
	//fmt.Println(txHash3, err3)
	txHash, err := b.AdminSetEnv()
	fmt.Println(txHash, err)
	txHash4, err4 := b.AdminSetResource(big.NewInt(1))
	fmt.Println(txHash4, err4)
}

func (b *BridgeTron) AdminSetEnv() (string, error) {

	_ = b.Ks.Unlock(*b.Ka, tron_keystore.KeyStorePassphrase)
	defer b.Ks.Lock(b.Ka.Address)
	sigNonce, err := tron.GetSigNonce(b.ContractAddress, OwnerAccount)
	if err != nil {
		return "", err
	}

	voteEth, _ := utils.TronToEth(ChainConfig.VoteContractAddress)
	signature, _ := abi.BridgeAdminSetEnvSignatureTron(
		sigNonce,
		ethCommon.HexToAddress(voteEth),
		big.NewInt(ChainConfig.BridgeId),
		big.NewInt(ChainConfig.ChainTypeId),
	)

	triggerData := fmt.Sprintf("[{\"address\":\"%s\"},{\"uint256\":\"%d\"},{\"uint256\":\"%d\"},{\"bytes\":\"%s\"}]",
		ChainConfig.VoteContractAddress,
		ChainConfig.BridgeId,
		ChainConfig.ChainTypeId,
		fmt.Sprintf("0x%x", signature),
	)
	fmt.Println("~~~~~~")
	fmt.Println(triggerData)
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
	_ = b.Ks.Unlock(*b.Ka, tron_keystore.KeyStorePassphrase)
	defer b.Ks.Lock(b.Ka.Address)
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

func (b *BridgeTron) GrantRole(role, addr string) (string, error) {
	_ = b.Ks.Unlock(*b.Ka, tron_keystore.KeyStorePassphrase)
	defer b.Ks.Lock(b.Ka.Address)
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
	_ = b.Ks.Unlock(*b.Ka, tron_keystore.KeyStorePassphrase)
	defer b.Ks.Lock(b.Ka.Address)
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
