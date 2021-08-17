const PawningShop = artifacts.require('PawningShop')

module.exports = (deployer) => {
    deployer.deployer(PawningShop)
}