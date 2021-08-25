import Vue from 'vue';
import App from './App.vue';
import './registerServiceWorker';
import router from './router';
import store from './store';
import web3 from './plugins/web3';
import vuetify from './plugins/vuetify';
import axios from './plugins/axios';

Vue.config.productionTip = false;
Vue.use(web3);
Vue.use(axios);

new Vue({
  router,
  store,
  vuetify,
  render: (h) => h(App),
}).$mount('#app');
