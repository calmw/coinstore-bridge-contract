package main

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/btcsuite/btcd/btcec/v2"
	"github.com/btcsuite/btcd/btcec/v2/ecdsa"
	"github.com/calmw/tron-sdk/pkg/client"
	"github.com/status-im/keycard-go/hexutils"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/proto"
	"os"
	"strings"
)

func TransferCoin() (string, error) {
	privateKey := os.Getenv("COINSTORE_BRIDGE_TRON_LOCAL")
	fromAddress := "TFBymbm7LrbRreGtByMPRD2HUyneKabsqb"
	toAddress := "TEwX7WKNQqsRpxd6KyHHQPMMigLg9c258y"
	amount := int64(1)

	c := client.NewGrpcClient("grpc.nile.trongrid.io:50051")
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
	hash := sha256.Sum256(rawData)
	privateKeyBytes, err := hex.DecodeString(privateKey)
	if err != nil {
		return "", err
	}
	pk, _ := btcec.PrivKeyFromBytes(privateKeyBytes)
	sigCompact := ecdsa.SignCompact(pk, hash[:], true)
	recoveryID := int(sigCompact[0]) - 27 // header减去27就是recoveryID
	r := sigCompact[1:33]                 // r部分
	s := sigCompact[33:65]                // s部分
	// 使用 v 作为标识符（27或28）
	v := byte(recoveryID + 27)
	// 拼接最终的签名（r + s + v）
	finalSig := append(r, s...)
	finalSig = append(finalSig, v)
	tx.Transaction.Signature = append(tx.Transaction.Signature, finalSig)
	res, err := c.Broadcast(tx.Transaction)
	fmt.Println(err)
	if err != nil || !res.Result {
		return "", errors.New("broadcast error")
	}
	fmt.Println(err, res)
	txHash := strings.ToLower(hexutils.BytesToHex(tx.Txid))
	fmt.Println(txHash)
	return txHash, nil
}

func TronTest() error {
	//privateKey := os.Getenv("COINSTORE_BRIDGE_TRON_LOCAL")
	//fromAddress := "TFBymbm7LrbRreGtByMPRD2HUyneKabsqb"
	fromAddress := "TTgY73yj5vzGM2HGHhVt7AR7avMW4jUx6n"
	privateKey := "4d69dca2ebc32e9d448a1c6d6fa3da61fc537077279c3ab6c6702adfe70dd8a2"
	//fromAddress := "TYMHM8fjbMJDDeyEXnmAntq5t4afzv4T8M"
	toAddress := "TEwX7WKNQqsRpxd6KyHHQPMMigLg9c258y"
	amount := int64(1)

	cli := client.NewGrpcClient("grpc.nile.trongrid.io:50051")
	err := cli.Start(grpc.WithInsecure())
	if err != nil {
		return err
	}
	tx, err := cli.Transfer(fromAddress, toAddress, amount)
	if err != nil {
		return err
	}
	rawData, err := proto.Marshal(tx.Transaction.GetRawData())
	if err != nil {
		return err
	}
	hash := sha256.Sum256(rawData)

	///
	privateKeyBytes, err := hex.DecodeString(privateKey)
	if err != nil {
		return err
	}
	pk, _ := btcec.PrivKeyFromBytes(privateKeyBytes)
	sigCompact := ecdsa.SignCompact(pk, hash[:], true)
	recoveryID := int(sigCompact[0]) - 27 // header减去27就是recoveryID
	r := sigCompact[1:33]                 // r部分
	s := sigCompact[33:65]                // s部分
	// 使用 v 作为标识符（27或28）
	v := byte(recoveryID + 27)
	// 拼接最终的签名（r + s + v）
	finalSig := append(r, s...)
	finalSig = append(finalSig, v)

	signature := finalSig
	//fmt.Println("本地签名 0 ", sigCompact[0])
	fmt.Println("本地签名长度", len(signature))
	fmt.Println("本地签名", signature)
	fmt.Println("本地签名 ", fmt.Sprintf("%x", signature))

	fromAddress = "TTgY73yj5vzGM2HGHhVt7AR7avMW4jUx6n"
	data := hexutils.BytesToHex(hash[:])
	taskID := RandInt(100, 1000)
	chianId := 728126428
	apiSecret := "ttbridge_9d8f7b6a5c4e3d2f1a0b9c8d7e6f5a4b3c2d1e0f"
	sigStr := fmt.Sprintf("%d%s%d%s%s",
		chianId,
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
		ChainID:     chianId,
		Fingerprint: fmt.Sprintf("%x", fingerprint),
	}
	res, err := RequestWithPem("https://18.141.210.154:8088/signature/sign", postData)
	//res, err := RequestWithPem("https://47.129.133.232:8088/signature/sign", postData)
	if err != nil {
		return err
	}
	var machineResp MachineResp
	err = json.Unmarshal(res, &machineResp)
	if err != nil {
		return err
	}
	if machineResp.Code != 200 {
		return errors.New("signature machine error")
	}
	signature2 := hexutils.HexToBytes(machineResp.Data)
	fmt.Println("签名机签名长度", len(signature2))
	fmt.Println("签名机签名 ", machineResp.Data)
	fmt.Println("签名机签名 ", signature2)
	signature2[64] = signature2[64] + 4
	fmt.Println("签名机签名最后一位+4:")
	fmt.Println("加4后签名机签名长度", len(signature2))
	fmt.Println("加4后签名机签名 ", fmt.Sprintf("%x", signature2))
	fmt.Println("加4后签名机签名 ", signature2)
	// 发送交易
	tx.Transaction.Signature = append(tx.Transaction.Signature, signature2)
	//tx.Transaction.Signature =  hexutils.HexToBytes(machineResp.Data)
	result, err := cli.Broadcast(tx.Transaction)
	if err != nil || !result.Result {
		return fmt.Errorf("broadcast error %v", err)
	}
	txHash := strings.ToLower(hexutils.BytesToHex(tx.Txid))
	fmt.Println(txHash)
	return nil
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
