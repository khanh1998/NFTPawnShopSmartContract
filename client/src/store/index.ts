import Vue from 'vue';
import Vuex from 'vuex';
import { PawnStateModule, PawnState } from './pawn';
import { State } from '@/store/state';

Vue.use(Vuex);

export default new Vuex.Store<State>({
  state: (): State => ({
    pawn: new PawnState(),
  }),
  mutations: {
  },
  actions: {
  },
  modules: {
    pawn: PawnStateModule,
  },
});
