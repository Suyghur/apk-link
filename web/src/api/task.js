import request from '@/utils/request'

export function getTaskList(param) {
  return request({
    url: '/task/searchTasks',
    method: 'get',
    params: param
  })
}

export function createTask(data) {
  return request({
    url: '/task/createTask',
    method: 'post',
    data
  })
}

export function modifyTask(data) {
  return request({
    url: '/task/modifyTask',
    method: 'put',
    data
  })
}

export function getTaskInfo(param) {
  return request({
    url: '/task/getTaskInfo',
    method: 'get',
    params: param
  })
}

export function modifyStatus(data) {
  return request({
    url: '/task/modifyStatus',
    method: 'put',
    data
  })
}

export function deleteTask(data) {
  return request({
    url: '/task/deleteTask',
    method: 'delete',
    data
  })
}
