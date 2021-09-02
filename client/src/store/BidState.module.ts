import {
  Module, VuexModule, Mutation, Action, getModule,
} from 'vuex-module-decorators';
import { Bid, ComputedBid } from './models/bid';
import { store } from '.';
import { IBidState } from './IBidState';
import { getRandomColor } from '@/utils/color';
import { convertSecondAndDurationToDateStr, convertSecondToDateStr } from '@/utils/time';

@Module({
  namespaced: true,
  name: 'bid',
  dynamic: true,
  store,
})
export class BidState extends VuexModule implements IBidState {
  loading = false;

  data: Array<Bid> = []

  error!: Error | null;

  get computedData(): Array<ComputedBid> {
    return this.data.map((bid) => ({
      ...bid,
      color: getRandomColor(),
      loan_start_time_str: convertSecondToDateStr(Number(bid.loan_start_time)),
      loan_end_time_str: convertSecondAndDurationToDateStr(
        Number(bid.loan_start_time),
        Number(bid.loan_duration),
      ),
    }));
  }

  @Mutation
  FIND_BID_REQUEST() {
    this.error = null;
    this.loading = true;
  }

  @Mutation
  FIND_BID_SUCCESS(data: Array<Bid>) {
    this.data = data;
    this.loading = false;
  }

  @Mutation
  FIND_BID_FAIL(error: Error) {
    this.error = error;
    this.loading = false;
  }

  @Action
  async findAllBy(query: string) {
    try {
      this.context.commit('FIND_BID_REQUEST');
      const res = await window.axios.get(`/bids?${query}`);
      this.context.commit('FIND_BID_SUCCESS', res.data);
    } catch (error) {
      this.context.commit('FIND_BID_FAIL', error);
    }
  }
}

export const bid = getModule(BidState);
