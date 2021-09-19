import _Vue from 'vue';
import { io } from 'socket.io-client';

export default async function myPlugin(Vue: typeof _Vue):Promise<void> {
  let instance;
  if (process.env.NODE_ENV === 'development') {
    const host = process.env.VUE_APP_SOCKET_HOST as string;
    console.log(host);
    instance = io(host);
  } else {
    const host = process.env.VUE_APP_HOST as string;
    const socketPath = process.env.VUE_APP_SOCKET_PATH as string;
    console.log(host, socketPath);
    instance = io(host, { path: `${socketPath}/socket.io` });
  }
  Vue.prototype.$socket = instance; //  eslint-disable-line
}
