import Vue from 'vue';
import App from './App.vue';
import './registerServiceWorker';
import router from './router';
import { store } from './store/index';
import web3 from './plugins/web3';
import vuetify from './plugins/vuetify';
import axios from './plugins/axios';
import socketio from './plugins/socket.io';

Vue.config.productionTip = false;
Vue.use(web3);
Vue.use(axios);
Vue.use(socketio);

new Vue({
  router,
  store,
  vuetify,
  render: (h) => h(App),
}).$mount('#app');
