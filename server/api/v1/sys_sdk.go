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
		response.OkWithData(fuseSdks, c)
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
	var bean model.SysFuseSdk
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
	err := service.ModifyFuseSdk(bean)
	if err != nil {
		global.GvaLog.Error(err.Error())
		response.FailWithMessage(fmt.Sprintf("%v", err), c)
	} else {
		response.OkWithMessage("修改聚合SDK成功", c)
	}
}

func CreateFuseSdk(c *gin.Context) {
	var bean model.SysFuseSdk
	_ = c.ShouldBindJSON(&bean)
	verifyRules := utils.Rules{
		"SdkVersion":  {utils.NotEmpty()},
		"SdkFileName": {utils.NotEmpty()},
		"SdkFileUrl":  {utils.NotEmpty()},
	}
	if verifyErr := utils.Verify(bean, verifyRules); verifyErr != nil {
		response.FailWithMessage(verifyErr.Error(), c)
		return
	}
	err := service.CreateFuseSdk(bean)
	if err != nil {
		global.GvaLog.Error(err.Error())
		response.FailWithMessage(fmt.Sprintf("%v", err), c)
	} else {
		response.OkWithMessage("创建聚合SDK成功", c)
	}
}

func GetChannelSdks(c *gin.Context) {
	var bean model.SysChannelSdk
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
		response.OkWithData(channelSdks, c)
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
	var bean model.SysChannelSdk
	_ = c.ShouldBindJSON(&bean)
	verifyRules := utils.Rules{
		"Cid":          {utils.NotEmpty()},
		"ChannelName":  {utils.NotEmpty()},
		"ChannelAlias": {utils.NotEmpty()},
		"SdkVersion":   {utils.NotEmpty()},
		"SdkFileName":  {utils.NotEmpty()},
		"SdkFileUrl":   {utils.NotEmpty()},
	}
	if verifyErr := utils.Verify(bean, verifyRules); verifyErr != nil {
		response.FailWithMessage(verifyErr.Error(), c)
		return
	}
	err := service.CreateChannelSdk(bean)
	if err != nil {
		global.GvaLog.Error(err.Error())
		response.FailWithMessage(fmt.Sprintf("%v", err), c)
	} else {
		response.OkWithMessage("创建渠道SDK成功", c)
	}
}

func ModifyChannelSdk(c *gin.Context) {
	var bean model.SysChannelSdk
	_ = c.ShouldBindJSON(&bean)
	verifyRules := utils.Rules{
		"Cid":         {utils.NotEmpty()},
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
	err := service.ModifyChannelSdk(bean)
	if err != nil {
		global.GvaLog.Error(err.Error())
		response.FailWithMessage(fmt.Sprintf("%v", err), c)
	} else {
		response.OkWithMessage("修改渠道SDK成功", c)
	}
}

func GetPluginSdks(c *gin.Context) {
	var bean model.SysPluginSdk
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
		response.OkWithData(pluginSdks, c)
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
	var bean model.SysPluginSdk
	_ = c.ShouldBindJSON(&bean)
	verifyRules := utils.Rules{
		"FormId":      {utils.NotEmpty()},
		"PluginName":  {utils.NotEmpty()},
		"PluginAlias": {utils.NotEmpty()},
		"SdkVersion":  {utils.NotEmpty()},
		"SdkFileName": {utils.NotEmpty()},
		"SdkFileUrl":  {utils.NotEmpty()},
	}
	if verifyErr := utils.Verify(bean, verifyRules); verifyErr != nil {
		response.FailWithMessage(verifyErr.Error(), c)
		return
	}
	err := service.ModifyPluginSdk(bean)
	if err != nil {
		global.GvaLog.Error(err.Error())
		response.FailWithMessage(fmt.Sprintf("%v", err), c)
	} else {
		response.OkWithMessage("修改插件SDK成功", c)
	}
}

func CreatePluginSdk(c *gin.Context) {
	var bean model.SysPluginSdk
	_ = c.ShouldBindJSON(&bean)
	verifyRules := utils.Rules{
		"FormId":      {utils.NotEmpty()},
		"PluginName":  {utils.NotEmpty()},
		"PluginAlias": {utils.NotEmpty()},
		"SdkVersion":  {utils.NotEmpty()},
		"SdkFileName": {utils.NotEmpty()},
		"SdkFileUrl":  {utils.NotEmpty()},
	}
	if verifyErr := utils.Verify(bean, verifyRules); verifyErr != nil {
		response.FailWithMessage(verifyErr.Error(), c)
		return
	}
	err := service.CreatePluginSdk(bean)
	if err != nil {
		global.GvaLog.Error(err.Error())
		response.FailWithMessage(fmt.Sprintf("%v", err), c)
	} else {
		response.OkWithMessage("创建插件SDK成功", c)
	}
}
