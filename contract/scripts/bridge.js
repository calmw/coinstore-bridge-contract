const {ethers, upgrades} = require("hardhat")
const {BigNumber} = require("ethers");
const {write_contract_address} = require("./fs");


async function main() {
    const contract_name = "Vote"

    const contract = await ethers.getContractFactory(contract_name)
    console.log("Deploying .........")

    const contractObj = await upgrades.deployProxy(contract, [
        '0x2ddae7885B6A069A4c2b46F429b467f638D6bD6d',
        '0x1c1E185fF58e34126613b67bd277C0180295bD27',
        BigNumber.from("50000000000000000000000"),
        BigNumber.from("100000000000000000000"),
    ], {initializer: "initialize"});

    console.log("Proxy address is ", contractObj.address)
    write_contract_address(contract_name, contractObj.address)
}

main()
    .then(() => process.exit(0))
    .catch(error => {
        console.error(error);
        process.exit(1);
    });