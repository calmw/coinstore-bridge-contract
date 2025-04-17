package main

func main() {

}

func ApproveSigTest() {
	// 连接到以太坊节点（例如Infura）
	//client, err := ethclient.Dial("https://rpc.tantin.com")
	//if err != nil {
	//	log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	//}

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
	//fromAddress := common.HexToAddress("to-address-here")                  // 请替换为接收地址
	//toAddress := common.HexToAddress("to-address-here")                    // 请替换为接收地址
	//value := big.NewInt(1000000000000000000)                               // 发送1 ETH（以Wei为单位）
	//gasLimit := uint64(21000)                                              // 设置Gas限制，根据你的需要进行调整
	//nonce, err := client.PendingNonceAt(context.Background(), fromAddress) // 获取nonce值
	//if err != nil {
	//	log.Fatalf("Failed to get nonce: %v", err)
	//}
	//gasPrice, err := client.SuggestGasPrice(context.Background()) // 获取当前Gas价格建议
	//if err != nil {
	//	log.Fatalf("Failed to suggest gas price: %v", err)
	//}
	//data, err := c.abi.Pack(method, params...)
	//tx := types.NewTx(&types.LegacyTx{
	//	Nonce:    nonce,
	//	To:       &toAddress,
	//	Value:    amount,
	//	Gas:      gasLimit,
	//	GasPrice: gasPrice,
	//	Data:     data,
	//})
	//
	//// 签名机返回的数据 sData
	//sData := "0xab...."
	//sData = strings.TrimPrefix(sData, "0x")
	//sDataByte := hexutils.HexToBytes(sData)
	//newTx := types.Transaction{}
	//
	//err := newTx.UnmarshalJSON(sDataByte)
	//if err != nil {
	//	return
	//}
	//// 发送交易到链上
	//err = client.SendTransaction(context.Background(), &newTx)
	//if err != nil {
	//	log.Fatalf("Failed to send transaction: %v", err)
	//}
	//fmt.Println("Transaction sent:", signedTx.Hash().Hex()) // 打印交易哈希

}
