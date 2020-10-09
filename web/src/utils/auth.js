import Cookies from 'js-cookie'

const TokenKey = 'Apk-Link-Token'

export function getToken() {
  return Cookies.get(TokenKey)
}

export function setToken(token) {
  return Cookies.set(TokenKey, token)
}

export function removeToken() {
  return Cookies.remove(TokenKey)
}

const UsernameKey = 'link_username'

export function getUsername() {
  return Cookies.get(UsernameKey)
}

export function setUsername(token) {
  return Cookies.set(UsernameKey, token)
}

export function removeUsername() {
  return Cookies.remove(UsernameKey)
}
