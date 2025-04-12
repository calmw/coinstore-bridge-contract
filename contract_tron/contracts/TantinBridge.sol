// SPDX-License-Identifier: MIT
pragma solidity ^0.8.22;

import "./interface/IERC20MintAble.sol";
import {IBridge} from "./interface/IBridge.sol";
import "@openzeppelin/contracts/utils/Address.sol";
import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import {ITantinBridge} from "./interface/ITantinBridge.sol";
import "@openzeppelin/contracts/utils/cryptography/ECDSA.sol";
import "@openzeppelin/contracts/utils/cryptography/MessageHashUtils.sol";
import {AccessControl} from "@openzeppelin/contracts/access/AccessControl.sol";

contract TantinBridge is AccessControl, ITantinBridge {
    using ECDSA for bytes32;
    using MessageHashUtils for bytes32;

    bytes32 public constant ADMIN_ROLE = keccak256("ADMIN_ROLE");
    bytes32 public constant BRIDGE_ROLE = keccak256("BRIDGE_ROLE");

    error ErrAssetsType(AssetsType assetsType);

    uint256 public sigNonce; // 签名nonce, parameter➕nonce➕chainID
    address private superAdminAddress;
    address private feeAddress;
    IBridge public Bridge; // bridge 合约
    uint256 public localNonce; // 跨链nonce
    mapping(address => mapping(uint256 => DepositRecord)) public depositRecord; // user => (depositNonce=> Deposit Record)
    mapping(address => bool) public blacklist; // 用户地址 => 是否在黑名单
    mapping(bytes32 => TokenInfo) public resourceIdToTokenInfo; //  resourceID => 设置的Token信息

    constructor() {
        _grantRole(DEFAULT_ADMIN_ROLE, msg.sender);
        superAdminAddress = 0x3942FdA93c573E2ce9e85B0bB00Ba98a144f27f6;
    }

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
    ) external onlyRole(ADMIN_ROLE) {
        require(
            checkAdminSetEnvSignature(signature_, feeAddress_,bridgeAddress_),
            "signature error"
        );
        feeAddress=feeAddress_;
        Bridge = IBridge(bridgeAddress_);
    }

    /**
        @notice 添加用户黑名单
        @param user 用户地址
        @param signature 签名
     */
    function adminAddBlacklist(
        address user,
        bytes memory signature
    ) external onlyRole(ADMIN_ROLE) {
        require(
            checkAdminAddBlacklistSignature(signature, user),
            "signature error"
        );
        blacklist[user] = true;
        emit AddBlacklist(user);
    }

    /**
        @notice 移除用户黑名单
        @param user 用户地址
        @param signature 签名
     */
    function adminRemoveBlacklist(
        address user,
        bytes memory signature
    ) external onlyRole(ADMIN_ROLE) {
        require(
            checkAdminRemoveBlacklistSignature(signature, user),
            "signature error"
        );
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
        @param signature 签名
     */
    function adminSetToken(
        bytes32 resourceID,
        AssetsType assetsType,
        address tokenAddress,
        bool burnable,
        bool mintable,
        bool pause,
        bytes memory signature
    ) external onlyRole(ADMIN_ROLE) {
        require(
            checkAdminSetTokenSignature(
                signature,
                resourceID,
                assetsType,
                tokenAddress,
                burnable,
                mintable,
                pause
            ),
            "signature error"
        );
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
        @param signature 签名，对资产接受地址的签名
     */
    function deposit(
        uint256 destinationChainId,
        bytes32 resourceId,
        address recipient,
        uint256 amount,
        bytes memory signature
    ) external payable {
        // 验证签名
        require(
            checkDepositSignature(signature, recipient, msg.sender),
            "signature error"
        );
        // 检测resource ID是否设置
        TokenInfo memory tokenInfo = resourceIdToTokenInfo[resourceId];
        require(uint8(tokenInfo.assetsType) > 0, "resourceId not exist");
        // 检测目标链ID
        uint256 chainId = Bridge.chainId();
        require(destinationChainId != chainId, "destinationChainId error");
        // 跨链费用比例，万分比
        uint256 fee = Bridge.getFeeByResourceId(resourceId);
        // 实际到账额度
        uint256 receiveAmount = amount - ((amount * fee) / 10000);
        // token地址
        address tokenAddress;
        if (tokenInfo.assetsType == AssetsType.Coin) {
            tokenAddress = address(0);
            require(msg.value == amount, "incorrect value supplied.");
        } else if (tokenInfo.assetsType == AssetsType.Erc20) {
            tokenAddress = tokenInfo.tokenAddress;
            IERC20 erc20 = IERC20(tokenAddress);
            if (tokenInfo.burnable) {
                erc20.transferFrom(msg.sender, address(0), receiveAmount);
            } else {
                erc20.transferFrom(msg.sender, address(this), receiveAmount);
            }
            erc20.transferFrom(
                msg.sender,
                feeAddress,
                amount - receiveAmount
            );
        } else {
            revert ErrAssetsType(tokenInfo.assetsType);
        }
        uint256 destId = destinationChainId;

        localNonce++;
        depositRecord[msg.sender][localNonce] = DepositRecord(
            tokenAddress,
            msg.sender,
            recipient,
            amount,
            fee,
            destId
        );
        // data
        bytes memory data = abi.encode(
            resourceId,
            chainId,
            msg.sender,
            recipient,
            receiveAmount,
            localNonce
        );
        Bridge.deposit(destId, resourceId, data);
        emit DepositEvent(
            msg.sender,
            recipient,
            amount,
            tokenAddress,
            localNonce,
            destId
        );
    }

//    /**
//        @notice 目标链执行到帐操作
//        @param data 跨链data, encode(originChainId,originDepositNonce,depositer,recipient,amount,resourceId)
//     */
//    function execute(bytes calldata data) public onlyRole(BRIDGE_ROLE) {
//        uint256 dataLength;
//        bytes32 resourceId;
//        uint256 originChainId;
//        address caller;
//        address recipient;
//        uint256 receiveAmount;
//        uint256 originNonce;
//        (
//            dataLength,
//            resourceId,
//            originChainId,
//            caller,
//            recipient,
//            receiveAmount,
//            originNonce
//        ) = abi.decode(
//            data,
//            (uint256, bytes32, uint256, address, address, uint256, uint256)
//        );
//
//        TokenInfo memory tokenInfo = resourceIdToTokenInfo[resourceId];
//        address tokenAddress = tokenInfo.tokenAddress;
//        if (tokenInfo.assetsType == AssetsType.Coin) {
//            Address.sendValue(payable(recipient), receiveAmount);
//        } else if (tokenInfo.assetsType == AssetsType.Erc20) {
//            if (tokenInfo.mintable) {
//                IERC20MintAble erc20 = IERC20MintAble(tokenAddress);
//                erc20.mint(recipient, receiveAmount);
//            } else {
//                IERC20 erc20 = IERC20(tokenAddress);
//                erc20.transfer(recipient, receiveAmount);
//            }
//        } else {
//            revert ErrAssetsType(tokenInfo.assetsType);
//        }
//
//        emit ExecuteEvent(
//            caller,
//            recipient,
//            receiveAmount,
//            tokenAddress,
//            originNonce,
//            originChainId
//        );
//    }

    /**
        @notice 查询跨链费用
        @param resourceId 跨链桥设置的resourceId
    */
    function getFee(bytes32 resourceId) external view returns (uint256) {
        return Bridge.getFeeByResourceId(resourceId);
    }

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
    ) public onlyRole(ADMIN_ROLE) {
        // 验证签名
        require(
            checkAdminWithdrawSignature(signature, tokenAddress, amount),
            "signature error"
        );
        if (tokenAddress == address(0)) {
            Address.sendValue(payable(msg.sender), amount);
        } else {
            IERC20 erc20 = IERC20(tokenAddress);
            erc20.transfer(msg.sender, amount);
        }
    }

    // 验证deposit签名
    function checkDepositSignature(
        bytes memory signature,
        address recipient,
        address sender
    ) private pure returns (bool) {
        bytes32 messageHash = keccak256(abi.encode(recipient));
        address recoverAddress = recoverSigner(
            toEthSignedMessageHash(messageHash),
            signature
        );

        return recoverAddress == sender;
    }

    // 验证adminSetEnv签名
    function checkAdminSetEnvSignature(
        bytes memory signature,
        address feeAddress,
        address bridgeAddress
    ) private returns (bool) {
        bytes32 messageHash = keccak256(abi.encode(sigNonce,feeAddress, bridgeAddress));
        address recoverAddress = recoverSigner(
            toEthSignedMessageHash(messageHash),
            signature
        );

        bool res = recoverAddress == superAdminAddress;
        if (res) {
            sigNonce++;
        }

        return res;
    }

    // 验证adminAddBlacklist签名
    function checkAdminAddBlacklistSignature(
        bytes memory signature,
        address user
    ) private returns (bool) {
        uint256 chainId = Bridge.chainId();
        bytes32 messageHash = keccak256(
            abi.encode(sigNonce, chainId, user, sigNonce)
        );
        address recoverAddress = recoverSigner(
            toEthSignedMessageHash(messageHash),
            signature
        );

        bool res = recoverAddress == superAdminAddress;
        if (res) {
            sigNonce++;
        }

        return res;
    }

    // 验证adminRemoveBlacklist签名
    function checkAdminRemoveBlacklistSignature(
        bytes memory signature,
        address user
    ) private returns (bool) {
        uint256 chainId = Bridge.chainId();
        bytes32 messageHash = keccak256(abi.encode(sigNonce, chainId, user));
        address recoverAddress = recoverSigner(
            toEthSignedMessageHash(messageHash),
            signature
        );

        bool res = recoverAddress == superAdminAddress;
        if (res) {
            sigNonce++;
        }

        return res;
    }

    // 验证adminSetTokenSignature签名
    function checkAdminSetTokenSignature(
        bytes memory signature,
        bytes32 resourceID,
        AssetsType assetsType,
        address tokenAddress,
        bool burnable,
        bool mintable,
        bool pause
    ) private returns (bool) {
        uint256 chainId = Bridge.chainId();
        bytes32 messageHash = keccak256(
            abi.encode(
                sigNonce,
                chainId,
                resourceID,
                assetsType,
                tokenAddress,
                burnable,
                mintable,
                pause
            )
        );
        address recoverAddress = recoverSigner(
            toEthSignedMessageHash(messageHash),
            signature
        );

        bool res = recoverAddress == superAdminAddress;
        if (res) {
            sigNonce++;
        }
        return res;
    }

    // 验证adminWithdrawSignature签名
    function checkAdminWithdrawSignature(
        bytes memory signature,
        address tokenAddress,
        uint256 amount
    ) private returns (bool) {
        uint256 chainId = Bridge.chainId();
        bytes32 messageHash = keccak256(
            abi.encode(sigNonce, chainId, tokenAddress, amount)
        );
        address recoverAddress = recoverSigner(
            toEthSignedMessageHash(messageHash),
            signature
        );

        bool res = recoverAddress == superAdminAddress;
        if (res) {
            sigNonce++;
        }
        return res;
    }

    function toEthSignedMessageHash(
        bytes32 hash
    ) public pure returns (bytes32) {
        return
            keccak256(
            abi.encodePacked("\x19Ethereum Signed Message:\n32", hash)
        );
    }

    function recoverSigner(
        bytes32 _msgHash,
        bytes memory _signature
    ) public pure returns (address) {
        require(_signature.length == 65, "invalid signature length");
        bytes32 r;
        bytes32 s;
        uint8 v;
        assembly {
            r := mload(add(_signature, 0x20))
            s := mload(add(_signature, 0x40))
            v := byte(0, mload(add(_signature, 0x60)))
        }
        return ecrecover(_msgHash, v, r, s);
    }
}
