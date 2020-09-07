import Layout from '@/layout/index'

export const ScriptRouterMap = {
  path: '/script',
  component: Layout,
  redirect: '/welcome',
  children: [
    {
      path: 'index',
      name: 'GameScript',
      component: () => import('@/views/script'),
      meta: { title: '游戏脚本管理', icon: 'table' }
    }
  ]
}
