/*
@Time : 9/2/2020
@Author : #Suyghur,
@File : sys_script
*/

package service

import (
	"errors"
	"linking-api/global"
	"linking-api/model"
	"linking-api/model/bean/request"
)

func ListScript(bean request.ReqListScriptBean) (err error, list interface{}, total int) {
	limit := bean.PageSize
	offset := bean.PageSize * (bean.Page - 1)
	//创建db
	db := global.GVA_DB.Model(&model.SysScript{})
	var scripts []model.SysScript
	if bean.GameGroup != "" {
		db = db.Where("game_group = ?", bean.GameGroup)
	}
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&scripts).Error
	return err, scripts, total

}

//func SearchScript(gameGroup string) (err error, script model.SysScript) {
//	err = global.GVA_DB.Where("game_group = ?", gameGroup).First(&script).Error
//	return
//}

func ModifyScript(bean model.SysScript) (err error) {
	var script model.SysScript
	//判断游戏脚本是否存在
	isExit := !global.GVA_DB.Where("game_group = ?", bean.GameGroup).First(&script).RecordNotFound()
	//isExit为true表明读到了，不能新建
	if isExit {
		err = global.GVA_DB.Where("game_group = ?", bean.GameGroup).First(&script).Updates(map[string]interface{}{
			"game_group":       bean.GameGroup,
			"script_file_name": bean.ScriptFileName,
			"script_file_url":  bean.ScriptFileUrl,
		}).Error
	} else {
		return errors.New("该游戏没有可修改的脚本")
	}
	return err
}

func CreateScript(bean model.SysScript) (err error) {
	var script model.SysScript
	//判断游戏脚本是否存在
	isExit := !global.GVA_DB.Where("game_group = ?", bean.GameGroup).First(&script).RecordNotFound()
	//isExit为true表明读到了，不能新建
	if isExit {
		return errors.New("该游戏已存在脚本文件")
	} else {
		err = global.GVA_DB.Create(&bean).Error
	}
	return err
}
