import _Vue from 'vue';
import axios from 'axios';

axios.defaults.baseURL = process.env.API_HOST;

export default async function myPlugin(Vue: typeof _Vue) {
  window.axios = axios;
}
