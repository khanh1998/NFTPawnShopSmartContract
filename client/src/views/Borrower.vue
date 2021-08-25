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
        <p>My pawns</p>
      </v-col>
    </v-row>
  </v-container>
</template>
<script lang="ts">
import { Vue, Component } from 'vue-property-decorator';
import { Contract } from 'web3-eth-contract';
import PawnCreator from '@/components/PawnCreator.vue';
import PawningShop from '../contracts/PawningShop.json';
import { IPawnState } from '@/store/IPawnState';

import { pawn, PawnState } from '@/store/pawn.vuex';

@Component({
  components: { PawnCreator },
  name: 'Borrower',
  data: () => ({
    pawn,
  }),
})
export default class extends Vue {
  accounts: string[] = [];

  networkId = 0;

  whiteList: string[] = [];

  data!: IPawnState[];

  error!: Error;

  localLoading = false;

  pawn!: PawnState;

  get loading(): boolean {
    return this.pawn.loading || this.localLoading;
  }

  getContractInstance(contractJson: any, networkId: number): Contract {
    const deployedNetwork = contractJson.networks[networkId];
    const instance: Contract = new this.$web3.eth.Contract(
      contractJson.abi,
      deployedNetwork && deployedNetwork.address,
    );
    return instance;
  }

  async createPawn(data: any): Promise<void> {
    this.localLoading = true;
    console.log(data);
    const pawningShop = this.getContractInstance(PawningShop, this.networkId);
    const res = await pawningShop.methods.createPawn(data.tokenAddress, data.tokenId)
      .send({ from: this.accounts[0] });
    console.log(res);
    this.localLoading = false;
  }

  async getNetworkId(): Promise<number> {
    return this.$web3.eth.net.getId();
  }

  async getAccounts(): Promise<string[]> {
    return this.$web3.eth.getAccounts();
  }

  async getWhiteList(): Promise<string[]> {
    const pawningShop = this.getContractInstance(PawningShop, this.networkId);
    const res: string[] = await pawningShop.methods.getWhiteList().call(); // eslint-disable-line
    console.log(res);
    return res;
  }

  async created() {
    this.accounts = await this.getAccounts();
    this.networkId = await this.getNetworkId();
    // this.whiteList = await this.getWhiteList();
    this.pawn.findAllByCreatorAddress(this.accounts[0]);
  }
}
</script>
