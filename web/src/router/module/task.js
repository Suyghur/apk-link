import Layout from '@/layout/index'

export const TaskRouterMap = {
  path: '/task',
  component: Layout,
  redirect: '/task/index',
  children: [
    {
      path: 'index',
      name: 'TaskInfos',
      component: () => import('@/views/task'),
      meta: { title: '打包任务管理', icon: 'el-icon-set-up' }
    },
    {
      path: 'create',
      name: 'TaskCreate',
      component: () => import('@/views/task/TaskCreate'),
      meta: { title: '创建任务', noCache: true, activeMenu: '/task/create' },
      hidden: true
    },
    {
      path: 'edit/:id(\\d+)',
      component: () => import('@/views/task/TaskEdit'),
      name: 'TaskEditor',
      meta: { title: '编辑任务', noCache: true, activeMenu: '/task/edit/:id(\\d+)' },
      hidden: true
    },
    {
      path: 'detail/:id(\\d+)',
      component: () => import('@/views/task/TaskDetail'),
      name: 'TaskDetail',
      meta: { title: '任务详情', noCache: true, activeMenu: '/task/detail/:id(\\d+)' },
      hidden: true
    }
  ]
}
