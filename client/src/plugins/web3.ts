import _Vue from 'vue';
import getWeb3 from './getWeb3';

// export class Web3Plugin {
//   async install(vue: typeof _Vue, options: any) {
//     vue.prototype.$web3 = await getWeb3(); //  eslint-disable-line
//   }
// };

export default async function myPlugin(Vue: typeof _Vue) {
  Vue.prototype.$web3 = await getWeb3(); //  eslint-disable-line
}
