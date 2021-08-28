import Vue from 'vue';
import Vuex from 'vuex';

Vue.use(Vuex);

export const store = new Vuex.Store({// eslint-disable-line
  // as you can see, there is no module declaration right here,
  // it because I am using dynamic module.
  // you can config a module to be dynamic by set param `dynamic` = true,
  // the param located inside anotation `@Module`.
});
