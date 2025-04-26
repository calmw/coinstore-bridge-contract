package main

import (
	"bytes"
	"context"
	"crypto/sha256"
	"crypto/tls"
	"encoding/json"
	"encoding/pem"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/status-im/keycard-go/hexutils"
	"io"
	"log"
	"math/big"
	"math/rand"
	"net/http"
	"os"
	"strings"
	"time"
)

func main() {
	ApproveSigTest()
}

type SigDataPost struct {
	FromAddress string `json:"fromAddress"`
	TxData      string `json:"txData"`
	TaskID      int    `json:"taskId"`
	ChainID     int    `json:"chainId"`
	Fingerprint string `json:"fingerprint"`
}

type MachineResp struct {
	Code    int    `json:"code"`
	Data    string `json:"data"`
	Message string `json:"message"`
}

func ApproveSigTest() error {
	//client, err := ethclient.Dial("https://rpc.tantin.com")
	//chainId:=202502
	chainId := 97
	client, err := ethclient.Dial("https://data-seed-prebsc-2-s3.bnbchain.org:8545")
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}
	//设置交易参数
	fromAddress := common.HexToAddress("0x1933ccd14cafe561d862e5f35d0de75322a55412") // Owner
	//toAddress := common.HexToAddress("0x2Bf013133aE838B6934B7F96fd43A10EE3FC3e18")   // USDT
	toAddress := common.HexToAddress("0x4b62Da623b5aAfE4BAEe909e1fBB321b96887B3D") // USDT
	value := big.NewInt(0)
	//gasLimit := uint64(21000)
	gasLimit := uint64(50000)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatalf("Failed to get nonce: %v", err)
		return err
	}
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatalf("Failed to suggest gas price: %v", err)
		return err
	}

	data, err := InputData(common.HexToAddress("0x982D3ef9Db6c2cb4AaDfD609EB69264F382e5c5d"), big.NewInt(1))
	if err != nil {
		log.Fatalf("Failed to suggest gas price2 : %v", err)
		return err
	}
	//gasPrice = gasPrice.Mul(gasPrice, big.NewInt(2))
	tx := types.NewTx(&types.LegacyTx{
		Nonce:    nonce,
		To:       &toAddress,
		Value:    value,
		Gas:      gasLimit,
		GasPrice: gasPrice,
		Data:     data,
	})
	marshalJSON, err := tx.MarshalJSON()
	if err != nil {
		return err
	}
	fmt.Println("txDataJson")
	fmt.Println(string(marshalJSON))
	taskID := RandInt(100, 10000)
	// 编码数据
	var buf bytes.Buffer
	err = tx.EncodeRLP(&buf)
	txDataRlp := fmt.Sprintf("%x", buf.Bytes())
	fmt.Println("txDataRlp")
	fmt.Println(txDataRlp)
	apiSecret := "ttbridge_9d8f7b6a5c4e3d2f1a0b9c8d7e6f5a4b3c2d1e0f"
	sigStr := fmt.Sprintf("%d%s%d%s%s",
		chainId,
		strings.ToLower(fromAddress.String()),
		taskID,
		txDataRlp,
		apiSecret,
	)

	fingerprint := sha256.Sum256([]byte(sigStr))
	fingerprint = sha256.Sum256(fingerprint[:])
	postData := SigDataPost{
		FromAddress: fromAddress.String(),
		TxData:      txDataRlp,
		TaskID:      taskID,
		ChainID:     chainId,
		Fingerprint: fmt.Sprintf("%x", fingerprint),
	}
	res, err := RequestWithPem("https://nlb.self.tantin.com:8088/signature/sign", postData)
	fmt.Println("RequestWithPem error:", err)
	var machineResp MachineResp
	err = json.Unmarshal(res, &machineResp)
	if err != nil {
		return err
	}
	if machineResp.Code != 200 {
		return errors.New("signature machine error")
	}
	// 发送交易
	txHash, err := SendTransactionFromRlpData(client, machineResp.Data)
	fmt.Println("SendTransactionFromRlpData:")
	fmt.Println(err, txHash)
	return err
}

