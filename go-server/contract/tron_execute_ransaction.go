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
		c.SignTxForSending()
	case transaction.Ledger:
		c.HardwareSignTxForSending()
	}
	err := SignTxForSending(c, chainId, fromAddress, apiSecret)
	if err != nil {
		return err
	}
	c.SendSignedTx()
	c.TxConfirmation()
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
