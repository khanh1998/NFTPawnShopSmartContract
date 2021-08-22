<template>
  <v-container>
    <v-row>
      <v-col>
        <h2>To query and interact with our ERC721 Test Token</h2>
        <p>Welcome {{ accounts[0] }}</p>
      </v-col>
    </v-row>
    <v-row v-if="!loading">
      <v-col>
        <v-card>
          <v-card-title> Balance Of </v-card-title>
          <v-card-text>
            <v-text-field label="account address" v-model="address" />
            <p>balance: {{ balance }}</p>
          </v-card-text>
          <v-card-actions>
            <v-btn @click="balanceOf">get</v-btn>
          </v-card-actions>
        </v-card>
      </v-col>
      <v-col>
        <v-card>
          <v-card-title> Mint </v-card-title>
          <v-card-text>
            <v-text-field label="account address" v-model="address" />
            <p>{{ mintMessage }}</p>
          </v-card-text>
          <v-card-actions>
            <v-btn @click="mint">Mint</v-btn>
          </v-card-actions>
        </v-card>
      </v-col>
    </v-row>
    <v-row v-if="!loading">
      <v-col>
        <v-card>
          <v-card-title> Owner Of </v-card-title>
          <v-card-text>
            <v-text-field label="Token ID" v-model="tokenId" />
            <p>{{ ownerOfMessage }}</p>
          </v-card-text>
          <v-card-actions>
            <v-btn @click="ownerOf">Get</v-btn>
          </v-card-actions>
        </v-card>
      </v-col>
      <v-col>
        <v-card>
          <v-card-title> Approve for all </v-card-title>
          <v-card-text>
            <v-text-field label="Approve to" v-model="approveTo" />
            <v-text-field label="Token ID" v-model="approveTokenId" />
            <p>{{ approveMessage }}</p>
          </v-card-text>
          <v-card-actions>
            <v-btn @click="approve">Approve</v-btn>
          </v-card-actions>
        </v-card>
      </v-col>
    </v-row>
    <v-row v-if="loading">
      <v-col>
        <p>Loading...</p>
      </v-col>
    </v-row>
  </v-container>
</template>
<script lang="ts">
import { Vue, Component } from 'vue-property-decorator';
import { Contract } from 'web3-eth-contract';
import TestToken from '../contracts/TestToken.json';

@Component({
  name: 'TestToken',
})
export default class extends Vue {
  loading = false;

  networkId = 5777;

  accounts: string[] = [];

  address = '';

  balance = -1;

  mintMessage = '';

  ownerOfMessage = '';

  tokenId = -1;

  approveTo = ''

  approveTokenId = -1

  approveMessage = ''

  getContractInstance(contractJson: any, networkId: number): Contract {
    const deployedNetwork = contractJson.networks[networkId];
    const instance: Contract = new this.$web3.eth.Contract(
      contractJson.abi,
      deployedNetwork && deployedNetwork.address,
    );
    return instance;
  }

  async balanceOf(): Promise<void> {
    this.balance = -1;
    this.loading = true;
    const testToken = this.getContractInstance(TestToken, this.networkId);
    const res: string = await testToken.methods.balanceOf(this.address).call(); // eslint-disable-line
    console.log(res);
    this.balance = Number(res);
    this.loading = false;
  }

  async mint(): Promise<void> {
    try {
      this.balance = -1;
      this.loading = true;
      this.mintMessage = '';
      const testToken = this.getContractInstance(TestToken, this.networkId);
      await testToken.methods.mint(this.address).send({ from: this.accounts[0] }); // eslint-disable-line
      this.mintMessage = `mint new token success for ${this.accounts[0]}`;
    } catch (error) {
      this.mintMessage = 'mint new token fail';
    }
    this.loading = false;
  }

  async approve(): Promise<void> {
    try {
      this.loading = true;
      this.approveMessage = '';
      const testToken = this.getContractInstance(TestToken, this.networkId);
      await testToken.methods.approve(this.approveTo, this.approveTokenId)
        .send({ from: this.accounts[0] }); // eslint-disable-line
      this.approveMessage = `approve success token ID ${this.approveTokenId} for ${this.approveTo}`;
    } catch (error) {
      this.approveMessage = 'fail to approve';
    }
    this.loading = false;
  }

  async ownerOf(): Promise<void> {
    try {
      this.loading = true;
      this.mintMessage = '';
      const testToken = this.getContractInstance(TestToken, this.networkId);
      const ownerOfToken = await testToken.methods.ownerOf(this.tokenId).call(); // eslint-disable-line
      if (ownerOfToken === this.accounts[0]) {
        this.ownerOfMessage = `Owner of token ID ${this.tokenId} is you :)))`;
      } else {
        this.ownerOfMessage = `Owner of token ID ${this.tokenId} is ${ownerOfToken}`;
      }
    } catch (error) {
      this.ownerOfMessage = 'Token ID is not existed';
    }
    this.loading = false;
  }

  async getNetworkId(): Promise<number> {
    return this.$web3.eth.net.getId();
  }

  async getAccounts(): Promise<string[]> {
    return this.$web3.eth.getAccounts();
  }

  async created() {
    this.loading = true;
    this.networkId = await this.getNetworkId();
    this.accounts = await this.getAccounts();
    [this.address] = this.accounts;
    this.loading = false;
  }
}
</script>
