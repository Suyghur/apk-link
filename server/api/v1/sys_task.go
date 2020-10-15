/*
@Time : 8/31/2020
@Author : #Suyghur,
@File : sys_task
*/

package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	//"go.uber.org/zap"
	"server/global"
	"server/global/response"
	"server/model"
	"server/model/bean/request"
	resp "server/model/bean/response"
	"server/service"
	"server/utils"
)

//@Tags task
//@Title SearchTasks
//@Summary 搜索任务
//@Produce application/json
//@Param data body response.PageResult true "搜索任务"
//@Success 200 {string} string "{"code":0,"data":{},"msg":"success"}"
//@Router /v1/task/searchTasks [get]
func SearchTasks(c *gin.Context) {
	var bean request.ReqListTaskBean
	_ = c.ShouldBindQuery(&bean)
	err, list, total := service.SearchTasks(&bean)
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

func DeleteTask(c *gin.Context) {
	var bean request.ReqTaskInfoBean
	_ = c.ShouldBindJSON(&bean)
	verifyRules := utils.Rules{
		"TaskId": {utils.NotEmpty()},
	}
	if verifyErr := utils.Verify(bean, verifyRules); verifyErr != nil {
		response.FailWithMessage(verifyErr.Error(), c)
		return
	}
	err := service.DeleteTask(bean.TaskId)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("删除任务失败，%v", err), c)
	} else {
		response.OkWithMessage("删除任务成功", c)
	}
}

func GetTaskInfo(c *gin.Context) {
	var bean request.ReqTaskInfoBean
	_ = c.ShouldBindQuery(&bean)
	global.GvaLog.Info("bean : ", zap.Any("bean", bean))
	verifyRules := utils.Rules{
		"TaskId": {utils.NotEmpty()},
	}
	if verifyErr := utils.Verify(bean, verifyRules); verifyErr != nil {
		response.FailWithMessage(verifyErr.Error(), c)
		return
	}
	err, taskInfo := service.GetTaskInfo(bean.TaskId)
	if err != nil {
		global.GvaLog.Error(err.Error())
		response.FailWithMessage(fmt.Sprintf("%v", err), c)
	} else {
		response.OkWithData(gin.H{"task_info": taskInfo}, c)
	}
}

func CreateTask(c *gin.Context) {
	var bean request.ReqCreateTaskBean
	_ = c.ShouldBindJSON(&bean)
	verifyRules := utils.Rules{
		"GameSite":          {utils.NotEmpty()},
		"GamePackageName":   {utils.NotEmpty()},
		"GameName":          {utils.NotEmpty()},
		"GameVersionCode":   {utils.NotEmpty()},
		"GameVersionName":   {utils.NotEmpty()},
		"GameFileName":      {utils.NotEmpty()},
		"ChannelName":       {utils.NotEmpty()},
		"ChannelSdkVersion": {utils.NotEmpty()},
		"KeystoreName":      {utils.NotEmpty()},
	}
	if verifyErr := utils.Verify(bean, verifyRules); verifyErr != nil {
		response.FailWithMessage(verifyErr.Error(), c)
		return
	}
	task := &model.SysTask{
		IsWhiteBag:            bean.IsWhiteBag,
		GameSite:              bean.GameSite,
		Gid:                   bean.Gid,
		Cid:                   bean.Cid,
		FormId:                bean.FormId,
		Aids:                  bean.Aids,
		ChannelParams:         bean.ChannelParams,
		PluginParams:          bean.PluginParams,
		GamePackageName:       bean.GamePackageName,
		GameName:              bean.GameName,
		GameVersionCode:       bean.GameVersionCode,
		GameVersionName:       bean.GameVersionName,
		GameIconUrl:           bean.GameIconUrl,
		GameLogoUrl:           bean.GameLogoUrl,
		GameSplashUrl:         bean.GameSplashUrl,
		GameLoginBgUrl:        bean.GameLoginBgUrl,
		GameFileName:          bean.GameFileName,
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
		KeystoreName:          bean.KeystoreName,
		KeystorePassword:      bean.KeystorePassword,
		KeystoreAlias:         bean.KeystoreAlias,
		KeystoreAliasPassword: bean.KeystoreAliasPassword,
		KeystoreFileUrl:       bean.KeystoreFileUrl,
		Ext:                   bean.Ext,
	}
	err := service.CreateTask(*task)
	if err != nil {
		global.GvaLog.Error(err.Error())
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
		"GamePackageName":   {utils.NotEmpty()},
		"GameName":          {utils.NotEmpty()},
		"GameVersionCode":   {utils.NotEmpty()},
		"GameVersionName":   {utils.NotEmpty()},
		"GameFileName":      {utils.NotEmpty()},
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
		IsWhiteBag:            bean.IsWhiteBag,
		IsPluginSdk:           bean.IsPluginSdk,
		GameSite:              bean.GameSite,
		Gid:                   bean.Gid,
		Aids:                  bean.Aids,
		ChannelParams:         bean.ChannelParams,
		PluginParams:          bean.PluginParams,
		GamePackageName:       bean.GamePackageName,
		GameName:              bean.GameName,
		GameVersionCode:       bean.GameVersionCode,
		GameVersionName:       bean.GameVersionName,
		GameOrientation:       bean.GameOrientation,
		StatusCode:            1000,
		StatusMsg:             "未执行",
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
		global.GvaLog.Error(err.Error())
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
	statusMsg := "未执行"
	switch bean.StatusCode {
	case 1000:
		statusMsg = "未执行"
	case 1001:
		statusMsg = "执行中"
	case 1002:
		statusMsg = "成功"
	case 1003:
		statusMsg = "失败"
	}
	err := service.ModifyTaskStatus(bean.TaskId, bean.StatusCode, statusMsg)
	if err != nil {
		global.GvaLog.Error(err.Error())
		response.FailWithMessage(fmt.Sprintf("%v", err), c)
	} else {
		response.OkWithMessage("修改打包任务状态成功", c)
	}
}
