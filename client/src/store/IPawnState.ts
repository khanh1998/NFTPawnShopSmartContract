import { Pawn } from '@/store/models/pawn';

export interface IPawnState {
  loading: boolean;
  data: Array<Pawn>;
  error: Error;
}