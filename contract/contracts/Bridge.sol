// SPDX-License-Identifier: MIT
pragma solidity ^0.8.22;

import "./interface/IAuth.sol";
import "./interface/IERC20.sol";
import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import {AccessControl} from "@openzeppelin/contracts/access/AccessControl.sol";
import "@openzeppelin/contracts/utils/cryptography/ECDSA.sol";
import "@openzeppelin/contracts/utils/cryptography/MessageHashUtils.sol";
//import "github.com/OpenZeppelin/openzeppelin-contracts/blob/release-v4.5/contracts/utils/cryptography/ECDSA.sol";

contract Bridge is AccessControl, Initializable {
    using ECDSA for bytes32;
    using MessageHashUtils for bytes32;

    struct ProposeInfo {
        uint256 Id;
        address creator;
    }

    struct VoteInfo {
        uint256 pId;
        uint256 voteDetail;
    }

    struct VoteDetails {
        uint256 pId;
        address user;
        uint256 isSupport;
    }

    //    event ProposeCreate(uint256 voteId, uint256 fee, address user); // 发起投票
    event ProposeCreate(
        uint256 voteId, // 提案ID
        uint256 chainId,
        uint256 fee,
        address user,
        uint256 ctime
    ); // 发起投票

    event Release(uint256 orderId, address user); // 返还投票发起人押金或者投票者投票所消耗的费用

    event VotePropose(
        uint256 orderId,
        uint256 fee,
        address user,
        bool isSupport
    ); // 用户投票

    event AdminSetProposeStatus(uint256 orderId, uint256 status); // 管理员设置提案状态
    event ReturnFee(
        uint256 orderId,
        address user,
        uint256 amount,
        uint256 IsCreator
    ); // 投票过期会被采纳后的退款

    IAuth public auth; // SBT认证合约
    IERC20 public tox; //TOX合约

    address public serverAddress; // 服务端签名地址

    uint256 public createFee; // 发起投票所需要的费用
    uint256 public voteFee; // 投票所需要的费用

    mapping(uint256 => uint256) public proposeCreatedAt; // 某个提案总的创建/支付时间
    mapping(uint256 => uint256) public proposeStatus; // 某个提案总的状态， 0 未开始/未支付；1进行中，2 被采纳，3 过期
    mapping(uint256 => uint256) public voteCount; // 某个提案总的投票数
    mapping(uint256 => uint256) public supportCount; // 某个提案投支持票的数量
    mapping(uint256 => uint256) public unSupportCount; // 某个提案投反对票的数量

    mapping(address => uint256) public userProposeCont; // 用户地址 => 用户投票的总数量
    mapping(address => uint256[]) public userProposeIds; // 用户地址 => 用户投票的所有提案ID

    mapping(address => mapping(uint256 => bool)) public userHasVote; // 用户地址 => (提案ID=>是否投过票)
    mapping(uint256 => address) public proposeCreator; // 提案ID => 用户地址
    mapping(uint256 => address[]) public proposeVoters; // 提案ID => 对该提案投票的所有用户地址集合
    mapping(address => mapping(uint256 => bool)) public userHasReturnVoteFee; // 用户地址 => (提案ID=>是否退还过投过票费用)

    mapping(address => uint256) public userVoteCount; // 用户钱包地址 => 用户累计投票的次数
    mapping(address => uint256[]) public userVoteProposes; // 用户钱包地址 => 用户参与的提案ID集合

    mapping(address => mapping(uint256 => uint256)) public userVoteDetail; // 用户地址 => (提案ID=>投的赞成还是反对) 1 赞成，2 反对

    mapping(address => mapping(uint256 => uint256)) public userVoteFee; // 用户地址 => (提案ID=>累计投票费用)
    uint256[] public allProposeIds; // 所有提案ID
    uint256 public countProposeIds; // 所有提案ID的数量

    mapping(uint256 => VoteDetails[]) public proposeVoteList; // 某提案的投票情况
    mapping(uint256 => bool) public hasCreatePropose; // 创建提案是否同步过
    mapping(uint256 => mapping(bytes32 => mapping(uint256 => bool)))
        public hasMatchCreatePropose; // 创建提案是否同步过
    bytes32 public constant SERVER_ROLE = keccak256("SERVER_ROLE");
    mapping(uint256 => uint256) public proposeCreatedChainId; // 某个提案总的创建/支付的链
    mapping(address => mapping(uint256 => bool)) public userHasReturnCreateFee; // 用户地址 => (提案ID=>是否退还创建提案费用)

    function initialize(
        address tox_,
        address serverAddress_,
        uint256 createFee_,
        uint256 voteFee_
    ) public initializer {
        _grantRole(DEFAULT_ADMIN_ROLE, msg.sender);
        tox = IERC20(tox_);
        createFee = createFee_;
        voteFee = voteFee_;
        serverAddress = serverAddress_;
    }

    function adminSetEnv(
        address tox_,
        address serverAddress_,
        uint256 createFee_,
        uint256 voteFee_
    ) public onlyRole(DEFAULT_ADMIN_ROLE) {
        tox = IERC20(tox_);
        createFee = createFee_;
        voteFee = voteFee_;
        serverAddress = serverAddress_;
    }

    // 创建/支付提案
    function createPropose(
        uint256 proposeId_,
        uint256 createTimestamp_,
        bytes calldata signature_
    ) external {
        require(
            checkCreateSignature(
                signature_,
                block.chainid,
                proposeId_,
                createTimestamp_
            ),
            "signature error"
        );
        require(!hasCreatePropose[proposeId_], "already created");
        require(
            checkCreateSignature(
                signature_,
                block.chainid,
                proposeId_,
                createTimestamp_
            ),
            "signature error"
        );
        require(block.timestamp - 570 <= createTimestamp_, "payment timeout");
        require(tox.balanceOf(msg.sender) >= createFee, "insufficient balance"); // 检查余额
        tox.transferFrom(msg.sender, address(this), createFee);

        proposeCreator[proposeId_] = msg.sender;
        proposeStatus[proposeId_] = 1;
        userProposeCont[msg.sender]++;
        userProposeIds[msg.sender].push(proposeId_);
        proposeCreatedAt[proposeId_] = block.timestamp;
        allProposeIds.push(proposeId_);
        countProposeIds++;
        hasCreatePropose[proposeId_] = true;

        emit ProposeCreate(
            proposeId_,
            block.chainid,
            createFee,
            msg.sender,
            block.timestamp
        );
    }

    // 对提案进行投票
    function vote(
        uint256 proposeId_,
        uint256 proposeStatus_,
        bool isSupport_,
        bytes calldata signature_
    ) external {
        require(
            checkVoteSignature(
                signature_,
                block.chainid,
                proposeId_,
                proposeStatus_,
                isSupport_
            ),
            "signature error"
        );
        require(proposeStatus_ == 1, "proposal status error"); // 检查提案状态
        require(userHasVote[msg.sender][proposeId_] == false, "already vote"); // 检查是否投过票
        require(tox.balanceOf(msg.sender) >= voteFee, "insufficient balance"); // 检查余额

        tox.transferFrom(msg.sender, address(this), voteFee);
        userVoteFee[msg.sender][proposeId_] += voteFee;
        voteCount[proposeId_]++;
        if (isSupport_) {
            supportCount[proposeId_]++;
            userVoteDetail[msg.sender][proposeId_] = 1;
            proposeVoteList[proposeId_].push(
                VoteDetails(proposeId_, msg.sender, 1)
            );
        } else {
            unSupportCount[proposeId_]++;
            userVoteDetail[msg.sender][proposeId_] = 2;
            proposeVoteList[proposeId_].push(
                VoteDetails(proposeId_, msg.sender, 0)
            );
        }

        userHasVote[msg.sender][proposeId_] = true;
        proposeVoters[proposeId_].push(msg.sender);

        userVoteCount[msg.sender]++;
        userVoteProposes[msg.sender].push(proposeId_);

        emit VotePropose(proposeId_, voteFee, msg.sender, isSupport_);
    }

    // 管理员采纳/通过/过期提案
    // 0 未开始/未支付；1进行中，3 过期，2 被采纳
    // 投票被平台方采纳后发起人退还,投票用户tox不退还
    // 当提案持续30天依然未被平台方未采纳则自动失效，冻结押金不退还，投票的用户消耗的 Tox退还
    function setProposeStatusBatch(
        uint256[] calldata proposeId_,
        uint256[] calldata status_
    ) external onlyRole(SERVER_ROLE) {
        for (uint256 i; i < proposeId_.length; i++) {
            require(
                status_[i] == 2 || status_[i] == 3,
                "status must be 2 or 3"
            );
            if (status_[i] == 2) {
                // 检测采纳标准 总票数超过1万票，支持大于反对
                require(
                    //                    voteCount[proposeId_[i]] > 10000,  // 正式
                    voteCount[proposeId_[i]] > 3, // 测试
                    "the total number of votes must be greater than 10000"
                );
                require(
                    supportCount[proposeId_[i]] >=
                        unSupportCount[proposeId_[i]],
                    "the number of supporters must be greater than the number of opponents"
                );
                // 发起人的押金退换
                address creator = proposeCreator[proposeId_[i]];
                tox.transfer(creator, createFee);
                emit ReturnFee(proposeId_[i], creator, createFee, 1);
            }
            proposeStatus[proposeId_[i]] = status_[i];

            emit AdminSetProposeStatus(proposeId_[i], status_[i]);
        }
    }

    // 当提案持续30天依然未被平台方未采纳则自动失效，冻结押金不退还，投票的用户消耗的 Tox退还。用户自己领取
    function claimVoteFee(
        bytes calldata signature_,
        uint256 proposeId_,
        uint256 proposeStatus_
    ) external {
        require(
            checkClaimSignature(
                signature_,
                block.chainid,
                proposeId_,
                proposeStatus_
            ),
            "signature error"
        );
        require(proposeStatus_ == 3, "the status of propose is not overtime"); // 检查是否投过票
        require(
            userHasVote[msg.sender][proposeId_],
            "this user not vote for this propose"
        ); // 检查是否投过票
        require(
            !userHasReturnVoteFee[msg.sender][proposeId_],
            "you have already claim"
        ); // 检查是否领过
        tox.transfer(msg.sender, voteFee);
        userHasReturnVoteFee[msg.sender][proposeId_] = true;
        emit ReturnFee(proposeId_, msg.sender, voteFee, 0);
    }

    function claimCreateFee(
        bytes calldata signature_,
        uint256 proposeId_,
        uint256 proposeStatus_
    ) external {
        require(
            checkClaimSignature(
                signature_,
                block.chainid,
                proposeId_,
                proposeStatus_
            ),
            "signature error"
        );

        require(proposeStatus_ == 2, "the status of propose is not accepted"); // 检查投过状态

        require(
            proposeCreator[proposeId_] == msg.sender,
            "you are not the creator of this propose"
        ); // 检查是否创建过该提案

        require(
            !userHasReturnCreateFee[msg.sender][proposeId_],
            "you have already claim"
        ); // 检查是否领过

        tox.transfer(msg.sender, createFee);
        userHasReturnCreateFee[msg.sender][proposeId_] = true;
        emit ReturnFee(proposeId_, msg.sender, createFee, 1);
    }

    // 验证创建提案签名
    function checkCreateSignature(
        bytes memory signature_,
        uint256 chainId_,
        uint256 proposeId_,
        uint256 createTimestamp_
    ) private view returns (bool) {
        bytes32 messageHash = keccak256(
            abi.encodePacked(chainId_, proposeId_, createTimestamp_)
        );
        address recoverAddress = messageHash.toEthSignedMessageHash().recover(
            signature_
        );

        return recoverAddress == serverAddress;
    }

    // 验证领取签名
    function checkClaimSignature(
        bytes memory signature_,
        uint256 chainId_,
        uint256 proposeId_,
        uint256 proposeStatus_
    ) private view returns (bool) {
        bytes32 messageHash = keccak256(
            abi.encodePacked(chainId_, proposeId_, proposeStatus_)
        );
        address recoverAddress = messageHash.toEthSignedMessageHash().recover(
            signature_
        );

        return recoverAddress == serverAddress;
    }

    // 验证投票签名
    function checkVoteSignature(
        bytes memory signature_,
        uint256 chainId_,
        uint256 proposeId_,
        uint256 proposeStatus_,
        bool isSupport_
    ) private view returns (bool) {
        bytes32 messageHash = keccak256(
            abi.encodePacked(chainId_, proposeId_, proposeStatus_, isSupport_)
        );
        address recoverAddress = messageHash.toEthSignedMessageHash().recover(
            signature_
        );

        return recoverAddress == serverAddress;
    }

    // 批量查询提案总的投票数量
    function voteCountBatch(
        uint256[] calldata proposeIds_
    ) public view returns (uint256[] memory) {
        uint256[] memory res = new uint256[](proposeIds_.length);
        for (uint256 i; i < proposeIds_.length; i++) {
            res[i] = voteCount[proposeIds_[i]];
        }

        return res;
    }

    // 批量查询提案总的赞成票数量
    function supportCountBatch(
        uint256[] calldata proposeIds_
    ) public view returns (uint256[] memory) {
        uint256[] memory res = new uint256[](proposeIds_.length);
        for (uint256 i; i < proposeIds_.length; i++) {
            res[i] = supportCount[proposeIds_[i]];
        }

        return res;
    }

    // 批量查询提案总的反对票数量
    function unSupportCountBatch(
        uint256[] calldata proposeIds_
    ) public view returns (uint256[] memory) {
        uint256[] memory res = new uint256[](proposeIds_.length);
        for (uint256 i; i < proposeIds_.length; i++) {
            res[i] = unSupportCount[proposeIds_[i]];
        }

        return res;
    }

    // 批量查询用户对提案是否投过票
    // userHasVote; // 用户地址 => (提案ID=>是否投过票)
    function userHasVoteBatch(
        address user_,
        uint256[] calldata proposeIds_
    ) public view returns (bool[] memory) {
        bool[] memory res = new bool[](proposeIds_.length);
        for (uint256 i; i < proposeIds_.length; i++) {
            res[i] = userHasVote[user_][proposeIds_[i]];
        }

        return res;
    }

    // 批量查询用户投票信息
    function userVoteLst(
        address user_
    ) public view returns (VoteInfo[] memory) {
        uint256[] memory ids = userVoteProposes[user_];
        VoteInfo[] memory res = new VoteInfo[](ids.length);
        for (uint256 i = 0; i < ids.length; i++) {
            res[i] = VoteInfo(ids[i], userVoteDetail[user_][ids[i]]);
        }

        return res;
    }

    // 批量查询用户发起的提案信息
    function userProposeLst(
        address user_
    ) public view returns (ProposeInfo[] memory) {
        uint256[] memory ids = userProposeIds[user_];
        ProposeInfo[] memory res = new ProposeInfo[](ids.length);
        for (uint256 i = 0; i < ids.length; i++) {
            res[i] = ProposeInfo(ids[i], proposeCreator[ids[i]]);
        }

        return res;
    }

    // 批量查询提案的支付/创建时间
    function proposeCreatedAtLst(
        uint256[] calldata proposeIds_
    ) public view returns (uint256[] memory) {
        uint256[] memory res = new uint256[](proposeIds_.length);
        for (uint256 i = 0; i < proposeIds_.length; i++) {
            res[i] = proposeCreatedAt[proposeIds_[i]];
        }

        return res;
    }

    // 批量查询提案的状态
    function proposeStatusLst(
        uint256[] calldata proposeIds_
    ) public view returns (uint256[] memory) {
        uint256[] memory res = new uint256[](proposeIds_.length);
        for (uint256 i = 0; i < proposeIds_.length; i++) {
            res[i] = proposeStatus[proposeIds_[i]];
        }

        return res;
    }
}
