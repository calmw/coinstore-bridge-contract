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
	"github.com/status-im/keycard-go/hexutils"
	"google.golang.org/grpc"
	"log"
	"math/big"
	"os"
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
	prvKey := utils.ThreeDesDecrypt("gZIMfo6LJm6GYXdClPhIMfo6", os.Getenv("COIN_STORE_BRIDGE_TRON"))
	ks, ka, err := tron_keystore.InitKeyStore(prvKey)
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
	//txHash1, err1 := t.GrantRole(AdminRole, OwnerAccount)
	//fmt.Println(txHash1, err1)
	//// 手工
	//txHash, err := t.AdminSetEnv(ChainConfig.BridgeContractAddress)
	//fmt.Println(txHash, err)
	//// 手工
	//txHash2, err2 := t.GrantRole(BridgeRole, ChainConfig.BridgeContractAddress)
	//fmt.Println(txHash2, err2)
	// 手工
	txHash3, err3 := t.AdminSetToken(ResourceIdUsdt, 2, ChainConfig.UsdtAddress, false, false, false)
	fmt.Println(txHash3, err3)
}

func (t *TanTinTron) AdminSetEnv(bridgeAddress string) (string, error) {
	_ = t.Ks.Unlock(*t.Ka, tron_keystore.KeyStorePassphrase)
	defer t.Ks.Lock(t.Ka.Address)
	sigNonce, err := tron.GetSigNonce(t.ContractAddress, OwnerAccount)
	if err != nil {
		return "", err
	}
	bridgeEth, _ := utils.TronToEth(bridgeAddress)
	signature, _ := abi.TantinAdminSetEnvSignatureTron(
		sigNonce,
		ethCommon.HexToAddress(bridgeEth),
	)
	triggerData := fmt.Sprintf("[{\"address\":\"%s\"},{\"bytes\":\"%s\"}]",
		ChainConfig.BridgeContractAddress,
		fmt.Sprintf("%x", signature),
	)
	fmt.Println(triggerData)
	tx, err := t.Cli.TriggerContract(OwnerAccount, t.ContractAddress, "adminSetEnv(address,bytes)", triggerData, 1500000000, 0, "", 0)
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

func (t *TanTinTron) GrantRole(role, addr string) (string, error) {
	_ = t.Ks.Unlock(*t.Ka, tron_keystore.KeyStorePassphrase)
	defer t.Ks.Lock(t.Ka.Address)
	triggerData := fmt.Sprintf("[{\"bytes32\":\"%s\"},{\"address\":\"%s\"}]", role, addr)
	tx, err := t.Cli.TriggerContract(OwnerAccount, t.ContractAddress, "grantRole(bytes32,address)", triggerData, 9500000000, 0, "", 0)
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

func (t *TanTinTron) AdminSetToken(resourceId string, assetsType uint8, tokenAddress string, burnable, mintable, pause bool) (string, error) {
	_ = t.Ks.Unlock(*t.Ka, tron_keystore.KeyStorePassphrase)
	defer t.Ks.Lock(t.Ka.Address)
	resourceIdBytes := hexutils.HexToBytes(strings.TrimPrefix(resourceId, "0x"))
	sigNonce, err := tron.GetSigNonce(t.ContractAddress, OwnerAccount)
	if err != nil {
		return "", err
	}
	tokenEth, _ := utils.TronToEth(tokenAddress)
	signature, _ := abi.TantinAdminSetTokenSignatureTron(
		sigNonce,
		big.NewInt(ChainConfig.BridgeId),
		[32]byte(resourceIdBytes),
		assetsType,
		ethCommon.HexToAddress(tokenEth),
		burnable,
		mintable,
		pause,
	)
	triggerData := fmt.Sprintf("[{\"bytes32\":\"%s\"},{\"uint8\":\"%d\"},{\"address\":\"%s\"},{\"bool\":%v},{\"bool\":%v},{\"bool\":%v},{\"bytes\":%s}]",
		fmt.Sprintf("%x", [32]byte(resourceIdBytes)),
		assetsType,
		tokenAddress,
		burnable,
		mintable,
		pause,
		fmt.Sprintf("%x", signature),
	)
	fmt.Println(triggerData)
	fmt.Println("~~~~~~~~~~")
	tx, err := t.Cli.TriggerContract(OwnerAccount, t.ContractAddress, "adminSetToken(bytes32,uint8,address,bool,bool,bool,bytes)", triggerData, 5000000000, 0, "", 0)
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
