import request from '@/utils/request'

export function getTaskList(params) {
  return request({
    url: '/get_task_list',
    method: 'get',
    params
  })
}
