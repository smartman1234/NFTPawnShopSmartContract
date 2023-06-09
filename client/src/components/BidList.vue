<template>
  <v-sheet>
    <v-list three-line>
      <v-subheader>Your bids</v-subheader>
      <v-progress-linear indeterminate v-if="loading"></v-progress-linear>
      <v-divider />
      <v-list-item v-for="bid in bids" :key="bid.id">
        <v-list-item-avatar :color="bid.color">
          <v-avatar>{{ bid.id }}</v-avatar>
        </v-list-item-avatar>
        <v-list-item-content>
          <v-list-item-title> Pawn id: {{ bid.pawn }} </v-list-item-title>
          <p>Loan amount: {{ bid.loan_amount }}</p>
          <p>Loan interest: {{ bid.interest }}</p>
          <p v-if="bid.status === 2">Repayment date: {{ bid.loan_end_time_str }}</p>
          <p v-if="viewer === borrower">Created by {{ bid.creator }}</p>
        </v-list-item-content>
        <v-list-item-action>
          <v-btn v-if="viewer === lender && bid.status === 0" @click="cancelBid(bid.id)">
            Cancel
          </v-btn>
          <v-btn v-if="viewer === borrower && bid.status == 0" @click="acceptBid(bid.id)">
            Accept
          </v-btn>
        </v-list-item-action>
      </v-list-item>
    </v-list>
  </v-sheet>
</template>
<script lang="ts">
import { Component, Prop, Vue } from 'vue-property-decorator';
import { Contract } from 'web3-eth-contract';
import { Bid } from '@/store/models/bid';
import { extractErrorObjFromMessage } from '@/utils/contract';

@Component({
  name: 'BidList',
})
export default class PawnList extends Vue {
  @Prop({ required: true, type: Array }) private bids!: Bid[];

  @Prop({ required: true, type: Contract }) private pawningShopContract!: Contract;

  @Prop({ required: true }) private accounts!: string[];

  @Prop({
    type: String,
    default: 'borrower',
    validator: (val) => ['borrower', 'lender'].includes(val),
  })
  private viewer!: string;

  borrower = 'borrower';

  lender = 'lender';

  localLoading = false;

  get loading(): boolean {
    return this.localLoading;
  }

  async cancelBid(bidId: string): Promise<void> {
    this.localLoading = true;
    try {
      const res = await this.pawningShopContract.methods
        .cancelBid(bidId)
        .send({ from: this.accounts[0] });
      console.log(res);
    } catch (error) {
      const err = extractErrorObjFromMessage(error.message);
      alert(err.reason);
    }
    this.localLoading = false;
  }

  async acceptBid(bidId: string): Promise<void> {
    this.localLoading = true;
    try {
      const res = await this.pawningShopContract.methods
        .acceptBid(bidId)
        .send({ from: this.accounts[0] });
      console.log(res);
    } catch (error) {
      const err = extractErrorObjFromMessage(error.message);
      alert(err.reason);
    }
    this.localLoading = false;
  }
}
</script>
