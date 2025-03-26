// SPDX-License-Identifier: MIT
pragma solidity ^0.8.22;

import "./interface/IERC20MintAble.sol";
import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";
import "@openzeppelin/contracts/utils/Address.sol";
import "@openzeppelin/contracts/utils/cryptography/ECDSA.sol";
import "@openzeppelin/contracts/utils/cryptography/MessageHashUtils.sol";
import {AccessControl} from "@openzeppelin/contracts/access/AccessControl.sol";
import {IBridge} from "./interface/IBridge.sol";
import {ITantinBridge} from "./interface/ITantinBridge.sol";

contract TestA {
    using SafeERC20 for IERC20;
    enum AssetsType {
        None,
        Coin,
        Erc20,
        Erc721,
        Erc1155
    }
    struct TokenInfo {
        AssetsType assetsType; // 跨链币种
        address tokenAddress; // 币种地址。coin的话，值为0地址
        bool burnable; // true burn;false lock
        bool mintable; // true mint;false release
        bool pause; // 该token是否暂停跨链
    }

    error ErrAssetsType(AssetsType assetsType);

    IBridge public Bridge; // bridge 合约
    mapping(bytes32 => TokenInfo) public resourceIdToTokenInfo; //  resourceID => 设置的Token信息

    constructor() {}

    /**
        @notice 设置
        @param bridgeAddress_ bridge合约地址
     */
    function adminSetEnv(address bridgeAddress_) external {
        Bridge = IBridge(bridgeAddress_);
    }

    function adminSetToken(
        bytes32 resourceID,
        AssetsType assetsType,
        address tokenAddress,
        bool burnable,
        bool mintable,
        bool pause
    ) external {
        resourceIdToTokenInfo[resourceID] = TokenInfo(
            assetsType,
            tokenAddress,
            burnable,
            mintable,
            pause
        );
    }

    function executeA(bytes32 resourceId, address recipient) public {
        TokenInfo memory tokenInfo = resourceIdToTokenInfo[resourceId];
        address tokenAddress = tokenInfo.tokenAddress;
        if (tokenInfo.assetsType == AssetsType.Coin) {
            Address.sendValue(payable(recipient), 1);
        } else if (tokenInfo.assetsType == AssetsType.Erc20) {
            if (tokenInfo.mintable) {
                IERC20MintAble erc20 = IERC20MintAble(tokenAddress);
                erc20.mint(recipient, 1);
            } else {
                IERC20 erc20 = IERC20(tokenAddress);
                //                erc20.safeTransfer(recipient, 1); // TODO 上线取消注释
                erc20.transfer(recipient, 1); // TODO 上线取消注释
            }
        } else {
            revert ErrAssetsType(tokenInfo.assetsType);
        }
    }
}
