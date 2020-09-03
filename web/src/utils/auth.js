import Cookies from 'js-cookie'

const TokenKey = 'linking_token'

export function getToken() {
  return Cookies.get(TokenKey)
}

export function setToken(token) {
  return Cookies.set(TokenKey, token)
}

export function removeToken() {
  return Cookies.remove(TokenKey)
}

const UsernameKey = 'linking_username'

export function getUsername() {
  return Cookies.get(UsernameKey)
}

export function setUsername(token) {
  return Cookies.set(UsernameKey, token)
}

export function removeUsername() {
  return Cookies.remove(UsernameKey)
}
