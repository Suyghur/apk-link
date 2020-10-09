import Layout from '@/layout/index'

export const SdkRouterMap = {
  path: '/sdk',
  component: Layout,
  redirect: '/welcome',
  name: 'Sdk',
  meta: { title: 'SDK管理', icon: 'el-icon-set-up', breadcrumb: false },
  children: [
    {
      path: 'fuse',
      name: 'FuseSdk',
      component: () => import('@/views/sdk/fuse'),
      meta: { title: '聚合SDK管理', icon: 'el-icon-star-off' }
    },
    {
      path: 'channel',
      name: 'ChannelSdk',
      component: () => import('@/views/sdk/channel'),
      meta: { title: '渠道SDK管理', icon: 'el-icon-star-off' }
    },
    {
      path: 'plugin',
      name: 'PluginSdk',
      component: () => import('@/views/sdk/plugin'),
      meta: { title: '插件SDK管理', icon: 'el-icon-star-off' }
    }
  ]
}
