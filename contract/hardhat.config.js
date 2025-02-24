require('@openzeppelin/hardhat-upgrades')
require("@nomiclabs/hardhat-waffle");
const PRIVATE_KEY = process.env.COINSTORE_BRIDGE

module.exports = {
    solidity: "0.8.22",
    settings: {
        optimizer: {
            enabled: true,
            runs: 2000
        }
    },
    networks: {
        match_test: {
            url: "http://52.195.158.175:8545",
            accounts: [`${PRIVATE_KEY}`],
            gasPrice: 10000000000
        }
    }
}
