<template>
  <v-container>
    <v-row>
      <v-col>
        <p>Hello borrower!!! {{ accounts[0] }}</p>
      </v-col>
    </v-row>
    <v-row v-if="!loading">
      <v-col>
        <pawn-creator @create-pawn="createPawn" :white-list="whiteList"/>
      </v-col>
      <v-col>
        <pawn-list
          :pawns="pawn.data"
          :accounts="accounts"
          :pawning-shop-contract="pawningShopContract"
        />
      </v-col>
    </v-row>
  </v-container>
</template>
<script lang="ts">
import { Vue, Component } from 'vue-property-decorator';
import { Contract } from 'web3-eth-contract';
import PawnCreator from '@/components/PawnCreator.vue';
import PawnList from '@/components/PawnList.vue';
import PawningShop from '@/contracts/PawningShop.json';
import { getContractInstance } from '@/utils/contract';

import { pawn, PawnState } from '@/store/PawnState.module';

@Component({
  components: { PawnCreator, PawnList },
  name: 'Borrower',
  data: () => ({
    pawn,
  }),
})
export default class extends Vue {
  accounts: string[] = [];

  networkId = 0;

  whiteList: string[] = [];

  localLoading = false;

  pawn!: PawnState;

  pawningShopContract!: Contract;

  get loading(): boolean {
    return this.pawn.loading || this.localLoading;
  }

  async createPawn(data: any): Promise<void> {
    this.localLoading = true;
    console.log(data);
    const res = await this.pawningShopContract.methods.createPawn(data.tokenAddress, data.tokenId)
      .send({ from: this.accounts[0] });
    console.log(res);
    setTimeout(() => this.pawn.findAllByCreatorAddress(this.accounts[0]), 2000);
    this.localLoading = false;
  }

  async getNetworkId(): Promise<number> {
    return this.$web3.eth.net.getId();
  }

  async getAccounts(): Promise<string[]> {
    return this.$web3.eth.getAccounts();
  }

  async getWhiteList(): Promise<string[]> {
    this.localLoading = true;
    const res: string[] = await this.pawningShopContract.methods.getWhiteList().call(); // eslint-disable-line
    this.localLoading = false;
    return res;
  }

  async created() {
    this.localLoading = true;
    this.accounts = await this.getAccounts();
    this.networkId = await this.getNetworkId();
    if (this.networkId !== 5777) {
      console.log('you are in wrong network babe :D');
    }
    this.pawningShopContract = getContractInstance(PawningShop, this.networkId, this.$web3);
    this.whiteList = await this.getWhiteList();
    this.pawn.findAllByCreatorAddress(this.accounts[0]);
    this.localLoading = false;
  }
}
</script>
