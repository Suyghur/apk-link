/*
@Time : 9/2/2020
@Author : #Suyghur,
@File : sys_script
*/

package service

import (
	"errors"
	"gorm.io/gorm"
	"server/global"
	"server/model"
	"server/model/bean/request"
)

func SearchScript(bean request.ReqListScriptBean) (err error, list interface{}, total int64) {
	limit := bean.PageSize
	offset := bean.PageSize * (bean.Page - 1)
	//创建db
	db := global.GvaDb.Model(&model.SysScript{})
	var scripts []model.SysScript
	if bean.GameSite != "" {
		db = db.Where("game_site = ?", bean.GameSite)
	}
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&scripts).Error
	return err, scripts, total

}

func ModifyScript(bean model.SysScript) (err error) {
	//判断游戏脚本是否存在
	isExist := !errors.Is(global.GvaDb.Where("game_site = ?", bean.GameSite).First(&model.SysScript{}).Error, gorm.ErrRecordNotFound)
	if isExist {
		err = global.GvaDb.Where("game_site = ?", bean.GameSite).First(&model.SysScript{}).Updates(map[string]interface{}{
			"game_site":        bean.GameSite,
			"script_file_name": bean.ScriptFileName,
			"script_file_url":  bean.ScriptFileUrl,
		}).Error
	} else {
		return errors.New("该游戏没有可修改的脚本")
	}
	return err
}

func CreateScript(bean model.SysScript) (err error) {
	//判断游戏脚本是否存在
	isExist := !errors.Is(global.GvaDb.Where("game_site = ?", bean.GameSite).First(&model.SysScript{}).Error, gorm.ErrRecordNotFound)
	if isExist {
		return errors.New("该游戏已存在脚本文件")
	} else {
		err = global.GvaDb.Create(&bean).Error
	}
	return err
}
