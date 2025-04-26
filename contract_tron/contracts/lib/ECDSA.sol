// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.10;

library ECDSA {
    function toEthSignedMessageHash(
        bytes32 hash
    ) public pure returns (bytes32) {
        return
            keccak256(
                abi.encodePacked("\x19TRON Signed Message:\n32", hash)
            );
    }

    function recoverSigner(
        bytes32 _msgHash,
        bytes memory _signature
    ) public pure returns (address) {
        require(_signature.length == 65, "invalid signature length");
        bytes32 r;
        bytes32 s;
        uint8 v;
        assembly {
            r := mload(add(_signature, 0x20))
            s := mload(add(_signature, 0x40))
            v := byte(0, mload(add(_signature, 0x60)))
        }
        return ecrecover(_msgHash, v, r, s);
    }
}
