import Vue from 'vue';
import VueRouter, { RouteConfig } from 'vue-router';
import Home from '@/views/Home.vue';
import Owner from '@/views/Owner.vue';
import Borrower from '@/views/Borrower.vue';
import TestToken from '@/views/TestToken.vue';
import Lender from '@/views/Lender.vue';

Vue.use(VueRouter);

const routes: Array<RouteConfig> = [
  {
    path: '/',
    name: 'Home',
    component: Home,
  },
  {
    path: '/about',
    name: 'About',
    // route level code-splitting
    // this generates a separate chunk (about.[hash].js) for this route
    // which is lazy-loaded when the route is visited.
    component: () => import(/* webpackChunkName: "about" */ '../views/About.vue'),
  },
  {
    path: '/owner',
    name: 'Owner',
    component: Owner,
  },
  {
    path: '/borrower',
    name: 'Borrower',
    component: Borrower,
  },
  {
    path: '/lender',
    name: 'Lender',
    component: Lender,
  },
  {
    path: '/test-token',
    name: 'TestToken',
    component: TestToken,
  },
];

const router = new VueRouter({
  mode: 'history',
  base: process.env.BASE_URL,
  routes,
});

export default router;
