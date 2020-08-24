import Layout from '@/layout/index'

export const ScriptRouterMap = {
  path: '/script',
  component: Layout,
  redirect: '/welcome',
  name: 'Script',
  meta: { title: '脚本管理', icon: 'el-icon-s-help', breadcrumb: false },
  children: [
    {
      path: 'fuse',
      name: 'FuseScript',
      component: () => import('@/views/script/FuseScript'),
      meta: { title: '聚合SDK脚本管理', icon: 'table' }
    },
    {
      path: 'channel',
      name: 'ChannelScript',
      component: () => import('@/views/script/ChannelScript'),
      meta: { title: '渠道SDK脚本管理', icon: 'table' }
    },
    {
      path: 'plugin',
      name: 'PluginScript',
      component: () => import('@/views/script/PluginScript'),
      meta: { title: '插件SDK脚本管理', icon: 'table' }
    },
    {
      path: 'game',
      name: 'GameScript',
      component: () => import('@/views/script/GameScript'),
      meta: { title: '游戏脚本管理', icon: 'table' }
    }
  ]
}
