import _Vue from 'vue';
import { io } from 'socket.io-client';

export default async function myPlugin(Vue: typeof _Vue):Promise<void> {
  const host = process.env.VUE_APP_SOCKET_HOST as string;
  const instance = io(host);
  Vue.prototype.$socket = instance; //  eslint-disable-line
}
