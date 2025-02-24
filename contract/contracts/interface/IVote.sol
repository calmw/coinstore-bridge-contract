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
}
