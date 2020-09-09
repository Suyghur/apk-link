/*
@Time : 2020/8/31
@Author : #Suyghur,
@File : sys_sdk
*/

package service

import (
	"errors"
	"server/global"
	"server/model"
	"server/model/bean/request"
	"server/model/bean/response"
)

func GetFuseSdks() (err error, sdks []response.SdksResponse) {
	err = global.GVA_DB.Model(&model.SysFuseSdk{}).Select("sdk_version , sdk_file_name , sdk_file_url").Scan(&sdks).Error
	return err, sdks
}

func SearchFuseSdk(bean request.ReqListFuseSdkBean) (err error, list interface{}, total int) {
	limit := bean.PageSize
	offset := bean.PageSize * (bean.Page - 1)
	//创建db
	db := global.GVA_DB.Model(&model.SysFuseSdk{})
	var fuseSdks []model.SysFuseSdk
	if bean.SdkVersion != "" {
		db = db.Where("sdk_version = ?", bean.SdkVersion)
	}
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&fuseSdks).Error
	return err, fuseSdks, total
}

func CreateFuseSdk(bean model.SysFuseSdk) (err error) {
	var fuseSdk model.SysFuseSdk
	//根据版本去判断是否存在
	isExit := !global.GVA_DB.Where("sdk_version = ?", bean.SdkVersion).First(&fuseSdk).RecordNotFound()
	//isExit为true表明读到了，不能新建
	if isExit {
		return errors.New("该版本聚合SDK已存在")
	} else {
		err = global.GVA_DB.Create(&bean).Error
	}
	return err
}

func ModifyFuseSdk(bean model.SysFuseSdk) (err error) {
	var fuseSdk model.SysFuseSdk
	//根据版本去判断是否存在
	isExit := !global.GVA_DB.Where("sdk_version = ?", bean.SdkVersion).First(&fuseSdk).RecordNotFound()
	//isExit为true表明读到了，不能新建
	if isExit {
		err = global.GVA_DB.Where("sdk_version = ?", bean.SdkVersion).First(&fuseSdk).Updates(map[string]interface{}{
			"sdk_name":      bean.SdkName,
			"sdk_version":   bean.SdkVersion,
			"sdk_file_name": bean.SdkFileName,
			"sdk_file_url":  bean.SdkFileUrl,
		}).Error
	} else {
		return errors.New("该版本聚合SDK不存在")
	}
	return err
}

func GetChannelSdks(channelName string) (err error, sdks []response.SdksResponse) {
	err = global.GVA_DB.Model(&model.SysChannelSdk{}).Select("sdk_version , sdk_file_name , sdk_file_url").Where("channel_name = ?", channelName).Scan(&sdks).Error
	return err, sdks
}

func SearchChannelSdk(bean request.ReqListChannelSdkBean) (err error, list interface{}, total int) {
	limit := bean.PageSize
	offset := bean.PageSize * (bean.Page - 1)
	//创建db
	db := global.GVA_DB.Model(&model.SysChannelSdk{})
	var channelSdks []model.SysChannelSdk
	if bean.ChannelName != "" {
		db = db.Where("channel_name = ?", bean.ChannelName)
	}
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&channelSdks).Error
	return err, channelSdks, total
}

func CreateChannelSdk(bean model.SysChannelSdk) (err error) {
	var channelSdk model.SysChannelSdk
	//根据版本去判断是否存在
	isExit := !global.GVA_DB.Where("channel_name = ? AND sdk_version = ?", bean.ChannelName, bean.SdkVersion).First(&channelSdk).RecordNotFound()
	//isExit为true表明读到了，不能新建
	if isExit {
		return errors.New("该版本渠道SDK已存在")
	} else {
		err = global.GVA_DB.Create(&bean).Error
	}
	return err
}

func ModifyChannelSdk(bean model.SysChannelSdk) (err error) {
	var channelSdk model.SysChannelSdk
	//根据版本去判断是否存在
	isExit := !global.GVA_DB.Where("channel_name = ? AND sdk_version = ?", bean.ChannelName, bean.SdkVersion).First(&channelSdk).RecordNotFound()
	//isExit为true表明读到了，不能新建
	if isExit {
		err = global.GVA_DB.Where("channel_name = ? AND sdk_version = ?", bean.ChannelName, bean.SdkVersion).First(&channelSdk).Updates(map[string]interface{}{
			"channel_name":  bean.ChannelName,
			"sdk_name":      bean.SdkName,
			"sdk_version":   bean.SdkVersion,
			"sdk_file_name": bean.SdkFileName,
			"sdk_file_url":  bean.SdkFileUrl,
		}).Error
	} else {
		return errors.New("该版本插件SDK不存在")
	}
	return err
}

func GetPluginSdks(pluginName string) (err error, sdks []response.SdksResponse) {
	global.GVA_LOG.Debug(pluginName)
	err = global.GVA_DB.Model(&model.SysPluginSdk{}).Select("sdk_version , sdk_file_name , sdk_file_url").Where("plugin_name = ?", pluginName).Scan(&sdks).Error
	return err, sdks
}

func SearchPluginSdk(bean request.ReqListPluginSdkBean) (err error, list interface{}, total int) {
	limit := bean.PageSize
	offset := bean.PageSize * (bean.Page - 1)
	//创建db
	db := global.GVA_DB.Model(&model.SysPluginSdk{})
	var pluginSdks []model.SysPluginSdk
	if bean.PluginName != "" {
		db = db.Where("plugin_name = ?", bean.PluginName)
	}
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&pluginSdks).Error
	return err, pluginSdks, total
}

func CreatePluginSdk(bean model.SysPluginSdk) (err error) {
	var pluginSdk model.SysPluginSdk
	//根据版本去判断是否存在
	isExit := !global.GVA_DB.Where("plugin_name = ? AND sdk_version = ?", bean.PluginName, bean.SdkVersion).First(&pluginSdk).RecordNotFound()
	//isExit为true表明读到了，不能新建
	if isExit {
		return errors.New("该版本插件SDK已存在")
	} else {
		err = global.GVA_DB.Create(&bean).Error
	}
	return err
}

func ModifyPluginSdk(bean model.SysPluginSdk) (err error) {
	var pluginSdk model.SysPluginSdk
	//根据版本去判断是否存在
	isExit := !global.GVA_DB.Where("plugin_name = ? AND sdk_version = ?", bean.PluginName, bean.SdkVersion).First(&pluginSdk).RecordNotFound()
	//isExit为true表明读到了，不能新建
	if isExit {
		err = global.GVA_DB.Where("plugin_name = ? AND sdk_version = ?", bean.PluginName, bean.SdkVersion).First(&pluginSdk).Updates(map[string]interface{}{
			"plugin_name":   bean.PluginName,
			"sdk_name":      bean.SdkName,
			"sdk_version":   bean.SdkVersion,
			"sdk_file_name": bean.SdkFileName,
			"sdk_file_url":  bean.SdkFileUrl,
		}).Error
	} else {
		return errors.New("该版本插件SDK不存在")
	}
	return err
}
