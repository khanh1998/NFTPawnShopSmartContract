import { Bid } from './bid';
import { User } from './user';

export class Pawn {// eslint-disable-line
  UUID!: string;

  id!: string;

  creator!: User;

  token_address!: string;// eslint-disable-line

  token_id!: string;// eslint-disable-line

  status!: number;

  acceptedBid!: Bid | undefined;

  bids!: Bid[];
}

export interface ComputedPawn extends Pawn {
  color : string;
  statusName: string;
}
