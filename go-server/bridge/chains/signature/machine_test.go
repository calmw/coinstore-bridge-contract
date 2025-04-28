package signature

import (
	"fmt"
	"github.com/calmw/tron-sdk/pkg/client"
	"github.com/calmw/tron-sdk/pkg/client/transaction"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/proto"
	"os"
	"testing"
)

func TestSignAndSendTxTron(t *testing.T) {
	cli := client.NewGrpcClient("grpc.nile.trongrid.io:50051")
	err := cli.Start(grpc.WithInsecure())
	if err != nil {
		fmt.Println(err)
		return
	}
	triggerData := fmt.Sprintf("[{\"address\":\"%s\"},{\"uint256\":\"%s\"}]",
		"TVWmdb8o1X5mdEhZsiHB8wSuuHcicbEJYA", "10",
	)

	tx, err := cli.TriggerContract("TTgY73yj5vzGM2HGHhVt7AR7avMW4jUx6n", "TXYZopYRdj2D9XRtbG411XZZ3kM5VkAeBf", "approve(address,uint256)", triggerData, 300000000, 0, "", 0)
	if err != nil {
		fmt.Println(err)
		return
	}
	ctrlr := transaction.NewController(cli, nil, nil, tx.Transaction)

	apiSecret := os.Getenv("API_SECRET_SIG_MACHINE")
	if len(apiSecret) <= 0 {
		apiSecret = "ttbridge_9d8f7b6a5c4e3d2f1a0b9c8d7e6f5a4b3c2d1e0f"
	}
	if err = ExecuteTronTransaction(ctrlr, 728126428, apiSecret); err != nil {
		fmt.Println(err)
		return
	}
}

func ExecuteTronTransaction(c *transaction.Controller, chainId int, apiSecret string) error {
	switch c.Behavior.SigningImpl {
	case transaction.Software:
		//c.SignTxForSending()
		fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~1 ", c.Tx.Signature)
		err := SignTxForSending(c, chainId, apiSecret)
		if err != nil {
			return err
		}
		fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~2 ", c.Tx.Signature)
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
	sig, err := SignAndSendTxTron(chainId, rawData, apiSecret)
	if err != nil {
		return err
	}
	c.Tx.Signature = append(c.Tx.Signature, sig)
	return nil
}
