// SPDX-License-Identifier: MIT
pragma solidity ^0.8.22;

import "./utils/Pausable.sol";
import {ECDSA} from "./lib/ECDSA.sol";
import {IVote} from "./interface/IVote.sol";
import {IBridge} from "./interface/IBridge.sol";
import {ITantinBridge} from "./interface/ITantinBridge.sol";
import {IERC20MintAble} from "./interface/IERC20MintAble.sol";
import {AccessControl} from "@openzeppelin/contracts/access/AccessControl.sol";

contract Bridge is IBridge, Pausable, AccessControl {
    using ECDSA for bytes32;

    bytes32 public constant ADMIN_ROLE = keccak256("ADMIN_ROLE");
    bytes32 public constant BRIDGE_ROLE = keccak256("BRIDGE_ROLE");

    uint256 public sigNonce; // 签名nonce, parameter➕nonce➕chainID
    address private superAdminAddress;
    IVote public Vote; // vote 合约
    uint256 public chainId; // 自定义链ID
    uint256 public chainType; // 自定义链类型， 1 EVM 2 Tron
    mapping(uint256 => uint256) public depositCounts; // destinationChainID => number of deposits
    mapping(bytes32 => TokenInfo) public resourceIdToTokenInfo; //  resourceID => 设置的Token信息
    mapping(uint256 => mapping(uint256 => DepositRecord)) public depositRecords; // depositNonce => (destinationChainId => Deposit Record)

    constructor() {
        _grantRole(DEFAULT_ADMIN_ROLE, msg.sender);
        superAdminAddress = 0x3942FdA93c573E2ce9e85B0bB00Ba98a144f27f6;
    }

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
    ) external onlyRole(ADMIN_ROLE) {
        require(
            checkAdminSetEnvSignature(
                signature_,
                voteAddress_,
                chainId_,
                chainType_
            ),
            "signature error"
        );
        Vote = IVote(voteAddress_);
        chainId = chainId_;
        chainType = chainType_;
    }

    /**
        @notice 暂停跨链、提案的的创建与投票和目标链执行操作
        @param signature 签名
     */
    function adminPauseTransfers(
        bytes memory signature
    ) external onlyRole(ADMIN_ROLE) {
        require(
            checkAdminPauseTransfersSignature(signature),
            "signature error"
        );
        _pause();
    }

    /**
        @notice 开启跨链、提案的的创建与投票和目标链执行操作
        @param signature 签名
     */
    function adminUnpauseTransfers(
        bytes memory signature
    ) external onlyRole(ADMIN_ROLE) {
        require(
            checkAdminUnpauseTransfersSignature(signature),
            "signature error"
        );
        _unpause();
    }

    /**
        @notice resource设置
        @param resourceID 跨链的resourceID
        @param assetsType 该币的类型
        @param tokenAddress 对应的token合约地址，coin为0地址
        @param decimal 该币的精度
        @param fee 该币的跨链费用,折合U的数量
        @param pause 该币种是否在黑名单中/是否允许跨链。币种黑名单/禁止该币种跨链
        @param burnable 该币是否burn
        @param mintable 该币是否mint
     */
    function adminSetResource(
        bytes32 resourceID,
        AssetsType assetsType,
        address tokenAddress,
        uint256 decimal,
        uint256 fee,
        bool pause,
        bool burnable,
        bool mintable,
        bytes memory signature
    ) external onlyRole(ADMIN_ROLE) {
        require(
            checkAdminSetResourceSignature(
                signature,
                resourceID,
                assetsType,
                tokenAddress,
                decimal,
                fee,
                pause,
                burnable,
                mintable
            ),
            "signature error"
        );
        resourceIdToTokenInfo[resourceID] = TokenInfo(
            assetsType,
            tokenAddress,
            pause,
            decimal,
            fee,
            burnable,
            mintable
        );

        emit SetResource(
            resourceID,
            tokenAddress,
            decimal,
            fee,
            pause,
            burnable,
            mintable
        );
    }

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
    ) external payable whenNotPaused onlyRole(BRIDGE_ROLE) {
        // 检测resource ID是否设置
        TokenInfo memory tokenInfo = resourceIdToTokenInfo[resourceId];
        require(uint8(tokenInfo.assetsType) > 0, "resourceId not exist");
        // 检测resourceId/token是否被暂停跨链
        require(!tokenInfo.pause, "service suspended");

        uint256 depositNonce = ++depositCounts[destinationChainId];

        depositRecords[destinationChainId][depositNonce] = DepositRecord(
            destinationChainId,
            msg.sender,
            resourceId,
            block.timestamp,
            data
        );

        emit Deposit(destinationChainId, resourceId, depositNonce, data);
    }

    // 获取跨链费用
    function getFeeByResourceId(
        bytes32 resourceId
    ) public view returns (uint256) {
        TokenInfo memory tokenInfo = resourceIdToTokenInfo[resourceId];
        require(uint8(tokenInfo.assetsType) > 0, "resourceId not exist");
        return tokenInfo.fee;
    }

    // 由resourceId获取token信息
    function getTokenInfoByResourceId(
        bytes32 resourceId
    ) public view returns (uint8, address, bool, uint256, uint256, bool, bool) {
        TokenInfo memory token = resourceIdToTokenInfo[resourceId];
        return (
            uint8(token.assetsType),
            token.tokenAddress,
            token.pause,
            token.decimal,
            token.fee,
            token.burnable,
            token.mintable
        );
    }

    // 验证adminSetEnv签名
    function checkAdminSetEnvSignature(
        bytes memory signature_,
        address voteAddress_,
        uint256 chainId_,
        uint256 chainType_
    ) private returns (bool) {
        bytes32 messageHash = keccak256(
            abi.encode(sigNonce, voteAddress_, chainId_, chainType_)
        );
        address recoverAddress = messageHash.toEthSignedMessageHash().recoverSigner(
            signature_
        );
        bool res = recoverAddress == superAdminAddress;
        if (res) {
            sigNonce++;
        }
        return res;
    }

    // 验证adminPauseTransfers签名
    function checkAdminPauseTransfersSignature(
        bytes memory signature
    ) private returns (bool) {
        bytes32 messageHash = keccak256(abi.encode(sigNonce, chainId));
        address recoverAddress = messageHash.toEthSignedMessageHash().recoverSigner(
            signature
        );
        bool res = recoverAddress == superAdminAddress;
        if (res) {
            sigNonce++;
        }
        return res;
    }

    // 验证adminUnpauseTransfers签名
    function checkAdminUnpauseTransfersSignature(
        bytes memory signature
    ) private returns (bool) {
        bytes32 messageHash = keccak256(abi.encode(sigNonce, chainId));
        address recoverAddress = messageHash.toEthSignedMessageHash().recoverSigner(
            signature
        );
        bool res = recoverAddress == superAdminAddress;
        if (res) {
            sigNonce++;
        }
        return res;
    }

    // 验证adminSetResource签名
    function checkAdminSetResourceSignature(
        bytes memory signature,
        bytes32 resourceID,
        AssetsType assetsType,
        address tokenAddress,
        uint256 decimal,
        uint256 fee,
        bool pause,
        bool burnable,
        bool mintable
    ) private returns (bool) {
        bytes32 messageHash = keccak256(
            abi.encode(
                sigNonce,
                chainId,
                resourceID,
                assetsType,
                tokenAddress,
                decimal,
                fee,
                pause,
                burnable,
                mintable
            )
        );
        address recoverAddress = messageHash.toEthSignedMessageHash().recoverSigner(
            signature
        );
        bool res = recoverAddress == superAdminAddress;
        if (res) {
            sigNonce++;
        }
        return res;
    }
}
