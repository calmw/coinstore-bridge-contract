package contract

import (
	"coinstore/bridge/chains/signature"
	"fmt"
	"github.com/calmw/tron-sdk/pkg/client/transaction"
	"google.golang.org/protobuf/proto"
)

func ExecuteTronTransaction(c *transaction.Controller, chainId int, fromAddress, apiSecret string) error {
	fmt.Println(c.Behavior.SigningImpl, "~~~~~~~~~")
	switch c.Behavior.SigningImpl {
	case transaction.Software:
		fmt.Println("签名1：", fmt.Sprintf("%x", c.Tx.Signature))
		c.SignTxForSending()
		fmt.Println("签名2：", fmt.Sprintf("%x", c.Tx.Signature))
		//err := SignTxForSending(c, chainId, fromAddress, apiSecret)
		//if err != nil {
		//	fmt.Println("从签名机获取签名错误", err)
		//	return err
		//}
	case transaction.Ledger:
		c.HardwareSignTxForSending()
	}
	//fmt.Println("签名2：", fmt.Sprintf("%x", c.Tx.Signature))
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
	sig, err := signature.SignAndSendTxTron(chainId, fromAddress, rawData, apiSecret)
	if err != nil {
		return err
	}
	c.Tx.Signature = append(c.Tx.Signature, sig)
	return nil
}
