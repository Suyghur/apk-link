/*
@Time : 2020/8/28
@Author : #Suyghur,
@File : sys_user
*/

package request

import uuid "github.com/satori/go.uuid"

type RegisterBean struct {
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password" gorm:"default:'abc123456'"`
	Nickname string `json:"nickname" form:"nickname" gorm:"default:'HoulangUser'"`
}

type LoginBean struct {
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
}

type SetUserAuth struct {
	UUID uuid.UUID `json:"uuid" form:"uuid"`
}

type UserInfoBean struct {
	Username string `json:"username" form:"username"`
}

type ChangePasswordBean struct {
	Username    string `json:"username" form:"username"`
	Password    string `json:"password" form:"password"`
	NewPassword string `json:"new_password" form:"new_password"`
}
