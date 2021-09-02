import { Bid } from './models/bid';

export interface IBidState {
  loading: boolean;
  data: Array<Bid>;
  error: Error | null;
}
