// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.22;

interface ITantinBridge {
    enum AssetsType {
        None,
        Coin,
        Erc20,
        Erc721,
        Erc1155
    }

    event AddBlacklist(address indexed user);

    event RemoveBlacklist(address indexed user);

    event DepositEvent(
        address indexed depositer,
        address indexed recipient,
        uint256 indexed amount,
        address tokenAddress,
        uint256 depositNonce,
        uint256 destinationChainId
    );

    event DepositNftEvent(
        address indexed depositer,
        address indexed recipient,
        uint256 indexed amount,
        uint256 tokenId,
        address tokenAddress,
        uint256 depositNonce,
        uint256 destinationChainId
    );

    event ExecuteEvent(
        address indexed depositer,
        address indexed recipient,
        uint256 indexed amount,
        address tokenAddress,
        uint256 originDepositNonce,
        uint256 originChainId
    );

    event ExecuteNftEvent(
        address indexed depositer,
        address indexed recipient,
        uint256 indexed amount,
        uint256 tokenId,
        address tokenAddress,
        uint256 originDepositNonce,
        uint256 originChainId
    );

    event SetTokenEvent(
        bytes32 indexed resourceID,
        AssetsType assetsType,
        address tokenAddress,
        bool burnable,
        bool mintable,
        bool pause // 该token是否暂停跨链
    );

    struct TokenInfo {
        AssetsType assetsType; // 跨链币种
        address tokenAddress; // 币种地址。coin的话，值为0地址
        bool burnable; // true burn;false lock
        bool mintable; // true mint;false release
        bool pause; // 该token是否暂停跨链
    }

    struct DepositRecord {
        address tokenAddress;
        address sender;
        address recipient;
        uint256 amount;
        uint256 fee;
        uint256 destinationChainId;
    }

    function execute(bytes calldata data) external;
}
