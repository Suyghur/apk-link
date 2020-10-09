/*
@Time : 2020/8/28
@Author : #Suyghur,
@File : sys_sdk
*/

package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"server/global"
	"server/global/response"
	"server/model"
	"server/model/bean/request"
	resp "server/model/bean/response"
	"server/service"
	"server/utils"
)

func GetFuseSdks(c *gin.Context) {
	err, fuseSdks := service.GetFuseSdks()
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("获取数据失败，%v", err), c)
	} else {
		response.OkWithData(gin.H{"fuse_sdks": fuseSdks}, c)
	}
}
func SearchFuseSdk(c *gin.Context) {
	var bean request.ReqListFuseSdkBean
	_ = c.ShouldBindQuery(&bean)
	err, list, total := service.SearchFuseSdk(bean)
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

func ModifyFuseSdk(c *gin.Context) {
	var bean request.ReqFuseSdkBean
	_ = c.ShouldBindJSON(&bean)
	verifyRules := utils.Rules{
		"SdkName":     {utils.NotEmpty()},
		"SdkVersion":  {utils.NotEmpty()},
		"SdkFileName": {utils.NotEmpty()},
		"SdkFileUrl":  {utils.NotEmpty()},
	}
	if verifyErr := utils.Verify(bean, verifyRules); verifyErr != nil {
		response.FailWithMessage(verifyErr.Error(), c)
		return
	}
	fuseSdk := &model.SysFuseSdk{
		SdkName:     bean.SdkName,
		SdkVersion:  bean.SdkVersion,
		SdkFileName: bean.SdkFileName,
		SdkFileUrl:  bean.SdkFileUrl,
	}
	err := service.ModifyFuseSdk(*fuseSdk)
	if err != nil {
		global.GVA_LOG.Error(err.Error())
		response.FailWithMessage(fmt.Sprintf("%v", err), c)
	} else {
		response.OkWithMessage("修改聚合SDK成功", c)
	}
}

func CreateFuseSdk(c *gin.Context) {
	var bean request.ReqFuseSdkBean
	_ = c.ShouldBindJSON(&bean)
	verifyRules := utils.Rules{
		"SdkName":     {utils.NotEmpty()},
		"SdkVersion":  {utils.NotEmpty()},
		"SdkFileName": {utils.NotEmpty()},
		"SdkFileUrl":  {utils.NotEmpty()},
	}
	if verifyErr := utils.Verify(bean, verifyRules); verifyErr != nil {
		response.FailWithMessage(verifyErr.Error(), c)
		return
	}
	fuseSdk := &model.SysFuseSdk{
		SdkName:     bean.SdkName,
		SdkVersion:  bean.SdkVersion,
		SdkFileName: bean.SdkFileName,
		SdkFileUrl:  bean.SdkFileUrl,
	}
	err := service.CreateFuseSdk(*fuseSdk)
	if err != nil {
		global.GVA_LOG.Error(err.Error())
		response.FailWithMessage(fmt.Sprintf("%v", err), c)
	} else {
		response.OkWithMessage("创建聚合SDK成功", c)
	}
}

func GetChannelSdks(c *gin.Context) {
	var bean request.ReqChannelSdksBean
	_ = c.ShouldBindQuery(&bean)
	verifyRules := utils.Rules{
		"ChannelName": {utils.NotEmpty()},
	}
	if verifyErr := utils.Verify(bean, verifyRules); verifyErr != nil {
		response.FailWithMessage(verifyErr.Error(), c)
		return
	}
	err, channelSdks := service.GetChannelSdks(bean.ChannelName)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("获取数据失败，%v", err), c)
	} else {
		response.OkWithData(gin.H{"channel_sdks": channelSdks}, c)
	}
}

