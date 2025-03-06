// SPDX-License-Identifier: MIT
pragma solidity ^0.8.22;

import "./interface/IBridge.sol";
import "./interface/IVote.sol";
import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import "@openzeppelin/contracts/utils/cryptography/ECDSA.sol";
import "@openzeppelin/contracts/utils/cryptography/MessageHashUtils.sol";
import {AccessControl} from "@openzeppelin/contracts/access/AccessControl.sol";
//import "github.com/OpenZeppelin/openzeppelin-contracts/blob/release-v4.5/contracts/utils/cryptography/ECDSA.sol";

contract Vote is IVote, AccessControl, Initializable {
    using ECDSA for bytes32;
    using MessageHashUtils for bytes32;

    bytes32 public constant ADMIN_ROLE = keccak256("ADMIN_ROLE");
    bytes32 public constant BRIDGE_ROLE = keccak256("BRIDGE_ROLE");
    bytes32 public constant RELAYER_ROLE = keccak256("RELAYER_ROLE");

    IBridge public Bridge; // bridge 合约
    uint256 public totalProposal; // 提案数量，每个天加1
    uint256 public totalRelayer; // 总的relayer账户数量
    uint256 public relayerThreshold; // 提案可以通过的最少投票数量
    uint256 public expiry; // 开始投票后经过 expiry 的块数量后投票过期
    mapping(uint72 => mapping(bytes32 => Proposal)) public proposals; // destinationChainID + depositNonce => dataHash => Proposal
    mapping(uint72 => mapping(bytes32 => mapping(address => bool)))
    public hasVotedOnProposal; // destinationChainID + depositNonce => dataHash => relayerAddress => bool

    function initialize() public initializer {
        _grantRole(DEFAULT_ADMIN_ROLE, msg.sender);
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
    ) external onlyRole(BRIDGE_ROLE) {
        uint72 nonceAndID = (uint72(originDepositNonce) << 8) | uint72(originChainId);
        Proposal storage proposal = proposals[nonceAndID][dataHash];

//        require(resourceIDToHandlerAddress[resourceID] != address(0), "no handler for resourceID");
        require(uint(proposal.status) <= 1, "proposal already passed/executed/cancelled");
        require(!hasVotedOnProposal[nonceAndID][dataHash][msg.sender], "relayer already voted");

        if (uint(proposal.status) == 0) {
            ++totalProposal;
            proposals[nonceAndID][dataHash] = Proposal(
                resourceId,
                dataHash,
                new address[](1),
                new address[](0),
                ProposalStatus.Active,
                block.number
            );

            proposal.yesVotes[0] = msg.sender; // 索引 0 是创建提案的relayer
            emit ProposalEvent(originChainId, originDepositNonce, ProposalStatus.Active, resourceId, dataHash);
        } else {
            if (block.number-proposal.proposedBlock > expiry) {
                // if the number of blocks that has passed since this proposal was
                // submitted exceeds the expiry threshold set, cancel the proposal
                proposal.status = ProposalStatus.Cancelled;
                emit ProposalEvent(originChainId, originDepositNonce, ProposalStatus.Cancelled, resourceId, dataHash);
            } else {
                require(dataHash == proposal.dataHash, "datahash mismatch");
                proposal.yesVotes.push(msg.sender);
            }

        }
        if (proposal.status != ProposalStatus.Cancelled) {
            hasVotedOnProposal[nonceAndID][dataHash][msg.sender] = true;
            emit ProposalVote(originChainId, originDepositNonce, proposal.status, resourceId);

            // If _depositThreshold is set to 1, then auto finalize
            // or if _relayerThreshold has been exceeded
            if (relayerThreshold <= 1 || proposal.yesVotes.length >= relayerThreshold) {
                proposal.status = ProposalStatus.Passed;

                emit ProposalEvent(originChainId, originDepositNonce, ProposalStatus.Passed, resourceId, dataHash);
            }
        }
    }

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
    ) public onlyRole(BRIDGE_ROLE) {}

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
    ) external onlyRole(BRIDGE_ROLE) {}

    // 获取投票信息
    function getProposal(
        uint8 originChainID,
        uint64 depositNonce,
        bytes32 dataHash
    ) external {}

    // 获取 relayerThreshold
    function getRelayerThreshold() public view returns (uint256) {
        return relayerThreshold;
    }

    /**
        @notice 检查某地址是否是relayer账户
        @param relayer地址
     */
    function isRelayer(address relayer) external view returns (bool) {
        return hasRole(RELAYER_ROLE, relayer);
    }

}
