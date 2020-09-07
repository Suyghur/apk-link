/*
@Time : 2020/9/5
@Author : #Suyghur,
@File : sys_game
*/

package service

import (
	"errors"
	"server/global"
	"server/model"
)

func CreatePlugin(bean model.SysPlugin) (err error) {
	var plugin model.SysPlugin
	//根据版本去判断是否存在
	isExit := !global.GVA_DB.Where("plugin_name = ?", bean.PluginName).First(&plugin).RecordNotFound()
	//isExit为true表明读到了，不能新建
	if isExit {
		return errors.New("该插件已存在")
	} else {
		err = global.GVA_DB.Create(&bean).Error
	}
	return err
}

func ListPlugin() (err error, games []string) {
	err = global.GVA_DB.Model(&model.SysGame{}).Pluck("game_group", &games).Error
	return err, games
}
