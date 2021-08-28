import { Contract } from 'web3-eth-contract';
import Web3 from 'web3';

export function getContractInstance(contractJson: any, networkId: number, web3: Web3): Contract { // eslint-disable-line
  const deployedNetwork = contractJson.networks[networkId];
  const instance: Contract = new web3.eth.Contract(
    contractJson.abi,
    deployedNetwork && deployedNetwork.address,
  );
  return instance;
}
