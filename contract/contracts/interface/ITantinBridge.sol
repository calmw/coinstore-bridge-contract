// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.10;

interface ITantinBridge {
    event DepositErc20Event(
        address indexed depositer,
        address indexed recipient,
        uint256 indexed amount,
        address tokenAddress,
        uint256 depositNonce
    );

    event DepositCoinEvent(
        address indexed depositer,
        address indexed recipient,
        uint256 indexed amount,
        uint256 depositNonce
    );

    event DepositErc721Event(
        address indexed depositer,
        address indexed recipient,
        uint256 indexed tokenId,
        address tokenAddress,
        uint256 depositNonce
    );

    event DepositErc1155Event(
        address indexed depositer,
        address indexed recipient,
        uint256 indexed amount,
        uint256 tokenId,
        address tokenAddress,
        uint256 depositNonce
    );
}
