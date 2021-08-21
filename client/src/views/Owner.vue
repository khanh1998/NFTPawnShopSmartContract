<template>
  <v-container>
    <v-row>
      <p v-if="loading">app loading...</p>
    <div v-if="!loading">
      <div>
        Current account<p v-for="acc in accounts" :key="acc">{{ acc }}</p>
      </div>
      <div class="d-flex flex-row">
        <v-card>
          <v-list>
            <v-list-item v-for="add in whiteList" :key="add">
              <v-list-item-title>{{ add }}</v-list-item-title>
            </v-list-item>
          </v-list>
        </v-card>
        <v-card>
          <v-card-title>Add address to white list</v-card-title>
          <v-text-field v-model="newAddress"/>
          <v-card-actions>
            <v-btn @click="addNewAddress">Submit</v-btn>
          </v-card-actions>
        </v-card>
        <v-card>
          <v-card-title>Remove address from white list</v-card-title>
          <v-text-field v-model="newAddress"/>
          <v-card-actions>
            <v-btn @click="removeAddress">Submit</v-btn>
          </v-card-actions>
        </v-card>
      </div>
      <div>
      </div>
    </div>
    </v-row>
  </v-container>
</template>
<script lang="ts">
import { Vue, Component } from 'vue-property-decorator';
import { Contract } from 'web3-eth-contract';
import PawningShop from '../contracts/PawningShop.json';

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

  getContractInstance(contractJson: any, networkId: number): Contract {
    const deployedNetwork = contractJson.networks[networkId];
    const instance: Contract = new this.$web3.eth.Contract(contractJson.abi,
      deployedNetwork && deployedNetwork.address);
    return instance;
  }

  async addNewAddress(): Promise<void> {
    this.loading = true;
    const pawningShop = this.getContractInstance(PawningShop, this.networkId);
    const res = await pawningShop.methods.addToWhiteList(this.newAddress)
      .send({ from: this.accounts[0] });
    console.log(res);
    this.loading = false;
    this.newAddress = '';
  }

  async getWhiteList(): Promise<string[]> {
    this.loading = true;
    const pawningShop = this.getContractInstance(PawningShop, this.networkId);
    const res: string[] = await pawningShop.methods.getWhiteList().call(); // eslint-disable-line
    console.log(res);
    this.loading = false;
    return res;
  }

  async removeAddress(): Promise<void> {
    this.loading = true;
    const pawningShop = this.getContractInstance(PawningShop, this.networkId);
    const res = await pawningShop.methods.removeFromWhiteList(this.newAddress)
      .send({ from: this.accounts[0] });
    console.log(res);
    this.loading = false;
    this.newAddress = '';
  }

  async getOwner(): Promise<string> {
    this.loading = true;
    const pawningShop = this.getContractInstance(PawningShop, this.networkId);
    console.log(pawningShop);
    const res: string = await pawningShop.methods.owner().call();
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
    this.owner = await this.getOwner();
    this.whiteList = await this.getWhiteList();
    this.loading = false;
  }
}

</script>
