import Layout from '@/layout/index'

export const LinkingRouterMap = {
  path: '/link',
  component: Layout,
  redirect: '/welcome',
  children: [
    {
      path: 'index',
      name: 'Linking',
      component: () => import('@/views/link/index'),
      meta: { title: '游戏分发包管理', icon: 'el-icon-set-up' }
    }
  ]
}
