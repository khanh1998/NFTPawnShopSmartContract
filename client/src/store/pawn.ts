// https://www.codeproject.com/Tips/5295301/Correctly-Typing-Vuex-with-TypeScript
// https://chiragrupani.medium.com/vuex-with-typescript-b83a62aa48a8
// https://stackoverflow.com/questions/62053067/how-to-get-intellisense-of-mapgetters-mapactions-vuex-and-typescript-without-cl
import { ActionContext } from 'vuex';
import { Pawn } from '@/store/models/pawn';
import { State as RootState } from '@/store/state';
import { IPawnState } from '@/store/IPawnState';

export const MutationNames = {
  FIND_ALL_PAWNS_BY_CREATOR_ADDRESS_REQUEST: 'FIND_ALL_PAWNS_BY_CREATOR_ADDRESS_REQUEST',
  FIND_ALL_PAWNS_BY_CREATOR_ADDRESS_SUCCESS: 'FIND_ALL_PAWNS_BY_CREATOR_ADDRESS_SUCCESS',
  FIND_ALL_PAWNS_BY_CREATOR_ADDRESS_FAIL: 'FIND_ALL_PAWNS_BY_CREATOR_ADDRESS_FAIL',
};

export class PawnState implements PawnState {
  loading: boolean;

  data: Array<Pawn>;

  error!: Error;

  constructor() {
    this.loading = false;
    this.data = Array<Pawn>();
  }
}

export const PawnStateModule = {
  state: () => new PawnState(),
  mutations: {
    FIND_ALL_PAWNS_BY_CREATOR_ADDRESS_REQUEST(state: IPawnState) {
      state.loading = true;
    },
    FIND_ALL_PAWNS_BY_CREATOR_ADDRESS_SUCCESS(state: IPawnState, pawns: Array<Pawn>) {
      state.data = [...pawns];
      state.loading = false;
    },
    FIND_ALL_PAWNS_BY_CREATOR_ADDRESS_FAIL(state: IPawnState, error: Error) {
      state.error = error;
      state.loading = false;
    },
  },
  actions: {
    async findAllByCreatorAddress(context: ActionContext<IPawnState, RootState>, address: string) {
      try {
        context.commit(MutationNames.FIND_ALL_PAWNS_BY_CREATOR_ADDRESS_REQUEST);
        const pawns = await window.axios.get(`/users/${address}/pawns`);
        if (pawns.status === 200) {
          context.commit(MutationNames.FIND_ALL_PAWNS_BY_CREATOR_ADDRESS_SUCCESS, pawns.data);
        }
      } catch (error) {
        context.commit(MutationNames.FIND_ALL_PAWNS_BY_CREATOR_ADDRESS_FAIL, error);
      }
    },
  },
  modules: {
  },
};
