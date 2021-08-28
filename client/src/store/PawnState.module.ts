import {
  Module, VuexModule, Mutation, Action, getModule,
} from 'vuex-module-decorators';
import { Pawn } from '@/store/models/pawn';
import { IPawnState } from '@/store/IPawnState';
import { store } from '@/store/index';

@Module({
  namespaced: true,
  store,
  name: 'pawn',
  dynamic: true,
})
export class PawnState extends VuexModule implements IPawnState {
  loading = false;

  data: Array<Pawn> = [];

  error: Error | null = null;

  @Mutation
  FIND_ALL_PAWNS_BY_CREATOR_ADDRESS_REQUEST() {
    this.error = null;
    this.loading = true;
  }

  @Mutation
  FIND_ALL_PAWNS_BY_CREATOR_ADDRESS_SUCCESS(pawns: Array<Pawn>) {
    this.data = [...pawns];
    this.loading = false;
  }

  @Mutation
  FIND_ALL_PAWNS_BY_CREATOR_ADDRESS_FAIL(error: Error) {
    this.error = error;
    this.loading = false;
  }

  @Action
  async findAllByCreatorAddress(address: string) {
    try {
      this.context.commit('FIND_ALL_PAWNS_BY_CREATOR_ADDRESS_REQUEST');
      const pawns = await window.axios.get(`/users/${address}/pawns`);
      if (pawns.status === 200) {
        this.context.commit('FIND_ALL_PAWNS_BY_CREATOR_ADDRESS_SUCCESS', pawns.data);
      }
    } catch (error) {
      this.context.commit('FIND_ALL_PAWNS_BY_CREATOR_ADDRESS_FAIL', error);
    }
  }
}

export const pawn = getModule(PawnState);
