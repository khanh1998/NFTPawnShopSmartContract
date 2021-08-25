import _Vue from 'vue';
import axios from 'axios';

const axiosConfig = {
  baseURL: process.env.VUE_APP_API_HOST,
  timeout: 30000,
};

console.log(process.env.VUE_APP_API_HOST);

const axiosInstance = axios.create(axiosConfig);

export default async function myPlugin(Vue: typeof _Vue) {
  window.axios = axiosInstance;
}
