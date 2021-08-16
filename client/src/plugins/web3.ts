import _Vue from 'vue';
import getWeb3 from './getWeb3';

export default {
  async install(vue: typeof _Vue, options: any) {
    vue.prototype.$web3 = await getWeb3(); //  eslint-disable-line
  },
};
