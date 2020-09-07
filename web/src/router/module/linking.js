import Layout from '@/layout/index'

export const LinkingRouterMap = {
  path: '/linking',
  component: Layout,
  redirect: '/welcome',
  children: [
    {
      path: 'index',
      name: 'Linking',
      component: () => import('@/views/linking/index'),
      meta: { title: '游戏分发包管理', icon: 'table' }
    }
  ]
}
