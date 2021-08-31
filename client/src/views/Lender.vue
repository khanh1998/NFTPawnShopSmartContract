<template>
  <v-container>
    <v-row v-if="!loading">
      <v-col>
        <bid-list
          :bids="bid.computedData"
          :pawning-shop-contract="pawningShopContract"
          :accounts="accounts"
          viewer="lender"
        />
      </v-col>
      <v-col>
        <pawn-list
          :pawns="pawn.computedData"
          :pawning-shop-contract="pawningShopContract"
          :accounts="accounts"
          viewer="lender"
        />
      </v-col>
    </v-row>
  </v-container>
</template>
<script lang="ts">
import { Vue, Component } from 'vue-property-decorator';
import { Contract } from 'web3-eth-contract';
import { pawn, PawnState } from '@/store/PawnState.module';
import { bid, BidState } from '@/store/BidState.module';
import PawnList from '@/components/PawnList.vue';
import BidList from '@/components/BidList.vue';
import { getContractInstance } from '@/utils/contract';
import PawningShop from '@/contracts/PawningShop.json';

@Component({
  name: 'Lender',
  components: { PawnList, BidList },
  data: () => ({
    pawn,
    bid,
  }),
})
export default class extends Vue {
  pawn!: PawnState;

  bid!: BidState;

  localLoading = false;

  pawningShopContract!: Contract;

  accounts: string[] = []

  networkId = -1

  get loading(): boolean {
    return this.pawn.loading || this.localLoading || this.bid.loading;
  }

  async getNetworkId(): Promise<number> {
    return this.$web3.eth.net.getId();
  }

  async getAccounts(): Promise<string[]> {
    return this.$web3.eth.getAccounts();
  }

  async created() {
    this.localLoading = true;
    this.pawn.findAllBy('status=0');
    this.accounts = await this.getAccounts();
    this.bid.findAllBy(`creator=${this.accounts[0]}`);
    this.networkId = await this.getNetworkId();
    if (this.networkId !== 5777) {
      console.log('you are in wrong network babe :D');
    }
    this.pawningShopContract = getContractInstance(PawningShop, this.networkId, this.$web3);
    this.localLoading = false;
  }
}
</script>
