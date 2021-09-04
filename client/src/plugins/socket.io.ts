import _Vue from 'vue';
import { io } from 'socket.io-client';

export default async function myPlugin(Vue: typeof _Vue) {
  // process.env.VUE_APP_SOCKET_HOST
  const instance = io('http://localhost:7789');
  Vue.prototype.$socket = instance; //  eslint-disable-line
}
