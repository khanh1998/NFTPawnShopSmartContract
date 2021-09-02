<template>
  <v-container>
    <v-row>
      <p v-if="loading">app loading...</p>
      <div>
        Current account<p v-for="acc in accounts" :key="acc">{{ acc }}</p>
      </div>
    </v-row>
    <v-row v-if="!loading">
      <v-col>
        <v-card>
          <v-list>
            <v-list-item v-for="add in whiteList" :key="add">
              <v-list-item-title>{{ add }}</v-list-item-title>
            </v-list-item>
          </v-list>
        </v-card>
      </v-col>
      <v-col>
        <v-card>
          <v-card-title>Add address to white list</v-card-title>
          <v-text-field v-model="newAddress"/>
          <v-card-actions>
            <v-btn @click="addNewAddress">Submit</v-btn>
          </v-card-actions>
        </v-card>
      </v-col>
      <v-col>
        <v-card>
          <v-card-title>Remove address from white list</v-card-title>
          <v-text-field v-model="newAddress"/>
          <v-card-actions>
            <v-btn @click="removeAddress">Submit</v-btn>
          </v-card-actions>
        </v-card>
      </v-col>
    </v-row>
  </v-container>
</template>
<script lang="ts">
import { Vue, Component } from 'vue-property-decorator';
import { Contract } from 'web3-eth-contract';
import PawningShop from '../contracts/PawningShop.json';
import { getContractInstance } from '@/utils/contract';

@Component({
  name: 'Owner',
})
export default class extends Vue {
  loading = false

  accounts: string[] = []

  whiteList: string[] = []

  networkId = 0

  newAddress = ''

  owner = ''

  pawningShopContract!: Contract;

  async addNewAddress(): Promise<void> {
    this.loading = true;
    const res = await this.pawningShopContract.methods.addToWhiteList(this.newAddress)
      .send({ from: this.accounts[0] });
    console.log(res);
    this.loading = false;
    this.newAddress = '';
  }

  async getWhiteList(): Promise<string[]> {
    this.loading = true;
    const res: string[] = await this.pawningShopContract.methods.getWhiteList().call(); // eslint-disable-line
    console.log(res);
    this.loading = false;
    return res;
  }

  async removeAddress(): Promise<void> {
    this.loading = true;
    const res = await this.pawningShopContract.methods.removeFromWhiteList(this.newAddress)
      .send({ from: this.accounts[0] });
    console.log(res);
    this.loading = false;
    this.newAddress = '';
  }

  async getOwner(): Promise<string> {
    this.loading = true;
    const res: string = await this.pawningShopContract.methods.owner().call();
    console.log(res);
    this.loading = false;
    return res;
  }

  async getNetworkId(): Promise<number> {
    return this.$web3.eth.net.getId();
  }

  async getAccounts(): Promise<string[]> {
    return this.$web3.eth.getAccounts();
  }

  async created() {
    this.loading = true;
    this.accounts = await this.getAccounts();
    this.networkId = await this.getNetworkId();
    this.pawningShopContract = getContractInstance(PawningShop, this.networkId, this.$web3);
    this.owner = await this.getOwner();
    this.whiteList = (await this.getWhiteList()).filter((item) => item !== '0x0000000000000000000000000000000000000000');
    this.loading = false;
  }
}

</script>
