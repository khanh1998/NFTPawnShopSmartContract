import Vue from 'vue';
import { Socket } from 'socket.io-client';

declare module 'vue/types/vue' {
  interface Vue {
    $socket: Socket;
  }
}
