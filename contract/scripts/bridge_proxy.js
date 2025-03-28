const {ethers, upgrades} = require("hardhat")
require('@openzeppelin/hardhat-upgrades')
const {read_contract_address} = require("./fs");

async function main() {

    const contract_name = "Vote"
    let proxy_address = read_contract_address(contract_name)

    console.log("ImplementationAddress is", await upgrades.erc1967.getImplementationAddress(proxy_address));
    console.log("ProxyAdmin is", await upgrades.erc1967.getAdminAddress(proxy_address));

    const factory = await ethers.getContractFactory(contract_name);
    const contract = await upgrades.upgradeProxy(proxy_address, factory);

    await contract.deployed();
    console.log("proxy address is ", contract.address)
    console.log("ImplementationAddress is", await upgrades.erc1967.getImplementationAddress(contract.address));
    console.log("ProxyAdmin is", await upgrades.erc1967.getAdminAddress(contract.address));
}

main()
    .then(() => process.exit(0))
    .catch(error => {
        console.error(error);
        process.exit(1);
    });