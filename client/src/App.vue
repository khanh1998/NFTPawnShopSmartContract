<template>
  <v-app id="inspire">
    <v-app-bar app color="white" flat>
      <v-avatar class="hidden-sm-and-up" :color="getRandomColor()" size="32" :title="fullName">
        {{ firstChar }}
      </v-avatar>

      <v-tabs centered class="ml-n9" color="grey darken-1">
        <v-tab v-for="link in links" :key="link[0]">
          <router-link :to="link[0]">{{ link[1] }}</router-link>
        </v-tab>
      </v-tabs>

      <v-avatar class="hidden-sm-and-down" :title="fullName" :color="getRandomColor()" size="32">
        {{ firstChar }}
      </v-avatar>
    </v-app-bar>
    <v-snackbar v-model="snackbar.show">
      {{ snackbar.text }}

      <template v-slot:action="{ attrs }">
        <v-btn color="red" text v-bind="attrs" @click="snackbar.show = false"> Close </v-btn>
      </template>
    </v-snackbar>
    <v-main class="grey lighten-3">
      <router-view />
    </v-main>
  </v-app>
</template>

<script lang="ts">
import { Vue, Component } from 'vue-property-decorator';
import { user, UserState } from '@/store/UserState.module';
import { getRandomColor } from '@/utils/color';
import {
  EventData, BidPayload, EventName, PawnPayload,
} from '@/store/models/eventData';

@Component({
  name: 'App',
  data: () => ({
    user,
  }),
  methods: {
    getRandomColor,
  },
})
export default class extends Vue {
  user!: UserState;

  account!: string;

  snackbar: any = {
    show: false,
    text: '',
  };

  links = [
    ['/', 'Home'],
    ['/owner', 'Owner'],
    ['/borrower', 'Borrower'],
    ['/lender', 'Lender'],
    ['/test-token', 'Test Token'],
  ];

  get isLoading() {
    return this.user.loading;
  }

  get fullName(): string | undefined {
    return this.user.data?.name;
  }

  get firstChar(): string | undefined {
    return this.fullName?.charAt(0).toUpperCase();
  }

  closeSnackbar() {
    this.snackbar.show = false;
    this.snackbar.text = '';
  }

  async mounted() {
    this.$socket.emit('data_update', 'hello from client');
    this.$socket.on('data_update', (args) => {
      console.log(args);
      const data = args as EventData;
      let bidPayload!: BidPayload;
      let pawnPayload!: PawnPayload;

      switch (data.code) {
        case EventName.BidCreated:
          bidPayload = JSON.parse(data.payload) as BidPayload;
          console.log(bidPayload);
          break;
        case EventName.PawnCreated:
          pawnPayload = JSON.parse(data.payload) as PawnPayload;
          console.log(pawnPayload);
          break;
        default:
      }
      this.snackbar.show = true;
      this.snackbar.text = data.message;
    });
    const accounts = await this.$web3.eth.getAccounts();
    [this.account] = accounts;
    this.user.findUserByAddress(accounts[0]);
  }
}
</script>
