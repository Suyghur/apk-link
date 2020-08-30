/*
@Time : 2020/8/27
@Author : #Suyghur,
@File : sys_user
*/

package model

import (
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

type SysUser struct {
	gorm.Model
	UUID     uuid.UUID `json:"uuid" gorm:"comment:'用户UUID'"`
	Username string    `json:"username" gorm:"comment:'用户登录名'"`
	Password string    `json:"-"  gorm:"comment:'用户登录密码'"`
	Nickname string    `json:"nickname" gorm:"default:'系统用户';comment:'用户昵称'" `
}
