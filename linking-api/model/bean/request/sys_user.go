/*
@Time : 2020/8/28
@Author : #Suyghur,
@File : sys_user
*/

package request

import uuid "github.com/satori/go.uuid"

type RegisterStruct struct {
	Username string `json:"username"`
	Password string `json:"password" gorm:"default:'abc123456'"`
	Nickname string `json:"nickname" gorm:"default:'HoulangUser'"`
}

type LoginStruct struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type SetUserAuth struct {
	UUID uuid.UUID `json:"uuid"`
}

type TestStruct struct {
	Username string `json:"username"`
	Password string `json:"password" gorm:"default:'abc123456'"`
	Nickname string `json:"nickname" gorm:"default:'HoulangUser'"`
}
