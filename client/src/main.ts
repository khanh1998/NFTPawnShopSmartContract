import Vue from 'vue';
import App from './App.vue';
import './registerServiceWorker';
import router from './router';
import store from './store';
import web3 from './plugins/web3';

Vue.config.productionTip = false;
Vue.use(web3);

new Vue({
  router,
  store,
  render: (h) => h(App),
}).$mount('#app');
