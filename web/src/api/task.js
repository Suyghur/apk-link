import request from '@/utils/request'

export function getTaskList(data) {
  return request({
    url: '/task/searchTasks',
    method: 'post',
    data
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
    method: 'post',
    data
  })
}

export function getTaskInfo(data) {
  return request({
    url: '/task/getTaskInfo',
    method: 'post',
    data
  })
}

export function modifyStatus(data) {
  return request({
    url: '/task/modifyStatus',
    method: 'post',
    data
  })
}

export function deleteTask(data) {
  return request({
    url: '/task/deleteTask',
    method: 'post',
    data
  })
}
