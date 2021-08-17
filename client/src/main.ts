import Vue from 'vue';
import App from './App.vue';
import './registerServiceWorker';
import router from './router';
import store from './store';
import Web3Plugin from './plugins/web3';
import vuetify from './plugins/vuetify';

Vue.config.productionTip = false;
Vue.use(Web3Plugin);

import web3 from "web3";

declare module 'vue/types/vue' {
  interface Vue {
    $web3: web3;
  }
}

new Vue({
  router,
  store,
  vuetify,
  render: (h) => h(App),
}).$mount('#app');
