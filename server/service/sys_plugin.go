/*
@Time : 2020/9/5
@Author : #Suyghur,
@File : sys_game
*/

package service

import (
	"errors"
	"gorm.io/gorm"
	"server/global"
	"server/model"
	"server/model/bean/request"
)

func CreatePlugin(bean model.SysPlugin) (err error) {
	//根据版本去判断是否存在
	isExit := !errors.Is(global.GvaDb.Where("plugin_name = ?", bean.PluginName).First(&model.SysPlugin{}).Error, gorm.ErrRecordNotFound)
	//isExit := !global.GvaDb.Where("plugin_name = ?", bean.PluginName).First(&plugin).RecordNotFound()
	//isExit为true表明读到了，不能新建
	if isExit {
		return errors.New("该插件已存在")
	} else {
		err = global.GvaDb.Create(&bean).Error
	}
	return err
}

func SearchPlugin(bean request.ReqPluginListBean) (err error, list interface{}, total int64) {
	limit := bean.PageSize
	offset := bean.PageSize * (bean.Page - 1)
	//创建db
	db := global.GvaDb.Model(&model.SysPlugin{})
	var plugins []model.SysPlugin
	if bean.PluginAlias != "" {
		db = db.Where("plugin_alias = ?", bean.PluginAlias)
	}
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&plugins).Error
	return err, plugins, total
}
