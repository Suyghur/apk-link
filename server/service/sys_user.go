/*
@Time : 2020/8/28
@Author : #Suyghur,
@File : sys_user
*/

package service

import (
	"errors"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
	"server/global"
	"server/model"
	"server/utils/crypto"
)

func Register(u model.SysUser) (err error, userInter model.SysUser) {
	//判断用户名是否注册
	notRegister := errors.Is(global.GvaDb.Where("username = ?", u.Username).First(&model.SysUser{}).Error, gorm.ErrRecordNotFound)
	//notRegister := global.GvaDb.Where("username = ?", u.Username).First(&user).RecordNotFound()
	//notRegister为false表明读取到了，不能注册
	if !notRegister {
		return errors.New("该用户已注册"), userInter
	} else {
		//否则附加uuid，密码md5入库
		u.Password = crypto.MD5([]byte(u.Password))
		u.UUID = uuid.NewV4()
		err = global.GvaDb.Create(&u).Error
	}
	return err, u
}

func Login(u *model.SysUser) (err error, userInter *model.SysUser) {
	var user model.SysUser
	u.Password = crypto.MD5([]byte(u.Password))
	err = global.GvaDb.Where("username = ? AND password = ?", u.Username, u.Password).First(&user).Error
	return err, &user
}

func Info(username string) (err error, userInter *model.SysUser) {
	var user model.SysUser
	err = global.GvaDb.Where("username = ?", username).First(&user).Error
	return err, &user
}

func ChangePassword(u *model.SysUser, newPassword string) (err error, userInter *model.SysUser) {
	var user model.SysUser
	u.Password = crypto.MD5([]byte(u.Password))
	err = global.GvaDb.Where("username = ? AND password = ?", u.Username, u.Password).First(&user).Update("password", crypto.MD5([]byte(newPassword))).Error
	return err, u
}
