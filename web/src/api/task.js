import request from '@/utils/request'

export function getTaskList(params) {
  return request({
    url: '/get_task_list',
    method: 'get',
    params
  })
}

export function getSelectOptions(params) {
  return request({
    url: '/get_select_options',
    method: 'get',
    params
  })
}
