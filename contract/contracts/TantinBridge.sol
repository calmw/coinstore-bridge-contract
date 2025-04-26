// SPDX-License-Identifier: MIT
pragma solidity ^0.8.22;

import "./interface/IERC20MintAble.sol";
import {IBridge} from "./interface/IBridge.sol";
import "@openzeppelin/contracts/utils/Address.sol";
import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import {ITantinBridge} from "./interface/ITantinBridge.sol";
import "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";
import "@openzeppelin/contracts/utils/cryptography/ECDSA.sol";
import "@openzeppelin/contracts/utils/cryptography/MessageHashUtils.sol";
import {AccessControl} from "@openzeppelin/contracts/access/AccessControl.sol";
import {Initializable} from "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";

contract TantinBridge is AccessControl, ITantinBridge, Initializable {
    using ECDSA for bytes32;
    using MessageHashUtils for bytes32;

    using Address for address;

    using SafeERC20 for IERC20;

    bytes32 public constant ADMIN_ROLE = keccak256("ADMIN_ROLE");
    bytes32 public constant BRIDGE_ROLE = keccak256("BRIDGE_ROLE");

    error ErrAssetsType(uint8 assetsType);

    uint256 public sigNonce; // 签名nonce, parameter➕nonce➕chainID
    address private superAdminAddress;
    address private serverAddress;
    address private feeAddress;
    IBridge public Bridge; // bridge 合约
    uint256 public localNonce; // 跨链nonce
    mapping(address => mapping(uint256 => DepositRecord)) public depositRecord; // user => (depositNonce=> Deposit Record)
    mapping(address => bool) public blacklist; // 用户地址 => 是否在黑名单

    function initialize() public initializer {
        _grantRole(DEFAULT_ADMIN_ROLE, msg.sender);
        superAdminAddress = 0xa47142f08f859aCeb2127C6Ab66eC8c8bc4FFBA9;
    }

    /**
        @notice 设置
        @param bridgeAddress_ bridge合约地址
        @param serverAddress_ 服务端价格签名地址
        @param feeAddress_ 跨链费接受地址
        @param signature_ 签名
     */
    function adminSetEnv(
        address feeAddress_,
        address serverAddress_,
        address bridgeAddress_,
        bytes memory signature_
    ) external onlyRole(ADMIN_ROLE) {
        require(
            checkAdminSetEnvSignature(
                signature_,
                feeAddress_,
                serverAddress_,
                bridgeAddress_
            ),
            "signature error"
        );
        feeAddress = feeAddress_;
        serverAddress = serverAddress_;
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
        @notice 发起跨链
        @param destinationChainId 目标链ID
        @param resourceId 跨链桥设置的resourceId
        @param recipient 目标链资产接受者地址
        @param amount 跨链金额
        @param price 当前token价格
        @param priceTimestamp 价格
        @param priceSignature 签名，对资产接受地址的签名
        @param recipientSignature 签名，对资产接受地址的签名
     */
    function deposit(
        uint256 destinationChainId,
        bytes32 resourceId,
        address recipient,
        uint256 amount,
        uint256 price,
        uint256 priceTimestamp,
        bytes memory priceSignature,
        bytes memory recipientSignature
    ) external payable {
        DepositData memory depositData;
        depositData.price = price;
        depositData.amount = amount;
        depositData.recipient = recipient;
        depositData.resourceId = resourceId;
        depositData.destinationChainId = destinationChainId;
        depositData.chainId = Bridge.chainId();
        // 验证接受地址签名
        require(
            checkDepositSignature(recipientSignature, recipient, msg.sender),
            "recipient signature error"
        );
        // 验证价格签名
        require(
            checkPriceSignature(
                priceSignature,
                depositData.chainId,
                price,
                priceTimestamp
            ),
            "price signature error"
        );
        // 验证价格签名时间
        require(block.timestamp - 30 < priceTimestamp, "please try again");
        // 检测resource ID是否设置
        (
            uint8 assetsType,
            address tokenAddress,
            bool pause,
            uint256 decimal,
            uint256 fee,
            bool burnable,
            bool mintable
        ) = Bridge.getTokenInfoByResourceId(depositData.resourceId);
        depositData.fee = fee;
        depositData.burnable = burnable;
        depositData.assetsType = assetsType;
        depositData.tokenAddress = tokenAddress;
        require(depositData.assetsType > 0, "resourceId not exist");
        // 检测目标链ID
        require(
            depositData.destinationChainId != depositData.chainId,
            "destinationChainId error"
        );

        // 实际到账额度
        depositData.feeAmount =
            ((depositData.price * depositData.fee) / 1e6) *
            decimal;
        require(depositData.amount > depositData.feeAmount, "amount too small");
        depositData.receiveAmount = depositData.amount - depositData.feeAmount;
        if (assetsType == uint8(AssetsType.Coin)) {
            tokenAddress = address(0);
            require(
                msg.value == depositData.amount,
                "incorrect value supplied."
            );
            Address.sendValue(payable(feeAddress), depositData.feeAmount);
            Address.sendValue(
                payable(address(this)),
                depositData.receiveAmount
            );
        } else if (assetsType == uint8(AssetsType.Erc20)) {
            IERC20 erc20 = IERC20(tokenAddress);
            if (burnable) {
                erc20.safeTransferFrom(
                    msg.sender,
                    address(0),
                    depositData.receiveAmount
                );
            } else {
                erc20.safeTransferFrom(
                    msg.sender,
                    address(this),
                    depositData.receiveAmount
                );
            }
            erc20.safeTransferFrom(
                msg.sender,
                feeAddress,
                depositData.feeAmount
            );
        } else {
            revert ErrAssetsType(assetsType);
        }

        localNonce++;
        depositRecord[msg.sender][localNonce] = DepositRecord(
            tokenAddress,
            msg.sender,
            depositData.recipient,
            amount,
            fee,
            depositData.destinationChainId
        );
        // data
        bytes memory data = abi.encode(
            depositData.resourceId,
            depositData.chainId,
            msg.sender,
            depositData.recipient,
            depositData.receiveAmount,
            localNonce
        );
        Bridge.deposit(
            depositData.destinationChainId,
            depositData.resourceId,
            data
        );
        emit DepositEvent(
            msg.sender,
            depositData.recipient,
            amount,
            tokenAddress,
            localNonce,
            depositData.destinationChainId
        );
    }

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
            erc20.safeTransfer(msg.sender, amount);
        }
    }

    // 验证deposit签名
    function checkDepositSignature(
        bytes memory signature,
        address recipient,
        address sender
    ) private pure returns (bool) {
        bytes32 messageHash = keccak256(abi.encode(recipient));
        address recoverAddress = messageHash.toEthSignedMessageHash().recover(
            signature
        );

        return recoverAddress == sender;
    }

    // 验证price签名
    function checkPriceSignature(
        bytes memory signature,
        uint256 chainId,
        uint256 price,
        uint256 priceTimestamp
    ) private view returns (bool) {
        bytes32 messageHash = keccak256(
            abi.encode(chainId, price, priceTimestamp)
        );
        address recoverAddress = messageHash.toEthSignedMessageHash().recover(
            signature
        );

        return recoverAddress == serverAddress;
    }

    function checkAdminSetEnvSignature(
        bytes memory signature_,
        address feeAddress_,
        address serverAddress_,
        address bridgeAddress_
    ) private returns (bool) {
        bytes32 messageHash = keccak256(
            abi.encode(sigNonce, feeAddress_, serverAddress_, bridgeAddress_)
        );
        address recoverAddress = messageHash.toEthSignedMessageHash().recover(
            signature_
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
        bytes32 messageHash = keccak256(abi.encode(sigNonce, chainId, user));
        address recoverAddress = messageHash.toEthSignedMessageHash().recover(
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
        address recoverAddress = messageHash.toEthSignedMessageHash().recover(
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
        address recoverAddress = messageHash.toEthSignedMessageHash().recover(
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
        address recoverAddress = messageHash.toEthSignedMessageHash().recover(
            signature
        );

        bool res = recoverAddress == superAdminAddress;
        if (res) {
            sigNonce++;
        }
        return res;
    }
}
