// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.10;

import "./interface/IBridge.sol";

contract A {
    IBridge public Bridge; // bridge 合约

    constructor(){}

    function adminSetEnv(
        address bridgeAddress_
    ) public {
        Bridge = IBridge(bridgeAddress_);
    }

    function executeProposal(
        uint256 originChainId,
        uint256 originDepositNonce,
        bytes calldata data
    ) public view returns (bytes32,uint72){
        uint72 nonceAndID = (uint72(originDepositNonce) << 8) |
                            uint72(originChainId);
        bytes32 dataHash = keccak256(abi.encodePacked(Bridge, data));

        return (dataHash,nonceAndID);
    }
}
