import Layout from '@/layout/index'

export const OriginRouterMap = {
  path: '/origin',
  component: Layout,
  redirect: '/welcome',
  children: [
    {
      path: 'index',
      name: 'Origin',
      component: () => import('@/views/origin/index'),
      meta: { title: '游戏母包管理', icon: 'el-icon-set-up' }
    }
  ]
}
