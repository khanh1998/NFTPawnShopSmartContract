const PawningShop = artifacts.require('PawningShop')
const TestToken = artifacts.require('TestToken')

module.exports = async (deployer) => {
    const walletAddress = await web3.eth.getAccounts()
    const [deployerAdd, beneficiaryAdd] = walletAddress
    await deployer.deploy(PawningShop)
    await deployer.deploy(TestToken)
    const TestTokenIns = await TestToken.deployed()
    await TestTokenIns.mint(beneficiaryAdd, { from: deployerAdd })
    await TestTokenIns.setApprovalForAll(PawningShop.address, true, { from: beneficiaryAdd})
}