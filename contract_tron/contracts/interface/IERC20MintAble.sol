// SPDX-License-Identifier: MIT
pragma solidity ^0.8.22;

interface IERC20MintAble {
    function mint(address account, uint256 amount) external;

    function burnFrom(address account, uint256 amount) external;
}
