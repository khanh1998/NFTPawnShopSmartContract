import { User } from '@/store/models/user';

export interface IUserState {
  loading: boolean;
  data: User | null;
  error: Error | null;
}
