package binding

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/btcsuite/btcd/btcec/v2"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/fbsobreira/gotron-sdk/pkg/address"
	"github.com/fbsobreira/gotron-sdk/pkg/client"
	"github.com/fbsobreira/gotron-sdk/pkg/common"
	"github.com/status-im/keycard-go/hexutils"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/proto"
	"io"
	"math/big"
	"net/http"
	"strings"
	"testing"
)

type Trc20 struct {
	Address string
}

func NewTrc20(address string) (*Trc20, error) {
	return &Trc20{Address: address}, nil
}

func (t *Trc20) Approve(to string, value *big.Int) (*types.Transaction, error) {
	return nil, nil
}

func TestSend(t *testing.T) {

}

func Aa() {
	// {
	//    "visible": false,
	//    "txID": "48d0ac3ff12adafc4267833ebec61a2c108891519b9244bf515c94ca43aac62b",
	//    "raw_data": {
	//        "contract": [
	//            {
	//                "parameter": {
	//                    "value": {
	//                        "data": "40c10f190000000000000000000000003942fda93c573e2ce9e85b0bb00ba98a144f27f6000000000000000000000000000000000000000000000000002386f26fc10000",
	//                        "owner_address": "413942fda93c573e2ce9e85b0bb00ba98a144f27f6",
	//                        "contract_address": "412ffd22a9021bd03f05ce0af413eb0516abe8ef00"
	//                    },
	//                    "type_url": "type.googleapis.com/protocol.TriggerSmartContract"
	//                },
	//                "type": "TriggerSmartContract"
	//            }
	//        ],
	//        "ref_block_bytes": "b4f1",
	//        "ref_block_hash": "36794c80995b9f68",
	//        "expiration": 1741599294000,
	//        "fee_limit": 10000000000,
	//        "timestamp": 1741599235387
	//    },
	//    "raw_data_hex": "0a02b4f1220836794c80995b9f6840b084a1fbd7325aae01081f12a9010a31747970652e676f6f676c65617069732e636f6d2f70726f746f636f6c2e54726967676572536d617274436f6e747261637412740a15413942fda93c573e2ce9e85b0bb00ba98a144f27f61215412ffd22a9021bd03f05ce0af413eb0516abe8ef00224440c10f190000000000000000000000003942fda93c573e2ce9e85b0bb00ba98a144f27f6000000000000000000000000000000000000000000000000002386f26fc1000070bbba9dfbd732900180c8afa025",
	//    "signature": [
	//        "98b705ee927bdcc5662a4b72ee90458265c673076846fbe899ba8c5a8f2f7bfd27487fe6a187cdb497f3c129af305146f86634a700411e753bbe2bfd2e6ef1611B"
	//    ]
	//}
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

func ContractCall(privateKey, signerAddress, contractAddress, method, jsonString string,
	feeLimit, tAmount int64, tTokenID string, tTokenAmount int64) (string, error) {
	//privateKeyBytes, _ := hex.DecodeString(privateKey)

	conn := client.NewGrpcClient("grpc.shasta.trongrid.io:50051")
	tx, err := conn.TriggerContract(
		signerAddress,
		contractAddress,
		args[1],
		param,
		feeLimit,
		valueInt,
		tTokenID,
		tokenInt,
	)

	return "", nil
}
