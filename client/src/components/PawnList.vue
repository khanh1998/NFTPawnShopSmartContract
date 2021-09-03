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
            Token ID: {{ pawn.token_id }}
          </v-list-item-title>
          <p>Token address: {{ pawn.token_address }}</p>
          <p>Pawn status: {{ pawn.statusName }}</p>
          <p>Bid: {{ pawn.bids.length }}</p>
          <p v-if="pawn.acceptedBid">{{ pawn.acceptedBid.loan_amount }}</p>
          <p v-if="viewer === lender">Created by {{ pawn.creator.name }}</p>
        </v-list-item-content>
        <v-list-item-action>
          <v-btn
            v-if="pawn.status === 0 && viewer === borrower"
            @click="cancelPawn(pawn.id)"
            small
          >
            Cancel
          </v-btn>
          <v-btn
            v-if="pawn.status === 0 && viewer === lender"
            @click="showBidDialog(pawn.id)"
            small
          >
            Bid
          </v-btn>
          <v-btn
            v-if="viewer === borrower && pawn.status == 2"
            @click="repaid(pawn)"
            small
          >
            Repay
          </v-btn>
          <v-btn
            v-if="viewer === lender && pawn.status == 2"
            @click="liquidate(pawn)"
            small
          >
            Liquidate
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
import { calculateRepaidAmount, extractErrorObjFromMessage } from '@/utils/contract';
import { ComputedPawn } from '@/store/models/pawn';

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

  async liquidate(pawn: ComputedPawn): Promise<void> {
    this.localLoading = false;
    try {
      const res = await this.pawningShopContract.methods.liquidate(pawn.id)
        .send({ from: this.accounts[0] });
      console.log(res);
    } catch (error) {
      const err = extractErrorObjFromMessage(error.message);
      alert(err.reason);
    }
    this.localLoading = true;
  }

  async repaid(pawn: ComputedPawn): Promise<void> {
    this.localLoading = false;
    const { acceptedBid } = pawn;
    if (!acceptedBid) {
      console.log(`Pawn ${pawn.id} has no accepted bid`);
      return;
    }
    try {
      const repayAmount = calculateRepaidAmount(pawn);
      console.log(repayAmount);
      const res = await this.pawningShopContract.methods.repaid(pawn.id)
        .send({ from: this.accounts[0], value: repayAmount });
      console.log(res);
    } catch (error) {
      const err = extractErrorObjFromMessage(error.message);
      alert(err.reason);
    }
    this.localLoading = true;
  }

  async cancelPawn(pawnId: string): Promise<void> {
    this.localLoading = true;
    try {
      const res = await this.pawningShopContract.methods.cancelPawn(pawnId)
        .send({ from: this.accounts[0] });
      console.log(res);
    } catch (error) {
      const err = extractErrorObjFromMessage(error.message);
      alert(err.reason);
    }
    this.localLoading = false;
  }

  async submitBid(data: BidCreate): Promise<void> {
    this.dialog = false;
    this.localLoading = true;
    try {
      const {
        loanAmount, isInterestProRated, loanDuration, interest,
      } = data;
      const res = await this.pawningShopContract.methods.createBid(
        interest, loanDuration, isInterestProRated, this.pawnId,
      )
        .send({ from: this.accounts[0], value: loanAmount });
      console.log({ data, res });
    } catch (error) {
      const err = extractErrorObjFromMessage(error.message);
      alert(err.reason);
    }
    this.localLoading = false;
  }

  showBidDialog(pawnId: string): void {
    this.dialog = true;
    this.pawnId = pawnId;
  }
}
</script>
