/*
@Time : 2020/8/31
@Author : #Suyghur,
@File : sys_task
*/

package service

import (
	"errors"
	"linking-api/global"
	"linking-api/model"
	"linking-api/model/bean/request"
)

func ListTask(bean request.ReqListTaskBean) (err error, list interface{}, total int) {
	limit := bean.PageSize
	offset := bean.PageSize * (bean.Page - 1)
	//创建db
	db := global.GVA_DB.Model(&model.SysTask{})
	var tasks []model.SysTask
	if bean.TaskId > 0 {
		db = db.Where("task_id = ?", bean.TaskId)
	}
	if bean.GameGroup != "" {
		db = db.Where("game_group = ?", bean.GameGroup)
	}
	if bean.ChannelName != "" {
		db = db.Where("channel_name = ?", bean.ChannelName)
	}
	if bean.PluginName != "" {
		db = db.Where("plugin_name = ?", bean.PluginName)
	}
	if bean.Status >= 0 {
		db = db.Where("status = ?", bean.Status)
	}
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&tasks).Error
	return err, tasks, total
}

func CreateTask(bean model.SysTask) (err error) {
	var task model.SysTask
	//判断游戏脚本是否存在
	isExit := !global.GVA_DB.Where("task_id = ?", bean.TaskId).First(&task).RecordNotFound()
	//isExit为true表明读到了，不能新建
	if isExit {
		return errors.New("该任务已存在")
	} else {
		err = global.GVA_DB.Create(&bean).Error
	}
	return err
}

func ModifyTask(bean model.SysTask) (err error) {
	var task model.SysTask
	//判断游戏脚本是否存在
	isExit := !global.GVA_DB.Where("task_id = ?", bean.TaskId).First(&task).RecordNotFound()
	//isExit为true表明读到了，不能新建
	if isExit {
		err = global.GVA_DB.Where("task_id = ?", bean.TaskId).First(&task).Updates(map[string]interface{}{
			"is_fuse_sdk":             bean.IsFuseSdk,
			"is_white_bag":            bean.IsWhiteBag,
			"is_plugin_sdk":           bean.IsPluginSdk,
			"game_group":              bean.GameGroup,
			"gid":                     bean.Gid,
			"aids":                    bean.Aids,
			"channel_params":          bean.ChannelParams,
			"plugin_params":           bean.PluginParams,
			"game_package_name":       bean.GamePackageName,
			"game_name":               bean.GameName,
			"game_version_code":       bean.GameVersionCode,
			"game_version_name":       bean.GameVersionName,
			"game_orientation":        bean.GameOrientation,
			"game_icon_url":           bean.GameIconUrl,
			"game_logo_url":           bean.GameLogoUrl,
			"game_splash_url":         bean.GameSplashUrl,
			"game_login_bg_url":       bean.GameLoginBgUrl,
			"status":                  bean.Status,
			"game_file_name":          bean.GameFileName,
			"game_file_url":           bean.GameFileUrl,
			"fuse_sdk_file_name":      bean.FuseSdkFileName,
			"fuse_sdk_version":        bean.FuseSdkVersion,
			"fuse_sdk_file_url":       bean.FuseSdkFileUrl,
			"channel_name":            bean.ChannelName,
			"channel_sdk_file_name":   bean.ChannelSdkFileName,
			"channel_sdk_version":     bean.ChannelSdkVersion,
			"channel_sdk_file_url":    bean.ChannelSdkFileUrl,
			"plugin_name":             bean.PluginName,
			"plugin_sdk_file_name":    bean.PluginSdkFileName,
			"plugin_sdk_version":      bean.PluginSdkVersion,
			"plugin_sdk_file_url":     bean.PluginSdkFileUrl,
			"script_file_name":        bean.ScriptFileName,
			"script_file_url":         bean.ScriptFileUrl,
			"keystore_name":           bean.KeystoreName,
			"keystore_password":       bean.KeystorePassword,
			"keystore_alias":          bean.KeystoreAlias,
			"keystore_alias_password": bean.KeystoreAliasPassword,
			"ext":                     bean.Ext,
		}).Error
	} else {
		return errors.New("没有可修改的任务")
	}
	return err
}

func ModifyTaskStatus(taskId uint, status int) (err error) {
	var task model.SysTask
	//判断游戏脚本是否存在
	isExit := !global.GVA_DB.Where("task_id = ?", taskId).First(&task).RecordNotFound()
	//isExit为true表明读到了，不能新建
	if isExit {
		err = global.GVA_DB.Where("task_id = ?", taskId).First(&task).Updates(map[string]interface{}{
			"status": status,
		}).Error
	} else {
		return errors.New("没有可修改的任务")
	}
	return err
}
