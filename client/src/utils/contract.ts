import { Contract } from 'web3-eth-contract';
import Web3 from 'web3';
import { ComputedPawn } from '@/store/models/pawn';

export function getContractInstance(contractJson: any, networkId: number, web3: Web3): Contract { // eslint-disable-line
  const deployedNetwork = contractJson.networks[networkId];
  const instance: Contract = new web3.eth.Contract(
    contractJson.abi,
    deployedNetwork && deployedNetwork.address,
  );
  return instance;
}

export function getStatusName(code: number) {
  switch (code) {
    case 0:
      return 'Created';
    case 1:
      return 'Cancelled';
    case 2:
      return 'Deal';
    case 3:
      return 'Liquidated';
    case 4:
      return 'Repaid';
    default:
      return 'Unknown';
  }
}

export function extractErrorObjFromMessage(message: string): any {
  const singleQuoteIdx = message.indexOf("'");
  const jsonStr = message.substring(singleQuoteIdx + 1, message.length - 1);
  const error = JSON.parse(jsonStr);
  const realErr = error.value.data.data;
  return realErr[Object.keys(realErr)[0]];
}

export function calculateRepaidAmount(pawn: ComputedPawn): number {
  if (pawn.acceptedBid) {
    const {
      loan_amount: loanAmount,
      interest,
      loan_duration: duration,
      pro_rated: isInterestProRated,
      loan_start_time: startTime,
    } = pawn.acceptedBid;
    let interestDue = Number(interest);
    if (isInterestProRated) {
      const interestPerDate = Math.floor(Number(interest) / Number(duration));
      const nowInSeconds = Date.now() / 1000;
      const aDayInSeconds = 24 * 60 * 60;
      const dayPassed = Math.ceil((nowInSeconds - Number(startTime)) / aDayInSeconds);
      interestDue = dayPassed * interestPerDate;
    }
    return Number(loanAmount) + Number(interestDue);
  }
  return -1;
}
