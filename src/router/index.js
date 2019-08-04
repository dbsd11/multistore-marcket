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
const explain = r => require.ensure([], () => r(require('@/page/explain')), 'explain');

export default new Router({
  routes: [
    {
      path: '/hello',
      name: 'HelloWorld',
      component: HelloWorld
    },
    {
      path: '/',
      component: manage
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
          meta: ["首页"],
        },
        {
          path: '/goodList',
          component: goodList,
          meta: ["商品管理", "商品列表"],
        },
        {
          path: '/marcketList',
          component: marcketList,
          meta: ["商品管理", "集市列表"],
        },
        {
          path: '/tradeList',
          component: tradeList,
          meta: ["交易", "交易记录"],
        },
        {
          path: '/initiateTrade',
          component: initiateTrade,
          meta: ["发起交易"],
        },{
          path: '/explain',
          component: explain,
          meta: ["说明"],
        }
      ]
    }
  ],
  strict: process.env.NODE_ENV !== 'production',
})
