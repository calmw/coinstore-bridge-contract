package tron

import "C"
import (
	"fmt"
	"github.com/fbsobreira/gotron-sdk/pkg/client/transaction"
)

func ExecuteTransaction(c *transaction.Controller) error {
	fmt.Println(c.Behavior.SigningImpl, "~~~~~~~~~")
	switch c.Behavior.SigningImpl {
	case transaction.Software:
		C.signTxForSending()
	case transaction.Ledger:
		C.hardwareSignTxForSending()
	}
	C.sendSignedTx()
	C.txConfirmation()
	return C.executionError
}
