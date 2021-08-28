<template>
  <v-card>
    <v-card-title>Create a new Pawn</v-card-title>
    <v-card-text>
      <v-select :items="whiteList" dense v-model="tokenAddress" label="token address" />
      <v-text-field v-model="tokenId" label="token id" />
    </v-card-text>
    <v-card-actions>
      <v-btn @click="createPawn">Submit</v-btn>
    </v-card-actions>
  </v-card>
</template>
<script lang="ts">
import {
  Vue, Component, Emit, Prop,
} from 'vue-property-decorator';

@Component({
  name: 'PawnCreator',
})
export default class extends Vue {
  @Prop({ type: Array, required: true }) whiteList!: Array<string>;

  tokenAddress = '';

  tokenId = '';

  @Emit()
  createPawn(): any {
    return {
      tokenAddress: this.tokenAddress,
      tokenId: this.tokenId,
    };
  }
}
</script>
