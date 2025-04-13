// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.10;

import "@openzeppelin/contracts/utils/cryptography/ECDSA.sol";
contract test_sig {
    using ECDSA for bytes32;

    uint256 public sigNonce; // 签名nonce, parameter➕nonce➕chainID
    address public superAdminAddress;
    uint256 public chainId; // 自定义链ID
    function adminSetEnv(address addr) external {
        superAdminAddress = addr;
    }

    function checkAdminSetEnvSignatureTest(
        bytes memory signature_,
        address voteAddress_,
        uint256 chainId_,
        uint256 chainType_
    ) public view returns (address, address, address, address, address) {
        bytes32 messageHash = keccak256(
            abi.encode(sigNonce, chainId_, voteAddress_, chainId_, chainType_)
        );
        address recoverAddress = recoverSigner(
            toEthSignedMessageHash(messageHash),
            signature_
        );
        address recoverAddress2 = recoverSigner(
            toEthSignedMessageHash2(messageHash),
            signature_
        );
        address recoverAddress3 = messageHash.recover(signature_);
        bytes memory signature2_ = signature_;
        address recoverAddress4 = recoverSigner(
            toEthSignedMessageHash3(messageHash),
            signature2_
        );
        address recoverAddress5 = recoverSigner(
            toEthSignedMessageHash4(messageHash),
            signature2_
        );

        return (
            recoverAddress,
            recoverAddress2,
            recoverAddress3,
            recoverAddress4,
            recoverAddress5
        );
    }

    function toEthSignedMessageHash(
        bytes32 hash
    ) public pure returns (bytes32) {
        return
            keccak256(abi.encodePacked("\x19TRON Signed Message:\n32", hash));
    }

    function toEthSignedMessageHash2(
        bytes32 hash
    ) public pure returns (bytes32) {
        return
            keccak256(
                abi.encodePacked("\x19Ethereum Signed Message:\n32", hash)
            );
    }

    function toEthSignedMessageHash3(
        bytes32 hash
    ) public pure returns (bytes32) {
        return keccak256(abi.encode("\x19TRON Signed Message:\n32", hash));
    }

    function toEthSignedMessageHash4(
        bytes32 hash
    ) public pure returns (bytes32) {
        return keccak256(abi.encode("\x19Ethereum Signed Message:\n32", hash));
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
