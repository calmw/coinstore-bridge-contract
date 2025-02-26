// SPDX-License-Identifier: MIT
pragma solidity ^0.8.22;

interface IERC20 {
    function balanceOf(address account) external view returns (uint256);

    function transfer(address to, uint256 amount) external returns (bool);

    function transferFrom(
        address from,
        address to,
        uint256 amount
    ) external returns (bool);

    function mint(address account, uint256 amount) external;

    function burnFrom(address account, uint256 amount) external;
}
