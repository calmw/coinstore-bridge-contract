package binding

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/btcsuite/btcd/btcec/v2"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/fbsobreira/gotron-sdk/pkg/address"
	"github.com/fbsobreira/gotron-sdk/pkg/client"
	"github.com/fbsobreira/gotron-sdk/pkg/client/transaction"
	"github.com/fbsobreira/gotron-sdk/pkg/common"
	"github.com/fbsobreira/gotron-sdk/pkg/keystore"
	"github.com/fbsobreira/gotron-sdk/pkg/store"
	"github.com/status-im/keycard-go/hexutils"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/proto"
	"io"
	"log"
	"math/big"
	"net/http"
	"os"
	"strings"
)

const (
	AccountName  = "my_account"
	Passphrase   = "account_pwd"
	ShastaGrpc   = "grpc.shasta.trongrid.io:50051"
	NileGrpc     = "grpc.nile.trongrid.io:50051"
	OwnerAccount = "TFBymbm7LrbRreGtByMPRD2HUyneKabsqb"
	UsdtAddress  = "TU6UjUJadm8TungBHvL4n9apv8Jns4wJiz"
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

func CreateTransaction() {
	url := "https://api.shasta.trongrid.io/wallet/createtransaction"

	payload := strings.NewReader("{\"owner_address\":\"TZ4UXDV5ZhNW7fb2AMSbgfAEZ7hWsnYS2g\",\"to_address\":\"TPswDDCAWhJAZGdHPidFg5nEf8TkNToDX1\",\"amount\":1000,\"visible\":true}")

	req, _ := http.NewRequest("POST", url, payload)

	req.Header.Add("accept", "application/json")
	req.Header.Add("content-type", "application/json")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

	fmt.Println(string(body))
}

func BroadcastHex() {
	url := "https://api.shasta.trongrid.io/wallet/broadcasthex"

	payload := strings.NewReader("{\"transaction\":\"0A8A010A0202DB2208C89D4811359A28004098A4E0A6B52D5A730802126F0A32747970652E676F6F676C65617069732E636F6D2F70726F746F636F6C2E5472616E736665724173736574436F6E747261637412390A07313030303030311215415A523B449890854C8FC460AB602DF9F31FE4293F1A15416B0580DA195542DDABE288FEC436C7D5AF769D24206412418BF3F2E492ED443607910EA9EF0A7EF79728DAAAAC0EE2BA6CB87DA38366DF9AC4ADE54B2912C1DEB0EE6666B86A07A6C7DF68F1F9DA171EEE6A370B3CA9CBBB00\"}")

	req, _ := http.NewRequest("POST", url, payload)

	req.Header.Add("accept", "application/json")
	req.Header.Add("content-type", "application/json")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

	fmt.Println(string(body))

}

func BroadcastTransaction() {
	url := "https://api.shasta.trongrid.io/wallet/broadcasttransaction"

	payload := strings.NewReader("{\"raw_data\":{\"contract\":[{\"parameter\":{\"value\":{\"amount\":1000,\"owner_address\":\"41608f8da72479edc7dd921e4c30bb7e7cddbe722e\",\"to_address\":\"41e9d79cc47518930bc322d9bf7cddd260a0260a8d\"},\"type_url\":\"type.googleapis.com/protocol.TransferContract\"},\"type\":\"TransferContract\"}],\"ref_block_bytes\":\"5e4b\",\"ref_block_hash\":\"47c9dc89341b300d\",\"expiration\":1591089627000,\"timestamp\":1591089567635},\"raw_data_hex\":\"0a025e4b220847c9dc89341b300d40f8fed3a2a72e5a66080112620a2d747970652e676f6f676c65617069732e636f6d2f70726f746f636f6c2e5472616e73666572436f6e747261637412310a1541608f8da72479edc7dd921e4c30bb7e7cddbe722e121541e9d79cc47518930bc322d9bf7cddd260a0260a8d18e8077093afd0a2a72e\"}")

	req, _ := http.NewRequest("POST", url, payload)

	req.Header.Add("accept", "application/json")
	req.Header.Add("content-type", "application/json")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

	fmt.Println(string(body))
}

func PrivateKeyToWalletAddress(pk string) (string, error) {
	privateKeyBytes, err := hex.DecodeString(pk)
	if err != nil {
		return "", err
	}
	if len(privateKeyBytes) != common.Secp256k1PrivateKeyBytesLength {
		fmt.Println(common.ErrBadKeyLength)
	}
	sk, _ := btcec.PrivKeyFromBytes(privateKeyBytes)
	//sk, pubKey := btcec.PrivKeyFromBytes(privateKeyBytes)
	// sk.PubKey().ToECDSA() == pubKey.ToECDSA() ,值一样
	addr := address.PubkeyToAddress(*sk.PubKey().ToECDSA())
	return addr.String(), nil
}

func TransferCoin(privateKey, fromAddress, toAddress string, amount int64) (string, error) {
	privateKeyBytes, _ := hex.DecodeString(privateKey)
	c := client.NewGrpcClient("grpc.shasta.trongrid.io:50051")
	err := c.Start(grpc.WithInsecure())
	if err != nil {
		return "", err
	}
	tx, err := c.Transfer(fromAddress, toAddress, amount)
	if err != nil {
		return "", err
	}
	rawData, err := proto.Marshal(tx.Transaction.GetRawData())
	if err != nil {
		return "", err
	}
	h256h := sha256.New()
	h256h.Write(rawData)
	hash := h256h.Sum(nil)
	sk, _ := btcec.PrivKeyFromBytes(privateKeyBytes)
	signature, err := crypto.Sign(hash, sk.ToECDSA())
	if err != nil {
		return "", err
	}
	tx.Transaction.Signature = append(tx.Transaction.Signature, signature)
	res, err := c.Broadcast(tx.Transaction)
	if err != nil || !res.Result {
		return "", errors.New("broadcast error")
	}
	txHash := strings.ToLower(hexutils.BytesToHex(tx.Txid))
	return txHash, nil
}

func GetKeyFromPrivateKey(privateKey, name, passphrase string) (*keystore.KeyStore, *keystore.Account, error) {
	privateKey = strings.TrimPrefix(privateKey, "0x")

	if store.DoesNamedAccountExist(name) {
		return nil, nil, fmt.Errorf("account %s already exists", name)
	}

	privateKeyBytes, err := hex.DecodeString(privateKey)
	if err != nil {
		return nil, nil, err
	}
	if len(privateKeyBytes) != common.Secp256k1PrivateKeyBytesLength {
		return nil, nil, common.ErrBadKeyLength
	}

	sk, _ := btcec.PrivKeyFromBytes(privateKeyBytes)
	ks := store.FromAccountName(name)
	account, err := ks.ImportECDSA(sk.ToECDSA(), passphrase)
	if err != nil {
		return nil, nil, err
	}
	return store.FromAccountName(name), &account, err
}

func TransferToken() (string, error) {
	// 携带的数据
	triggerData := "[{\"address\":\"TEy2BtGxixqhbcM7w65rJiotAerSBFR48W\"},{\"uint256\":\"1\"}]"
	cli := client.NewGrpcClient(NileGrpc)
	err := cli.Start(grpc.WithInsecure())
	if err != nil {
		return "", err
	}
	from := "TFBymbm7LrbRreGtByMPRD2HUyneKabsqb"
	tokenAddress := "TU6UjUJadm8TungBHvL4n9apv8Jns4wJiz"
	// 构建Tx，注意，此处只是构造，tx并没有被发送
	// 参数分别为 发送者地址 合约地址 调用方法(正常字符串) gas Tx发送的TRX数量 Tx发送的TRC20代币 代币数量
	tx, err := cli.TriggerContract(from, tokenAddress, "transfer(address,uint256)", triggerData, 300000000, 0, "", 0)
	if err != nil {
		return "", err
	}
	privateKey := os.Getenv("COINSTORE_BRIDGE_TRON")
	_, _, err = GetKeyFromPrivateKey(privateKey, AccountName, Passphrase)
	//if strings.Contains(err.Error(),"already exists")
	if err != nil && !strings.Contains(err.Error(), "already exists") {
		return "", err
	}
	// 获得keystore与account
	ks, acct, err := store.UnlockedKeystore(from, Passphrase)
	if err != nil {
		return "", err
	}
	// 封装Tx
	ctrlr := transaction.NewController(cli, ks, acct, tx.Transaction)
	// 真正执行Tx，并判断执行结果
	if err = ctrlr.ExecuteTransaction(); err != nil {
		return "", err
	}
	// 此时Tx才上链
	log.Println("tx hash: ", common.BytesToHexString(tx.GetTxid()))
	//log.Println(6, common.BytesToHexString(tx.GetResult().GetMessage()))
	return common.BytesToHexString(tx.GetTxid()), nil
}

func AA() {
	toAddress, err := address.Base58ToAddress("TPrEMmYc2nz5bHbjs3M2f1gZ9PtWsLzr8A")
	fmt.Println(toAddress, err)
	fmt.Println(toAddress.Hex())
}
