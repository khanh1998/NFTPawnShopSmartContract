import Vue from 'vue';
import web3 from 'web3';

declare module 'vue/types/vue' {
  interface Vue {
    $web3: web3;
  }
}
