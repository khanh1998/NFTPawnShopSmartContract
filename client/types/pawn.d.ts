import { User } from './user'
import { Bid } from './bid'

export interface Pawn {
  UUID: string,
  id: string,
  creator: User,
  token_address: string,
  token_id: string,
  status: Number,
  bids: Bid[],
}