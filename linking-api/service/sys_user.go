/*
@Time : 2020/8/28
@Author : #Suyghur,
@File : sys_user
*/

package service

import (
	"linking-api/global"
	"linking-api/model"
	"linking-api/utils/crypto"
)

func Login(u *model.SysUser) (err error, userInter *model.SysUser) {
	var user model.SysUser
	u.Password = crypto.MD5([]byte(u.Password))
	err = global.GVA_DB.Where("username = ? AND password = ?", u.Username, u.Password).First(&user).Error
	return err, &user
}
