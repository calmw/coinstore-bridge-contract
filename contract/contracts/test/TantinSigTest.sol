// SPDX-License-Identifier: MIT
pragma solidity ^0.8.22;

import "@openzeppelin/contracts/utils/cryptography/ECDSA.sol";
import "@openzeppelin/contracts/utils/cryptography/MessageHashUtils.sol";

contract TantinBridgeSigTest {
    using ECDSA for bytes32;
    using MessageHashUtils for bytes32;

    // 验证deposit签名
    function testDepositSignature(
        bytes memory signature,
        address recipient
    ) public pure returns (address) {
        bytes32 messageHash = keccak256(abi.encode(recipient));
        address recoverAddress = messageHash.toEthSignedMessageHash().recover(
            signature
        );

        return recoverAddress;
    }
}
