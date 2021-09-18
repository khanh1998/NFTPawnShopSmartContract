import _Vue from 'vue';
import axios from 'axios';

const apiUri = `${process.env.VUE_APP_HOST}${process.env.VUE_APP_API_PATH}`;
const axiosConfig = {
  baseURL: apiUri,
  timeout: 30000,
};

console.log(apiUri);

const axiosInstance = axios.create(axiosConfig);

export default async function myPlugin(Vue: typeof _Vue) {
  window.axios = axiosInstance;
}
