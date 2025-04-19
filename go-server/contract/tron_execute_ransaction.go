package contract

import (
	"fmt"
	"github.com/calmw/tron-sdk/pkg/client/transaction"
	"google.golang.org/protobuf/proto"
)

func ExecuteTronTransaction(c *transaction.Controller) error {
	fmt.Println(c.Behavior.SigningImpl, "~~~~~~~~~")
	switch c.Behavior.SigningImpl {
	case transaction.Software:
		c.SignTxForSending()
	case transaction.Ledger:
		c.HardwareSignTxForSending()
	}
	c.SendSignedTx()
	c.TxConfirmation()
	return c.ExecutionError
	return nil
}

func SignTxForSending(c *transaction.Controller) error {

	///
	rawData, err := proto.Marshal(c.Tx.GetRawData())
	if err != nil {
		return err
	}
	// 请求 让我打他

	///
	if c.ExecutionError != nil {
		return
	}
	signedTransaction, err :=
		c.Sender.Ks.SignTx(*c.Sender.Account, c.Tx)
	if err != nil {
		c.ExecutionError = err
		return
	}
	c.Tx = signedTransaction
}
