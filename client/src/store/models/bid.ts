export interface Bid {
creator: string;
id: string;
interest: string;
loan_amount: string; // eslint-disable-line
loan_duration: string; // eslint-disable-line
loan_start_time: string; // eslint-disable-line
pawn: string;
pro_rated: boolean; // eslint-disable-line
status: number;
}

export interface BidCreate {
loanAmount: number
interest: number
loanStartTime: number
loanDuration: number
isInterestProRated: boolean
}

export interface ComputedBid extends Bid {
color: string;
loan_start_time_str: string; // eslint-disable-line
loan_end_time_str: string; // eslint-disable-line
}

export const BidStatus = {
  CREATED: 0,
  CANCELLED: 1,
  ACCEPTED: 2,
};
