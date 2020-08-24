import Layout from '@/layout/index'

export const GameRouterMap = {
  path: '/game',
  component: Layout,
  redirect: '/welcome',
  name: 'Bag',
  meta: { title: '游戏管理', icon: 'el-icon-s-help', breadcrumb: false },
  children: [
    {
      path: 'origin',
      name: 'OriginBag',
      component: () => import('@/views/game/OriginBag'),
      meta: { title: '游戏母包管理', icon: 'table' }
    },
    {
      path: 'linking',
      name: 'LinkingBag',
      component: () => import('@/views/game/LinkingBag'),
      meta: { title: '游戏分发包管理', icon: 'table' }
    },
    {
      path: 'keystore',
      name: 'Keystore',
      component: () => import('@/views/game/Keystore'),
      meta: { title: 'Keystore签名文件管理', icon: 'table' }
    }
  ]
}
