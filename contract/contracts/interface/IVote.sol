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

    struct Proposal {
        bytes32 resourceId;
        bytes32 dataHash;
        address[] yesVotes;
        address[] noVotes;
        ProposalStatus status;
        uint256 proposedBlock;
    }

    event ProposalEvent(
        uint8 indexed originChainID,
        uint64 indexed depositNonce,
        ProposalStatus indexed status,
        bytes32 resourceID,
        bytes32 dataHash
    );

    event ProposalVote(
        uint8 indexed originChainID,
        uint64 indexed depositNonce,
        ProposalStatus indexed status,
        bytes32 resourceID
    );

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
        uint64 originDepositNonce,
        bytes calldata data,
        bytes32 resourceID
    ) external;

    function getProposal(
        uint8 originChainID,
        uint64 depositNonce,
        bytes32 dataHash
    ) external;
}
