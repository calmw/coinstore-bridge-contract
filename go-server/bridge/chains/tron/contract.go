package tron

func Call() {
	// 携带的数据
	//triggerData := "[{\"address\":\"TWRhFW2AQg8Ae2j8BynFnGNfj7xg95qhvQ\"},{\"uint256\":\"100000\"}]"
	//
	//// 构建Tx，注意，此处只是构造，tx并没有被发送
	//// 参数分别为 发送者地址 合约地址 调用方法(正常字符串) gas Tx发送的TRX数量 Tx发送的TRC20代币 代币数量
	//tx, _ := rpcClient.TriggerContract(wallet, contract, "transfer(address,uint256)", triggerData, 300000000, 0, "", 0)
	//
	//// 获得keystore与account
	//ks, acct, _ := store.UnlockedKeystore(wallet, accountPassword)
	//// 封装Tx
	//ctrlr := transaction.NewController(rpcClient, ks, acct, tx.Transaction)
	//// 真正执行Tx，并判断执行结果
	//if err = ctrlr.ExecuteTransaction(); err != nil {
	//	log.Println(err.Error())
	//	return
	//}
	//// 此时Tx才上链
	//log.Println(common.BytesToHexString(tx.GetTxid()))
	//log.Println(common.BytesToHexString(tx.GetResult().GetMessage()))
}
