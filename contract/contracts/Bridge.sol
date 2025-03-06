// SPDX-License-Identifier: MIT
pragma solidity ^0.8.22;

import "./interface/IBridge.sol";
import "./interface/IERC20.sol";
import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import "@openzeppelin/contracts/utils/Pausable.sol";
import {AccessControl} from "@openzeppelin/contracts/access/AccessControl.sol";
import {IVote} from "./interface/IVote.sol";

contract Bridge is IBridge, Pausable, AccessControl, Initializable {
    bytes32 public constant ADMIN_ROLE = keccak256("ADMIN_ROLE");
    bytes32 public constant VOTE_ROLE = keccak256("VOTE_ROLE");
//    bytes32 public constant RELAYER_ROLE = keccak256("RELAYER_ROLE");

    IVote public Vote; // vote 合约
    uint256 public chainId; // 自定义链ID
//    uint256 public relayerThreshold; // 提案可以通过的最少投票数量
//    uint256 public expiry; // 开始投票后经过 expiry 的块数量后投票过期
    mapping(uint256 => uint64) public depositCounts; // destinationChainID => number of deposits
    mapping(bytes32 => address) public resourceIdToContractAddress; // resourceID => 业务合约地址(tantin address)
    mapping(address => bytes32) public contractAddressToResourceID; // 业务合约地址(tantin address) => resourceID
    mapping(bytes32 => TokenInfo) public resourceIdToTokenInfo; //  resourceID => 设置的Token信息
    mapping(bytes32 => bytes4) public resourceIdToExecuteSig; //  resourceID => tantin execute sig
    mapping(uint8 => mapping(uint64 => DepositRecord)) public depositRecords; // depositNonce => Deposit Record

    function initialize() public initializer {
        _grantRole(DEFAULT_ADMIN_ROLE, msg.sender);
    }

    /**
        @notice 暂停跨链、提案的的创建与投票和目标链执行操作
     */
    function adminPauseTransfers() external onlyRole(ADMIN_ROLE) {
        _pause();
    }

    /**
        @notice 开启跨链、提案的的创建与投票和目标链执行操作
     */
    function adminUnpauseTransfers() external onlyRole(ADMIN_ROLE) {
        _unpause();
    }

    /**
        @notice 提取跨链桥coin资产
        @param recipient 资产接受者地址
        @param amount 提取数量,单位wei
     */
    function adminWithdraw(
        address recipient,
        uint256 amount
    ) external onlyRole(ADMIN_ROLE) {}

    /**
        @notice resource设置
        @param resourceID 跨链的resourceID
        @param assetsType 该币的类型
        @param tokenAddress 对应的token合约地址，coin为0地址
        @param fee 该币的跨链费用
        @param pause 该币种是否在黑名单中/是否允许跨链。币种黑名单/禁止该币种跨链
        @param tantinAddress 对应的tantin业务合约地址
        @param executeFunctionSig tantin业务合约执行到帐操作的方法签名
     */
    function adminSetResource(
        bytes32 resourceID,
        AssetsType assetsType,
        address tokenAddress,
        uint256 fee,
        bool pause,
        address tantinAddress,
        bytes4 executeFunctionSig
    ) external onlyRole(ADMIN_ROLE) {
        resourceIdToTokenInfo[resourceID] = TokenInfo(
            assetsType,
            tokenAddress,
            pause,
            fee
        );
        resourceIdToContractAddress[resourceID] = tantinAddress;
        resourceIdToExecuteSig[resourceID] = executeFunctionSig;

        emit SetResource(
            resourceID,
            tokenAddress,
            fee,
            pause,
            tantinAddress,
            executeFunctionSig
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
    ) external payable whenNotPaused {
        // 检测resource ID是否设置
        TokenInfo memory tokenInfo = resourceIdToTokenInfo[resourceId];
        require(uint8(tokenInfo.assetsType) > 0, "resourceId not exist");
        // 检测跨链费用
        require(tokenInfo.fee == msg.value, "incorrect fee supplied");
        // 检测resourceId/token是否暂停跨链
        require(!tokenInfo.pause, "service suspended");

        uint64 depositNonce = ++depositCounts[destinationChainId];

        emit Deposit(destinationChainId, resourceId, depositNonce, data);
    }

    /**
        @notice relayer执行投票通过后的到帐操作
        @param originChainID 源链ID
        @param originDepositNonce 源链nonce
        @param resourceID 跨链的resourceID
        @param data 跨链data
     */
    function execute(
        uint256 originChainID,
        uint64 originDepositNonce,
        bytes calldata data,
        bytes32 resourceID
    ) external onlyRole(VOTE_ROLE) whenNotPaused {}

    // 获取自定义链ID
    function getChainId() public view returns (uint256) {
        return chainId;
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
    function getToeknInfoByResourceId(
        bytes32 resourceID
    ) public view returns (uint256, address, bool) {
        return (
            uint256(resourceIdToTokenInfo[resourceID].assetsType),
            resourceIdToTokenInfo[resourceID].tokenAddress,
            resourceIdToTokenInfo[resourceID].pause
        );
    }

}
