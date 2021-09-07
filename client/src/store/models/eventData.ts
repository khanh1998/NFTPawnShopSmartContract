export interface EventData {
  bidid: string
  borrower: string
  code: string
  pawnid: string
  lender: string
  payload: string
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
  PawnCreatedName      : "PawnCreated",
	PawnCancelledName    : "PawnCancelled",
	WhiteListAddedName   : "WhiteListAdded",
	WhiteListRemovedName : "WhiteListRemoved",
	BidCreatedName       : "BidCreated",
	BidCancelledName     : "BidCancelled",
	BidAcceptedName      : "BidAccepted",
	PawnRepaidName       : "PawnRepaid",
	PawnLiquidatedName   : "PawnLiquidated",
}