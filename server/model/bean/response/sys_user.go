/*
@Time : 2020/8/28
@Author : #Suyghur,
@File : sys_user
*/

package response

import "server/model"

type SysUserResponse struct {
	User model.SysUser `json:"user"`
}

type LoginResponse struct {
	User      model.SysUser `json:"user"`
	Token     string        `json:"token"`
	ExpiresAt int64         `json:"expires_at"`
}
