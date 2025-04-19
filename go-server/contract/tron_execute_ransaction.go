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
		//c.SignTxForSending()
		err := SignTxForSending(c, chainId, fromAddress, apiSecret)
		fmt.Println("!!!!!!!!!!!!", err)
		if err != nil {
			return err
		}
	case transaction.Ledger:
		c.HardwareSignTxForSending()
	}
	err := SignTxForSending(c, chainId, fromAddress, apiSecret)
	fmt.Println("!!!!!!!!!!!!~~", err)
	fmt.Println("Signature !! 1", c.Tx.Signature)
	if err != nil {
		return err
	}
	c.SendSignedTx()
	c.TxConfirmation()
	fmt.Println("!!!!!!!!!!!!~~###", c.ExecutionError)
	return c.ExecutionError
}

func SignTxForSending(c *transaction.Controller, chainId int, fromAddress, apiSecret string) error {
	rawData, err := proto.Marshal(c.Tx.GetRawData())
	fmt.Println("rawData !! 1", fmt.Sprintf("%x", rawData))
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
