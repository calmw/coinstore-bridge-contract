// SPDX-License-Identifier: MIT
pragma solidity ^0.8.22;

contract Vote  {
    struct Proposal {
        bytes32 resourceId;
        bytes32 dataHash;
        address[] yesVotes;
        address[] noVotes;
        ProposalStatus status;
        uint256 proposedBlock;
    }
    enum Vote {
        No,
        Yes
    }

    enum ProposalStatus {
        Inactive,
        Active,
        Passed,
        Executed,
        Cancelled
    }
    /**
        @notice 设置
        @param bridgeAddress_ Bridge合约地址
        @param expiry_ 提案过期的块高差
        @param relayerThreshold_ 提案通过的投票数量
        @param signature_ 签名
     */
    function adminSetEnv(
        address bridgeAddress_,
        uint256 expiry_,
        uint256 relayerThreshold_,
        bytes memory signature_
    ) external  {}

    /**
        @notice 设置投票可通过时的最小投票数量
        @param newThreshold 投票可通过时的最小投票数量
        @param signature 签名
     */
    function adminChangeRelayerThreshold(
        uint256 newThreshold,
        bytes memory signature
    ) external {}

    /**
        @notice 添加relayer账户
        @notice Only callable by an address that currently has the admin role.
        @param relayerAddress Address of relayer to be added.
        @notice Emits {RelayerAdded} event.
        @param signature 签名
     */
    function adminAddRelayer(
        address relayerAddress,
        bytes memory signature
    ) external  {}

    /**
        @notice 删除relayer账户
        @notice Only callable by an address that currently has the admin role.
        @param relayerAddress Address of relayer to be removed.
        @notice Emits {RelayerRemoved} event.
        @param signature 签名
     */
    function adminRemoveRelayer(
        address relayerAddress,
        bytes memory signature
    ) external  {}

    /**
        @notice relayer执行投票通过后的到帐操作
        @param originChainId 源链ID
        @param originDepositNonce 源链nonce
        @param resourceId 跨链的resourceID
        @param dataHash dataHash
     */
    function voteProposal(
        uint256 originChainId,
        uint256 originDepositNonce,
        bytes32 resourceId,
        bytes32 dataHash
    ) external  {}

    /**
        @notice relayer执行投票通过后的到帐操作
        @param originChainID 源链ID
        @param originDepositNonce 源链nonce
        @param dataHash dataHash
     */
    function cancelProposal(
        uint256 originChainID,
        uint256 originDepositNonce,
        bytes32 dataHash
    ) public  {}

    /**
        @notice relayer执行投票通过后的到帐操作
        @param originChainId 源链ID
        @param originDepositNonce 源链nonce
        @param data 跨链data
     */
    function executeProposal(
        uint256 originChainId,
        uint256 originDepositNonce,
        bytes calldata data
    ) external  {}

    /**
        @notice 目标链执行到帐操作
        @param data 跨链data, encode(originChainId,originDepositNonce,depositer,recipient,amount,resourceId)
     */
    function execute(bytes calldata data) public {}

    // 获取投票信息
    function getProposal(
        uint256 originChainID,
        uint256 depositNonce,
        bytes32 dataHash
    ) external view returns (Proposal memory) {}

    // 验证adminSetEnv签名
    function checkAdminSetEnvSignature(
        bytes memory signature_,
        address bridgeAddress_,
        uint256 expiry_,
        uint256 relayerThreshold_
    ) private returns (bool) {}

    // 验证adminChangeRelayerThreshold签名
    function checkAdminChangeRelayerThresholdSignature(
        bytes memory signature_,
        uint256 newThreshold
    ) private returns (bool) {}

    // 验证adminAddRelayer签名
    function checkAdminAddRelayerSignature(
        bytes memory signature_,
        address relayerAddress
    ) private returns (bool) {}

    // 验证adminRemoveRelayer签名
    function checkAdminRemoveRelayerSignature(
        bytes memory signature_,
        address relayerAddress
    ) private returns (bool) {}
}
