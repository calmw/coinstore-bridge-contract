package main

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"encoding/pem"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
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
func ApproveSigTest() {
	client, err := ethclient.Dial("https://rpc.tantin.com")
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}

	fmt.Println(123)
	//设置交易参数
	fromAddress := common.HexToAddress("0x1933ccd14cafe561d862e5f35d0de75322a55412") // Owner
	toAddress := common.HexToAddress("0x2Bf013133aE838B6934B7F96fd43A10EE3FC3e18")   // USDT
	value := big.NewInt(0)
	gasLimit := uint64(21000)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress) // 获取nonce值
	if err != nil {
		log.Fatalf("Failed to get nonce: %v", err)
	}
	gasPrice, err := client.SuggestGasPrice(context.Background()) // 获取当前Gas价格建议
	if err != nil {
		log.Fatalf("Failed to suggest gas price: %v", err)
	}
	fmt.Println(123)
	data, err := InputData(common.HexToAddress("0x982D3ef9Db6c2cb4AaDfD609EB69264F382e5c5d"), big.NewInt(1))
	if err != nil {
		fmt.Println(11, err)
		return
	}
	//data, err := c.abi.Pack(method, params...)
	tx := types.NewTx(&types.LegacyTx{
		Nonce:    nonce,
		To:       &toAddress,
		Value:    value,
		Gas:      gasLimit,
		GasPrice: gasPrice,
		Data:     data,
	})
	//taskID := RandInt(100, 10000)
	taskID := 6687
	//txData:=fmt.Sprintf("%x", tx.Data())
	txData := "095ea7b3000000000000000000000000982d3ef9db6c2cb4aadfd609eb69264f382e5c5d0000000000000000000000000000000000000000000000000000000000000001"
	apiSecret := "ttbridge_9d8f7b6a5c4e3d2f1a0b9c8d7e6f5a4b3c2d1e0f"
	sigStr := fmt.Sprintf("%d%s%d%s%s",
		202502,
		strings.ToLower(fromAddress.String()),
		taskID,
		txData,
		apiSecret,
	)

	fmt.Println(sigStr)
	fingerprint := Keccak256([]byte(sigStr))
	fingerprint = Keccak256(fingerprint[:])

	fmt.Println(fmt.Sprintf("%x", fingerprint))
	return
	sigData, err := RequestWithPem("https://10.234.99.69:8088/signature/sign", SigDataPost{
		FromAddress: fromAddress.String(),
		TxData:      fmt.Sprintf("%x", tx.Data()),
		TaskID:      RandInt(100, 10000),
		ChainID:     202502,
		Fingerprint: fmt.Sprintf("%x", fingerprint),
	})
	postData := SigDataPost{
		FromAddress: fromAddress.String(),
		TxData:      fmt.Sprintf("%x", tx.Data()),
		TaskID:      RandInt(100, 10000),
		ChainID:     202502,
		Fingerprint: fmt.Sprintf("%x", fingerprint),
	}
	marshal, err := json.Marshal(postData)

	fmt.Println("post data:", string(marshal))
	fmt.Println("")
	fmt.Println(sigData, err)

	// 签名机返回的数据 sData
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

/// https://www.cnblogs.com/paulwhw/p/14015824.html

// SigDataPost
// {"fromAddress":"TTgY73yj5vzGM2HGHhVt7AR7avMW4jUx6n","txData":"c914f2fa2a214bf1c6bf80de09ecda76b9e7bc379c16b33e484f5ccf47e92b0a","taskId":2,"chainId":728126428,"fingerprint":"00b3a7749898d575ea92577f0f3ed3689355ea3dc276076fd2f716ead07d15b6"}
type SigDataPost struct {
	FromAddress string `json:"fromAddress"`
	TxData      string `json:"txData"`
	TaskID      int    `json:"taskId"`
	ChainID     int    `json:"chainId"`
	Fingerprint string `json:"fingerprint"`
}

func RequestWithPem(url string, data SigDataPost) ([]byte, error) {

	b, err := os.ReadFile("../sig1.pem")
	if err != nil {
		log.Fatal(err)
	}
	//pem.Decode(b)
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
	//keyString := string(pkey)
	//CertString := string(bytes)
	//fmt.Printf("Cert :\n %s \n Key:\n %s \n ", CertString, keyString)
	c, _ := tls.X509KeyPair(bytes, pkey)
	cfg := &tls.Config{
		Certificates:       []tls.Certificate{c},
		InsecureSkipVerify: true,
	}
	tr := &http.Transport{
		TLSClientConfig: cfg,
	}
	client := &http.Client{Transport: tr}
	strings.NewReader(fmt.Sprintf(`{"fromAddress":"%s","txData":"%s","taskId":%d,"chainId":%d,"fingerprint":"%s"}`,
		data.FromAddress, data.TxData, data.TaskID, data.ChainID, data.Fingerprint,
	))
	request, _ := http.NewRequest("POST", url, nil)
	request.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(request)
	if err != nil {
		return nil, err
	} else {
		dataa, _ := io.ReadAll(resp.Body)
		fmt.Println("~~~~~")
		fmt.Println(string(dataa))
		return dataa, nil
	}
}

func RandInt(min, max int) int {
	rand.NewSource(time.Now().UnixNano())
	randomNum := rand.Intn(max-min+1) + min
	return randomNum
}

func Keccak256(data []byte) [32]byte {
	return crypto.Keccak256Hash(data)
}
