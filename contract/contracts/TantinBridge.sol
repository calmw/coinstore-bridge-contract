// SPDX-License-Identifier: MIT
pragma solidity ^0.8.22;


contract TantinBridge {

    /**
        @notice 设置
        @param bridgeAddress_ bridge合约地址
        @param feeAddress_ 跨链费接受地址
        @param signature_ 签名
     */
    function adminSetEnv(
        address feeAddress_,
        address bridgeAddress_,
        bytes memory signature_
    ) external  {}

    /**
        @notice 添加用户黑名单
        @param user 用户地址
        @param signature 签名
     */
    function adminAddBlacklist(
        address user,
        bytes memory signature
    ) external {}

    /**
        @notice 移除用户黑名单
        @param user 用户地址
        @param signature 签名
     */
    function adminRemoveBlacklist(
        address user,
        bytes memory signature
    ) external  {}

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

    /**
        @notice 查询跨链费用
        @param resourceId 跨链桥设置的resourceId
    */
    function getFee(bytes32 resourceId) external view returns (uint256) {}

    /**
        @notice 提取跨链桥资产
        @param tokenAddress 币种地址，coin为0地址
        @param amount 提取数量
        @param signature 签名
     */
    function adminWithdraw(
        address tokenAddress,
        uint256 amount,
        bytes memory signature
    ) public {}

    // 验证deposit签名
    function checkDepositSignature(
        bytes memory signature,
        address recipient,
        address sender
    ) private pure returns (bool) {}

    function checkAdminSetEnvSignature(
        bytes memory signature_,
        address feeAddress_,
        address bridgeAddress_
    ) private returns (bool) {}

    // 验证adminAddBlacklist签名
    function checkAdminAddBlacklistSignature(
        bytes memory signature,
        address user
    ) private returns (bool) {}

    // 验证adminRemoveBlacklist签名
    function checkAdminRemoveBlacklistSignature(
        bytes memory signature,
        address user
    ) private returns (bool) {}

    // 验证adminSetTokenSignature签名
    function checkAdminSetTokenSignature(
        bytes memory signature,
        bytes32 resourceID,
        AssetsType assetsType,
        address tokenAddress,
        bool burnable,
        bool mintable,
        bool pause
    ) private returns (bool) {}

    // 验证adminWithdrawSignature签名
    function checkAdminWithdrawSignature(
        bytes memory signature,
        address tokenAddress,
        uint256 amount
    ) private returns (bool) {}
}
