import web3 from 'web3';
import Vue from 'vue';
import App from './App.vue';
import './registerServiceWorker';
import router from './router';
import store from './store';
import myPlugin from './plugins/web3';
import vuetify from './plugins/vuetify';

Vue.config.productionTip = false;
Vue.use(myPlugin);

new Vue({
  router,
  store,
  vuetify,
  render: (h) => h(App),
}).$mount('#app');
