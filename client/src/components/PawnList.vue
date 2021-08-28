<template>
  <v-list three-line>
    <v-subheader>Your pawns</v-subheader>
    <v-divider/>
    <v-list-item v-for="pawn in colorPawns" :key="pawn.id">
      <v-list-item-avatar :color="pawn.color">
        <v-avatar>{{ pawn.id }}</v-avatar>
      </v-list-item-avatar>
      <v-list-item-content>
        <v-list-item-title>
          Token address: {{ pawn.token_address }}, Token ID: {{ pawn.token_id }}
        </v-list-item-title>
        <v-list-item-subtitle>
          Pawn status: {{ pawn.status }}
        </v-list-item-subtitle>
      </v-list-item-content>
      <v-list-item-action>
        <v-btn v-if="pawn.status === 0" @click="cancelPawn(pawn.id)">
          Cancel
        </v-btn>
      </v-list-item-action>
    </v-list-item>
  </v-list>
</template>
<script lang="ts">
import { Component, Prop, Vue } from 'vue-property-decorator';
import { Contract } from 'web3-eth-contract';
import { getRandomColor } from '@/utils/color';
import { Pawn } from '@/store/models/pawn';

interface ColorPawn extends Pawn {
  color : string;
}

@Component
export default class PawnList extends Vue {
  @Prop({ required: true, type: Array }) private pawns!: Pawn[];

  @Prop({ required: true, type: Contract }) private pawningShopContract!: Contract;

  @Prop({ required: true }) private accounts!: string[]

  localLoading = false;

  get colorPawns(): ColorPawn[] {
    return this.pawns.map((pawn) => ({
      ...pawn,
      color: getRandomColor(),
    }));
  }

  async cancelPawn(pawnId: string): Promise<void> {
    this.localLoading = true;
    const res = await this.pawningShopContract.methods.cancelPawn(pawnId)
      .send({ from: this.accounts[0] });
    console.log(res);
    this.localLoading = false;
  }
}
</script>
