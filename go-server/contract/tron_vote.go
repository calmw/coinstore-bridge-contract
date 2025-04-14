package contract

import (
	"coinstore/abi"
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
	"time"
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
	prvKey := utils.ThreeDesDecrypt("gZIMfo6LJm6GYXdClPhIMfo6", os.Getenv("COIN_STORE_BRIDGE_TRON"))
	ks, ka, err := tron_keystore.InitKeyStore(prvKey)
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
	txHash1, err1 := v.GrantRole(AdminRole, OwnerAccount)
	fmt.Println(txHash1, err1)
	time.Sleep(time.Second)
	txHash2, err2 := v.GrantRole(BridgeRole, ChainConfig.BridgeContractAddress)
	fmt.Println(txHash2, err2)
	time.Sleep(time.Second)
	txHash3, err3 := v.GrantRole(RelayerRole, OwnerAccount) // TODO 线上更改
	fmt.Println(txHash3, err3)
	time.Sleep(time.Second)
	txHash4, err4 := v.GrantRole(RelayerRole, Realyer1Account)
	fmt.Println(txHash4, err4)
	time.Sleep(time.Second)
	txHash5, err5 := v.GrantRole(RelayerRole, Realyer2Account)
	fmt.Println(txHash5, err5)
	time.Sleep(time.Second)
	txHash6, err6 := v.GrantRole(RelayerRole, Realyer3Account)
	fmt.Println(txHash6, err6)
	time.Sleep(time.Second)
	txHash, err := v.AdminSetEnv(ChainConfig.BridgeContractAddress, big.NewInt(100000), big.NewInt(1))
	fmt.Println(txHash, err)
}

func (v *VoteTron) AdminSetEnv(bridgeAddress string, expiry *big.Int, relayerThreshold *big.Int) (string, error) {
	_ = v.Ks.Unlock(*v.Ka, tron_keystore.KeyStorePassphrase)
	defer v.Ks.Lock(v.Ka.Address)
	//sigNonce, err := tron.GetSigNonce(v.ContractAddress, OwnerAccount)
	//if err != nil {
	//	return "", err
	//}
	sigNonce := big.NewInt(0)
	bridgeEth, _ := utils.TronToEth(bridgeAddress)
	//tantinEth, _ := utils.TronToEth(tantinAddress)
	signature, _ := abi.VoteAdminSetEnvSignatureTron(
		sigNonce,
		ethCommon.HexToAddress(bridgeEth),
		expiry,
		relayerThreshold,
	)
	triggerData := fmt.Sprintf("[{\"address\":\"%s\"},{\"uint256\":\"%s\"},{\"uint256\":\"%s\"},{\"bytes\":\"%s\"}]",
		ChainConfig.BridgeContractAddress,
		expiry.String(),
		relayerThreshold.String(),
		fmt.Sprintf("%x", signature),
	)
	fmt.Println(triggerData)
	tx, err := v.Cli.TriggerContract(OwnerAccount, ChainConfig.VoteContractAddress, "adminSetEnv(address,uint256,uint256,bytes)", triggerData, 300000000, 0, "", 0)
	fmt.Println("!!!!", err)
	if err != nil {
		return "", err
	}
	ctrlr := transaction.NewController(v.Cli, v.Ks, v.Ka, tx.Transaction)
	if err = ctrlr.ExecuteTransaction(); err != nil {
		fmt.Println("!!!! 2 ", err)
		return "", err
	}
	log.Println("tx hash: ", common.BytesToHexString(tx.GetTxid()))
	return common.BytesToHexString(tx.GetTxid()), nil
}

func (v *VoteTron) GrantRole(role, addr string) (string, error) {
	_ = v.Ks.Unlock(*v.Ka, tron_keystore.KeyStorePassphrase)
	defer v.Ks.Lock(v.Ka.Address)
	triggerData := fmt.Sprintf("[{\"bytes32\":\"%s\"},{\"address\":\"%s\"}]", role, addr)
	tx, err := v.Cli.TriggerContract(OwnerAccount, v.ContractAddress, "grantRole(bytes32,address)", triggerData, 9500000000, 0, "", 0)
	if err != nil {
		return "", err
	}
	ctrlr := transaction.NewController(v.Cli, v.Ks, v.Ka, tx.Transaction)
	if err = ctrlr.ExecuteTransaction(); err != nil {
		return "", err
	}
	log.Println("tx hash: ", common.BytesToHexString(tx.GetTxid()))
	return common.BytesToHexString(tx.GetTxid()), nil
}
