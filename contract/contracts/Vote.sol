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

    bytes32 public constant BRIDGE_ROLE = keccak256("BRIDGE_ROLE");

    IBridge public Bridge; // bridge 合约
    mapping(uint72 => mapping(bytes32 => Proposal)) public proposals; // destinationChainID + depositNonce => dataHash => Proposal
    mapping(uint72 => mapping(bytes32 => mapping(address => bool)))
        public hasVotedOnProposal; // destinationChainID + depositNonce => dataHash => relayerAddress => bool

    function initialize() public initializer {
        _grantRole(DEFAULT_ADMIN_ROLE, msg.sender);
    }

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
    ) external onlyRole(BRIDGE_ROLE) {}

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

    function getProposal(
        uint8 originChainID,
        uint64 depositNonce,
        bytes32 dataHash
    ) external {}
}
