// SPDX-License-Identifier: MIT
pragma solidity ^0.8.22;

interface IBridge {
    enum AssetsType {
        None,
        Coin,
        Erc20,
        Erc721,
        Erc1155
    }

    event Deposit(
        uint256 indexed destinationChainId,
        bytes32 indexed resourceID,
        uint256 indexed depositNonce,
        bytes data
    );

    event SetResource(
        bytes32 indexed resourceID,
        address tokenAddress,
        uint256 fee,
        bool pause, // 该resourceID是否被暂停交易
        address tantinAddress
    );

    // 跨链币种信息
    struct TokenInfo {
        AssetsType assetsType; // 跨链币种
        address tokenAddress; // 币种地址。coin的话，值为0地址
        bool pause; // 该token是否暂停跨链
        uint256 fee; // 跨链费用,对跨链币种按比例收取，此处为万分比
    }

    struct DepositRecord {
        uint256 destinationChainId;
        address sender; // 某个业务合约的地址，可以有多个业务合约
        bytes32 resourceID;
        uint256 ctime;
        bytes data;
    }

    function deposit(
        uint256 destinationChainId,
        bytes32 resourceID,
        bytes calldata data
    ) external payable;

    function chainId() external view returns (uint256);

    function getTokenInfoByResourceId(
        bytes32 resourceID
    ) external view returns (uint256, address, bool);

    function getFeeByResourceId(
        bytes32 resourceId
    ) external view returns (uint256);

    function getContractAddressByResourceId(
        bytes32 resourceId
    ) external view returns (address);

    function execute(bytes32 resourceId, bytes calldata data) external;
}
