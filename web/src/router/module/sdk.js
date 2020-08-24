import Layout from '@/layout/index'

export const SdkRouterMap = {
  path: '/sdk',
  component: Layout,
  redirect: '/welcome',
  name: 'Sdk',
  meta: { title: 'SDK管理', icon: 'el-icon-s-help', breadcrumb: false },
  children: [
    {
      path: 'fuse',
      name: 'FuseSdk',
      component: () => import('@/views/sdk/FuseSdk'),
      meta: { title: '聚合SDK管理', icon: 'table' }
    },
    {
      path: 'channel',
      name: 'ChannelSdk',
      component: () => import('@/views/sdk/ChannelSdk'),
      meta: { title: '渠道SDK管理', icon: 'table' }
    },
    {
      path: 'plugin',
      name: 'PluginSdk',
      component: () => import('@/views/sdk/PluginSdk'),
      meta: { title: '插件SDK管理', icon: 'table' }
    }
  ]
}
