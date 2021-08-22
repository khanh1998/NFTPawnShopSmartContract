<template>
  <v-container>
    <v-row>
      <v-col>
        <p>Hello borrower!!! {{ accounts[0] }}</p>
      </v-col>
    </v-row>
    <v-row v-if="!loading">
      <pawn-creator @create-pawn="createPawn" :white-list="whiteList"/>
    </v-row>
  </v-container>
</template>
<script lang="ts">
import { Vue, Component } from 'vue-property-decorator';
import { Contract } from 'web3-eth-contract';
import PawnCreator from '@/components/PawnCreator.vue';
import PawningShop from '../contracts/PawningShop.json';

@Component({
  components: { PawnCreator },
  name: 'Borrower',
})
export default class extends Vue {
  accounts: string[] = [];

  networkId = 0;

  whiteList: string[] = [];

  loading = false;

  getContractInstance(contractJson: any, networkId: number): Contract {
    const deployedNetwork = contractJson.networks[networkId];
    const instance: Contract = new this.$web3.eth.Contract(
      contractJson.abi,
      deployedNetwork && deployedNetwork.address,
    );
    return instance;
  }

  async createPawn(data: any): Promise<void> {
    this.loading = true;
    console.log(data);
    const pawningShop = this.getContractInstance(PawningShop, this.networkId);
    const res = await pawningShop.methods.createPawn(data.tokenAddress, data.tokenId)
      .send({ from: this.accounts[0] });
    console.log(res);
    this.loading = false;
  }

  async getNetworkId(): Promise<number> {
    return this.$web3.eth.net.getId();
  }

  async getAccounts(): Promise<string[]> {
    return this.$web3.eth.getAccounts();
  }

  async getWhiteList(): Promise<string[]> {
    this.loading = true;
    const pawningShop = this.getContractInstance(PawningShop, this.networkId);
    const res: string[] = await pawningShop.methods.getWhiteList().call(); // eslint-disable-line
    console.log(res);
    this.loading = false;
    return res;
  }

  async created() {
    this.loading = true;
    this.accounts = await this.getAccounts();
    this.networkId = await this.getNetworkId();
    this.whiteList = await this.getWhiteList();
    this.loading = false;
  }
}
</script>
