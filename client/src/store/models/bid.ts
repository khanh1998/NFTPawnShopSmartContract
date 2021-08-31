export class Bid {
  creator!: string;
	id!: string;
	interest!: string;
	loan_amount!: string;
	loan_duration!: string;
	loan_start_time!: string;
	pawn!: string;
	pro_rated!: boolean;
}

export class BidCreate {
  loanAmount!: number
	interest!: number
	loanStartTime!: number
	loanDuration!: number
	isInterestProRated!: boolean
}
