<template>
  <div>
    
  </div> 
</template>
<script lang="ts">
import { Vue, Component } from 'vue-property-decorator';

@Component({
  name: "Owner",
})
export default class extends Vue {
  loading: boolean = false
  async getWhiteList() {
    this.loading = true
    const accounts = await this.$web3.eth.getAccounts()
    console.log(accounts)
    this.loading = false
  }

  created() {
    this.getWhiteList()
  }
}

</script>