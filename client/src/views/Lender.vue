<template>
  <v-container>
    <v-row v-if="!loading">
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
import PawnList from '@/components/PawnList.vue';
import { getContractInstance } from '@/utils/contract';
import PawningShop from '@/contracts/PawningShop.json';

@Component({
  name: 'Lender',
  components: { PawnList },
  data: () => ({
    pawn,
  }),
})
export default class extends Vue {
  pawn!: PawnState;

  localLoading = false;

  pawningShopContract!: Contract;

  accounts: string[] = []

  networkId = -1

  get loading(): boolean {
    return this.pawn.loading || this.localLoading;
  }

  async getNetworkId(): Promise<number> {
    return this.$web3.eth.net.getId();
  }

  async getAccounts(): Promise<string[]> {
    return this.$web3.eth.getAccounts();
  }

  async created() {
    this.pawn.findAllBy('status=0');
    this.localLoading = true;
    this.accounts = await this.getAccounts();
    this.networkId = await this.getNetworkId();
    if (this.networkId !== 5777) {
      console.log('you are in wrong network babe :D');
    }
    this.pawningShopContract = getContractInstance(PawningShop, this.networkId, this.$web3);
    this.localLoading = false;
  }
}
</script>
