import { Bid } from './bid';
import { User } from './user';

export class Pawn {
  UUID!: string;

  id!: string;

  creator!: User;

  token_address!: string;

  token_id!: string;

  status!: number;

  bids!: Bid[];
}
