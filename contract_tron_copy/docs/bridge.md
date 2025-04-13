#### vote合约

- 合约地址：

#### 合约API

- getChainId 获取自定义链ID
- adminPauseTransfers 暂停跨链
- adminUnpauseTransfers 开启跨链
- adminAddRelayer 添加relayer账户
- adminRemoveRelayer 移除relayer账户
- adminChangeRelayerThreshold 设置投票可通过时的最小投票数量
- adminSetResource 跨链设置
    - resource设置
    - resourceID 跨链的resourceID
    - assetsType 该币的类型
    - tokenAddress 对应的token合约地址，coin为0地址
    - fee 该币的跨链费用
    - burnable true burn;false lock
    - mintable true mint;false release
    - blacklist 该币种是否在黑名单中/是否允许跨链。币种黑名单/禁止该币种跨链
    - tantinAddress 对应的tantin业务合约地址
    - executeFunctionSig tantin业务合约执行到帐操作的方法签名
- getToeknInfoByResourceId 由resourceId获取跨链的token信息
- deposit 跨链操作
- isRelayer 检查某地址是否是relayer账户
- voteProposal relayer投票
- cancelProposal 取消投票
- executeProposal 执行投票通过后的到帐操作

#### ABI 文件

