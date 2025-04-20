package contract

import (
	"coinstore/bridge/chains/signature"
	"coinstore/utils"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"github.com/btcsuite/btcd/btcec/v2"
	"github.com/calmw/tron-sdk/pkg/client"
	"github.com/calmw/tron-sdk/pkg/client/transaction"
	"github.com/calmw/tron-sdk/pkg/proto/api"
	"github.com/ethereum/go-ethereum/crypto"
	"google.golang.org/protobuf/proto"
	"os"
)

func ExecuteTronTransaction(c *transaction.Controller, chainId int, fromAddress, apiSecret string) error {
	fmt.Println(c.Behavior.SigningImpl, "~~~~~~~~~")
	switch c.Behavior.SigningImpl {
	case transaction.Software:
		//fmt.Println("签名1：", fmt.Sprintf("%x", c.Tx.String()))
		//fmt.Println("签名1：", fmt.Sprintf("%x", c.Tx.Signature))
		//c.SignTxForSending()
		//fmt.Println("签名2：", fmt.Sprintf("%x", c.Tx.Signature))
		err := SignTxForSending(c, chainId, fromAddress, apiSecret)
		if err != nil {
			fmt.Println("从签名机获取签名错误", err)
			return err
		}
	case transaction.Ledger:
		c.HardwareSignTxForSending()
	}
	fmt.Println("签名：", fmt.Sprintf("%x", c.Tx.Signature))
	c.SendSignedTx()
	c.TxConfirmation()
	fmt.Println("交易结果：", c.ExecutionError)
	return c.ExecutionError
}

func SignTxForSending(c *transaction.Controller, chainId int, fromAddress, apiSecret string) error {

	rawData, err := proto.Marshal(c.Tx.GetRawData())
	if err != nil {
		return err
	}
	fmt.Println("rawData:")
	fmt.Println(fmt.Sprintf("%x", rawData))
	h256h := sha256.New()
	h256h.Write(rawData)
	hash := h256h.Sum(nil)
	fmt.Println("hash:")
	fmt.Println(fmt.Sprintf("%x", hash))
	//
	//err := TestTx(c.Client,c.Tx,hash)

	fmt.Println(err)
	return nil

	sig, err := signature.SignAndSendTxTron(chainId, fromAddress, hash, apiSecret)
	if err != nil {
		return err
	}
	c.Tx.Signature = append(c.Tx.Signature, sig)
	return nil
}

func TestTx(c *client.GrpcClient, tx *api.TransactionExtention, hash []byte) error {
	pk := os.Getenv("COIN_STORE_BRIDGE_TRON")

	privateKeyStr := utils.ThreeDesDecrypt("gZIMfo6LJm6GYXdClPhIMfo6", pk)

	privateKeyBytes, _ := hex.DecodeString(privateKeyStr)
	sk, _ := btcec.PrivKeyFromBytes(privateKeyBytes)
	sig, err := crypto.Sign(hash, sk.ToECDSA())
	if err != nil {
		return err
	}
	fmt.Println(fmt.Sprintf("sssss1:%x", tx.Transaction.Signature))
	tx.Transaction.Signature = append(tx.Transaction.Signature, sig)
	fmt.Println(fmt.Sprintf("sssss2:%x", tx.Transaction.Signature))
	//res, err := c.Broadcast(tx.Transaction)
	//if err != nil || !res.Result {
	//	return fmt.Errorf("broadcast error %v", err)
	//}
	//txHash := strings.ToLower(hexutils.BytesToHex(tx.Txid))
	//fmt.Println("txHash:", txHash)
	return nil
}
