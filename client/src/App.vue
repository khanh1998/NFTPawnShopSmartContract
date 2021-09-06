<template>
  <v-app id="inspire">
    <v-app-bar app color="white" flat>
      <v-avatar
        class="hidden-sm-and-up"
        :color="getRandomColor()"
        size="32"
        :title="fullName"
      >
      {{ firstChar }}
      </v-avatar>

      <v-tabs
        centered
        class="ml-n9"
        color="grey darken-1"
      >
        <v-tab
          v-for="link in links"
          :key="link[0]"
        >
          <router-link :to="link[0]">{{link[1]}}</router-link>
        </v-tab>
      </v-tabs>

      <v-avatar class="hidden-sm-and-down" :title="fullName" :color="getRandomColor()" size="32">
        {{ firstChar }}
      </v-avatar>
    </v-app-bar>

    <v-main class="grey lighten-3">
      <router-view />
    </v-main>
  </v-app>
</template>

<script lang="ts">
import { Vue, Component } from 'vue-property-decorator';
import { user, UserState } from '@/store/UserState.module';
import { getRandomColor } from '@/utils/color';

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

  async mounted() {
    this.$socket.emit('data_update', 'hello from client');
    this.$socket.on('data_update', (args) => console.log(args));
    const accounts = await this.$web3.eth.getAccounts();
    console.log(accounts[0]);
    this.user.findUserByAddress(accounts[0]);
  }
}
</script>
