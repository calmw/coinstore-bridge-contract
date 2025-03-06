// SPDX-License-Identifier: MIT
pragma solidity ^0.8.22;

import {IBridge} from "./interface/IBridge.sol";
import {IERC20} from "./interface/IERC20.sol";
import {ITantinBridge} from "./interface/ITantinBridge.sol";
import "@openzeppelin/contracts/utils/Address.sol";
import {Initializable} from "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import {AccessControl} from "@openzeppelin/contracts/access/AccessControl.sol";

/// ERC20/Coin跨链

contract TantinBridge is AccessControl, ITantinBridge, Initializable {
    bytes32 public constant ADMIN_ROLE = keccak256("ADMIN_ROLE");
    bytes32 public constant BRIDGE_ROLE = keccak256("BRIDGE_ROLE");

    IBridge public Bridge; // bridge 合约
    mapping(address => uint64) public userDepositNonce; // 用户跨链nonce
    mapping(address => mapping(uint256 => DepositRecord)) public depositRecord; // user => (depositNonce=> Deposit Record)
    mapping(address => bool) public blacklist; // 用户地址 => 是否在黑名单
    mapping(bytes32 => TokenInfo) public resourceIdToTokenInfo; //  resourceID => 设置的Token信息

    function initialize() public initializer {
        _grantRole(DEFAULT_ADMIN_ROLE, msg.sender);
    }

    /**
        @notice 添加用户黑名单
        @param user 用户地址
     */
    function adminAddBlacklist(address user) external onlyRole(ADMIN_ROLE) {
        blacklist[user] = true;
        emit AddBlacklist(user);
    }

    /**
        @notice 移除用户黑名单
        @param user 用户地址
     */
    function adminRemoveBlacklist(address user) external onlyRole(ADMIN_ROLE) {
        blacklist[user] = false;
        emit RemoveBlacklist(user);
    }

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
    ) external onlyRole(ADMIN_ROLE) {
        resourceIdToTokenInfo[resourceID] = TokenInfo(
            assetsType,
            tokenAddress,
            burnable,
            mintable,
            pause
        );

        emit SetTokenEvent(
            resourceID,
            assetsType,
            tokenAddress,
            burnable,
            mintable,
            pause
        );
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
    ) external payable {
        // 检测resource ID是否设置
        TokenInfo memory tokenInfo = resourceIdToTokenInfo[resourceId];
        require(uint8(tokenInfo.assetsType) > 0, "resourceId not exist");
        // 检测目标链ID
        uint256 chainId = Bridge.getChainId();
        require(destinationChainId != chainId, "destinationChainId error");
        // 跨链费用
        uint256 fee = Bridge.getFeeByResourceId(resourceId);
        userDepositNonce[msg.sender]++;
        // 跨链token
        address tokenAddress;
        bytes memory data = abi.encode(
            resourceId,
            chainId,
            msg.sender,
            recipient,
            msg.value,
            userDepositNonce[msg.sender]
        );
        if (tokenInfo.assetsType == AssetsType.Coin) {
            tokenAddress = address(0);
            require(msg.value == fee + amount, "coin value error");
        }
        if (tokenInfo.assetsType == AssetsType.Erc20) {
            tokenAddress = tokenInfo.tokenAddress;
            IERC20 erc20 = IERC20(tokenAddress);
            if (tokenInfo.burnable) {
                erc20.transferFrom(msg.sender, address(0), amount);
            } else {
                erc20.transferFrom(msg.sender, address(this), amount);
            }
        }
        depositRecord[msg.sender][userDepositNonce[msg.sender]] = DepositRecord(
            tokenAddress,
            msg.sender,
            recipient,
            amount,
            destinationChainId
        );
        Bridge.deposit{value: fee}(destinationChainId, resourceId, data);

        emit DepositEvent(
            msg.sender,
            recipient,
            amount,
            tokenAddress,
            userDepositNonce[msg.sender],
            destinationChainId
        );
    }

    /**
        @notice 目标链执行到帐操作
        @param data 跨链data, encode(originChainId,originDepositNonce,depositer,recipient,amount,resourceId)
     */
    function execute(bytes calldata data) public onlyRole(BRIDGE_ROLE) {
        bytes32 resourceId;
        address sender;
        address recipient;
        uint256 amount;
        uint256 originChainId;
        uint256 userNonce;
        (resourceId, originChainId, sender, recipient, amount, userNonce) = abi
            .decode(
                data,
                (bytes32, uint256, address, address, uint256, uint256)
            );

        TokenInfo memory tokenInfo = resourceIdToTokenInfo[resourceId];
        address tokenAddress = tokenInfo.tokenAddress;
        if (tokenInfo.assetsType == AssetsType.Coin) {
            Address.sendValue(payable(recipient), amount);
        }
        if (tokenInfo.assetsType == AssetsType.Erc20) {
            IERC20 erc20 = IERC20(tokenAddress);
            erc20.transfer(recipient, amount);
            if (tokenInfo.mintable) {
                erc20.mint(recipient, amount);
            } else {
                erc20.transfer(recipient, amount);
            }
        }
        emit ExecuteEvent(
            sender,
            recipient,
            amount,
            tokenAddress,
            userNonce,
            originChainId
        );
    }

    /**
        @notice 获取跨链记录
        @param user_ 用户地址
        @param userDepositNonce_ nonce
    */
    function getDepositRecord(
        address user_,
        uint256 userDepositNonce_
    ) external view returns (DepositRecord memory) {
        return depositRecord[user_][userDepositNonce_];
    }
}
