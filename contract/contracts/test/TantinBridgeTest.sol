// SPDX-License-Identifier: MIT
pragma solidity ^0.8.22;

contract TantinBridgeTest {
    enum AssetsType {
        None,
        Coin,
        Erc20,
        Erc721,
        Erc1155
    }

    /**
        @notice 设置
        @param bridgeAddress_ bridge合约地址
     */
    function adminSetEnv(address bridgeAddress_) public {}

    /**
        @notice 添加用户黑名单
        @param user 用户地址
     */
    function adminAddBlacklist(address user) public {}

    /**
        @notice 移除用户黑名单
        @param user 用户地址
     */
    function adminRemoveBlacklist(address user) public {}

    /**
        @notice token/coin设置
        @param resourceID 跨链的resourceID。resourceID和币对关联，不是和币关联的。 resourceID 1 =>(tokenA <=> token B);resourceID 2 =>(tokenA <=> token C)
        @param assetsType 该币的类型
        @param tokenAddress 对应的token合约地址，coin为0地址
        @param burnable true burn;false lock
        @param mintable  true mint;false release
        @param pause 是否暂停该币种跨链
     */
    function adminSetToken(
        bytes32 resourceID,
        AssetsType assetsType,
        address tokenAddress,
        bool burnable,
        bool mintable,
        bool pause
    ) public {}

    /**
        @notice 发起跨链
        @param destinationChainId 目标链ID
        @param resourceId 跨链桥设置的resourceId
        @param recipient 目标链资产接受者地址
        @param amount 跨链金额
        @param signature 签名，对资产接受地址的签名
     */
    function deposit(
        uint256 destinationChainId,
        bytes32 resourceId,
        address recipient,
        uint256 amount,
        bytes memory signature
    ) external payable {}

    // 验证签名
    function checkDepositSignature(
        bytes memory signature,
        address recipient,
        address sender
    ) private pure returns (bool) {}

    /**
        @notice 查询跨链费用
        @param resourceId 跨链桥设置的resourceId
    */
    function getFee(bytes32 resourceId) external view returns (uint256) {}

    /**
        @notice 提取跨链桥资产
        @param tokenAddress 币种地址，coin为0地址
        @param amount 提取数量
     */
    function adminWithdraw(address tokenAddress, uint256 amount) public {}
}
