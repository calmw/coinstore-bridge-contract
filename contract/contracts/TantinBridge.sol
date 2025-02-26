// SPDX-License-Identifier: MIT
pragma solidity ^0.8.22;

import {IBridge} from "./interface/IBridge.sol";
import {IERC20} from "./interface/IERC20.sol";
import {ITantinBridge} from "./interface/ITantinBridge.sol";
import {Initializable} from "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import {AccessControl} from "@openzeppelin/contracts/access/AccessControl.sol";

/// ERC20跨链 demo

contract TantinBridge is AccessControl, ITantinBridge, Initializable {
    bytes32 public constant BRIDGE_ROLE = keccak256("BRIDGE_ROLE");

    IBridge public Bridge; // bridge 合约
    uint256 public depositNonce; // 跨链nonce
    mapping(uint256 => mapping(uint256 => DepositErc20Record))
        public depositRecord; // destinationChainId => (depositNonce=> Deposit Record)

    function initialize() public initializer {
        _grantRole(DEFAULT_ADMIN_ROLE, msg.sender);
    }

    /**
        @notice 发起跨链
        @param destinationChainId 目标链ID
        @param resourceId 跨链桥设置的resourceId
        @param recipient 目标链资产接受者地址
        @param amount 跨链金额
     */
    function deposit(
        uint256 destinationChainId,
        bytes32 resourceId,
        address recipient,
        uint256 amount
    ) external {}

    /**
        @notice 目标链执行到帐操作
        @param data 跨链data, encode(originChainId,originDepositNonce,depositer,recipient,amount,resourceId)
     */
    function execute(bytes calldata data) public onlyRole(BRIDGE_ROLE) {}

    /**
        @notice 锁定ERC20资产
        @param tokenAddress token地址
        @param amount 锁定金额
     */
    function lockERC20(address tokenAddress, uint256 amount) internal {
        IERC20 erc20 = IERC20(tokenAddress);
        erc20.transferFrom(msg.sender, address(this), amount);
    }

    /**
        @notice 释放RC20资产
        @param tokenAddress token地址
        @param recipient 接收者地址
        @param amount 释放金额
     */
    function releaseERC20(
        address tokenAddress,
        address recipient,
        uint256 amount
    ) internal {
        IERC20 erc20 = IERC20(tokenAddress);
        erc20.transfer(recipient, amount);
    }

    /**
        @notice 铸造ERC20资产,需要知道该token具体使用的铸造方法以及权限问题
        @param tokenAddress token地址
        @param recipient 接收者地址
        @param amount 铸造金额
     */
    function mintERC20(
        address tokenAddress,
        address recipient,
        uint256 amount
    ) internal {
        IERC20 erc20 = IERC20(tokenAddress);
        erc20.mint(recipient, amount);
    }

    /**
        @notice 销毁ERC20资产,需要知道该token具体使用的销毁方法以及权限问题
        @param tokenAddress token地址
        @param owner token所有者地址
        @param amount 销毁金额
     */
    function burnERC20(
        address tokenAddress,
        address owner,
        uint256 amount
    ) internal {
        IERC20 erc20 = IERC20(tokenAddress);
        erc20.burnFrom(owner, amount);
    }

    /**
        @notice 获取跨链记录
        @param depositNonce_ 跨链nonce
        @param destinationChainId_ 目标链ID
    */
    function getDepositRecord(
        uint256 depositNonce_,
        uint256 destinationChainId_
    ) external view returns (DepositErc20Record memory) {
        return depositRecord[destinationChainId_][depositNonce_];
    }
}
