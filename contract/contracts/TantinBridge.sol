// SPDX-License-Identifier: MIT
pragma solidity ^0.8.22;

import "./interface/IBridge.sol";
import "./interface/ITantinBridge.sol";
import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import {AccessControl} from "@openzeppelin/contracts/access/AccessControl.sol";

/// ERC20跨链 demo

contract TantinBridge is AccessControl, ITantinBridge, Initializable {
    bytes32 public constant BRIDGE_ROLE = keccak256("BRIDGE_ROLE");

    IBridge public bridge; // bridge 合约
    uint256 public depositNonce; // 自增跨链nonce

    function initialize() public initializer {}

    /**
        @notice 发起跨链 demo
        @param destinationChainId 目标链ID
        @param resourceId 跨链桥设置的resourceId.
        @param recipient 目标链资产接受者地址.
        @param amount 跨链金额.
     */
    function deposit(
        uint256 destinationChainId,
        bytes32 resourceId,
        address recipient,
        uint256 amount
    ) external {}

    /**
        @notice 跨链目标链执行 demo
        @param originChainId 源链ID
        @param originDepositNonce 源链ID depositNonce
        @param recipient 目标链资产接受者地址.
        @param amount 跨链金额.
     */
    function execute(
        uint256 originChainId,
        uint256 originDepositNonce,
        bytes32 resourceID,
        bytes calldata data
    ) public onlyRole(BRIDGE_ROLE) {}
}
