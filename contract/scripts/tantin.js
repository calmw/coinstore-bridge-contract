const {ethers, upgrades} = require("hardhat")
const {BigNumber} = require("ethers");
const {write_contract_address} = require("./fs");


async function main() {
    const contract_name = "TantinBridge"

    const contract = await ethers.getContractFactory(contract_name)
    console.log("Deploying .........")

    const contractObj = await upgrades.deployProxy(contract, [], {initializer: "initialize"});

    console.log("Proxy address is ", contractObj.address)
    write_contract_address(contract_name, contractObj.address)
}

main()
    .then(() => process.exit(0))
    .catch(error => {
        console.error(error);
        process.exit(1);
    });