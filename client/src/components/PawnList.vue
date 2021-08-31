<template>
  <v-sheet>
    <bid-create-dialog v-if="selectedPawn" v-model="dialog" @submit-bid="submitBid"
      :pawn="selectedPawn"
    />
    <v-list three-line>
      <v-subheader>Your pawns</v-subheader>
      <v-progress-linear indeterminate v-if="loading"></v-progress-linear>
      <v-divider/>
      <v-list-item v-for="pawn in pawns" :key="pawn.id">
        <v-list-item-avatar :color="pawn.color">
          <v-avatar>{{ pawn.id }}</v-avatar>
        </v-list-item-avatar>
        <v-list-item-content>
          <v-list-item-title>
            Token address: {{ pawn.token_address }}, Token ID: {{ pawn.token_id }}
          </v-list-item-title>
          <v-list-item-subtitle>
            <p>Pawn status: {{ pawn.statusName }}</p>
            <p v-if="viewer === lender">Created by {{ pawn.creator.name }}</p>
          </v-list-item-subtitle>
        </v-list-item-content>
        <v-list-item-action>
          <v-btn
            v-if="pawn.status === 0 && viewer === borrower"
            @click="cancelPawn(pawn.id)"
          >
            Cancel
          </v-btn>
          <v-btn
            v-if="pawn.status === 0 && viewer === lender"
            @click="showBidDialog(pawn.id)"
          >
            Bid
          </v-btn>
        </v-list-item-action>
      </v-list-item>
    </v-list>
  </v-sheet>
</template>
<script lang="ts">
import { Component, Prop, Vue } from 'vue-property-decorator';
import { Contract } from 'web3-eth-contract';
import BidCreateDialog from '@/components/BidCreateDialog.vue';
import { BidCreate } from '@/store/models/bid';
import { ComputedPawn } from '@/store/PawnState.module';

@Component({
  name: 'PawnList',
  components: { BidCreateDialog },
})
export default class PawnList extends Vue {
  @Prop({ required: true, type: Array }) private pawns!: ComputedPawn[];

  @Prop({ required: true, type: Contract }) private pawningShopContract!: Contract;

  @Prop({ required: true }) private accounts!: string[];

  @Prop({
    type: String,
    default: 'borrower',
    validator: (val) => ['borrower', 'lender'].includes(val),
  }) private viewer!: string;

  borrower = 'borrower';

  lender = 'lender';

  localLoading = false;

  dialog = false;

  pawnId = '';

  get loading(): boolean {
    return this.localLoading;
  }

  get selectedPawn(): ComputedPawn | undefined {
    return this.pawns.find((pawn) => pawn.id === this.pawnId);
  }

  async cancelPawn(pawnId: string): Promise<void> {
    this.localLoading = true;
    const res = await this.pawningShopContract.methods.cancelPawn(pawnId)
      .send({ from: this.accounts[0] });
    console.log(res);
    this.localLoading = false;
  }

  async submitBid(data: BidCreate): Promise<void> {
    this.dialog = false;
    this.localLoading = true;
    const {
      loanAmount, isInterestProRated, loanDuration, interest,
    } = data;
    const res = await this.pawningShopContract.methods.createBid(
      interest, loanDuration, isInterestProRated, this.pawnId,
    )
      .send({ from: this.accounts[0], value: loanAmount });
    console.log({ data, res });
    this.localLoading = false;
  }

  showBidDialog(pawnId: string): void {
    this.dialog = true;
    this.pawnId = pawnId;
  }
}
</script>
