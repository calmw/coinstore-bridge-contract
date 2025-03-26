// SPDX-License-Identifier: MIT
pragma solidity ^0.8.22;

import "./interface/IBridge.sol";
import "./interface/IVote.sol";
import {AccessControl} from "@openzeppelin/contracts/access/AccessControl.sol";

contract Vote is IVote, AccessControl {

    IBridge public Bridge; // bridge 合约
    address public br;
    uint256 public totalRelayer; // 总的relayer账户数量
    uint256 public relayerThreshold; // 提案可以通过的最少投票数量
    uint256 public expiry; // 开始投票后经过 expiry 的块数量后投票过期
    mapping(uint72 => mapping(bytes32 => Proposal)) public proposals; // destinationChainID + depositNonce => dataHash => Proposal
    mapping(uint72 => mapping(bytes32 => mapping(address => bool)))
    public hasVotedOnProposal; // destinationChainID + depositNonce => dataHash => relayerAddress => bool

    constructor() {}

    /**
        @notice 设置
        @param bridgeAddress_ bridge合约地址
        @param expiry_ 提案过期的块高差
        @param relayerThreshold_ 提案通过的投票数量
     */
    function adminSetEnv(
        address bridgeAddress_
    ) external {
        Bridge = IBridge(bridgeAddress_);
        br = bridgeAddress_;
    }


    function executeProposal(
        uint256 originChainId,
        uint256 originDepositNonce,
        bytes calldata data,
        bytes32 resourceId
    ) external returns(bytes32,uint72) {
        uint72 nonceAndID = (uint72(originDepositNonce) << 8) |
                            uint72(originChainId);
        bytes32 dataHash = keccak256(abi.encodePacked(Bridge, data));
       returns(dataHash,nonceAndID);
    }

    function TestdataHash(
        uint256 originChainID,
        uint256 originDepositNonce,
        bytes memory data
    ) public view returns (uint72, bytes32) {
        uint72 nonceAndID = (uint72(originDepositNonce) << 8) |
                            uint72(originChainID);
        bytes32 dataHash = keccak256(abi.encodePacked(Bridge, data));
        return (nonceAndID, dataHash);
    }

    // 获取投票信息
    function getProposal(
        uint256 originChainID,
        uint256 depositNonce,
        bytes32 dataHash
    ) external view returns (Proposal memory) {
        uint72 nonceAndID = (uint72(depositNonce) << 8) | uint72(originChainID);
        return proposals[nonceAndID][dataHash];
    }
}
