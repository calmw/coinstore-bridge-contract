package contract

import (
	"coinstore/binding"
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/status-im/keycard-go/hexutils"
	"log"
	"math/big"
	"strings"
	"time"
)

type Erc20 struct {
	Cli      *ethclient.Client
	Contract *binding.Erc20
}

func NewErc20(addr common.Address) (*Erc20, error) {
	err, cli := Client(ChainConfig)
	if err != nil {
		return nil, err
	}
	contractObj, err := binding.NewErc20(addr, cli)
	if err != nil {
		return nil, err
	}
	return &Erc20{
		Cli:      cli,
		Contract: contractObj,
	}, nil
}

func (c Erc20) Approve(amount *big.Int, address string) {
	var res *types.Transaction
	for {
		err, txOpts := GetAuth(c.Cli)
		if err != nil {
			log.Println(err)
			return
		}
		res, err = c.Contract.Approve(
			txOpts,
			common.HexToAddress(address),
			amount,
		)
		if err == nil {
			break
		} else {
			log.Println(fmt.Sprintf("Approve error: %v", err))
		}
		time.Sleep(3 * time.Second)
	}
	log.Println(fmt.Sprintf("Approve 成功"))
	for {
		receipt, err := c.Cli.TransactionReceipt(context.Background(), res.Hash())
		if err == nil && receipt.Status == 1 {
			break
		}
		log.Println(fmt.Sprintf("Approve 失败"), err)
		time.Sleep(time.Second * 2)
	}

	log.Println(fmt.Sprintf("Approve 确认成功 %s", res.Hash()))
}

func (c Erc20) Transfer(amount *big.Int, address string) {
	var res *types.Transaction
	for {
		err, txOpts := GetAuth(c.Cli)
		if err != nil {
			log.Println(err)
			return
		}
		res, err = c.Contract.Transfer(
			txOpts,
			common.HexToAddress(address),
			amount,
		)
		if err == nil {
			break
		} else {
			log.Println(fmt.Sprintf("Transfer error: %v", err))
		}
		time.Sleep(3 * time.Second)
	}
	log.Println(fmt.Sprintf("Transfer 成功"))
	for {
		receipt, err := c.Cli.TransactionReceipt(context.Background(), res.Hash())
		if err == nil && receipt.Status == 1 {
			break
		}
		time.Sleep(time.Second * 2)
	}

	log.Println(fmt.Sprintf("Transfer 确认成功 %s", res.Hash()))
}

func ApproveSigTest() {
	// 连接到以太坊节点（例如Infura）
	client, err := ethclient.Dial("https://mainnet.infura.io/v3/your-project-id")
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}

	// 私钥和地址（确保使用你的私钥和地址）
	//privateKey, err := crypto.HexToECDSA("your-private-key-here") // 请替换为你的私钥
	//if err != nil {
	//	log.Fatalf("Invalid private key: %v", err)
	//}
	//publicKey := privateKey.Public()
	//publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	//if !ok {
	//	log.Fatal("Error recovering public key")
	//}
	//fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA) // 计算公钥对应的地址

	// 设置交易参数
	fromAddress := common.HexToAddress("to-address-here")                  // 请替换为接收地址
	toAddress := common.HexToAddress("to-address-here")                    // 请替换为接收地址
	value := big.NewInt(1000000000000000000)                               // 发送1 ETH（以Wei为单位）
	gasLimit := uint64(21000)                                              // 设置Gas限制，根据你的需要进行调整
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress) // 获取nonce值
	if err != nil {
		log.Fatalf("Failed to get nonce: %v", err)
	}
	gasPrice, err := client.SuggestGasPrice(context.Background()) // 获取当前Gas价格建议
	if err != nil {
		log.Fatalf("Failed to suggest gas price: %v", err)
	}
	data, err := c.abi.Pack(method, params...)
	tx := types.NewTx(&types.LegacyTx{
		Nonce:    nonce,
		To:       &toAddress,
		Value:    amount,
		Gas:      gasLimit,
		GasPrice: gasPrice,
		Data:     data,
	})

	// 签名机返回的数据 sData
	sData := "0xab...."
	sData = strings.TrimPrefix(sData, "0x")
	sDataByte := hexutils.HexToBytes(sData)
	newTx := types.Transaction{}

	err := newTx.UnmarshalJSON(sDataByte)
	if err != nil {
		return
	}
	// 发送交易到链上
	err = client.SendTransaction(context.Background(), &newTx)
	if err != nil {
		log.Fatalf("Failed to send transaction: %v", err)
	}
	fmt.Println("Transaction sent:", signedTx.Hash().Hex()) // 打印交易哈希

}