// SendTransactionFromRlpData rawTx rlpDataHex
func SendTransactionFromRlpData(cli *ethclient.Client, rlpDataHex string) (string, error) {
	fmt.Println(rlpDataHex)
	fmt.Println(strings.TrimPrefix(rlpDataHex, "0x"))
	rawTxBytes := hexutils.HexToBytes(strings.TrimPrefix(rlpDataHex, "0x"))

	//bb := hexutils.HexToBytes("f8ac808504a817c8008261a8942bf013133ae838b6934b7f96fd43a10ee3fc3e1880b844095ea7b3000000000000000000000000982d3ef9db6c2cb4aadfd609eb69264f382e5c5d000000000000000000000000000000000000000000000000000000000000000183062e2fa00a41d1514e447d8a71590f824b27dcf4aabd8387eb5975e540e84f163ccef1a2a04b2e16da5994099e1c82043109f6c608cc03ee76595b55c2e042301f49655003")
	tx := new(types.Transaction)
	//err := rlp.DecodeBytes(bb, &tx)
	err := rlp.DecodeBytes(rawTxBytes, &tx)
	if err != nil {
		return "", err
	}
	err = cli.SendTransaction(context.Background(), tx)
	if err != nil {
		return "", err
	}

	fmt.Printf("tx sent: %s", tx.Hash().Hex())
	return tx.Hash().String(), nil

	// tx sent: 0xc429e5f128387d224ba8bed6885e86525e14bfdc2eb24b5e9c3351a1176fd81f
	//sData := "0xab...."
	//sData = strings.TrimPrefix(sData, "0x")
	//sDataByte := hexutils.HexToBytes(sData)
	//newTx := types.Transaction{}
	//
	//err := newTx.UnmarshalJSON(sDataByte)
	//if err != nil {
	//	return
	//}
	//// 发送交易到链上
	//err = client.SendTransaction(context.Background(), &newTx)
	//if err != nil {
	//	log.Fatalf("Failed to send transaction: %v", err)
	//}
	//fmt.Println("Transaction sent:", signedTx.Hash().Hex()) // 打印交易哈希

}

func InputData(spender common.Address, value *big.Int) ([]byte, error) {
	abiJson := `[{
    "inputs": [
      {
        "internalType": "address",
        "name": "spender",
        "type": "address"
      },
      {
        "internalType": "uint256",
        "name": "value",
        "type": "uint256"
      }
    ],
    "name": "approve",
    "outputs": [
      {
        "internalType": "bool",
        "name": "",
        "type": "bool"
      }
    ],
    "stateMutability": "nonpayable",
    "type": "function"
  }]`
	contractAbi, _ := abi.JSON(strings.NewReader(abiJson))
	return contractAbi.Pack("approve",
		spender,
		value,
	)
}

func RequestWithPem(url string, data SigDataPost) ([]byte, error) {

	b, err := os.ReadFile("../sig1.pem")
	if err != nil {
		log.Fatal(err)
	}
	var pemBlocks []*pem.Block
	var v *pem.Block
	var pkey []byte
	for {
		v, b = pem.Decode(b)
		if v == nil {
			break
		}
		if v.Type == "PRIVATE KEY" {
			pkey = pem.EncodeToMemory(v)
		} else {
			pemBlocks = append(pemBlocks, v)
		}
	}

	bytes := pem.EncodeToMemory(pemBlocks[0])
	c, _ := tls.X509KeyPair(bytes, pkey)
	cfg := &tls.Config{
		Certificates:       []tls.Certificate{c},
		InsecureSkipVerify: true,
	}
	tr := &http.Transport{
		TLSClientConfig: cfg,
	}
	client := &http.Client{Transport: tr}
	postData := fmt.Sprintf(`{"fromAddress":"%s","txData":"%s","taskId":%d,"chainId":%d,"fingerprint":"%s"}`,
		data.FromAddress, data.TxData, data.TaskID, data.ChainID, data.Fingerprint,
	)
	fmt.Println("post data:")
	fmt.Println(postData)
	strings.NewReader(postData)
	d := strings.NewReader(postData)
	request, _ := http.NewRequest("POST", url, d)
	request.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(request)
	if err != nil {
		return nil, err
	} else {
		resData, _ := io.ReadAll(resp.Body)
		fmt.Println("post response:")
		fmt.Println(string(resData))
		return resData, nil
	}
}

func RandInt(min, max int) int {
	rand.NewSource(time.Now().UnixNano())
	randomNum := rand.Intn(max-min+1) + min
	return randomNum
}
