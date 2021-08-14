require('dotenv').config({ path: '../.env'})
const ERC721 = artifacts.require('ERC721')

const chai = require('./setupChai.js')
const BN = web3.utils.BN;
expect = chai.expect

contract('erc721', (accounts) => {
    const [deployer, recipient, other] = accounts
    
})