require('@openzeppelin/hardhat-upgrades')
require("@nomiclabs/hardhat-waffle");
// const PRIVATE_KEY = process.env.TT_BRIDGE_SIGN
const PRIVATE_KEY = process.env.TB_PROD
module.exports = {
    solidity: "0.8.22",
    settings: {
        optimizer: {
            enabled: true,
            runs: 2000
        }
    },
    networks: {
        bsc: {
            // url: "https://bsc-mainnet.infura.io/v3/59ec080dc74d4af893ea04bfe2b168b5",
            url: "https://data-seed-prebsc-2-s3.bnbchain.org:8545",
            accounts: [`${PRIVATE_KEY}`]
            // gasPrice: 10000000000
        },
        bsc_mainnet: {
            url: "https://bsc-mainnet.infura.io/v3/59ec080dc74d4af893ea04bfe2b168b5",
            accounts: [`${PRIVATE_KEY}`]
            // gasPrice: 10000000000
        },
        ethereum: {
            url: "https://mainnet.infura.io/v3/59ec080dc74d4af893ea04bfe2b168b5",
            accounts: [`${PRIVATE_KEY}`]
            // gasPrice: 10000000000
        },
        tantin_test: {
            url: "https://testrpc.tantin.com",
            // url: "https://testrpcdex.tantin.com",
            accounts: [`${PRIVATE_KEY}`],
            gasPrice: 142861904761900
        },
        tantin: {
            url: "https://rpc.tantin.com",
            accounts: [`${PRIVATE_KEY}`],
            gasPrice: 100000000000000
        },
        open_bnb: {
            url: "https://opbnb-testnet-rpc.bnbchain.org",
            accounts: [`${PRIVATE_KEY}`]
            // gasPrice: 10000000000
        },
        sepolia: {
            // url: "https://sepolia.infura.io/v3/732f6502b35c486fb07e333b32e89c04",
            // url: "https://endpoints.omniatech.io/v1/eth/sepolia/public",
            url: "https://sepolia.drpc.org",
            // url: "https://1rpc.io/sepolia",
            // url: "https://linea-sepolia.infura.io",
            accounts: [`${PRIVATE_KEY}`],
            gasPrice: 1000000000000
        }
    },
    etherscan: {
        apiKey: {
            open_bnb: "95d7c8f518a549b1a5a844c552f3725e"
        },
        customChains: [
            {
                network: "open_bnb",
                chainId: 5611,
                urls: {
                    apiURL:
                        "https://open-platform.nodereal.io/95d7c8f518a549b1a5a844c552f3725e/op-bnb-testnet/contract/",
                    browserURL: "https://testnet.opbnbscan.com/"
                }
            }
        ]
    }
}
