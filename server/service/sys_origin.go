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

func GetOrigins(gameSite string) (err error, origins []response.OriginsResponse) {
	err = global.GvaDb.Model(&model.SysOrigin{}).Select("game_file_name , game_orientation , apk_url").Where("game_site = ?", gameSite).Scan(&origins).Error
	return err, origins
}

func SearchOrigin(bean request.ReqSearchOriginBean) (err error, origins []model.SysOrigin, total int64) {
	limit := bean.PageSize
	offset := bean.PageSize * (bean.Page - 1)
	//创建db
	db := global.GvaDb.Model(&model.SysOrigin{})
	if bean.GameSite != "" {
		db = db.Where("game_site = ?", bean.GameSite)
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
	isExist := !errors.Is(global.GvaDb.Where("game_site = ?  AND game_file_name = ? AND game_version_code = ?", bean.GameSite, bean.GameFileName, bean.GameVersionCode).First(&model.SysOrigin{}).Error, gorm.ErrRecordNotFound)
	if isExist {
		err = global.GvaDb.Where("game_site = ? AND game_version_code = ?", bean.GameSite, bean.GameVersionCode).First(&model.SysOrigin{}).Updates(map[string]interface{}{
			"game_site":         bean.GameSite,
			"sdk_version":       bean.SdkVersion,
			"game_file_name":    bean.GameFileName,
			"game_version_code": bean.GameVersionCode,
			"game_version_name": bean.GameVersionName,
			"game_orientation":  bean.GameOrientation,
			"keystore_name":     bean.KeystoreName,
			"apk_url":           bean.ApkUrl,
		}).Error
	} else {
		return errors.New("该游戏没有可修改的母包")
	}
	return err
}

func CreateOrigin(bean model.SysOrigin) (err error) {
	//判断游戏母包是否存在
	isExist := !errors.Is(global.GvaDb.Where("game_site = ? AND game_file_name = ? AND game_version_code = ?", bean.GameSite, bean.GameFileName, bean.GameVersionCode).First(&model.SysOrigin{}).Error, gorm.ErrRecordNotFound)
	if isExist {
		return errors.New("该游戏已存在对应版本的母包")
	} else {
		err = global.GvaDb.Create(&bean).Error
	}
	return err
}
