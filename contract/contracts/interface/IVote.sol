// SPDX-License-Identifier: MIT
pragma solidity ^0.8.22;

interface IVote {
    enum Vote {
        No,
        Yes
    }

    enum ProposalStatus {
        Inactive,
        Active,
        Passed,
        Executed,
        Cancelled
    }

    event RelayerThresholdChanged(uint indexed newThreshold);
    event RelayerAdded(address indexed relayer);
    event RelayerRemoved(address indexed relayer);

    event ProposalEvent(
        uint256 indexed originChainID,
        uint256 indexed depositNonce,
        ProposalStatus indexed status,
        bytes32 resourceID,
        bytes32 dataHash
    );

    event ProposalVote(
        uint256 indexed originChainID,
        uint256 indexed depositNonce,
        ProposalStatus indexed status,
        bytes32 resourceID
    );

    event ExecuteEvent(
        address indexed depositer,
        address indexed recipient,
        uint256 indexed amount,
        address tokenAddress,
        uint256 originDepositNonce,
        uint256 originChainId
    );

    struct Proposal {
        bytes32 resourceId;
        bytes32 dataHash;
        address[] yesVotes;
        address[] noVotes;
        ProposalStatus status;
        uint256 proposedBlock;
    }

    function voteProposal(
        uint256 originChainID,
        uint256 originDepositNonce,
        bytes32 resourceID,
        bytes32 dataHash
    ) external;

    function cancelProposal(
        uint256 originChainID,
        uint256 originDepositNonce,
        bytes32 dataHash
    ) external;

    function executeProposal(
        uint256 originChainID,
        uint256 originDepositNonce,
        bytes calldata data
    ) external;

    function getProposal(
        uint256 originChainID,
        uint256 depositNonce,
        bytes32 dataHash
    ) external returns (Proposal memory);
}
