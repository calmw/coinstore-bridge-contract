// SPDX-License-Identifier: MIT
pragma solidity ^0.8.22;

import "./interface/IBridge.sol";
import "./interface/IERC20.sol";
import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import "@openzeppelin/contracts/utils/Pausable.sol";
import "@openzeppelin/contracts/utils/cryptography/ECDSA.sol";
import "@openzeppelin/contracts/utils/cryptography/MessageHashUtils.sol";
import {AccessControl} from "@openzeppelin/contracts/access/AccessControl.sol";
import {IVote} from "./interface/IVote.sol";
//import "github.com/OpenZeppelin/openzeppelin-contracts/blob/release-v4.5/contracts/utils/cryptography/ECDSA.sol";

contract Bridge is IBridge, Pausable, AccessControl, Initializable {
    using ECDSA for bytes32;
    using MessageHashUtils for bytes32;

    bytes32 public constant ADMIN_ROLE = keccak256("ADMIN_ROLE");
    bytes32 public constant VOTE_ROLE = keccak256("VOTE_ROLE");
    bytes32 public constant RELAYER_ROLE = keccak256("RELAYER_ROLE");

    IVote public Vote; // vote 合约
    uint256 public chainId; // 自定义链ID
    uint256 public fee; // 跨链费用,当前设置的收手续费模式为固定数量的coin
    uint256 public totalRelayer; // 总的relayer账户数量
    uint256 public relayerThreshold; // 提案可以通过的最少投票数量
    uint256 public expiry; // 开始投票后经过 expiry 的块数量后投票过期
    mapping(bytes32 => address) public resourceIdToContractAddress; // resourceID => 业务合约地址(tantin address)
    mapping(address => bytes32) public contractAddressToResourceID; // 业务合约地址(tantin address) => resourceID
    mapping(bytes32 => TokenInfo) public resourceIdToTokenInfo; // 业务合约地址(tantin address) => resourceID
    mapping(address => bool) public contractWhitelist; // token contract address => 是否在白名单中
    mapping(uint8 => mapping(uint64 => DepositRecord)) public depositRecords; // depositNonce => Deposit Record

    modifier onlyAdminOrRelayer() {
        require(
            hasRole(ADMIN_ROLE, msg.sender) ||
                hasRole(RELAYER_ROLE, msg.sender),
            "sender is not relayer or admin"
        );
        _;
    }

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
        @notice 添加relayer账户
        @notice Only callable by an address that currently has the admin role.
        @param relayerAddress Address of relayer to be added.
        @notice Emits {RelayerAdded} event.
     */
    function adminAddRelayer(
        address relayerAddress
    ) external onlyRole(ADMIN_ROLE) {
        require(
            !hasRole(RELAYER_ROLE, relayerAddress),
            "addr already has relayer role!"
        );
        grantRole(RELAYER_ROLE, relayerAddress);
        emit RelayerAdded(relayerAddress);
        totalRelayer++;
    }

    /**
        @notice 删除relayer账户
        @notice Only callable by an address that currently has the admin role.
        @param relayerAddress Address of relayer to be removed.
        @notice Emits {RelayerRemoved} event.
     */
    function adminRemoveRelayer(
        address relayerAddress
    ) external onlyRole(ADMIN_ROLE) {
        require(
            hasRole(RELAYER_ROLE, relayerAddress),
            "addr doesn't have relayer role!"
        );
        revokeRole(RELAYER_ROLE, relayerAddress);
        emit RelayerRemoved(relayerAddress);
        totalRelayer--;
    }

    /**
        @notice 设置投票可通过时的最小投票数量
        @param newThreshold 投票可通过时的最小投票数量
     */
    function adminChangeRelayerThreshold(
        uint newThreshold
    ) external onlyRole(ADMIN_ROLE) {
        relayerThreshold = newThreshold;
        emit RelayerThresholdChanged(newThreshold);
    }

    /**
        @notice 管理员设置跨链费用
        @param newFee 跨链手续费,单位wei
     */
    function adminChangeFee(uint256 newFee) external onlyRole(ADMIN_ROLE) {}

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
        @param tantinAddress 对应的tantin业务合约地址
        @param tokenAddress 对应的token合约地址，coin为0地址
        @param executeFunctionSig tantin业务合约执行到帐操作的方法签名
     */
    function adminSetResource(
        bytes32 resourceID,
        address tantinAddress,
        address tokenAddress,
        bytes4 executeFunctionSig
    ) external onlyRole(ADMIN_ROLE) {}

    /**
        @notice 资产跨链
        @param destinationChainId 目标链ID
        @param resourceID 跨链的resourceID
        @param data   跨链data
     */
    function deposit(
        uint256 destinationChainId,
        bytes32 resourceID,
        bytes calldata data
    ) external whenNotPaused {}

    /**
        @notice relayer执行投票通过后的到帐操作
        @param originChainID 源链ID
        @param originDepositNonce 源链nonce
        @param resourceID 跨链的resourceID
        @param dataHash dataHash
     */
    function voteProposal(
        uint256 originChainID,
        uint256 originDepositNonce,
        bytes32 resourceID,
        bytes32 dataHash
    ) external onlyRole(RELAYER_ROLE) whenNotPaused {}

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
    ) public onlyAdminOrRelayer {}

    /**
        @notice relayer执行投票通过后的到帐操作
        @param originChainID 源链ID
        @param originDepositNonce 源链nonce
        @param resourceID 跨链的resourceID
        @param data 跨链data
     */
    function executeProposal(
        uint256 originChainID,
        uint64 originDepositNonce,
        bytes calldata data,
        bytes32 resourceID
    ) external onlyRole(RELAYER_ROLE) whenNotPaused {}

    // 获取自定义链ID
    function getChainId() public view returns (uint256) {
        return chainId;
    }

    // 由resourceId获取token信息
    function getToeknInfoByResourceId(
        bytes32 resourceID
    ) public view returns (uint256, address, bool) {
        return (
            uint256(resourceIdToTokenInfo[resourceID].assetsType),
            resourceIdToTokenInfo[resourceID].tokenAddress,
            resourceIdToTokenInfo[resourceID].burnable
        );
    }

    /**
        @notice 检查某地址是否是relayer账户
        @param relayer地址
     */
    function isRelayer(address relayer) external view returns (bool) {
        return hasRole(RELAYER_ROLE, relayer);
    }
}
