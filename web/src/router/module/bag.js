import Layout from '@/layout/index'

export const BagRouterMap = {
  path: '/bag',
  component: Layout,
  redirect: '/welcome',
  name: 'Bag',
  meta: { title: '游戏包管理', icon: 'el-icon-s-help', breadcrumb: false },
  children: [
    {
      path: 'origin',
      name: 'OriginBag',
      component: () => import('@/views/bag/OriginBag'),
      meta: { title: '游戏母包管理', icon: 'table' }
    },
    {
      path: 'linking',
      name: 'LinkingBag',
      component: () => import('@/views/bag/LinkingBag'),
      meta: { title: '游戏分发包管理', icon: 'table' }
    },
    {
      path: 'keystore',
      name: 'KeyStore',
      component: () => import('@/views/bag/KeyStore'),
      meta: { title: 'KeyStore签名文件管理', icon: 'table' }
    }
  ]
}
