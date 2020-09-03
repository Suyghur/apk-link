/*
@Time : 8/31/2020
@Author : #Suyghur,
@File : sjys_task
*/

package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"linking-api/global"
	"linking-api/global/response"
	"linking-api/model"
	"linking-api/model/bean/request"
	resp "linking-api/model/bean/response"
	"linking-api/service"
	"linking-api/utils"
)

func ListTask(c *gin.Context) {
	var bean request.ReqListTaskBean
	_ = c.ShouldBindJSON(&bean)
	err, list, total := service.ListTask(bean)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("获取数据失败，%v", err), c)
	} else {
		response.OkWithData(resp.PageResult{
			List:     list,
			Total:    total,
			Page:     bean.Page,
			PageSize: bean.PageSize,
		}, c)
	}
}

func CreateTask(c *gin.Context) {
	var bean request.ReqTaskBean
	_ = c.ShouldBindJSON(&bean)
	verifyRules := utils.Rules{
		"GameGroup":         {utils.NotEmpty()},
		"Aids":              {utils.NotEmpty()},
		"GamePackageName":   {utils.NotEmpty()},
		"GameName":          {utils.NotEmpty()},
		"GameVersionCode":   {utils.NotEmpty()},
		"GameVersionName":   {utils.NotEmpty()},
		"GameFileName":      {utils.NotEmpty()},
		"FuseSdkVersion":    {utils.NotEmpty()},
		"ChannelName":       {utils.NotEmpty()},
		"ChannelSdkVersion": {utils.NotEmpty()},
		"KeystoreName":      {utils.NotEmpty()},
	}
	if verifyErr := utils.Verify(bean, verifyRules); verifyErr != nil {
		response.FailWithMessage(verifyErr.Error(), c)
		return
	}
	task := &model.SysTask{
		TaskId:                bean.TaskId,
		IsFuseSdk:             bean.IsFuseSdk,
		IsWhiteBag:            bean.IsWhiteBag,
		IsPluginSdk:           bean.IsFuseSdk,
		GameGroup:             bean.GameGroup,
		Gid:                   bean.Gid,
		Aids:                  bean.Aids,
		ChannelParams:         bean.ChannelParams,
		PluginParams:          bean.PluginParams,
		GamePackageName:       bean.GamePackageName,
		GameName:              bean.GameName,
		GameVersionCode:       bean.GameVersionCode,
		GameVersionName:       bean.GameVersionName,
		GameOrientation:       bean.GameOrientation,
		GameIconUrl:           bean.GameIconUrl,
		GameLogoUrl:           bean.GameLogoUrl,
		GameSplashUrl:         bean.GameSplashUrl,
		GameLoginBgUrl:        bean.GameLoginBgUrl,
		GameFileName:          bean.GameFileName,
		GameFileUrl:           bean.GameFileUrl,
		FuseSdkFileName:       bean.FuseSdkFileName,
		FuseSdkVersion:        bean.FuseSdkVersion,
		FuseSdkFileUrl:        bean.FuseSdkFileUrl,
		ChannelName:           bean.ChannelName,
		ChannelSdkFileName:    bean.ChannelSdkFileName,
		ChannelSdkVersion:     bean.ChannelSdkVersion,
		ChannelSdkFileUrl:     bean.ChannelSdkFileUrl,
		PluginName:            bean.PluginName,
		PluginSdkFileName:     bean.PluginSdkFileName,
		PluginSdkVersion:      bean.PluginSdkVersion,
		PluginSdkFileUrl:      bean.PluginSdkFileUrl,
		ScriptFileName:        bean.ScriptFileName,
		ScriptFileUrl:         bean.ScriptFileUrl,
		KeystoreName:          bean.KeystoreName,
		KeystorePassword:      bean.KeystorePassword,
		KeystoreAlias:         bean.KeystoreAlias,
		KeystoreAliasPassword: bean.KeystoreAliasPassword,
		Ext:                   bean.Ext,
	}
	err := service.CreateTask(*task)
	if err != nil {
		global.GVA_LOG.Error(err.Error())
		response.FailWithMessage(fmt.Sprintf("%v", err), c)
	} else {
		response.OkWithMessage("创建打包任务成功", c)
	}
}

