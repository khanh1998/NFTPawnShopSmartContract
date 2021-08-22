<template>
  <div>
    <v-select :items="whiteList" dense v-model="tokenAddress" label="token address" />
    <v-text-field v-model="tokenId" label="token id" />
    <v-btn @click="createPawn">Submit</v-btn>
  </div>
</template>
<script lang="ts">
import {
  Vue, Component, Emit, Prop,
} from 'vue-property-decorator';

@Component({
  name: 'PawnCreator',
})
export default class extends Vue {
  @Prop({ type: Array, required: true }) whiteList = [];

  tokenAddress = ''

  tokenId = ''

  @Emit()
  createPawn(): any {
    return {
      tokenAddress: this.tokenAddress,
      tokenId: this.tokenId,
    };
  }
}
</script>
