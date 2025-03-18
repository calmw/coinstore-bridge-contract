// SPDX-License-Identifier: MIT
pragma solidity ^0.8.22;

import "./interface/IBridge.sol";
import "./interface/IERC20.sol";
import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import "./utils/Pausable.sol";
import {AccessControl} from "@openzeppelin/contracts/access/AccessControl.sol";
import {IVote} from "./interface/IVote.sol";

contract Bridge is IBridge, Pausable, AccessControl, Initializable {
    bytes32 public constant ADMIN_ROLE = keccak256("ADMIN_ROLE");
    bytes32 public constant VOTE_ROLE = keccak256("VOTE_ROLE");

    IVote public Vote; // vote 合约
    uint256 public chainId; // 自定义链ID
    uint256 public chainType; // 自定义链类型， 1 EVM 2 Tron
    mapping(uint256 => uint256) public depositCounts; // destinationChainID => number of deposits
    mapping(bytes32 => address) public resourceIdToContractAddress; // resourceID => 业务合约地址(tantin address)
    mapping(address => bytes32) public contractAddressToResourceID; // 业务合约地址(tantin address) => resourceID
    mapping(bytes32 => TokenInfo) public resourceIdToTokenInfo; //  resourceID => 设置的Token信息
    mapping(bytes32 => bytes4) public resourceIdToExecuteSig; //  resourceID => tantin execute sig
    mapping(uint256 => mapping(uint256 => DepositRecord)) public depositRecords; // depositNonce => (destinationChainId => Deposit Record)

    function initialize() public initializer {
        _grantRole(ADMIN_ROLE, msg.sender);
        _grantRole(DEFAULT_ADMIN_ROLE, msg.sender);
    }

    /**
        @notice 设置
        @param voteAddress_ 投票合约地址
        @param chainId_ 链ID
     */
    function adminSetEnv(
        address voteAddress_,
        uint256 chainId_,
        uint256 chainType_
    ) external onlyRole(ADMIN_ROLE) {
        Vote = IVote(voteAddress_);
        chainId = chainId_;
        chainType = chainType_;
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
        // 检测resourceId/token是否被暂停跨链
        require(!tokenInfo.pause, "service suspended");

        uint256 depositNonce = ++depositCounts[destinationChainId];

        depositRecords[destinationChainId][depositNonce] = DepositRecord(
            destinationChainId,
            msg.sender,
            resourceId,
            data
        );

        emit Deposit(destinationChainId, resourceId, depositNonce, data);
    }

    /**
        @notice relayer执行投票通过后的到帐操作
        @param originChainId 源链ID
        @param originDepositNonce 源链nonce
        @param resourceId 跨链的resourceID
        @param data 跨链data
     */
    function execute(
        uint256 originChainId,
        bytes32 resourceId,
        uint256 originDepositNonce,
        bytes calldata data
    ) external onlyRole(VOTE_ROLE) whenNotPaused {
        bytes32 resourceId;
        uint256 originChainId;
        address caller;
        address recipient;
        uint256 receiveAmount;
        uint256 originNonce;
        (resourceId, originChainId, caller, recipient, receiveAmount, originNonce) = abi.decode(data, (bytes32, uint256, address, address, uint256, uint256));

        //
//    resourceIdToContractAddress
    }

    // 获取自定义链ID
    function getChainId() public view returns (uint256) {
        return chainId;
    }

    // 获取自定义链类型ID
    function getChainTypeId() public view returns (uint256) {
        return chainType;
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

    // 由resourceId获取tantin address信息
    function getContractAddressByResourceId(
        bytes32 resourceId
    ) public view returns (address) {
        return resourceIdToContractAddress[resourceId];
    }
}
