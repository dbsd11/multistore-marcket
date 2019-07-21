import Vue from 'vue'
import Router from 'vue-router'
import HelloWorld from '@/components/HelloWorld'

Vue.use(Router)

const login = r => require.ensure([], () => r(require('@/page/login')), 'login');
const manage = r => require.ensure([], () => r(require('@/page/manage')), 'manage');
const home = r => require.ensure([], () => r(require('@/page/home')), 'home');
const goodList = r => require.ensure([], () => r(require('@/page/goodList')), 'goodList');
const marcketList = r => require.ensure([], () => r(require('@/page/marcketList')), 'marcketList');
const tradeList = r => require.ensure([], () => r(require('@/page/tradeList')), 'tradeList');
const initiateTrade = r => require.ensure([], () => r(require('@/page/initiateTrade')), 'initiateTrade');

export default new Router({
  routes: [
    {
      path: '/hello',
      name: 'HelloWorld',
      component: HelloWorld
    },
    {
      path: '/',
      component: login
    },
    {
      path: '/index',
      component: manage
    },
    {
      path: '/manage',
      component: manage,
      children:[
        {
          path: '',
          component: home,
          meta: [],
        },
        {
          path: '/goodList',
          component: goodList,
          meta: [],
        },
        {
          path: '/marcketList',
          component: marcketList,
          meta: [],
        },
        {
          path: '/tradeList',
          component: tradeList,
          meta: [],
        },
        {
          path: '/initiateTrade',
          component: initiateTrade,
          meta: [],
        }
      ]
    }
  ],
  strict: process.env.NODE_ENV !== 'production',
})
