/*
@Time : 2020/9/3
@Author : #Suyghur,
@File : sys_origin
*/

package service

import (
	"errors"
	"gorm.io/gorm"
	"server/global"
	"server/model"
	"server/model/bean/request"
	"server/model/bean/response"
)

func GetOrigins(gameGroup string) (err error, origins []response.OriginsResponse) {
	err = global.GvaDb.Model(&model.SysOrigin{}).Select("game_file_name , is_fuse_sdk , game_orientation , apk_url").Where("game_group = ?", gameGroup).Scan(&origins).Error
	return err, origins
}

func SearchOrigin(bean request.ReqSearchOriginBean) (err error, list interface{}, total int64) {
	limit := bean.PageSize
	offset := bean.PageSize * (bean.Page - 1)
	//创建db
	db := global.GvaDb.Model(&model.SysOrigin{})
	var origins []model.SysOrigin
	if bean.GameGroup != "" {
		db = db.Where("game_group = ?", bean.GameGroup)
	}
	if bean.SdkVersion != "" {
		db = db.Where("sdk_version = ?", bean.SdkVersion)
	}
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&origins).Error
	return err, origins, total
}

func ModifyOrigin(bean model.SysOrigin) (err error) {
	//判断游戏母包是否存在
	isExit := !errors.Is(global.GvaDb.Where("game_group = ? AND game_file_name = ?", bean.GameGroup, bean.GameFileName).First(&model.SysOrigin{}).Error, gorm.ErrRecordNotFound)
	//isExit := !global.GvaDb.Where("game_group = ? AND game_file_name = ?", bean.GameGroup, bean.GameFileName).First(&origin).RecordNotFound()
	//isExit为true表明读到了，不能新建
	if isExit {
		err = global.GvaDb.Where("game_group = ? AND game_file_name = ?", bean.GameGroup, bean.GameFileName).First(&model.SysOrigin{}).Updates(map[string]interface{}{
			"game_group":        bean.GameGroup,
			"gid":               bean.Gid,
			"is_fuse_sdk":       bean.IsFuseSdk,
			"sdk_version":       bean.SdkVersion,
			"game_file_name":    bean.GameFileName,
			"game_version_code": bean.GameVersionCode,
			"game_version_name": bean.GameVersionName,
			"game_orientation":  bean.GameOrientation,
			"keystore_name":     bean.KeystoreName,
			"icon_url":          bean.IconUrl,
			"apk_url":           bean.ApkUrl,
		}).Error
	} else {
		return errors.New("该游戏没有可修改的母包")
	}
	return err
}

func CreateOrigin(bean model.SysOrigin) (err error) {
	//判断游戏母包是否存在
	isExit := !errors.Is(global.GvaDb.Where("game_group = ? AND game_file_name = ?", bean.GameGroup, bean.GameFileName).First(&model.SysOrigin{}).Error, gorm.ErrRecordNotFound)
	//isExit := !global.GvaDb.Where("game_group = ? AND game_file_name = ?", bean.GameGroup, bean.GameFileName).First(&origin).RecordNotFound()
	//isExit为true表明读到了，不能新建
	if isExit {
		return errors.New("该游戏已存在对应版本的母包")
	} else {
		err = global.GvaDb.Create(&bean).Error
	}
	return err
}
