// SPDX-License-Identifier: MIT
pragma solidity ^0.8.22;

contract TantinBridgeSigTest {
    function getChainId() public view returns (uint256) {
        return block.chainid;
    }
}
