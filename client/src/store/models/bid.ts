export class Bid {
  creator!: string;
}

export class BidCreate {
  loanAmount!: number
	interest!: number
	loanStartTime!: number
	loanDuration!: number
	isInterestProRated!: boolean
}
