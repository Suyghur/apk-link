/*
@Time : 2020/8/27
@Author : #Suyghur,
@File : sys_jwt_blacklist
*/

package model

import "github.com/jinzhu/gorm"

type JwtBlacklist struct {
	gorm.Model
	Jwt string `gorm:"type:text;comment:'jwt'"`
}
