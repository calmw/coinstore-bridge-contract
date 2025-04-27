package contract

import (
	"coinstore/abi"
	"coinstore/bridge/chains/tron/trigger"
	"coinstore/tron_keystore"
	"coinstore/utils"
	"fmt"
	"github.com/calmw/tron-sdk/pkg/client"
	"github.com/calmw/tron-sdk/pkg/client/transaction"
	"github.com/calmw/tron-sdk/pkg/common"
	"github.com/calmw/tron-sdk/pkg/keystore"
	ethCommon "github.com/ethereum/go-ethereum/common"
	"github.com/status-im/keycard-go/hexutils"
	"google.golang.org/grpc"
	"log"
	"math/big"
	"strings"
	"time"
)

type TanTinTron struct {
	ContractAddress string
	Ka              *keystore.Account
	Ks              *keystore.KeyStore
	Cli             *client.GrpcClient
}

func NewTanTinTron(ka *keystore.Account, ks *keystore.KeyStore) (*TanTinTron, error) {
	endpoint := ChainConfig.RPC
	cli := client.NewGrpcClient(endpoint)
	err := cli.Start(grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	//prvKey := utils.ThreeDesDecrypt("gZIMfo6LJm6GYXdClPhIMfo6", os.Getenv("COIN_STORE_BRIDGE_TRON"))
	////prvKey := "3f9f4b92d709f167b8ba98b9f89a5ec5272973aeb8f1affd11d5d2c67c5acf62"
	//ks, ka, err := tron_keystore.InitKeyStore(prvKey)
	//if err != nil {
	//	panic(fmt.Sprintf("private key conversion failed %v", err))
	//}
	return &TanTinTron{
		Ks:              ks,
		Ka:              ka,
		Cli:             cli,
		ContractAddress: ChainConfig.TantinContractAddress,
	}, nil
}

func (t *TanTinTron) Init(adminAddress, feeAddress, serverAddress string) {
	txHash1, err1 := t.GrantRole(AdminRole, adminAddress)
	fmt.Println(txHash1, err1)
	time.Sleep(time.Second)
	txHash2, err2 := t.GrantRole(BridgeRole, ChainConfig.VoteContractAddress)
	fmt.Println(txHash2, err2)
	time.Sleep(time.Second)
	txHash, err := t.AdminSetEnv(feeAddress, serverAddress, serverAddress)
	fmt.Println(txHash, err)
}

func (t *TanTinTron) AdminSetEnv(feeAddress, serverAddress, bridgeAddress string) (string, error) {
	_ = t.Ks.Unlock(*t.Ka, tron_keystore.KeyStorePassphrase)
	defer t.Ks.Lock(t.Ka.Address)
	sigNonce, err := trigger.GetSigNonce(t.ContractAddress, OwnerAccount)
	if err != nil {
		return "", err
	}
	//sigNonce := big.NewInt(0)
	bridgeEth, _ := utils.TronToEth(bridgeAddress)
	feeEth, _ := utils.TronToEth(feeAddress)
	serverEth, _ := utils.TronToEth(serverAddress)
	signature, _ := abi.TantinAdminSetEnvSignatureTron(
		sigNonce,
		ethCommon.HexToAddress(feeEth),
		ethCommon.HexToAddress(serverEth),
		ethCommon.HexToAddress(bridgeEth),
	)
	triggerData := fmt.Sprintf("[{\"address\":\"%s\"},{\"address\":\"%s\"},{\"bytes\":\"%s\"}]",
		feeAddress,
		ChainConfig.BridgeContractAddress,
		fmt.Sprintf("%x", signature),
	)
	fmt.Println(triggerData)
	tx, err := t.Cli.TriggerContract(OwnerAccount, t.ContractAddress, "adminSetEnv(address,address,bytes)", triggerData, 1500000000, 0, "", 0)
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
	sigNonce, err := trigger.GetSigNonce(t.ContractAddress, OwnerAccount)
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
	triggerData := fmt.Sprintf("[{\"bytes32\":\"%s\"},{\"uint8\":\"%d\"},{\"address\":\"%s\"},{\"bool\":%v},{\"bool\":%v},{\"bool\":%v},{\"bytes\":\"%s\"}]",
		fmt.Sprintf("%x", [32]byte(resourceIdBytes)),
		assetsType,
		tokenAddress,
		burnable,
		mintable,
		pause,
		fmt.Sprintf("%x", signature),
	)
	fmt.Println(triggerData)
	tx, err := t.Cli.TriggerContract(OwnerAccount, t.ContractAddress, "adminSetToken(bytes32,uint8,address,bool,bool,bool,bytes)", triggerData, 6000000000, 0, "", 0)
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

func (t *TanTinTron) Deposit(destinationChainId, amount *big.Int, resourceId, recipient string) (string, error) {
	_ = t.Ks.Unlock(*t.Ka, tron_keystore.KeyStorePassphrase)
	defer t.Ks.Lock(t.Ka.Address)
	recipientEth, err := utils.TronToEth(recipient)
	signature, err := abi.TantinDepositSignatureTron(
		ethCommon.HexToAddress(recipientEth),
	)
	triggerData := fmt.Sprintf("[{\"uint256\":\"%s\"},{\"bytes32\":\"%s\"},{\"address\":\"%s\"},{\"uint256\":\"%s\"},{\"bytes\":\"%s\"}]",
		destinationChainId.String(), strings.TrimPrefix(resourceId, "0x"), recipient, amount.String(), fmt.Sprintf("%x", signature),
	)
	fmt.Println(triggerData)
	//tx, err := t.Cli.TriggerContract(OwnerAccount, t.ContractAddress, "deposit(uint256,bytes32,address,uint256,bytes)", triggerData, 300000000, 0, "", 0)
	tx, err := t.Cli.TriggerContract("TFBymbm7LrbRreGtByMPRD2HUyneKabsqb", t.ContractAddress, "deposit(uint256,bytes32,address,uint256,bytes)", triggerData, 300000000, 0, "", 0)
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
