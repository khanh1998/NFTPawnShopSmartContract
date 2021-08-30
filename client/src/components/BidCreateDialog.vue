<template>
  <v-dialog v-model="value" persistent max-width="290">
    <v-card>
      <v-card-title class="text-h5"> Make a bid for Pawn id: {{ pawn.id }} </v-card-title>
      <v-card-text>
        <v-text-field v-model="data.loanAmount" label="Loan amount" type="number" suffix="wei" />
        <v-text-field v-model="data.interest" label="Interest" type="number" suffix="wei" />
        <v-switch
          v-model="data.isInterestProRated"
          inset
          label="Is interest pro rated"
        ></v-switch>
        <v-text-field
          v-model="data.loanDuration"
          label="Loan duration"
          type="number"
          suffix="Day"
        />
        <v-menu
          v-model="datePicker"
          :close-on-content-click="false"
          transition="scale-transition"
          offset-y
          max-width="290px"
          min-width="auto"
        >
          <template v-slot:activator="{ on, attrs }">
            <v-text-field
              v-model="dateString"
              label="Loan start date"
              hint="MM/DD/YYYY format"
              persistent-hint
              prepend-icon="mdi-calendar"
              readonly
              v-bind="attrs"
              v-on="on"
            ></v-text-field>
          </template>
          <v-date-picker
            v-model="dateString"
            no-title
            @input="datePicker = false"
          ></v-date-picker>
        </v-menu>
      </v-card-text>
      <v-card-subtitle>
        Repay amount: {{ repayAmount }} wei, in {{ repayDate }}
      </v-card-subtitle>
      <v-card-actions>
        <v-spacer></v-spacer>
        <v-btn color="green darken-1" text @click="closeDialog"> Cancel </v-btn>
        <v-btn color="green darken-1" text @click="submitBid"> Agree </v-btn>
      </v-card-actions>
    </v-card>
  </v-dialog>
</template>
<script lang="ts">
import {
  Component, Emit, Prop, Vue, Watch,
} from 'vue-property-decorator';
import { BidCreate } from '@/store/models/bid';
import { ComputedPawn } from '@/store/PawnState.module';

@Component({
  name: 'BidCreateDialog',
})
export default class BidCreateDialog extends Vue {
  @Prop({ type: Boolean, default: false, required: true }) value!: boolean;

  @Prop({ required: true }) pawn!: ComputedPawn;

  data: BidCreate = {
    loanAmount: 0,
    interest: 0,
    loanStartTime: 0,
    loanDuration: 0,
    isInterestProRated: false,
  }

  datePicker = false;

  dateString = ''

  get repayAmount() {
    return Number(this.data.loanAmount) + Number(this.data.interest);
  }

  get repayDate(): string {
    const startDate = new Date(this.dateString);
    startDate.setDate(startDate.getDate() + Number(this.data.loanDuration));
    return startDate.toLocaleDateString('vi');
  }

  closeDialog():void {
    this.input();
  }

  @Watch('dateString')
  dateStringChange(val: string):void {
    this.data.loanStartTime = new Date(val).getTime() / 1000;
  }

  @Emit()
  input(): boolean {
    return !this.value;
  }

  @Emit()
  submitBid(): BidCreate {
    this.closeDialog();
    return this.data;
  }
}
</script>
