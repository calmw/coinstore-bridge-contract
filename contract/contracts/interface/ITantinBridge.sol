// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.22;

interface ITantinBridge {
    event DepositCoinEvent(
        address indexed depositer,
        address indexed recipient,
        uint256 indexed amount,
        uint256 depositNonce,
        uint256 destinationChainId
    );

    event DepositErc20Event(
        address indexed depositer,
        address indexed recipient,
        uint256 indexed amount,
        address tokenAddress,
        uint256 depositNonce,
        uint256 destinationChainId
    );

    event DepositErc721Event(
        address indexed depositer,
        address indexed recipient,
        uint256 indexed tokenId,
        address tokenAddress,
        uint256 depositNonce,
        uint256 destinationChainId
    );

    event DepositErc1155Event(
        address indexed depositer,
        address indexed recipient,
        uint256 indexed amount,
        uint256 tokenId,
        address tokenAddress,
        uint256 depositNonce,
        uint256 destinationChainId
    );

    event ExecuteCoinEvent(
        address indexed depositer,
        address indexed recipient,
        uint256 indexed amount,
        uint256 originDepositNonce,
        uint256 originChainId
    );

    event ExecuteErc20Event(
        address indexed depositer,
        address indexed recipient,
        uint256 indexed amount,
        address tokenAddress,
        uint256 originDepositNonce,
        uint256 originChainId
    );

    event ExecuteErc721Event(
        address indexed depositer,
        address indexed recipient,
        uint256 indexed tokenId,
        address tokenAddress,
        uint256 originDepositNonce,
        uint256 originChainId
    );

    event ExecuteErc1155Event(
        address indexed depositer,
        address indexed recipient,
        uint256 indexed amount,
        uint256 tokenId,
        address tokenAddress,
        uint256 originDepositNonce,
        uint256 originChainId
    );

    struct DepositErc20Record {
        address tokenAddress;
        address sender;
        address recipient;
        uint256 amount;
        uint256 depositNonce;
        uint256 destinationChainId;
    }
}
