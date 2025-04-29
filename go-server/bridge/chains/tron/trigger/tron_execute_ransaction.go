package trigger

import (
	"coinstore/bridge/chains/signature"
	"fmt"
	"github.com/calmw/tron-sdk/pkg/client/transaction"
	"google.golang.org/protobuf/proto"
)

func ExecuteTronTransaction(c *transaction.Controller, chainId int, apiSecret string) error {
	switch c.Behavior.SigningImpl {
	case transaction.Software:
		err := SignTxForSending(c, chainId, apiSecret)
		if err != nil {
			return err
		}
	case transaction.Ledger:
		c.HardwareSignTxForSending()
	}
	fmt.Println("Signature:", fmt.Sprintf("%x", c.Tx.Signature))
	c.SendSignedTx()
	c.TxConfirmation()
	return c.ExecutionError
}

func SignTxForSending(c *transaction.Controller, chainId int, apiSecret string) error {
	rawData, err := proto.Marshal(c.Tx.GetRawData())
	if err != nil {
		return err
	}
	sig, err := signature.SignAndSendTxTron(chainId, rawData, apiSecret)
	if err != nil {
		return err
	}
	c.Tx.Signature = append(c.Tx.Signature, sig)
	return nil
}
