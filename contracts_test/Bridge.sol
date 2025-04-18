// SPDX-License-Identifier: MIT
pragma solidity ^0.8.22;


contract Bridge {

    /**
        @notice 设置
        @param voteAddress_ 投票合约地址
        @param chainId_ 链ID
        @param signature_ 签名
     */
    function adminSetEnv(
        address voteAddress_,
        uint256 chainId_,
        uint256 chainType_,
        bytes memory signature_
    ) external {}

    /**
        @notice 暂停跨链、提案的的创建与投票和目标链执行操作
        @param signature 签名
     */
    function adminPauseTransfers(
        bytes memory signature
    ) external {}

    /**
        @notice 开启跨链、提案的的创建与投票和目标链执行操作
        @param signature 签名
     */
    function adminUnpauseTransfers(
        bytes memory signature
    ) external {}

    /**
        @notice resource设置
        @param resourceID 跨链的resourceID
        @param assetsType 该币的类型
        @param tokenAddress 对应的token合约地址，coin为0地址
        @param fee 该币的跨链费用
        @param pause 该币种是否在黑名单中/是否允许跨链。币种黑名单/禁止该币种跨链
        @param tantinAddress 对应的tantin业务合约地址
     */
    function adminSetResource(
        bytes32 resourceID,
        AssetsType assetsType,
        address tokenAddress,
        uint256 fee,
        bool pause,
        bool burnable, // true burn;false lock
        bool mintable,
        address tantinAddress,
        bytes memory signature
    ) external {}

    /**
        @notice 资产跨链
        @param destinationChainId 目标链ID
        @param resourceId 跨链的resourceID
        @param data   跨链data
     */
    function deposit(
        uint256 destinationChainId,
        bytes32 resourceId,
        bytes calldata data
    ) external payable {}

    // 获取跨链费用
    function getFeeByResourceId(
        bytes32 resourceId
    ) public view returns (uint256) {}

    // 由resourceId获取tantin address信息
    function getContractAddressByResourceId(
        bytes32 resourceId
    ) public view returns (address) {}

    // 由resourceId获取token信息
    function getTokenInfoByResourceId(
        bytes32 resourceId
    ) public view returns (uint8, address, bool, uint256, bool, bool) {}

    // 验证adminSetEnv签名
    function checkAdminSetEnvSignature(
        bytes memory signature_,
        address voteAddress_,
        uint256 chainId_,
        uint256 chainType_
    ) private returns (bool) {}

    // 验证adminPauseTransfers签名
    function checkAdminPauseTransfersSignature(
        bytes memory signature
    ) private returns (bool) {}

    // 验证adminUnpauseTransfers签名
    function checkAdminUnpauseTransfersSignature(
        bytes memory signature
    ) private returns (bool) {}

    // 验证adminSetResource签名
    function checkAdminSetResourceSignature(
        bytes memory signature,
        bytes32 resourceID,
        AssetsType assetsType,
        address tokenAddress,
        uint256 fee,
        bool pause,
        bool burnable,
        bool mintable,
        address tantinAddress
    ) private returns (bool) {}
}
