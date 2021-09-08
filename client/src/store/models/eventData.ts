export interface EventData {
  bidid: string
  borrower: string
  code: string
  pawnid: string
  lender: string
  payload: string
  message: string
}

export interface PawnPayload {
  Creator: string
  ContractAddress: string
  TokenId: string
  Status: number
}

export interface BidPayload {
  Creator: string
  LoanAmount: string
  Interest: string
  LoanStartTime: string
  LoanDuration: string
  isInterestProRated: boolean,
}

export const EventName = {
  PawnCreated: 'PawnCreated',
  PawnCancelled: 'PawnCancelled',
  WhiteListAdded: 'WhiteListAdded',
  WhiteListRemoved: 'WhiteListRemoved',
  BidCreated: 'BidCreated',
  BidCancelled: 'BidCancelled',
  BidAccepted: 'BidAccepted',
  PawnRepaid: 'PawnRepaid',
  PawnLiquidated: 'PawnLiquidated',
};
