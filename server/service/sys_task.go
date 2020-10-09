/*
@Time : 2020/8/31
@Author : #Suyghur,
@File : sys_task
*/

package service

import (
	"errors"
	"server/global"
	"server/model"
	"server/model/bean/request"
)

func SearchTasks(bean request.ReqListTaskBean) (err error, list interface{}, total int) {
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
	if bean.StatusCode >= 0 {
		db = db.Where("status_code = ?", bean.StatusCode)
	}
	err = db.Count(&total).Error
	err = db.Order("task_id desc").Limit(limit).Offset(offset).Find(&tasks).Error
	return err, tasks, total
}

func GetTaskInfo(taskId uint) (err error, taskInfo model.SysTask) {

	isExit := !global.GVA_DB.Where("task_id = ?", taskId).First(&taskInfo).RecordNotFound()
	if isExit {
		err = global.GVA_DB.Where("task_id = ?", taskId).First(&taskInfo).Error
	} else {
		return errors.New("该任务不存在"), taskInfo
	}
	return err, taskInfo
}

func DeleteTask(taskId uint) (err error) {
	var taskInfo model.SysTask
	isExit := !global.GVA_DB.Where("task_id = ?", taskId).First(&taskInfo).RecordNotFound()
	if isExit {
		err = global.GVA_DB.Where("task_id = ?", taskId).Unscoped().Delete(&taskInfo).Error
	} else {
		return errors.New("该任务不存在")
	}
	return err
}

func CreateTask(bean model.SysTask) (err error) {
	var task model.SysTask
	//判断游戏脚本是否存在
	isExit := !global.GVA_DB.Where("task_id = ?", bean.TaskId).First(&task).RecordNotFound()
	//isExit为true表明读到了，不能新建
	if isExit {
		return errors.New("该任务已存在")
	} else {
		tx := global.GVA_DB.Begin()
		if bean.PluginName == "" {
			bean.IsPluginSdk = 0
		} else {
			bean.IsPluginSdk = 1
		}
		type scriptBean struct {
			ScriptFileName string
			ScriptFileUrl  string
		}
		var script scriptBean
		if err = global.GVA_DB.Model(&model.SysScript{}).Select("script_file_name , script_file_url").Where("game_group = ?", bean.GameGroup).Scan(&script).Error; err != nil {
			bean.ScriptFileName = ""
			bean.ScriptFileUrl = ""
		} else {
			bean.ScriptFileName = script.ScriptFileName
			bean.ScriptFileUrl = script.ScriptFileUrl
		}

		if err = tx.Create(&bean).Error; err != nil {
			tx.Rollback()
			return err
		}
		tx.Commit()
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
			"status_code":             bean.StatusCode,
			"status_msg":              bean.StatusMsg,
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

func ModifyTaskStatus(taskId uint, statusCode int, statusMsg string) (err error) {
	var task model.SysTask
	isExit := !global.GVA_DB.Where("task_id = ?", taskId).First(&task).RecordNotFound()
	//isExit为true表明读到了
	if isExit {
		err = global.GVA_DB.Where("task_id = ?", taskId).First(&task).Updates(map[string]interface{}{
			"status_code": statusCode,
			"status_msg":  statusMsg,
		}).Error
	} else {
		return errors.New("没有可修改的任务")
	}
	return err
}
