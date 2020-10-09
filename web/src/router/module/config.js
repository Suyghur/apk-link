import Layout from '@/layout/index'

export const ConfigRouterMap = {
  path: '/config',
  component: Layout,
  redirect: '/welcome',
  name: 'Config',
  meta: { title: '配置管理', icon: 'el-icon-set-up', breadcrumb: false },
  children: [
    {
      path: 'game',
      name: 'Game',
      component: () => import('@/views/config/game/index'),
      meta: { title: '游戏配置管理', icon: 'el-icon-star-off' }
    },
    {
      path: 'channel',
      name: 'Channel',
      component: () => import('@/views/config/channel'),
      meta: { title: '渠道配置管理', icon: 'el-icon-star-off' }
    },
    {
      path: 'plugin',
      name: 'Plugin',
      component: () => import('@/views/config/plugin'),
      meta: { title: '插件配置管理', icon: 'el-icon-star-off' }
    }
  ]
}
