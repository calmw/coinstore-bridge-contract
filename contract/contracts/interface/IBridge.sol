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
        uint256 indexed depositNonce
    );
    event RelayerThresholdChanged(uint indexed newThreshold);
    event RelayerAdded(address indexed relayer);
    event RelayerRemoved(address indexed relayer);

    // 跨链币种信息
    struct TokenInfo {
        AssetsType assetsType; // 跨链币种
        address tokenAddress; // 币种地址。coin的话，值为0地址
        bool burnable; // true burn;false lock
        bool mintable; // true mint;false release
        bool blacklist; // 该币种是否在黑名单中/是否允许跨链。币种黑名单/禁止该币种跨链
        uint256 fee; // 跨链费用,当前设置的收手续费模式为固定数量的coin
    }

    struct DepositRecord {
        uint256 destinationChainId;
        address sender; // 某个业务合约的地址，可以有多个业务合约
        bytes32 resourceID;
        bytes data;
    }

    function deposit(
        uint256 destinationChainId,
        bytes32 resourceID,
        bytes calldata data
    ) external;

    function getChainId() external view returns (uint256);

    function getToeknInfoByResourceId(
        bytes32 resourceID
    ) external view returns (uint256, address, bool);
}
