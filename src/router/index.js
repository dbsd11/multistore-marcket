import Vue from 'vue'
import Router from 'vue-router'
import HelloWorld from '@/components/HelloWorld'

Vue.use(Router)

const login = r => require.ensure([], () => r(require('@/page/login')), 'login');
const manage = r => require.ensure([], () => r(require('@/page/manage')), 'manage');
const home = r => require.ensure([], () => r(require('@/page/home')), 'home');

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
        }
      ]
    }
  ],
  strict: process.env.NODE_ENV !== 'production',
})
