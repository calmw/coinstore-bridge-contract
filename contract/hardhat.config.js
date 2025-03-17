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
        open_bnb: {
            url: "https://opbnb-testnet-rpc.bnbchain.org",
            accounts: [`${PRIVATE_KEY}`]
            // gasPrice: 10000000000
        },
        tantin_testnet: {
            url: "https://rpc.tantin.com",
            accounts: [`${PRIVATE_KEY}`]
            // gasPrice: 10000000000
        }
    }
}