func SearchChannelSdk(c *gin.Context) {
	var bean request.ReqListChannelSdkBean
	_ = c.ShouldBindQuery(&bean)
	err, list, total := service.SearchChannelSdk(bean)
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

func CreateChannelSdk(c *gin.Context) {
	var bean request.ReqChannelSdkBean
	_ = c.ShouldBindJSON(&bean)
	verifyRules := utils.Rules{
		"ChannelName": {utils.NotEmpty()},
		"SdkName":     {utils.NotEmpty()},
		"SdkVersion":  {utils.NotEmpty()},
		"SdkFileName": {utils.NotEmpty()},
		"SdkFileUrl":  {utils.NotEmpty()},
	}
	if verifyErr := utils.Verify(bean, verifyRules); verifyErr != nil {
		response.FailWithMessage(verifyErr.Error(), c)
		return
	}
	channelSdk := &model.SysChannelSdk{
		ChannelName: bean.ChannelName,
		SdkName:     bean.SdkName,
		SdkVersion:  bean.SdkVersion,
		SdkFileName: bean.SdkFileName,
		SdkFileUrl:  bean.SdkFileUrl,
	}
	err := service.CreateChannelSdk(*channelSdk)
	if err != nil {
		global.GVA_LOG.Error(err.Error())
		response.FailWithMessage(fmt.Sprintf("%v", err), c)
	} else {
		response.OkWithMessage("创建渠道SDK成功", c)
	}
}

func ModifyChannelSdk(c *gin.Context) {
	var bean request.ReqChannelSdkBean
	_ = c.ShouldBindJSON(&bean)
	verifyRules := utils.Rules{
		"ChannelName": {utils.NotEmpty()},
		"SdkName":     {utils.NotEmpty()},
		"SdkVersion":  {utils.NotEmpty()},
		"SdkFileName": {utils.NotEmpty()},
		"SdkFileUrl":  {utils.NotEmpty()},
	}
	if verifyErr := utils.Verify(bean, verifyRules); verifyErr != nil {
		response.FailWithMessage(verifyErr.Error(), c)
		return
	}
	channelSdk := &model.SysChannelSdk{
		ChannelName: bean.ChannelName,
		SdkName:     bean.SdkName,
		SdkVersion:  bean.SdkVersion,
		SdkFileName: bean.SdkFileName,
		SdkFileUrl:  bean.SdkFileUrl,
	}
	err := service.ModifyChannelSdk(*channelSdk)
	if err != nil {
		global.GVA_LOG.Error(err.Error())
		response.FailWithMessage(fmt.Sprintf("%v", err), c)
	} else {
		response.OkWithMessage("修改渠道SDK成功", c)
	}
}

func GetPluginSdks(c *gin.Context) {
	var bean request.ReqPluginSdksBean
	_ = c.ShouldBindQuery(&bean)
	verifyRules := utils.Rules{
		"PluginName": {utils.NotEmpty()},
	}
	if verifyErr := utils.Verify(bean, verifyRules); verifyErr != nil {
		response.FailWithMessage(verifyErr.Error(), c)
		return
	}
	err, pluginSdks := service.GetPluginSdks(bean.PluginName)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("获取数据失败，%v", err), c)
	} else {
		response.OkWithData(gin.H{"plugin_sdks": pluginSdks}, c)
	}
}

func SearchPluginSdk(c *gin.Context) {
	var bean request.ReqListPluginSdkBean
	_ = c.ShouldBindQuery(&bean)
	err, list, total := service.SearchPluginSdk(bean)
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

func ModifyPluginSdk(c *gin.Context) {
	var bean request.ReqPluginSdkBean
	_ = c.ShouldBindJSON(&bean)
	verifyRules := utils.Rules{
		"PluginName":  {utils.NotEmpty()},
		"SdkName":     {utils.NotEmpty()},
		"SdkVersion":  {utils.NotEmpty()},
		"SdkFileName": {utils.NotEmpty()},
		"SdkFileUrl":  {utils.NotEmpty()},
	}
	if verifyErr := utils.Verify(bean, verifyRules); verifyErr != nil {
		response.FailWithMessage(verifyErr.Error(), c)
		return
	}
	pluginSdk := &model.SysPluginSdk{
		PluginName:  bean.PluginName,
		SdkName:     bean.SdkName,
		SdkVersion:  bean.SdkVersion,
		SdkFileName: bean.SdkFileName,
		SdkFileUrl:  bean.SdkFileUrl,
	}
	err := service.ModifyPluginSdk(*pluginSdk)
	if err != nil {
		global.GVA_LOG.Error(err.Error())
		response.FailWithMessage(fmt.Sprintf("%v", err), c)
	} else {
		response.OkWithMessage("修改插件SDK成功", c)
	}
}

func CreatePluginSdk(c *gin.Context) {
	var bean request.ReqPluginSdkBean
	_ = c.ShouldBindJSON(&bean)
	verifyRules := utils.Rules{
		"PluginName":  {utils.NotEmpty()},
		"SdkName":     {utils.NotEmpty()},
		"SdkVersion":  {utils.NotEmpty()},
		"SdkFileName": {utils.NotEmpty()},
		"SdkFileUrl":  {utils.NotEmpty()},
	}
	if verifyErr := utils.Verify(bean, verifyRules); verifyErr != nil {
		response.FailWithMessage(verifyErr.Error(), c)
		return
	}
	pluginSdk := &model.SysPluginSdk{
		PluginName:  bean.PluginName,
		SdkName:     bean.SdkName,
		SdkVersion:  bean.SdkVersion,
		SdkFileName: bean.SdkFileName,
		SdkFileUrl:  bean.SdkFileUrl,
	}
	err := service.CreatePluginSdk(*pluginSdk)
	if err != nil {
		global.GVA_LOG.Error(err.Error())
		response.FailWithMessage(fmt.Sprintf("%v", err), c)
	} else {
		response.OkWithMessage("创建插件SDK成功", c)
	}
}
