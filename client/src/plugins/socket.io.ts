import _Vue from 'vue';
import { io } from 'socket.io-client';

export default async function myPlugin(Vue: typeof _Vue):Promise<void> {
  const host = process.env.VUE_APP_HOST as string;
  const socketPath = process.env.VUE_APP_SOCKET_PATH as string;
  console.log(host);
  const instance = io(host, { path: `${socketPath}/socket.io` });
  Vue.prototype.$socket = instance; //  eslint-disable-line
}
