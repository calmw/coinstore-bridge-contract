package main

import (
	"crypto/sha256"
	"crypto/tls"
	"encoding/hex"
	"encoding/json"
	"encoding/pem"
	"errors"
	"fmt"
	"github.com/btcsuite/btcd/btcec/v2"
	"github.com/calmw/tron-sdk/pkg/client"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/status-im/keycard-go/hexutils"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/proto"
	"io"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strings"
	"time"
)

func main() {
	//TronTest()
	TransferCoin("af056e353eaddc842029822cc4c7285ef379880eaae25a1f8d8a16da5f41ff2b", "TFBymbm7LrbRreGtByMPRD2HUyneKabsqb", "TEwX7WKNQqsRpxd6KyHHQPMMigLg9c258y", 2)
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
	fmt.Println(err, res)
	txHash := strings.ToLower(hexutils.BytesToHex(tx.Txid))
	fmt.Println(txHash)
	return txHash, nil
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

//func ApproveSigTest() error {
//	client, err := ethclient.Dial("https://testrpcdex.tantin.com")
//	//chainId:=202502
//	chainId := 12302
//	//client, err := ethclient.Dial("https://data-seed-prebsc-2-s3.bnbchain.org:8545")
//	if err != nil {
//		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
//	}
//	//设置交易参数
//	fromAddress := common.HexToAddress("0x1933ccd14cafe561d862e5f35d0de75322a55412") // Owner
//	toAddress := common.HexToAddress("0x53F1BAA532710FC1FEE8a66433bE6c6fE823fCE9")   // USDT
//	value := big.NewInt(0)
//	//gasLimit := uint64(21000)
//	gasLimit := uint64(50000)
//	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
//	if err != nil {
//		log.Fatalf("Failed to get nonce: %v", err)
//		return err
//	}
//	gasPrice, err := client.SuggestGasPrice(context.Background())
//	if err != nil {
//		log.Fatalf("Failed to suggest gas price: %v", err)
//		return err
//	}
//
//	data, err := InputData(toAddress, big.NewInt(1))
//	if err != nil {
//		log.Fatalf("Failed to suggest gas price2 : %v", err)
//		return err
//	}
//	//gasPrice = gasPrice.Mul(gasPrice, big.NewInt(2))
//	tx := types.NewTx(&types.LegacyTx{
//		Nonce:    nonce,
//		To:       &toAddress,
//		Value:    value,
//		Gas:      gasLimit,
//		GasPrice: gasPrice,
//		Data:     data,
//	})
//	marshalJSON, err := tx.MarshalJSON()
//	if err != nil {
//		return err
//	}
//	fmt.Println("txDataJson")
//	fmt.Println(string(marshalJSON))
//	taskID := RandInt(100, 10000)
//	// 编码数据
//	var buf bytes.Buffer
//	err = tx.EncodeRLP(&buf)
//	txDataRlp := fmt.Sprintf("%x", buf.Bytes())
//	fmt.Println("txDataRlp")
//	fmt.Println(txDataRlp)
//	apiSecret := "ttbridge_9d8f7b6a5c4e3d2f1a0b9c8d7e6f5a4b3c2d1e0f"
//	sigStr := fmt.Sprintf("%d%s%d%s%s",
//		chainId,
//		strings.ToLower(fromAddress.String()),
//		taskID,
//		txDataRlp,
//		apiSecret,
//	)
//
//	fingerprint := sha256.Sum256([]byte(sigStr))
//	fingerprint = sha256.Sum256(fingerprint[:])
//	postData := SigDataPost{
//		FromAddress: fromAddress.String(),
//		TxData:      txDataRlp,
//		TaskID:      taskID,
//		ChainID:     chainId,
//		Fingerprint: fmt.Sprintf("%x", fingerprint),
//	}
//	res, err := RequestWithPem("https://18.141.210.154:8088/signature/sign", postData)
//	fmt.Println("RequestWithPem error:", err)
//	var machineResp MachineResp
//	err = json.Unmarshal(res, &machineResp)
//	if err != nil {
//		return err
//	}
//	if machineResp.Code != 200 {
//		return errors.New("signature machine error")
//	}
//	// 发送交易
//	txHash, err := SendTransactionFromRlpData(client, machineResp.Data)
//	fmt.Println("SendTransactionFromRlpData:")
//	fmt.Println(err, txHash)
//	return err
//}
//
//func ApproveSigTest2() {
//	apiSecret := "ttbridge_9d8f7b6a5c4e3d2f1a0b9c8d7e6f5a4b3c2d1e0f"
//	sigStr := fmt.Sprintf("%d%s%d%s%s",
//		728126428,
//		strings.ToLower("TFBymbm7LrbRreGtByMPRD2HUyneKabsqb"),
//		4791,
//		"0a02eed722086c7bbb83f01730e040d8faedd6e7325af001081f12eb010a31747970652e676f6f676c65617069732e636f6d2f70726f746f636f6c2e54726967676572536d617274436f6e747261637412b5010a15413942fda93c573e2ce9e85b0bb00ba98a144f27f6121541d66285d10311a4ef8d788d68ef679a3d6bc41bf1228401cff3da8a00000000000000000000000000000000000000000000000000000000000000010000000000000000000000000000000000000000000000000000000000000023ac589789ed8c9d2c61f17b13369864b5f181e58eba230a6ee4ec4c3e7750cd1d038ad295563f130f1c44636c6d7dafb77241b984361c736df7b3b064f9e0d12f70ead2ead6e732900180c6868f01",
//		apiSecret,
//	)
//
//	fingerprint := sha256.Sum256([]byte(sigStr))
//	fingerprint = sha256.Sum256(fingerprint[:])
//	fmt.Println(fmt.Sprintf("%x", fingerprint))
//
//}

// SendTransactionFromRlpData rawTx rlpDataHex
//func SendTransactionFromRlpData(cli *ethclient.Client, rlpDataHex string) (string, error) {
//	fmt.Println(rlpDataHex)
//	fmt.Println(strings.TrimPrefix(rlpDataHex, "0x"))
//	rawTxBytes := hexutils.HexToBytes(strings.TrimPrefix(rlpDataHex, "0x"))
//
//	//bb := hexutils.HexToBytes("f8ac808504a817c8008261a8942bf013133ae838b6934b7f96fd43a10ee3fc3e1880b844095ea7b3000000000000000000000000982d3ef9db6c2cb4aadfd609eb69264f382e5c5d000000000000000000000000000000000000000000000000000000000000000183062e2fa00a41d1514e447d8a71590f824b27dcf4aabd8387eb5975e540e84f163ccef1a2a04b2e16da5994099e1c82043109f6c608cc03ee76595b55c2e042301f49655003")
//	tx := new(types.Transaction)
//	//err := rlp.DecodeBytes(bb, &tx)
//	err := rlp.DecodeBytes(rawTxBytes, &tx)
//	if err != nil {
//		return "", err
//	}
//	err = cli.SendTransaction(context.Background(), tx)
//	if err != nil {
//		return "", err
//	}
//
//	fmt.Printf("tx sent: %s", tx.Hash().Hex())
//	return tx.Hash().String(), nil
//	//tx sent: 0xc429e5f128387d224ba8bed6885e86525e14bfdc2eb24b5e9c3351a1176fd81f
//	//sData := "0xab...."
//	//sData = strings.TrimPrefix(sData, "0x")
//	//sDataByte := hexutils.HexToBytes(sData)
//	//newTx := types.Transaction{}
//	//err := newTx.UnmarshalJSON(sDataByte)
//	//if err != nil {
//	//	return
//	//}
//	//// 发送交易到链上
//	//err = client.SendTransaction(context.Background(), &newTx)
//	//if err != nil {
//	//	log.Fatalf("Failed to send transaction: %v", err)
//	//}
//	//fmt.Println("Transaction sent:", signedTx.Hash().Hex()) // 打印交易哈希
//
//}
//
//func InputData(spender common.Address, value *big.Int) ([]byte, error) {
//	abiJson := `[{
//    "inputs": [
//      {
//        "internalType": "address",
//        "name": "spender",
//        "type": "address"
//      },
//      {
//        "internalType": "uint256",
//        "name": "value",
//        "type": "uint256"
//      }
//    ],
//    "name": "approve",
//    "outputs": [
//      {
//        "internalType": "bool",
//        "name": "",
//        "type": "bool"
//      }
//    ],
//    "stateMutability": "nonpayable",
//    "type": "function"
//  }]`
//	contractAbi, _ := abi.JSON(strings.NewReader(abiJson))
//	return contractAbi.Pack("approve",
//		spender,
//		value,
//	)
//}

func RequestWithPem(url string, data SigDataPost) ([]byte, error) {

	b, err := os.ReadFile("./sig1.pem")
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

func TronTest() error {
	cli := client.NewGrpcClient("grpc.nile.trongrid.io:50051")
	err := cli.Start(grpc.WithInsecure())
	if err != nil {
		fmt.Println(err)
		return err
	}
	fromAddress := "TTgY73yj5vzGM2HGHhVt7AR7avMW4jUx6n"
	//fromAddress := "TSARBFH6PW6jEuf8chd1DxZGW6JEmHuv6g"
	//fromAddress := "TPA22F33SaBwZAVkYw1GsJ7uz2CwFMWBAx"
	toAddress := "TEwX7WKNQqsRpxd6KyHHQPMMigLg9c258y"
	//chainId := 3448148188
	chainId := 728126428
	taskID := RandInt(100, 1000)
	tx, err := cli.Transfer(fromAddress, toAddress, 2)
	if err != nil {
		fmt.Println(err)
		return err
	}
	rawData, err := proto.Marshal(tx.Transaction.GetRawData())
	if err != nil {
		fmt.Println(err)
		return err
	}
	h256h := sha256.New()
	h256h.Write(rawData)
	hash := h256h.Sum(nil)

	data := hexutils.BytesToHex(hash)
	apiSecret := "ttbridge_9d8f7b6a5c4e3d2f1a0b9c8d7e6f5a4b3c2d1e0f"
	sigStr := fmt.Sprintf("%d%s%d%s%s",
		chainId,
		strings.ToLower(fromAddress),
		taskID,
		data,
		apiSecret,
	)

	fingerprint := sha256.Sum256([]byte(sigStr))
	fingerprint = sha256.Sum256(fingerprint[:])
	postData := SigDataPost{
		FromAddress: fromAddress,
		TxData:      data,
		TaskID:      taskID,
		ChainID:     chainId,
		Fingerprint: fmt.Sprintf("%x", fingerprint),
	}
	res, err := RequestWithPem("https://18.141.210.154:8088/signature/sign", postData)
	fmt.Println("RequestWithPem error:", err)
	var machineResp MachineResp
	err = json.Unmarshal(res, &machineResp)
	if err != nil {
		return err
	}
	if machineResp.Code != 200 {
		return errors.New("signature machine error")
	}
	fmt.Println(len(hexutils.HexToBytes(machineResp.Data)), "~~~")
	// 发送交易
	tx.Transaction.Signature = append(tx.Transaction.Signature, hexutils.HexToBytes(machineResp.Data))
	//tx.Transaction.Signature =  hexutils.HexToBytes(machineResp.Data)
	result, err := cli.Broadcast(tx.Transaction)
	fmt.Println(result, err)
	return nil
}
