import request from '@/utils/request'

export function getList(params) {
  return request({
    url: '/table/list',
    method: 'get',
    params
  })
}

export function getTasks(params) {
  return request({
    url: '/table/tasks',
    method: 'get',
    params
  })
}
