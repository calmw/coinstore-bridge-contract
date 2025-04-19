package contract

import (
	"fmt"
	"github.com/calmw/tron-sdk/pkg/client"
	"github.com/calmw/tron-sdk/pkg/client/transaction"
	"github.com/calmw/tron-sdk/pkg/common"
	"google.golang.org/grpc"
	"log"
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
	//Ka              *keystore.Account
	//Ks              *keystore.KeyStore
	Cli *client.GrpcClient
}

func NewBridgeTron() (*BridgeTron, error) {
	cli := client.NewGrpcClient(NileGrpc)
	err := cli.Start(grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	//prvKey := utils.ThreeDesDecrypt("gZIMfo6LJm6GYXdClPhIMfo6", os.Getenv("COIN_STORE_BRIDGE_TRON"))
	//ks, ka, err := tron_keystore.InitKeyStore(prvKey)
	//if err != nil {
	//	panic(fmt.Sprintf("private key conversion failed %v", err))
	//}
	return &BridgeTron{
		//Ks:              ks,
		//Ka:              ka,
		Cli:             cli,
		ContractAddress: ChainConfig.BridgeContractAddress,
	}, nil
}

func (b *BridgeTron) Init() {
	txHash20, err20 := b.GrantRoleTest(AdminRole, OwnerAccount)
	fmt.Println(txHash20, err20)
	//txHash2, err2 := b.GrantRole(AdminRole, OwnerAccount)
	//fmt.Println(txHash2, err2)
	//time.Sleep(time.Second)
	//txHash3, err3 := b.GrantRole(VoteRole, ChainConfig.VoteContractAddress)
	//fmt.Println(txHash3, err3)
	//time.Sleep(time.Second)
	//txHash, err := b.AdminSetEnv()
	//fmt.Println(txHash, err)
	//time.Sleep(time.Second)
	//txHash4, err4 := b.AdminSetResource(ResourceIdUsdt, 2, ChainConfig.UsdtAddress, big.NewInt(100), false, false, false)
	//fmt.Println(txHash4, err4)
	//time.Sleep(time.Second)
	//txHash5, err5 := b.AdminSetResource(ResourceIdUsdc, 2, ChainConfig.UsdcAddress, big.NewInt(100), false, false, false)
	//fmt.Println(txHash5, err5)
	//time.Sleep(time.Second)
	//txHash6, err6 := b.AdminSetResource(ResourceIdEth, 2, ChainConfig.WEthAddress, big.NewInt(100), false, false, false)
	//fmt.Println(txHash6, err6)
}

//func (b *BridgeTron) AdminSetEnv() (string, error) {
//
//	_ = b.Ks.Unlock(*b.Ka, tron_keystore.KeyStorePassphrase)
//	defer b.Ks.Lock(b.Ka.Address)
//	sigNonce, err := tron.GetSigNonce(b.ContractAddress, OwnerAccount)
//	if err != nil {
//		return "", err
//	}
//	//sigNonce := big.NewInt(0)
//
//	voteEth, _ := utils.TronToEth(ChainConfig.VoteContractAddress)
//	signature, _ := abi.BridgeAdminSetEnvSignatureTron(
//		sigNonce,
//		ethCommon.HexToAddress(voteEth),
//		big.NewInt(ChainConfig.BridgeId),
//		big.NewInt(ChainConfig.ChainTypeId),
//	)
//
//	triggerData := fmt.Sprintf("[{\"address\":\"%s\"},{\"uint256\":\"%d\"},{\"uint256\":\"%d\"},{\"bytes\":\"%s\"}]",
//		ChainConfig.VoteContractAddress,
//		ChainConfig.BridgeId,
//		ChainConfig.ChainTypeId,
//		fmt.Sprintf("%x", signature),
//	)
//	fmt.Println(triggerData)
//	tx, err := b.Cli.TriggerContract(OwnerAccount, b.ContractAddress, "adminSetEnv(address,uint256,uint256,bytes)", triggerData, 300000000, 0, "", 0)
//	if err != nil {
//		return "", err
//	}
//	ctrlr := transaction.NewController(b.Cli, b.Ks, b.Ka, tx.Transaction)
//	if err = ctrlr.ExecuteTransaction(); err != nil {
//		return "", err
//	}
//	return common.BytesToHexString(tx.GetTxid()), nil
//}

//func (b *BridgeTron) GrantRole(role, addr string) (string, error) {
//	_ = b.Ks.Unlock(*b.Ka, tron_keystore.KeyStorePassphrase)
//	defer b.Ks.Lock(b.Ka.Address)
//	triggerData := fmt.Sprintf("[{\"bytes32\":\"%s\"},{\"address\":\"%s\"}]", role, addr)
//	tx, err := b.Cli.TriggerContract(OwnerAccount, b.ContractAddress, "grantRole(bytes32,address)", triggerData, 9500000000, 0, "", 0)
//
//	fmt.Println(111, b.ContractAddress, err)
//	if err != nil {
//		return "", err
//	}
//	ctrlr := transaction.NewController(b.Cli, b.Ks, b.Ka, tx.Transaction)
//	if err = ctrlr.ExecuteTransaction(); err != nil {
//		return "", err
//	}
//	log.Println("tx hash: ", common.BytesToHexString(tx.GetTxid()))
//	return common.BytesToHexString(tx.GetTxid()), nil
//}

func (b *BridgeTron) GrantRoleTest(role, addr string) (string, error) {
	account := "TEz4CMzy3mgtVECcYxu5ui9nJfgv3oXhyx"
	triggerData := fmt.Sprintf("[{\"bytes32\":\"%s\"},{\"address\":\"%s\"}]", role, addr)
	tx, err := b.Cli.TriggerContract(account, b.ContractAddress, "grantRole(bytes32,address)", triggerData, 9500000000, 0, "", 0)

	fmt.Println(111, b.ContractAddress, err)
	if err != nil {
		return "", err
	}
	ctrlr := transaction.NewController(b.Cli, nil, nil, tx.Transaction)
	if err = ExecuteTronTransaction(ctrlr, 3448148188, account, "ttbridge_9d8f7b6a5c4e3d2f1a0b9c8d7e6f5a4b3c2d1e0f"); err != nil {
		return "", err
	}
	log.Println("tx hash: ", common.BytesToHexString(tx.GetTxid()))
	return common.BytesToHexString(tx.GetTxid()), nil
}

//func (b *BridgeTron) AdminSetResource(resourceId string, assetsType uint8, tokenAddress string, fee *big.Int, pause bool, burnable bool, mintable bool) (string, error) {
//	_ = b.Ks.Unlock(*b.Ka, tron_keystore.KeyStorePassphrase)
//	defer b.Ks.Lock(b.Ka.Address)
//	resourceIdBytes := hexutils.HexToBytes(strings.TrimPrefix(resourceId, "0x"))
//	sigNonce, err := tron.GetSigNonce(b.ContractAddress, OwnerAccount)
//	if err != nil {
//		return "", err
//	}
//	//sigNonce := big.NewInt(1)
//	tokenEth, _ := utils.TronToEth(tokenAddress)
//	tantinEth, _ := utils.TronToEth(ChainConfig.TantinContractAddress)
//	signature, _ := abi.BridgeAdminSetResourceSignatureTron(
//		sigNonce,
//		big.NewInt(ChainConfig.BridgeId),
//		[32]byte(resourceIdBytes),
//		assetsType,
//		ethCommon.HexToAddress(tokenEth),
//		ethCommon.HexToAddress(tantinEth),
//		fee,
//		pause,
//		burnable,
//		mintable,
//	)
//	triggerData := fmt.Sprintf("[{\"bytes32\":\"%s\"},{\"uint8\":\"%d\"},{\"address\":\"%s\"},{\"uint256\":\"%s\"},{\"bool\":%v},{\"bool\":%v},{\"bool\":%v},{\"address\":\"%s\"},{\"bytes\":\"%s\"}]",
//		strings.TrimPrefix(resourceId, "0x"),
//		assetsType,
//		tokenAddress,
//		fee,
//		false,
//		false,
//		false,
//		ChainConfig.TantinContractAddress,
//		fmt.Sprintf("%x", signature),
//	)
//	fmt.Println(triggerData)
//	tx, err := b.Cli.TriggerContract(OwnerAccount, b.ContractAddress, "adminSetResource(bytes32,uint8,address,uint256,bool,bool,bool,address,bytes)", triggerData, 300000000, 0, "", 0)
//	if err != nil {
//		return "", err
//	}
//	ctrlr := transaction.NewController(b.Cli, b.Ks, b.Ka, tx.Transaction)
//	if err = ctrlr.ExecuteTransaction(); err != nil {
//		return "", err
//	}
//	log.Println("tx hash: ", common.BytesToHexString(tx.GetTxid()))
//	return common.BytesToHexString(tx.GetTxid()), nil
//}
