// SPDX-License-Identifier: MIT
pragma solidity ^0.8.22;


contract ERC20TokenTest {

    uint8 public decimal;

    constructor() {}

    function decimals() public view returns (uint8) {
        return decimal;
    }

    function approve(address spender, uint256 value) external returns (bool){}

    function transfer(address to, uint256 value) external returns (bool){}

    function allowance(address owner, address spender) external view returns (uint256){}
}
