package main

import (
	"bytes"
	"context"
	"crypto/ecdsa"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/status-im/keycard-go/hexutils"
	"log"
	"math/big"
	"strings"
)

func ApproveSigTest() error {
	client, err := ethclient.Dial("https://testrpcdex.tantin.com")
	//chainId:=202502
	chainId := 12302
	//client, err := ethclient.Dial("https://data-seed-prebsc-2-s3.bnbchain.org:8545")
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}
	//设置交易参数
	fromAddress := common.HexToAddress("0x1933ccd14cafe561d862e5f35d0de75322a55412") // Owner
	toAddress := common.HexToAddress("0x53F1BAA532710FC1FEE8a66433bE6c6fE823fCE9")   // USDT
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

	data, err := InputData(toAddress, big.NewInt(1))
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

	taskID := RandInt(100, 10000)
	// 编码数据
	var buf bytes.Buffer
	err = tx.EncodeRLP(&buf)
	txDataRlp := fmt.Sprintf("%x", buf.Bytes())
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
	res, err := RequestWithPem("https://18.141.210.154:8088/signature/sign", postData)
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
	// 发送交易
	txHash, err := SendTransactionFromRlpData(client, machineResp.Data)
	fmt.Println("SendTransactionFromRlpData:")
	fmt.Println(err, txHash)
	return err
}

func ApproveSigTest2() {
	apiSecret := "ttbridge_9d8f7b6a5c4e3d2f1a0b9c8d7e6f5a4b3c2d1e0f"
	sigStr := fmt.Sprintf("%d%s%d%s%s",
		728126428,
		strings.ToLower("TFBymbm7LrbRreGtByMPRD2HUyneKabsqb"),
		4791,
		"0a02eed722086c7bbb83f01730e040d8faedd6e7325af001081f12eb010a31747970652e676f6f676c65617069732e636f6d2f70726f746f636f6c2e54726967676572536d617274436f6e747261637412b5010a15413942fda93c573e2ce9e85b0bb00ba98a144f27f6121541d66285d10311a4ef8d788d68ef679a3d6bc41bf1228401cff3da8a00000000000000000000000000000000000000000000000000000000000000010000000000000000000000000000000000000000000000000000000000000023ac589789ed8c9d2c61f17b13369864b5f181e58eba230a6ee4ec4c3e7750cd1d038ad295563f130f1c44636c6d7dafb77241b984361c736df7b3b064f9e0d12f70ead2ead6e732900180c6868f01",
		apiSecret,
	)

	fingerprint := sha256.Sum256([]byte(sigStr))
	fingerprint = sha256.Sum256(fingerprint[:])
	fmt.Println(fmt.Sprintf("%x", fingerprint))

}

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

	return tx.Hash().String(), nil
	//tx sent: 0xc429e5f128387d224ba8bed6885e86525e14bfdc2eb24b5e9c3351a1176fd81f
	//sData := "0xab...."
	//sData = strings.TrimPrefix(sData, "0x")
	//sDataByte := hexutils.HexToBytes(sData)
	//newTx := types.Transaction{}
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

func GenerateEvmAccount() {
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		log.Fatal(err)
	}

	// 2. 获取私钥（hex编码）
	privateKeyBytes := crypto.FromECDSA(privateKey)
	fmt.Println("Private Key:", hex.EncodeToString(privateKeyBytes))

	// 3. 生成公钥
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("无法转换公钥")
	}

	// 4. 获取公钥（hex编码）
	publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)
	fmt.Println("Public Key:", hex.EncodeToString(publicKeyBytes))

	// 5. 根据公钥生成钱包地址
	address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()
	fmt.Println("Address:", address)
}
