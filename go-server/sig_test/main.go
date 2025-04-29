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

func main() {
	//tx, err := TransferCoin()
	//fmt.Println(tx, err)

	err := TronTest()
	fmt.Println(err)
}

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
	privateKey := os.Getenv("COINSTORE_BRIDGE_TRON_LOCAL")
	fromAddress := "TFBymbm7LrbRreGtByMPRD2HUyneKabsqb"
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
	fmt.Println("本地签名长度", len(signature))
	fmt.Println("本地签名 ", fmt.Sprintf("%x", signature))
	///

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
	//signature = hexutils.HexToBytes(machineResp.Data)
	fmt.Println("签名机签名长度", len(signature))
	fmt.Println("签名机签名 ", machineResp.Data)
	// 发送交易
	tx.Transaction.Signature = append(tx.Transaction.Signature, signature)
	//tx.Transaction.Signature =  hexutils.HexToBytes(machineResp.Data)
	result, err := cli.Broadcast(tx.Transaction)
	if err != nil || !result.Result {
		return fmt.Errorf("broadcast error %v", err)
	}
	txHash := strings.ToLower(hexutils.BytesToHex(tx.Txid))
	fmt.Println(txHash)
	return nil
}
