import {
  Module, VuexModule, Mutation, Action, getModule,
} from 'vuex-module-decorators';
import { ComputedPawn, Pawn } from '@/store/models/pawn';
import { getRandomColor } from '@/utils/color';
import { getStatusName } from '@/utils/contract';
import { IPawnState } from '@/store/IPawnState';
import { store } from '@/store/index';
import { Bid, BidStatus } from './models/bid';

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

  get computedData(): Array<ComputedPawn> {
    return this.data.map((pawn) => ({
      ...pawn,
      color: getRandomColor(),
      statusName: getStatusName(pawn.status),
      acceptedBid: pawn.bids.find((bid: Bid) => bid.status === BidStatus.ACCEPTED),
    }));
  }

  @Mutation
  FIND_PAWNS_REQUEST() {
    this.error = null;
    this.loading = true;
  }

  @Mutation
  FIND_PAWNS_SUCCESS(pawns: Array<Pawn>) {
    this.data = [...pawns];
    this.loading = false;
  }

  @Mutation
  FIND_PAWNS_FAIL(error: Error) {
    this.error = error;
    this.loading = false;
  }

  @Action
  async findAllByCreatorAddress(address: string) {
    try {
      this.context.commit('FIND_PAWNS_REQUEST');
      const pawns = await window.axios.get(`/users/${address}/pawns`);
      if (pawns.status === 200) {
        this.context.commit('FIND_PAWNS_SUCCESS', pawns.data);
      }
    } catch (error) {
      this.context.commit('FIND_PAWNS_FAIL', error);
    }
  }

  @Action
  async findAllBy(query: string) {
    try {
      this.context.commit('FIND_PAWNS_REQUEST');
      const pawns = await window.axios.get(`/pawns?${query}`);
      if (pawns.status === 200) {
        this.context.commit('FIND_PAWNS_SUCCESS', pawns.data);
      }
    } catch (error) {
      this.context.commit('FIND_PAWNS_FAIL', error);
    }
  }
}

export const pawn = getModule(PawnState);
