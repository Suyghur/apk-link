/*
@Time : 2020/9/4
@Author : #Suyghur,
@File : sys_options
*/

package service

import (
	"server/global"
	"server/model"
	"server/model/bean/response"
)

func GetOptions() (err error, options response.OptionsResponse) {
	tx := global.GvaDb.Begin()
	err = tx.Model(&model.SysGid{}).Pluck("gid", &options.GameOptions).Error
	if err != nil {
		tx.Rollback()
		return err, options
	}
	err = tx.Model(&model.SysChannel{}).Pluck("channel_name", &options.ChannelOptions).Error
	if err != nil {
		tx.Rollback()
		return err, options
	}
	err = tx.Model(&model.SysPlugin{}).Pluck("plugin_name", &options.PluginOptions).Error
	if err != nil {
		tx.Rollback()
		return err, options
	}
	tx.Commit()
	return err, options
}
