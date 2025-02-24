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

    IVote public vote; // vote 合约
    uint256 public chainId; // 自定义链ID
    uint256 public fee; // 跨链费用
    uint256 public totalRelayer; // 总的relayer数量
    uint256 public relayerThreshold; // 投票可以通过的最少relayer数量
    uint256 public expiry; // 开始投票后经过 expiry 的块数量后投票过期
    mapping(bytes32 => address) public resourceIdToContractAddress; // resourceID => token contract address
    mapping(address => bytes32) public contractAddressToResourceID; // token contract address => resourceID
    mapping(address => bool) public contractWhitelist; // token contract address => 是否在白名单中
    mapping(uint8 => mapping(uint64 => DepositRecord)) public depositRecords; // depositNonce => Deposit Record

    function initialize() public initializer {
        _grantRole(DEFAULT_ADMIN_ROLE, msg.sender);
    }

    function deposit(
        uint256 destinationChainId,
        bytes32 resourceID,
        bytes calldata data
    ) external whenNotPaused {}

//    function getChainId() public view returns (uint256) {
//        return chainId;
//    }

    function getToeknAddressByResourceID(bytes32 resourceID) public view returns (address) {
        return resourceIdToContractAddress[resourceID];
    }

    function vote() public  {}

    function voteProposal() public  onlyRole(RELAYER_ROLE){}
    function cancelProposal() public  onlyRole(ADMIN_ROLE){}
    function executeProposal() public  {}
    function getProposal(uint8 originChainID, uint64 depositNonce, bytes32 dataHash) external view returns (Proposal memory) {}

    function adminChangeFee(uint newFee)external onlyRole(ADMIN_ROLE) {}

    function adminWithdraw(
        address handlerAddress,
        address tokenAddress,
        address recipient,
        uint256 amountOrTokenID
    ) external onlyRole(ADMIN_ROLE) {}
}
