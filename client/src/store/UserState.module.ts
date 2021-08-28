import {
  Module, VuexModule, Mutation, Action, getModule,
} from 'vuex-module-decorators';
import { IUserState } from './IUserState';
import { User } from './models/user';
import { store } from '.';

@Module({
  namespaced: true,
  name: 'user',
  dynamic: true,
  store,
})
export class UserState extends VuexModule implements IUserState {
  loading = false;

  data: User | null = null;

  error!: Error | null;

  @Mutation
  FIND_USER_BY_ADDRESS_REQUEST() {
    this.error = null;
    this.loading = true;
  }

  @Mutation
  FIND_USER_BY_ADDRESS_SUCCESS(data: User) {
    this.data = data;
    this.loading = false;
  }

  @Mutation
  FIND_USER_BY_ADDRESS_FAIL(error: Error) {
    this.error = error;
    this.loading = false;
  }

  @Action
  async findUserByAddress(address: string) {
    try {
      this.context.commit('FIND_USER_BY_ADDRESS_REQUEST');
      const res = await window.axios.get(`/users/${address}`);
      this.context.commit('FIND_USER_BY_ADDRESS_SUCCESS', res.data);
    } catch (error) {
      this.context.commit('FIND_USER_BY_ADDRESS_FAIL', error);
    }
  }
}

export const user = getModule(UserState);
