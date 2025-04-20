package contract

import (
	"coinstore/bridge/chains/signature"
	"crypto/sha256"
	"fmt"
	"github.com/calmw/tron-sdk/pkg/client/transaction"
	"google.golang.org/protobuf/proto"
)

func ExecuteTronTransaction(c *transaction.Controller, chainId int, fromAddress, apiSecret string) error {
	fmt.Println(c.Behavior.SigningImpl, "~~~~~~~~~")
	switch c.Behavior.SigningImpl {
	case transaction.Software:
		fmt.Println("签名1：", fmt.Sprintf("%x", c.Tx.String()))
		fmt.Println("签名1：", fmt.Sprintf("%x", c.Tx.Signature))
		c.SignTxForSending()
		//fmt.Println("签名2：", fmt.Sprintf("%x", c.Tx.Signature))
		err := SignTxForSending(c, chainId, fromAddress, apiSecret)
		if err != nil {
			fmt.Println("从签名机获取签名错误", err)
			return err
		}
	case transaction.Ledger:
		c.HardwareSignTxForSending()
	}
	fmt.Println("签名2：", fmt.Sprintf("%x", c.Tx.Signature))
	c.SendSignedTx()
	c.TxConfirmation()
	fmt.Println("!!!!!!!!!!!!~~###", c.ExecutionError)
	return c.ExecutionError
}

func SignTxForSending(c *transaction.Controller, chainId int, fromAddress, apiSecret string) error {
	rawData, err := proto.Marshal(c.Tx.GetRawData())
	if err != nil {
		return err
	}
	h256h := sha256.New()
	h256h.Write(rawData)
	hash := h256h.Sum(nil)
	sig, err := signature.SignAndSendTxTron(chainId, fromAddress, hash, apiSecret)
	if err != nil {
		return err
	}
	c.Tx.Signature = append(c.Tx.Signature, sig)
	return nil
}
