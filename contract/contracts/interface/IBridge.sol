// SPDX-License-Identifier: MIT
pragma solidity ^0.8.22;

interface IBridge {
    event Deposit(
        uint256 indexed destinationChainId,
        bytes32 indexed resourceID,
        uint256 indexed depositNonce
    );
    event RelayerThresholdChanged(uint indexed newThreshold);
    event RelayerAdded(address indexed relayer);
    event RelayerRemoved(address indexed relayer);

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

    function getToeknAddress(
        bytes32 resourceID
    ) external view returns (address);
}
