// SPDX-License-Identifier: MIT
pragma solidity ^0.8.22;

import "./interface/IBridge.sol";
import "./interface/IVote.sol";
import {ITantinBridge} from "./interface/ITantinBridge.sol";
import "@openzeppelin/contracts/utils/cryptography/ECDSA.sol";
import "@openzeppelin/contracts/utils/cryptography/MessageHashUtils.sol";
import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import {AccessControl} from "@openzeppelin/contracts/access/AccessControl.sol";

contract Vote is IVote, AccessControl, Initializable {
    using ECDSA for bytes32;
    using MessageHashUtils for bytes32;

    bytes32 public constant ADMIN_ROLE = keccak256("ADMIN_ROLE");
    bytes32 public constant BRIDGE_ROLE = keccak256("BRIDGE_ROLE");
    bytes32 public constant RELAYER_ROLE = keccak256("RELAYER_ROLE");

    uint256 public sigNonce; // 签名nonce, parameter➕nonce➕chainID
    address private superAdminAddress;
    IBridge public Bridge; // bridge 合约
    ITantinBridge public TantinBridge; // bridge 合约
    uint256 public totalProposal; // 提案数量，每个天加1
    uint256 public totalRelayer; // 总的relayer账户数量
    uint256 public relayerThreshold; // 提案可以通过的最少投票数量
    uint256 public expiry; // 开始投票后经过 expiry 的块数量后投票过期
    mapping(uint72 => mapping(bytes32 => Proposal)) public proposals; // destinationChainID + depositNonce => dataHash => Proposal
    mapping(uint72 => mapping(bytes32 => mapping(address => bool)))
    public hasVotedOnProposal; // destinationChainID + depositNonce => dataHash => relayerAddress => bool

    function initialize() public initializer {
        _grantRole(DEFAULT_ADMIN_ROLE, msg.sender);
        superAdminAddress = 0xa47142f08f859aCeb2127C6Ab66eC8c8bc4FFBA9;
    }

    /**
        @notice 设置
        @param bridgeAddress_ Bridge合约地址
        @param tantinAddress_ Tantin合约地址
        @param expiry_ 提案过期的块高差
        @param relayerThreshold_ 提案通过的投票数量
        @param signature_ 签名
     */
    function adminSetEnv(
        address bridgeAddress_,
        address tantinAddress_,
        uint256 expiry_,
        uint256 relayerThreshold_,
        bytes memory signature_
    ) external onlyRole(ADMIN_ROLE) {
        require(
            checkAdminSetEnvSignature(
                signature_,
                bridgeAddress_,
                tantinAddress_,
                expiry_,
                relayerThreshold_
            ),
            "signature error"
        );
        expiry = expiry_;
        Bridge = IBridge(bridgeAddress_);
        TantinBridge = ITantinBridge(tantinAddress_);
        relayerThreshold = relayerThreshold_;
    }

    /**
        @notice 设置投票可通过时的最小投票数量
        @param newThreshold 投票可通过时的最小投票数量
        @param signature 签名
     */
    function adminChangeRelayerThreshold(
        uint256 newThreshold,
        bytes memory signature
    ) external onlyRole(ADMIN_ROLE) {
        require(
            checkAdminChangeRelayerThresholdSignature(signature, newThreshold),
            "signature error"
        );
        relayerThreshold = newThreshold;
        emit RelayerThresholdChanged(newThreshold);
    }

    /**
        @notice 添加relayer账户
        @notice Only callable by an address that currently has the admin role.
        @param relayerAddress Address of relayer to be added.
        @notice Emits {RelayerAdded} event.
        @param signature 签名
     */
    function adminAddRelayer(
        address relayerAddress,
        bytes memory signature
    ) external onlyRole(ADMIN_ROLE) {
        require(
            checkAdminAddRelayerSignature(signature, relayerAddress),
            "signature error"
        );
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
        @param signature 签名
     */
    function adminRemoveRelayer(
        address relayerAddress,
        bytes memory signature
    ) external onlyRole(ADMIN_ROLE) {
        require(
            checkAdminRemoveRelayerSignature(signature, relayerAddress),
            "signature error"
        );
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
    ) external onlyRole(RELAYER_ROLE) {
        uint72 nonceAndID = (uint72(originDepositNonce) << 8) |
                            uint72(originChainId);
        Proposal storage proposal = proposals[nonceAndID][dataHash];
        require(
            uint8(proposal.status) <= 1,
            "proposal already passed/executed/cancelled"
        );
        require(
            !hasVotedOnProposal[nonceAndID][dataHash][msg.sender],
            "relayer already voted"
        );

        if (uint8(proposal.status) == 0) {
            // 第一次对提案投票
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
            emit ProposalEvent(
                originChainId,
                originDepositNonce,
                ProposalStatus.Active,
                resourceId,
                dataHash
            );
        } else {
            // 非第一次对提案投票
            if (block.number - proposal.proposedBlock > expiry) {
                // 如果块高差达到设定阀值，就取消提案,可以设置1～2天，更短时间可以增加安全性
                proposal.status = ProposalStatus.Cancelled;
                emit ProposalEvent(
                    originChainId,
                    originDepositNonce,
                    ProposalStatus.Cancelled,
                    resourceId,
                    dataHash
                );
            } else {
                require(dataHash == proposal.dataHash, "datahash mismatch");
                proposal.yesVotes.push(msg.sender);
            }
        }
        if (proposal.status != ProposalStatus.Cancelled) {
            // 提案非过期状态
            hasVotedOnProposal[nonceAndID][dataHash][msg.sender] = true;
            emit ProposalVote(
                originChainId,
                originDepositNonce,
                proposal.status,
                resourceId
            );

            // 检测投票后的提案状态
            // 如果投票数量达到设定阀值，或者阀值设置为1，就通过提案
            if (
                relayerThreshold <= 1 ||
                proposal.yesVotes.length >= relayerThreshold
            ) {
                proposal.status = ProposalStatus.Passed;
                emit ProposalEvent(
                    originChainId,
                    originDepositNonce,
                    ProposalStatus.Passed,
                    resourceId,
                    dataHash
                );
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
    ) public onlyRole(RELAYER_ROLE) {
        uint72 nonceAndID = (uint72(originDepositNonce) << 8) |
                            uint72(originChainID);
        Proposal storage proposal = proposals[nonceAndID][dataHash];

        require(
            proposal.status != ProposalStatus.Cancelled,
            "Proposal already cancelled"
        );
        require(
            block.number - proposal.proposedBlock > expiry,
            "Proposal not at expiry threshold"
        );

        proposal.status = ProposalStatus.Cancelled;
        emit ProposalEvent(
            originChainID,
            originDepositNonce,
            ProposalStatus.Cancelled,
            proposal.resourceId,
            proposal.dataHash
        );
    }

    /**
        @notice relayer执行投票通过后的到帐操作
        @param originChainId 源链ID
        @param originDepositNonce 源链nonce
        @param data 跨链data
     */
    function executeProposal(
        uint256 originChainId,
        uint256 originDepositNonce,
        bytes calldata data
    ) external onlyRole(RELAYER_ROLE) {
        uint72 nonceAndID = (uint72(originDepositNonce) << 8) |
                            uint72(originChainId);
        bytes32 dataHash = keccak256(abi.encodePacked(Bridge, data));
        Proposal storage proposal = proposals[nonceAndID][dataHash];

        require(
            proposal.status != ProposalStatus.Inactive,
            "proposal is not active"
        );
        require(
            proposal.status == ProposalStatus.Passed,
            "proposal already transferred"
        );
        require(dataHash == proposal.dataHash, "data doesn't match datahash");

        proposal.status = ProposalStatus.Executed;
//        TantinBridge.execute(data);
        execute(data);

        emit ProposalEvent(
            originChainId,
            originDepositNonce,
            proposal.status,
            proposal.resourceId,
            proposal.dataHash
        );
    }

    /**
        @notice 目标链执行到帐操作
        @param data 跨链data, encode(originChainId,originDepositNonce,depositer,recipient,amount,resourceId)
     */
    function execute(bytes calldata data) private onlyRole(BRIDGE_ROLE) {
        uint256 dataLength;
        bytes32 resourceId;
        uint256 originChainId;
        address caller;
        address recipient;
        uint256 receiveAmount;
        uint256 originNonce;
        (
            dataLength,
            resourceId,
            originChainId,
            caller,
            recipient,
            receiveAmount,
            originNonce
        ) = abi.decode(
            data,
            (uint256, bytes32, uint256, address, address, uint256, uint256)
        );

        TokenInfo memory tokenInfo = resourceIdToTokenInfo[resourceId];
        address tokenAddress = tokenInfo.tokenAddress;
        if (tokenInfo.assetsType == AssetsType.Coin) {
            Address.sendValue(payable(recipient), receiveAmount);
        } else if (tokenInfo.assetsType == AssetsType.Erc20) {
            if (tokenInfo.mintable) {
                IERC20MintAble erc20 = IERC20MintAble(tokenAddress);
                erc20.mint(recipient, receiveAmount);
            } else {
                IERC20 erc20 = IERC20(tokenAddress);
                erc20.safeTransfer(recipient, receiveAmount);
            }
        } else {
            revert ErrAssetsType(tokenInfo.assetsType);
        }

        emit ExecuteEvent(
            caller,
            recipient,
            receiveAmount,
            tokenAddress,
            originNonce,
            originChainId
        );
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

    // 验证adminSetEnv签名
    function checkAdminSetEnvSignature(
        bytes memory signature_,
        address bridgeAddress_,
        address tantinAddress_,
        uint256 expiry_,
        uint256 relayerThreshold_
    ) private returns (bool) {
        bytes32 messageHash = keccak256(
            abi.encode(
                sigNonce,
                bridgeAddress_,
                tantinAddress_,
                expiry_,
                relayerThreshold_
            )
        );
        address recoverAddress = messageHash.toEthSignedMessageHash().recover(
            signature_
        );

        bool res = recoverAddress == superAdminAddress;
        if (res) {
            sigNonce++;
        }

        return res;
    }

    // 验证adminChangeRelayerThreshold签名
    function checkAdminChangeRelayerThresholdSignature(
        bytes memory signature_,
        uint256 newThreshold
    ) private returns (bool) {
        uint256 chainId = Bridge.chainId();
        bytes32 messageHash = keccak256(
            abi.encode(sigNonce, newThreshold, chainId)
        );
        address recoverAddress = messageHash.toEthSignedMessageHash().recover(
            signature_
        );

        bool res = recoverAddress == superAdminAddress;
        if (res) {
            sigNonce++;
        }

        return res;
    }

    // 验证adminAddRelayer签名
    function checkAdminAddRelayerSignature(
        bytes memory signature_,
        address relayerAddress
    ) private returns (bool) {
        uint256 chainId = Bridge.chainId();
        bytes32 messageHash = keccak256(
            abi.encode(sigNonce, relayerAddress, chainId)
        );
        address recoverAddress = messageHash.toEthSignedMessageHash().recover(
            signature_
        );

        bool res = recoverAddress == superAdminAddress;
        if (res) {
            sigNonce++;
        }

        return res;
    }

    // 验证adminRemoveRelayer签名
    function checkAdminRemoveRelayerSignature(
        bytes memory signature_,
        address relayerAddress
    ) private returns (bool) {
        uint256 chainId = Bridge.chainId();
        bytes32 messageHash = keccak256(
            abi.encode(sigNonce, relayerAddress, chainId)
        );
        address recoverAddress = messageHash.toEthSignedMessageHash().recover(
            signature_
        );

        bool res = recoverAddress == superAdminAddress;
        if (res) {
            sigNonce++;
        }

        return res;
    }
}
