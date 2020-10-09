import Layout from '@/layout/index'

export const KeystoreRouterMap = {
  path: '/keystore',
  component: Layout,
  redirect: '/welcome',
  children: [
    {
      path: 'index',
      name: 'Keystore',
      component: () => import('@/views/keystore/index'),
      meta: { title: '签名文件管理', icon: 'el-icon-set-up' }
    }
  ]
}