func ModifyTask(c *gin.Context) {
	var bean request.ReqTaskBean
	_ = c.ShouldBindJSON(&bean)
	verifyRules := utils.Rules{
		"GameGroup":         {utils.NotEmpty()},
		"Aids":              {utils.NotEmpty()},
		"GamePackageName":   {utils.NotEmpty()},
		"GameName":          {utils.NotEmpty()},
		"GameVersionCode":   {utils.NotEmpty()},
		"GameVersionName":   {utils.NotEmpty()},
		"GameFileName":      {utils.NotEmpty()},
		"FuseSdkVersion":    {utils.NotEmpty()},
		"ChannelName":       {utils.NotEmpty()},
		"ChannelSdkVersion": {utils.NotEmpty()},
		"KeystoreName":      {utils.NotEmpty()},
	}
	if verifyErr := utils.Verify(bean, verifyRules); verifyErr != nil {
		response.FailWithMessage(verifyErr.Error(), c)
		return
	}
	task := &model.SysTask{
		TaskId:                bean.TaskId,
		IsFuseSdk:             bean.IsFuseSdk,
		IsWhiteBag:            bean.IsWhiteBag,
		IsPluginSdk:           bean.IsFuseSdk,
		GameGroup:             bean.GameGroup,
		Gid:                   bean.Gid,
		Aids:                  bean.Aids,
		ChannelParams:         bean.ChannelParams,
		PluginParams:          bean.PluginParams,
		GamePackageName:       bean.GamePackageName,
		GameName:              bean.GameName,
		GameVersionCode:       bean.GameVersionCode,
		GameVersionName:       bean.GameVersionName,
		GameOrientation:       bean.GameOrientation,
		GameIconUrl:           bean.GameIconUrl,
		GameLogoUrl:           bean.GameLogoUrl,
		GameSplashUrl:         bean.GameSplashUrl,
		GameLoginBgUrl:        bean.GameLoginBgUrl,
		GameFileName:          bean.GameFileName,
		GameFileUrl:           bean.GameFileUrl,
		FuseSdkFileName:       bean.FuseSdkFileName,
		FuseSdkVersion:        bean.FuseSdkVersion,
		FuseSdkFileUrl:        bean.FuseSdkFileUrl,
		ChannelName:           bean.ChannelName,
		ChannelSdkFileName:    bean.ChannelSdkFileName,
		ChannelSdkVersion:     bean.ChannelSdkVersion,
		ChannelSdkFileUrl:     bean.ChannelSdkFileUrl,
		PluginName:            bean.PluginName,
		PluginSdkFileName:     bean.PluginSdkFileName,
		PluginSdkVersion:      bean.PluginSdkVersion,
		PluginSdkFileUrl:      bean.PluginSdkFileUrl,
		ScriptFileName:        bean.ScriptFileName,
		ScriptFileUrl:         bean.ScriptFileUrl,
		KeystoreName:          bean.KeystoreName,
		KeystorePassword:      bean.KeystorePassword,
		KeystoreAlias:         bean.KeystoreAlias,
		KeystoreAliasPassword: bean.KeystoreAliasPassword,
		Ext:                   bean.Ext,
	}
	err := service.ModifyTask(*task)
	if err != nil {
		global.GVA_LOG.Error(err.Error())
		response.FailWithMessage(fmt.Sprintf("%v", err), c)
	} else {
		response.OkWithMessage("修改打包任务成功", c)
	}
}

func ModifyTaskStatus(c *gin.Context) {
	var bean request.ReqTaskStatusBean
	_ = c.ShouldBindJSON(&bean)
	verifyRules := utils.Rules{
		"task_id": {utils.NotEmpty()},
	}
	if verifyErr := utils.Verify(bean, verifyRules); verifyErr != nil {
		response.FailWithMessage(verifyErr.Error(), c)
		return
	}
	err := service.ModifyTaskStatus(bean.TaskId, bean.Status)
	if err != nil {
		global.GVA_LOG.Error(err.Error())
		response.FailWithMessage(fmt.Sprintf("%v", err), c)
	} else {
		response.OkWithMessage("修改打包任务状态成功", c)
	}
}
